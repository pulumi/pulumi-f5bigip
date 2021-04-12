// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.VirtualAddress` Configures Virtual Server
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewVirtualAddress(ctx, "vsVa", &ltm.VirtualAddressArgs{
// 			AdvertizeRoute: pulumi.String("true"),
// 			Name:           pulumi.String("/Common/vs_va"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VirtualAddress struct {
	pulumi.CustomResourceState

	// Enabled dynamic routing of the address
	AdvertizeRoute pulumi.StringPtrOutput `pulumi:"advertizeRoute"`
	// Enable or disable ARP for the virtual address
	Arp pulumi.BoolPtrOutput `pulumi:"arp"`
	// Automatically delete the virtual address with the virtual server
	AutoDelete pulumi.BoolPtrOutput `pulumi:"autoDelete"`
	// Max number of connections for virtual address
	ConnLimit pulumi.IntPtrOutput `pulumi:"connLimit"`
	// Enable or disable the virtual address
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Enable/Disable ICMP response to the virtual address
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
	// Enabled dynamic routing of the address
	AdvertizeRoute *string `pulumi:"advertizeRoute"`
	// Enable or disable ARP for the virtual address
	Arp *bool `pulumi:"arp"`
	// Automatically delete the virtual address with the virtual server
	AutoDelete *bool `pulumi:"autoDelete"`
	// Max number of connections for virtual address
	ConnLimit *int `pulumi:"connLimit"`
	// Enable or disable the virtual address
	Enabled *bool `pulumi:"enabled"`
	// Enable/Disable ICMP response to the virtual address
	IcmpEcho *string `pulumi:"icmpEcho"`
	// Name of the virtual address
	Name *string `pulumi:"name"`
	// Specify the partition and traffic group
	TrafficGroup *string `pulumi:"trafficGroup"`
}

type VirtualAddressState struct {
	// Enabled dynamic routing of the address
	AdvertizeRoute pulumi.StringPtrInput
	// Enable or disable ARP for the virtual address
	Arp pulumi.BoolPtrInput
	// Automatically delete the virtual address with the virtual server
	AutoDelete pulumi.BoolPtrInput
	// Max number of connections for virtual address
	ConnLimit pulumi.IntPtrInput
	// Enable or disable the virtual address
	Enabled pulumi.BoolPtrInput
	// Enable/Disable ICMP response to the virtual address
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
	// Enabled dynamic routing of the address
	AdvertizeRoute *string `pulumi:"advertizeRoute"`
	// Enable or disable ARP for the virtual address
	Arp *bool `pulumi:"arp"`
	// Automatically delete the virtual address with the virtual server
	AutoDelete *bool `pulumi:"autoDelete"`
	// Max number of connections for virtual address
	ConnLimit *int `pulumi:"connLimit"`
	// Enable or disable the virtual address
	Enabled *bool `pulumi:"enabled"`
	// Enable/Disable ICMP response to the virtual address
	IcmpEcho *string `pulumi:"icmpEcho"`
	// Name of the virtual address
	Name string `pulumi:"name"`
	// Specify the partition and traffic group
	TrafficGroup *string `pulumi:"trafficGroup"`
}

// The set of arguments for constructing a VirtualAddress resource.
type VirtualAddressArgs struct {
	// Enabled dynamic routing of the address
	AdvertizeRoute pulumi.StringPtrInput
	// Enable or disable ARP for the virtual address
	Arp pulumi.BoolPtrInput
	// Automatically delete the virtual address with the virtual server
	AutoDelete pulumi.BoolPtrInput
	// Max number of connections for virtual address
	ConnLimit pulumi.IntPtrInput
	// Enable or disable the virtual address
	Enabled pulumi.BoolPtrInput
	// Enable/Disable ICMP response to the virtual address
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
	return reflect.TypeOf((*VirtualAddress)(nil))
}

func (i *VirtualAddress) ToVirtualAddressOutput() VirtualAddressOutput {
	return i.ToVirtualAddressOutputWithContext(context.Background())
}

func (i *VirtualAddress) ToVirtualAddressOutputWithContext(ctx context.Context) VirtualAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualAddressOutput)
}

func (i *VirtualAddress) ToVirtualAddressPtrOutput() VirtualAddressPtrOutput {
	return i.ToVirtualAddressPtrOutputWithContext(context.Background())
}

func (i *VirtualAddress) ToVirtualAddressPtrOutputWithContext(ctx context.Context) VirtualAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualAddressPtrOutput)
}

type VirtualAddressPtrInput interface {
	pulumi.Input

	ToVirtualAddressPtrOutput() VirtualAddressPtrOutput
	ToVirtualAddressPtrOutputWithContext(ctx context.Context) VirtualAddressPtrOutput
}

