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

// `FastHttpsApp` This resource will create and manage FAST HTTPS applications on BIG-IP
//
// [FAST documentation](https://clouddocs.f5.com/products/extensions/f5-appsvcs-templates/latest/)
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
//			_, err := f5bigip.NewFastHttpsApp(ctx, "fast_https_app", &f5bigip.FastHttpsAppArgs{
//				Tenant:      pulumi.String("fasthttpstenant"),
//				Application: pulumi.String("fasthttpsapp"),
//				VirtualServer: &f5bigip.FastHttpsAppVirtualServerArgs{
//					Ip:   pulumi.String("10.30.40.44"),
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
//
// ### With Service Discovery
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip"
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/fast"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			TC3, err := fast.GetAzureServiceDiscovery(ctx, &fast.GetAzureServiceDiscoveryArgs{
//				ResourceGroup:  "testazurerg",
//				SubscriptionId: "testazuresid",
//				TagKey:         pulumi.StringRef("testazuretag"),
//				TagValue:       pulumi.StringRef("testazurevalue"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			TC3GetGceServiceDiscovery, err := fast.GetGceServiceDiscovery(ctx, &fast.GetGceServiceDiscoveryArgs{
//				TagKey:   "testgcetag",
//				TagValue: "testgcevalue",
//				Region:   "testgceregion",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = f5bigip.NewFastHttpsApp(ctx, "fast_https_app", &f5bigip.FastHttpsAppArgs{
//				Tenant:      pulumi.String("fasthttpstenant"),
//				Application: pulumi.String("fasthttpsapp"),
//				VirtualServer: &f5bigip.FastHttpsAppVirtualServerArgs{
//					Ip:   pulumi.String("10.30.40.44"),
//					Port: pulumi.Int(443),
//				},
//				PoolMembers: f5bigip.FastHttpsAppPoolMemberArray{
//					&f5bigip.FastHttpsAppPoolMemberArgs{
//						Addresses: pulumi.StringArray{
//							pulumi.String("10.11.40.120"),
//							pulumi.String("10.11.30.121"),
//							pulumi.String("10.11.30.122"),
//						},
//						Port: pulumi.Int(80),
//					},
//				},
//				ServiceDiscoveries: pulumi.StringArray{
//					pulumi.String(TC3GetGceServiceDiscovery.GceSdJson),
//					pulumi.String(TC3.AzureSdJson),
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
type FastHttpsApp struct {
	pulumi.CustomResourceState

	// Name of the FAST HTTPS application.
	Application pulumi.StringOutput `pulumi:"application"`
	// List of LTM Policies to be applied FAST HTTPS Application.
	EndpointLtmPolicies pulumi.StringArrayOutput `pulumi:"endpointLtmPolicies"`
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrOutput `pulumi:"existingMonitor"`
	// Name of an existing BIG-IP pool.
	ExistingPool pulumi.StringPtrOutput `pulumi:"existingPool"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrOutput `pulumi:"existingSnatPool"`
	// Name of an existing TLS client profile.
	ExistingTlsClientProfile pulumi.StringPtrOutput `pulumi:"existingTlsClientProfile"`
	// Name of an existing TLS server profile.
	ExistingTlsServerProfile pulumi.StringPtrOutput `pulumi:"existingTlsServerProfile"`
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy pulumi.StringPtrOutput `pulumi:"existingWafSecurityPolicy"`
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence pulumi.StringPtrOutput `pulumi:"fallbackPersistence"`
	// Json payload for FAST HTTPS application.
	FastHttpsJson pulumi.StringOutput `pulumi:"fastHttpsJson"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrOutput `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastHttpsAppMonitorPtrOutput `pulumi:"monitor"`
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile pulumi.StringPtrOutput `pulumi:"persistenceProfile"`
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType pulumi.StringPtrOutput `pulumi:"persistenceType"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastHttpsAppPoolMemberArrayOutput `pulumi:"poolMembers"`
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
	// `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
	// See TLS Client Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsClientProfile` or `tlsClientProfile` used for encrypt server-side connections.
	TlsClientProfile FastHttpsAppTlsClientProfilePtrOutput `pulumi:"tlsClientProfile"`
	// `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
	// See TLS Server Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsServerProfile` or `tlsServerProfile` used for decrypt client-side connections.
	TlsServerProfile FastHttpsAppTlsServerProfilePtrOutput `pulumi:"tlsServerProfile"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastHttpsAppVirtualServerPtrOutput `pulumi:"virtualServer"`
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy FastHttpsAppWafSecurityPolicyPtrOutput `pulumi:"wafSecurityPolicy"`
}

// NewFastHttpsApp registers a new resource with the given unique name, arguments, and options.
func NewFastHttpsApp(ctx *pulumi.Context,
	name string, args *FastHttpsAppArgs, opts ...pulumi.ResourceOption) (*FastHttpsApp, error) {
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
	var resource FastHttpsApp
	err := ctx.RegisterResource("f5bigip:index/fastHttpsApp:FastHttpsApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFastHttpsApp gets an existing FastHttpsApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFastHttpsApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FastHttpsAppState, opts ...pulumi.ResourceOption) (*FastHttpsApp, error) {
	var resource FastHttpsApp
	err := ctx.ReadResource("f5bigip:index/fastHttpsApp:FastHttpsApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FastHttpsApp resources.
type fastHttpsAppState struct {
	// Name of the FAST HTTPS application.
	Application *string `pulumi:"application"`
	// List of LTM Policies to be applied FAST HTTPS Application.
	EndpointLtmPolicies []string `pulumi:"endpointLtmPolicies"`
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor *string `pulumi:"existingMonitor"`
	// Name of an existing BIG-IP pool.
	ExistingPool *string `pulumi:"existingPool"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool *string `pulumi:"existingSnatPool"`
	// Name of an existing TLS client profile.
	ExistingTlsClientProfile *string `pulumi:"existingTlsClientProfile"`
	// Name of an existing TLS server profile.
	ExistingTlsServerProfile *string `pulumi:"existingTlsServerProfile"`
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy *string `pulumi:"existingWafSecurityPolicy"`
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence *string `pulumi:"fallbackPersistence"`
	// Json payload for FAST HTTPS application.
	FastHttpsJson *string `pulumi:"fastHttpsJson"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor *FastHttpsAppMonitor `pulumi:"monitor"`
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile *string `pulumi:"persistenceProfile"`
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType *string `pulumi:"persistenceType"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers []FastHttpsAppPoolMember `pulumi:"poolMembers"`
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
	// `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
	// See TLS Client Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsClientProfile` or `tlsClientProfile` used for encrypt server-side connections.
	TlsClientProfile *FastHttpsAppTlsClientProfile `pulumi:"tlsClientProfile"`
	// `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
	// See TLS Server Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsServerProfile` or `tlsServerProfile` used for decrypt client-side connections.
	TlsServerProfile *FastHttpsAppTlsServerProfile `pulumi:"tlsServerProfile"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer *FastHttpsAppVirtualServer `pulumi:"virtualServer"`
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy *FastHttpsAppWafSecurityPolicy `pulumi:"wafSecurityPolicy"`
}

type FastHttpsAppState struct {
	// Name of the FAST HTTPS application.
	Application pulumi.StringPtrInput
	// List of LTM Policies to be applied FAST HTTPS Application.
	EndpointLtmPolicies pulumi.StringArrayInput
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrInput
	// Name of an existing BIG-IP pool.
	ExistingPool pulumi.StringPtrInput
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrInput
	// Name of an existing TLS client profile.
	ExistingTlsClientProfile pulumi.StringPtrInput
	// Name of an existing TLS server profile.
	ExistingTlsServerProfile pulumi.StringPtrInput
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy pulumi.StringPtrInput
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence pulumi.StringPtrInput
	// Json payload for FAST HTTPS application.
	FastHttpsJson pulumi.StringPtrInput
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrInput
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastHttpsAppMonitorPtrInput
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile pulumi.StringPtrInput
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType pulumi.StringPtrInput
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastHttpsAppPoolMemberArrayInput
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
	// `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
	// See TLS Client Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsClientProfile` or `tlsClientProfile` used for encrypt server-side connections.
	TlsClientProfile FastHttpsAppTlsClientProfilePtrInput
	// `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
	// See TLS Server Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsServerProfile` or `tlsServerProfile` used for decrypt client-side connections.
	TlsServerProfile FastHttpsAppTlsServerProfilePtrInput
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastHttpsAppVirtualServerPtrInput
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy FastHttpsAppWafSecurityPolicyPtrInput
}

