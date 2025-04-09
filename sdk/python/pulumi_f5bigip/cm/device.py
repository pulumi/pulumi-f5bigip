# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['DeviceArgs', 'Device']

@pulumi.input_type
class DeviceArgs:
    def __init__(__self__, *,
                 configsync_ip: pulumi.Input[builtins.str],
                 name: pulumi.Input[builtins.str],
                 mirror_ip: Optional[pulumi.Input[builtins.str]] = None,
                 mirror_secondary_ip: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Device resource.
        :param pulumi.Input[builtins.str] configsync_ip: IP address used for config sync
        :param pulumi.Input[builtins.str] name: Address of the Device which needs to be Deviceensed
        :param pulumi.Input[builtins.str] mirror_ip: IP address used for state mirroring
        :param pulumi.Input[builtins.str] mirror_secondary_ip: Secondary IP address used for state mirroring
        """
        pulumi.set(__self__, "configsync_ip", configsync_ip)
        pulumi.set(__self__, "name", name)
        if mirror_ip is not None:
            pulumi.set(__self__, "mirror_ip", mirror_ip)
        if mirror_secondary_ip is not None:
            pulumi.set(__self__, "mirror_secondary_ip", mirror_secondary_ip)

    @property
    @pulumi.getter(name="configsyncIp")
    def configsync_ip(self) -> pulumi.Input[builtins.str]:
        """
        IP address used for config sync
        """
        return pulumi.get(self, "configsync_ip")

    @configsync_ip.setter
    def configsync_ip(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "configsync_ip", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Address of the Device which needs to be Deviceensed
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="mirrorIp")
    def mirror_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        IP address used for state mirroring
        """
        return pulumi.get(self, "mirror_ip")

    @mirror_ip.setter
    def mirror_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "mirror_ip", value)

    @property
    @pulumi.getter(name="mirrorSecondaryIp")
    def mirror_secondary_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Secondary IP address used for state mirroring
        """
        return pulumi.get(self, "mirror_secondary_ip")

    @mirror_secondary_ip.setter
    def mirror_secondary_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "mirror_secondary_ip", value)


@pulumi.input_type
class _DeviceState:
    def __init__(__self__, *,
                 configsync_ip: Optional[pulumi.Input[builtins.str]] = None,
                 mirror_ip: Optional[pulumi.Input[builtins.str]] = None,
                 mirror_secondary_ip: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Device resources.
        :param pulumi.Input[builtins.str] configsync_ip: IP address used for config sync
        :param pulumi.Input[builtins.str] mirror_ip: IP address used for state mirroring
        :param pulumi.Input[builtins.str] mirror_secondary_ip: Secondary IP address used for state mirroring
        :param pulumi.Input[builtins.str] name: Address of the Device which needs to be Deviceensed
        """
        if configsync_ip is not None:
            pulumi.set(__self__, "configsync_ip", configsync_ip)
        if mirror_ip is not None:
            pulumi.set(__self__, "mirror_ip", mirror_ip)
        if mirror_secondary_ip is not None:
            pulumi.set(__self__, "mirror_secondary_ip", mirror_secondary_ip)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="configsyncIp")
    def configsync_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        IP address used for config sync
        """
        return pulumi.get(self, "configsync_ip")

    @configsync_ip.setter
    def configsync_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "configsync_ip", value)

    @property
    @pulumi.getter(name="mirrorIp")
    def mirror_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        IP address used for state mirroring
        """
        return pulumi.get(self, "mirror_ip")

    @mirror_ip.setter
    def mirror_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "mirror_ip", value)

    @property
    @pulumi.getter(name="mirrorSecondaryIp")
    def mirror_secondary_ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Secondary IP address used for state mirroring
        """
        return pulumi.get(self, "mirror_secondary_ip")

    @mirror_secondary_ip.setter
    def mirror_secondary_ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "mirror_secondary_ip", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Address of the Device which needs to be Deviceensed
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


class Device(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configsync_ip: Optional[pulumi.Input[builtins.str]] = None,
                 mirror_ip: Optional[pulumi.Input[builtins.str]] = None,
                 mirror_secondary_ip: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        `cm.Device` provides details about a specific bigip

        This resource is helpful when configuring the BIG-IP device in cluster or in HA mode.
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        my_new_device = f5bigip.cm.Device("my_new_device",
            name="bigip300.f5.com",
            configsync_ip="2.2.2.2",
            mirror_ip="10.10.10.10",
            mirror_secondary_ip="11.11.11.11")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] configsync_ip: IP address used for config sync
        :param pulumi.Input[builtins.str] mirror_ip: IP address used for state mirroring
        :param pulumi.Input[builtins.str] mirror_secondary_ip: Secondary IP address used for state mirroring
        :param pulumi.Input[builtins.str] name: Address of the Device which needs to be Deviceensed
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `cm.Device` provides details about a specific bigip

        This resource is helpful when configuring the BIG-IP device in cluster or in HA mode.
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        my_new_device = f5bigip.cm.Device("my_new_device",
            name="bigip300.f5.com",
            configsync_ip="2.2.2.2",
            mirror_ip="10.10.10.10",
            mirror_secondary_ip="11.11.11.11")
        ```

        :param str resource_name: The name of the resource.
        :param DeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 configsync_ip: Optional[pulumi.Input[builtins.str]] = None,
                 mirror_ip: Optional[pulumi.Input[builtins.str]] = None,
                 mirror_secondary_ip: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceArgs.__new__(DeviceArgs)

            if configsync_ip is None and not opts.urn:
                raise TypeError("Missing required property 'configsync_ip'")
            __props__.__dict__["configsync_ip"] = configsync_ip
            __props__.__dict__["mirror_ip"] = mirror_ip
            __props__.__dict__["mirror_secondary_ip"] = mirror_secondary_ip
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
        super(Device, __self__).__init__(
            'f5bigip:cm/device:Device',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            configsync_ip: Optional[pulumi.Input[builtins.str]] = None,
            mirror_ip: Optional[pulumi.Input[builtins.str]] = None,
            mirror_secondary_ip: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None) -> 'Device':
        """
        Get an existing Device resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] configsync_ip: IP address used for config sync
        :param pulumi.Input[builtins.str] mirror_ip: IP address used for state mirroring
        :param pulumi.Input[builtins.str] mirror_secondary_ip: Secondary IP address used for state mirroring
        :param pulumi.Input[builtins.str] name: Address of the Device which needs to be Deviceensed
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceState.__new__(_DeviceState)

        __props__.__dict__["configsync_ip"] = configsync_ip
        __props__.__dict__["mirror_ip"] = mirror_ip
        __props__.__dict__["mirror_secondary_ip"] = mirror_secondary_ip
        __props__.__dict__["name"] = name
        return Device(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configsyncIp")
    def configsync_ip(self) -> pulumi.Output[builtins.str]:
        """
        IP address used for config sync
        """
        return pulumi.get(self, "configsync_ip")

    @property
    @pulumi.getter(name="mirrorIp")
    def mirror_ip(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        IP address used for state mirroring
        """
        return pulumi.get(self, "mirror_ip")

    @property
    @pulumi.getter(name="mirrorSecondaryIp")
    def mirror_secondary_ip(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Secondary IP address used for state mirroring
        """
        return pulumi.get(self, "mirror_secondary_ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Address of the Device which needs to be Deviceensed
        """
        return pulumi.get(self, "name")

