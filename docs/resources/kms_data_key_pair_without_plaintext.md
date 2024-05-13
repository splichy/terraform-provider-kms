---
page_title: "kms_data_key_pair_without_plaintext Resource - kms"
subcategory: ""
description: |-
 Create encrypted KMS data key pair.
---

# kms_data_key_pair_without_plaintext (Resource)

Generates a asymmetric data key pair, encrypted with the specified KMS key, without returning the plaintext key to Terraform.
This ensures that the key is both properly random, and prevents Terraform from learning anything about the key material itself.


## Example Usage

```terraform
resource "aws_kms_key" "a" {
  description = "KMS key 1"
}

resource "kms_data_key_pair_without_plaintext" "a" {
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
- **key_spec** (String) Specify that the data key will be used for the given asymmetric encryption algorithm, which is used to determine the size of the data key.
  Valid values: `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, `ECC_SECG_P256K1` or `SM2`.

### Read-Only

- **id** (String) The ID of this resource. It is automatically derived from the data key itself, and so will change if the data key is regenerated.
- **ciphertext_blob** (String) The encrypted data key, encoded using base64.
- **public_key** (String) Public key, base64 encoded.
