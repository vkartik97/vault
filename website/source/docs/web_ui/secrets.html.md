---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-secrets"
description: |-
  This
---


## Secrets

![](/images/vault-ui-guide/vault-ui-secrets0.png)

The **Secrets** tab is the default destination when signed into the UI.

From here, you'll be able to work with your enabled [Secrets
Engines](/docs/secrets/index.html) and the secrets they contain or enable a new
Secrets Engine. You'll note that by default and as shown in the screenshot, two
Secrets Engines are enabled by default — a [KV Secrets
Engine](/docs/secrets/kv/index.html) and a [Cubbyhole Secrets
Engine](/docs/secrets/cubbyhole/index.html).

For the purposes of secret interaction from scratch and so that we do not
disturb the existing/default KV Secrets Engine mounted at `secret/`, we will
enable a new [KV version 2 Secrets Engine](/docs/secrets/kv/kv-v2.html) mounted
at `secret-demo/` and focus our interactions on *secret-demo* throughout this
section of the guide.

Select **Enable new engine** to begin enabling the demonstration Secrets Engine.

### Enable a Secrets Engine

![](/images/vault-ui-guide/vault-ui-secrets-enable0.png)

As you will notice in the screenshot Navigation from the **Secrets** tab via the
**Enable new engine** link actually brings you to the **Settings** section where
new Secrets Engines are enabled. Note also, that Vault provides support for a
wide array of [Secrets Engines](/docs/secrets/index.html) which all have their
own documentation.

You might be curious as to why some of the Secrets Engines such as Cassandra,
MySQL, and PostgreSQL are marked as *(deprecated)*. This is due to the presence
of a newer [Databases Secrets Engine](/docs/secrets/databases/index.html). It
provides support for all of the deprecated legacy Secrets Engines, and is now
the preferred way to enable a database specific Secrets Engine.

![](/images/vault-ui-guide/vault-ui-secrets-enable1.png)

From the **Enable a secrets engine** dialog let's enable our *secret-demo* KV
version 2 Secrets Engine!

1. Select *KV* for the **Secrets engine type**
2. Set **Path** to `secret-demo`
3. Add an optional description
4. Select *Version 2* from the **Version** dropdown

This is the minimum which is required, but we'll discuss the remaining options as well.

![](/images/vault-ui-guide/vault-ui-secrets-enable2.png)

- **Local**: This option is for environments using Vaults [Enterprise Replication](/docs/enterprise/replication/index.html) and selecting this option limits the mountpoint to use by the local Vault cluster only, and its associated data will not be replicated to other Vault clusters which are participating in replication.
- **Seal Wrap**: This option enables [Seal Wrap](/docs/enterprise/sealwrap/index.html) on the mountpoint. It can only be selected at mount time.
- **More options**
  - **Default Lease TTL**: This is the *role* based default lease TTL. It can be set a value lower than the system default TTL, but not higher. See [Token Time-To-Live, Periodic Tokens, and Explicit Max TTLs](/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) for more details.
  - **Maximum Lease TTL**: This is the *role* based maximum lease TTL. It can be set a value lower than the system default TTL, but not higher. See [Token Time-To-Live, Periodic Tokens, and Explicit Max TTLs](/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) for more details.

When ready to enable the Secretes Engine, select **Enable Engine**.

![](/images/vault-ui-guide/vault-ui-secrets-enable3.png)

You'll note that for the version of Vault used in this demonstration (Enterprise Premium v0.10.2), we receive both a success message for the enabling of our Secrets Engine, but also an informational message that lets us know that this version of Vault does not yet contain full KV version 2 support, so it only operates on the latest version of a secret when using KV version 2:

> "secret-demo/" is a newer version of the KV backend. The Vault UI does not currently support the additional versioning features. All actions taken through the UI in this engine will operate on the most recent version of a secret.

### Secrets Engine Configuration

![](/images/vault-ui-guide/vault-ui-secret-demo0.png)

As you'll note from the screenshot, you can determine useful configuration about a Secrets Engine instance, by selecting **Configuration** when on the page for any given Secrets Engine.

### Create a Secret

![](/images/vault-ui-guide/vault-ui-create-secret0.png)

To begin the process of creating a secret make sure that you are on the **Secrets** tab and then select **Create secret**; let's break down some components of the secret:
