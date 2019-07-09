# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProfileOneConnect(pulumi.CustomResource):
    defaults_from: pulumi.Output[str]
    """
    Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
    """
    idle_timeout_override: pulumi.Output[str]
    """
    Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
    """
    max_age: pulumi.Output[float]
    """
    Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
    """
    max_reuse: pulumi.Output[float]
    """
    Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
    """
    max_size: pulumi.Output[float]
    """
    Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
    """
    name: pulumi.Output[str]
    """
    Name of the profile_oneconnect
    """
    partition: pulumi.Output[str]
    """
    Displays the administrative partition within which this profile resides
    """
    share_pools: pulumi.Output[str]
    """
    Specify if you want to share the pool, default value is "disabled"
    """
    source_mask: pulumi.Output[str]
    """
    Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
    """
    def __init__(__self__, resource_name, opts=None, defaults_from=None, idle_timeout_override=None, max_age=None, max_reuse=None, max_size=None, name=None, partition=None, share_pools=None, source_mask=None, __name__=None, __opts__=None):
        """
        `bigip_ltm_profile_oneconnect` Configures a custom profile_oneconnect for use by health checks.
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] idle_timeout_override: Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
        :param pulumi.Input[float] max_age: Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
        :param pulumi.Input[float] max_reuse: Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
        :param pulumi.Input[float] max_size: Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
        :param pulumi.Input[str] name: Name of the profile_oneconnect
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] share_pools: Specify if you want to share the pool, default value is "disabled"
        :param pulumi.Input[str] source_mask: Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_oneconnect.html.markdown.
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

        __props__['defaults_from'] = defaults_from

        __props__['idle_timeout_override'] = idle_timeout_override

        __props__['max_age'] = max_age

        __props__['max_reuse'] = max_reuse

        __props__['max_size'] = max_size

        if name is None:
            raise TypeError("Missing required property 'name'")
        __props__['name'] = name

        __props__['partition'] = partition

        __props__['share_pools'] = share_pools

        __props__['source_mask'] = source_mask

        super(ProfileOneConnect, __self__).__init__(
            'f5bigip:ltm/profileOneConnect:ProfileOneConnect',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

