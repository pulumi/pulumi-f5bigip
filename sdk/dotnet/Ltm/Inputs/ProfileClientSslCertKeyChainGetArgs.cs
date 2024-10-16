// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm.Inputs
{

    public sealed class ProfileClientSslCertKeyChainGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a cert name for use.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
        /// </summary>
        [Input("chain")]
        public Input<string>? Chain { get; set; }

        /// <summary>
        /// Contains a key name
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Specifies the name of the profile.Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-clientssl-profile`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("passphrase")]
        private Input<string>? _passphrase;

        /// <summary>
        /// Key passphrase
        /// </summary>
        public Input<string>? Passphrase
        {
            get => _passphrase;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _passphrase = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProfileClientSslCertKeyChainGetArgs()
        {
        }
        public static new ProfileClientSslCertKeyChainGetArgs Empty => new ProfileClientSslCertKeyChainGetArgs();
    }
}