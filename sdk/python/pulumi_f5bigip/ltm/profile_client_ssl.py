# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ProfileClientSsl(pulumi.CustomResource):
    alert_timeout: pulumi.Output[str]
    """
    Alert time out
    """
    allow_non_ssl: pulumi.Output[str]
    """
    Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
    """
    authenticate: pulumi.Output[str]
    """
    Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
    When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
    """
    authenticate_depth: pulumi.Output[float]
    """
    Specifies the maximum number of certificates to be traversed in a client certificate chain
    """
    ca_file: pulumi.Output[str]
    """
    Client certificate file path. Default None.
    """
    cache_size: pulumi.Output[float]
    """
    Cache size (sessions).
    """
    cache_timeout: pulumi.Output[float]
    """
    Cache time out
    """
    cert: pulumi.Output[str]
    """
    Specifies a cert name for use.
    """
    cert_extension_includes: pulumi.Output[list]
    """
    Cert extension includes for ssl forward proxy
    """
    cert_key_chains: pulumi.Output[list]
    cert_life_span: pulumi.Output[float]
    """
    Life span of the certificate in days for ssl forward proxy
    """
    cert_lookup_by_ipaddr_port: pulumi.Output[str]
    """
    Cert lookup by ip address and port enabled / disabled
    """
    chain: pulumi.Output[str]
    """
    Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
    """
    ciphers: pulumi.Output[str]
    """
    Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
    """
    client_cert_ca: pulumi.Output[str]
    """
    client certificate name
    """
    crl_file: pulumi.Output[str]
    """
    Certificate revocation file name
    """
    defaults_from: pulumi.Output[str]
    """
    The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
    """
    forward_proxy_bypass_default_action: pulumi.Output[str]
    """
    Forward proxy bypass default action. (enabled / disabled)
    """
    full_path: pulumi.Output[str]
    """
    full path of the profile
    """
    generation: pulumi.Output[float]
    """
    generation
    """
    generic_alert: pulumi.Output[str]
    """
    Generic alerts enabled / disabled.
    """
    handshake_timeout: pulumi.Output[str]
    """
    Handshake time out (seconds)
    """
    inherit_cert_keychain: pulumi.Output[str]
    """
    Inherit cert key chain
    """
    key: pulumi.Output[str]
    """
    Contains a key name
    """
    mod_ssl_methods: pulumi.Output[str]
    """
    ModSSL Methods enabled / disabled. Default is disabled.
    """
    mode: pulumi.Output[str]
    """
    ModSSL Methods enabled / disabled. Default is disabled.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the profile. (type `string`)
    """
    partition: pulumi.Output[str]
    """
    Device partition to manage resources on.
    """
    passphrase: pulumi.Output[str]
    """
    Client Certificate Constrained Delegation CA passphrase
    """
    peer_cert_mode: pulumi.Output[str]
    """
    Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
    """
    proxy_ca_cert: pulumi.Output[str]
    """
    Proxy CA Cert
    """
    proxy_ca_key: pulumi.Output[str]
    """
    Proxy CA Key
    """
    proxy_ca_passphrase: pulumi.Output[str]
    """
    Proxy CA Passphrase
    """
    proxy_ssl: pulumi.Output[str]
    """
    Proxy SSL enabled / disabled. Default is disabled.
    """
    proxy_ssl_passthrough: pulumi.Output[str]
    """
    Proxy SSL passthrough enabled / disabled. Default is disabled.
    """
    renegotiate_period: pulumi.Output[str]
    """
    Renogotiate Period (seconds)
    """
    renegotiate_size: pulumi.Output[str]
    """
    Renogotiate Size
    """
    renegotiation: pulumi.Output[str]
    """
    Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
    """
    retain_certificate: pulumi.Output[str]
    """
    When `true`, client certificate is retained in SSL session.
    """
    secure_renegotiation: pulumi.Output[str]
    """
    Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
    When `request` is set the system request secure renegotation of SSL connections.
    `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
    The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
    """
    server_name: pulumi.Output[str]
    """
    Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
    """
    session_mirroring: pulumi.Output[str]
    """
    Session Mirroring (enabled / disabled)
    """
    session_ticket: pulumi.Output[str]
    """
    Session Ticket (enabled / disabled)
    """
    sni_default: pulumi.Output[str]
    """
    Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
    There can be only one SSL profile with this setting enabled.
    """
    sni_require: pulumi.Output[str]
    """
    Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
    """
    ssl_forward_proxy: pulumi.Output[str]
    """
    SSL forward Proxy (enabled / disabled)
    """
    ssl_forward_proxy_bypass: pulumi.Output[str]
    """
    SSL forward Proxy Bypass (enabled / disabled)
    """
    ssl_sign_hash: pulumi.Output[str]
    """
    SSL sign hash (any, sha1, sha256, sha384)
    """
    strict_resume: pulumi.Output[str]
    """
    Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
    """
    tm_options: pulumi.Output[list]
    unclean_shutdown: pulumi.Output[str]
    """
    Unclean Shutdown (enabled / disabled)
    """
    def __init__(__self__, resource_name, opts=None, alert_timeout=None, allow_non_ssl=None, authenticate=None, authenticate_depth=None, ca_file=None, cache_size=None, cache_timeout=None, cert=None, cert_extension_includes=None, cert_key_chains=None, cert_life_span=None, cert_lookup_by_ipaddr_port=None, chain=None, ciphers=None, client_cert_ca=None, crl_file=None, defaults_from=None, forward_proxy_bypass_default_action=None, full_path=None, generation=None, generic_alert=None, handshake_timeout=None, inherit_cert_keychain=None, key=None, mod_ssl_methods=None, mode=None, name=None, partition=None, passphrase=None, peer_cert_mode=None, proxy_ca_cert=None, proxy_ca_key=None, proxy_ca_passphrase=None, proxy_ssl=None, proxy_ssl_passthrough=None, renegotiate_period=None, renegotiate_size=None, renegotiation=None, retain_certificate=None, secure_renegotiation=None, server_name=None, session_mirroring=None, session_ticket=None, sni_default=None, sni_require=None, ssl_forward_proxy=None, ssl_forward_proxy_bypass=None, ssl_sign_hash=None, strict_resume=None, tm_options=None, unclean_shutdown=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.ProfileClientSsl` Manages client SSL profiles on a BIG-IP





        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_profile_client_ssl.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_timeout: Alert time out
        :param pulumi.Input[str] allow_non_ssl: Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
        :param pulumi.Input[str] authenticate: Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
               When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
        :param pulumi.Input[float] authenticate_depth: Specifies the maximum number of certificates to be traversed in a client certificate chain
        :param pulumi.Input[str] ca_file: Client certificate file path. Default None.
        :param pulumi.Input[float] cache_size: Cache size (sessions).
        :param pulumi.Input[float] cache_timeout: Cache time out
        :param pulumi.Input[str] cert: Specifies a cert name for use.
        :param pulumi.Input[list] cert_extension_includes: Cert extension includes for ssl forward proxy
        :param pulumi.Input[float] cert_life_span: Life span of the certificate in days for ssl forward proxy
        :param pulumi.Input[str] cert_lookup_by_ipaddr_port: Cert lookup by ip address and port enabled / disabled
        :param pulumi.Input[str] chain: Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
        :param pulumi.Input[str] ciphers: Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        :param pulumi.Input[str] client_cert_ca: client certificate name
        :param pulumi.Input[str] crl_file: Certificate revocation file name
        :param pulumi.Input[str] defaults_from: The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
        :param pulumi.Input[str] forward_proxy_bypass_default_action: Forward proxy bypass default action. (enabled / disabled)
        :param pulumi.Input[str] full_path: full path of the profile
        :param pulumi.Input[float] generation: generation
        :param pulumi.Input[str] generic_alert: Generic alerts enabled / disabled.
        :param pulumi.Input[str] handshake_timeout: Handshake time out (seconds)
        :param pulumi.Input[str] inherit_cert_keychain: Inherit cert key chain
        :param pulumi.Input[str] key: Contains a key name
        :param pulumi.Input[str] mod_ssl_methods: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] mode: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] name: Specifies the name of the profile. (type `string`)
        :param pulumi.Input[str] partition: Device partition to manage resources on.
        :param pulumi.Input[str] passphrase: Client Certificate Constrained Delegation CA passphrase
        :param pulumi.Input[str] peer_cert_mode: Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        :param pulumi.Input[str] proxy_ca_cert: Proxy CA Cert
        :param pulumi.Input[str] proxy_ca_key: Proxy CA Key
        :param pulumi.Input[str] proxy_ca_passphrase: Proxy CA Passphrase
        :param pulumi.Input[str] proxy_ssl: Proxy SSL enabled / disabled. Default is disabled.
        :param pulumi.Input[str] proxy_ssl_passthrough: Proxy SSL passthrough enabled / disabled. Default is disabled.
        :param pulumi.Input[str] renegotiate_period: Renogotiate Period (seconds)
        :param pulumi.Input[str] renegotiate_size: Renogotiate Size
        :param pulumi.Input[str] renegotiation: Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
        :param pulumi.Input[str] retain_certificate: When `true`, client certificate is retained in SSL session.
        :param pulumi.Input[str] secure_renegotiation: Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
               When `request` is set the system request secure renegotation of SSL connections.
               `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
               The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
        :param pulumi.Input[str] server_name: Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
        :param pulumi.Input[str] session_mirroring: Session Mirroring (enabled / disabled)
        :param pulumi.Input[str] session_ticket: Session Ticket (enabled / disabled)
        :param pulumi.Input[str] sni_default: Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
               There can be only one SSL profile with this setting enabled.
        :param pulumi.Input[str] sni_require: Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
        :param pulumi.Input[str] ssl_forward_proxy: SSL forward Proxy (enabled / disabled)
        :param pulumi.Input[str] ssl_forward_proxy_bypass: SSL forward Proxy Bypass (enabled / disabled)
        :param pulumi.Input[str] ssl_sign_hash: SSL sign hash (any, sha1, sha256, sha384)
        :param pulumi.Input[str] strict_resume: Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
        :param pulumi.Input[str] unclean_shutdown: Unclean Shutdown (enabled / disabled)

        The **cert_key_chains** object supports the following:

          * `cert` (`pulumi.Input[str]`) - Specifies a cert name for use.
          * `chain` (`pulumi.Input[str]`) - Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
          * `key` (`pulumi.Input[str]`) - Contains a key name
          * `name` (`pulumi.Input[str]`) - Specifies the name of the profile. (type `string`)
          * `passphrase` (`pulumi.Input[str]`)
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['alert_timeout'] = alert_timeout
            __props__['allow_non_ssl'] = allow_non_ssl
            __props__['authenticate'] = authenticate
            __props__['authenticate_depth'] = authenticate_depth
            __props__['ca_file'] = ca_file
            __props__['cache_size'] = cache_size
            __props__['cache_timeout'] = cache_timeout
            __props__['cert'] = cert
            __props__['cert_extension_includes'] = cert_extension_includes
            __props__['cert_key_chains'] = cert_key_chains
            __props__['cert_life_span'] = cert_life_span
            __props__['cert_lookup_by_ipaddr_port'] = cert_lookup_by_ipaddr_port
            __props__['chain'] = chain
            __props__['ciphers'] = ciphers
            __props__['client_cert_ca'] = client_cert_ca
            __props__['crl_file'] = crl_file
            __props__['defaults_from'] = defaults_from
            __props__['forward_proxy_bypass_default_action'] = forward_proxy_bypass_default_action
            __props__['full_path'] = full_path
            __props__['generation'] = generation
            __props__['generic_alert'] = generic_alert
            __props__['handshake_timeout'] = handshake_timeout
            __props__['inherit_cert_keychain'] = inherit_cert_keychain
            __props__['key'] = key
            __props__['mod_ssl_methods'] = mod_ssl_methods
            __props__['mode'] = mode
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['partition'] = partition
            __props__['passphrase'] = passphrase
            __props__['peer_cert_mode'] = peer_cert_mode
            __props__['proxy_ca_cert'] = proxy_ca_cert
            __props__['proxy_ca_key'] = proxy_ca_key
            __props__['proxy_ca_passphrase'] = proxy_ca_passphrase
            __props__['proxy_ssl'] = proxy_ssl
            __props__['proxy_ssl_passthrough'] = proxy_ssl_passthrough
            __props__['renegotiate_period'] = renegotiate_period
            __props__['renegotiate_size'] = renegotiate_size
            __props__['renegotiation'] = renegotiation
            __props__['retain_certificate'] = retain_certificate
            __props__['secure_renegotiation'] = secure_renegotiation
            __props__['server_name'] = server_name
            __props__['session_mirroring'] = session_mirroring
            __props__['session_ticket'] = session_ticket
            __props__['sni_default'] = sni_default
            __props__['sni_require'] = sni_require
            __props__['ssl_forward_proxy'] = ssl_forward_proxy
            __props__['ssl_forward_proxy_bypass'] = ssl_forward_proxy_bypass
            __props__['ssl_sign_hash'] = ssl_sign_hash
            __props__['strict_resume'] = strict_resume
            __props__['tm_options'] = tm_options
            __props__['unclean_shutdown'] = unclean_shutdown
        super(ProfileClientSsl, __self__).__init__(
            'f5bigip:ltm/profileClientSsl:ProfileClientSsl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, alert_timeout=None, allow_non_ssl=None, authenticate=None, authenticate_depth=None, ca_file=None, cache_size=None, cache_timeout=None, cert=None, cert_extension_includes=None, cert_key_chains=None, cert_life_span=None, cert_lookup_by_ipaddr_port=None, chain=None, ciphers=None, client_cert_ca=None, crl_file=None, defaults_from=None, forward_proxy_bypass_default_action=None, full_path=None, generation=None, generic_alert=None, handshake_timeout=None, inherit_cert_keychain=None, key=None, mod_ssl_methods=None, mode=None, name=None, partition=None, passphrase=None, peer_cert_mode=None, proxy_ca_cert=None, proxy_ca_key=None, proxy_ca_passphrase=None, proxy_ssl=None, proxy_ssl_passthrough=None, renegotiate_period=None, renegotiate_size=None, renegotiation=None, retain_certificate=None, secure_renegotiation=None, server_name=None, session_mirroring=None, session_ticket=None, sni_default=None, sni_require=None, ssl_forward_proxy=None, ssl_forward_proxy_bypass=None, ssl_sign_hash=None, strict_resume=None, tm_options=None, unclean_shutdown=None):
        """
        Get an existing ProfileClientSsl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_timeout: Alert time out
        :param pulumi.Input[str] allow_non_ssl: Enables or disables acceptance of non-SSL connections, When creating a new profile, the setting is provided by the parent profile
        :param pulumi.Input[str] authenticate: Specifies the frequency of client authentication for an SSL session.When `once`,specifies that the system authenticates the client once for an SSL session.
               When `always`, specifies that the system authenticates the client once for an SSL session and also upon reuse of that session.
        :param pulumi.Input[float] authenticate_depth: Specifies the maximum number of certificates to be traversed in a client certificate chain
        :param pulumi.Input[str] ca_file: Client certificate file path. Default None.
        :param pulumi.Input[float] cache_size: Cache size (sessions).
        :param pulumi.Input[float] cache_timeout: Cache time out
        :param pulumi.Input[str] cert: Specifies a cert name for use.
        :param pulumi.Input[list] cert_extension_includes: Cert extension includes for ssl forward proxy
        :param pulumi.Input[float] cert_life_span: Life span of the certificate in days for ssl forward proxy
        :param pulumi.Input[str] cert_lookup_by_ipaddr_port: Cert lookup by ip address and port enabled / disabled
        :param pulumi.Input[str] chain: Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
        :param pulumi.Input[str] ciphers: Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        :param pulumi.Input[str] client_cert_ca: client certificate name
        :param pulumi.Input[str] crl_file: Certificate revocation file name
        :param pulumi.Input[str] defaults_from: The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is the `clientssl` parent on the `Common` partition.
        :param pulumi.Input[str] forward_proxy_bypass_default_action: Forward proxy bypass default action. (enabled / disabled)
        :param pulumi.Input[str] full_path: full path of the profile
        :param pulumi.Input[float] generation: generation
        :param pulumi.Input[str] generic_alert: Generic alerts enabled / disabled.
        :param pulumi.Input[str] handshake_timeout: Handshake time out (seconds)
        :param pulumi.Input[str] inherit_cert_keychain: Inherit cert key chain
        :param pulumi.Input[str] key: Contains a key name
        :param pulumi.Input[str] mod_ssl_methods: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] mode: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] name: Specifies the name of the profile. (type `string`)
        :param pulumi.Input[str] partition: Device partition to manage resources on.
        :param pulumi.Input[str] passphrase: Client Certificate Constrained Delegation CA passphrase
        :param pulumi.Input[str] peer_cert_mode: Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        :param pulumi.Input[str] proxy_ca_cert: Proxy CA Cert
        :param pulumi.Input[str] proxy_ca_key: Proxy CA Key
        :param pulumi.Input[str] proxy_ca_passphrase: Proxy CA Passphrase
        :param pulumi.Input[str] proxy_ssl: Proxy SSL enabled / disabled. Default is disabled.
        :param pulumi.Input[str] proxy_ssl_passthrough: Proxy SSL passthrough enabled / disabled. Default is disabled.
        :param pulumi.Input[str] renegotiate_period: Renogotiate Period (seconds)
        :param pulumi.Input[str] renegotiate_size: Renogotiate Size
        :param pulumi.Input[str] renegotiation: Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
        :param pulumi.Input[str] retain_certificate: When `true`, client certificate is retained in SSL session.
        :param pulumi.Input[str] secure_renegotiation: Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
               When `request` is set the system request secure renegotation of SSL connections.
               `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
               The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
        :param pulumi.Input[str] server_name: Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
        :param pulumi.Input[str] session_mirroring: Session Mirroring (enabled / disabled)
        :param pulumi.Input[str] session_ticket: Session Ticket (enabled / disabled)
        :param pulumi.Input[str] sni_default: Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
               There can be only one SSL profile with this setting enabled.
        :param pulumi.Input[str] sni_require: Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
        :param pulumi.Input[str] ssl_forward_proxy: SSL forward Proxy (enabled / disabled)
        :param pulumi.Input[str] ssl_forward_proxy_bypass: SSL forward Proxy Bypass (enabled / disabled)
        :param pulumi.Input[str] ssl_sign_hash: SSL sign hash (any, sha1, sha256, sha384)
        :param pulumi.Input[str] strict_resume: Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile. 
        :param pulumi.Input[str] unclean_shutdown: Unclean Shutdown (enabled / disabled)

        The **cert_key_chains** object supports the following:

          * `cert` (`pulumi.Input[str]`) - Specifies a cert name for use.
          * `chain` (`pulumi.Input[str]`) - Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
          * `key` (`pulumi.Input[str]`) - Contains a key name
          * `name` (`pulumi.Input[str]`) - Specifies the name of the profile. (type `string`)
          * `passphrase` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alert_timeout"] = alert_timeout
        __props__["allow_non_ssl"] = allow_non_ssl
        __props__["authenticate"] = authenticate
        __props__["authenticate_depth"] = authenticate_depth
        __props__["ca_file"] = ca_file
        __props__["cache_size"] = cache_size
        __props__["cache_timeout"] = cache_timeout
        __props__["cert"] = cert
        __props__["cert_extension_includes"] = cert_extension_includes
        __props__["cert_key_chains"] = cert_key_chains
        __props__["cert_life_span"] = cert_life_span
        __props__["cert_lookup_by_ipaddr_port"] = cert_lookup_by_ipaddr_port
        __props__["chain"] = chain
        __props__["ciphers"] = ciphers
        __props__["client_cert_ca"] = client_cert_ca
        __props__["crl_file"] = crl_file
        __props__["defaults_from"] = defaults_from
        __props__["forward_proxy_bypass_default_action"] = forward_proxy_bypass_default_action
        __props__["full_path"] = full_path
        __props__["generation"] = generation
        __props__["generic_alert"] = generic_alert
        __props__["handshake_timeout"] = handshake_timeout
        __props__["inherit_cert_keychain"] = inherit_cert_keychain
        __props__["key"] = key
        __props__["mod_ssl_methods"] = mod_ssl_methods
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["partition"] = partition
        __props__["passphrase"] = passphrase
        __props__["peer_cert_mode"] = peer_cert_mode
        __props__["proxy_ca_cert"] = proxy_ca_cert
        __props__["proxy_ca_key"] = proxy_ca_key
        __props__["proxy_ca_passphrase"] = proxy_ca_passphrase
        __props__["proxy_ssl"] = proxy_ssl
        __props__["proxy_ssl_passthrough"] = proxy_ssl_passthrough
        __props__["renegotiate_period"] = renegotiate_period
        __props__["renegotiate_size"] = renegotiate_size
        __props__["renegotiation"] = renegotiation
        __props__["retain_certificate"] = retain_certificate
        __props__["secure_renegotiation"] = secure_renegotiation
        __props__["server_name"] = server_name
        __props__["session_mirroring"] = session_mirroring
        __props__["session_ticket"] = session_ticket
        __props__["sni_default"] = sni_default
        __props__["sni_require"] = sni_require
        __props__["ssl_forward_proxy"] = ssl_forward_proxy
        __props__["ssl_forward_proxy_bypass"] = ssl_forward_proxy_bypass
        __props__["ssl_sign_hash"] = ssl_sign_hash
        __props__["strict_resume"] = strict_resume
        __props__["tm_options"] = tm_options
        __props__["unclean_shutdown"] = unclean_shutdown
        return ProfileClientSsl(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

