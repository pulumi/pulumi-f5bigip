// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `cm.DeviceGroup` A device group is a collection of BIG-IP devices that are configured to securely synchronize their BIG-IP configuration data, and fail over when needed.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/cm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cm.NewDeviceGroup(ctx, "myNewDevicegroup", &cm.DeviceGroupArgs{
// 			AutoSync: pulumi.String("enabled"),
// 			Devices: cm.DeviceGroupDeviceArray{
// 				&cm.DeviceGroupDeviceArgs{
// 					Name: pulumi.String("bigip1.cisco.com"),
// 				},
// 				&cm.DeviceGroupDeviceArgs{
// 					Name: pulumi.String("bigip200.f5.com"),
// 				},
// 			},
// 			FullLoadOnSync: pulumi.String("true"),
// 			Name:           pulumi.String("sanjose_devicegroup"),
// 			Type:           pulumi.String("sync-only"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DeviceGroup struct {
	pulumi.CustomResourceState

	// Specifies if the device-group will automatically sync configuration data to its members
	AutoSync pulumi.StringPtrOutput `pulumi:"autoSync"`
	// Description of Device group
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the device to be included in device group, this need to be configured before using devicegroup resource
	Devices DeviceGroupDeviceArrayOutput `pulumi:"devices"`
	// Specifies if the device-group will perform a full-load upon sync
	FullLoadOnSync pulumi.StringPtrOutput `pulumi:"fullLoadOnSync"`
	// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
	IncrementalConfig pulumi.IntPtrOutput `pulumi:"incrementalConfig"`
	// Is the name of the device Group
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Specifies if the device-group will use a network connection for failover
	NetworkFailover pulumi.StringPtrOutput `pulumi:"networkFailover"`
	// Device administrative partition
	Partition pulumi.StringPtrOutput `pulumi:"partition"`
	// Specifies whether the configuration should be saved upon auto-sync.
	SaveOnAutoSync pulumi.StringPtrOutput `pulumi:"saveOnAutoSync"`
	// Specifies if the device-group will be used for failover or resource syncing
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewDeviceGroup registers a new resource with the given unique name, arguments, and options.
func NewDeviceGroup(ctx *pulumi.Context,
	name string, args *DeviceGroupArgs, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	if args == nil {
		args = &DeviceGroupArgs{}
	}

	var resource DeviceGroup
	err := ctx.RegisterResource("f5bigip:cm/deviceGroup:DeviceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceGroup gets an existing DeviceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceGroupState, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	var resource DeviceGroup
	err := ctx.ReadResource("f5bigip:cm/deviceGroup:DeviceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceGroup resources.
type deviceGroupState struct {
	// Specifies if the device-group will automatically sync configuration data to its members
	AutoSync *string `pulumi:"autoSync"`
	// Description of Device group
	Description *string `pulumi:"description"`
	// Name of the device to be included in device group, this need to be configured before using devicegroup resource
	Devices []DeviceGroupDevice `pulumi:"devices"`
	// Specifies if the device-group will perform a full-load upon sync
	FullLoadOnSync *string `pulumi:"fullLoadOnSync"`
	// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
	IncrementalConfig *int `pulumi:"incrementalConfig"`
	// Is the name of the device Group
	Name *string `pulumi:"name"`
	// Specifies if the device-group will use a network connection for failover
	NetworkFailover *string `pulumi:"networkFailover"`
	// Device administrative partition
	Partition *string `pulumi:"partition"`
	// Specifies whether the configuration should be saved upon auto-sync.
	SaveOnAutoSync *string `pulumi:"saveOnAutoSync"`
	// Specifies if the device-group will be used for failover or resource syncing
	Type *string `pulumi:"type"`
}

type DeviceGroupState struct {
	// Specifies if the device-group will automatically sync configuration data to its members
	AutoSync pulumi.StringPtrInput
	// Description of Device group
	Description pulumi.StringPtrInput
	// Name of the device to be included in device group, this need to be configured before using devicegroup resource
	Devices DeviceGroupDeviceArrayInput
	// Specifies if the device-group will perform a full-load upon sync
	FullLoadOnSync pulumi.StringPtrInput
	// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
	IncrementalConfig pulumi.IntPtrInput
	// Is the name of the device Group
	Name pulumi.StringPtrInput
	// Specifies if the device-group will use a network connection for failover
	NetworkFailover pulumi.StringPtrInput
	// Device administrative partition
	Partition pulumi.StringPtrInput
	// Specifies whether the configuration should be saved upon auto-sync.
	SaveOnAutoSync pulumi.StringPtrInput
	// Specifies if the device-group will be used for failover or resource syncing
	Type pulumi.StringPtrInput
}

func (DeviceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupState)(nil)).Elem()
}

type deviceGroupArgs struct {
	// Specifies if the device-group will automatically sync configuration data to its members
	AutoSync *string `pulumi:"autoSync"`
	// Description of Device group
	Description *string `pulumi:"description"`
	// Name of the device to be included in device group, this need to be configured before using devicegroup resource
	Devices []DeviceGroupDevice `pulumi:"devices"`
	// Specifies if the device-group will perform a full-load upon sync
	FullLoadOnSync *string `pulumi:"fullLoadOnSync"`
	// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
	IncrementalConfig *int `pulumi:"incrementalConfig"`
	// Is the name of the device Group
	Name *string `pulumi:"name"`
	// Specifies if the device-group will use a network connection for failover
	NetworkFailover *string `pulumi:"networkFailover"`
	// Device administrative partition
	Partition *string `pulumi:"partition"`
	// Specifies whether the configuration should be saved upon auto-sync.
	SaveOnAutoSync *string `pulumi:"saveOnAutoSync"`
	// Specifies if the device-group will be used for failover or resource syncing
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a DeviceGroup resource.
type DeviceGroupArgs struct {
	// Specifies if the device-group will automatically sync configuration data to its members
	AutoSync pulumi.StringPtrInput
	// Description of Device group
	Description pulumi.StringPtrInput
	// Name of the device to be included in device group, this need to be configured before using devicegroup resource
	Devices DeviceGroupDeviceArrayInput
	// Specifies if the device-group will perform a full-load upon sync
	FullLoadOnSync pulumi.StringPtrInput
	// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
	IncrementalConfig pulumi.IntPtrInput
	// Is the name of the device Group
	Name pulumi.StringPtrInput
	// Specifies if the device-group will use a network connection for failover
	NetworkFailover pulumi.StringPtrInput
	// Device administrative partition
	Partition pulumi.StringPtrInput
	// Specifies whether the configuration should be saved upon auto-sync.
	SaveOnAutoSync pulumi.StringPtrInput
	// Specifies if the device-group will be used for failover or resource syncing
	Type pulumi.StringPtrInput
}

func (DeviceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupArgs)(nil)).Elem()
}

