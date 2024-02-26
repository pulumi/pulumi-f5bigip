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

// `ltm.VirtualServer` Configures Virtual Server
//
// For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource (example: `/Common/test-virtualserver` ) or `partition + directory + name` of the resource (example: `/Common/test/test-virtualserver` ).
// When including directory in `fullpath` we have to make sure it is created in the given partition before using it.
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
//			_, err := ltm.NewVirtualServer(ctx, "http", &ltm.VirtualServerArgs{
//				Name:        pulumi.String("/Common/terraform_vs_http"),
//				Destination: pulumi.String("10.12.12.12"),
//				Port:        pulumi.Int(80),
//				Pool:        pulumi.String("/Common/the-default-pool"),
//			})
//			if err != nil {
//				return err
//			}
//			// A Virtual server with SSL enabled
//			_, err = ltm.NewVirtualServer(ctx, "httpsVirtualServer", &ltm.VirtualServerArgs{
//				Name:        pulumi.String("/Common/terraform_vs_https"),
//				Destination: pulumi.Any(_var.Vip_ip),
//				Description: pulumi.String("VirtualServer-test"),
//				Port:        pulumi.Int(443),
//				Pool:        pulumi.Any(_var.Pool),
//				Profiles: pulumi.StringArray{
//					pulumi.String("/Common/tcp"),
//					pulumi.String("/Common/my-awesome-ssl-cert"),
//					pulumi.String("/Common/http"),
//				},
//				SourceAddressTranslation: pulumi.String("automap"),
//				TranslateAddress:         pulumi.String("enabled"),
//				TranslatePort:            pulumi.String("enabled"),
//			})
//			if err != nil {
//				return err
//			}
//			// A Virtual server with separate client and server profiles
//			_, err = ltm.NewVirtualServer(ctx, "httpsLtm/virtualServerVirtualServer", &ltm.VirtualServerArgs{
//				Name:        pulumi.String("/Common/terraform_vs_https"),
//				Destination: pulumi.String("10.255.255.254"),
//				Description: pulumi.String("VirtualServer-test"),
//				Port:        pulumi.Int(443),
//				ClientProfiles: pulumi.StringArray{
//					pulumi.String("/Common/clientssl"),
//				},
//				ServerProfiles: pulumi.StringArray{
//					pulumi.String("/Common/serverssl"),
//				},
//				SecurityLogProfiles: pulumi.StringArray{
//					pulumi.String("/Common/global-network"),
//				},
//				SourceAddressTranslation: pulumi.String("automap"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VirtualServer struct {
	pulumi.CustomResourceState

	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles            pulumi.StringArrayOutput `pulumi:"clientProfiles"`
	DefaultPersistenceProfile pulumi.StringOutput      `pulumi:"defaultPersistenceProfile"`
	// Description of Virtual server
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Destination IP
	Destination pulumi.StringPtrOutput `pulumi:"destination"`
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile pulumi.StringOutput `pulumi:"fallbackPersistenceProfile"`
	// Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
	FirewallEnforcedPolicy pulumi.StringOutput `pulumi:"firewallEnforcedPolicy"`
	// Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
	IpProtocol pulumi.StringPtrOutput `pulumi:"ipProtocol"`
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules pulumi.StringArrayOutput `pulumi:"irules"`
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask pulumi.StringOutput `pulumi:"mask"`
	// Name of the virtual server
	Name                       pulumi.StringOutput `pulumi:"name"`
	PerFlowRequestAccessPolicy pulumi.StringOutput `pulumi:"perFlowRequestAccessPolicy"`
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles pulumi.StringArrayOutput `pulumi:"persistenceProfiles"`
	// Specifies the policies for the virtual server.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Default pool name
	Pool pulumi.StringPtrOutput `pulumi:"pool"`
	// Listen port for the virtual server
	Port pulumi.IntOutput `pulumi:"port"`
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles pulumi.StringArrayOutput `pulumi:"profiles"`
	// Specifies the log profile applied to the virtual server.
	SecurityLogProfiles pulumi.StringArrayOutput `pulumi:"securityLogProfiles"`
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles pulumi.StringArrayOutput `pulumi:"serverProfiles"`
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
	Snatpool pulumi.StringOutput `pulumi:"snatpool"`
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source pulumi.StringOutput `pulumi:"source"`
	// Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
	SourceAddressTranslation pulumi.StringOutput `pulumi:"sourceAddressTranslation"`
	// Specifies whether the system preserves the source port of the connection. The default is `preserve`.
	SourcePort pulumi.StringOutput `pulumi:"sourcePort"`
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Specifies destination traffic matching information to which the virtual server sends traffic
	TrafficmatchingCriteria pulumi.StringOutput `pulumi:"trafficmatchingCriteria"`
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress pulumi.StringPtrOutput `pulumi:"translateAddress"`
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort pulumi.StringPtrOutput `pulumi:"translatePort"`
	// The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
	Vlans pulumi.StringArrayOutput `pulumi:"vlans"`
	// Enables the virtual server on the VLANs specified by the `vlans` option.
	// By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
	VlansEnabled pulumi.BoolPtrOutput `pulumi:"vlansEnabled"`
}

