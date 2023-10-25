# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FastTemplateArgs', 'FastTemplate']

@pulumi.input_type
class FastTemplateArgs:
    def __init__(__self__, *,
                 md5_hash: pulumi.Input[str],
                 source: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FastTemplate resource.
        :param pulumi.Input[str] md5_hash: MD5 hash of the zip archive file containing FAST template
        :param pulumi.Input[str] source: Path to the zip archive file containing FAST template set on Local Disk
        :param pulumi.Input[str] name: Name of the FAST template set to be created on to BIGIP
        """
        FastTemplateArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            md5_hash=md5_hash,
            source=source,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             md5_hash: Optional[pulumi.Input[str]] = None,
             source: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if md5_hash is None and 'md5Hash' in kwargs:
            md5_hash = kwargs['md5Hash']
        if md5_hash is None:
            raise TypeError("Missing 'md5_hash' argument")
        if source is None:
            raise TypeError("Missing 'source' argument")

        _setter("md5_hash", md5_hash)
        _setter("source", source)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="md5Hash")
    def md5_hash(self) -> pulumi.Input[str]:
        """
        MD5 hash of the zip archive file containing FAST template
        """
        return pulumi.get(self, "md5_hash")

    @md5_hash.setter
    def md5_hash(self, value: pulumi.Input[str]):
        pulumi.set(self, "md5_hash", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        Path to the zip archive file containing FAST template set on Local Disk
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the FAST template set to be created on to BIGIP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FastTemplateState:
    def __init__(__self__, *,
                 md5_hash: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FastTemplate resources.
        :param pulumi.Input[str] md5_hash: MD5 hash of the zip archive file containing FAST template
        :param pulumi.Input[str] name: Name of the FAST template set to be created on to BIGIP
        :param pulumi.Input[str] source: Path to the zip archive file containing FAST template set on Local Disk
        """
        _FastTemplateState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            md5_hash=md5_hash,
            name=name,
            source=source,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             md5_hash: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             source: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if md5_hash is None and 'md5Hash' in kwargs:
            md5_hash = kwargs['md5Hash']

        if md5_hash is not None:
            _setter("md5_hash", md5_hash)
        if name is not None:
            _setter("name", name)
        if source is not None:
            _setter("source", source)

    @property
    @pulumi.getter(name="md5Hash")
    def md5_hash(self) -> Optional[pulumi.Input[str]]:
        """
        MD5 hash of the zip archive file containing FAST template
        """
        return pulumi.get(self, "md5_hash")

    @md5_hash.setter
    def md5_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "md5_hash", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the FAST template set to be created on to BIGIP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the zip archive file containing FAST template set on Local Disk
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)


class FastTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 md5_hash: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `FastTemplate` This resource will import and create FAST template sets on BIG-IP LTM.
        Template set can be imported from zip archive files on the local disk.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] md5_hash: MD5 hash of the zip archive file containing FAST template
        :param pulumi.Input[str] name: Name of the FAST template set to be created on to BIGIP
        :param pulumi.Input[str] source: Path to the zip archive file containing FAST template set on Local Disk
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FastTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `FastTemplate` This resource will import and create FAST template sets on BIG-IP LTM.
        Template set can be imported from zip archive files on the local disk.

        :param str resource_name: The name of the resource.
        :param FastTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FastTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            FastTemplateArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 md5_hash: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FastTemplateArgs.__new__(FastTemplateArgs)

            if md5_hash is None and not opts.urn:
                raise TypeError("Missing required property 'md5_hash'")
            __props__.__dict__["md5_hash"] = md5_hash
            __props__.__dict__["name"] = name
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
        super(FastTemplate, __self__).__init__(
            'f5bigip:index/fastTemplate:FastTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            md5_hash: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None) -> 'FastTemplate':
        """
        Get an existing FastTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] md5_hash: MD5 hash of the zip archive file containing FAST template
        :param pulumi.Input[str] name: Name of the FAST template set to be created on to BIGIP
        :param pulumi.Input[str] source: Path to the zip archive file containing FAST template set on Local Disk
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FastTemplateState.__new__(_FastTemplateState)

        __props__.__dict__["md5_hash"] = md5_hash
        __props__.__dict__["name"] = name
        __props__.__dict__["source"] = source
        return FastTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="md5Hash")
    def md5_hash(self) -> pulumi.Output[str]:
        """
        MD5 hash of the zip archive file containing FAST template
        """
        return pulumi.get(self, "md5_hash")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the FAST template set to be created on to BIGIP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        Path to the zip archive file containing FAST template set on Local Disk
        """
        return pulumi.get(self, "source")

