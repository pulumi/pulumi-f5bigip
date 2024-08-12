// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.PoolAttachmentArgs;
import com.pulumi.f5bigip.ltm.inputs.PoolAttachmentState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.PoolAttachment` Manages nodes membership in pools
 * 
 * ## Example Usage
 * 
 * There are two ways to use `f5bigip.ltm.PoolAttachment` resource for `node` attribute
 * 
 * * It can be reference from `f5bigip.ltm.Node` (or)
 * * It can be specify directly with `ipv4:port`/`fqdn:port`/`ipv6.port` which will also create node and attach member to pool.
 * 
 * &gt; For adding IPv6 node/member to pool it should be specific in `node` attribute in format like `ipv6_address.port`.
 * IPv4 should be specified as `ipv4_address:port`
 * 
 * ### Usage Pool attachment with node/member directly attaching to pool.
 * 
 * node can be specified in format `ipv4:port` / `fqdn:port` / `ipv6.port`
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.Monitor;
 * import com.pulumi.f5bigip.ltm.MonitorArgs;
 * import com.pulumi.f5bigip.ltm.Pool;
 * import com.pulumi.f5bigip.ltm.PoolArgs;
 * import com.pulumi.f5bigip.ltm.PoolAttachment;
 * import com.pulumi.f5bigip.ltm.PoolAttachmentArgs;
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
 *         var monitor = new Monitor("monitor", MonitorArgs.builder()
 *             .name("/Common/terraform_monitor")
 *             .parent("/Common/http")
 *             .send("""
 * GET /some/path
 *             """)
 *             .timeout("999")
 *             .interval("998")
 *             .build());
 * 
 *         var pool = new Pool("pool", PoolArgs.builder()
 *             .name("/Common/terraform-pool")
 *             .loadBalancingMode("round-robin")
 *             .monitors(monitor.name())
 *             .allowSnat("yes")
 *             .allowNat("yes")
 *             .build());
 * 
 *         // attaching ipv4 address with service port
 *         var ipv4NodeAttach = new PoolAttachment("ipv4NodeAttach", PoolAttachmentArgs.builder()
 *             .pool(pool.name())
 *             .node("1.1.1.1:80")
 *             .build());
 * 
 *         // attaching ipv6 address with service port
 *         var ipv6NodeAttach = new PoolAttachment("ipv6NodeAttach", PoolAttachmentArgs.builder()
 *             .pool(pool.name())
 *             .node("2003::4.80")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Usage Pool attachment with node referenced from `f5bigip.ltm.Node`
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ltm.Monitor;
 * import com.pulumi.f5bigip.ltm.MonitorArgs;
 * import com.pulumi.f5bigip.ltm.Pool;
 * import com.pulumi.f5bigip.ltm.PoolArgs;
 * import com.pulumi.f5bigip.ltm.Node;
 * import com.pulumi.f5bigip.ltm.NodeArgs;
 * import com.pulumi.f5bigip.ltm.PoolAttachment;
 * import com.pulumi.f5bigip.ltm.PoolAttachmentArgs;
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
 *         var monitor = new Monitor("monitor", MonitorArgs.builder()
 *             .name("/Common/terraform_monitor")
 *             .parent("/Common/http")
 *             .send("""
 * GET /some/path
 *             """)
 *             .timeout("999")
 *             .interval("998")
 *             .build());
 * 
 *         var pool = new Pool("pool", PoolArgs.builder()
 *             .name("/Common/terraform-pool")
 *             .loadBalancingMode("round-robin")
 *             .monitors(monitor.name())
 *             .allowSnat("yes")
 *             .allowNat("yes")
 *             .build());
 * 
 *         var node = new Node("node", NodeArgs.builder()
 *             .name("/Common/terraform_node")
 *             .address("192.168.30.2")
 *             .build());
 * 
 *         var attachNode = new PoolAttachment("attachNode", PoolAttachmentArgs.builder()
 *             .pool(pool.name())
 *             .node(node.name().applyValue(name -> String.format("%s:80", name)))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Importing
 * 
 * An existing pool attachment (i.e. pool membership) can be imported into this resource by supplying both the pool full path, and the node full path with the relevant port. If the pool or node membership is not found, an error will be returned. An example is below:
 * 
 * ```sh
 * $ terraform import bigip_ltm_pool_attachment.node-pool-attach \
 * 	&#39;{&#34;pool&#34;: &#34;/Common/terraform-pool&#34;, &#34;node&#34;: &#34;/Common/node1:80&#34;}&#39;
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/poolAttachment:PoolAttachment")
public class PoolAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
     * 
     */
    @Export(name="connectionLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectionLimit;

    /**
     * @return Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
     * 
     */
    public Output<Integer> connectionLimit() {
        return this.connectionLimit;
    }
    /**
     * Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
     * 
     */
    @Export(name="connectionRateLimit", refs={String.class}, tree="[0]")
    private Output<String> connectionRateLimit;

    /**
     * @return Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
     * 
     */
    public Output<String> connectionRateLimit() {
        return this.connectionRateLimit;
    }
    /**
     * Specifies the fixed ratio value used for a node during ratio load balancing.
     * 
     */
    @Export(name="dynamicRatio", refs={Integer.class}, tree="[0]")
    private Output<Integer> dynamicRatio;

    /**
     * @return Specifies the fixed ratio value used for a node during ratio load balancing.
     * 
     */
    public Output<Integer> dynamicRatio() {
        return this.dynamicRatio;
    }
    /**
     * Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
     * 
     */
    @Export(name="fqdnAutopopulate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> fqdnAutopopulate;

    /**
     * @return Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
     * 
     */
    public Output<Optional<String>> fqdnAutopopulate() {
        return Codegen.optional(this.fqdnAutopopulate);
    }
    /**
     * Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
     * 
     */
    @Export(name="monitor", refs={String.class}, tree="[0]")
    private Output<String> monitor;

    /**
     * @return Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
     * 
     */
    public Output<String> monitor() {
        return this.monitor;
    }
    /**
     * Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
     * 
     */
    @Export(name="node", refs={String.class}, tree="[0]")
    private Output<String> node;

    /**
     * @return Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
     * 
     */
    public Output<String> node() {
        return this.node;
    }
    /**
     * Name of the pool to which members should be attached,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
     * 
     */
    @Export(name="pool", refs={String.class}, tree="[0]")
    private Output<String> pool;

    /**
     * @return Name of the pool to which members should be attached,it should be &#34;full path&#34;.The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
     * 
     */
    public Output<String> pool() {
        return this.pool;
    }
    /**
     * Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
     * 
     */
    @Export(name="priorityGroup", refs={Integer.class}, tree="[0]")
    private Output<Integer> priorityGroup;

    /**
     * @return Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
     * 
     */
    public Output<Integer> priorityGroup() {
        return this.priorityGroup;
    }
    /**
     * &#34;Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.&#34;.
     * 
     */
    @Export(name="ratio", refs={Integer.class}, tree="[0]")
    private Output<Integer> ratio;

    /**
     * @return &#34;Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.&#34;.
     * 
     */
    public Output<Integer> ratio() {
        return this.ratio;
    }
    /**
     * Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PoolAttachment(java.lang.String name) {
        this(name, PoolAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PoolAttachment(java.lang.String name, PoolAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PoolAttachment(java.lang.String name, PoolAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/poolAttachment:PoolAttachment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PoolAttachment(java.lang.String name, Output<java.lang.String> id, @Nullable PoolAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/poolAttachment:PoolAttachment", name, state, makeResourceOptions(options, id), false);
    }

    private static PoolAttachmentArgs makeArgs(PoolAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PoolAttachmentArgs.Empty : args;
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
    public static PoolAttachment get(java.lang.String name, Output<java.lang.String> id, @Nullable PoolAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PoolAttachment(name, id, state, options);
    }
}
