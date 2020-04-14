// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Outputs
{

    [OutputType]
    public sealed class NodeFqdn
    {
        /// <summary>
        /// Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
        /// </summary>
        public readonly string? AddressFamily;
        public readonly string? Autopopulate;
        public readonly int? Downinterval;
        /// <summary>
        /// Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
        /// </summary>
        public readonly string? Interval;
        /// <summary>
        /// Name of the node
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private NodeFqdn(
            string? addressFamily,

            string? autopopulate,

            int? downinterval,

            string? interval,

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
