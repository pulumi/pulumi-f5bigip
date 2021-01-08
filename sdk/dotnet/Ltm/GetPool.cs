// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    public static class GetPool
    {
        public static Task<GetPoolResult> InvokeAsync(GetPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPoolResult>("f5bigip:ltm/getPool:getPool", args ?? new GetPoolArgs(), options.WithVersion());
    }


    public sealed class GetPoolArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("partition", required: true)]
        public string Partition { get; set; } = null!;

        public GetPoolArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPoolResult
    {
        public readonly string FullPath;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Partition;

        [OutputConstructor]
        private GetPoolResult(
            string fullPath,

            string id,

            string name,

            string partition)
        {
            FullPath = fullPath;
            Id = id;
            Name = name;
            Partition = partition;
        }
    }
}