// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `FastHttpApp` This resource will create and manage FAST HTTP applications on BIG-IP
//
// [FAST documentation](https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/)
type FastHttpApp struct {
	pulumi.CustomResourceState

	// Name of the FAST HTTPS application.
	Application pulumi.StringOutput `pulumi:"application"`
	// List of LTM Policies to be applied FAST HTTP Application.
	EndpointLtmPolicies pulumi.StringArrayOutput `pulumi:"endpointLtmPolicies"`
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrOutput `pulumi:"existingMonitor"`
	// Select an existing BIG-IP Pool
	ExistingPool pulumi.StringPtrOutput `pulumi:"existingPool"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrOutput `pulumi:"existingSnatPool"`
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy pulumi.StringPtrOutput `pulumi:"existingWafSecurityPolicy"`
	// Json payload for FAST HTTP application.
	FastHttpJson pulumi.StringOutput `pulumi:"fastHttpJson"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrOutput `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastHttpAppMonitorPtrOutput `pulumi:"monitor"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastHttpAppPoolMemberArrayOutput `pulumi:"poolMembers"`
	// List of security log profiles to be used for FAST application
	SecurityLogProfiles pulumi.StringArrayOutput `pulumi:"securityLogProfiles"`
	// List of different cloud service discovery config provided as string, provided `serviceDiscovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
	ServiceDiscoveries pulumi.StringArrayOutput `pulumi:"serviceDiscoveries"`
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime pulumi.IntPtrOutput `pulumi:"slowRampTime"`
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses pulumi.StringArrayOutput `pulumi:"snatPoolAddresses"`
	// Name of the FAST HTTPS application tenant.
	Tenant pulumi.StringOutput `pulumi:"tenant"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastHttpAppVirtualServerPtrOutput `pulumi:"virtualServer"`
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy FastHttpAppWafSecurityPolicyPtrOutput `pulumi:"wafSecurityPolicy"`
}

// NewFastHttpApp registers a new resource with the given unique name, arguments, and options.
func NewFastHttpApp(ctx *pulumi.Context,
	name string, args *FastHttpAppArgs, opts ...pulumi.ResourceOption) (*FastHttpApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Application == nil {
		return nil, errors.New("invalid value for required argument 'Application'")
	}
	if args.Tenant == nil {
		return nil, errors.New("invalid value for required argument 'Tenant'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FastHttpApp
	err := ctx.RegisterResource("f5bigip:index/fastHttpApp:FastHttpApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFastHttpApp gets an existing FastHttpApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFastHttpApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FastHttpAppState, opts ...pulumi.ResourceOption) (*FastHttpApp, error) {
	var resource FastHttpApp
	err := ctx.ReadResource("f5bigip:index/fastHttpApp:FastHttpApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FastHttpApp resources.
type fastHttpAppState struct {
	// Name of the FAST HTTPS application.
	Application *string `pulumi:"application"`
	// List of LTM Policies to be applied FAST HTTP Application.
	EndpointLtmPolicies []string `pulumi:"endpointLtmPolicies"`
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor *string `pulumi:"existingMonitor"`
	// Select an existing BIG-IP Pool
	ExistingPool *string `pulumi:"existingPool"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool *string `pulumi:"existingSnatPool"`
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy *string `pulumi:"existingWafSecurityPolicy"`
	// Json payload for FAST HTTP application.
	FastHttpJson *string `pulumi:"fastHttpJson"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor *FastHttpAppMonitor `pulumi:"monitor"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers []FastHttpAppPoolMember `pulumi:"poolMembers"`
	// List of security log profiles to be used for FAST application
	SecurityLogProfiles []string `pulumi:"securityLogProfiles"`
	// List of different cloud service discovery config provided as string, provided `serviceDiscovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
	ServiceDiscoveries []string `pulumi:"serviceDiscoveries"`
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime *int `pulumi:"slowRampTime"`
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses []string `pulumi:"snatPoolAddresses"`
	// Name of the FAST HTTPS application tenant.
	Tenant *string `pulumi:"tenant"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer *FastHttpAppVirtualServer `pulumi:"virtualServer"`
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy *FastHttpAppWafSecurityPolicy `pulumi:"wafSecurityPolicy"`
}

type FastHttpAppState struct {
	// Name of the FAST HTTPS application.
	Application pulumi.StringPtrInput
	// List of LTM Policies to be applied FAST HTTP Application.
	EndpointLtmPolicies pulumi.StringArrayInput
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrInput
	// Select an existing BIG-IP Pool
	ExistingPool pulumi.StringPtrInput
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrInput
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy pulumi.StringPtrInput
	// Json payload for FAST HTTP application.
	FastHttpJson pulumi.StringPtrInput
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrInput
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastHttpAppMonitorPtrInput
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastHttpAppPoolMemberArrayInput
	// List of security log profiles to be used for FAST application
	SecurityLogProfiles pulumi.StringArrayInput
	// List of different cloud service discovery config provided as string, provided `serviceDiscovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
	ServiceDiscoveries pulumi.StringArrayInput
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime pulumi.IntPtrInput
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses pulumi.StringArrayInput
	// Name of the FAST HTTPS application tenant.
	Tenant pulumi.StringPtrInput
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastHttpAppVirtualServerPtrInput
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy FastHttpAppWafSecurityPolicyPtrInput
}

