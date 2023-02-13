// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `NetTunnel` Manages a tunnel configuration
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
//			_, err := f5bigip.NewNetTunnel(ctx, "example1", &f5bigip.NetTunnelArgs{
//				LocalAddress: pulumi.String("192.16.81.240"),
//				Name:         pulumi.String("example1"),
//				Profile:      pulumi.String("/Common/dslite"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type NetTunnel struct {
	pulumi.CustomResourceState

	// The application service that the object belongs to
	AppService pulumi.StringPtrOutput `pulumi:"appService"`
	// Specifies whether auto lasthop is enabled or not
	AutoLastHop pulumi.StringPtrOutput `pulumi:"autoLastHop"`
	// User defined description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies an idle timeout for wildcard tunnels in seconds
	IdleTimeout pulumi.IntPtrOutput `pulumi:"idleTimeout"`
	// The key field may represent different values depending on the type of the tunnel
	Key pulumi.IntPtrOutput `pulumi:"key"`
	// Specifies a local IP address. This option is required
	LocalAddress pulumi.StringOutput `pulumi:"localAddress"`
	// Specifies how the tunnel carries traffic
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Specifies the maximum transmission unit (MTU) of the tunnel
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// Name of the tunnel
	Name pulumi.StringOutput `pulumi:"name"`
	// Displays the admin-partition within which this component resides
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// Specifies the profile that you want to associate with the tunnel
	Profile pulumi.StringOutput `pulumi:"profile"`
	// Specifies a remote IP address
	RemoteAddress pulumi.StringPtrOutput `pulumi:"remoteAddress"`
	// Specifies a secondary non-floating IP address when the local-address is set to a floating address
	SecondaryAddress pulumi.StringPtrOutput `pulumi:"secondaryAddress"`
	// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
	Tos pulumi.StringPtrOutput `pulumi:"tos"`
	// Specifies a traffic-group for use with the tunnel
	TrafficGroup pulumi.StringPtrOutput `pulumi:"trafficGroup"`
	// Enables or disables the tunnel to be transparent
	Transparent pulumi.StringPtrOutput `pulumi:"transparent"`
	// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
	UsePmtu pulumi.StringPtrOutput `pulumi:"usePmtu"`
}

