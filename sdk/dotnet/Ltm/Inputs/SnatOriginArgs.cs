// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class SnatOriginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// app service
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SnatOriginArgs()
        {
        }
        public static new SnatOriginArgs Empty => new SnatOriginArgs();
    }
}
