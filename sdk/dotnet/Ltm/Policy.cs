// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.Policy` Configures ltm policies to manage traffic assigned to a virtual server
    /// 
    /// For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource. For example `/Common/test-policy`.
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
    ///     var mypool = new F5BigIP.Ltm.Pool("mypool", new()
    ///     {
    ///         Name = "/Common/test-pool",
    ///         AllowNat = "yes",
    ///         AllowSnat = "yes",
    ///         LoadBalancingMode = "round-robin",
    ///     });
    /// 
    ///     var test_policy = new F5BigIP.Ltm.Policy("test-policy", new()
    ///     {
    ///         Name = "/Common/test-policy",
    ///         Strategy = "first-match",
    ///         Requires = new[]
    ///         {
    ///             "http",
    ///         },
    ///         Controls = new[]
    ///         {
    ///             "forwarding",
    ///         },
    ///         Rules = new[]
    ///         {
    ///             new F5BigIP.Ltm.Inputs.PolicyRuleArgs
    ///             {
    ///                 Name = "rule6",
    ///                 Actions = new[]
    ///                 {
    ///                     new F5BigIP.Ltm.Inputs.PolicyRuleActionArgs
    ///                     {
    ///                         Forward = true,
    ///                         Connection = false,
    ///                         Pool = mypool.Name,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             mypool,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/policy:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the controls
        /// </summary>
        [Output("controls")]
        public Output<ImmutableArray<string>> Controls { get; private set; } = null!;

        /// <summary>
        /// Name of Rule to be applied in policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// If you want to publish the policy else it will be deployed in Drafts mode.
        /// </summary>
        [Output("publishedCopy")]
        public Output<string?> PublishedCopy { get; private set; } = null!;

        /// <summary>
        /// Specifies the protocol
        /// </summary>
        [Output("requires")]
        public Output<ImmutableArray<string>> Requires { get; private set; } = null!;

        /// <summary>
        /// List of Rules can be applied using the policy. Each rule is block type with following arguments.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.PolicyRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// Specifies the match strategy
        /// </summary>
        [Output("strategy")]
        public Output<string?> Strategy { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/policy:Policy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("controls")]
        private InputList<string>? _controls;

        /// <summary>
        /// Specifies the controls
        /// </summary>
        public InputList<string> Controls
        {
            get => _controls ?? (_controls = new InputList<string>());
            set => _controls = value;
        }

        /// <summary>
        /// Name of Rule to be applied in policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// If you want to publish the policy else it will be deployed in Drafts mode.
        /// </summary>
        [Input("publishedCopy")]
        public Input<string>? PublishedCopy { get; set; }

        [Input("requires")]
        private InputList<string>? _requires;

        /// <summary>
        /// Specifies the protocol
        /// </summary>
        public InputList<string> Requires
        {
            get => _requires ?? (_requires = new InputList<string>());
            set => _requires = value;
        }

        [Input("rules")]
        private InputList<Inputs.PolicyRuleArgs>? _rules;

        /// <summary>
        /// List of Rules can be applied using the policy. Each rule is block type with following arguments.
        /// </summary>
        public InputList<Inputs.PolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.PolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the match strategy
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }

    public sealed class PolicyState : global::Pulumi.ResourceArgs
    {
        [Input("controls")]
        private InputList<string>? _controls;

        /// <summary>
        /// Specifies the controls
        /// </summary>
        public InputList<string> Controls
        {
            get => _controls ?? (_controls = new InputList<string>());
            set => _controls = value;
        }

        /// <summary>
        /// Name of Rule to be applied in policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If you want to publish the policy else it will be deployed in Drafts mode.
        /// </summary>
        [Input("publishedCopy")]
        public Input<string>? PublishedCopy { get; set; }

        [Input("requires")]
        private InputList<string>? _requires;

        /// <summary>
        /// Specifies the protocol
        /// </summary>
        public InputList<string> Requires
        {
            get => _requires ?? (_requires = new InputList<string>());
            set => _requires = value;
        }

        [Input("rules")]
        private InputList<Inputs.PolicyRuleGetArgs>? _rules;

        /// <summary>
        /// List of Rules can be applied using the policy. Each rule is block type with following arguments.
        /// </summary>
        public InputList<Inputs.PolicyRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.PolicyRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Specifies the match strategy
        /// </summary>
        [Input("strategy")]
        public Input<string>? Strategy { get; set; }

        public PolicyState()
        {
        }
        public static new PolicyState Empty => new PolicyState();
    }
}
