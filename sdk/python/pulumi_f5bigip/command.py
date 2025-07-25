# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['CommandArgs', 'Command']

@pulumi.input_type
class CommandArgs:
    def __init__(__self__, *,
                 commands: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]],
                 command_results: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 when: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a Command resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] commands: The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] command_results: The resulting output from the `commands` executed.
        """
        pulumi.set(__self__, "commands", commands)
        if command_results is not None:
            pulumi.set(__self__, "command_results", command_results)
        if when is not None:
            pulumi.set(__self__, "when", when)

    @_builtins.property
    @pulumi.getter
    def commands(self) -> pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]:
        """
        The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        return pulumi.get(self, "commands")

    @commands.setter
    def commands(self, value: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]):
        pulumi.set(self, "commands", value)

    @_builtins.property
    @pulumi.getter(name="commandResults")
    def command_results(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        The resulting output from the `commands` executed.
        """
        return pulumi.get(self, "command_results")

    @command_results.setter
    def command_results(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "command_results", value)

    @_builtins.property
    @pulumi.getter
    def when(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "when")

    @when.setter
    def when(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "when", value)


@pulumi.input_type
class _CommandState:
    def __init__(__self__, *,
                 command_results: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 commands: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 when: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering Command resources.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] command_results: The resulting output from the `commands` executed.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] commands: The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        if command_results is not None:
            pulumi.set(__self__, "command_results", command_results)
        if commands is not None:
            pulumi.set(__self__, "commands", commands)
        if when is not None:
            pulumi.set(__self__, "when", when)

    @_builtins.property
    @pulumi.getter(name="commandResults")
    def command_results(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        The resulting output from the `commands` executed.
        """
        return pulumi.get(self, "command_results")

    @command_results.setter
    def command_results(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "command_results", value)

    @_builtins.property
    @pulumi.getter
    def commands(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        return pulumi.get(self, "commands")

    @commands.setter
    def commands(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "commands", value)

    @_builtins.property
    @pulumi.getter
    def when(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "when")

    @when.setter
    def when(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "when", value)


@pulumi.type_token("f5bigip:index/command:Command")
class Command(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command_results: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 commands: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 when: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        `Command` Run TMSH commands on F5 devices

        This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] command_results: The resulting output from the `commands` executed.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] commands: The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CommandArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `Command` Run TMSH commands on F5 devices

        This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device

        :param str resource_name: The name of the resource.
        :param CommandArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CommandArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command_results: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 commands: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 when: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CommandArgs.__new__(CommandArgs)

            __props__.__dict__["command_results"] = command_results
            if commands is None and not opts.urn:
                raise TypeError("Missing required property 'commands'")
            __props__.__dict__["commands"] = commands
            __props__.__dict__["when"] = when
        super(Command, __self__).__init__(
            'f5bigip:index/command:Command',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            command_results: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            commands: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            when: Optional[pulumi.Input[_builtins.str]] = None) -> 'Command':
        """
        Get an existing Command resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] command_results: The resulting output from the `commands` executed.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] commands: The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CommandState.__new__(_CommandState)

        __props__.__dict__["command_results"] = command_results
        __props__.__dict__["commands"] = commands
        __props__.__dict__["when"] = when
        return Command(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="commandResults")
    def command_results(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        The resulting output from the `commands` executed.
        """
        return pulumi.get(self, "command_results")

    @_builtins.property
    @pulumi.getter
    def commands(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        """
        return pulumi.get(self, "commands")

    @_builtins.property
    @pulumi.getter
    def when(self) -> pulumi.Output[Optional[_builtins.str]]:
        return pulumi.get(self, "when")

