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
    public sealed class FastTcpAppMonitor
    {
        /// <summary>
        /// Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        /// </summary>
        public readonly int? Interval;

        [OutputConstructor]
        private FastTcpAppMonitor(int? interval)
        {
            Interval = interval;
        }
    }
}
