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

// `FastUdpApp` This resource will create and manage FAST UDP applications on BIG-IP from provided JSON declaration.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := f5bigip.NewFastUdpApp(ctx, "fast-udp-app", &f5bigip.FastUdpAppArgs{
//				Application: pulumi.String("udp_app_2"),
//				PoolMembers: f5bigip.FastUdpAppPoolMemberArray{
//					&f5bigip.FastUdpAppPoolMemberArgs{
//						Addresses: pulumi.StringArray{
//							pulumi.String("10.11.34.65"),
//							pulumi.String("56.43.23.76"),
//						},
//						ConnectionLimit: pulumi.Int(4),
//						Port:            pulumi.Int(443),
//						PriorityGroup:   pulumi.Int(1),
//						ShareNodes:      pulumi.Bool(true),
//					},
//				},
//				Tenant: pulumi.String("udp_app_tenant"),
//				VirtualServer: &f5bigip.FastUdpAppVirtualServerArgs{
//					Ip:   pulumi.String("11.12.16.30"),
//					Port: pulumi.Int(443),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type FastUdpApp struct {
	pulumi.CustomResourceState

	// Name of the FAST UDP application.
	Application pulumi.StringOutput `pulumi:"application"`
	// Enables use of FastL4 profiles.
	EnableFastl4 pulumi.BoolPtrOutput `pulumi:"enableFastl4"`
	// Name of an existing BIG-IP UDP pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrOutput `pulumi:"existingMonitor"`
	// Name of an existing BIG-IP pool.
	ExistingPool pulumi.StringPtrOutput `pulumi:"existingPool"`
	// Name of an existing BIG-IP FastL4 or UDP profile.
	ExistingProfile pulumi.StringPtrOutput `pulumi:"existingProfile"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrOutput `pulumi:"existingSnatPool"`
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence pulumi.StringPtrOutput `pulumi:"fallbackPersistence"`
	// Json payload for FAST UDP application.
	FastUdpJson pulumi.StringOutput `pulumi:"fastUdpJson"`
	// Irules to attach to Virtual Server.
	Irules pulumi.StringArrayOutput `pulumi:"irules"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrOutput `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastUdpAppMonitorPtrOutput `pulumi:"monitor"`
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile pulumi.StringPtrOutput `pulumi:"persistenceProfile"`
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType pulumi.StringPtrOutput `pulumi:"persistenceType"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastUdpAppPoolMemberArrayOutput `pulumi:"poolMembers"`
	// Existing security log profiles to enable.
	SecurityLogProfiles pulumi.StringArrayOutput `pulumi:"securityLogProfiles"`
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime pulumi.IntPtrOutput `pulumi:"slowRampTime"`
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses pulumi.StringArrayOutput `pulumi:"snatPoolAddresses"`
	// Name of the FAST UDP application tenant.
	Tenant pulumi.StringOutput `pulumi:"tenant"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastUdpAppVirtualServerPtrOutput `pulumi:"virtualServer"`
	// Names of existing VLANs to allow.
	VlansAlloweds pulumi.StringArrayOutput `pulumi:"vlansAlloweds"`
	// Names of existing VLANs to reject.
	VlansRejecteds pulumi.StringArrayOutput `pulumi:"vlansRejecteds"`
}

// NewFastUdpApp registers a new resource with the given unique name, arguments, and options.
func NewFastUdpApp(ctx *pulumi.Context,
	name string, args *FastUdpAppArgs, opts ...pulumi.ResourceOption) (*FastUdpApp, error) {
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
	var resource FastUdpApp
	err := ctx.RegisterResource("f5bigip:index/fastUdpApp:FastUdpApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFastUdpApp gets an existing FastUdpApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFastUdpApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FastUdpAppState, opts ...pulumi.ResourceOption) (*FastUdpApp, error) {
	var resource FastUdpApp
	err := ctx.ReadResource("f5bigip:index/fastUdpApp:FastUdpApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FastUdpApp resources.
type fastUdpAppState struct {
	// Name of the FAST UDP application.
	Application *string `pulumi:"application"`
	// Enables use of FastL4 profiles.
	EnableFastl4 *bool `pulumi:"enableFastl4"`
	// Name of an existing BIG-IP UDP pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor *string `pulumi:"existingMonitor"`
	// Name of an existing BIG-IP pool.
	ExistingPool *string `pulumi:"existingPool"`
	// Name of an existing BIG-IP FastL4 or UDP profile.
	ExistingProfile *string `pulumi:"existingProfile"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool *string `pulumi:"existingSnatPool"`
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence *string `pulumi:"fallbackPersistence"`
	// Json payload for FAST UDP application.
	FastUdpJson *string `pulumi:"fastUdpJson"`
	// Irules to attach to Virtual Server.
	Irules []string `pulumi:"irules"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor *FastUdpAppMonitor `pulumi:"monitor"`
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile *string `pulumi:"persistenceProfile"`
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType *string `pulumi:"persistenceType"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers []FastUdpAppPoolMember `pulumi:"poolMembers"`
	// Existing security log profiles to enable.
	SecurityLogProfiles []string `pulumi:"securityLogProfiles"`
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime *int `pulumi:"slowRampTime"`
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses []string `pulumi:"snatPoolAddresses"`
	// Name of the FAST UDP application tenant.
	Tenant *string `pulumi:"tenant"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer *FastUdpAppVirtualServer `pulumi:"virtualServer"`
	// Names of existing VLANs to allow.
	VlansAlloweds []string `pulumi:"vlansAlloweds"`
	// Names of existing VLANs to reject.
	VlansRejecteds []string `pulumi:"vlansRejecteds"`
}

type FastUdpAppState struct {
	// Name of the FAST UDP application.
	Application pulumi.StringPtrInput
	// Enables use of FastL4 profiles.
	EnableFastl4 pulumi.BoolPtrInput
	// Name of an existing BIG-IP UDP pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrInput
	// Name of an existing BIG-IP pool.
	ExistingPool pulumi.StringPtrInput
	// Name of an existing BIG-IP FastL4 or UDP profile.
	ExistingProfile pulumi.StringPtrInput
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrInput
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence pulumi.StringPtrInput
	// Json payload for FAST UDP application.
	FastUdpJson pulumi.StringPtrInput
	// Irules to attach to Virtual Server.
	Irules pulumi.StringArrayInput
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrInput
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastUdpAppMonitorPtrInput
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile pulumi.StringPtrInput
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType pulumi.StringPtrInput
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastUdpAppPoolMemberArrayInput
	// Existing security log profiles to enable.
	SecurityLogProfiles pulumi.StringArrayInput
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime pulumi.IntPtrInput
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses pulumi.StringArrayInput
	// Name of the FAST UDP application tenant.
	Tenant pulumi.StringPtrInput
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastUdpAppVirtualServerPtrInput
	// Names of existing VLANs to allow.
	VlansAlloweds pulumi.StringArrayInput
	// Names of existing VLANs to reject.
	VlansRejecteds pulumi.StringArrayInput
}

func (FastUdpAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*fastUdpAppState)(nil)).Elem()
}

type fastUdpAppArgs struct {
	// Name of the FAST UDP application.
	Application string `pulumi:"application"`
	// Enables use of FastL4 profiles.
	EnableFastl4 *bool `pulumi:"enableFastl4"`
	// Name of an existing BIG-IP UDP pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor *string `pulumi:"existingMonitor"`
	// Name of an existing BIG-IP pool.
	ExistingPool *string `pulumi:"existingPool"`
	// Name of an existing BIG-IP FastL4 or UDP profile.
	ExistingProfile *string `pulumi:"existingProfile"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool *string `pulumi:"existingSnatPool"`
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence *string `pulumi:"fallbackPersistence"`
	// Irules to attach to Virtual Server.
	Irules []string `pulumi:"irules"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor *FastUdpAppMonitor `pulumi:"monitor"`
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile *string `pulumi:"persistenceProfile"`
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType *string `pulumi:"persistenceType"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers []FastUdpAppPoolMember `pulumi:"poolMembers"`
	// Existing security log profiles to enable.
	SecurityLogProfiles []string `pulumi:"securityLogProfiles"`
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime *int `pulumi:"slowRampTime"`
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses []string `pulumi:"snatPoolAddresses"`
	// Name of the FAST UDP application tenant.
	Tenant string `pulumi:"tenant"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer *FastUdpAppVirtualServer `pulumi:"virtualServer"`
	// Names of existing VLANs to allow.
	VlansAlloweds []string `pulumi:"vlansAlloweds"`
	// Names of existing VLANs to reject.
	VlansRejecteds []string `pulumi:"vlansRejecteds"`
}

// The set of arguments for constructing a FastUdpApp resource.
type FastUdpAppArgs struct {
	// Name of the FAST UDP application.
	Application pulumi.StringInput
	// Enables use of FastL4 profiles.
	EnableFastl4 pulumi.BoolPtrInput
	// Name of an existing BIG-IP UDP pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrInput
	// Name of an existing BIG-IP pool.
	ExistingPool pulumi.StringPtrInput
	// Name of an existing BIG-IP FastL4 or UDP profile.
	ExistingProfile pulumi.StringPtrInput
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrInput
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence pulumi.StringPtrInput
	// Irules to attach to Virtual Server.
	Irules pulumi.StringArrayInput
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrInput
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastUdpAppMonitorPtrInput
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile pulumi.StringPtrInput
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType pulumi.StringPtrInput
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastUdpAppPoolMemberArrayInput
	// Existing security log profiles to enable.
	SecurityLogProfiles pulumi.StringArrayInput
	// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
	SlowRampTime pulumi.IntPtrInput
	// List of address to be used for FAST-Generated SNAT Pool.
	SnatPoolAddresses pulumi.StringArrayInput
	// Name of the FAST UDP application tenant.
	Tenant pulumi.StringInput
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastUdpAppVirtualServerPtrInput
	// Names of existing VLANs to allow.
	VlansAlloweds pulumi.StringArrayInput
	// Names of existing VLANs to reject.
	VlansRejecteds pulumi.StringArrayInput
}

func (FastUdpAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fastUdpAppArgs)(nil)).Elem()
}

type FastUdpAppInput interface {
	pulumi.Input

	ToFastUdpAppOutput() FastUdpAppOutput
	ToFastUdpAppOutputWithContext(ctx context.Context) FastUdpAppOutput
}

func (*FastUdpApp) ElementType() reflect.Type {
	return reflect.TypeOf((**FastUdpApp)(nil)).Elem()
}

func (i *FastUdpApp) ToFastUdpAppOutput() FastUdpAppOutput {
	return i.ToFastUdpAppOutputWithContext(context.Background())
}

func (i *FastUdpApp) ToFastUdpAppOutputWithContext(ctx context.Context) FastUdpAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastUdpAppOutput)
}

// FastUdpAppArrayInput is an input type that accepts FastUdpAppArray and FastUdpAppArrayOutput values.
// You can construct a concrete instance of `FastUdpAppArrayInput` via:
//
//	FastUdpAppArray{ FastUdpAppArgs{...} }
type FastUdpAppArrayInput interface {
	pulumi.Input

	ToFastUdpAppArrayOutput() FastUdpAppArrayOutput
	ToFastUdpAppArrayOutputWithContext(context.Context) FastUdpAppArrayOutput
}

type FastUdpAppArray []FastUdpAppInput

func (FastUdpAppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastUdpApp)(nil)).Elem()
}

func (i FastUdpAppArray) ToFastUdpAppArrayOutput() FastUdpAppArrayOutput {
	return i.ToFastUdpAppArrayOutputWithContext(context.Background())
}

func (i FastUdpAppArray) ToFastUdpAppArrayOutputWithContext(ctx context.Context) FastUdpAppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastUdpAppArrayOutput)
}

// FastUdpAppMapInput is an input type that accepts FastUdpAppMap and FastUdpAppMapOutput values.
// You can construct a concrete instance of `FastUdpAppMapInput` via:
//
//	FastUdpAppMap{ "key": FastUdpAppArgs{...} }
type FastUdpAppMapInput interface {
	pulumi.Input

	ToFastUdpAppMapOutput() FastUdpAppMapOutput
	ToFastUdpAppMapOutputWithContext(context.Context) FastUdpAppMapOutput
}

type FastUdpAppMap map[string]FastUdpAppInput

func (FastUdpAppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastUdpApp)(nil)).Elem()
}

func (i FastUdpAppMap) ToFastUdpAppMapOutput() FastUdpAppMapOutput {
	return i.ToFastUdpAppMapOutputWithContext(context.Background())
}

func (i FastUdpAppMap) ToFastUdpAppMapOutputWithContext(ctx context.Context) FastUdpAppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastUdpAppMapOutput)
}

type FastUdpAppOutput struct{ *pulumi.OutputState }

func (FastUdpAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FastUdpApp)(nil)).Elem()
}

func (o FastUdpAppOutput) ToFastUdpAppOutput() FastUdpAppOutput {
	return o
}

func (o FastUdpAppOutput) ToFastUdpAppOutputWithContext(ctx context.Context) FastUdpAppOutput {
	return o
}

// Name of the FAST UDP application.
func (o FastUdpAppOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// Enables use of FastL4 profiles.
func (o FastUdpAppOutput) EnableFastl4() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.BoolPtrOutput { return v.EnableFastl4 }).(pulumi.BoolPtrOutput)
}

// Name of an existing BIG-IP UDP pool monitor. Monitors are used to determine the health of the application on each server.
func (o FastUdpAppOutput) ExistingMonitor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.ExistingMonitor }).(pulumi.StringPtrOutput)
}

// Name of an existing BIG-IP pool.
func (o FastUdpAppOutput) ExistingPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.ExistingPool }).(pulumi.StringPtrOutput)
}

// Name of an existing BIG-IP FastL4 or UDP profile.
func (o FastUdpAppOutput) ExistingProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.ExistingProfile }).(pulumi.StringPtrOutput)
}

// Name of an existing BIG-IP SNAT pool.
func (o FastUdpAppOutput) ExistingSnatPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.ExistingSnatPool }).(pulumi.StringPtrOutput)
}

// Type of fallback persistence record to be created for each new client connection.
func (o FastUdpAppOutput) FallbackPersistence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.FallbackPersistence }).(pulumi.StringPtrOutput)
}

// Json payload for FAST UDP application.
func (o FastUdpAppOutput) FastUdpJson() pulumi.StringOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringOutput { return v.FastUdpJson }).(pulumi.StringOutput)
}

// Irules to attach to Virtual Server.
func (o FastUdpAppOutput) Irules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringArrayOutput { return v.Irules }).(pulumi.StringArrayOutput)
}

// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
func (o FastUdpAppOutput) LoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.LoadBalancingMode }).(pulumi.StringPtrOutput)
}

// `monitor` block takes input for FAST-Generated Pool Monitor.
// See Pool Monitor below for more details.
func (o FastUdpAppOutput) Monitor() FastUdpAppMonitorPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) FastUdpAppMonitorPtrOutput { return v.Monitor }).(FastUdpAppMonitorPtrOutput)
}

// Name of an existing BIG-IP persistence profile to be used.
func (o FastUdpAppOutput) PersistenceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.PersistenceProfile }).(pulumi.StringPtrOutput)
}

// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
func (o FastUdpAppOutput) PersistenceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringPtrOutput { return v.PersistenceType }).(pulumi.StringPtrOutput)
}

// `poolMembers` block takes input for FAST-Generated Pool.
// See Pool Members below for more details.
func (o FastUdpAppOutput) PoolMembers() FastUdpAppPoolMemberArrayOutput {
	return o.ApplyT(func(v *FastUdpApp) FastUdpAppPoolMemberArrayOutput { return v.PoolMembers }).(FastUdpAppPoolMemberArrayOutput)
}

// Existing security log profiles to enable.
func (o FastUdpAppOutput) SecurityLogProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringArrayOutput { return v.SecurityLogProfiles }).(pulumi.StringArrayOutput)
}

// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
func (o FastUdpAppOutput) SlowRampTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.IntPtrOutput { return v.SlowRampTime }).(pulumi.IntPtrOutput)
}

// List of address to be used for FAST-Generated SNAT Pool.
func (o FastUdpAppOutput) SnatPoolAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringArrayOutput { return v.SnatPoolAddresses }).(pulumi.StringArrayOutput)
}

// Name of the FAST UDP application tenant.
func (o FastUdpAppOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
// See virtual server below for more details.
func (o FastUdpAppOutput) VirtualServer() FastUdpAppVirtualServerPtrOutput {
	return o.ApplyT(func(v *FastUdpApp) FastUdpAppVirtualServerPtrOutput { return v.VirtualServer }).(FastUdpAppVirtualServerPtrOutput)
}

// Names of existing VLANs to allow.
func (o FastUdpAppOutput) VlansAlloweds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringArrayOutput { return v.VlansAlloweds }).(pulumi.StringArrayOutput)
}

// Names of existing VLANs to reject.
func (o FastUdpAppOutput) VlansRejecteds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastUdpApp) pulumi.StringArrayOutput { return v.VlansRejecteds }).(pulumi.StringArrayOutput)
}

type FastUdpAppArrayOutput struct{ *pulumi.OutputState }

func (FastUdpAppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastUdpApp)(nil)).Elem()
}

func (o FastUdpAppArrayOutput) ToFastUdpAppArrayOutput() FastUdpAppArrayOutput {
	return o
}

func (o FastUdpAppArrayOutput) ToFastUdpAppArrayOutputWithContext(ctx context.Context) FastUdpAppArrayOutput {
	return o
}

func (o FastUdpAppArrayOutput) Index(i pulumi.IntInput) FastUdpAppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FastUdpApp {
		return vs[0].([]*FastUdpApp)[vs[1].(int)]
	}).(FastUdpAppOutput)
}

type FastUdpAppMapOutput struct{ *pulumi.OutputState }

func (FastUdpAppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastUdpApp)(nil)).Elem()
}

func (o FastUdpAppMapOutput) ToFastUdpAppMapOutput() FastUdpAppMapOutput {
	return o
}

func (o FastUdpAppMapOutput) ToFastUdpAppMapOutputWithContext(ctx context.Context) FastUdpAppMapOutput {
	return o
}

func (o FastUdpAppMapOutput) MapIndex(k pulumi.StringInput) FastUdpAppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FastUdpApp {
		return vs[0].(map[string]*FastUdpApp)[vs[1].(string)]
	}).(FastUdpAppOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FastUdpAppInput)(nil)).Elem(), &FastUdpApp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastUdpAppArrayInput)(nil)).Elem(), FastUdpAppArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastUdpAppMapInput)(nil)).Elem(), FastUdpAppMap{})
	pulumi.RegisterOutputType(FastUdpAppOutput{})
	pulumi.RegisterOutputType(FastUdpAppArrayOutput{})
	pulumi.RegisterOutputType(FastUdpAppMapOutput{})
}
