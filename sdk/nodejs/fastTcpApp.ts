// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `f5bigip.FastTcpApp` This resource will create and manage FAST TCP applications on BIG-IP from provided JSON declaration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const fast_tcp_app = new f5bigip.FastTcpApp("fast-tcp-app", {
 *     application: "tcp_app_2",
 *     poolMembers: [{
 *         addresses: [
 *             "10.11.34.65",
 *             "56.43.23.76",
 *         ],
 *         connectionLimit: 4,
 *         port: 443,
 *         priorityGroup: 1,
 *         shareNodes: true,
 *     }],
 *     tenant: "tcp_app_tenant",
 *     virtualServer: {
 *         ip: "11.12.16.30",
 *         port: 443,
 *     },
 * });
 * ```
 */
export class FastTcpApp extends pulumi.CustomResource {
    /**
     * Get an existing FastTcpApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FastTcpAppState, opts?: pulumi.CustomResourceOptions): FastTcpApp {
        return new FastTcpApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/fastTcpApp:FastTcpApp';

    /**
     * Returns true if the given object is an instance of FastTcpApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FastTcpApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FastTcpApp.__pulumiType;
    }

    /**
     * Name of the FAST TCP application.
     */
    public readonly application!: pulumi.Output<string>;
    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     */
    public readonly existingMonitor!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing BIG-IP pool.
     */
    public readonly existingPool!: pulumi.Output<string | undefined>;
    /**
     * Name of an existing BIG-IP SNAT pool.
     */
    public readonly existingSnatPool!: pulumi.Output<string | undefined>;
    /**
     * Json payload for FAST TCP application.
     */
    public /*out*/ readonly fastTcpJson!: pulumi.Output<string>;
    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     */
    public readonly loadBalancingMode!: pulumi.Output<string | undefined>;
    /**
     * block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     */
    public readonly monitor!: pulumi.Output<outputs.FastTcpAppMonitor | undefined>;
    /**
     * block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     */
    public readonly poolMembers!: pulumi.Output<outputs.FastTcpAppPoolMember[] | undefined>;
    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     */
    public readonly slowRampTime!: pulumi.Output<number | undefined>;
    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     */
    public readonly snatPoolAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the FAST TCP application tenant.
     */
    public readonly tenant!: pulumi.Output<string>;
    /**
     * block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     */
    public readonly virtualServer!: pulumi.Output<outputs.FastTcpAppVirtualServer | undefined>;

    /**
     * Create a FastTcpApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FastTcpAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FastTcpAppArgs | FastTcpAppState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FastTcpAppState | undefined;
            resourceInputs["application"] = state ? state.application : undefined;
            resourceInputs["existingMonitor"] = state ? state.existingMonitor : undefined;
            resourceInputs["existingPool"] = state ? state.existingPool : undefined;
            resourceInputs["existingSnatPool"] = state ? state.existingSnatPool : undefined;
            resourceInputs["fastTcpJson"] = state ? state.fastTcpJson : undefined;
            resourceInputs["loadBalancingMode"] = state ? state.loadBalancingMode : undefined;
            resourceInputs["monitor"] = state ? state.monitor : undefined;
            resourceInputs["poolMembers"] = state ? state.poolMembers : undefined;
            resourceInputs["slowRampTime"] = state ? state.slowRampTime : undefined;
            resourceInputs["snatPoolAddresses"] = state ? state.snatPoolAddresses : undefined;
            resourceInputs["tenant"] = state ? state.tenant : undefined;
            resourceInputs["virtualServer"] = state ? state.virtualServer : undefined;
        } else {
            const args = argsOrState as FastTcpAppArgs | undefined;
            if ((!args || args.application === undefined) && !opts.urn) {
                throw new Error("Missing required property 'application'");
            }
            if ((!args || args.tenant === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenant'");
            }
            resourceInputs["application"] = args ? args.application : undefined;
            resourceInputs["existingMonitor"] = args ? args.existingMonitor : undefined;
            resourceInputs["existingPool"] = args ? args.existingPool : undefined;
            resourceInputs["existingSnatPool"] = args ? args.existingSnatPool : undefined;
            resourceInputs["loadBalancingMode"] = args ? args.loadBalancingMode : undefined;
            resourceInputs["monitor"] = args ? args.monitor : undefined;
            resourceInputs["poolMembers"] = args ? args.poolMembers : undefined;
            resourceInputs["slowRampTime"] = args ? args.slowRampTime : undefined;
            resourceInputs["snatPoolAddresses"] = args ? args.snatPoolAddresses : undefined;
            resourceInputs["tenant"] = args ? args.tenant : undefined;
            resourceInputs["virtualServer"] = args ? args.virtualServer : undefined;
            resourceInputs["fastTcpJson"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FastTcpApp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FastTcpApp resources.
 */
export interface FastTcpAppState {
    /**
     * Name of the FAST TCP application.
     */
    application?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     */
    existingMonitor?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP pool.
     */
    existingPool?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP SNAT pool.
     */
    existingSnatPool?: pulumi.Input<string>;
    /**
     * Json payload for FAST TCP application.
     */
    fastTcpJson?: pulumi.Input<string>;
    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     */
    loadBalancingMode?: pulumi.Input<string>;
    /**
     * block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     */
    monitor?: pulumi.Input<inputs.FastTcpAppMonitor>;
    /**
     * block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     */
    poolMembers?: pulumi.Input<pulumi.Input<inputs.FastTcpAppPoolMember>[]>;
    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     */
    slowRampTime?: pulumi.Input<number>;
    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     */
    snatPoolAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the FAST TCP application tenant.
     */
    tenant?: pulumi.Input<string>;
    /**
     * block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     */
    virtualServer?: pulumi.Input<inputs.FastTcpAppVirtualServer>;
}

/**
 * The set of arguments for constructing a FastTcpApp resource.
 */
export interface FastTcpAppArgs {
    /**
     * Name of the FAST TCP application.
     */
    application: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP HTTPS pool monitor. Monitors are used to determine the health of the application on each server.
     */
    existingMonitor?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP pool.
     */
    existingPool?: pulumi.Input<string>;
    /**
     * Name of an existing BIG-IP SNAT pool.
     */
    existingSnatPool?: pulumi.Input<string>;
    /**
     * A `load balancing method` is an algorithm that the BIG-IP system uses to select a pool member for processing a request. F5 recommends the Least Connections load balancing method
     */
    loadBalancingMode?: pulumi.Input<string>;
    /**
     * block takes input for FAST-Generated Pool Monitor.
     * See Pool Monitor below for more details.
     */
    monitor?: pulumi.Input<inputs.FastTcpAppMonitor>;
    /**
     * block takes input for FAST-Generated Pool.
     * See Pool Members below for more details.
     */
    poolMembers?: pulumi.Input<pulumi.Input<inputs.FastTcpAppPoolMember>[]>;
    /**
     * Slow ramp temporarily throttles the number of connections to a new pool member. The recommended value is 300 seconds
     */
    slowRampTime?: pulumi.Input<number>;
    /**
     * List of address to be used for FAST-Generated SNAT Pool.
     */
    snatPoolAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the FAST TCP application tenant.
     */
    tenant: pulumi.Input<string>;
    /**
     * block will provide `ip` and `port` options to be used for virtual server.
     * See virtual server below for more details.
     */
    virtualServer?: pulumi.Input<inputs.FastTcpAppVirtualServer>;
}