func (FastHttpsAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*fastHttpsAppState)(nil)).Elem()
}

type fastHttpsAppArgs struct {
	// Name of the FAST HTTPS application.
	Application string `pulumi:"application"`
	// List of LTM Policies to be applied FAST HTTPS Application.
	EndpointLtmPolicies []string `pulumi:"endpointLtmPolicies"`
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor *string `pulumi:"existingMonitor"`
	// Name of an existing BIG-IP pool.
	ExistingPool *string `pulumi:"existingPool"`
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool *string `pulumi:"existingSnatPool"`
	// Name of an existing TLS client profile.
	ExistingTlsClientProfile *string `pulumi:"existingTlsClientProfile"`
	// Name of an existing TLS server profile.
	ExistingTlsServerProfile *string `pulumi:"existingTlsServerProfile"`
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy *string `pulumi:"existingWafSecurityPolicy"`
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence *string `pulumi:"fallbackPersistence"`
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode *string `pulumi:"loadBalancingMode"`
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor *FastHttpsAppMonitor `pulumi:"monitor"`
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile *string `pulumi:"persistenceProfile"`
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType *string `pulumi:"persistenceType"`
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers []FastHttpsAppPoolMember `pulumi:"poolMembers"`
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
	// `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
	// See TLS Client Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsClientProfile` or `tlsClientProfile` used for encrypt server-side connections.
	TlsClientProfile *FastHttpsAppTlsClientProfile `pulumi:"tlsClientProfile"`
	// `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
	// See TLS Server Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsServerProfile` or `tlsServerProfile` used for decrypt client-side connections.
	TlsServerProfile *FastHttpsAppTlsServerProfile `pulumi:"tlsServerProfile"`
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer *FastHttpsAppVirtualServer `pulumi:"virtualServer"`
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy *FastHttpsAppWafSecurityPolicy `pulumi:"wafSecurityPolicy"`
}