// NewNetTunnel registers a new resource with the given unique name, arguments, and options.
func NewNetTunnel(ctx *pulumi.Context,
	name string, args *NetTunnelArgs, opts ...pulumi.ResourceOption) (*NetTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LocalAddress == nil {
		return nil, errors.New("invalid value for required argument 'LocalAddress'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Profile == nil {
		return nil, errors.New("invalid value for required argument 'Profile'")
	}
	var resource NetTunnel
	err := ctx.RegisterResource("f5bigip:index/netTunnel:NetTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetTunnel gets an existing NetTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetTunnelState, opts ...pulumi.ResourceOption) (*NetTunnel, error) {
	var resource NetTunnel
	err := ctx.ReadResource("f5bigip:index/netTunnel:NetTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetTunnel resources.
type netTunnelState struct {
	// The application service that the object belongs to
	AppService *string `pulumi:"appService"`
	// Specifies whether auto lasthop is enabled or not
	AutoLastHop *string `pulumi:"autoLastHop"`
	// User defined description
	Description *string `pulumi:"description"`
	// Specifies an idle timeout for wildcard tunnels in seconds
	IdleTimeout *int `pulumi:"idleTimeout"`
	// The key field may represent different values depending on the type of the tunnel
	Key *int `pulumi:"key"`
	// Specifies a local IP address. This option is required
	LocalAddress *string `pulumi:"localAddress"`
	// Specifies how the tunnel carries traffic
	Mode *string `pulumi:"mode"`
	// Specifies the maximum transmission unit (MTU) of the tunnel
	Mtu *int `pulumi:"mtu"`
	// Name of the tunnel
	Name *string `pulumi:"name"`
	// Displays the admin-partition within which this component resides
	Partition *string `pulumi:"partition"`
	// Specifies the profile that you want to associate with the tunnel
	Profile *string `pulumi:"profile"`
	// Specifies a remote IP address
	RemoteAddress *string `pulumi:"remoteAddress"`
	// Specifies a secondary non-floating IP address when the local-address is set to a floating address
	SecondaryAddress *string `pulumi:"secondaryAddress"`
	// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
	Tos *string `pulumi:"tos"`
	// Specifies a traffic-group for use with the tunnel
	TrafficGroup *string `pulumi:"trafficGroup"`
	// Enables or disables the tunnel to be transparent
	Transparent *string `pulumi:"transparent"`
	// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
	UsePmtu *string `pulumi:"usePmtu"`
}

type NetTunnelState struct {
	// The application service that the object belongs to
	AppService pulumi.StringPtrInput
	// Specifies whether auto lasthop is enabled or not
	AutoLastHop pulumi.StringPtrInput
	// User defined description
	Description pulumi.StringPtrInput
	// Specifies an idle timeout for wildcard tunnels in seconds
	IdleTimeout pulumi.IntPtrInput
	// The key field may represent different values depending on the type of the tunnel
	Key pulumi.IntPtrInput
	// Specifies a local IP address. This option is required
	LocalAddress pulumi.StringPtrInput
	// Specifies how the tunnel carries traffic
	Mode pulumi.StringPtrInput
	// Specifies the maximum transmission unit (MTU) of the tunnel
	Mtu pulumi.IntPtrInput
	// Name of the tunnel
	Name pulumi.StringPtrInput
	// Displays the admin-partition within which this component resides
	Partition pulumi.StringPtrInput
	// Specifies the profile that you want to associate with the tunnel
	Profile pulumi.StringPtrInput
	// Specifies a remote IP address
	RemoteAddress pulumi.StringPtrInput
	// Specifies a secondary non-floating IP address when the local-address is set to a floating address
	SecondaryAddress pulumi.StringPtrInput
	// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
	Tos pulumi.StringPtrInput
	// Specifies a traffic-group for use with the tunnel
	TrafficGroup pulumi.StringPtrInput
	// Enables or disables the tunnel to be transparent
	Transparent pulumi.StringPtrInput
	// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
	UsePmtu pulumi.StringPtrInput
}

func (NetTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*netTunnelState)(nil)).Elem()
}

type netTunnelArgs struct {
	// The application service that the object belongs to
	AppService *string `pulumi:"appService"`
	// Specifies whether auto lasthop is enabled or not
	AutoLastHop *string `pulumi:"autoLastHop"`
	// User defined description
	Description *string `pulumi:"description"`
	// Specifies an idle timeout for wildcard tunnels in seconds
	IdleTimeout *int `pulumi:"idleTimeout"`
	// The key field may represent different values depending on the type of the tunnel
	Key *int `pulumi:"key"`
	// Specifies a local IP address. This option is required
	LocalAddress string `pulumi:"localAddress"`
	// Specifies how the tunnel carries traffic
	Mode *string `pulumi:"mode"`
	// Specifies the maximum transmission unit (MTU) of the tunnel
	Mtu *int `pulumi:"mtu"`
	// Name of the tunnel
	Name string `pulumi:"name"`
	// Displays the admin-partition within which this component resides
	Partition *string `pulumi:"partition"`
	// Specifies the profile that you want to associate with the tunnel
	Profile string `pulumi:"profile"`
	// Specifies a remote IP address
	RemoteAddress *string `pulumi:"remoteAddress"`
	// Specifies a secondary non-floating IP address when the local-address is set to a floating address
	SecondaryAddress *string `pulumi:"secondaryAddress"`
	// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
	Tos *string `pulumi:"tos"`
	// Specifies a traffic-group for use with the tunnel
	TrafficGroup *string `pulumi:"trafficGroup"`
	// Enables or disables the tunnel to be transparent
	Transparent *string `pulumi:"transparent"`
	// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
	UsePmtu *string `pulumi:"usePmtu"`
}

// The set of arguments for constructing a NetTunnel resource.
type NetTunnelArgs struct {
	// The application service that the object belongs to
	AppService pulumi.StringPtrInput
	// Specifies whether auto lasthop is enabled or not
	AutoLastHop pulumi.StringPtrInput
	// User defined description
	Description pulumi.StringPtrInput
	// Specifies an idle timeout for wildcard tunnels in seconds
	IdleTimeout pulumi.IntPtrInput
	// The key field may represent different values depending on the type of the tunnel
	Key pulumi.IntPtrInput
	// Specifies a local IP address. This option is required
	LocalAddress pulumi.StringInput
	// Specifies how the tunnel carries traffic
	Mode pulumi.StringPtrInput
	// Specifies the maximum transmission unit (MTU) of the tunnel
	Mtu pulumi.IntPtrInput
	// Name of the tunnel
	Name pulumi.StringInput
	// Displays the admin-partition within which this component resides
	Partition pulumi.StringPtrInput
	// Specifies the profile that you want to associate with the tunnel
	Profile pulumi.StringInput
	// Specifies a remote IP address
	RemoteAddress pulumi.StringPtrInput
	// Specifies a secondary non-floating IP address when the local-address is set to a floating address
	SecondaryAddress pulumi.StringPtrInput
	// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
	Tos pulumi.StringPtrInput
	// Specifies a traffic-group for use with the tunnel
	TrafficGroup pulumi.StringPtrInput
	// Enables or disables the tunnel to be transparent
	Transparent pulumi.StringPtrInput
	// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
	UsePmtu pulumi.StringPtrInput
}

func (NetTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*netTunnelArgs)(nil)).Elem()
}

