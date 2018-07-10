---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-config"
description: |-
  Enabling the Vault Web UI through server configuration file.
---

## Activating the Web UI

An update to the Vault server configuration is required before you can use the
Vault UI. This is documented in the [Vault
UI](/docs/configuration/ui/index.html) configuration
documentation. Configuration essentially amounts to placing one option in the
global scope of Vault's configuration file:

```hcl
ui = true

listener "tcp" {
  address = "127.0.0.1:8200"
}
```

This configuration change requires restarting the Vault process.


### Starting the server

Start the server by specifying the configuration file with `-config` flag:

**Example:**

```plaintext
$ vault server -config=config.hcl

==> Vault server configuration:

                     Cgo: disabled
              Listener 1: tcp (addr: "127.0.0.1:8200", cluster address: "127.0.0.1:8201", tls: "disabled")
               Log Level: info
                   Mlock: supported: false, enabled: false
                 Storage: file
                 Version: Vault v0.10.2
             Version Sha: 3ee0802ed08cb7f4046c2151ec4671a076b76166

==> Vault server started! Log data will stream in below:
```

### Launch the Web UI

The UI runs on the same port as the Vault listener. To access the UI, use the
following URL: `<Vault_address>/ui`

**Example:**

```plaintext
http://127.0.0.1:8200/ui
```

Refer to the [Accessing the Vault UI](/docs/configuration/ui/index.html#accessing-the-vault-ui)
documentation for more details.

## Next Step

When a Vault server was started for the first time, it needs to be
[initialized](/docs/web_ui/init.html).
