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
    public sealed class GetDataGroupRecordResult
    {
        public readonly string? Data;
        /// <summary>
        /// Name of the datagroup
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetDataGroupRecordResult(
            string? data,

            string name)
        {
            Data = data;
            Name = name;
        }
    }
}
