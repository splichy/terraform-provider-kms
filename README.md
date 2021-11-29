This is a (hopefully temporary) [Terraform](https:/terraform.io) provider for working with [AWS KMS](https://aws.amazon.com/kms/), particularly for generating [data keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-keys).
It attempts to correct a [deficiency](https://github.com/hashicorp/terraform-provider-aws/issues/21912) in the [AWS Terraform provider](https://registry.terraform.io/providers/hashicorp/aws/latest) in this area.

As this is not intended to be a permanent provider, it is not particularly general-purpose.
We will accept reasonable PRs to expand the provider's scope, however large-scale surgery is beyond the intended scope.


# Installation

It should, ideally, be as simple as telling your Terraform module to use the provider:

```hcl
terraform {
  required_providers {
    kms = {
      source  = "cipherstash/kms"
      version = "~> 0.1"
    }
  }
}
```

Then a quick `terraform init` should see it installed and ready to use.


# Usage

See the [registry docs](https://registry.terraform.io/providers/cipherstash/kms/latest) for all the gory details.


# Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md).


# Licence

Unless otherwise stated, everything in this repo is covered by the following
copyright notice:

    Copyright (C) 2021  CipherStash Pty Ltd.

    This program is free software: you can redistribute it and/or modify it
    under the terms of the GNU General Public License version 3, as
    published by the Free Software Foundation.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
