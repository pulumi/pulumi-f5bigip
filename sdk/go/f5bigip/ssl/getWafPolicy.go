// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`WafPolicy`) to get the details of exist WAF policy BIG-IP.
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
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssl.GetWafPolicy(ctx, &ssl.GetWafPolicyArgs{
//				PolicyId: "xxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetWafPolicy(ctx *pulumi.Context, args *GetWafPolicyArgs, opts ...pulumi.InvokeOption) (*GetWafPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWafPolicyResult
	err := ctx.Invoke("f5bigip:ssl/getWafPolicy:getWafPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWafPolicy.
type GetWafPolicyArgs struct {
	// ID of the WAF policy deployed in the BIG-IP.
	PolicyId string `pulumi:"policyId"`
	// Exported WAF policy JSON
	PolicyJson *string `pulumi:"policyJson"`
}

// A collection of values returned by getWafPolicy.
type GetWafPolicyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	PolicyId string `pulumi:"policyId"`
	// Exported WAF policy JSON
	PolicyJson string `pulumi:"policyJson"`
}

func GetWafPolicyOutput(ctx *pulumi.Context, args GetWafPolicyOutputArgs, opts ...pulumi.InvokeOption) GetWafPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWafPolicyResultOutput, error) {
			args := v.(GetWafPolicyArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetWafPolicyResult
			secret, err := ctx.InvokePackageRaw("f5bigip:ssl/getWafPolicy:getWafPolicy", args, &rv, "", opts...)
			if err != nil {
				return GetWafPolicyResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetWafPolicyResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetWafPolicyResultOutput), nil
			}
			return output, nil
		}).(GetWafPolicyResultOutput)
}

// A collection of arguments for invoking getWafPolicy.
type GetWafPolicyOutputArgs struct {
	// ID of the WAF policy deployed in the BIG-IP.
	PolicyId pulumi.StringInput `pulumi:"policyId"`
	// Exported WAF policy JSON
	PolicyJson pulumi.StringPtrInput `pulumi:"policyJson"`
}

func (GetWafPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getWafPolicy.
type GetWafPolicyResultOutput struct{ *pulumi.OutputState }

func (GetWafPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafPolicyResult)(nil)).Elem()
}

func (o GetWafPolicyResultOutput) ToGetWafPolicyResultOutput() GetWafPolicyResultOutput {
	return o
}

func (o GetWafPolicyResultOutput) ToGetWafPolicyResultOutputWithContext(ctx context.Context) GetWafPolicyResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetWafPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWafPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

// Exported WAF policy JSON
func (o GetWafPolicyResultOutput) PolicyJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPolicyResult) string { return v.PolicyJson }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWafPolicyResultOutput{})
}
