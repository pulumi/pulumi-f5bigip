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
    public sealed class SnatOrigin
    {
        public readonly string? AppService;
        /// <summary>
        /// Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private SnatOrigin(
            string? appService,

            string? name)
        {
            AppService = appService;
            Name = name;
        }
    }
}
