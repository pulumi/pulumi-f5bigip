// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// `f5bigip.TrafficSelector` Manage IPSec Traffic Selectors on BIG-IP
    /// 
    /// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-selector)
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test_selector = new F5BigIP.TrafficSelector("test-selector", new()
    ///     {
    ///         DestinationAddress = "3.10.11.2/32",
    ///         Name = "/Common/test-selector",
    ///         SourceAddress = "2.10.11.12/32",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/trafficSelector:TrafficSelector")]
    public partial class TrafficSelector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the traffic selector.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        /// </summary>
        [Output("destinationAddress")]
        public Output<string> DestinationAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the IP port used by the application. The default value is `All Ports (0)`
        /// </summary>
        [Output("destinationPort")]
        public Output<int> DestinationPort { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        /// </summary>
        [Output("ipProtocol")]
        public Output<int> IpProtocol { get; private set; } = null!;

        /// <summary>
        /// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        /// </summary>
        [Output("ipsecPolicy")]
        public Output<string> IpsecPolicy { get; private set; } = null!;

        /// <summary>
        /// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
        /// When creating a new traffic selector, if this parameter is not specified, the default is `last`
        /// </summary>
        [Output("order")]
        public Output<int> Order { get; private set; } = null!;

        /// <summary>
        /// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        /// </summary>
        [Output("sourceAddress")]
        public Output<string> SourceAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the IP port used by the application. The default value is `All Ports (0)`.
        /// </summary>
        [Output("sourcePort")]
        public Output<int> SourcePort { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficSelector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficSelector(string name, TrafficSelectorArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/trafficSelector:TrafficSelector", name, args ?? new TrafficSelectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficSelector(string name, Input<string> id, TrafficSelectorState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/trafficSelector:TrafficSelector", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TrafficSelector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficSelector Get(string name, Input<string> id, TrafficSelectorState? state = null, CustomResourceOptions? options = null)
        {
            return new TrafficSelector(name, id, state, options);
        }
    }

    public sealed class TrafficSelectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the traffic selector.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        /// </summary>
        [Input("destinationAddress", required: true)]
        public Input<string> DestinationAddress { get; set; } = null!;

        /// <summary>
        /// Specifies the IP port used by the application. The default value is `All Ports (0)`
        /// </summary>
        [Input("destinationPort")]
        public Input<int>? DestinationPort { get; set; }

        /// <summary>
        /// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        /// </summary>
        [Input("ipProtocol")]
        public Input<int>? IpProtocol { get; set; }

        /// <summary>
        /// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        /// </summary>
        [Input("ipsecPolicy")]
        public Input<string>? IpsecPolicy { get; set; }

        /// <summary>
        /// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
        /// When creating a new traffic selector, if this parameter is not specified, the default is `last`
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        /// </summary>
        [Input("sourceAddress", required: true)]
        public Input<string> SourceAddress { get; set; } = null!;

        /// <summary>
        /// Specifies the IP port used by the application. The default value is `All Ports (0)`.
        /// </summary>
        [Input("sourcePort")]
        public Input<int>? SourcePort { get; set; }

        public TrafficSelectorArgs()
        {
        }
        public static new TrafficSelectorArgs Empty => new TrafficSelectorArgs();
    }

    public sealed class TrafficSelectorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the traffic selector.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        /// </summary>
        [Input("destinationAddress")]
        public Input<string>? DestinationAddress { get; set; }

        /// <summary>
        /// Specifies the IP port used by the application. The default value is `All Ports (0)`
        /// </summary>
        [Input("destinationPort")]
        public Input<int>? DestinationPort { get; set; }

        /// <summary>
        /// Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        /// </summary>
        [Input("ipProtocol")]
        public Input<int>? IpProtocol { get; set; }

        /// <summary>
        /// Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        /// </summary>
        [Input("ipsecPolicy")]
        public Input<string>? IpsecPolicy { get; set; }

        /// <summary>
        /// Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
        /// When creating a new traffic selector, if this parameter is not specified, the default is `last`
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        /// </summary>
        [Input("sourceAddress")]
        public Input<string>? SourceAddress { get; set; }

        /// <summary>
        /// Specifies the IP port used by the application. The default value is `All Ports (0)`.
        /// </summary>
        [Input("sourcePort")]
        public Input<int>? SourcePort { get; set; }

        public TrafficSelectorState()
        {
        }
        public static new TrafficSelectorState Empty => new TrafficSelectorState();
    }
}
