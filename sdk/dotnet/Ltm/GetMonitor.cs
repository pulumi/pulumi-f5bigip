// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    public static class GetMonitor
    {
        /// <summary>
        /// Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
        ///  
        ///  
        /// </summary>
        public static Task<GetMonitorResult> InvokeAsync(GetMonitorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMonitorResult>("f5bigip:ltm/getMonitor:getMonitor", args ?? new GetMonitorArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ltm.Monitor`) to get the ltm monitor details available on BIG-IP
        ///  
        ///  
        /// </summary>
        public static Output<GetMonitorResult> Invoke(GetMonitorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMonitorResult>("f5bigip:ltm/getMonitor:getMonitor", args ?? new GetMonitorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMonitorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ltm monitor
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// partition of the ltm monitor
        /// </summary>
        [Input("partition", required: true)]
        public string Partition { get; set; } = null!;

        public GetMonitorArgs()
        {
        }
        public static new GetMonitorArgs Empty => new GetMonitorArgs();
    }

    public sealed class GetMonitorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the ltm monitor
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// partition of the ltm monitor
        /// </summary>
        [Input("partition", required: true)]
        public Input<string> Partition { get; set; } = null!;

        public GetMonitorInvokeArgs()
        {
        }
        public static new GetMonitorInvokeArgs Empty => new GetMonitorInvokeArgs();
    }


    [OutputType]
    public sealed class GetMonitorResult
    {
        /// <summary>
        /// Displays whether adaptive response time monitoring is enabled for this monitor.
        /// </summary>
        public readonly string Adaptive;
        /// <summary>
        /// Displays whether adaptive response time monitoring is enabled for this monitor.
        /// </summary>
        public readonly int AdaptiveLimit;
        public readonly string Database;
        public readonly string DefaultsFrom;
        /// <summary>
        /// id will be full path name of ltm monitor.
        /// </summary>
        public readonly string Destination;
        public readonly string Filename;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown.
        /// </summary>
        public readonly int Interval;
        /// <summary>
        /// Displays the differentiated services code point (DSCP). DSCP is a 6-bit value in the Differentiated Services (DS) field of the IP header.
        /// </summary>
        public readonly int IpDscp;
        /// <summary>
        /// Displays whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
        /// </summary>
        public readonly string ManualResume;
        public readonly string Mode;
        public readonly string Name;
        public readonly string Partition;
        public readonly string ReceiveDisable;
        /// <summary>
        /// Instructs the system to mark the target resource down when the test is successful.
        /// </summary>
        public readonly string Reverse;
        public readonly int TimeUntilUp;
        public readonly int Timeout;
        /// <summary>
        /// Displays whether the monitor operates in transparent mode.
        /// </summary>
        public readonly string Transparent;
        public readonly string Username;

        [OutputConstructor]
        private GetMonitorResult(
            string adaptive,

            int adaptiveLimit,

            string database,

            string defaultsFrom,

            string destination,

            string filename,

            string id,

            int interval,

            int ipDscp,

            string manualResume,

            string mode,

            string name,

            string partition,

            string receiveDisable,

            string reverse,

            int timeUntilUp,

            int timeout,

            string transparent,

            string username)
        {
            Adaptive = adaptive;
            AdaptiveLimit = adaptiveLimit;
            Database = database;
            DefaultsFrom = defaultsFrom;
            Destination = destination;
            Filename = filename;
            Id = id;
            Interval = interval;
            IpDscp = ipDscp;
            ManualResume = manualResume;
            Mode = mode;
            Name = name;
            Partition = partition;
            ReceiveDisable = receiveDisable;
            Reverse = reverse;
            TimeUntilUp = timeUntilUp;
            Timeout = timeout;
            Transparent = transparent;
            Username = username;
        }
    }
}
