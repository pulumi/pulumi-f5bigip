module github.com/pulumi/pulumi-f5bigip/provider/v3

go 1.16

require (
	github.com/F5Networks/terraform-provider-bigip v1.3.3
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.25.1
	github.com/pulumi/pulumi/sdk/v3 v3.35.0
)

replace github.com/F5Networks/terraform-provider-bigip => github.com/pulumi/terraform-provider-bigip v1.1.1-0.20220526124114-f5b0215f2250
