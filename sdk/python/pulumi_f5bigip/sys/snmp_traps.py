# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SnmpTrapsArgs', 'SnmpTraps']

@pulumi.input_type
class SnmpTrapsArgs:
    def __init__(__self__, *,
                 auth_passwordencrypted: Optional[pulumi.Input[str]] = None,
                 auth_protocol: Optional[pulumi.Input[str]] = None,
                 community: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 privacy_password: Optional[pulumi.Input[str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[str]] = None,
                 privacy_protocol: Optional[pulumi.Input[str]] = None,
                 security_level: Optional[pulumi.Input[str]] = None,
                 security_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SnmpTraps resource.
        :param pulumi.Input[str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[str] community: Specifies the community string used for this trap.
        :param pulumi.Input[str] description: The port that the trap will be sent to.
        :param pulumi.Input[str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[str] host: The host the trap will be sent to.
        :param pulumi.Input[str] name: Name of the snmp trap.
        :param pulumi.Input[int] port: User defined description.
        :param pulumi.Input[str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[str] version: SNMP version used for sending the trap.
        """
        SnmpTrapsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auth_passwordencrypted=auth_passwordencrypted,
            auth_protocol=auth_protocol,
            community=community,
            description=description,
            engine_id=engine_id,
            host=host,
            name=name,
            port=port,
            privacy_password=privacy_password,
            privacy_password_encrypted=privacy_password_encrypted,
            privacy_protocol=privacy_protocol,
            security_level=security_level,
            security_name=security_name,
            version=version,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auth_passwordencrypted: Optional[pulumi.Input[str]] = None,
             auth_protocol: Optional[pulumi.Input[str]] = None,
             community: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             engine_id: Optional[pulumi.Input[str]] = None,
             host: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             port: Optional[pulumi.Input[int]] = None,
             privacy_password: Optional[pulumi.Input[str]] = None,
             privacy_password_encrypted: Optional[pulumi.Input[str]] = None,
             privacy_protocol: Optional[pulumi.Input[str]] = None,
             security_level: Optional[pulumi.Input[str]] = None,
             security_name: Optional[pulumi.Input[str]] = None,
             version: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if auth_passwordencrypted is None and 'authPasswordencrypted' in kwargs:
            auth_passwordencrypted = kwargs['authPasswordencrypted']
        if auth_protocol is None and 'authProtocol' in kwargs:
            auth_protocol = kwargs['authProtocol']
        if engine_id is None and 'engineId' in kwargs:
            engine_id = kwargs['engineId']
        if privacy_password is None and 'privacyPassword' in kwargs:
            privacy_password = kwargs['privacyPassword']
        if privacy_password_encrypted is None and 'privacyPasswordEncrypted' in kwargs:
            privacy_password_encrypted = kwargs['privacyPasswordEncrypted']
        if privacy_protocol is None and 'privacyProtocol' in kwargs:
            privacy_protocol = kwargs['privacyProtocol']
        if security_level is None and 'securityLevel' in kwargs:
            security_level = kwargs['securityLevel']
        if security_name is None and 'securityName' in kwargs:
            security_name = kwargs['securityName']

        if auth_passwordencrypted is not None:
            _setter("auth_passwordencrypted", auth_passwordencrypted)
        if auth_protocol is not None:
            _setter("auth_protocol", auth_protocol)
        if community is not None:
            _setter("community", community)
        if description is not None:
            _setter("description", description)
        if engine_id is not None:
            _setter("engine_id", engine_id)
        if host is not None:
            _setter("host", host)
        if name is not None:
            _setter("name", name)
        if port is not None:
            _setter("port", port)
        if privacy_password is not None:
            _setter("privacy_password", privacy_password)
        if privacy_password_encrypted is not None:
            _setter("privacy_password_encrypted", privacy_password_encrypted)
        if privacy_protocol is not None:
            _setter("privacy_protocol", privacy_protocol)
        if security_level is not None:
            _setter("security_level", security_level)
        if security_name is not None:
            _setter("security_name", security_name)
        if version is not None:
            _setter("version", version)

    @property
    @pulumi.getter(name="authPasswordencrypted")
    def auth_passwordencrypted(self) -> Optional[pulumi.Input[str]]:
        """
        Encrypted password
        """
        return pulumi.get(self, "auth_passwordencrypted")

    @auth_passwordencrypted.setter
    def auth_passwordencrypted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_passwordencrypted", value)

    @property
    @pulumi.getter(name="authProtocol")
    def auth_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the protocol used to authenticate the user.
        """
        return pulumi.get(self, "auth_protocol")

    @auth_protocol.setter
    def auth_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_protocol", value)

    @property
    @pulumi.getter
    def community(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the community string used for this trap.
        """
        return pulumi.get(self, "community")

    @community.setter
    def community(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "community", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The port that the trap will be sent to.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the authoritative security engine for SNMPv3.
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The host the trap will be sent to.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the snmp trap.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        User defined description.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privacyPassword")
    def privacy_password(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        """
        return pulumi.get(self, "privacy_password")

    @privacy_password.setter
    def privacy_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_password", value)

    @property
    @pulumi.getter(name="privacyPasswordEncrypted")
    def privacy_password_encrypted(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the encrypted password used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_password_encrypted")

    @privacy_password_encrypted.setter
    def privacy_password_encrypted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_password_encrypted", value)

    @property
    @pulumi.getter(name="privacyProtocol")
    def privacy_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the protocol used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_protocol")

    @privacy_protocol.setter
    def privacy_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_protocol", value)

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether or not traffic is encrypted and whether or not authentication is required.
        """
        return pulumi.get(self, "security_level")

    @security_level.setter
    def security_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_level", value)

    @property
    @pulumi.getter(name="securityName")
    def security_name(self) -> Optional[pulumi.Input[str]]:
        """
        Security name used in conjunction with SNMPv3.
        """
        return pulumi.get(self, "security_name")

    @security_name.setter
    def security_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        SNMP version used for sending the trap.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _SnmpTrapsState:
    def __init__(__self__, *,
                 auth_passwordencrypted: Optional[pulumi.Input[str]] = None,
                 auth_protocol: Optional[pulumi.Input[str]] = None,
                 community: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 privacy_password: Optional[pulumi.Input[str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[str]] = None,
                 privacy_protocol: Optional[pulumi.Input[str]] = None,
                 security_level: Optional[pulumi.Input[str]] = None,
                 security_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SnmpTraps resources.
        :param pulumi.Input[str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[str] community: Specifies the community string used for this trap.
        :param pulumi.Input[str] description: The port that the trap will be sent to.
        :param pulumi.Input[str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[str] host: The host the trap will be sent to.
        :param pulumi.Input[str] name: Name of the snmp trap.
        :param pulumi.Input[int] port: User defined description.
        :param pulumi.Input[str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[str] version: SNMP version used for sending the trap.
        """
        _SnmpTrapsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auth_passwordencrypted=auth_passwordencrypted,
            auth_protocol=auth_protocol,
            community=community,
            description=description,
            engine_id=engine_id,
            host=host,
            name=name,
            port=port,
            privacy_password=privacy_password,
            privacy_password_encrypted=privacy_password_encrypted,
            privacy_protocol=privacy_protocol,
            security_level=security_level,
            security_name=security_name,
            version=version,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auth_passwordencrypted: Optional[pulumi.Input[str]] = None,
             auth_protocol: Optional[pulumi.Input[str]] = None,
             community: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             engine_id: Optional[pulumi.Input[str]] = None,
             host: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             port: Optional[pulumi.Input[int]] = None,
             privacy_password: Optional[pulumi.Input[str]] = None,
             privacy_password_encrypted: Optional[pulumi.Input[str]] = None,
             privacy_protocol: Optional[pulumi.Input[str]] = None,
             security_level: Optional[pulumi.Input[str]] = None,
             security_name: Optional[pulumi.Input[str]] = None,
             version: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if auth_passwordencrypted is None and 'authPasswordencrypted' in kwargs:
            auth_passwordencrypted = kwargs['authPasswordencrypted']
        if auth_protocol is None and 'authProtocol' in kwargs:
            auth_protocol = kwargs['authProtocol']
        if engine_id is None and 'engineId' in kwargs:
            engine_id = kwargs['engineId']
        if privacy_password is None and 'privacyPassword' in kwargs:
            privacy_password = kwargs['privacyPassword']
        if privacy_password_encrypted is None and 'privacyPasswordEncrypted' in kwargs:
            privacy_password_encrypted = kwargs['privacyPasswordEncrypted']
        if privacy_protocol is None and 'privacyProtocol' in kwargs:
            privacy_protocol = kwargs['privacyProtocol']
        if security_level is None and 'securityLevel' in kwargs:
            security_level = kwargs['securityLevel']
        if security_name is None and 'securityName' in kwargs:
            security_name = kwargs['securityName']

        if auth_passwordencrypted is not None:
            _setter("auth_passwordencrypted", auth_passwordencrypted)
        if auth_protocol is not None:
            _setter("auth_protocol", auth_protocol)
        if community is not None:
            _setter("community", community)
        if description is not None:
            _setter("description", description)
        if engine_id is not None:
            _setter("engine_id", engine_id)
        if host is not None:
            _setter("host", host)
        if name is not None:
            _setter("name", name)
        if port is not None:
            _setter("port", port)
        if privacy_password is not None:
            _setter("privacy_password", privacy_password)
        if privacy_password_encrypted is not None:
            _setter("privacy_password_encrypted", privacy_password_encrypted)
        if privacy_protocol is not None:
            _setter("privacy_protocol", privacy_protocol)
        if security_level is not None:
            _setter("security_level", security_level)
        if security_name is not None:
            _setter("security_name", security_name)
        if version is not None:
            _setter("version", version)

    @property
    @pulumi.getter(name="authPasswordencrypted")
    def auth_passwordencrypted(self) -> Optional[pulumi.Input[str]]:
        """
        Encrypted password
        """
        return pulumi.get(self, "auth_passwordencrypted")

    @auth_passwordencrypted.setter
    def auth_passwordencrypted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_passwordencrypted", value)

    @property
    @pulumi.getter(name="authProtocol")
    def auth_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the protocol used to authenticate the user.
        """
        return pulumi.get(self, "auth_protocol")

    @auth_protocol.setter
    def auth_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_protocol", value)

    @property
    @pulumi.getter
    def community(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the community string used for this trap.
        """
        return pulumi.get(self, "community")

    @community.setter
    def community(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "community", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The port that the trap will be sent to.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the authoritative security engine for SNMPv3.
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The host the trap will be sent to.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the snmp trap.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        User defined description.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privacyPassword")
    def privacy_password(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        """
        return pulumi.get(self, "privacy_password")

    @privacy_password.setter
    def privacy_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_password", value)

    @property
    @pulumi.getter(name="privacyPasswordEncrypted")
    def privacy_password_encrypted(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the encrypted password used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_password_encrypted")

    @privacy_password_encrypted.setter
    def privacy_password_encrypted(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_password_encrypted", value)

    @property
    @pulumi.getter(name="privacyProtocol")
    def privacy_protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the protocol used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_protocol")

    @privacy_protocol.setter
    def privacy_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy_protocol", value)

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether or not traffic is encrypted and whether or not authentication is required.
        """
        return pulumi.get(self, "security_level")

    @security_level.setter
    def security_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_level", value)

    @property
    @pulumi.getter(name="securityName")
    def security_name(self) -> Optional[pulumi.Input[str]]:
        """
        Security name used in conjunction with SNMPv3.
        """
        return pulumi.get(self, "security_name")

    @security_name.setter
    def security_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        SNMP version used for sending the trap.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class SnmpTraps(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_passwordencrypted: Optional[pulumi.Input[str]] = None,
                 auth_protocol: Optional[pulumi.Input[str]] = None,
                 community: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 privacy_password: Optional[pulumi.Input[str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[str]] = None,
                 privacy_protocol: Optional[pulumi.Input[str]] = None,
                 security_level: Optional[pulumi.Input[str]] = None,
                 security_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `sys.SnmpTraps` provides details bout how to enable snmp_traps resource on BIG-IP

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[str] community: Specifies the community string used for this trap.
        :param pulumi.Input[str] description: The port that the trap will be sent to.
        :param pulumi.Input[str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[str] host: The host the trap will be sent to.
        :param pulumi.Input[str] name: Name of the snmp trap.
        :param pulumi.Input[int] port: User defined description.
        :param pulumi.Input[str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[str] version: SNMP version used for sending the trap.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SnmpTrapsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `sys.SnmpTraps` provides details bout how to enable snmp_traps resource on BIG-IP

        :param str resource_name: The name of the resource.
        :param SnmpTrapsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnmpTrapsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            SnmpTrapsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_passwordencrypted: Optional[pulumi.Input[str]] = None,
                 auth_protocol: Optional[pulumi.Input[str]] = None,
                 community: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 privacy_password: Optional[pulumi.Input[str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[str]] = None,
                 privacy_protocol: Optional[pulumi.Input[str]] = None,
                 security_level: Optional[pulumi.Input[str]] = None,
                 security_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnmpTrapsArgs.__new__(SnmpTrapsArgs)

            __props__.__dict__["auth_passwordencrypted"] = auth_passwordencrypted
            __props__.__dict__["auth_protocol"] = auth_protocol
            __props__.__dict__["community"] = community
            __props__.__dict__["description"] = description
            __props__.__dict__["engine_id"] = engine_id
            __props__.__dict__["host"] = host
            __props__.__dict__["name"] = name
            __props__.__dict__["port"] = port
            __props__.__dict__["privacy_password"] = privacy_password
            __props__.__dict__["privacy_password_encrypted"] = privacy_password_encrypted
            __props__.__dict__["privacy_protocol"] = privacy_protocol
            __props__.__dict__["security_level"] = security_level
            __props__.__dict__["security_name"] = security_name
            __props__.__dict__["version"] = version
        super(SnmpTraps, __self__).__init__(
            'f5bigip:sys/snmpTraps:SnmpTraps',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_passwordencrypted: Optional[pulumi.Input[str]] = None,
            auth_protocol: Optional[pulumi.Input[str]] = None,
            community: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            engine_id: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            privacy_password: Optional[pulumi.Input[str]] = None,
            privacy_password_encrypted: Optional[pulumi.Input[str]] = None,
            privacy_protocol: Optional[pulumi.Input[str]] = None,
            security_level: Optional[pulumi.Input[str]] = None,
            security_name: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'SnmpTraps':
        """
        Get an existing SnmpTraps resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[str] community: Specifies the community string used for this trap.
        :param pulumi.Input[str] description: The port that the trap will be sent to.
        :param pulumi.Input[str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[str] host: The host the trap will be sent to.
        :param pulumi.Input[str] name: Name of the snmp trap.
        :param pulumi.Input[int] port: User defined description.
        :param pulumi.Input[str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[str] version: SNMP version used for sending the trap.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnmpTrapsState.__new__(_SnmpTrapsState)

        __props__.__dict__["auth_passwordencrypted"] = auth_passwordencrypted
        __props__.__dict__["auth_protocol"] = auth_protocol
        __props__.__dict__["community"] = community
        __props__.__dict__["description"] = description
        __props__.__dict__["engine_id"] = engine_id
        __props__.__dict__["host"] = host
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["privacy_password"] = privacy_password
        __props__.__dict__["privacy_password_encrypted"] = privacy_password_encrypted
        __props__.__dict__["privacy_protocol"] = privacy_protocol
        __props__.__dict__["security_level"] = security_level
        __props__.__dict__["security_name"] = security_name
        __props__.__dict__["version"] = version
        return SnmpTraps(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authPasswordencrypted")
    def auth_passwordencrypted(self) -> pulumi.Output[Optional[str]]:
        """
        Encrypted password
        """
        return pulumi.get(self, "auth_passwordencrypted")

    @property
    @pulumi.getter(name="authProtocol")
    def auth_protocol(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the protocol used to authenticate the user.
        """
        return pulumi.get(self, "auth_protocol")

    @property
    @pulumi.getter
    def community(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the community string used for this trap.
        """
        return pulumi.get(self, "community")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The port that the trap will be sent to.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the authoritative security engine for SNMPv3.
        """
        return pulumi.get(self, "engine_id")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[str]]:
        """
        The host the trap will be sent to.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the snmp trap.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        User defined description.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privacyPassword")
    def privacy_password(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        """
        return pulumi.get(self, "privacy_password")

    @property
    @pulumi.getter(name="privacyPasswordEncrypted")
    def privacy_password_encrypted(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the encrypted password used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_password_encrypted")

    @property
    @pulumi.getter(name="privacyProtocol")
    def privacy_protocol(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the protocol used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_protocol")

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> pulumi.Output[str]:
        """
        Specifies whether or not traffic is encrypted and whether or not authentication is required.
        """
        return pulumi.get(self, "security_level")

    @property
    @pulumi.getter(name="securityName")
    def security_name(self) -> pulumi.Output[Optional[str]]:
        """
        Security name used in conjunction with SNMPv3.
        """
        return pulumi.get(self, "security_name")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        SNMP version used for sending the trap.
        """
        return pulumi.get(self, "version")

