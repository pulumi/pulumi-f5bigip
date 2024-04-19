# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NetTunnelArgs', 'NetTunnel']

@pulumi.input_type
class NetTunnelArgs:
    def __init__(__self__, *,
                 local_address: pulumi.Input[str],
                 name: pulumi.Input[str],
                 profile: pulumi.Input[str],
                 app_service: Optional[pulumi.Input[str]] = None,
                 auto_last_hop: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 key: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 remote_address: Optional[pulumi.Input[str]] = None,
                 secondary_address: Optional[pulumi.Input[str]] = None,
                 tos: Optional[pulumi.Input[str]] = None,
                 traffic_group: Optional[pulumi.Input[str]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 use_pmtu: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetTunnel resource.
        :param pulumi.Input[str] local_address: Specifies a local IP address. This option is required
        :param pulumi.Input[str] name: Name of the tunnel
        :param pulumi.Input[str] profile: Specifies the profile that you want to associate with the tunnel
        :param pulumi.Input[str] app_service: The application service that the object belongs to
        :param pulumi.Input[str] auto_last_hop: Specifies whether auto lasthop is enabled or not
        :param pulumi.Input[str] description: User defined description
        :param pulumi.Input[int] idle_timeout: Specifies an idle timeout for wildcard tunnels in seconds
        :param pulumi.Input[int] key: The key field may represent different values depending on the type of the tunnel
        :param pulumi.Input[str] mode: Specifies how the tunnel carries traffic
        :param pulumi.Input[int] mtu: Specifies the maximum transmission unit (MTU) of the tunnel
        :param pulumi.Input[str] partition: Displays the admin-partition within which this component resides
        :param pulumi.Input[str] remote_address: Specifies a remote IP address
        :param pulumi.Input[str] secondary_address: Specifies a secondary non-floating IP address when the local-address is set to a floating address
        :param pulumi.Input[str] tos: Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        :param pulumi.Input[str] traffic_group: Specifies a traffic-group for use with the tunnel
        :param pulumi.Input[str] transparent: Enables or disables the tunnel to be transparent
        :param pulumi.Input[str] use_pmtu: Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        """
        pulumi.set(__self__, "local_address", local_address)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "profile", profile)
        if app_service is not None:
            pulumi.set(__self__, "app_service", app_service)
        if auto_last_hop is not None:
            pulumi.set(__self__, "auto_last_hop", auto_last_hop)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if idle_timeout is not None:
            pulumi.set(__self__, "idle_timeout", idle_timeout)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if remote_address is not None:
            pulumi.set(__self__, "remote_address", remote_address)
        if secondary_address is not None:
            pulumi.set(__self__, "secondary_address", secondary_address)
        if tos is not None:
            pulumi.set(__self__, "tos", tos)
        if traffic_group is not None:
            pulumi.set(__self__, "traffic_group", traffic_group)
        if transparent is not None:
            pulumi.set(__self__, "transparent", transparent)
        if use_pmtu is not None:
            pulumi.set(__self__, "use_pmtu", use_pmtu)

    @property
    @pulumi.getter(name="localAddress")
    def local_address(self) -> pulumi.Input[str]:
        """
        Specifies a local IP address. This option is required
        """
        return pulumi.get(self, "local_address")

    @local_address.setter
    def local_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_address", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the tunnel
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Input[str]:
        """
        Specifies the profile that you want to associate with the tunnel
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: pulumi.Input[str]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[pulumi.Input[str]]:
        """
        The application service that the object belongs to
        """
        return pulumi.get(self, "app_service")

    @app_service.setter
    def app_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_service", value)

    @property
    @pulumi.getter(name="autoLastHop")
    def auto_last_hop(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether auto lasthop is enabled or not
        """
        return pulumi.get(self, "auto_last_hop")

    @auto_last_hop.setter
    def auto_last_hop(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_last_hop", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User defined description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies an idle timeout for wildcard tunnels in seconds
        """
        return pulumi.get(self, "idle_timeout")

    @idle_timeout.setter
    def idle_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "idle_timeout", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[int]]:
        """
        The key field may represent different values depending on the type of the tunnel
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how the tunnel carries traffic
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum transmission unit (MTU) of the tunnel
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Displays the admin-partition within which this component resides
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter(name="remoteAddress")
    def remote_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a remote IP address
        """
        return pulumi.get(self, "remote_address")

    @remote_address.setter
    def remote_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_address", value)

    @property
    @pulumi.getter(name="secondaryAddress")
    def secondary_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a secondary non-floating IP address when the local-address is set to a floating address
        """
        return pulumi.get(self, "secondary_address")

    @secondary_address.setter
    def secondary_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_address", value)

    @property
    @pulumi.getter
    def tos(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        """
        return pulumi.get(self, "tos")

    @tos.setter
    def tos(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tos", value)

    @property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a traffic-group for use with the tunnel
        """
        return pulumi.get(self, "traffic_group")

    @traffic_group.setter
    def traffic_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_group", value)

    @property
    @pulumi.getter
    def transparent(self) -> Optional[pulumi.Input[str]]:
        """
        Enables or disables the tunnel to be transparent
        """
        return pulumi.get(self, "transparent")

    @transparent.setter
    def transparent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transparent", value)

    @property
    @pulumi.getter(name="usePmtu")
    def use_pmtu(self) -> Optional[pulumi.Input[str]]:
        """
        Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        """
        return pulumi.get(self, "use_pmtu")

    @use_pmtu.setter
    def use_pmtu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_pmtu", value)


@pulumi.input_type
class _NetTunnelState:
    def __init__(__self__, *,
                 app_service: Optional[pulumi.Input[str]] = None,
                 auto_last_hop: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 key: Optional[pulumi.Input[int]] = None,
                 local_address: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 remote_address: Optional[pulumi.Input[str]] = None,
                 secondary_address: Optional[pulumi.Input[str]] = None,
                 tos: Optional[pulumi.Input[str]] = None,
                 traffic_group: Optional[pulumi.Input[str]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 use_pmtu: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetTunnel resources.
        :param pulumi.Input[str] app_service: The application service that the object belongs to
        :param pulumi.Input[str] auto_last_hop: Specifies whether auto lasthop is enabled or not
        :param pulumi.Input[str] description: User defined description
        :param pulumi.Input[int] idle_timeout: Specifies an idle timeout for wildcard tunnels in seconds
        :param pulumi.Input[int] key: The key field may represent different values depending on the type of the tunnel
        :param pulumi.Input[str] local_address: Specifies a local IP address. This option is required
        :param pulumi.Input[str] mode: Specifies how the tunnel carries traffic
        :param pulumi.Input[int] mtu: Specifies the maximum transmission unit (MTU) of the tunnel
        :param pulumi.Input[str] name: Name of the tunnel
        :param pulumi.Input[str] partition: Displays the admin-partition within which this component resides
        :param pulumi.Input[str] profile: Specifies the profile that you want to associate with the tunnel
        :param pulumi.Input[str] remote_address: Specifies a remote IP address
        :param pulumi.Input[str] secondary_address: Specifies a secondary non-floating IP address when the local-address is set to a floating address
        :param pulumi.Input[str] tos: Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        :param pulumi.Input[str] traffic_group: Specifies a traffic-group for use with the tunnel
        :param pulumi.Input[str] transparent: Enables or disables the tunnel to be transparent
        :param pulumi.Input[str] use_pmtu: Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        """
        if app_service is not None:
            pulumi.set(__self__, "app_service", app_service)
        if auto_last_hop is not None:
            pulumi.set(__self__, "auto_last_hop", auto_last_hop)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if idle_timeout is not None:
            pulumi.set(__self__, "idle_timeout", idle_timeout)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if local_address is not None:
            pulumi.set(__self__, "local_address", local_address)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if remote_address is not None:
            pulumi.set(__self__, "remote_address", remote_address)
        if secondary_address is not None:
            pulumi.set(__self__, "secondary_address", secondary_address)
        if tos is not None:
            pulumi.set(__self__, "tos", tos)
        if traffic_group is not None:
            pulumi.set(__self__, "traffic_group", traffic_group)
        if transparent is not None:
            pulumi.set(__self__, "transparent", transparent)
        if use_pmtu is not None:
            pulumi.set(__self__, "use_pmtu", use_pmtu)

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[pulumi.Input[str]]:
        """
        The application service that the object belongs to
        """
        return pulumi.get(self, "app_service")

    @app_service.setter
    def app_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_service", value)

    @property
    @pulumi.getter(name="autoLastHop")
    def auto_last_hop(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether auto lasthop is enabled or not
        """
        return pulumi.get(self, "auto_last_hop")

    @auto_last_hop.setter
    def auto_last_hop(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_last_hop", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User defined description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies an idle timeout for wildcard tunnels in seconds
        """
        return pulumi.get(self, "idle_timeout")

    @idle_timeout.setter
    def idle_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "idle_timeout", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[int]]:
        """
        The key field may represent different values depending on the type of the tunnel
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="localAddress")
    def local_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a local IP address. This option is required
        """
        return pulumi.get(self, "local_address")

    @local_address.setter
    def local_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_address", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how the tunnel carries traffic
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum transmission unit (MTU) of the tunnel
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the tunnel
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Displays the admin-partition within which this component resides
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the profile that you want to associate with the tunnel
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter(name="remoteAddress")
    def remote_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a remote IP address
        """
        return pulumi.get(self, "remote_address")

    @remote_address.setter
    def remote_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_address", value)

    @property
    @pulumi.getter(name="secondaryAddress")
    def secondary_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a secondary non-floating IP address when the local-address is set to a floating address
        """
        return pulumi.get(self, "secondary_address")

    @secondary_address.setter
    def secondary_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_address", value)

    @property
    @pulumi.getter
    def tos(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        """
        return pulumi.get(self, "tos")

    @tos.setter
    def tos(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tos", value)

    @property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a traffic-group for use with the tunnel
        """
        return pulumi.get(self, "traffic_group")

    @traffic_group.setter
    def traffic_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_group", value)

    @property
    @pulumi.getter
    def transparent(self) -> Optional[pulumi.Input[str]]:
        """
        Enables or disables the tunnel to be transparent
        """
        return pulumi.get(self, "transparent")

    @transparent.setter
    def transparent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transparent", value)

    @property
    @pulumi.getter(name="usePmtu")
    def use_pmtu(self) -> Optional[pulumi.Input[str]]:
        """
        Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        """
        return pulumi.get(self, "use_pmtu")

    @use_pmtu.setter
    def use_pmtu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "use_pmtu", value)


class NetTunnel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service: Optional[pulumi.Input[str]] = None,
                 auto_last_hop: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 key: Optional[pulumi.Input[int]] = None,
                 local_address: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 remote_address: Optional[pulumi.Input[str]] = None,
                 secondary_address: Optional[pulumi.Input[str]] = None,
                 tos: Optional[pulumi.Input[str]] = None,
                 traffic_group: Optional[pulumi.Input[str]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 use_pmtu: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `NetTunnel` Manages a tunnel configuration

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        example1 = f5bigip.NetTunnel("example1",
            name="example1",
            local_address="192.16.81.240",
            profile="/Common/dslite")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service: The application service that the object belongs to
        :param pulumi.Input[str] auto_last_hop: Specifies whether auto lasthop is enabled or not
        :param pulumi.Input[str] description: User defined description
        :param pulumi.Input[int] idle_timeout: Specifies an idle timeout for wildcard tunnels in seconds
        :param pulumi.Input[int] key: The key field may represent different values depending on the type of the tunnel
        :param pulumi.Input[str] local_address: Specifies a local IP address. This option is required
        :param pulumi.Input[str] mode: Specifies how the tunnel carries traffic
        :param pulumi.Input[int] mtu: Specifies the maximum transmission unit (MTU) of the tunnel
        :param pulumi.Input[str] name: Name of the tunnel
        :param pulumi.Input[str] partition: Displays the admin-partition within which this component resides
        :param pulumi.Input[str] profile: Specifies the profile that you want to associate with the tunnel
        :param pulumi.Input[str] remote_address: Specifies a remote IP address
        :param pulumi.Input[str] secondary_address: Specifies a secondary non-floating IP address when the local-address is set to a floating address
        :param pulumi.Input[str] tos: Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        :param pulumi.Input[str] traffic_group: Specifies a traffic-group for use with the tunnel
        :param pulumi.Input[str] transparent: Enables or disables the tunnel to be transparent
        :param pulumi.Input[str] use_pmtu: Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetTunnelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `NetTunnel` Manages a tunnel configuration

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        example1 = f5bigip.NetTunnel("example1",
            name="example1",
            local_address="192.16.81.240",
            profile="/Common/dslite")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param NetTunnelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetTunnelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service: Optional[pulumi.Input[str]] = None,
                 auto_last_hop: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 key: Optional[pulumi.Input[int]] = None,
                 local_address: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 remote_address: Optional[pulumi.Input[str]] = None,
                 secondary_address: Optional[pulumi.Input[str]] = None,
                 tos: Optional[pulumi.Input[str]] = None,
                 traffic_group: Optional[pulumi.Input[str]] = None,
                 transparent: Optional[pulumi.Input[str]] = None,
                 use_pmtu: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetTunnelArgs.__new__(NetTunnelArgs)

            __props__.__dict__["app_service"] = app_service
            __props__.__dict__["auto_last_hop"] = auto_last_hop
            __props__.__dict__["description"] = description
            __props__.__dict__["idle_timeout"] = idle_timeout
            __props__.__dict__["key"] = key
            if local_address is None and not opts.urn:
                raise TypeError("Missing required property 'local_address'")
            __props__.__dict__["local_address"] = local_address
            __props__.__dict__["mode"] = mode
            __props__.__dict__["mtu"] = mtu
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["partition"] = partition
            if profile is None and not opts.urn:
                raise TypeError("Missing required property 'profile'")
            __props__.__dict__["profile"] = profile
            __props__.__dict__["remote_address"] = remote_address
            __props__.__dict__["secondary_address"] = secondary_address
            __props__.__dict__["tos"] = tos
            __props__.__dict__["traffic_group"] = traffic_group
            __props__.__dict__["transparent"] = transparent
            __props__.__dict__["use_pmtu"] = use_pmtu
        super(NetTunnel, __self__).__init__(
            'f5bigip:index/netTunnel:NetTunnel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service: Optional[pulumi.Input[str]] = None,
            auto_last_hop: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            idle_timeout: Optional[pulumi.Input[int]] = None,
            key: Optional[pulumi.Input[int]] = None,
            local_address: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            mtu: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition: Optional[pulumi.Input[str]] = None,
            profile: Optional[pulumi.Input[str]] = None,
            remote_address: Optional[pulumi.Input[str]] = None,
            secondary_address: Optional[pulumi.Input[str]] = None,
            tos: Optional[pulumi.Input[str]] = None,
            traffic_group: Optional[pulumi.Input[str]] = None,
            transparent: Optional[pulumi.Input[str]] = None,
            use_pmtu: Optional[pulumi.Input[str]] = None) -> 'NetTunnel':
        """
        Get an existing NetTunnel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_service: The application service that the object belongs to
        :param pulumi.Input[str] auto_last_hop: Specifies whether auto lasthop is enabled or not
        :param pulumi.Input[str] description: User defined description
        :param pulumi.Input[int] idle_timeout: Specifies an idle timeout for wildcard tunnels in seconds
        :param pulumi.Input[int] key: The key field may represent different values depending on the type of the tunnel
        :param pulumi.Input[str] local_address: Specifies a local IP address. This option is required
        :param pulumi.Input[str] mode: Specifies how the tunnel carries traffic
        :param pulumi.Input[int] mtu: Specifies the maximum transmission unit (MTU) of the tunnel
        :param pulumi.Input[str] name: Name of the tunnel
        :param pulumi.Input[str] partition: Displays the admin-partition within which this component resides
        :param pulumi.Input[str] profile: Specifies the profile that you want to associate with the tunnel
        :param pulumi.Input[str] remote_address: Specifies a remote IP address
        :param pulumi.Input[str] secondary_address: Specifies a secondary non-floating IP address when the local-address is set to a floating address
        :param pulumi.Input[str] tos: Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        :param pulumi.Input[str] traffic_group: Specifies a traffic-group for use with the tunnel
        :param pulumi.Input[str] transparent: Enables or disables the tunnel to be transparent
        :param pulumi.Input[str] use_pmtu: Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetTunnelState.__new__(_NetTunnelState)

        __props__.__dict__["app_service"] = app_service
        __props__.__dict__["auto_last_hop"] = auto_last_hop
        __props__.__dict__["description"] = description
        __props__.__dict__["idle_timeout"] = idle_timeout
        __props__.__dict__["key"] = key
        __props__.__dict__["local_address"] = local_address
        __props__.__dict__["mode"] = mode
        __props__.__dict__["mtu"] = mtu
        __props__.__dict__["name"] = name
        __props__.__dict__["partition"] = partition
        __props__.__dict__["profile"] = profile
        __props__.__dict__["remote_address"] = remote_address
        __props__.__dict__["secondary_address"] = secondary_address
        __props__.__dict__["tos"] = tos
        __props__.__dict__["traffic_group"] = traffic_group
        __props__.__dict__["transparent"] = transparent
        __props__.__dict__["use_pmtu"] = use_pmtu
        return NetTunnel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> pulumi.Output[Optional[str]]:
        """
        The application service that the object belongs to
        """
        return pulumi.get(self, "app_service")

    @property
    @pulumi.getter(name="autoLastHop")
    def auto_last_hop(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether auto lasthop is enabled or not
        """
        return pulumi.get(self, "auto_last_hop")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        User defined description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies an idle timeout for wildcard tunnels in seconds
        """
        return pulumi.get(self, "idle_timeout")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[int]]:
        """
        The key field may represent different values depending on the type of the tunnel
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="localAddress")
    def local_address(self) -> pulumi.Output[str]:
        """
        Specifies a local IP address. This option is required
        """
        return pulumi.get(self, "local_address")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies how the tunnel carries traffic
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the maximum transmission unit (MTU) of the tunnel
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the tunnel
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> pulumi.Output[Optional[str]]:
        """
        Displays the admin-partition within which this component resides
        """
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[str]:
        """
        Specifies the profile that you want to associate with the tunnel
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter(name="remoteAddress")
    def remote_address(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a remote IP address
        """
        return pulumi.get(self, "remote_address")

    @property
    @pulumi.getter(name="secondaryAddress")
    def secondary_address(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a secondary non-floating IP address when the local-address is set to a floating address
        """
        return pulumi.get(self, "secondary_address")

    @property
    @pulumi.getter
    def tos(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a value for insertion into the Type of Service (ToS) octet within the IP header of the encapsulating header of transmitted packets
        """
        return pulumi.get(self, "tos")

    @property
    @pulumi.getter(name="trafficGroup")
    def traffic_group(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a traffic-group for use with the tunnel
        """
        return pulumi.get(self, "traffic_group")

    @property
    @pulumi.getter
    def transparent(self) -> pulumi.Output[Optional[str]]:
        """
        Enables or disables the tunnel to be transparent
        """
        return pulumi.get(self, "transparent")

    @property
    @pulumi.getter(name="usePmtu")
    def use_pmtu(self) -> pulumi.Output[Optional[str]]:
        """
        Enables or disables the tunnel to use the PMTU (Path MTU) information provided by ICMP NeedFrag error messages
        """
        return pulumi.get(self, "use_pmtu")

