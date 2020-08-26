// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.VirtualServer` Configures Virtual Server
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
// 		_, err := ltm.NewVirtualServer(ctx, "http", &ltm.VirtualServerArgs{
// 			Name:        pulumi.String("/Common/terraform_vs_http"),
// 			Destination: pulumi.String("10.12.12.12"),
// 			Port:        pulumi.Int(80),
// 			Pool:        pulumi.String("/Common/the-default-pool"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewVirtualServer(ctx, "httpsVirtualServer", &ltm.VirtualServerArgs{
// 			Name:        pulumi.String("/Common/terraform_vs_https"),
// 			Destination: pulumi.Any(_var.Vip_ip),
// 			Description: pulumi.String("VirtualServer-test"),
// 			Port:        pulumi.Int(443),
// 			Pool:        pulumi.Any(_var.Pool),
// 			Profiles: pulumi.StringArray{
// 				pulumi.String("/Common/tcp"),
// 				pulumi.String("/Common/my-awesome-ssl-cert"),
// 				pulumi.String("/Common/http"),
// 			},
// 			SourceAddressTranslation: pulumi.String("automap"),
// 			TranslateAddress:         pulumi.String("enabled"),
// 			TranslatePort:            pulumi.String("enabled"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ltm.NewVirtualServer(ctx, "httpsLtm_virtualServerVirtualServer", &ltm.VirtualServerArgs{
// 			Name:        pulumi.String("/Common/terraform_vs_https"),
// 			Destination: pulumi.String("10.255.255.254"),
// 			Description: pulumi.String("VirtualServer-test"),
// 			Port:        pulumi.Int(443),
// 			ClientProfiles: pulumi.StringArray{
// 				pulumi.String("/Common/clientssl"),
// 			},
// 			ServerProfiles: pulumi.StringArray{
// 				pulumi.String("/Common/serverssl"),
// 			},
// 			SourceAddressTranslation: pulumi.String("automap"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VirtualServer struct {
	pulumi.CustomResourceState

	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles            pulumi.StringArrayOutput `pulumi:"clientProfiles"`
	DefaultPersistenceProfile pulumi.StringPtrOutput   `pulumi:"defaultPersistenceProfile"`
	// Description of Virtual server
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Destination IP
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile pulumi.StringOutput `pulumi:"fallbackPersistenceProfile"`
	// Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules pulumi.StringArrayOutput `pulumi:"irules"`
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask pulumi.StringOutput `pulumi:"mask"`
	// Name of the virtual server
	Name pulumi.StringOutput `pulumi:"name"`
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles pulumi.StringArrayOutput `pulumi:"persistenceProfiles"`
	Policies            pulumi.StringArrayOutput `pulumi:"policies"`
	// Default pool name
	Pool pulumi.StringPtrOutput `pulumi:"pool"`
	// Listen port for the virtual server
	Port pulumi.IntOutput `pulumi:"port"`
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles pulumi.StringArrayOutput `pulumi:"profiles"`
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles pulumi.StringArrayOutput `pulumi:"serverProfiles"`
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
	Snatpool pulumi.StringOutput `pulumi:"snatpool"`
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source pulumi.StringOutput `pulumi:"source"`
	// Can be either omitted for none or the values automap or snat
	SourceAddressTranslation pulumi.StringOutput `pulumi:"sourceAddressTranslation"`
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress pulumi.StringOutput `pulumi:"translateAddress"`
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort pulumi.StringOutput `pulumi:"translatePort"`
	// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
	Vlans pulumi.StringArrayOutput `pulumi:"vlans"`
	// Enables the virtual server on the VLANs specified by the VLANs option.
	VlansEnabled pulumi.BoolOutput `pulumi:"vlansEnabled"`
}

