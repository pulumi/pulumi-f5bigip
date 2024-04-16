// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`ssl.getVWanConfig`) to get the vWAN site config from Azure VWAN Site
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ssl"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssl.GetVWanConfig(ctx, &ssl.GetVWanConfigArgs{
//				AzureVwanResourcegroup: "azurevwan-bigip-rg-9c8d",
//				AzureVwanName:          "azurevwan-bigip-vwan-9c8d",
//				AzureVwanVpnsite:       "azurevwan-bigip-vsite-9c8d",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Pre-required Environment Settings:
//
// * `AZURE_CLIENT_ID` - (Required) Set this environment variable with the Azure app client ID to use.
//
// * `AZURE_CLIENT_SECRET` - (Required) Set this environment variable with the Azure app secret to use.
//
// * `AZURE_SUBSCRIPTION_ID` - (Required) Set this environment variable with the Azure subscription ID to use.
//
// * `AZURE_TENANT_ID` - (Required) Set this environment variable with the Tenant ID to which to authenticate.
//
// * `STORAGE_ACCOUNT_NAME` - (Required) Set this environment variable with the storage account for download config.
//
// * `STORAGE_ACCOUNT_KEY` - (Required) Specifies the storage account key to authenticate,set this Environment variable with account key value.
func GetVWanConfig(ctx *pulumi.Context, args *GetVWanConfigArgs, opts ...pulumi.InvokeOption) (*GetVWanConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVWanConfigResult
	err := ctx.Invoke("f5bigip:ssl/getVWanConfig:getVWanConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVWanConfig.
type GetVWanConfigArgs struct {
	// Name of the Azure vWAN Name
	AzureVwanName string `pulumi:"azureVwanName"`
	// Name of the Azure vWAN resource group
	AzureVwanResourcegroup string `pulumi:"azureVwanResourcegroup"`
	// Name of the Azure vWAN VPN site from which configuration to be download
	AzureVwanVpnsite string `pulumi:"azureVwanVpnsite"`
}

// A collection of values returned by getVWanConfig.
type GetVWanConfigResult struct {
	AzureVwanName          string `pulumi:"azureVwanName"`
	AzureVwanResourcegroup string `pulumi:"azureVwanResourcegroup"`
	AzureVwanVpnsite       string `pulumi:"azureVwanVpnsite"`
	// (type `string`) provides IP address of BIGIP G/W for IPSec Endpoint.
	BigipGwIp string `pulumi:"bigipGwIp"`
	// (type `string`) Provides IP Address space used on vWAN Hub.
	HubAddressSpace string `pulumi:"hubAddressSpace"`
	// (type `list`) Provides Subnets connected to vWAN Hub.
	HubConnectedSubnets []string `pulumi:"hubConnectedSubnets"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (type `string`) provides pre-shared-key used for IPSec Tunnel creation.
	PresharedKey string `pulumi:"presharedKey"`
	// (type `list`) Provides vWAN Gateway Address for IPSec End point
	VwanGwAddresses []string `pulumi:"vwanGwAddresses"`
}

func GetVWanConfigOutput(ctx *pulumi.Context, args GetVWanConfigOutputArgs, opts ...pulumi.InvokeOption) GetVWanConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVWanConfigResult, error) {
			args := v.(GetVWanConfigArgs)
			r, err := GetVWanConfig(ctx, &args, opts...)
			var s GetVWanConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVWanConfigResultOutput)
}

// A collection of arguments for invoking getVWanConfig.
type GetVWanConfigOutputArgs struct {
	// Name of the Azure vWAN Name
	AzureVwanName pulumi.StringInput `pulumi:"azureVwanName"`
	// Name of the Azure vWAN resource group
	AzureVwanResourcegroup pulumi.StringInput `pulumi:"azureVwanResourcegroup"`
	// Name of the Azure vWAN VPN site from which configuration to be download
	AzureVwanVpnsite pulumi.StringInput `pulumi:"azureVwanVpnsite"`
}

func (GetVWanConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVWanConfigArgs)(nil)).Elem()
}

// A collection of values returned by getVWanConfig.
type GetVWanConfigResultOutput struct{ *pulumi.OutputState }

func (GetVWanConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVWanConfigResult)(nil)).Elem()
}

func (o GetVWanConfigResultOutput) ToGetVWanConfigResultOutput() GetVWanConfigResultOutput {
	return o
}

func (o GetVWanConfigResultOutput) ToGetVWanConfigResultOutputWithContext(ctx context.Context) GetVWanConfigResultOutput {
	return o
}

func (o GetVWanConfigResultOutput) AzureVwanName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVWanConfigResult) string { return v.AzureVwanName }).(pulumi.StringOutput)
}

func (o GetVWanConfigResultOutput) AzureVwanResourcegroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetVWanConfigResult) string { return v.AzureVwanResourcegroup }).(pulumi.StringOutput)
}

func (o GetVWanConfigResultOutput) AzureVwanVpnsite() pulumi.StringOutput {
	return o.ApplyT(func(v GetVWanConfigResult) string { return v.AzureVwanVpnsite }).(pulumi.StringOutput)
}

// (type `string`) provides IP address of BIGIP G/W for IPSec Endpoint.
func (o GetVWanConfigResultOutput) BigipGwIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetVWanConfigResult) string { return v.BigipGwIp }).(pulumi.StringOutput)
}

// (type `string`) Provides IP Address space used on vWAN Hub.
func (o GetVWanConfigResultOutput) HubAddressSpace() pulumi.StringOutput {
	return o.ApplyT(func(v GetVWanConfigResult) string { return v.HubAddressSpace }).(pulumi.StringOutput)
}

// (type `list`) Provides Subnets connected to vWAN Hub.
func (o GetVWanConfigResultOutput) HubConnectedSubnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVWanConfigResult) []string { return v.HubConnectedSubnets }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVWanConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVWanConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// (type `string`) provides pre-shared-key used for IPSec Tunnel creation.
func (o GetVWanConfigResultOutput) PresharedKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetVWanConfigResult) string { return v.PresharedKey }).(pulumi.StringOutput)
}

// (type `list`) Provides vWAN Gateway Address for IPSec End point
func (o GetVWanConfigResultOutput) VwanGwAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVWanConfigResult) []string { return v.VwanGwAddresses }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVWanConfigResultOutput{})
}
