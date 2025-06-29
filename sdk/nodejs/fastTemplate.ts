// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.FastTemplate` This resource will import and create FAST template sets on BIG-IP LTM.
 * Template set can be imported from zip archive files on the local disk.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 * import * as std from "@pulumi/std";
 *
 * const foo_template = new f5bigip.FastTemplate("foo-template", {
 *     name: "foo_template",
 *     source: "foo_template.zip",
 *     md5Hash: std.filemd5({
 *         input: "foo_template.zip",
 *     }).then(invoke => invoke.result),
 * });
 * ```
 */
export class FastTemplate extends pulumi.CustomResource {
    /**
     * Get an existing FastTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FastTemplateState, opts?: pulumi.CustomResourceOptions): FastTemplate {
        return new FastTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/fastTemplate:FastTemplate';

    /**
     * Returns true if the given object is an instance of FastTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FastTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FastTemplate.__pulumiType;
    }

    /**
     * MD5 hash of the zip archive file containing FAST template
     */
    public readonly md5Hash!: pulumi.Output<string>;
    /**
     * Name of the FAST template set to be created on to BIGIP
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Path to the zip archive file containing FAST template set on Local Disk
     */
    public readonly source!: pulumi.Output<string>;

    /**
     * Create a FastTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FastTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FastTemplateArgs | FastTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FastTemplateState | undefined;
            resourceInputs["md5Hash"] = state ? state.md5Hash : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["source"] = state ? state.source : undefined;
        } else {
            const args = argsOrState as FastTemplateArgs | undefined;
            if ((!args || args.md5Hash === undefined) && !opts.urn) {
                throw new Error("Missing required property 'md5Hash'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["md5Hash"] = args ? args.md5Hash : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FastTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FastTemplate resources.
 */
export interface FastTemplateState {
    /**
     * MD5 hash of the zip archive file containing FAST template
     */
    md5Hash?: pulumi.Input<string>;
    /**
     * Name of the FAST template set to be created on to BIGIP
     */
    name?: pulumi.Input<string>;
    /**
     * Path to the zip archive file containing FAST template set on Local Disk
     */
    source?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FastTemplate resource.
 */
export interface FastTemplateArgs {
    /**
     * MD5 hash of the zip archive file containing FAST template
     */
    md5Hash: pulumi.Input<string>;
    /**
     * Name of the FAST template set to be created on to BIGIP
     */
    name?: pulumi.Input<string>;
    /**
     * Path to the zip archive file containing FAST template set on Local Disk
     */
    source: pulumi.Input<string>;
}
