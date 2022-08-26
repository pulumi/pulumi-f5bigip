// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Inputs
{

    public sealed class WafPolicyPolicyBuilderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// learning mode setting for policy-builder, possible options: [`automatic`,`disabled`, `manual`]
        /// </summary>
        [Input("learningMode")]
        public Input<string>? LearningMode { get; set; }

        public WafPolicyPolicyBuilderArgs()
        {
        }
        public static new WafPolicyPolicyBuilderArgs Empty => new WafPolicyPolicyBuilderArgs();
    }
}
