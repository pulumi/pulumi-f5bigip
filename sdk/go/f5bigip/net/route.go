// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `net.Route` Manages a route configuration
//
// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/net"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := net.NewRoute(ctx, "route2", &net.RouteArgs{
// 			Gw:      pulumi.String("1.1.1.2"),
// 			Name:    pulumi.String("external-route"),
// 			Network: pulumi.String("10.10.10.0/24"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Route struct {
	pulumi.CustomResourceState

	// Gateway address
	Gw pulumi.StringPtrOutput `pulumi:"gw"`
	// Name of the route
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a gateway address for the route.
	Network pulumi.StringOutput `pulumi:"network"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	if args == nil {
		args = &RouteArgs{}
	}
	var resource Route
	err := ctx.RegisterResource("f5bigip:net/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("f5bigip:net/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// Gateway address
	Gw *string `pulumi:"gw"`
	// Name of the route
	Name *string `pulumi:"name"`
	// Specifies a gateway address for the route.
	Network *string `pulumi:"network"`
}

type RouteState struct {
	// Gateway address
	Gw pulumi.StringPtrInput
	// Name of the route
	Name pulumi.StringPtrInput
	// Specifies a gateway address for the route.
	Network pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// Gateway address
	Gw *string `pulumi:"gw"`
	// Name of the route
	Name string `pulumi:"name"`
	// Specifies a gateway address for the route.
	Network string `pulumi:"network"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Gateway address
	Gw pulumi.StringPtrInput
	// Name of the route
	Name pulumi.StringInput
	// Specifies a gateway address for the route.
	Network pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}
