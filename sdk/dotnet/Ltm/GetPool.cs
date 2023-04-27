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
        /// <summary>
        /// Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
        ///  
        ///  
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pool_Example = F5BigIP.Ltm.GetPool.Invoke(new()
        ///     {
        ///         Name = "example-pool",
        ///         Partition = "Common",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPoolResult> InvokeAsync(GetPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPoolResult>("f5bigip:ltm/getPool:getPool", args ?? new GetPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ltm.Pool`) to get the ltm monitor details available on BIG-IP
        ///  
        ///  
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pool_Example = F5BigIP.Ltm.GetPool.Invoke(new()
        ///     {
        ///         Name = "example-pool",
        ///         Partition = "Common",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPoolResult> Invoke(GetPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPoolResult>("f5bigip:ltm/getPool:getPool", args ?? new GetPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPoolArgs : global::Pulumi.InvokeArgs
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

        public GetPoolArgs()
        {
        }
        public static new GetPoolArgs Empty => new GetPoolArgs();
    }

    public sealed class GetPoolInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetPoolInvokeArgs()
        {
        }
        public static new GetPoolInvokeArgs Empty => new GetPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetPoolResult
    {
        /// <summary>
        /// Full path to the pool.
        /// </summary>
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
