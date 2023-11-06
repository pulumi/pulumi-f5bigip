# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['NodeArgs', 'Node']

@pulumi.input_type
class NodeArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 name: pulumi.Input[str],
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn: Optional[pulumi.Input['NodeFqdnArgs']] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 rate_limit: Optional[pulumi.Input[str]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Node resource.
        :param pulumi.Input[str] address: IP or hostname of the node
        :param pulumi.Input[str] name: Name of the node
        :param pulumi.Input[int] connection_limit: Specifies the maximum number of connections allowed for the node or node address.
        :param pulumi.Input[str] description: User-defined description give ltm_node
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] monitor: specifies the name of the monitor or monitor rule that you want to associate with the node.
        :param pulumi.Input[str] rate_limit: Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        :param pulumi.Input[int] ratio: Sets the ratio number for the node.
        :param pulumi.Input[str] session: Enables or disables the node for new sessions. The default value is user-enabled.
        :param pulumi.Input[str] state: Default is "user-up" you can set to "user-down" if you want to disable
               
               > *NOTE* Below attributes needs to be configured under fqdn option.
        """
        NodeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address=address,
            name=name,
            connection_limit=connection_limit,
            description=description,
            dynamic_ratio=dynamic_ratio,
            fqdn=fqdn,
            monitor=monitor,
            rate_limit=rate_limit,
            ratio=ratio,
            session=session,
            state=state,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             connection_limit: Optional[pulumi.Input[int]] = None,
             description: Optional[pulumi.Input[str]] = None,
             dynamic_ratio: Optional[pulumi.Input[int]] = None,
             fqdn: Optional[pulumi.Input['NodeFqdnArgs']] = None,
             monitor: Optional[pulumi.Input[str]] = None,
             rate_limit: Optional[pulumi.Input[str]] = None,
             ratio: Optional[pulumi.Input[int]] = None,
             session: Optional[pulumi.Input[str]] = None,
             state: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if address is None:
            raise TypeError("Missing 'address' argument")
        if name is None:
            raise TypeError("Missing 'name' argument")
        if connection_limit is None and 'connectionLimit' in kwargs:
            connection_limit = kwargs['connectionLimit']
        if dynamic_ratio is None and 'dynamicRatio' in kwargs:
            dynamic_ratio = kwargs['dynamicRatio']
        if rate_limit is None and 'rateLimit' in kwargs:
            rate_limit = kwargs['rateLimit']

        _setter("address", address)
        _setter("name", name)
        if connection_limit is not None:
            _setter("connection_limit", connection_limit)
        if description is not None:
            _setter("description", description)
        if dynamic_ratio is not None:
            _setter("dynamic_ratio", dynamic_ratio)
        if fqdn is not None:
            _setter("fqdn", fqdn)
        if monitor is not None:
            _setter("monitor", monitor)
        if rate_limit is not None:
            _setter("rate_limit", rate_limit)
        if ratio is not None:
            _setter("ratio", ratio)
        if session is not None:
            _setter("session", session)
        if state is not None:
            _setter("state", state)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        IP or hostname of the node
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the node
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of connections allowed for the node or node address.
        """
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined description give ltm_node
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dynamicRatio")
    def dynamic_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the fixed ratio value used for a node during ratio load balancing.
        """
        return pulumi.get(self, "dynamic_ratio")

    @dynamic_ratio.setter
    def dynamic_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dynamic_ratio", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input['NodeFqdnArgs']]:
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input['NodeFqdnArgs']]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter
    def monitor(self) -> Optional[pulumi.Input[str]]:
        """
        specifies the name of the monitor or monitor rule that you want to associate with the node.
        """
        return pulumi.get(self, "monitor")

    @monitor.setter
    def monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor", value)

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        """
        return pulumi.get(self, "rate_limit")

    @rate_limit.setter
    def rate_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rate_limit", value)

    @property
    @pulumi.getter
    def ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Sets the ratio number for the node.
        """
        return pulumi.get(self, "ratio")

    @ratio.setter
    def ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ratio", value)

    @property
    @pulumi.getter
    def session(self) -> Optional[pulumi.Input[str]]:
        """
        Enables or disables the node for new sessions. The default value is user-enabled.
        """
        return pulumi.get(self, "session")

    @session.setter
    def session(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Default is "user-up" you can set to "user-down" if you want to disable

        > *NOTE* Below attributes needs to be configured under fqdn option.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class _NodeState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn: Optional[pulumi.Input['NodeFqdnArgs']] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_limit: Optional[pulumi.Input[str]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Node resources.
        :param pulumi.Input[str] address: IP or hostname of the node
        :param pulumi.Input[int] connection_limit: Specifies the maximum number of connections allowed for the node or node address.
        :param pulumi.Input[str] description: User-defined description give ltm_node
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] monitor: specifies the name of the monitor or monitor rule that you want to associate with the node.
        :param pulumi.Input[str] name: Name of the node
        :param pulumi.Input[str] rate_limit: Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        :param pulumi.Input[int] ratio: Sets the ratio number for the node.
        :param pulumi.Input[str] session: Enables or disables the node for new sessions. The default value is user-enabled.
        :param pulumi.Input[str] state: Default is "user-up" you can set to "user-down" if you want to disable
               
               > *NOTE* Below attributes needs to be configured under fqdn option.
        """
        _NodeState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            address=address,
            connection_limit=connection_limit,
            description=description,
            dynamic_ratio=dynamic_ratio,
            fqdn=fqdn,
            monitor=monitor,
            name=name,
            rate_limit=rate_limit,
            ratio=ratio,
            session=session,
            state=state,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             address: Optional[pulumi.Input[str]] = None,
             connection_limit: Optional[pulumi.Input[int]] = None,
             description: Optional[pulumi.Input[str]] = None,
             dynamic_ratio: Optional[pulumi.Input[int]] = None,
             fqdn: Optional[pulumi.Input['NodeFqdnArgs']] = None,
             monitor: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             rate_limit: Optional[pulumi.Input[str]] = None,
             ratio: Optional[pulumi.Input[int]] = None,
             session: Optional[pulumi.Input[str]] = None,
             state: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if connection_limit is None and 'connectionLimit' in kwargs:
            connection_limit = kwargs['connectionLimit']
        if dynamic_ratio is None and 'dynamicRatio' in kwargs:
            dynamic_ratio = kwargs['dynamicRatio']
        if rate_limit is None and 'rateLimit' in kwargs:
            rate_limit = kwargs['rateLimit']

        if address is not None:
            _setter("address", address)
        if connection_limit is not None:
            _setter("connection_limit", connection_limit)
        if description is not None:
            _setter("description", description)
        if dynamic_ratio is not None:
            _setter("dynamic_ratio", dynamic_ratio)
        if fqdn is not None:
            _setter("fqdn", fqdn)
        if monitor is not None:
            _setter("monitor", monitor)
        if name is not None:
            _setter("name", name)
        if rate_limit is not None:
            _setter("rate_limit", rate_limit)
        if ratio is not None:
            _setter("ratio", ratio)
        if session is not None:
            _setter("session", session)
        if state is not None:
            _setter("state", state)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        IP or hostname of the node
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of connections allowed for the node or node address.
        """
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined description give ltm_node
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dynamicRatio")
    def dynamic_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the fixed ratio value used for a node during ratio load balancing.
        """
        return pulumi.get(self, "dynamic_ratio")

    @dynamic_ratio.setter
    def dynamic_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dynamic_ratio", value)

    @property
    @pulumi.getter
    def fqdn(self) -> Optional[pulumi.Input['NodeFqdnArgs']]:
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input['NodeFqdnArgs']]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter
    def monitor(self) -> Optional[pulumi.Input[str]]:
        """
        specifies the name of the monitor or monitor rule that you want to associate with the node.
        """
        return pulumi.get(self, "monitor")

    @monitor.setter
    def monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the node
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        """
        return pulumi.get(self, "rate_limit")

    @rate_limit.setter
    def rate_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rate_limit", value)

    @property
    @pulumi.getter
    def ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Sets the ratio number for the node.
        """
        return pulumi.get(self, "ratio")

    @ratio.setter
    def ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ratio", value)

    @property
    @pulumi.getter
    def session(self) -> Optional[pulumi.Input[str]]:
        """
        Enables or disables the node for new sessions. The default value is user-enabled.
        """
        return pulumi.get(self, "session")

    @session.setter
    def session(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Default is "user-up" you can set to "user-down" if you want to disable

        > *NOTE* Below attributes needs to be configured under fqdn option.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class Node(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn: Optional[pulumi.Input[pulumi.InputType['NodeFqdnArgs']]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_limit: Optional[pulumi.Input[str]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `ltm.Node` Manages a node configuration

        For resources should be named with their `full path`.The full path is the combination of the `partition + name` of the resource( example: `/Common/my-node` ) or `partition + Direcroty + name` of the resource ( example: `/Common/test/my-node` ).
        When including directory in `full path` we have to make sure it is created in the given partition before using it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        node = f5bigip.ltm.Node("node",
            address="192.168.30.1",
            connection_limit=0,
            description="Test-Node",
            dynamic_ratio=1,
            fqdn=f5bigip.ltm.NodeFqdnArgs(
                address_family="ipv4",
                interval="3000",
            ),
            monitor="/Common/icmp",
            name="/Common/terraform_node1",
            rate_limit="disabled")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IP or hostname of the node
        :param pulumi.Input[int] connection_limit: Specifies the maximum number of connections allowed for the node or node address.
        :param pulumi.Input[str] description: User-defined description give ltm_node
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] monitor: specifies the name of the monitor or monitor rule that you want to associate with the node.
        :param pulumi.Input[str] name: Name of the node
        :param pulumi.Input[str] rate_limit: Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        :param pulumi.Input[int] ratio: Sets the ratio number for the node.
        :param pulumi.Input[str] session: Enables or disables the node for new sessions. The default value is user-enabled.
        :param pulumi.Input[str] state: Default is "user-up" you can set to "user-down" if you want to disable
               
               > *NOTE* Below attributes needs to be configured under fqdn option.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.Node` Manages a node configuration

        For resources should be named with their `full path`.The full path is the combination of the `partition + name` of the resource( example: `/Common/my-node` ) or `partition + Direcroty + name` of the resource ( example: `/Common/test/my-node` ).
        When including directory in `full path` we have to make sure it is created in the given partition before using it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        node = f5bigip.ltm.Node("node",
            address="192.168.30.1",
            connection_limit=0,
            description="Test-Node",
            dynamic_ratio=1,
            fqdn=f5bigip.ltm.NodeFqdnArgs(
                address_family="ipv4",
                interval="3000",
            ),
            monitor="/Common/icmp",
            name="/Common/terraform_node1",
            rate_limit="disabled")
        ```

        :param str resource_name: The name of the resource.
        :param NodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            NodeArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn: Optional[pulumi.Input[pulumi.InputType['NodeFqdnArgs']]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_limit: Optional[pulumi.Input[str]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 session: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NodeArgs.__new__(NodeArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["connection_limit"] = connection_limit
            __props__.__dict__["description"] = description
            __props__.__dict__["dynamic_ratio"] = dynamic_ratio
            if fqdn is not None and not isinstance(fqdn, NodeFqdnArgs):
                fqdn = fqdn or {}
                def _setter(key, value):
                    fqdn[key] = value
                NodeFqdnArgs._configure(_setter, **fqdn)
            __props__.__dict__["fqdn"] = fqdn
            __props__.__dict__["monitor"] = monitor
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["rate_limit"] = rate_limit
            __props__.__dict__["ratio"] = ratio
            __props__.__dict__["session"] = session
            __props__.__dict__["state"] = state
        super(Node, __self__).__init__(
            'f5bigip:ltm/node:Node',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            connection_limit: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamic_ratio: Optional[pulumi.Input[int]] = None,
            fqdn: Optional[pulumi.Input[pulumi.InputType['NodeFqdnArgs']]] = None,
            monitor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rate_limit: Optional[pulumi.Input[str]] = None,
            ratio: Optional[pulumi.Input[int]] = None,
            session: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'Node':
        """
        Get an existing Node resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IP or hostname of the node
        :param pulumi.Input[int] connection_limit: Specifies the maximum number of connections allowed for the node or node address.
        :param pulumi.Input[str] description: User-defined description give ltm_node
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] monitor: specifies the name of the monitor or monitor rule that you want to associate with the node.
        :param pulumi.Input[str] name: Name of the node
        :param pulumi.Input[str] rate_limit: Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        :param pulumi.Input[int] ratio: Sets the ratio number for the node.
        :param pulumi.Input[str] session: Enables or disables the node for new sessions. The default value is user-enabled.
        :param pulumi.Input[str] state: Default is "user-up" you can set to "user-down" if you want to disable
               
               > *NOTE* Below attributes needs to be configured under fqdn option.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NodeState.__new__(_NodeState)

        __props__.__dict__["address"] = address
        __props__.__dict__["connection_limit"] = connection_limit
        __props__.__dict__["description"] = description
        __props__.__dict__["dynamic_ratio"] = dynamic_ratio
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["monitor"] = monitor
        __props__.__dict__["name"] = name
        __props__.__dict__["rate_limit"] = rate_limit
        __props__.__dict__["ratio"] = ratio
        __props__.__dict__["session"] = session
        __props__.__dict__["state"] = state
        return Node(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        IP or hostname of the node
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> pulumi.Output[int]:
        """
        Specifies the maximum number of connections allowed for the node or node address.
        """
        return pulumi.get(self, "connection_limit")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        User-defined description give ltm_node
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamicRatio")
    def dynamic_ratio(self) -> pulumi.Output[int]:
        """
        Specifies the fixed ratio value used for a node during ratio load balancing.
        """
        return pulumi.get(self, "dynamic_ratio")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[Optional['outputs.NodeFqdn']]:
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def monitor(self) -> pulumi.Output[Optional[str]]:
        """
        specifies the name of the monitor or monitor rule that you want to associate with the node.
        """
        return pulumi.get(self, "monitor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the node
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> pulumi.Output[str]:
        """
        Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        """
        return pulumi.get(self, "rate_limit")

    @property
    @pulumi.getter
    def ratio(self) -> pulumi.Output[int]:
        """
        Sets the ratio number for the node.
        """
        return pulumi.get(self, "ratio")

    @property
    @pulumi.getter
    def session(self) -> pulumi.Output[str]:
        """
        Enables or disables the node for new sessions. The default value is user-enabled.
        """
        return pulumi.get(self, "session")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Default is "user-up" you can set to "user-down" if you want to disable

        > *NOTE* Below attributes needs to be configured under fqdn option.
        """
        return pulumi.get(self, "state")

