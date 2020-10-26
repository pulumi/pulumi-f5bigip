// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.Pool` Manages a pool configuration.
 *
 * Resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const monitor = new f5bigip.ltm.Monitor("monitor", {
 *     name: "/Common/terraform_monitor",
 *     parent: "/Common/http",
 * });
 * const pool = new f5bigip.ltm.Pool("pool", {
 *     name: "/Common/Axiom_Environment_APP1_Pool",
 *     loadBalancingMode: "round-robin",
 *     minimumActiveMembers: 1,
 *     monitors: [monitor.name],
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
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PoolState, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/pool:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    /**
     * Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
     */
    public readonly allowNat!: pulumi.Output<string>;
    /**
     * Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
     */
    public readonly allowSnat!: pulumi.Output<string>;
    /**
     * Specifies descriptive text that identifies the pool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the load balancing method. The default is Round Robin.
     */
    public readonly loadBalancingMode!: pulumi.Output<string>;
    /**
     * Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
     */
    public readonly minimumActiveMembers!: pulumi.Output<number>;
    /**
     * List of monitor names to associate with the pool
     */
    public readonly monitors!: pulumi.Output<string[]>;
    /**
     * Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the number of times the system tries to contact a new pool member after a passive failure.
     */
    public readonly reselectTries!: pulumi.Output<number>;
    /**
     * Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
     */
    public readonly serviceDownAction!: pulumi.Output<string>;
    /**
     * Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
     */
    public readonly slowRampTime!: pulumi.Output<number>;

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
            const state = argsOrState as PoolState | undefined;
            inputs["allowNat"] = state ? state.allowNat : undefined;
            inputs["allowSnat"] = state ? state.allowSnat : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["loadBalancingMode"] = state ? state.loadBalancingMode : undefined;
            inputs["minimumActiveMembers"] = state ? state.minimumActiveMembers : undefined;
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
            inputs["description"] = args ? args.description : undefined;
            inputs["loadBalancingMode"] = args ? args.loadBalancingMode : undefined;
            inputs["minimumActiveMembers"] = args ? args.minimumActiveMembers : undefined;
            inputs["monitors"] = args ? args.monitors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["reselectTries"] = args ? args.reselectTries : undefined;
            inputs["serviceDownAction"] = args ? args.serviceDownAction : undefined;
            inputs["slowRampTime"] = args ? args.slowRampTime : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Pool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pool resources.
 */
export interface PoolState {
    /**
     * Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
     */
    readonly allowNat?: pulumi.Input<string>;
    /**
     * Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
     */
    readonly allowSnat?: pulumi.Input<string>;
    /**
     * Specifies descriptive text that identifies the pool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the load balancing method. The default is Round Robin.
     */
    readonly loadBalancingMode?: pulumi.Input<string>;
    /**
     * Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
     */
    readonly minimumActiveMembers?: pulumi.Input<number>;
    /**
     * List of monitor names to associate with the pool
     */
    readonly monitors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the number of times the system tries to contact a new pool member after a passive failure.
     */
    readonly reselectTries?: pulumi.Input<number>;
    /**
     * Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
     */
    readonly serviceDownAction?: pulumi.Input<string>;
    /**
     * Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
     */
    readonly slowRampTime?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
     */
    readonly allowNat?: pulumi.Input<string>;
    /**
     * Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
     */
    readonly allowSnat?: pulumi.Input<string>;
    /**
     * Specifies descriptive text that identifies the pool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the load balancing method. The default is Round Robin.
     */
    readonly loadBalancingMode?: pulumi.Input<string>;
    /**
     * Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
     */
    readonly minimumActiveMembers?: pulumi.Input<number>;
    /**
     * List of monitor names to associate with the pool
     */
    readonly monitors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies the number of times the system tries to contact a new pool member after a passive failure.
     */
    readonly reselectTries?: pulumi.Input<number>;
    /**
     * Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
     */
    readonly serviceDownAction?: pulumi.Input<string>;
    /**
     * Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
     */
    readonly slowRampTime?: pulumi.Input<number>;
}
