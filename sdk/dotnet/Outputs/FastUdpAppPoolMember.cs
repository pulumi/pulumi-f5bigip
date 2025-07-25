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
    public sealed class FastUdpAppPoolMember
    {
        /// <summary>
        /// List of server address to be used for FAST-Generated Pool.
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        /// <summary>
        /// connectionLimit value to be used for FAST-Generated Pool.
        /// </summary>
        public readonly int? ConnectionLimit;
        /// <summary>
        /// port number of serviceport to be used for FAST-Generated Pool.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// priorityGroup value to be used for FAST-Generated Pool.
        /// </summary>
        public readonly int? PriorityGroup;
        /// <summary>
        /// shareNodes value to be used for FAST-Generated Pool.
        /// </summary>
        public readonly bool? ShareNodes;

        [OutputConstructor]
        private FastUdpAppPoolMember(
            ImmutableArray<string> addresses,

            int? connectionLimit,

            int? port,

            int? priorityGroup,

            bool? shareNodes)
        {
            Addresses = addresses;
            ConnectionLimit = connectionLimit;
            Port = port;
            PriorityGroup = priorityGroup;
            ShareNodes = shareNodes;
        }
    }
}
