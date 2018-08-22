package alicloud

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	aliCloudAuth "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/providers"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault-plugin-auth-alicloud/tools"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/command/agent/auth"
)

const defaultCredCheckFreqSeconds = 60

func NewAliCloudAuthMethod(conf *auth.AuthConfig) (auth.AuthMethod, error) {
	if conf == nil {
		return nil, errors.New("empty config")
	}
	if conf.Config == nil {
		return nil, errors.New("empty config data")
	}

	a := &alicloudMethod{
		logger:     conf.Logger,
		mountPath:  conf.MountPath,
		credsFound: make(chan struct{}),
		stopCh:     make(chan struct{}),
	}

	// Build the required information we'll need to create a client.
	if roleRaw, ok := conf.Config["role"]; !ok {
		return nil, errors.New("'role' is required but is not provided")
	} else {
		if a.role, ok = roleRaw.(string); !ok {
			return nil, errors.New("could not convert 'role' config value to string")
		}
	}
	if regionRaw, ok := conf.Config["region"]; !ok {
		return nil, errors.New("'region' is required but is not provided")
	} else {
		if a.region, ok = regionRaw.(string); !ok {
			return nil, errors.New("could not convert 'region' config value to string")
		}
	}

	// Check for an optional custom frequency at which we should poll for creds.
	if checkFreqRaw, ok := conf.Config["cred_check_freq_seconds"]; ok {
		if credFreq, ok := checkFreqRaw.(int); !ok {
			a.credCheckFreqSec = defaultCredCheckFreqSeconds
		} else {
			a.credCheckFreqSec = credFreq
		}
	}

	// Build the optional, configuration-based piece of the credential chain.
	credConfig := &providers.Configuration{}

	if accessKeyRaw, ok := conf.Config["access_key"]; ok {
		if credConfig.AccessKeyID, ok = accessKeyRaw.(string); !ok {
			return nil, errors.New("could not convert 'access_key' config value to string")
		}
	}

	if accessSecretRaw, ok := conf.Config["access_secret"]; ok {
		if credConfig.AccessKeySecret, ok = accessSecretRaw.(string); !ok {
			return nil, errors.New("could not convert 'access_secret' config value to string")
		}
	}

	if accessTokenRaw, ok := conf.Config["access_token"]; ok {
		if credConfig.AccessKeyStsToken, ok = accessTokenRaw.(string); !ok {
			return nil, errors.New("could not convert 'access_token' config value to string")
		}
	}

	if roleArnRaw, ok := conf.Config["role_arn"]; ok {
		if credConfig.RoleArn, ok = roleArnRaw.(string); !ok {
			return nil, errors.New("could not convert 'role_arn' config value to string")
		}
	}

	if roleSessionNameRaw, ok := conf.Config["role_session_name"]; ok {
		if credConfig.RoleSessionName, ok = roleSessionNameRaw.(string); !ok {
			return nil, errors.New("could not convert 'role_session_name' config value to string")
		}
	}

	if roleSessionExpirationRaw, ok := conf.Config["role_session_expiration"]; ok {
		if roleSessionExpiration, ok := roleSessionExpirationRaw.(int); !ok {
			return nil, errors.New("could not convert 'role_session_expiration' config value to int")
		} else {
			credConfig.RoleSessionExpiration = &roleSessionExpiration
		}
	}

	if privateKeyRaw, ok := conf.Config["private_key"]; ok {
		if credConfig.PrivateKey, ok = privateKeyRaw.(string); !ok {
			return nil, errors.New("could not convert 'private_key' config value to string")
		}
	}

	if publicKeyIdRaw, ok := conf.Config["public_key_id"]; ok {
		if credConfig.PublicKeyID, ok = publicKeyIdRaw.(string); !ok {
			return nil, errors.New("could not convert 'public_key_id' config value to string")
		}
	}

	if sessionExpirationRaw, ok := conf.Config["session_expiration"]; ok {
		if sessionExpiration, ok := sessionExpirationRaw.(int); !ok {
			return nil, errors.New("could not convert 'session_expiration' config value to int")
		} else {
			credConfig.SessionExpiration = &sessionExpiration
		}
	}

	if roleNameRaw, ok := conf.Config["role_name"]; ok {
		if credConfig.RoleName, ok = roleNameRaw.(string); !ok {
			return nil, errors.New("could not convert 'role_name' config value to string")
		}
	}

	credentialChain := []providers.Provider{
		providers.NewEnvCredentialProvider(),
		providers.NewConfigurationCredentialProvider(credConfig),
		providers.NewInstanceMetadataProvider(),
	}
	a.credentialChain = providers.NewChainProvider(credentialChain)

	go a.pollForNewCreds()

	return a, nil
}

type alicloudMethod struct {
	logger           hclog.Logger
	mountPath        string
	credCheckFreqSec int

	role            string
	credentialChain providers.Provider
	region          string

	credLock  sync.Mutex
	lastCreds aliCloudAuth.Credential

	credsFound chan struct{}
	stopCh     chan struct{}
}

func (m *alicloudMethod) Authenticate(context.Context, *api.Client) (string, map[string]interface{}, error) {
	m.logger.Trace("beginning authentication")

	m.credLock.Lock()
	if m.lastCreds == nil {
		// This is our first time authenticating.
		creds, err := m.credentialChain.Retrieve()
		if err != nil {
			return "", nil, err
		}
		m.lastCreds = creds
	}
	data, err := tools.GenerateLoginData(m.role, m.lastCreds, m.region)
	if err != nil {
		return "", nil, err
	}
	m.credLock.Unlock()
	return fmt.Sprintf("%s/login", m.mountPath), data, nil
}

func (m *alicloudMethod) NewCreds() chan struct{} {
	return m.credsFound
}

func (m *alicloudMethod) CredSuccess() {}

func (m *alicloudMethod) Shutdown() {
	close(m.stopCh)
}

func (m *alicloudMethod) pollForNewCreds() {
	timer := time.NewTimer(time.Duration(m.credCheckFreqSec) * time.Second)
	for {
		select {

		case <-m.stopCh:
			// Shutdown has been called.
			return

		case <-timer.C:
			m.credLock.Lock()
			currentCreds, err := m.credentialChain.Retrieve()
			if err != nil {
				m.logger.Warn("unable to retrieve current creds, retaining previous ones", err)
				continue
			}
			if currentCreds == m.lastCreds {
				continue
			}
			m.lastCreds = currentCreds
			m.credLock.Unlock()

			// Let the outer context know it should run the Authenticate method again.
			m.credsFound <- struct{}{}
		}
	}
}
