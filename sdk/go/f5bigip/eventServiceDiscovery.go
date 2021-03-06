// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type EventServiceDiscovery struct {
	pulumi.CustomResourceState

	Nodes EventServiceDiscoveryNodeArrayOutput `pulumi:"nodes"`
	// Name of the partition/tenant
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
	Nodes []EventServiceDiscoveryNode `pulumi:"nodes"`
	// Name of the partition/tenant
	Taskid *string `pulumi:"taskid"`
}

type EventServiceDiscoveryState struct {
	Nodes EventServiceDiscoveryNodeArrayInput
	// Name of the partition/tenant
	Taskid pulumi.StringPtrInput
}

func (EventServiceDiscoveryState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventServiceDiscoveryState)(nil)).Elem()
}

type eventServiceDiscoveryArgs struct {
	Nodes []EventServiceDiscoveryNode `pulumi:"nodes"`
	// Name of the partition/tenant
	Taskid string `pulumi:"taskid"`
}

// The set of arguments for constructing a EventServiceDiscovery resource.
type EventServiceDiscoveryArgs struct {
	Nodes EventServiceDiscoveryNodeArrayInput
	// Name of the partition/tenant
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
	return reflect.TypeOf((*EventServiceDiscovery)(nil))
}

func (i *EventServiceDiscovery) ToEventServiceDiscoveryOutput() EventServiceDiscoveryOutput {
	return i.ToEventServiceDiscoveryOutputWithContext(context.Background())
}

func (i *EventServiceDiscovery) ToEventServiceDiscoveryOutputWithContext(ctx context.Context) EventServiceDiscoveryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryOutput)
}

func (i *EventServiceDiscovery) ToEventServiceDiscoveryPtrOutput() EventServiceDiscoveryPtrOutput {
	return i.ToEventServiceDiscoveryPtrOutputWithContext(context.Background())
}

func (i *EventServiceDiscovery) ToEventServiceDiscoveryPtrOutputWithContext(ctx context.Context) EventServiceDiscoveryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryPtrOutput)
}

type EventServiceDiscoveryPtrInput interface {
	pulumi.Input

	ToEventServiceDiscoveryPtrOutput() EventServiceDiscoveryPtrOutput
	ToEventServiceDiscoveryPtrOutputWithContext(ctx context.Context) EventServiceDiscoveryPtrOutput
}

type eventServiceDiscoveryPtrType EventServiceDiscoveryArgs

func (*eventServiceDiscoveryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventServiceDiscovery)(nil))
}

func (i *eventServiceDiscoveryPtrType) ToEventServiceDiscoveryPtrOutput() EventServiceDiscoveryPtrOutput {
	return i.ToEventServiceDiscoveryPtrOutputWithContext(context.Background())
}

func (i *eventServiceDiscoveryPtrType) ToEventServiceDiscoveryPtrOutputWithContext(ctx context.Context) EventServiceDiscoveryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryPtrOutput)
}

// EventServiceDiscoveryArrayInput is an input type that accepts EventServiceDiscoveryArray and EventServiceDiscoveryArrayOutput values.
// You can construct a concrete instance of `EventServiceDiscoveryArrayInput` via:
//
//          EventServiceDiscoveryArray{ EventServiceDiscoveryArgs{...} }
type EventServiceDiscoveryArrayInput interface {
	pulumi.Input

	ToEventServiceDiscoveryArrayOutput() EventServiceDiscoveryArrayOutput
	ToEventServiceDiscoveryArrayOutputWithContext(context.Context) EventServiceDiscoveryArrayOutput
}

type EventServiceDiscoveryArray []EventServiceDiscoveryInput

func (EventServiceDiscoveryArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EventServiceDiscovery)(nil))
}

func (i EventServiceDiscoveryArray) ToEventServiceDiscoveryArrayOutput() EventServiceDiscoveryArrayOutput {
	return i.ToEventServiceDiscoveryArrayOutputWithContext(context.Background())
}

func (i EventServiceDiscoveryArray) ToEventServiceDiscoveryArrayOutputWithContext(ctx context.Context) EventServiceDiscoveryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryArrayOutput)
}

// EventServiceDiscoveryMapInput is an input type that accepts EventServiceDiscoveryMap and EventServiceDiscoveryMapOutput values.
// You can construct a concrete instance of `EventServiceDiscoveryMapInput` via:
//
//          EventServiceDiscoveryMap{ "key": EventServiceDiscoveryArgs{...} }
type EventServiceDiscoveryMapInput interface {
	pulumi.Input

	ToEventServiceDiscoveryMapOutput() EventServiceDiscoveryMapOutput
	ToEventServiceDiscoveryMapOutputWithContext(context.Context) EventServiceDiscoveryMapOutput
}

