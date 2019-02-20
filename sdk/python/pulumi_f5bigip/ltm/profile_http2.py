# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProfileHttp2(pulumi.CustomResource):
    activation_modes: pulumi.Output[list]
    """
    Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
    """
    concurrent_streams_per_connection: pulumi.Output[int]
    """
    Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
    """
    connection_idle_timeout: pulumi.Output[int]
    """
    Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
    """
    defaults_from: pulumi.Output[str]
    """
    Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
    """
    header_table_size: pulumi.Output[int]
    name: pulumi.Output[str]
    """
    Name of the profile_http2
    """
    def __init__(__self__, resource_name, opts=None, activation_modes=None, concurrent_streams_per_connection=None, connection_idle_timeout=None, defaults_from=None, header_table_size=None, name=None, __name__=None, __opts__=None):
        """
        `bigip_ltm_profile_http2` Configures a custom profile_http2 for use by health checks.
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] activation_modes: Specifies what will cause an incoming connection to be handled as a HTTP/2 connection. The default values npn and alpn specify that the TLS next-protocol-negotiation and application-layer-protocol-negotiation extensions will be used.
        :param pulumi.Input[int] concurrent_streams_per_connection: Specifies how many concurrent requests are allowed to be outstanding on a single HTTP/2 connection.
        :param pulumi.Input[int] connection_idle_timeout: Specifies the number of seconds that a connection is idle before the connection is eligible for deletion..
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[int] header_table_size
        :param pulumi.Input[str] name: Name of the profile_http2
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['activation_modes'] = activation_modes

        __props__['concurrent_streams_per_connection'] = concurrent_streams_per_connection

        __props__['connection_idle_timeout'] = connection_idle_timeout

        __props__['defaults_from'] = defaults_from

        __props__['header_table_size'] = header_table_size

        if name is None:
            raise TypeError('Missing required property name')
        __props__['name'] = name

        super(ProfileHttp2, __self__).__init__(
            'f5bigip:ltm/profileHttp2:ProfileHttp2',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

