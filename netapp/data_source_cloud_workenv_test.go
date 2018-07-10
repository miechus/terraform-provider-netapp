package netapp

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetAppCloudWorkingEnvironmentDataSource(t *testing.T) {
	envName := os.Getenv("NETAPP_AWSHA_WORKENV_NAME")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetAppCloudWorkingEnvironmentDataSource_config(envName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWoringEnvironmentDataSourceID("data.netapp_cloud_workenv.aws-ha-env"),
					resource.TestCheckResourceAttr(
						"data.netapp_cloud_workenv.aws-ha-env", "name", envName),
					resource.TestCheckResourceAttrSet(
						"data.netapp_cloud_workenv.aws-ha-env", "public_id"),
					resource.TestCheckResourceAttrSet(
						"data.netapp_cloud_workenv.aws-ha-env", "tenant_id"),
					resource.TestCheckResourceAttrSet(
						"data.netapp_cloud_workenv.aws-ha-env", "svm_name"),
					resource.TestCheckResourceAttrSet(
						"data.netapp_cloud_workenv.aws-ha-env", "is_ha"),
				),
			},
		},
	})
}

func testAccCheckWoringEnvironmentDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find network data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Working environment data source ID not set")
		}

		return nil
	}
}

func testAccNetAppCloudWorkingEnvironmentDataSource_config(envName string) string {
	c := `
data "netapp_cloud_working_environment" "aws-ha-env" {
        name = "%s"
}
`
	return FormatString(c, envName)
}
