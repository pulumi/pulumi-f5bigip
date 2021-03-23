module github.com/pulumi/pulumi-f5bigip/provider/v2

go 1.16

require (
	github.com/F5Networks/terraform-provider-bigip v1.3.3
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.22.1
	github.com/pulumi/pulumi/sdk/v2 v2.22.1-0.20210310211618-1f16423ede4c
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/F5Networks/terraform-provider-bigip => github.com/pulumi/terraform-provider-bigip v1.1.1-0.20210218140212-259a15df7533
)
