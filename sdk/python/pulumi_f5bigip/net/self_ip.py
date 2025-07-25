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
from .. import _utilities

__all__ = ['SelfIpArgs', 'SelfIp']

@pulumi.input_type
class SelfIpArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[_builtins.str],
                 name: pulumi.Input[_builtins.str],
                 vlan: pulumi.Input[_builtins.str],
                 port_lockdowns: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 traffic_group: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a SelfIp resource.
        :param pulumi.Input[_builtins.str] ip: The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
        :param pulumi.Input[_builtins.str] name: Name of the selfip
        :param pulumi.Input[_builtins.str] vlan: Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] port_lockdowns: Specifies the port lockdown, defaults to `Allow None` if not specified.
        :param pulumi.Input[_builtins.str] traffic_group: Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "vlan", vlan)
        if port_lockdowns is not None:
            pulumi.set(__self__, "port_lockdowns", port_lockdowns)
        if traffic_group is not None:
            pulumi.set(__self__, "traffic_group", traffic_group)

    @_builtins.property
    @pulumi.getter
    def ip(self) -> pulumi.Input[_builtins.str]:
        """
        The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "ip", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the selfip
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def vlan(self) -> pulumi.Input[_builtins.str]:
        """
        Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "vlan", value)

    @_builtins.property
    @pulumi.getter(name="portLockdowns")
    def port_lockdowns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the port lockdown, defaults to `Allow None` if not specified.
        """
        return pulumi.get(self, "port_lockdowns")

    @port_lockdowns.setter
    def port_lockdowns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "port_lockdowns", value)

    @_builtins.property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        """
        return pulumi.get(self, "traffic_group")

    @traffic_group.setter
    def traffic_group(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "traffic_group", value)


