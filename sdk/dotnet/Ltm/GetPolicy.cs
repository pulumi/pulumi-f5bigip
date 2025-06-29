// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    public static class GetPolicy
    {
        /// <summary>
        /// Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
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
        ///     var test = F5BigIP.Ltm.GetPolicy.Invoke(new()
        ///     {
        ///         Name = "/Common/test-policy",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipPolicy"] = test.Apply(getPolicyResult =&gt; getPolicyResult.Rules),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetPolicyResult> InvokeAsync(GetPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPolicyResult>("f5bigip:ltm/getPolicy:getPolicy", args ?? new GetPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
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
        ///     var test = F5BigIP.Ltm.GetPolicy.Invoke(new()
        ///     {
        ///         Name = "/Common/test-policy",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipPolicy"] = test.Apply(getPolicyResult =&gt; getPolicyResult.Rules),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetPolicyResult> Invoke(GetPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPolicyResult>("f5bigip:ltm/getPolicy:getPolicy", args ?? new GetPolicyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source (`f5bigip.ltm.Policy`) to get the ltm policy details available on BIG-IP
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
        ///     var test = F5BigIP.Ltm.GetPolicy.Invoke(new()
        ///     {
        ///         Name = "/Common/test-policy",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["bigipPolicy"] = test.Apply(getPolicyResult =&gt; getPolicyResult.Rules),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetPolicyResult> Invoke(GetPolicyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPolicyResult>("f5bigip:ltm/getPolicy:getPolicy", args ?? new GetPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("controls")]
        private List<string>? _controls;

        /// <summary>
        /// Specifies the controls.
        /// </summary>
        public List<string> Controls
        {
            get => _controls ?? (_controls = new List<string>());
            set => _controls = value;
        }

        /// <summary>
        /// Name of the policy which includes partion ( /partition/policy-name )
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("publishedCopy")]
        public string? PublishedCopy { get; set; }

        [Input("requires")]
        private List<string>? _requires;

        /// <summary>
        /// Specifies the protocol.
        /// </summary>
        public List<string> Requires
        {
            get => _requires ?? (_requires = new List<string>());
            set => _requires = value;
        }

        [Input("rules")]
        private List<Inputs.GetPolicyRuleArgs>? _rules;

        /// <summary>
        /// Rules defined in the policy.
        /// </summary>
        public List<Inputs.GetPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new List<Inputs.GetPolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the match strategy.
        /// </summary>
        [Input("strategy")]
        public string? Strategy { get; set; }

        public GetPolicyArgs()
        {
        }
        public static new GetPolicyArgs Empty => new GetPolicyArgs();
    }

    public sealed class GetPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("controls")]
        private InputList<string>? _controls;

        /// <summary>
        /// Specifies the controls.
        /// </summary>
        public InputList<string> Controls
        {
            get => _controls ?? (_controls = new InputList<string>());
            set => _controls = value;
        }

        /// <summary>
        /// Name of the policy which includes partion ( /partition/policy-name )
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("publishedCopy")]
        public Input<string>? PublishedCopy { get; set; }

        [Input("requires")]
        private InputList<string>? _requires;

        /// <summary>
        /// Specifies the protocol.
        /// </summary>
        public InputList<string> Requires
        {
            get => _requires ?? (_requires = new InputList<string>());
            set => _requires = value;
        }

        [Input("rules")]
        private InputList<Inputs.GetPolicyRuleInputArgs>? _rules;

        /// <summary>
        /// Rules defined in the policy.
        /// </summary>
        public InputList<Inputs.GetPolicyRuleInputArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.GetPolicyRuleInputArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the match strategy.
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        public GetPolicyInvokeArgs()
        {
        }
        public static new GetPolicyInvokeArgs Empty => new GetPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPolicyResult
    {
        /// <summary>
        /// Specifies the controls.
        /// </summary>
        public readonly ImmutableArray<string> Controls;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the policy.
        /// </summary>
        public readonly string Name;
        public readonly string? PublishedCopy;
        /// <summary>
        /// Specifies the protocol.
        /// </summary>
        public readonly ImmutableArray<string> Requires;
        /// <summary>
        /// Rules defined in the policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicyRuleResult> Rules;
        /// <summary>
        /// Specifies the match strategy.
        /// </summary>
        public readonly string? Strategy;

        [OutputConstructor]
        private GetPolicyResult(
            ImmutableArray<string> controls,

            string id,

            string name,

            string? publishedCopy,

            ImmutableArray<string> requires,

            ImmutableArray<Outputs.GetPolicyRuleResult> rules,

            string? strategy)
        {
            Controls = controls;
            Id = id;
            Name = name;
            PublishedCopy = publishedCopy;
            Requires = requires;
            Rules = rules;
            Strategy = strategy;
        }
    }
}
