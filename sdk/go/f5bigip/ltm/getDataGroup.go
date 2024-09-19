// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`ltm.DataGroup`) to get the data group details available on BIG-IP
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
//			_, err := ltm.LookupDataGroup(ctx, &ltm.LookupDataGroupArgs{
//				Name:      "test-dg",
//				Partition: "Common",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDataGroup(ctx *pulumi.Context, args *LookupDataGroupArgs, opts ...pulumi.InvokeOption) (*LookupDataGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataGroupResult
	err := ctx.Invoke("f5bigip:ltm/getDataGroup:getDataGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataGroup.
type LookupDataGroupArgs struct {
	// Name of the datagroup
	Name string `pulumi:"name"`
	// partition of the datagroup
	Partition string `pulumi:"partition"`
	// Specifies record of type (string/ip/integer)
	Records []GetDataGroupRecord `pulumi:"records"`
	// The Data Group type (string, ip, integer)"
	Type *string `pulumi:"type"`
}

// A collection of values returned by getDataGroup.
type LookupDataGroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	Name      string `pulumi:"name"`
	Partition string `pulumi:"partition"`
	// Specifies record of type (string/ip/integer)
	Records []GetDataGroupRecord `pulumi:"records"`
	// The Data Group type (string, ip, integer)"
	Type string `pulumi:"type"`
}

func LookupDataGroupOutput(ctx *pulumi.Context, args LookupDataGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDataGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataGroupResultOutput, error) {
			args := v.(LookupDataGroupArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupDataGroupResult
			secret, err := ctx.InvokePackageRaw("f5bigip:ltm/getDataGroup:getDataGroup", args, &rv, "", opts...)
			if err != nil {
				return LookupDataGroupResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupDataGroupResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupDataGroupResultOutput), nil
			}
			return output, nil
		}).(LookupDataGroupResultOutput)
}

// A collection of arguments for invoking getDataGroup.
type LookupDataGroupOutputArgs struct {
	// Name of the datagroup
	Name pulumi.StringInput `pulumi:"name"`
	// partition of the datagroup
	Partition pulumi.StringInput `pulumi:"partition"`
	// Specifies record of type (string/ip/integer)
	Records GetDataGroupRecordArrayInput `pulumi:"records"`
	// The Data Group type (string, ip, integer)"
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (LookupDataGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataGroupArgs)(nil)).Elem()
}

// A collection of values returned by getDataGroup.
type LookupDataGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDataGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataGroupResult)(nil)).Elem()
}

func (o LookupDataGroupResultOutput) ToLookupDataGroupResultOutput() LookupDataGroupResultOutput {
	return o
}

func (o LookupDataGroupResultOutput) ToLookupDataGroupResultOutputWithContext(ctx context.Context) LookupDataGroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDataGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataGroupResultOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataGroupResult) string { return v.Partition }).(pulumi.StringOutput)
}

// Specifies record of type (string/ip/integer)
func (o LookupDataGroupResultOutput) Records() GetDataGroupRecordArrayOutput {
	return o.ApplyT(func(v LookupDataGroupResult) []GetDataGroupRecord { return v.Records }).(GetDataGroupRecordArrayOutput)
}

// The Data Group type (string, ip, integer)"
func (o LookupDataGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataGroupResultOutput{})
}
