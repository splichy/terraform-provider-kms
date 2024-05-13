package kms

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDataKeyPairWithoutPlaintext() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataKeyPairWithoutPlaintextCreate,
		Read:   schema.Noop,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"ciphertext_blob": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"keepers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_pair_spec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      false,
				ValidateFunc: validation.StringInSlice(kms.DataKeyPairSpec_Values(), false),
				ForceNew:     true,
			},
		},
	}
}

func resourceDataKeyPairWithoutPlaintextCreate(d *schema.ResourceData, meta interface{}) error {
	conn := kms.New(session.Must(session.NewSession()))

	input := &kms.GenerateDataKeyPairWithoutPlaintextInput{
		KeyId: aws.String(d.Get("key_id").(string)),
	}

	if v, ok := d.GetOk("key_pair_spec"); ok {
		input.KeyPairSpec = aws.String(v.(string))
	}

	log.Printf("[DEBUG] Generating KMS data key pair without plaintext")

	outputRaw, err := conn.GenerateDataKeyPairWithoutPlaintext(input)

	if err != nil {
		return fmt.Errorf("error generating KMS data key pair: %w", err)
	}

	b64blob := base64.RawStdEncoding.EncodeToString(outputRaw.PrivateKeyCiphertextBlob)
	d.Set("ciphertext_blob", b64blob)
	pkBlob := base64.RawStdEncoding.EncodeToString(outputRaw.PublicKey)
	d.Set("public_key", pkBlob)
	d.SetId(b64blob[6:35])

	return nil
}
