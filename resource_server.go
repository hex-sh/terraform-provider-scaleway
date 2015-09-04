package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
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

const baseUrl = "https://api.scaleway.com"

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(*Config)
	name := d.Get("name").(string)
	image := d.Get("image").(string)
	fields := map[string]interface{}{
		"organization": config.Organization,
		"name":         name,
		"image":        image,
		"volumes":      nil,
	}

	data, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	res, err := http.Post(baseUrl+"/servers", "application/json", bytes.NewReader(data))

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode >= 400 {
		return fmt.Errorf("%s:\n%s", res.Status, body)
	}

	if err != nil {
		return err
	}

	var resData map[string]interface{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return err
	}

	id := resData["id"].(string)
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
