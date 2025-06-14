// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type VlanInterface struct {
	// Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
	Tagged *bool `pulumi:"tagged"`
	// Physical or virtual port used for traffic
	Vlanport *string `pulumi:"vlanport"`
}

// VlanInterfaceInput is an input type that accepts VlanInterfaceArgs and VlanInterfaceOutput values.
// You can construct a concrete instance of `VlanInterfaceInput` via:
//
//	VlanInterfaceArgs{...}
type VlanInterfaceInput interface {
	pulumi.Input

	ToVlanInterfaceOutput() VlanInterfaceOutput
	ToVlanInterfaceOutputWithContext(context.Context) VlanInterfaceOutput
}

type VlanInterfaceArgs struct {
	// Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
	Tagged pulumi.BoolPtrInput `pulumi:"tagged"`
	// Physical or virtual port used for traffic
	Vlanport pulumi.StringPtrInput `pulumi:"vlanport"`
}

func (VlanInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VlanInterface)(nil)).Elem()
}

func (i VlanInterfaceArgs) ToVlanInterfaceOutput() VlanInterfaceOutput {
	return i.ToVlanInterfaceOutputWithContext(context.Background())
}

func (i VlanInterfaceArgs) ToVlanInterfaceOutputWithContext(ctx context.Context) VlanInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanInterfaceOutput)
}

// VlanInterfaceArrayInput is an input type that accepts VlanInterfaceArray and VlanInterfaceArrayOutput values.
// You can construct a concrete instance of `VlanInterfaceArrayInput` via:
//
//	VlanInterfaceArray{ VlanInterfaceArgs{...} }
type VlanInterfaceArrayInput interface {
	pulumi.Input

	ToVlanInterfaceArrayOutput() VlanInterfaceArrayOutput
	ToVlanInterfaceArrayOutputWithContext(context.Context) VlanInterfaceArrayOutput
}

type VlanInterfaceArray []VlanInterfaceInput

func (VlanInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VlanInterface)(nil)).Elem()
}

func (i VlanInterfaceArray) ToVlanInterfaceArrayOutput() VlanInterfaceArrayOutput {
	return i.ToVlanInterfaceArrayOutputWithContext(context.Background())
}

func (i VlanInterfaceArray) ToVlanInterfaceArrayOutputWithContext(ctx context.Context) VlanInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VlanInterfaceArrayOutput)
}

type VlanInterfaceOutput struct{ *pulumi.OutputState }

func (VlanInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VlanInterface)(nil)).Elem()
}

func (o VlanInterfaceOutput) ToVlanInterfaceOutput() VlanInterfaceOutput {
	return o
}

func (o VlanInterfaceOutput) ToVlanInterfaceOutputWithContext(ctx context.Context) VlanInterfaceOutput {
	return o
}

// Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
func (o VlanInterfaceOutput) Tagged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VlanInterface) *bool { return v.Tagged }).(pulumi.BoolPtrOutput)
}

// Physical or virtual port used for traffic
func (o VlanInterfaceOutput) Vlanport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VlanInterface) *string { return v.Vlanport }).(pulumi.StringPtrOutput)
}

type VlanInterfaceArrayOutput struct{ *pulumi.OutputState }

func (VlanInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VlanInterface)(nil)).Elem()
}

func (o VlanInterfaceArrayOutput) ToVlanInterfaceArrayOutput() VlanInterfaceArrayOutput {
	return o
}

func (o VlanInterfaceArrayOutput) ToVlanInterfaceArrayOutputWithContext(ctx context.Context) VlanInterfaceArrayOutput {
	return o
}

func (o VlanInterfaceArrayOutput) Index(i pulumi.IntInput) VlanInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VlanInterface {
		return vs[0].([]VlanInterface)[vs[1].(int)]
	}).(VlanInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VlanInterfaceInput)(nil)).Elem(), VlanInterfaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VlanInterfaceArrayInput)(nil)).Elem(), VlanInterfaceArray{})
	pulumi.RegisterOutputType(VlanInterfaceOutput{})
	pulumi.RegisterOutputType(VlanInterfaceArrayOutput{})
}
