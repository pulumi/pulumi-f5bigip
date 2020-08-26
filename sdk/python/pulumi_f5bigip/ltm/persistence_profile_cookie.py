# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['PersistenceProfileCookie']


class PersistenceProfileCookie(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 always_send: Optional[pulumi.Input[str]] = None,
                 app_service: Optional[pulumi.Input[str]] = None,
                 cookie_encryption: Optional[pulumi.Input[str]] = None,
                 cookie_encryption_passphrase: Optional[pulumi.Input[str]] = None,
                 cookie_name: Optional[pulumi.Input[str]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 hash_length: Optional[pulumi.Input[float]] = None,
                 hash_offset: Optional[pulumi.Input[float]] = None,
                 httponly: Optional[pulumi.Input[str]] = None,
                 match_across_pools: Optional[pulumi.Input[str]] = None,
                 match_across_services: Optional[pulumi.Input[str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[str]] = None,
                 mirror: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_conn_limit: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configures a cookie persistence profile

        ## Example

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_ppcookie = f5bigip.ltm.PersistenceProfileCookie("testPpcookie",
            name="/Common/terraform_cookie",
            defaults_from="/Common/cookie",
            match_across_pools="enabled",
            match_across_services="enabled",
            match_across_virtuals="enabled",
            timeout=3600,
            override_conn_limit="enabled",
            always_send="enabled",
            cookie_encryption="required",
            cookie_encryption_passphrase="iam",
            cookie_name="ham",
            expiration="1:0:0",
            hash_length=0)
        ```

        ## Reference

        `name` - (Required) Name of the virtual address

        `defaults_from` - (Required) Parent cookie persistence profile

        `match_across_pools` (Optional) (enabled or disabled) match across pools with given persistence record

        `match_across_services` (Optional) (enabled or disabled) match across services with given persistence record

        `match_across_virtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record

        `mirror` (Optional) (enabled or disabled) mirror persistence record

        `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds

        `override_conn_limit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.

        `always_send` (Optional) (enabled or disabled) always send cookies

        `cookie_encryption` (Optional) (required, preferred, or disabled) To required, preferred, or disabled policy for cookie encryption

        `cookie_encryption_passphrase` (Optional) (required, preferred, or disabled) Passphrase for encrypted cookies. The field is encrypted on the server and will always return differently then set.
        If this is configured specify `ignore_changes` under the `lifecycle` block to ignore returned encrypted value.

        `cookie_name` (Optional) Name of the cookie to track persistence

        `expiration` (Optional) Expiration TTL for cookie specified in DAY:HOUR:MIN:SECONDS (Examples: 1:0:0:0 one day, 1:0:0 one hour, 30:0 thirty minutes)

        `hash_length` (Optional) (Integer) Length of hash to apply to cookie

        `hash_offset` (Optional) (Integer) Number of characters to skip in the cookie for the hash

        `httponly` (Optional) (enabled or disabled) Sending only over http

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] always_send: To enable _ disable always sending cookies
        :param pulumi.Input[str] cookie_encryption: To required, preferred, or disabled policy for cookie encryption
        :param pulumi.Input[str] cookie_encryption_passphrase: Passphrase for encrypted cookies
        :param pulumi.Input[str] cookie_name: Name of the cookie to track persistence
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] expiration: Expiration TTL for cookie specified in D:H:M:S or in seconds
        :param pulumi.Input[float] hash_length: Length of hash to apply to cookie
        :param pulumi.Input[float] hash_offset: Number of characters to skip in the cookie for the hash
        :param pulumi.Input[str] httponly: To enable _ disable sending only over http
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across virtual servers with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[float] timeout: Timeout for persistence of the session
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

            __props__['always_send'] = always_send
            __props__['app_service'] = app_service
            __props__['cookie_encryption'] = cookie_encryption
            __props__['cookie_encryption_passphrase'] = cookie_encryption_passphrase
            __props__['cookie_name'] = cookie_name
            if defaults_from is None:
                raise TypeError("Missing required property 'defaults_from'")
            __props__['defaults_from'] = defaults_from
            __props__['expiration'] = expiration
            __props__['hash_length'] = hash_length
            __props__['hash_offset'] = hash_offset
            __props__['httponly'] = httponly
            __props__['match_across_pools'] = match_across_pools
            __props__['match_across_services'] = match_across_services
            __props__['match_across_virtuals'] = match_across_virtuals
            __props__['mirror'] = mirror
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['override_conn_limit'] = override_conn_limit
            __props__['timeout'] = timeout
        super(PersistenceProfileCookie, __self__).__init__(
            'f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            always_send: Optional[pulumi.Input[str]] = None,
            app_service: Optional[pulumi.Input[str]] = None,
            cookie_encryption: Optional[pulumi.Input[str]] = None,
            cookie_encryption_passphrase: Optional[pulumi.Input[str]] = None,
            cookie_name: Optional[pulumi.Input[str]] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            expiration: Optional[pulumi.Input[str]] = None,
            hash_length: Optional[pulumi.Input[float]] = None,
            hash_offset: Optional[pulumi.Input[float]] = None,
            httponly: Optional[pulumi.Input[str]] = None,
            match_across_pools: Optional[pulumi.Input[str]] = None,
            match_across_services: Optional[pulumi.Input[str]] = None,
            match_across_virtuals: Optional[pulumi.Input[str]] = None,
            mirror: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            override_conn_limit: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[float]] = None) -> 'PersistenceProfileCookie':
        """
        Get an existing PersistenceProfileCookie resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] always_send: To enable _ disable always sending cookies
        :param pulumi.Input[str] cookie_encryption: To required, preferred, or disabled policy for cookie encryption
        :param pulumi.Input[str] cookie_encryption_passphrase: Passphrase for encrypted cookies
        :param pulumi.Input[str] cookie_name: Name of the cookie to track persistence
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] expiration: Expiration TTL for cookie specified in D:H:M:S or in seconds
        :param pulumi.Input[float] hash_length: Length of hash to apply to cookie
        :param pulumi.Input[float] hash_offset: Number of characters to skip in the cookie for the hash
        :param pulumi.Input[str] httponly: To enable _ disable sending only over http
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across virtual servers with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[float] timeout: Timeout for persistence of the session
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["always_send"] = always_send
        __props__["app_service"] = app_service
        __props__["cookie_encryption"] = cookie_encryption
        __props__["cookie_encryption_passphrase"] = cookie_encryption_passphrase
        __props__["cookie_name"] = cookie_name
        __props__["defaults_from"] = defaults_from
        __props__["expiration"] = expiration
        __props__["hash_length"] = hash_length
        __props__["hash_offset"] = hash_offset
        __props__["httponly"] = httponly
        __props__["match_across_pools"] = match_across_pools
        __props__["match_across_services"] = match_across_services
        __props__["match_across_virtuals"] = match_across_virtuals
        __props__["mirror"] = mirror
        __props__["name"] = name
        __props__["override_conn_limit"] = override_conn_limit
        __props__["timeout"] = timeout
        return PersistenceProfileCookie(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alwaysSend")
    def always_send(self) -> Optional[str]:
        """
        To enable _ disable always sending cookies
        """
        return pulumi.get(self, "always_send")

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[str]:
        return pulumi.get(self, "app_service")

    @property
    @pulumi.getter(name="cookieEncryption")
    def cookie_encryption(self) -> Optional[str]:
        """
        To required, preferred, or disabled policy for cookie encryption
        """
        return pulumi.get(self, "cookie_encryption")

    @property
    @pulumi.getter(name="cookieEncryptionPassphrase")
    def cookie_encryption_passphrase(self) -> Optional[str]:
        """
        Passphrase for encrypted cookies
        """
        return pulumi.get(self, "cookie_encryption_passphrase")

    @property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> Optional[str]:
        """
        Name of the cookie to track persistence
        """
        return pulumi.get(self, "cookie_name")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> str:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter
    def expiration(self) -> Optional[str]:
        """
        Expiration TTL for cookie specified in D:H:M:S or in seconds
        """
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter(name="hashLength")
    def hash_length(self) -> Optional[float]:
        """
        Length of hash to apply to cookie
        """
        return pulumi.get(self, "hash_length")

    @property
    @pulumi.getter(name="hashOffset")
    def hash_offset(self) -> Optional[float]:
        """
        Number of characters to skip in the cookie for the hash
        """
        return pulumi.get(self, "hash_offset")

    @property
    @pulumi.getter
    def httponly(self) -> Optional[str]:
        """
        To enable _ disable sending only over http
        """
        return pulumi.get(self, "httponly")

    @property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> Optional[str]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> Optional[str]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> Optional[str]:
        """
        To enable _ disable match across virtual servers with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @property
    @pulumi.getter
    def mirror(self) -> Optional[str]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> Optional[str]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[float]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

