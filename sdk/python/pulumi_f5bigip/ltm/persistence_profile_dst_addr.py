# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['PersistenceProfileDstAddrArgs', 'PersistenceProfileDstAddr']

@pulumi.input_type
class PersistenceProfileDstAddrArgs:
    def __init__(__self__, *,
                 defaults_from: pulumi.Input[_builtins.str],
                 name: pulumi.Input[_builtins.str],
                 app_service: Optional[pulumi.Input[_builtins.str]] = None,
                 hash_algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 mask: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_pools: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_services: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 override_conn_limit: Optional[pulumi.Input[_builtins.str]] = None,
                 timeout: Optional[pulumi.Input[_builtins.int]] = None):
        """
        The set of arguments for constructing a PersistenceProfileDstAddr resource.
        :param pulumi.Input[_builtins.str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[_builtins.str] name: Name of the persistence profile
        :param pulumi.Input[_builtins.str] hash_algorithm: Specify the hash algorithm
        :param pulumi.Input[_builtins.str] mask: Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
               when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        :param pulumi.Input[_builtins.str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[_builtins.str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] mirror: To enable _ disable
        :param pulumi.Input[_builtins.str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[_builtins.int] timeout: Timeout for persistence of the session
        """
        pulumi.set(__self__, "defaults_from", defaults_from)
        pulumi.set(__self__, "name", name)
        if app_service is not None:
            pulumi.set(__self__, "app_service", app_service)
        if hash_algorithm is not None:
            pulumi.set(__self__, "hash_algorithm", hash_algorithm)
        if mask is not None:
            pulumi.set(__self__, "mask", mask)
        if match_across_pools is not None:
            pulumi.set(__self__, "match_across_pools", match_across_pools)
        if match_across_services is not None:
            pulumi.set(__self__, "match_across_services", match_across_services)
        if match_across_virtuals is not None:
            pulumi.set(__self__, "match_across_virtuals", match_across_virtuals)
        if mirror is not None:
            pulumi.set(__self__, "mirror", mirror)
        if override_conn_limit is not None:
            pulumi.set(__self__, "override_conn_limit", override_conn_limit)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @_builtins.property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> pulumi.Input[_builtins.str]:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @defaults_from.setter
    def defaults_from(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "defaults_from", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "app_service")

    @app_service.setter
    def app_service(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "app_service", value)

    @_builtins.property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specify the hash algorithm
        """
        return pulumi.get(self, "hash_algorithm")

    @hash_algorithm.setter
    def hash_algorithm(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "hash_algorithm", value)

    @_builtins.property
    @pulumi.getter
    def mask(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
        when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        """
        return pulumi.get(self, "mask")

    @mask.setter
    def mask(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mask", value)

    @_builtins.property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @match_across_pools.setter
    def match_across_pools(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "match_across_pools", value)

    @_builtins.property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @match_across_services.setter
    def match_across_services(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "match_across_services", value)

    @_builtins.property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @match_across_virtuals.setter
    def match_across_virtuals(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "match_across_virtuals", value)

    @_builtins.property
    @pulumi.getter
    def mirror(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @mirror.setter
    def mirror(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mirror", value)

    @_builtins.property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @override_conn_limit.setter
    def override_conn_limit(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "override_conn_limit", value)

    @_builtins.property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _PersistenceProfileDstAddrState:
    def __init__(__self__, *,
                 app_service: Optional[pulumi.Input[_builtins.str]] = None,
                 defaults_from: Optional[pulumi.Input[_builtins.str]] = None,
                 hash_algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 mask: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_pools: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_services: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 override_conn_limit: Optional[pulumi.Input[_builtins.str]] = None,
                 timeout: Optional[pulumi.Input[_builtins.int]] = None):
        """
        Input properties used for looking up and filtering PersistenceProfileDstAddr resources.
        :param pulumi.Input[_builtins.str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[_builtins.str] hash_algorithm: Specify the hash algorithm
        :param pulumi.Input[_builtins.str] mask: Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
               when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        :param pulumi.Input[_builtins.str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[_builtins.str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] mirror: To enable _ disable
        :param pulumi.Input[_builtins.str] name: Name of the persistence profile
        :param pulumi.Input[_builtins.str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[_builtins.int] timeout: Timeout for persistence of the session
        """
        if app_service is not None:
            pulumi.set(__self__, "app_service", app_service)
        if defaults_from is not None:
            pulumi.set(__self__, "defaults_from", defaults_from)
        if hash_algorithm is not None:
            pulumi.set(__self__, "hash_algorithm", hash_algorithm)
        if mask is not None:
            pulumi.set(__self__, "mask", mask)
        if match_across_pools is not None:
            pulumi.set(__self__, "match_across_pools", match_across_pools)
        if match_across_services is not None:
            pulumi.set(__self__, "match_across_services", match_across_services)
        if match_across_virtuals is not None:
            pulumi.set(__self__, "match_across_virtuals", match_across_virtuals)
        if mirror is not None:
            pulumi.set(__self__, "mirror", mirror)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if override_conn_limit is not None:
            pulumi.set(__self__, "override_conn_limit", override_conn_limit)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @_builtins.property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "app_service")

    @app_service.setter
    def app_service(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "app_service", value)

    @_builtins.property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @defaults_from.setter
    def defaults_from(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "defaults_from", value)

    @_builtins.property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specify the hash algorithm
        """
        return pulumi.get(self, "hash_algorithm")

    @hash_algorithm.setter
    def hash_algorithm(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "hash_algorithm", value)

    @_builtins.property
    @pulumi.getter
    def mask(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
        when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        """
        return pulumi.get(self, "mask")

    @mask.setter
    def mask(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mask", value)

    @_builtins.property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @match_across_pools.setter
    def match_across_pools(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "match_across_pools", value)

    @_builtins.property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @match_across_services.setter
    def match_across_services(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "match_across_services", value)

    @_builtins.property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @match_across_virtuals.setter
    def match_across_virtuals(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "match_across_virtuals", value)

    @_builtins.property
    @pulumi.getter
    def mirror(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @mirror.setter
    def mirror(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mirror", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @override_conn_limit.setter
    def override_conn_limit(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "override_conn_limit", value)

    @_builtins.property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "timeout", value)


@pulumi.type_token("f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr")
class PersistenceProfileDstAddr(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service: Optional[pulumi.Input[_builtins.str]] = None,
                 defaults_from: Optional[pulumi.Input[_builtins.str]] = None,
                 hash_algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 mask: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_pools: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_services: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 override_conn_limit: Optional[pulumi.Input[_builtins.str]] = None,
                 timeout: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        """
        Configures a cookie persistence profile

        ## Example

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        dstaddr = f5bigip.ltm.PersistenceProfileDstAddr("dstaddr",
            name="/Common/terraform_ppdstaddr",
            defaults_from="/Common/dest_addr",
            match_across_pools="enabled",
            match_across_services="enabled",
            match_across_virtuals="enabled",
            mirror="enabled",
            timeout=3600,
            override_conn_limit="enabled",
            hash_algorithm="carp",
            mask="255.255.255.255")
        ```

        ## Reference

        `name` - (Required) Name of the virtual address

        `defaults_from` - (Optional) Specifies the existing profile from which the system imports settings for the new profile.

        `match_across_pools` (Optional) (enabled or disabled) match across pools with given persistence record

        `match_across_services` (Optional) (enabled or disabled) match across services with given persistence record

        `match_across_virtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record

        `mirror` (Optional) (enabled or disabled) mirror persistence record

        `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds

        `override_conn_limit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.

        ## Importing

        An dest-addr persistence profile can be imported into this resource by supplying the Name in `full path` as `id`.
        An example is below:
        ```sh
        $ terraform import bigip_ltm_persistence_profile_dstaddr.dstaddr "/Common/terraform_ppdstaddr"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[_builtins.str] hash_algorithm: Specify the hash algorithm
        :param pulumi.Input[_builtins.str] mask: Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
               when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        :param pulumi.Input[_builtins.str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[_builtins.str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] mirror: To enable _ disable
        :param pulumi.Input[_builtins.str] name: Name of the persistence profile
        :param pulumi.Input[_builtins.str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[_builtins.int] timeout: Timeout for persistence of the session
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PersistenceProfileDstAddrArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configures a cookie persistence profile

        ## Example

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        dstaddr = f5bigip.ltm.PersistenceProfileDstAddr("dstaddr",
            name="/Common/terraform_ppdstaddr",
            defaults_from="/Common/dest_addr",
            match_across_pools="enabled",
            match_across_services="enabled",
            match_across_virtuals="enabled",
            mirror="enabled",
            timeout=3600,
            override_conn_limit="enabled",
            hash_algorithm="carp",
            mask="255.255.255.255")
        ```

        ## Reference

        `name` - (Required) Name of the virtual address

        `defaults_from` - (Optional) Specifies the existing profile from which the system imports settings for the new profile.

        `match_across_pools` (Optional) (enabled or disabled) match across pools with given persistence record

        `match_across_services` (Optional) (enabled or disabled) match across services with given persistence record

        `match_across_virtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record

        `mirror` (Optional) (enabled or disabled) mirror persistence record

        `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds

        `override_conn_limit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.

        ## Importing

        An dest-addr persistence profile can be imported into this resource by supplying the Name in `full path` as `id`.
        An example is below:
        ```sh
        $ terraform import bigip_ltm_persistence_profile_dstaddr.dstaddr "/Common/terraform_ppdstaddr"
        ```

        :param str resource_name: The name of the resource.
        :param PersistenceProfileDstAddrArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PersistenceProfileDstAddrArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service: Optional[pulumi.Input[_builtins.str]] = None,
                 defaults_from: Optional[pulumi.Input[_builtins.str]] = None,
                 hash_algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 mask: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_pools: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_services: Optional[pulumi.Input[_builtins.str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 override_conn_limit: Optional[pulumi.Input[_builtins.str]] = None,
                 timeout: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PersistenceProfileDstAddrArgs.__new__(PersistenceProfileDstAddrArgs)

            __props__.__dict__["app_service"] = app_service
            if defaults_from is None and not opts.urn:
                raise TypeError("Missing required property 'defaults_from'")
            __props__.__dict__["defaults_from"] = defaults_from
            __props__.__dict__["hash_algorithm"] = hash_algorithm
            __props__.__dict__["mask"] = mask
            __props__.__dict__["match_across_pools"] = match_across_pools
            __props__.__dict__["match_across_services"] = match_across_services
            __props__.__dict__["match_across_virtuals"] = match_across_virtuals
            __props__.__dict__["mirror"] = mirror
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["override_conn_limit"] = override_conn_limit
            __props__.__dict__["timeout"] = timeout
        super(PersistenceProfileDstAddr, __self__).__init__(
            'f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service: Optional[pulumi.Input[_builtins.str]] = None,
            defaults_from: Optional[pulumi.Input[_builtins.str]] = None,
            hash_algorithm: Optional[pulumi.Input[_builtins.str]] = None,
            mask: Optional[pulumi.Input[_builtins.str]] = None,
            match_across_pools: Optional[pulumi.Input[_builtins.str]] = None,
            match_across_services: Optional[pulumi.Input[_builtins.str]] = None,
            match_across_virtuals: Optional[pulumi.Input[_builtins.str]] = None,
            mirror: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            override_conn_limit: Optional[pulumi.Input[_builtins.str]] = None,
            timeout: Optional[pulumi.Input[_builtins.int]] = None) -> 'PersistenceProfileDstAddr':
        """
        Get an existing PersistenceProfileDstAddr resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[_builtins.str] hash_algorithm: Specify the hash algorithm
        :param pulumi.Input[_builtins.str] mask: Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
               when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        :param pulumi.Input[_builtins.str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[_builtins.str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[_builtins.str] mirror: To enable _ disable
        :param pulumi.Input[_builtins.str] name: Name of the persistence profile
        :param pulumi.Input[_builtins.str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[_builtins.int] timeout: Timeout for persistence of the session
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PersistenceProfileDstAddrState.__new__(_PersistenceProfileDstAddrState)

        __props__.__dict__["app_service"] = app_service
        __props__.__dict__["defaults_from"] = defaults_from
        __props__.__dict__["hash_algorithm"] = hash_algorithm
        __props__.__dict__["mask"] = mask
        __props__.__dict__["match_across_pools"] = match_across_pools
        __props__.__dict__["match_across_services"] = match_across_services
        __props__.__dict__["match_across_virtuals"] = match_across_virtuals
        __props__.__dict__["mirror"] = mirror
        __props__.__dict__["name"] = name
        __props__.__dict__["override_conn_limit"] = override_conn_limit
        __props__.__dict__["timeout"] = timeout
        return PersistenceProfileDstAddr(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="appService")
    def app_service(self) -> pulumi.Output[_builtins.str]:
        return pulumi.get(self, "app_service")

    @_builtins.property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> pulumi.Output[_builtins.str]:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @_builtins.property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> pulumi.Output[_builtins.str]:
        """
        Specify the hash algorithm
        """
        return pulumi.get(self, "hash_algorithm")

    @_builtins.property
    @pulumi.getter
    def mask(self) -> pulumi.Output[_builtins.str]:
        """
        Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
        when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
        """
        return pulumi.get(self, "mask")

    @_builtins.property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> pulumi.Output[_builtins.str]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @_builtins.property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> pulumi.Output[_builtins.str]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @_builtins.property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> pulumi.Output[_builtins.str]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @_builtins.property
    @pulumi.getter
    def mirror(self) -> pulumi.Output[_builtins.str]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> pulumi.Output[_builtins.str]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @_builtins.property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[_builtins.int]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

