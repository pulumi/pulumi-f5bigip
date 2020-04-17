module github.com/pulumi/pulumi-f5bigip/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.0
	github.com/terraform-providers/terraform-provider-bigip v1.1.2
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
