// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.WafPolicy`) to get the details of exist WAF policy BIG-IP.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const existpolicy = f5bigip.ssl.getWafPolicy({
 *     policyId: "xxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getWafPolicy(args: GetWafPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetWafPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:ssl/getWafPolicy:getWafPolicy", {
        "policyId": args.policyId,
        "policyJson": args.policyJson,
    }, opts);
}

/**
 * A collection of arguments for invoking getWafPolicy.
 */
export interface GetWafPolicyArgs {
    /**
     * ID of the WAF policy deployed in the BIG-IP.
     */
    policyId: string;
    /**
     * Exported WAF policy JSON
     */
    policyJson?: string;
}

/**
 * A collection of values returned by getWafPolicy.
 */
export interface GetWafPolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly policyId: string;
    /**
     * Exported WAF policy JSON
     */
    readonly policyJson: string;
}
/**
 * Use this data source (`f5bigip.WafPolicy`) to get the details of exist WAF policy BIG-IP.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const existpolicy = f5bigip.ssl.getWafPolicy({
 *     policyId: "xxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getWafPolicyOutput(args: GetWafPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWafPolicyResult> {
    return pulumi.output(args).apply((a: any) => getWafPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getWafPolicy.
 */
export interface GetWafPolicyOutputArgs {
    /**
     * ID of the WAF policy deployed in the BIG-IP.
     */
    policyId: pulumi.Input<string>;
    /**
     * Exported WAF policy JSON
     */
    policyJson?: pulumi.Input<string>;
}
