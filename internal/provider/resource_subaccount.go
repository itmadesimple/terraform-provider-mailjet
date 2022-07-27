package provider

import (
	"context"
	"strconv"
	"time"

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

		Schema: map[string]*schema.Schema {
			"name": {
				Description: "User readable name for this API Key.",
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
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
			"created_at": {
				Description: "Timestamp indicating when the API Key was created.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"api_key": {
				Description: "The unique alphanumeric API Key itself.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"secret_key": {
				Description: "The unique alphanumeric Secret Key (password) linked to this API Key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"primary": {
				Description: "Indicates whether this is the Primary API Key or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"quarantine_value": {
				Description: "Indicates the sending limitation applied on this API Key in terms of messages per hour, when the API Key is under quarantine. 0 means no limitation is applie",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"run_level": {
				Description: "Runlevel, used to indicate data is migrated and reduced performance is expected. Possible values: \"Normal\", \"Softlock\", \"Hardlock\". Default value: \"Normal\".",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"track_host": {
				Description: "Host to be used when tracking clicks, opens, unsubscribe requests for this API Key. Default value: \"r.mailjet.com\".",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"user_id": {
				Description: "User ID. The Primary API Key and all sub-account API Keys will have the same user ID.",
				Type:        schema.TypeInt,
				Computed:    true,
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
		return diag.FromErr(err)
	}

	idFromAPI := strconv.FormatInt(data[0].ID, 10)
	d.SetId(idFromAPI)

	tflog.Trace(ctx, "created a Mailjet sub account with ID " + idFromAPI)

	return resourceMailjetSubaccountRead(ctx, d, meta)
}

func resourceMailjetSubaccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mailjetClient := meta.(*mailjet.Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		tflog.Error(ctx, err.Error())
		return diag.FromErr(err)
	}

	var data []resources.Apikey
	mr := &mailjet.Request{
		Resource: "apikey",
		ID:       id,
	}
	err = mailjetClient.Get(mr, &data)
	if err != nil {
		tflog.Error(ctx, err.Error())
		return diag.FromErr(err)
	}

	d.Set("active", data[0].IsActive)
	d.Set(("acl"), data[0].ACL)
	d.Set(("created_at"), data[0].CreatedAt.Format(`"` + time.RFC3339 + `"`))
	d.Set(("api_key"), data[0].APIKey)
	d.Set(("secret_key"), data[0].SecretKey)
	d.Set(("primary"), data[0].IsMaster)
	d.Set(("quarantine_value"), data[0].QuarantineValue)
	d.Set(("run_level"), data[0].Runlevel)
	d.Set(("track_host"), data[0].TrackHost)
	d.Set(("user_id"), data[0].UserID)

	return nil
}

func resourceMailjetSubaccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mailjetClient := meta.(*mailjet.Client)

	active := d.Get("active").(bool)
	acl := d.Get("acl").(string)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		tflog.Error(ctx, err.Error())
		return diag.FromErr(err)
	}
	mr := &mailjet.Request{
		Resource: "apikey",
		ID:       id,
	}
	fmr := &mailjet.FullRequest{
		Info: mr,
		Payload: &resources.Apikey{
			IsActive: active,
			ACL: acl,
		},
	}
	err = mailjetClient.Put(fmr, nil)
	if err != nil {
		tflog.Error(ctx, err.Error())
		return diag.FromErr(err)
	}

	return resourceMailjetSubaccountRead(ctx, d, meta)
}

func resourceMailjetSubaccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
