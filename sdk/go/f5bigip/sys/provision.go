// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `sys.Provision` Manage BIG-IP module provisioning. This resource will only provision at the standard levels of Dedicated, Nominal, and Minimum.
type Provision struct {
	pulumi.CustomResourceState

	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	CpuRatio pulumi.IntPtrOutput `pulumi:"cpuRatio"`
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	DiskRatio pulumi.IntPtrOutput `pulumi:"diskRatio"`
	FullPath  pulumi.StringOutput `pulumi:"fullPath"`
	// Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
	// default is `nominal`
	// possible options:
	// * nominal
	// * minimum
	// * none
	// * dedicated
	Level pulumi.StringPtrOutput `pulumi:"level"`
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	MemoryRatio pulumi.IntPtrOutput `pulumi:"memoryRatio"`
	// Name of module to provision in BIG-IP.
	// possible options:
	// * afm
	// * am
	// * apm
	// * cgnat
	// * asm
	// * avr
	// * dos
	// * fps
	// * gtm
	// * ilx
	// * lc
	// * ltm
	// * pem
	// * sslo
	// * swg
	// * urldb
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	CpuRatio *int `pulumi:"cpuRatio"`
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	DiskRatio *int    `pulumi:"diskRatio"`
	FullPath  *string `pulumi:"fullPath"`
	// Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
	// default is `nominal`
	// possible options:
	// * nominal
	// * minimum
	// * none
	// * dedicated
	Level *string `pulumi:"level"`
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	MemoryRatio *int `pulumi:"memoryRatio"`
	// Name of module to provision in BIG-IP.
	// possible options:
	// * afm
	// * am
	// * apm
	// * cgnat
	// * asm
	// * avr
	// * dos
	// * fps
	// * gtm
	// * ilx
	// * lc
	// * ltm
	// * pem
	// * sslo
	// * swg
	// * urldb
	Name *string `pulumi:"name"`
}

type ProvisionState struct {
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	CpuRatio pulumi.IntPtrInput
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	DiskRatio pulumi.IntPtrInput
	FullPath  pulumi.StringPtrInput
	// Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
	// default is `nominal`
	// possible options:
	// * nominal
	// * minimum
	// * none
	// * dedicated
	Level pulumi.StringPtrInput
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	MemoryRatio pulumi.IntPtrInput
	// Name of module to provision in BIG-IP.
	// possible options:
	// * afm
	// * am
	// * apm
	// * cgnat
	// * asm
	// * avr
	// * dos
	// * fps
	// * gtm
	// * ilx
	// * lc
	// * ltm
	// * pem
	// * sslo
	// * swg
	// * urldb
	Name pulumi.StringPtrInput
}

func (ProvisionState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionState)(nil)).Elem()
}

type provisionArgs struct {
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	CpuRatio *int `pulumi:"cpuRatio"`
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	DiskRatio *int    `pulumi:"diskRatio"`
	FullPath  *string `pulumi:"fullPath"`
	// Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
	// default is `nominal`
	// possible options:
	// * nominal
	// * minimum
	// * none
	// * dedicated
	Level *string `pulumi:"level"`
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	MemoryRatio *int `pulumi:"memoryRatio"`
	// Name of module to provision in BIG-IP.
	// possible options:
	// * afm
	// * am
	// * apm
	// * cgnat
	// * asm
	// * avr
	// * dos
	// * fps
	// * gtm
	// * ilx
	// * lc
	// * ltm
	// * pem
	// * sslo
	// * swg
	// * urldb
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a Provision resource.
type ProvisionArgs struct {
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	CpuRatio pulumi.IntPtrInput
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	DiskRatio pulumi.IntPtrInput
	FullPath  pulumi.StringPtrInput
	// Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
	// default is `nominal`
	// possible options:
	// * nominal
	// * minimum
	// * none
	// * dedicated
	Level pulumi.StringPtrInput
	// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
	MemoryRatio pulumi.IntPtrInput
	// Name of module to provision in BIG-IP.
	// possible options:
	// * afm
	// * am
	// * apm
	// * cgnat
	// * asm
	// * avr
	// * dos
	// * fps
	// * gtm
	// * ilx
	// * lc
	// * ltm
	// * pem
	// * sslo
	// * swg
	// * urldb
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
	return reflect.TypeOf((**Provision)(nil)).Elem()
}

func (i *Provision) ToProvisionOutput() ProvisionOutput {
	return i.ToProvisionOutputWithContext(context.Background())
}

func (i *Provision) ToProvisionOutputWithContext(ctx context.Context) ProvisionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionOutput)
}

func (i *Provision) ToOutput(ctx context.Context) pulumix.Output[*Provision] {
	return pulumix.Output[*Provision]{
		OutputState: i.ToProvisionOutputWithContext(ctx).OutputState,
	}
}

// ProvisionArrayInput is an input type that accepts ProvisionArray and ProvisionArrayOutput values.
// You can construct a concrete instance of `ProvisionArrayInput` via:
//
//	ProvisionArray{ ProvisionArgs{...} }
type ProvisionArrayInput interface {
	pulumi.Input

	ToProvisionArrayOutput() ProvisionArrayOutput
	ToProvisionArrayOutputWithContext(context.Context) ProvisionArrayOutput
}

