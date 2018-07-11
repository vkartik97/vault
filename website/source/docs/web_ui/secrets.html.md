---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-secrets"
description: |-
  This
---


## Secrets

The **Secrets** tab is the default destination when signed into the UI.

![Secrets](/assets/images/vault-ui-guide/vault-ui-secrets0.png)

A [Cubbyhole Secrets Engine](/docs/secrets/cubbyhole/index.html) and a
[Key/Value Secrets Engine](/docs/secrets/kv/index.html) are enabled by default
and ready to use.

Using the UI, you can managed the enabled [Secrets
Engines](/docs/secrets/index.html) and enable a new Secrets Engine.


### Enabling a new secrets engine

1. Select **Enable new engine** to begin enabling the demonstration Secrets Engine.
    ![](/assets/images/vault-ui-guide/vault-ui-secrets-enable0-1.png)

1. The **Secrets engine type** drop-down list shows available secret engines to choose from.
    ![](/assets/images/vault-ui-guide/vault-ui-secrets-enable0.png)

    Vault provides support for a wide array of [Secrets
    Engines](/docs/secrets/index.html) which all have their own documentation.

    > Those secrets engines marked as "deprecated" (Cassandra, MySQL,
    PostgreSQL, etc.) are now supported by the [Databases Secrets
    Engine](/docs/secrets/databases/index.html). The database secret engine is
    now the preferred secret engine to interact with those databases.

1. Select **KV** for the **Secrets engine type**.

1. Enter `secret-demo` in the **Path** field, and enter some description.

1. Select **Version 2** from the **Version** dropdown.
    ![](/assets/images/vault-ui-guide/vault-ui-secrets-enable1.png)
    There are some optional parameters to set on the secret engine.
    ![](/assets/images/vault-ui-guide/vault-ui-secrets-enable2.png)
    - **Local**: This option is for environments using Vaults [Enterprise Replication](/docs/enterprise/replication/index.html) and selecting this option limits the mountpoint to use by the local Vault cluster only, and its associated data will not be replicated to other Vault clusters which are participating in replication.
    - **Seal Wrap**: This option enables [Seal Wrap](/docs/enterprise/sealwrap/index.html) on the mountpoint. It can only be selected at mount time.
    - **More options**
      - **Default Lease TTL**: This is the *role* based default lease TTL. It can be set a value lower than the system default TTL, but not higher. See [Token Time-To-Live, Periodic Tokens, and Explicit Max TTLs](/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) for more details.
      - **Maximum Lease TTL**: This is the *role* based maximum lease TTL. It can be set a value lower than the system default TTL, but not higher. See [Token Time-To-Live, Periodic Tokens, and Explicit Max TTLs](/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) for more details.

1. Select **Enable Engine**.
    ![](/assets/images/vault-ui-guide/vault-ui-secrets-enable3.png)

    ~> An informational message indicates that this version of Vault does not yet contain full KV version 2 support, so it only operates on the latest version of a secret when using KV version 2.



### View secrets engine configuration

Select a secret engine, and click **Configuration** to view its details.

![](/assets/images/vault-ui-guide/vault-ui-secret-demo0.png)



### Create a secret

To create a secret,  make sure that you are on the **Secrets** tab and then select **Create secret**; let's break down some components of the secret:

1. In the **Secrets** tab, select `secret-demo` and then click **Create secret**.
    ![](/assets/images/vault-ui-guide/vault-ui-create-secret1.png)

1. Enter the desired path for the secret in the **PATH FOR THIS SECRET** field.

1. Enter the key-value pair secrets that you wish to store.
    ![](/assets/images/vault-ui-guide/vault-ui-create-secret0.png)

1. Click **Save**.



## Next Step
