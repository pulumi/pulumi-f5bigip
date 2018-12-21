// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type PersistenceProfileSsl struct {
	s *pulumi.ResourceState
}

// NewPersistenceProfileSsl registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileSsl(ctx *pulumi.Context,
	name string, args *PersistenceProfileSslArgs, opts ...pulumi.ResourceOpt) (*PersistenceProfileSsl, error) {
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
		inputs["matchAcrossPools"] = args.MatchAcrossPools
		inputs["matchAcrossServices"] = args.MatchAcrossServices
		inputs["matchAcrossVirtuals"] = args.MatchAcrossVirtuals
		inputs["mirror"] = args.Mirror
		inputs["name"] = args.Name
		inputs["overrideConnLimit"] = args.OverrideConnLimit
		inputs["timeout"] = args.Timeout
	}
	s, err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PersistenceProfileSsl{s: s}, nil
}

// GetPersistenceProfileSsl gets an existing PersistenceProfileSsl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileSsl(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PersistenceProfileSslState, opts ...pulumi.ResourceOpt) (*PersistenceProfileSsl, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["appService"] = state.AppService
		inputs["defaultsFrom"] = state.DefaultsFrom
		inputs["matchAcrossPools"] = state.MatchAcrossPools
		inputs["matchAcrossServices"] = state.MatchAcrossServices
		inputs["matchAcrossVirtuals"] = state.MatchAcrossVirtuals
		inputs["mirror"] = state.Mirror
		inputs["name"] = state.Name
		inputs["overrideConnLimit"] = state.OverrideConnLimit
		inputs["timeout"] = state.Timeout
	}
	s, err := ctx.ReadResource("f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PersistenceProfileSsl{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PersistenceProfileSsl) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PersistenceProfileSsl) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *PersistenceProfileSsl) AppService() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["appService"])
}

// Inherit defaults from parent profile
func (r *PersistenceProfileSsl) DefaultsFrom() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultsFrom"])
}

// To enable _ disable match across pools with given persistence record
func (r *PersistenceProfileSsl) MatchAcrossPools() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["matchAcrossPools"])
}

// To enable _ disable match across services with given persistence record
func (r *PersistenceProfileSsl) MatchAcrossServices() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["matchAcrossServices"])
}

// To enable _ disable match across services with given persistence record
func (r *PersistenceProfileSsl) MatchAcrossVirtuals() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["matchAcrossVirtuals"])
}

// To enable _ disable
func (r *PersistenceProfileSsl) Mirror() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["mirror"])
}

// Name of the persistence profile
func (r *PersistenceProfileSsl) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
// limits remain hard limits and are not overridden.
func (r *PersistenceProfileSsl) OverrideConnLimit() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["overrideConnLimit"])
}

// Timeout for persistence of the session
func (r *PersistenceProfileSsl) Timeout() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["timeout"])
}

// Input properties used for looking up and filtering PersistenceProfileSsl resources.
type PersistenceProfileSslState struct {
	AppService interface{}
	// Inherit defaults from parent profile
	DefaultsFrom interface{}
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

// The set of arguments for constructing a PersistenceProfileSsl resource.
type PersistenceProfileSslArgs struct {
	AppService interface{}
	// Inherit defaults from parent profile
	DefaultsFrom interface{}
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
