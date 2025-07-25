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
from . import _utilities

__all__ = ['PartitionArgs', 'Partition']

@pulumi.input_type
class PartitionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[_builtins.str],
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 route_domain_id: Optional[pulumi.Input[_builtins.int]] = None):
        """
        The set of arguments for constructing a Partition resource.
        :param pulumi.Input[_builtins.str] name: Name of the partition.
        :param pulumi.Input[_builtins.str] description: Description of the partition.
        :param pulumi.Input[_builtins.int] route_domain_id: Route domain id of the partition.
        """
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if route_domain_id is not None:
            pulumi.set(__self__, "route_domain_id", route_domain_id)

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the partition.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Description of the partition.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter(name="routeDomainId")
    def route_domain_id(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Route domain id of the partition.
        """
        return pulumi.get(self, "route_domain_id")

    @route_domain_id.setter
    def route_domain_id(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "route_domain_id", value)


@pulumi.input_type
class _PartitionState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 route_domain_id: Optional[pulumi.Input[_builtins.int]] = None):
        """
        Input properties used for looking up and filtering Partition resources.
        :param pulumi.Input[_builtins.str] description: Description of the partition.
        :param pulumi.Input[_builtins.str] name: Name of the partition.
        :param pulumi.Input[_builtins.int] route_domain_id: Route domain id of the partition.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if route_domain_id is not None:
            pulumi.set(__self__, "route_domain_id", route_domain_id)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Description of the partition.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the partition.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter(name="routeDomainId")
    def route_domain_id(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Route domain id of the partition.
        """
        return pulumi.get(self, "route_domain_id")

    @route_domain_id.setter
    def route_domain_id(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "route_domain_id", value)


@pulumi.type_token("f5bigip:index/partition:Partition")
class Partition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 route_domain_id: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        """
        `Partition` Manages F5 BIG-IP partitions

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_partition = f5bigip.Partition("test-partition",
            name="test-partition",
            description="created by terraform",
            route_domain_id=2)
        ```

        ## Importing

        An existing cipher group can be imported into this resource by supplying the cipher rule full path name ex : `/partition/name`

        An example is below:

        ```sh
        $ terraform import bigip_partition.test_partition test_partition

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: Description of the partition.
        :param pulumi.Input[_builtins.str] name: Name of the partition.
        :param pulumi.Input[_builtins.int] route_domain_id: Route domain id of the partition.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PartitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `Partition` Manages F5 BIG-IP partitions

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_partition = f5bigip.Partition("test-partition",
            name="test-partition",
            description="created by terraform",
            route_domain_id=2)
        ```

        ## Importing

        An existing cipher group can be imported into this resource by supplying the cipher rule full path name ex : `/partition/name`

        An example is below:

        ```sh
        $ terraform import bigip_partition.test_partition test_partition

        ```

        :param str resource_name: The name of the resource.
        :param PartitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PartitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 route_domain_id: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PartitionArgs.__new__(PartitionArgs)

            __props__.__dict__["description"] = description
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["route_domain_id"] = route_domain_id
        super(Partition, __self__).__init__(
            'f5bigip:index/partition:Partition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            route_domain_id: Optional[pulumi.Input[_builtins.int]] = None) -> 'Partition':
        """
        Get an existing Partition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: Description of the partition.
        :param pulumi.Input[_builtins.str] name: Name of the partition.
        :param pulumi.Input[_builtins.int] route_domain_id: Route domain id of the partition.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PartitionState.__new__(_PartitionState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["route_domain_id"] = route_domain_id
        return Partition(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Description of the partition.
        """
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the partition.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter(name="routeDomainId")
    def route_domain_id(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        Route domain id of the partition.
        """
        return pulumi.get(self, "route_domain_id")

