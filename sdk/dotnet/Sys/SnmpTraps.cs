// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5bigip.Sys
{
    /// <summary>
    /// `f5bigip.sys.SnmpTraps` provides details bout how to enable snmp_traps resource on BIG-IP
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/sys_snmp_traps.html.markdown.
    /// </summary>
    public partial class SnmpTraps : Pulumi.CustomResource
    {
        /// <summary>
        /// Encrypted password
        /// </summary>
        [Output("authPasswordencrypted")]
        public Output<string?> AuthPasswordencrypted { get; private set; } = null!;

        /// <summary>
        /// Specifies the protocol used to authenticate the user.
        /// </summary>
        [Output("authProtocol")]
        public Output<string?> AuthProtocol { get; private set; } = null!;

        /// <summary>
        /// Specifies the community string used for this trap.
        /// </summary>
        [Output("community")]
        public Output<string?> Community { get; private set; } = null!;

        /// <summary>
        /// The port that the trap will be sent to.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies the authoritative security engine for SNMPv3.
        /// </summary>
        [Output("engineId")]
        public Output<string?> EngineId { get; private set; } = null!;

        /// <summary>
        /// The host the trap will be sent to.
        /// </summary>
        [Output("host")]
        public Output<string?> Host { get; private set; } = null!;

        /// <summary>
        /// Name of the snmp trap.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// User defined description.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        /// </summary>
        [Output("privacyPassword")]
        public Output<string?> PrivacyPassword { get; private set; } = null!;

        /// <summary>
        /// Specifies the encrypted password used to encrypt traffic.
        /// </summary>
        [Output("privacyPasswordEncrypted")]
        public Output<string?> PrivacyPasswordEncrypted { get; private set; } = null!;

        /// <summary>
        /// Specifies the protocol used to encrypt traffic.
        /// </summary>
        [Output("privacyProtocol")]
        public Output<string?> PrivacyProtocol { get; private set; } = null!;

        /// <summary>
        /// Specifies whether or not traffic is encrypted and whether or not authentication is required.
        /// </summary>
        [Output("securityLevel")]
        public Output<string?> SecurityLevel { get; private set; } = null!;

        /// <summary>
        /// Security name used in conjunction with SNMPv3.
        /// </summary>
        [Output("securityName")]
        public Output<string?> SecurityName { get; private set; } = null!;

        /// <summary>
        /// SNMP version used for sending the trap.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a SnmpTraps resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SnmpTraps(string name, SnmpTrapsArgs? args = null, CustomResourceOptions? options = null)
            : base("f5bigip:sys/snmpTraps:SnmpTraps", name, args, MakeResourceOptions(options, ""))
        {
        }

        private SnmpTraps(string name, Input<string> id, SnmpTrapsState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:sys/snmpTraps:SnmpTraps", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SnmpTraps resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SnmpTraps Get(string name, Input<string> id, SnmpTrapsState? state = null, CustomResourceOptions? options = null)
        {
            return new SnmpTraps(name, id, state, options);
        }
    }

    public sealed class SnmpTrapsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encrypted password
        /// </summary>
        [Input("authPasswordencrypted")]
        public Input<string>? AuthPasswordencrypted { get; set; }

        /// <summary>
        /// Specifies the protocol used to authenticate the user.
        /// </summary>
        [Input("authProtocol")]
        public Input<string>? AuthProtocol { get; set; }

        /// <summary>
        /// Specifies the community string used for this trap.
        /// </summary>
        [Input("community")]
        public Input<string>? Community { get; set; }

        /// <summary>
        /// The port that the trap will be sent to.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the authoritative security engine for SNMPv3.
        /// </summary>
        [Input("engineId")]
        public Input<string>? EngineId { get; set; }

        /// <summary>
        /// The host the trap will be sent to.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Name of the snmp trap.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User defined description.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        /// </summary>
        [Input("privacyPassword")]
        public Input<string>? PrivacyPassword { get; set; }

        /// <summary>
        /// Specifies the encrypted password used to encrypt traffic.
        /// </summary>
        [Input("privacyPasswordEncrypted")]
        public Input<string>? PrivacyPasswordEncrypted { get; set; }

        /// <summary>
        /// Specifies the protocol used to encrypt traffic.
        /// </summary>
        [Input("privacyProtocol")]
        public Input<string>? PrivacyProtocol { get; set; }

        /// <summary>
        /// Specifies whether or not traffic is encrypted and whether or not authentication is required.
        /// </summary>
        [Input("securityLevel")]
        public Input<string>? SecurityLevel { get; set; }

        /// <summary>
        /// Security name used in conjunction with SNMPv3.
        /// </summary>
        [Input("securityName")]
        public Input<string>? SecurityName { get; set; }

        /// <summary>
        /// SNMP version used for sending the trap.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public SnmpTrapsArgs()
        {
        }
    }

    public sealed class SnmpTrapsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encrypted password
        /// </summary>
        [Input("authPasswordencrypted")]
        public Input<string>? AuthPasswordencrypted { get; set; }

        /// <summary>
        /// Specifies the protocol used to authenticate the user.
        /// </summary>
        [Input("authProtocol")]
        public Input<string>? AuthProtocol { get; set; }

        /// <summary>
        /// Specifies the community string used for this trap.
        /// </summary>
        [Input("community")]
        public Input<string>? Community { get; set; }

        /// <summary>
        /// The port that the trap will be sent to.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies the authoritative security engine for SNMPv3.
        /// </summary>
        [Input("engineId")]
        public Input<string>? EngineId { get; set; }

        /// <summary>
        /// The host the trap will be sent to.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Name of the snmp trap.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User defined description.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        /// </summary>
        [Input("privacyPassword")]
        public Input<string>? PrivacyPassword { get; set; }

        /// <summary>
        /// Specifies the encrypted password used to encrypt traffic.
        /// </summary>
        [Input("privacyPasswordEncrypted")]
        public Input<string>? PrivacyPasswordEncrypted { get; set; }

        /// <summary>
        /// Specifies the protocol used to encrypt traffic.
        /// </summary>
        [Input("privacyProtocol")]
        public Input<string>? PrivacyProtocol { get; set; }

        /// <summary>
        /// Specifies whether or not traffic is encrypted and whether or not authentication is required.
        /// </summary>
        [Input("securityLevel")]
        public Input<string>? SecurityLevel { get; set; }

        /// <summary>
        /// Security name used in conjunction with SNMPv3.
        /// </summary>
        [Input("securityName")]
        public Input<string>? SecurityName { get; set; }

        /// <summary>
        /// SNMP version used for sending the trap.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public SnmpTrapsState()
        {
        }
    }
}
