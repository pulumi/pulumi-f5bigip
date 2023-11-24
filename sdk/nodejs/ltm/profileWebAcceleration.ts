// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileWebAcceleration` Configures a custom web-acceleration profile for use.
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/sample-resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const sample_resource = new f5bigip.ltm.ProfileWebAcceleration("sample-resource", {
 *     cacheMaxEntries: 201,
 *     cacheSize: 101,
 *     defaultsFrom: "/Common/test2",
 *     name: "/Common/sample-resource",
 * });
 * ```
 */
export class ProfileWebAcceleration extends pulumi.CustomResource {
    /**
     * Get an existing ProfileWebAcceleration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileWebAccelerationState, opts?: pulumi.CustomResourceOptions): ProfileWebAcceleration {
        return new ProfileWebAcceleration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileWebAcceleration:ProfileWebAcceleration';

    /**
     * Returns true if the given object is an instance of ProfileWebAcceleration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileWebAcceleration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileWebAcceleration.__pulumiType;
    }

    /**
     * Specifies how quickly the system ages a cache entry. The aging rate ranges from 0 (slowest aging) to 10 (fastest aging). The default value is `9`.
     */
    public readonly cacheAgingRate!: pulumi.Output<number>;
    /**
     * Specifies which cache disabling headers sent by clients the system ignores. The default value is `all`.
     */
    public readonly cacheClientCacheControlMode!: pulumi.Output<string>;
    /**
     * Inserts Age and Date headers in the response. The default value is `enabled`.
     */
    public readonly cacheInsertAgeHeader!: pulumi.Output<string>;
    /**
     * Specifies how long the system considers the cached content to be valid. The default value is `3600 seconds`.
     */
    public readonly cacheMaxAge!: pulumi.Output<number>;
    /**
     * Specifies the maximum number of entries that can be in the cache. The default value is `0` (zero), which means that the system does not limit the maximum entries.
     */
    public readonly cacheMaxEntries!: pulumi.Output<number>;
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     */
    public readonly cacheObjectMaxSize!: pulumi.Output<number>;
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     */
    public readonly cacheObjectMinSize!: pulumi.Output<number>;
    /**
     * Specifies the maximum size for the cache. When the cache reaches the maximum size, the system starts removing the oldest entries. The default value is `100 megabytes`.
     */
    public readonly cacheSize!: pulumi.Output<number>;
    /**
     * Configures a list of URIs to exclude from the cache. The default value of `none` specifies no URIs are excluded.
     */
    public readonly cacheUriExcludes!: pulumi.Output<string[]>;
    /**
     * Configures a list of URIs to include in the cache even if they would normally be excluded due to factors like object size or HTTP request type. The default value of none specifies no URIs are to be forced into the cache.
     */
    public readonly cacheUriIncludeOverrides!: pulumi.Output<string[]>;
    /**
     * Configures a list of URIs to include in the cache. The default value of `.*` specifies that all URIs are cacheable.
     */
    public readonly cacheUriIncludes!: pulumi.Output<string[]>;
    /**
     * Configures a list of URIs to keep in the cache. The pinning process keeps URIs in cache when they would normally be evicted to make room for more active URIs.
     */
    public readonly cacheUriPinneds!: pulumi.Output<string[]>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Specifies the name of the web acceleration profile service ,name of Profile should be full path. Full path is the combination of the `partition + web acceleration profile name`,For example `/Common/sample-resource`.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ProfileWebAcceleration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileWebAccelerationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileWebAccelerationArgs | ProfileWebAccelerationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileWebAccelerationState | undefined;
            resourceInputs["cacheAgingRate"] = state ? state.cacheAgingRate : undefined;
            resourceInputs["cacheClientCacheControlMode"] = state ? state.cacheClientCacheControlMode : undefined;
            resourceInputs["cacheInsertAgeHeader"] = state ? state.cacheInsertAgeHeader : undefined;
            resourceInputs["cacheMaxAge"] = state ? state.cacheMaxAge : undefined;
            resourceInputs["cacheMaxEntries"] = state ? state.cacheMaxEntries : undefined;
            resourceInputs["cacheObjectMaxSize"] = state ? state.cacheObjectMaxSize : undefined;
            resourceInputs["cacheObjectMinSize"] = state ? state.cacheObjectMinSize : undefined;
            resourceInputs["cacheSize"] = state ? state.cacheSize : undefined;
            resourceInputs["cacheUriExcludes"] = state ? state.cacheUriExcludes : undefined;
            resourceInputs["cacheUriIncludeOverrides"] = state ? state.cacheUriIncludeOverrides : undefined;
            resourceInputs["cacheUriIncludes"] = state ? state.cacheUriIncludes : undefined;
            resourceInputs["cacheUriPinneds"] = state ? state.cacheUriPinneds : undefined;
            resourceInputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ProfileWebAccelerationArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["cacheAgingRate"] = args ? args.cacheAgingRate : undefined;
            resourceInputs["cacheClientCacheControlMode"] = args ? args.cacheClientCacheControlMode : undefined;
            resourceInputs["cacheInsertAgeHeader"] = args ? args.cacheInsertAgeHeader : undefined;
            resourceInputs["cacheMaxAge"] = args ? args.cacheMaxAge : undefined;
            resourceInputs["cacheMaxEntries"] = args ? args.cacheMaxEntries : undefined;
            resourceInputs["cacheObjectMaxSize"] = args ? args.cacheObjectMaxSize : undefined;
            resourceInputs["cacheObjectMinSize"] = args ? args.cacheObjectMinSize : undefined;
            resourceInputs["cacheSize"] = args ? args.cacheSize : undefined;
            resourceInputs["cacheUriExcludes"] = args ? args.cacheUriExcludes : undefined;
            resourceInputs["cacheUriIncludeOverrides"] = args ? args.cacheUriIncludeOverrides : undefined;
            resourceInputs["cacheUriIncludes"] = args ? args.cacheUriIncludes : undefined;
            resourceInputs["cacheUriPinneds"] = args ? args.cacheUriPinneds : undefined;
            resourceInputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProfileWebAcceleration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileWebAcceleration resources.
 */
export interface ProfileWebAccelerationState {
    /**
     * Specifies how quickly the system ages a cache entry. The aging rate ranges from 0 (slowest aging) to 10 (fastest aging). The default value is `9`.
     */
    cacheAgingRate?: pulumi.Input<number>;
    /**
     * Specifies which cache disabling headers sent by clients the system ignores. The default value is `all`.
     */
    cacheClientCacheControlMode?: pulumi.Input<string>;
    /**
     * Inserts Age and Date headers in the response. The default value is `enabled`.
     */
    cacheInsertAgeHeader?: pulumi.Input<string>;
    /**
     * Specifies how long the system considers the cached content to be valid. The default value is `3600 seconds`.
     */
    cacheMaxAge?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of entries that can be in the cache. The default value is `0` (zero), which means that the system does not limit the maximum entries.
     */
    cacheMaxEntries?: pulumi.Input<number>;
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     */
    cacheObjectMaxSize?: pulumi.Input<number>;
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     */
    cacheObjectMinSize?: pulumi.Input<number>;
    /**
     * Specifies the maximum size for the cache. When the cache reaches the maximum size, the system starts removing the oldest entries. The default value is `100 megabytes`.
     */
    cacheSize?: pulumi.Input<number>;
    /**
     * Configures a list of URIs to exclude from the cache. The default value of `none` specifies no URIs are excluded.
     */
    cacheUriExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configures a list of URIs to include in the cache even if they would normally be excluded due to factors like object size or HTTP request type. The default value of none specifies no URIs are to be forced into the cache.
     */
    cacheUriIncludeOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configures a list of URIs to include in the cache. The default value of `.*` specifies that all URIs are cacheable.
     */
    cacheUriIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configures a list of URIs to keep in the cache. The pinning process keeps URIs in cache when they would normally be evicted to make room for more active URIs.
     */
    cacheUriPinneds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies the name of the web acceleration profile service ,name of Profile should be full path. Full path is the combination of the `partition + web acceleration profile name`,For example `/Common/sample-resource`.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileWebAcceleration resource.
 */
export interface ProfileWebAccelerationArgs {
    /**
     * Specifies how quickly the system ages a cache entry. The aging rate ranges from 0 (slowest aging) to 10 (fastest aging). The default value is `9`.
     */
    cacheAgingRate?: pulumi.Input<number>;
    /**
     * Specifies which cache disabling headers sent by clients the system ignores. The default value is `all`.
     */
    cacheClientCacheControlMode?: pulumi.Input<string>;
    /**
     * Inserts Age and Date headers in the response. The default value is `enabled`.
     */
    cacheInsertAgeHeader?: pulumi.Input<string>;
    /**
     * Specifies how long the system considers the cached content to be valid. The default value is `3600 seconds`.
     */
    cacheMaxAge?: pulumi.Input<number>;
    /**
     * Specifies the maximum number of entries that can be in the cache. The default value is `0` (zero), which means that the system does not limit the maximum entries.
     */
    cacheMaxEntries?: pulumi.Input<number>;
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     */
    cacheObjectMaxSize?: pulumi.Input<number>;
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     */
    cacheObjectMinSize?: pulumi.Input<number>;
    /**
     * Specifies the maximum size for the cache. When the cache reaches the maximum size, the system starts removing the oldest entries. The default value is `100 megabytes`.
     */
    cacheSize?: pulumi.Input<number>;
    /**
     * Configures a list of URIs to exclude from the cache. The default value of `none` specifies no URIs are excluded.
     */
    cacheUriExcludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configures a list of URIs to include in the cache even if they would normally be excluded due to factors like object size or HTTP request type. The default value of none specifies no URIs are to be forced into the cache.
     */
    cacheUriIncludeOverrides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configures a list of URIs to include in the cache. The default value of `.*` specifies that all URIs are cacheable.
     */
    cacheUriIncludes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Configures a list of URIs to keep in the cache. The pinning process keeps URIs in cache when they would normally be evicted to make room for more active URIs.
     */
    cacheUriPinneds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies the name of the web acceleration profile service ,name of Profile should be full path. Full path is the combination of the `partition + web acceleration profile name`,For example `/Common/sample-resource`.
     */
    name: pulumi.Input<string>;
}
