// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `bigip_ltm_pool` Manages a pool configuration.
 * 
 * Resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 * 
 * 
 * ## Example Usage
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * 
 * const bigip_ltm_pool_pool = new f5bigip.ltm.Pool("pool", {
 *     allowNat: "yes",
 *     allowSnat: "yes",
 *     loadBalancingMode: "round-robin",
 *     monitors: [
 *         bigip_ltm_monitor_monitor.name,
 *         bigip_ltm_monitor_monitor2.name,
 *     ],
 *     name: "/Common/terraform-pool",
 * });
 * ```
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /**
     * Allow NAT
     */
    public readonly allowNat: pulumi.Output<string | undefined>;
    /**
     * Allow SNAT
     */
    public readonly allowSnat: pulumi.Output<string | undefined>;
    /**
     * Possible values: round-robin, ...
     */
    public readonly loadBalancingMode: pulumi.Output<string | undefined>;
    /**
     * List of monitor names to associate with the pool
     */
    public readonly monitors: pulumi.Output<string[] | undefined>;
    /**
     * Name of the pool
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Number of times the system tries to select a new pool member after a failure.
     */
    public readonly reselectTries: pulumi.Output<number | undefined>;
    /**
     * Possible values: none, reset, reselect, drop
     */
    public readonly serviceDownAction: pulumi.Output<string | undefined>;
    /**
     * Slow ramp time for pool members
     */
    public readonly slowRampTime: pulumi.Output<number | undefined>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PoolArgs | PoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: PoolState = argsOrState as PoolState | undefined;
            inputs["allowNat"] = state ? state.allowNat : undefined;
            inputs["allowSnat"] = state ? state.allowSnat : undefined;
            inputs["loadBalancingMode"] = state ? state.loadBalancingMode : undefined;
            inputs["monitors"] = state ? state.monitors : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["reselectTries"] = state ? state.reselectTries : undefined;
            inputs["serviceDownAction"] = state ? state.serviceDownAction : undefined;
            inputs["slowRampTime"] = state ? state.slowRampTime : undefined;
        } else {
            const args = argsOrState as PoolArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["allowNat"] = args ? args.allowNat : undefined;
            inputs["allowSnat"] = args ? args.allowSnat : undefined;
            inputs["loadBalancingMode"] = args ? args.loadBalancingMode : undefined;
            inputs["monitors"] = args ? args.monitors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["reselectTries"] = args ? args.reselectTries : undefined;
            inputs["serviceDownAction"] = args ? args.serviceDownAction : undefined;
            inputs["slowRampTime"] = args ? args.slowRampTime : undefined;
        }
        super("f5bigip:ltm/pool:Pool", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    /**
     * Allow NAT
     */
    readonly allowNat?: pulumi.Input<string>;
    /**
     * Allow SNAT
     */
    readonly allowSnat?: pulumi.Input<string>;
    /**
     * Possible values: round-robin, ...
     */
    readonly loadBalancingMode?: pulumi.Input<string>;
    /**
     * List of monitor names to associate with the pool
     */
    readonly monitors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the pool
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Number of times the system tries to select a new pool member after a failure.
     */
    readonly reselectTries?: pulumi.Input<number>;
    /**
     * Possible values: none, reset, reselect, drop
     */
    readonly serviceDownAction?: pulumi.Input<string>;
    /**
     * Slow ramp time for pool members
     */
    readonly slowRampTime?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * Allow NAT
     */
    readonly allowNat?: pulumi.Input<string>;
    /**
     * Allow SNAT
     */
    readonly allowSnat?: pulumi.Input<string>;
    /**
     * Possible values: round-robin, ...
     */
    readonly loadBalancingMode?: pulumi.Input<string>;
    /**
     * List of monitor names to associate with the pool
     */
    readonly monitors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the pool
     */
    readonly name: pulumi.Input<string>;
    /**
     * Number of times the system tries to select a new pool member after a failure.
     */
    readonly reselectTries?: pulumi.Input<number>;
    /**
     * Possible values: none, reset, reselect, drop
     */
    readonly serviceDownAction?: pulumi.Input<string>;
    /**
     * Slow ramp time for pool members
     */
    readonly slowRampTime?: pulumi.Input<number>;
}
