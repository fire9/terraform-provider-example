package main

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceService() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCreate,
		Read:   resourceServiceRead,
		Update: resourceServiceUpdate,
		Delete: resourceServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(fmt.Sprintf("%x", rand.Int()))
	return resourceServiceRead(d, m)
}

func resourceServiceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServiceUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServiceRead(d, m)
}

func resourceServiceDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
