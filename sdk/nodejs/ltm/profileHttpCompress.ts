// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileHttpCompress`  Virtual server HTTP compression profile configuration
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const sjhttpcompression = new f5bigip.ltm.ProfileHttpCompress("sjhttpcompression", {
 *     contentTypeExcludes: ["nicecontentexclude.com"],
 *     contentTypeIncludes: ["nicecontent.com"],
 *     defaultsFrom: "/Common/httpcompression",
 *     name: "/Common/sjhttpcompression2",
 *     uriExcludes: [
 *         "www.abc.f5.com",
 *         "www.abc2.f5.com",
 *     ],
 *     uriIncludes: ["www.xyzbc.cisco.com"],
 * });
 * ```
 */
export class ProfileHttpCompress extends pulumi.CustomResource {
    /**
     * Get an existing ProfileHttpCompress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileHttpCompressState, opts?: pulumi.CustomResourceOptions): ProfileHttpCompress {
        return new ProfileHttpCompress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileHttpCompress:ProfileHttpCompress';

    /**
     * Returns true if the given object is an instance of ProfileHttpCompress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileHttpCompress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileHttpCompress.__pulumiType;
    }

    /**
     * Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     */
    public readonly contentTypeExcludes!: pulumi.Output<string[]>;
    /**
     * Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     */
    public readonly contentTypeIncludes!: pulumi.Output<string[]>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Name of the profile_httpcompress
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
     */
    public readonly uriExcludes!: pulumi.Output<string[]>;
    /**
     * Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
     */
    public readonly uriIncludes!: pulumi.Output<string[]>;

    /**
     * Create a ProfileHttpCompress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileHttpCompressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileHttpCompressArgs | ProfileHttpCompressState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileHttpCompressState | undefined;
            inputs["contentTypeExcludes"] = state ? state.contentTypeExcludes : undefined;
            inputs["contentTypeIncludes"] = state ? state.contentTypeIncludes : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["uriExcludes"] = state ? state.uriExcludes : undefined;
            inputs["uriIncludes"] = state ? state.uriIncludes : undefined;
        } else {
            const args = argsOrState as ProfileHttpCompressArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["contentTypeExcludes"] = args ? args.contentTypeExcludes : undefined;
            inputs["contentTypeIncludes"] = args ? args.contentTypeIncludes : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["uriExcludes"] = args ? args.uriExcludes : undefined;
            inputs["uriIncludes"] = args ? args.uriIncludes : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ProfileHttpCompress.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileHttpCompress resources.
 */
export interface ProfileHttpCompressState {
    /**
     * Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     */
    contentTypeExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     */
    contentTypeIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Name of the profile_httpcompress
     */
    name?: pulumi.Input<string>;
    /**
     * Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
     */
    uriExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
     */
    uriIncludes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ProfileHttpCompress resource.
 */
export interface ProfileHttpCompressArgs {
    /**
     * Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     */
    contentTypeExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     */
    contentTypeIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Name of the profile_httpcompress
     */
    name: pulumi.Input<string>;
    /**
     * Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
     */
    uriExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
     */
    uriIncludes?: pulumi.Input<pulumi.Input<string>[]>;
}