type DeviceGroupInput interface {
	pulumi.Input

	ToDeviceGroupOutput() DeviceGroupOutput
	ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput
}

func (*DeviceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil)).Elem()
}

func (i *DeviceGroup) ToDeviceGroupOutput() DeviceGroupOutput {
	return i.ToDeviceGroupOutputWithContext(context.Background())
}

func (i *DeviceGroup) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupOutput)
}

// DeviceGroupArrayInput is an input type that accepts DeviceGroupArray and DeviceGroupArrayOutput values.
// You can construct a concrete instance of `DeviceGroupArrayInput` via:
//
//          DeviceGroupArray{ DeviceGroupArgs{...} }
type DeviceGroupArrayInput interface {
	pulumi.Input

	ToDeviceGroupArrayOutput() DeviceGroupArrayOutput
	ToDeviceGroupArrayOutputWithContext(context.Context) DeviceGroupArrayOutput
}

type DeviceGroupArray []DeviceGroupInput

func (DeviceGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceGroup)(nil)).Elem()
}

func (i DeviceGroupArray) ToDeviceGroupArrayOutput() DeviceGroupArrayOutput {
	return i.ToDeviceGroupArrayOutputWithContext(context.Background())
}

func (i DeviceGroupArray) ToDeviceGroupArrayOutputWithContext(ctx context.Context) DeviceGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupArrayOutput)
}

// DeviceGroupMapInput is an input type that accepts DeviceGroupMap and DeviceGroupMapOutput values.
// You can construct a concrete instance of `DeviceGroupMapInput` via:
//
//          DeviceGroupMap{ "key": DeviceGroupArgs{...} }
type DeviceGroupMapInput interface {
	pulumi.Input

	ToDeviceGroupMapOutput() DeviceGroupMapOutput
	ToDeviceGroupMapOutputWithContext(context.Context) DeviceGroupMapOutput
}

type DeviceGroupMap map[string]DeviceGroupInput

func (DeviceGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceGroup)(nil)).Elem()
}

func (i DeviceGroupMap) ToDeviceGroupMapOutput() DeviceGroupMapOutput {
	return i.ToDeviceGroupMapOutputWithContext(context.Background())
}

func (i DeviceGroupMap) ToDeviceGroupMapOutputWithContext(ctx context.Context) DeviceGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupMapOutput)
}

type DeviceGroupOutput struct{ *pulumi.OutputState }

func (DeviceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil)).Elem()
}

func (o DeviceGroupOutput) ToDeviceGroupOutput() DeviceGroupOutput {
	return o
}

func (o DeviceGroupOutput) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return o
}

type DeviceGroupArrayOutput struct{ *pulumi.OutputState }

func (DeviceGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceGroup)(nil)).Elem()
}

func (o DeviceGroupArrayOutput) ToDeviceGroupArrayOutput() DeviceGroupArrayOutput {
	return o
}

func (o DeviceGroupArrayOutput) ToDeviceGroupArrayOutputWithContext(ctx context.Context) DeviceGroupArrayOutput {
	return o
}

func (o DeviceGroupArrayOutput) Index(i pulumi.IntInput) DeviceGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceGroup {
		return vs[0].([]*DeviceGroup)[vs[1].(int)]
	}).(DeviceGroupOutput)
}

type DeviceGroupMapOutput struct{ *pulumi.OutputState }

func (DeviceGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceGroup)(nil)).Elem()
}

func (o DeviceGroupMapOutput) ToDeviceGroupMapOutput() DeviceGroupMapOutput {
	return o
}

func (o DeviceGroupMapOutput) ToDeviceGroupMapOutputWithContext(ctx context.Context) DeviceGroupMapOutput {
	return o
}

func (o DeviceGroupMapOutput) MapIndex(k pulumi.StringInput) DeviceGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceGroup {
		return vs[0].(map[string]*DeviceGroup)[vs[1].(string)]
	}).(DeviceGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceGroupInput)(nil)).Elem(), &DeviceGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceGroupArrayInput)(nil)).Elem(), DeviceGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceGroupMapInput)(nil)).Elem(), DeviceGroupMap{})
	pulumi.RegisterOutputType(DeviceGroupOutput{})
	pulumi.RegisterOutputType(DeviceGroupArrayOutput{})
	pulumi.RegisterOutputType(DeviceGroupMapOutput{})
}
