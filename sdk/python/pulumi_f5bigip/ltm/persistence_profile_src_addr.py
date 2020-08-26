# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['PersistenceProfileSrcAddr']


class PersistenceProfileSrcAddr(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service: Optional[pulumi.Input[str]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 hash_algorithm: Optional[pulumi.Input[str]] = None,
                 map_proxies: Optional[pulumi.Input[str]] = None,
                 mask: Optional[pulumi.Input[str]] = None,
                 match_across_pools: Optional[pulumi.Input[str]] = None,
                 match_across_services: Optional[pulumi.Input[str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[str]] = None,
                 mirror: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_conn_limit: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configures a source address persistence profile

        ## Example

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        srcaddr = f5bigip.ltm.PersistenceProfileSrcAddr("srcaddr",
            defaults_from="/Common/source_addr",
            hash_algorithm="carp",
            map_proxies="enabled",
            mask="255.255.255.255",
            match_across_pools="enabled",
            match_across_services="enabled",
            match_across_virtuals="enabled",
            mirror="enabled",
            name="/Common/terraform_srcaddr",
            override_conn_limit="enabled",
            timeout=3600)
        ```

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
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] hash_algorithm: Specify the hash algorithm
        :param pulumi.Input[str] map_proxies: To enable _ disable directs all to the same single pool member
        :param pulumi.Input[str] mask: Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
               when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[float] timeout: Timeout for persistence of the session
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service: Optional[pulumi.Input[str]] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            hash_algorithm: Optional[pulumi.Input[str]] = None,
            map_proxies: Optional[pulumi.Input[str]] = None,
            mask: Optional[pulumi.Input[str]] = None,
            match_across_pools: Optional[pulumi.Input[str]] = None,
            match_across_services: Optional[pulumi.Input[str]] = None,
            match_across_virtuals: Optional[pulumi.Input[str]] = None,
            mirror: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            override_conn_limit: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[float]] = None) -> 'PersistenceProfileSrcAddr':
        """
        Get an existing PersistenceProfileSrcAddr resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] hash_algorithm: Specify the hash algorithm
        :param pulumi.Input[str] map_proxies: To enable _ disable directs all to the same single pool member
        :param pulumi.Input[str] mask: Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
               when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[float] timeout: Timeout for persistence of the session
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

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[str]:
        return pulumi.get(self, "app_service")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> str:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> Optional[str]:
        """
        Specify the hash algorithm
        """
        return pulumi.get(self, "hash_algorithm")

    @property
    @pulumi.getter(name="mapProxies")
    def map_proxies(self) -> Optional[str]:
        """
        To enable _ disable directs all to the same single pool member
        """
        return pulumi.get(self, "map_proxies")

    @property
    @pulumi.getter
    def mask(self) -> Optional[str]:
        """
        Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
        when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        """
        return pulumi.get(self, "mask")

    @property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> Optional[str]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> Optional[str]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> Optional[str]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @property
    @pulumi.getter
    def mirror(self) -> Optional[str]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> Optional[str]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[float]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

