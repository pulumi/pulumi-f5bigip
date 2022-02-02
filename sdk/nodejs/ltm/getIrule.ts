// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test = pulumi.output(f5bigip.ltm.getIrule({
 *     name: "terraform_irule",
 *     partition: "Common",
 * }));
 *
 * export const bigipIrule = test.irule;
 * ```
 */
export function getIrule(args: GetIruleArgs, opts?: pulumi.InvokeOptions): Promise<GetIruleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("f5bigip:ltm/getIrule:getIrule", {
        "name": args.name,
        "partition": args.partition,
    }, opts);
}

/**
 * A collection of arguments for invoking getIrule.
 */
export interface GetIruleArgs {
    /**
     * Name of the irule
     */
    name: string;
    /**
     * partition of the ltm irule
     */
    partition: string;
}

/**
 * A collection of values returned by getIrule.
 */
export interface GetIruleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Irule configured on bigip
     */
    readonly irule: string;
    /**
     * Name of irule configured on bigip with full path
     */
    readonly name: string;
    /**
     * Bigip partition in which rule is configured
     */
    readonly partition: string;
}

export function getIruleOutput(args: GetIruleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIruleResult> {
    return pulumi.output(args).apply(a => getIrule(a, opts))
}

/**
 * A collection of arguments for invoking getIrule.
 */
export interface GetIruleOutputArgs {
    /**
     * Name of the irule
     */
    name: pulumi.Input<string>;
    /**
     * partition of the ltm irule
     */
    partition: pulumi.Input<string>;
}
