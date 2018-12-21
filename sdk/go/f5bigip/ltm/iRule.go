// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type IRule struct {
	s *pulumi.ResourceState
}

// NewIRule registers a new resource with the given unique name, arguments, and options.
func NewIRule(ctx *pulumi.Context,
	name string, args *IRuleArgs, opts ...pulumi.ResourceOpt) (*IRule, error) {
	if args == nil || args.Irule == nil {
		return nil, errors.New("missing required argument 'Irule'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["irule"] = nil
		inputs["name"] = nil
	} else {
		inputs["irule"] = args.Irule
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("f5bigip:ltm/iRule:IRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IRule{s: s}, nil
}

// GetIRule gets an existing IRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IRuleState, opts ...pulumi.ResourceOpt) (*IRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["irule"] = state.Irule
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("f5bigip:ltm/iRule:IRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &IRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *IRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *IRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The iRule body
func (r *IRule) Irule() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["irule"])
}

// Name of the iRule
func (r *IRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering IRule resources.
type IRuleState struct {
	// The iRule body
	Irule interface{}
	// Name of the iRule
	Name interface{}
}

// The set of arguments for constructing a IRule resource.
type IRuleArgs struct {
	// The iRule body
	Irule interface{}
	// Name of the iRule
	Name interface{}
}
