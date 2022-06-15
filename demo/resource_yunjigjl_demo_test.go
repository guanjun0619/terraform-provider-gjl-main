package demo

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccYunjiDemo_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPackResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testYunjiDemoPackConfig,
				Check: resource.ComposeTestCheckFunc(
					testEndpointExist("yunjigjl_demo.test"),
					resource.TestCheckResourceAttr("yunjigjl_demo.test", "instance_name", "guanguan"),
					resource.TestCheckResourceAttr("yunjigjl_demo.test", "disk_size", "100"),
					resource.TestCheckResourceAttr("yunjigjl_demo.test", "tags", "test"),
				),
			},
		},
	})
}

var testYunjiDemoPackConfig = `
resource "yunjigjl_demo" "test" {
  instance_name  = "yunji"
  disk_size = 100
}
`

func testEndpointExist(rn string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[rn]
		if !ok {
			return fmt.Errorf("resource not found: %s", rn)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("resource id not set")
		}
		return nil
	}
}

func testAccCheckPackResourceDestroy(s *terraform.State) error {
	return nil
}