// The set of arguments for constructing a FastHttpsApp resource.
type FastHttpsAppArgs struct {
	// Name of the FAST HTTPS application.
	Application pulumi.StringInput
	// List of LTM Policies to be applied FAST HTTPS Application.
	EndpointLtmPolicies pulumi.StringArrayInput
	// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
	ExistingMonitor pulumi.StringPtrInput
	// Name of an existing BIG-IP pool.
	ExistingPool pulumi.StringPtrInput
	// Name of an existing BIG-IP SNAT pool.
	ExistingSnatPool pulumi.StringPtrInput
	// Name of an existing TLS client profile.
	ExistingTlsClientProfile pulumi.StringPtrInput
	// Name of an existing TLS server profile.
	ExistingTlsServerProfile pulumi.StringPtrInput
	// Name of an existing WAF Security policy.
	ExistingWafSecurityPolicy pulumi.StringPtrInput
	// Type of fallback persistence record to be created for each new client connection.
	FallbackPersistence pulumi.StringPtrInput
	// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
	LoadBalancingMode pulumi.StringPtrInput
	// `monitor` block takes input for FAST-Generated Pool Monitor.
	// See Pool Monitor below for more details.
	Monitor FastHttpsAppMonitorPtrInput
	// Name of an existing BIG-IP persistence profile to be used.
	PersistenceProfile pulumi.StringPtrInput
	// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
	PersistenceType pulumi.StringPtrInput
	// `poolMembers` block takes input for FAST-Generated Pool.
	// See Pool Members below for more details.
	PoolMembers FastHttpsAppPoolMemberArrayInput
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
	// `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
	// See TLS Client Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsClientProfile` or `tlsClientProfile` used for encrypt server-side connections.
	TlsClientProfile FastHttpsAppTlsClientProfilePtrInput
	// `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
	// See TLS Server Profile below for more details.
	//
	// > **NOTE** Profile provided by `existingTlsServerProfile` or `tlsServerProfile` used for decrypt client-side connections.
	TlsServerProfile FastHttpsAppTlsServerProfilePtrInput
	// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
	// See virtual server below for more details.
	VirtualServer FastHttpsAppVirtualServerPtrInput
	// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
	// See WAF Security Policy below for more details.
	WafSecurityPolicy FastHttpsAppWafSecurityPolicyPtrInput
}

