// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getIrule(args: GetIruleArgs, opts?: pulumi.InvokeOptions): Promise<GetIruleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("f5bigip:ltm/getIrule:getIrule", {
        "name": args.name,
        "partition": args.partition,
    }, opts);
}

/**
 * A collection of arguments for invoking getIrule.
 */
export interface GetIruleArgs {
    readonly name: string;
    readonly partition: string;
}

/**
 * A collection of values returned by getIrule.
 */
export interface GetIruleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly irule: string;
    readonly name: string;
    readonly partition: string;
}