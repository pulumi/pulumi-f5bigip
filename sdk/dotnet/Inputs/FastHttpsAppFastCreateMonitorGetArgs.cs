// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Inputs
{

    public sealed class FastHttpsAppFastCreateMonitorGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
        /// </summary>
        [Input("monitorAuth")]
        public Input<bool>? MonitorAuth { get; set; }

        /// <summary>
        /// password for web access on FAST-Generated Pool Monitor.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The presence of this string anywhere in the HTTP response implies availability.
        /// </summary>
        [Input("response")]
        public Input<string>? Response { get; set; }

        /// <summary>
        /// Specify data to be sent during each health check for FAST-Generated Pool Monitor.
        /// </summary>
        [Input("sendString")]
        public Input<string>? SendString { get; set; }

        /// <summary>
        /// username for web access on FAST-Generated Pool Monitor.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public FastHttpsAppFastCreateMonitorGetArgs()
        {
        }
        public static new FastHttpsAppFastCreateMonitorGetArgs Empty => new FastHttpsAppFastCreateMonitorGetArgs();
    }
}
