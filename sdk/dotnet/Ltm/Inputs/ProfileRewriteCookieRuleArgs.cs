// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class ProfileRewriteCookieRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientDomain", required: true)]
        public Input<string> ClientDomain { get; set; } = null!;

        [Input("clientPath", required: true)]
        public Input<string> ClientPath { get; set; } = null!;

        /// <summary>
        /// Name of the cookie rewrite rule.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        [Input("serverDomain", required: true)]
        public Input<string> ServerDomain { get; set; } = null!;

        [Input("serverPath", required: true)]
        public Input<string> ServerPath { get; set; } = null!;

        public ProfileRewriteCookieRuleArgs()
        {
        }
        public static new ProfileRewriteCookieRuleArgs Empty => new ProfileRewriteCookieRuleArgs();
    }
}
