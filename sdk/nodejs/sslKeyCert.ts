// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.SslKeyCert` This resource will import SSL certificate and key on BIG-IP LTM.
 * The certificate and the key can be imported from files on the local disk, in PEM format
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * as std from "@pulumi/std";
 *
 * const testkeycert = new f5bigip.SslKeyCert("testkeycert", {
 *     partition: "Common",
 *     keyName: "ssl-test-key",
 *     keyContent: std.file({
 *         input: "key.pem",
 *     }).then(invoke => invoke.result),
 *     certName: "ssl-test-cert",
 *     certContent: std.file({
 *         input: "certificate.pem",
 *     }).then(invoke => invoke.result),
 * });
 * ```
 */
export class SslKeyCert extends pulumi.CustomResource {
    /**
     * Get an existing SslKeyCert resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SslKeyCertState, opts?: pulumi.CustomResourceOptions): SslKeyCert {
        return new SslKeyCert(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/sslKeyCert:SslKeyCert';

    /**
     * Returns true if the given object is an instance of SslKeyCert.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SslKeyCert {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SslKeyCert.__pulumiType;
    }

    /**
     * The content of the cert.
     */
    public readonly certContent!: pulumi.Output<string>;
    /**
     * full path of the SSL certificate on the BIGIP.
     */
    public readonly certFullPath!: pulumi.Output<string>;
    /**
     * Specifies the type of monitoring used.
     */
    public readonly certMonitoringType!: pulumi.Output<string | undefined>;
    /**
     * Name of the SSL certificate to be Imported on to BIGIP.
     */
    public readonly certName!: pulumi.Output<string>;
    /**
     * Specifies the OCSP responder.
     */
    public readonly certOcsp!: pulumi.Output<string | undefined>;
    /**
     * Specifies the issuer certificate.
     */
    public readonly issuerCert!: pulumi.Output<string | undefined>;
    /**
     * The content of the key.
     */
    public readonly keyContent!: pulumi.Output<string>;
    /**
     * full path of the SSL key on the BIGIP.
     */
    public readonly keyFullPath!: pulumi.Output<string>;
    /**
     * Name of the SSL key to be Imported on to BIGIP.
     */
    public readonly keyName!: pulumi.Output<string>;
    /**
     * Partition on to SSL certificate and key to be imported.
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * Passphrase on the SSL key.
     */
    public readonly passphrase!: pulumi.Output<string | undefined>;

    /**
     * Create a SslKeyCert resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslKeyCertArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SslKeyCertArgs | SslKeyCertState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SslKeyCertState | undefined;
            resourceInputs["certContent"] = state ? state.certContent : undefined;
            resourceInputs["certFullPath"] = state ? state.certFullPath : undefined;
            resourceInputs["certMonitoringType"] = state ? state.certMonitoringType : undefined;
            resourceInputs["certName"] = state ? state.certName : undefined;
            resourceInputs["certOcsp"] = state ? state.certOcsp : undefined;
            resourceInputs["issuerCert"] = state ? state.issuerCert : undefined;
            resourceInputs["keyContent"] = state ? state.keyContent : undefined;
            resourceInputs["keyFullPath"] = state ? state.keyFullPath : undefined;
            resourceInputs["keyName"] = state ? state.keyName : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["passphrase"] = state ? state.passphrase : undefined;
        } else {
            const args = argsOrState as SslKeyCertArgs | undefined;
            if ((!args || args.certContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certContent'");
            }
            if ((!args || args.certName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certName'");
            }
            if ((!args || args.keyContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyContent'");
            }
            if ((!args || args.keyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyName'");
            }
            resourceInputs["certContent"] = args?.certContent ? pulumi.secret(args.certContent) : undefined;
            resourceInputs["certFullPath"] = args ? args.certFullPath : undefined;
            resourceInputs["certMonitoringType"] = args ? args.certMonitoringType : undefined;
            resourceInputs["certName"] = args ? args.certName : undefined;
            resourceInputs["certOcsp"] = args ? args.certOcsp : undefined;
            resourceInputs["issuerCert"] = args ? args.issuerCert : undefined;
            resourceInputs["keyContent"] = args?.keyContent ? pulumi.secret(args.keyContent) : undefined;
            resourceInputs["keyFullPath"] = args ? args.keyFullPath : undefined;
            resourceInputs["keyName"] = args ? args.keyName : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["passphrase"] = args?.passphrase ? pulumi.secret(args.passphrase) : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["certContent", "keyContent", "passphrase"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SslKeyCert.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SslKeyCert resources.
 */
export interface SslKeyCertState {
    /**
     * The content of the cert.
     */
    certContent?: pulumi.Input<string>;
    /**
     * full path of the SSL certificate on the BIGIP.
     */
    certFullPath?: pulumi.Input<string>;
    /**
     * Specifies the type of monitoring used.
     */
    certMonitoringType?: pulumi.Input<string>;
    /**
     * Name of the SSL certificate to be Imported on to BIGIP.
     */
    certName?: pulumi.Input<string>;
    /**
     * Specifies the OCSP responder.
     */
    certOcsp?: pulumi.Input<string>;
    /**
     * Specifies the issuer certificate.
     */
    issuerCert?: pulumi.Input<string>;
    /**
     * The content of the key.
     */
    keyContent?: pulumi.Input<string>;
    /**
     * full path of the SSL key on the BIGIP.
     */
    keyFullPath?: pulumi.Input<string>;
    /**
     * Name of the SSL key to be Imported on to BIGIP.
     */
    keyName?: pulumi.Input<string>;
    /**
     * Partition on to SSL certificate and key to be imported.
     */
    partition?: pulumi.Input<string>;
    /**
     * Passphrase on the SSL key.
     */
    passphrase?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SslKeyCert resource.
 */
export interface SslKeyCertArgs {
    /**
     * The content of the cert.
     */
    certContent: pulumi.Input<string>;
    /**
     * full path of the SSL certificate on the BIGIP.
     */
    certFullPath?: pulumi.Input<string>;
    /**
     * Specifies the type of monitoring used.
     */
    certMonitoringType?: pulumi.Input<string>;
    /**
     * Name of the SSL certificate to be Imported on to BIGIP.
     */
    certName: pulumi.Input<string>;
    /**
     * Specifies the OCSP responder.
     */
    certOcsp?: pulumi.Input<string>;
    /**
     * Specifies the issuer certificate.
     */
    issuerCert?: pulumi.Input<string>;
    /**
     * The content of the key.
     */
    keyContent: pulumi.Input<string>;
    /**
     * full path of the SSL key on the BIGIP.
     */
    keyFullPath?: pulumi.Input<string>;
    /**
     * Name of the SSL key to be Imported on to BIGIP.
     */
    keyName: pulumi.Input<string>;
    /**
     * Partition on to SSL certificate and key to be imported.
     */
    partition?: pulumi.Input<string>;
    /**
     * Passphrase on the SSL key.
     */
    passphrase?: pulumi.Input<string>;
}
