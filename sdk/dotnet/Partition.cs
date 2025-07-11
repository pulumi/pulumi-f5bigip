// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// `f5bigip.Partition` Manages F5 BIG-IP partitions
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
    ///     var test_partition = new F5BigIP.Partition("test-partition", new()
    ///     {
    ///         Name = "test-partition",
    ///         Description = "created by terraform",
    ///         RouteDomainId = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Importing
    /// 
    /// An existing cipher group can be imported into this resource by supplying the cipher rule full path name ex : `/partition/name`
    /// 
    /// An example is below:
    /// 
    /// ```sh
    /// $ terraform import bigip_partition.test_partition test_partition
    /// 
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/partition:Partition")]
    public partial class Partition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the partition.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the partition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Route domain id of the partition.
        /// </summary>
        [Output("routeDomainId")]
        public Output<int?> RouteDomainId { get; private set; } = null!;


        /// <summary>
        /// Create a Partition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Partition(string name, PartitionArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/partition:Partition", name, args ?? new PartitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Partition(string name, Input<string> id, PartitionState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/partition:Partition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Partition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Partition Get(string name, Input<string> id, PartitionState? state = null, CustomResourceOptions? options = null)
        {
            return new Partition(name, id, state, options);
        }
    }

    public sealed class PartitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the partition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the partition.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Route domain id of the partition.
        /// </summary>
        [Input("routeDomainId")]
        public Input<int>? RouteDomainId { get; set; }

        public PartitionArgs()
        {
        }
        public static new PartitionArgs Empty => new PartitionArgs();
    }

    public sealed class PartitionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the partition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the partition.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Route domain id of the partition.
        /// </summary>
        [Input("routeDomainId")]
        public Input<int>? RouteDomainId { get; set; }

        public PartitionState()
        {
        }
        public static new PartitionState Empty => new PartitionState();
    }
}
