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
    public sealed class ProfileRewriteUriRulesClient
    {
        /// <summary>
        /// Host part of the uri, e.g. `www.foo.com`.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Path part of the uri, must always end with `/`. Default value is: `/`
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Port part of the uri. Default value is: `none`
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// Scheme part of the uri, e.g. `https`, `ftp`.
        /// </summary>
        public readonly string Scheme;

        [OutputConstructor]
        private ProfileRewriteUriRulesClient(
            string host,

            string? path,

            string? port,

            string scheme)
        {
            Host = host;
            Path = path;
            Port = port;
            Scheme = scheme;
        }
    }
}