---
layout: "docs"
page_title: "Web UI"
sidebar_current: "docs-web_ui-replication-dr"
description: |-
  This
---

## Disaster Recovery (DR) Replication

### Enable DR Primary Replication

Open a web browser and launch the Vault UI (e.g.
https://cluster-A.example.com:8200/ui) and then login.

1. Select **Replication** and check the **Disaster Recovery (DR)** radio button.
  ![DR Replication - primary](/assets/images/vault-dr-1.png)

1. Click **Enable replication**.

1. Select the **Secondaries** tab, and then click **Add**.
  ![DR Replication - primary](/assets/images/vault-dr-2.png)

1. Populate the **Secondary ID** field, and click **Generate token**.
  ![DR Replication - primary](/assets/images/vault-dr-3.png)

1. Click **Copy** to copy the token which you will need to enable the DR secondary cluster.
  ![DR Replication - primary](/assets/images/vault-dr-4.png)


<br>

### Enable DR Secondary Replication

The following operations must be performed on the DR secondary cluster.


1. Now, launch the Vault UI for the **secondary** cluster (e.g. https://cluster-B.example.com:8200/ui) and click **Replication**.

1. Check the **Disaster Recovery (DR)** radio button and select **secondary** under the **Cluster mode**. Paste the token you copied from the primary in the **Secondary activation token** field.
  ![DR Replication - secondary](/assets/images/vault-dr-5.png)

1. Click **Enable replication**.
  ![DR Replication - secondary](/assets/images/vault-dr-5.2.png)

  !> **NOTE:** This will immediately clear all data in the secondary cluster.



### Promote DR Secondary to Primary

This step walks you through the promotion of the secondary cluster to become the
new primary when a catastrophic failure causes the primary cluster to be
inoperable. Refer to the [_Important Note about Automated DR
Failover_](#important) section for more background information.

First, you must generate a **DR operation token** which you need to promote the
secondary cluster. The process, outlined below using API calls, is the similar to [_Generating a Root Token (via CLI)_](/guides/operations/generate-root.html).


1. Click on **Generate OTP** to generate an OTP.  Then click **Copy OTP**.
    ![DR Replication - secondary](/assets/images/vault-dr-6.png)

1. Click **Generate Operation Token**.

1. A quorum of unseal keys must be entered to create a new operation token for
the DR secondary.

    ![DR Replication - secondary](/assets/images/vault-dr-7.png)

    -> This operation must be performed by each unseal-key holder.


1. Once the quorum has been reached, the output displays the encoded DR operation token.  Click **Copy CLI command**.

    ![DR Replication - secondary](/assets/images/vault-dr-8.png)

1. Execute the CLI command from a terminal to generate a DR operation token
using the OTP generated earlier. (Be sure to enter your OTP in the command.)

    **Example:**

    ```
    $ vault operator generate-root \
            -otp="vZpZZf5UI1nvB3A5/7Xq9A==" \          
            -decode="cuplaFGYduDEY6ZVC5IfaA=="

    cf703c0d-afcc-55b9-2b64-d66cf427f59c
    ```

1. Now, click **Promote** tab, and then enter the generated DR operation token.

    ![DR Replication - secondary](/assets/images/vault-dr-9-1.png)

1. Click **Promote cluster**.

    When you prompted, "_Are you sure you want to promote this cluster?_", click **Promote cluster** again to complete.

    ![DR Replication - secondary](/assets/images/vault-dr-9.png)

<br>

> Once the secondary cluster was successfully promoted, you should be able to
log in using the original primary cluster's root token or via configured
authentication method.  If desired, generate a [new root
token](/guides/operations/generate-root.html).



### Demote DR Primary to Secondary

If the _original_ DR primary cluster becomes operational again, you may want to
utilize the cluster by making it a DR secondary cluster. This step explains how
to demote the original DR primary cluster to a secondary.

~> Remember that there is only **one** primary cluster available to the clients
in DR replication.


Select **Replication** and click **Demote cluster**.

![DR Replication - demotion](/assets/images/vault-dr-10.png)

When you prompted, "_Are you sure you want to demote this cluster?_", click
**Demote cluster** again to complete.

![DR Replication - demotion](/assets/images/vault-dr-12.png)


### Disable DR Primary

Once the DR secondary cluster was promoted to be the **new primary**, you may
want to disable the DR replication on the _original_ primary when it becomes
operational again.

~> Remember that there is only **one** primary cluster available to the clients
in DR replication.


Select **Replication** and click **Disable replication**.

![DR Replication - demotion](/assets/images/vault-dr-11.png)

When you prompted, "_Are you sure you want to disable replication on this
cluster?_", click **Disable** again to complete.

![DR Replication - demotion](/assets/images/vault-dr-13.png)

Any secondaries will no longer be able to connect.

!> **Caution:** Once this is done, re-enabling the DR replication as a primary
will change the cluster's ID.  Its connecting secondaries will require a wipe of
the underlying storage even if they have connected before. If re-enabling DR
replication as a secondary, its underlying storage will be wiped when connected
to a primary.
