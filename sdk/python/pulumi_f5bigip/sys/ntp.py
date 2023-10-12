# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NtpArgs', 'Ntp']

@pulumi.input_type
class NtpArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 servers: pulumi.Input[Sequence[pulumi.Input[str]]],
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ntp resource.
        :param pulumi.Input[str] description: User defined description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] servers: Specifies the time servers that the system uses to update the system time.
        :param pulumi.Input[str] timezone: Specifies the time zone that you want to use for the system time.
        """
        NtpArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            servers=servers,
            timezone=timezone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: pulumi.Input[str],
             servers: pulumi.Input[Sequence[pulumi.Input[str]]],
             timezone: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("description", description)
        _setter("servers", servers)
        if timezone is not None:
            _setter("timezone", timezone)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        User defined description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies the time servers that the system uses to update the system time.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "servers", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the time zone that you want to use for the system time.
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


@pulumi.input_type
class _NtpState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ntp resources.
        :param pulumi.Input[str] description: User defined description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] servers: Specifies the time servers that the system uses to update the system time.
        :param pulumi.Input[str] timezone: Specifies the time zone that you want to use for the system time.
        """
        _NtpState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            servers=servers,
            timezone=timezone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             timezone: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if description is not None:
            _setter("description", description)
        if servers is not None:
            _setter("servers", servers)
        if timezone is not None:
            _setter("timezone", timezone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User defined description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the time servers that the system uses to update the system time.
        """
        return pulumi.get(self, "servers")

    @servers.setter
    def servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "servers", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the time zone that you want to use for the system time.
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


class Ntp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `sys.Ntp` resource is helpful when configuring NTP server on the BIG-IP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        ntp1 = f5bigip.sys.Ntp("ntp1",
            description="/Common/NTP1",
            servers=["time.facebook.com"],
            timezone="America/Los_Angeles")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User defined description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] servers: Specifies the time servers that the system uses to update the system time.
        :param pulumi.Input[str] timezone: Specifies the time zone that you want to use for the system time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NtpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `sys.Ntp` resource is helpful when configuring NTP server on the BIG-IP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        ntp1 = f5bigip.sys.Ntp("ntp1",
            description="/Common/NTP1",
            servers=["time.facebook.com"],
            timezone="America/Los_Angeles")
        ```

        :param str resource_name: The name of the resource.
        :param NtpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NtpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            NtpArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NtpArgs.__new__(NtpArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if servers is None and not opts.urn:
                raise TypeError("Missing required property 'servers'")
            __props__.__dict__["servers"] = servers
            __props__.__dict__["timezone"] = timezone
        super(Ntp, __self__).__init__(
            'f5bigip:sys/ntp:Ntp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            timezone: Optional[pulumi.Input[str]] = None) -> 'Ntp':
        """
        Get an existing Ntp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User defined description.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] servers: Specifies the time servers that the system uses to update the system time.
        :param pulumi.Input[str] timezone: Specifies the time zone that you want to use for the system time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NtpState.__new__(_NtpState)

        __props__.__dict__["description"] = description
        __props__.__dict__["servers"] = servers
        __props__.__dict__["timezone"] = timezone
        return Ntp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        User defined description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the time servers that the system uses to update the system time.
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the time zone that you want to use for the system time.
        """
        return pulumi.get(self, "timezone")

