// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileFastHttpArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileFastHttpState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileFastHttp` Configures a custom profile_fasthttp for use by health checks.
 * 
 * For resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
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
 * import com.pulumi.f5bigip.ltm.ProfileFastHttp;
 * import com.pulumi.f5bigip.ltm.ProfileFastHttpArgs;
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
 *         var sjfasthttpprofile = new ProfileFastHttp("sjfasthttpprofile", ProfileFastHttpArgs.builder()
 *             .name("/Common/sjfasthttpprofile")
 *             .defaultsFrom("/Common/fasthttp")
 *             .idleTimeout(300)
 *             .connpoolidleTimeoutoverride(0)
 *             .connpoolMaxreuse(2)
 *             .connpoolMaxsize(2048)
 *             .connpoolMinsize(0)
 *             .connpoolReplenish("enabled")
 *             .connpoolStep(4)
 *             .forcehttp10response("disabled")
 *             .maxheaderSize(32768)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:ltm/profileFastHttp:ProfileFastHttp")
public class ProfileFastHttp extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
     * 
     */
    @Export(name="connpoolMaxreuse", refs={Integer.class}, tree="[0]")
    private Output<Integer> connpoolMaxreuse;

    /**
     * @return Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
     * 
     */
    public Output<Integer> connpoolMaxreuse() {
        return this.connpoolMaxreuse;
    }
    /**
     * Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
     * 
     */
    @Export(name="connpoolMaxsize", refs={Integer.class}, tree="[0]")
    private Output<Integer> connpoolMaxsize;

    /**
     * @return Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
     * 
     */
    public Output<Integer> connpoolMaxsize() {
        return this.connpoolMaxsize;
    }
    /**
     * Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
     * 
     */
    @Export(name="connpoolMinsize", refs={Integer.class}, tree="[0]")
    private Output<Integer> connpoolMinsize;

    /**
     * @return Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
     * 
     */
    public Output<Integer> connpoolMinsize() {
        return this.connpoolMinsize;
    }
    /**
     * The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
     * 
     */
    @Export(name="connpoolReplenish", refs={String.class}, tree="[0]")
    private Output<String> connpoolReplenish;

    /**
     * @return The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
     * 
     */
    public Output<String> connpoolReplenish() {
        return this.connpoolReplenish;
    }
    /**
     * Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
     * 
     */
    @Export(name="connpoolStep", refs={Integer.class}, tree="[0]")
    private Output<Integer> connpoolStep;

    /**
     * @return Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
     * 
     */
    public Output<Integer> connpoolStep() {
        return this.connpoolStep;
    }
    /**
     * Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
     * 
     */
    @Export(name="connpoolidleTimeoutoverride", refs={Integer.class}, tree="[0]")
    private Output<Integer> connpoolidleTimeoutoverride;

    /**
     * @return Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
     * 
     */
    public Output<Integer> connpoolidleTimeoutoverride() {
        return this.connpoolidleTimeoutoverride;
    }
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    @Export(name="defaultsFrom", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultsFrom;

    /**
     * @return Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     * 
     */
    public Output<Optional<String>> defaultsFrom() {
        return Codegen.optional(this.defaultsFrom);
    }
    /**
     * Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
     * 
     */
    @Export(name="forcehttp10response", refs={String.class}, tree="[0]")
    private Output<String> forcehttp10response;

    /**
     * @return Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
     * 
     */
    public Output<String> forcehttp10response() {
        return this.forcehttp10response;
    }
    /**
     * Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     * 
     */
    @Export(name="idleTimeout", refs={Integer.class}, tree="[0]")
    private Output<Integer> idleTimeout;

    /**
     * @return Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
     * 
     */
    public Output<Integer> idleTimeout() {
        return this.idleTimeout;
    }
    /**
     * Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
     * 
     */
    @Export(name="maxheaderSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxheaderSize;

    /**
     * @return Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
     * 
     */
    public Output<Integer> maxheaderSize() {
        return this.maxheaderSize;
    }
    /**
     * Name of the profile_fasthttp
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the profile_fasthttp
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileFastHttp(java.lang.String name) {
        this(name, ProfileFastHttpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileFastHttp(java.lang.String name, ProfileFastHttpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileFastHttp(java.lang.String name, ProfileFastHttpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileFastHttp:ProfileFastHttp", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ProfileFastHttp(java.lang.String name, Output<java.lang.String> id, @Nullable ProfileFastHttpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileFastHttp:ProfileFastHttp", name, state, makeResourceOptions(options, id), false);
    }

    private static ProfileFastHttpArgs makeArgs(ProfileFastHttpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProfileFastHttpArgs.Empty : args;
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
    public static ProfileFastHttp get(java.lang.String name, Output<java.lang.String> id, @Nullable ProfileFastHttpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileFastHttp(name, id, state, options);
    }
}