// NewVirtualServer registers a new resource with the given unique name, arguments, and options.
func NewVirtualServer(ctx *pulumi.Context,
	name string, args *VirtualServerArgs, opts ...pulumi.ResourceOption) (*VirtualServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualServer
	err := ctx.RegisterResource("f5bigip:ltm/virtualServer:VirtualServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualServer gets an existing VirtualServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualServerState, opts ...pulumi.ResourceOption) (*VirtualServer, error) {
	var resource VirtualServer
	err := ctx.ReadResource("f5bigip:ltm/virtualServer:VirtualServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualServer resources.
type virtualServerState struct {
	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles            []string `pulumi:"clientProfiles"`
	DefaultPersistenceProfile *string  `pulumi:"defaultPersistenceProfile"`
	// Description of Virtual server
	Description *string `pulumi:"description"`
	// Destination IP
	Destination *string `pulumi:"destination"`
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile *string `pulumi:"fallbackPersistenceProfile"`
	// Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
	FirewallEnforcedPolicy *string `pulumi:"firewallEnforcedPolicy"`
	// Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
	IpProtocol *string `pulumi:"ipProtocol"`
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules []string `pulumi:"irules"`
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask *string `pulumi:"mask"`
	// Name of the virtual server
	Name                       *string `pulumi:"name"`
	PerFlowRequestAccessPolicy *string `pulumi:"perFlowRequestAccessPolicy"`
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles []string `pulumi:"persistenceProfiles"`
	// Specifies the policies for the virtual server.
	Policies []string `pulumi:"policies"`
	// Default pool name
	Pool *string `pulumi:"pool"`
	// Listen port for the virtual server
	Port *int `pulumi:"port"`
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles []string `pulumi:"profiles"`
	// Specifies the log profile applied to the virtual server.
	SecurityLogProfiles []string `pulumi:"securityLogProfiles"`
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles []string `pulumi:"serverProfiles"`
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
	Snatpool *string `pulumi:"snatpool"`
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source *string `pulumi:"source"`
	// Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
	SourceAddressTranslation *string `pulumi:"sourceAddressTranslation"`
	// Specifies whether the system preserves the source port of the connection. The default is `preserve`.
	SourcePort *string `pulumi:"sourcePort"`
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State *string `pulumi:"state"`
	// Specifies destination traffic matching information to which the virtual server sends traffic
	TrafficmatchingCriteria *string `pulumi:"trafficmatchingCriteria"`
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress *string `pulumi:"translateAddress"`
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort *string `pulumi:"translatePort"`
	// The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
	Vlans []string `pulumi:"vlans"`
	// Enables the virtual server on the VLANs specified by the `vlans` option.
	// By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
	VlansEnabled *bool `pulumi:"vlansEnabled"`
}

type VirtualServerState struct {
	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles            pulumi.StringArrayInput
	DefaultPersistenceProfile pulumi.StringPtrInput
	// Description of Virtual server
	Description pulumi.StringPtrInput
	// Destination IP
	Destination pulumi.StringPtrInput
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile pulumi.StringPtrInput
	// Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
	FirewallEnforcedPolicy pulumi.StringPtrInput
	// Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
	IpProtocol pulumi.StringPtrInput
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules pulumi.StringArrayInput
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask pulumi.StringPtrInput
	// Name of the virtual server
	Name                       pulumi.StringPtrInput
	PerFlowRequestAccessPolicy pulumi.StringPtrInput
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles pulumi.StringArrayInput
	// Specifies the policies for the virtual server.
	Policies pulumi.StringArrayInput
	// Default pool name
	Pool pulumi.StringPtrInput
	// Listen port for the virtual server
	Port pulumi.IntPtrInput
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles pulumi.StringArrayInput
	// Specifies the log profile applied to the virtual server.
	SecurityLogProfiles pulumi.StringArrayInput
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles pulumi.StringArrayInput
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
	Snatpool pulumi.StringPtrInput
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source pulumi.StringPtrInput
	// Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
	SourceAddressTranslation pulumi.StringPtrInput
	// Specifies whether the system preserves the source port of the connection. The default is `preserve`.
	SourcePort pulumi.StringPtrInput
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State pulumi.StringPtrInput
	// Specifies destination traffic matching information to which the virtual server sends traffic
	TrafficmatchingCriteria pulumi.StringPtrInput
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress pulumi.StringPtrInput
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort pulumi.StringPtrInput
	// The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
	Vlans pulumi.StringArrayInput
	// Enables the virtual server on the VLANs specified by the `vlans` option.
	// By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
	VlansEnabled pulumi.BoolPtrInput
}

func (VirtualServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServerState)(nil)).Elem()
}

type virtualServerArgs struct {
	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles            []string `pulumi:"clientProfiles"`
	DefaultPersistenceProfile *string  `pulumi:"defaultPersistenceProfile"`
	// Description of Virtual server
	Description *string `pulumi:"description"`
	// Destination IP
	Destination *string `pulumi:"destination"`
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile *string `pulumi:"fallbackPersistenceProfile"`
	// Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
	FirewallEnforcedPolicy *string `pulumi:"firewallEnforcedPolicy"`
	// Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
	IpProtocol *string `pulumi:"ipProtocol"`
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules []string `pulumi:"irules"`
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask *string `pulumi:"mask"`
	// Name of the virtual server
	Name                       string  `pulumi:"name"`
	PerFlowRequestAccessPolicy *string `pulumi:"perFlowRequestAccessPolicy"`
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles []string `pulumi:"persistenceProfiles"`
	// Specifies the policies for the virtual server.
	Policies []string `pulumi:"policies"`
	// Default pool name
	Pool *string `pulumi:"pool"`
	// Listen port for the virtual server
	Port *int `pulumi:"port"`
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles []string `pulumi:"profiles"`
	// Specifies the log profile applied to the virtual server.
	SecurityLogProfiles []string `pulumi:"securityLogProfiles"`
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles []string `pulumi:"serverProfiles"`
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
	Snatpool *string `pulumi:"snatpool"`
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source *string `pulumi:"source"`
	// Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
	SourceAddressTranslation *string `pulumi:"sourceAddressTranslation"`
	// Specifies whether the system preserves the source port of the connection. The default is `preserve`.
	SourcePort *string `pulumi:"sourcePort"`
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State *string `pulumi:"state"`
	// Specifies destination traffic matching information to which the virtual server sends traffic
	TrafficmatchingCriteria *string `pulumi:"trafficmatchingCriteria"`
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress *string `pulumi:"translateAddress"`
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort *string `pulumi:"translatePort"`
	// The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
	Vlans []string `pulumi:"vlans"`
	// Enables the virtual server on the VLANs specified by the `vlans` option.
	// By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
	VlansEnabled *bool `pulumi:"vlansEnabled"`
}

// The set of arguments for constructing a VirtualServer resource.
type VirtualServerArgs struct {
	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles            pulumi.StringArrayInput
	DefaultPersistenceProfile pulumi.StringPtrInput
	// Description of Virtual server
	Description pulumi.StringPtrInput
	// Destination IP
	Destination pulumi.StringPtrInput
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile pulumi.StringPtrInput
	// Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
	FirewallEnforcedPolicy pulumi.StringPtrInput
	// Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
	IpProtocol pulumi.StringPtrInput
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules pulumi.StringArrayInput
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask pulumi.StringPtrInput
	// Name of the virtual server
	Name                       pulumi.StringInput
	PerFlowRequestAccessPolicy pulumi.StringPtrInput
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles pulumi.StringArrayInput
	// Specifies the policies for the virtual server.
	Policies pulumi.StringArrayInput
	// Default pool name
	Pool pulumi.StringPtrInput
	// Listen port for the virtual server
	Port pulumi.IntPtrInput
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles pulumi.StringArrayInput
	// Specifies the log profile applied to the virtual server.
	SecurityLogProfiles pulumi.StringArrayInput
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles pulumi.StringArrayInput
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
	Snatpool pulumi.StringPtrInput
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source pulumi.StringPtrInput
	// Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
	SourceAddressTranslation pulumi.StringPtrInput
	// Specifies whether the system preserves the source port of the connection. The default is `preserve`.
	SourcePort pulumi.StringPtrInput
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State pulumi.StringPtrInput
	// Specifies destination traffic matching information to which the virtual server sends traffic
	TrafficmatchingCriteria pulumi.StringPtrInput
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress pulumi.StringPtrInput
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort pulumi.StringPtrInput
	// The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
	Vlans pulumi.StringArrayInput
	// Enables the virtual server on the VLANs specified by the `vlans` option.
	// By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
	VlansEnabled pulumi.BoolPtrInput
}

func (VirtualServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServerArgs)(nil)).Elem()
}

