# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TrafficSelectorArgs', 'TrafficSelector']

@pulumi.input_type
class TrafficSelectorArgs:
    def __init__(__self__, *,
                 destination_address: pulumi.Input[str],
                 name: pulumi.Input[str],
                 source_address: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[int]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[int]] = None,
                 ipsec_policy: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 source_port: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a TrafficSelector resource.
        :param pulumi.Input[str] destination_address: Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[str] name: Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        :param pulumi.Input[str] source_address: Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[str] description: Description of the traffic selector.
        :param pulumi.Input[int] destination_port: Specifies the IP port used by the application. The default value is `All Ports (0)`
        :param pulumi.Input[str] direction: Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        :param pulumi.Input[int] ip_protocol: Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        :param pulumi.Input[str] ipsec_policy: Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        :param pulumi.Input[int] order: Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
               When creating a new traffic selector, if this parameter is not specified, the default is `last`
        :param pulumi.Input[int] source_port: Specifies the IP port used by the application. The default value is `All Ports (0)`.
        """
        pulumi.set(__self__, "destination_address", destination_address)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "source_address", source_address)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_port is not None:
            pulumi.set(__self__, "destination_port", destination_port)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if ip_protocol is not None:
            pulumi.set(__self__, "ip_protocol", ip_protocol)
        if ipsec_policy is not None:
            pulumi.set(__self__, "ipsec_policy", ipsec_policy)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if source_port is not None:
            pulumi.set(__self__, "source_port", source_port)

    @property
    @pulumi.getter(name="destinationAddress")
    def destination_address(self) -> pulumi.Input[str]:
        """
        Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        """
        return pulumi.get(self, "destination_address")

    @destination_address.setter
    def destination_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_address", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sourceAddress")
    def source_address(self) -> pulumi.Input[str]:
        """
        Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        """
        return pulumi.get(self, "source_address")

    @source_address.setter
    def source_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "source_address", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the traffic selector.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the IP port used by the application. The default value is `All Ports (0)`
        """
        return pulumi.get(self, "destination_port")

    @destination_port.setter
    def destination_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "destination_port", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter(name="ipsecPolicy")
    def ipsec_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        """
        return pulumi.get(self, "ipsec_policy")

    @ipsec_policy.setter
    def ipsec_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_policy", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
        When creating a new traffic selector, if this parameter is not specified, the default is `last`
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the IP port used by the application. The default value is `All Ports (0)`.
        """
        return pulumi.get(self, "source_port")

    @source_port.setter
    def source_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "source_port", value)


@pulumi.input_type
class _TrafficSelectorState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_address: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[int]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[int]] = None,
                 ipsec_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 source_address: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering TrafficSelector resources.
        :param pulumi.Input[str] description: Description of the traffic selector.
        :param pulumi.Input[str] destination_address: Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[int] destination_port: Specifies the IP port used by the application. The default value is `All Ports (0)`
        :param pulumi.Input[str] direction: Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        :param pulumi.Input[int] ip_protocol: Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        :param pulumi.Input[str] ipsec_policy: Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        :param pulumi.Input[str] name: Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        :param pulumi.Input[int] order: Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
               When creating a new traffic selector, if this parameter is not specified, the default is `last`
        :param pulumi.Input[str] source_address: Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[int] source_port: Specifies the IP port used by the application. The default value is `All Ports (0)`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_address is not None:
            pulumi.set(__self__, "destination_address", destination_address)
        if destination_port is not None:
            pulumi.set(__self__, "destination_port", destination_port)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if ip_protocol is not None:
            pulumi.set(__self__, "ip_protocol", ip_protocol)
        if ipsec_policy is not None:
            pulumi.set(__self__, "ipsec_policy", ipsec_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if source_address is not None:
            pulumi.set(__self__, "source_address", source_address)
        if source_port is not None:
            pulumi.set(__self__, "source_port", source_port)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the traffic selector.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationAddress")
    def destination_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        """
        return pulumi.get(self, "destination_address")

    @destination_address.setter
    def destination_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_address", value)

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the IP port used by the application. The default value is `All Ports (0)`
        """
        return pulumi.get(self, "destination_port")

    @destination_port.setter
    def destination_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "destination_port", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        """
        return pulumi.get(self, "ip_protocol")

    @ip_protocol.setter
    def ip_protocol(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ip_protocol", value)

    @property
    @pulumi.getter(name="ipsecPolicy")
    def ipsec_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        """
        return pulumi.get(self, "ipsec_policy")

    @ipsec_policy.setter
    def ipsec_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipsec_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
        When creating a new traffic selector, if this parameter is not specified, the default is `last`
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter(name="sourceAddress")
    def source_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        """
        return pulumi.get(self, "source_address")

    @source_address.setter
    def source_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_address", value)

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the IP port used by the application. The default value is `All Ports (0)`.
        """
        return pulumi.get(self, "source_port")

    @source_port.setter
    def source_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "source_port", value)


