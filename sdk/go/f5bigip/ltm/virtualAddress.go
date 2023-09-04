// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.VirtualAddress` Configures Virtual Server
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/virtual_server.
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
//			_, err := ltm.NewVirtualAddress(ctx, "vsVa", &ltm.VirtualAddressArgs{
//				AdvertizeRoute: pulumi.String("enabled"),
//				Name:           pulumi.String("/Common/xxxxx"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VirtualAddress struct {
	pulumi.CustomResourceState

	// Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
	AdvertizeRoute pulumi.StringPtrOutput `pulumi:"advertizeRoute"`
	// Enable or disable ARP for the virtual address
	Arp pulumi.BoolPtrOutput `pulumi:"arp"`
	// Automatically delete the virtual address with the virtual server
	AutoDelete pulumi.BoolPtrOutput `pulumi:"autoDelete"`
	// Max number of connections for virtual address
	ConnLimit pulumi.IntPtrOutput `pulumi:"connLimit"`
	// Enable or disable the virtual address
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
	IcmpEcho pulumi.StringPtrOutput `pulumi:"icmpEcho"`
	// Name of the virtual address
	Name pulumi.StringOutput `pulumi:"name"`
	// Specify the partition and traffic group
	TrafficGroup pulumi.StringPtrOutput `pulumi:"trafficGroup"`
}

// NewVirtualAddress registers a new resource with the given unique name, arguments, and options.
func NewVirtualAddress(ctx *pulumi.Context,
	name string, args *VirtualAddressArgs, opts ...pulumi.ResourceOption) (*VirtualAddress, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualAddress
	err := ctx.RegisterResource("f5bigip:ltm/virtualAddress:VirtualAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualAddress gets an existing VirtualAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualAddressState, opts ...pulumi.ResourceOption) (*VirtualAddress, error) {
	var resource VirtualAddress
	err := ctx.ReadResource("f5bigip:ltm/virtualAddress:VirtualAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualAddress resources.
type virtualAddressState struct {
	// Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
	AdvertizeRoute *string `pulumi:"advertizeRoute"`
	// Enable or disable ARP for the virtual address
	Arp *bool `pulumi:"arp"`
	// Automatically delete the virtual address with the virtual server
	AutoDelete *bool `pulumi:"autoDelete"`
	// Max number of connections for virtual address
	ConnLimit *int `pulumi:"connLimit"`
	// Enable or disable the virtual address
	Enabled *bool `pulumi:"enabled"`
	// Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
	IcmpEcho *string `pulumi:"icmpEcho"`
	// Name of the virtual address
	Name *string `pulumi:"name"`
	// Specify the partition and traffic group
	TrafficGroup *string `pulumi:"trafficGroup"`
}

type VirtualAddressState struct {
	// Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
	AdvertizeRoute pulumi.StringPtrInput
	// Enable or disable ARP for the virtual address
	Arp pulumi.BoolPtrInput
	// Automatically delete the virtual address with the virtual server
	AutoDelete pulumi.BoolPtrInput
	// Max number of connections for virtual address
	ConnLimit pulumi.IntPtrInput
	// Enable or disable the virtual address
	Enabled pulumi.BoolPtrInput
	// Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
	IcmpEcho pulumi.StringPtrInput
	// Name of the virtual address
	Name pulumi.StringPtrInput
	// Specify the partition and traffic group
	TrafficGroup pulumi.StringPtrInput
}

func (VirtualAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualAddressState)(nil)).Elem()
}

type virtualAddressArgs struct {
	// Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
	AdvertizeRoute *string `pulumi:"advertizeRoute"`
	// Enable or disable ARP for the virtual address
	Arp *bool `pulumi:"arp"`
	// Automatically delete the virtual address with the virtual server
	AutoDelete *bool `pulumi:"autoDelete"`
	// Max number of connections for virtual address
	ConnLimit *int `pulumi:"connLimit"`
	// Enable or disable the virtual address
	Enabled *bool `pulumi:"enabled"`
	// Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
	IcmpEcho *string `pulumi:"icmpEcho"`
	// Name of the virtual address
	Name string `pulumi:"name"`
	// Specify the partition and traffic group
	TrafficGroup *string `pulumi:"trafficGroup"`
}

// The set of arguments for constructing a VirtualAddress resource.
type VirtualAddressArgs struct {
	// Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
	AdvertizeRoute pulumi.StringPtrInput
	// Enable or disable ARP for the virtual address
	Arp pulumi.BoolPtrInput
	// Automatically delete the virtual address with the virtual server
	AutoDelete pulumi.BoolPtrInput
	// Max number of connections for virtual address
	ConnLimit pulumi.IntPtrInput
	// Enable or disable the virtual address
	Enabled pulumi.BoolPtrInput
	// Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
	IcmpEcho pulumi.StringPtrInput
	// Name of the virtual address
	Name pulumi.StringInput
	// Specify the partition and traffic group
	TrafficGroup pulumi.StringPtrInput
}

func (VirtualAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualAddressArgs)(nil)).Elem()
}

type VirtualAddressInput interface {
	pulumi.Input

	ToVirtualAddressOutput() VirtualAddressOutput
	ToVirtualAddressOutputWithContext(ctx context.Context) VirtualAddressOutput
}

func (*VirtualAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualAddress)(nil)).Elem()
}

func (i *VirtualAddress) ToVirtualAddressOutput() VirtualAddressOutput {
	return i.ToVirtualAddressOutputWithContext(context.Background())
}

func (i *VirtualAddress) ToVirtualAddressOutputWithContext(ctx context.Context) VirtualAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualAddressOutput)
}