type EventServiceDiscoveryMap map[string]EventServiceDiscoveryInput

func (EventServiceDiscoveryMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EventServiceDiscovery)(nil))
}

func (i EventServiceDiscoveryMap) ToEventServiceDiscoveryMapOutput() EventServiceDiscoveryMapOutput {
	return i.ToEventServiceDiscoveryMapOutputWithContext(context.Background())
}

func (i EventServiceDiscoveryMap) ToEventServiceDiscoveryMapOutputWithContext(ctx context.Context) EventServiceDiscoveryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventServiceDiscoveryMapOutput)
}

type EventServiceDiscoveryOutput struct {
	*pulumi.OutputState
}

func (EventServiceDiscoveryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventServiceDiscovery)(nil))
}

func (o EventServiceDiscoveryOutput) ToEventServiceDiscoveryOutput() EventServiceDiscoveryOutput {
	return o
}

func (o EventServiceDiscoveryOutput) ToEventServiceDiscoveryOutputWithContext(ctx context.Context) EventServiceDiscoveryOutput {
	return o
}

func (o EventServiceDiscoveryOutput) ToEventServiceDiscoveryPtrOutput() EventServiceDiscoveryPtrOutput {
	return o.ToEventServiceDiscoveryPtrOutputWithContext(context.Background())
}

func (o EventServiceDiscoveryOutput) ToEventServiceDiscoveryPtrOutputWithContext(ctx context.Context) EventServiceDiscoveryPtrOutput {
	return o.ApplyT(func(v EventServiceDiscovery) *EventServiceDiscovery {
		return &v
	}).(EventServiceDiscoveryPtrOutput)
}

type EventServiceDiscoveryPtrOutput struct {
	*pulumi.OutputState
}

func (EventServiceDiscoveryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventServiceDiscovery)(nil))
}

func (o EventServiceDiscoveryPtrOutput) ToEventServiceDiscoveryPtrOutput() EventServiceDiscoveryPtrOutput {
	return o
}

func (o EventServiceDiscoveryPtrOutput) ToEventServiceDiscoveryPtrOutputWithContext(ctx context.Context) EventServiceDiscoveryPtrOutput {
	return o
}

type EventServiceDiscoveryArrayOutput struct{ *pulumi.OutputState }

func (EventServiceDiscoveryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventServiceDiscovery)(nil))
}

func (o EventServiceDiscoveryArrayOutput) ToEventServiceDiscoveryArrayOutput() EventServiceDiscoveryArrayOutput {
	return o
}

func (o EventServiceDiscoveryArrayOutput) ToEventServiceDiscoveryArrayOutputWithContext(ctx context.Context) EventServiceDiscoveryArrayOutput {
	return o
}

func (o EventServiceDiscoveryArrayOutput) Index(i pulumi.IntInput) EventServiceDiscoveryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventServiceDiscovery {
		return vs[0].([]EventServiceDiscovery)[vs[1].(int)]
	}).(EventServiceDiscoveryOutput)
}

type EventServiceDiscoveryMapOutput struct{ *pulumi.OutputState }

func (EventServiceDiscoveryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventServiceDiscovery)(nil))
}

func (o EventServiceDiscoveryMapOutput) ToEventServiceDiscoveryMapOutput() EventServiceDiscoveryMapOutput {
	return o
}

func (o EventServiceDiscoveryMapOutput) ToEventServiceDiscoveryMapOutputWithContext(ctx context.Context) EventServiceDiscoveryMapOutput {
	return o
}

func (o EventServiceDiscoveryMapOutput) MapIndex(k pulumi.StringInput) EventServiceDiscoveryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventServiceDiscovery {
		return vs[0].(map[string]EventServiceDiscovery)[vs[1].(string)]
	}).(EventServiceDiscoveryOutput)
}

func init() {
	pulumi.RegisterOutputType(EventServiceDiscoveryOutput{})
	pulumi.RegisterOutputType(EventServiceDiscoveryPtrOutput{})
	pulumi.RegisterOutputType(EventServiceDiscoveryArrayOutput{})
	pulumi.RegisterOutputType(EventServiceDiscoveryMapOutput{})
}
