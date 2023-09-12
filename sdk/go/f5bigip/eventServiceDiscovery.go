// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := f5bigip.NewEventServiceDiscovery(ctx, "test", &f5bigip.EventServiceDiscoveryArgs{
//				Nodes: f5bigip.EventServiceDiscoveryNodeArray{
//					&f5bigip.EventServiceDiscoveryNodeArgs{
//						Id:   pulumi.String("newNode1"),
//						Ip:   pulumi.String("192.168.2.3"),
//						Port: pulumi.Int(8080),
//					},
//					&f5bigip.EventServiceDiscoveryNodeArgs{
//						Id:   pulumi.String("newNode2"),
//						Ip:   pulumi.String("192.168.2.4"),
//						Port: pulumi.Int(8080),
//					},
//				},
//				Taskid: pulumi.String("~Sample_event_sd~My_app~My_pool"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type EventServiceDiscovery struct {
	pulumi.CustomResourceState

	// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
	//
	// For more information, please refer below document
	// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
	//
	// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
	//
	// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
	//
	// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
	//
	// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
	//
	// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
	// First we show the initial declaration to POST to the BIG-IP system.
	//
	// {
	// "class": "ADC",
	// "schemaVersion": "3.9.0",
	// "id": "Pool",
	// "Sample_event_sd": {
	// "class": "Tenant",
	// "My_app": {
	// "class": "Application",
	// "My_pool": {
	// "class": "Pool",
	// "members": [
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "static",
	// "serverAddresses": [
	// "192.0.2.2"
	// ]
	// },
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "event"
	// }
	// ]
	// }
	// }
	// }
	// }
	//
	// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
	Nodes EventServiceDiscoveryNodeArrayOutput `pulumi:"nodes"`
	// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
	Taskid pulumi.StringOutput `pulumi:"taskid"`
}

// NewEventServiceDiscovery registers a new resource with the given unique name, arguments, and options.
func NewEventServiceDiscovery(ctx *pulumi.Context,
	name string, args *EventServiceDiscoveryArgs, opts ...pulumi.ResourceOption) (*EventServiceDiscovery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Taskid == nil {
		return nil, errors.New("invalid value for required argument 'Taskid'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventServiceDiscovery
	err := ctx.RegisterResource("f5bigip:index/eventServiceDiscovery:EventServiceDiscovery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventServiceDiscovery gets an existing EventServiceDiscovery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventServiceDiscovery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventServiceDiscoveryState, opts ...pulumi.ResourceOption) (*EventServiceDiscovery, error) {
	var resource EventServiceDiscovery
	err := ctx.ReadResource("f5bigip:index/eventServiceDiscovery:EventServiceDiscovery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventServiceDiscovery resources.
type eventServiceDiscoveryState struct {
	// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
	//
	// For more information, please refer below document
	// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
	//
	// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
	//
	// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
	//
	// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
	//
	// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
	//
	// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
	// First we show the initial declaration to POST to the BIG-IP system.
	//
	// {
	// "class": "ADC",
	// "schemaVersion": "3.9.0",
	// "id": "Pool",
	// "Sample_event_sd": {
	// "class": "Tenant",
	// "My_app": {
	// "class": "Application",
	// "My_pool": {
	// "class": "Pool",
	// "members": [
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "static",
	// "serverAddresses": [
	// "192.0.2.2"
	// ]
	// },
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "event"
	// }
	// ]
	// }
	// }
	// }
	// }
	//
	// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
	Nodes []EventServiceDiscoveryNode `pulumi:"nodes"`
	// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
	Taskid *string `pulumi:"taskid"`
}

type EventServiceDiscoveryState struct {
	// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
	//
	// For more information, please refer below document
	// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
	//
	// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
	//
	// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
	//
	// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
	//
	// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
	//
	// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
	// First we show the initial declaration to POST to the BIG-IP system.
	//
	// {
	// "class": "ADC",
	// "schemaVersion": "3.9.0",
	// "id": "Pool",
	// "Sample_event_sd": {
	// "class": "Tenant",
	// "My_app": {
	// "class": "Application",
	// "My_pool": {
	// "class": "Pool",
	// "members": [
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "static",
	// "serverAddresses": [
	// "192.0.2.2"
	// ]
	// },
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "event"
	// }
	// ]
	// }
	// }
	// }
	// }
	//
	// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
	Nodes EventServiceDiscoveryNodeArrayInput
	// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
	Taskid pulumi.StringPtrInput
}

func (EventServiceDiscoveryState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventServiceDiscoveryState)(nil)).Elem()
}

type eventServiceDiscoveryArgs struct {
	// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
	//
	// For more information, please refer below document
	// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
	//
	// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
	//
	// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
	//
	// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
	//
	// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
	//
	// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
	// First we show the initial declaration to POST to the BIG-IP system.
	//
	// {
	// "class": "ADC",
	// "schemaVersion": "3.9.0",
	// "id": "Pool",
	// "Sample_event_sd": {
	// "class": "Tenant",
	// "My_app": {
	// "class": "Application",
	// "My_pool": {
	// "class": "Pool",
	// "members": [
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "static",
	// "serverAddresses": [
	// "192.0.2.2"
	// ]
	// },
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "event"
	// }
	// ]
	// }
	// }
	// }
	// }
	//
	// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
	Nodes []EventServiceDiscoveryNode `pulumi:"nodes"`
	// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
	Taskid string `pulumi:"taskid"`
}

// The set of arguments for constructing a EventServiceDiscovery resource.
type EventServiceDiscoveryArgs struct {
	// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
	//
	// For more information, please refer below document
	// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
	//
	// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
	//
	// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
	//
	// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
	//
	// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
	//
	// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
	// First we show the initial declaration to POST to the BIG-IP system.
	//
	// {
	// "class": "ADC",
	// "schemaVersion": "3.9.0",
	// "id": "Pool",
	// "Sample_event_sd": {
	// "class": "Tenant",
	// "My_app": {
	// "class": "Application",
	// "My_pool": {
	// "class": "Pool",
	// "members": [
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "static",
	// "serverAddresses": [
	// "192.0.2.2"
	// ]
	// },
	// {
	// "servicePort": 8080,
	// "addressDiscovery": "event"
	// }
	// ]
	// }
	// }
	// }
	// }
	//
	// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
	Nodes EventServiceDiscoveryNodeArrayInput
	// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
	Taskid pulumi.StringInput
}

func (EventServiceDiscoveryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventServiceDiscoveryArgs)(nil)).Elem()
}

type EventServiceDiscoveryInput interface {
	pulumi.Input

	ToEventServiceDiscoveryOutput() EventServiceDiscoveryOutput
	ToEventServiceDiscoveryOutputWithContext(ctx context.Context) EventServiceDiscoveryOutput
}

func (*EventServiceDiscovery) ElementType() reflect.Type {
	return reflect.TypeOf((**EventServiceDiscovery)(nil)).Elem()
}

func (i *EventServiceDiscovery) ToEventServiceDiscoveryOutput() EventServiceDiscoveryOutput {
	return i.ToEventServiceDiscoveryOutputWithContext(context.Background())
}

func (i *EventServiceDiscovery) ToEventServiceDiscoveryOutputWithContext(ctx context.Context) EventServiceDiscoveryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryOutput)
}

