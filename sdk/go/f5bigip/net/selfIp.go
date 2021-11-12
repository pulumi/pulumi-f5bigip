// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `net.SelfIp` Manages a selfip configuration
//
// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-selfip.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/net"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := net.NewSelfIp(ctx, "selfip1", &net.SelfIpArgs{
// 			Name:         pulumi.String("/Common/internalselfIP"),
// 			Ip:           pulumi.String("11.1.1.1/24"),
// 			Vlan:         pulumi.String("/Common/internal"),
// 			TrafficGroup: pulumi.String("traffic-group-1"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			bigip_net_vlan.Vlan1,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SelfIp struct {
	pulumi.CustomResourceState

	// The Self IP's address and netmask.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Name of the selfip
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
	TrafficGroup pulumi.StringPtrOutput `pulumi:"trafficGroup"`
	// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
	Vlan pulumi.StringOutput `pulumi:"vlan"`
}

// NewSelfIp registers a new resource with the given unique name, arguments, and options.
func NewSelfIp(ctx *pulumi.Context,
	name string, args *SelfIpArgs, opts ...pulumi.ResourceOption) (*SelfIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Vlan == nil {
		return nil, errors.New("invalid value for required argument 'Vlan'")
	}
	var resource SelfIp
	err := ctx.RegisterResource("f5bigip:net/selfIp:SelfIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfIp gets an existing SelfIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfIpState, opts ...pulumi.ResourceOption) (*SelfIp, error) {
	var resource SelfIp
	err := ctx.ReadResource("f5bigip:net/selfIp:SelfIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfIp resources.
type selfIpState struct {
	// The Self IP's address and netmask.
	Ip *string `pulumi:"ip"`
	// Name of the selfip
	Name *string `pulumi:"name"`
	// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
	TrafficGroup *string `pulumi:"trafficGroup"`
	// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
	Vlan *string `pulumi:"vlan"`
}

type SelfIpState struct {
	// The Self IP's address and netmask.
	Ip pulumi.StringPtrInput
	// Name of the selfip
	Name pulumi.StringPtrInput
	// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
	TrafficGroup pulumi.StringPtrInput
	// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
	Vlan pulumi.StringPtrInput
}

func (SelfIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfIpState)(nil)).Elem()
}

type selfIpArgs struct {
	// The Self IP's address and netmask.
	Ip string `pulumi:"ip"`
	// Name of the selfip
	Name string `pulumi:"name"`
	// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
	TrafficGroup *string `pulumi:"trafficGroup"`
	// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
	Vlan string `pulumi:"vlan"`
}

// The set of arguments for constructing a SelfIp resource.
type SelfIpArgs struct {
	// The Self IP's address and netmask.
	Ip pulumi.StringInput
	// Name of the selfip
	Name pulumi.StringInput
	// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
	TrafficGroup pulumi.StringPtrInput
	// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
	Vlan pulumi.StringInput
}

func (SelfIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfIpArgs)(nil)).Elem()
}

type SelfIpInput interface {
	pulumi.Input

	ToSelfIpOutput() SelfIpOutput
	ToSelfIpOutputWithContext(ctx context.Context) SelfIpOutput
}

func (*SelfIp) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfIp)(nil))
}

func (i *SelfIp) ToSelfIpOutput() SelfIpOutput {
	return i.ToSelfIpOutputWithContext(context.Background())
}

func (i *SelfIp) ToSelfIpOutputWithContext(ctx context.Context) SelfIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfIpOutput)
}

func (i *SelfIp) ToSelfIpPtrOutput() SelfIpPtrOutput {
	return i.ToSelfIpPtrOutputWithContext(context.Background())
}

func (i *SelfIp) ToSelfIpPtrOutputWithContext(ctx context.Context) SelfIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfIpPtrOutput)
}

type SelfIpPtrInput interface {
	pulumi.Input

	ToSelfIpPtrOutput() SelfIpPtrOutput
	ToSelfIpPtrOutputWithContext(ctx context.Context) SelfIpPtrOutput
}

type selfIpPtrType SelfIpArgs

func (*selfIpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfIp)(nil))
}

func (i *selfIpPtrType) ToSelfIpPtrOutput() SelfIpPtrOutput {
	return i.ToSelfIpPtrOutputWithContext(context.Background())
}

func (i *selfIpPtrType) ToSelfIpPtrOutputWithContext(ctx context.Context) SelfIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfIpPtrOutput)
}

