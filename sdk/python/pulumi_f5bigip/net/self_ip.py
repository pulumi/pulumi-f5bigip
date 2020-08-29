# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['SelfIp']


class SelfIp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 traffic_group: Optional[pulumi.Input[str]] = None,
                 vlan: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `net.SelfIp` Manages a selfip configuration

        Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-selfip.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1/24",
            vlan="/Common/internal",
            traffic_group="traffic-group-1",
            opts=ResourceOptions(depends_on=[bigip_net_vlan["vlan1"]]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip: The Self IP's address and netmask.
        :param pulumi.Input[str] name: Name of the selfip
        :param pulumi.Input[str] traffic_group: Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        :param pulumi.Input[str] vlan: Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
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

            if ip is None:
                raise TypeError("Missing required property 'ip'")
            __props__['ip'] = ip
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['traffic_group'] = traffic_group
            if vlan is None:
                raise TypeError("Missing required property 'vlan'")
            __props__['vlan'] = vlan
        super(SelfIp, __self__).__init__(
            'f5bigip:net/selfIp:SelfIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            traffic_group: Optional[pulumi.Input[str]] = None,
            vlan: Optional[pulumi.Input[str]] = None) -> 'SelfIp':
        """
        Get an existing SelfIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip: The Self IP's address and netmask.
        :param pulumi.Input[str] name: Name of the selfip
        :param pulumi.Input[str] traffic_group: Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        :param pulumi.Input[str] vlan: Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ip"] = ip
        __props__["name"] = name
        __props__["traffic_group"] = traffic_group
        __props__["vlan"] = vlan
        return SelfIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        The Self IP's address and netmask.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the selfip
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        """
        return pulumi.get(self, "traffic_group")

    @property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[str]:
        """
        Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        return pulumi.get(self, "vlan")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

