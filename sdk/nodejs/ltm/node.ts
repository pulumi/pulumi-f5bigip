// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Node extends pulumi.CustomResource {
    /**
     * Get an existing Node resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * Default is "user-up" you can set to "user-down" if you want to disable
     */
    public readonly state!: pulumi.Output<string | undefined>;

    /**
     * Create a Node resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodeArgs | NodeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NodeState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["connectionLimit"] = state ? state.connectionLimit : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dynamicRatio"] = state ? state.dynamicRatio : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["monitor"] = state ? state.monitor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rateLimit"] = state ? state.rateLimit : undefined;
            inputs["ratio"] = state ? state.ratio : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as NodeArgs | undefined;
            if (!args || args.address === undefined) {
                throw new Error("Missing required property 'address'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["address"] = args ? args.address : undefined;
            inputs["connectionLimit"] = args ? args.connectionLimit : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dynamicRatio"] = args ? args.dynamicRatio : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["monitor"] = args ? args.monitor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rateLimit"] = args ? args.rateLimit : undefined;
            inputs["ratio"] = args ? args.ratio : undefined;
            inputs["state"] = args ? args.state : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Node.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Node resources.
 */
export interface NodeState {
    /**
     * IP or hostname of the node
     */
    readonly address?: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections allowed for the node or node address.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * User-defined description give ltm_node
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the fixed ratio value used for a node during ratio load balancing.
     */
    readonly dynamicRatio?: pulumi.Input<number>;
    readonly fqdn?: pulumi.Input<inputs.ltm.NodeFqdn>;
    /**
     * specifies the name of the monitor or monitor rule that you want to associate with the node.
     */
    readonly monitor?: pulumi.Input<string>;
    /**
     * Name of the node
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
     */
    readonly rateLimit?: pulumi.Input<string>;
    /**
     * Sets the ratio number for the node.
     */
    readonly ratio?: pulumi.Input<number>;
    /**
     * Default is "user-up" you can set to "user-down" if you want to disable
     */
    readonly state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Node resource.
 */
export interface NodeArgs {
    /**
     * IP or hostname of the node
     */
    readonly address: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections allowed for the node or node address.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * User-defined description give ltm_node
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the fixed ratio value used for a node during ratio load balancing.
     */
    readonly dynamicRatio?: pulumi.Input<number>;
    readonly fqdn?: pulumi.Input<inputs.ltm.NodeFqdn>;
    /**
     * specifies the name of the monitor or monitor rule that you want to associate with the node.
     */
    readonly monitor?: pulumi.Input<string>;
    /**
     * Name of the node
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
     */
    readonly rateLimit?: pulumi.Input<string>;
    /**
     * Sets the ratio number for the node.
     */
    readonly ratio?: pulumi.Input<number>;
    /**
     * Default is "user-up" you can set to "user-down" if you want to disable
     */
    readonly state?: pulumi.Input<string>;
}
