// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.Partition` Manages F5 BIG-IP partitions
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test_partition = new f5bigip.Partition("test-partition", {
 *     description: "created by terraform",
 *     name: "test-partition",
 *     routeDomainId: 2,
 * });
 * ```
 */
export class Partition extends pulumi.CustomResource {
    /**
     * Get an existing Partition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PartitionState, opts?: pulumi.CustomResourceOptions): Partition {
        return new Partition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/partition:Partition';

    /**
     * Returns true if the given object is an instance of Partition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Partition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Partition.__pulumiType;
    }

    /**
     * Description of the partition.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the partition.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Route domain id of the partition.
     */
    public readonly routeDomainId!: pulumi.Output<number | undefined>;

    /**
     * Create a Partition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PartitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PartitionArgs | PartitionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PartitionState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["routeDomainId"] = state ? state.routeDomainId : undefined;
        } else {
            const args = argsOrState as PartitionArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["routeDomainId"] = args ? args.routeDomainId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Partition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Partition resources.
 */
export interface PartitionState {
    /**
     * Description of the partition.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the partition.
     */
    name?: pulumi.Input<string>;
    /**
     * Route domain id of the partition.
     */
    routeDomainId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Partition resource.
 */
export interface PartitionArgs {
    /**
     * Description of the partition.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the partition.
     */
    name: pulumi.Input<string>;
    /**
     * Route domain id of the partition.
     */
    routeDomainId?: pulumi.Input<number>;
}
