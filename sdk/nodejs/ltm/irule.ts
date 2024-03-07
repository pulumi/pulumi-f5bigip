// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.IRule` Creates iRule on BIG-IP F5 device
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * as fs from "fs";
 *
 * // Loading from a file is the preferred method
 * const rule = new f5bigip.ltm.IRule("rule", {
 *     name: "/Common/terraform_irule",
 *     irule: fs.readFileSync("myirule.tcl", "utf8"),
 * });
 * const rule2 = new f5bigip.ltm.IRule("rule2", {
 *     name: "/Common/terraform_irule2",
 *     irule: `when CLIENT_ACCEPTED {
 *      log local0. "test"
 *    }
 * `,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class IRule extends pulumi.CustomResource {
    /**
     * Get an existing IRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IRuleState, opts?: pulumi.CustomResourceOptions): IRule {
        return new IRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/iRule:IRule';

    /**
     * Returns true if the given object is an instance of IRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IRule.__pulumiType;
    }

    /**
     * Body of the iRule
     */
    public readonly irule!: pulumi.Output<string>;
    /**
     * Name of the iRule
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a IRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IRuleArgs | IRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IRuleState | undefined;
            resourceInputs["irule"] = state ? state.irule : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as IRuleArgs | undefined;
            if ((!args || args.irule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'irule'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["irule"] = args ? args.irule : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IRule resources.
 */
export interface IRuleState {
    /**
     * Body of the iRule
     */
    irule?: pulumi.Input<string>;
    /**
     * Name of the iRule
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IRule resource.
 */
export interface IRuleArgs {
    /**
     * Body of the iRule
     */
    irule: pulumi.Input<string>;
    /**
     * Name of the iRule
     */
    name: pulumi.Input<string>;
}
