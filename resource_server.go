package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/scaleway/scaleway-cli/pkg/api"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"volumes": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}

}

const baseUrl = "https://api.scaleway.com"

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	s := m.(*api.ScalewayAPI)
	name := d.Get("name").(string)
	image := d.Get("image").(string)

	id, err := s.PostServer(api.ScalewayServerDefinition{
		Name:  name,
		Image: &image,
		Volumes: map[string]string{
			"0": "1",
		},
	})

	if err != nil {
		return err
	}
	d.SetId(id)
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
