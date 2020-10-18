# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SnmpTraps']


class SnmpTraps(pulumi.CustomResource):
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
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `sys.SnmpTraps` provides details bout how to enable snmp_traps resource on BIG-IP
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        snmp_traps = f5bigip.sys.SnmpTraps("snmpTraps",
            community="f5community",
            description="Setup snmp traps",
            host="195.10.10.1",
            name="snmptraps",
            port=111)
        ```

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

            __props__['auth_passwordencrypted'] = auth_passwordencrypted
            __props__['auth_protocol'] = auth_protocol
            __props__['community'] = community
            __props__['description'] = description
            __props__['engine_id'] = engine_id
            __props__['host'] = host
            __props__['name'] = name
            __props__['port'] = port
            __props__['privacy_password'] = privacy_password
            __props__['privacy_password_encrypted'] = privacy_password_encrypted
            __props__['privacy_protocol'] = privacy_protocol
            __props__['security_level'] = security_level
            __props__['security_name'] = security_name
            __props__['version'] = version
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

        __props__ = dict()

        __props__["auth_passwordencrypted"] = auth_passwordencrypted
        __props__["auth_protocol"] = auth_protocol
        __props__["community"] = community
        __props__["description"] = description
        __props__["engine_id"] = engine_id
        __props__["host"] = host
        __props__["name"] = name
        __props__["port"] = port
        __props__["privacy_password"] = privacy_password
        __props__["privacy_password_encrypted"] = privacy_password_encrypted
        __props__["privacy_protocol"] = privacy_protocol
        __props__["security_level"] = security_level
        __props__["security_name"] = security_name
        __props__["version"] = version
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
    def security_level(self) -> pulumi.Output[Optional[str]]:
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
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        SNMP version used for sending the trap.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

