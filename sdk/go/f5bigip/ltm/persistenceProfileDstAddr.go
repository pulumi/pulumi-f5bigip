// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures a cookie persistence profile
// 
// ## Reference
// 
// `name` - (Required) Name of the virtual address
// 
// `defaultsFrom` - (Optional) Specifies the existing profile from which the system imports settings for the new profile.
// 
// `matchAcrossPools` (Optional) (enabled or disabled) match across pools with given persistence record
// 
// `matchAcrossServices` (Optional) (enabled or disabled) match across services with given persistence record
// 
// `matchAcrossVirtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
// 
// `mirror` (Optional) (enabled or disabled) mirror persistence record
// 
// `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
// 
// `overrideConnLimit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_persistence_profile_dstaddr.html.markdown.
type PersistenceProfileDstAddr struct {
	s *pulumi.ResourceState
}

// NewPersistenceProfileDstAddr registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileDstAddr(ctx *pulumi.Context,
	name string, args *PersistenceProfileDstAddrArgs, opts ...pulumi.ResourceOpt) (*PersistenceProfileDstAddr, error) {
	if args == nil || args.DefaultsFrom == nil {
		return nil, errors.New("missing required argument 'DefaultsFrom'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["appService"] = nil
		inputs["defaultsFrom"] = nil
		inputs["hashAlgorithm"] = nil
		inputs["mask"] = nil
		inputs["matchAcrossPools"] = nil
		inputs["matchAcrossServices"] = nil
		inputs["matchAcrossVirtuals"] = nil
		inputs["mirror"] = nil
		inputs["name"] = nil
		inputs["overrideConnLimit"] = nil
		inputs["timeout"] = nil
	} else {
		inputs["appService"] = args.AppService
		inputs["defaultsFrom"] = args.DefaultsFrom
		inputs["hashAlgorithm"] = args.HashAlgorithm
		inputs["mask"] = args.Mask
		inputs["matchAcrossPools"] = args.MatchAcrossPools
		inputs["matchAcrossServices"] = args.MatchAcrossServices
		inputs["matchAcrossVirtuals"] = args.MatchAcrossVirtuals
		inputs["mirror"] = args.Mirror
		inputs["name"] = args.Name
		inputs["overrideConnLimit"] = args.OverrideConnLimit
		inputs["timeout"] = args.Timeout
	}
	s, err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PersistenceProfileDstAddr{s: s}, nil
}

// GetPersistenceProfileDstAddr gets an existing PersistenceProfileDstAddr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileDstAddr(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PersistenceProfileDstAddrState, opts ...pulumi.ResourceOpt) (*PersistenceProfileDstAddr, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appService"] = state.AppService
		inputs["defaultsFrom"] = state.DefaultsFrom
		inputs["hashAlgorithm"] = state.HashAlgorithm
		inputs["mask"] = state.Mask
		inputs["matchAcrossPools"] = state.MatchAcrossPools
		inputs["matchAcrossServices"] = state.MatchAcrossServices
		inputs["matchAcrossVirtuals"] = state.MatchAcrossVirtuals
		inputs["mirror"] = state.Mirror
		inputs["name"] = state.Name
		inputs["overrideConnLimit"] = state.OverrideConnLimit
		inputs["timeout"] = state.Timeout
	}
	s, err := ctx.ReadResource("f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PersistenceProfileDstAddr{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PersistenceProfileDstAddr) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PersistenceProfileDstAddr) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *PersistenceProfileDstAddr) AppService() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["appService"])
}

// Inherit defaults from parent profile
func (r *PersistenceProfileDstAddr) DefaultsFrom() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["defaultsFrom"])
}

// Specify the hash algorithm
func (r *PersistenceProfileDstAddr) HashAlgorithm() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["hashAlgorithm"])
}

// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
func (r *PersistenceProfileDstAddr) Mask() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mask"])
}

// To enable _ disable match across pools with given persistence record
func (r *PersistenceProfileDstAddr) MatchAcrossPools() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["matchAcrossPools"])
}

// To enable _ disable match across services with given persistence record
func (r *PersistenceProfileDstAddr) MatchAcrossServices() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["matchAcrossServices"])
}

// To enable _ disable match across services with given persistence record
func (r *PersistenceProfileDstAddr) MatchAcrossVirtuals() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["matchAcrossVirtuals"])
}

// To enable _ disable
func (r *PersistenceProfileDstAddr) Mirror() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["mirror"])
}

// Name of the persistence profile
func (r *PersistenceProfileDstAddr) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
// limits remain hard limits and are not overridden.
func (r *PersistenceProfileDstAddr) OverrideConnLimit() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["overrideConnLimit"])
}

// Timeout for persistence of the session
func (r *PersistenceProfileDstAddr) Timeout() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["timeout"])
}

// Input properties used for looking up and filtering PersistenceProfileDstAddr resources.
type PersistenceProfileDstAddrState struct {
	AppService interface{}
	// Inherit defaults from parent profile
	DefaultsFrom interface{}
	// Specify the hash algorithm
	HashAlgorithm interface{}
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask interface{}
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools interface{}
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices interface{}
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals interface{}
	// To enable _ disable
	Mirror interface{}
	// Name of the persistence profile
	Name interface{}
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit interface{}
	// Timeout for persistence of the session
	Timeout interface{}
}

// The set of arguments for constructing a PersistenceProfileDstAddr resource.
type PersistenceProfileDstAddrArgs struct {
	AppService interface{}
	// Inherit defaults from parent profile
	DefaultsFrom interface{}
	// Specify the hash algorithm
	HashAlgorithm interface{}
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask interface{}
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools interface{}
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices interface{}
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals interface{}
	// To enable _ disable
	Mirror interface{}
	// Name of the persistence profile
	Name interface{}
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit interface{}
	// Timeout for persistence of the session
	Timeout interface{}
}
