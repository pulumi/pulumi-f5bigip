// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    public static class GetIrule
    {
        /// <summary>
        /// Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
        ///  
        ///  
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = F5BigIP.Ltm.GetIrule.Invoke(new()
        ///     {
        ///         Name = "terraform_irule",
        ///         Partition = "Common",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipIrule"] = test.Apply(getIruleResult =&gt; getIruleResult.Irule),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetIruleResult> InvokeAsync(GetIruleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIruleResult>("f5bigip:ltm/getIrule:getIrule", args ?? new GetIruleArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
        ///  
        ///  
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = F5BigIP.Ltm.GetIrule.Invoke(new()
        ///     {
        ///         Name = "terraform_irule",
        ///         Partition = "Common",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipIrule"] = test.Apply(getIruleResult =&gt; getIruleResult.Irule),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetIruleResult> Invoke(GetIruleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIruleResult>("f5bigip:ltm/getIrule:getIrule", args ?? new GetIruleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ltm.IRule`) to get the ltm irule details available on BIG-IP
        ///  
        ///  
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = F5BigIP.Ltm.GetIrule.Invoke(new()
        ///     {
        ///         Name = "terraform_irule",
        ///         Partition = "Common",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipIrule"] = test.Apply(getIruleResult =&gt; getIruleResult.Irule),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetIruleResult> Invoke(GetIruleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIruleResult>("f5bigip:ltm/getIrule:getIrule", args ?? new GetIruleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIruleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Irule configured on bigip
        /// </summary>
        [Input("irule")]
        public string? Irule { get; set; }

        /// <summary>
        /// Name of the irule
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// partition of the ltm irule
        /// </summary>
        [Input("partition", required: true)]
        public string Partition { get; set; } = null!;

        public GetIruleArgs()
        {
        }
        public static new GetIruleArgs Empty => new GetIruleArgs();
    }

    public sealed class GetIruleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Irule configured on bigip
        /// </summary>
        [Input("irule")]
        public Input<string>? Irule { get; set; }

        /// <summary>
        /// Name of the irule
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// partition of the ltm irule
        /// </summary>
        [Input("partition", required: true)]
        public Input<string> Partition { get; set; } = null!;

        public GetIruleInvokeArgs()
        {
        }
        public static new GetIruleInvokeArgs Empty => new GetIruleInvokeArgs();
    }


    [OutputType]
    public sealed class GetIruleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Irule configured on bigip
        /// </summary>
        public readonly string? Irule;
        /// <summary>
        /// Name of irule configured on bigip with full path
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Bigip partition in which rule is configured
        /// </summary>
        public readonly string Partition;

        [OutputConstructor]
        private GetIruleResult(
            string id,

            string? irule,

            string name,

            string partition)
        {
            Id = id;
            Irule = irule;
            Name = name;
            Partition = partition;
        }
    }
}
