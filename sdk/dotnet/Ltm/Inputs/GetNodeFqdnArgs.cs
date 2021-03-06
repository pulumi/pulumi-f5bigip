// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class GetNodeFqdnArgs : Pulumi.InvokeArgs
    {
        [Input("addressFamily")]
        public string? AddressFamily { get; set; }

        [Input("autopopulate", required: true)]
        public string Autopopulate { get; set; } = null!;

        [Input("downinterval", required: true)]
        public int Downinterval { get; set; }

        [Input("interval", required: true)]
        public string Interval { get; set; } = null!;

        [Input("name")]
        public string? Name { get; set; }

        public GetNodeFqdnArgs()
        {
        }
    }
}
