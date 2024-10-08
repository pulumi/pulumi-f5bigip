// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileHttpCompressArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileHttpCompressState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileHttpCompress`  Virtual server HTTP compression profile configuration
 * 
 * Resources should be named with their `full path`.The full path is the combination of the `partition + name` (example: `/Common/my-httpcompresprofile` ) or  `partition + directory + name` of the resource  (example: `/Common/test/my-httpcompresprofile`)
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
 * import com.pulumi.f5bigip.ltm.ProfileHttpCompress;
 * import com.pulumi.f5bigip.ltm.ProfileHttpCompressArgs;
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
 *         var sjhttpcompression = new ProfileHttpCompress("sjhttpcompression", ProfileHttpCompressArgs.builder()
 *             .name("/Common/sjhttpcompression2")
 *             .defaultsFrom("/Common/httpcompression")
 *             .uriExcludes(            
 *                 "www.abc.f5.com",
 *                 "www.abc2.f5.com")
 *             .uriIncludes("www.xyzbc.cisco.com")
 *             .contentTypeIncludes("nicecontent.com")
 *             .contentTypeExcludes("nicecontentexclude.com")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * BIG-IP LTM HTTP Compress profiles can be imported using the `name`, e.g.
 * 
 * ```sh
 * $ pulumi import f5bigip:ltm/profileHttpCompress:ProfileHttpCompress test-httpcomprs_import /Common/test-httpcomprs
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/profileHttpCompress:ProfileHttpCompress")
public class ProfileHttpCompress extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the maximum number of compressed bytes that the system buffers before inserting a Content-Length header (which specifies the compressed size) into the response. The default is `4096` bytes.
     * 
     */
    @Export(name="compressionBuffersize", refs={Integer.class}, tree="[0]")
    private Output<Integer> compressionBuffersize;

    /**
     * @return Specifies the maximum number of compressed bytes that the system buffers before inserting a Content-Length header (which specifies the compressed size) into the response. The default is `4096` bytes.
     * 
     */
    public Output<Integer> compressionBuffersize() {
        return this.compressionBuffersize;
    }
    /**
     * Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     * 
     */
    @Export(name="contentTypeExcludes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> contentTypeExcludes;

    /**
     * @return Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     * 
     */
    public Output<List<String>> contentTypeExcludes() {
        return this.contentTypeExcludes;
    }
    /**
     * Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     * 
     */
    @Export(name="contentTypeIncludes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> contentTypeIncludes;

    /**
     * @return Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
     * 
     */
    public Output<List<String>> contentTypeIncludes() {
        return this.contentTypeIncludes;
    }
    /**
     * Specifies, when checked (enabled), that the system monitors the percent CPU usage and adjusts compression rates automatically when the CPU usage reaches either the CPU Saver High Threshold or the CPU Saver Low Threshold. The default is `enabled`.
     * 
     */
    @Export(name="cpuSaver", refs={String.class}, tree="[0]")
    private Output<String> cpuSaver;

    /**
     * @return Specifies, when checked (enabled), that the system monitors the percent CPU usage and adjusts compression rates automatically when the CPU usage reaches either the CPU Saver High Threshold or the CPU Saver Low Threshold. The default is `enabled`.
     * 
     */
    public Output<String> cpuSaver() {
        return this.cpuSaver;
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
     * Specifies the degree to which the system compresses the content. Higher compression levels cause the compression process to be slower. The default is 1 - Least Compression (Fastest)
     * 
     */
    @Export(name="gzipCompressionLevel", refs={Integer.class}, tree="[0]")
    private Output<Integer> gzipCompressionLevel;

    /**
     * @return Specifies the degree to which the system compresses the content. Higher compression levels cause the compression process to be slower. The default is 1 - Least Compression (Fastest)
     * 
     */
    public Output<Integer> gzipCompressionLevel() {
        return this.gzipCompressionLevel;
    }
    /**
     * Specifies the number of bytes of memory that the system uses for internal compression buffers when compressing a server response. The default is `8 kilobytes/8192 bytes`.
     * 
     */
    @Export(name="gzipMemoryLevel", refs={Integer.class}, tree="[0]")
    private Output<Integer> gzipMemoryLevel;

    /**
     * @return Specifies the number of bytes of memory that the system uses for internal compression buffers when compressing a server response. The default is `8 kilobytes/8192 bytes`.
     * 
     */
    public Output<Integer> gzipMemoryLevel() {
        return this.gzipMemoryLevel;
    }
    /**
     * Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The default is `16` kilobytes
     * 
     */
    @Export(name="gzipWindowSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> gzipWindowSize;

    /**
     * @return Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The default is `16` kilobytes
     * 
     */
    public Output<Integer> gzipWindowSize() {
        return this.gzipWindowSize;
    }
    /**
     * Specifies, when checked (enabled), that the system does not remove the Accept-Encoding: header from an HTTP request. The default is `disabled`.
     * 
     */
    @Export(name="keepAcceptEncoding", refs={String.class}, tree="[0]")
    private Output<String> keepAcceptEncoding;

    /**
     * @return Specifies, when checked (enabled), that the system does not remove the Accept-Encoding: header from an HTTP request. The default is `disabled`.
     * 
     */
    public Output<String> keepAcceptEncoding() {
        return this.keepAcceptEncoding;
    }
    /**
     * Name of the LTM http compress profile,named with their `full path`.The full path is the combination of the `partition + name` (example: `/Common/my-httpcompresprofile` ) or  `partition + directory + name` of the resource  (example: `my-httpcompresprofile`)
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the LTM http compress profile,named with their `full path`.The full path is the combination of the `partition + name` (example: `/Common/my-httpcompresprofile` ) or  `partition + directory + name` of the resource  (example: `my-httpcompresprofile`)
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
     * 
     */
    @Export(name="uriExcludes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> uriExcludes;

    /**
     * @return Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
     * 
     */
    public Output<List<String>> uriExcludes() {
        return this.uriExcludes;
    }
    /**
     * Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
     * 
     */
    @Export(name="uriIncludes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> uriIncludes;

    /**
     * @return Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
     * 
     */
    public Output<List<String>> uriIncludes() {
        return this.uriIncludes;
    }
    /**
     * Specifies, when checked (enabled), that the system inserts a Vary header into cacheable server responses. The default is `enabled`.
     * 
     */
    @Export(name="varyHeader", refs={String.class}, tree="[0]")
    private Output<String> varyHeader;

    /**
     * @return Specifies, when checked (enabled), that the system inserts a Vary header into cacheable server responses. The default is `enabled`.
     * 
     */
    public Output<String> varyHeader() {
        return this.varyHeader;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileHttpCompress(java.lang.String name) {
        this(name, ProfileHttpCompressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileHttpCompress(java.lang.String name, ProfileHttpCompressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileHttpCompress(java.lang.String name, ProfileHttpCompressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileHttpCompress:ProfileHttpCompress", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProfileHttpCompress(java.lang.String name, Output<java.lang.String> id, @Nullable ProfileHttpCompressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileHttpCompress:ProfileHttpCompress", name, state, makeResourceOptions(options, id), false);
    }

    private static ProfileHttpCompressArgs makeArgs(ProfileHttpCompressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProfileHttpCompressArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ProfileHttpCompress get(java.lang.String name, Output<java.lang.String> id, @Nullable ProfileHttpCompressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileHttpCompress(name, id, state, options);
    }
}
