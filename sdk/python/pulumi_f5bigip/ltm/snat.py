# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Snat(pulumi.CustomResource):
    autolasthop: pulumi.Output[str]
    full_path: pulumi.Output[str]
    mirror: pulumi.Output[str]
    """
    Enables or disables mirroring of SNAT connections.
    """
    name: pulumi.Output[str]
    """
    Name of the snat
    """
    origins: pulumi.Output[list]
    """
    IP or hostname of the snat
    """
    partition: pulumi.Output[str]
    """
    Displays the administrative partition within which this profile resides
    """
    snatpool: pulumi.Output[str]
    """
    Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
    """
    sourceport: pulumi.Output[str]
    """
    Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
    """
    translation: pulumi.Output[str]
    """
    Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
    """
    vlansdisabled: pulumi.Output[bool]
    """
    Disables the SNAT on all VLANs.
    """
    def __init__(__self__, __name__, __opts__=None, autolasthop=None, full_path=None, mirror=None, name=None, origins=None, partition=None, snatpool=None, sourceport=None, translation=None, vlansdisabled=None):
        """
        `bigip_ltm_snat` Manages a snat configuration
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] autolasthop
        :param pulumi.Input[str] full_path
        :param pulumi.Input[str] mirror: Enables or disables mirroring of SNAT connections.
        :param pulumi.Input[str] name: Name of the snat
        :param pulumi.Input[list] origins: IP or hostname of the snat
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] snatpool: Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
        :param pulumi.Input[str] sourceport: Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
        :param pulumi.Input[str] translation: Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
        :param pulumi.Input[bool] vlansdisabled: Disables the SNAT on all VLANs.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['autolasthop'] = autolasthop

        __props__['full_path'] = full_path

        __props__['mirror'] = mirror

        if not name:
            raise TypeError('Missing required property name')
        __props__['name'] = name

        if not origins:
            raise TypeError('Missing required property origins')
        __props__['origins'] = origins

        __props__['partition'] = partition

        __props__['snatpool'] = snatpool

        __props__['sourceport'] = sourceport

        __props__['translation'] = translation

        __props__['vlansdisabled'] = vlansdisabled

        super(Snat, __self__).__init__(
            'f5bigip:ltm/snat:Snat',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

