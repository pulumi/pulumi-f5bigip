# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['PoolAttachment']


class PoolAttachment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 connection_rate_limit: Optional[pulumi.Input[str]] = None,
                 dynamic_ratio: Optional[pulumi.Input[int]] = None,
                 fqdn_autopopulate: Optional[pulumi.Input[str]] = None,
                 node: Optional[pulumi.Input[str]] = None,
                 pool: Optional[pulumi.Input[str]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 ratio: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.PoolAttachment` Manages nodes membership in pools

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        monitor = f5bigip.ltm.Monitor("monitor",
            name="/Common/terraform_monitor",
            parent="/Common/http",
            send="GET /some/path\n",
            timeout=999,
            interval=998)
        pool = f5bigip.ltm.Pool("pool",
            name="/Common/terraform-pool",
            load_balancing_mode="round-robin",
            monitors=[monitor.name],
            allow_snat="yes",
            allow_nat="yes")
        attach_node = f5bigip.ltm.PoolAttachment("attachNode",
            pool=pool.name,
            node="1.1.1.1:80",
            ratio=2,
            connection_limit=2,
            connection_rate_limit="2",
            priority_group=2,
            dynamic_ratio=3)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] connection_limit: Specifies a maximum established connection limit for a pool member or node.The default is 0, meaning that there is no limit to the number of connections.
        :param pulumi.Input[str] connection_rate_limit: Specifies the maximum number of connections-per-second allowed for a pool member,The default is 0.
        :param pulumi.Input[int] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] fqdn_autopopulate: Specifies whether the system automatically creates ephemeral nodes using the IP addresses returned by the resolution of a DNS query for a node defined by an FQDN. The default is enabled
        :param pulumi.Input[str] node: Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        :param pulumi.Input[str] pool: Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
        :param pulumi.Input[int] priority_group: Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        :param pulumi.Input[int] ratio: "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
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

            __props__['connection_limit'] = connection_limit
            __props__['connection_rate_limit'] = connection_rate_limit
            __props__['dynamic_ratio'] = dynamic_ratio
            __props__['fqdn_autopopulate'] = fqdn_autopopulate
            if node is None and not opts.urn:
                raise TypeError("Missing required property 'node'")
            __props__['node'] = node
            if pool is None and not opts.urn:
                raise TypeError("Missing required property 'pool'")
            __props__['pool'] = pool
            __props__['priority_group'] = priority_group
            __props__['ratio'] = ratio
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
            node: Optional[pulumi.Input[str]] = None,
            pool: Optional[pulumi.Input[str]] = None,
            priority_group: Optional[pulumi.Input[int]] = None,
            ratio: Optional[pulumi.Input[int]] = None) -> 'PoolAttachment':
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
        :param pulumi.Input[str] node: Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        :param pulumi.Input[str] pool: Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
        :param pulumi.Input[int] priority_group: Specifies a number representing the priority group for the pool member. The default is 0, meaning that the member has no priority
        :param pulumi.Input[int] ratio: "Specifies the ratio weight to assign to the pool member. Valid values range from 1 through 65535. The default is 1, which means that each pool member has an equal ratio proportion.".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["connection_limit"] = connection_limit
        __props__["connection_rate_limit"] = connection_rate_limit
        __props__["dynamic_ratio"] = dynamic_ratio
        __props__["fqdn_autopopulate"] = fqdn_autopopulate
        __props__["node"] = node
        __props__["pool"] = pool
        __props__["priority_group"] = priority_group
        __props__["ratio"] = ratio
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
    def node(self) -> pulumi.Output[str]:
        """
        Pool member address/fqdn with service port, (ex: `1.1.1.1:80/www.google.com:80`). (Note: Member will be in same partition of Pool)
        """
        return pulumi.get(self, "node")

    @property
    @pulumi.getter
    def pool(self) -> pulumi.Output[str]:
        """
        Name of the pool to which members should be attached,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