func (FastHttpAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*fastHttpAppState)(nil)).Elem()
}

type fastHttpAppArgs struct {
	// Name of the FAST HTTPS application.
	Application string `pulumi:"application"`
	// List of LTM Policies to be applied FAST HTTP Application.
	EndpointLtmPolicies []string `pulumi:"endpointLtmPolicies"`
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor *string `pulumi:"existingMonitor"`
	// Select an existing BIG-IP Pool
	ExistingPool *string `pulumi:"existingPool"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool *string `pulumi:"existingSnatPool"`
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy *string `pulumi:"existingWafSecurityPolicy"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor *FastHttpAppMonitor `pulumi:"monitor"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers []FastHttpAppPoolMember `pulumi:"poolMembers"`
	// List of security log profiles to be used for FAST application
	SecurityLogProfiles []string `pulumi:"securityLogProfiles"`
	// List of different cloud service discovery config provided as string, provided `serviceDiscovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
	ServiceDiscoveries []string `pulumi:"serviceDiscoveries"`
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime *int `pulumi:"slowRampTime"`
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses []string `pulumi:"snatPoolAddresses"`
	// Name of the FAST HTTPS application tenant.
	Tenant string `pulumi:"tenant"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer *FastHttpAppVirtualServer `pulumi:"virtualServer"`
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy *FastHttpAppWafSecurityPolicy `pulumi:"wafSecurityPolicy"`
}

// The set of arguments for constructing a FastHttpApp resource.
type FastHttpAppArgs struct {
	// Name of the FAST HTTPS application.
	Application pulumi.StringInput
	// List of LTM Policies to be applied FAST HTTP Application.
	EndpointLtmPolicies pulumi.StringArrayInput
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrInput
	// Select an existing BIG-IP Pool
	ExistingPool pulumi.StringPtrInput
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrInput
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy pulumi.StringPtrInput
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrInput
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastHttpAppMonitorPtrInput
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastHttpAppPoolMemberArrayInput
	// List of security log profiles to be used for FAST application
	SecurityLogProfiles pulumi.StringArrayInput
	// List of different cloud service discovery config provided as string, provided `serviceDiscovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
	ServiceDiscoveries pulumi.StringArrayInput
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime pulumi.IntPtrInput
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses pulumi.StringArrayInput
	// Name of the FAST HTTPS application tenant.
	Tenant pulumi.StringInput
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastHttpAppVirtualServerPtrInput
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy FastHttpAppWafSecurityPolicyPtrInput
}

func (FastHttpAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fastHttpAppArgs)(nil)).Elem()
}

type FastHttpAppInput interface {
	pulumi.Input

	ToFastHttpAppOutput() FastHttpAppOutput
	ToFastHttpAppOutputWithContext(ctx context.Context) FastHttpAppOutput
}

func (*FastHttpApp) ElementType() reflect.Type {
	return reflect.TypeOf((**FastHttpApp)(nil)).Elem()
}

func (i *FastHttpApp) ToFastHttpAppOutput() FastHttpAppOutput {
	return i.ToFastHttpAppOutputWithContext(context.Background())
}

func (i *FastHttpApp) ToFastHttpAppOutputWithContext(ctx context.Context) FastHttpAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastHttpAppOutput)
}

func (i *FastHttpApp) ToOutput(ctx context.Context) pulumix.Output[*FastHttpApp] {
	return pulumix.Output[*FastHttpApp]{
		OutputState: i.ToFastHttpAppOutputWithContext(ctx).OutputState,
	}
}

