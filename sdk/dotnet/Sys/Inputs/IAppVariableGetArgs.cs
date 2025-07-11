// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys.Inputs
{

    public sealed class IAppVariableGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of origin
        /// </summary>
        [Input("encrypted")]
        public Input<string>? Encrypted { get; set; }

        /// <summary>
        /// Name of the iApp.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of origin
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IAppVariableGetArgs()
        {
        }
        public static new IAppVariableGetArgs Empty => new IAppVariableGetArgs();
    }
}
