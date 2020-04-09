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
    /// `f5bigip.ltm.Snat` Manages a snat configuration
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_snat.html.markdown.
    /// </summary>
    public partial class Snat : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether to automatically map last hop for pools or not. The default is to use next level's defaul
        /// </summary>
        [Output("autolasthop")]
        public Output<string?> Autolasthop { get; private set; } = null!;

        /// <summary>
        /// Fullpath
        /// </summary>
        [Output("fullPath")]
        public Output<string?> FullPath { get; private set; } = null!;

        /// <summary>
        /// Enables or disables mirroring of SNAT connections.
        /// </summary>
        [Output("mirror")]
        public Output<string?> Mirror { get; private set; } = null!;

        /// <summary>
        /// Name of the snat
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// IP or hostname of the snat
        /// </summary>
        [Output("origins")]
        public Output<ImmutableArray<Outputs.SnatOrigins>> Origins { get; private set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
        /// </summary>
        [Output("snatpool")]
        public Output<string?> Snatpool { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
        /// </summary>
        [Output("sourceport")]
        public Output<string?> Sourceport { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
        /// </summary>
        [Output("translation")]
        public Output<string?> Translation { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
        /// </summary>
        [Output("vlans")]
        public Output<ImmutableArray<string>> Vlans { get; private set; } = null!;

        /// <summary>
        /// Disables the SNAT on all VLANs.
        /// </summary>
        [Output("vlansdisabled")]
        public Output<bool?> Vlansdisabled { get; private set; } = null!;


        /// <summary>
        /// Create a Snat resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Snat(string name, SnatArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/snat:Snat", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Snat(string name, Input<string> id, SnatState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/snat:Snat", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Snat resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Snat Get(string name, Input<string> id, SnatState? state = null, CustomResourceOptions? options = null)
        {
            return new Snat(name, id, state, options);
        }
    }

    public sealed class SnatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to automatically map last hop for pools or not. The default is to use next level's defaul
        /// </summary>
        [Input("autolasthop")]
        public Input<string>? Autolasthop { get; set; }

        /// <summary>
        /// Fullpath
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        /// <summary>
        /// Enables or disables mirroring of SNAT connections.
        /// </summary>
        [Input("mirror")]
        public Input<string>? Mirror { get; set; }

        /// <summary>
        /// Name of the snat
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("origins", required: true)]
        private InputList<Inputs.SnatOriginsArgs>? _origins;

        /// <summary>
        /// IP or hostname of the snat
        /// </summary>
        public InputList<Inputs.SnatOriginsArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.SnatOriginsArgs>());
            set => _origins = value;
        }

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
        /// </summary>
        [Input("snatpool")]
        public Input<string>? Snatpool { get; set; }

        /// <summary>
        /// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
        /// </summary>
        [Input("sourceport")]
        public Input<string>? Sourceport { get; set; }

        /// <summary>
        /// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
        /// </summary>
        [Input("translation")]
        public Input<string>? Translation { get; set; }

        [Input("vlans")]
        private InputList<string>? _vlans;

        /// <summary>
        /// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
        /// </summary>
        public InputList<string> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<string>());
            set => _vlans = value;
        }

        /// <summary>
        /// Disables the SNAT on all VLANs.
        /// </summary>
        [Input("vlansdisabled")]
        public Input<bool>? Vlansdisabled { get; set; }

        public SnatArgs()
        {
        }
    }

    public sealed class SnatState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to automatically map last hop for pools or not. The default is to use next level's defaul
        /// </summary>
        [Input("autolasthop")]
        public Input<string>? Autolasthop { get; set; }

        /// <summary>
        /// Fullpath
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        /// <summary>
        /// Enables or disables mirroring of SNAT connections.
        /// </summary>
        [Input("mirror")]
        public Input<string>? Mirror { get; set; }

        /// <summary>
        /// Name of the snat
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("origins")]
        private InputList<Inputs.SnatOriginsGetArgs>? _origins;

        /// <summary>
        /// IP or hostname of the snat
        /// </summary>
        public InputList<Inputs.SnatOriginsGetArgs> Origins
        {
            get => _origins ?? (_origins = new InputList<Inputs.SnatOriginsGetArgs>());
            set => _origins = value;
        }

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
        /// </summary>
        [Input("snatpool")]
        public Input<string>? Snatpool { get; set; }

        /// <summary>
        /// Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
        /// </summary>
        [Input("sourceport")]
        public Input<string>? Sourceport { get; set; }

        /// <summary>
        /// Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
        /// </summary>
        [Input("translation")]
        public Input<string>? Translation { get; set; }

        [Input("vlans")]
        private InputList<string>? _vlans;

        /// <summary>
        /// Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
        /// </summary>
        public InputList<string> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<string>());
            set => _vlans = value;
        }

        /// <summary>
        /// Disables the SNAT on all VLANs.
        /// </summary>
        [Input("vlansdisabled")]
        public Input<bool>? Vlansdisabled { get; set; }

        public SnatState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SnatOriginsArgs : Pulumi.ResourceArgs
    {
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Name of the snat
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SnatOriginsArgs()
        {
        }
    }

    public sealed class SnatOriginsGetArgs : Pulumi.ResourceArgs
    {
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Name of the snat
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SnatOriginsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SnatOrigins
    {
        public readonly string? AppService;
        /// <summary>
        /// Name of the snat
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private SnatOrigins(
            string? appService,
            string? name)
        {
            AppService = appService;
            Name = name;
        }
    }
    }
}
