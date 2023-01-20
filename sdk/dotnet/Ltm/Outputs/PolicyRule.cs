// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Outputs
{

    [OutputType]
    public sealed class PolicyRule
    {
        /// <summary>
        /// Block type. See action block for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyRuleAction> Actions;
        /// <summary>
        /// Block type. See condition block for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.PolicyRuleCondition> Conditions;
        /// <summary>
        /// Name of Rule to be applied in policy.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private PolicyRule(
            ImmutableArray<Outputs.PolicyRuleAction> actions,

            ImmutableArray<Outputs.PolicyRuleCondition> conditions,

            string name)
        {
            Actions = actions;
            Conditions = conditions;
            Name = name;
        }
    }
}
