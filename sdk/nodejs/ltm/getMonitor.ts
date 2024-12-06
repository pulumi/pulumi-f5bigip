// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const monitor-TC1 = f5bigip.ltm.getMonitor({
 *     name: "test-monitor",
 *     partition: "Common",
 * });
 * ```
 */
export function getMonitor(args: GetMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:ltm/getMonitor:getMonitor", {
        "name": args.name,
        "partition": args.partition,
    }, opts);
}

/**
 * A collection of arguments for invoking getMonitor.
 */
export interface GetMonitorArgs {
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
 * A collection of values returned by getMonitor.
 */
export interface GetMonitorResult {
    /**
     * Displays whether adaptive response time monitoring is enabled for this monitor.
     */
    readonly adaptive: string;
    /**
     * Displays whether adaptive response time monitoring is enabled for this monitor.
     */
    readonly adaptiveLimit: number;
    readonly base: string;
    readonly chaseReferrals: string;
    readonly database: string;
    readonly defaultsFrom: string;
    /**
     * id will be full path name of ltm monitor.
     */
    readonly destination: string;
    readonly filename: string;
    readonly filter: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown.
     */
    readonly interval: number;
    /**
     * Displays the differentiated services code point (DSCP). DSCP is a 6-bit value in the Differentiated Services (DS) field of the IP header.
     */
    readonly ipDscp: number;
    readonly mandatoryAttributes: string;
    /**
     * Displays whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
     */
    readonly manualResume: string;
    readonly mode: string;
    readonly name: string;
    readonly partition: string;
    readonly receiveDisable: string;
    /**
     * Instructs the system to mark the target resource down when the test is successful.
     */
    readonly reverse: string;
    readonly security: string;
    readonly timeUntilUp: number;
    readonly timeout: number;
    /**
     * Displays whether the monitor operates in transparent mode.
     */
    readonly transparent: string;
    readonly username: string;
}
/**
 * Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const monitor-TC1 = f5bigip.ltm.getMonitor({
 *     name: "test-monitor",
 *     partition: "Common",
 * });
 * ```
 */
export function getMonitorOutput(args: GetMonitorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMonitorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("f5bigip:ltm/getMonitor:getMonitor", {
        "name": args.name,
        "partition": args.partition,
    }, opts);
}

/**
 * A collection of arguments for invoking getMonitor.
 */
export interface GetMonitorOutputArgs {
    /**
     * Name of the ltm monitor
     */
    name: pulumi.Input<string>;
    /**
     * partition of the ltm monitor
     */
    partition: pulumi.Input<string>;
}
