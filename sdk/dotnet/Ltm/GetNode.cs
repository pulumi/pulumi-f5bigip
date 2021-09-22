// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    public static class GetNode
    {
        /// <summary>
        /// Use this data source (`f5bigip.ltm.Node`) to get the ltm node details available on BIG-IP
        /// </summary>
        public static Task<GetNodeResult> InvokeAsync(GetNodeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNodeResult>("f5bigip:ltm/getNode:getNode", args ?? new GetNodeArgs(), options.WithVersion());
    }


    public sealed class GetNodeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The address of the node.
        /// </summary>
        [Input("address")]
        public string? Address { get; set; }

        /// <summary>
        /// User defined description of the node.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        [Input("fqdn")]
        public Inputs.GetNodeFqdnArgs? Fqdn { get; set; }

        /// <summary>
        /// Full path of the node (partition and name)
        /// </summary>
        [Input("fullPath")]
        public string? FullPath { get; set; }

        /// <summary>
        /// Name of the node.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// partition of the node.
        /// </summary>
        [Input("partition", required: true)]
        public string Partition { get; set; } = null!;

        public GetNodeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNodeResult
    {
        /// <summary>
        /// The address of the node.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// Node connection limit.
        /// </summary>
        public readonly int ConnectionLimit;
        /// <summary>
        /// User defined description of the node.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The dynamic ratio number for the node.
        /// </summary>
        public readonly int DynamicRatio;
        public readonly Outputs.GetNodeFqdnResult Fqdn;
        /// <summary>
        /// Full path of the node (partition and name)
        /// </summary>
        public readonly string? FullPath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies the health monitors the system currently uses to monitor this node.
        /// </summary>
        public readonly string Monitor;
        public readonly string Name;
        public readonly string Partition;
        /// <summary>
        /// Node rate limit.
        /// </summary>
        public readonly string RateLimit;
        /// <summary>
        /// Node ratio weight.
        /// </summary>
        public readonly int Ratio;
        public readonly string Session;
        /// <summary>
        /// The current state of the node.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetNodeResult(
            string? address,

            int connectionLimit,

            string? description,

            int dynamicRatio,

            Outputs.GetNodeFqdnResult fqdn,

            string? fullPath,

            string id,

            string monitor,

            string name,

            string partition,

            string rateLimit,

            int ratio,

            string session,

            string state)
        {
            Address = address;
            ConnectionLimit = connectionLimit;
            Description = description;
            DynamicRatio = dynamicRatio;
            Fqdn = fqdn;
            FullPath = fullPath;
            Id = id;
            Monitor = monitor;
            Name = name;
            Partition = partition;
            RateLimit = rateLimit;
            Ratio = ratio;
            Session = session;
            State = state;
        }
    }
}
