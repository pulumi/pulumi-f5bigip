// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const TC3 = f5bigip.fast.getAzureServiceDiscovery({
 *     resourceGroup: "testazurerg",
 *     subscriptionId: "testazuresid",
 *     tagKey: "testazuretag",
 *     tagValue: "testazurevalue",
 * });
 * ```
 */
export function getAzureServiceDiscovery(args: GetAzureServiceDiscoveryArgs, opts?: pulumi.InvokeOptions): Promise<GetAzureServiceDiscoveryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery", {
        "addressRealm": args.addressRealm,
        "credentialUpdate": args.credentialUpdate,
        "minimumMonitors": args.minimumMonitors,
        "port": args.port,
        "resourceGroup": args.resourceGroup,
        "subscriptionId": args.subscriptionId,
        "tagKey": args.tagKey,
        "tagValue": args.tagValue,
        "type": args.type,
        "undetectableAction": args.undetectableAction,
        "updateInterval": args.updateInterval,
    }, opts);
}

/**
 * A collection of arguments for invoking getAzureServiceDiscovery.
 */
export interface GetAzureServiceDiscoveryArgs {
    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     */
    addressRealm?: string;
    /**
     * Specifies whether you are updating your credentials,default `false`.
     */
    credentialUpdate?: boolean;
    /**
     * Member is down when fewer than minimum monitors report it healthy.
     */
    minimumMonitors?: string;
    /**
     * Port to be used for Azure service discovery,default `80`.
     */
    port?: number;
    /**
     * Azure Resource Group name.
     */
    resourceGroup: string;
    /**
     * Azure subscription ID.
     */
    subscriptionId: string;
    /**
     * The tag key associated with the node to add to this pool.
     */
    tagKey?: string;
    /**
     * The tag value associated with the node to add to this pool.
     */
    tagValue?: string;
    type?: string;
    /**
     * Action to take when node cannot be detected,default `remove`.
     */
    undetectableAction?: string;
    /**
     * Update interval for service discovery.
     */
    updateInterval?: string;
}

/**
 * A collection of values returned by getAzureServiceDiscovery.
 */
export interface GetAzureServiceDiscoveryResult {
    readonly addressRealm?: string;
    /**
     * The JSON for Azure service discovery block.
     */
    readonly azureSdJson: string;
    readonly credentialUpdate?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly minimumMonitors?: string;
    readonly port?: number;
    readonly resourceGroup: string;
    readonly subscriptionId: string;
    readonly tagKey?: string;
    readonly tagValue?: string;
    readonly type?: string;
    readonly undetectableAction?: string;
    readonly updateInterval?: string;
}
/**
 * Use this data source (`f5bigip.fast.getAzureServiceDiscovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const TC3 = f5bigip.fast.getAzureServiceDiscovery({
 *     resourceGroup: "testazurerg",
 *     subscriptionId: "testazuresid",
 *     tagKey: "testazuretag",
 *     tagValue: "testazurevalue",
 * });
 * ```
 */
export function getAzureServiceDiscoveryOutput(args: GetAzureServiceDiscoveryOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAzureServiceDiscoveryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery", {
        "addressRealm": args.addressRealm,
        "credentialUpdate": args.credentialUpdate,
        "minimumMonitors": args.minimumMonitors,
        "port": args.port,
        "resourceGroup": args.resourceGroup,
        "subscriptionId": args.subscriptionId,
        "tagKey": args.tagKey,
        "tagValue": args.tagValue,
        "type": args.type,
        "undetectableAction": args.undetectableAction,
        "updateInterval": args.updateInterval,
    }, opts);
}

/**
 * A collection of arguments for invoking getAzureServiceDiscovery.
 */
export interface GetAzureServiceDiscoveryOutputArgs {
    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     */
    addressRealm?: pulumi.Input<string>;
    /**
     * Specifies whether you are updating your credentials,default `false`.
     */
    credentialUpdate?: pulumi.Input<boolean>;
    /**
     * Member is down when fewer than minimum monitors report it healthy.
     */
    minimumMonitors?: pulumi.Input<string>;
    /**
     * Port to be used for Azure service discovery,default `80`.
     */
    port?: pulumi.Input<number>;
    /**
     * Azure Resource Group name.
     */
    resourceGroup: pulumi.Input<string>;
    /**
     * Azure subscription ID.
     */
    subscriptionId: pulumi.Input<string>;
    /**
     * The tag key associated with the node to add to this pool.
     */
    tagKey?: pulumi.Input<string>;
    /**
     * The tag value associated with the node to add to this pool.
     */
    tagValue?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    /**
     * Action to take when node cannot be detected,default `remove`.
     */
    undetectableAction?: pulumi.Input<string>;
    /**
     * Update interval for service discovery.
     */
    updateInterval?: pulumi.Input<string>;
}
