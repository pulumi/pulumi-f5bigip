# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PersistenceProfileSslArgs', 'PersistenceProfileSsl']

@pulumi.input_type
class PersistenceProfileSslArgs:
    def __init__(__self__, *,
                 defaults_from: pulumi.Input[str],
                 name: pulumi.Input[str],
                 app_service: Optional[pulumi.Input[str]] = None,
                 match_across_pools: Optional[pulumi.Input[str]] = None,
                 match_across_services: Optional[pulumi.Input[str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[str]] = None,
                 mirror: Optional[pulumi.Input[str]] = None,
                 override_conn_limit: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a PersistenceProfileSsl resource.
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[int] timeout: Timeout for persistence of the session
        """
        PersistenceProfileSslArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            defaults_from=defaults_from,
            name=name,
            app_service=app_service,
            match_across_pools=match_across_pools,
            match_across_services=match_across_services,
            match_across_virtuals=match_across_virtuals,
            mirror=mirror,
            override_conn_limit=override_conn_limit,
            timeout=timeout,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             defaults_from: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             app_service: Optional[pulumi.Input[str]] = None,
             match_across_pools: Optional[pulumi.Input[str]] = None,
             match_across_services: Optional[pulumi.Input[str]] = None,
             match_across_virtuals: Optional[pulumi.Input[str]] = None,
             mirror: Optional[pulumi.Input[str]] = None,
             override_conn_limit: Optional[pulumi.Input[str]] = None,
             timeout: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if defaults_from is None and 'defaultsFrom' in kwargs:
            defaults_from = kwargs['defaultsFrom']
        if defaults_from is None:
            raise TypeError("Missing 'defaults_from' argument")
        if name is None:
            raise TypeError("Missing 'name' argument")
        if app_service is None and 'appService' in kwargs:
            app_service = kwargs['appService']
        if match_across_pools is None and 'matchAcrossPools' in kwargs:
            match_across_pools = kwargs['matchAcrossPools']
        if match_across_services is None and 'matchAcrossServices' in kwargs:
            match_across_services = kwargs['matchAcrossServices']
        if match_across_virtuals is None and 'matchAcrossVirtuals' in kwargs:
            match_across_virtuals = kwargs['matchAcrossVirtuals']
        if override_conn_limit is None and 'overrideConnLimit' in kwargs:
            override_conn_limit = kwargs['overrideConnLimit']

        _setter("defaults_from", defaults_from)
        _setter("name", name)
        if app_service is not None:
            _setter("app_service", app_service)
        if match_across_pools is not None:
            _setter("match_across_pools", match_across_pools)
        if match_across_services is not None:
            _setter("match_across_services", match_across_services)
        if match_across_virtuals is not None:
            _setter("match_across_virtuals", match_across_virtuals)
        if mirror is not None:
            _setter("mirror", mirror)
        if override_conn_limit is not None:
            _setter("override_conn_limit", override_conn_limit)
        if timeout is not None:
            _setter("timeout", timeout)

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> pulumi.Input[str]:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @defaults_from.setter
    def defaults_from(self, value: pulumi.Input[str]):
        pulumi.set(self, "defaults_from", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "app_service")

    @app_service.setter
    def app_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_service", value)

    @property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @match_across_pools.setter
    def match_across_pools(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_across_pools", value)

    @property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @match_across_services.setter
    def match_across_services(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_across_services", value)

    @property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @match_across_virtuals.setter
    def match_across_virtuals(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_across_virtuals", value)

    @property
    @pulumi.getter
    def mirror(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @mirror.setter
    def mirror(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mirror", value)

    @property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @override_conn_limit.setter
    def override_conn_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_conn_limit", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class _PersistenceProfileSslState:
    def __init__(__self__, *,
                 app_service: Optional[pulumi.Input[str]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 match_across_pools: Optional[pulumi.Input[str]] = None,
                 match_across_services: Optional[pulumi.Input[str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[str]] = None,
                 mirror: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_conn_limit: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering PersistenceProfileSsl resources.
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[int] timeout: Timeout for persistence of the session
        """
        _PersistenceProfileSslState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            app_service=app_service,
            defaults_from=defaults_from,
            match_across_pools=match_across_pools,
            match_across_services=match_across_services,
            match_across_virtuals=match_across_virtuals,
            mirror=mirror,
            name=name,
            override_conn_limit=override_conn_limit,
            timeout=timeout,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             app_service: Optional[pulumi.Input[str]] = None,
             defaults_from: Optional[pulumi.Input[str]] = None,
             match_across_pools: Optional[pulumi.Input[str]] = None,
             match_across_services: Optional[pulumi.Input[str]] = None,
             match_across_virtuals: Optional[pulumi.Input[str]] = None,
             mirror: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             override_conn_limit: Optional[pulumi.Input[str]] = None,
             timeout: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if app_service is None and 'appService' in kwargs:
            app_service = kwargs['appService']
        if defaults_from is None and 'defaultsFrom' in kwargs:
            defaults_from = kwargs['defaultsFrom']
        if match_across_pools is None and 'matchAcrossPools' in kwargs:
            match_across_pools = kwargs['matchAcrossPools']
        if match_across_services is None and 'matchAcrossServices' in kwargs:
            match_across_services = kwargs['matchAcrossServices']
        if match_across_virtuals is None and 'matchAcrossVirtuals' in kwargs:
            match_across_virtuals = kwargs['matchAcrossVirtuals']
        if override_conn_limit is None and 'overrideConnLimit' in kwargs:
            override_conn_limit = kwargs['overrideConnLimit']

        if app_service is not None:
            _setter("app_service", app_service)
        if defaults_from is not None:
            _setter("defaults_from", defaults_from)
        if match_across_pools is not None:
            _setter("match_across_pools", match_across_pools)
        if match_across_services is not None:
            _setter("match_across_services", match_across_services)
        if match_across_virtuals is not None:
            _setter("match_across_virtuals", match_across_virtuals)
        if mirror is not None:
            _setter("mirror", mirror)
        if name is not None:
            _setter("name", name)
        if override_conn_limit is not None:
            _setter("override_conn_limit", override_conn_limit)
        if timeout is not None:
            _setter("timeout", timeout)

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "app_service")

    @app_service.setter
    def app_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_service", value)

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> Optional[pulumi.Input[str]]:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @defaults_from.setter
    def defaults_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "defaults_from", value)

    @property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @match_across_pools.setter
    def match_across_pools(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_across_pools", value)

    @property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @match_across_services.setter
    def match_across_services(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_across_services", value)

    @property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @match_across_virtuals.setter
    def match_across_virtuals(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match_across_virtuals", value)

    @property
    @pulumi.getter
    def mirror(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @mirror.setter
    def mirror(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mirror", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> Optional[pulumi.Input[str]]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @override_conn_limit.setter
    def override_conn_limit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "override_conn_limit", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


class PersistenceProfileSsl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service: Optional[pulumi.Input[str]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 match_across_pools: Optional[pulumi.Input[str]] = None,
                 match_across_services: Optional[pulumi.Input[str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[str]] = None,
                 mirror: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_conn_limit: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Configures an SSL persistence profile

        ## Example

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        ppssl = f5bigip.ltm.PersistenceProfileSsl("ppssl",
            defaults_from="/Common/ssl",
            match_across_pools="enabled",
            match_across_services="enabled",
            match_across_virtuals="enabled",
            mirror="enabled",
            name="/Common/terraform_ssl",
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

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[int] timeout: Timeout for persistence of the session
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PersistenceProfileSslArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configures an SSL persistence profile

        ## Example

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        ppssl = f5bigip.ltm.PersistenceProfileSsl("ppssl",
            defaults_from="/Common/ssl",
            match_across_pools="enabled",
            match_across_services="enabled",
            match_across_virtuals="enabled",
            mirror="enabled",
            name="/Common/terraform_ssl",
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

        :param str resource_name: The name of the resource.
        :param PersistenceProfileSslArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PersistenceProfileSslArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            PersistenceProfileSslArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_service: Optional[pulumi.Input[str]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 match_across_pools: Optional[pulumi.Input[str]] = None,
                 match_across_services: Optional[pulumi.Input[str]] = None,
                 match_across_virtuals: Optional[pulumi.Input[str]] = None,
                 mirror: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 override_conn_limit: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PersistenceProfileSslArgs.__new__(PersistenceProfileSslArgs)

            __props__.__dict__["app_service"] = app_service
            if defaults_from is None and not opts.urn:
                raise TypeError("Missing required property 'defaults_from'")
            __props__.__dict__["defaults_from"] = defaults_from
            __props__.__dict__["match_across_pools"] = match_across_pools
            __props__.__dict__["match_across_services"] = match_across_services
            __props__.__dict__["match_across_virtuals"] = match_across_virtuals
            __props__.__dict__["mirror"] = mirror
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["override_conn_limit"] = override_conn_limit
            __props__.__dict__["timeout"] = timeout
        super(PersistenceProfileSsl, __self__).__init__(
            'f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_service: Optional[pulumi.Input[str]] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            match_across_pools: Optional[pulumi.Input[str]] = None,
            match_across_services: Optional[pulumi.Input[str]] = None,
            match_across_virtuals: Optional[pulumi.Input[str]] = None,
            mirror: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            override_conn_limit: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None) -> 'PersistenceProfileSsl':
        """
        Get an existing PersistenceProfileSsl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defaults_from: Inherit defaults from parent profile
        :param pulumi.Input[str] match_across_pools: To enable _ disable match across pools with given persistence record
        :param pulumi.Input[str] match_across_services: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] match_across_virtuals: To enable _ disable match across services with given persistence record
        :param pulumi.Input[str] mirror: To enable _ disable
        :param pulumi.Input[str] name: Name of the persistence profile
        :param pulumi.Input[str] override_conn_limit: To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
               limits remain hard limits and are not overridden.
        :param pulumi.Input[int] timeout: Timeout for persistence of the session
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PersistenceProfileSslState.__new__(_PersistenceProfileSslState)

        __props__.__dict__["app_service"] = app_service
        __props__.__dict__["defaults_from"] = defaults_from
        __props__.__dict__["match_across_pools"] = match_across_pools
        __props__.__dict__["match_across_services"] = match_across_services
        __props__.__dict__["match_across_virtuals"] = match_across_virtuals
        __props__.__dict__["mirror"] = mirror
        __props__.__dict__["name"] = name
        __props__.__dict__["override_conn_limit"] = override_conn_limit
        __props__.__dict__["timeout"] = timeout
        return PersistenceProfileSsl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "app_service")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> pulumi.Output[str]:
        """
        Inherit defaults from parent profile
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter(name="matchAcrossPools")
    def match_across_pools(self) -> pulumi.Output[str]:
        """
        To enable _ disable match across pools with given persistence record
        """
        return pulumi.get(self, "match_across_pools")

    @property
    @pulumi.getter(name="matchAcrossServices")
    def match_across_services(self) -> pulumi.Output[str]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_services")

    @property
    @pulumi.getter(name="matchAcrossVirtuals")
    def match_across_virtuals(self) -> pulumi.Output[str]:
        """
        To enable _ disable match across services with given persistence record
        """
        return pulumi.get(self, "match_across_virtuals")

    @property
    @pulumi.getter
    def mirror(self) -> pulumi.Output[str]:
        """
        To enable _ disable
        """
        return pulumi.get(self, "mirror")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the persistence profile
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overrideConnLimit")
    def override_conn_limit(self) -> pulumi.Output[str]:
        """
        To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
        limits remain hard limits and are not overridden.
        """
        return pulumi.get(self, "override_conn_limit")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        Timeout for persistence of the session
        """
        return pulumi.get(self, "timeout")

