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

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	scaleway := m.(*api.ScalewayAPI)

	image := d.Get("image").(string)

	var volumes map[string]string

	for k, v := range d.Get("volumes").(map[string]interface{}) {
		volumes[k] = v.(string)
	}

	id, err := scaleway.PostServer(api.ScalewayServerDefinition{
		Name:    d.Get("name").(string),
		Image:   &image,
		Volumes: volumes,
	})
	if err != nil {
		return err
	}

	err = scaleway.PostServerAction(id, "poweron")
	if err != nil {
		return err
	}

	d.SetId(id)
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	scaleway := m.(*api.ScalewayAPI)

	_, err := scaleway.GetServer(d.Id())

	if err != nil {
		serr := err.(api.ScalewayAPIError)

		// if the resource was destroyed, destroy the resource locally
		if serr.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}
	// TODO: set fields

	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	scaleway := m.(*api.ScalewayAPI)

	// TODO: does this delete associated volumes???
	err := scaleway.DeleteServerSafe(d.Id())
	if err != nil {
		return err
	}
	return nil
}
