// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.Node` Manages a node configuration
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// </summary>
    public partial class Node : Pulumi.CustomResource
    {
        /// <summary>
        /// IP or hostname of the node
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum number of connections allowed for the node or node address.
        /// </summary>
        [Output("connectionLimit")]
        public Output<int> ConnectionLimit { get; private set; } = null!;

        /// <summary>
        /// User-defined description give ltm_node
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the fixed ratio value used for a node during ratio load balancing.
        /// </summary>
        [Output("dynamicRatio")]
        public Output<int> DynamicRatio { get; private set; } = null!;

        [Output("fqdn")]
        public Output<Outputs.NodeFqdn?> Fqdn { get; private set; } = null!;

        /// <summary>
        /// specifies the name of the monitor or monitor rule that you want to associate with the node.
        /// </summary>
        [Output("monitor")]
        public Output<string?> Monitor { get; private set; } = null!;

        /// <summary>
        /// Name of the node
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
        /// 'disabled'.
        /// </summary>
        [Output("rateLimit")]
        public Output<string> RateLimit { get; private set; } = null!;

        /// <summary>
        /// Sets the ratio number for the node.
        /// </summary>
        [Output("ratio")]
        public Output<int> Ratio { get; private set; } = null!;

        /// <summary>
        /// Default is "user-up" you can set to "user-down" if you want to disable
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;


        /// <summary>
        /// Create a Node resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Node(string name, NodeArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/node:Node", name, args ?? new NodeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Node(string name, Input<string> id, NodeState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/node:Node", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Node resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Node Get(string name, Input<string> id, NodeState? state = null, CustomResourceOptions? options = null)
        {
            return new Node(name, id, state, options);
        }
    }

    public sealed class NodeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP or hostname of the node
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// Specifies the maximum number of connections allowed for the node or node address.
        /// </summary>
        [Input("connectionLimit")]
        public Input<int>? ConnectionLimit { get; set; }

        /// <summary>
        /// User-defined description give ltm_node
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the fixed ratio value used for a node during ratio load balancing.
        /// </summary>
        [Input("dynamicRatio")]
        public Input<int>? DynamicRatio { get; set; }

        [Input("fqdn")]
        public Input<Inputs.NodeFqdnArgs>? Fqdn { get; set; }

        /// <summary>
        /// specifies the name of the monitor or monitor rule that you want to associate with the node.
        /// </summary>
        [Input("monitor")]
        public Input<string>? Monitor { get; set; }

        /// <summary>
        /// Name of the node
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
        /// 'disabled'.
        /// </summary>
        [Input("rateLimit")]
        public Input<string>? RateLimit { get; set; }

        /// <summary>
        /// Sets the ratio number for the node.
        /// </summary>
        [Input("ratio")]
        public Input<int>? Ratio { get; set; }

        /// <summary>
        /// Default is "user-up" you can set to "user-down" if you want to disable
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public NodeArgs()
        {
        }
    }

    public sealed class NodeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP or hostname of the node
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections allowed for the node or node address.
        /// </summary>
        [Input("connectionLimit")]
        public Input<int>? ConnectionLimit { get; set; }

        /// <summary>
        /// User-defined description give ltm_node
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the fixed ratio value used for a node during ratio load balancing.
        /// </summary>
        [Input("dynamicRatio")]
        public Input<int>? DynamicRatio { get; set; }

        [Input("fqdn")]
        public Input<Inputs.NodeFqdnGetArgs>? Fqdn { get; set; }

        /// <summary>
        /// specifies the name of the monitor or monitor rule that you want to associate with the node.
        /// </summary>
        [Input("monitor")]
        public Input<string>? Monitor { get; set; }

        /// <summary>
        /// Name of the node
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections per second allowed for a node or node address. The default value is
        /// 'disabled'.
        /// </summary>
        [Input("rateLimit")]
        public Input<string>? RateLimit { get; set; }

        /// <summary>
        /// Sets the ratio number for the node.
        /// </summary>
        [Input("ratio")]
        public Input<int>? Ratio { get; set; }

        /// <summary>
        /// Default is "user-up" you can set to "user-down" if you want to disable
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public NodeState()
        {
        }
    }
}
