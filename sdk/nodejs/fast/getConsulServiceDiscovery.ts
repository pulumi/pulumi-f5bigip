// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const TC2 = f5bigip.fast.getConsulServiceDiscovery({
 *     uri: "https://192.0.2.100:8500/v1/catalog/nodes",
 *     port: 8080,
 * });
 * ```
 */
export function getConsulServiceDiscovery(args: GetConsulServiceDiscoveryArgs, opts?: pulumi.InvokeOptions): Promise<GetConsulServiceDiscoveryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", {
        "addressRealm": args.addressRealm,
        "credentialUpdate": args.credentialUpdate,
        "encodedToken": args.encodedToken,
        "jmesPathQuery": args.jmesPathQuery,
        "minimumMonitors": args.minimumMonitors,
        "port": args.port,
        "rejectUnauthorized": args.rejectUnauthorized,
        "trustCa": args.trustCa,
        "type": args.type,
        "undetectableAction": args.undetectableAction,
        "updateInterval": args.updateInterval,
        "uri": args.uri,
    }, opts);
}

/**
 * A collection of arguments for invoking getConsulServiceDiscovery.
 */
export interface GetConsulServiceDiscoveryArgs {
    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     */
    addressRealm?: string;
    /**
     * Specifies whether you are updating your credentials,default `false`.
     */
    credentialUpdate?: boolean;
    /**
     * Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
     */
    encodedToken?: string;
    /**
     * Custom JMESPath Query.
     */
    jmesPathQuery?: string;
    /**
     * Member is down when fewer than minimum monitors report it healthy.
     */
    minimumMonitors?: string;
    /**
     * Port to be used for AWS service discovery,default `80`.
     */
    port: number;
    /**
     * If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
     */
    rejectUnauthorized?: boolean;
    /**
     * CA Bundle to validate server certificates.
     */
    trustCa?: string;
    type?: string;
    /**
     * Action to take when node cannot be detected,default `remove`.
     */
    undetectableAction?: string;
    /**
     * Update interval for service discovery.
     */
    updateInterval?: string;
    /**
     * The location of the node data.
     */
    uri: string;
}

/**
 * A collection of values returned by getConsulServiceDiscovery.
 */
export interface GetConsulServiceDiscoveryResult {
    readonly addressRealm?: string;
    readonly consulSdJson: string;
    readonly credentialUpdate?: boolean;
    readonly encodedToken?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly jmesPathQuery?: string;
    readonly minimumMonitors?: string;
    readonly port: number;
    readonly rejectUnauthorized?: boolean;
    readonly trustCa?: string;
    readonly type?: string;
    readonly undetectableAction?: string;
    readonly updateInterval?: string;
    readonly uri: string;
}
/**
 * Use this data source (`f5bigip.fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const TC2 = f5bigip.fast.getConsulServiceDiscovery({
 *     uri: "https://192.0.2.100:8500/v1/catalog/nodes",
 *     port: 8080,
 * });
 * ```
 */
export function getConsulServiceDiscoveryOutput(args: GetConsulServiceDiscoveryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConsulServiceDiscoveryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", {
        "addressRealm": args.addressRealm,
        "credentialUpdate": args.credentialUpdate,
        "encodedToken": args.encodedToken,
        "jmesPathQuery": args.jmesPathQuery,
        "minimumMonitors": args.minimumMonitors,
        "port": args.port,
        "rejectUnauthorized": args.rejectUnauthorized,
        "trustCa": args.trustCa,
        "type": args.type,
        "undetectableAction": args.undetectableAction,
        "updateInterval": args.updateInterval,
        "uri": args.uri,
    }, opts);
}

/**
 * A collection of arguments for invoking getConsulServiceDiscovery.
 */
export interface GetConsulServiceDiscoveryOutputArgs {
    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     */
    addressRealm?: pulumi.Input<string>;
    /**
     * Specifies whether you are updating your credentials,default `false`.
     */
    credentialUpdate?: pulumi.Input<boolean>;
    /**
     * Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
     */
    encodedToken?: pulumi.Input<string>;
    /**
     * Custom JMESPath Query.
     */
    jmesPathQuery?: pulumi.Input<string>;
    /**
     * Member is down when fewer than minimum monitors report it healthy.
     */
    minimumMonitors?: pulumi.Input<string>;
    /**
     * Port to be used for AWS service discovery,default `80`.
     */
    port: pulumi.Input<number>;
    /**
     * If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
     */
    rejectUnauthorized?: pulumi.Input<boolean>;
    /**
     * CA Bundle to validate server certificates.
     */
    trustCa?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    /**
     * Action to take when node cannot be detected,default `remove`.
     */
    undetectableAction?: pulumi.Input<string>;
    /**
     * Update interval for service discovery.
     */
    updateInterval?: pulumi.Input<string>;
    /**
     * The location of the node data.
     */
    uri: pulumi.Input<string>;
}
