# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ProfileFastHttp']


class ProfileFastHttp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connpool_maxreuse: Optional[pulumi.Input[int]] = None,
                 connpool_maxsize: Optional[pulumi.Input[int]] = None,
                 connpool_minsize: Optional[pulumi.Input[int]] = None,
                 connpool_replenish: Optional[pulumi.Input[str]] = None,
                 connpool_step: Optional[pulumi.Input[int]] = None,
                 connpoolidle_timeoutoverride: Optional[pulumi.Input[int]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 forcehttp10response: Optional[pulumi.Input[str]] = None,
                 idle_timeout: Optional[pulumi.Input[int]] = None,
                 maxheader_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.ProfileFastHttp` Configures a custom profile_fasthttp for use by health checks.

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        sjfasthttpprofile = f5bigip.ltm.ProfileFastHttp("sjfasthttpprofile",
            connpool_maxreuse=2,
            connpool_maxsize=2048,
            connpool_minsize=0,
            connpool_replenish="enabled",
            connpool_step=4,
            connpoolidle_timeoutoverride=0,
            defaults_from="/Common/fasthttp",
            forcehttp10response="disabled",
            idle_timeout=300,
            maxheader_size=32768,
            name="/Common/sjfasthttpprofile")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] connpool_maxreuse: Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        :param pulumi.Input[int] connpool_maxsize: Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        :param pulumi.Input[int] connpool_minsize: Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        :param pulumi.Input[str] connpool_replenish: The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        :param pulumi.Input[int] connpool_step: Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        :param pulumi.Input[int] connpoolidle_timeoutoverride: Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] forcehttp10response: Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        :param pulumi.Input[int] idle_timeout: Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        :param pulumi.Input[int] maxheader_size: Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        :param pulumi.Input[str] name: Name of the profile_fasthttp
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

            __props__['connpool_maxreuse'] = connpool_maxreuse
            __props__['connpool_maxsize'] = connpool_maxsize
            __props__['connpool_minsize'] = connpool_minsize
            __props__['connpool_replenish'] = connpool_replenish
            __props__['connpool_step'] = connpool_step
            __props__['connpoolidle_timeoutoverride'] = connpoolidle_timeoutoverride
            __props__['defaults_from'] = defaults_from
            __props__['forcehttp10response'] = forcehttp10response
            __props__['idle_timeout'] = idle_timeout
            __props__['maxheader_size'] = maxheader_size
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
        super(ProfileFastHttp, __self__).__init__(
            'f5bigip:ltm/profileFastHttp:ProfileFastHttp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connpool_maxreuse: Optional[pulumi.Input[int]] = None,
            connpool_maxsize: Optional[pulumi.Input[int]] = None,
            connpool_minsize: Optional[pulumi.Input[int]] = None,
            connpool_replenish: Optional[pulumi.Input[str]] = None,
            connpool_step: Optional[pulumi.Input[int]] = None,
            connpoolidle_timeoutoverride: Optional[pulumi.Input[int]] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            forcehttp10response: Optional[pulumi.Input[str]] = None,
            idle_timeout: Optional[pulumi.Input[int]] = None,
            maxheader_size: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ProfileFastHttp':
        """
        Get an existing ProfileFastHttp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] connpool_maxreuse: Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        :param pulumi.Input[int] connpool_maxsize: Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        :param pulumi.Input[int] connpool_minsize: Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        :param pulumi.Input[str] connpool_replenish: The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        :param pulumi.Input[int] connpool_step: Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        :param pulumi.Input[int] connpoolidle_timeoutoverride: Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] forcehttp10response: Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        :param pulumi.Input[int] idle_timeout: Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        :param pulumi.Input[int] maxheader_size: Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        :param pulumi.Input[str] name: Name of the profile_fasthttp
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["connpool_maxreuse"] = connpool_maxreuse
        __props__["connpool_maxsize"] = connpool_maxsize
        __props__["connpool_minsize"] = connpool_minsize
        __props__["connpool_replenish"] = connpool_replenish
        __props__["connpool_step"] = connpool_step
        __props__["connpoolidle_timeoutoverride"] = connpoolidle_timeoutoverride
        __props__["defaults_from"] = defaults_from
        __props__["forcehttp10response"] = forcehttp10response
        __props__["idle_timeout"] = idle_timeout
        __props__["maxheader_size"] = maxheader_size
        __props__["name"] = name
        return ProfileFastHttp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connpoolMaxreuse")
    def connpool_maxreuse(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        """
        return pulumi.get(self, "connpool_maxreuse")

    @property
    @pulumi.getter(name="connpoolMaxsize")
    def connpool_maxsize(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        """
        return pulumi.get(self, "connpool_maxsize")

    @property
    @pulumi.getter(name="connpoolMinsize")
    def connpool_minsize(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        """
        return pulumi.get(self, "connpool_minsize")

    @property
    @pulumi.getter(name="connpoolReplenish")
    def connpool_replenish(self) -> pulumi.Output[Optional[str]]:
        """
        The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        """
        return pulumi.get(self, "connpool_replenish")

    @property
    @pulumi.getter(name="connpoolStep")
    def connpool_step(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        """
        return pulumi.get(self, "connpool_step")

    @property
    @pulumi.getter(name="connpoolidleTimeoutoverride")
    def connpoolidle_timeoutoverride(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        """
        return pulumi.get(self, "connpoolidle_timeoutoverride")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter
    def forcehttp10response(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        """
        return pulumi.get(self, "forcehttp10response")

    @property
    @pulumi.getter(name="idleTimeout")
    def idle_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        """
        return pulumi.get(self, "idle_timeout")

    @property
    @pulumi.getter(name="maxheaderSize")
    def maxheader_size(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        """
        return pulumi.get(self, "maxheader_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the profile_fasthttp
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

