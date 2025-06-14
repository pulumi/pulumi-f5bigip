// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`ssl.getWafPbSuggestions`) to create JSON for WAF URL to later use with an existing WAF policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ssl"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// _, err := ssl.GetWafEntityUrl(ctx, &ssl.GetWafEntityUrlArgs{
// Name: "/foobar",
// Description: pulumi.StringRef("this is a test"),
// Type: pulumi.StringRef("explicit"),
// Protocol: pulumi.StringRef("HTTP"),
// PerformStaging: pulumi.BoolRef(true),
// SignatureOverridesDisables: interface{}{
// 12345678,
// 87654321,
// },
// MethodOverrides: []ssl.GetWafEntityUrlMethodOverride{
// {
// Allow: false,
// Method: "BCOPY",
// },
// {
// Allow: true,
// Method: "BDELETE",
// },
// },
// CrossOriginRequestsEnforcements: []ssl.GetWafEntityUrlCrossOriginRequestsEnforcement{
// {
// IncludeSubdomains: pulumi.BoolRef(true),
// OriginName: "app1.com",
// OriginPort: "80",
// OriginProtocol: "http",
// },
// {
// IncludeSubdomains: pulumi.BoolRef(true),
// OriginName: "app2.com",
// OriginPort: "443",
// OriginProtocol: "http",
// },
// },
// }, nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func GetWafEntityUrl(ctx *pulumi.Context, args *GetWafEntityUrlArgs, opts ...pulumi.InvokeOption) (*GetWafEntityUrlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWafEntityUrlResult
	err := ctx.Invoke("f5bigip:ssl/getWafEntityUrl:getWafEntityUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWafEntityUrl.
type GetWafEntityUrlArgs struct {
	// A list of options that enables your web-application to share data with a website hosted on a
	// different domain.
	CrossOriginRequestsEnforcements []GetWafEntityUrlCrossOriginRequestsEnforcement `pulumi:"crossOriginRequestsEnforcements"`
	// A description of the URL.
	Description *string `pulumi:"description"`
	// Select a Method for the URL to create an API endpoint. Default is : *.
	Method *string `pulumi:"method"`
	// A list of methods that are allowed or disallowed for a specific URL.
	MethodOverrides []GetWafEntityUrlMethodOverride `pulumi:"methodOverrides"`
	// WAF entity URL name.
	Name string `pulumi:"name"`
	// If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
	PerformStaging *bool `pulumi:"performStaging"`
	// Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
	Protocol *string `pulumi:"protocol"`
	// List of Attack Signature Ids which are disabled for this particular URL.
	SignatureOverridesDisables []int `pulumi:"signatureOverridesDisables"`
	// Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getWafEntityUrl.
type GetWafEntityUrlResult struct {
	CrossOriginRequestsEnforcements []GetWafEntityUrlCrossOriginRequestsEnforcement `pulumi:"crossOriginRequestsEnforcements"`
	Description                     *string                                         `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Json string representing created WAF entity URL declaration in JSON format
	Json                       string                          `pulumi:"json"`
	Method                     *string                         `pulumi:"method"`
	MethodOverrides            []GetWafEntityUrlMethodOverride `pulumi:"methodOverrides"`
	Name                       string                          `pulumi:"name"`
	PerformStaging             *bool                           `pulumi:"performStaging"`
	Protocol                   *string                         `pulumi:"protocol"`
	SignatureOverridesDisables []int                           `pulumi:"signatureOverridesDisables"`
	Type                       *string                         `pulumi:"type"`
}

func GetWafEntityUrlOutput(ctx *pulumi.Context, args GetWafEntityUrlOutputArgs, opts ...pulumi.InvokeOption) GetWafEntityUrlResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetWafEntityUrlResultOutput, error) {
			args := v.(GetWafEntityUrlArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("f5bigip:ssl/getWafEntityUrl:getWafEntityUrl", args, GetWafEntityUrlResultOutput{}, options).(GetWafEntityUrlResultOutput), nil
		}).(GetWafEntityUrlResultOutput)
}

// A collection of arguments for invoking getWafEntityUrl.
type GetWafEntityUrlOutputArgs struct {
	// A list of options that enables your web-application to share data with a website hosted on a
	// different domain.
	CrossOriginRequestsEnforcements GetWafEntityUrlCrossOriginRequestsEnforcementArrayInput `pulumi:"crossOriginRequestsEnforcements"`
	// A description of the URL.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Select a Method for the URL to create an API endpoint. Default is : *.
	Method pulumi.StringPtrInput `pulumi:"method"`
	// A list of methods that are allowed or disallowed for a specific URL.
	MethodOverrides GetWafEntityUrlMethodOverrideArrayInput `pulumi:"methodOverrides"`
	// WAF entity URL name.
	Name pulumi.StringInput `pulumi:"name"`
	// If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
	PerformStaging pulumi.BoolPtrInput `pulumi:"performStaging"`
	// Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// List of Attack Signature Ids which are disabled for this particular URL.
	SignatureOverridesDisables pulumi.IntArrayInput `pulumi:"signatureOverridesDisables"`
	// Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetWafEntityUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafEntityUrlArgs)(nil)).Elem()
}

// A collection of values returned by getWafEntityUrl.
type GetWafEntityUrlResultOutput struct{ *pulumi.OutputState }

func (GetWafEntityUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafEntityUrlResult)(nil)).Elem()
}

func (o GetWafEntityUrlResultOutput) ToGetWafEntityUrlResultOutput() GetWafEntityUrlResultOutput {
	return o
}

func (o GetWafEntityUrlResultOutput) ToGetWafEntityUrlResultOutputWithContext(ctx context.Context) GetWafEntityUrlResultOutput {
	return o
}

func (o GetWafEntityUrlResultOutput) CrossOriginRequestsEnforcements() GetWafEntityUrlCrossOriginRequestsEnforcementArrayOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) []GetWafEntityUrlCrossOriginRequestsEnforcement {
		return v.CrossOriginRequestsEnforcements
	}).(GetWafEntityUrlCrossOriginRequestsEnforcementArrayOutput)
}

func (o GetWafEntityUrlResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWafEntityUrlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) string { return v.Id }).(pulumi.StringOutput)
}

// Json string representing created WAF entity URL declaration in JSON format
func (o GetWafEntityUrlResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetWafEntityUrlResultOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o GetWafEntityUrlResultOutput) MethodOverrides() GetWafEntityUrlMethodOverrideArrayOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) []GetWafEntityUrlMethodOverride { return v.MethodOverrides }).(GetWafEntityUrlMethodOverrideArrayOutput)
}

func (o GetWafEntityUrlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetWafEntityUrlResultOutput) PerformStaging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) *bool { return v.PerformStaging }).(pulumi.BoolPtrOutput)
}

func (o GetWafEntityUrlResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o GetWafEntityUrlResultOutput) SignatureOverridesDisables() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) []int { return v.SignatureOverridesDisables }).(pulumi.IntArrayOutput)
}

func (o GetWafEntityUrlResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWafEntityUrlResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWafEntityUrlResultOutput{})
}
