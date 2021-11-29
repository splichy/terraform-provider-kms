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

func resourceDataKeyWithoutPlaintext() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataKeyWithoutPlaintextCreate,
		Read:   schema.Noop,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"ciphertext_blob": {
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
			"key_spec": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      false,
				ValidateFunc: validation.StringInSlice(kms.DataKeySpec_Values(), false),
				ForceNew:     true,
			},
			"number_of_bytes": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 1024),
				ForceNew:     true,
			},
		},
	}
}

func resourceDataKeyWithoutPlaintextCreate(d *schema.ResourceData, meta interface{}) error {
	conn := kms.New(session.Must(session.NewSession()))

	input := &kms.GenerateDataKeyWithoutPlaintextInput{
		KeyId: aws.String(d.Get("key_id").(string)),
	}

	if v, ok := d.GetOk("key_spec"); ok {
		input.KeySpec = aws.String(v.(string))
	}

	if v, ok := d.GetOk("number_of_bytes"); ok {
		input.NumberOfBytes = aws.Int64(int64(v.(int)))
	}

	log.Printf("[DEBUG] Generating KMS data key without plaintext")

	outputRaw, err := conn.GenerateDataKeyWithoutPlaintext(input)

	if err != nil {
		return fmt.Errorf("error generating KMS data key: %w", err)
	}

	b64blob := base64.RawStdEncoding.EncodeToString(outputRaw.CiphertextBlob)
	d.Set("ciphertext_blob", b64blob)
	d.SetId(b64blob[6:35])

	return nil
}
