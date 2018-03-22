package netapp

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccCloudVolume_nfs_import(t *testing.T) {
	envName := os.Getenv("NETAPP_VSA_WORKENV_NAME")
	resourceName := "netapp_cloud_volume.vsa-nfs-volume"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCloudVolume_nfs_vsa, envName),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudVolume_cifs_import(t *testing.T) {
	envName := os.Getenv("NETAPP_VSA_WORKENV_NAME")
	resourceName := "netapp_cloud_volume.vsa-cifs-volume"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCloudVolume_cifs_vsa, envName),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