type NetTunnelInput interface {
	pulumi.Input

	ToNetTunnelOutput() NetTunnelOutput
	ToNetTunnelOutputWithContext(ctx context.Context) NetTunnelOutput
}

func (*NetTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**NetTunnel)(nil)).Elem()
}

func (i *NetTunnel) ToNetTunnelOutput() NetTunnelOutput {
	return i.ToNetTunnelOutputWithContext(context.Background())
}

func (i *NetTunnel) ToNetTunnelOutputWithContext(ctx context.Context) NetTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetTunnelOutput)
}

// NetTunnelArrayInput is an input type that accepts NetTunnelArray and NetTunnelArrayOutput values.
// You can construct a concrete instance of `NetTunnelArrayInput` via:
//
//	NetTunnelArray{ NetTunnelArgs{...} }
type NetTunnelArrayInput interface {
	pulumi.Input

	ToNetTunnelArrayOutput() NetTunnelArrayOutput
	ToNetTunnelArrayOutputWithContext(context.Context) NetTunnelArrayOutput
}

type NetTunnelArray []NetTunnelInput

func (NetTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetTunnel)(nil)).Elem()
}

func (i NetTunnelArray) ToNetTunnelArrayOutput() NetTunnelArrayOutput {
	return i.ToNetTunnelArrayOutputWithContext(context.Background())
}

func (i NetTunnelArray) ToNetTunnelArrayOutputWithContext(ctx context.Context) NetTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetTunnelArrayOutput)
}

// NetTunnelMapInput is an input type that accepts NetTunnelMap and NetTunnelMapOutput values.
// You can construct a concrete instance of `NetTunnelMapInput` via:
//
//	NetTunnelMap{ "key": NetTunnelArgs{...} }
type NetTunnelMapInput interface {
	pulumi.Input

	ToNetTunnelMapOutput() NetTunnelMapOutput
	ToNetTunnelMapOutputWithContext(context.Context) NetTunnelMapOutput
}

type NetTunnelMap map[string]NetTunnelInput

func (NetTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetTunnel)(nil)).Elem()
}

func (i NetTunnelMap) ToNetTunnelMapOutput() NetTunnelMapOutput {
	return i.ToNetTunnelMapOutputWithContext(context.Background())
}

func (i NetTunnelMap) ToNetTunnelMapOutputWithContext(ctx context.Context) NetTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetTunnelMapOutput)
}

type NetTunnelOutput struct{ *pulumi.OutputState }

func (NetTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetTunnel)(nil)).Elem()
}

func (o NetTunnelOutput) ToNetTunnelOutput() NetTunnelOutput {
	return o
}

func (o NetTunnelOutput) ToNetTunnelOutputWithContext(ctx context.Context) NetTunnelOutput {
	return o
}

