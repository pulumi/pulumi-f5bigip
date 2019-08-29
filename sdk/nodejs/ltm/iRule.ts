// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.IRule` Creates iRule on BIG-IP F5 device
 * 
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_irule.html.markdown.
 */
export class IRule extends pulumi.CustomResource {
    /**
     * Get an existing IRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IRuleState | undefined;
            inputs["irule"] = state ? state.irule : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as IRuleArgs | undefined;
            if (!args || args.irule === undefined) {
                throw new Error("Missing required property 'irule'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["irule"] = args ? args.irule : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IRule resources.
 */
export interface IRuleState {
    /**
     * Body of the iRule
     */
    readonly irule?: pulumi.Input<string>;
    /**
     * Name of the iRule
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IRule resource.
 */
export interface IRuleArgs {
    /**
     * Body of the iRule
     */
    readonly irule: pulumi.Input<string>;
    /**
     * Name of the iRule
     */
    readonly name: pulumi.Input<string>;
}
