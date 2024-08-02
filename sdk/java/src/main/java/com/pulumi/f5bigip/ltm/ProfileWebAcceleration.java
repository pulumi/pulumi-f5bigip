// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileWebAccelerationArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileWebAccelerationState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileWebAcceleration` Configures a custom web-acceleration profile for use.
 * 
 * For resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name of the resource. For example /Common/sample-resource.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.ProfileWebAcceleration;
 * import com.pulumi.f5bigip.ltm.ProfileWebAccelerationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var sample_resource = new ProfileWebAcceleration("sample-resource", ProfileWebAccelerationArgs.builder()
 *             .name("/Common/sample-resource")
 *             .defaultsFrom("/Common/test2")
 *             .cacheSize(101)
 *             .cacheMaxEntries(201)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:ltm/profileWebAcceleration:ProfileWebAcceleration")
public class ProfileWebAcceleration extends com.pulumi.resources.CustomResource {
    /**
     * Specifies how quickly the system ages a cache entry. The aging rate ranges from 0 (slowest aging) to 10 (fastest aging). The default value is `9`.
     * 
     */
    @Export(name="cacheAgingRate", refs={Integer.class}, tree="[0]")
    private Output<Integer> cacheAgingRate;

    /**
     * @return Specifies how quickly the system ages a cache entry. The aging rate ranges from 0 (slowest aging) to 10 (fastest aging). The default value is `9`.
     * 
     */
    public Output<Integer> cacheAgingRate() {
        return this.cacheAgingRate;
    }
    /**
     * Specifies which cache disabling headers sent by clients the system ignores. The default value is `all`.
     * 
     */
    @Export(name="cacheClientCacheControlMode", refs={String.class}, tree="[0]")
    private Output<String> cacheClientCacheControlMode;

    /**
     * @return Specifies which cache disabling headers sent by clients the system ignores. The default value is `all`.
     * 
     */
    public Output<String> cacheClientCacheControlMode() {
        return this.cacheClientCacheControlMode;
    }
    /**
     * Inserts Age and Date headers in the response. The default value is `enabled`.
     * 
     */
    @Export(name="cacheInsertAgeHeader", refs={String.class}, tree="[0]")
    private Output<String> cacheInsertAgeHeader;

    /**
     * @return Inserts Age and Date headers in the response. The default value is `enabled`.
     * 
     */
    public Output<String> cacheInsertAgeHeader() {
        return this.cacheInsertAgeHeader;
    }
    /**
     * Specifies how long the system considers the cached content to be valid. The default value is `3600 seconds`.
     * 
     */
    @Export(name="cacheMaxAge", refs={Integer.class}, tree="[0]")
    private Output<Integer> cacheMaxAge;

    /**
     * @return Specifies how long the system considers the cached content to be valid. The default value is `3600 seconds`.
     * 
     */
    public Output<Integer> cacheMaxAge() {
        return this.cacheMaxAge;
    }
    /**
     * Specifies the maximum number of entries that can be in the cache. The default value is `0` (zero), which means that the system does not limit the maximum entries.
     * 
     */
    @Export(name="cacheMaxEntries", refs={Integer.class}, tree="[0]")
    private Output<Integer> cacheMaxEntries;

    /**
     * @return Specifies the maximum number of entries that can be in the cache. The default value is `0` (zero), which means that the system does not limit the maximum entries.
     * 
     */
    public Output<Integer> cacheMaxEntries() {
        return this.cacheMaxEntries;
    }
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     * 
     */
    @Export(name="cacheObjectMaxSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> cacheObjectMaxSize;

    /**
     * @return Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     * 
     */
    public Output<Integer> cacheObjectMaxSize() {
        return this.cacheObjectMaxSize;
    }
    /**
     * Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     * 
     */
    @Export(name="cacheObjectMinSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> cacheObjectMinSize;

    /**
     * @return Specifies the smallest object that the system considers eligible for caching. The default value is `500 bytes`.
     * 
     */
    public Output<Integer> cacheObjectMinSize() {
        return this.cacheObjectMinSize;
    }
    /**
     * Specifies the maximum size for the cache. When the cache reaches the maximum size, the system starts removing the oldest entries. The default value is `100 megabytes`.
     * 
     */
    @Export(name="cacheSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> cacheSize;

    /**
     * @return Specifies the maximum size for the cache. When the cache reaches the maximum size, the system starts removing the oldest entries. The default value is `100 megabytes`.
     * 
     */
    public Output<Integer> cacheSize() {
        return this.cacheSize;
    }
    /**
     * Configures a list of URIs to exclude from the cache. The default value of `none` specifies no URIs are excluded.
     * 
     */
    @Export(name="cacheUriExcludes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> cacheUriExcludes;

    /**
     * @return Configures a list of URIs to exclude from the cache. The default value of `none` specifies no URIs are excluded.
     * 
     */
    public Output<List<String>> cacheUriExcludes() {
        return this.cacheUriExcludes;
    }
    /**
     * Configures a list of URIs to include in the cache even if they would normally be excluded due to factors like object size or HTTP request type. The default value of none specifies no URIs are to be forced into the cache.
     * 
     */
    @Export(name="cacheUriIncludeOverrides", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> cacheUriIncludeOverrides;

    /**
     * @return Configures a list of URIs to include in the cache even if they would normally be excluded due to factors like object size or HTTP request type. The default value of none specifies no URIs are to be forced into the cache.
     * 
     */
    public Output<List<String>> cacheUriIncludeOverrides() {
        return this.cacheUriIncludeOverrides;
    }
    /**
     * Configures a list of URIs to include in the cache. The default value of `.*` specifies that all URIs are cacheable.
     * 
     */
    @Export(name="cacheUriIncludes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> cacheUriIncludes;

    /**
     * @return Configures a list of URIs to include in the cache. The default value of `.*` specifies that all URIs are cacheable.
     * 
     */
    public Output<List<String>> cacheUriIncludes() {
        return this.cacheUriIncludes;
    }
    /**
     * Configures a list of URIs to keep in the cache. The pinning process keeps URIs in cache when they would normally be evicted to make room for more active URIs.
     * 
     */
    @Export(name="cacheUriPinneds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> cacheUriPinneds;

    /**
     * @return Configures a list of URIs to keep in the cache. The pinning process keeps URIs in cache when they would normally be evicted to make room for more active URIs.
     * 
     */
    public Output<List<String>> cacheUriPinneds() {
        return this.cacheUriPinneds;
    }
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    @Export(name="defaultsFrom", refs={String.class}, tree="[0]")
    private Output<String> defaultsFrom;

    /**
     * @return Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    public Output<String> defaultsFrom() {
        return this.defaultsFrom;
    }
    /**
     * Specifies the name of the web acceleration profile service ,name of Profile should be full path. Full path is the combination of the `partition + web acceleration profile name`,For example `/Common/sample-resource`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the name of the web acceleration profile service ,name of Profile should be full path. Full path is the combination of the `partition + web acceleration profile name`,For example `/Common/sample-resource`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileWebAcceleration(String name) {
        this(name, ProfileWebAccelerationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileWebAcceleration(String name, ProfileWebAccelerationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileWebAcceleration(String name, ProfileWebAccelerationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileWebAcceleration:ProfileWebAcceleration", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileWebAcceleration(String name, Output<String> id, @Nullable ProfileWebAccelerationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileWebAcceleration:ProfileWebAcceleration", name, state, makeResourceOptions(options, id));
    }

    private static ProfileWebAccelerationArgs makeArgs(ProfileWebAccelerationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProfileWebAccelerationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static ProfileWebAcceleration get(String name, Output<String> id, @Nullable ProfileWebAccelerationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileWebAcceleration(name, id, state, options);
    }
}