@pulumi.input_type
class _SelfIpState:
    def __init__(__self__, *,
                 ip: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 port_lockdowns: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 traffic_group: Optional[pulumi.Input[_builtins.str]] = None,
                 vlan: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering SelfIp resources.
        :param pulumi.Input[_builtins.str] ip: The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
        :param pulumi.Input[_builtins.str] name: Name of the selfip
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] port_lockdowns: Specifies the port lockdown, defaults to `Allow None` if not specified.
        :param pulumi.Input[_builtins.str] traffic_group: Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        :param pulumi.Input[_builtins.str] vlan: Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if port_lockdowns is not None:
            pulumi.set(__self__, "port_lockdowns", port_lockdowns)
        if traffic_group is not None:
            pulumi.set(__self__, "traffic_group", traffic_group)
        if vlan is not None:
            pulumi.set(__self__, "vlan", vlan)

    @_builtins.property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "ip", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the selfip
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="portLockdowns")
    def port_lockdowns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the port lockdown, defaults to `Allow None` if not specified.
        """
        return pulumi.get(self, "port_lockdowns")

    @port_lockdowns.setter
    def port_lockdowns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "port_lockdowns", value)

    @_builtins.property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        """
        return pulumi.get(self, "traffic_group")

    @traffic_group.setter
    def traffic_group(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "traffic_group", value)

    @_builtins.property
    @pulumi.getter
    def vlan(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        return pulumi.get(self, "vlan")

    @vlan.setter
    def vlan(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "vlan", value)


@pulumi.type_token("f5bigip:net/selfIp:SelfIp")
class SelfIp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 port_lockdowns: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 traffic_group: Optional[pulumi.Input[_builtins.str]] = None,
                 vlan: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        `net.SelfIp` Manages a selfip configuration

        Resource should be named with their `full path`. The full path is the combination of the `partition + name of the resource`, for example `/Common/my-selfip`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        vlan1 = f5bigip.net.Vlan("vlan1",
            name="/Common/Internal",
            tag=101,
            interfaces=[{
                "vlanport": "1.2",
                "tagged": False,
            }])
        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1/24",
            vlan="/Common/internal",
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```
        ### Example usage with `port_lockdown`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1/24",
            vlan="/Common/internal",
            traffic_group="traffic-group-1",
            port_lockdowns=[
                "tcp:4040",
                "udp:5050",
                "egp:0",
            ],
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```

        ### Example usage with `port_lockdown` set to `["none"]`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1/24",
            vlan="/Common/internal",
            traffic_group="traffic-group-1",
            port_lockdowns=["none"],
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```

        ### Example usage with route domain embedded in the `ip`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1%4/24",
            vlan="/Common/internal",
            traffic_group="traffic-group-1",
            port_lockdowns=["none"],
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] ip: The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
        :param pulumi.Input[_builtins.str] name: Name of the selfip
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] port_lockdowns: Specifies the port lockdown, defaults to `Allow None` if not specified.
        :param pulumi.Input[_builtins.str] traffic_group: Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        :param pulumi.Input[_builtins.str] vlan: Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SelfIpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `net.SelfIp` Manages a selfip configuration

        Resource should be named with their `full path`. The full path is the combination of the `partition + name of the resource`, for example `/Common/my-selfip`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        vlan1 = f5bigip.net.Vlan("vlan1",
            name="/Common/Internal",
            tag=101,
            interfaces=[{
                "vlanport": "1.2",
                "tagged": False,
            }])
        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1/24",
            vlan="/Common/internal",
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```
        ### Example usage with `port_lockdown`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1/24",
            vlan="/Common/internal",
            traffic_group="traffic-group-1",
            port_lockdowns=[
                "tcp:4040",
                "udp:5050",
                "egp:0",
            ],
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```

        ### Example usage with `port_lockdown` set to `["none"]`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1/24",
            vlan="/Common/internal",
            traffic_group="traffic-group-1",
            port_lockdowns=["none"],
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```

        ### Example usage with route domain embedded in the `ip`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        selfip1 = f5bigip.net.SelfIp("selfip1",
            name="/Common/internalselfIP",
            ip="11.1.1.1%4/24",
            vlan="/Common/internal",
            traffic_group="traffic-group-1",
            port_lockdowns=["none"],
            opts = pulumi.ResourceOptions(depends_on=[vlan1]))
        ```

        :param str resource_name: The name of the resource.
        :param SelfIpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SelfIpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 port_lockdowns: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 traffic_group: Optional[pulumi.Input[_builtins.str]] = None,
                 vlan: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SelfIpArgs.__new__(SelfIpArgs)

            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["port_lockdowns"] = port_lockdowns
            __props__.__dict__["traffic_group"] = traffic_group
            if vlan is None and not opts.urn:
                raise TypeError("Missing required property 'vlan'")
            __props__.__dict__["vlan"] = vlan
        super(SelfIp, __self__).__init__(
            'f5bigip:net/selfIp:SelfIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            port_lockdowns: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            traffic_group: Optional[pulumi.Input[_builtins.str]] = None,
            vlan: Optional[pulumi.Input[_builtins.str]] = None) -> 'SelfIp':
        """
        Get an existing SelfIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] ip: The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
        :param pulumi.Input[_builtins.str] name: Name of the selfip
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] port_lockdowns: Specifies the port lockdown, defaults to `Allow None` if not specified.
        :param pulumi.Input[_builtins.str] traffic_group: Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        :param pulumi.Input[_builtins.str] vlan: Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SelfIpState.__new__(_SelfIpState)

        __props__.__dict__["ip"] = ip
        __props__.__dict__["name"] = name
        __props__.__dict__["port_lockdowns"] = port_lockdowns
        __props__.__dict__["traffic_group"] = traffic_group
        __props__.__dict__["vlan"] = vlan
        return SelfIp(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def ip(self) -> pulumi.Output[_builtins.str]:
        """
        The Self IP's address and netmask. The IP address could also contain the route domain, e.g. `10.12.13.14%4/24`.
        """
        return pulumi.get(self, "ip")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the selfip
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="portLockdowns")
    def port_lockdowns(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Specifies the port lockdown, defaults to `Allow None` if not specified.
        """
        return pulumi.get(self, "port_lockdowns")

    @_builtins.property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the traffic group, defaults to `traffic-group-local-only` if not specified.
        """
        return pulumi.get(self, "traffic_group")

    @_builtins.property
    @pulumi.getter
    def vlan(self) -> pulumi.Output[_builtins.str]:
        """
        Specifies the VLAN for which you are setting a self IP address. This setting must be provided when a self IP is created.
        """
        return pulumi.get(self, "vlan")

