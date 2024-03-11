// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`ssl.getWafPbSuggestions`) to export PB suggestions from an existing WAF policy.
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
//			_, err := ssl.GetWafPbSuggestions(ctx, &ssl.GetWafPbSuggestionsArgs{
//				MinimumLearningScore: 20,
//				Partition:            "Common",
//				PolicyName:           "protect_me_policy",
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
func GetWafPbSuggestions(ctx *pulumi.Context, args *GetWafPbSuggestionsArgs, opts ...pulumi.InvokeOption) (*GetWafPbSuggestionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWafPbSuggestionsResult
	err := ctx.Invoke("f5bigip:ssl/getWafPbSuggestions:getWafPbSuggestions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWafPbSuggestions.
type GetWafPbSuggestionsArgs struct {
	// The minimum learning score for suggestions.
	MinimumLearningScore int `pulumi:"minimumLearningScore"`
	// Partition on which WAF policy is located.
	Partition string `pulumi:"partition"`
	// System generated id of the WAF policy
	PolicyId *string `pulumi:"policyId"`
	// WAF policy name from which PB suggestions should be exported.
	PolicyName string `pulumi:"policyName"`
}

// A collection of values returned by getWafPbSuggestions.
type GetWafPbSuggestionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Json string representing exported PB suggestions ready to be used in WAF policy declaration
	Json                 string `pulumi:"json"`
	MinimumLearningScore int    `pulumi:"minimumLearningScore"`
	Partition            string `pulumi:"partition"`
	// System generated id of the WAF policy
	PolicyId   string `pulumi:"policyId"`
	PolicyName string `pulumi:"policyName"`
}

func GetWafPbSuggestionsOutput(ctx *pulumi.Context, args GetWafPbSuggestionsOutputArgs, opts ...pulumi.InvokeOption) GetWafPbSuggestionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWafPbSuggestionsResult, error) {
			args := v.(GetWafPbSuggestionsArgs)
			r, err := GetWafPbSuggestions(ctx, &args, opts...)
			var s GetWafPbSuggestionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWafPbSuggestionsResultOutput)
}

// A collection of arguments for invoking getWafPbSuggestions.
type GetWafPbSuggestionsOutputArgs struct {
	// The minimum learning score for suggestions.
	MinimumLearningScore pulumi.IntInput `pulumi:"minimumLearningScore"`
	// Partition on which WAF policy is located.
	Partition pulumi.StringInput `pulumi:"partition"`
	// System generated id of the WAF policy
	PolicyId pulumi.StringPtrInput `pulumi:"policyId"`
	// WAF policy name from which PB suggestions should be exported.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
}

func (GetWafPbSuggestionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafPbSuggestionsArgs)(nil)).Elem()
}

// A collection of values returned by getWafPbSuggestions.
type GetWafPbSuggestionsResultOutput struct{ *pulumi.OutputState }

func (GetWafPbSuggestionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWafPbSuggestionsResult)(nil)).Elem()
}

func (o GetWafPbSuggestionsResultOutput) ToGetWafPbSuggestionsResultOutput() GetWafPbSuggestionsResultOutput {
	return o
}

func (o GetWafPbSuggestionsResultOutput) ToGetWafPbSuggestionsResultOutputWithContext(ctx context.Context) GetWafPbSuggestionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetWafPbSuggestionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPbSuggestionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Json string representing exported PB suggestions ready to be used in WAF policy declaration
func (o GetWafPbSuggestionsResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPbSuggestionsResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetWafPbSuggestionsResultOutput) MinimumLearningScore() pulumi.IntOutput {
	return o.ApplyT(func(v GetWafPbSuggestionsResult) int { return v.MinimumLearningScore }).(pulumi.IntOutput)
}

func (o GetWafPbSuggestionsResultOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPbSuggestionsResult) string { return v.Partition }).(pulumi.StringOutput)
}

// System generated id of the WAF policy
func (o GetWafPbSuggestionsResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPbSuggestionsResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o GetWafPbSuggestionsResultOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v GetWafPbSuggestionsResult) string { return v.PolicyName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWafPbSuggestionsResultOutput{})
}
