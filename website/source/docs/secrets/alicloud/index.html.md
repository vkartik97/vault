---
layout: "docs"
page_title: "AliCloud - Secrets Engines"
sidebar_current: "docs-secrets-alicloud"
description: |-
  The AliCloud secrets engine for Vault generates access tokens or STS credentials 
  dynamically based on RAM policies or roles.
---

# AliCloud Secrets Engine

The AliCloud secrets engine dynamically generates AliCloud access tokens based on RAM
policies, or AliCloud STS credentials based on RAM roles. This generally 
makes working with AliCloud easier, since it does not involve clicking in the web UI. 
The AliCloud access tokens are time-based and are automatically revoked when the Vault 
lease expires. STS credentials are short-lived, non-renewable expire on their own.

## Setup

Most secrets engines must be configured in advance before they can perform their
functions. These steps are usually completed by an operator or configuration
management tool.

1. Enable the AliCloud secrets engine:

    ```text
    $ vault secrets enable alicloud
    Success! Enabled the alicloud secrets engine at: alicloud/
    ```

    By default, the secrets engine will mount at the name of the engine. To
    enable the secrets engine at a different path, use the `-path` argument.
    
1. [Create a custom policy](https://www.alibabacloud.com/help/doc-detail/28640.htm) 
in AliCloud that will be used for the access key you will give Vault. See "Example 
RAM Policy for Vault".

1. [Create a user](https://www.alibabacloud.com/help/faq-detail/28637.htm) in AliCloud 
with a name like "hashicorp-vault", and directly apply the new custom policy to that user
in the "User Authorization Policies" section.

1. Create an access key for that user in AliCloud, which is an action available in
AliCloud's UI on the user's page.

1. Configure that access key as the credentials that Vault will use to communicate with 
AliCloud to generate credentials:

    ```text
    $ vault write alicloud/config \
        access_key=AKIAJWVN5Z4FOFT7NLNA \
        secret_key=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
    ```

1. Configure a role describing how credentials will be granted. 

    To generate access tokens using only policies that have already been created in AliCloud:
    
        ```text
        $ vault write alicloud/role/policy-based \
            remote_policies='name:AliyunOSSReadOnlyAccess,type:System' \
            remote_policies='name:AliyunRDSReadOnlyAccess,type:System'
        ```
    To generate access tokens using only policies that will be dynamically created in AliCloud by
    Vault:
    
        ```text
        $ vault write alicloud/role/policy-based \
            inline_policies=-<<EOF
        [
            {
              "Statement": [
                {
                  "Action": "rds:Describe*",
                  "Effect": "Allow",
                  "Resource": "*"
                }
              ],
              "Version": "1"
            },
            {...}
        ]
        EOF
        ```
    Both `inline_policies` and `remote_policies` may be used together. However, neither may be
    used configuring how to generate STS credentials, like so:
    
        ```text
        $ vault write alibaba/role/role-based \
              role_arn='acs:ram::5138828231865461:role/hastrustedactors'
        ```
    Any `role_arn` specified must have added "trusted actors" when it was being created. These
    can only be added at role creation time. Trusted actors are entities that can assume the role.
    Since we will be assuming the role to gain credentials, the `access_key` and `secret_key` in
    the config must qualify as a trusted actor.
    
    - [More on roles](https://www.alibabacloud.com/help/doc-detail/28649.htm)
    - [More on policies](https://www.alibabacloud.com/help/doc-detail/28652.htm)
    
### Example RAM Policy for Vault

The `alicloud/config` credentials need permissions that vary based on the ways
roles are configured.

This is an example RAM policy that would allow you to create credentials using
any type of role:

TODO - add this and test it too
```json
{
  "Statement": [
    {
      "Action": [
        "ram:CreateAccessKey",
        "ram:DeleteAccessKey",
        "ram:CreatePolicy",
        "ram:DeletePolicy",
        "ram:AttachPolicyToUser",
        "ram:DetachPolicyFromUser",
        "ram:CreateUser",
        "ram:DeleteUser",
        "sts:AssumeRole"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ],
  "Version": "1"
}
```
However, the policy you use should only allow the actions you actually need
for how your roles are configured.

If any roles are using `inline_policies`, you need the following actions:
- `"ram:CreateAccessKey"`
- `"ram:DeleteAccessKey"`
- `"ram:AttachPolicyToUser"`
- `"ram:DetachPolicyFromUser"`
- `"ram:CreateUser"`
- `"ram:DeleteUser"`

If any roles are using `remote_policies`, you need the following actions:
- All listed for `inline_policies`
- `"ram:CreatePolicy"`
- `"ram:DeletePolicy"`

If any roles are using `role_arn`, you need the following actions:
- `"sts:AssumeRole"`

## Usage

After the secrets engine is configured and a user/machine has a Vault token with
the proper permission, it can generate credentials.

1. Generate a new access key by reading from the `/creds` endpoint with the name
of the role:

    ```text
    TODO check these through, same with below
    $ vault read alicloud/creds/policy-based
    Key                Value
    ---                -----
    lease_id           alicloud/creds/policy-based/f3e92392-7d9c-09c8-c921-575d62fe80d8
    lease_duration     768h
    lease_renewable    true
    access_key         AKIAIOWQXTLW36DV7IEA
    secret_key         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
    ```

    The `access_key` and `secret_key` returned are also known is an 
    `"AccessKeyId"`and `"AccessKeySecret"`, respectively, in the Alibaba's
    docs. 
    
    Retrieving creds for a role using a `role_arn` will carry the additional 
    fields of `expiration` and `security_token`, like so:
    
    ```text
        $ vault read alicloud/creds/role-based
        Key                Value
        ---                -----
        lease_id           alicloud/creds/my-role/f3e92392-7d9c-09c8-c921-575d62fe80d9
        lease_duration     768h
        lease_renewable    false
        access_key         AKIAIOWQXTLW36DV7IEA
        secret_key         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
        security_token     CAESrAIIARKAAShQquMnLIlbvEcIxO6wCoqJufs8sWwieUxu45hS9A...
        expiration         TODO I need an example of this
        ```

## API

The AliCloud secrets engine has a full HTTP API. Please see the
[AliCloud secrets engine API](/api/secret/alicloud/index.html) for more
details.
