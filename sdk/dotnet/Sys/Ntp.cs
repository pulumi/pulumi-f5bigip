// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys
{
    /// <summary>
    /// `f5bigip.sys.Ntp` provides details about a specific bigip
    /// 
    /// This resource is helpful when configuring NTP server on the BIG-IP.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/sys_ntp.html.markdown.
    /// </summary>
    public partial class Ntp : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the ntp Servers
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
        /// </summary>
        [Output("servers")]
        public Output<ImmutableArray<string>> Servers { get; private set; } = null!;

        /// <summary>
        /// Specifies the time zone that you want to use for the system time.
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a Ntp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ntp(string name, NtpArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:sys/ntp:Ntp", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Ntp(string name, Input<string> id, NtpState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:sys/ntp:Ntp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ntp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ntp Get(string name, Input<string> id, NtpState? state = null, CustomResourceOptions? options = null)
        {
            return new Ntp(name, id, state, options);
        }
    }

    public sealed class NtpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the ntp Servers
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("servers")]
        private InputList<string>? _servers;

        /// <summary>
        /// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
        /// </summary>
        public InputList<string> Servers
        {
            get => _servers ?? (_servers = new InputList<string>());
            set => _servers = value;
        }

        /// <summary>
        /// Specifies the time zone that you want to use for the system time.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public NtpArgs()
        {
        }
    }

    public sealed class NtpState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the ntp Servers
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("servers")]
        private InputList<string>? _servers;

        /// <summary>
        /// Adds NTP servers to or deletes NTP servers from the BIG-IP system.
        /// </summary>
        public InputList<string> Servers
        {
            get => _servers ?? (_servers = new InputList<string>());
            set => _servers = value;
        }

        /// <summary>
        /// Specifies the time zone that you want to use for the system time.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public NtpState()
        {
        }
    }
}
