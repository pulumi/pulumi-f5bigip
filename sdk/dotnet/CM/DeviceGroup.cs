// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.CM
{
    /// <summary>
    /// `f5bigip.cm.DeviceGroup` A device group is a collection of BIG-IP devices that are configured to securely synchronize their BIG-IP configuration data, and fail over when needed.
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
    ///     var myNewDevicegroup = new F5BigIP.CM.DeviceGroup("my_new_devicegroup", new()
    ///     {
    ///         Name = "sanjose_devicegroup",
    ///         AutoSync = "enabled",
    ///         FullLoadOnSync = "true",
    ///         Type = "sync-only",
    ///         Devices = new[]
    ///         {
    ///             new F5BigIP.CM.Inputs.DeviceGroupDeviceArgs
    ///             {
    ///                 Name = "bigip1.cisco.com",
    ///             },
    ///             new F5BigIP.CM.Inputs.DeviceGroupDeviceArgs
    ///             {
    ///                 Name = "bigip200.f5.com",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:cm/deviceGroup:DeviceGroup")]
    public partial class DeviceGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies if the device-group will automatically sync configuration data to its members
        /// </summary>
        [Output("autoSync")]
        public Output<string?> AutoSync { get; private set; } = null!;

        /// <summary>
        /// Description of Device group
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the device to be included in device group, this need to be configured before using devicegroup resource
        /// </summary>
        [Output("devices")]
        public Output<ImmutableArray<Outputs.DeviceGroupDevice>> Devices { get; private set; } = null!;

        /// <summary>
        /// Specifies if the device-group will perform a full-load upon sync
        /// </summary>
        [Output("fullLoadOnSync")]
        public Output<string?> FullLoadOnSync { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
        /// </summary>
        [Output("incrementalConfig")]
        public Output<int?> IncrementalConfig { get; private set; } = null!;

        /// <summary>
        /// Is the name of the device Group
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies if the device-group will use a network connection for failover
        /// </summary>
        [Output("networkFailover")]
        public Output<string?> NetworkFailover { get; private set; } = null!;

        /// <summary>
        /// Device administrative partition
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the configuration should be saved upon auto-sync.
        /// </summary>
        [Output("saveOnAutoSync")]
        public Output<string?> SaveOnAutoSync { get; private set; } = null!;

        /// <summary>
        /// Specifies if the device-group will be used for failover or resource syncing
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceGroup(string name, DeviceGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("f5bigip:cm/deviceGroup:DeviceGroup", name, args ?? new DeviceGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceGroup(string name, Input<string> id, DeviceGroupState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:cm/deviceGroup:DeviceGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeviceGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceGroup Get(string name, Input<string> id, DeviceGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new DeviceGroup(name, id, state, options);
        }
    }

    public sealed class DeviceGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies if the device-group will automatically sync configuration data to its members
        /// </summary>
        [Input("autoSync")]
        public Input<string>? AutoSync { get; set; }

        /// <summary>
        /// Description of Device group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("devices")]
        private InputList<Inputs.DeviceGroupDeviceArgs>? _devices;

        /// <summary>
        /// Name of the device to be included in device group, this need to be configured before using devicegroup resource
        /// </summary>
        public InputList<Inputs.DeviceGroupDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.DeviceGroupDeviceArgs>());
            set => _devices = value;
        }

        /// <summary>
        /// Specifies if the device-group will perform a full-load upon sync
        /// </summary>
        [Input("fullLoadOnSync")]
        public Input<string>? FullLoadOnSync { get; set; }

        /// <summary>
        /// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
        /// </summary>
        [Input("incrementalConfig")]
        public Input<int>? IncrementalConfig { get; set; }

        /// <summary>
        /// Is the name of the device Group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies if the device-group will use a network connection for failover
        /// </summary>
        [Input("networkFailover")]
        public Input<string>? NetworkFailover { get; set; }

        /// <summary>
        /// Device administrative partition
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies whether the configuration should be saved upon auto-sync.
        /// </summary>
        [Input("saveOnAutoSync")]
        public Input<string>? SaveOnAutoSync { get; set; }

        /// <summary>
        /// Specifies if the device-group will be used for failover or resource syncing
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DeviceGroupArgs()
        {
        }
        public static new DeviceGroupArgs Empty => new DeviceGroupArgs();
    }

    public sealed class DeviceGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies if the device-group will automatically sync configuration data to its members
        /// </summary>
        [Input("autoSync")]
        public Input<string>? AutoSync { get; set; }

        /// <summary>
        /// Description of Device group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("devices")]
        private InputList<Inputs.DeviceGroupDeviceGetArgs>? _devices;

        /// <summary>
        /// Name of the device to be included in device group, this need to be configured before using devicegroup resource
        /// </summary>
        public InputList<Inputs.DeviceGroupDeviceGetArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.DeviceGroupDeviceGetArgs>());
            set => _devices = value;
        }

        /// <summary>
        /// Specifies if the device-group will perform a full-load upon sync
        /// </summary>
        [Input("fullLoadOnSync")]
        public Input<string>? FullLoadOnSync { get; set; }

        /// <summary>
        /// Specifies the maximum size (in KB) to devote to incremental config sync cached transactions. The default is 1024 KB.
        /// </summary>
        [Input("incrementalConfig")]
        public Input<int>? IncrementalConfig { get; set; }

        /// <summary>
        /// Is the name of the device Group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies if the device-group will use a network connection for failover
        /// </summary>
        [Input("networkFailover")]
        public Input<string>? NetworkFailover { get; set; }

        /// <summary>
        /// Device administrative partition
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies whether the configuration should be saved upon auto-sync.
        /// </summary>
        [Input("saveOnAutoSync")]
        public Input<string>? SaveOnAutoSync { get; set; }

        /// <summary>
        /// Specifies if the device-group will be used for failover or resource syncing
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DeviceGroupState()
        {
        }
        public static new DeviceGroupState Empty => new DeviceGroupState();
    }
}
