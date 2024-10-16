// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `bigipLtmRewriteProfile` Configures ltm policies to manage traffic assigned to a virtual server
 *
 * For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-profile`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test_profile = new f5bigip.ltm.ProfileRewrite("test-profile", {
 *     name: "/Common/tf_profile",
 *     defaultsFrom: "/Common/rewrite",
 *     bypassLists: ["http://notouch.com"],
 *     rewriteLists: ["http://some.com"],
 *     rewriteMode: "portal",
 *     cacheType: "cache-img-css-js",
 *     caFile: "/Common/ca-bundle.crt",
 *     crlFile: "none",
 *     signingCert: "/Common/default.crt",
 *     signingKey: "/Common/default.key",
 *     splitTunneling: "true",
 * });
 * const test_profile2 = new f5bigip.ltm.ProfileRewrite("test-profile2", {
 *     name: "/Common/tf_profile_translate",
 *     defaultsFrom: "/Common/rewrite",
 *     rewriteMode: "uri-translation",
 *     requests: [{
 *         insertXfwdFor: "enabled",
 *         insertXfwdHost: "disabled",
 *         insertXfwdProtocol: "enabled",
 *         rewriteHeaders: "disabled",
 *     }],
 *     responses: [{
 *         rewriteContent: "enabled",
 *         rewriteHeaders: "disabled",
 *     }],
 *     cookieRules: [
 *         {
 *             ruleName: "cookie1",
 *             clientDomain: "wrong.com",
 *             clientPath: "/this/",
 *             serverDomain: "wrong.com",
 *             serverPath: "/this/",
 *         },
 *         {
 *             ruleName: "cookie2",
 *             clientDomain: "incorrect.com",
 *             clientPath: "/this/",
 *             serverDomain: "absolute.com",
 *             serverPath: "/this/",
 *         },
 *     ],
 * });
 * ```
 */
export class ProfileRewrite extends pulumi.CustomResource {
    /**
     * Get an existing ProfileRewrite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileRewriteState, opts?: pulumi.CustomResourceOptions): ProfileRewrite {
        return new ProfileRewrite(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileRewrite:ProfileRewrite';

    /**
     * Returns true if the given object is an instance of ProfileRewrite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileRewrite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileRewrite.__pulumiType;
    }

    /**
     * Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
     */
    public readonly bypassLists!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
     */
    public readonly caFile!: pulumi.Output<string>;
    /**
     * Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
     */
    public readonly cacheType!: pulumi.Output<string | undefined>;
    /**
     * Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
     */
    public readonly cookieRules!: pulumi.Output<outputs.ltm.ProfileRewriteCookieRule[] | undefined>;
    /**
     * Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
     */
    public readonly crlFile!: pulumi.Output<string | undefined>;
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
     */
    public readonly defaultsFrom!: pulumi.Output<string | undefined>;
    /**
     * Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Block type. Each request is block type with following arguments.
     */
    public readonly requests!: pulumi.Output<outputs.ltm.ProfileRewriteRequest[]>;
    /**
     * Block type. Each request is block type with following arguments.
     */
    public readonly responses!: pulumi.Output<outputs.ltm.ProfileRewriteResponse[]>;
    /**
     * Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
     */
    public readonly rewriteLists!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
     */
    public readonly rewriteMode!: pulumi.Output<string>;
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
     */
    public readonly signingCert!: pulumi.Output<string>;
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
     */
    public readonly signingKey!: pulumi.Output<string>;
    /**
     * Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
     */
    public readonly signingKeyPassword!: pulumi.Output<string>;
    /**
     * Specifies the type of Client caching. Valid choices are: `true, false`
     */
    public readonly splitTunneling!: pulumi.Output<string>;

