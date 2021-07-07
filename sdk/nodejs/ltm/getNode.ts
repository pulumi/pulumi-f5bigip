// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getNode(args: GetNodeArgs, opts?: pulumi.InvokeOptions): Promise<GetNodeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("f5bigip:ltm/getNode:getNode", {
        "address": args.address,
        "description": args.description,
        "fqdn": args.fqdn,
        "name": args.name,
        "partition": args.partition,
    }, opts);
}

/**
 * A collection of arguments for invoking getNode.
 */
export interface GetNodeArgs {
    address?: string;
    description?: string;
    fqdn?: inputs.ltm.GetNodeFqdn;
    name: string;
    partition: string;
}

/**
 * A collection of values returned by getNode.
 */
export interface GetNodeResult {
    readonly address?: string;
    readonly connectionLimit: number;
    readonly description?: string;
    readonly dynamicRatio: number;
    readonly fqdn: outputs.ltm.GetNodeFqdn;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly monitor: string;
    readonly name: string;
    readonly partition: string;
    readonly rateLimit: string;
    readonly ratio: number;
    readonly session: string;
    readonly state: string;
}
