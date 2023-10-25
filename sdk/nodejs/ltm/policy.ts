// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.Policy` Configures ltm policies to manage traffic assigned to a virtual server
 *
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-policy`.
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
 *             connection: false,
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
     * Specifies descriptive text that identifies the irule attached to policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of Rule to be applied in policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
     *
     * @deprecated This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
     */
    public readonly publishedCopy!: pulumi.Output<string | undefined>;
    /**
     * Specifies the protocol
     */
    public readonly requires!: pulumi.Output<string[] | undefined>;
    /**
     * List of Rules can be applied using the policy. Each rule is block type with following arguments.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            resourceInputs["controls"] = state ? state.controls : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publishedCopy"] = state ? state.publishedCopy : undefined;
            resourceInputs["requires"] = state ? state.requires : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["strategy"] = state ? state.strategy : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["controls"] = args ? args.controls : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publishedCopy"] = args ? args.publishedCopy : undefined;
            resourceInputs["requires"] = args ? args.requires : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["strategy"] = args ? args.strategy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Specifies the controls
     */
    controls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies descriptive text that identifies the irule attached to policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of Rule to be applied in policy.
     */
    name?: pulumi.Input<string>;
    /**
     * If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
     *
     * @deprecated This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
     */
    publishedCopy?: pulumi.Input<string>;
    /**
     * Specifies the protocol
     */
    requires?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Rules can be applied using the policy. Each rule is block type with following arguments.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.ltm.PolicyRule>[]>;
    /**
     * Specifies the match strategy
     */
    strategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Specifies the controls
     */
    controls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies descriptive text that identifies the irule attached to policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of Rule to be applied in policy.
     */
    name: pulumi.Input<string>;
    /**
     * If you want to publish the policy else it will be deployed in Drafts mode. This attribute is deprecated and will be removed in a future release.
     *
     * @deprecated This attribute is not required anymore because the resource automatically publishes the policy, for that reason this field is deprecated and will be removed in a future release.
     */
    publishedCopy?: pulumi.Input<string>;
    /**
     * Specifies the protocol
     */
    requires?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Rules can be applied using the policy. Each rule is block type with following arguments.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.ltm.PolicyRule>[]>;
    /**
     * Specifies the match strategy
     */
    strategy?: pulumi.Input<string>;
}
