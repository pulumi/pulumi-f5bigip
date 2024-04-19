// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys
{
    /// <summary>
    /// `f5bigip.sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var simplehttp = new F5BigIP.Sys.IApp("simplehttp", new()
    ///     {
    ///         Name = "simplehttp",
    ///         Jsonfile = File.ReadAllText("simplehttp.json"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Json File
    /// </summary>
    [F5BigIPResourceType("f5bigip:sys/iApp:IApp")]
    public partial class IApp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User defined description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// BIG-IP password
        /// </summary>
        [Output("devicegroup")]
        public Output<string> Devicegroup { get; private set; } = null!;

        /// <summary>
        /// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `execute_action` attribute take precedence over `json` value
        /// </summary>
        [Output("executeAction")]
        public Output<string> ExecuteAction { get; private set; } = null!;

        /// <summary>
        /// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
        /// </summary>
        [Output("inheritedDevicegroup")]
        public Output<string?> InheritedDevicegroup { get; private set; } = null!;

        /// <summary>
        /// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
        /// </summary>
        [Output("inheritedTrafficGroup")]
        public Output<string?> InheritedTrafficGroup { get; private set; } = null!;

        /// <summary>
        /// Refer to the Json file which will be deployed on F5 BIG-IP.
        /// </summary>
        [Output("jsonfile")]
        public Output<string> Jsonfile { get; private set; } = null!;

        /// <summary>
        /// string values
        /// </summary>
        [Output("lists")]
        public Output<ImmutableArray<Outputs.IAppList>> Lists { get; private set; } = null!;

        /// <summary>
        /// User defined generic data for the application service. It is a name and value pair.
        /// </summary>
        [Output("metadatas")]
        public Output<ImmutableArray<Outputs.IAppMetadata>> Metadatas { get; private set; } = null!;

        /// <summary>
        /// Name of the iApp.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which the application resides.
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;

        /// <summary>
        /// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
        /// </summary>
        [Output("strictUpdates")]
        public Output<string?> StrictUpdates { get; private set; } = null!;

        [Output("tables")]
        public Output<ImmutableArray<Outputs.IAppTable>> Tables { get; private set; } = null!;

        /// <summary>
        /// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
        /// </summary>
        [Output("template")]
        public Output<string?> Template { get; private set; } = null!;

        /// <summary>
        /// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
        /// </summary>
        [Output("templateModified")]
        public Output<string?> TemplateModified { get; private set; } = null!;

        /// <summary>
        /// Indicates any missing prerequisites associated with the template that defines this application.
        /// </summary>
        [Output("templatePrerequisiteErrors")]
        public Output<string?> TemplatePrerequisiteErrors { get; private set; } = null!;

        /// <summary>
        /// The name of the traffic group that the application service is assigned to.
        /// </summary>
        [Output("trafficGroup")]
        public Output<string?> TrafficGroup { get; private set; } = null!;

        [Output("variables")]
        public Output<ImmutableArray<Outputs.IAppVariable>> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a IApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IApp(string name, IAppArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:sys/iApp:IApp", name, args ?? new IAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IApp(string name, Input<string> id, IAppState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:sys/iApp:IApp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IApp Get(string name, Input<string> id, IAppState? state = null, CustomResourceOptions? options = null)
        {
            return new IApp(name, id, state, options);
        }
    }

    public sealed class IAppArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User defined description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// BIG-IP password
        /// </summary>
        [Input("devicegroup")]
        public Input<string>? Devicegroup { get; set; }

        /// <summary>
        /// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `execute_action` attribute take precedence over `json` value
        /// </summary>
        [Input("executeAction")]
        public Input<string>? ExecuteAction { get; set; }

        /// <summary>
        /// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
        /// </summary>
        [Input("inheritedDevicegroup")]
        public Input<string>? InheritedDevicegroup { get; set; }

        /// <summary>
        /// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
        /// </summary>
        [Input("inheritedTrafficGroup")]
        public Input<string>? InheritedTrafficGroup { get; set; }

        /// <summary>
        /// Refer to the Json file which will be deployed on F5 BIG-IP.
        /// </summary>
        [Input("jsonfile", required: true)]
        public Input<string> Jsonfile { get; set; } = null!;

        [Input("lists")]
        private InputList<Inputs.IAppListArgs>? _lists;

        /// <summary>
        /// string values
        /// </summary>
        public InputList<Inputs.IAppListArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.IAppListArgs>());
            set => _lists = value;
        }

        [Input("metadatas")]
        private InputList<Inputs.IAppMetadataArgs>? _metadatas;

        /// <summary>
        /// User defined generic data for the application service. It is a name and value pair.
        /// </summary>
        public InputList<Inputs.IAppMetadataArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.IAppMetadataArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Name of the iApp.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which the application resides.
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
        /// </summary>
        [Input("strictUpdates")]
        public Input<string>? StrictUpdates { get; set; }

        [Input("tables")]
        private InputList<Inputs.IAppTableArgs>? _tables;
        public InputList<Inputs.IAppTableArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.IAppTableArgs>());
            set => _tables = value;
        }

        /// <summary>
        /// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
        /// </summary>
        [Input("templateModified")]
        public Input<string>? TemplateModified { get; set; }

        /// <summary>
        /// Indicates any missing prerequisites associated with the template that defines this application.
        /// </summary>
        [Input("templatePrerequisiteErrors")]
        public Input<string>? TemplatePrerequisiteErrors { get; set; }

        /// <summary>
        /// The name of the traffic group that the application service is assigned to.
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        [Input("variables")]
        private InputList<Inputs.IAppVariableArgs>? _variables;
        public InputList<Inputs.IAppVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.IAppVariableArgs>());
            set => _variables = value;
        }

        public IAppArgs()
        {
        }
        public static new IAppArgs Empty => new IAppArgs();
    }

    public sealed class IAppState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User defined description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// BIG-IP password
        /// </summary>
        [Input("devicegroup")]
        public Input<string>? Devicegroup { get; set; }

        /// <summary>
        /// Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `execute_action` attribute take precedence over `json` value
        /// </summary>
        [Input("executeAction")]
        public Input<string>? ExecuteAction { get; set; }

        /// <summary>
        /// Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
        /// </summary>
        [Input("inheritedDevicegroup")]
        public Input<string>? InheritedDevicegroup { get; set; }

        /// <summary>
        /// Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
        /// </summary>
        [Input("inheritedTrafficGroup")]
        public Input<string>? InheritedTrafficGroup { get; set; }

        /// <summary>
        /// Refer to the Json file which will be deployed on F5 BIG-IP.
        /// </summary>
        [Input("jsonfile")]
        public Input<string>? Jsonfile { get; set; }

        [Input("lists")]
        private InputList<Inputs.IAppListGetArgs>? _lists;

        /// <summary>
        /// string values
        /// </summary>
        public InputList<Inputs.IAppListGetArgs> Lists
        {
            get => _lists ?? (_lists = new InputList<Inputs.IAppListGetArgs>());
            set => _lists = value;
        }

        [Input("metadatas")]
        private InputList<Inputs.IAppMetadataGetArgs>? _metadatas;

        /// <summary>
        /// User defined generic data for the application service. It is a name and value pair.
        /// </summary>
        public InputList<Inputs.IAppMetadataGetArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.IAppMetadataGetArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Name of the iApp.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Displays the administrative partition within which the application resides.
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
        /// </summary>
        [Input("strictUpdates")]
        public Input<string>? StrictUpdates { get; set; }

        [Input("tables")]
        private InputList<Inputs.IAppTableGetArgs>? _tables;
        public InputList<Inputs.IAppTableGetArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.IAppTableGetArgs>());
            set => _tables = value;
        }

        /// <summary>
        /// The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        /// <summary>
        /// Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
        /// </summary>
        [Input("templateModified")]
        public Input<string>? TemplateModified { get; set; }

        /// <summary>
        /// Indicates any missing prerequisites associated with the template that defines this application.
        /// </summary>
        [Input("templatePrerequisiteErrors")]
        public Input<string>? TemplatePrerequisiteErrors { get; set; }

        /// <summary>
        /// The name of the traffic group that the application service is assigned to.
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        [Input("variables")]
        private InputList<Inputs.IAppVariableGetArgs>? _variables;
        public InputList<Inputs.IAppVariableGetArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.IAppVariableGetArgs>());
            set => _variables = value;
        }

        public IAppState()
        {
        }
        public static new IAppState Empty => new IAppState();
    }
}
