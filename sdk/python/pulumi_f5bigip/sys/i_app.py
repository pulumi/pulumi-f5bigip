# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['IApp']


class IApp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 devicegroup: Optional[pulumi.Input[str]] = None,
                 execute_action: Optional[pulumi.Input[str]] = None,
                 inherited_devicegroup: Optional[pulumi.Input[str]] = None,
                 inherited_traffic_group: Optional[pulumi.Input[str]] = None,
                 jsonfile: Optional[pulumi.Input[str]] = None,
                 lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppListArgs']]]]] = None,
                 metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppMetadataArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 strict_updates: Optional[pulumi.Input[str]] = None,
                 tables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppTableArgs']]]]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 template_modified: Optional[pulumi.Input[str]] = None,
                 template_prerequisite_errors: Optional[pulumi.Input[str]] = None,
                 traffic_group: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppVariableArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        simplehttp = f5bigip.sys.IApp("simplehttp",
            name="simplehttp",
            jsonfile=(lambda path: open(path).read())("simplehttp.json"))
        ```
        ### Json File
        ```python
        import pulumi
        ```

         * `description` - User defined description.
         * `deviceGroup` - The name of the device group that the application service is assigned to.
         * `executeAction` - Run the specified template action associated with the application.
         * `inheritedDevicegroup`- Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
         * `inheritedTrafficGroup` - Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
         * `partition` - Displays the administrative partition within which the application resides.
         * `strictUpdates` - Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
         * `template` - The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
         * `templateModified` - Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
         * `templatePrerequisiteErrors` - Indicates any missing prerequisites associated with the template that defines this application.
         * `trafficGroup` - The name of the traffic group that the application service is assigned to.
         * `lists` - string values
         * `metadata` - User defined generic data for the application service. It is a name and value pair.
         * `tables` - Values provided like pool name, nodes etc.
         * `variables` - Name, values, encrypted or not

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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            devicegroup: Optional[pulumi.Input[str]] = None,
            execute_action: Optional[pulumi.Input[str]] = None,
            inherited_devicegroup: Optional[pulumi.Input[str]] = None,
            inherited_traffic_group: Optional[pulumi.Input[str]] = None,
            jsonfile: Optional[pulumi.Input[str]] = None,
            lists: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppListArgs']]]]] = None,
            metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppMetadataArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition: Optional[pulumi.Input[str]] = None,
            strict_updates: Optional[pulumi.Input[str]] = None,
            tables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppTableArgs']]]]] = None,
            template: Optional[pulumi.Input[str]] = None,
            template_modified: Optional[pulumi.Input[str]] = None,
            template_prerequisite_errors: Optional[pulumi.Input[str]] = None,
            traffic_group: Optional[pulumi.Input[str]] = None,
            variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IAppVariableArgs']]]]] = None) -> 'IApp':
        """
        Get an existing IApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Address of the Iapp which needs to be Iappensed
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def devicegroup(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "devicegroup")

    @property
    @pulumi.getter(name="executeAction")
    def execute_action(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "execute_action")

    @property
    @pulumi.getter(name="inheritedDevicegroup")
    def inherited_devicegroup(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "inherited_devicegroup")

    @property
    @pulumi.getter(name="inheritedTrafficGroup")
    def inherited_traffic_group(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "inherited_traffic_group")

    @property
    @pulumi.getter
    def jsonfile(self) -> pulumi.Output[Optional[str]]:
        """
        Refer to the Json file which will be deployed on F5 BIG-IP.
        """
        return pulumi.get(self, "jsonfile")

    @property
    @pulumi.getter
    def lists(self) -> pulumi.Output[Optional[Sequence['outputs.IAppList']]]:
        return pulumi.get(self, "lists")

    @property
    @pulumi.getter
    def metadatas(self) -> pulumi.Output[Optional[Sequence['outputs.IAppMetadata']]]:
        return pulumi.get(self, "metadatas")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the iApp.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> pulumi.Output[Optional[str]]:
        """
        Address of the Iapp which needs to be Iappensed
        """
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter(name="strictUpdates")
    def strict_updates(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "strict_updates")

    @property
    @pulumi.getter
    def tables(self) -> pulumi.Output[Optional[Sequence['outputs.IAppTable']]]:
        return pulumi.get(self, "tables")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter(name="templateModified")
    def template_modified(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "template_modified")

    @property
    @pulumi.getter(name="templatePrerequisiteErrors")
    def template_prerequisite_errors(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "template_prerequisite_errors")

    @property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> pulumi.Output[Optional[str]]:
        """
        BIG-IP password
        """
        return pulumi.get(self, "traffic_group")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[Sequence['outputs.IAppVariable']]]:
        return pulumi.get(self, "variables")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

