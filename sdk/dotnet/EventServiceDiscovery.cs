// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new F5BigIP.EventServiceDiscovery("test", new()
    ///     {
    ///         Taskid = "~Sample_event_sd~My_app~My_pool",
    ///         Nodes = new[]
    ///         {
    ///             new F5BigIP.Inputs.EventServiceDiscoveryNodeArgs
    ///             {
    ///                 Id = "newNode1",
    ///                 Ip = "192.168.2.3",
    ///                 Port = 8080,
    ///             },
    ///             new F5BigIP.Inputs.EventServiceDiscoveryNodeArgs
    ///             {
    ///                 Id = "newNode2",
    ///                 Ip = "192.168.2.4",
    ///                 Port = 8080,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/eventServiceDiscovery:EventServiceDiscovery")]
    public partial class EventServiceDiscovery : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
        /// 
        /// For more information, please refer below document
        /// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
        /// 
        /// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
        /// 
        /// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
        /// 
        /// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://&lt;host&gt;/mgmt/shared/service-discovery/task/~&lt;tenant name&gt;~&lt;application name&gt;~&lt;pool name&gt;/nodes.
        /// 
        /// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
        /// 
        /// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
        /// First we show the initial declaration to POST to the BIG-IP system.
        /// 
        /// {
        /// "class": "ADC",
        /// "schemaVersion": "3.9.0",
        /// "id": "Pool",
        /// "Sample_event_sd": {
        /// "class": "Tenant",
        /// "My_app": {
        /// "class": "Application",
        /// "My_pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 8080,
        /// "addressDiscovery": "static",
        /// "serverAddresses": [
        /// "192.0.2.2"
        /// ]
        /// },
        /// {
        /// "servicePort": 8080,
        /// "addressDiscovery": "event"
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// 
        /// 
        /// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<Outputs.EventServiceDiscoveryNode>> Nodes { get; private set; } = null!;

        /// <summary>
        /// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        /// </summary>
        [Output("taskid")]
        public Output<string> Taskid { get; private set; } = null!;


        /// <summary>
        /// Create a EventServiceDiscovery resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventServiceDiscovery(string name, EventServiceDiscoveryArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/eventServiceDiscovery:EventServiceDiscovery", name, args ?? new EventServiceDiscoveryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventServiceDiscovery(string name, Input<string> id, EventServiceDiscoveryState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/eventServiceDiscovery:EventServiceDiscovery", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventServiceDiscovery resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventServiceDiscovery Get(string name, Input<string> id, EventServiceDiscoveryState? state = null, CustomResourceOptions? options = null)
        {
            return new EventServiceDiscovery(name, id, state, options);
        }
    }

    public sealed class EventServiceDiscoveryArgs : global::Pulumi.ResourceArgs
    {
        [Input("nodes")]
        private InputList<Inputs.EventServiceDiscoveryNodeArgs>? _nodes;

        /// <summary>
        /// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
        /// 
        /// For more information, please refer below document
        /// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
        /// 
        /// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
        /// 
        /// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
        /// 
        /// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://&lt;host&gt;/mgmt/shared/service-discovery/task/~&lt;tenant name&gt;~&lt;application name&gt;~&lt;pool name&gt;/nodes.
        /// 
        /// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
        /// 
        /// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
        /// First we show the initial declaration to POST to the BIG-IP system.
        /// 
        /// {
        /// "class": "ADC",
        /// "schemaVersion": "3.9.0",
        /// "id": "Pool",
        /// "Sample_event_sd": {
        /// "class": "Tenant",
        /// "My_app": {
        /// "class": "Application",
        /// "My_pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 8080,
        /// "addressDiscovery": "static",
        /// "serverAddresses": [
        /// "192.0.2.2"
        /// ]
        /// },
        /// {
        /// "servicePort": 8080,
        /// "addressDiscovery": "event"
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// 
        /// 
        /// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        /// </summary>
        public InputList<Inputs.EventServiceDiscoveryNodeArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.EventServiceDiscoveryNodeArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        /// </summary>
        [Input("taskid", required: true)]
        public Input<string> Taskid { get; set; } = null!;

        public EventServiceDiscoveryArgs()
        {
        }
        public static new EventServiceDiscoveryArgs Empty => new EventServiceDiscoveryArgs();
    }

    public sealed class EventServiceDiscoveryState : global::Pulumi.ResourceArgs
    {
        [Input("nodes")]
        private InputList<Inputs.EventServiceDiscoveryNodeGetArgs>? _nodes;

        /// <summary>
        /// Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
        /// 
        /// For more information, please refer below document
        /// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
        /// 
        /// Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
        /// 
        /// With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
        /// 
        /// When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://&lt;host&gt;/mgmt/shared/service-discovery/task/~&lt;tenant name&gt;~&lt;application name&gt;~&lt;pool name&gt;/nodes.
        /// 
        /// For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
        /// 
        /// Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
        /// First we show the initial declaration to POST to the BIG-IP system.
        /// 
        /// {
        /// "class": "ADC",
        /// "schemaVersion": "3.9.0",
        /// "id": "Pool",
        /// "Sample_event_sd": {
        /// "class": "Tenant",
        /// "My_app": {
        /// "class": "Application",
        /// "My_pool": {
        /// "class": "Pool",
        /// "members": [
        /// {
        /// "servicePort": 8080,
        /// "addressDiscovery": "static",
        /// "serverAddresses": [
        /// "192.0.2.2"
        /// ]
        /// },
        /// {
        /// "servicePort": 8080,
        /// "addressDiscovery": "event"
        /// }
        /// ]
        /// }
        /// }
        /// }
        /// }
        /// 
        /// 
        /// Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        /// </summary>
        public InputList<Inputs.EventServiceDiscoveryNodeGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.EventServiceDiscoveryNodeGetArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        /// </summary>
        [Input("taskid")]
        public Input<string>? Taskid { get; set; }

        public EventServiceDiscoveryState()
        {
        }
        public static new EventServiceDiscoveryState Empty => new EventServiceDiscoveryState();
    }
}
