// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package net

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/net"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := net.NewRoute(ctx, "route2", &net.RouteArgs{
// 			Gw:      pulumi.String("1.1.1.2"),
// 			Name:    pulumi.String("/Common/external-route"),
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

	// Specifies a gateway address for the route.
	Gw pulumi.StringPtrOutput `pulumi:"gw"`
	// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The destination subnet and netmask for the route.
	Network pulumi.StringOutput `pulumi:"network"`
	// reject route
	Reject pulumi.BoolPtrOutput `pulumi:"reject"`
	// tunnel_ref to route traffic
	TunnelRef pulumi.StringPtrOutput `pulumi:"tunnelRef"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
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
	// Specifies a gateway address for the route.
	Gw *string `pulumi:"gw"`
	// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
	Name *string `pulumi:"name"`
	// The destination subnet and netmask for the route.
	Network *string `pulumi:"network"`
	// reject route
	Reject *bool `pulumi:"reject"`
	// tunnel_ref to route traffic
	TunnelRef *string `pulumi:"tunnelRef"`
}

type RouteState struct {
	// Specifies a gateway address for the route.
	Gw pulumi.StringPtrInput
	// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
	Name pulumi.StringPtrInput
	// The destination subnet and netmask for the route.
	Network pulumi.StringPtrInput
	// reject route
	Reject pulumi.BoolPtrInput
	// tunnel_ref to route traffic
	TunnelRef pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// Specifies a gateway address for the route.
	Gw *string `pulumi:"gw"`
	// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
	Name string `pulumi:"name"`
	// The destination subnet and netmask for the route.
	Network string `pulumi:"network"`
	// reject route
	Reject *bool `pulumi:"reject"`
	// tunnel_ref to route traffic
	TunnelRef *string `pulumi:"tunnelRef"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Specifies a gateway address for the route.
	Gw pulumi.StringPtrInput
	// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
	Name pulumi.StringInput
	// The destination subnet and netmask for the route.
	Network pulumi.StringInput
	// reject route
	Reject pulumi.BoolPtrInput
	// tunnel_ref to route traffic
	TunnelRef pulumi.StringPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i *Route) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePtrOutput)
}

type RoutePtrInput interface {
	pulumi.Input

	ToRoutePtrOutput() RoutePtrOutput
	ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput
}

type routePtrType RouteArgs

func (*routePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil))
}

func (i *routePtrType) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i *routePtrType) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePtrOutput)
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//          RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Route)(nil))
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//          RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Route)(nil))
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

type RouteOutput struct {
	*pulumi.OutputState
}

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o.ToRoutePtrOutputWithContext(context.Background())
}

func (o RouteOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o.ApplyT(func(v Route) *Route {
		return &v
	}).(RoutePtrOutput)
}

type RoutePtrOutput struct {
	*pulumi.OutputState
}

func (RoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil))
}

func (o RoutePtrOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o
}

func (o RoutePtrOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Route)(nil))
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Route {
		return vs[0].([]Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Route)(nil))
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Route {
		return vs[0].(map[string]Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RoutePtrOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
