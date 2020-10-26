# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Pool']


class Pool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_nat: Optional[pulumi.Input[str]] = None,
                 allow_snat: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 load_balancing_mode: Optional[pulumi.Input[str]] = None,
                 minimum_active_members: Optional[pulumi.Input[int]] = None,
                 monitors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 reselect_tries: Optional[pulumi.Input[int]] = None,
                 service_down_action: Optional[pulumi.Input[str]] = None,
                 slow_ramp_time: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.Pool` Manages a pool configuration.

        Resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        monitor = f5bigip.ltm.Monitor("monitor",
            name="/Common/terraform_monitor",
            parent="/Common/http")
        pool = f5bigip.ltm.Pool("pool",
            name="/Common/Axiom_Environment_APP1_Pool",
            load_balancing_mode="round-robin",
            minimum_active_members=1,
            monitors=[monitor.name])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allow_nat: Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
        :param pulumi.Input[str] allow_snat: Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
        :param pulumi.Input[str] description: Specifies descriptive text that identifies the pool.
        :param pulumi.Input[str] load_balancing_mode: Specifies the load balancing method. The default is Round Robin.
        :param pulumi.Input[int] minimum_active_members: Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitors: List of monitor names to associate with the pool
        :param pulumi.Input[str] name: Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
        :param pulumi.Input[int] reselect_tries: Specifies the number of times the system tries to contact a new pool member after a passive failure.
        :param pulumi.Input[str] service_down_action: Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
        :param pulumi.Input[int] slow_ramp_time: Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
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

            __props__['allow_nat'] = allow_nat
            __props__['allow_snat'] = allow_snat
            __props__['description'] = description
            __props__['load_balancing_mode'] = load_balancing_mode
            __props__['minimum_active_members'] = minimum_active_members
            __props__['monitors'] = monitors
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['reselect_tries'] = reselect_tries
            __props__['service_down_action'] = service_down_action
            __props__['slow_ramp_time'] = slow_ramp_time
        super(Pool, __self__).__init__(
            'f5bigip:ltm/pool:Pool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_nat: Optional[pulumi.Input[str]] = None,
            allow_snat: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            load_balancing_mode: Optional[pulumi.Input[str]] = None,
            minimum_active_members: Optional[pulumi.Input[int]] = None,
            monitors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            reselect_tries: Optional[pulumi.Input[int]] = None,
            service_down_action: Optional[pulumi.Input[str]] = None,
            slow_ramp_time: Optional[pulumi.Input[int]] = None) -> 'Pool':
        """
        Get an existing Pool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allow_nat: Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
        :param pulumi.Input[str] allow_snat: Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
        :param pulumi.Input[str] description: Specifies descriptive text that identifies the pool.
        :param pulumi.Input[str] load_balancing_mode: Specifies the load balancing method. The default is Round Robin.
        :param pulumi.Input[int] minimum_active_members: Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] monitors: List of monitor names to associate with the pool
        :param pulumi.Input[str] name: Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
        :param pulumi.Input[int] reselect_tries: Specifies the number of times the system tries to contact a new pool member after a passive failure.
        :param pulumi.Input[str] service_down_action: Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
        :param pulumi.Input[int] slow_ramp_time: Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_nat"] = allow_nat
        __props__["allow_snat"] = allow_snat
        __props__["description"] = description
        __props__["load_balancing_mode"] = load_balancing_mode
        __props__["minimum_active_members"] = minimum_active_members
        __props__["monitors"] = monitors
        __props__["name"] = name
        __props__["reselect_tries"] = reselect_tries
        __props__["service_down_action"] = service_down_action
        __props__["slow_ramp_time"] = slow_ramp_time
        return Pool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowNat")
    def allow_nat(self) -> pulumi.Output[str]:
        """
        Specifies whether NATs are automatically enabled or disabled for any connections using this pool, [ Default : `yes`, Possible Values `yes` or `no`].
        """
        return pulumi.get(self, "allow_nat")

    @property
    @pulumi.getter(name="allowSnat")
    def allow_snat(self) -> pulumi.Output[str]:
        """
        Specifies whether SNATs are automatically enabled or disabled for any connections using this pool,[ Default : `yes`, Possible Values `yes` or `no`].
        """
        return pulumi.get(self, "allow_snat")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies descriptive text that identifies the pool.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="loadBalancingMode")
    def load_balancing_mode(self) -> pulumi.Output[str]:
        """
        Specifies the load balancing method. The default is Round Robin.
        """
        return pulumi.get(self, "load_balancing_mode")

    @property
    @pulumi.getter(name="minimumActiveMembers")
    def minimum_active_members(self) -> pulumi.Output[int]:
        """
        Specifies whether the system load balances traffic according to the priority number assigned to the pool member,Default Value is `0` meaning `disabled`.
        """
        return pulumi.get(self, "minimum_active_members")

    @property
    @pulumi.getter
    def monitors(self) -> pulumi.Output[Sequence[str]]:
        """
        List of monitor names to associate with the pool
        """
        return pulumi.get(self, "monitors")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the pool,it should be "full path".The full path is the combination of the partition + name of the pool.(For example `/Common/my-pool`)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="reselectTries")
    def reselect_tries(self) -> pulumi.Output[int]:
        """
        Specifies the number of times the system tries to contact a new pool member after a passive failure.
        """
        return pulumi.get(self, "reselect_tries")

    @property
    @pulumi.getter(name="serviceDownAction")
    def service_down_action(self) -> pulumi.Output[str]:
        """
        Specifies how the system should respond when the target pool member becomes unavailable. The default is `None`, Possible values: `[none, reset, reselect, drop]`.
        """
        return pulumi.get(self, "service_down_action")

    @property
    @pulumi.getter(name="slowRampTime")
    def slow_ramp_time(self) -> pulumi.Output[int]:
        """
        Specifies the duration during which the system sends less traffic to a newly-enabled pool member.
        """
        return pulumi.get(self, "slow_ramp_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

