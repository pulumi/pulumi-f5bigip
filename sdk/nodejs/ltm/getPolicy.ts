// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test = f5bigip.ltm.getPolicy({
 *     name: "/Common/test-policy",
 * });
 * export const bigipPolicy = test.then(test => test.rules);
 * ```
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:ltm/getPolicy:getPolicy", {
        "controls": args.controls,
        "name": args.name,
        "publishedCopy": args.publishedCopy,
        "requires": args.requires,
        "rules": args.rules,
        "strategy": args.strategy,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    /**
     * Specifies the controls.
     */
    controls?: string[];
    /**
     * Name of the policy which includes partion ( /partition/policy-name )
     */
    name: string;
    publishedCopy?: string;
    /**
     * Specifies the protocol.
     */
    requires?: string[];
    /**
     * Rules defined in the policy.
     */
    rules?: inputs.ltm.GetPolicyRule[];
    /**
     * Specifies the match strategy.
     */
    strategy?: string;
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    /**
     * Specifies the controls.
     */
    readonly controls?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the policy.
     */
    readonly name: string;
    readonly publishedCopy?: string;
    /**
     * Specifies the protocol.
     */
    readonly requires?: string[];
    /**
     * Rules defined in the policy.
     */
    readonly rules?: outputs.ltm.GetPolicyRule[];
    /**
     * Specifies the match strategy.
     */
    readonly strategy?: string;
}
/**
 * Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test = f5bigip.ltm.getPolicy({
 *     name: "/Common/test-policy",
 * });
 * export const bigipPolicy = test.then(test => test.rules);
 * ```
 */
export function getPolicyOutput(args: GetPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyResult> {
    return pulumi.output(args).apply((a: any) => getPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyOutputArgs {
    /**
     * Specifies the controls.
     */
    controls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the policy which includes partion ( /partition/policy-name )
     */
    name: pulumi.Input<string>;
    publishedCopy?: pulumi.Input<string>;
    /**
     * Specifies the protocol.
     */
    requires?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Rules defined in the policy.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.ltm.GetPolicyRuleArgs>[]>;
    /**
     * Specifies the match strategy.
     */
    strategy?: pulumi.Input<string>;
}
