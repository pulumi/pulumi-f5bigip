// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceState, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, <any>state, { ...opts, id: id });
    }

    /**
     * IP address used for config sync
     */
    public readonly configsyncIp: pulumi.Output<string>;
    /**
     * IP address used for state mirroring
     */
    public readonly mirrorIp: pulumi.Output<string | undefined>;
    /**
     * Secondary IP address used for state mirroring
     */
    public readonly mirrorSecondaryIp: pulumi.Output<string | undefined>;
    /**
     * Address of the Device which needs to be Deviceensed
     */
    public readonly name: pulumi.Output<string>;

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
            const state: DeviceState = argsOrState as DeviceState | undefined;
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
        super("f5bigip:cm/device:Device", name, inputs, opts);
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
