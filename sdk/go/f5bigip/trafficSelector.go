// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `TrafficSelector` Manage IPSec Traffic Selectors on BIG-IP
//
// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-selector)
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
//			_, err := f5bigip.NewTrafficSelector(ctx, "test-selector", &f5bigip.TrafficSelectorArgs{
//				Name:               pulumi.String("/Common/test-selector"),
//				DestinationAddress: pulumi.String("3.10.11.2/32"),
//				SourceAddress:      pulumi.String("2.10.11.12/32"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TrafficSelector struct {
	pulumi.CustomResourceState

	// Description of the traffic selector.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
	DestinationAddress pulumi.StringOutput `pulumi:"destinationAddress"`
	// Specifies the IP port used by the application. The default value is `All Ports (0)`
	DestinationPort pulumi.IntOutput `pulumi:"destinationPort"`
	// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
	IpProtocol pulumi.IntOutput `pulumi:"ipProtocol"`
	// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
	IpsecPolicy pulumi.StringOutput `pulumi:"ipsecPolicy"`
	// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
	// When creating a new traffic selector, if this parameter is not specified, the default is `last`
	Order pulumi.IntOutput `pulumi:"order"`
	// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
	SourceAddress pulumi.StringOutput `pulumi:"sourceAddress"`
	// Specifies the IP port used by the application. The default value is `All Ports (0)`.
	SourcePort pulumi.IntOutput `pulumi:"sourcePort"`
}

// NewTrafficSelector registers a new resource with the given unique name, arguments, and options.
func NewTrafficSelector(ctx *pulumi.Context,
	name string, args *TrafficSelectorArgs, opts ...pulumi.ResourceOption) (*TrafficSelector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationAddress == nil {
		return nil, errors.New("invalid value for required argument 'DestinationAddress'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.SourceAddress == nil {
		return nil, errors.New("invalid value for required argument 'SourceAddress'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrafficSelector
	err := ctx.RegisterResource("f5bigip:index/trafficSelector:TrafficSelector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficSelector gets an existing TrafficSelector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficSelector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficSelectorState, opts ...pulumi.ResourceOption) (*TrafficSelector, error) {
	var resource TrafficSelector
	err := ctx.ReadResource("f5bigip:index/trafficSelector:TrafficSelector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficSelector resources.
type trafficSelectorState struct {
	// Description of the traffic selector.
	Description *string `pulumi:"description"`
	// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
	DestinationAddress *string `pulumi:"destinationAddress"`
	// Specifies the IP port used by the application. The default value is `All Ports (0)`
	DestinationPort *int `pulumi:"destinationPort"`
	// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
	Direction *string `pulumi:"direction"`
	// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
	IpProtocol *int `pulumi:"ipProtocol"`
	// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
	IpsecPolicy *string `pulumi:"ipsecPolicy"`
	// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
	Name *string `pulumi:"name"`
	// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
	// When creating a new traffic selector, if this parameter is not specified, the default is `last`
	Order *int `pulumi:"order"`
	// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
	SourceAddress *string `pulumi:"sourceAddress"`
	// Specifies the IP port used by the application. The default value is `All Ports (0)`.
	SourcePort *int `pulumi:"sourcePort"`
}

type TrafficSelectorState struct {
	// Description of the traffic selector.
	Description pulumi.StringPtrInput
	// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
	DestinationAddress pulumi.StringPtrInput
	// Specifies the IP port used by the application. The default value is `All Ports (0)`
	DestinationPort pulumi.IntPtrInput
	// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
	Direction pulumi.StringPtrInput
	// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
	IpProtocol pulumi.IntPtrInput
	// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
	IpsecPolicy pulumi.StringPtrInput
	// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
	Name pulumi.StringPtrInput
	// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
	// When creating a new traffic selector, if this parameter is not specified, the default is `last`
	Order pulumi.IntPtrInput
	// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
	SourceAddress pulumi.StringPtrInput
	// Specifies the IP port used by the application. The default value is `All Ports (0)`.
	SourcePort pulumi.IntPtrInput
}

func (TrafficSelectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficSelectorState)(nil)).Elem()
}

type trafficSelectorArgs struct {
	// Description of the traffic selector.
	Description *string `pulumi:"description"`
	// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
	DestinationAddress string `pulumi:"destinationAddress"`
	// Specifies the IP port used by the application. The default value is `All Ports (0)`
	DestinationPort *int `pulumi:"destinationPort"`
	// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
	Direction *string `pulumi:"direction"`
	// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
	IpProtocol *int `pulumi:"ipProtocol"`
	// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
	IpsecPolicy *string `pulumi:"ipsecPolicy"`
	// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
	Name string `pulumi:"name"`
	// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
	// When creating a new traffic selector, if this parameter is not specified, the default is `last`
	Order *int `pulumi:"order"`
	// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
	SourceAddress string `pulumi:"sourceAddress"`
	// Specifies the IP port used by the application. The default value is `All Ports (0)`.
	SourcePort *int `pulumi:"sourcePort"`
}

// The set of arguments for constructing a TrafficSelector resource.
type TrafficSelectorArgs struct {
	// Description of the traffic selector.
	Description pulumi.StringPtrInput
	// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
	DestinationAddress pulumi.StringInput
	// Specifies the IP port used by the application. The default value is `All Ports (0)`
	DestinationPort pulumi.IntPtrInput
	// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
	Direction pulumi.StringPtrInput
	// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
	IpProtocol pulumi.IntPtrInput
	// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
	IpsecPolicy pulumi.StringPtrInput
	// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
	Name pulumi.StringInput
	// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
	// When creating a new traffic selector, if this parameter is not specified, the default is `last`
	Order pulumi.IntPtrInput
	// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
	SourceAddress pulumi.StringInput
	// Specifies the IP port used by the application. The default value is `All Ports (0)`.
	SourcePort pulumi.IntPtrInput
}

func (TrafficSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficSelectorArgs)(nil)).Elem()
}

type TrafficSelectorInput interface {
	pulumi.Input

	ToTrafficSelectorOutput() TrafficSelectorOutput
	ToTrafficSelectorOutputWithContext(ctx context.Context) TrafficSelectorOutput
}

func (*TrafficSelector) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficSelector)(nil)).Elem()
}

