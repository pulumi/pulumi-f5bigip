// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.Policy` Configures ltm policies to manage traffic assigned to a virtual server
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const mypool = new f5bigip.ltm.Pool("mypool", {
 *     name: "/Common/test-pool",
 *     allowNat: "yes",
 *     allowSnat: "yes",
 *     loadBalancingMode: "round-robin",
 * });
 * const test_policy = new f5bigip.ltm.Policy("test-policy", {
 *     name: "/Common/test-policy",
 *     strategy: "first-match",
 *     requires: ["http"],
 *     controls: ["forwarding"],
 *     rules: [{
 *         name: "rule6",
 *         actions: [{
 *             forward: true,
 *             pool: mypool.name,
 *         }],
 *     }],
 * }, {
 *     dependsOn: [mypool],
 * });
 * ```
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Specifies the controls
     */
    public readonly controls!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If you want to publish the policy else it will be deployed in Drafts mode.
     */
    public readonly publishedCopy!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol
     */
    public readonly requires!: pulumi.Output<string[] | undefined>;
    /**
     * Rules can be applied using the policy
     */
    public readonly rules!: pulumi.Output<outputs.ltm.PolicyRule[] | undefined>;
    /**
     * Specifies the match strategy
     */
    public readonly strategy!: pulumi.Output<string | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["controls"] = state ? state.controls : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publishedCopy"] = state ? state.publishedCopy : undefined;
            inputs["requires"] = state ? state.requires : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["strategy"] = state ? state.strategy : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["controls"] = args ? args.controls : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publishedCopy"] = args ? args.publishedCopy : undefined;
            inputs["requires"] = args ? args.requires : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["strategy"] = args ? args.strategy : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Specifies the controls
     */
    readonly controls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
     */
    readonly name?: pulumi.Input<string>;
    /**
     * If you want to publish the policy else it will be deployed in Drafts mode.
     */
    readonly publishedCopy?: pulumi.Input<string>;
    /**
     * Specifies the protocol
     */
    readonly requires?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Rules can be applied using the policy
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.ltm.PolicyRule>[]>;
    /**
     * Specifies the match strategy
     */
    readonly strategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Specifies the controls
     */
    readonly controls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
     */
    readonly name: pulumi.Input<string>;
    /**
     * If you want to publish the policy else it will be deployed in Drafts mode.
     */
    readonly publishedCopy?: pulumi.Input<string>;
    /**
     * Specifies the protocol
     */
    readonly requires?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Rules can be applied using the policy
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.ltm.PolicyRule>[]>;
    /**
     * Specifies the match strategy
     */
    readonly strategy?: pulumi.Input<string>;
}