type VirtualServerInput interface {
	pulumi.Input

	ToVirtualServerOutput() VirtualServerOutput
	ToVirtualServerOutputWithContext(ctx context.Context) VirtualServerOutput
}

func (*VirtualServer) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualServer)(nil)).Elem()
}

func (i *VirtualServer) ToVirtualServerOutput() VirtualServerOutput {
	return i.ToVirtualServerOutputWithContext(context.Background())
}

func (i *VirtualServer) ToVirtualServerOutputWithContext(ctx context.Context) VirtualServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServerOutput)
}

// VirtualServerArrayInput is an input type that accepts VirtualServerArray and VirtualServerArrayOutput values.
// You can construct a concrete instance of `VirtualServerArrayInput` via:
//
//	VirtualServerArray{ VirtualServerArgs{...} }
type VirtualServerArrayInput interface {
	pulumi.Input

	ToVirtualServerArrayOutput() VirtualServerArrayOutput
	ToVirtualServerArrayOutputWithContext(context.Context) VirtualServerArrayOutput
}

type VirtualServerArray []VirtualServerInput

func (VirtualServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualServer)(nil)).Elem()
}

func (i VirtualServerArray) ToVirtualServerArrayOutput() VirtualServerArrayOutput {
	return i.ToVirtualServerArrayOutputWithContext(context.Background())
}

