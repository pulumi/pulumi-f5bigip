// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.ProfileOneConnectArgs;
import com.pulumi.f5bigip.ltm.inputs.ProfileOneConnectState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.ProfileOneConnect` Configures a custom profile_oneconnect for use by health checks.
 * 
 * Resources should be named with their &#34;full path&#34;. The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
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
 * import com.pulumi.f5bigip.ltm.ProfileOneConnect;
 * import com.pulumi.f5bigip.ltm.ProfileOneConnectArgs;
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
 *         var test_oneconnect = new ProfileOneConnect("test-oneconnect", ProfileOneConnectArgs.builder()
 *             .name("/Common/test-oneconnect")
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
 * BIG-IP LTM oneconnect profiles can be imported using the `name` , e.g.
 * 
 * ```sh
 * $ pulumi import f5bigip:ltm/profileOneConnect:ProfileOneConnect test-oneconnect /Common/test-oneconnect
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/profileOneConnect:ProfileOneConnect")
public class ProfileOneConnect extends com.pulumi.resources.CustomResource {
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
     * Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
     * 
     */
    @Export(name="idleTimeoutOverride", refs={String.class}, tree="[0]")
    private Output<String> idleTimeoutOverride;

    /**
     * @return Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
     * 
     */
    public Output<String> idleTimeoutOverride() {
        return this.idleTimeoutOverride;
    }
    /**
     * Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
     * 
     */
    @Export(name="limitType", refs={String.class}, tree="[0]")
    private Output<String> limitType;

    /**
     * @return Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
     * 
     */
    public Output<String> limitType() {
        return this.limitType;
    }
    /**
     * Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
     * 
     */
    @Export(name="maxAge", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxAge;

    /**
     * @return Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
     * 
     */
    public Output<Integer> maxAge() {
        return this.maxAge;
    }
    /**
     * Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
     * 
     */
    @Export(name="maxReuse", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxReuse;

    /**
     * @return Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
     * 
     */
    public Output<Integer> maxReuse() {
        return this.maxReuse;
    }
    /**
     * Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
     * 
     */
    @Export(name="maxSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxSize;

    /**
     * @return Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
     * 
     */
    public Output<Integer> maxSize() {
        return this.maxSize;
    }
    /**
     * Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Displays the administrative partition within which this profile resides
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output<String> partition;

    /**
     * @return Displays the administrative partition within which this profile resides
     * 
     */
    public Output<String> partition() {
        return this.partition;
    }
    /**
     * Specify if you want to share the pool, default value is `disabled`.
     * 
     */
    @Export(name="sharePools", refs={String.class}, tree="[0]")
    private Output<String> sharePools;

    /**
     * @return Specify if you want to share the pool, default value is `disabled`.
     * 
     */
    public Output<String> sharePools() {
        return this.sharePools;
    }
    /**
     * Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1&#39;s in binary), causes the system to share only those reused connections originating from the same client IP address.
     * 
     */
    @Export(name="sourceMask", refs={String.class}, tree="[0]")
    private Output<String> sourceMask;

    /**
     * @return Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1&#39;s in binary), causes the system to share only those reused connections originating from the same client IP address.
     * 
     */
    public Output<String> sourceMask() {
        return this.sourceMask;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ProfileOneConnect(String name) {
        this(name, ProfileOneConnectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ProfileOneConnect(String name, ProfileOneConnectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ProfileOneConnect(String name, ProfileOneConnectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileOneConnect:ProfileOneConnect", name, args == null ? ProfileOneConnectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ProfileOneConnect(String name, Output<String> id, @Nullable ProfileOneConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/profileOneConnect:ProfileOneConnect", name, state, makeResourceOptions(options, id));
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
    public static ProfileOneConnect get(String name, Output<String> id, @Nullable ProfileOneConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ProfileOneConnect(name, id, state, options);
    }
}
