// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source (`f5bigip.ssl.getWafSignatures`) to get the details of attack signatures available on BIG-IP WAF
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const WAFSIG1 = f5bigip.ssl.getWafSignatures({
 *     signatureId: 200104004,
 * });
 * ```
 */
export function getWafSignatures(args: GetWafSignaturesArgs, opts?: pulumi.InvokeOptions): Promise<GetWafSignaturesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("f5bigip:ssl/getWafSignatures:getWafSignatures", {
        "accuracy": args.accuracy,
        "description": args.description,
        "enabled": args.enabled,
        "name": args.name,
        "performStaging": args.performStaging,
        "risk": args.risk,
        "signatureId": args.signatureId,
        "systemSignatureId": args.systemSignatureId,
        "tag": args.tag,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getWafSignatures.
 */
export interface GetWafSignaturesArgs {
    /**
     * The relative detection accuracy of the signature.
     */
    accuracy?: string;
    /**
     * Description of the signature.
     */
    description?: string;
    enabled?: boolean;
    /**
     * Name of the signature as configured on the system.
     */
    name?: string;
    performStaging?: boolean;
    /**
     * The relative risk level of the attack that matches this signature.
     */
    risk?: string;
    /**
     * ID of the signature in the BIG-IP WAF database.
     */
    signatureId: number;
    /**
     * System generated ID of the signature.
     */
    systemSignatureId?: string;
    tag?: string;
    /**
     * Type of the signature.
     */
    type?: string;
}

/**
 * A collection of values returned by getWafSignatures.
 */
export interface GetWafSignaturesResult {
    /**
     * The relative detection accuracy of the signature.
     */
    readonly accuracy: string;
    /**
     * Description of the signature.
     */
    readonly description?: string;
    readonly enabled?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly json: string;
    /**
     * Name of the signature as configured on the system.
     */
    readonly name: string;
    readonly performStaging?: boolean;
    /**
     * The relative risk level of the attack that matches this signature.
     */
    readonly risk: string;
    /**
     * ID of the signature in the database.
     */
    readonly signatureId: number;
    /**
     * System generated ID of the signature.
     */
    readonly systemSignatureId: string;
    readonly tag: string;
    /**
     * Type of the signature.
     */
    readonly type: string;
}
/**
 * Use this data source (`f5bigip.ssl.getWafSignatures`) to get the details of attack signatures available on BIG-IP WAF
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const WAFSIG1 = f5bigip.ssl.getWafSignatures({
 *     signatureId: 200104004,
 * });
 * ```
 */
export function getWafSignaturesOutput(args: GetWafSignaturesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWafSignaturesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("f5bigip:ssl/getWafSignatures:getWafSignatures", {
        "accuracy": args.accuracy,
        "description": args.description,
        "enabled": args.enabled,
        "name": args.name,
        "performStaging": args.performStaging,
        "risk": args.risk,
        "signatureId": args.signatureId,
        "systemSignatureId": args.systemSignatureId,
        "tag": args.tag,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getWafSignatures.
 */
export interface GetWafSignaturesOutputArgs {
    /**
     * The relative detection accuracy of the signature.
     */
    accuracy?: pulumi.Input<string>;
    /**
     * Description of the signature.
     */
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the signature as configured on the system.
     */
    name?: pulumi.Input<string>;
    performStaging?: pulumi.Input<boolean>;
    /**
     * The relative risk level of the attack that matches this signature.
     */
    risk?: pulumi.Input<string>;
    /**
     * ID of the signature in the BIG-IP WAF database.
     */
    signatureId: pulumi.Input<number>;
    /**
     * System generated ID of the signature.
     */
    systemSignatureId?: pulumi.Input<string>;
    tag?: pulumi.Input<string>;
    /**
     * Type of the signature.
     */
    type?: pulumi.Input<string>;
}