func (i VirtualServerArray) ToVirtualServerArrayOutputWithContext(ctx context.Context) VirtualServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServerArrayOutput)
}

// VirtualServerMapInput is an input type that accepts VirtualServerMap and VirtualServerMapOutput values.
// You can construct a concrete instance of `VirtualServerMapInput` via:
//
//	VirtualServerMap{ "key": VirtualServerArgs{...} }
type VirtualServerMapInput interface {
	pulumi.Input

	ToVirtualServerMapOutput() VirtualServerMapOutput
	ToVirtualServerMapOutputWithContext(context.Context) VirtualServerMapOutput
}

type VirtualServerMap map[string]VirtualServerInput

func (VirtualServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualServer)(nil)).Elem()
}

func (i VirtualServerMap) ToVirtualServerMapOutput() VirtualServerMapOutput {
	return i.ToVirtualServerMapOutputWithContext(context.Background())
}

func (i VirtualServerMap) ToVirtualServerMapOutputWithContext(ctx context.Context) VirtualServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServerMapOutput)
}

type VirtualServerOutput struct{ *pulumi.OutputState }

func (VirtualServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualServer)(nil)).Elem()
}

func (o VirtualServerOutput) ToVirtualServerOutput() VirtualServerOutput {
	return o
}

func (o VirtualServerOutput) ToVirtualServerOutputWithContext(ctx context.Context) VirtualServerOutput {
	return o
}

// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
func (o VirtualServerOutput) ClientProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.ClientProfiles }).(pulumi.StringArrayOutput)
}

func (o VirtualServerOutput) DefaultPersistenceProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.DefaultPersistenceProfile }).(pulumi.StringOutput)
}

// Description of Virtual server
func (o VirtualServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Destination IP
func (o VirtualServerOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringPtrOutput { return v.Destination }).(pulumi.StringPtrOutput)
}

// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
func (o VirtualServerOutput) FallbackPersistenceProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.FallbackPersistenceProfile }).(pulumi.StringOutput)
}

// Applies the specified AFM policy to the virtual in an enforcing way,when creating a new virtual, if this parameter is not specified, the enforced is disabled.This should be in full path ex: `/Common/afm-test-policy`.
func (o VirtualServerOutput) FirewallEnforcedPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.FirewallEnforcedPolicy }).(pulumi.StringOutput)
}

// Specifies a network protocol name you want the system to use to direct traffic on this virtual server. The default is `tcp`. valid options are [`any`,`udp`,`tcp`]
func (o VirtualServerOutput) IpProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringPtrOutput { return v.IpProtocol }).(pulumi.StringPtrOutput)
}

// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
func (o VirtualServerOutput) Irules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.Irules }).(pulumi.StringArrayOutput)
}

// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
func (o VirtualServerOutput) Mask() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.Mask }).(pulumi.StringOutput)
}

