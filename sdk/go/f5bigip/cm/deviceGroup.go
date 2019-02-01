// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cm

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `bigip_cm_devicegroup` A device group is a collection of BIG-IP devices that are configured to securely synchronize their BIG-IP configuration data, and fail over when needed.
type DeviceGroup struct {
	s *pulumi.ResourceState
}

// NewDeviceGroup registers a new resource with the given unique name, arguments, and options.
func NewDeviceGroup(ctx *pulumi.Context,
	name string, args *DeviceGroupArgs, opts ...pulumi.ResourceOpt) (*DeviceGroup, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoSync"] = nil
		inputs["description"] = nil
		inputs["devices"] = nil
		inputs["fullLoadOnSync"] = nil
		inputs["incrementalConfig"] = nil
		inputs["name"] = nil
		inputs["networkFailover"] = nil
		inputs["partition"] = nil
		inputs["saveOnAutoSync"] = nil
		inputs["type"] = nil
	} else {
		inputs["autoSync"] = args.AutoSync
		inputs["description"] = args.Description
		inputs["devices"] = args.Devices
		inputs["fullLoadOnSync"] = args.FullLoadOnSync
		inputs["incrementalConfig"] = args.IncrementalConfig
		inputs["name"] = args.Name
		inputs["networkFailover"] = args.NetworkFailover
		inputs["partition"] = args.Partition
		inputs["saveOnAutoSync"] = args.SaveOnAutoSync
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("f5bigip:cm/deviceGroup:DeviceGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DeviceGroup{s: s}, nil
}

// GetDeviceGroup gets an existing DeviceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DeviceGroupState, opts ...pulumi.ResourceOpt) (*DeviceGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoSync"] = state.AutoSync
		inputs["description"] = state.Description
		inputs["devices"] = state.Devices
		inputs["fullLoadOnSync"] = state.FullLoadOnSync
		inputs["incrementalConfig"] = state.IncrementalConfig
		inputs["name"] = state.Name
		inputs["networkFailover"] = state.NetworkFailover
		inputs["partition"] = state.Partition
		inputs["saveOnAutoSync"] = state.SaveOnAutoSync
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("f5bigip:cm/deviceGroup:DeviceGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DeviceGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DeviceGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DeviceGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies if the device-group will automatically sync configuration data to its members
func (r *DeviceGroup) AutoSync() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["autoSync"])
}

// Description of Device group
func (r *DeviceGroup) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Name of the device to be included in device group, this need to be configured before using devicegroup resource
func (r *DeviceGroup) Devices() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["devices"])
}

// Specifies if the device-group will perform a full-load upon sync
func (r *DeviceGroup) FullLoadOnSync() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fullLoadOnSync"])
}

// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
func (r *DeviceGroup) IncrementalConfig() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["incrementalConfig"])
}

// Is the name of the device Group
func (r *DeviceGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specifies if the device-group will use a network connection for failover
func (r *DeviceGroup) NetworkFailover() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkFailover"])
}

// Device administrative partition
func (r *DeviceGroup) Partition() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["partition"])
}

// Specifies whether the configuration should be saved upon auto-sync.
func (r *DeviceGroup) SaveOnAutoSync() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["saveOnAutoSync"])
}

// Specifies if the device-group will be used for failover or resource syncing
func (r *DeviceGroup) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering DeviceGroup resources.
type DeviceGroupState struct {
	// Specifies if the device-group will automatically sync configuration data to its members
	AutoSync interface{}
	// Description of Device group
	Description interface{}
	// Name of the device to be included in device group, this need to be configured before using devicegroup resource
	Devices interface{}
	// Specifies if the device-group will perform a full-load upon sync
	FullLoadOnSync interface{}
	// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
	IncrementalConfig interface{}
	// Is the name of the device Group
	Name interface{}
	// Specifies if the device-group will use a network connection for failover
	NetworkFailover interface{}
	// Device administrative partition
	Partition interface{}
	// Specifies whether the configuration should be saved upon auto-sync.
	SaveOnAutoSync interface{}
	// Specifies if the device-group will be used for failover or resource syncing
	Type interface{}
}

// The set of arguments for constructing a DeviceGroup resource.
type DeviceGroupArgs struct {
	// Specifies if the device-group will automatically sync configuration data to its members
	AutoSync interface{}
	// Description of Device group
	Description interface{}
	// Name of the device to be included in device group, this need to be configured before using devicegroup resource
	Devices interface{}
	// Specifies if the device-group will perform a full-load upon sync
	FullLoadOnSync interface{}
	// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
	IncrementalConfig interface{}
	// Is the name of the device Group
	Name interface{}
	// Specifies if the device-group will use a network connection for failover
	NetworkFailover interface{}
	// Device administrative partition
	Partition interface{}
	// Specifies whether the configuration should be saved upon auto-sync.
	SaveOnAutoSync interface{}
	// Specifies if the device-group will be used for failover or resource syncing
	Type interface{}
}
