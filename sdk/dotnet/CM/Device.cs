// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.CM
{
    /// <summary>
    /// `f5bigip.cm.Device` provides details about a specific bigip
    /// 
    /// This resource is helpful when configuring the BIG-IP device in cluster or in HA mode.
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myNewDevice = new F5BigIP.CM.Device("my_new_device", new()
    ///     {
    ///         Name = "bigip300.f5.com",
    ///         ConfigsyncIp = "2.2.2.2",
    ///         MirrorIp = "10.10.10.10",
    ///         MirrorSecondaryIp = "11.11.11.11",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [F5BigIPResourceType("f5bigip:cm/device:Device")]
    public partial class Device : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP address used for config sync
        /// </summary>
        [Output("configsyncIp")]
        public Output<string> ConfigsyncIp { get; private set; } = null!;

        /// <summary>
        /// IP address used for state mirroring
        /// </summary>
        [Output("mirrorIp")]
        public Output<string?> MirrorIp { get; private set; } = null!;

        /// <summary>
        /// Secondary IP address used for state mirroring
        /// </summary>
        [Output("mirrorSecondaryIp")]
        public Output<string?> MirrorSecondaryIp { get; private set; } = null!;

        /// <summary>
        /// Address of the Device which needs to be Deviceensed
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Device resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Device(string name, DeviceArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:cm/device:Device", name, args ?? new DeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Device(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:cm/device:Device", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Device resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Device Get(string name, Input<string> id, DeviceState? state = null, CustomResourceOptions? options = null)
        {
            return new Device(name, id, state, options);
        }
    }

    public sealed class DeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address used for config sync
        /// </summary>
        [Input("configsyncIp", required: true)]
        public Input<string> ConfigsyncIp { get; set; } = null!;

        /// <summary>
        /// IP address used for state mirroring
        /// </summary>
        [Input("mirrorIp")]
        public Input<string>? MirrorIp { get; set; }

        /// <summary>
        /// Secondary IP address used for state mirroring
        /// </summary>
        [Input("mirrorSecondaryIp")]
        public Input<string>? MirrorSecondaryIp { get; set; }

        /// <summary>
        /// Address of the Device which needs to be Deviceensed
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public DeviceArgs()
        {
        }
        public static new DeviceArgs Empty => new DeviceArgs();
    }

    public sealed class DeviceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address used for config sync
        /// </summary>
        [Input("configsyncIp")]
        public Input<string>? ConfigsyncIp { get; set; }

        /// <summary>
        /// IP address used for state mirroring
        /// </summary>
        [Input("mirrorIp")]
        public Input<string>? MirrorIp { get; set; }

        /// <summary>
        /// Secondary IP address used for state mirroring
        /// </summary>
        [Input("mirrorSecondaryIp")]
        public Input<string>? MirrorSecondaryIp { get; set; }

        /// <summary>
        /// Address of the Device which needs to be Deviceensed
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DeviceState()
        {
        }
        public static new DeviceState Empty => new DeviceState();
    }
}
