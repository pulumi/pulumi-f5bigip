// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ssl.Inputs
{

    public sealed class GetWafEntityParameterUrlInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("method", required: true)]
        public Input<string> Method { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetWafEntityParameterUrlInputArgs()
        {
        }
        public static new GetWafEntityParameterUrlInputArgs Empty => new GetWafEntityParameterUrlInputArgs();
    }
}