// Name of the virtual server
func (o VirtualServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualServerOutput) PerFlowRequestAccessPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.PerFlowRequestAccessPolicy }).(pulumi.StringOutput)
}

// List of persistence profiles associated with the Virtual Server.
func (o VirtualServerOutput) PersistenceProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.PersistenceProfiles }).(pulumi.StringArrayOutput)
}

// Specifies the policies for the virtual server.
func (o VirtualServerOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// Default pool name
func (o VirtualServerOutput) Pool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringPtrOutput { return v.Pool }).(pulumi.StringPtrOutput)
}

// Listen port for the virtual server
func (o VirtualServerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
func (o VirtualServerOutput) Profiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.Profiles }).(pulumi.StringArrayOutput)
}

// Specifies the log profile applied to the virtual server.
func (o VirtualServerOutput) SecurityLogProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.SecurityLogProfiles }).(pulumi.StringArrayOutput)
}

// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
func (o VirtualServerOutput) ServerProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.ServerProfiles }).(pulumi.StringArrayOutput)
}

// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs.
func (o VirtualServerOutput) Snatpool() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.Snatpool }).(pulumi.StringOutput)
}

// Specifies an IP address or network from which the virtual server will accept traffic.
func (o VirtualServerOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// Can be either omitted for `none` or the values `automap` options : [`snat`,`automap`,`none`].
func (o VirtualServerOutput) SourceAddressTranslation() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.SourceAddressTranslation }).(pulumi.StringOutput)
}

// Specifies whether the system preserves the source port of the connection. The default is `preserve`.
func (o VirtualServerOutput) SourcePort() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.SourcePort }).(pulumi.StringOutput)
}

// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
func (o VirtualServerOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Specifies destination traffic matching information to which the virtual server sends traffic
func (o VirtualServerOutput) TrafficmatchingCriteria() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringOutput { return v.TrafficmatchingCriteria }).(pulumi.StringOutput)
}

// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
func (o VirtualServerOutput) TranslateAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringPtrOutput { return v.TranslateAddress }).(pulumi.StringPtrOutput)
}

// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
func (o VirtualServerOutput) TranslatePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringPtrOutput { return v.TranslatePort }).(pulumi.StringPtrOutput)
}

// The virtual server is enabled/disabled on this set of VLANs,enable/disabled will be desided by attribute `vlanEnabled`
func (o VirtualServerOutput) Vlans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.StringArrayOutput { return v.Vlans }).(pulumi.StringArrayOutput)
}

// Enables the virtual server on the VLANs specified by the `vlans` option.
// By default it is `false` i.e vlanDisabled on specified vlans, if we want enable virtual server on VLANs specified by `vlans`, mark this attribute to `true`.
func (o VirtualServerOutput) VlansEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualServer) pulumi.BoolPtrOutput { return v.VlansEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualServerArrayOutput struct{ *pulumi.OutputState }

func (VirtualServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualServer)(nil)).Elem()
}

func (o VirtualServerArrayOutput) ToVirtualServerArrayOutput() VirtualServerArrayOutput {
	return o
}

func (o VirtualServerArrayOutput) ToVirtualServerArrayOutputWithContext(ctx context.Context) VirtualServerArrayOutput {
	return o
}

func (o VirtualServerArrayOutput) Index(i pulumi.IntInput) VirtualServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualServer {
		return vs[0].([]*VirtualServer)[vs[1].(int)]
	}).(VirtualServerOutput)
}

type VirtualServerMapOutput struct{ *pulumi.OutputState }

func (VirtualServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualServer)(nil)).Elem()
}

func (o VirtualServerMapOutput) ToVirtualServerMapOutput() VirtualServerMapOutput {
	return o
}

func (o VirtualServerMapOutput) ToVirtualServerMapOutputWithContext(ctx context.Context) VirtualServerMapOutput {
	return o
}

func (o VirtualServerMapOutput) MapIndex(k pulumi.StringInput) VirtualServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualServer {
		return vs[0].(map[string]*VirtualServer)[vs[1].(string)]
	}).(VirtualServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualServerInput)(nil)).Elem(), &VirtualServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualServerArrayInput)(nil)).Elem(), VirtualServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualServerMapInput)(nil)).Elem(), VirtualServerMap{})
	pulumi.RegisterOutputType(VirtualServerOutput{})
	pulumi.RegisterOutputType(VirtualServerArrayOutput{})
	pulumi.RegisterOutputType(VirtualServerMapOutput{})
}
