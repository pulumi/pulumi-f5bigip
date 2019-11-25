// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `ltm.Policy` Configures Virtual Server
// 
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_policy.html.markdown.
type Policy struct {
	s *pulumi.ResourceState
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOpt) (*Policy, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["controls"] = nil
		inputs["name"] = nil
		inputs["publishedCopy"] = nil
		inputs["requires"] = nil
		inputs["rules"] = nil
		inputs["strategy"] = nil
	} else {
		inputs["controls"] = args.Controls
		inputs["name"] = args.Name
		inputs["publishedCopy"] = args.PublishedCopy
		inputs["requires"] = args.Requires
		inputs["rules"] = args.Rules
		inputs["strategy"] = args.Strategy
	}
	s, err := ctx.RegisterResource("f5bigip:ltm/policy:Policy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyState, opts ...pulumi.ResourceOpt) (*Policy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["controls"] = state.Controls
		inputs["name"] = state.Name
		inputs["publishedCopy"] = state.PublishedCopy
		inputs["requires"] = state.Requires
		inputs["rules"] = state.Rules
		inputs["strategy"] = state.Strategy
	}
	s, err := ctx.ReadResource("f5bigip:ltm/policy:Policy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Policy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Policy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the controls
func (r *Policy) Controls() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["controls"])
}

// Name of the Policy
func (r *Policy) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// If you want to publish the policy else it will be deployed in Drafts mode.
func (r *Policy) PublishedCopy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["publishedCopy"])
}

// Specifies the protocol
func (r *Policy) Requires() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["requires"])
}

// Rules can be applied using the policy
func (r *Policy) Rules() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["rules"])
}

// Specifies the match strategy
func (r *Policy) Strategy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["strategy"])
}

// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	// Specifies the controls
	Controls interface{}
	// Name of the Policy
	Name interface{}
	// If you want to publish the policy else it will be deployed in Drafts mode.
	PublishedCopy interface{}
	// Specifies the protocol
	Requires interface{}
	// Rules can be applied using the policy
	Rules interface{}
	// Specifies the match strategy
	Strategy interface{}
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Specifies the controls
	Controls interface{}
	// Name of the Policy
	Name interface{}
	// If you want to publish the policy else it will be deployed in Drafts mode.
	PublishedCopy interface{}
	// Specifies the protocol
	Requires interface{}
	// Rules can be applied using the policy
	Rules interface{}
	// Specifies the match strategy
	Strategy interface{}
}
