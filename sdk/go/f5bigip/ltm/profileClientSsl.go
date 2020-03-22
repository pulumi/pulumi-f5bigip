// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `ltm.ProfileClientSsl` Manages client SSL profiles on a BIG-IP
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_profile_client_ssl.html.markdown.
type ProfileClientSsl struct {
	pulumi.CustomResourceState

	// Alert time out
	AlertTimeout pulumi.StringOutput `pulumi:"alertTimeout"`
	// Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
	AllowNonSsl pulumi.StringOutput `pulumi:"allowNonSsl"`
	// Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
	// When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
	Authenticate pulumi.StringOutput `pulumi:"authenticate"`
	// Specifies the maximum number of certificates to be traversed in a client certificate chain
	AuthenticateDepth pulumi.IntOutput `pulumi:"authenticateDepth"`
	// Client certificate file path. Default None.
	CaFile pulumi.StringOutput `pulumi:"caFile"`
	// Cache size (sessions).
	CacheSize pulumi.IntOutput `pulumi:"cacheSize"`
	// Cache time out
	CacheTimeout pulumi.IntOutput `pulumi:"cacheTimeout"`
	// Specifies a cert name for use.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Cert extension includes for ssl forward proxy
	CertExtensionIncludes pulumi.StringArrayOutput `pulumi:"certExtensionIncludes"`
	CertKeyChains ProfileClientSslCertKeyChainArrayOutput `pulumi:"certKeyChains"`
	// Life span of the certificate in days for ssl forward proxy
	CertLifeSpan pulumi.IntOutput `pulumi:"certLifeSpan"`
	// Cert lookup by ip address and port enabled / disabled
	CertLookupByIpaddrPort pulumi.StringOutput `pulumi:"certLookupByIpaddrPort"`
	// Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
	Chain pulumi.StringOutput `pulumi:"chain"`
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers pulumi.StringOutput `pulumi:"ciphers"`
	// client certificate name
	ClientCertCa pulumi.StringOutput `pulumi:"clientCertCa"`
	// Certificate revocation file name
	CrlFile pulumi.StringOutput `pulumi:"crlFile"`
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Forward proxy bypass default action. (enabled / disabled)
	ForwardProxyBypassDefaultAction pulumi.StringOutput `pulumi:"forwardProxyBypassDefaultAction"`
	// full path of the profile
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// generation
	Generation pulumi.IntOutput `pulumi:"generation"`
	// Generic alerts enabled / disabled.
	GenericAlert pulumi.StringOutput `pulumi:"genericAlert"`
	// Handshake time out (seconds)
	HandshakeTimeout pulumi.StringOutput `pulumi:"handshakeTimeout"`
	// Inherit cert key chain
	InheritCertKeychain pulumi.StringOutput `pulumi:"inheritCertKeychain"`
	// Contains a key name
	Key pulumi.StringOutput `pulumi:"key"`
	// ModSSL Methods enabled / disabled. Default is disabled.
	ModSslMethods pulumi.StringOutput `pulumi:"modSslMethods"`
	// ModSSL Methods enabled / disabled. Default is disabled.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Specifies the name of the profile. (type `string`)
	Name pulumi.StringOutput `pulumi:"name"`
	// Device partition to manage resources on.
	Partition pulumi.StringOutput `pulumi:"partition"`
	// Client Certificate Constrained Delegation CA passphrase
	Passphrase pulumi.StringOutput `pulumi:"passphrase"`
	// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
	PeerCertMode pulumi.StringOutput `pulumi:"peerCertMode"`
	// Proxy CA Cert
	ProxyCaCert pulumi.StringOutput `pulumi:"proxyCaCert"`
	// Proxy CA Key
	ProxyCaKey pulumi.StringOutput `pulumi:"proxyCaKey"`
	// Proxy CA Passphrase
	ProxyCaPassphrase pulumi.StringOutput `pulumi:"proxyCaPassphrase"`
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl pulumi.StringOutput `pulumi:"proxySsl"`
	// Proxy SSL passthrough enabled / disabled. Default is disabled.
	ProxySslPassthrough pulumi.StringOutput `pulumi:"proxySslPassthrough"`
	// Renogotiate Period (seconds)
	RenegotiatePeriod pulumi.StringOutput `pulumi:"renegotiatePeriod"`
	// Renogotiate Size
	RenegotiateSize pulumi.StringOutput `pulumi:"renegotiateSize"`
	// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
	Renegotiation pulumi.StringOutput `pulumi:"renegotiation"`
	// When `true`, client certificate is retained in SSL session.
	RetainCertificate pulumi.StringOutput `pulumi:"retainCertificate"`
	// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
	// When `request` is set the system request secure renegotation of SSL connections.
	// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
	// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
	SecureRenegotiation pulumi.StringOutput `pulumi:"secureRenegotiation"`
	// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Session Mirroring (enabled / disabled)
	SessionMirroring pulumi.StringOutput `pulumi:"sessionMirroring"`
	// Session Ticket (enabled / disabled)
	SessionTicket pulumi.StringOutput `pulumi:"sessionTicket"`
	// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
	// There can be only one SSL profile with this setting enabled.
	SniDefault pulumi.StringOutput `pulumi:"sniDefault"`
	// Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
	SniRequire pulumi.StringOutput `pulumi:"sniRequire"`
	// SSL forward Proxy (enabled / disabled)
	SslForwardProxy pulumi.StringOutput `pulumi:"sslForwardProxy"`
	// SSL forward Proxy Bypass (enabled / disabled)
	SslForwardProxyBypass pulumi.StringOutput `pulumi:"sslForwardProxyBypass"`
	// SSL sign hash (any, sha1, sha256, sha384)
	SslSignHash pulumi.StringOutput `pulumi:"sslSignHash"`
	// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
	StrictResume pulumi.StringOutput `pulumi:"strictResume"`
	TmOptions pulumi.StringArrayOutput `pulumi:"tmOptions"`
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown pulumi.StringOutput `pulumi:"uncleanShutdown"`
}

