package opc

import (
	"fmt"

	"github.com/hashicorp/go-oracle-terraform/client"
	"github.com/hashicorp/go-oracle-terraform/compute"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIPReservation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPReservationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"permanent": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_pool": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsComputedSchema(),
			"ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"used": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIPReservationRead(d *schema.ResourceData, meta interface{}) error {
	computeClient := meta.(*Client).computeClient.IPReservations()
	name := d.Get("name").(string)

	input := compute.GetIPReservationInput{
		Name: name,
	}

	result, err := computeClient.GetIPReservation(&input)
	if err != nil {
		// IP Reservation does not exist
		if client.WasNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error reading ip reservation %s: %s", d.Id(), err)
	}

	if result == nil {
		d.SetId("")
		return nil
	}

	d.SetId(name)
	d.Set("parent_pool", result.ParentPool)
	d.Set("permanent", result.Permanent)

	if err := setStringList(d, "tags", result.Tags); err != nil {
		return err
	}

	d.Set("ip", result.IP)
	d.Set("used", result.Used)
	d.Set("fqdn", result.FQDN)
	return nil
}
