// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.ProfileServerSsl` Manages server SSL profiles on a BIG-IP
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewProfileServerSsl(ctx, "test_ServerSsl", &ltm.ProfileServerSslArgs{
// 			Authenticate: pulumi.String("always"),
// 			Ciphers:      pulumi.String("DEFAULT"),
// 			DefaultsFrom: pulumi.String("/Common/serverssl"),
// 			Name:         pulumi.String("/Common/test-ServerSsl"),
// 			Partition:    pulumi.String("Common"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ProfileServerSsl struct {
	pulumi.CustomResourceState

	// Alert time out
	AlertTimeout pulumi.StringOutput `pulumi:"alertTimeout"`
	// Server authentication once / always (default is once).
	Authenticate pulumi.StringOutput `pulumi:"authenticate"`
	// Client certificate chain traversal depth. Default 9.
	AuthenticateDepth pulumi.IntOutput `pulumi:"authenticateDepth"`
	// Client certificate file path. Default None.
	CaFile pulumi.StringOutput `pulumi:"caFile"`
	// Cache size (sessions).
	CacheSize pulumi.IntOutput `pulumi:"cacheSize"`
	// Cache time out
	CacheTimeout pulumi.IntOutput `pulumi:"cacheTimeout"`
	// Specifies the name of the certificate that the system uses for server-side SSL processing.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Specifies the certificates-key chain to associate with the SSL profile
	Chain pulumi.StringOutput `pulumi:"chain"`
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers pulumi.StringOutput `pulumi:"ciphers"`
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
	DefaultsFrom pulumi.StringPtrOutput `pulumi:"defaultsFrom"`
	// Response if the cert is expired (drop / ignore).
	ExpireCertResponseControl pulumi.StringOutput `pulumi:"expireCertResponseControl"`
	// full path of the profile
	FullPath pulumi.StringOutput `pulumi:"fullPath"`
	// generation
	Generation pulumi.IntOutput `pulumi:"generation"`
	// Generic alerts enabled / disabled.
	GenericAlert pulumi.StringOutput `pulumi:"genericAlert"`
	// Handshake time out (seconds)
	HandshakeTimeout pulumi.StringOutput `pulumi:"handshakeTimeout"`
	// Specifies the file name of the SSL key.
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
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl pulumi.StringOutput `pulumi:"proxySsl"`
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
	StrictResume pulumi.StringOutput      `pulumi:"strictResume"`
	TmOptions    pulumi.StringArrayOutput `pulumi:"tmOptions"`
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown pulumi.StringOutput `pulumi:"uncleanShutdown"`
	// Unclean Shutdown (drop / ignore)
	UntrustedCertResponseControl pulumi.StringOutput `pulumi:"untrustedCertResponseControl"`
}

// NewProfileServerSsl registers a new resource with the given unique name, arguments, and options.
func NewProfileServerSsl(ctx *pulumi.Context,
	name string, args *ProfileServerSslArgs, opts ...pulumi.ResourceOption) (*ProfileServerSsl, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &ProfileServerSslArgs{}
	}
	var resource ProfileServerSsl
	err := ctx.RegisterResource("f5bigip:ltm/profileServerSsl:ProfileServerSsl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfileServerSsl gets an existing ProfileServerSsl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfileServerSsl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileServerSslState, opts ...pulumi.ResourceOption) (*ProfileServerSsl, error) {
	var resource ProfileServerSsl
	err := ctx.ReadResource("f5bigip:ltm/profileServerSsl:ProfileServerSsl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfileServerSsl resources.
type profileServerSslState struct {
	// Alert time out
	AlertTimeout *string `pulumi:"alertTimeout"`
	// Server authentication once / always (default is once).
	Authenticate *string `pulumi:"authenticate"`
	// Client certificate chain traversal depth. Default 9.
	AuthenticateDepth *int `pulumi:"authenticateDepth"`
	// Client certificate file path. Default None.
	CaFile *string `pulumi:"caFile"`
	// Cache size (sessions).
	CacheSize *int `pulumi:"cacheSize"`
	// Cache time out
	CacheTimeout *int `pulumi:"cacheTimeout"`
	// Specifies the name of the certificate that the system uses for server-side SSL processing.
	Cert *string `pulumi:"cert"`
	// Specifies the certificates-key chain to associate with the SSL profile
	Chain *string `pulumi:"chain"`
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers *string `pulumi:"ciphers"`
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Response if the cert is expired (drop / ignore).
	ExpireCertResponseControl *string `pulumi:"expireCertResponseControl"`
	// full path of the profile
	FullPath *string `pulumi:"fullPath"`
	// generation
	Generation *int `pulumi:"generation"`
	// Generic alerts enabled / disabled.
	GenericAlert *string `pulumi:"genericAlert"`
	// Handshake time out (seconds)
	HandshakeTimeout *string `pulumi:"handshakeTimeout"`
	// Specifies the file name of the SSL key.
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
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl *string `pulumi:"proxySsl"`
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
	StrictResume *string  `pulumi:"strictResume"`
	TmOptions    []string `pulumi:"tmOptions"`
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown *string `pulumi:"uncleanShutdown"`
	// Unclean Shutdown (drop / ignore)
	UntrustedCertResponseControl *string `pulumi:"untrustedCertResponseControl"`
}

type ProfileServerSslState struct {
	// Alert time out
	AlertTimeout pulumi.StringPtrInput
	// Server authentication once / always (default is once).
	Authenticate pulumi.StringPtrInput
	// Client certificate chain traversal depth. Default 9.
	AuthenticateDepth pulumi.IntPtrInput
	// Client certificate file path. Default None.
	CaFile pulumi.StringPtrInput
	// Cache size (sessions).
	CacheSize pulumi.IntPtrInput
	// Cache time out
	CacheTimeout pulumi.IntPtrInput
	// Specifies the name of the certificate that the system uses for server-side SSL processing.
	Cert pulumi.StringPtrInput
	// Specifies the certificates-key chain to associate with the SSL profile
	Chain pulumi.StringPtrInput
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers pulumi.StringPtrInput
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
	DefaultsFrom pulumi.StringPtrInput
	// Response if the cert is expired (drop / ignore).
	ExpireCertResponseControl pulumi.StringPtrInput
	// full path of the profile
	FullPath pulumi.StringPtrInput
	// generation
	Generation pulumi.IntPtrInput
	// Generic alerts enabled / disabled.
	GenericAlert pulumi.StringPtrInput
	// Handshake time out (seconds)
	HandshakeTimeout pulumi.StringPtrInput
	// Specifies the file name of the SSL key.
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
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl pulumi.StringPtrInput
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
	TmOptions    pulumi.StringArrayInput
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown pulumi.StringPtrInput
	// Unclean Shutdown (drop / ignore)
	UntrustedCertResponseControl pulumi.StringPtrInput
}

func (ProfileServerSslState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileServerSslState)(nil)).Elem()
}

type profileServerSslArgs struct {
	// Alert time out
	AlertTimeout *string `pulumi:"alertTimeout"`
	// Server authentication once / always (default is once).
	Authenticate *string `pulumi:"authenticate"`
	// Client certificate chain traversal depth. Default 9.
	AuthenticateDepth *int `pulumi:"authenticateDepth"`
	// Client certificate file path. Default None.
	CaFile *string `pulumi:"caFile"`
	// Cache size (sessions).
	CacheSize *int `pulumi:"cacheSize"`
	// Cache time out
	CacheTimeout *int `pulumi:"cacheTimeout"`
	// Specifies the name of the certificate that the system uses for server-side SSL processing.
	Cert *string `pulumi:"cert"`
	// Specifies the certificates-key chain to associate with the SSL profile
	Chain *string `pulumi:"chain"`
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers *string `pulumi:"ciphers"`
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Response if the cert is expired (drop / ignore).
	ExpireCertResponseControl *string `pulumi:"expireCertResponseControl"`
	// full path of the profile
	FullPath *string `pulumi:"fullPath"`
	// generation
	Generation *int `pulumi:"generation"`
	// Generic alerts enabled / disabled.
	GenericAlert *string `pulumi:"genericAlert"`
	// Handshake time out (seconds)
	HandshakeTimeout *string `pulumi:"handshakeTimeout"`
	// Specifies the file name of the SSL key.
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
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl *string `pulumi:"proxySsl"`
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
	StrictResume *string  `pulumi:"strictResume"`
	TmOptions    []string `pulumi:"tmOptions"`
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown *string `pulumi:"uncleanShutdown"`
	// Unclean Shutdown (drop / ignore)
	UntrustedCertResponseControl *string `pulumi:"untrustedCertResponseControl"`
}

// The set of arguments for constructing a ProfileServerSsl resource.
type ProfileServerSslArgs struct {
	// Alert time out
	AlertTimeout pulumi.StringPtrInput
	// Server authentication once / always (default is once).
	Authenticate pulumi.StringPtrInput
	// Client certificate chain traversal depth. Default 9.
	AuthenticateDepth pulumi.IntPtrInput
	// Client certificate file path. Default None.
	CaFile pulumi.StringPtrInput
	// Cache size (sessions).
	CacheSize pulumi.IntPtrInput
	// Cache time out
	CacheTimeout pulumi.IntPtrInput
	// Specifies the name of the certificate that the system uses for server-side SSL processing.
	Cert pulumi.StringPtrInput
	// Specifies the certificates-key chain to associate with the SSL profile
	Chain pulumi.StringPtrInput
	// Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
	Ciphers pulumi.StringPtrInput
	// The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
	DefaultsFrom pulumi.StringPtrInput
	// Response if the cert is expired (drop / ignore).
	ExpireCertResponseControl pulumi.StringPtrInput
	// full path of the profile
	FullPath pulumi.StringPtrInput
	// generation
	Generation pulumi.IntPtrInput
	// Generic alerts enabled / disabled.
	GenericAlert pulumi.StringPtrInput
	// Handshake time out (seconds)
	HandshakeTimeout pulumi.StringPtrInput
	// Specifies the file name of the SSL key.
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
	// Proxy SSL enabled / disabled. Default is disabled.
	ProxySsl pulumi.StringPtrInput
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
	TmOptions    pulumi.StringArrayInput
	// Unclean Shutdown (enabled / disabled)
	UncleanShutdown pulumi.StringPtrInput
	// Unclean Shutdown (drop / ignore)
	UntrustedCertResponseControl pulumi.StringPtrInput
}

func (ProfileServerSslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileServerSslArgs)(nil)).Elem()
}
