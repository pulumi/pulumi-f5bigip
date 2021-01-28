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
    /// `f5bigip.ltm.DataGroup` Manages internal (in-line) datagroup configuration
    /// 
    /// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-datagroup.
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
    ///         var datagroup = new F5BigIP.Ltm.DataGroup("datagroup", new F5BigIP.Ltm.DataGroupArgs
    ///         {
    ///             Name = "/Common/dgx2",
    ///             Records = 
    ///             {
    ///                 new F5BigIP.Ltm.Inputs.DataGroupRecordArgs
    ///                 {
    ///                     Data = "pool1",
    ///                     Name = "abc.com",
    ///                 },
    ///                 new F5BigIP.Ltm.Inputs.DataGroupRecordArgs
    ///                 {
    ///                     Data = "123",
    ///                     Name = "test",
    ///                 },
    ///             },
    ///             Type = "string",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/dataGroup:DataGroup")]
    public partial class DataGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        /// </summary>
        [Output("records")]
        public Output<ImmutableArray<Outputs.DataGroupRecord>> Records { get; private set; } = null!;

        /// <summary>
        /// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DataGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataGroup(string name, DataGroupArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/dataGroup:DataGroup", name, args ?? new DataGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataGroup(string name, Input<string> id, DataGroupState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/dataGroup:DataGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataGroup Get(string name, Input<string> id, DataGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new DataGroup(name, id, state, options);
        }
    }

    public sealed class DataGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("records")]
        private InputList<Inputs.DataGroupRecordArgs>? _records;

        /// <summary>
        /// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        /// </summary>
        public InputList<Inputs.DataGroupRecordArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.DataGroupRecordArgs>());
            set => _records = value;
        }

        /// <summary>
        /// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DataGroupArgs()
        {
        }
    }

    public sealed class DataGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("records")]
        private InputList<Inputs.DataGroupRecordGetArgs>? _records;

        /// <summary>
        /// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        /// </summary>
        public InputList<Inputs.DataGroupRecordGetArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.DataGroupRecordGetArgs>());
            set => _records = value;
        }

        /// <summary>
        /// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DataGroupState()
        {
        }
    }
}
