package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token" & schema.Schema{},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scaleway_organization":   organizationResource,
			"scaleway_server":         serverResource,
			"scaleway_volume":         volumeResource,
			"scaleway_image":          imageResource,
			"scaleway_ip":             ipResource,
			"scaleway_security_group": scalewaySecurityGroup,
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerFunc(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token: d.Get("token")(string),
	}
	return config.Client()
}
