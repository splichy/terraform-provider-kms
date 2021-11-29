---
page_title: "kms_data_key_without_plaintext Resource - kms"
subcategory: ""
description: |-
 Create encrypted KMS data keys.
---

# kms_data_key_without_plaintext (Resource)

Generates a symmetric data key, encrypted with the specified KMS key, without returning the plaintext key to Terraform.
This ensures that the key is both properly random, and prevents Terraform from learning anything about the key material itself.


## Example Usage

```terraform
resource "aws_kms_key" "a" {
  description = "KMS key 1"
}

resource "aws_kms_data_key_without_plaintext" "a" {
  key_id   = aws_kms_key.a.id
  key_spec = "AES_256"
}
```

## Resource "Keepers"

This resource will, by default, remember the data key produced by previous Terraform runs and re-use it.
A new data key will be generated on the first run of the resource, or if the value of the `keepers` argument changes.

## Schema

### Required

- **key_id** (String) The ID (or ARN) of the KMS key which will be used to encrypt the generated data key.

### Optional

- **keepers** (Map of String) Arbitrary map of values that, when changed, will trigger the generation of a new data key.
- **key_spec** (String) Specify that the data key will be used for the given symmetric encryption algorithm, which is used to determine the size of the data key.
  Valid values: `AES_128`, `AES_256`.
- **number_of_bytes** (Number) Manually specify the size of the data key, in bytes. Only useful as an alternative to `key_spec`.

### Read-Only

- **id** (String) The ID of this resource. It is automatically derived from the data key itself, and so will change if the data key is regenerated.
- **ciphertext_blob** (String) The encrypted data key, encoded using base64.
