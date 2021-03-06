# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Ntp']


class Ntp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `sys.Ntp` provides details about a specific bigip

        This resource is helpful when configuring NTP server on the BIG-IP.
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
        :param pulumi.Input[str] description: Name of the ntp Servers
        :param pulumi.Input[Sequence[pulumi.Input[str]]] servers: Adds NTP servers to or deletes NTP servers from the BIG-IP system.
        :param pulumi.Input[str] timezone: Specifies the time zone that you want to use for the system time.
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

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            __props__['servers'] = servers
            __props__['timezone'] = timezone
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
        :param pulumi.Input[str] description: Name of the ntp Servers
        :param pulumi.Input[Sequence[pulumi.Input[str]]] servers: Adds NTP servers to or deletes NTP servers from the BIG-IP system.
        :param pulumi.Input[str] timezone: Specifies the time zone that you want to use for the system time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["servers"] = servers
        __props__["timezone"] = timezone
        return Ntp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Name of the ntp Servers
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Adds NTP servers to or deletes NTP servers from the BIG-IP system.
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the time zone that you want to use for the system time.
        """
        return pulumi.get(self, "timezone")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

