// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    [F5BigIPResourceType("f5bigip:ltm/profileBotDefense:ProfileBotDefense")]
    public partial class ProfileBotDefense : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log`
        /// profile
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// User defined description for Bot Defense profile
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Select the enforcement mode, possible values are `transparent` and `blocking`.
        /// </summary>
        [Output("enforcementMode")]
        public Output<string> EnforcementMode { get; private set; } = null!;

        /// <summary>
        /// Name of the Bot Defense profile
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Profile templates specify Mitigation and Verification Settings default values. possible ptions `balanced`,`relaxed` and
        /// `strict`
        /// </summary>
        [Output("template")]
        public Output<string> Template { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileBotDefense resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileBotDefense(string name, ProfileBotDefenseArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileBotDefense:ProfileBotDefense", name, args ?? new ProfileBotDefenseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileBotDefense(string name, Input<string> id, ProfileBotDefenseState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileBotDefense:ProfileBotDefense", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileBotDefense resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileBotDefense Get(string name, Input<string> id, ProfileBotDefenseState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileBotDefense(name, id, state, options);
        }
    }

    public sealed class ProfileBotDefenseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log`
        /// profile
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// User defined description for Bot Defense profile
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Select the enforcement mode, possible values are `transparent` and `blocking`.
        /// </summary>
        [Input("enforcementMode")]
        public Input<string>? EnforcementMode { get; set; }

        /// <summary>
        /// Name of the Bot Defense profile
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Profile templates specify Mitigation and Verification Settings default values. possible ptions `balanced`,`relaxed` and
        /// `strict`
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public ProfileBotDefenseArgs()
        {
        }
        public static new ProfileBotDefenseArgs Empty => new ProfileBotDefenseArgs();
    }

    public sealed class ProfileBotDefenseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `request-log`
        /// profile
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// User defined description for Bot Defense profile
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Select the enforcement mode, possible values are `transparent` and `blocking`.
        /// </summary>
        [Input("enforcementMode")]
        public Input<string>? EnforcementMode { get; set; }

        /// <summary>
        /// Name of the Bot Defense profile
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Profile templates specify Mitigation and Verification Settings default values. possible ptions `balanced`,`relaxed` and
        /// `strict`
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public ProfileBotDefenseState()
        {
        }
        public static new ProfileBotDefenseState Empty => new ProfileBotDefenseState();
    }
}
