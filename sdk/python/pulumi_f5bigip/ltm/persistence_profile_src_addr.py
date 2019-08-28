# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PersistenceProfileSrcAddr(pulumi.CustomResource):
    app_service: pulumi.Output[str]
    defaults_from: pulumi.Output[str]
    hash_algorithm: pulumi.Output[str]
    map_proxies: pulumi.Output[str]
    mask: pulumi.Output[str]
    match_across_pools: pulumi.Output[str]
    match_across_services: pulumi.Output[str]
    match_across_virtuals: pulumi.Output[str]
    mirror: pulumi.Output[str]
    name: pulumi.Output[str]
    override_conn_limit: pulumi.Output[str]
    timeout: pulumi.Output[float]
    def __init__(__self__, resource_name, opts=None, app_service=None, defaults_from=None, hash_algorithm=None, map_proxies=None, mask=None, match_across_pools=None, match_across_services=None, match_across_virtuals=None, mirror=None, name=None, override_conn_limit=None, timeout=None, __props__=None, __name__=None, __opts__=None):
        """
        Configures a source address persistence profile
        
        ## Reference
        
        `name` - (Required) Name of the virtual address
        
        `defaults_from` - (Required) Parent cookie persistence profile
        
        `match_across_pools` (Optional) (enabled or disabled) match across pools with given persistence record
        
        `match_across_services` (Optional) (enabled or disabled) match across services with given persistence record
        
        `match_across_virtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
        
        `mirror` (Optional) (enabled or disabled) mirror persistence record
        
        `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
        
        `override_conn_limit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
        
        `hash_algorithm` (Optional) Specify the hash algorithm
        
        `mask` (Optional) Identify a range of source IP addresses to manage together as a single source address affinity persistent connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        
        `map_proxies` (Optional) (enabled or disabled) Directs all to the same single pool member
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_persistence_profile_srcaddr.html.markdown.
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

            __props__['app_service'] = app_service
            if defaults_from is None:
                raise TypeError("Missing required property 'defaults_from'")
            __props__['defaults_from'] = defaults_from
            __props__['hash_algorithm'] = hash_algorithm
            __props__['map_proxies'] = map_proxies
            __props__['mask'] = mask
            __props__['match_across_pools'] = match_across_pools
            __props__['match_across_services'] = match_across_services
            __props__['match_across_virtuals'] = match_across_virtuals
            __props__['mirror'] = mirror
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['override_conn_limit'] = override_conn_limit
            __props__['timeout'] = timeout
        super(PersistenceProfileSrcAddr, __self__).__init__(
            'f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_service=None, defaults_from=None, hash_algorithm=None, map_proxies=None, mask=None, match_across_pools=None, match_across_services=None, match_across_virtuals=None, mirror=None, name=None, override_conn_limit=None, timeout=None):
        """
        Get an existing PersistenceProfileSrcAddr resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_persistence_profile_srcaddr.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["app_service"] = app_service
        __props__["defaults_from"] = defaults_from
        __props__["hash_algorithm"] = hash_algorithm
        __props__["map_proxies"] = map_proxies
        __props__["mask"] = mask
        __props__["match_across_pools"] = match_across_pools
        __props__["match_across_services"] = match_across_services
        __props__["match_across_virtuals"] = match_across_virtuals
        __props__["mirror"] = mirror
        __props__["name"] = name
        __props__["override_conn_limit"] = override_conn_limit
        __props__["timeout"] = timeout
        return PersistenceProfileSrcAddr(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

