// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ssl.Certificate` This resource will import SSL certificates on BIG-IP LTM.
 * Certificates can be imported from certificate files on the local disk, in PEM format
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * from "fs";
 *
 * const test_cert = new f5bigip.ssl.Certificate("test-cert", {
 *     name: "servercert.crt",
 *     content: fs.readFileSync("servercert.crt"),
 *     partition: "Common",
 * });
 * ```
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateState, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ssl/certificate:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Content of certificate on Disk
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Name of the SSL Certificate to be Imported on to BIGIP
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Partition on to SSL Certificate to be imported
     */
    public readonly partition!: pulumi.Output<string | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs | CertificateState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CertificateState | undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partition"] = state ? state.partition : undefined;
        } else {
            const args = argsOrState as CertificateArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["content"] = args ? args.content : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partition"] = args ? args.partition : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Certificate resources.
 */
export interface CertificateState {
    /**
     * Content of certificate on Disk
     */
    readonly content?: pulumi.Input<string>;
    /**
     * Name of the SSL Certificate to be Imported on to BIGIP
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Partition on to SSL Certificate to be imported
     */
    readonly partition?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Content of certificate on Disk
     */
    readonly content: pulumi.Input<string>;
    /**
     * Name of the SSL Certificate to be Imported on to BIGIP
     */
    readonly name: pulumi.Input<string>;
    /**
     * Partition on to SSL Certificate to be imported
     */
    readonly partition?: pulumi.Input<string>;
}
