package kms

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"kms_data_key_without_plaintext":      resourceDataKeyWithoutPlaintext(),
			"kms_data_key_pair_without_plaintext": resourceDataKeyPairWithoutPlaintext(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
