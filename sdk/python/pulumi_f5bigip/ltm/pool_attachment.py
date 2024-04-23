# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PoolAttachmentArgs', 'PoolAttachment']

@pulumi.input_type
class PoolAttachmentArgs:
    def __init__(__self__, *,
                 node: pulumi.Input[str],
                 pool: pulumi.Input[str],
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 connection_rate_limit: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn_autopopulate: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PoolAttachment resource.
        :param pulumi.Input[str] node: Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        :param pulumi.Input[str] pool: Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
        :param pulumi.Input[int] connection_limit: Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        :param pulumi.Input[str] connection_rate_limit: Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] fqdn_autopopulate: Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        :param pulumi.Input[str] monitor: Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
        :param pulumi.Input[int] priority_group: Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        :param pulumi.Input[int] ratio: "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        :param pulumi.Input[str] state: Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
        """
        pulumi.set(__self__, "node", node)
        pulumi.set(__self__, "pool", pool)
        if connection_limit is not None:
            pulumi.set(__self__, "connection_limit", connection_limit)
        if connection_rate_limit is not None:
            pulumi.set(__self__, "connection_rate_limit", connection_rate_limit)
        if dynamic_ratio is not None:
            pulumi.set(__self__, "dynamic_ratio", dynamic_ratio)
        if fqdn_autopopulate is not None:
            pulumi.set(__self__, "fqdn_autopopulate", fqdn_autopopulate)
        if monitor is not None:
            pulumi.set(__self__, "monitor", monitor)
        if priority_group is not None:
            pulumi.set(__self__, "priority_group", priority_group)
        if ratio is not None:
            pulumi.set(__self__, "ratio", ratio)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def node(self) -> pulumi.Input[str]:
        """
        Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        """
        return pulumi.get(self, "node")

    @node.setter
    def node(self, value: pulumi.Input[str]):
        pulumi.set(self, "node", value)

    @property
    @pulumi.getter
    def pool(self) -> pulumi.Input[str]:
        """
        Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
        """
        return pulumi.get(self, "pool")

    @pool.setter
    def pool(self, value: pulumi.Input[str]):
        pulumi.set(self, "pool", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        """
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter(name="connectionRateLimit")
    def connection_rate_limit(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        """
        return pulumi.get(self, "connection_rate_limit")

    @connection_rate_limit.setter
    def connection_rate_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_rate_limit", value)

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
    @pulumi.getter(name="fqdnAutopopulate")
    def fqdn_autopopulate(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        """
        return pulumi.get(self, "fqdn_autopopulate")

    @fqdn_autopopulate.setter
    def fqdn_autopopulate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn_autopopulate", value)

    @property
    @pulumi.getter
    def monitor(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
        """
        return pulumi.get(self, "monitor")

    @monitor.setter
    def monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor", value)

    @property
    @pulumi.getter(name="priorityGroup")
    def priority_group(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        """
        return pulumi.get(self, "priority_group")

    @priority_group.setter
    def priority_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority_group", value)

    @property
    @pulumi.getter
    def ratio(self) -> Optional[pulumi.Input[int]]:
        """
        "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        """
        return pulumi.get(self, "ratio")

    @ratio.setter
    def ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ratio", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class _PoolAttachmentState:
    def __init__(__self__, *,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 connection_rate_limit: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn_autopopulate: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[str]] = None,
                 pool: Optional[pulumi.Input[str]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PoolAttachment resources.
        :param pulumi.Input[int] connection_limit: Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        :param pulumi.Input[str] connection_rate_limit: Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] fqdn_autopopulate: Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        :param pulumi.Input[str] monitor: Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
        :param pulumi.Input[str] node: Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        :param pulumi.Input[str] pool: Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
        :param pulumi.Input[int] priority_group: Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        :param pulumi.Input[int] ratio: "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        :param pulumi.Input[str] state: Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
        """
        if connection_limit is not None:
            pulumi.set(__self__, "connection_limit", connection_limit)
        if connection_rate_limit is not None:
            pulumi.set(__self__, "connection_rate_limit", connection_rate_limit)
        if dynamic_ratio is not None:
            pulumi.set(__self__, "dynamic_ratio", dynamic_ratio)
        if fqdn_autopopulate is not None:
            pulumi.set(__self__, "fqdn_autopopulate", fqdn_autopopulate)
        if monitor is not None:
            pulumi.set(__self__, "monitor", monitor)
        if node is not None:
            pulumi.set(__self__, "node", node)
        if pool is not None:
            pulumi.set(__self__, "pool", pool)
        if priority_group is not None:
            pulumi.set(__self__, "priority_group", priority_group)
        if ratio is not None:
            pulumi.set(__self__, "ratio", ratio)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        """
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter(name="connectionRateLimit")
    def connection_rate_limit(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        """
        return pulumi.get(self, "connection_rate_limit")

    @connection_rate_limit.setter
    def connection_rate_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_rate_limit", value)

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
    @pulumi.getter(name="fqdnAutopopulate")
    def fqdn_autopopulate(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        """
        return pulumi.get(self, "fqdn_autopopulate")

    @fqdn_autopopulate.setter
    def fqdn_autopopulate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn_autopopulate", value)

    @property
    @pulumi.getter
    def monitor(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
        """
        return pulumi.get(self, "monitor")

    @monitor.setter
    def monitor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "monitor", value)

    @property
    @pulumi.getter
    def node(self) -> Optional[pulumi.Input[str]]:
        """
        Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        """
        return pulumi.get(self, "node")

    @node.setter
    def node(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node", value)

    @property
    @pulumi.getter
    def pool(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
        """
        return pulumi.get(self, "pool")

    @pool.setter
    def pool(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pool", value)

    @property
    @pulumi.getter(name="priorityGroup")
    def priority_group(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        """
        return pulumi.get(self, "priority_group")

    @priority_group.setter
    def priority_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority_group", value)

    @property
    @pulumi.getter
    def ratio(self) -> Optional[pulumi.Input[int]]:
        """
        "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        """
        return pulumi.get(self, "ratio")

    @ratio.setter
    def ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ratio", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class PoolAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 connection_rate_limit: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn_autopopulate: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[str]] = None,
                 pool: Optional[pulumi.Input[str]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `ltm.PoolAttachment` Manages nodes membership in pools

        ## Example Usage

        There are two ways to use `ltm.PoolAttachment` resource for `node` attribute

        * It can be reference from `ltm.Node` (or)
        * It can be specify directly with `ipv4:port`/`fqdn:port`/`ipv6.port` which will also create node and attach member to pool.

        > For adding IPv6 node/member to pool it should be specific in `node` attribute in format like `ipv6_address.port`.
        IPv4 should be specified as `ipv4_address:port`

        ### Usage Pool attachment with node/member directly attaching to pool.

        node can be specified in format `ipv4:port` / `fqdn:port` / `ipv6.port`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        monitor = f5bigip.ltm.Monitor("monitor",
            name="/Common/terraform_monitor",
            parent="/Common/http",
            send="GET /some/path\\x0d\\n",
            timeout=999,
            interval=998)
        pool = f5bigip.ltm.Pool("pool",
            name="/Common/terraform-pool",
            load_balancing_mode="round-robin",
            monitors=[monitor.name],
            allow_snat="yes",
            allow_nat="yes")
        # attaching ipv4 address with service port
        ipv4_node_attach = f5bigip.ltm.PoolAttachment("ipv4_node_attach",
            pool=pool.name,
            node="1.1.1.1:80")
        # attaching ipv6 address with service port
        ipv6_node_attach = f5bigip.ltm.PoolAttachment("ipv6_node_attach",
            pool=pool.name,
            node="2003::4.80")
        ```

        ### Usage Pool attachment with node referenced from `ltm.Node`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        monitor = f5bigip.ltm.Monitor("monitor",
            name="/Common/terraform_monitor",
            parent="/Common/http",
            send="GET /some/path\\x0d\\n",
            timeout=999,
            interval=998)
        pool = f5bigip.ltm.Pool("pool",
            name="/Common/terraform-pool",
            load_balancing_mode="round-robin",
            monitors=[monitor.name],
            allow_snat="yes",
            allow_nat="yes")
        node = f5bigip.ltm.Node("node",
            name="/Common/terraform_node",
            address="192.168.30.2")
        attach_node = f5bigip.ltm.PoolAttachment("attach_node",
            pool=pool.name,
            node=node.name.apply(lambda name: f"{name}:80"))
        ```

        ## Importing

        An existing pool attachment (i.e. pool membership) can be imported into this resource by supplying both the pool full path, and the node full path with the relevant port. If the pool or node membership is not found, an error will be returned. An example is below:

        ```sh
        $ terraform import bigip_ltm_pool_attachment.node-pool-attach \\
        	'{"pool": "/Common/terraform-pool", "node": "/Common/node1:80"}'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] connection_limit: Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        :param pulumi.Input[str] connection_rate_limit: Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] fqdn_autopopulate: Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        :param pulumi.Input[str] monitor: Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
        :param pulumi.Input[str] node: Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        :param pulumi.Input[str] pool: Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
        :param pulumi.Input[int] priority_group: Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        :param pulumi.Input[int] ratio: "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        :param pulumi.Input[str] state: Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PoolAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.PoolAttachment` Manages nodes membership in pools

        ## Example Usage

        There are two ways to use `ltm.PoolAttachment` resource for `node` attribute

        * It can be reference from `ltm.Node` (or)
        * It can be specify directly with `ipv4:port`/`fqdn:port`/`ipv6.port` which will also create node and attach member to pool.

        > For adding IPv6 node/member to pool it should be specific in `node` attribute in format like `ipv6_address.port`.
        IPv4 should be specified as `ipv4_address:port`

        ### Usage Pool attachment with node/member directly attaching to pool.

        node can be specified in format `ipv4:port` / `fqdn:port` / `ipv6.port`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        monitor = f5bigip.ltm.Monitor("monitor",
            name="/Common/terraform_monitor",
            parent="/Common/http",
            send="GET /some/path\\x0d\\n",
            timeout=999,
            interval=998)
        pool = f5bigip.ltm.Pool("pool",
            name="/Common/terraform-pool",
            load_balancing_mode="round-robin",
            monitors=[monitor.name],
            allow_snat="yes",
            allow_nat="yes")
        # attaching ipv4 address with service port
        ipv4_node_attach = f5bigip.ltm.PoolAttachment("ipv4_node_attach",
            pool=pool.name,
            node="1.1.1.1:80")
        # attaching ipv6 address with service port
        ipv6_node_attach = f5bigip.ltm.PoolAttachment("ipv6_node_attach",
            pool=pool.name,
            node="2003::4.80")
        ```

        ### Usage Pool attachment with node referenced from `ltm.Node`

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        monitor = f5bigip.ltm.Monitor("monitor",
            name="/Common/terraform_monitor",
            parent="/Common/http",
            send="GET /some/path\\x0d\\n",
            timeout=999,
            interval=998)
        pool = f5bigip.ltm.Pool("pool",
            name="/Common/terraform-pool",
            load_balancing_mode="round-robin",
            monitors=[monitor.name],
            allow_snat="yes",
            allow_nat="yes")
        node = f5bigip.ltm.Node("node",
            name="/Common/terraform_node",
            address="192.168.30.2")
        attach_node = f5bigip.ltm.PoolAttachment("attach_node",
            pool=pool.name,
            node=node.name.apply(lambda name: f"{name}:80"))
        ```

        ## Importing

        An existing pool attachment (i.e. pool membership) can be imported into this resource by supplying both the pool full path, and the node full path with the relevant port. If the pool or node membership is not found, an error will be returned. An example is below:

        ```sh
        $ terraform import bigip_ltm_pool_attachment.node-pool-attach \\
        	'{"pool": "/Common/terraform-pool", "node": "/Common/node1:80"}'
        ```

        :param str resource_name: The name of the resource.
        :param PoolAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PoolAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 connection_rate_limit: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn_autopopulate: Optional[pulumi.Input[str]] = None,
                 monitor: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[str]] = None,
                 pool: Optional[pulumi.Input[str]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PoolAttachmentArgs.__new__(PoolAttachmentArgs)

            __props__.__dict__["connection_limit"] = connection_limit
            __props__.__dict__["connection_rate_limit"] = connection_rate_limit
            __props__.__dict__["dynamic_ratio"] = dynamic_ratio
            __props__.__dict__["fqdn_autopopulate"] = fqdn_autopopulate
            __props__.__dict__["monitor"] = monitor
            if node is None and not opts.urn:
                raise TypeError("Missing required property 'node'")
            __props__.__dict__["node"] = node
            if pool is None and not opts.urn:
                raise TypeError("Missing required property 'pool'")
            __props__.__dict__["pool"] = pool
            __props__.__dict__["priority_group"] = priority_group
            __props__.__dict__["ratio"] = ratio
            __props__.__dict__["state"] = state
        super(PoolAttachment, __self__).__init__(
            'f5bigip:ltm/poolAttachment:PoolAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_limit: Optional[pulumi.Input[int]] = None,
            connection_rate_limit: Optional[pulumi.Input[str]] = None,
            dynamic_ratio: Optional[pulumi.Input[int]] = None,
            fqdn_autopopulate: Optional[pulumi.Input[str]] = None,
            monitor: Optional[pulumi.Input[str]] = None,
            node: Optional[pulumi.Input[str]] = None,
            pool: Optional[pulumi.Input[str]] = None,
            priority_group: Optional[pulumi.Input[int]] = None,
            ratio: Optional[pulumi.Input[int]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'PoolAttachment':
        """
        Get an existing PoolAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] connection_limit: Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        :param pulumi.Input[str] connection_rate_limit: Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] fqdn_autopopulate: Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        :param pulumi.Input[str] monitor: Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
        :param pulumi.Input[str] node: Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        :param pulumi.Input[str] pool: Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
        :param pulumi.Input[int] priority_group: Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        :param pulumi.Input[int] ratio: "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        :param pulumi.Input[str] state: Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PoolAttachmentState.__new__(_PoolAttachmentState)

        __props__.__dict__["connection_limit"] = connection_limit
        __props__.__dict__["connection_rate_limit"] = connection_rate_limit
        __props__.__dict__["dynamic_ratio"] = dynamic_ratio
        __props__.__dict__["fqdn_autopopulate"] = fqdn_autopopulate
        __props__.__dict__["monitor"] = monitor
        __props__.__dict__["node"] = node
        __props__.__dict__["pool"] = pool
        __props__.__dict__["priority_group"] = priority_group
        __props__.__dict__["ratio"] = ratio
        __props__.__dict__["state"] = state
        return PoolAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> pulumi.Output[int]:
        """
        Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        """
        return pulumi.get(self, "connection_limit")

    @property
    @pulumi.getter(name="connectionRateLimit")
    def connection_rate_limit(self) -> pulumi.Output[str]:
        """
        Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        """
        return pulumi.get(self, "connection_rate_limit")

    @property
    @pulumi.getter(name="dynamicRatio")
    def dynamic_ratio(self) -> pulumi.Output[int]:
        """
        Specifies the fixed ratio value used for a node during ratio load balancing.
        """
        return pulumi.get(self, "dynamic_ratio")

    @property
    @pulumi.getter(name="fqdnAutopopulate")
    def fqdn_autopopulate(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        """
        return pulumi.get(self, "fqdn_autopopulate")

    @property
    @pulumi.getter
    def monitor(self) -> pulumi.Output[str]:
        """
        Specifies the health monitors that the system uses to monitor this pool member,value can be `none` (or) `default` (or) list of monitors joined with and ( ex: `/Common/test_monitor_pa_tc1 and /Common/gateway_icmp`).
        """
        return pulumi.get(self, "monitor")

    @property
    @pulumi.getter
    def node(self) -> pulumi.Output[str]:
        """
        Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        """
        return pulumi.get(self, "node")

    @property
    @pulumi.getter
    def pool(self) -> pulumi.Output[str]:
        """
        Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`) or partition + directory + name of the pool (For example `/Common/test/my-pool`).When including directory in fullpath we have to make sure it is created in the given partition before using it.
        """
        return pulumi.get(self, "pool")

    @property
    @pulumi.getter(name="priorityGroup")
    def priority_group(self) -> pulumi.Output[int]:
        """
        Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        """
        return pulumi.get(self, "priority_group")

    @property
    @pulumi.getter
    def ratio(self) -> pulumi.Output[int]:
        """
        "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        """
        return pulumi.get(self, "ratio")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the state the pool member should be in,value can be `enabled` (or) `disabled` (or) `forced_offline`).
        """
        return pulumi.get(self, "state")

