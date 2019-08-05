# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Snmp(pulumi.CustomResource):
    allowedaddresses: pulumi.Output[list]
    """
    Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
    """
    sys_contact: pulumi.Output[str]
    """
    Specifies the contact information for the system administrator.
    """
    sys_location: pulumi.Output[str]
    """
    Describes the system's physical location.
    """
    def __init__(__self__, resource_name, opts=None, allowedaddresses=None, sys_contact=None, sys_location=None, __name__=None, __opts__=None):
        """
        `bigip_sys_snmp` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowedaddresses: Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
        :param pulumi.Input[str] sys_contact: Specifies the contact information for the system administrator.
        :param pulumi.Input[str] sys_location: Describes the system's physical location.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/sys_snmp.html.markdown.
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

        __props__['allowedaddresses'] = allowedaddresses

        __props__['sys_contact'] = sys_contact

        __props__['sys_location'] = sys_location

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Snmp, __self__).__init__(
            'f5bigip:sys/snmp:Snmp',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

