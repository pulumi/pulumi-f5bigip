// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys.Outputs
{

    [OutputType]
    public sealed class IAppMetadata
    {
        /// <summary>
        /// Name of origin
        /// </summary>
        public readonly string? Persists;
        /// <summary>
        /// Name of origin
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private IAppMetadata(
            string? persists,

            string? value)
        {
            Persists = persists;
            Value = value;
        }
    }
}
