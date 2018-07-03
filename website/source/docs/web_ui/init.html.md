---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-init"
description: |-
  This
---

## Initialize

If you are not using the default Shamir's secret sharing algorithm to split the
Master Key as described in the [Architecture
documentation](https://www.vaultproject.io/docs/internals/architecture.html)
instead of one of the other [seal
types](https://www.vaultproject.io/docs/configuration/seal/index.html), the
Vault initialization process can be done through the UI.

Here's how:

![](/assets/images/vault-ui-guide/vault-ui-init0.png)

**FIXME: REPLACE SCREENSHOT**

1. After enabling the UI, start the Vault instance(s)
1. Navigate to `$VAULT_ADDR` in a web browser, where `$VAULT_ADDR` represents the
full URL to your Vault with port, like `https://vault.example.com:8200`, for
example.
1. You should observe a dialog like that in the above screenshot
1. Enter the number of key shares the master key will be split across into the
**Key Shares** field
1. Enter the number of key shares required to reconstruct the master key into the
**Key Threshold** field

If you wish to encrypt and hex encode the key share output with PGP, you can
either select the PGP key from a file or enter it as text in the **Encrypt
Output with PGP** dialog.

Likewise, if you wish to encrypt and hex encode the initial root token output
with PGP you can do th same in the **Encrypt Root Token with PGP** dialog.

Once you've configured the master key generation, use the **Initialize** button
to complete Vault initialization.

~> **NOTE:** The next dialog after initialization with display sensitive
information, including the key shares and initial root token!

![](/assets/images/vault-ui-guide/vault-ui-init1.png)

You should of course use the values required by your organizations policies and
practices, but for demonstration here is the simplest possible case, which is to
specify a value for 1 in both the **Key Shares** and **Key Threshold** fields
and then **Initialize**.

![](/assets/images/vault-ui-guide/vault-ui-init2.png)

After Vault is initialized you'll be presented with the key share(s) and initial
root token to note for later use. You will also have the options to download the
keys and proceed on to unsealing Vault.

If you choose **Download Keys**, your browser will prompt you to download JSON
file containing the key share(s) and initial root token with a filename format
like this:


```
vault-cluster-vault-$RFC3339_TIMESTAMP.json
```


## Next

Once Vault is initialized, it is ready to be unsealed. [Let's do that now!](/docs/web_ui/unseal.html)
