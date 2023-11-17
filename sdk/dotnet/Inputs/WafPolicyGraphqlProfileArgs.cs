// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Inputs
{

    public sealed class WafPolicyGraphqlProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies when checked (enabled) that you want attack signatures and threat campaigns to be detected on this GraphQL profile and possibly override the security policy settings of an attack signature or threat campaign specifically for this GraphQL profile. After you enable this setting, the system displays a list of attack signatures and and threat campaigns. The default is enabled.
        /// </summary>
        [Input("attackSignaturesCheck")]
        public Input<bool>? AttackSignaturesCheck { get; set; }

        [Input("defenseAttributes")]
        private InputList<Inputs.WafPolicyGraphqlProfileDefenseAttributeArgs>? _defenseAttributes;

        /// <summary>
        /// block settings for GraphQl policy.See defense attributes below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyGraphqlProfileDefenseAttributeArgs> DefenseAttributes
        {
            get => _defenseAttributes ?? (_defenseAttributes = new InputList<Inputs.WafPolicyGraphqlProfileDefenseAttributeArgs>());
            set => _defenseAttributes = value;
        }

        /// <summary>
        /// Specifies when checked (enabled) that the system enforces the security policy settings of a meta character for the GraphQL profile. After you enable this setting, the system displays a list of meta characters. The default is enabled.
        /// </summary>
        [Input("metacharElementcheck")]
        public Input<bool>? MetacharElementcheck { get; set; }

        /// <summary>
        /// The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public WafPolicyGraphqlProfileArgs()
        {
        }
        public static new WafPolicyGraphqlProfileArgs Empty => new WafPolicyGraphqlProfileArgs();
    }
}
