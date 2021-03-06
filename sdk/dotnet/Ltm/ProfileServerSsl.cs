// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.ProfileServerSsl` Manages server SSL profiles on a BIG-IP
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var test_ServerSsl = new F5BigIP.Ltm.ProfileServerSsl("test-ServerSsl", new F5BigIP.Ltm.ProfileServerSslArgs
    ///         {
    ///             Authenticate = "always",
    ///             Ciphers = "DEFAULT",
    ///             DefaultsFrom = "/Common/serverssl",
    ///             Name = "/Common/test-ServerSsl",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/profileServerSsl:ProfileServerSsl")]
    public partial class ProfileServerSsl : Pulumi.CustomResource
    {
        /// <summary>
        /// Alert time out
        /// </summary>
        [Output("alertTimeout")]
        public Output<string> AlertTimeout { get; private set; } = null!;

        /// <summary>
        /// Server authentication once / always (default is once).
        /// </summary>
        [Output("authenticate")]
        public Output<string> Authenticate { get; private set; } = null!;

        /// <summary>
        /// Client certificate chain traversal depth. Default 9.
        /// </summary>
        [Output("authenticateDepth")]
        public Output<int> AuthenticateDepth { get; private set; } = null!;

        /// <summary>
        /// Client certificate file path. Default None.
        /// </summary>
        [Output("caFile")]
        public Output<string> CaFile { get; private set; } = null!;

        /// <summary>
        /// Cache size (sessions).
        /// </summary>
        [Output("cacheSize")]
        public Output<int> CacheSize { get; private set; } = null!;

        /// <summary>
        /// Cache time out
        /// </summary>
        [Output("cacheTimeout")]
        public Output<int> CacheTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the certificate that the system uses for server-side SSL processing.
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// Specifies the certificates-key chain to associate with the SSL profile
        /// </summary>
        [Output("chain")]
        public Output<string> Chain { get; private set; } = null!;

        /// <summary>
        /// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        /// </summary>
        [Output("ciphers")]
        public Output<string> Ciphers { get; private set; } = null!;

        /// <summary>
        /// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Response if the cert is expired (drop / ignore).
        /// </summary>
        [Output("expireCertResponseControl")]
        public Output<string> ExpireCertResponseControl { get; private set; } = null!;

        /// <summary>
        /// full path of the profile
        /// </summary>
        [Output("fullPath")]
        public Output<string> FullPath { get; private set; } = null!;

        /// <summary>
        /// generation
        /// </summary>
        [Output("generation")]
        public Output<int> Generation { get; private set; } = null!;

        /// <summary>
        /// Generic alerts enabled / disabled.
        /// </summary>
        [Output("genericAlert")]
        public Output<string> GenericAlert { get; private set; } = null!;

        /// <summary>
        /// Handshake time out (seconds)
        /// </summary>
        [Output("handshakeTimeout")]
        public Output<string> HandshakeTimeout { get; private set; } = null!;

        /// <summary>
        /// Specifies the file name of the SSL key.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// ModSSL Methods enabled / disabled. Default is disabled.
        /// </summary>
        [Output("modSslMethods")]
        public Output<string> ModSslMethods { get; private set; } = null!;

        /// <summary>
        /// ModSSL Methods enabled / disabled. Default is disabled.
        /// </summary>
        [Output("mode")]
        public Output<string> Mode { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the profile. (type `string`)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Device partition to manage resources on.
        /// </summary>
        [Output("partition")]
        public Output<string> Partition { get; private set; } = null!;

        /// <summary>
        /// Client Certificate Constrained Delegation CA passphrase
        /// </summary>
        [Output("passphrase")]
        public Output<string> Passphrase { get; private set; } = null!;

        /// <summary>
        /// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        /// </summary>
        [Output("peerCertMode")]
        public Output<string> PeerCertMode { get; private set; } = null!;

        /// <summary>
        /// Proxy CA Cert
        /// </summary>
        [Output("proxyCaCert")]
        public Output<string> ProxyCaCert { get; private set; } = null!;

        /// <summary>
        /// Proxy CA Key
        /// </summary>
        [Output("proxyCaKey")]
        public Output<string> ProxyCaKey { get; private set; } = null!;

        /// <summary>
        /// Proxy SSL enabled / disabled. Default is disabled.
        /// </summary>
        [Output("proxySsl")]
        public Output<string> ProxySsl { get; private set; } = null!;

        /// <summary>
        /// Renogotiate Period (seconds)
        /// </summary>
        [Output("renegotiatePeriod")]
        public Output<string> RenegotiatePeriod { get; private set; } = null!;

        /// <summary>
        /// Renogotiate Size
        /// </summary>
        [Output("renegotiateSize")]
        public Output<string> RenegotiateSize { get; private set; } = null!;

        /// <summary>
        /// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
        /// </summary>
        [Output("renegotiation")]
        public Output<string> Renegotiation { get; private set; } = null!;

        /// <summary>
        /// When `true`, client certificate is retained in SSL session.
        /// </summary>
        [Output("retainCertificate")]
        public Output<string> RetainCertificate { get; private set; } = null!;

        /// <summary>
        /// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
        /// When `request` is set the system request secure renegotation of SSL connections.
        /// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
        /// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
        /// </summary>
        [Output("secureRenegotiation")]
        public Output<string> SecureRenegotiation { get; private set; } = null!;

        /// <summary>
        /// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
        /// </summary>
        [Output("serverName")]
        public Output<string> ServerName { get; private set; } = null!;

        /// <summary>
        /// Session Mirroring (enabled / disabled)
        /// </summary>
        [Output("sessionMirroring")]
        public Output<string> SessionMirroring { get; private set; } = null!;

        /// <summary>
        /// Session Ticket (enabled / disabled)
        /// </summary>
        [Output("sessionTicket")]
        public Output<string> SessionTicket { get; private set; } = null!;

        /// <summary>
        /// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
        /// There can be only one SSL profile with this setting enabled.
        /// </summary>
        [Output("sniDefault")]
        public Output<string> SniDefault { get; private set; } = null!;

        /// <summary>
        /// Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
        /// </summary>
        [Output("sniRequire")]
        public Output<string> SniRequire { get; private set; } = null!;

        /// <summary>
        /// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
        /// </summary>
        [Output("sslForwardProxy")]
        public Output<string> SslForwardProxy { get; private set; } = null!;

        /// <summary>
        /// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
        /// </summary>
        [Output("sslForwardProxyBypass")]
        public Output<string> SslForwardProxyBypass { get; private set; } = null!;

        /// <summary>
        /// SSL sign hash (any, sha1, sha256, sha384)
        /// </summary>
        [Output("sslSignHash")]
        public Output<string> SslSignHash { get; private set; } = null!;

        /// <summary>
        /// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.
        /// </summary>
        [Output("strictResume")]
        public Output<string> StrictResume { get; private set; } = null!;

        [Output("tmOptions")]
        public Output<ImmutableArray<string>> TmOptions { get; private set; } = null!;

        /// <summary>
        /// Unclean Shutdown (enabled / disabled)
        /// </summary>
        [Output("uncleanShutdown")]
        public Output<string> UncleanShutdown { get; private set; } = null!;

        /// <summary>
        /// Unclean Shutdown (drop / ignore)
        /// </summary>
        [Output("untrustedCertResponseControl")]
        public Output<string> UntrustedCertResponseControl { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileServerSsl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileServerSsl(string name, ProfileServerSslArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileServerSsl:ProfileServerSsl", name, args ?? new ProfileServerSslArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileServerSsl(string name, Input<string> id, ProfileServerSslState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileServerSsl:ProfileServerSsl", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProfileServerSsl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileServerSsl Get(string name, Input<string> id, ProfileServerSslState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileServerSsl(name, id, state, options);
        }
    }

    public sealed class ProfileServerSslArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert time out
        /// </summary>
        [Input("alertTimeout")]
        public Input<string>? AlertTimeout { get; set; }

        /// <summary>
        /// Server authentication once / always (default is once).
        /// </summary>
        [Input("authenticate")]
        public Input<string>? Authenticate { get; set; }

        /// <summary>
        /// Client certificate chain traversal depth. Default 9.
        /// </summary>
        [Input("authenticateDepth")]
        public Input<int>? AuthenticateDepth { get; set; }

        /// <summary>
        /// Client certificate file path. Default None.
        /// </summary>
        [Input("caFile")]
        public Input<string>? CaFile { get; set; }

        /// <summary>
        /// Cache size (sessions).
        /// </summary>
        [Input("cacheSize")]
        public Input<int>? CacheSize { get; set; }

        /// <summary>
        /// Cache time out
        /// </summary>
        [Input("cacheTimeout")]
        public Input<int>? CacheTimeout { get; set; }

        /// <summary>
        /// Specifies the name of the certificate that the system uses for server-side SSL processing.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Specifies the certificates-key chain to associate with the SSL profile
        /// </summary>
        [Input("chain")]
        public Input<string>? Chain { get; set; }

        /// <summary>
        /// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        /// </summary>
        [Input("ciphers")]
        public Input<string>? Ciphers { get; set; }

        /// <summary>
        /// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Response if the cert is expired (drop / ignore).
        /// </summary>
        [Input("expireCertResponseControl")]
        public Input<string>? ExpireCertResponseControl { get; set; }

        /// <summary>
        /// full path of the profile
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        /// <summary>
        /// generation
        /// </summary>
        [Input("generation")]
        public Input<int>? Generation { get; set; }

        /// <summary>
        /// Generic alerts enabled / disabled.
        /// </summary>
        [Input("genericAlert")]
        public Input<string>? GenericAlert { get; set; }

        /// <summary>
        /// Handshake time out (seconds)
        /// </summary>
        [Input("handshakeTimeout")]
        public Input<string>? HandshakeTimeout { get; set; }

        /// <summary>
        /// Specifies the file name of the SSL key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// ModSSL Methods enabled / disabled. Default is disabled.
        /// </summary>
        [Input("modSslMethods")]
        public Input<string>? ModSslMethods { get; set; }

        /// <summary>
        /// ModSSL Methods enabled / disabled. Default is disabled.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the name of the profile. (type `string`)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Device partition to manage resources on.
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Client Certificate Constrained Delegation CA passphrase
        /// </summary>
        [Input("passphrase")]
        public Input<string>? Passphrase { get; set; }

        /// <summary>
        /// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        /// </summary>
        [Input("peerCertMode")]
        public Input<string>? PeerCertMode { get; set; }

        /// <summary>
        /// Proxy CA Cert
        /// </summary>
        [Input("proxyCaCert")]
        public Input<string>? ProxyCaCert { get; set; }

        /// <summary>
        /// Proxy CA Key
        /// </summary>
        [Input("proxyCaKey")]
        public Input<string>? ProxyCaKey { get; set; }

        /// <summary>
        /// Proxy SSL enabled / disabled. Default is disabled.
        /// </summary>
        [Input("proxySsl")]
        public Input<string>? ProxySsl { get; set; }

        /// <summary>
        /// Renogotiate Period (seconds)
        /// </summary>
        [Input("renegotiatePeriod")]
        public Input<string>? RenegotiatePeriod { get; set; }

        /// <summary>
        /// Renogotiate Size
        /// </summary>
        [Input("renegotiateSize")]
        public Input<string>? RenegotiateSize { get; set; }

        /// <summary>
        /// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
        /// </summary>
        [Input("renegotiation")]
        public Input<string>? Renegotiation { get; set; }

        /// <summary>
        /// When `true`, client certificate is retained in SSL session.
        /// </summary>
        [Input("retainCertificate")]
        public Input<string>? RetainCertificate { get; set; }

        /// <summary>
        /// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
        /// When `request` is set the system request secure renegotation of SSL connections.
        /// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
        /// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
        /// </summary>
        [Input("secureRenegotiation")]
        public Input<string>? SecureRenegotiation { get; set; }

        /// <summary>
        /// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// Session Mirroring (enabled / disabled)
        /// </summary>
        [Input("sessionMirroring")]
        public Input<string>? SessionMirroring { get; set; }

        /// <summary>
        /// Session Ticket (enabled / disabled)
        /// </summary>
        [Input("sessionTicket")]
        public Input<string>? SessionTicket { get; set; }

        /// <summary>
        /// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
        /// There can be only one SSL profile with this setting enabled.
        /// </summary>
        [Input("sniDefault")]
        public Input<string>? SniDefault { get; set; }

        /// <summary>
        /// Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
        /// </summary>
        [Input("sniRequire")]
        public Input<string>? SniRequire { get; set; }

        /// <summary>
        /// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
        /// </summary>
        [Input("sslForwardProxy")]
        public Input<string>? SslForwardProxy { get; set; }

        /// <summary>
        /// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
        /// </summary>
        [Input("sslForwardProxyBypass")]
        public Input<string>? SslForwardProxyBypass { get; set; }

        /// <summary>
        /// SSL sign hash (any, sha1, sha256, sha384)
        /// </summary>
        [Input("sslSignHash")]
        public Input<string>? SslSignHash { get; set; }

        /// <summary>
        /// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.
        /// </summary>
        [Input("strictResume")]
        public Input<string>? StrictResume { get; set; }

        [Input("tmOptions")]
        private InputList<string>? _tmOptions;
        public InputList<string> TmOptions
        {
            get => _tmOptions ?? (_tmOptions = new InputList<string>());
            set => _tmOptions = value;
        }

        /// <summary>
        /// Unclean Shutdown (enabled / disabled)
        /// </summary>
        [Input("uncleanShutdown")]
        public Input<string>? UncleanShutdown { get; set; }

        /// <summary>
        /// Unclean Shutdown (drop / ignore)
        /// </summary>
        [Input("untrustedCertResponseControl")]
        public Input<string>? UntrustedCertResponseControl { get; set; }

        public ProfileServerSslArgs()
        {
        }
    }

    public sealed class ProfileServerSslState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert time out
        /// </summary>
        [Input("alertTimeout")]
        public Input<string>? AlertTimeout { get; set; }

        /// <summary>
        /// Server authentication once / always (default is once).
        /// </summary>
        [Input("authenticate")]
        public Input<string>? Authenticate { get; set; }

        /// <summary>
        /// Client certificate chain traversal depth. Default 9.
        /// </summary>
        [Input("authenticateDepth")]
        public Input<int>? AuthenticateDepth { get; set; }

        /// <summary>
        /// Client certificate file path. Default None.
        /// </summary>
        [Input("caFile")]
        public Input<string>? CaFile { get; set; }

        /// <summary>
        /// Cache size (sessions).
        /// </summary>
        [Input("cacheSize")]
        public Input<int>? CacheSize { get; set; }

        /// <summary>
        /// Cache time out
        /// </summary>
        [Input("cacheTimeout")]
        public Input<int>? CacheTimeout { get; set; }

        /// <summary>
        /// Specifies the name of the certificate that the system uses for server-side SSL processing.
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Specifies the certificates-key chain to associate with the SSL profile
        /// </summary>
        [Input("chain")]
        public Input<string>? Chain { get; set; }

        /// <summary>
        /// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        /// </summary>
        [Input("ciphers")]
        public Input<string>? Ciphers { get; set; }

        /// <summary>
        /// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Response if the cert is expired (drop / ignore).
        /// </summary>
        [Input("expireCertResponseControl")]
        public Input<string>? ExpireCertResponseControl { get; set; }

        /// <summary>
        /// full path of the profile
        /// </summary>
        [Input("fullPath")]
        public Input<string>? FullPath { get; set; }

        /// <summary>
        /// generation
        /// </summary>
        [Input("generation")]
        public Input<int>? Generation { get; set; }

        /// <summary>
        /// Generic alerts enabled / disabled.
        /// </summary>
        [Input("genericAlert")]
        public Input<string>? GenericAlert { get; set; }

        /// <summary>
        /// Handshake time out (seconds)
        /// </summary>
        [Input("handshakeTimeout")]
        public Input<string>? HandshakeTimeout { get; set; }

        /// <summary>
        /// Specifies the file name of the SSL key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// ModSSL Methods enabled / disabled. Default is disabled.
        /// </summary>
        [Input("modSslMethods")]
        public Input<string>? ModSslMethods { get; set; }

        /// <summary>
        /// ModSSL Methods enabled / disabled. Default is disabled.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Specifies the name of the profile. (type `string`)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Device partition to manage resources on.
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Client Certificate Constrained Delegation CA passphrase
        /// </summary>
        [Input("passphrase")]
        public Input<string>? Passphrase { get; set; }

        /// <summary>
        /// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        /// </summary>
        [Input("peerCertMode")]
        public Input<string>? PeerCertMode { get; set; }

        /// <summary>
        /// Proxy CA Cert
        /// </summary>
        [Input("proxyCaCert")]
        public Input<string>? ProxyCaCert { get; set; }

        /// <summary>
        /// Proxy CA Key
        /// </summary>
        [Input("proxyCaKey")]
        public Input<string>? ProxyCaKey { get; set; }

        /// <summary>
        /// Proxy SSL enabled / disabled. Default is disabled.
        /// </summary>
        [Input("proxySsl")]
        public Input<string>? ProxySsl { get; set; }

        /// <summary>
        /// Renogotiate Period (seconds)
        /// </summary>
        [Input("renegotiatePeriod")]
        public Input<string>? RenegotiatePeriod { get; set; }

        /// <summary>
        /// Renogotiate Size
        /// </summary>
        [Input("renegotiateSize")]
        public Input<string>? RenegotiateSize { get; set; }

        /// <summary>
        /// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
        /// </summary>
        [Input("renegotiation")]
        public Input<string>? Renegotiation { get; set; }

        /// <summary>
        /// When `true`, client certificate is retained in SSL session.
        /// </summary>
        [Input("retainCertificate")]
        public Input<string>? RetainCertificate { get; set; }

        /// <summary>
        /// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
        /// When `request` is set the system request secure renegotation of SSL connections.
        /// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
        /// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
        /// </summary>
        [Input("secureRenegotiation")]
        public Input<string>? SecureRenegotiation { get; set; }

        /// <summary>
        /// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// Session Mirroring (enabled / disabled)
        /// </summary>
        [Input("sessionMirroring")]
        public Input<string>? SessionMirroring { get; set; }

        /// <summary>
        /// Session Ticket (enabled / disabled)
        /// </summary>
        [Input("sessionTicket")]
        public Input<string>? SessionTicket { get; set; }

        /// <summary>
        /// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
        /// There can be only one SSL profile with this setting enabled.
        /// </summary>
        [Input("sniDefault")]
        public Input<string>? SniDefault { get; set; }

        /// <summary>
        /// Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
        /// </summary>
        [Input("sniRequire")]
        public Input<string>? SniRequire { get; set; }

        /// <summary>
        /// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
        /// </summary>
        [Input("sslForwardProxy")]
        public Input<string>? SslForwardProxy { get; set; }

        /// <summary>
        /// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
        /// </summary>
        [Input("sslForwardProxyBypass")]
        public Input<string>? SslForwardProxyBypass { get; set; }

        /// <summary>
        /// SSL sign hash (any, sha1, sha256, sha384)
        /// </summary>
        [Input("sslSignHash")]
        public Input<string>? SslSignHash { get; set; }

        /// <summary>
        /// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.
        /// </summary>
        [Input("strictResume")]
        public Input<string>? StrictResume { get; set; }

        [Input("tmOptions")]
        private InputList<string>? _tmOptions;
        public InputList<string> TmOptions
        {
            get => _tmOptions ?? (_tmOptions = new InputList<string>());
            set => _tmOptions = value;
        }

        /// <summary>
        /// Unclean Shutdown (enabled / disabled)
        /// </summary>
        [Input("uncleanShutdown")]
        public Input<string>? UncleanShutdown { get; set; }

        /// <summary>
        /// Unclean Shutdown (drop / ignore)
        /// </summary>
        [Input("untrustedCertResponseControl")]
        public Input<string>? UntrustedCertResponseControl { get; set; }

        public ProfileServerSslState()
        {
        }
    }
}
