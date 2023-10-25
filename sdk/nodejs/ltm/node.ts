// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.Node` Manages a node configuration
 *
 * For resources should be named with their `full path`.The full path is the combination of the `partition + name` of the resource( example: `/Common/my-node` ) or `partition + Direcroty + name` of the resource ( example: `/Common/test/my-node` ).
 * When including directory in `full path` we have to make sure it is created in the given partition before using it.
 */
export class Node extends pulumi.CustomResource {
    /**
     * Get an existing Node resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodeState, opts?: pulumi.CustomResourceOptions): Node {
        return new Node(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/node:Node';

    /**
     * Returns true if the given object is an instance of Node.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Node {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Node.__pulumiType;
    }

    /**
     * IP or hostname of the node
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * Specifies the maximum number of connections allowed for the node or node address.
     */
    public readonly connectionLimit!: pulumi.Output<number>;
    /**
     * User-defined description give ltm_node
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the fixed ratio value used for a node during ratio load balancing.
     */
    public readonly dynamicRatio!: pulumi.Output<number>;
    public readonly fqdn!: pulumi.Output<outputs.ltm.NodeFqdn | undefined>;
    /**
     * specifies the name of the monitor or monitor rule that you want to associate with the node.
     */
    public readonly monitor!: pulumi.Output<string | undefined>;
    /**
     * Name of the node
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
     */
    public readonly rateLimit!: pulumi.Output<string>;
    /**
     * Sets the ratio number for the node.
     */
    public readonly ratio!: pulumi.Output<number>;
    /**
     * Enables or disables the node for new sessions. The default value is user-enabled.
     */
    public readonly session!: pulumi.Output<string>;
    /**
     * Default is "user-up" you can set to "user-down" if you want to disable
     *
     * > *NOTE* Below attributes needs to be configured under fqdn option.
     */
    public readonly state!: pulumi.Output<string>;

    /**
     * Create a Node resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeArgs | NodeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodeState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["connectionLimit"] = state ? state.connectionLimit : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicRatio"] = state ? state.dynamicRatio : undefined;
            resourceInputs["fqdn"] = state ? state.fqdn : undefined;
            resourceInputs["monitor"] = state ? state.monitor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rateLimit"] = state ? state.rateLimit : undefined;
            resourceInputs["ratio"] = state ? state.ratio : undefined;
            resourceInputs["session"] = state ? state.session : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as NodeArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["connectionLimit"] = args ? args.connectionLimit : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicRatio"] = args ? args.dynamicRatio : undefined;
            resourceInputs["fqdn"] = args ? args.fqdn : undefined;
            resourceInputs["monitor"] = args ? args.monitor : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rateLimit"] = args ? args.rateLimit : undefined;
            resourceInputs["ratio"] = args ? args.ratio : undefined;
            resourceInputs["session"] = args ? args.session : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Node.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Node resources.
 */
export interface NodeState {
    /**
     * IP or hostname of the node
     */
    address?: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections allowed for the node or node address.
     */
    connectionLimit?: pulumi.Input<number>;
    /**
     * User-defined description give ltm_node
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the fixed ratio value used for a node during ratio load balancing.
     */
    dynamicRatio?: pulumi.Input<number>;
    fqdn?: pulumi.Input<inputs.ltm.NodeFqdn>;
    /**
     * specifies the name of the monitor or monitor rule that you want to associate with the node.
     */
    monitor?: pulumi.Input<string>;
    /**
     * Name of the node
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
     */
    rateLimit?: pulumi.Input<string>;
    /**
     * Sets the ratio number for the node.
     */
    ratio?: pulumi.Input<number>;
    /**
     * Enables or disables the node for new sessions. The default value is user-enabled.
     */
    session?: pulumi.Input<string>;
    /**
     * Default is "user-up" you can set to "user-down" if you want to disable
     *
     * > *NOTE* Below attributes needs to be configured under fqdn option.
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Node resource.
 */
export interface NodeArgs {
    /**
     * IP or hostname of the node
     */
    address: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections allowed for the node or node address.
     */
    connectionLimit?: pulumi.Input<number>;
    /**
     * User-defined description give ltm_node
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the fixed ratio value used for a node during ratio load balancing.
     */
    dynamicRatio?: pulumi.Input<number>;
    fqdn?: pulumi.Input<inputs.ltm.NodeFqdn>;
    /**
     * specifies the name of the monitor or monitor rule that you want to associate with the node.
     */
    monitor?: pulumi.Input<string>;
    /**
     * Name of the node
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
     */
    rateLimit?: pulumi.Input<string>;
    /**
     * Sets the ratio number for the node.
     */
    ratio?: pulumi.Input<number>;
    /**
     * Enables or disables the node for new sessions. The default value is user-enabled.
     */
    session?: pulumi.Input<string>;
    /**
     * Default is "user-up" you can set to "user-down" if you want to disable
     *
     * > *NOTE* Below attributes needs to be configured under fqdn option.
     */
    state?: pulumi.Input<string>;
}
