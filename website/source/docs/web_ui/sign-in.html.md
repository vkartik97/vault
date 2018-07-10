---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-sign-in"
description: |-
  This
---


## Sign in

Once Vault is successfully unsealed, you can sign in. This dialog provides a
means to authenticate to Vault by singing in with enabled authentication
methods. The default authentication method is **Token**, which accepts a valid
Vault token.

If the Vault server has been configured with other [authentication
methods](/docs/auth/index.html), select the desired auth method to sign in.


### Token

To sign into Vault with a token, enter a _valid Vault token_ in the **Token**
field and select **Sign in**.

![Token](/assets/images/vault-ui-guide/vault-ui-signin0.png)

### Userpass

To sign into Vault with the **Userpass** auth method, enter a valid username and
password into the **Username** and **Password** fields, and an optional mount
path entered into the **Mount path** field if the Userpass auth method is
mounted at a non-default path.

![Userpass](/assets/images/vault-ui-guide/vault-ui-signin1.png)

~> If you do not have this information, please ask your Vault administrator.

### LDAP

To sign into Vault with the **LDAP** auth method, enter a valid LDAP or Active
Directory username and password into the **Username** and **Password** fields,
and an optional mount path entered into the **Mount path** field if the Userpass
auth method is mounted at a non-default path.

![LDAP](/assets/images/vault-ui-guide/vault-ui-signin2.png)

~> If you do not have this information, please ask your Vault administrator.

### Okta

To sign into Vault with the **Okta** auth method, enter a valid Okta username
and password into the **Username** and **Password** fields, and an optional
mount path entered into the **Mount path** field if the Userpass auth method is
mounted at a non-default path.

![Okta](/assets/images/vault-ui-guide/vault-ui-signin3.png)

~> If you do not have this information, please ask your Vault administrator.

### GitHub

To sign into Vault with the **GitHub** auth method, enter a valid GitHub user
token into the **GitHub Token** field, and an optional mount path entered into
the **Mount path** field if the Userpass auth method is mounted at a non-default
path.

![Github](/assets/images/vault-ui-guide/vault-ui-signin4.png)

~> If you do not have this information, please ask your Vault administrator.

<br>
---

Upon successful sign in, the UI should be available for use. You will note that the screenshot is displaying a warning because signin was with a token and that token was a *root token*; usage of root tokens should be extremely guarded, so a warning is displayed:

> Attention
> You have logged in with a root token. As a security precaution, this root token will not be stored by your browser and you will need to re-authenticate after the window is closed or refreshed.

![](/assets/images/vault-ui-guide/vault-ui-signin5.png)

This particular warning dialog will not be displayed when signing in with _any token that is not a root token_.

See the [Root Tokens section](/docs/concepts/tokens.html#root-tokens) of Vault's [Tokens documentation](/docs/concepts/tokens.html) for more details about the significance of root tokens with Vault.

## Next Step

Now, you are ready to write some [secrets](/docs/web_ui/secrets.html). 
