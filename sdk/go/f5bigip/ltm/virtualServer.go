// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `bigip_ltm_virtual_server` Configures Virtual Server
// 
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
type VirtualServer struct {
	s *pulumi.ResourceState
}

// NewVirtualServer registers a new resource with the given unique name, arguments, and options.
func NewVirtualServer(ctx *pulumi.Context,
	name string, args *VirtualServerArgs, opts ...pulumi.ResourceOpt) (*VirtualServer, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clientProfiles"] = nil
		inputs["destination"] = nil
		inputs["fallbackPersistenceProfile"] = nil
		inputs["ipProtocol"] = nil
		inputs["irules"] = nil
		inputs["mask"] = nil
		inputs["name"] = nil
		inputs["persistenceProfiles"] = nil
		inputs["policies"] = nil
		inputs["pool"] = nil
		inputs["port"] = nil
		inputs["profiles"] = nil
		inputs["serverProfiles"] = nil
		inputs["snatpool"] = nil
		inputs["source"] = nil
		inputs["sourceAddressTranslation"] = nil
		inputs["translateAddress"] = nil
		inputs["translatePort"] = nil
		inputs["vlans"] = nil
		inputs["vlansEnabled"] = nil
	} else {
		inputs["clientProfiles"] = args.ClientProfiles
		inputs["destination"] = args.Destination
		inputs["fallbackPersistenceProfile"] = args.FallbackPersistenceProfile
		inputs["ipProtocol"] = args.IpProtocol
		inputs["irules"] = args.Irules
		inputs["mask"] = args.Mask
		inputs["name"] = args.Name
		inputs["persistenceProfiles"] = args.PersistenceProfiles
		inputs["policies"] = args.Policies
		inputs["pool"] = args.Pool
		inputs["port"] = args.Port
		inputs["profiles"] = args.Profiles
		inputs["serverProfiles"] = args.ServerProfiles
		inputs["snatpool"] = args.Snatpool
		inputs["source"] = args.Source
		inputs["sourceAddressTranslation"] = args.SourceAddressTranslation
		inputs["translateAddress"] = args.TranslateAddress
		inputs["translatePort"] = args.TranslatePort
		inputs["vlans"] = args.Vlans
		inputs["vlansEnabled"] = args.VlansEnabled
	}
	s, err := ctx.RegisterResource("f5bigip:ltm/virtualServer:VirtualServer", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualServer{s: s}, nil
}

// GetVirtualServer gets an existing VirtualServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualServer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VirtualServerState, opts ...pulumi.ResourceOpt) (*VirtualServer, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["clientProfiles"] = state.ClientProfiles
		inputs["destination"] = state.Destination
		inputs["fallbackPersistenceProfile"] = state.FallbackPersistenceProfile
		inputs["ipProtocol"] = state.IpProtocol
		inputs["irules"] = state.Irules
		inputs["mask"] = state.Mask
		inputs["name"] = state.Name
		inputs["persistenceProfiles"] = state.PersistenceProfiles
		inputs["policies"] = state.Policies
		inputs["pool"] = state.Pool
		inputs["port"] = state.Port
		inputs["profiles"] = state.Profiles
		inputs["serverProfiles"] = state.ServerProfiles
		inputs["snatpool"] = state.Snatpool
		inputs["source"] = state.Source
		inputs["sourceAddressTranslation"] = state.SourceAddressTranslation
		inputs["translateAddress"] = state.TranslateAddress
		inputs["translatePort"] = state.TranslatePort
		inputs["vlans"] = state.Vlans
		inputs["vlansEnabled"] = state.VlansEnabled
	}
	s, err := ctx.ReadResource("f5bigip:ltm/virtualServer:VirtualServer", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VirtualServer{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VirtualServer) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VirtualServer) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
func (r *VirtualServer) ClientProfiles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["clientProfiles"])
}

// Destination IP
func (r *VirtualServer) Destination() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["destination"])
}

// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
func (r *VirtualServer) FallbackPersistenceProfile() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fallbackPersistenceProfile"])
}

// all, tcp, udp
func (r *VirtualServer) IpProtocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipProtocol"])
}

func (r *VirtualServer) Irules() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["irules"])
}

// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
func (r *VirtualServer) Mask() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["mask"])
}

// Name of the virtual server
func (r *VirtualServer) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// List of persistence profiles associated with the Virtual Server.
func (r *VirtualServer) PersistenceProfiles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["persistenceProfiles"])
}

func (r *VirtualServer) Policies() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policies"])
}

// Default pool name
func (r *VirtualServer) Pool() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["pool"])
}

// Listen port for the virtual server
func (r *VirtualServer) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
func (r *VirtualServer) Profiles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["profiles"])
}

// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
func (r *VirtualServer) ServerProfiles() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["serverProfiles"])
}

// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
func (r *VirtualServer) Snatpool() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["snatpool"])
}

// Specifies an IP address or network from which the virtual server will accept traffic.
func (r *VirtualServer) Source() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["source"])
}

// Can be either omitted for none or the values automap or snat
func (r *VirtualServer) SourceAddressTranslation() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceAddressTranslation"])
}

// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
func (r *VirtualServer) TranslateAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["translateAddress"])
}

// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
func (r *VirtualServer) TranslatePort() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["translatePort"])
}

// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
func (r *VirtualServer) Vlans() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["vlans"])
}

// Enables the virtual server on the VLANs specified by the VLANs option.
func (r *VirtualServer) VlansEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["vlansEnabled"])
}

// Input properties used for looking up and filtering VirtualServer resources.
type VirtualServerState struct {
	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles interface{}
	// Destination IP
	Destination interface{}
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile interface{}
	// all, tcp, udp
	IpProtocol interface{}
	Irules interface{}
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask interface{}
	// Name of the virtual server
	Name interface{}
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles interface{}
	Policies interface{}
	// Default pool name
	Pool interface{}
	// Listen port for the virtual server
	Port interface{}
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles interface{}
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles interface{}
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
	Snatpool interface{}
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source interface{}
	// Can be either omitted for none or the values automap or snat
	SourceAddressTranslation interface{}
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress interface{}
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort interface{}
	// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
	Vlans interface{}
	// Enables the virtual server on the VLANs specified by the VLANs option.
	VlansEnabled interface{}
}

// The set of arguments for constructing a VirtualServer resource.
type VirtualServerArgs struct {
	// List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
	ClientProfiles interface{}
	// Destination IP
	Destination interface{}
	// Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
	FallbackPersistenceProfile interface{}
	// all, tcp, udp
	IpProtocol interface{}
	Irules interface{}
	// Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
	Mask interface{}
	// Name of the virtual server
	Name interface{}
	// List of persistence profiles associated with the Virtual Server.
	PersistenceProfiles interface{}
	Policies interface{}
	// Default pool name
	Pool interface{}
	// Listen port for the virtual server
	Port interface{}
	// List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
	Profiles interface{}
	// List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
	ServerProfiles interface{}
	// Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
	Snatpool interface{}
	// Specifies an IP address or network from which the virtual server will accept traffic.
	Source interface{}
	// Can be either omitted for none or the values automap or snat
	SourceAddressTranslation interface{}
	// Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
	TranslateAddress interface{}
	// Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
	TranslatePort interface{}
	// The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
	Vlans interface{}
	// Enables the virtual server on the VLANs specified by the VLANs option.
	VlansEnabled interface{}
}
