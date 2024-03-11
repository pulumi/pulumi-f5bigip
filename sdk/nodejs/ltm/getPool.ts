// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const pool-Example = f5bigip.ltm.getPool({
 *     name: "example-pool",
 *     partition: "Common",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPool(args: GetPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetPoolResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:ltm/getPool:getPool", {
        "name": args.name,
        "partition": args.partition,
    }, opts);
}

/**
 * A collection of arguments for invoking getPool.
 */
export interface GetPoolArgs {
    /**
     * Name of the ltm monitor
     */
    name: string;
    /**
     * partition of the ltm monitor
     */
    partition: string;
}

/**
 * A collection of values returned by getPool.
 */
export interface GetPoolResult {
    /**
     * Full path to the pool.
     */
    readonly fullPath: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly partition: string;
}
/**
 * Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const pool-Example = f5bigip.ltm.getPool({
 *     name: "example-pool",
 *     partition: "Common",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPoolOutput(args: GetPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPoolResult> {
    return pulumi.output(args).apply((a: any) => getPool(a, opts))
}

/**
 * A collection of arguments for invoking getPool.
 */
export interface GetPoolOutputArgs {
    /**
     * Name of the ltm monitor
     */
    name: pulumi.Input<string>;
    /**
     * partition of the ltm monitor
     */
    partition: pulumi.Input<string>;
}
