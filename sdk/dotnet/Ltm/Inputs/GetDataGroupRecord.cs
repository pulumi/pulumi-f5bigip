// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class GetDataGroupRecordArgs : global::Pulumi.InvokeArgs
    {
        [Input("data")]
        public string? Data { get; set; }

        /// <summary>
        /// Name of the datagroup
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDataGroupRecordArgs()
        {
        }
        public static new GetDataGroupRecordArgs Empty => new GetDataGroupRecordArgs();
    }
}
