// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Outputs
{

    [OutputType]
    public sealed class GetNodeFqdnResult
    {
        /// <summary>
        /// The FQDN node's address family.
        /// </summary>
        public readonly string? AddressFamily;
        /// <summary>
        /// Specifies if the node should scale to the IP address set returned by DNS.
        /// </summary>
        public readonly string Autopopulate;
        /// <summary>
        /// The number of attempts to resolve a domain name.
        /// </summary>
        public readonly int Downinterval;
        /// <summary>
        /// The amount of time before sending the next DNS query.
        /// </summary>
        public readonly string Interval;
        /// <summary>
        /// Name of the node.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private GetNodeFqdnResult(
            string? addressFamily,

            string autopopulate,

            int downinterval,

            string interval,

            string? name)
        {
            AddressFamily = addressFamily;
            Autopopulate = autopopulate;
            Downinterval = downinterval;
            Interval = interval;
            Name = name;
        }
    }
}
