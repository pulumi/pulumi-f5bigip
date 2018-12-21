# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProfileOneConnect(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, defaults_from=None, idle_timeout_override=None, max_age=None, max_reuse=None, max_size=None, name=None, partition=None, share_pools=None, source_mask=None):
        """Create a ProfileOneConnect resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['defaults_from'] = defaults_from

        __props__['idle_timeout_override'] = idle_timeout_override

        __props__['max_age'] = max_age

        __props__['max_reuse'] = max_reuse

        __props__['max_size'] = max_size

        if not name:
            raise TypeError('Missing required property name')
        __props__['name'] = name

        __props__['partition'] = partition

        __props__['share_pools'] = share_pools

        __props__['source_mask'] = source_mask

        super(ProfileOneConnect, __self__).__init__(
            'f5bigip:ltm/profileOneConnect:ProfileOneConnect',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

