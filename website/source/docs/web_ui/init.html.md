---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-init"
description: |-
  Initialization is the process configuring the Vault. This only happens once when the server is started against a new backend that has never been used with Vault before. When running in HA mode, this happens once per cluster, not per server.
---

## Initialize

If you are not using the default Shamir's secret sharing algorithm to split the
Master Key as described in the [Architecture
documentation](/docs/internals/architecture.html) instead of one of the other
[seal types](/docs/configuration/seal/index.html), the Vault initialization
process can be done through the UI.


1. Launch the UI using your prefer web browser (e.g. `https://vault.example.com:8200`)
    ![](/assets/images/vault-ui-guide/vault-ui-init0.png)

1. Enter the number of key shares the master key will be split across into the
**Key Shares** field
1. Enter the number of key shares required to reconstruct the master key into the
**Key Threshold** field
    ![](/assets/images/vault-ui-guide/vault-ui-init0-2.png)

    > If you wish to encrypt and hex encode the key share output with PGP, you can
    either select the PGP key from a file or enter it as text in the **Encrypt
    Output with PGP** dialog. Likewise, if you wish to encrypt and hex encode the
    initial root token output with PGP you can do th same in the **Encrypt Root
    Token with PGP** dialog.

    ![](/assets/images/vault-ui-guide/vault-ui-init1-2.png)

1. Click **Initialize**
    ![](/assets/images/vault-ui-guide/vault-ui-init2.png)

    > **NOTE:** The output displays the key share(s) and
    initial root token to note for later use. You will also have the options to
    download the keys and proceed on to unsealing Vault.

1. Click **Download Keys** to download JSON file containing the key share(s) and
initial root token with a file name: `vault-cluster-vault-$RFC3339_TIMESTAMP.json`


## Next Step

Once Vault is initialized, it is ready to be [unsealed](/docs/web_ui/unseal.html).
