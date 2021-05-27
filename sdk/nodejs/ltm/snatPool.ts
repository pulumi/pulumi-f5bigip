// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.SnatPool` Collections of SNAT translation addresses
 *
 * Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-snatpool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const snatpoolSanjose = new f5bigip.ltm.SnatPool("snatpool_sanjose", {
 *     members: [
 *         "191.1.1.1",
 *         "194.2.2.2",
 *     ],
 *     name: "/Common/snatpool_sanjose",
 * });
 * ```
 */
export class SnatPool extends pulumi.CustomResource {
    /**
     * Get an existing SnatPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnatPoolState, opts?: pulumi.CustomResourceOptions): SnatPool {
        return new SnatPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/snatPool:SnatPool';

    /**
     * Returns true if the given object is an instance of SnatPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnatPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnatPool.__pulumiType;
    }

    /**
     * Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     */
    public readonly members!: pulumi.Output<string[]>;
    /**
     * Name of the snatpool
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a SnatPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnatPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnatPoolArgs | SnatPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnatPoolState | undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SnatPoolArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SnatPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnatPool resources.
 */
export interface SnatPoolState {
    /**
     * Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the snatpool
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnatPool resource.
 */
export interface SnatPoolArgs {
    /**
     * Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     */
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the snatpool
     */
    name: pulumi.Input<string>;
}
