package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scaleway/scaleway-cli/pkg/api"
	"github.com/scaleway/scaleway-cli/pkg/scwversion"
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
		},
		ResourcesMap: map[string]*schema.Resource{
			"scaleway_server": resourceServer(),
			"scaleway_ip":     resourceIp(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return api.NewScalewayAPI(
				d.Get("organization").(string),
				d.Get("token").(string),
				scwversion.UserAgent(),
			)
		},
	}
}
