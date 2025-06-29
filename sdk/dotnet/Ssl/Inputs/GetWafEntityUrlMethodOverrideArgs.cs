// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl.Inputs
{

    public sealed class GetWafEntityUrlMethodOverrideInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies that the system allows or disallows a method for this URL
        /// </summary>
        [Input("allow", required: true)]
        public Input<bool> Allow { get; set; } = null!;

        /// <summary>
        /// Specifies an HTTP method.
        /// </summary>
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        public GetWafEntityUrlMethodOverrideInputArgs()
        {
        }
        public static new GetWafEntityUrlMethodOverrideInputArgs Empty => new GetWafEntityUrlMethodOverrideInputArgs();
    }
}
