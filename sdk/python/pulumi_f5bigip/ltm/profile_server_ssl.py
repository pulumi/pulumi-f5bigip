# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ProfileServerSsl']


class ProfileServerSsl(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_timeout: Optional[pulumi.Input[str]] = None,
                 authenticate: Optional[pulumi.Input[str]] = None,
                 authenticate_depth: Optional[pulumi.Input[float]] = None,
                 ca_file: Optional[pulumi.Input[str]] = None,
                 cache_size: Optional[pulumi.Input[float]] = None,
                 cache_timeout: Optional[pulumi.Input[float]] = None,
                 cert: Optional[pulumi.Input[str]] = None,
                 chain: Optional[pulumi.Input[str]] = None,
                 ciphers: Optional[pulumi.Input[str]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 expire_cert_response_control: Optional[pulumi.Input[str]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 generation: Optional[pulumi.Input[float]] = None,
                 generic_alert: Optional[pulumi.Input[str]] = None,
                 handshake_timeout: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 mod_ssl_methods: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 peer_cert_mode: Optional[pulumi.Input[str]] = None,
                 proxy_ssl: Optional[pulumi.Input[str]] = None,
                 renegotiate_period: Optional[pulumi.Input[str]] = None,
                 renegotiate_size: Optional[pulumi.Input[str]] = None,
                 renegotiation: Optional[pulumi.Input[str]] = None,
                 retain_certificate: Optional[pulumi.Input[str]] = None,
                 secure_renegotiation: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 session_mirroring: Optional[pulumi.Input[str]] = None,
                 session_ticket: Optional[pulumi.Input[str]] = None,
                 sni_default: Optional[pulumi.Input[str]] = None,
                 sni_require: Optional[pulumi.Input[str]] = None,
                 ssl_forward_proxy: Optional[pulumi.Input[str]] = None,
                 ssl_forward_proxy_bypass: Optional[pulumi.Input[str]] = None,
                 ssl_sign_hash: Optional[pulumi.Input[str]] = None,
                 strict_resume: Optional[pulumi.Input[str]] = None,
                 tm_options: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 unclean_shutdown: Optional[pulumi.Input[str]] = None,
                 untrusted_cert_response_control: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.ProfileServerSsl` Manages server SSL profiles on a BIG-IP

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test__server_ssl = f5bigip.ltm.ProfileServerSsl("test-ServerSsl",
            authenticate="always",
            ciphers="DEFAULT",
            defaults_from="/Common/serverssl",
            name="/Common/test-ServerSsl",
            partition="Common")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_timeout: Alert time out
        :param pulumi.Input[str] authenticate: Server authentication once / always (default is once).
        :param pulumi.Input[float] authenticate_depth: Client certificate chain traversal depth. Default 9.
        :param pulumi.Input[str] ca_file: Client certificate file path. Default None.
        :param pulumi.Input[float] cache_size: Cache size (sessions).
        :param pulumi.Input[float] cache_timeout: Cache time out
        :param pulumi.Input[str] cert: Specifies the name of the certificate that the system uses for server-side SSL processing.
        :param pulumi.Input[str] chain: Specifies the certificates-key chain to associate with the SSL profile
        :param pulumi.Input[str] ciphers: Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        :param pulumi.Input[str] defaults_from: The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
        :param pulumi.Input[str] expire_cert_response_control: Response if the cert is expired (drop / ignore).
        :param pulumi.Input[str] full_path: full path of the profile
        :param pulumi.Input[float] generation: generation
        :param pulumi.Input[str] generic_alert: Generic alerts enabled / disabled.
        :param pulumi.Input[str] handshake_timeout: Handshake time out (seconds)
        :param pulumi.Input[str] key: Specifies the file name of the SSL key.
        :param pulumi.Input[str] mod_ssl_methods: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] mode: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] name: Specifies the name of the profile. (type `string`)
        :param pulumi.Input[str] partition: Device partition to manage resources on.
        :param pulumi.Input[str] passphrase: Client Certificate Constrained Delegation CA passphrase
        :param pulumi.Input[str] peer_cert_mode: Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        :param pulumi.Input[str] proxy_ssl: Proxy SSL enabled / disabled. Default is disabled.
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
        :param pulumi.Input[str] untrusted_cert_response_control: Unclean Shutdown (drop / ignore)
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['alert_timeout'] = alert_timeout
            __props__['authenticate'] = authenticate
            __props__['authenticate_depth'] = authenticate_depth
            __props__['ca_file'] = ca_file
            __props__['cache_size'] = cache_size
            __props__['cache_timeout'] = cache_timeout
            __props__['cert'] = cert
            __props__['chain'] = chain
            __props__['ciphers'] = ciphers
            __props__['defaults_from'] = defaults_from
            __props__['expire_cert_response_control'] = expire_cert_response_control
            __props__['full_path'] = full_path
            __props__['generation'] = generation
            __props__['generic_alert'] = generic_alert
            __props__['handshake_timeout'] = handshake_timeout
            __props__['key'] = key
            __props__['mod_ssl_methods'] = mod_ssl_methods
            __props__['mode'] = mode
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['partition'] = partition
            __props__['passphrase'] = passphrase
            __props__['peer_cert_mode'] = peer_cert_mode
            __props__['proxy_ssl'] = proxy_ssl
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
            __props__['untrusted_cert_response_control'] = untrusted_cert_response_control
        super(ProfileServerSsl, __self__).__init__(
            'f5bigip:ltm/profileServerSsl:ProfileServerSsl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_timeout: Optional[pulumi.Input[str]] = None,
            authenticate: Optional[pulumi.Input[str]] = None,
            authenticate_depth: Optional[pulumi.Input[float]] = None,
            ca_file: Optional[pulumi.Input[str]] = None,
            cache_size: Optional[pulumi.Input[float]] = None,
            cache_timeout: Optional[pulumi.Input[float]] = None,
            cert: Optional[pulumi.Input[str]] = None,
            chain: Optional[pulumi.Input[str]] = None,
            ciphers: Optional[pulumi.Input[str]] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            expire_cert_response_control: Optional[pulumi.Input[str]] = None,
            full_path: Optional[pulumi.Input[str]] = None,
            generation: Optional[pulumi.Input[float]] = None,
            generic_alert: Optional[pulumi.Input[str]] = None,
            handshake_timeout: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            mod_ssl_methods: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition: Optional[pulumi.Input[str]] = None,
            passphrase: Optional[pulumi.Input[str]] = None,
            peer_cert_mode: Optional[pulumi.Input[str]] = None,
            proxy_ssl: Optional[pulumi.Input[str]] = None,
            renegotiate_period: Optional[pulumi.Input[str]] = None,
            renegotiate_size: Optional[pulumi.Input[str]] = None,
            renegotiation: Optional[pulumi.Input[str]] = None,
            retain_certificate: Optional[pulumi.Input[str]] = None,
            secure_renegotiation: Optional[pulumi.Input[str]] = None,
            server_name: Optional[pulumi.Input[str]] = None,
            session_mirroring: Optional[pulumi.Input[str]] = None,
            session_ticket: Optional[pulumi.Input[str]] = None,
            sni_default: Optional[pulumi.Input[str]] = None,
            sni_require: Optional[pulumi.Input[str]] = None,
            ssl_forward_proxy: Optional[pulumi.Input[str]] = None,
            ssl_forward_proxy_bypass: Optional[pulumi.Input[str]] = None,
            ssl_sign_hash: Optional[pulumi.Input[str]] = None,
            strict_resume: Optional[pulumi.Input[str]] = None,
            tm_options: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            unclean_shutdown: Optional[pulumi.Input[str]] = None,
            untrusted_cert_response_control: Optional[pulumi.Input[str]] = None) -> 'ProfileServerSsl':
        """
        Get an existing ProfileServerSsl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_timeout: Alert time out
        :param pulumi.Input[str] authenticate: Server authentication once / always (default is once).
        :param pulumi.Input[float] authenticate_depth: Client certificate chain traversal depth. Default 9.
        :param pulumi.Input[str] ca_file: Client certificate file path. Default None.
        :param pulumi.Input[float] cache_size: Cache size (sessions).
        :param pulumi.Input[float] cache_timeout: Cache time out
        :param pulumi.Input[str] cert: Specifies the name of the certificate that the system uses for server-side SSL processing.
        :param pulumi.Input[str] chain: Specifies the certificates-key chain to associate with the SSL profile
        :param pulumi.Input[str] ciphers: Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        :param pulumi.Input[str] defaults_from: The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
        :param pulumi.Input[str] expire_cert_response_control: Response if the cert is expired (drop / ignore).
        :param pulumi.Input[str] full_path: full path of the profile
        :param pulumi.Input[float] generation: generation
        :param pulumi.Input[str] generic_alert: Generic alerts enabled / disabled.
        :param pulumi.Input[str] handshake_timeout: Handshake time out (seconds)
        :param pulumi.Input[str] key: Specifies the file name of the SSL key.
        :param pulumi.Input[str] mod_ssl_methods: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] mode: ModSSL Methods enabled / disabled. Default is disabled.
        :param pulumi.Input[str] name: Specifies the name of the profile. (type `string`)
        :param pulumi.Input[str] partition: Device partition to manage resources on.
        :param pulumi.Input[str] passphrase: Client Certificate Constrained Delegation CA passphrase
        :param pulumi.Input[str] peer_cert_mode: Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        :param pulumi.Input[str] proxy_ssl: Proxy SSL enabled / disabled. Default is disabled.
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
        :param pulumi.Input[str] untrusted_cert_response_control: Unclean Shutdown (drop / ignore)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alert_timeout"] = alert_timeout
        __props__["authenticate"] = authenticate
        __props__["authenticate_depth"] = authenticate_depth
        __props__["ca_file"] = ca_file
        __props__["cache_size"] = cache_size
        __props__["cache_timeout"] = cache_timeout
        __props__["cert"] = cert
        __props__["chain"] = chain
        __props__["ciphers"] = ciphers
        __props__["defaults_from"] = defaults_from
        __props__["expire_cert_response_control"] = expire_cert_response_control
        __props__["full_path"] = full_path
        __props__["generation"] = generation
        __props__["generic_alert"] = generic_alert
        __props__["handshake_timeout"] = handshake_timeout
        __props__["key"] = key
        __props__["mod_ssl_methods"] = mod_ssl_methods
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["partition"] = partition
        __props__["passphrase"] = passphrase
        __props__["peer_cert_mode"] = peer_cert_mode
        __props__["proxy_ssl"] = proxy_ssl
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
        __props__["untrusted_cert_response_control"] = untrusted_cert_response_control
        return ProfileServerSsl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertTimeout")
    def alert_timeout(self) -> str:
        """
        Alert time out
        """
        return pulumi.get(self, "alert_timeout")

    @property
    @pulumi.getter
    def authenticate(self) -> str:
        """
        Server authentication once / always (default is once).
        """
        return pulumi.get(self, "authenticate")

    @property
    @pulumi.getter(name="authenticateDepth")
    def authenticate_depth(self) -> float:
        """
        Client certificate chain traversal depth. Default 9.
        """
        return pulumi.get(self, "authenticate_depth")

    @property
    @pulumi.getter(name="caFile")
    def ca_file(self) -> str:
        """
        Client certificate file path. Default None.
        """
        return pulumi.get(self, "ca_file")

    @property
    @pulumi.getter(name="cacheSize")
    def cache_size(self) -> float:
        """
        Cache size (sessions).
        """
        return pulumi.get(self, "cache_size")

    @property
    @pulumi.getter(name="cacheTimeout")
    def cache_timeout(self) -> float:
        """
        Cache time out
        """
        return pulumi.get(self, "cache_timeout")

    @property
    @pulumi.getter
    def cert(self) -> str:
        """
        Specifies the name of the certificate that the system uses for server-side SSL processing.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def chain(self) -> str:
        """
        Specifies the certificates-key chain to associate with the SSL profile
        """
        return pulumi.get(self, "chain")

    @property
    @pulumi.getter
    def ciphers(self) -> str:
        """
        Specifies the list of ciphers that the system supports. When creating a new profile, the default cipher list is provided by the parent profile.
        """
        return pulumi.get(self, "ciphers")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> Optional[str]:
        """
        The parent template of this monitor template. Once this value has been set, it cannot be changed. By default, this value is `/Common/serverssl`.
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter(name="expireCertResponseControl")
    def expire_cert_response_control(self) -> str:
        """
        Response if the cert is expired (drop / ignore).
        """
        return pulumi.get(self, "expire_cert_response_control")

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> str:
        """
        full path of the profile
        """
        return pulumi.get(self, "full_path")

    @property
    @pulumi.getter
    def generation(self) -> float:
        """
        generation
        """
        return pulumi.get(self, "generation")

    @property
    @pulumi.getter(name="genericAlert")
    def generic_alert(self) -> str:
        """
        Generic alerts enabled / disabled.
        """
        return pulumi.get(self, "generic_alert")

    @property
    @pulumi.getter(name="handshakeTimeout")
    def handshake_timeout(self) -> str:
        """
        Handshake time out (seconds)
        """
        return pulumi.get(self, "handshake_timeout")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Specifies the file name of the SSL key.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="modSslMethods")
    def mod_ssl_methods(self) -> str:
        """
        ModSSL Methods enabled / disabled. Default is disabled.
        """
        return pulumi.get(self, "mod_ssl_methods")

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        ModSSL Methods enabled / disabled. Default is disabled.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the profile. (type `string`)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> str:
        """
        Device partition to manage resources on.
        """
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter
    def passphrase(self) -> str:
        """
        Client Certificate Constrained Delegation CA passphrase
        """
        return pulumi.get(self, "passphrase")

    @property
    @pulumi.getter(name="peerCertMode")
    def peer_cert_mode(self) -> str:
        """
        Specifies the way the system handles client certificates.When ignore, specifies that the system ignores certificates from client systems.When require, specifies that the system requires a client to present a valid certificate.When request, specifies that the system requests a valid certificate from a client but always authenticate the client.
        """
        return pulumi.get(self, "peer_cert_mode")

    @property
    @pulumi.getter(name="proxySsl")
    def proxy_ssl(self) -> str:
        """
        Proxy SSL enabled / disabled. Default is disabled.
        """
        return pulumi.get(self, "proxy_ssl")

    @property
    @pulumi.getter(name="renegotiatePeriod")
    def renegotiate_period(self) -> str:
        """
        Renogotiate Period (seconds)
        """
        return pulumi.get(self, "renegotiate_period")

    @property
    @pulumi.getter(name="renegotiateSize")
    def renegotiate_size(self) -> str:
        """
        Renogotiate Size
        """
        return pulumi.get(self, "renegotiate_size")

    @property
    @pulumi.getter
    def renegotiation(self) -> str:
        """
        Enables or disables SSL renegotiation.When creating a new profile, the setting is provided by the parent profile
        """
        return pulumi.get(self, "renegotiation")

    @property
    @pulumi.getter(name="retainCertificate")
    def retain_certificate(self) -> str:
        """
        When `true`, client certificate is retained in SSL session.
        """
        return pulumi.get(self, "retain_certificate")

    @property
    @pulumi.getter(name="secureRenegotiation")
    def secure_renegotiation(self) -> str:
        """
        Specifies the method of secure renegotiations for SSL connections. When creating a new profile, the setting is provided by the parent profile.
        When `request` is set the system request secure renegotation of SSL connections.
        `require` is a default setting and when set the system permits initial SSL handshakes from clients but terminates renegotiations from unpatched clients.
        The `require-strict` setting the system requires strict renegotiation of SSL connections. In this mode the system refuses connections to insecure servers, and terminates existing SSL connections to insecure servers
        """
        return pulumi.get(self, "secure_renegotiation")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> str:
        """
        Specifies the fully qualified DNS hostname of the server used in Server Name Indication communications. When creating a new profile, the setting is provided by the parent profile.The server name can also be a wildcard string containing the asterisk `*` character.
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter(name="sessionMirroring")
    def session_mirroring(self) -> str:
        """
        Session Mirroring (enabled / disabled)
        """
        return pulumi.get(self, "session_mirroring")

    @property
    @pulumi.getter(name="sessionTicket")
    def session_ticket(self) -> str:
        """
        Session Ticket (enabled / disabled)
        """
        return pulumi.get(self, "session_ticket")

    @property
    @pulumi.getter(name="sniDefault")
    def sni_default(self) -> str:
        """
        Indicates that the system uses this profile as the default SSL profile when there is no match to the server name, or when the client provides no SNI extension support.When creating a new profile, the setting is provided by the parent profile.
        There can be only one SSL profile with this setting enabled.
        """
        return pulumi.get(self, "sni_default")

    @property
    @pulumi.getter(name="sniRequire")
    def sni_require(self) -> str:
        """
        Requires that the network peers also provide SNI support, this setting only takes effect when `sni_default` is set to `true`.When creating a new profile, the setting is provided by the parent profile
        """
        return pulumi.get(self, "sni_require")

    @property
    @pulumi.getter(name="sslForwardProxy")
    def ssl_forward_proxy(self) -> str:
        """
        SSL forward Proxy (enabled / disabled)
        """
        return pulumi.get(self, "ssl_forward_proxy")

    @property
    @pulumi.getter(name="sslForwardProxyBypass")
    def ssl_forward_proxy_bypass(self) -> str:
        """
        SSL forward Proxy Bypass (enabled / disabled)
        """
        return pulumi.get(self, "ssl_forward_proxy_bypass")

    @property
    @pulumi.getter(name="sslSignHash")
    def ssl_sign_hash(self) -> str:
        """
        SSL sign hash (any, sha1, sha256, sha384)
        """
        return pulumi.get(self, "ssl_sign_hash")

    @property
    @pulumi.getter(name="strictResume")
    def strict_resume(self) -> str:
        """
        Enables or disables the resumption of SSL sessions after an unclean shutdown.When creating a new profile, the setting is provided by the parent profile.
        """
        return pulumi.get(self, "strict_resume")

    @property
    @pulumi.getter(name="tmOptions")
    def tm_options(self) -> List[str]:
        return pulumi.get(self, "tm_options")

    @property
    @pulumi.getter(name="uncleanShutdown")
    def unclean_shutdown(self) -> str:
        """
        Unclean Shutdown (enabled / disabled)
        """
        return pulumi.get(self, "unclean_shutdown")

    @property
    @pulumi.getter(name="untrustedCertResponseControl")
    def untrusted_cert_response_control(self) -> str:
        """
        Unclean Shutdown (drop / ignore)
        """
        return pulumi.get(self, "untrusted_cert_response_control")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

