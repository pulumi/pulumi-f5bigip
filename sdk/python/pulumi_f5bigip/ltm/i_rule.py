# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IRule(pulumi.CustomResource):
    irule: pulumi.Output[str]
    """
    Body of the iRule
    """
    name: pulumi.Output[str]
    """
    Name of the iRule
    """
    def __init__(__self__, resource_name, opts=None, irule=None, name=None, __name__=None, __opts__=None):
        """
        `bigip_ltm_irule` Creates iRule on BIG-IP F5 device
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] irule: Body of the iRule
        :param pulumi.Input[str] name: Name of the iRule
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

        if irule is None:
            raise TypeError("Missing required property 'irule'")
        __props__['irule'] = irule

        if name is None:
            raise TypeError("Missing required property 'name'")
        __props__['name'] = name

        super(IRule, __self__).__init__(
            'f5bigip:ltm/iRule:IRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

