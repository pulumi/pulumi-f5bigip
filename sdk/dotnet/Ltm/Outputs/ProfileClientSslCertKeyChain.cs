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
    public sealed class ProfileClientSslCertKeyChain
    {
        /// <summary>
        /// Specifies a cert name for use.
        /// </summary>
        public readonly string? Cert;
        /// <summary>
        /// Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
        /// </summary>
        public readonly string? Chain;
        /// <summary>
        /// Contains a key name
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// Specifies the name of the profile.Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-clientssl-profile`.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Key passphrase
        /// </summary>
        public readonly string? Passphrase;

        [OutputConstructor]
        private ProfileClientSslCertKeyChain(
            string? cert,

            string? chain,

            string? key,

            string? name,

            string? passphrase)
        {
            Cert = cert;
            Chain = chain;
            Key = key;
            Name = name;
            Passphrase = passphrase;
        }
    }
}
