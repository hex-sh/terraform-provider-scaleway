package main

import (
	"bytes"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
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
			},
		},
	}

}

var baseUrl = "https://api.scaleway.com"

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	name := d.Get("name").(string)
	imag := d.Get("image").(string)
	fields := map[string]interface{}{
		"organization": config.Organization,
		"name":         name,
		"image":        image,
	}

	data, err := json.Marshal(fields)
	if err {
		return err
	}

	resp, err := http.Post(baseUrl+"/servers", "application/json", bytes.NewReader(data))
	if err {
		return err
	}

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
