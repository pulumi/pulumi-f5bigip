// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Net
{
    /// <summary>
    /// `f5bigip.net.SelfIp` Manages a selfip configuration
    /// 
    /// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-selfip.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_net_selfip.html.markdown.
    /// </summary>
    public partial class SelfIp : Pulumi.CustomResource
    {
        /// <summary>
        /// The Self IP's address and netmask.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Name of the selfip
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        /// </summary>
        [Output("trafficGroup")]
        public Output<string?> TrafficGroup { get; private set; } = null!;

        /// <summary>
        /// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        /// </summary>
        [Output("vlan")]
        public Output<string> Vlan { get; private set; } = null!;


        /// <summary>
        /// Create a SelfIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SelfIp(string name, SelfIpArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:net/selfIp:SelfIp", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private SelfIp(string name, Input<string> id, SelfIpState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:net/selfIp:SelfIp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SelfIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SelfIp Get(string name, Input<string> id, SelfIpState? state = null, CustomResourceOptions? options = null)
        {
            return new SelfIp(name, id, state, options);
        }
    }

    public sealed class SelfIpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Self IP's address and netmask.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// Name of the selfip
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        /// <summary>
        /// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        /// </summary>
        [Input("vlan", required: true)]
        public Input<string> Vlan { get; set; } = null!;

        public SelfIpArgs()
        {
        }
    }

    public sealed class SelfIpState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Self IP's address and netmask.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Name of the selfip
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        /// <summary>
        /// Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        /// </summary>
        [Input("vlan")]
        public Input<string>? Vlan { get; set; }

        public SelfIpState()
        {
        }
    }
}
