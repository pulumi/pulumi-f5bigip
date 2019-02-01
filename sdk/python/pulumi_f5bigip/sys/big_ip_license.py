# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BigIpLicense(pulumi.CustomResource):
    command: pulumi.Output[str]
    registration_key: pulumi.Output[str]
    def __init__(__self__, __name__, __opts__=None, command=None, registration_key=None):
        """
        Create a BigIpLicense resource with the given unique name, props, and options.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] command
        :param pulumi.Input[str] registration_key
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not command:
            raise TypeError('Missing required property command')
        __props__['command'] = command

        if not registration_key:
            raise TypeError('Missing required property registration_key')
        __props__['registration_key'] = registration_key

        super(BigIpLicense, __self__).__init__(
            'f5bigip:sys/bigIpLicense:BigIpLicense',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

