// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class SnatOriginGetArgs : Pulumi.ResourceArgs
    {
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Name of the snat
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SnatOriginGetArgs()
        {
        }
    }
}
