// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.net.SelfIp` Manages a selfip configuration
 *
 * Resource should be named with their `full path`. The full path is the combination of the `partition + name of the resource`, for example `/Common/my-selfip`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const vlan1 = new f5bigip.net.Vlan("vlan1", {
 *     name: "/Common/Internal",
 *     tag: 101,
 *     interfaces: [{
 *         vlanport: "1.2",
 *         tagged: false,
 *     }],
 * });
 * const selfip1 = new f5bigip.net.SelfIp("selfip1", {
 *     name: "/Common/internalselfIP",
 *     ip: "11.1.1.1/24",
 *     vlan: "/Common/internal",
 * }, {
 *     dependsOn: [vlan1],
 * });
 * ```
 * ### Example usage with `portLockdown`
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
 *     portLockdowns: [
 *         "tcp:4040",
 *         "udp:5050",
 *         "egp:0",
 *     ],
 * }, {
 *     dependsOn: [bigip_net_vlan.vlan1],
 * });
 * ```
 *
 * ### Example usage with `portLockdown` set to `["none"]`
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
 *     portLockdowns: ["none"],
 * }, {
 *     dependsOn: [bigip_net_vlan.vlan1],
 * });
 * ```
 *
 * ### Example usage with route domain embedded in the `ip`
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const selfip1 = new f5bigip.net.SelfIp("selfip1", {
 *     name: "/Common/internalselfIP",
 *     ip: "11.1.1.1%4/24",
 *     vlan: "/Common/internal",
 *     trafficGroup: "traffic-group-1",
 *     portLockdowns: ["none"],
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
     * The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Name of the selfip
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the port lockdown, defaults to `Allow None` if not specified.
     */
    public readonly portLockdowns!: pulumi.Output<string[] | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SelfIpState | undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["portLockdowns"] = state ? state.portLockdowns : undefined;
            resourceInputs["trafficGroup"] = state ? state.trafficGroup : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
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
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["portLockdowns"] = args ? args.portLockdowns : undefined;
            resourceInputs["trafficGroup"] = args ? args.trafficGroup : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SelfIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SelfIp resources.
 */
export interface SelfIpState {
    /**
     * The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
     */
    ip?: pulumi.Input<string>;
    /**
     * Name of the selfip
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the port lockdown, defaults to `Allow None` if not specified.
     */
    portLockdowns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     */
    trafficGroup?: pulumi.Input<string>;
    /**
     * Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     */
    vlan?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SelfIp resource.
 */
export interface SelfIpArgs {
    /**
     * The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
     */
    ip: pulumi.Input<string>;
    /**
     * Name of the selfip
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the port lockdown, defaults to `Allow None` if not specified.
     */
    portLockdowns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     */
    trafficGroup?: pulumi.Input<string>;
    /**
     * Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     */
    vlan: pulumi.Input<string>;
}
