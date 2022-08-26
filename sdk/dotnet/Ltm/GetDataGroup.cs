// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    public static class GetDataGroup
    {
        /// <summary>
        /// Use this data source (`f5bigip.ltm.DataGroup`) to get the data group details available on BIG-IP
        ///  
        ///  
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var dG_TC3 = F5BigIP.Ltm.GetDataGroup.Invoke(new()
        ///     {
        ///         Name = "test-dg",
        ///         Partition = "Common",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDataGroupResult> InvokeAsync(GetDataGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataGroupResult>("f5bigip:ltm/getDataGroup:getDataGroup", args ?? new GetDataGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ltm.DataGroup`) to get the data group details available on BIG-IP
        ///  
        ///  
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using F5BigIP = Pulumi.F5BigIP;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var dG_TC3 = F5BigIP.Ltm.GetDataGroup.Invoke(new()
        ///     {
        ///         Name = "test-dg",
        ///         Partition = "Common",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDataGroupResult> Invoke(GetDataGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataGroupResult>("f5bigip:ltm/getDataGroup:getDataGroup", args ?? new GetDataGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the datagroup
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// partition of the datagroup
        /// </summary>
        [Input("partition", required: true)]
        public string Partition { get; set; } = null!;

        [Input("records")]
        private List<Inputs.GetDataGroupRecordArgs>? _records;

        /// <summary>
        /// Specifies record of type (string/ip/integer)
        /// </summary>
        public List<Inputs.GetDataGroupRecordArgs> Records
        {
            get => _records ?? (_records = new List<Inputs.GetDataGroupRecordArgs>());
            set => _records = value;
        }

        /// <summary>
        /// The Data Group type (string, ip, integer)"
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public GetDataGroupArgs()
        {
        }
        public static new GetDataGroupArgs Empty => new GetDataGroupArgs();
    }

    public sealed class GetDataGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the datagroup
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// partition of the datagroup
        /// </summary>
        [Input("partition", required: true)]
        public Input<string> Partition { get; set; } = null!;

        [Input("records")]
        private InputList<Inputs.GetDataGroupRecordInputArgs>? _records;

        /// <summary>
        /// Specifies record of type (string/ip/integer)
        /// </summary>
        public InputList<Inputs.GetDataGroupRecordInputArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.GetDataGroupRecordInputArgs>());
            set => _records = value;
        }

        /// <summary>
        /// The Data Group type (string, ip, integer)"
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GetDataGroupInvokeArgs()
        {
        }
        public static new GetDataGroupInvokeArgs Empty => new GetDataGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataGroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Partition;
        /// <summary>
        /// Specifies record of type (string/ip/integer)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDataGroupRecordResult> Records;
        /// <summary>
        /// The Data Group type (string, ip, integer)"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDataGroupResult(
            string id,

            string name,

            string partition,

            ImmutableArray<Outputs.GetDataGroupRecordResult> records,

            string type)
        {
            Id = id;
            Name = name;
            Partition = partition;
            Records = records;
            Type = type;
        }
    }
}