func (i *TrafficSelector) ToTrafficSelectorOutput() TrafficSelectorOutput {
	return i.ToTrafficSelectorOutputWithContext(context.Background())
}

func (i *TrafficSelector) ToTrafficSelectorOutputWithContext(ctx context.Context) TrafficSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSelectorOutput)
}

// TrafficSelectorArrayInput is an input type that accepts TrafficSelectorArray and TrafficSelectorArrayOutput values.
// You can construct a concrete instance of `TrafficSelectorArrayInput` via:
//
//	TrafficSelectorArray{ TrafficSelectorArgs{...} }
type TrafficSelectorArrayInput interface {
	pulumi.Input

	ToTrafficSelectorArrayOutput() TrafficSelectorArrayOutput
	ToTrafficSelectorArrayOutputWithContext(context.Context) TrafficSelectorArrayOutput
}

type TrafficSelectorArray []TrafficSelectorInput

func (TrafficSelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficSelector)(nil)).Elem()
}

func (i TrafficSelectorArray) ToTrafficSelectorArrayOutput() TrafficSelectorArrayOutput {
	return i.ToTrafficSelectorArrayOutputWithContext(context.Background())
}

func (i TrafficSelectorArray) ToTrafficSelectorArrayOutputWithContext(ctx context.Context) TrafficSelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSelectorArrayOutput)
}

// TrafficSelectorMapInput is an input type that accepts TrafficSelectorMap and TrafficSelectorMapOutput values.
// You can construct a concrete instance of `TrafficSelectorMapInput` via:
//
//	TrafficSelectorMap{ "key": TrafficSelectorArgs{...} }
type TrafficSelectorMapInput interface {
	pulumi.Input

	ToTrafficSelectorMapOutput() TrafficSelectorMapOutput
	ToTrafficSelectorMapOutputWithContext(context.Context) TrafficSelectorMapOutput
}

