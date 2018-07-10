package netapp

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudVolume_nfs_import(t *testing.T) {
	envName := os.Getenv("NETAPP_VSA_WORKENV_NAME")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCloudVolume_nfs_vsa_config(envName),
			},
		},
	})
}

func TestAccCloudVolume_cifs_import(t *testing.T) {
	envName := os.Getenv("NETAPP_VSA_WORKENV_NAME")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCloudVolume_cifs_vsa_config(envName),
			},
		},
	})
}
