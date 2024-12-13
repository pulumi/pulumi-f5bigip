// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl
{
    public static class GetWafPolicy
    {
        /// <summary>
        /// Use this data source (`f5bigip.WafPolicy`) to get the details of exist WAF policy BIG-IP.
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
        ///     var existpolicy = F5BigIP.Ssl.GetWafPolicy.Invoke(new()
        ///     {
        ///         PolicyId = "xxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetWafPolicyResult> InvokeAsync(GetWafPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWafPolicyResult>("f5bigip:ssl/getWafPolicy:getWafPolicy", args ?? new GetWafPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.WafPolicy`) to get the details of exist WAF policy BIG-IP.
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
        ///     var existpolicy = F5BigIP.Ssl.GetWafPolicy.Invoke(new()
        ///     {
        ///         PolicyId = "xxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetWafPolicyResult> Invoke(GetWafPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWafPolicyResult>("f5bigip:ssl/getWafPolicy:getWafPolicy", args ?? new GetWafPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.WafPolicy`) to get the details of exist WAF policy BIG-IP.
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
        ///     var existpolicy = F5BigIP.Ssl.GetWafPolicy.Invoke(new()
        ///     {
        ///         PolicyId = "xxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetWafPolicyResult> Invoke(GetWafPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWafPolicyResult>("f5bigip:ssl/getWafPolicy:getWafPolicy", args ?? new GetWafPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWafPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the WAF policy deployed in the BIG-IP.
        /// </summary>
        [Input("policyId", required: true)]
        public string PolicyId { get; set; } = null!;

        /// <summary>
        /// Exported WAF policy JSON
        /// </summary>
        [Input("policyJson")]
        public string? PolicyJson { get; set; }

        public GetWafPolicyArgs()
        {
        }
        public static new GetWafPolicyArgs Empty => new GetWafPolicyArgs();
    }

    public sealed class GetWafPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the WAF policy deployed in the BIG-IP.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// Exported WAF policy JSON
        /// </summary>
        [Input("policyJson")]
        public Input<string>? PolicyJson { get; set; }

        public GetWafPolicyInvokeArgs()
        {
        }
        public static new GetWafPolicyInvokeArgs Empty => new GetWafPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetWafPolicyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string PolicyId;
        /// <summary>
        /// Exported WAF policy JSON
        /// </summary>
        public readonly string PolicyJson;

        [OutputConstructor]
        private GetWafPolicyResult(
            string id,

            string policyId,

            string policyJson)
        {
            Id = id;
            PolicyId = policyId;
            PolicyJson = policyJson;
        }
    }
}
