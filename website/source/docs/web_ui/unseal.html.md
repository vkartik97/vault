---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-unseal"
description: |-
  Unsealing is the process of constructing the master key necessary to read the
  decryption key to decrypt the data, allowing access to the Vault.
---


## Unseal

When Vault is sealed, the UI indicates that **Vault is sealed** as shown below:

![](/assets/images/vault-ui-guide/vault-ui-init3.png)


You need to unseal Vault using the unseal keys generated during the [initialization](/docs/web_ui/init.html.md).

![](/assets/images/vault-ui-guide/vault-ui-init4.png)

Each unseal key holder must enter his/her unseal key in the **Key** field, and
then click **Unseal**.

Once the configured threshold has been reached, the UI displays the **Sign
in to Vault** dialog.


## Next Step

Now, you are ready to [sign in](/docs/web_ui/sign-in.html).
