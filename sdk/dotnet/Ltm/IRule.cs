// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.IRule` Creates iRule on BIG-IP F5 device
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// using Std = Pulumi.Std;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Loading from a file is the preferred method
    ///     var rule = new F5BigIP.Ltm.IRule("rule", new()
    ///     {
    ///         Name = "/Common/terraform_irule",
    ///         Irule = Std.File.Invoke(new()
    ///         {
    ///             Input = "myirule.tcl",
    ///         }).Apply(invoke =&gt; invoke.Result),
    ///     });
    /// 
    ///     var rule2 = new F5BigIP.Ltm.IRule("rule2", new()
    ///     {
    ///         Name = "/Common/terraform_irule2",
    ///         Irule = @"when CLIENT_ACCEPTED {
    ///      log local0. ""test""
    ///    }
    /// ",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ##myirule.tcl
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/iRule:IRule")]
    public partial class IRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Body of the iRule
        /// </summary>
        [Output("irule")]
        public Output<string> Irule { get; private set; } = null!;

        /// <summary>
        /// Name of the iRule
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a IRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IRule(string name, IRuleArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/iRule:IRule", name, args ?? new IRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IRule(string name, Input<string> id, IRuleState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/iRule:IRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IRule Get(string name, Input<string> id, IRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new IRule(name, id, state, options);
        }
    }

    public sealed class IRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Body of the iRule
        /// </summary>
        [Input("irule", required: true)]
        public Input<string> Irule { get; set; } = null!;

        /// <summary>
        /// Name of the iRule
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public IRuleArgs()
        {
        }
        public static new IRuleArgs Empty => new IRuleArgs();
    }

    public sealed class IRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Body of the iRule
        /// </summary>
        [Input("irule")]
        public Input<string>? Irule { get; set; }

        /// <summary>
        /// Name of the iRule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IRuleState()
        {
        }
        public static new IRuleState Empty => new IRuleState();
    }
}
