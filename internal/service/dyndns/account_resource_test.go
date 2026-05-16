package dyndns_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDyndnsAccountResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDyndnsAccountResourceConfig("acc-test-1", false, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "description", "acc-test-1"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "service", "dyndns2"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "server", "members.dyndns.org"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "username", "acc-user"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "checkip", "web_dyndns"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "hostnames.#", "1"),
					resource.TestCheckTypeSetElemAttr("opnsense_dyndns_account.test", "hostnames.*", "test.example.com"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "wildcard", "false"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "force_ssl", "false"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "ttl", "300"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "checkip_timeout", "10"),
					resource.TestCheckResourceAttrSet("opnsense_dyndns_account.test", "id"),
				),
			},
			{
				ResourceName:            "opnsense_dyndns_account.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: testAccDyndnsAccountResourceConfig("acc-test-2", true, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "description", "acc-test-2"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "wildcard", "true"),
					resource.TestCheckResourceAttr("opnsense_dyndns_account.test", "force_ssl", "true"),
				),
			},
		},
	})
}

func TestAccDyndnsAccountDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDyndnsAccountDataSourceConfig("acc-ds-test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.opnsense_dyndns_account.test", "id",
						"opnsense_dyndns_account.test", "id"),
					resource.TestCheckResourceAttr("data.opnsense_dyndns_account.test", "description", "acc-ds-test"),
					resource.TestCheckResourceAttr("data.opnsense_dyndns_account.test", "service", "dyndns2"),
					resource.TestCheckResourceAttr("data.opnsense_dyndns_account.test", "ttl", "300"),
				),
			},
		},
	})
}

func testAccDyndnsAccountResourceConfig(description string, wildcard, forceSSL bool) string {
	return fmt.Sprintf(`
resource "opnsense_dyndns_account" "test" {
  description = %[1]q
  service     = "dyndns2"
  server      = "members.dyndns.org"
  username    = "acc-user"
  password    = "acc-pass"
  hostnames   = ["test.example.com"]
  checkip     = "web_dyndns"
  wildcard    = %[2]t
  force_ssl   = %[3]t
}
`, description, wildcard, forceSSL)
}

func testAccDyndnsAccountDataSourceConfig(description string) string {
	return fmt.Sprintf(`
resource "opnsense_dyndns_account" "test" {
  description = %[1]q
  service     = "dyndns2"
  server      = "members.dyndns.org"
  username    = "acc-user"
  password    = "acc-pass"
  hostnames   = ["ds.example.com"]
  checkip     = "web_dyndns"
}

data "opnsense_dyndns_account" "test" {
  id = opnsense_dyndns_account.test.id
}
`, description)
}