// NewVirtualServer registers a new resource with the given unique name, arguments, and options.
func NewVirtualServer(ctx *pulumi.Context,
	name string, args *VirtualServerArgs, opts ...pulumi.ResourceOption) (*VirtualServer, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil {
		args = &VirtualServerArgs{}
	}
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
	// Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
	IpProtocol *string `pulumi:"ipProtocol"`
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules []string `pulumi:"irules"`
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask *string `pulumi:"mask"`
	// Name of the virtual server
	Name *string `pulumi:"name"`
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles []string `pulumi:"persistenceProfiles"`
	Policies            []string `pulumi:"policies"`
	// Default pool name
	Pool *string `pulumi:"pool"`
	// Listen port for the virtual server
	Port *int `pulumi:"port"`
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles []string `pulumi:"profiles"`
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles []string `pulumi:"serverProfiles"`
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
	Snatpool *string `pulumi:"snatpool"`
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source *string `pulumi:"source"`
	// Can be either omitted for none or the values automap or snat
	SourceAddressTranslation *string `pulumi:"sourceAddressTranslation"`
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State *string `pulumi:"state"`
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress *string `pulumi:"translateAddress"`
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort *string `pulumi:"translatePort"`
	// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
	Vlans []string `pulumi:"vlans"`
	// Enables the virtual server on the VLANs specified by the VLANs option.
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
	// Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
	IpProtocol pulumi.StringPtrInput
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules pulumi.StringArrayInput
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask pulumi.StringPtrInput
	// Name of the virtual server
	Name pulumi.StringPtrInput
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles pulumi.StringArrayInput
	Policies            pulumi.StringArrayInput
	// Default pool name
	Pool pulumi.StringPtrInput
	// Listen port for the virtual server
	Port pulumi.IntPtrInput
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles pulumi.StringArrayInput
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles pulumi.StringArrayInput
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
	Snatpool pulumi.StringPtrInput
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source pulumi.StringPtrInput
	// Can be either omitted for none or the values automap or snat
	SourceAddressTranslation pulumi.StringPtrInput
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State pulumi.StringPtrInput
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress pulumi.StringPtrInput
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort pulumi.StringPtrInput
	// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
	Vlans pulumi.StringArrayInput
	// Enables the virtual server on the VLANs specified by the VLANs option.
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
	Destination string `pulumi:"destination"`
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile *string `pulumi:"fallbackPersistenceProfile"`
	// Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
	IpProtocol *string `pulumi:"ipProtocol"`
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules []string `pulumi:"irules"`
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask *string `pulumi:"mask"`
	// Name of the virtual server
	Name string `pulumi:"name"`
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles []string `pulumi:"persistenceProfiles"`
	Policies            []string `pulumi:"policies"`
	// Default pool name
	Pool *string `pulumi:"pool"`
	// Listen port for the virtual server
	Port int `pulumi:"port"`
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles []string `pulumi:"profiles"`
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles []string `pulumi:"serverProfiles"`
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
	Snatpool *string `pulumi:"snatpool"`
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source *string `pulumi:"source"`
	// Can be either omitted for none or the values automap or snat
	SourceAddressTranslation *string `pulumi:"sourceAddressTranslation"`
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State *string `pulumi:"state"`
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress *string `pulumi:"translateAddress"`
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort *string `pulumi:"translatePort"`
	// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
	Vlans []string `pulumi:"vlans"`
	// Enables the virtual server on the VLANs specified by the VLANs option.
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
	Destination pulumi.StringInput
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile pulumi.StringPtrInput
	// Specify the IP protocol to use with the the virtual server (all, tcp, or udp are valid)
	IpProtocol pulumi.StringPtrInput
	// The iRules list you want run on this virtual server. iRules help automate the intercepting, processing, and routing of application traffic.
	Irules pulumi.StringArrayInput
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask pulumi.StringPtrInput
	// Name of the virtual server
	Name pulumi.StringInput
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles pulumi.StringArrayInput
	Policies            pulumi.StringArrayInput
	// Default pool name
	Pool pulumi.StringPtrInput
	// Listen port for the virtual server
	Port pulumi.IntInput
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles pulumi.StringArrayInput
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles pulumi.StringArrayInput
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
	Snatpool pulumi.StringPtrInput
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source pulumi.StringPtrInput
	// Can be either omitted for none or the values automap or snat
	SourceAddressTranslation pulumi.StringPtrInput
	// Specifies whether the virtual server and its resources are available for load balancing. The default is Enabled
	State pulumi.StringPtrInput
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress pulumi.StringPtrInput
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort pulumi.StringPtrInput
	// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
	Vlans pulumi.StringArrayInput
	// Enables the virtual server on the VLANs specified by the VLANs option.
	VlansEnabled pulumi.BoolPtrInput
}

func (VirtualServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServerArgs)(nil)).Elem()
}
