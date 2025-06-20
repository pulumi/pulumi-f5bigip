// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `Partition` Manages F5 BIG-IP partitions
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := f5bigip.NewPartition(ctx, "test-partition", &f5bigip.PartitionArgs{
//				Name:          pulumi.String("test-partition"),
//				Description:   pulumi.String("created by terraform"),
//				RouteDomainId: pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Importing
//
// An existing cipher group can be imported into this resource by supplying the cipher rule full path name ex : `/partition/name`
//
// An example is below:
//
// ```sh
// $ terraform import bigip_partition.test_partition test_partition
//
// ```
type Partition struct {
	pulumi.CustomResourceState

	// Description of the partition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the partition.
	Name pulumi.StringOutput `pulumi:"name"`
	// Route domain id of the partition.
	RouteDomainId pulumi.IntPtrOutput `pulumi:"routeDomainId"`
}

// NewPartition registers a new resource with the given unique name, arguments, and options.
func NewPartition(ctx *pulumi.Context,
	name string, args *PartitionArgs, opts ...pulumi.ResourceOption) (*Partition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Partition
	err := ctx.RegisterResource("f5bigip:index/partition:Partition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartition gets an existing Partition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartitionState, opts ...pulumi.ResourceOption) (*Partition, error) {
	var resource Partition
	err := ctx.ReadResource("f5bigip:index/partition:Partition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Partition resources.
type partitionState struct {
	// Description of the partition.
	Description *string `pulumi:"description"`
	// Name of the partition.
	Name *string `pulumi:"name"`
	// Route domain id of the partition.
	RouteDomainId *int `pulumi:"routeDomainId"`
}

type PartitionState struct {
	// Description of the partition.
	Description pulumi.StringPtrInput
	// Name of the partition.
	Name pulumi.StringPtrInput
	// Route domain id of the partition.
	RouteDomainId pulumi.IntPtrInput
}

func (PartitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionState)(nil)).Elem()
}

type partitionArgs struct {
	// Description of the partition.
	Description *string `pulumi:"description"`
	// Name of the partition.
	Name string `pulumi:"name"`
	// Route domain id of the partition.
	RouteDomainId *int `pulumi:"routeDomainId"`
}

// The set of arguments for constructing a Partition resource.
type PartitionArgs struct {
	// Description of the partition.
	Description pulumi.StringPtrInput
	// Name of the partition.
	Name pulumi.StringInput
	// Route domain id of the partition.
	RouteDomainId pulumi.IntPtrInput
}

func (PartitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partitionArgs)(nil)).Elem()
}

type PartitionInput interface {
	pulumi.Input

	ToPartitionOutput() PartitionOutput
	ToPartitionOutputWithContext(ctx context.Context) PartitionOutput
}

func (*Partition) ElementType() reflect.Type {
	return reflect.TypeOf((**Partition)(nil)).Elem()
}

func (i *Partition) ToPartitionOutput() PartitionOutput {
	return i.ToPartitionOutputWithContext(context.Background())
}

func (i *Partition) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionOutput)
}

// PartitionArrayInput is an input type that accepts PartitionArray and PartitionArrayOutput values.
// You can construct a concrete instance of `PartitionArrayInput` via:
//
//	PartitionArray{ PartitionArgs{...} }
type PartitionArrayInput interface {
	pulumi.Input

	ToPartitionArrayOutput() PartitionArrayOutput
	ToPartitionArrayOutputWithContext(context.Context) PartitionArrayOutput
}

type PartitionArray []PartitionInput

func (PartitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Partition)(nil)).Elem()
}

func (i PartitionArray) ToPartitionArrayOutput() PartitionArrayOutput {
	return i.ToPartitionArrayOutputWithContext(context.Background())
}

func (i PartitionArray) ToPartitionArrayOutputWithContext(ctx context.Context) PartitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionArrayOutput)
}

// PartitionMapInput is an input type that accepts PartitionMap and PartitionMapOutput values.
// You can construct a concrete instance of `PartitionMapInput` via:
//
//	PartitionMap{ "key": PartitionArgs{...} }
type PartitionMapInput interface {
	pulumi.Input

	ToPartitionMapOutput() PartitionMapOutput
	ToPartitionMapOutputWithContext(context.Context) PartitionMapOutput
}

type PartitionMap map[string]PartitionInput

func (PartitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Partition)(nil)).Elem()
}

func (i PartitionMap) ToPartitionMapOutput() PartitionMapOutput {
	return i.ToPartitionMapOutputWithContext(context.Background())
}

func (i PartitionMap) ToPartitionMapOutputWithContext(ctx context.Context) PartitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionMapOutput)
}

type PartitionOutput struct{ *pulumi.OutputState }

func (PartitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Partition)(nil)).Elem()
}

func (o PartitionOutput) ToPartitionOutput() PartitionOutput {
	return o
}

func (o PartitionOutput) ToPartitionOutputWithContext(ctx context.Context) PartitionOutput {
	return o
}

// Description of the partition.
func (o PartitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Partition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the partition.
func (o PartitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Partition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Route domain id of the partition.
func (o PartitionOutput) RouteDomainId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Partition) pulumi.IntPtrOutput { return v.RouteDomainId }).(pulumi.IntPtrOutput)
}

type PartitionArrayOutput struct{ *pulumi.OutputState }

func (PartitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Partition)(nil)).Elem()
}

func (o PartitionArrayOutput) ToPartitionArrayOutput() PartitionArrayOutput {
	return o
}

func (o PartitionArrayOutput) ToPartitionArrayOutputWithContext(ctx context.Context) PartitionArrayOutput {
	return o
}

func (o PartitionArrayOutput) Index(i pulumi.IntInput) PartitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Partition {
		return vs[0].([]*Partition)[vs[1].(int)]
	}).(PartitionOutput)
}

type PartitionMapOutput struct{ *pulumi.OutputState }

func (PartitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Partition)(nil)).Elem()
}

func (o PartitionMapOutput) ToPartitionMapOutput() PartitionMapOutput {
	return o
}

func (o PartitionMapOutput) ToPartitionMapOutputWithContext(ctx context.Context) PartitionMapOutput {
	return o
}

func (o PartitionMapOutput) MapIndex(k pulumi.StringInput) PartitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Partition {
		return vs[0].(map[string]*Partition)[vs[1].(string)]
	}).(PartitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionInput)(nil)).Elem(), &Partition{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionArrayInput)(nil)).Elem(), PartitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PartitionMapInput)(nil)).Elem(), PartitionMap{})
	pulumi.RegisterOutputType(PartitionOutput{})
	pulumi.RegisterOutputType(PartitionArrayOutput{})
	pulumi.RegisterOutputType(PartitionMapOutput{})
}
