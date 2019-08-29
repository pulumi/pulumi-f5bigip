# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ProfileFastHttp(pulumi.CustomResource):
    connpool_maxreuse: pulumi.Output[float]
    """
    Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
    """
    connpool_maxsize: pulumi.Output[float]
    """
    Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
    """
    connpool_minsize: pulumi.Output[float]
    """
    Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
    """
    connpool_replenish: pulumi.Output[str]
    """
    The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
    """
    connpool_step: pulumi.Output[float]
    """
    Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
    """
    connpoolidle_timeoutoverride: pulumi.Output[float]
    """
    Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
    """
    defaults_from: pulumi.Output[str]
    """
    Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
    """
    forcehttp10response: pulumi.Output[str]
    """
    Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
    """
    idle_timeout: pulumi.Output[float]
    """
    Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
    """
    maxheader_size: pulumi.Output[float]
    """
    Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
    """
    name: pulumi.Output[str]
    """
    Name of the profile_fasthttp
    """
    def __init__(__self__, resource_name, opts=None, connpool_maxreuse=None, connpool_maxsize=None, connpool_minsize=None, connpool_replenish=None, connpool_step=None, connpoolidle_timeoutoverride=None, defaults_from=None, forcehttp10response=None, idle_timeout=None, maxheader_size=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.ProfileFastHttp` Configures a custom profile_fasthttp for use by health checks.
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] connpool_maxreuse: Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        :param pulumi.Input[float] connpool_maxsize: Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        :param pulumi.Input[float] connpool_minsize: Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        :param pulumi.Input[str] connpool_replenish: The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        :param pulumi.Input[float] connpool_step: Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        :param pulumi.Input[float] connpoolidle_timeoutoverride: Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] forcehttp10response: Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        :param pulumi.Input[float] idle_timeout: Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        :param pulumi.Input[float] maxheader_size: Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        :param pulumi.Input[str] name: Name of the profile_fasthttp

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_fasthttp.html.markdown.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, connpool_maxreuse=None, connpool_maxsize=None, connpool_minsize=None, connpool_replenish=None, connpool_step=None, connpoolidle_timeoutoverride=None, defaults_from=None, forcehttp10response=None, idle_timeout=None, maxheader_size=None, name=None):
        """
        Get an existing ProfileFastHttp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] connpool_maxreuse: Specifies the maximum number of times that the system can re-use a current connection. The default value is 0 (zero).
        :param pulumi.Input[float] connpool_maxsize: Specifies the maximum number of connections to a load balancing pool. A setting of 0 specifies that a pool can accept an unlimited number of connections. The default value is 2048.
        :param pulumi.Input[float] connpool_minsize: Specifies the minimum number of connections to a load balancing pool. A setting of 0 specifies that there is no minimum. The default value is 10.
        :param pulumi.Input[str] connpool_replenish: The default value is enabled. When this option is enabled, the system replenishes the number of connections to a load balancing pool to the number of connections that existed when the server closed the connection to the pool. When disabled, the system replenishes the connection that was closed by the server, only when there are fewer connections to the pool than the number of connections set in the connpool-min-size connections option. Also see the connpool-min-size option..
        :param pulumi.Input[float] connpool_step: Specifies the increment in which the system makes additional connections available, when all available connections are in use. The default value is 4.
        :param pulumi.Input[float] connpoolidle_timeoutoverride: Specifies the number of seconds after which a server-side connection in a OneConnect pool is eligible for deletion, when the connection has no traffic.The value of this option overrides the idle-timeout value that you specify. The default value is 0 (zero) seconds, which disables the override setting.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] forcehttp10response: Specifies whether to rewrite the HTTP version in the status line of the server to HTTP 1.0 to discourage the client from pipelining or chunking data. The default value is disabled.
        :param pulumi.Input[float] idle_timeout: Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        :param pulumi.Input[float] maxheader_size: Specifies the maximum amount of HTTP header data that the system buffers before making a load balancing decision. The default setting is 32768.
        :param pulumi.Input[str] name: Name of the profile_fasthttp

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_fasthttp.html.markdown.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