type TrafficSelectorMap map[string]TrafficSelectorInput

func (TrafficSelectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficSelector)(nil)).Elem()
}

func (i TrafficSelectorMap) ToTrafficSelectorMapOutput() TrafficSelectorMapOutput {
	return i.ToTrafficSelectorMapOutputWithContext(context.Background())
}

func (i TrafficSelectorMap) ToTrafficSelectorMapOutputWithContext(ctx context.Context) TrafficSelectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficSelectorMapOutput)
}

type TrafficSelectorOutput struct{ *pulumi.OutputState }

func (TrafficSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficSelector)(nil)).Elem()
}

func (o TrafficSelectorOutput) ToTrafficSelectorOutput() TrafficSelectorOutput {
	return o
}

func (o TrafficSelectorOutput) ToTrafficSelectorOutputWithContext(ctx context.Context) TrafficSelectorOutput {
	return o
}

// Description of the traffic selector.
func (o TrafficSelectorOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
func (o TrafficSelectorOutput) DestinationAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.StringOutput { return v.DestinationAddress }).(pulumi.StringOutput)
}

// Specifies the IP port used by the application. The default value is `All Ports (0)`
func (o TrafficSelectorOutput) DestinationPort() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.IntOutput { return v.DestinationPort }).(pulumi.IntOutput)
}

// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
func (o TrafficSelectorOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
func (o TrafficSelectorOutput) IpProtocol() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.IntOutput { return v.IpProtocol }).(pulumi.IntOutput)
}

// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
func (o TrafficSelectorOutput) IpsecPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.StringOutput { return v.IpsecPolicy }).(pulumi.StringOutput)
}

// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
func (o TrafficSelectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
// When creating a new traffic selector, if this parameter is not specified, the default is `last`
func (o TrafficSelectorOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
func (o TrafficSelectorOutput) SourceAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.StringOutput { return v.SourceAddress }).(pulumi.StringOutput)
}

// Specifies the IP port used by the application. The default value is `All Ports (0)`.
func (o TrafficSelectorOutput) SourcePort() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficSelector) pulumi.IntOutput { return v.SourcePort }).(pulumi.IntOutput)
}

type TrafficSelectorArrayOutput struct{ *pulumi.OutputState }

func (TrafficSelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficSelector)(nil)).Elem()
}

func (o TrafficSelectorArrayOutput) ToTrafficSelectorArrayOutput() TrafficSelectorArrayOutput {
	return o
}

func (o TrafficSelectorArrayOutput) ToTrafficSelectorArrayOutputWithContext(ctx context.Context) TrafficSelectorArrayOutput {
	return o
}

func (o TrafficSelectorArrayOutput) Index(i pulumi.IntInput) TrafficSelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficSelector {
		return vs[0].([]*TrafficSelector)[vs[1].(int)]
	}).(TrafficSelectorOutput)
}

type TrafficSelectorMapOutput struct{ *pulumi.OutputState }

func (TrafficSelectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficSelector)(nil)).Elem()
}

func (o TrafficSelectorMapOutput) ToTrafficSelectorMapOutput() TrafficSelectorMapOutput {
	return o
}

func (o TrafficSelectorMapOutput) ToTrafficSelectorMapOutputWithContext(ctx context.Context) TrafficSelectorMapOutput {
	return o
}

func (o TrafficSelectorMapOutput) MapIndex(k pulumi.StringInput) TrafficSelectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficSelector {
		return vs[0].(map[string]*TrafficSelector)[vs[1].(string)]
	}).(TrafficSelectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficSelectorInput)(nil)).Elem(), &TrafficSelector{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficSelectorArrayInput)(nil)).Elem(), TrafficSelectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficSelectorMapInput)(nil)).Elem(), TrafficSelectorMap{})
	pulumi.RegisterOutputType(TrafficSelectorOutput{})
	pulumi.RegisterOutputType(TrafficSelectorArrayOutput{})
	pulumi.RegisterOutputType(TrafficSelectorMapOutput{})
}
