package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

func resourceMailjetSubaccount() *schema.Resource {
	return &schema.Resource{
		Description: "Mailjet Sub Account Resource",

		CreateContext: resourceMailjetSubaccountCreate,
		ReadContext:   resourceMailjetSubaccountRead,
		UpdateContext: resourceMailjetSubaccountUpdate,
		DeleteContext: resourceMailjetSubaccountDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "User readable name for this API Key.",
				Type:        schema.TypeString,
				Optional:    false,
				ForceNew:    false,
			},
			"active": {
				Description: "Indicates whether this API Key is active or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    false,
			},
			"acl": {
				Description: "Access Control List. Indicates permissions attached to a resource.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
			},
		},
	}
}

func resourceMailjetSubaccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mailjetClient := meta.(*mailjet.Client)

	name := d.Get("name").(string)
	active := d.Get("active").(bool)
	acl := d.Get("acl").(string)

	var data []resources.Apikey
	mr := &mailjet.Request{
		Resource: "apikey",
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.Apikey{
			Name: name,
			IsActive: active,
			ACL: acl,
		},
	}
	err := mailjetClient.Post(fmr, &data)
	if err != nil {
		tflog.Error(ctx, err.Error())
	}

	idFromAPI := strconv.FormatInt(data[0].ID, 10)
	d.SetId(idFromAPI)

	tflog.Trace(ctx, "created a Mailjet sub account")

	return resourceMailjetSubaccountRead(ctx, d, meta)
}

func resourceMailjetSubaccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceMailjetSubaccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceMailjetSubaccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
