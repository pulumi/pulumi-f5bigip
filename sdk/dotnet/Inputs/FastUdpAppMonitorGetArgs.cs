// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Inputs
{

    public sealed class FastUdpAppMonitorGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The presence of this optional string is required in the response, if specified it confirms availability.
        /// </summary>
        [Input("expectedResponse")]
        public Input<string>? ExpectedResponse { get; set; }

        /// <summary>
        /// Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// Optional data to be sent during each health check.
        /// </summary>
        [Input("sendString")]
        public Input<string>? SendString { get; set; }

        public FastUdpAppMonitorGetArgs()
        {
        }
        public static new FastUdpAppMonitorGetArgs Empty => new FastUdpAppMonitorGetArgs();
    }
}
