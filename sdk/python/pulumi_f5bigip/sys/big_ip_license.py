# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BigIpLicenseArgs', 'BigIpLicense']

@pulumi.input_type
class BigIpLicenseArgs:
    def __init__(__self__, *,
                 command: pulumi.Input[str],
                 registration_key: pulumi.Input[str]):
        """
        The set of arguments for constructing a BigIpLicense resource.
        :param pulumi.Input[str] command: Tmsh command to execute tmsh commands like install
        :param pulumi.Input[str] registration_key: A unique Key F5 provides for Licensing BIG-IP
        """
        pulumi.set(__self__, "command", command)
        pulumi.set(__self__, "registration_key", registration_key)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Input[str]:
        """
        Tmsh command to execute tmsh commands like install
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: pulumi.Input[str]):
        pulumi.set(self, "command", value)

    @property
    @pulumi.getter(name="registrationKey")
    def registration_key(self) -> pulumi.Input[str]:
        """
        A unique Key F5 provides for Licensing BIG-IP
        """
        return pulumi.get(self, "registration_key")

    @registration_key.setter
    def registration_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "registration_key", value)


@pulumi.input_type
class _BigIpLicenseState:
    def __init__(__self__, *,
                 command: Optional[pulumi.Input[str]] = None,
                 registration_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BigIpLicense resources.
        :param pulumi.Input[str] command: Tmsh command to execute tmsh commands like install
        :param pulumi.Input[str] registration_key: A unique Key F5 provides for Licensing BIG-IP
        """
        if command is not None:
            pulumi.set(__self__, "command", command)
        if registration_key is not None:
            pulumi.set(__self__, "registration_key", registration_key)

    @property
    @pulumi.getter
    def command(self) -> Optional[pulumi.Input[str]]:
        """
        Tmsh command to execute tmsh commands like install
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command", value)

    @property
    @pulumi.getter(name="registrationKey")
    def registration_key(self) -> Optional[pulumi.Input[str]]:
        """
        A unique Key F5 provides for Licensing BIG-IP
        """
        return pulumi.get(self, "registration_key")

    @registration_key.setter
    def registration_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registration_key", value)


class BigIpLicense(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command: Optional[pulumi.Input[str]] = None,
                 registration_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BigIpLicense resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command: Tmsh command to execute tmsh commands like install
        :param pulumi.Input[str] registration_key: A unique Key F5 provides for Licensing BIG-IP
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BigIpLicenseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BigIpLicense resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BigIpLicenseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BigIpLicenseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 command: Optional[pulumi.Input[str]] = None,
                 registration_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BigIpLicenseArgs.__new__(BigIpLicenseArgs)

            if command is None and not opts.urn:
                raise TypeError("Missing required property 'command'")
            __props__.__dict__["command"] = command
            if registration_key is None and not opts.urn:
                raise TypeError("Missing required property 'registration_key'")
            __props__.__dict__["registration_key"] = registration_key
        super(BigIpLicense, __self__).__init__(
            'f5bigip:sys/bigIpLicense:BigIpLicense',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            command: Optional[pulumi.Input[str]] = None,
            registration_key: Optional[pulumi.Input[str]] = None) -> 'BigIpLicense':
        """
        Get an existing BigIpLicense resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] command: Tmsh command to execute tmsh commands like install
        :param pulumi.Input[str] registration_key: A unique Key F5 provides for Licensing BIG-IP
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BigIpLicenseState.__new__(_BigIpLicenseState)

        __props__.__dict__["command"] = command
        __props__.__dict__["registration_key"] = registration_key
        return BigIpLicense(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def command(self) -> pulumi.Output[str]:
        """
        Tmsh command to execute tmsh commands like install
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter(name="registrationKey")
    def registration_key(self) -> pulumi.Output[str]:
        """
        A unique Key F5 provides for Licensing BIG-IP
        """
        return pulumi.get(self, "registration_key")

