# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['SnmpTrapsArgs', 'SnmpTraps']

@pulumi.input_type
class SnmpTrapsArgs:
    def __init__(__self__, *,
                 auth_passwordencrypted: Optional[pulumi.Input[builtins.str]] = None,
                 auth_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 community: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 engine_id: Optional[pulumi.Input[builtins.str]] = None,
                 host: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 privacy_password: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 security_level: Optional[pulumi.Input[builtins.str]] = None,
                 security_name: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a SnmpTraps resource.
        :param pulumi.Input[builtins.str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[builtins.str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[builtins.str] community: Specifies the community string used for this trap.
        :param pulumi.Input[builtins.str] description: The port that the trap will be sent to.
        :param pulumi.Input[builtins.str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[builtins.str] host: The host the trap will be sent to.
        :param pulumi.Input[builtins.str] name: Name of the snmp trap.
        :param pulumi.Input[builtins.int] port: User defined description.
        :param pulumi.Input[builtins.str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[builtins.str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[builtins.str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[builtins.str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[builtins.str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[builtins.str] version: SNMP version used for sending the trap.
        """
        if auth_passwordencrypted is not None:
            pulumi.set(__self__, "auth_passwordencrypted", auth_passwordencrypted)
        if auth_protocol is not None:
            pulumi.set(__self__, "auth_protocol", auth_protocol)
        if community is not None:
            pulumi.set(__self__, "community", community)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_id is not None:
            pulumi.set(__self__, "engine_id", engine_id)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if privacy_password is not None:
            pulumi.set(__self__, "privacy_password", privacy_password)
        if privacy_password_encrypted is not None:
            pulumi.set(__self__, "privacy_password_encrypted", privacy_password_encrypted)
        if privacy_protocol is not None:
            pulumi.set(__self__, "privacy_protocol", privacy_protocol)
        if security_level is not None:
            pulumi.set(__self__, "security_level", security_level)
        if security_name is not None:
            pulumi.set(__self__, "security_name", security_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="authPasswordencrypted")
    def auth_passwordencrypted(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Encrypted password
        """
        return pulumi.get(self, "auth_passwordencrypted")

    @auth_passwordencrypted.setter
    def auth_passwordencrypted(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auth_passwordencrypted", value)

    @property
    @pulumi.getter(name="authProtocol")
    def auth_protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the protocol used to authenticate the user.
        """
        return pulumi.get(self, "auth_protocol")

    @auth_protocol.setter
    def auth_protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auth_protocol", value)

    @property
    @pulumi.getter
    def community(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the community string used for this trap.
        """
        return pulumi.get(self, "community")

    @community.setter
    def community(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "community", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The port that the trap will be sent to.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the authoritative security engine for SNMPv3.
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The host the trap will be sent to.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the snmp trap.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        User defined description.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privacyPassword")
    def privacy_password(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        """
        return pulumi.get(self, "privacy_password")

    @privacy_password.setter
    def privacy_password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "privacy_password", value)

    @property
    @pulumi.getter(name="privacyPasswordEncrypted")
    def privacy_password_encrypted(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the encrypted password used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_password_encrypted")

    @privacy_password_encrypted.setter
    def privacy_password_encrypted(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "privacy_password_encrypted", value)

    @property
    @pulumi.getter(name="privacyProtocol")
    def privacy_protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the protocol used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_protocol")

    @privacy_protocol.setter
    def privacy_protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "privacy_protocol", value)

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies whether or not traffic is encrypted and whether or not authentication is required.
        """
        return pulumi.get(self, "security_level")

    @security_level.setter
    def security_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_level", value)

    @property
    @pulumi.getter(name="securityName")
    def security_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Security name used in conjunction with SNMPv3.
        """
        return pulumi.get(self, "security_name")

    @security_name.setter
    def security_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        SNMP version used for sending the trap.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _SnmpTrapsState:
    def __init__(__self__, *,
                 auth_passwordencrypted: Optional[pulumi.Input[builtins.str]] = None,
                 auth_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 community: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 engine_id: Optional[pulumi.Input[builtins.str]] = None,
                 host: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 privacy_password: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 security_level: Optional[pulumi.Input[builtins.str]] = None,
                 security_name: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering SnmpTraps resources.
        :param pulumi.Input[builtins.str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[builtins.str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[builtins.str] community: Specifies the community string used for this trap.
        :param pulumi.Input[builtins.str] description: The port that the trap will be sent to.
        :param pulumi.Input[builtins.str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[builtins.str] host: The host the trap will be sent to.
        :param pulumi.Input[builtins.str] name: Name of the snmp trap.
        :param pulumi.Input[builtins.int] port: User defined description.
        :param pulumi.Input[builtins.str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[builtins.str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[builtins.str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[builtins.str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[builtins.str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[builtins.str] version: SNMP version used for sending the trap.
        """
        if auth_passwordencrypted is not None:
            pulumi.set(__self__, "auth_passwordencrypted", auth_passwordencrypted)
        if auth_protocol is not None:
            pulumi.set(__self__, "auth_protocol", auth_protocol)
        if community is not None:
            pulumi.set(__self__, "community", community)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_id is not None:
            pulumi.set(__self__, "engine_id", engine_id)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if privacy_password is not None:
            pulumi.set(__self__, "privacy_password", privacy_password)
        if privacy_password_encrypted is not None:
            pulumi.set(__self__, "privacy_password_encrypted", privacy_password_encrypted)
        if privacy_protocol is not None:
            pulumi.set(__self__, "privacy_protocol", privacy_protocol)
        if security_level is not None:
            pulumi.set(__self__, "security_level", security_level)
        if security_name is not None:
            pulumi.set(__self__, "security_name", security_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="authPasswordencrypted")
    def auth_passwordencrypted(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Encrypted password
        """
        return pulumi.get(self, "auth_passwordencrypted")

    @auth_passwordencrypted.setter
    def auth_passwordencrypted(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auth_passwordencrypted", value)

    @property
    @pulumi.getter(name="authProtocol")
    def auth_protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the protocol used to authenticate the user.
        """
        return pulumi.get(self, "auth_protocol")

    @auth_protocol.setter
    def auth_protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auth_protocol", value)

    @property
    @pulumi.getter
    def community(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the community string used for this trap.
        """
        return pulumi.get(self, "community")

    @community.setter
    def community(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "community", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The port that the trap will be sent to.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the authoritative security engine for SNMPv3.
        """
        return pulumi.get(self, "engine_id")

    @engine_id.setter
    def engine_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "engine_id", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The host the trap will be sent to.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the snmp trap.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        User defined description.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="privacyPassword")
    def privacy_password(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        """
        return pulumi.get(self, "privacy_password")

    @privacy_password.setter
    def privacy_password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "privacy_password", value)

    @property
    @pulumi.getter(name="privacyPasswordEncrypted")
    def privacy_password_encrypted(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the encrypted password used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_password_encrypted")

    @privacy_password_encrypted.setter
    def privacy_password_encrypted(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "privacy_password_encrypted", value)

    @property
    @pulumi.getter(name="privacyProtocol")
    def privacy_protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the protocol used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_protocol")

    @privacy_protocol.setter
    def privacy_protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "privacy_protocol", value)

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies whether or not traffic is encrypted and whether or not authentication is required.
        """
        return pulumi.get(self, "security_level")

    @security_level.setter
    def security_level(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_level", value)

    @property
    @pulumi.getter(name="securityName")
    def security_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Security name used in conjunction with SNMPv3.
        """
        return pulumi.get(self, "security_name")

    @security_name.setter
    def security_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "security_name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        SNMP version used for sending the trap.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "version", value)


class SnmpTraps(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_passwordencrypted: Optional[pulumi.Input[builtins.str]] = None,
                 auth_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 community: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 engine_id: Optional[pulumi.Input[builtins.str]] = None,
                 host: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 privacy_password: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 security_level: Optional[pulumi.Input[builtins.str]] = None,
                 security_name: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        `sys.SnmpTraps` provides details bout how to enable snmp_traps resource on BIG-IP
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        snmp_traps = f5bigip.sys.SnmpTraps("snmp_traps",
            name="snmptraps",
            community="f5community",
            host="195.10.10.1",
            description="Setup snmp traps",
            port=111)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[builtins.str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[builtins.str] community: Specifies the community string used for this trap.
        :param pulumi.Input[builtins.str] description: The port that the trap will be sent to.
        :param pulumi.Input[builtins.str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[builtins.str] host: The host the trap will be sent to.
        :param pulumi.Input[builtins.str] name: Name of the snmp trap.
        :param pulumi.Input[builtins.int] port: User defined description.
        :param pulumi.Input[builtins.str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[builtins.str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[builtins.str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[builtins.str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[builtins.str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[builtins.str] version: SNMP version used for sending the trap.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SnmpTrapsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `sys.SnmpTraps` provides details bout how to enable snmp_traps resource on BIG-IP
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        snmp_traps = f5bigip.sys.SnmpTraps("snmp_traps",
            name="snmptraps",
            community="f5community",
            host="195.10.10.1",
            description="Setup snmp traps",
            port=111)
        ```

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
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_passwordencrypted: Optional[pulumi.Input[builtins.str]] = None,
                 auth_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 community: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 engine_id: Optional[pulumi.Input[builtins.str]] = None,
                 host: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 port: Optional[pulumi.Input[builtins.int]] = None,
                 privacy_password: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_password_encrypted: Optional[pulumi.Input[builtins.str]] = None,
                 privacy_protocol: Optional[pulumi.Input[builtins.str]] = None,
                 security_level: Optional[pulumi.Input[builtins.str]] = None,
                 security_name: Optional[pulumi.Input[builtins.str]] = None,
                 version: Optional[pulumi.Input[builtins.str]] = None,
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
            auth_passwordencrypted: Optional[pulumi.Input[builtins.str]] = None,
            auth_protocol: Optional[pulumi.Input[builtins.str]] = None,
            community: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            engine_id: Optional[pulumi.Input[builtins.str]] = None,
            host: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            port: Optional[pulumi.Input[builtins.int]] = None,
            privacy_password: Optional[pulumi.Input[builtins.str]] = None,
            privacy_password_encrypted: Optional[pulumi.Input[builtins.str]] = None,
            privacy_protocol: Optional[pulumi.Input[builtins.str]] = None,
            security_level: Optional[pulumi.Input[builtins.str]] = None,
            security_name: Optional[pulumi.Input[builtins.str]] = None,
            version: Optional[pulumi.Input[builtins.str]] = None) -> 'SnmpTraps':
        """
        Get an existing SnmpTraps resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] auth_passwordencrypted: Encrypted password
        :param pulumi.Input[builtins.str] auth_protocol: Specifies the protocol used to authenticate the user.
        :param pulumi.Input[builtins.str] community: Specifies the community string used for this trap.
        :param pulumi.Input[builtins.str] description: The port that the trap will be sent to.
        :param pulumi.Input[builtins.str] engine_id: Specifies the authoritative security engine for SNMPv3.
        :param pulumi.Input[builtins.str] host: The host the trap will be sent to.
        :param pulumi.Input[builtins.str] name: Name of the snmp trap.
        :param pulumi.Input[builtins.int] port: User defined description.
        :param pulumi.Input[builtins.str] privacy_password: Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        :param pulumi.Input[builtins.str] privacy_password_encrypted: Specifies the encrypted password used to encrypt traffic.
        :param pulumi.Input[builtins.str] privacy_protocol: Specifies the protocol used to encrypt traffic.
        :param pulumi.Input[builtins.str] security_level: Specifies whether or not traffic is encrypted and whether or not authentication is required.
        :param pulumi.Input[builtins.str] security_name: Security name used in conjunction with SNMPv3.
        :param pulumi.Input[builtins.str] version: SNMP version used for sending the trap.
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
    def auth_passwordencrypted(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Encrypted password
        """
        return pulumi.get(self, "auth_passwordencrypted")

    @property
    @pulumi.getter(name="authProtocol")
    def auth_protocol(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the protocol used to authenticate the user.
        """
        return pulumi.get(self, "auth_protocol")

    @property
    @pulumi.getter
    def community(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the community string used for this trap.
        """
        return pulumi.get(self, "community")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The port that the trap will be sent to.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="engineId")
    def engine_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the authoritative security engine for SNMPv3.
        """
        return pulumi.get(self, "engine_id")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The host the trap will be sent to.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Name of the snmp trap.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        User defined description.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privacyPassword")
    def privacy_password(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the clear text password used to encrypt traffic. This field will not be displayed.
        """
        return pulumi.get(self, "privacy_password")

    @property
    @pulumi.getter(name="privacyPasswordEncrypted")
    def privacy_password_encrypted(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the encrypted password used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_password_encrypted")

    @property
    @pulumi.getter(name="privacyProtocol")
    def privacy_protocol(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the protocol used to encrypt traffic.
        """
        return pulumi.get(self, "privacy_protocol")

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> pulumi.Output[builtins.str]:
        """
        Specifies whether or not traffic is encrypted and whether or not authentication is required.
        """
        return pulumi.get(self, "security_level")

    @property
    @pulumi.getter(name="securityName")
    def security_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Security name used in conjunction with SNMPv3.
        """
        return pulumi.get(self, "security_name")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[builtins.str]:
        """
        SNMP version used for sending the trap.
        """
        return pulumi.get(self, "version")