// The application service that the object belongs to
func (o NetTunnelOutput) AppService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.AppService }).(pulumi.StringPtrOutput)
}

// Specifies whether auto lasthop is enabled or not
func (o NetTunnelOutput) AutoLastHop() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.AutoLastHop }).(pulumi.StringPtrOutput)
}

// User defined description
func (o NetTunnelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies an idle timeout for wildcard tunnels in seconds
func (o NetTunnelOutput) IdleTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.IntPtrOutput { return v.IdleTimeout }).(pulumi.IntPtrOutput)
}

// The key field may represent different values depending on the type of the tunnel
func (o NetTunnelOutput) Key() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.IntPtrOutput { return v.Key }).(pulumi.IntPtrOutput)
}

// Specifies a local IP address. This option is required
func (o NetTunnelOutput) LocalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringOutput { return v.LocalAddress }).(pulumi.StringOutput)
}

// Specifies how the tunnel carries traffic
func (o NetTunnelOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// Specifies the maximum transmission unit (MTU) of the tunnel
func (o NetTunnelOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.IntPtrOutput { return v.Mtu }).(pulumi.IntPtrOutput)
}

// Name of the tunnel
func (o NetTunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Displays the admin-partition within which this component resides
func (o NetTunnelOutput) Partition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.Partition }).(pulumi.StringPtrOutput)
}

// Specifies the profile that you want to associate with the tunnel
func (o NetTunnelOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringOutput { return v.Profile }).(pulumi.StringOutput)
}

// Specifies a remote IP address
func (o NetTunnelOutput) RemoteAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.RemoteAddress }).(pulumi.StringPtrOutput)
}

// Specifies a secondary non-floating IP address when the local-address is set to a floating address
func (o NetTunnelOutput) SecondaryAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.SecondaryAddress }).(pulumi.StringPtrOutput)
}

// Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
func (o NetTunnelOutput) Tos() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.Tos }).(pulumi.StringPtrOutput)
}

// Specifies a traffic-group for use with the tunnel
func (o NetTunnelOutput) TrafficGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.TrafficGroup }).(pulumi.StringPtrOutput)
}

// Enables or disables the tunnel to be transparent
func (o NetTunnelOutput) Transparent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.Transparent }).(pulumi.StringPtrOutput)
}

// Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
func (o NetTunnelOutput) UsePmtu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetTunnel) pulumi.StringPtrOutput { return v.UsePmtu }).(pulumi.StringPtrOutput)
}

type NetTunnelArrayOutput struct{ *pulumi.OutputState }

func (NetTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetTunnel)(nil)).Elem()
}

func (o NetTunnelArrayOutput) ToNetTunnelArrayOutput() NetTunnelArrayOutput {
	return o
}

func (o NetTunnelArrayOutput) ToNetTunnelArrayOutputWithContext(ctx context.Context) NetTunnelArrayOutput {
	return o
}

func (o NetTunnelArrayOutput) Index(i pulumi.IntInput) NetTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetTunnel {
		return vs[0].([]*NetTunnel)[vs[1].(int)]
	}).(NetTunnelOutput)
}

type NetTunnelMapOutput struct{ *pulumi.OutputState }

func (NetTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetTunnel)(nil)).Elem()
}

func (o NetTunnelMapOutput) ToNetTunnelMapOutput() NetTunnelMapOutput {
	return o
}

func (o NetTunnelMapOutput) ToNetTunnelMapOutputWithContext(ctx context.Context) NetTunnelMapOutput {
	return o
}

func (o NetTunnelMapOutput) MapIndex(k pulumi.StringInput) NetTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetTunnel {
		return vs[0].(map[string]*NetTunnel)[vs[1].(string)]
	}).(NetTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetTunnelInput)(nil)).Elem(), &NetTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetTunnelArrayInput)(nil)).Elem(), NetTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetTunnelMapInput)(nil)).Elem(), NetTunnelMap{})
	pulumi.RegisterOutputType(NetTunnelOutput{})
	pulumi.RegisterOutputType(NetTunnelArrayOutput{})
	pulumi.RegisterOutputType(NetTunnelMapOutput{})
}
