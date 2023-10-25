// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source (`ltm.IRule`) to get the ltm irule details available on BIG-IP
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := ltm.GetIrule(ctx, &ltm.GetIruleArgs{
//				Name:      "terraform_irule",
//				Partition: "Common",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("bigipIrule", test.Irule)
//			return nil
//		})
//	}
//
// ```
func GetIrule(ctx *pulumi.Context, args *GetIruleArgs, opts ...pulumi.InvokeOption) (*GetIruleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIruleResult
	err := ctx.Invoke("f5bigip:ltm/getIrule:getIrule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIrule.
type GetIruleArgs struct {
	// Irule configured on bigip
	Irule *string `pulumi:"irule"`
	// Name of the irule
	Name string `pulumi:"name"`
	// partition of the ltm irule
	Partition string `pulumi:"partition"`
}

// A collection of values returned by getIrule.
type GetIruleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Irule configured on bigip
	Irule *string `pulumi:"irule"`
	// Name of irule configured on bigip with full path
	Name string `pulumi:"name"`
	// Bigip partition in which rule is configured
	Partition string `pulumi:"partition"`
}

func GetIruleOutput(ctx *pulumi.Context, args GetIruleOutputArgs, opts ...pulumi.InvokeOption) GetIruleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIruleResult, error) {
			args := v.(GetIruleArgs)
			r, err := GetIrule(ctx, &args, opts...)
			var s GetIruleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIruleResultOutput)
}

// A collection of arguments for invoking getIrule.
type GetIruleOutputArgs struct {
	// Irule configured on bigip
	Irule pulumi.StringPtrInput `pulumi:"irule"`
	// Name of the irule
	Name pulumi.StringInput `pulumi:"name"`
	// partition of the ltm irule
	Partition pulumi.StringInput `pulumi:"partition"`
}

func (GetIruleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIruleArgs)(nil)).Elem()
}

// A collection of values returned by getIrule.
type GetIruleResultOutput struct{ *pulumi.OutputState }

func (GetIruleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIruleResult)(nil)).Elem()
}

func (o GetIruleResultOutput) ToGetIruleResultOutput() GetIruleResultOutput {
	return o
}

func (o GetIruleResultOutput) ToGetIruleResultOutputWithContext(ctx context.Context) GetIruleResultOutput {
	return o
}

func (o GetIruleResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetIruleResult] {
	return pulumix.Output[GetIruleResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetIruleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIruleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Irule configured on bigip
func (o GetIruleResultOutput) Irule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIruleResult) *string { return v.Irule }).(pulumi.StringPtrOutput)
}

// Name of irule configured on bigip with full path
func (o GetIruleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIruleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Bigip partition in which rule is configured
func (o GetIruleResultOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v GetIruleResult) string { return v.Partition }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIruleResultOutput{})
}
