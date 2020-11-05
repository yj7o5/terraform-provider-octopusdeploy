package octopusdeploy

import (
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAWSAccountBasic(t *testing.T) {
	localName := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	prefix := constOctopusDeployAWSAccount + "." + localName

	accessKey := acctest.RandString(10)
	name := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	secretKey := acctest.RandString(10)
	tenantedDeploymentParticipation := octopusdeploy.TenantedDeploymentModeTenantedOrUntenanted

	resource.Test(t, resource.TestCase{
		CheckDestroy: testAccountDestroy,
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAWSAccountBasic(localName, name, accessKey, secretKey, tenantedDeploymentParticipation),
				Check: resource.ComposeTestCheckFunc(
					testAccountExists(prefix),
					resource.TestCheckResourceAttr(prefix, constAccessKey, accessKey),
					resource.TestCheckResourceAttr(prefix, constName, name),
					resource.TestCheckResourceAttr(prefix, constSecretKey, secretKey),
					resource.TestCheckResourceAttr(prefix, constTenantedDeploymentParticipation, string(tenantedDeploymentParticipation)),
				),
			},
		},
	})
}

func testAWSAccountBasic(localName string, name string, accessKey string, secretKey string, tenantedDeploymentParticipation octopusdeploy.TenantedDeploymentMode) string {
	return fmt.Sprintf(`resource "%s" "%s" {
		access_key = "%s"
		name = "%s"
		secret_key = "%s"
		tenanted_deployment_participation = "%s"
	}`, constOctopusDeployAWSAccount, localName, accessKey, name, secretKey, tenantedDeploymentParticipation)
}