// SelfIpArrayInput is an input type that accepts SelfIpArray and SelfIpArrayOutput values.
// You can construct a concrete instance of `SelfIpArrayInput` via:
//
//          SelfIpArray{ SelfIpArgs{...} }
type SelfIpArrayInput interface {
	pulumi.Input

	ToSelfIpArrayOutput() SelfIpArrayOutput
	ToSelfIpArrayOutputWithContext(context.Context) SelfIpArrayOutput
}

type SelfIpArray []SelfIpInput

func (SelfIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfIp)(nil)).Elem()
}

func (i SelfIpArray) ToSelfIpArrayOutput() SelfIpArrayOutput {
	return i.ToSelfIpArrayOutputWithContext(context.Background())
}

func (i SelfIpArray) ToSelfIpArrayOutputWithContext(ctx context.Context) SelfIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfIpArrayOutput)
}

// SelfIpMapInput is an input type that accepts SelfIpMap and SelfIpMapOutput values.
// You can construct a concrete instance of `SelfIpMapInput` via:
//
//          SelfIpMap{ "key": SelfIpArgs{...} }
type SelfIpMapInput interface {
	pulumi.Input

	ToSelfIpMapOutput() SelfIpMapOutput
	ToSelfIpMapOutputWithContext(context.Context) SelfIpMapOutput
}

type SelfIpMap map[string]SelfIpInput

func (SelfIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfIp)(nil)).Elem()
}

func (i SelfIpMap) ToSelfIpMapOutput() SelfIpMapOutput {
	return i.ToSelfIpMapOutputWithContext(context.Background())
}

func (i SelfIpMap) ToSelfIpMapOutputWithContext(ctx context.Context) SelfIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfIpMapOutput)
}

type SelfIpOutput struct{ *pulumi.OutputState }

func (SelfIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfIp)(nil))
}

func (o SelfIpOutput) ToSelfIpOutput() SelfIpOutput {
	return o
}

func (o SelfIpOutput) ToSelfIpOutputWithContext(ctx context.Context) SelfIpOutput {
	return o
}

func (o SelfIpOutput) ToSelfIpPtrOutput() SelfIpPtrOutput {
	return o.ToSelfIpPtrOutputWithContext(context.Background())
}

func (o SelfIpOutput) ToSelfIpPtrOutputWithContext(ctx context.Context) SelfIpPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SelfIp) *SelfIp {
		return &v
	}).(SelfIpPtrOutput)
}

type SelfIpPtrOutput struct{ *pulumi.OutputState }

func (SelfIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfIp)(nil))
}

func (o SelfIpPtrOutput) ToSelfIpPtrOutput() SelfIpPtrOutput {
	return o
}

func (o SelfIpPtrOutput) ToSelfIpPtrOutputWithContext(ctx context.Context) SelfIpPtrOutput {
	return o
}

func (o SelfIpPtrOutput) Elem() SelfIpOutput {
	return o.ApplyT(func(v *SelfIp) SelfIp {
		if v != nil {
			return *v
		}
		var ret SelfIp
		return ret
	}).(SelfIpOutput)
}

type SelfIpArrayOutput struct{ *pulumi.OutputState }

func (SelfIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelfIp)(nil))
}

func (o SelfIpArrayOutput) ToSelfIpArrayOutput() SelfIpArrayOutput {
	return o
}

func (o SelfIpArrayOutput) ToSelfIpArrayOutputWithContext(ctx context.Context) SelfIpArrayOutput {
	return o
}

func (o SelfIpArrayOutput) Index(i pulumi.IntInput) SelfIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelfIp {
		return vs[0].([]SelfIp)[vs[1].(int)]
	}).(SelfIpOutput)
}

type SelfIpMapOutput struct{ *pulumi.OutputState }

func (SelfIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SelfIp)(nil))
}

func (o SelfIpMapOutput) ToSelfIpMapOutput() SelfIpMapOutput {
	return o
}

func (o SelfIpMapOutput) ToSelfIpMapOutputWithContext(ctx context.Context) SelfIpMapOutput {
	return o
}

func (o SelfIpMapOutput) MapIndex(k pulumi.StringInput) SelfIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SelfIp {
		return vs[0].(map[string]SelfIp)[vs[1].(string)]
	}).(SelfIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SelfIpInput)(nil)).Elem(), &SelfIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfIpPtrInput)(nil)).Elem(), &SelfIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfIpArrayInput)(nil)).Elem(), SelfIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfIpMapInput)(nil)).Elem(), SelfIpMap{})
	pulumi.RegisterOutputType(SelfIpOutput{})
	pulumi.RegisterOutputType(SelfIpPtrOutput{})
	pulumi.RegisterOutputType(SelfIpArrayOutput{})
	pulumi.RegisterOutputType(SelfIpMapOutput{})
}