// VirtualAddressArrayInput is an input type that accepts VirtualAddressArray and VirtualAddressArrayOutput values.
// You can construct a concrete instance of `VirtualAddressArrayInput` via:
//
//	VirtualAddressArray{ VirtualAddressArgs{...} }
type VirtualAddressArrayInput interface {
	pulumi.Input

	ToVirtualAddressArrayOutput() VirtualAddressArrayOutput
	ToVirtualAddressArrayOutputWithContext(context.Context) VirtualAddressArrayOutput
}

type VirtualAddressArray []VirtualAddressInput

func (VirtualAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualAddress)(nil)).Elem()
}

func (i VirtualAddressArray) ToVirtualAddressArrayOutput() VirtualAddressArrayOutput {
	return i.ToVirtualAddressArrayOutputWithContext(context.Background())
}

func (i VirtualAddressArray) ToVirtualAddressArrayOutputWithContext(ctx context.Context) VirtualAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualAddressArrayOutput)
}

// VirtualAddressMapInput is an input type that accepts VirtualAddressMap and VirtualAddressMapOutput values.
// You can construct a concrete instance of `VirtualAddressMapInput` via:
//
//	VirtualAddressMap{ "key": VirtualAddressArgs{...} }
type VirtualAddressMapInput interface {
	pulumi.Input

	ToVirtualAddressMapOutput() VirtualAddressMapOutput
	ToVirtualAddressMapOutputWithContext(context.Context) VirtualAddressMapOutput
}

type VirtualAddressMap map[string]VirtualAddressInput

func (VirtualAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualAddress)(nil)).Elem()
}

func (i VirtualAddressMap) ToVirtualAddressMapOutput() VirtualAddressMapOutput {
	return i.ToVirtualAddressMapOutputWithContext(context.Background())
}

func (i VirtualAddressMap) ToVirtualAddressMapOutputWithContext(ctx context.Context) VirtualAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualAddressMapOutput)
}

type VirtualAddressOutput struct{ *pulumi.OutputState }

func (VirtualAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualAddress)(nil)).Elem()
}

func (o VirtualAddressOutput) ToVirtualAddressOutput() VirtualAddressOutput {
	return o
}

func (o VirtualAddressOutput) ToVirtualAddressOutputWithContext(ctx context.Context) VirtualAddressOutput {
	return o
}

// Enabled dynamic routing of the address ( In versions prior to BIG-IP 13.0.0 HF1, you can configure the Route Advertisement option for a virtual address to be either Enabled or Disabled only. Beginning with BIG-IP 13.0.0 HF1, F5 added more settings for the Route Advertisement option. In addition, the Enabled setting is deprecated and replaced by the Selective setting. For more information, please look into KB article https://support.f5.com/csp/article/K85543242 )
func (o VirtualAddressOutput) AdvertizeRoute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.StringPtrOutput { return v.AdvertizeRoute }).(pulumi.StringPtrOutput)
}

// Enable or disable ARP for the virtual address
func (o VirtualAddressOutput) Arp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.BoolPtrOutput { return v.Arp }).(pulumi.BoolPtrOutput)
}

// Automatically delete the virtual address with the virtual server
func (o VirtualAddressOutput) AutoDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.BoolPtrOutput { return v.AutoDelete }).(pulumi.BoolPtrOutput)
}

// Max number of connections for virtual address
func (o VirtualAddressOutput) ConnLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.IntPtrOutput { return v.ConnLimit }).(pulumi.IntPtrOutput)
}

// Enable or disable the virtual address
func (o VirtualAddressOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies how the system sends responses to ICMP echo requests on a per-virtual address basis.
func (o VirtualAddressOutput) IcmpEcho() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.StringPtrOutput { return v.IcmpEcho }).(pulumi.StringPtrOutput)
}

// Name of the virtual address
func (o VirtualAddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specify the partition and traffic group
func (o VirtualAddressOutput) TrafficGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualAddress) pulumi.StringPtrOutput { return v.TrafficGroup }).(pulumi.StringPtrOutput)
}

type VirtualAddressArrayOutput struct{ *pulumi.OutputState }

func (VirtualAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualAddress)(nil)).Elem()
}

func (o VirtualAddressArrayOutput) ToVirtualAddressArrayOutput() VirtualAddressArrayOutput {
	return o
}

func (o VirtualAddressArrayOutput) ToVirtualAddressArrayOutputWithContext(ctx context.Context) VirtualAddressArrayOutput {
	return o
}

func (o VirtualAddressArrayOutput) Index(i pulumi.IntInput) VirtualAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualAddress {
		return vs[0].([]*VirtualAddress)[vs[1].(int)]
	}).(VirtualAddressOutput)
}

type VirtualAddressMapOutput struct{ *pulumi.OutputState }

func (VirtualAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualAddress)(nil)).Elem()
}

func (o VirtualAddressMapOutput) ToVirtualAddressMapOutput() VirtualAddressMapOutput {
	return o
}

func (o VirtualAddressMapOutput) ToVirtualAddressMapOutputWithContext(ctx context.Context) VirtualAddressMapOutput {
	return o
}

func (o VirtualAddressMapOutput) MapIndex(k pulumi.StringInput) VirtualAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualAddress {
		return vs[0].(map[string]*VirtualAddress)[vs[1].(string)]
	}).(VirtualAddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualAddressInput)(nil)).Elem(), &VirtualAddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualAddressArrayInput)(nil)).Elem(), VirtualAddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualAddressMapInput)(nil)).Elem(), VirtualAddressMap{})
	pulumi.RegisterOutputType(VirtualAddressOutput{})
	pulumi.RegisterOutputType(VirtualAddressArrayOutput{})
	pulumi.RegisterOutputType(VirtualAddressMapOutput{})
}
