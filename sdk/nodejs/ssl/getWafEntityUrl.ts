// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
 * const WAFURL1 = f5bigip.ssl.getWafEntityUrl({
 *     name: "/foobar",
 *     description: "this is a test",
 *     type: "explicit",
 *     protocol: "HTTP",
 *     performStaging: true,
 *     signatureOverridesDisables: [
 *         12345678,
 *         87654321,
 *     ],
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
 *     crossOriginRequestsEnforcements: [
 *         {
 *             includeSubdomains: true,
 *             originName: "app1.com",
 *             originPort: "80",
 *             originProtocol: "http",
 *         },
 *         {
 *             includeSubdomains: true,
 *             originName: "app2.com",
 *             originPort: "443",
 *             originProtocol: "http",
 *         },
 *     ],
 * });
 * ```
 */
export function getWafEntityUrl(args: GetWafEntityUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetWafEntityUrlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:ssl/getWafEntityUrl:getWafEntityUrl", {
        "crossOriginRequestsEnforcements": args.crossOriginRequestsEnforcements,
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
     * A list of options that enables your web-application to share data with a website hosted on a
     * different domain.
     */
    crossOriginRequestsEnforcements?: inputs.ssl.GetWafEntityUrlCrossOriginRequestsEnforcement[];
    /**
     * A description of the URL.
     */
    description?: string;
    /**
     * Select a Method for the URL to create an API endpoint. Default is : *.
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
    readonly crossOriginRequestsEnforcements?: outputs.ssl.GetWafEntityUrlCrossOriginRequestsEnforcement[];
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
/**
 * Use this data source (`f5bigip.ssl.getWafPbSuggestions`) to create JSON for WAF URL to later use with an existing WAF policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const WAFURL1 = f5bigip.ssl.getWafEntityUrl({
 *     name: "/foobar",
 *     description: "this is a test",
 *     type: "explicit",
 *     protocol: "HTTP",
 *     performStaging: true,
 *     signatureOverridesDisables: [
 *         12345678,
 *         87654321,
 *     ],
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
 *     crossOriginRequestsEnforcements: [
 *         {
 *             includeSubdomains: true,
 *             originName: "app1.com",
 *             originPort: "80",
 *             originProtocol: "http",
 *         },
 *         {
 *             includeSubdomains: true,
 *             originName: "app2.com",
 *             originPort: "443",
 *             originProtocol: "http",
 *         },
 *     ],
 * });
 * ```
 */
export function getWafEntityUrlOutput(args: GetWafEntityUrlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWafEntityUrlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("f5bigip:ssl/getWafEntityUrl:getWafEntityUrl", {
        "crossOriginRequestsEnforcements": args.crossOriginRequestsEnforcements,
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
export interface GetWafEntityUrlOutputArgs {
    /**
     * A list of options that enables your web-application to share data with a website hosted on a
     * different domain.
     */
    crossOriginRequestsEnforcements?: pulumi.Input<pulumi.Input<inputs.ssl.GetWafEntityUrlCrossOriginRequestsEnforcementArgs>[]>;
    /**
     * A description of the URL.
     */
    description?: pulumi.Input<string>;
    /**
     * Select a Method for the URL to create an API endpoint. Default is : *.
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