    /**
     * Create a ProfileRewrite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileRewriteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileRewriteArgs | ProfileRewriteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileRewriteState | undefined;
            resourceInputs["bypassLists"] = state ? state.bypassLists : undefined;
            resourceInputs["caFile"] = state ? state.caFile : undefined;
            resourceInputs["cacheType"] = state ? state.cacheType : undefined;
            resourceInputs["cookieRules"] = state ? state.cookieRules : undefined;
            resourceInputs["crlFile"] = state ? state.crlFile : undefined;
            resourceInputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["requests"] = state ? state.requests : undefined;
            resourceInputs["responses"] = state ? state.responses : undefined;
            resourceInputs["rewriteLists"] = state ? state.rewriteLists : undefined;
            resourceInputs["rewriteMode"] = state ? state.rewriteMode : undefined;
            resourceInputs["signingCert"] = state ? state.signingCert : undefined;
            resourceInputs["signingKey"] = state ? state.signingKey : undefined;
            resourceInputs["signingKeyPassword"] = state ? state.signingKeyPassword : undefined;
            resourceInputs["splitTunneling"] = state ? state.splitTunneling : undefined;
        } else {
            const args = argsOrState as ProfileRewriteArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.rewriteMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rewriteMode'");
            }
            resourceInputs["bypassLists"] = args ? args.bypassLists : undefined;
            resourceInputs["caFile"] = args ? args.caFile : undefined;
            resourceInputs["cacheType"] = args ? args.cacheType : undefined;
            resourceInputs["cookieRules"] = args ? args.cookieRules : undefined;
            resourceInputs["crlFile"] = args ? args.crlFile : undefined;
            resourceInputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["requests"] = args ? args.requests : undefined;
            resourceInputs["responses"] = args ? args.responses : undefined;
            resourceInputs["rewriteLists"] = args ? args.rewriteLists : undefined;
            resourceInputs["rewriteMode"] = args ? args.rewriteMode : undefined;
            resourceInputs["signingCert"] = args ? args.signingCert : undefined;
            resourceInputs["signingKey"] = args ? args.signingKey : undefined;
            resourceInputs["signingKeyPassword"] = args?.signingKeyPassword ? pulumi.secret(args.signingKeyPassword) : undefined;
            resourceInputs["splitTunneling"] = args ? args.splitTunneling : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["signingKeyPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProfileRewrite.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileRewrite resources.
 */
export interface ProfileRewriteState {
    /**
     * Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
     */
    bypassLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
     */
    caFile?: pulumi.Input<string>;
    /**
     * Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
     */
    cacheType?: pulumi.Input<string>;
    /**
     * Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
     */
    cookieRules?: pulumi.Input<pulumi.Input<inputs.ltm.ProfileRewriteCookieRule>[]>;
    /**
     * Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
     */
    crlFile?: pulumi.Input<string>;
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
     */
    name?: pulumi.Input<string>;
    /**
     * Block type. Each request is block type with following arguments.
     */
    requests?: pulumi.Input<pulumi.Input<inputs.ltm.ProfileRewriteRequest>[]>;
    /**
     * Block type. Each request is block type with following arguments.
     */
    responses?: pulumi.Input<pulumi.Input<inputs.ltm.ProfileRewriteResponse>[]>;
    /**
     * Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
     */
    rewriteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
     */
    rewriteMode?: pulumi.Input<string>;
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
     */
    signingCert?: pulumi.Input<string>;
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
     */
    signingKey?: pulumi.Input<string>;
    /**
     * Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
     */
    signingKeyPassword?: pulumi.Input<string>;
    /**
     * Specifies the type of Client caching. Valid choices are: `true, false`
     */
    splitTunneling?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileRewrite resource.
 */
export interface ProfileRewriteArgs {
    /**
     * Specifies a list of URIs to bypass inside a web page when the page is accessed using Portal Access.
     */
    bypassLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies a CA against which to verify signed Java applets signatures. (name should be in full path which is combination of partition and CA file name )
     */
    caFile?: pulumi.Input<string>;
    /**
     * Specifies the type of Client caching. Valid choices are: `cache-css-js, cache-all, no-cache, cache-img-css-js`. Default value: `cache-img-css-js`
     */
    cacheType?: pulumi.Input<string>;
    /**
     * Specifies the cookie rewrite rules. Block type. Each request is block type with following arguments.
     */
    cookieRules?: pulumi.Input<pulumi.Input<inputs.ltm.ProfileRewriteCookieRule>[]>;
    /**
     * Specifies a CRL against which to verify signed Java applets signature certificates. The default option is `none`.
     */
    crlFile?: pulumi.Input<string>;
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `rewrite` profile.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Name of the rewrite profile. ( profile name should be in full path which is combination of partition and profile name )
     */
    name: pulumi.Input<string>;
    /**
     * Block type. Each request is block type with following arguments.
     */
    requests?: pulumi.Input<pulumi.Input<inputs.ltm.ProfileRewriteRequest>[]>;
    /**
     * Block type. Each request is block type with following arguments.
     */
    responses?: pulumi.Input<pulumi.Input<inputs.ltm.ProfileRewriteResponse>[]>;
    /**
     * Specifies a list of URIs to rewrite inside a web page when the page is accessed using Portal Access.
     */
    rewriteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the type of Client caching. Valid choices are: `portal, uri-translation`
     */
    rewriteMode: pulumi.Input<string>;
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and certificate name )
     */
    signingCert?: pulumi.Input<string>;
    /**
     * Specifies a certificate to use for re-signing of signed Java applets after patching. (name should be in full path which is combination of partition and key name )
     */
    signingKey?: pulumi.Input<string>;
    /**
     * Specifies a pass phrase to use for encrypting the private signing key. Since it's a sensitive entity idempotency will fail in the update call.
     */
    signingKeyPassword?: pulumi.Input<string>;
    /**
     * Specifies the type of Client caching. Valid choices are: `true, false`
     */
    splitTunneling?: pulumi.Input<string>;
}