# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Snmp']


class Snmp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowedaddresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 sys_contact: Optional[pulumi.Input[str]] = None,
                 sys_location: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `sys.Snmp` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        snmp = f5bigip.sys.Snmp("snmp",
            allowedaddresses=["202.10.10.2"],
            sys_contact=" NetOPsAdmin s.shitole@f5.com",
            sys_location="SeattleHQ")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] allowedaddresses: Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
        :param pulumi.Input[str] sys_contact: Specifies the contact information for the system administrator.
        :param pulumi.Input[str] sys_location: Describes the system's physical location.
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

            __props__['allowedaddresses'] = allowedaddresses
            __props__['sys_contact'] = sys_contact
            __props__['sys_location'] = sys_location
        super(Snmp, __self__).__init__(
            'f5bigip:sys/snmp:Snmp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowedaddresses: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            sys_contact: Optional[pulumi.Input[str]] = None,
            sys_location: Optional[pulumi.Input[str]] = None) -> 'Snmp':
        """
        Get an existing Snmp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] allowedaddresses: Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
        :param pulumi.Input[str] sys_contact: Specifies the contact information for the system administrator.
        :param pulumi.Input[str] sys_location: Describes the system's physical location.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowedaddresses"] = allowedaddresses
        __props__["sys_contact"] = sys_contact
        __props__["sys_location"] = sys_location
        return Snmp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def allowedaddresses(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
        """
        return pulumi.get(self, "allowedaddresses")

    @property
    @pulumi.getter(name="sysContact")
    def sys_contact(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the contact information for the system administrator.
        """
        return pulumi.get(self, "sys_contact")

    @property
    @pulumi.getter(name="sysLocation")
    def sys_location(self) -> pulumi.Output[Optional[str]]:
        """
        Describes the system's physical location.
        """
        return pulumi.get(self, "sys_location")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

