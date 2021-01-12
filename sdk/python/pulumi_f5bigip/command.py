# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Command']


class Command(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command_results: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 commands: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 when: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `Command` Run TMSH commands on F5 devices

        This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        #create ltm node
        test_command = f5bigip.Command("test-command",
            commands=["delete ltm node 10.10.10.70"],
            when="destroy")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] command_results: The resulting output from the `commands` executed
        :param pulumi.Input[Sequence[pulumi.Input[str]]] commands: The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
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

            __props__['command_results'] = command_results
            if commands is None and not opts.urn:
                raise TypeError("Missing required property 'commands'")
            __props__['commands'] = commands
            __props__['when'] = when
        super(Command, __self__).__init__(
            'f5bigip:index/command:Command',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            command_results: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            commands: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            when: Optional[pulumi.Input[str]] = None) -> 'Command':
        """
        Get an existing Command resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] command_results: The resulting output from the `commands` executed
        :param pulumi.Input[Sequence[pulumi.Input[str]]] commands: The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["command_results"] = command_results
        __props__["commands"] = commands
        __props__["when"] = when
        return Command(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="commandResults")
    def command_results(self) -> pulumi.Output[Sequence[str]]:
        """
        The resulting output from the `commands` executed
        """
        return pulumi.get(self, "command_results")

    @property
    @pulumi.getter
    def commands(self) -> pulumi.Output[Sequence[str]]:
        """
        The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        return pulumi.get(self, "commands")

    @property
    @pulumi.getter
    def when(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "when")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

