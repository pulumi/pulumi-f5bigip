// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
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
	// Proxy CA Cert
	ProxyCaCert pulumi.StringOutput `pulumi:"proxyCaCert"`
	// Proxy CA Key
	ProxyCaKey pulumi.StringOutput `pulumi:"proxyCaKey"`
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
	// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
	SslForwardProxy pulumi.StringOutput `pulumi:"sslForwardProxy"`
	// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
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
	// Proxy CA Cert
	ProxyCaCert *string `pulumi:"proxyCaCert"`
	// Proxy CA Key
	ProxyCaKey *string `pulumi:"proxyCaKey"`
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
	// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
	SslForwardProxy *string `pulumi:"sslForwardProxy"`
	// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
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
	// Proxy CA Cert
	ProxyCaCert pulumi.StringPtrInput
	// Proxy CA Key
	ProxyCaKey pulumi.StringPtrInput
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
	// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
	SslForwardProxy pulumi.StringPtrInput
	// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
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
	// Proxy CA Cert
	ProxyCaCert *string `pulumi:"proxyCaCert"`
	// Proxy CA Key
	ProxyCaKey *string `pulumi:"proxyCaKey"`
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
	// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
	SslForwardProxy *string `pulumi:"sslForwardProxy"`
	// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
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
	// Proxy CA Cert
	ProxyCaCert pulumi.StringPtrInput
	// Proxy CA Key
	ProxyCaKey pulumi.StringPtrInput
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
	// Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled.
	SslForwardProxy pulumi.StringPtrInput
	// Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled.
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

type ProfileServerSslInput interface {
	pulumi.Input

	ToProfileServerSslOutput() ProfileServerSslOutput
	ToProfileServerSslOutputWithContext(ctx context.Context) ProfileServerSslOutput
}

func (*ProfileServerSsl) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileServerSsl)(nil))
}

func (i *ProfileServerSsl) ToProfileServerSslOutput() ProfileServerSslOutput {
	return i.ToProfileServerSslOutputWithContext(context.Background())
}

func (i *ProfileServerSsl) ToProfileServerSslOutputWithContext(ctx context.Context) ProfileServerSslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileServerSslOutput)
}

func (i *ProfileServerSsl) ToProfileServerSslPtrOutput() ProfileServerSslPtrOutput {
	return i.ToProfileServerSslPtrOutputWithContext(context.Background())
}

func (i *ProfileServerSsl) ToProfileServerSslPtrOutputWithContext(ctx context.Context) ProfileServerSslPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileServerSslPtrOutput)
}

type ProfileServerSslPtrInput interface {
	pulumi.Input

	ToProfileServerSslPtrOutput() ProfileServerSslPtrOutput
	ToProfileServerSslPtrOutputWithContext(ctx context.Context) ProfileServerSslPtrOutput
}

type profileServerSslPtrType ProfileServerSslArgs

func (*profileServerSslPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileServerSsl)(nil))
}

func (i *profileServerSslPtrType) ToProfileServerSslPtrOutput() ProfileServerSslPtrOutput {
	return i.ToProfileServerSslPtrOutputWithContext(context.Background())
}

func (i *profileServerSslPtrType) ToProfileServerSslPtrOutputWithContext(ctx context.Context) ProfileServerSslPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileServerSslPtrOutput)
}

// ProfileServerSslArrayInput is an input type that accepts ProfileServerSslArray and ProfileServerSslArrayOutput values.
// You can construct a concrete instance of `ProfileServerSslArrayInput` via:
//
//          ProfileServerSslArray{ ProfileServerSslArgs{...} }
type ProfileServerSslArrayInput interface {
	pulumi.Input

	ToProfileServerSslArrayOutput() ProfileServerSslArrayOutput
	ToProfileServerSslArrayOutputWithContext(context.Context) ProfileServerSslArrayOutput
}

type ProfileServerSslArray []ProfileServerSslInput

func (ProfileServerSslArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ProfileServerSsl)(nil))
}

func (i ProfileServerSslArray) ToProfileServerSslArrayOutput() ProfileServerSslArrayOutput {
	return i.ToProfileServerSslArrayOutputWithContext(context.Background())
}

