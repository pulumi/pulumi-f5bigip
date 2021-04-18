module github.com/pulumi/pulumi-f5bigip/provider/v3

go 1.16

require (
	github.com/F5Networks/terraform-provider-bigip v1.3.3
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0-beta.1
	github.com/pulumi/pulumi/sdk/v3 v3.0.0-beta.2
)

replace (
	github.com/F5Networks/terraform-provider-bigip => github.com/pulumi/terraform-provider-bigip v1.1.1-0.20210401070204-bf6448713f79
)
