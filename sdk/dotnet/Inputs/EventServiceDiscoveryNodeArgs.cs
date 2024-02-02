// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Inputs
{

    public sealed class EventServiceDiscoveryNodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of node
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// ip of nonde
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// port
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public EventServiceDiscoveryNodeArgs()
        {
        }
        public static new EventServiceDiscoveryNodeArgs Empty => new EventServiceDiscoveryNodeArgs();
    }
}