// NewProfileClientSsl registers a new resource with the given unique name, arguments, and options.
func NewProfileClientSsl(ctx *pulumi.Context,
	name string, args *ProfileClientSslArgs, opts ...pulumi.ResourceOption) (*ProfileClientSsl, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &ProfileClientSslArgs{}
	}
	var resource ProfileClientSsl
	err := ctx.RegisterResource("f5bigip:ltm/profileClientSsl:ProfileClientSsl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileClientSsl gets an existing ProfileClientSsl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileClientSsl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileClientSslState, opts ...pulumi.ResourceOption) (*ProfileClientSsl, error) {
	var resource ProfileClientSsl
	err := ctx.ReadResource("f5bigip:ltm/profileClientSsl:ProfileClientSsl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileClientSsl resources.
type profileClientSslState struct {
	// Alert time out
	AlertTimeout *string `pulumi:"alertTimeout"`
	// Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
	AllowNonSsl *string `pulumi:"allowNonSsl"`
	// Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
	// When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
	Authenticate *string `pulumi:"authenticate"`
	// Specifies the maximum number of certificates to be traversed in a client certificate chain
	AuthenticateDepth *int `pulumi:"authenticateDepth"`
	// Client certificate file path. Default None.
	CaFile *string `pulumi:"caFile"`
	// Cache size (sessions).
	CacheSize *int `pulumi:"cacheSize"`
	// Cache time out
	CacheTimeout *int `pulumi:"cacheTimeout"`
	// Specifies a cert name for use.
	Cert *string `pulumi:"cert"`
	// Cert extension includes for ssl forward proxy
	CertExtensionIncludes []string `pulumi:"certExtensionIncludes"`
	CertKeyChains []ProfileClientSslCertKeyChain `pulumi:"certKeyChains"`
	// Life span of the certificate in days for ssl forward proxy
	CertLifeSpan *int `pulumi:"certLifeSpan"`
	// Cert lookup by ip address and port enabled / disabled
	CertLookupByIpaddrPort *string `pulumi:"certLookupByIpaddrPort"`
	// Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
	Chain *string `pulumi:"chain"`
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers *string `pulumi:"ciphers"`
	// client certificate name
	ClientCertCa *string `pulumi:"clientCertCa"`
	// Certificate revocation file name
	CrlFile *string `pulumi:"crlFile"`
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Forward proxy bypass default action. (enabled / disabled)
	ForwardProxyBypassDefaultAction *string `pulumi:"forwardProxyBypassDefaultAction"`
	// full path of the profile
	FullPath *string `pulumi:"fullPath"`
	// generation
	Generation *int `pulumi:"generation"`
	// Generic alerts enabled / disabled.
	GenericAlert *string `pulumi:"genericAlert"`
	// Handshake time out (seconds)
	HandshakeTimeout *string `pulumi:"handshakeTimeout"`
	// Inherit cert key chain
	InheritCertKeychain *string `pulumi:"inheritCertKeychain"`
	// Contains a key name
	Key *string `pulumi:"key"`
	// ModSSL Methods enabled / disabled. Default is disabled.
	ModSslMethods *string `pulumi:"modSslMethods"`
	// ModSSL Methods enabled / disabled. Default is disabled.
	Mode *string `pulumi:"mode"`
	// Specifies the name of the profile. (type `string`)
	Name *string `pulumi:"name"`
	// Device partition to manage resources on.
	Partition *string `pulumi:"partition"`
	// Client Certificate Constrained Delegation CA passphrase
	Passphrase *string `pulumi:"passphrase"`
	// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
	PeerCertMode *string `pulumi:"peerCertMode"`
	// Proxy CA Cert
	ProxyCaCert *string `pulumi:"proxyCaCert"`
	// Proxy CA Key
	ProxyCaKey *string `pulumi:"proxyCaKey"`
	// Proxy CA Passphrase
	ProxyCaPassphrase *string `pulumi:"proxyCaPassphrase"`
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl *string `pulumi:"proxySsl"`
	// Proxy SSL passthrough enabled / disabled. Default is disabled.
	ProxySslPassthrough *string `pulumi:"proxySslPassthrough"`
	// Renogotiate Period (seconds)
	RenegotiatePeriod *string `pulumi:"renegotiatePeriod"`
	// Renogotiate Size
	RenegotiateSize *string `pulumi:"renegotiateSize"`
	// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
	Renegotiation *string `pulumi:"renegotiation"`
	// When `true`, client certificate is retained in SSL session.
	RetainCertificate *string `pulumi:"retainCertificate"`
	// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
	// When `request` is set the system request secure renegotation of SSL connections.
	// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
	// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
	SecureRenegotiation *string `pulumi:"secureRenegotiation"`
	// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
	ServerName *string `pulumi:"serverName"`
	// Session Mirroring (enabled / disabled)
	SessionMirroring *string `pulumi:"sessionMirroring"`
	// Session Ticket (enabled / disabled)
	SessionTicket *string `pulumi:"sessionTicket"`
	// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
	// There can be only one SSL profile with this setting enabled.
	SniDefault *string `pulumi:"sniDefault"`
	// Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
	SniRequire *string `pulumi:"sniRequire"`
	// SSL forward Proxy (enabled / disabled)
	SslForwardProxy *string `pulumi:"sslForwardProxy"`
	// SSL forward Proxy Bypass (enabled / disabled)
	SslForwardProxyBypass *string `pulumi:"sslForwardProxyBypass"`
	// SSL sign hash (any, sha1, sha256, sha384)
	SslSignHash *string `pulumi:"sslSignHash"`
	// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
	StrictResume *string `pulumi:"strictResume"`
	TmOptions []string `pulumi:"tmOptions"`
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown *string `pulumi:"uncleanShutdown"`
}

type ProfileClientSslState struct {
	// Alert time out
	AlertTimeout pulumi.StringPtrInput
	// Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
	AllowNonSsl pulumi.StringPtrInput
	// Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
	// When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
	Authenticate pulumi.StringPtrInput
	// Specifies the maximum number of certificates to be traversed in a client certificate chain
	AuthenticateDepth pulumi.IntPtrInput
	// Client certificate file path. Default None.
	CaFile pulumi.StringPtrInput
	// Cache size (sessions).
	CacheSize pulumi.IntPtrInput
	// Cache time out
	CacheTimeout pulumi.IntPtrInput
	// Specifies a cert name for use.
	Cert pulumi.StringPtrInput
	// Cert extension includes for ssl forward proxy
	CertExtensionIncludes pulumi.StringArrayInput
	CertKeyChains ProfileClientSslCertKeyChainArrayInput
	// Life span of the certificate in days for ssl forward proxy
	CertLifeSpan pulumi.IntPtrInput
	// Cert lookup by ip address and port enabled / disabled
	CertLookupByIpaddrPort pulumi.StringPtrInput
	// Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
	Chain pulumi.StringPtrInput
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers pulumi.StringPtrInput
	// client certificate name
	ClientCertCa pulumi.StringPtrInput
	// Certificate revocation file name
	CrlFile pulumi.StringPtrInput
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
	DefaultsFrom pulumi.StringPtrInput
	// Forward proxy bypass default action. (enabled / disabled)
	ForwardProxyBypassDefaultAction pulumi.StringPtrInput
	// full path of the profile
	FullPath pulumi.StringPtrInput
	// generation
	Generation pulumi.IntPtrInput
	// Generic alerts enabled / disabled.
	GenericAlert pulumi.StringPtrInput
	// Handshake time out (seconds)
	HandshakeTimeout pulumi.StringPtrInput
	// Inherit cert key chain
	InheritCertKeychain pulumi.StringPtrInput
	// Contains a key name
	Key pulumi.StringPtrInput
	// ModSSL Methods enabled / disabled. Default is disabled.
	ModSslMethods pulumi.StringPtrInput
	// ModSSL Methods enabled / disabled. Default is disabled.
	Mode pulumi.StringPtrInput
	// Specifies the name of the profile. (type `string`)
	Name pulumi.StringPtrInput
	// Device partition to manage resources on.
	Partition pulumi.StringPtrInput
	// Client Certificate Constrained Delegation CA passphrase
	Passphrase pulumi.StringPtrInput
	// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
	PeerCertMode pulumi.StringPtrInput
	// Proxy CA Cert
	ProxyCaCert pulumi.StringPtrInput
	// Proxy CA Key
	ProxyCaKey pulumi.StringPtrInput
	// Proxy CA Passphrase
	ProxyCaPassphrase pulumi.StringPtrInput
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl pulumi.StringPtrInput
	// Proxy SSL passthrough enabled / disabled. Default is disabled.
	ProxySslPassthrough pulumi.StringPtrInput
	// Renogotiate Period (seconds)
	RenegotiatePeriod pulumi.StringPtrInput
	// Renogotiate Size
	RenegotiateSize pulumi.StringPtrInput
	// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
	Renegotiation pulumi.StringPtrInput
	// When `true`, client certificate is retained in SSL session.
	RetainCertificate pulumi.StringPtrInput
	// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
	// When `request` is set the system request secure renegotation of SSL connections.
	// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
	// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
	SecureRenegotiation pulumi.StringPtrInput
	// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
	ServerName pulumi.StringPtrInput
	// Session Mirroring (enabled / disabled)
	SessionMirroring pulumi.StringPtrInput
	// Session Ticket (enabled / disabled)
	SessionTicket pulumi.StringPtrInput
	// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
	// There can be only one SSL profile with this setting enabled.
	SniDefault pulumi.StringPtrInput
	// Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
	SniRequire pulumi.StringPtrInput
	// SSL forward Proxy (enabled / disabled)
	SslForwardProxy pulumi.StringPtrInput
	// SSL forward Proxy Bypass (enabled / disabled)
	SslForwardProxyBypass pulumi.StringPtrInput
	// SSL sign hash (any, sha1, sha256, sha384)
	SslSignHash pulumi.StringPtrInput
	// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
	StrictResume pulumi.StringPtrInput
	TmOptions pulumi.StringArrayInput
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown pulumi.StringPtrInput
}

func (ProfileClientSslState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileClientSslState)(nil)).Elem()
}

