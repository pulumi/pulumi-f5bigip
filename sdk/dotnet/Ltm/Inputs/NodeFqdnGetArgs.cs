// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class NodeFqdnGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
        /// </summary>
        [Input("addressFamily")]
        public Input<string>? AddressFamily { get; set; }

        [Input("autopopulate")]
        public Input<string>? Autopopulate { get; set; }

        [Input("downinterval")]
        public Input<int>? Downinterval { get; set; }

        /// <summary>
        /// Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
        /// </summary>
        [Input("interval")]
        public Input<string>? Interval { get; set; }

        /// <summary>
        /// Name of the node
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public NodeFqdnGetArgs()
        {
        }
    }
}
