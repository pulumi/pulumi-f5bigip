// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `net.Route` Manages a route configuration
// 
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/net_route.html.markdown.
type Route struct {
	s *pulumi.ResourceState
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOpt) (*Route, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["gw"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
	} else {
		inputs["gw"] = args.Gw
		inputs["name"] = args.Name
		inputs["network"] = args.Network
	}
	s, err := ctx.RegisterResource("f5bigip:net/route:Route", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Route{s: s}, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RouteState, opts ...pulumi.ResourceOpt) (*Route, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["gw"] = state.Gw
		inputs["name"] = state.Name
		inputs["network"] = state.Network
	}
	s, err := ctx.ReadResource("f5bigip:net/route:Route", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Route{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Route) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Route) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Gateway address
func (r *Route) Gw() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["gw"])
}

// Name of the route
func (r *Route) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Specifies a gateway address for the route.
func (r *Route) Network() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["network"])
}

// Input properties used for looking up and filtering Route resources.
type RouteState struct {
	// Gateway address
	Gw interface{}
	// Name of the route
	Name interface{}
	// Specifies a gateway address for the route.
	Network interface{}
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Gateway address
	Gw interface{}
	// Name of the route
	Name interface{}
	// Specifies a gateway address for the route.
	Network interface{}
}
