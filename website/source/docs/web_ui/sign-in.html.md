---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-sign-in"
description: |-
  This
---


## Sign in

![](/assets/images/vault-ui-guide/vault-ui-signin0.png)

When Vault is unsealed, the sign in dialog is displayed.

This dialog provides a means to authenticate to Vault by singing in with enabled authentication methods. As shown in the screenshot, the default authentication method is **Token**, which accepts a valid Vault token. If you have other authentication methods configured, you'll notice them listed here for use as well.

Note that in the example, there are five total authentication methods enabled, including **Token**, **Userpass**, **LDAP**, **Okta**, and **GitHub**. We'll review each type for the purposes of detailing the differences so that you have some idea of what to expect from each.

### Token

To sign into Vault with a token, enter a _valid Vault token_ in the **Token** field and select **Sign in**.

### Userpass

![](/assets/images/vault-ui-guide/vault-ui-signin1.png)

To sign into Vault with the **Userpass** auth method, enter a valid username and password into the **Username** and **Password** fields, and an optional mount path entered into the **Mount path** field if the Userpass auth method is mounted at a non-default path.

If you do not have this information, please ask your Vault administrator.

### LDAP

![](/assets/images/vault-ui-guide/vault-ui-signin2.png)

To sign into Vault with the **LDAP** auth method, enter a valid LDAP or Active Directory username and password into the **Username** and **Password** fields, and an optional mount path entered into the **Mount path** field if the Userpass auth method is mounted at a non-default path.

If you do not have this information, please ask your Vault administrator.

### Okta

![](/assets/images/vault-ui-guide/vault-ui-signin3.png)

To sign into Vault with the **Okta** auth method, enter a valid Okta username and password into the **Username** and **Password** fields, and an optional mount path entered into the **Mount path** field if the Userpass auth method is mounted at a non-default path.

If you do not have this information, please ask your Vault administrator.

### GitHub

![](/assets/images/vault-ui-guide/vault-ui-signin4.png)

To sign into Vault with the **GitHub** auth method, enter a valid GitHub user token into the **GitHub Token** field, and an optional mount path entered into the **Mount path** field if the Userpass auth method is mounted at a non-default path.

If you do not have this information, please ask your Vault administrator.

![](/assets/images/vault-ui-guide/vault-ui-signin5.png)

Upon successful sign in, the UI should be available for use. You will note that the screenshot is displaying a warning because signin was with a token and that token was a *root token*; usage of root tokens should be extremely guarded, so a warning is displayed:

> Attention
> You have logged in with a root token. As a security precaution, this root token will not be stored by your browser and you will need to re-authenticate after the window is closed or refreshed.

This particular warning dialog will not be displayed when signing in with _any token that is not a root token_.

See the [Root Tokens section](https://www.vaultproject.io/docs/concepts/tokens.html#root-tokens) of Vault's [Tokens documentation](https://www.vaultproject.io/docs/concepts/tokens.html) for more details about the significance of root tokens with Vault.
