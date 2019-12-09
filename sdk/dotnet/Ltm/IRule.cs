// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_irule.html.markdown.
    /// </summary>
    public partial class IRule : Pulumi.CustomResource
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
            : base("f5bigip:ltm/iRule:IRule", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class IRuleArgs : Pulumi.ResourceArgs
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
    }

    public sealed class IRuleState : Pulumi.ResourceArgs
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
    }
}
