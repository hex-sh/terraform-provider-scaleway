package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"api_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://api.scaleway.com/",
			},
			"account_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://account.scaleway.com/",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scaleway_server": resourceServer(),
			"scaleway_ip":     resourceIp(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return api.NewScalewayAPI(
				d.Get("api_endpoint").(string),
				d.Get("account_endpoint").(string),
				d.Get("organization").(string),
				d.Get("token").(string),
			)
		},
	}
}
