// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.NodeArgs;
import com.pulumi.f5bigip.ltm.inputs.NodeState;
import com.pulumi.f5bigip.ltm.outputs.NodeFqdn;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.Node` Manages a node configuration
 * 
 * For resources should be named with their `full path`.The full path is the combination of the `partition + name` of the resource( example: `/Common/my-node` ) or `partition + Direcroty + name` of the resource ( example: `/Common/test/my-node` ).
 * When including directory in `full path` we have to make sure it is created in the given partition before using it.
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
 * import com.pulumi.f5bigip.ltm.Node;
 * import com.pulumi.f5bigip.ltm.NodeArgs;
 * import com.pulumi.f5bigip.ltm.inputs.NodeFqdnArgs;
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
 *         var node = new Node(&#34;node&#34;, NodeArgs.builder()        
 *             .name(&#34;/Common/terraform_node1&#34;)
 *             .address(&#34;192.168.30.1&#34;)
 *             .connectionLimit(&#34;0&#34;)
 *             .dynamicRatio(&#34;1&#34;)
 *             .monitor(&#34;/Common/icmp&#34;)
 *             .description(&#34;Test-Node&#34;)
 *             .rateLimit(&#34;disabled&#34;)
 *             .fqdn(NodeFqdnArgs.builder()
 *                 .addressFamily(&#34;ipv4&#34;)
 *                 .interval(&#34;3000&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;      
 * 
 * ## Importing
 * 
 * An existing Node can be imported into this resource by supplying Node Name in `full path` as `id`.
 * An example is below:
 * ```sh
 * $ terraform import bigip_ltm_node.site2_node &#34;/TEST/testnode&#34;
 *             (or)
 * $ terraform import bigip_ltm_node.site2_node &#34;/Common/3.3.3.3&#34;
 * 
 * ```
 * 
 */
@ResourceType(type="f5bigip:ltm/node:Node")
public class Node extends com.pulumi.resources.CustomResource {
    /**
     * IP or hostname of the node
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return IP or hostname of the node
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * Specifies the maximum number of connections allowed for the node or node address.
     * 
     */
    @Export(name="connectionLimit", refs={Integer.class}, tree="[0]")
    private Output<Integer> connectionLimit;

    /**
     * @return Specifies the maximum number of connections allowed for the node or node address.
     * 
     */
    public Output<Integer> connectionLimit() {
        return this.connectionLimit;
    }
    /**
     * User-defined description give ltm_node
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User-defined description give ltm_node
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
    @Export(name="fqdn", refs={NodeFqdn.class}, tree="[0]")
    private Output</* @Nullable */ NodeFqdn> fqdn;

    public Output<Optional<NodeFqdn>> fqdn() {
        return Codegen.optional(this.fqdn);
    }
    /**
     * specifies the name of the monitor or monitor rule that you want to associate with the node.
     * 
     */
    @Export(name="monitor", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> monitor;

    /**
     * @return specifies the name of the monitor or monitor rule that you want to associate with the node.
     * 
     */
    public Output<Optional<String>> monitor() {
        return Codegen.optional(this.monitor);
    }
    /**
     * Name of the node
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the node
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Specifies the maximum number of connections per second allowed for a node or node address. The default value is &#39;disabled&#39;.
     * 
     */
    @Export(name="rateLimit", refs={String.class}, tree="[0]")
    private Output<String> rateLimit;

    /**
     * @return Specifies the maximum number of connections per second allowed for a node or node address. The default value is &#39;disabled&#39;.
     * 
     */
    public Output<String> rateLimit() {
        return this.rateLimit;
    }
    /**
     * Sets the ratio number for the node.
     * 
     */
    @Export(name="ratio", refs={Integer.class}, tree="[0]")
    private Output<Integer> ratio;

    /**
     * @return Sets the ratio number for the node.
     * 
     */
    public Output<Integer> ratio() {
        return this.ratio;
    }
    /**
     * Enables or disables the node for new sessions. The default value is user-enabled.
     * 
     */
    @Export(name="session", refs={String.class}, tree="[0]")
    private Output<String> session;

    /**
     * @return Enables or disables the node for new sessions. The default value is user-enabled.
     * 
     */
    public Output<String> session() {
        return this.session;
    }
    /**
     * Default is &#34;user-up&#34; you can set to &#34;user-down&#34; if you want to disable
     * 
     * &gt; *NOTE* Below attributes needs to be configured under fqdn option.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Default is &#34;user-up&#34; you can set to &#34;user-down&#34; if you want to disable
     * 
     * &gt; *NOTE* Below attributes needs to be configured under fqdn option.
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Node(String name) {
        this(name, NodeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Node(String name, NodeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Node(String name, NodeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/node:Node", name, args == null ? NodeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Node(String name, Output<String> id, @Nullable NodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/node:Node", name, state, makeResourceOptions(options, id));
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
    public static Node get(String name, Output<String> id, @Nullable NodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Node(name, id, state, options);
    }
}
