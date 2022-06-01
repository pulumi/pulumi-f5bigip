// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.ssl.getWafPbSuggestions`) to create JSON for WAF URL to later use with an existing WAF policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const wAFURL1 = pulumi.output(f5bigip.ssl.getWafEntityUrl({
 *     description: "this is a test",
 *     methodOverrides: [
 *         {
 *             allow: false,
 *             method: "BCOPY",
 *         },
 *         {
 *             allow: true,
 *             method: "BDELETE",
 *         },
 *     ],
 *     name: "/foobar",
 *     performStaging: true,
 *     protocol: "HTTP",
 *     signatureOverridesDisables: [
 *         12345678,
 *         87654321,
 *     ],
 *     type: "explicit",
 * }));
 * ```
 */
export function getWafEntityUrl(args: GetWafEntityUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetWafEntityUrlResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("f5bigip:ssl/getWafEntityUrl:getWafEntityUrl", {
        "description": args.description,
        "method": args.method,
        "methodOverrides": args.methodOverrides,
        "name": args.name,
        "performStaging": args.performStaging,
        "protocol": args.protocol,
        "signatureOverridesDisables": args.signatureOverridesDisables,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getWafEntityUrl.
 */
export interface GetWafEntityUrlArgs {
    /**
     * A description of the URL.
     */
    description?: string;
    /**
     * Specifies an HTTP method.
     */
    method?: string;
    /**
     * A list of methods that are allowed or disallowed for a specific URL.
     */
    methodOverrides?: inputs.ssl.GetWafEntityUrlMethodOverride[];
    /**
     * WAF entity URL name.
     */
    name: string;
    /**
     * If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
     */
    performStaging?: boolean;
    /**
     * Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
     */
    protocol?: string;
    /**
     * List of Attack Signature Ids which are disabled for this particular URL.
     */
    signatureOverridesDisables?: number[];
    /**
     * Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
     */
    type?: string;
}

/**
 * A collection of values returned by getWafEntityUrl.
 */
export interface GetWafEntityUrlResult {
    readonly description?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Json string representing created WAF entity URL declaration in JSON format
     */
    readonly json: string;
    readonly method?: string;
    readonly methodOverrides?: outputs.ssl.GetWafEntityUrlMethodOverride[];
    readonly name: string;
    readonly performStaging?: boolean;
    readonly protocol?: string;
    readonly signatureOverridesDisables?: number[];
    readonly type?: string;
}

export function getWafEntityUrlOutput(args: GetWafEntityUrlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWafEntityUrlResult> {
    return pulumi.output(args).apply(a => getWafEntityUrl(a, opts))
}

/**
 * A collection of arguments for invoking getWafEntityUrl.
 */
export interface GetWafEntityUrlOutputArgs {
    /**
     * A description of the URL.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies an HTTP method.
     */
    method?: pulumi.Input<string>;
    /**
     * A list of methods that are allowed or disallowed for a specific URL.
     */
    methodOverrides?: pulumi.Input<pulumi.Input<inputs.ssl.GetWafEntityUrlMethodOverrideArgs>[]>;
    /**
     * WAF entity URL name.
     */
    name: pulumi.Input<string>;
    /**
     * If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
     */
    performStaging?: pulumi.Input<boolean>;
    /**
     * Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
     */
    protocol?: pulumi.Input<string>;
    /**
     * List of Attack Signature Ids which are disabled for this particular URL.
     */
    signatureOverridesDisables?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
     */
    type?: pulumi.Input<string>;
}