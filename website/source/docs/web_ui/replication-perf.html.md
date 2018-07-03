---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-replication-perf"
description: |-
  This
---


## Performance Replication
#### Web UI

1. Select **Replication** and check the **Performance** radio button.
  ![Performance Replication - primary](/assets/images/vault-mount-filter-5.png)

1. Click **Enable replication**.

1. Select the **Secondaries** tab, and then click **Add**.
  ![Performance Replication - primary](/assets/images/vault-mount-filter-6.png)

1. Populate the **Secondary ID** field, and then select **Configure performance
mount filtering** to set your mount filter options.  You can filter by
whitelisting or blacklisting. For this example, select **Blacklist**.

1. Check **EU_GDPR_data** to prevent it from being replicated to the secondary
cluster.
  ![Performance Replication - primary](/assets/images/vault-mount-filter-7.png)

1. Click **Generate token**.
  ![Performance Replication - primary](/assets/images/vault-mount-filter-8.png)

1. Click **Copy** to copy the token.

1. Now, launch the Vault UI for the secondary cluster (e.g. https://us-central.compute.com:8201/ui), and then click **Replication**.

1. Check the **Performance** radio button, and then select **secondary** under the **Cluster mode**. Paste the token you copied from the primary.
  ![Performance Replication - secondary](/assets/images/vault-mount-filter-9.png)

1. Click **Enable replication**.

<br>

~> **NOTE:** At this point, the secondary cluster must be unsealed using the
**primary cluster's unseal key**. If the secondary is in an HA cluster, ensure
that each standby is sealed and unsealed with the primary’s unseal keys. The
secondary cluster mirrors the configuration of its primary cluster's backends
such as auth methods, secret engines, audit devices, etc. It uses the primary as
the _source of truth_ and passes token requests to the primary.


Restart the secondary vault server (e.g. `https://us-central.compute.com:8201`)
and unseal it with the primary cluster's unseal key.

```plaintext
$ vault operator unseal
Unseal Key (will be hidden): <primary_cluster_unseal_key>
```

The initial root token on the secondary no longer works. Use the auth methods
configured on the primary cluster to log into the secondary.

**Example:**

Enable and configure the userpass auth method on the **primary** cluster and
create a new username and password.

```shell
# Enable the userpass auth method on the primary
$ vault auth enable userpass

# Create a user with admin policy
$ vault write auth/userpass/users/james password="passw0rd" policy="admin"
```

-> Alternatively, you can [generate a new root token](/guides/operations/generate-root.html)
using the primary cluster's unseal key. However, it is recommended that root
tokens are only used for just enough initial setup or in emergencies.


Log into the **secondary** cluster using the enabled auth method.

```plaintext
$ vault login -method=userpass username=james password="passw0rd"
```


### <a name="step3"></a>Step 3: Verify the replication mount filter

Once the replication completes, verify that the secrets stored in the
`EU_GDPR_data` never get replicated to the US cluster.


On the **EU** cluster, select **EU_GDPR_data** > **Create secret**:

![Secrets](/assets/images/vault-mount-filter-12.png)

Enter the values and click **Save**.  Repeat the step to write some secrets at
the **US_NON_GDPR_data** path as well.


On the **US** cluster, select **US_NON_GDPR_data**. You should be able to see
the `apikey` under `US_NON_GDPR_data/secret`.

![Secrets](/assets/images/vault-mount-filter-13.png)

The **EU_GDPR_data** data is not replicated, so you won't be able to see the
secrets.


### <a name="step4"></a>Step 4: Enable a local secret engine

When replication is enabled, you can mark the secrets engine local only.  Local
secret engines are not replicated or removed by replication.

Login to the **secondary** cluster and enable key/value secret engine at
`US_ONLY_data` to store secrets only valid for the US region.


Be sure to select the check box for **Local** to keep it mounted locally within
the cluster.

![Local Secret](/assets/images/vault-mount-filter-10.png)

<br>
