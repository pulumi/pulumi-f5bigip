module github.com/pulumi/pulumi-f5bigip

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.8.0
	github.com/pulumi/pulumi-terraform-bridge v1.5.2
	github.com/terraform-providers/terraform-provider-bigip v1.1.1
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
