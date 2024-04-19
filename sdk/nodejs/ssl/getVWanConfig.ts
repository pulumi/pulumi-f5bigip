// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.ssl.getVWanConfig`) to get the vWAN site config from Azure VWAN Site
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const vwanconfig = f5bigip.ssl.getVWanConfig({
 *     azureVwanName: "azurevwan-bigip-vwan-9c8d",
 *     azureVwanResourcegroup: "azurevwan-bigip-rg-9c8d",
 *     azureVwanVpnsite: "azurevwan-bigip-vsite-9c8d",
 * });
 * ```
 *
 * ## Pre-required Environment Settings:
 *
 * * `AZURE_CLIENT_ID` - (Required) Set this environment variable with the Azure app client ID to use.
 *
 * * `AZURE_CLIENT_SECRET` - (Required) Set this environment variable with the Azure app secret to use.
 *
 * * `AZURE_SUBSCRIPTION_ID` - (Required) Set this environment variable with the Azure subscription ID to use.
 *
 * * `AZURE_TENANT_ID` - (Required) Set this environment variable with the Tenant ID to which to authenticate.
 *
 * * `STORAGE_ACCOUNT_NAME` - (Required) Set this environment variable with the storage account for download config.
 *
 * * `STORAGE_ACCOUNT_KEY` - (Required) Specifies the storage account key to authenticate,set this Environment variable with account key value.
 */
export function getVWanConfig(args: GetVWanConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetVWanConfigResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:ssl/getVWanConfig:getVWanConfig", {
        "azureVwanName": args.azureVwanName,
        "azureVwanResourcegroup": args.azureVwanResourcegroup,
        "azureVwanVpnsite": args.azureVwanVpnsite,
    }, opts);
}

/**
 * A collection of arguments for invoking getVWanConfig.
 */
export interface GetVWanConfigArgs {
    /**
     * Name of the Azure vWAN Name
     */
    azureVwanName: string;
    /**
     * Name of the Azure vWAN resource group
     */
    azureVwanResourcegroup: string;
    /**
     * Name of the Azure vWAN VPN site from which configuration to be download
     */
    azureVwanVpnsite: string;
}

/**
 * A collection of values returned by getVWanConfig.
 */
export interface GetVWanConfigResult {
    readonly azureVwanName: string;
    readonly azureVwanResourcegroup: string;
    readonly azureVwanVpnsite: string;
    /**
     * (type `string`) provides IP address of BIGIP G/W for IPSec Endpoint.
     */
    readonly bigipGwIp: string;
    /**
     * (type `string`) Provides IP Address space used on vWAN Hub.
     */
    readonly hubAddressSpace: string;
    /**
     * (type `list`) Provides Subnets connected to vWAN Hub.
     */
    readonly hubConnectedSubnets: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (type `string`) provides pre-shared-key used for IPSec Tunnel creation.
     */
    readonly presharedKey: string;
    /**
     * (type `list`) Provides vWAN Gateway Address for IPSec End point
     */
    readonly vwanGwAddresses: string[];
}
/**
 * Use this data source (`f5bigip.ssl.getVWanConfig`) to get the vWAN site config from Azure VWAN Site
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const vwanconfig = f5bigip.ssl.getVWanConfig({
 *     azureVwanName: "azurevwan-bigip-vwan-9c8d",
 *     azureVwanResourcegroup: "azurevwan-bigip-rg-9c8d",
 *     azureVwanVpnsite: "azurevwan-bigip-vsite-9c8d",
 * });
 * ```
 *
 * ## Pre-required Environment Settings:
 *
 * * `AZURE_CLIENT_ID` - (Required) Set this environment variable with the Azure app client ID to use.
 *
 * * `AZURE_CLIENT_SECRET` - (Required) Set this environment variable with the Azure app secret to use.
 *
 * * `AZURE_SUBSCRIPTION_ID` - (Required) Set this environment variable with the Azure subscription ID to use.
 *
 * * `AZURE_TENANT_ID` - (Required) Set this environment variable with the Tenant ID to which to authenticate.
 *
 * * `STORAGE_ACCOUNT_NAME` - (Required) Set this environment variable with the storage account for download config.
 *
 * * `STORAGE_ACCOUNT_KEY` - (Required) Specifies the storage account key to authenticate,set this Environment variable with account key value.
 */
export function getVWanConfigOutput(args: GetVWanConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVWanConfigResult> {
    return pulumi.output(args).apply((a: any) => getVWanConfig(a, opts))
}

/**
 * A collection of arguments for invoking getVWanConfig.
 */
export interface GetVWanConfigOutputArgs {
    /**
     * Name of the Azure vWAN Name
     */
    azureVwanName: pulumi.Input<string>;
    /**
     * Name of the Azure vWAN resource group
     */
    azureVwanResourcegroup: pulumi.Input<string>;
    /**
     * Name of the Azure vWAN VPN site from which configuration to be download
     */
    azureVwanVpnsite: pulumi.Input<string>;
}
