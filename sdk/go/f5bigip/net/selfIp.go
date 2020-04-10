// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `net.SelfIp` Manages a selfip configuration
//
// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-selfip.
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
	if args == nil || args.Ip == nil {
		return nil, errors.New("missing required argument 'Ip'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Vlan == nil {
		return nil, errors.New("missing required argument 'Vlan'")
	}
	if args == nil {
		args = &SelfIpArgs{}
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