func (FastHttpsAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fastHttpsAppArgs)(nil)).Elem()
}

type FastHttpsAppInput interface {
	pulumi.Input

	ToFastHttpsAppOutput() FastHttpsAppOutput
	ToFastHttpsAppOutputWithContext(ctx context.Context) FastHttpsAppOutput
}

func (*FastHttpsApp) ElementType() reflect.Type {
	return reflect.TypeOf((**FastHttpsApp)(nil)).Elem()
}

func (i *FastHttpsApp) ToFastHttpsAppOutput() FastHttpsAppOutput {
	return i.ToFastHttpsAppOutputWithContext(context.Background())
}

func (i *FastHttpsApp) ToFastHttpsAppOutputWithContext(ctx context.Context) FastHttpsAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastHttpsAppOutput)
}

// FastHttpsAppArrayInput is an input type that accepts FastHttpsAppArray and FastHttpsAppArrayOutput values.
// You can construct a concrete instance of `FastHttpsAppArrayInput` via:
//
//	FastHttpsAppArray{ FastHttpsAppArgs{...} }
type FastHttpsAppArrayInput interface {
	pulumi.Input

	ToFastHttpsAppArrayOutput() FastHttpsAppArrayOutput
	ToFastHttpsAppArrayOutputWithContext(context.Context) FastHttpsAppArrayOutput
}

type FastHttpsAppArray []FastHttpsAppInput

func (FastHttpsAppArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastHttpsApp)(nil)).Elem()
}

func (i FastHttpsAppArray) ToFastHttpsAppArrayOutput() FastHttpsAppArrayOutput {
	return i.ToFastHttpsAppArrayOutputWithContext(context.Background())
}

func (i FastHttpsAppArray) ToFastHttpsAppArrayOutputWithContext(ctx context.Context) FastHttpsAppArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastHttpsAppArrayOutput)
}

// FastHttpsAppMapInput is an input type that accepts FastHttpsAppMap and FastHttpsAppMapOutput values.
// You can construct a concrete instance of `FastHttpsAppMapInput` via:
//
//	FastHttpsAppMap{ "key": FastHttpsAppArgs{...} }
type FastHttpsAppMapInput interface {
	pulumi.Input

	ToFastHttpsAppMapOutput() FastHttpsAppMapOutput
	ToFastHttpsAppMapOutputWithContext(context.Context) FastHttpsAppMapOutput
}

type FastHttpsAppMap map[string]FastHttpsAppInput

func (FastHttpsAppMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastHttpsApp)(nil)).Elem()
}

func (i FastHttpsAppMap) ToFastHttpsAppMapOutput() FastHttpsAppMapOutput {
	return i.ToFastHttpsAppMapOutputWithContext(context.Background())
}

func (i FastHttpsAppMap) ToFastHttpsAppMapOutputWithContext(ctx context.Context) FastHttpsAppMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FastHttpsAppMapOutput)
}

type FastHttpsAppOutput struct{ *pulumi.OutputState }

func (FastHttpsAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FastHttpsApp)(nil)).Elem()
}

func (o FastHttpsAppOutput) ToFastHttpsAppOutput() FastHttpsAppOutput {
	return o
}

func (o FastHttpsAppOutput) ToFastHttpsAppOutputWithContext(ctx context.Context) FastHttpsAppOutput {
	return o
}