type ProvisionArray []ProvisionInput

func (ProvisionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Provision)(nil)).Elem()
}

func (i ProvisionArray) ToProvisionArrayOutput() ProvisionArrayOutput {
	return i.ToProvisionArrayOutputWithContext(context.Background())
}

func (i ProvisionArray) ToProvisionArrayOutputWithContext(ctx context.Context) ProvisionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionArrayOutput)
}

func (i ProvisionArray) ToOutput(ctx context.Context) pulumix.Output[[]*Provision] {
	return pulumix.Output[[]*Provision]{
		OutputState: i.ToProvisionArrayOutputWithContext(ctx).OutputState,
	}
}

// ProvisionMapInput is an input type that accepts ProvisionMap and ProvisionMapOutput values.
// You can construct a concrete instance of `ProvisionMapInput` via:
//
//	ProvisionMap{ "key": ProvisionArgs{...} }
type ProvisionMapInput interface {
	pulumi.Input

	ToProvisionMapOutput() ProvisionMapOutput
	ToProvisionMapOutputWithContext(context.Context) ProvisionMapOutput
}

type ProvisionMap map[string]ProvisionInput

func (ProvisionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Provision)(nil)).Elem()
}

func (i ProvisionMap) ToProvisionMapOutput() ProvisionMapOutput {
	return i.ToProvisionMapOutputWithContext(context.Background())
}

func (i ProvisionMap) ToProvisionMapOutputWithContext(ctx context.Context) ProvisionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisionMapOutput)
}

func (i ProvisionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Provision] {
	return pulumix.Output[map[string]*Provision]{
		OutputState: i.ToProvisionMapOutputWithContext(ctx).OutputState,
	}
}

type ProvisionOutput struct{ *pulumi.OutputState }

func (ProvisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provision)(nil)).Elem()
}

func (o ProvisionOutput) ToProvisionOutput() ProvisionOutput {
	return o
}

func (o ProvisionOutput) ToProvisionOutputWithContext(ctx context.Context) ProvisionOutput {
	return o
}

func (o ProvisionOutput) ToOutput(ctx context.Context) pulumix.Output[*Provision] {
	return pulumix.Output[*Provision]{
		OutputState: o.OutputState,
	}
}

// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
func (o ProvisionOutput) CpuRatio() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Provision) pulumi.IntPtrOutput { return v.CpuRatio }).(pulumi.IntPtrOutput)
}

// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
func (o ProvisionOutput) DiskRatio() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Provision) pulumi.IntPtrOutput { return v.DiskRatio }).(pulumi.IntPtrOutput)
}

func (o ProvisionOutput) FullPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Provision) pulumi.StringOutput { return v.FullPath }).(pulumi.StringOutput)
}

// Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
// default is `nominal`
// possible options:
// * nominal
// * minimum
// * none
// * dedicated
func (o ProvisionOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provision) pulumi.StringPtrOutput { return v.Level }).(pulumi.StringPtrOutput)
}

// Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
func (o ProvisionOutput) MemoryRatio() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Provision) pulumi.IntPtrOutput { return v.MemoryRatio }).(pulumi.IntPtrOutput)
}

// Name of module to provision in BIG-IP.
// possible options:
// * afm
// * am
// * apm
// * cgnat
// * asm
// * avr
// * dos
// * fps
// * gtm
// * ilx
// * lc
// * ltm
// * pem
// * sslo
// * swg
// * urldb
func (o ProvisionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Provision) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type ProvisionArrayOutput struct{ *pulumi.OutputState }

func (ProvisionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Provision)(nil)).Elem()
}

func (o ProvisionArrayOutput) ToProvisionArrayOutput() ProvisionArrayOutput {
	return o
}

func (o ProvisionArrayOutput) ToProvisionArrayOutputWithContext(ctx context.Context) ProvisionArrayOutput {
	return o
}

func (o ProvisionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Provision] {
	return pulumix.Output[[]*Provision]{
		OutputState: o.OutputState,
	}
}

func (o ProvisionArrayOutput) Index(i pulumi.IntInput) ProvisionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Provision {
		return vs[0].([]*Provision)[vs[1].(int)]
	}).(ProvisionOutput)
}

type ProvisionMapOutput struct{ *pulumi.OutputState }

func (ProvisionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Provision)(nil)).Elem()
}

func (o ProvisionMapOutput) ToProvisionMapOutput() ProvisionMapOutput {
	return o
}

func (o ProvisionMapOutput) ToProvisionMapOutputWithContext(ctx context.Context) ProvisionMapOutput {
	return o
}

func (o ProvisionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Provision] {
	return pulumix.Output[map[string]*Provision]{
		OutputState: o.OutputState,
	}
}

func (o ProvisionMapOutput) MapIndex(k pulumi.StringInput) ProvisionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Provision {
		return vs[0].(map[string]*Provision)[vs[1].(string)]
	}).(ProvisionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionInput)(nil)).Elem(), &Provision{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionArrayInput)(nil)).Elem(), ProvisionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisionMapInput)(nil)).Elem(), ProvisionMap{})
	pulumi.RegisterOutputType(ProvisionOutput{})
	pulumi.RegisterOutputType(ProvisionArrayOutput{})
	pulumi.RegisterOutputType(ProvisionMapOutput{})
}
