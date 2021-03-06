// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `sys.Provision` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/sys"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sys.NewProvision(ctx, "test_provision", &sys.ProvisionArgs{
// 			CpuRatio:    pulumi.Int(0),
// 			DiskRatio:   pulumi.Int(0),
// 			FullPath:    pulumi.String("asm"),
// 			Level:       pulumi.String("none"),
// 			MemoryRatio: pulumi.Int(0),
// 			Name:        pulumi.String("TEST_ASM_PROVISION_NAME"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Provision struct {
	pulumi.CustomResourceState

	// cpu Ratio
	CpuRatio pulumi.IntPtrOutput `pulumi:"cpuRatio"`
	// disk Ratio
	DiskRatio pulumi.IntPtrOutput `pulumi:"diskRatio"`
	// path
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// what level nominal or dedicated
	Level pulumi.StringPtrOutput `pulumi:"level"`
	// memory Ratio
	MemoryRatio pulumi.IntPtrOutput `pulumi:"memoryRatio"`
	// Name of the module to be provisioned
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewProvision registers a new resource with the given unique name, arguments, and options.
func NewProvision(ctx *pulumi.Context,
	name string, args *ProvisionArgs, opts ...pulumi.ResourceOption) (*Provision, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource Provision
	err := ctx.RegisterResource("f5bigip:sys/provision:Provision", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProvision gets an existing Provision resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvision(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisionState, opts ...pulumi.ResourceOption) (*Provision, error) {
	var resource Provision
	err := ctx.ReadResource("f5bigip:sys/provision:Provision", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Provision resources.
type provisionState struct {
	// cpu Ratio
	CpuRatio *int `pulumi:"cpuRatio"`
	// disk Ratio
	DiskRatio *int `pulumi:"diskRatio"`
	// path
	FullPath *string `pulumi:"fullPath"`
	// what level nominal or dedicated
	Level *string `pulumi:"level"`
	// memory Ratio
	MemoryRatio *int `pulumi:"memoryRatio"`
	// Name of the module to be provisioned
	Name *string `pulumi:"name"`
}

type ProvisionState struct {
	// cpu Ratio
	CpuRatio pulumi.IntPtrInput
	// disk Ratio
	DiskRatio pulumi.IntPtrInput
	// path
	FullPath pulumi.StringPtrInput
	// what level nominal or dedicated
	Level pulumi.StringPtrInput
	// memory Ratio
	MemoryRatio pulumi.IntPtrInput
	// Name of the module to be provisioned
	Name pulumi.StringPtrInput
}

func (ProvisionState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionState)(nil)).Elem()
}

type provisionArgs struct {
	// cpu Ratio
	CpuRatio *int `pulumi:"cpuRatio"`
	// disk Ratio
	DiskRatio *int `pulumi:"diskRatio"`
	// path
	FullPath *string `pulumi:"fullPath"`
	// what level nominal or dedicated
	Level *string `pulumi:"level"`
	// memory Ratio
	MemoryRatio *int `pulumi:"memoryRatio"`
	// Name of the module to be provisioned
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a Provision resource.
type ProvisionArgs struct {
	// cpu Ratio
	CpuRatio pulumi.IntPtrInput
	// disk Ratio
	DiskRatio pulumi.IntPtrInput
	// path
	FullPath pulumi.StringPtrInput
	// what level nominal or dedicated
	Level pulumi.StringPtrInput
	// memory Ratio
	MemoryRatio pulumi.IntPtrInput
	// Name of the module to be provisioned
	Name pulumi.StringInput
}

func (ProvisionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionArgs)(nil)).Elem()
}

type ProvisionInput interface {
	pulumi.Input

	ToProvisionOutput() ProvisionOutput
	ToProvisionOutputWithContext(ctx context.Context) ProvisionOutput
}

func (*Provision) ElementType() reflect.Type {
	return reflect.TypeOf((*Provision)(nil))
}

func (i *Provision) ToProvisionOutput() ProvisionOutput {
	return i.ToProvisionOutputWithContext(context.Background())
}

func (i *Provision) ToProvisionOutputWithContext(ctx context.Context) ProvisionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionOutput)
}

func (i *Provision) ToProvisionPtrOutput() ProvisionPtrOutput {
	return i.ToProvisionPtrOutputWithContext(context.Background())
}

func (i *Provision) ToProvisionPtrOutputWithContext(ctx context.Context) ProvisionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionPtrOutput)
}

type ProvisionPtrInput interface {
	pulumi.Input

	ToProvisionPtrOutput() ProvisionPtrOutput
	ToProvisionPtrOutputWithContext(ctx context.Context) ProvisionPtrOutput
}