// Name of the FAST HTTPS application.
func (o FastHttpsAppOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// List of LTM Policies to be applied FAST HTTPS Application.
func (o FastHttpsAppOutput) EndpointLtmPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringArrayOutput { return v.EndpointLtmPolicies }).(pulumi.StringArrayOutput)
}

// Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
func (o FastHttpsAppOutput) ExistingMonitor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.ExistingMonitor }).(pulumi.StringPtrOutput)
}

// Name of an existing BIG-IP pool.
func (o FastHttpsAppOutput) ExistingPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.ExistingPool }).(pulumi.StringPtrOutput)
}

// Name of an existing BIG-IP SNAT pool.
func (o FastHttpsAppOutput) ExistingSnatPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.ExistingSnatPool }).(pulumi.StringPtrOutput)
}

// Name of an existing TLS client profile.
func (o FastHttpsAppOutput) ExistingTlsClientProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.ExistingTlsClientProfile }).(pulumi.StringPtrOutput)
}

// Name of an existing TLS server profile.
func (o FastHttpsAppOutput) ExistingTlsServerProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.ExistingTlsServerProfile }).(pulumi.StringPtrOutput)
}

// Name of an existing WAF Security policy.
func (o FastHttpsAppOutput) ExistingWafSecurityPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.ExistingWafSecurityPolicy }).(pulumi.StringPtrOutput)
}

// Type of fallback persistence record to be created for each new client connection.
func (o FastHttpsAppOutput) FallbackPersistence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.FallbackPersistence }).(pulumi.StringPtrOutput)
}

// Json payload for FAST HTTPS application.
func (o FastHttpsAppOutput) FastHttpsJson() pulumi.StringOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringOutput { return v.FastHttpsJson }).(pulumi.StringOutput)
}

// A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
func (o FastHttpsAppOutput) LoadBalancingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.LoadBalancingMode }).(pulumi.StringPtrOutput)
}

// `monitor` block takes input for FAST-Generated Pool Monitor.
// See Pool Monitor below for more details.
func (o FastHttpsAppOutput) Monitor() FastHttpsAppMonitorPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) FastHttpsAppMonitorPtrOutput { return v.Monitor }).(FastHttpsAppMonitorPtrOutput)
}

// Name of an existing BIG-IP persistence profile to be used.
func (o FastHttpsAppOutput) PersistenceProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.PersistenceProfile }).(pulumi.StringPtrOutput)
}

// Type of persistence profile to be created. Using this option will enable use of FAST generated persistence profiles.
func (o FastHttpsAppOutput) PersistenceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringPtrOutput { return v.PersistenceType }).(pulumi.StringPtrOutput)
}

// `poolMembers` block takes input for FAST-Generated Pool.
// See Pool Members below for more details.
func (o FastHttpsAppOutput) PoolMembers() FastHttpsAppPoolMemberArrayOutput {
	return o.ApplyT(func(v *FastHttpsApp) FastHttpsAppPoolMemberArrayOutput { return v.PoolMembers }).(FastHttpsAppPoolMemberArrayOutput)
}

// List of security log profiles to be used for FAST application
func (o FastHttpsAppOutput) SecurityLogProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringArrayOutput { return v.SecurityLogProfiles }).(pulumi.StringArrayOutput)
}

// List of different cloud service discovery config provided as string, provided `serviceDiscovery` block to Automatically Discover Pool Members with Service Discovery on different clouds.
func (o FastHttpsAppOutput) ServiceDiscoveries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringArrayOutput { return v.ServiceDiscoveries }).(pulumi.StringArrayOutput)
}

// Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
func (o FastHttpsAppOutput) SlowRampTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.IntPtrOutput { return v.SlowRampTime }).(pulumi.IntPtrOutput)
}

// List of address to be used for FAST-Generated SNAT Pool.
func (o FastHttpsAppOutput) SnatPoolAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringArrayOutput { return v.SnatPoolAddresses }).(pulumi.StringArrayOutput)
}

