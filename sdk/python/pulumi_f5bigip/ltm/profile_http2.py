# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ProfileHttp2(pulumi.CustomResource):
    activation_modes: pulumi.Output[list]
    """
    Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
    """
    concurrent_streams_per_connection: pulumi.Output[float]
    """
    Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
    """
    connection_idle_timeout: pulumi.Output[float]
    """
    Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
    """
    defaults_from: pulumi.Output[str]
    """
    Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
    """
    header_table_size: pulumi.Output[float]
    """
    Use the parent Http2 profile
    """
    name: pulumi.Output[str]
    """
    Name of the profile_http2
    """
    def __init__(__self__, resource_name, opts=None, activation_modes=None, concurrent_streams_per_connection=None, connection_idle_timeout=None, defaults_from=None, header_table_size=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.ProfileHttp2` Configures a custom profile_http2 for use by health checks.

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        nyhttp2 = f5bigip.ltm.ProfileHttp2("nyhttp2",
            activation_modes=[
                "alpn",
                "npn",
            ],
            concurrent_streams_per_connection=10,
            connection_idle_timeout=30,
            defaults_from="/Common/http2",
            name="/Common/NewYork_http2")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] activation_modes: Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
        :param pulumi.Input[float] concurrent_streams_per_connection: Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
        :param pulumi.Input[float] connection_idle_timeout: Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[float] header_table_size: Use the parent Http2 profile
        :param pulumi.Input[str] name: Name of the profile_http2
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

            __props__['activation_modes'] = activation_modes
            __props__['concurrent_streams_per_connection'] = concurrent_streams_per_connection
            __props__['connection_idle_timeout'] = connection_idle_timeout
            __props__['defaults_from'] = defaults_from
            __props__['header_table_size'] = header_table_size
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
        super(ProfileHttp2, __self__).__init__(
            'f5bigip:ltm/profileHttp2:ProfileHttp2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, activation_modes=None, concurrent_streams_per_connection=None, connection_idle_timeout=None, defaults_from=None, header_table_size=None, name=None):
        """
        Get an existing ProfileHttp2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] activation_modes: Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
        :param pulumi.Input[float] concurrent_streams_per_connection: Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
        :param pulumi.Input[float] connection_idle_timeout: Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[float] header_table_size: Use the parent Http2 profile
        :param pulumi.Input[str] name: Name of the profile_http2
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["activation_modes"] = activation_modes
        __props__["concurrent_streams_per_connection"] = concurrent_streams_per_connection
        __props__["connection_idle_timeout"] = connection_idle_timeout
        __props__["defaults_from"] = defaults_from
        __props__["header_table_size"] = header_table_size
        __props__["name"] = name
        return ProfileHttp2(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

