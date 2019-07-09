// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `bigip_net_vlan` Manages a vlan configuration
// 
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/net_vlan.html.markdown.
type Vlan struct {
	s *pulumi.ResourceState
}

// NewVlan registers a new resource with the given unique name, arguments, and options.
func NewVlan(ctx *pulumi.Context,
	name string, args *VlanArgs, opts ...pulumi.ResourceOpt) (*Vlan, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["interfaces"] = nil
		inputs["name"] = nil
		inputs["tag"] = nil
	} else {
		inputs["interfaces"] = args.Interfaces
		inputs["name"] = args.Name
		inputs["tag"] = args.Tag
	}
	s, err := ctx.RegisterResource("f5bigip:net/vlan:Vlan", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vlan{s: s}, nil
}

// GetVlan gets an existing Vlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVlan(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VlanState, opts ...pulumi.ResourceOpt) (*Vlan, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["interfaces"] = state.Interfaces
		inputs["name"] = state.Name
		inputs["tag"] = state.Tag
	}
	s, err := ctx.ReadResource("f5bigip:net/vlan:Vlan", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Vlan{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Vlan) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Vlan) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies which interfaces you want this VLAN to use for traffic management.
func (r *Vlan) Interfaces() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["interfaces"])
}

// Name of the vlan
func (r *Vlan) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specifies a number that the system adds into the header of any frame passing through the VLAN.
func (r *Vlan) Tag() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["tag"])
}

// Input properties used for looking up and filtering Vlan resources.
type VlanState struct {
	// Specifies which interfaces you want this VLAN to use for traffic management.
	Interfaces interface{}
	// Name of the vlan
	Name interface{}
	// Specifies a number that the system adds into the header of any frame passing through the VLAN.
	Tag interface{}
}

// The set of arguments for constructing a Vlan resource.
type VlanArgs struct {
	// Specifies which interfaces you want this VLAN to use for traffic management.
	Interfaces interface{}
	// Name of the vlan
	Name interface{}
	// Specifies a number that the system adds into the header of any frame passing through the VLAN.
	Tag interface{}
}
