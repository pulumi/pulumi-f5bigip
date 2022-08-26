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
    /// `f5bigip.ltm.ProfileFtp` Configures a custom profile_ftp.
    /// 
    /// Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
    /// 
    /// ## Example Usage
    /// ### For Bigip versions (14.x - 16.x)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sanjose_ftp_profile = new F5BigIP.Ltm.ProfileFtp("sanjose-ftp-profile", new()
    ///     {
    ///         AllowActiveMode = "enabled",
    ///         DefaultsFrom = "/Common/ftp",
    ///         Description = "test-tftp-profile",
    ///         EnforceTlssessionReuse = "enabled",
    ///         FtpsMode = "allow",
    ///         Name = "/Common/sanjose-ftp-profile",
    ///         Port = 2020,
    ///     });
    /// 
    /// });
    /// ```
    /// ### For Bigip versions (12.x - 13.x)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sanjose_ftp_profile = new F5BigIP.Ltm.ProfileFtp("sanjose-ftp-profile", new()
    ///     {
    ///         AllowFtps = "enabled",
    ///         DefaultsFrom = "/Common/ftp",
    ///         Description = "test-tftp-profile",
    ///         Name = "/Common/sanjose-ftp-profile",
    ///         Port = 2020,
    ///         TranslateExtended = "enabled",
    ///     });
    /// 
    /// });
    /// ```
    /// ## Common Arguments for all versions
    /// 
    /// * `security` - (Optional)Specifies, when checked (enabled), that the system inspects FTP traffic for security vulnerabilities using an FTP security profile. This option is available only on systems licensed for BIG-IP ASM.
    /// 
    /// * `port` - (Optional)Allows you to configure the FTP service to run on an alternate port. The default is 20.
    /// 
    /// * `log_profile` - (Optional)Configures the ALG log profile that controls logging
    /// 
    /// * `log_publisher` - (Optional)Configures the log publisher that handles events logging for this profile
    /// 
    /// *  `inherit_parent_profile` - (Optional)Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses FastL4 only.
    /// 
    /// * `description` - (Optional)User defined description for FTP profile
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/profileFtp:ProfileFtp")]
    public partial class ProfileFtp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled
        /// </summary>
        [Output("allowActiveMode")]
        public Output<string?> AllowActiveMode { get; private set; } = null!;

        /// <summary>
        /// Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.
        /// </summary>
        [Output("allowFtps")]
        public Output<string?> AllowFtps { get; private set; } = null!;

        /// <summary>
        /// The application service to which the object belongs.
        /// </summary>
        [Output("appService")]
        public Output<string?> AppService { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// User defined description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)
        /// </summary>
        [Output("enforceTlssessionReuse")]
        public Output<string?> EnforceTlssessionReuse { get; private set; } = null!;

        /// <summary>
        /// Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow
        /// </summary>
        [Output("ftpsMode")]
        public Output<string?> FtpsMode { get; private set; } = null!;

        /// <summary>
        /// Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
        /// FastL4 only.
        /// </summary>
        [Output("inheritParentProfile")]
        public Output<string?> InheritParentProfile { get; private set; } = null!;

        /// <summary>
        /// inherent vlan list
        /// </summary>
        [Output("inheritVlanList")]
        public Output<string?> InheritVlanList { get; private set; } = null!;

        /// <summary>
        /// Configures the ALG log profile that controls logging
        /// </summary>
        [Output("logProfile")]
        public Output<string> LogProfile { get; private set; } = null!;

        /// <summary>
        /// Configures the log publisher that handles events logging for this profile
        /// </summary>
        [Output("logPublisher")]
        public Output<string> LogPublisher { get; private set; } = null!;

        /// <summary>
        /// Name of the profile_ftp
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Output("partition")]
        public Output<string> Partition { get; private set; } = null!;

        /// <summary>
        /// Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
        /// system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
        /// </summary>
        [Output("security")]
        public Output<string> Security { get; private set; } = null!;

        /// <summary>
        /// Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled).
        /// </summary>
        [Output("translateExtended")]
        public Output<string?> TranslateExtended { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileFtp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileFtp(string name, ProfileFtpArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFtp:ProfileFtp", name, args ?? new ProfileFtpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileFtp(string name, Input<string> id, ProfileFtpState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileFtp:ProfileFtp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileFtp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileFtp Get(string name, Input<string> id, ProfileFtpState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileFtp(name, id, state, options);
        }
    }

    public sealed class ProfileFtpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled
        /// </summary>
        [Input("allowActiveMode")]
        public Input<string>? AllowActiveMode { get; set; }

        /// <summary>
        /// Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.
        /// </summary>
        [Input("allowFtps")]
        public Input<string>? AllowFtps { get; set; }

        /// <summary>
        /// The application service to which the object belongs.
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)
        /// </summary>
        [Input("enforceTlssessionReuse")]
        public Input<string>? EnforceTlssessionReuse { get; set; }

        /// <summary>
        /// Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow
        /// </summary>
        [Input("ftpsMode")]
        public Input<string>? FtpsMode { get; set; }

        /// <summary>
        /// Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
        /// FastL4 only.
        /// </summary>
        [Input("inheritParentProfile")]
        public Input<string>? InheritParentProfile { get; set; }

        /// <summary>
        /// inherent vlan list
        /// </summary>
        [Input("inheritVlanList")]
        public Input<string>? InheritVlanList { get; set; }

        /// <summary>
        /// Configures the ALG log profile that controls logging
        /// </summary>
        [Input("logProfile")]
        public Input<string>? LogProfile { get; set; }

        /// <summary>
        /// Configures the log publisher that handles events logging for this profile
        /// </summary>
        [Input("logPublisher")]
        public Input<string>? LogPublisher { get; set; }

        /// <summary>
        /// Name of the profile_ftp
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
        /// system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
        /// </summary>
        [Input("security")]
        public Input<string>? Security { get; set; }

        /// <summary>
        /// Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled).
        /// </summary>
        [Input("translateExtended")]
        public Input<string>? TranslateExtended { get; set; }

        public ProfileFtpArgs()
        {
        }
        public static new ProfileFtpArgs Empty => new ProfileFtpArgs();
    }

    public sealed class ProfileFtpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled
        /// </summary>
        [Input("allowActiveMode")]
        public Input<string>? AllowActiveMode { get; set; }

        /// <summary>
        /// Allow explicit FTPS negotiation. The default is disabled.When enabled (selected), that the system allows explicit FTPS negotiation for SSL or TLS.
        /// </summary>
        [Input("allowFtps")]
        public Input<string>? AllowFtps { get; set; }

        /// <summary>
        /// The application service to which the object belongs.
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default value is unchecked (disabled)
        /// </summary>
        [Input("enforceTlssessionReuse")]
        public Input<string>? EnforceTlssessionReuse { get; set; }

        /// <summary>
        /// Specifies if you want to Disallow, Allow, or Require FTPS mode. The default is Disallow
        /// </summary>
        [Input("ftpsMode")]
        public Input<string>? FtpsMode { get; set; }

        /// <summary>
        /// Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
        /// FastL4 only.
        /// </summary>
        [Input("inheritParentProfile")]
        public Input<string>? InheritParentProfile { get; set; }

        /// <summary>
        /// inherent vlan list
        /// </summary>
        [Input("inheritVlanList")]
        public Input<string>? InheritVlanList { get; set; }

        /// <summary>
        /// Configures the ALG log profile that controls logging
        /// </summary>
        [Input("logProfile")]
        public Input<string>? LogProfile { get; set; }

        /// <summary>
        /// Configures the log publisher that handles events logging for this profile
        /// </summary>
        [Input("logPublisher")]
        public Input<string>? LogPublisher { get; set; }

        /// <summary>
        /// Name of the profile_ftp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
        /// system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
        /// </summary>
        [Input("security")]
        public Input<string>? Security { get; set; }

        /// <summary>
        /// Specifies, when selected (enabled), that the system uses ensures compatibility between IP version 4 and IP version 6 clients and servers when using the FTP protocol. The default is selected (enabled).
        /// </summary>
        [Input("translateExtended")]
        public Input<string>? TranslateExtended { get; set; }

        public ProfileFtpState()
        {
        }
        public static new ProfileFtpState Empty => new ProfileFtpState();
    }
}