type profileClientSslArgs struct {
	// Alert time out
	AlertTimeout *string `pulumi:"alertTimeout"`
	// Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
	AllowNonSsl *string `pulumi:"allowNonSsl"`
	// Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
	// When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
	Authenticate *string `pulumi:"authenticate"`
	// Specifies the maximum number of certificates to be traversed in a client certificate chain
	AuthenticateDepth *int `pulumi:"authenticateDepth"`
	// Client certificate file path. Default None.
	CaFile *string `pulumi:"caFile"`
	// Cache size (sessions).
	CacheSize *int `pulumi:"cacheSize"`
	// Cache time out
	CacheTimeout *int `pulumi:"cacheTimeout"`
	// Specifies a cert name for use.
	Cert *string `pulumi:"cert"`
	// Cert extension includes for ssl forward proxy
	CertExtensionIncludes []string `pulumi:"certExtensionIncludes"`
	CertKeyChains []ProfileClientSslCertKeyChain `pulumi:"certKeyChains"`
	// Life span of the certificate in days for ssl forward proxy
	CertLifeSpan *int `pulumi:"certLifeSpan"`
	// Cert lookup by ip address and port enabled / disabled
	CertLookupByIpaddrPort *string `pulumi:"certLookupByIpaddrPort"`
	// Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
	Chain *string `pulumi:"chain"`
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers *string `pulumi:"ciphers"`
	// client certificate name
	ClientCertCa *string `pulumi:"clientCertCa"`
	// Certificate revocation file name
	CrlFile *string `pulumi:"crlFile"`
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Forward proxy bypass default action. (enabled / disabled)
	ForwardProxyBypassDefaultAction *string `pulumi:"forwardProxyBypassDefaultAction"`
	// full path of the profile
	FullPath *string `pulumi:"fullPath"`
	// generation
	Generation *int `pulumi:"generation"`
	// Generic alerts enabled / disabled.
	GenericAlert *string `pulumi:"genericAlert"`
	// Handshake time out (seconds)
	HandshakeTimeout *string `pulumi:"handshakeTimeout"`
	// Inherit cert key chain
	InheritCertKeychain *string `pulumi:"inheritCertKeychain"`
	// Contains a key name
	Key *string `pulumi:"key"`
	// ModSSL Methods enabled / disabled. Default is disabled.
	ModSslMethods *string `pulumi:"modSslMethods"`
	// ModSSL Methods enabled / disabled. Default is disabled.
	Mode *string `pulumi:"mode"`
	// Specifies the name of the profile. (type `string`)
	Name string `pulumi:"name"`
	// Device partition to manage resources on.
	Partition *string `pulumi:"partition"`
	// Client Certificate Constrained Delegation CA passphrase
	Passphrase *string `pulumi:"passphrase"`
	// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
	PeerCertMode *string `pulumi:"peerCertMode"`
	// Proxy CA Cert
	ProxyCaCert *string `pulumi:"proxyCaCert"`
	// Proxy CA Key
	ProxyCaKey *string `pulumi:"proxyCaKey"`
	// Proxy CA Passphrase
	ProxyCaPassphrase *string `pulumi:"proxyCaPassphrase"`
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl *string `pulumi:"proxySsl"`
	// Proxy SSL passthrough enabled / disabled. Default is disabled.
	ProxySslPassthrough *string `pulumi:"proxySslPassthrough"`
	// Renogotiate Period (seconds)
	RenegotiatePeriod *string `pulumi:"renegotiatePeriod"`
	// Renogotiate Size
	RenegotiateSize *string `pulumi:"renegotiateSize"`
	// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
	Renegotiation *string `pulumi:"renegotiation"`
	// When `true`, client certificate is retained in SSL session.
	RetainCertificate *string `pulumi:"retainCertificate"`
	// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
	// When `request` is set the system request secure renegotation of SSL connections.
	// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
	// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
	SecureRenegotiation *string `pulumi:"secureRenegotiation"`
	// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
	ServerName *string `pulumi:"serverName"`
	// Session Mirroring (enabled / disabled)
	SessionMirroring *string `pulumi:"sessionMirroring"`
	// Session Ticket (enabled / disabled)
	SessionTicket *string `pulumi:"sessionTicket"`
	// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
	// There can be only one SSL profile with this setting enabled.
	SniDefault *string `pulumi:"sniDefault"`
	// Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
	SniRequire *string `pulumi:"sniRequire"`
	// SSL forward Proxy (enabled / disabled)
	SslForwardProxy *string `pulumi:"sslForwardProxy"`
	// SSL forward Proxy Bypass (enabled / disabled)
	SslForwardProxyBypass *string `pulumi:"sslForwardProxyBypass"`
	// SSL sign hash (any, sha1, sha256, sha384)
	SslSignHash *string `pulumi:"sslSignHash"`
	// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
	StrictResume *string `pulumi:"strictResume"`
	TmOptions []string `pulumi:"tmOptions"`
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown *string `pulumi:"uncleanShutdown"`
}

