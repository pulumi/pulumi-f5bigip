// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Sys.Inputs
{

    public sealed class IAppListGetArgs : Pulumi.ResourceArgs
    {
        [Input("encrypted")]
        public Input<string>? Encrypted { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public IAppListGetArgs()
        {
        }
    }
}