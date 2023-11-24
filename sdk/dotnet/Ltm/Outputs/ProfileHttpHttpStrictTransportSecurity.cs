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
    public sealed class ProfileHttpHttpStrictTransportSecurity
    {
        /// <summary>
        /// The Include Subdomains setting applies the HSTS policy to the HSTS host and its subdomains. The default is "enabled". If no string is specified during Create, then default value will be assigned by BigIp. If include_subdomains is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "enabled" explicitly.
        /// </summary>
        public readonly string? IncludeSubdomains;
        /// <summary>
        /// The Maximum Age value specifies the length of time, in seconds, that HSTS functionality requests that clients only use HTTPS to connect to the current host and any subdomains of the current host's domain name.  The default is 16070400 seconds. If no value is specified during Create, then default value will be assigned by BigIp. If maximum_age is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value , we need to pass 16070400 explicitly.
        /// </summary>
        public readonly int? MaximumAge;
        /// <summary>
        /// The Mode setting enables and disables HSTS functionality within the HTTP profile. The default is "disabled". If no string is specified during Create, then default value will be assigned by BigIp. If mode is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "disabled" explicitly.
        /// </summary>
        public readonly string? Mode;
        /// <summary>
        /// An HSTS preload list is a list of domains built into a web browser. When you enable the Preload setting, the domain for the web site that this HTTP profile is associated with is submitted for inclusion in the browser's preload list. The default is "disabled". If no string is specified during Create, then default value will be assigned by BigIp. If preload is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "disabled" explicitly.
        /// </summary>
        public readonly string? Preload;

        [OutputConstructor]
        private ProfileHttpHttpStrictTransportSecurity(
            string? includeSubdomains,

            int? maximumAge,

            string? mode,

            string? preload)
        {
            IncludeSubdomains = includeSubdomains;
            MaximumAge = maximumAge;
            Mode = mode;
            Preload = preload;
        }
    }
}
