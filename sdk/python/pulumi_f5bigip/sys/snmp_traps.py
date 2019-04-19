# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SnmpTraps(pulumi.CustomResource):
    auth_passwordencrypted: pulumi.Output[str]
    auth_protocol: pulumi.Output[str]
    community: pulumi.Output[str]
    """
    Specifies the community string used for this trap.
    """
    description: pulumi.Output[str]
    """
    The port that the trap will be sent to.
    """
    engine_id: pulumi.Output[str]
    host: pulumi.Output[str]
    """
    The host the trap will be sent to.
    """
    name: pulumi.Output[str]
    """
    Name of the snmp trap.
    """
    port: pulumi.Output[float]
    """
    User defined description.
    """
    privacy_password: pulumi.Output[str]
    privacy_password_encrypted: pulumi.Output[str]
    privacy_protocol: pulumi.Output[str]
    security_level: pulumi.Output[str]
    security_name: pulumi.Output[str]
    version: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, auth_passwordencrypted=None, auth_protocol=None, community=None, description=None, engine_id=None, host=None, name=None, port=None, privacy_password=None, privacy_password_encrypted=None, privacy_protocol=None, security_level=None, security_name=None, version=None, __name__=None, __opts__=None):
        """
        `bigip_sys_snmp_traps` provides details bout how to enable snmp_traps resource on BIG-IP
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] community: Specifies the community string used for this trap.
        :param pulumi.Input[str] description: The port that the trap will be sent to.
        :param pulumi.Input[str] host: The host the trap will be sent to.
        :param pulumi.Input[str] name: Name of the snmp trap.
        :param pulumi.Input[float] port: User defined description.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

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


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

