# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProfileOneConnectArgs', 'ProfileOneConnect']

@pulumi.input_type
class ProfileOneConnectArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 idle_timeout_override: Optional[pulumi.Input[str]] = None,
                 limit_type: Optional[pulumi.Input[str]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_reuse: Optional[pulumi.Input[int]] = None,
                 max_size: Optional[pulumi.Input[int]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 share_pools: Optional[pulumi.Input[str]] = None,
                 source_mask: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProfileOneConnect resource.
        :param pulumi.Input[str] name: Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] idle_timeout_override: Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
        :param pulumi.Input[str] limit_type: Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
        :param pulumi.Input[int] max_age: Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
        :param pulumi.Input[int] max_reuse: Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
        :param pulumi.Input[int] max_size: Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] share_pools: Specify if you want to share the pool, default value is `disabled`.
        :param pulumi.Input[str] source_mask: Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        """
        pulumi.set(__self__, "name", name)
        if defaults_from is not None:
            pulumi.set(__self__, "defaults_from", defaults_from)
        if idle_timeout_override is not None:
            pulumi.set(__self__, "idle_timeout_override", idle_timeout_override)
        if limit_type is not None:
            pulumi.set(__self__, "limit_type", limit_type)
        if max_age is not None:
            pulumi.set(__self__, "max_age", max_age)
        if max_reuse is not None:
            pulumi.set(__self__, "max_reuse", max_reuse)
        if max_size is not None:
            pulumi.set(__self__, "max_size", max_size)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if share_pools is not None:
            pulumi.set(__self__, "share_pools", share_pools)
        if source_mask is not None:
            pulumi.set(__self__, "source_mask", source_mask)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        """
        return pulumi.get(self, "defaults_from")

    @defaults_from.setter
    def defaults_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "defaults_from", value)

    @property
    @pulumi.getter(name="idleTimeoutOverride")
    def idle_timeout_override(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
        """
        return pulumi.get(self, "idle_timeout_override")

    @idle_timeout_override.setter
    def idle_timeout_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idle_timeout_override", value)

    @property
    @pulumi.getter(name="limitType")
    def limit_type(self) -> Optional[pulumi.Input[str]]:
        """
        Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
        """
        return pulumi.get(self, "limit_type")

    @limit_type.setter
    def limit_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_type", value)

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
        """
        return pulumi.get(self, "max_age")

    @max_age.setter
    def max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age", value)

    @property
    @pulumi.getter(name="maxReuse")
    def max_reuse(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
        """
        return pulumi.get(self, "max_reuse")

    @max_reuse.setter
    def max_reuse(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_reuse", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
        """
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Displays the administrative partition within which this profile resides
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter(name="sharePools")
    def share_pools(self) -> Optional[pulumi.Input[str]]:
        """
        Specify if you want to share the pool, default value is `disabled`.
        """
        return pulumi.get(self, "share_pools")

    @share_pools.setter
    def share_pools(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_pools", value)

    @property
    @pulumi.getter(name="sourceMask")
    def source_mask(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        """
        return pulumi.get(self, "source_mask")

    @source_mask.setter
    def source_mask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_mask", value)


@pulumi.input_type
class _ProfileOneConnectState:
    def __init__(__self__, *,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 idle_timeout_override: Optional[pulumi.Input[str]] = None,
                 limit_type: Optional[pulumi.Input[str]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_reuse: Optional[pulumi.Input[int]] = None,
                 max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 share_pools: Optional[pulumi.Input[str]] = None,
                 source_mask: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProfileOneConnect resources.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] idle_timeout_override: Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
        :param pulumi.Input[str] limit_type: Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
        :param pulumi.Input[int] max_age: Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
        :param pulumi.Input[int] max_reuse: Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
        :param pulumi.Input[int] max_size: Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
        :param pulumi.Input[str] name: Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] share_pools: Specify if you want to share the pool, default value is `disabled`.
        :param pulumi.Input[str] source_mask: Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        """
        if defaults_from is not None:
            pulumi.set(__self__, "defaults_from", defaults_from)
        if idle_timeout_override is not None:
            pulumi.set(__self__, "idle_timeout_override", idle_timeout_override)
        if limit_type is not None:
            pulumi.set(__self__, "limit_type", limit_type)
        if max_age is not None:
            pulumi.set(__self__, "max_age", max_age)
        if max_reuse is not None:
            pulumi.set(__self__, "max_reuse", max_reuse)
        if max_size is not None:
            pulumi.set(__self__, "max_size", max_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if share_pools is not None:
            pulumi.set(__self__, "share_pools", share_pools)
        if source_mask is not None:
            pulumi.set(__self__, "source_mask", source_mask)

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        """
        return pulumi.get(self, "defaults_from")

    @defaults_from.setter
    def defaults_from(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "defaults_from", value)

    @property
    @pulumi.getter(name="idleTimeoutOverride")
    def idle_timeout_override(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
        """
        return pulumi.get(self, "idle_timeout_override")

    @idle_timeout_override.setter
    def idle_timeout_override(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idle_timeout_override", value)

    @property
    @pulumi.getter(name="limitType")
    def limit_type(self) -> Optional[pulumi.Input[str]]:
        """
        Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
        """
        return pulumi.get(self, "limit_type")

    @limit_type.setter
    def limit_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "limit_type", value)

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
        """
        return pulumi.get(self, "max_age")

    @max_age.setter
    def max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_age", value)

    @property
    @pulumi.getter(name="maxReuse")
    def max_reuse(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
        """
        return pulumi.get(self, "max_reuse")

    @max_reuse.setter
    def max_reuse(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_reuse", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
        """
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Displays the administrative partition within which this profile resides
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter(name="sharePools")
    def share_pools(self) -> Optional[pulumi.Input[str]]:
        """
        Specify if you want to share the pool, default value is `disabled`.
        """
        return pulumi.get(self, "share_pools")

    @share_pools.setter
    def share_pools(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "share_pools", value)

    @property
    @pulumi.getter(name="sourceMask")
    def source_mask(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        """
        return pulumi.get(self, "source_mask")

    @source_mask.setter
    def source_mask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_mask", value)


class ProfileOneConnect(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 idle_timeout_override: Optional[pulumi.Input[str]] = None,
                 limit_type: Optional[pulumi.Input[str]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_reuse: Optional[pulumi.Input[int]] = None,
                 max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 share_pools: Optional[pulumi.Input[str]] = None,
                 source_mask: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `ltm.ProfileOneConnect` Configures a custom profile_oneconnect for use by health checks.

        Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_oneconnect = f5bigip.ltm.ProfileOneConnect("test-oneconnect", name="/Common/test-oneconnect")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        BIG-IP LTM oneconnect profiles can be imported using the `name` , e.g.

        ```sh
        $ pulumi import f5bigip:ltm/profileOneConnect:ProfileOneConnect test-oneconnect /Common/test-oneconnect
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] idle_timeout_override: Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
        :param pulumi.Input[str] limit_type: Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
        :param pulumi.Input[int] max_age: Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
        :param pulumi.Input[int] max_reuse: Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
        :param pulumi.Input[int] max_size: Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
        :param pulumi.Input[str] name: Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] share_pools: Specify if you want to share the pool, default value is `disabled`.
        :param pulumi.Input[str] source_mask: Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProfileOneConnectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.ProfileOneConnect` Configures a custom profile_oneconnect for use by health checks.

        Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_oneconnect = f5bigip.ltm.ProfileOneConnect("test-oneconnect", name="/Common/test-oneconnect")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        BIG-IP LTM oneconnect profiles can be imported using the `name` , e.g.

        ```sh
        $ pulumi import f5bigip:ltm/profileOneConnect:ProfileOneConnect test-oneconnect /Common/test-oneconnect
        ```

        :param str resource_name: The name of the resource.
        :param ProfileOneConnectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProfileOneConnectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 idle_timeout_override: Optional[pulumi.Input[str]] = None,
                 limit_type: Optional[pulumi.Input[str]] = None,
                 max_age: Optional[pulumi.Input[int]] = None,
                 max_reuse: Optional[pulumi.Input[int]] = None,
                 max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 share_pools: Optional[pulumi.Input[str]] = None,
                 source_mask: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProfileOneConnectArgs.__new__(ProfileOneConnectArgs)

            __props__.__dict__["defaults_from"] = defaults_from
            __props__.__dict__["idle_timeout_override"] = idle_timeout_override
            __props__.__dict__["limit_type"] = limit_type
            __props__.__dict__["max_age"] = max_age
            __props__.__dict__["max_reuse"] = max_reuse
            __props__.__dict__["max_size"] = max_size
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["partition"] = partition
            __props__.__dict__["share_pools"] = share_pools
            __props__.__dict__["source_mask"] = source_mask
        super(ProfileOneConnect, __self__).__init__(
            'f5bigip:ltm/profileOneConnect:ProfileOneConnect',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            idle_timeout_override: Optional[pulumi.Input[str]] = None,
            limit_type: Optional[pulumi.Input[str]] = None,
            max_age: Optional[pulumi.Input[int]] = None,
            max_reuse: Optional[pulumi.Input[int]] = None,
            max_size: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition: Optional[pulumi.Input[str]] = None,
            share_pools: Optional[pulumi.Input[str]] = None,
            source_mask: Optional[pulumi.Input[str]] = None) -> 'ProfileOneConnect':
        """
        Get an existing ProfileOneConnect resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] idle_timeout_override: Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
        :param pulumi.Input[str] limit_type: Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
        :param pulumi.Input[int] max_age: Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
        :param pulumi.Input[int] max_reuse: Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
        :param pulumi.Input[int] max_size: Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
        :param pulumi.Input[str] name: Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] share_pools: Specify if you want to share the pool, default value is `disabled`.
        :param pulumi.Input[str] source_mask: Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProfileOneConnectState.__new__(_ProfileOneConnectState)

        __props__.__dict__["defaults_from"] = defaults_from
        __props__.__dict__["idle_timeout_override"] = idle_timeout_override
        __props__.__dict__["limit_type"] = limit_type
        __props__.__dict__["max_age"] = max_age
        __props__.__dict__["max_reuse"] = max_reuse
        __props__.__dict__["max_size"] = max_size
        __props__.__dict__["name"] = name
        __props__.__dict__["partition"] = partition
        __props__.__dict__["share_pools"] = share_pools
        __props__.__dict__["source_mask"] = source_mask
        return ProfileOneConnect(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> pulumi.Output[str]:
        """
        Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter(name="idleTimeoutOverride")
    def idle_timeout_override(self) -> pulumi.Output[str]:
        """
        Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are `disabled`, `indefinite`, or a numeric value that you specify. The default value is `disabled`
        """
        return pulumi.get(self, "idle_timeout_override")

    @property
    @pulumi.getter(name="limitType")
    def limit_type(self) -> pulumi.Output[str]:
        """
        Controls how connection limits are enforced in conjunction with OneConnect. The default is `None`. Supported Values: `[None,idle,strict]`
        """
        return pulumi.get(self, "limit_type")

    @property
    @pulumi.getter(name="maxAge")
    def max_age(self) -> pulumi.Output[int]:
        """
        Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is `86400`.
        """
        return pulumi.get(self, "max_age")

    @property
    @pulumi.getter(name="maxReuse")
    def max_reuse(self) -> pulumi.Output[int]:
        """
        Specifies the maximum number of times that a server-side connection can be reused. The default value is `1000`.
        """
        return pulumi.get(self, "max_reuse")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Output[int]:
        """
        Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is `10000`.
        """
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of Profile should be full path.The full path is the combination of the `partition + profile_name`,For example `/Common/test-oneconnect-profile`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> pulumi.Output[str]:
        """
        Displays the administrative partition within which this profile resides
        """
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter(name="sharePools")
    def share_pools(self) -> pulumi.Output[str]:
        """
        Specify if you want to share the pool, default value is `disabled`.
        """
        return pulumi.get(self, "share_pools")

    @property
    @pulumi.getter(name="sourceMask")
    def source_mask(self) -> pulumi.Output[str]:
        """
        Specifies a source IP mask. The default value is `0.0.0.0`. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        """
        return pulumi.get(self, "source_mask")

