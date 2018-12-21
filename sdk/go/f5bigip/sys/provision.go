// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Provision struct {
	s *pulumi.ResourceState
}

// NewProvision registers a new resource with the given unique name, arguments, and options.
func NewProvision(ctx *pulumi.Context,
	name string, args *ProvisionArgs, opts ...pulumi.ResourceOpt) (*Provision, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cpuRatio"] = nil
		inputs["diskRatio"] = nil
		inputs["fullPath"] = nil
		inputs["level"] = nil
		inputs["memoryRatio"] = nil
		inputs["name"] = nil
	} else {
		inputs["cpuRatio"] = args.CpuRatio
		inputs["diskRatio"] = args.DiskRatio
		inputs["fullPath"] = args.FullPath
		inputs["level"] = args.Level
		inputs["memoryRatio"] = args.MemoryRatio
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("f5bigip:sys/provision:Provision", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Provision{s: s}, nil
}

// GetProvision gets an existing Provision resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvision(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ProvisionState, opts ...pulumi.ResourceOpt) (*Provision, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cpuRatio"] = state.CpuRatio
		inputs["diskRatio"] = state.DiskRatio
		inputs["fullPath"] = state.FullPath
		inputs["level"] = state.Level
		inputs["memoryRatio"] = state.MemoryRatio
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("f5bigip:sys/provision:Provision", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Provision{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Provision) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Provision) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// cpu Ratio
func (r *Provision) CpuRatio() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cpuRatio"])
}

// disk Ratio
func (r *Provision) DiskRatio() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["diskRatio"])
}

// path
func (r *Provision) FullPath() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fullPath"])
}

// what level nominal or dedicated
func (r *Provision) Level() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["level"])
}

// memory Ratio
func (r *Provision) MemoryRatio() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["memoryRatio"])
}

// Name of the module to be provisioned
func (r *Provision) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering Provision resources.
type ProvisionState struct {
	// cpu Ratio
	CpuRatio interface{}
	// disk Ratio
	DiskRatio interface{}
	// path
	FullPath interface{}
	// what level nominal or dedicated
	Level interface{}
	// memory Ratio
	MemoryRatio interface{}
	// Name of the module to be provisioned
	Name interface{}
}

// The set of arguments for constructing a Provision resource.
type ProvisionArgs struct {
	// cpu Ratio
	CpuRatio interface{}
	// disk Ratio
	DiskRatio interface{}
	// path
	FullPath interface{}
	// what level nominal or dedicated
	Level interface{}
	// memory Ratio
	MemoryRatio interface{}
	// Name of the module to be provisioned
	Name interface{}
}
