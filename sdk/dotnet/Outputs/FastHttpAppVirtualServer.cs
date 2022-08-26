// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Outputs
{

    [OutputType]
    public sealed class FastHttpAppVirtualServer
    {
        /// <summary>
        /// IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// -(Optional , `int`) Port number to used for accessing virtual server/application
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private FastHttpAppVirtualServer(
            string ip,

            int port)
        {
            Ip = ip;
            Port = port;
        }
    }
}