func (i ProfileServerSslArray) ToProfileServerSslArrayOutputWithContext(ctx context.Context) ProfileServerSslArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileServerSslArrayOutput)
}

// ProfileServerSslMapInput is an input type that accepts ProfileServerSslMap and ProfileServerSslMapOutput values.
// You can construct a concrete instance of `ProfileServerSslMapInput` via:
//
//          ProfileServerSslMap{ "key": ProfileServerSslArgs{...} }
type ProfileServerSslMapInput interface {
	pulumi.Input

	ToProfileServerSslMapOutput() ProfileServerSslMapOutput
	ToProfileServerSslMapOutputWithContext(context.Context) ProfileServerSslMapOutput
}

type ProfileServerSslMap map[string]ProfileServerSslInput

func (ProfileServerSslMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ProfileServerSsl)(nil))
}

func (i ProfileServerSslMap) ToProfileServerSslMapOutput() ProfileServerSslMapOutput {
	return i.ToProfileServerSslMapOutputWithContext(context.Background())
}

func (i ProfileServerSslMap) ToProfileServerSslMapOutputWithContext(ctx context.Context) ProfileServerSslMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfileServerSslMapOutput)
}

type ProfileServerSslOutput struct {
	*pulumi.OutputState
}

func (ProfileServerSslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfileServerSsl)(nil))
}

func (o ProfileServerSslOutput) ToProfileServerSslOutput() ProfileServerSslOutput {
	return o
}

func (o ProfileServerSslOutput) ToProfileServerSslOutputWithContext(ctx context.Context) ProfileServerSslOutput {
	return o
}

func (o ProfileServerSslOutput) ToProfileServerSslPtrOutput() ProfileServerSslPtrOutput {
	return o.ToProfileServerSslPtrOutputWithContext(context.Background())
}

func (o ProfileServerSslOutput) ToProfileServerSslPtrOutputWithContext(ctx context.Context) ProfileServerSslPtrOutput {
	return o.ApplyT(func(v ProfileServerSsl) *ProfileServerSsl {
		return &v
	}).(ProfileServerSslPtrOutput)
}

type ProfileServerSslPtrOutput struct {
	*pulumi.OutputState
}

func (ProfileServerSslPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProfileServerSsl)(nil))
}

func (o ProfileServerSslPtrOutput) ToProfileServerSslPtrOutput() ProfileServerSslPtrOutput {
	return o
}

func (o ProfileServerSslPtrOutput) ToProfileServerSslPtrOutputWithContext(ctx context.Context) ProfileServerSslPtrOutput {
	return o
}

type ProfileServerSslArrayOutput struct{ *pulumi.OutputState }

func (ProfileServerSslArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProfileServerSsl)(nil))
}

func (o ProfileServerSslArrayOutput) ToProfileServerSslArrayOutput() ProfileServerSslArrayOutput {
	return o
}

func (o ProfileServerSslArrayOutput) ToProfileServerSslArrayOutputWithContext(ctx context.Context) ProfileServerSslArrayOutput {
	return o
}

func (o ProfileServerSslArrayOutput) Index(i pulumi.IntInput) ProfileServerSslOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProfileServerSsl {
		return vs[0].([]ProfileServerSsl)[vs[1].(int)]
	}).(ProfileServerSslOutput)
}

type ProfileServerSslMapOutput struct{ *pulumi.OutputState }

func (ProfileServerSslMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ProfileServerSsl)(nil))
}

func (o ProfileServerSslMapOutput) ToProfileServerSslMapOutput() ProfileServerSslMapOutput {
	return o
}

func (o ProfileServerSslMapOutput) ToProfileServerSslMapOutputWithContext(ctx context.Context) ProfileServerSslMapOutput {
	return o
}

func (o ProfileServerSslMapOutput) MapIndex(k pulumi.StringInput) ProfileServerSslOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ProfileServerSsl {
		return vs[0].(map[string]ProfileServerSsl)[vs[1].(string)]
	}).(ProfileServerSslOutput)
}

func init() {
	pulumi.RegisterOutputType(ProfileServerSslOutput{})
	pulumi.RegisterOutputType(ProfileServerSslPtrOutput{})
	pulumi.RegisterOutputType(ProfileServerSslArrayOutput{})
	pulumi.RegisterOutputType(ProfileServerSslMapOutput{})
}
