// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.VCMP
{
    /// <summary>
    /// `f5bigip.vcmp.Guest` Manages a vCMP guest configuration
    /// 
    /// Resource does not wait for vCMP guest to reach the desired state, it only ensures that a desired configuration is set on the target device.
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
    ///     var vcmp_test = new F5BigIP.VCMP.Guest("vcmp-test", new()
    ///     {
    ///         Name = "tf_guest",
    ///         InitialImage = "12.1.2.iso",
    ///         MgmtNetwork = "bridged",
    ///         MgmtAddress = "10.1.1.1/24",
    ///         MgmtRoute = "none",
    ///         State = "provisioned",
    ///         CoresPerSlot = 2,
    ///         NumberOfSlots = 1,
    ///         MinNumberOfSlots = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:vcmp/guest:Guest")]
    public partial class Guest : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Contains those slots to which the guest is allowed to be assigned.
        /// </summary>
        [Output("allowedSlots")]
        public Output<ImmutableArray<int>> AllowedSlots { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of cores the system allocates to the guest.
        /// </summary>
        [Output("coresPerSlot")]
        public Output<int> CoresPerSlot { get; private set; } = null!;

        /// <summary>
        /// Indicates if virtual disk associated with vCMP guest should be removed during remove operation.  The default is `true`
        /// </summary>
        [Output("deleteVirtualDisk")]
        public Output<bool?> DeleteVirtualDisk { get; private set; } = null!;

        /// <summary>
        /// Resource name including prepended partition path.
        /// </summary>
        [Output("fullPath")]
        public Output<string> FullPath { get; private set; } = null!;

        /// <summary>
        /// Specifies the hotfix ISO image file which is applied on top of the base image.
        /// </summary>
        [Output("initialHotfix")]
        public Output<string> InitialHotfix { get; private set; } = null!;

        /// <summary>
        /// Specifies the base software release ISO image file for installing the TMOS hypervisor instance.
        /// </summary>
        [Output("initialImage")]
        public Output<string> InitialImage { get; private set; } = null!;

        /// <summary>
        /// Specifies the IP address and subnet or subnet mask you use to access the guest when you want to manage a module running within the guest.
        /// </summary>
        [Output("mgmtAddress")]
        public Output<string> MgmtAddress { get; private set; } = null!;

        /// <summary>
        /// Specifies the method by which the management address is used in the vCMP guest. options : [`bridged`,`isolated`,`host-only`].
        /// </summary>
        [Output("mgmtNetwork")]
        public Output<string> MgmtNetwork { get; private set; } = null!;

        /// <summary>
        /// Specifies the gateway address for the `mgmt_address`. Can be set to `none` to remove the value from the configuration.
        /// </summary>
        [Output("mgmtRoute")]
        public Output<string> MgmtRoute { get; private set; } = null!;

        /// <summary>
        /// Specifies the minimum number of slots the guest must be assigned to in order to deploy.
        /// </summary>
        [Output("minNumberOfSlots")]
        public Output<int> MinNumberOfSlots { get; private set; } = null!;

        /// <summary>
        /// Name of the vCMP guest
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of slots for the system to use when creating the guest.
        /// </summary>
        [Output("numberOfSlots")]
        public Output<int> NumberOfSlots { get; private set; } = null!;

        /// <summary>
        /// Specifies the state of the vCMP guest on the system. options : [`configured`,`provisioned`,`deployed`].
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Virtual disk associated with vCMP guest.
        /// </summary>
        [Output("virtualDisk")]
        public Output<string> VirtualDisk { get; private set; } = null!;

        /// <summary>
        /// Specifies the list of VLANs the vCMP guest uses to communicate with other guests, the host, and with the external network. The naming format must be the combination of the partition + name. For example /Common/my-vlan
        /// </summary>
        [Output("vlans")]
        public Output<ImmutableArray<string>> Vlans { get; private set; } = null!;


        /// <summary>
        /// Create a Guest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Guest(string name, GuestArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:vcmp/guest:Guest", name, args ?? new GuestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Guest(string name, Input<string> id, GuestState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:vcmp/guest:Guest", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Guest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Guest Get(string name, Input<string> id, GuestState? state = null, CustomResourceOptions? options = null)
        {
            return new Guest(name, id, state, options);
        }
    }

    public sealed class GuestArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedSlots")]
        private InputList<int>? _allowedSlots;

        /// <summary>
        /// Contains those slots to which the guest is allowed to be assigned.
        /// </summary>
        public InputList<int> AllowedSlots
        {
            get => _allowedSlots ?? (_allowedSlots = new InputList<int>());
            set => _allowedSlots = value;
        }

        /// <summary>
        /// Specifies the number of cores the system allocates to the guest.
        /// </summary>
        [Input("coresPerSlot")]
        public Input<int>? CoresPerSlot { get; set; }

        /// <summary>
        /// Indicates if virtual disk associated with vCMP guest should be removed during remove operation.  The default is `true`
        /// </summary>
        [Input("deleteVirtualDisk")]
        public Input<bool>? DeleteVirtualDisk { get; set; }

        /// <summary>
        /// Specifies the hotfix ISO image file which is applied on top of the base image.
        /// </summary>
        [Input("initialHotfix")]
        public Input<string>? InitialHotfix { get; set; }

        /// <summary>
        /// Specifies the base software release ISO image file for installing the TMOS hypervisor instance.
        /// </summary>
        [Input("initialImage")]
        public Input<string>? InitialImage { get; set; }

        /// <summary>
        /// Specifies the IP address and subnet or subnet mask you use to access the guest when you want to manage a module running within the guest.
        /// </summary>
        [Input("mgmtAddress")]
        public Input<string>? MgmtAddress { get; set; }

        /// <summary>
        /// Specifies the method by which the management address is used in the vCMP guest. options : [`bridged`,`isolated`,`host-only`].
        /// </summary>
        [Input("mgmtNetwork")]
        public Input<string>? MgmtNetwork { get; set; }

        /// <summary>
        /// Specifies the gateway address for the `mgmt_address`. Can be set to `none` to remove the value from the configuration.
        /// </summary>
        [Input("mgmtRoute")]
        public Input<string>? MgmtRoute { get; set; }

        /// <summary>
        /// Specifies the minimum number of slots the guest must be assigned to in order to deploy.
        /// </summary>
        [Input("minNumberOfSlots")]
        public Input<int>? MinNumberOfSlots { get; set; }

        /// <summary>
        /// Name of the vCMP guest
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the number of slots for the system to use when creating the guest.
        /// </summary>
        [Input("numberOfSlots")]
        public Input<int>? NumberOfSlots { get; set; }

        /// <summary>
        /// Specifies the state of the vCMP guest on the system. options : [`configured`,`provisioned`,`deployed`].
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("vlans")]
        private InputList<string>? _vlans;

        /// <summary>
        /// Specifies the list of VLANs the vCMP guest uses to communicate with other guests, the host, and with the external network. The naming format must be the combination of the partition + name. For example /Common/my-vlan
        /// </summary>
        public InputList<string> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<string>());
            set => _vlans = value;
        }

        public GuestArgs()
        {
        }
        public static new GuestArgs Empty => new GuestArgs();
    }

    public sealed class GuestState : global::Pulumi.ResourceArgs
    {
        [Input("allowedSlots")]
        private InputList<int>? _allowedSlots;

        /// <summary>
        /// Contains those slots to which the guest is allowed to be assigned.
        /// </summary>
        public InputList<int> AllowedSlots
        {
            get => _allowedSlots ?? (_allowedSlots = new InputList<int>());
            set => _allowedSlots = value;
        }

        /// <summary>
        /// Specifies the number of cores the system allocates to the guest.
        /// </summary>
        [Input("coresPerSlot")]
        public Input<int>? CoresPerSlot { get; set; }

        /// <summary>
        /// Indicates if virtual disk associated with vCMP guest should be removed during remove operation.  The default is `true`
        /// </summary>
        [Input("deleteVirtualDisk")]
        public Input<bool>? DeleteVirtualDisk { get; set; }

        /// <summary>
        /// Resource name including prepended partition path.
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        /// <summary>
        /// Specifies the hotfix ISO image file which is applied on top of the base image.
        /// </summary>
        [Input("initialHotfix")]
        public Input<string>? InitialHotfix { get; set; }

        /// <summary>
        /// Specifies the base software release ISO image file for installing the TMOS hypervisor instance.
        /// </summary>
        [Input("initialImage")]
        public Input<string>? InitialImage { get; set; }

        /// <summary>
        /// Specifies the IP address and subnet or subnet mask you use to access the guest when you want to manage a module running within the guest.
        /// </summary>
        [Input("mgmtAddress")]
        public Input<string>? MgmtAddress { get; set; }

        /// <summary>
        /// Specifies the method by which the management address is used in the vCMP guest. options : [`bridged`,`isolated`,`host-only`].
        /// </summary>
        [Input("mgmtNetwork")]
        public Input<string>? MgmtNetwork { get; set; }

        /// <summary>
        /// Specifies the gateway address for the `mgmt_address`. Can be set to `none` to remove the value from the configuration.
        /// </summary>
        [Input("mgmtRoute")]
        public Input<string>? MgmtRoute { get; set; }

        /// <summary>
        /// Specifies the minimum number of slots the guest must be assigned to in order to deploy.
        /// </summary>
        [Input("minNumberOfSlots")]
        public Input<int>? MinNumberOfSlots { get; set; }

        /// <summary>
        /// Name of the vCMP guest
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the number of slots for the system to use when creating the guest.
        /// </summary>
        [Input("numberOfSlots")]
        public Input<int>? NumberOfSlots { get; set; }

        /// <summary>
        /// Specifies the state of the vCMP guest on the system. options : [`configured`,`provisioned`,`deployed`].
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Virtual disk associated with vCMP guest.
        /// </summary>
        [Input("virtualDisk")]
        public Input<string>? VirtualDisk { get; set; }

        [Input("vlans")]
        private InputList<string>? _vlans;

        /// <summary>
        /// Specifies the list of VLANs the vCMP guest uses to communicate with other guests, the host, and with the external network. The naming format must be the combination of the partition + name. For example /Common/my-vlan
        /// </summary>
        public InputList<string> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<string>());
            set => _vlans = value;
        }

        public GuestState()
        {
        }
        public static new GuestState Empty => new GuestState();
    }
}