type virtualAddressPtrType VirtualAddressArgs

func (*virtualAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualAddress)(nil))
}

func (i *virtualAddressPtrType) ToVirtualAddressPtrOutput() VirtualAddressPtrOutput {
	return i.ToVirtualAddressPtrOutputWithContext(context.Background())
}

func (i *virtualAddressPtrType) ToVirtualAddressPtrOutputWithContext(ctx context.Context) VirtualAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualAddressPtrOutput)
}

// VirtualAddressArrayInput is an input type that accepts VirtualAddressArray and VirtualAddressArrayOutput values.
// You can construct a concrete instance of `VirtualAddressArrayInput` via:
//
//          VirtualAddressArray{ VirtualAddressArgs{...} }
type VirtualAddressArrayInput interface {
	pulumi.Input

	ToVirtualAddressArrayOutput() VirtualAddressArrayOutput
	ToVirtualAddressArrayOutputWithContext(context.Context) VirtualAddressArrayOutput
}

type VirtualAddressArray []VirtualAddressInput

func (VirtualAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VirtualAddress)(nil))
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
//          VirtualAddressMap{ "key": VirtualAddressArgs{...} }
type VirtualAddressMapInput interface {
	pulumi.Input

	ToVirtualAddressMapOutput() VirtualAddressMapOutput
	ToVirtualAddressMapOutputWithContext(context.Context) VirtualAddressMapOutput
}

type VirtualAddressMap map[string]VirtualAddressInput

func (VirtualAddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VirtualAddress)(nil))
}

func (i VirtualAddressMap) ToVirtualAddressMapOutput() VirtualAddressMapOutput {
	return i.ToVirtualAddressMapOutputWithContext(context.Background())
}

func (i VirtualAddressMap) ToVirtualAddressMapOutputWithContext(ctx context.Context) VirtualAddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualAddressMapOutput)
}

type VirtualAddressOutput struct {
	*pulumi.OutputState
}

func (VirtualAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualAddress)(nil))
}

func (o VirtualAddressOutput) ToVirtualAddressOutput() VirtualAddressOutput {
	return o
}

func (o VirtualAddressOutput) ToVirtualAddressOutputWithContext(ctx context.Context) VirtualAddressOutput {
	return o
}

func (o VirtualAddressOutput) ToVirtualAddressPtrOutput() VirtualAddressPtrOutput {
	return o.ToVirtualAddressPtrOutputWithContext(context.Background())
}

func (o VirtualAddressOutput) ToVirtualAddressPtrOutputWithContext(ctx context.Context) VirtualAddressPtrOutput {
	return o.ApplyT(func(v VirtualAddress) *VirtualAddress {
		return &v
	}).(VirtualAddressPtrOutput)
}

type VirtualAddressPtrOutput struct {
	*pulumi.OutputState
}

func (VirtualAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualAddress)(nil))
}

func (o VirtualAddressPtrOutput) ToVirtualAddressPtrOutput() VirtualAddressPtrOutput {
	return o
}

func (o VirtualAddressPtrOutput) ToVirtualAddressPtrOutputWithContext(ctx context.Context) VirtualAddressPtrOutput {
	return o
}

type VirtualAddressArrayOutput struct{ *pulumi.OutputState }

func (VirtualAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualAddress)(nil))
}

func (o VirtualAddressArrayOutput) ToVirtualAddressArrayOutput() VirtualAddressArrayOutput {
	return o
}

func (o VirtualAddressArrayOutput) ToVirtualAddressArrayOutputWithContext(ctx context.Context) VirtualAddressArrayOutput {
	return o
}

func (o VirtualAddressArrayOutput) Index(i pulumi.IntInput) VirtualAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualAddress {
		return vs[0].([]VirtualAddress)[vs[1].(int)]
	}).(VirtualAddressOutput)
}

type VirtualAddressMapOutput struct{ *pulumi.OutputState }

func (VirtualAddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VirtualAddress)(nil))
}

func (o VirtualAddressMapOutput) ToVirtualAddressMapOutput() VirtualAddressMapOutput {
	return o
}

func (o VirtualAddressMapOutput) ToVirtualAddressMapOutputWithContext(ctx context.Context) VirtualAddressMapOutput {
	return o
}

func (o VirtualAddressMapOutput) MapIndex(k pulumi.StringInput) VirtualAddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VirtualAddress {
		return vs[0].(map[string]VirtualAddress)[vs[1].(string)]
	}).(VirtualAddressOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualAddressOutput{})
	pulumi.RegisterOutputType(VirtualAddressPtrOutput{})
	pulumi.RegisterOutputType(VirtualAddressArrayOutput{})
	pulumi.RegisterOutputType(VirtualAddressMapOutput{})
}
