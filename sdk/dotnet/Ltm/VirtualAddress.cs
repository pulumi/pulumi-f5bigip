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
    /// `f5bigip.ltm.VirtualAddress` Configures Virtual Server
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
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
    ///         var vsVa = new F5BigIP.Ltm.VirtualAddress("vsVa", new F5BigIP.Ltm.VirtualAddressArgs
    ///         {
    ///             AdvertizeRoute = "true",
    ///             Name = "/Common/vs_va",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/virtualAddress:VirtualAddress")]
    public partial class VirtualAddress : Pulumi.CustomResource
    {
        /// <summary>
        /// Enabled dynamic routing of the address
        /// </summary>
        [Output("advertizeRoute")]
        public Output<string?> AdvertizeRoute { get; private set; } = null!;

        /// <summary>
        /// Enable or disable ARP for the virtual address
        /// </summary>
        [Output("arp")]
        public Output<bool?> Arp { get; private set; } = null!;

        /// <summary>
        /// Automatically delete the virtual address with the virtual server
        /// </summary>
        [Output("autoDelete")]
        public Output<bool?> AutoDelete { get; private set; } = null!;

        /// <summary>
        /// Max number of connections for virtual address
        /// </summary>
        [Output("connLimit")]
        public Output<int?> ConnLimit { get; private set; } = null!;

        /// <summary>
        /// Enable or disable the virtual address
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Enable/Disable ICMP response to the virtual address
        /// </summary>
        [Output("icmpEcho")]
        public Output<bool?> IcmpEcho { get; private set; } = null!;

        /// <summary>
        /// Name of the virtual address
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specify the partition and traffic group
        /// </summary>
        [Output("trafficGroup")]
        public Output<string?> TrafficGroup { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualAddress(string name, VirtualAddressArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/virtualAddress:VirtualAddress", name, args ?? new VirtualAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualAddress(string name, Input<string> id, VirtualAddressState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/virtualAddress:VirtualAddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualAddress Get(string name, Input<string> id, VirtualAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new VirtualAddress(name, id, state, options);
        }
    }

    public sealed class VirtualAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enabled dynamic routing of the address
        /// </summary>
        [Input("advertizeRoute")]
        public Input<string>? AdvertizeRoute { get; set; }

        /// <summary>
        /// Enable or disable ARP for the virtual address
        /// </summary>
        [Input("arp")]
        public Input<bool>? Arp { get; set; }

        /// <summary>
        /// Automatically delete the virtual address with the virtual server
        /// </summary>
        [Input("autoDelete")]
        public Input<bool>? AutoDelete { get; set; }

        /// <summary>
        /// Max number of connections for virtual address
        /// </summary>
        [Input("connLimit")]
        public Input<int>? ConnLimit { get; set; }

        /// <summary>
        /// Enable or disable the virtual address
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Enable/Disable ICMP response to the virtual address
        /// </summary>
        [Input("icmpEcho")]
        public Input<bool>? IcmpEcho { get; set; }

        /// <summary>
        /// Name of the virtual address
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specify the partition and traffic group
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        public VirtualAddressArgs()
        {
        }
    }

    public sealed class VirtualAddressState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enabled dynamic routing of the address
        /// </summary>
        [Input("advertizeRoute")]
        public Input<string>? AdvertizeRoute { get; set; }

        /// <summary>
        /// Enable or disable ARP for the virtual address
        /// </summary>
        [Input("arp")]
        public Input<bool>? Arp { get; set; }

        /// <summary>
        /// Automatically delete the virtual address with the virtual server
        /// </summary>
        [Input("autoDelete")]
        public Input<bool>? AutoDelete { get; set; }

        /// <summary>
        /// Max number of connections for virtual address
        /// </summary>
        [Input("connLimit")]
        public Input<int>? ConnLimit { get; set; }

        /// <summary>
        /// Enable or disable the virtual address
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Enable/Disable ICMP response to the virtual address
        /// </summary>
        [Input("icmpEcho")]
        public Input<bool>? IcmpEcho { get; set; }

        /// <summary>
        /// Name of the virtual address
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specify the partition and traffic group
        /// </summary>
        [Input("trafficGroup")]
        public Input<string>? TrafficGroup { get; set; }

        public VirtualAddressState()
        {
        }
    }
}