class TrafficSelector(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_address: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[int]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[int]] = None,
                 ipsec_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 source_address: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        `TrafficSelector` Manage IPSec Traffic Selectors on BIG-IP

        Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-selector)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_selector = f5bigip.TrafficSelector("test-selector",
            name="/Common/test-selector",
            destination_address="3.10.11.2/32",
            source_address="2.10.11.12/32")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the traffic selector.
        :param pulumi.Input[str] destination_address: Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[int] destination_port: Specifies the IP port used by the application. The default value is `All Ports (0)`
        :param pulumi.Input[str] direction: Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        :param pulumi.Input[int] ip_protocol: Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        :param pulumi.Input[str] ipsec_policy: Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        :param pulumi.Input[str] name: Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        :param pulumi.Input[int] order: Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
               When creating a new traffic selector, if this parameter is not specified, the default is `last`
        :param pulumi.Input[str] source_address: Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[int] source_port: Specifies the IP port used by the application. The default value is `All Ports (0)`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrafficSelectorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `TrafficSelector` Manage IPSec Traffic Selectors on BIG-IP

        Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-selector)

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_selector = f5bigip.TrafficSelector("test-selector",
            name="/Common/test-selector",
            destination_address="3.10.11.2/32",
            source_address="2.10.11.12/32")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param TrafficSelectorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficSelectorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_address: Optional[pulumi.Input[str]] = None,
                 destination_port: Optional[pulumi.Input[int]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 ip_protocol: Optional[pulumi.Input[int]] = None,
                 ipsec_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 source_address: Optional[pulumi.Input[str]] = None,
                 source_port: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrafficSelectorArgs.__new__(TrafficSelectorArgs)

            __props__.__dict__["description"] = description
            if destination_address is None and not opts.urn:
                raise TypeError("Missing required property 'destination_address'")
            __props__.__dict__["destination_address"] = destination_address
            __props__.__dict__["destination_port"] = destination_port
            __props__.__dict__["direction"] = direction
            __props__.__dict__["ip_protocol"] = ip_protocol
            __props__.__dict__["ipsec_policy"] = ipsec_policy
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["order"] = order
            if source_address is None and not opts.urn:
                raise TypeError("Missing required property 'source_address'")
            __props__.__dict__["source_address"] = source_address
            __props__.__dict__["source_port"] = source_port
        super(TrafficSelector, __self__).__init__(
            'f5bigip:index/trafficSelector:TrafficSelector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination_address: Optional[pulumi.Input[str]] = None,
            destination_port: Optional[pulumi.Input[int]] = None,
            direction: Optional[pulumi.Input[str]] = None,
            ip_protocol: Optional[pulumi.Input[int]] = None,
            ipsec_policy: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            order: Optional[pulumi.Input[int]] = None,
            source_address: Optional[pulumi.Input[str]] = None,
            source_port: Optional[pulumi.Input[int]] = None) -> 'TrafficSelector':
        """
        Get an existing TrafficSelector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the traffic selector.
        :param pulumi.Input[str] destination_address: Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[int] destination_port: Specifies the IP port used by the application. The default value is `All Ports (0)`
        :param pulumi.Input[str] direction: Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        :param pulumi.Input[int] ip_protocol: Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        :param pulumi.Input[str] ipsec_policy: Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        :param pulumi.Input[str] name: Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        :param pulumi.Input[int] order: Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
               When creating a new traffic selector, if this parameter is not specified, the default is `last`
        :param pulumi.Input[str] source_address: Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        :param pulumi.Input[int] source_port: Specifies the IP port used by the application. The default value is `All Ports (0)`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrafficSelectorState.__new__(_TrafficSelectorState)

        __props__.__dict__["description"] = description
        __props__.__dict__["destination_address"] = destination_address
        __props__.__dict__["destination_port"] = destination_port
        __props__.__dict__["direction"] = direction
        __props__.__dict__["ip_protocol"] = ip_protocol
        __props__.__dict__["ipsec_policy"] = ipsec_policy
        __props__.__dict__["name"] = name
        __props__.__dict__["order"] = order
        __props__.__dict__["source_address"] = source_address
        __props__.__dict__["source_port"] = source_port
        return TrafficSelector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the traffic selector.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationAddress")
    def destination_address(self) -> pulumi.Output[str]:
        """
        Specifies the host or network IP address to which the application traffic is destined.When creating a new traffic selector, this parameter is required.
        """
        return pulumi.get(self, "destination_address")

    @property
    @pulumi.getter(name="destinationPort")
    def destination_port(self) -> pulumi.Output[int]:
        """
        Specifies the IP port used by the application. The default value is `All Ports (0)`
        """
        return pulumi.get(self, "destination_port")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[str]:
        """
        Specifies whether the traffic selector applies to inbound or outbound traffic, or both. The default value is `Both`.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> pulumi.Output[int]:
        """
        Specifies the network protocol to use for this traffic. The default value is `All Protocols (255)`
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="ipsecPolicy")
    def ipsec_policy(self) -> pulumi.Output[str]:
        """
        Specifies the IPsec policy that tells the BIG-IP system how to handle the packets.When creating a new traffic selector, if this parameter is not specified, the default is `default-ipsec-policy`.
        """
        return pulumi.get(self, "ipsec_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the IPSec traffic-selector,it should be "full path".The full path is the combination of the partition + name of the IPSec traffic-selector.(For example `/Common/test-selector`)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> pulumi.Output[int]:
        """
        Specifies the order in which traffic is matched, if traffic can be matched to multiple traffic selectors.Traffic is matched to the traffic selector with the highest priority (lowest order number).
        When creating a new traffic selector, if this parameter is not specified, the default is `last`
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="sourceAddress")
    def source_address(self) -> pulumi.Output[str]:
        """
        Specifies the host or network IP address from which the application traffic originates.When creating a new traffic selector, this parameter is required.
        """
        return pulumi.get(self, "source_address")

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> pulumi.Output[int]:
        """
        Specifies the IP port used by the application. The default value is `All Ports (0)`.
        """
        return pulumi.get(self, "source_port")

