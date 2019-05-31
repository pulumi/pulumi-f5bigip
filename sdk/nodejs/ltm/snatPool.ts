// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `bigip_ltm_snatpool` Collections of SNAT translation addresses
 * 
 * Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-snatpool. 
 * 
 * 
 * ## Example Usage
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * 
 * const snatpoolSanjose = new f5bigip.LtmSnatpoolpool("snatpool_sanjose", {
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
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnatPoolState, opts?: pulumi.CustomResourceOptions): SnatPool {
        return new SnatPool(name, <any>state, { ...opts, id: id });
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
        if (opts && opts.id) {
            const state = argsOrState as SnatPoolState | undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as SnatPoolArgs | undefined;
            if (!args || args.members === undefined) {
                throw new Error("Missing required property 'members'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super("f5bigip:ltm/snatPool:SnatPool", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnatPool resources.
 */
export interface SnatPoolState {
    /**
     * Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the snatpool
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnatPool resource.
 */
export interface SnatPoolArgs {
    /**
     * Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
     */
    readonly members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the snatpool
     */
    readonly name: pulumi.Input<string>;
}
