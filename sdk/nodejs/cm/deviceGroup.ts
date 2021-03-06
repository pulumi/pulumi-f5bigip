// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * `f5bigip.cm.DeviceGroup` A device group is a collection of BIG-IP devices that are configured to securely synchronize their BIG-IP configuration data, and fail over when needed.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const myNewDevicegroup = new f5bigip.cm.DeviceGroup("my_new_devicegroup", {
 *     autoSync: "enabled",
 *     devices: [
 *         {
 *             name: "bigip1.cisco.com",
 *         },
 *         {
 *             name: "bigip200.f5.com",
 *         },
 *     ],
 *     fullLoadOnSync: "true",
 *     name: "sanjose_devicegroup",
 *     type: "sync-only",
 * });
 * ```
 */
export class DeviceGroup extends pulumi.CustomResource {
    /**
     * Get an existing DeviceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceGroupState, opts?: pulumi.CustomResourceOptions): DeviceGroup {
        return new DeviceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:cm/deviceGroup:DeviceGroup';

    /**
     * Returns true if the given object is an instance of DeviceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceGroup.__pulumiType;
    }

    /**
     * Specifies if the device-group will automatically sync configuration data to its members
     */
    public readonly autoSync!: pulumi.Output<string | undefined>;
    /**
     * Description of Device group
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the device to be included in device group, this need to be configured before using devicegroup resource
     */
    public readonly devices!: pulumi.Output<outputs.cm.DeviceGroupDevice[] | undefined>;
    /**
     * Specifies if the device-group will perform a full-load upon sync
     */
    public readonly fullLoadOnSync!: pulumi.Output<string | undefined>;
    /**
     * Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
     */
    public readonly incrementalConfig!: pulumi.Output<number | undefined>;
    /**
     * Is the name of the device Group
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Specifies if the device-group will use a network connection for failover
     */
    public readonly networkFailover!: pulumi.Output<string | undefined>;
    /**
     * Device administrative partition
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the configuration should be saved upon auto-sync.
     */
    public readonly saveOnAutoSync!: pulumi.Output<string | undefined>;
    /**
     * Specifies if the device-group will be used for failover or resource syncing
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a DeviceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DeviceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceGroupArgs | DeviceGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceGroupState | undefined;
            inputs["autoSync"] = state ? state.autoSync : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["devices"] = state ? state.devices : undefined;
            inputs["fullLoadOnSync"] = state ? state.fullLoadOnSync : undefined;
            inputs["incrementalConfig"] = state ? state.incrementalConfig : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkFailover"] = state ? state.networkFailover : undefined;
            inputs["partition"] = state ? state.partition : undefined;
            inputs["saveOnAutoSync"] = state ? state.saveOnAutoSync : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DeviceGroupArgs | undefined;
            inputs["autoSync"] = args ? args.autoSync : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["devices"] = args ? args.devices : undefined;
            inputs["fullLoadOnSync"] = args ? args.fullLoadOnSync : undefined;
            inputs["incrementalConfig"] = args ? args.incrementalConfig : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkFailover"] = args ? args.networkFailover : undefined;
            inputs["partition"] = args ? args.partition : undefined;
            inputs["saveOnAutoSync"] = args ? args.saveOnAutoSync : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeviceGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceGroup resources.
 */
export interface DeviceGroupState {
    /**
     * Specifies if the device-group will automatically sync configuration data to its members
     */
    readonly autoSync?: pulumi.Input<string>;
    /**
     * Description of Device group
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the device to be included in device group, this need to be configured before using devicegroup resource
     */
    readonly devices?: pulumi.Input<pulumi.Input<inputs.cm.DeviceGroupDevice>[]>;
    /**
     * Specifies if the device-group will perform a full-load upon sync
     */
    readonly fullLoadOnSync?: pulumi.Input<string>;
    /**
     * Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
     */
    readonly incrementalConfig?: pulumi.Input<number>;
    /**
     * Is the name of the device Group
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies if the device-group will use a network connection for failover
     */
    readonly networkFailover?: pulumi.Input<string>;
    /**
     * Device administrative partition
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Specifies whether the configuration should be saved upon auto-sync.
     */
    readonly saveOnAutoSync?: pulumi.Input<string>;
    /**
     * Specifies if the device-group will be used for failover or resource syncing
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeviceGroup resource.
 */
export interface DeviceGroupArgs {
    /**
     * Specifies if the device-group will automatically sync configuration data to its members
     */
    readonly autoSync?: pulumi.Input<string>;
    /**
     * Description of Device group
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the device to be included in device group, this need to be configured before using devicegroup resource
     */
    readonly devices?: pulumi.Input<pulumi.Input<inputs.cm.DeviceGroupDevice>[]>;
    /**
     * Specifies if the device-group will perform a full-load upon sync
     */
    readonly fullLoadOnSync?: pulumi.Input<string>;
    /**
     * Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
     */
    readonly incrementalConfig?: pulumi.Input<number>;
    /**
     * Is the name of the device Group
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies if the device-group will use a network connection for failover
     */
    readonly networkFailover?: pulumi.Input<string>;
    /**
     * Device administrative partition
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * Specifies whether the configuration should be saved upon auto-sync.
     */
    readonly saveOnAutoSync?: pulumi.Input<string>;
    /**
     * Specifies if the device-group will be used for failover or resource syncing
     */
    readonly type?: pulumi.Input<string>;
}
