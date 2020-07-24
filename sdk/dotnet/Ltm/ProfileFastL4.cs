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
    /// `f5bigip.ltm.ProfileFastL4` Configures a custom profile_fastl4 for use by health checks.
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
    ///         var profileFastl4 = new F5BigIP.Ltm.ProfileFastL4("profileFastl4", new F5BigIP.Ltm.ProfileFastL4Args
    ///         {
    ///             ClientTimeout = 40,
    ///             DefaultsFrom = "/Common/fastL4",
    ///             ExplicitflowMigration = "enabled",
    ///             HardwareSyncookie = "enabled",
    ///             IdleTimeout = "200",
    ///             IptosToclient = "pass-through",
    ///             IptosToserver = "pass-through",
    ///             KeepaliveInterval = "disabled",
    ///             Name = "/Common/sjfastl4profile",
    ///             Partition = "Common",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ProfileFastL4 : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
        /// </summary>
        [Output("clientTimeout")]
        public Output<int> ClientTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
        /// </summary>
        [Output("explicitflowMigration")]
        public Output<string> ExplicitflowMigration { get; private set; } = null!;

        /// <summary>
        /// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
        /// </summary>
        [Output("hardwareSyncookie")]
        public Output<string> HardwareSyncookie { get; private set; } = null!;

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Output("idleTimeout")]
        public Output<string> IdleTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Output("iptosToclient")]
        public Output<string> IptosToclient { get; private set; } = null!;

        /// <summary>
        /// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Output("iptosToserver")]
        public Output<string> IptosToserver { get; private set; } = null!;

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
        /// </summary>
        [Output("keepaliveInterval")]
        public Output<string> KeepaliveInterval { get; private set; } = null!;

        /// <summary>
        /// Name of the profile_fastl4
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Output("partition")]
        public Output<string> Partition { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileFastL4 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileFastL4(string name, ProfileFastL4Args args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFastL4:ProfileFastL4", name, args ?? new ProfileFastL4Args(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileFastL4(string name, Input<string> id, ProfileFastL4State? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFastL4:ProfileFastL4", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileFastL4 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileFastL4 Get(string name, Input<string> id, ProfileFastL4State? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileFastL4(name, id, state, options);
        }
    }

    public sealed class ProfileFastL4Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
        /// </summary>
        [Input("clientTimeout")]
        public Input<int>? ClientTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
        /// </summary>
        [Input("explicitflowMigration")]
        public Input<string>? ExplicitflowMigration { get; set; }

        /// <summary>
        /// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
        /// </summary>
        [Input("hardwareSyncookie")]
        public Input<string>? HardwareSyncookie { get; set; }

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<string>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToclient")]
        public Input<string>? IptosToclient { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToserver")]
        public Input<string>? IptosToserver { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<string>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Name of the profile_fastl4
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        public ProfileFastL4Args()
        {
        }
    }

    public sealed class ProfileFastL4State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
        /// </summary>
        [Input("clientTimeout")]
        public Input<int>? ClientTimeout { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
        /// </summary>
        [Input("explicitflowMigration")]
        public Input<string>? ExplicitflowMigration { get; set; }

        /// <summary>
        /// Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
        /// </summary>
        [Input("hardwareSyncookie")]
        public Input<string>? HardwareSyncookie { get; set; }

        /// <summary>
        /// Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        /// </summary>
        [Input("idleTimeout")]
        public Input<string>? IdleTimeout { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToclient")]
        public Input<string>? IptosToclient { get; set; }

        /// <summary>
        /// Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
        /// </summary>
        [Input("iptosToserver")]
        public Input<string>? IptosToserver { get; set; }

        /// <summary>
        /// Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
        /// </summary>
        [Input("keepaliveInterval")]
        public Input<string>? KeepaliveInterval { get; set; }

        /// <summary>
        /// Name of the profile_fastl4
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        public ProfileFastL4State()
        {
        }
    }
}
