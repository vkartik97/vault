---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-config"
description: |-
  This
---

## Configure Vault

An update to the Vault server configuration is required before you can use the Vault UI. This is documented in the [Vault UI](https://www.vaultproject.io/docs/configuration/ui/index.html) configuration documentation. Configuration essentially amounts to placing one option in the global scope of Vault's configuration file:

```
ui = true

...
```

This configuration change requires restarting the Vault process.
