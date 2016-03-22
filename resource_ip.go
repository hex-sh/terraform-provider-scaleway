package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

func resourceIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpCreate,
		Read:   resourceIpRead,
		Update: resourceIpUpdate,
		Delete: resourceIpDelete,
		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIpCreate(d *schema.ResourceData, m interface{}) error {
	scaleway := m.(*api.ScalewayAPI)
	ip, err := scaleway.NewIP()
	if err != nil {
		return err
	}
	d.SetId(ip.IP.ID)
	return resourceIpUpdate(d, m)
}

func resourceIpRead(d *schema.ResourceData, m interface{}) error {
	scaleway := m.(*api.ScalewayAPI)
	ip, err := scaleway.GetIP(d.Id())
	if err != nil {
		return err
	}
	d.Set("ip", ip.IP.Address)
	return nil
}

func resourceIpUpdate(d *schema.ResourceData, m interface{}) error {
	server := m.(*api.ScalewayAPI)
	server.AttachIP(d.Id(), d.Get("server").(string))
	return resourceIpRead(d, m)
}

func resourceIpDelete(d *schema.ResourceData, m interface{}) error {
	server := m.(*api.ScalewayAPI)
	err := server.DeleteIP(d.Id())
	if err != nil {
		return err
	}
	return nil
}
