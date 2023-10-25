# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 login_ref: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 teem_disable: Optional[pulumi.Input[bool]] = None,
                 token_auth: Optional[pulumi.Input[bool]] = None,
                 token_value: Optional[pulumi.Input[str]] = None,
                 trusted_cert_path: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 validate_certs_disable: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] address: Domain name/IP of the BigIP
        :param pulumi.Input[str] login_ref: Login reference for token authentication (see BIG-IP REST docs for details)
        :param pulumi.Input[str] password: The user's password. Leave empty if using token_value
        :param pulumi.Input[str] port: Management Port to connect to Bigip
        :param pulumi.Input[bool] teem_disable: If this flag set to true,sending telemetry data to TEEM will be disabled
        :param pulumi.Input[bool] token_auth: Enable to use an external authentication source (LDAP, TACACS, etc)
        :param pulumi.Input[str] token_value: A token generated outside the provider, in place of password
        :param pulumi.Input[str] trusted_cert_path: Valid Trusted Certificate path
        :param pulumi.Input[str] username: Username with API access to the BigIP
        :param pulumi.Input[bool] validate_certs_disable: If set to true, Disables TLS certificate check on BIG-IP. Default : True
        """
        ProviderArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address=address,
            login_ref=login_ref,
            password=password,
            port=port,
            teem_disable=teem_disable,
            token_auth=token_auth,
            token_value=token_value,
            trusted_cert_path=trusted_cert_path,
            username=username,
            validate_certs_disable=validate_certs_disable,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address: Optional[pulumi.Input[str]] = None,
             login_ref: Optional[pulumi.Input[str]] = None,
             password: Optional[pulumi.Input[str]] = None,
             port: Optional[pulumi.Input[str]] = None,
             teem_disable: Optional[pulumi.Input[bool]] = None,
             token_auth: Optional[pulumi.Input[bool]] = None,
             token_value: Optional[pulumi.Input[str]] = None,
             trusted_cert_path: Optional[pulumi.Input[str]] = None,
             username: Optional[pulumi.Input[str]] = None,
             validate_certs_disable: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if login_ref is None and 'loginRef' in kwargs:
            login_ref = kwargs['loginRef']
        if teem_disable is None and 'teemDisable' in kwargs:
            teem_disable = kwargs['teemDisable']
        if token_auth is None and 'tokenAuth' in kwargs:
            token_auth = kwargs['tokenAuth']
        if token_value is None and 'tokenValue' in kwargs:
            token_value = kwargs['tokenValue']
        if trusted_cert_path is None and 'trustedCertPath' in kwargs:
            trusted_cert_path = kwargs['trustedCertPath']
        if validate_certs_disable is None and 'validateCertsDisable' in kwargs:
            validate_certs_disable = kwargs['validateCertsDisable']

        if address is not None:
            _setter("address", address)
        if login_ref is not None:
            _setter("login_ref", login_ref)
        if password is not None:
            _setter("password", password)
        if port is not None:
            _setter("port", port)
        if teem_disable is not None:
            _setter("teem_disable", teem_disable)
        if token_auth is not None:
            _setter("token_auth", token_auth)
        if token_value is not None:
            _setter("token_value", token_value)
        if trusted_cert_path is not None:
            _setter("trusted_cert_path", trusted_cert_path)
        if username is not None:
            _setter("username", username)
        if validate_certs_disable is not None:
            _setter("validate_certs_disable", validate_certs_disable)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Domain name/IP of the BigIP
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="loginRef")
    def login_ref(self) -> Optional[pulumi.Input[str]]:
        """
        Login reference for token authentication (see BIG-IP REST docs for details)
        """
        return pulumi.get(self, "login_ref")

    @login_ref.setter
    def login_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_ref", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The user's password. Leave empty if using token_value
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        Management Port to connect to Bigip
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="teemDisable")
    def teem_disable(self) -> Optional[pulumi.Input[bool]]:
        """
        If this flag set to true,sending telemetry data to TEEM will be disabled
        """
        return pulumi.get(self, "teem_disable")

    @teem_disable.setter
    def teem_disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "teem_disable", value)

    @property
    @pulumi.getter(name="tokenAuth")
    def token_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable to use an external authentication source (LDAP, TACACS, etc)
        """
        return pulumi.get(self, "token_auth")

    @token_auth.setter
    def token_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "token_auth", value)

    @property
    @pulumi.getter(name="tokenValue")
    def token_value(self) -> Optional[pulumi.Input[str]]:
        """
        A token generated outside the provider, in place of password
        """
        return pulumi.get(self, "token_value")

    @token_value.setter
    def token_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_value", value)

    @property
    @pulumi.getter(name="trustedCertPath")
    def trusted_cert_path(self) -> Optional[pulumi.Input[str]]:
        """
        Valid Trusted Certificate path
        """
        return pulumi.get(self, "trusted_cert_path")

    @trusted_cert_path.setter
    def trusted_cert_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trusted_cert_path", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Username with API access to the BigIP
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="validateCertsDisable")
    def validate_certs_disable(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, Disables TLS certificate check on BIG-IP. Default : True
        """
        return pulumi.get(self, "validate_certs_disable")

    @validate_certs_disable.setter
    def validate_certs_disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_certs_disable", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 login_ref: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 teem_disable: Optional[pulumi.Input[bool]] = None,
                 token_auth: Optional[pulumi.Input[bool]] = None,
                 token_value: Optional[pulumi.Input[str]] = None,
                 trusted_cert_path: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 validate_certs_disable: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The provider type for the bigip package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Domain name/IP of the BigIP
        :param pulumi.Input[str] login_ref: Login reference for token authentication (see BIG-IP REST docs for details)
        :param pulumi.Input[str] password: The user's password. Leave empty if using token_value
        :param pulumi.Input[str] port: Management Port to connect to Bigip
        :param pulumi.Input[bool] teem_disable: If this flag set to true,sending telemetry data to TEEM will be disabled
        :param pulumi.Input[bool] token_auth: Enable to use an external authentication source (LDAP, TACACS, etc)
        :param pulumi.Input[str] token_value: A token generated outside the provider, in place of password
        :param pulumi.Input[str] trusted_cert_path: Valid Trusted Certificate path
        :param pulumi.Input[str] username: Username with API access to the BigIP
        :param pulumi.Input[bool] validate_certs_disable: If set to true, Disables TLS certificate check on BIG-IP. Default : True
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the bigip package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProviderArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 login_ref: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 teem_disable: Optional[pulumi.Input[bool]] = None,
                 token_auth: Optional[pulumi.Input[bool]] = None,
                 token_value: Optional[pulumi.Input[str]] = None,
                 trusted_cert_path: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 validate_certs_disable: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["address"] = address
            __props__.__dict__["login_ref"] = login_ref
            __props__.__dict__["password"] = password
            __props__.__dict__["port"] = port
            __props__.__dict__["teem_disable"] = pulumi.Output.from_input(teem_disable).apply(pulumi.runtime.to_json) if teem_disable is not None else None
            __props__.__dict__["token_auth"] = pulumi.Output.from_input(token_auth).apply(pulumi.runtime.to_json) if token_auth is not None else None
            __props__.__dict__["token_value"] = token_value
            __props__.__dict__["trusted_cert_path"] = trusted_cert_path
            __props__.__dict__["username"] = username
            __props__.__dict__["validate_certs_disable"] = pulumi.Output.from_input(validate_certs_disable).apply(pulumi.runtime.to_json) if validate_certs_disable is not None else None
        super(Provider, __self__).__init__(
            'f5bigip',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[Optional[str]]:
        """
        Domain name/IP of the BigIP
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="loginRef")
    def login_ref(self) -> pulumi.Output[Optional[str]]:
        """
        Login reference for token authentication (see BIG-IP REST docs for details)
        """
        return pulumi.get(self, "login_ref")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The user's password. Leave empty if using token_value
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[str]]:
        """
        Management Port to connect to Bigip
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="tokenValue")
    def token_value(self) -> pulumi.Output[Optional[str]]:
        """
        A token generated outside the provider, in place of password
        """
        return pulumi.get(self, "token_value")

    @property
    @pulumi.getter(name="trustedCertPath")
    def trusted_cert_path(self) -> pulumi.Output[Optional[str]]:
        """
        Valid Trusted Certificate path
        """
        return pulumi.get(self, "trusted_cert_path")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        Username with API access to the BigIP
        """
        return pulumi.get(self, "username")

