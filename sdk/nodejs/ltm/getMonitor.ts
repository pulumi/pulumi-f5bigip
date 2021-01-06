// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getMonitor(args: GetMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("f5bigip:ltm/getMonitor:getMonitor", {
        "name": args.name,
        "partition": args.partition,
    }, opts);
}

/**
 * A collection of arguments for invoking getMonitor.
 */
export interface GetMonitorArgs {
    readonly name: string;
    readonly partition: string;
}

/**
 * A collection of values returned by getMonitor.
 */
export interface GetMonitorResult {
    readonly adaptive: string;
    readonly adaptiveLimit: number;
    readonly database: string;
    readonly defaultsFrom: string;
    readonly destination: string;
    readonly filename: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly interval: number;
    readonly ipDscp: number;
    readonly manualResume: string;
    readonly mode: string;
    readonly name: string;
    readonly partition: string;
    readonly receiveDisable: string;
    readonly reverse: string;
    readonly timeUntilUp: number;
    readonly timeout: number;
    readonly transparent: string;
    readonly username: string;
}