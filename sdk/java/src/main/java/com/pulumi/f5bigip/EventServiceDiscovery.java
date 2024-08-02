// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.EventServiceDiscoveryArgs;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.inputs.EventServiceDiscoveryState;
import com.pulumi.f5bigip.outputs.EventServiceDiscoveryNode;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.f5bigip.EventServiceDiscovery;
 * import com.pulumi.f5bigip.EventServiceDiscoveryArgs;
 * import com.pulumi.f5bigip.inputs.EventServiceDiscoveryNodeArgs;
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
 *         var test = new EventServiceDiscovery("test", EventServiceDiscoveryArgs.builder()
 *             .taskid("~Sample_event_sd~My_app~My_pool")
 *             .nodes(            
 *                 EventServiceDiscoveryNodeArgs.builder()
 *                     .id("newNode1")
 *                     .ip("192.168.2.3")
 *                     .port(8080)
 *                     .build(),
 *                 EventServiceDiscoveryNodeArgs.builder()
 *                     .id("newNode2")
 *                     .ip("192.168.2.4")
 *                     .port(8080)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="f5bigip:index/eventServiceDiscovery:EventServiceDiscovery")
public class EventServiceDiscovery extends com.pulumi.resources.CustomResource {
    /**
     * Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     * 
     * For more information, please refer below document
     * https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
     * 
     * Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
     * 
     * With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
     * 
     * When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://&lt;host&gt;/mgmt/shared/service-discovery/task/~&lt;tenant name&gt;~&lt;application name&gt;~&lt;pool name&gt;/nodes.
     * 
     * For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
     * 
     * Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
     * First we show the initial declaration to POST to the BIG-IP system.
     * 
     * {
     * &#34;class&#34;: &#34;ADC&#34;,
     * &#34;schemaVersion&#34;: &#34;3.9.0&#34;,
     * &#34;id&#34;: &#34;Pool&#34;,
     * &#34;Sample_event_sd&#34;: {
     * &#34;class&#34;: &#34;Tenant&#34;,
     * &#34;My_app&#34;: {
     * &#34;class&#34;: &#34;Application&#34;,
     * &#34;My_pool&#34;: {
     * &#34;class&#34;: &#34;Pool&#34;,
     * &#34;members&#34;: [
     * {
     * &#34;servicePort&#34;: 8080,
     * &#34;addressDiscovery&#34;: &#34;static&#34;,
     * &#34;serverAddresses&#34;: [
     * &#34;192.0.2.2&#34;
     * ]
     * },
     * {
     * &#34;servicePort&#34;: 8080,
     * &#34;addressDiscovery&#34;: &#34;event&#34;
     * }
     * ]
     * }
     * }
     * }
     * }
     * 
     * Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool&#34; ) and node list for the resource to dynamically update the node list.
     * 
     */
    @Export(name="nodes", refs={List.class,EventServiceDiscoveryNode.class}, tree="[0,1]")
    private Output</* @Nullable */ List<EventServiceDiscoveryNode>> nodes;

    /**
     * @return Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
     * 
     * For more information, please refer below document
     * https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
     * 
     * Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
     * 
     * With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
     * 
     * When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://&lt;host&gt;/mgmt/shared/service-discovery/task/~&lt;tenant name&gt;~&lt;application name&gt;~&lt;pool name&gt;/nodes.
     * 
     * For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
     * 
     * Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
     * First we show the initial declaration to POST to the BIG-IP system.
     * 
     * {
     * &#34;class&#34;: &#34;ADC&#34;,
     * &#34;schemaVersion&#34;: &#34;3.9.0&#34;,
     * &#34;id&#34;: &#34;Pool&#34;,
     * &#34;Sample_event_sd&#34;: {
     * &#34;class&#34;: &#34;Tenant&#34;,
     * &#34;My_app&#34;: {
     * &#34;class&#34;: &#34;Application&#34;,
     * &#34;My_pool&#34;: {
     * &#34;class&#34;: &#34;Pool&#34;,
     * &#34;members&#34;: [
     * {
     * &#34;servicePort&#34;: 8080,
     * &#34;addressDiscovery&#34;: &#34;static&#34;,
     * &#34;serverAddresses&#34;: [
     * &#34;192.0.2.2&#34;
     * ]
     * },
     * {
     * &#34;servicePort&#34;: 8080,
     * &#34;addressDiscovery&#34;: &#34;event&#34;
     * }
     * ]
     * }
     * }
     * }
     * }
     * 
     * Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool&#34; ) and node list for the resource to dynamically update the node list.
     * 
     */
    public Output<Optional<List<EventServiceDiscoveryNode>>> nodes() {
        return Codegen.optional(this.nodes);
    }
    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     * 
     */
    @Export(name="taskid", refs={String.class}, tree="[0]")
    private Output<String> taskid;

    /**
     * @return servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     * 
     */
    public Output<String> taskid() {
        return this.taskid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EventServiceDiscovery(String name) {
        this(name, EventServiceDiscoveryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EventServiceDiscovery(String name, EventServiceDiscoveryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EventServiceDiscovery(String name, EventServiceDiscoveryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/eventServiceDiscovery:EventServiceDiscovery", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private EventServiceDiscovery(String name, Output<String> id, @Nullable EventServiceDiscoveryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/eventServiceDiscovery:EventServiceDiscovery", name, state, makeResourceOptions(options, id));
    }

    private static EventServiceDiscoveryArgs makeArgs(EventServiceDiscoveryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EventServiceDiscoveryArgs.Empty : args;
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
    public static EventServiceDiscovery get(String name, Output<String> id, @Nullable EventServiceDiscoveryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EventServiceDiscovery(name, id, state, options);
    }
}
