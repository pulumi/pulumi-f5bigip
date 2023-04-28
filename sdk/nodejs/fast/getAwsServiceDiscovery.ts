// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const tC2 = f5bigip.fast.getAwsServiceDiscovery({
 *     tagKey: "testawstagkey",
 *     tagValue: "testawstagvalue",
 * });
 * ```
 */
export function getAwsServiceDiscovery(args: GetAwsServiceDiscoveryArgs, opts?: pulumi.InvokeOptions): Promise<GetAwsServiceDiscoveryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:fast/getAwsServiceDiscovery:getAwsServiceDiscovery", {
        "addressRealm": args.addressRealm,
        "awsAccessKey": args.awsAccessKey,
        "awsRegion": args.awsRegion,
        "awsSecretAccessKey": args.awsSecretAccessKey,
        "credentialUpdate": args.credentialUpdate,
        "externalId": args.externalId,
        "minimumMonitors": args.minimumMonitors,
        "port": args.port,
        "roleArn": args.roleArn,
        "tagKey": args.tagKey,
        "tagValue": args.tagValue,
        "type": args.type,
        "undetectableAction": args.undetectableAction,
        "updateInterval": args.updateInterval,
    }, opts);
}

/**
 * A collection of arguments for invoking getAwsServiceDiscovery.
 */
export interface GetAwsServiceDiscoveryArgs {
    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     */
    addressRealm?: string;
    /**
     * Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
     */
    awsAccessKey?: string;
    /**
     * AWS region in which ADC is running,default Empty string.
     */
    awsRegion?: string;
    /**
     * Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
     */
    awsSecretAccessKey?: string;
    /**
     * Specifies whether you are updating your credentials,default `false`.
     */
    credentialUpdate?: boolean;
    /**
     * AWS externalID field.
     */
    externalId?: string;
    /**
     * Member is down when fewer than minimum monitors report it healthy.
     */
    minimumMonitors?: string;
    /**
     * Port to be used for AWS service discovery,default `80`.
     */
    port?: number;
    /**
     * Assume a role (also requires the `externalId` field)
     */
    roleArn?: string;
    /**
     * The tag key associated with the node to add to this pool.
     */
    tagKey: string;
    /**
     * The tag value associated with the node to add to this pool.
     */
    tagValue: string;
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
 * A collection of values returned by getAwsServiceDiscovery.
 */
export interface GetAwsServiceDiscoveryResult {
    readonly addressRealm?: string;
    readonly awsAccessKey?: string;
    readonly awsRegion: string;
    /**
     * The JSON for AWS service discovery block.
     */
    readonly awsSdJson: string;
    readonly awsSecretAccessKey?: string;
    readonly credentialUpdate?: boolean;
    readonly externalId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly minimumMonitors?: string;
    readonly port?: number;
    readonly roleArn: string;
    readonly tagKey: string;
    readonly tagValue: string;
    readonly type?: string;
    readonly undetectableAction?: string;
    readonly updateInterval?: string;
}
/**
 * Use this data source (`f5bigip.fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const tC2 = f5bigip.fast.getAwsServiceDiscovery({
 *     tagKey: "testawstagkey",
 *     tagValue: "testawstagvalue",
 * });
 * ```
 */
export function getAwsServiceDiscoveryOutput(args: GetAwsServiceDiscoveryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAwsServiceDiscoveryResult> {
    return pulumi.output(args).apply((a: any) => getAwsServiceDiscovery(a, opts))
}

/**
 * A collection of arguments for invoking getAwsServiceDiscovery.
 */
export interface GetAwsServiceDiscoveryOutputArgs {
    /**
     * Specifies whether to look for public or private IP addresses,default `private`.
     */
    addressRealm?: pulumi.Input<string>;
    /**
     * Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
     */
    awsAccessKey?: pulumi.Input<string>;
    /**
     * AWS region in which ADC is running,default Empty string.
     */
    awsRegion?: pulumi.Input<string>;
    /**
     * Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
     */
    awsSecretAccessKey?: pulumi.Input<string>;
    /**
     * Specifies whether you are updating your credentials,default `false`.
     */
    credentialUpdate?: pulumi.Input<boolean>;
    /**
     * AWS externalID field.
     */
    externalId?: pulumi.Input<string>;
    /**
     * Member is down when fewer than minimum monitors report it healthy.
     */
    minimumMonitors?: pulumi.Input<string>;
    /**
     * Port to be used for AWS service discovery,default `80`.
     */
    port?: pulumi.Input<number>;
    /**
     * Assume a role (also requires the `externalId` field)
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The tag key associated with the node to add to this pool.
     */
    tagKey: pulumi.Input<string>;
    /**
     * The tag value associated with the node to add to this pool.
     */
    tagValue: pulumi.Input<string>;
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