// The set of arguments for constructing a ProfileClientSsl resource.
type ProfileClientSslArgs struct {
	// Alert time out
	AlertTimeout pulumi.StringPtrInput
	// Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
	AllowNonSsl pulumi.StringPtrInput
	// Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
	// When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
	Authenticate pulumi.StringPtrInput
	// Specifies the maximum number of certificates to be traversed in a client certificate chain
	AuthenticateDepth pulumi.IntPtrInput
	// Client certificate file path. Default None.
	CaFile pulumi.StringPtrInput
	// Cache size (sessions).
	CacheSize pulumi.IntPtrInput
	// Cache time out
	CacheTimeout pulumi.IntPtrInput
	// Specifies a cert name for use.
	Cert pulumi.StringPtrInput
	// Cert extension includes for ssl forward proxy
	CertExtensionIncludes pulumi.StringArrayInput
	CertKeyChains ProfileClientSslCertKeyChainArrayInput
	// Life span of the certificate in days for ssl forward proxy
	CertLifeSpan pulumi.IntPtrInput
	// Cert lookup by ip address and port enabled / disabled
	CertLookupByIpaddrPort pulumi.StringPtrInput
	// Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
	Chain pulumi.StringPtrInput
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers pulumi.StringPtrInput
	// client certificate name
	ClientCertCa pulumi.StringPtrInput
	// Certificate revocation file name
	CrlFile pulumi.StringPtrInput
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
	DefaultsFrom pulumi.StringPtrInput
	// Forward proxy bypass default action. (enabled / disabled)
	ForwardProxyBypassDefaultAction pulumi.StringPtrInput
	// full path of the profile
	FullPath pulumi.StringPtrInput
	// generation
	Generation pulumi.IntPtrInput
	// Generic alerts enabled / disabled.
	GenericAlert pulumi.StringPtrInput
	// Handshake time out (seconds)
	HandshakeTimeout pulumi.StringPtrInput
	// Inherit cert key chain
	InheritCertKeychain pulumi.StringPtrInput
	// Contains a key name
	Key pulumi.StringPtrInput
	// ModSSL Methods enabled / disabled. Default is disabled.
	ModSslMethods pulumi.StringPtrInput
	// ModSSL Methods enabled / disabled. Default is disabled.
	Mode pulumi.StringPtrInput
	// Specifies the name of the profile. (type `string`)
	Name pulumi.StringInput
	// Device partition to manage resources on.
	Partition pulumi.StringPtrInput
	// Client Certificate Constrained Delegation CA passphrase
	Passphrase pulumi.StringPtrInput
	// Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
	PeerCertMode pulumi.StringPtrInput
	// Proxy CA Cert
	ProxyCaCert pulumi.StringPtrInput
	// Proxy CA Key
	ProxyCaKey pulumi.StringPtrInput
	// Proxy CA Passphrase
	ProxyCaPassphrase pulumi.StringPtrInput
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl pulumi.StringPtrInput
	// Proxy SSL passthrough enabled / disabled. Default is disabled.
	ProxySslPassthrough pulumi.StringPtrInput
	// Renogotiate Period (seconds)
	RenegotiatePeriod pulumi.StringPtrInput
	// Renogotiate Size
	RenegotiateSize pulumi.StringPtrInput
	// Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
	Renegotiation pulumi.StringPtrInput
	// When `true`, client certificate is retained in SSL session.
	RetainCertificate pulumi.StringPtrInput
	// Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
	// When `request` is set the system request secure renegotation of SSL connections.
	// `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
	// The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
	SecureRenegotiation pulumi.StringPtrInput
	// Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
	ServerName pulumi.StringPtrInput
	// Session Mirroring (enabled / disabled)
	SessionMirroring pulumi.StringPtrInput
	// Session Ticket (enabled / disabled)
	SessionTicket pulumi.StringPtrInput
	// Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
	// There can be only one SSL profile with this setting enabled.
	SniDefault pulumi.StringPtrInput
	// Requires that the network peers also provide SNI support, this setting only takes effect when `sniDefault` is set to `true`.When creating a new profile, the setting is provided by the parent profile
	SniRequire pulumi.StringPtrInput
	// SSL forward Proxy (enabled / disabled)
	SslForwardProxy pulumi.StringPtrInput
	// SSL forward Proxy Bypass (enabled / disabled)
	SslForwardProxyBypass pulumi.StringPtrInput
	// SSL sign hash (any, sha1, sha256, sha384)
	SslSignHash pulumi.StringPtrInput
	// Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
	StrictResume pulumi.StringPtrInput
	TmOptions pulumi.StringArrayInput
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown pulumi.StringPtrInput
}

func (ProfileClientSslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileClientSslArgs)(nil)).Elem()
}

