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
    /// `f5bigip.ltm.SnatPool` Collections of SNAT translation addresses
    /// 
    /// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-snatpool.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var snatpoolSanjose = new F5BigIP.Ltm.SnatPool("snatpoolSanjose", new F5BigIP.Ltm.SnatPoolArgs
    ///         {
    ///             Members = 
    ///             {
    ///                 "191.1.1.1",
    ///                 "194.2.2.2",
    ///             },
    ///             Name = "/Common/snatpool_sanjose",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SnatPool : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Name of the snatpool
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a SnatPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnatPool(string name, SnatPoolArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/snatPool:SnatPool", name, args ?? new SnatPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SnatPool(string name, Input<string> id, SnatPoolState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/snatPool:SnatPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnatPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnatPool Get(string name, Input<string> id, SnatPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new SnatPool(name, id, state, options);
        }
    }

    public sealed class SnatPoolArgs : Pulumi.ResourceArgs
    {
        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Name of the snatpool
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public SnatPoolArgs()
        {
        }
    }

    public sealed class SnatPoolState : Pulumi.ResourceArgs
    {
        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Name of the snatpool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SnatPoolState()
        {
        }
    }
}
