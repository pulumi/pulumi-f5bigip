// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Inputs
{

    public sealed class WafPolicyHostNameArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public WafPolicyHostNameArgs()
        {
        }
        public static new WafPolicyHostNameArgs Empty => new WafPolicyHostNameArgs();
    }
}
