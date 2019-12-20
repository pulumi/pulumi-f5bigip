// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Net
{
    /// <summary>
    /// `f5bigip.net.Vlan` Manages a vlan configuration
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/net_vlan.html.markdown.
    /// </summary>
    public partial class Vlan : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies which interfaces you want this VLAN to use for traffic management.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.VlanInterfaces>> Interfaces { get; private set; } = null!;

        /// <summary>
        /// Name of the vlan
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies a number that the system adds into the header of any frame passing through the VLAN.
        /// </summary>
        [Output("tag")]
        public Output<int?> Tag { get; private set; } = null!;


        /// <summary>
        /// Create a Vlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vlan(string name, VlanArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:net/vlan:Vlan", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Vlan(string name, Input<string> id, VlanState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:net/vlan:Vlan", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vlan Get(string name, Input<string> id, VlanState? state = null, CustomResourceOptions? options = null)
        {
            return new Vlan(name, id, state, options);
        }
    }

    public sealed class VlanArgs : Pulumi.ResourceArgs
    {
        [Input("interfaces")]
        private InputList<Inputs.VlanInterfacesArgs>? _interfaces;

        /// <summary>
        /// Specifies which interfaces you want this VLAN to use for traffic management.
        /// </summary>
        public InputList<Inputs.VlanInterfacesArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.VlanInterfacesArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Name of the vlan
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies a number that the system adds into the header of any frame passing through the VLAN.
        /// </summary>
        [Input("tag")]
        public Input<int>? Tag { get; set; }

        public VlanArgs()
        {
        }
    }

    public sealed class VlanState : Pulumi.ResourceArgs
    {
        [Input("interfaces")]
        private InputList<Inputs.VlanInterfacesGetArgs>? _interfaces;

        /// <summary>
        /// Specifies which interfaces you want this VLAN to use for traffic management.
        /// </summary>
        public InputList<Inputs.VlanInterfacesGetArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.VlanInterfacesGetArgs>());
            set => _interfaces = value;
        }

        /// <summary>
        /// Name of the vlan
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies a number that the system adds into the header of any frame passing through the VLAN.
        /// </summary>
        [Input("tag")]
        public Input<int>? Tag { get; set; }

        public VlanState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class VlanInterfacesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        /// </summary>
        [Input("tagged")]
        public Input<bool>? Tagged { get; set; }

        /// <summary>
        /// Physical or virtual port used for traffic
        /// </summary>
        [Input("vlanport")]
        public Input<string>? Vlanport { get; set; }

        public VlanInterfacesArgs()
        {
        }
    }

    public sealed class VlanInterfacesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        /// </summary>
        [Input("tagged")]
        public Input<bool>? Tagged { get; set; }

        /// <summary>
        /// Physical or virtual port used for traffic
        /// </summary>
        [Input("vlanport")]
        public Input<string>? Vlanport { get; set; }

        public VlanInterfacesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class VlanInterfaces
    {
        /// <summary>
        /// Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        /// </summary>
        public readonly bool? Tagged;
        /// <summary>
        /// Physical or virtual port used for traffic
        /// </summary>
        public readonly string? Vlanport;

        [OutputConstructor]
        private VlanInterfaces(
            bool? tagged,
            string? vlanport)
        {
            Tagged = tagged;
            Vlanport = vlanport;
        }
    }
    }
}