// FastHttpAppArrayInput is an input type that accepts FastHttpAppArray and FastHttpAppArrayOutput values.
// You can construct a concrete instance of `FastHttpAppArrayInput` via:
//
//	FastHttpAppArray{ FastHttpAppArgs{...} }
type FastHttpAppArrayInput interface {
	pulumi.Input

	ToFastHttpAppArrayOutput() FastHttpAppArrayOutput
	ToFastHttpAppArrayOutputWithContext(context.Context) FastHttpAppArrayOutput
}

type FastHttpAppArray []FastHttpAppInput

func (FastHttpAppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastHttpApp)(nil)).Elem()
}

func (i FastHttpAppArray) ToFastHttpAppArrayOutput() FastHttpAppArrayOutput {
	return i.ToFastHttpAppArrayOutputWithContext(context.Background())
}

func (i FastHttpAppArray) ToFastHttpAppArrayOutputWithContext(ctx context.Context) FastHttpAppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastHttpAppArrayOutput)
}

func (i FastHttpAppArray) ToOutput(ctx context.Context) pulumix.Output[[]*FastHttpApp] {
	return pulumix.Output[[]*FastHttpApp]{
		OutputState: i.ToFastHttpAppArrayOutputWithContext(ctx).OutputState,
	}
}

// FastHttpAppMapInput is an input type that accepts FastHttpAppMap and FastHttpAppMapOutput values.
// You can construct a concrete instance of `FastHttpAppMapInput` via:
//
//	FastHttpAppMap{ "key": FastHttpAppArgs{...} }
type FastHttpAppMapInput interface {
	pulumi.Input

	ToFastHttpAppMapOutput() FastHttpAppMapOutput
	ToFastHttpAppMapOutputWithContext(context.Context) FastHttpAppMapOutput
}

type FastHttpAppMap map[string]FastHttpAppInput

func (FastHttpAppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastHttpApp)(nil)).Elem()
}

func (i FastHttpAppMap) ToFastHttpAppMapOutput() FastHttpAppMapOutput {
	return i.ToFastHttpAppMapOutputWithContext(context.Background())
}

func (i FastHttpAppMap) ToFastHttpAppMapOutputWithContext(ctx context.Context) FastHttpAppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastHttpAppMapOutput)
}

func (i FastHttpAppMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FastHttpApp] {
	return pulumix.Output[map[string]*FastHttpApp]{
		OutputState: i.ToFastHttpAppMapOutputWithContext(ctx).OutputState,
	}
}

type FastHttpAppOutput struct{ *pulumi.OutputState }

func (FastHttpAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FastHttpApp)(nil)).Elem()
}

func (o FastHttpAppOutput) ToFastHttpAppOutput() FastHttpAppOutput {
	return o
}

func (o FastHttpAppOutput) ToFastHttpAppOutputWithContext(ctx context.Context) FastHttpAppOutput {
	return o
}

func (o FastHttpAppOutput) ToOutput(ctx context.Context) pulumix.Output[*FastHttpApp] {
	return pulumix.Output[*FastHttpApp]{
		OutputState: o.OutputState,
	}
}

// Name of the FAST HTTPS application.
func (o FastHttpAppOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// List of LTM Policies to be applied FAST HTTP Application.
func (o FastHttpAppOutput) EndpointLtmPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringArrayOutput { return v.EndpointLtmPolicies }).(pulumi.StringArrayOutput)
}

// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
func (o FastHttpAppOutput) ExistingMonitor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringPtrOutput { return v.ExistingMonitor }).(pulumi.StringPtrOutput)
}

// Select an existing BIG-IP Pool
func (o FastHttpAppOutput) ExistingPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringPtrOutput { return v.ExistingPool }).(pulumi.StringPtrOutput)
}

// Name of an existing BIG-IP SNAT pool.
func (o FastHttpAppOutput) ExistingSnatPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringPtrOutput { return v.ExistingSnatPool }).(pulumi.StringPtrOutput)
}

// Name of an existing WAF Security policy.
func (o FastHttpAppOutput) ExistingWafSecurityPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringPtrOutput { return v.ExistingWafSecurityPolicy }).(pulumi.StringPtrOutput)
}

// Json payload for FAST HTTP application.
func (o FastHttpAppOutput) FastHttpJson() pulumi.StringOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringOutput { return v.FastHttpJson }).(pulumi.StringOutput)
}

// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
func (o FastHttpAppOutput) LoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringPtrOutput { return v.LoadBalancingMode }).(pulumi.StringPtrOutput)
}

