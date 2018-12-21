# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class IApp(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, description=None, devicegroup=None, execute_action=None, inherited_devicegroup=None, inherited_traffic_group=None, jsonfile=None, lists=None, metadatas=None, name=None, partition=None, strict_updates=None, tables=None, template=None, template_modified=None, template_prerequisite_errors=None, traffic_group=None, variables=None):
        """Create a IApp resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['devicegroup'] = devicegroup

        __props__['execute_action'] = execute_action

        __props__['inherited_devicegroup'] = inherited_devicegroup

        __props__['inherited_traffic_group'] = inherited_traffic_group

        __props__['jsonfile'] = jsonfile

        __props__['lists'] = lists

        __props__['metadatas'] = metadatas

        __props__['name'] = name

        __props__['partition'] = partition

        __props__['strict_updates'] = strict_updates

        __props__['tables'] = tables

        __props__['template'] = template

        __props__['template_modified'] = template_modified

        __props__['template_prerequisite_errors'] = template_prerequisite_errors

        __props__['traffic_group'] = traffic_group

        __props__['variables'] = variables

        super(IApp, __self__).__init__(
            'f5bigip:sys/iApp:IApp',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

