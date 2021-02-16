module github.com/pulumi/pulumi-f5bigip/provider/v2

go 1.14

require (
	github.com/F5Networks/terraform-provider-bigip v1.3.3
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.19.0
	github.com/pulumi/pulumi/sdk/v2 v2.20.1-0.20210212181059-f4b0fa86fedc
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/F5Networks/terraform-provider-bigip => github.com/pulumi/terraform-provider-bigip v1.1.1-0.20210106174146-f28ccf6108e8
)