// Name of the FAST HTTPS application tenant.
func (o FastHttpsAppOutput) Tenant() pulumi.StringOutput {
	return o.ApplyT(func(v *FastHttpsApp) pulumi.StringOutput { return v.Tenant }).(pulumi.StringOutput)
}

// `tlsClientProfile` block takes input for FAST-Generated TLS client Profile.
// See TLS Client Profile below for more details.
//
// > **NOTE** Profile provided by `existingTlsClientProfile` or `tlsClientProfile` used for encrypt server-side connections.
func (o FastHttpsAppOutput) TlsClientProfile() FastHttpsAppTlsClientProfilePtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) FastHttpsAppTlsClientProfilePtrOutput { return v.TlsClientProfile }).(FastHttpsAppTlsClientProfilePtrOutput)
}

// `tlsServerProfile` block takes input for FAST-Generated TLS Server Profile.
// See TLS Server Profile below for more details.
//
// > **NOTE** Profile provided by `existingTlsServerProfile` or `tlsServerProfile` used for decrypt client-side connections.
func (o FastHttpsAppOutput) TlsServerProfile() FastHttpsAppTlsServerProfilePtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) FastHttpsAppTlsServerProfilePtrOutput { return v.TlsServerProfile }).(FastHttpsAppTlsServerProfilePtrOutput)
}

// `virtualServer` block will provide `ip` and `port` options to be used for virtual server.
// See virtual server below for more details.
func (o FastHttpsAppOutput) VirtualServer() FastHttpsAppVirtualServerPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) FastHttpsAppVirtualServerPtrOutput { return v.VirtualServer }).(FastHttpsAppVirtualServerPtrOutput)
}

// `wafSecurityPolicy` block takes input for FAST-Generated WAF Security Policy.
// See WAF Security Policy below for more details.
func (o FastHttpsAppOutput) WafSecurityPolicy() FastHttpsAppWafSecurityPolicyPtrOutput {
	return o.ApplyT(func(v *FastHttpsApp) FastHttpsAppWafSecurityPolicyPtrOutput { return v.WafSecurityPolicy }).(FastHttpsAppWafSecurityPolicyPtrOutput)
}

type FastHttpsAppArrayOutput struct{ *pulumi.OutputState }

func (FastHttpsAppArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FastHttpsApp)(nil)).Elem()
}

func (o FastHttpsAppArrayOutput) ToFastHttpsAppArrayOutput() FastHttpsAppArrayOutput {
	return o
}

func (o FastHttpsAppArrayOutput) ToFastHttpsAppArrayOutputWithContext(ctx context.Context) FastHttpsAppArrayOutput {
	return o
}

func (o FastHttpsAppArrayOutput) Index(i pulumi.IntInput) FastHttpsAppOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FastHttpsApp {
		return vs[0].([]*FastHttpsApp)[vs[1].(int)]
	}).(FastHttpsAppOutput)
}

type FastHttpsAppMapOutput struct{ *pulumi.OutputState }

func (FastHttpsAppMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FastHttpsApp)(nil)).Elem()
}

func (o FastHttpsAppMapOutput) ToFastHttpsAppMapOutput() FastHttpsAppMapOutput {
	return o
}

func (o FastHttpsAppMapOutput) ToFastHttpsAppMapOutputWithContext(ctx context.Context) FastHttpsAppMapOutput {
	return o
}

func (o FastHttpsAppMapOutput) MapIndex(k pulumi.StringInput) FastHttpsAppOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FastHttpsApp {
		return vs[0].(map[string]*FastHttpsApp)[vs[1].(string)]
	}).(FastHttpsAppOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FastHttpsAppInput)(nil)).Elem(), &FastHttpsApp{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastHttpsAppArrayInput)(nil)).Elem(), FastHttpsAppArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FastHttpsAppMapInput)(nil)).Elem(), FastHttpsAppMap{})
	pulumi.RegisterOutputType(FastHttpsAppOutput{})
	pulumi.RegisterOutputType(FastHttpsAppArrayOutput{})
	pulumi.RegisterOutputType(FastHttpsAppMapOutput{})
}
