// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Outputs
{

    [OutputType]
    public sealed class EventServiceDiscoveryNode
    {
        /// <summary>
        /// name of node
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// ip of nonde
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// port
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private EventServiceDiscoveryNode(
            string? id,

            string? ip,

            int? port)
        {
            Id = id;
            Ip = ip;
            Port = port;
        }
    }
}
