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
				Config: fmt.Sprintf(testAccNetAppCloudWorkingEnvironmentDataSource, envName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckWoringEnvironmentDataSourceID("data.netapp_cloud_workenv.aws-ha-env"),
					resource.TestCheckResourceAttr(
						"data.netapp_cloud_workenv.aws-ha-env", "name", envName),
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

const testAccNetAppCloudWorkingEnvironmentDataSource = `
data "netapp_cloud_workenv" "aws-ha-env" {
        name = "%s"
}
`
