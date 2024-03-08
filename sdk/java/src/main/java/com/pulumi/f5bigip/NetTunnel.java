// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.NetTunnelArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.NetTunnelState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.NetTunnel` Manages a tunnel configuration
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.NetTunnel;
 * import com.pulumi.f5bigip.NetTunnelArgs;
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
 *         var example1 = new NetTunnel(&#34;example1&#34;, NetTunnelArgs.builder()        
 *             .localAddress(&#34;192.16.81.240&#34;)
 *             .name(&#34;example1&#34;)
 *             .profile(&#34;/Common/dslite&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:index/netTunnel:NetTunnel")
public class NetTunnel extends com.pulumi.resources.CustomResource {
    /**
     * The application service that the object belongs to
     * 
     */
    @Export(name="appService", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> appService;

    /**
     * @return The application service that the object belongs to
     * 
     */
    public Output<Optional<String>> appService() {
        return Codegen.optional(this.appService);
    }
    /**
     * Specifies whether auto lasthop is enabled or not
     * 
     */
    @Export(name="autoLastHop", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> autoLastHop;

    /**
     * @return Specifies whether auto lasthop is enabled or not
     * 
     */
    public Output<Optional<String>> autoLastHop() {
        return Codegen.optional(this.autoLastHop);
    }
    /**
     * User defined description
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User defined description
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Specifies an idle timeout for wildcard tunnels in seconds
     * 
     */
    @Export(name="idleTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> idleTimeout;

    /**
     * @return Specifies an idle timeout for wildcard tunnels in seconds
     * 
     */
    public Output<Optional<Integer>> idleTimeout() {
        return Codegen.optional(this.idleTimeout);
    }
    /**
     * The key field may represent different values depending on the type of the tunnel
     * 
     */
    @Export(name="key", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> key;

    /**
     * @return The key field may represent different values depending on the type of the tunnel
     * 
     */
    public Output<Optional<Integer>> key() {
        return Codegen.optional(this.key);
    }
    /**
     * Specifies a local IP address. This option is required
     * 
     */
    @Export(name="localAddress", refs={String.class}, tree="[0]")
    private Output<String> localAddress;

    /**
     * @return Specifies a local IP address. This option is required
     * 
     */
    public Output<String> localAddress() {
        return this.localAddress;
    }
    /**
     * Specifies how the tunnel carries traffic
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mode;

    /**
     * @return Specifies how the tunnel carries traffic
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
    }
    /**
     * Specifies the maximum transmission unit (MTU) of the tunnel
     * 
     */
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> mtu;

    /**
     * @return Specifies the maximum transmission unit (MTU) of the tunnel
     * 
     */
    public Output<Optional<Integer>> mtu() {
        return Codegen.optional(this.mtu);
    }
    /**
     * Name of the tunnel
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the tunnel
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Displays the admin-partition within which this component resides
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partition;

    /**
     * @return Displays the admin-partition within which this component resides
     * 
     */
    public Output<Optional<String>> partition() {
        return Codegen.optional(this.partition);
    }
    /**
     * Specifies the profile that you want to associate with the tunnel
     * 
     */
    @Export(name="profile", refs={String.class}, tree="[0]")
    private Output<String> profile;

    /**
     * @return Specifies the profile that you want to associate with the tunnel
     * 
     */
    public Output<String> profile() {
        return this.profile;
    }
    /**
     * Specifies a remote IP address
     * 
     */
    @Export(name="remoteAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remoteAddress;

    /**
     * @return Specifies a remote IP address
     * 
     */
    public Output<Optional<String>> remoteAddress() {
        return Codegen.optional(this.remoteAddress);
    }
    /**
     * Specifies a secondary non-floating IP address when the local-address is set to a floating address
     * 
     */
    @Export(name="secondaryAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secondaryAddress;

    /**
     * @return Specifies a secondary non-floating IP address when the local-address is set to a floating address
     * 
     */
    public Output<Optional<String>> secondaryAddress() {
        return Codegen.optional(this.secondaryAddress);
    }
    /**
     * Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
     * 
     */
    @Export(name="tos", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tos;

    /**
     * @return Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
     * 
     */
    public Output<Optional<String>> tos() {
        return Codegen.optional(this.tos);
    }
    /**
     * Specifies a traffic-group for use with the tunnel
     * 
     */
    @Export(name="trafficGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trafficGroup;

    /**
     * @return Specifies a traffic-group for use with the tunnel
     * 
     */
    public Output<Optional<String>> trafficGroup() {
        return Codegen.optional(this.trafficGroup);
    }
    /**
     * Enables or disables the tunnel to be transparent
     * 
     */
    @Export(name="transparent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> transparent;

    /**
     * @return Enables or disables the tunnel to be transparent
     * 
     */
    public Output<Optional<String>> transparent() {
        return Codegen.optional(this.transparent);
    }
    /**
     * Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
     * 
     */
    @Export(name="usePmtu", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> usePmtu;

    /**
     * @return Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
     * 
     */
    public Output<Optional<String>> usePmtu() {
        return Codegen.optional(this.usePmtu);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetTunnel(String name) {
        this(name, NetTunnelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetTunnel(String name, NetTunnelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetTunnel(String name, NetTunnelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/netTunnel:NetTunnel", name, args == null ? NetTunnelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetTunnel(String name, Output<String> id, @Nullable NetTunnelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/netTunnel:NetTunnel", name, state, makeResourceOptions(options, id));
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
    public static NetTunnel get(String name, Output<String> id, @Nullable NetTunnelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetTunnel(name, id, state, options);
    }
}
