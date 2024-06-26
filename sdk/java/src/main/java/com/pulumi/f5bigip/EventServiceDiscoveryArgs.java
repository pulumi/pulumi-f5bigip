// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.f5bigip.inputs.EventServiceDiscoveryNodeArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EventServiceDiscoveryArgs extends com.pulumi.resources.ResourceArgs {

    public static final EventServiceDiscoveryArgs Empty = new EventServiceDiscoveryArgs();

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
    @Import(name="nodes")
    private @Nullable Output<List<EventServiceDiscoveryNodeArgs>> nodes;

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
    public Optional<Output<List<EventServiceDiscoveryNodeArgs>>> nodes() {
        return Optional.ofNullable(this.nodes);
    }

    /**
     * servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     * 
     */
    @Import(name="taskid", required=true)
    private Output<String> taskid;

    /**
     * @return servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
     * 
     */
    public Output<String> taskid() {
        return this.taskid;
    }

    private EventServiceDiscoveryArgs() {}

    private EventServiceDiscoveryArgs(EventServiceDiscoveryArgs $) {
        this.nodes = $.nodes;
        this.taskid = $.taskid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EventServiceDiscoveryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EventServiceDiscoveryArgs $;

        public Builder() {
            $ = new EventServiceDiscoveryArgs();
        }

        public Builder(EventServiceDiscoveryArgs defaults) {
            $ = new EventServiceDiscoveryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param nodes Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
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
         * @return builder
         * 
         */
        public Builder nodes(@Nullable Output<List<EventServiceDiscoveryNodeArgs>> nodes) {
            $.nodes = nodes;
            return this;
        }

        /**
         * @param nodes Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
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
         * @return builder
         * 
         */
        public Builder nodes(List<EventServiceDiscoveryNodeArgs> nodes) {
            return nodes(Output.of(nodes));
        }

        /**
         * @param nodes Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
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
         * @return builder
         * 
         */
        public Builder nodes(EventServiceDiscoveryNodeArgs... nodes) {
            return nodes(List.of(nodes));
        }

        /**
         * @param taskid servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
         * 
         * @return builder
         * 
         */
        public Builder taskid(Output<String> taskid) {
            $.taskid = taskid;
            return this;
        }

        /**
         * @param taskid servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
         * 
         * @return builder
         * 
         */
        public Builder taskid(String taskid) {
            return taskid(Output.of(taskid));
        }

        public EventServiceDiscoveryArgs build() {
            if ($.taskid == null) {
                throw new MissingRequiredPropertyException("EventServiceDiscoveryArgs", "taskid");
            }
            return $;
        }
    }

}