// `monitor` block takes input for FAST-Generated Pool Monitor.
// See Pool Monitor below for more details.
func (o FastHttpAppOutput) Monitor() FastHttpAppMonitorPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) FastHttpAppMonitorPtrOutput { return v.Monitor }).(FastHttpAppMonitorPtrOutput)
}

// `poolMembers` block takes input for FAST-Generated Pool.
// See Pool Members below for more details.
func (o FastHttpAppOutput) PoolMembers() FastHttpAppPoolMemberArrayOutput {
	return o.ApplyT(func(v *FastHttpApp) FastHttpAppPoolMemberArrayOutput { return v.PoolMembers }).(FastHttpAppPoolMemberArrayOutput)
}

// List of security log profiles to be used for FAST application
func (o FastHttpAppOutput) SecurityLogProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringArrayOutput { return v.SecurityLogProfiles }).(pulumi.StringArrayOutput)
}

// List of different cloud service discovery config provided as string, provided `serviceDiscovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
func (o FastHttpAppOutput) ServiceDiscoveries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringArrayOutput { return v.ServiceDiscoveries }).(pulumi.StringArrayOutput)
}

// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
func (o FastHttpAppOutput) SlowRampTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.IntPtrOutput { return v.SlowRampTime }).(pulumi.IntPtrOutput)
}

// List of address to be used for FAST-Generated SNAT Pool.
func (o FastHttpAppOutput) SnatPoolAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringArrayOutput { return v.SnatPoolAddresses }).(pulumi.StringArrayOutput)
}

// Name of the FAST HTTPS application tenant.
func (o FastHttpAppOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *FastHttpApp) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
// See virtual server below for more details.
func (o FastHttpAppOutput) VirtualServer() FastHttpAppVirtualServerPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) FastHttpAppVirtualServerPtrOutput { return v.VirtualServer }).(FastHttpAppVirtualServerPtrOutput)
}

// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
// See WAF Security Policy below for more details.
func (o FastHttpAppOutput) WafSecurityPolicy() FastHttpAppWafSecurityPolicyPtrOutput {
	return o.ApplyT(func(v *FastHttpApp) FastHttpAppWafSecurityPolicyPtrOutput { return v.WafSecurityPolicy }).(FastHttpAppWafSecurityPolicyPtrOutput)
}

type FastHttpAppArrayOutput struct{ *pulumi.OutputState }

func (FastHttpAppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastHttpApp)(nil)).Elem()
}

func (o FastHttpAppArrayOutput) ToFastHttpAppArrayOutput() FastHttpAppArrayOutput {
	return o
}

func (o FastHttpAppArrayOutput) ToFastHttpAppArrayOutputWithContext(ctx context.Context) FastHttpAppArrayOutput {
	return o
}

func (o FastHttpAppArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FastHttpApp] {
	return pulumix.Output[[]*FastHttpApp]{
		OutputState: o.OutputState,
	}
}

func (o FastHttpAppArrayOutput) Index(i pulumi.IntInput) FastHttpAppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FastHttpApp {
		return vs[0].([]*FastHttpApp)[vs[1].(int)]
	}).(FastHttpAppOutput)
}

type FastHttpAppMapOutput struct{ *pulumi.OutputState }

func (FastHttpAppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastHttpApp)(nil)).Elem()
}

func (o FastHttpAppMapOutput) ToFastHttpAppMapOutput() FastHttpAppMapOutput {
	return o
}

func (o FastHttpAppMapOutput) ToFastHttpAppMapOutputWithContext(ctx context.Context) FastHttpAppMapOutput {
	return o
}

func (o FastHttpAppMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FastHttpApp] {
	return pulumix.Output[map[string]*FastHttpApp]{
		OutputState: o.OutputState,
	}
}

func (o FastHttpAppMapOutput) MapIndex(k pulumi.StringInput) FastHttpAppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FastHttpApp {
		return vs[0].(map[string]*FastHttpApp)[vs[1].(string)]
	}).(FastHttpAppOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FastHttpAppInput)(nil)).Elem(), &FastHttpApp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastHttpAppArrayInput)(nil)).Elem(), FastHttpAppArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastHttpAppMapInput)(nil)).Elem(), FastHttpAppMap{})
	pulumi.RegisterOutputType(FastHttpAppOutput{})
	pulumi.RegisterOutputType(FastHttpAppArrayOutput{})
	pulumi.RegisterOutputType(FastHttpAppMapOutput{})
}
