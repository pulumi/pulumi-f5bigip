// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class PolicyRuleArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.PolicyRuleActionArgs>? _actions;
        public InputList<Inputs.PolicyRuleActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.PolicyRuleActionArgs>());
            set => _actions = value;
        }

        [Input("conditions")]
        private InputList<Inputs.PolicyRuleConditionArgs>? _conditions;
        public InputList<Inputs.PolicyRuleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.PolicyRuleConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public PolicyRuleArgs()
        {
        }
    }
}