type provisionPtrType ProvisionArgs

func (*provisionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Provision)(nil))
}

func (i *provisionPtrType) ToProvisionPtrOutput() ProvisionPtrOutput {
	return i.ToProvisionPtrOutputWithContext(context.Background())
}

func (i *provisionPtrType) ToProvisionPtrOutputWithContext(ctx context.Context) ProvisionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionPtrOutput)
}

// ProvisionArrayInput is an input type that accepts ProvisionArray and ProvisionArrayOutput values.
// You can construct a concrete instance of `ProvisionArrayInput` via:
//
//          ProvisionArray{ ProvisionArgs{...} }
type ProvisionArrayInput interface {
	pulumi.Input

	ToProvisionArrayOutput() ProvisionArrayOutput
	ToProvisionArrayOutputWithContext(context.Context) ProvisionArrayOutput
}

type ProvisionArray []ProvisionInput

func (ProvisionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Provision)(nil))
}

func (i ProvisionArray) ToProvisionArrayOutput() ProvisionArrayOutput {
	return i.ToProvisionArrayOutputWithContext(context.Background())
}

func (i ProvisionArray) ToProvisionArrayOutputWithContext(ctx context.Context) ProvisionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionArrayOutput)
}

// ProvisionMapInput is an input type that accepts ProvisionMap and ProvisionMapOutput values.
// You can construct a concrete instance of `ProvisionMapInput` via:
//
//          ProvisionMap{ "key": ProvisionArgs{...} }
type ProvisionMapInput interface {
	pulumi.Input

	ToProvisionMapOutput() ProvisionMapOutput
	ToProvisionMapOutputWithContext(context.Context) ProvisionMapOutput
}

type ProvisionMap map[string]ProvisionInput

func (ProvisionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Provision)(nil))
}

func (i ProvisionMap) ToProvisionMapOutput() ProvisionMapOutput {
	return i.ToProvisionMapOutputWithContext(context.Background())
}

func (i ProvisionMap) ToProvisionMapOutputWithContext(ctx context.Context) ProvisionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionMapOutput)
}

type ProvisionOutput struct {
	*pulumi.OutputState
}

func (ProvisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provision)(nil))
}

func (o ProvisionOutput) ToProvisionOutput() ProvisionOutput {
	return o
}

func (o ProvisionOutput) ToProvisionOutputWithContext(ctx context.Context) ProvisionOutput {
	return o
}

func (o ProvisionOutput) ToProvisionPtrOutput() ProvisionPtrOutput {
	return o.ToProvisionPtrOutputWithContext(context.Background())
}

func (o ProvisionOutput) ToProvisionPtrOutputWithContext(ctx context.Context) ProvisionPtrOutput {
	return o.ApplyT(func(v Provision) *Provision {
		return &v
	}).(ProvisionPtrOutput)
}

type ProvisionPtrOutput struct {
	*pulumi.OutputState
}

func (ProvisionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provision)(nil))
}

func (o ProvisionPtrOutput) ToProvisionPtrOutput() ProvisionPtrOutput {
	return o
}

func (o ProvisionPtrOutput) ToProvisionPtrOutputWithContext(ctx context.Context) ProvisionPtrOutput {
	return o
}

type ProvisionArrayOutput struct{ *pulumi.OutputState }

func (ProvisionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Provision)(nil))
}

func (o ProvisionArrayOutput) ToProvisionArrayOutput() ProvisionArrayOutput {
	return o
}

func (o ProvisionArrayOutput) ToProvisionArrayOutputWithContext(ctx context.Context) ProvisionArrayOutput {
	return o
}

func (o ProvisionArrayOutput) Index(i pulumi.IntInput) ProvisionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Provision {
		return vs[0].([]Provision)[vs[1].(int)]
	}).(ProvisionOutput)
}

type ProvisionMapOutput struct{ *pulumi.OutputState }

func (ProvisionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Provision)(nil))
}

func (o ProvisionMapOutput) ToProvisionMapOutput() ProvisionMapOutput {
	return o
}

func (o ProvisionMapOutput) ToProvisionMapOutputWithContext(ctx context.Context) ProvisionMapOutput {
	return o
}

func (o ProvisionMapOutput) MapIndex(k pulumi.StringInput) ProvisionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Provision {
		return vs[0].(map[string]Provision)[vs[1].(string)]
	}).(ProvisionOutput)
}

func init() {
	pulumi.RegisterOutputType(ProvisionOutput{})
	pulumi.RegisterOutputType(ProvisionPtrOutput{})
	pulumi.RegisterOutputType(ProvisionArrayOutput{})
	pulumi.RegisterOutputType(ProvisionMapOutput{})
}
