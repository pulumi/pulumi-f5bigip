# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IApp(pulumi.CustomResource):
    description: pulumi.Output[str]
    devicegroup: pulumi.Output[str]
    execute_action: pulumi.Output[str]
    inherited_devicegroup: pulumi.Output[str]
    inherited_traffic_group: pulumi.Output[str]
    jsonfile: pulumi.Output[str]
    """
    Refer to the Json file which will be deployed on F5 BIG-IP.
    """
    lists: pulumi.Output[list]
    metadatas: pulumi.Output[list]
    name: pulumi.Output[str]
    """
    Name of the iApp.
    """
    partition: pulumi.Output[str]
    strict_updates: pulumi.Output[str]
    tables: pulumi.Output[list]
    template: pulumi.Output[str]
    template_modified: pulumi.Output[str]
    template_prerequisite_errors: pulumi.Output[str]
    traffic_group: pulumi.Output[str]
    variables: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, description=None, devicegroup=None, execute_action=None, inherited_devicegroup=None, inherited_traffic_group=None, jsonfile=None, lists=None, metadatas=None, name=None, partition=None, strict_updates=None, tables=None, template=None, template_modified=None, template_prerequisite_errors=None, traffic_group=None, variables=None, __name__=None, __opts__=None):
        """
        `bigip_sys_iapp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.  
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] jsonfile: Refer to the Json file which will be deployed on F5 BIG-IP.
        :param pulumi.Input[str] name: Name of the iApp.
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
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

