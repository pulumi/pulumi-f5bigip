# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IApp(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Address of the Iapp which needs to be Iappensed
    """
    devicegroup: pulumi.Output[str]
    """
    BIG-IP password
    """
    execute_action: pulumi.Output[str]
    """
    BIG-IP password
    """
    inherited_devicegroup: pulumi.Output[str]
    """
    BIG-IP password
    """
    inherited_traffic_group: pulumi.Output[str]
    """
    BIG-IP password
    """
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
    """
    Address of the Iapp which needs to be Iappensed
    """
    strict_updates: pulumi.Output[str]
    """
    BIG-IP password
    """
    tables: pulumi.Output[list]
    template: pulumi.Output[str]
    """
    BIG-IP password
    """
    template_modified: pulumi.Output[str]
    """
    BIG-IP password
    """
    template_prerequisite_errors: pulumi.Output[str]
    """
    BIG-IP password
    """
    traffic_group: pulumi.Output[str]
    """
    BIG-IP password
    """
    variables: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, description=None, devicegroup=None, execute_action=None, inherited_devicegroup=None, inherited_traffic_group=None, jsonfile=None, lists=None, metadatas=None, name=None, partition=None, strict_updates=None, tables=None, template=None, template_modified=None, template_prerequisite_errors=None, traffic_group=None, variables=None, __props__=None, __name__=None, __opts__=None):
        """
        `sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.  



        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_sys_iapp.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Address of the Iapp which needs to be Iappensed
        :param pulumi.Input[str] devicegroup: BIG-IP password
        :param pulumi.Input[str] execute_action: BIG-IP password
        :param pulumi.Input[str] inherited_devicegroup: BIG-IP password
        :param pulumi.Input[str] inherited_traffic_group: BIG-IP password
        :param pulumi.Input[str] jsonfile: Refer to the Json file which will be deployed on F5 BIG-IP.
        :param pulumi.Input[str] name: Name of the iApp.
        :param pulumi.Input[str] partition: Address of the Iapp which needs to be Iappensed
        :param pulumi.Input[str] strict_updates: BIG-IP password
        :param pulumi.Input[str] template: BIG-IP password
        :param pulumi.Input[str] template_modified: BIG-IP password
        :param pulumi.Input[str] template_prerequisite_errors: BIG-IP password
        :param pulumi.Input[str] traffic_group: BIG-IP password

        The **lists** object supports the following:

          * `encrypted` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **metadatas** object supports the following:

          * `persists` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **tables** object supports the following:

          * `columnNames` (`pulumi.Input[list]`)
          * `encryptedColumns` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - Name of the iApp.
          * `rows` (`pulumi.Input[list]`)
            * `rows` (`pulumi.Input[list]`)

        The **variables** object supports the following:

          * `encrypted` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - Name of the iApp.
          * `value` (`pulumi.Input[str]`)
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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

    @staticmethod
    def get(resource_name, id, opts=None, description=None, devicegroup=None, execute_action=None, inherited_devicegroup=None, inherited_traffic_group=None, jsonfile=None, lists=None, metadatas=None, name=None, partition=None, strict_updates=None, tables=None, template=None, template_modified=None, template_prerequisite_errors=None, traffic_group=None, variables=None):
        """
        Get an existing IApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Address of the Iapp which needs to be Iappensed
        :param pulumi.Input[str] devicegroup: BIG-IP password
        :param pulumi.Input[str] execute_action: BIG-IP password
        :param pulumi.Input[str] inherited_devicegroup: BIG-IP password
        :param pulumi.Input[str] inherited_traffic_group: BIG-IP password
        :param pulumi.Input[str] jsonfile: Refer to the Json file which will be deployed on F5 BIG-IP.
        :param pulumi.Input[str] name: Name of the iApp.
        :param pulumi.Input[str] partition: Address of the Iapp which needs to be Iappensed
        :param pulumi.Input[str] strict_updates: BIG-IP password
        :param pulumi.Input[str] template: BIG-IP password
        :param pulumi.Input[str] template_modified: BIG-IP password
        :param pulumi.Input[str] template_prerequisite_errors: BIG-IP password
        :param pulumi.Input[str] traffic_group: BIG-IP password

        The **lists** object supports the following:

          * `encrypted` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **metadatas** object supports the following:

          * `persists` (`pulumi.Input[str]`)
          * `value` (`pulumi.Input[str]`)

        The **tables** object supports the following:

          * `columnNames` (`pulumi.Input[list]`)
          * `encryptedColumns` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - Name of the iApp.
          * `rows` (`pulumi.Input[list]`)
            * `rows` (`pulumi.Input[list]`)

        The **variables** object supports the following:

          * `encrypted` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - Name of the iApp.
          * `value` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["devicegroup"] = devicegroup
        __props__["execute_action"] = execute_action
        __props__["inherited_devicegroup"] = inherited_devicegroup
        __props__["inherited_traffic_group"] = inherited_traffic_group
        __props__["jsonfile"] = jsonfile
        __props__["lists"] = lists
        __props__["metadatas"] = metadatas
        __props__["name"] = name
        __props__["partition"] = partition
        __props__["strict_updates"] = strict_updates
        __props__["tables"] = tables
        __props__["template"] = template
        __props__["template_modified"] = template_modified
        __props__["template_prerequisite_errors"] = template_prerequisite_errors
        __props__["traffic_group"] = traffic_group
        __props__["variables"] = variables
        return IApp(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

