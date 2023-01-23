// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.net;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.net.SelfIpArgs;
import com.pulumi.f5bigip.net.inputs.SelfIpState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.net.SelfIp` Manages a selfip configuration
 * 
 * Resource should be named with their `full path`. The full path is the combination of the `partition + name of the resource`, for example `/Common/my-selfip`.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.net.Vlan;
 * import com.pulumi.f5bigip.net.VlanArgs;
 * import com.pulumi.f5bigip.net.inputs.VlanInterfaceArgs;
 * import com.pulumi.f5bigip.net.SelfIp;
 * import com.pulumi.f5bigip.net.SelfIpArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var vlan1 = new Vlan(&#34;vlan1&#34;, VlanArgs.builder()        
 *             .name(&#34;/Common/Internal&#34;)
 *             .tag(101)
 *             .interfaces(VlanInterfaceArgs.builder()
 *                 .vlanport(1.2)
 *                 .tagged(false)
 *                 .build())
 *             .build());
 * 
 *         var selfip1 = new SelfIp(&#34;selfip1&#34;, SelfIpArgs.builder()        
 *             .name(&#34;/Common/internalselfIP&#34;)
 *             .ip(&#34;11.1.1.1/24&#34;)
 *             .vlan(&#34;/Common/internal&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vlan1)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Example usage with `port_lockdown`
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.net.SelfIp;
 * import com.pulumi.f5bigip.net.SelfIpArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var selfip1 = new SelfIp(&#34;selfip1&#34;, SelfIpArgs.builder()        
 *             .name(&#34;/Common/internalselfIP&#34;)
 *             .ip(&#34;11.1.1.1/24&#34;)
 *             .vlan(&#34;/Common/internal&#34;)
 *             .trafficGroup(&#34;traffic-group-1&#34;)
 *             .portLockdowns(            
 *                 &#34;tcp:4040&#34;,
 *                 &#34;udp:5050&#34;,
 *                 &#34;egp:0&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(bigip_net_vlan.vlan1())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Example usage with `port_lockdown` set to `[&#34;none&#34;]`
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.net.SelfIp;
 * import com.pulumi.f5bigip.net.SelfIpArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var selfip1 = new SelfIp(&#34;selfip1&#34;, SelfIpArgs.builder()        
 *             .name(&#34;/Common/internalselfIP&#34;)
 *             .ip(&#34;11.1.1.1/24&#34;)
 *             .vlan(&#34;/Common/internal&#34;)
 *             .trafficGroup(&#34;traffic-group-1&#34;)
 *             .portLockdowns(&#34;none&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(bigip_net_vlan.vlan1())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Example usage with route domain embedded in the `ip`
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.net.SelfIp;
 * import com.pulumi.f5bigip.net.SelfIpArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var selfip1 = new SelfIp(&#34;selfip1&#34;, SelfIpArgs.builder()        
 *             .name(&#34;/Common/internalselfIP&#34;)
 *             .ip(&#34;11.1.1.1%4/24&#34;)
 *             .vlan(&#34;/Common/internal&#34;)
 *             .trafficGroup(&#34;traffic-group-1&#34;)
 *             .portLockdowns(&#34;none&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(bigip_net_vlan.vlan1())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="f5bigip:net/selfIp:SelfIp")
public class SelfIp extends com.pulumi.resources.CustomResource {
    /**
     * The Self IP&#39;s address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
     * 
     */
    @Export(name="ip", type=String.class, parameters={})
    private Output<String> ip;

    /**
     * @return The Self IP&#39;s address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * Name of the selfip
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the selfip
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the port lockdown, defaults to `Allow None` if not specified.
     * 
     */
    @Export(name="portLockdowns", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> portLockdowns;

    /**
     * @return Specifies the port lockdown, defaults to `Allow None` if not specified.
     * 
     */
    public Output<Optional<List<String>>> portLockdowns() {
        return Codegen.optional(this.portLockdowns);
    }
    /**
     * Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     * 
     */
    @Export(name="trafficGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> trafficGroup;

    /**
     * @return Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
     * 
     */
    public Output<Optional<String>> trafficGroup() {
        return Codegen.optional(this.trafficGroup);
    }
    /**
     * Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     * 
     */
    @Export(name="vlan", type=String.class, parameters={})
    private Output<String> vlan;

    /**
     * @return Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
     * 
     */
    public Output<String> vlan() {
        return this.vlan;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SelfIp(String name) {
        this(name, SelfIpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SelfIp(String name, SelfIpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SelfIp(String name, SelfIpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:net/selfIp:SelfIp", name, args == null ? SelfIpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SelfIp(String name, Output<String> id, @Nullable SelfIpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:net/selfIp:SelfIp", name, state, makeResourceOptions(options, id));
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
    public static SelfIp get(String name, Output<String> id, @Nullable SelfIpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SelfIp(name, id, state, options);
    }
}