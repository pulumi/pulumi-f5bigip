// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.cm.Device` provides details about a specific bigip
 *
 * This resource is helpful when configuring the BIG-IP device in cluster or in HA mode.
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const myNewDevice = new f5bigip.cm.Device("my_new_device", {
 *     configsyncIp: "2.2.2.2",
 *     mirrorIp: "10.10.10.10",
 *     mirrorSecondaryIp: "11.11.11.11",
 *     name: "bigip300.f5.com",
 * });
 * ```
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceState, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:cm/device:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * IP address used for config sync
     */
    public readonly configsyncIp!: pulumi.Output<string>;
    /**
     * IP address used for state mirroring
     */
    public readonly mirrorIp!: pulumi.Output<string | undefined>;
    /**
     * Secondary IP address used for state mirroring
     */
    public readonly mirrorSecondaryIp!: pulumi.Output<string | undefined>;
    /**
     * Address of the Device which needs to be Deviceensed
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceArgs | DeviceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DeviceState | undefined;
            inputs["configsyncIp"] = state ? state.configsyncIp : undefined;
            inputs["mirrorIp"] = state ? state.mirrorIp : undefined;
            inputs["mirrorSecondaryIp"] = state ? state.mirrorSecondaryIp : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DeviceArgs | undefined;
            if (!args || args.configsyncIp === undefined) {
                throw new Error("Missing required property 'configsyncIp'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["configsyncIp"] = args ? args.configsyncIp : undefined;
            inputs["mirrorIp"] = args ? args.mirrorIp : undefined;
            inputs["mirrorSecondaryIp"] = args ? args.mirrorSecondaryIp : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Device.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Device resources.
 */
export interface DeviceState {
    /**
     * IP address used for config sync
     */
    readonly configsyncIp?: pulumi.Input<string>;
    /**
     * IP address used for state mirroring
     */
    readonly mirrorIp?: pulumi.Input<string>;
    /**
     * Secondary IP address used for state mirroring
     */
    readonly mirrorSecondaryIp?: pulumi.Input<string>;
    /**
     * Address of the Device which needs to be Deviceensed
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * IP address used for config sync
     */
    readonly configsyncIp: pulumi.Input<string>;
    /**
     * IP address used for state mirroring
     */
    readonly mirrorIp?: pulumi.Input<string>;
    /**
     * Secondary IP address used for state mirroring
     */
    readonly mirrorSecondaryIp?: pulumi.Input<string>;
    /**
     * Address of the Device which needs to be Deviceensed
     */
    readonly name: pulumi.Input<string>;
}
