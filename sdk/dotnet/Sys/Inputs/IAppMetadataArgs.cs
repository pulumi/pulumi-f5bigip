// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys.Inputs
{

    public sealed class IAppMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of origin
        /// </summary>
        [Input("persists")]
        public Input<string>? Persists { get; set; }

        /// <summary>
        /// Name of origin
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IAppMetadataArgs()
        {
        }
        public static new IAppMetadataArgs Empty => new IAppMetadataArgs();
    }
}