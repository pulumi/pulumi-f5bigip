// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.net.SelfIp` Manages a selfip configuration
 *
 * Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-selfip.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const selfip1 = new f5bigip.net.SelfIp("selfip1", {
 *     name: "/Common/internalselfIP",
 *     ip: "11.1.1.1/24",
 *     vlan: "/Common/internal",
 *     trafficGroup: "traffic-group-1",
 * }, {
 *     dependsOn: [bigip_net_vlan.vlan1],
 * });
 * ```
 */
export class SelfIp extends pulumi.CustomResource {
    /**
     * Get an existing SelfIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SelfIpState, opts?: pulumi.CustomResourceOptions): SelfIp {
        return new SelfIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:net/selfIp:SelfIp';

    /**
     * Returns true if the given object is an instance of SelfIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SelfIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SelfIp.__pulumiType;
    }

    /**
     * The Self IP's address and netmask.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Name of the selfip
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     */
    public readonly trafficGroup!: pulumi.Output<string | undefined>;
    /**
     * Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     */
    public readonly vlan!: pulumi.Output<string>;

    /**
     * Create a SelfIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SelfIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SelfIpArgs | SelfIpState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SelfIpState | undefined;
            inputs["ip"] = state ? state.ip : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["trafficGroup"] = state ? state.trafficGroup : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as SelfIpArgs | undefined;
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.vlan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vlan'");
            }
            inputs["ip"] = args ? args.ip : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["trafficGroup"] = args ? args.trafficGroup : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SelfIp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SelfIp resources.
 */
export interface SelfIpState {
    /**
     * The Self IP's address and netmask.
     */
    readonly ip?: pulumi.Input<string>;
    /**
     * Name of the selfip
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     */
    readonly trafficGroup?: pulumi.Input<string>;
    /**
     * Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     */
    readonly vlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SelfIp resource.
 */
export interface SelfIpArgs {
    /**
     * The Self IP's address and netmask.
     */
    readonly ip: pulumi.Input<string>;
    /**
     * Name of the selfip
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     */
    readonly trafficGroup?: pulumi.Input<string>;
    /**
     * Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     */
    readonly vlan: pulumi.Input<string>;
}
