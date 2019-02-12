# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Pool(pulumi.CustomResource):
    allow_nat: pulumi.Output[str]
    allow_snat: pulumi.Output[str]
    load_balancing_mode: pulumi.Output[str]
    monitors: pulumi.Output[list]
    """
    List of monitor names to associate with the pool
    """
    name: pulumi.Output[str]
    """
    Name of the pool
    """
    reselect_tries: pulumi.Output[int]
    service_down_action: pulumi.Output[str]
    slow_ramp_time: pulumi.Output[int]
    def __init__(__self__, resource_name, opts=None, allow_nat=None, allow_snat=None, load_balancing_mode=None, monitors=None, name=None, reselect_tries=None, service_down_action=None, slow_ramp_time=None, __name__=None, __opts__=None):
        """
        `bigip_ltm_pool` Manages a pool configuration.
        
        Resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allow_nat
        :param pulumi.Input[str] allow_snat
        :param pulumi.Input[str] load_balancing_mode
        :param pulumi.Input[list] monitors: List of monitor names to associate with the pool
        :param pulumi.Input[str] name: Name of the pool
        :param pulumi.Input[int] reselect_tries
        :param pulumi.Input[str] service_down_action
        :param pulumi.Input[int] slow_ramp_time
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

        __props__['allow_nat'] = allow_nat

        __props__['allow_snat'] = allow_snat

        __props__['load_balancing_mode'] = load_balancing_mode

        __props__['monitors'] = monitors

        if name is None:
            raise TypeError('Missing required property name')
        __props__['name'] = name

        __props__['reselect_tries'] = reselect_tries

        __props__['service_down_action'] = service_down_action

        __props__['slow_ramp_time'] = slow_ramp_time

        super(Pool, __self__).__init__(
            'f5bigip:ltm/pool:Pool',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

