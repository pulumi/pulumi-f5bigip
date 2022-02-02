// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.SnatPool` Collections of SNAT translation addresses
//
// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-snatpool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewSnatPool(ctx, "snatpoolSanjose", &ltm.SnatPoolArgs{
// 			Members: pulumi.StringArray{
// 				pulumi.String("191.1.1.1"),
// 				pulumi.String("194.2.2.2"),
// 			},
// 			Name: pulumi.String("/Common/snatpool_sanjose"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SnatPool struct {
	pulumi.CustomResourceState

	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Name of the snatpool
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSnatPool registers a new resource with the given unique name, arguments, and options.
func NewSnatPool(ctx *pulumi.Context,
	name string, args *SnatPoolArgs, opts ...pulumi.ResourceOption) (*SnatPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource SnatPool
	err := ctx.RegisterResource("f5bigip:ltm/snatPool:SnatPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnatPool gets an existing SnatPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnatPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnatPoolState, opts ...pulumi.ResourceOption) (*SnatPool, error) {
	var resource SnatPool
	err := ctx.ReadResource("f5bigip:ltm/snatPool:SnatPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnatPool resources.
type snatPoolState struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members []string `pulumi:"members"`
	// Name of the snatpool
	Name *string `pulumi:"name"`
}

type SnatPoolState struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members pulumi.StringArrayInput
	// Name of the snatpool
	Name pulumi.StringPtrInput
}

func (SnatPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*snatPoolState)(nil)).Elem()
}

type snatPoolArgs struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members []string `pulumi:"members"`
	// Name of the snatpool
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a SnatPool resource.
type SnatPoolArgs struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members pulumi.StringArrayInput
	// Name of the snatpool
	Name pulumi.StringInput
}

func (SnatPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snatPoolArgs)(nil)).Elem()
}

type SnatPoolInput interface {
	pulumi.Input

	ToSnatPoolOutput() SnatPoolOutput
	ToSnatPoolOutputWithContext(ctx context.Context) SnatPoolOutput
}

func (*SnatPool) ElementType() reflect.Type {
	return reflect.TypeOf((**SnatPool)(nil)).Elem()
}

func (i *SnatPool) ToSnatPoolOutput() SnatPoolOutput {
	return i.ToSnatPoolOutputWithContext(context.Background())
}

func (i *SnatPool) ToSnatPoolOutputWithContext(ctx context.Context) SnatPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatPoolOutput)
}

// SnatPoolArrayInput is an input type that accepts SnatPoolArray and SnatPoolArrayOutput values.
// You can construct a concrete instance of `SnatPoolArrayInput` via:
//
//          SnatPoolArray{ SnatPoolArgs{...} }
type SnatPoolArrayInput interface {
	pulumi.Input

	ToSnatPoolArrayOutput() SnatPoolArrayOutput
	ToSnatPoolArrayOutputWithContext(context.Context) SnatPoolArrayOutput
}

type SnatPoolArray []SnatPoolInput

func (SnatPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnatPool)(nil)).Elem()
}

func (i SnatPoolArray) ToSnatPoolArrayOutput() SnatPoolArrayOutput {
	return i.ToSnatPoolArrayOutputWithContext(context.Background())
}

func (i SnatPoolArray) ToSnatPoolArrayOutputWithContext(ctx context.Context) SnatPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatPoolArrayOutput)
}

// SnatPoolMapInput is an input type that accepts SnatPoolMap and SnatPoolMapOutput values.
// You can construct a concrete instance of `SnatPoolMapInput` via:
//
//          SnatPoolMap{ "key": SnatPoolArgs{...} }
type SnatPoolMapInput interface {
	pulumi.Input

	ToSnatPoolMapOutput() SnatPoolMapOutput
	ToSnatPoolMapOutputWithContext(context.Context) SnatPoolMapOutput
}

type SnatPoolMap map[string]SnatPoolInput

func (SnatPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnatPool)(nil)).Elem()
}

func (i SnatPoolMap) ToSnatPoolMapOutput() SnatPoolMapOutput {
	return i.ToSnatPoolMapOutputWithContext(context.Background())
}

func (i SnatPoolMap) ToSnatPoolMapOutputWithContext(ctx context.Context) SnatPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatPoolMapOutput)
}

type SnatPoolOutput struct{ *pulumi.OutputState }

func (SnatPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnatPool)(nil)).Elem()
}

func (o SnatPoolOutput) ToSnatPoolOutput() SnatPoolOutput {
	return o
}

func (o SnatPoolOutput) ToSnatPoolOutputWithContext(ctx context.Context) SnatPoolOutput {
	return o
}

type SnatPoolArrayOutput struct{ *pulumi.OutputState }

func (SnatPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnatPool)(nil)).Elem()
}

func (o SnatPoolArrayOutput) ToSnatPoolArrayOutput() SnatPoolArrayOutput {
	return o
}

func (o SnatPoolArrayOutput) ToSnatPoolArrayOutputWithContext(ctx context.Context) SnatPoolArrayOutput {
	return o
}

func (o SnatPoolArrayOutput) Index(i pulumi.IntInput) SnatPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnatPool {
		return vs[0].([]*SnatPool)[vs[1].(int)]
	}).(SnatPoolOutput)
}

type SnatPoolMapOutput struct{ *pulumi.OutputState }

func (SnatPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnatPool)(nil)).Elem()
}

func (o SnatPoolMapOutput) ToSnatPoolMapOutput() SnatPoolMapOutput {
	return o
}

func (o SnatPoolMapOutput) ToSnatPoolMapOutputWithContext(ctx context.Context) SnatPoolMapOutput {
	return o
}

func (o SnatPoolMapOutput) MapIndex(k pulumi.StringInput) SnatPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnatPool {
		return vs[0].(map[string]*SnatPool)[vs[1].(string)]
	}).(SnatPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnatPoolInput)(nil)).Elem(), &SnatPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatPoolArrayInput)(nil)).Elem(), SnatPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatPoolMapInput)(nil)).Elem(), SnatPoolMap{})
	pulumi.RegisterOutputType(SnatPoolOutput{})
	pulumi.RegisterOutputType(SnatPoolArrayOutput{})
	pulumi.RegisterOutputType(SnatPoolMapOutput{})
}
