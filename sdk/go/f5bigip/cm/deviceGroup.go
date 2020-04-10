// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cm

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `cm.DeviceGroup` A device group is a collection of BIG-IP devices that are configured to securely synchronize their BIG-IP configuration data, and fail over when needed.
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
