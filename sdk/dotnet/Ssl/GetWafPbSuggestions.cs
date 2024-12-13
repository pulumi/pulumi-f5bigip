// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl
{
    public static class GetWafPbSuggestions
    {
        /// <summary>
        /// Use this data source (`f5bigip.ssl.getWafPbSuggestions`) to export PB suggestions from an existing WAF policy.
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
        ///     var PBWAF1 = F5BigIP.Ssl.GetWafPbSuggestions.Invoke(new()
        ///     {
        ///         PolicyName = "protect_me_policy",
        ///         Partition = "Common",
        ///         MinimumLearningScore = 20,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetWafPbSuggestionsResult> InvokeAsync(GetWafPbSuggestionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWafPbSuggestionsResult>("f5bigip:ssl/getWafPbSuggestions:getWafPbSuggestions", args ?? new GetWafPbSuggestionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ssl.getWafPbSuggestions`) to export PB suggestions from an existing WAF policy.
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
        ///     var PBWAF1 = F5BigIP.Ssl.GetWafPbSuggestions.Invoke(new()
        ///     {
        ///         PolicyName = "protect_me_policy",
        ///         Partition = "Common",
        ///         MinimumLearningScore = 20,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetWafPbSuggestionsResult> Invoke(GetWafPbSuggestionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWafPbSuggestionsResult>("f5bigip:ssl/getWafPbSuggestions:getWafPbSuggestions", args ?? new GetWafPbSuggestionsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ssl.getWafPbSuggestions`) to export PB suggestions from an existing WAF policy.
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
        ///     var PBWAF1 = F5BigIP.Ssl.GetWafPbSuggestions.Invoke(new()
        ///     {
        ///         PolicyName = "protect_me_policy",
        ///         Partition = "Common",
        ///         MinimumLearningScore = 20,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetWafPbSuggestionsResult> Invoke(GetWafPbSuggestionsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetWafPbSuggestionsResult>("f5bigip:ssl/getWafPbSuggestions:getWafPbSuggestions", args ?? new GetWafPbSuggestionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWafPbSuggestionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The minimum learning score for suggestions.
        /// </summary>
        [Input("minimumLearningScore", required: true)]
        public int MinimumLearningScore { get; set; }

        /// <summary>
        /// Partition on which WAF policy is located.
        /// </summary>
        [Input("partition", required: true)]
        public string Partition { get; set; } = null!;

        /// <summary>
        /// System generated id of the WAF policy
        /// </summary>
        [Input("policyId")]
        public string? PolicyId { get; set; }

        /// <summary>
        /// WAF policy name from which PB suggestions should be exported.
        /// </summary>
        [Input("policyName", required: true)]
        public string PolicyName { get; set; } = null!;

        public GetWafPbSuggestionsArgs()
        {
        }
        public static new GetWafPbSuggestionsArgs Empty => new GetWafPbSuggestionsArgs();
    }

    public sealed class GetWafPbSuggestionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The minimum learning score for suggestions.
        /// </summary>
        [Input("minimumLearningScore", required: true)]
        public Input<int> MinimumLearningScore { get; set; } = null!;

        /// <summary>
        /// Partition on which WAF policy is located.
        /// </summary>
        [Input("partition", required: true)]
        public Input<string> Partition { get; set; } = null!;

        /// <summary>
        /// System generated id of the WAF policy
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// WAF policy name from which PB suggestions should be exported.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        public GetWafPbSuggestionsInvokeArgs()
        {
        }
        public static new GetWafPbSuggestionsInvokeArgs Empty => new GetWafPbSuggestionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetWafPbSuggestionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Json string representing exported PB suggestions ready to be used in WAF policy declaration
        /// </summary>
        public readonly string Json;
        public readonly int MinimumLearningScore;
        public readonly string Partition;
        /// <summary>
        /// System generated id of the WAF policy
        /// </summary>
        public readonly string PolicyId;
        public readonly string PolicyName;

        [OutputConstructor]
        private GetWafPbSuggestionsResult(
            string id,

            string json,

            int minimumLearningScore,

            string partition,

            string policyId,

            string policyName)
        {
            Id = id;
            Json = json;
            MinimumLearningScore = minimumLearningScore;
            Partition = partition;
            PolicyId = policyId;
            PolicyName = policyName;
        }
    }
}
