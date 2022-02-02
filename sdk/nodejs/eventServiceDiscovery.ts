// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class EventServiceDiscovery extends pulumi.CustomResource {
    /**
     * Get an existing EventServiceDiscovery resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventServiceDiscoveryState, opts?: pulumi.CustomResourceOptions): EventServiceDiscovery {
        return new EventServiceDiscovery(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/eventServiceDiscovery:EventServiceDiscovery';

    /**
     * Returns true if the given object is an instance of EventServiceDiscovery.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventServiceDiscovery {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventServiceDiscovery.__pulumiType;
    }

    /**
     * Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     */
    public readonly nodes!: pulumi.Output<outputs.EventServiceDiscoveryNode[] | undefined>;
    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     */
    public readonly taskid!: pulumi.Output<string>;

    /**
     * Create a EventServiceDiscovery resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventServiceDiscoveryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventServiceDiscoveryArgs | EventServiceDiscoveryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventServiceDiscoveryState | undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["taskid"] = state ? state.taskid : undefined;
        } else {
            const args = argsOrState as EventServiceDiscoveryArgs | undefined;
            if ((!args || args.taskid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskid'");
            }
            resourceInputs["nodes"] = args ? args.nodes : undefined;
            resourceInputs["taskid"] = args ? args.taskid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventServiceDiscovery.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventServiceDiscovery resources.
 */
export interface EventServiceDiscoveryState {
    /**
     * Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.EventServiceDiscoveryNode>[]>;
    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     */
    taskid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventServiceDiscovery resource.
 */
export interface EventServiceDiscoveryArgs {
    /**
     * Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.EventServiceDiscoveryNode>[]>;
    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     */
    taskid: pulumi.Input<string>;
}