func (i *EventServiceDiscovery) ToOutput(ctx context.Context) pulumix.Output[*EventServiceDiscovery] {
	return pulumix.Output[*EventServiceDiscovery]{
		OutputState: i.ToEventServiceDiscoveryOutputWithContext(ctx).OutputState,
	}
}

// EventServiceDiscoveryArrayInput is an input type that accepts EventServiceDiscoveryArray and EventServiceDiscoveryArrayOutput values.
// You can construct a concrete instance of `EventServiceDiscoveryArrayInput` via:
//
//	EventServiceDiscoveryArray{ EventServiceDiscoveryArgs{...} }
type EventServiceDiscoveryArrayInput interface {
	pulumi.Input

	ToEventServiceDiscoveryArrayOutput() EventServiceDiscoveryArrayOutput
	ToEventServiceDiscoveryArrayOutputWithContext(context.Context) EventServiceDiscoveryArrayOutput
}

type EventServiceDiscoveryArray []EventServiceDiscoveryInput

func (EventServiceDiscoveryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventServiceDiscovery)(nil)).Elem()
}

func (i EventServiceDiscoveryArray) ToEventServiceDiscoveryArrayOutput() EventServiceDiscoveryArrayOutput {
	return i.ToEventServiceDiscoveryArrayOutputWithContext(context.Background())
}

func (i EventServiceDiscoveryArray) ToEventServiceDiscoveryArrayOutputWithContext(ctx context.Context) EventServiceDiscoveryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryArrayOutput)
}

func (i EventServiceDiscoveryArray) ToOutput(ctx context.Context) pulumix.Output[[]*EventServiceDiscovery] {
	return pulumix.Output[[]*EventServiceDiscovery]{
		OutputState: i.ToEventServiceDiscoveryArrayOutputWithContext(ctx).OutputState,
	}
}

// EventServiceDiscoveryMapInput is an input type that accepts EventServiceDiscoveryMap and EventServiceDiscoveryMapOutput values.
// You can construct a concrete instance of `EventServiceDiscoveryMapInput` via:
//
//	EventServiceDiscoveryMap{ "key": EventServiceDiscoveryArgs{...} }
type EventServiceDiscoveryMapInput interface {
	pulumi.Input

	ToEventServiceDiscoveryMapOutput() EventServiceDiscoveryMapOutput
	ToEventServiceDiscoveryMapOutputWithContext(context.Context) EventServiceDiscoveryMapOutput
}

type EventServiceDiscoveryMap map[string]EventServiceDiscoveryInput

func (EventServiceDiscoveryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventServiceDiscovery)(nil)).Elem()
}

func (i EventServiceDiscoveryMap) ToEventServiceDiscoveryMapOutput() EventServiceDiscoveryMapOutput {
	return i.ToEventServiceDiscoveryMapOutputWithContext(context.Background())
}

