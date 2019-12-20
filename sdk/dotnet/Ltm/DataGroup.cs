// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_datagroup.html.markdown.
    /// </summary>
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
        public Output<ImmutableArray<Outputs.DataGroupRecords>> Records { get; private set; } = null!;

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
            : base("f5bigip:ltm/dataGroup:DataGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
        private InputList<Inputs.DataGroupRecordsArgs>? _records;

        /// <summary>
        /// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        /// </summary>
        public InputList<Inputs.DataGroupRecordsArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.DataGroupRecordsArgs>());
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
        private InputList<Inputs.DataGroupRecordsGetArgs>? _records;

        /// <summary>
        /// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        /// </summary>
        public InputList<Inputs.DataGroupRecordsGetArgs> Records
        {
            get => _records ?? (_records = new InputList<Inputs.DataGroupRecordsGetArgs>());
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

    namespace Inputs
    {

    public sealed class DataGroupRecordsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public DataGroupRecordsArgs()
        {
        }
    }

    public sealed class DataGroupRecordsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public DataGroupRecordsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DataGroupRecords
    {
        /// <summary>
        /// , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
        /// </summary>
        public readonly string? Data;
        /// <summary>
        /// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private DataGroupRecords(
            string? data,
            string name)
        {
            Data = data;
            Name = name;
        }
    }
    }
}