func (i EventServiceDiscoveryMap) ToEventServiceDiscoveryMapOutputWithContext(ctx context.Context) EventServiceDiscoveryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryMapOutput)
}

func (i EventServiceDiscoveryMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*EventServiceDiscovery] {
	return pulumix.Output[map[string]*EventServiceDiscovery]{
		OutputState: i.ToEventServiceDiscoveryMapOutputWithContext(ctx).OutputState,
	}
}

type EventServiceDiscoveryOutput struct{ *pulumi.OutputState }

func (EventServiceDiscoveryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventServiceDiscovery)(nil)).Elem()
}

func (o EventServiceDiscoveryOutput) ToEventServiceDiscoveryOutput() EventServiceDiscoveryOutput {
	return o
}

func (o EventServiceDiscoveryOutput) ToEventServiceDiscoveryOutputWithContext(ctx context.Context) EventServiceDiscoveryOutput {
	return o
}

func (o EventServiceDiscoveryOutput) ToOutput(ctx context.Context) pulumix.Output[*EventServiceDiscovery] {
	return pulumix.Output[*EventServiceDiscovery]{
		OutputState: o.OutputState,
	}
}

// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
//
// For more information, please refer below document
// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
//
// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
//
// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
//
// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
//
// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
//
// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
// First we show the initial declaration to POST to the BIG-IP system.
//
// {
// "class": "ADC",
// "schemaVersion": "3.9.0",
// "id": "Pool",
// "Sample_event_sd": {
// "class": "Tenant",
// "My_app": {
// "class": "Application",
// "My_pool": {
// "class": "Pool",
// "members": [
// {
// "servicePort": 8080,
// "addressDiscovery": "static",
// "serverAddresses": [
// "192.0.2.2"
// ]
// },
// {
// "servicePort": 8080,
// "addressDiscovery": "event"
// }
// ]
// }
// }
// }
// }
//
// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
func (o EventServiceDiscoveryOutput) Nodes() EventServiceDiscoveryNodeArrayOutput {
	return o.ApplyT(func(v *EventServiceDiscovery) EventServiceDiscoveryNodeArrayOutput { return v.Nodes }).(EventServiceDiscoveryNodeArrayOutput)
}

// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
func (o EventServiceDiscoveryOutput) Taskid() pulumi.StringOutput {
	return o.ApplyT(func(v *EventServiceDiscovery) pulumi.StringOutput { return v.Taskid }).(pulumi.StringOutput)
}

type EventServiceDiscoveryArrayOutput struct{ *pulumi.OutputState }

func (EventServiceDiscoveryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventServiceDiscovery)(nil)).Elem()
}

func (o EventServiceDiscoveryArrayOutput) ToEventServiceDiscoveryArrayOutput() EventServiceDiscoveryArrayOutput {
	return o
}

func (o EventServiceDiscoveryArrayOutput) ToEventServiceDiscoveryArrayOutputWithContext(ctx context.Context) EventServiceDiscoveryArrayOutput {
	return o
}

func (o EventServiceDiscoveryArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*EventServiceDiscovery] {
	return pulumix.Output[[]*EventServiceDiscovery]{
		OutputState: o.OutputState,
	}
}

func (o EventServiceDiscoveryArrayOutput) Index(i pulumi.IntInput) EventServiceDiscoveryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventServiceDiscovery {
		return vs[0].([]*EventServiceDiscovery)[vs[1].(int)]
	}).(EventServiceDiscoveryOutput)
}

type EventServiceDiscoveryMapOutput struct{ *pulumi.OutputState }

func (EventServiceDiscoveryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventServiceDiscovery)(nil)).Elem()
}

func (o EventServiceDiscoveryMapOutput) ToEventServiceDiscoveryMapOutput() EventServiceDiscoveryMapOutput {
	return o
}

func (o EventServiceDiscoveryMapOutput) ToEventServiceDiscoveryMapOutputWithContext(ctx context.Context) EventServiceDiscoveryMapOutput {
	return o
}

func (o EventServiceDiscoveryMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*EventServiceDiscovery] {
	return pulumix.Output[map[string]*EventServiceDiscovery]{
		OutputState: o.OutputState,
	}
}

func (o EventServiceDiscoveryMapOutput) MapIndex(k pulumi.StringInput) EventServiceDiscoveryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventServiceDiscovery {
		return vs[0].(map[string]*EventServiceDiscovery)[vs[1].(string)]
	}).(EventServiceDiscoveryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventServiceDiscoveryInput)(nil)).Elem(), &EventServiceDiscovery{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventServiceDiscoveryArrayInput)(nil)).Elem(), EventServiceDiscoveryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventServiceDiscoveryMapInput)(nil)).Elem(), EventServiceDiscoveryMap{})
	pulumi.RegisterOutputType(EventServiceDiscoveryOutput{})
	pulumi.RegisterOutputType(EventServiceDiscoveryArrayOutput{})
	pulumi.RegisterOutputType(EventServiceDiscoveryMapOutput{})
}
