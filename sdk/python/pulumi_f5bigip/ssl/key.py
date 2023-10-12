# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 name: pulumi.Input[str],
                 full_path: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[str] content: Content of SSL certificate key present on local Disk
        :param pulumi.Input[str] name: Name of the SSL Certificate key to be Imported on to BIGIP
        :param pulumi.Input[str] full_path: Full Path Name of ssl key
        :param pulumi.Input[str] partition: Partition of ssl certificate key
        :param pulumi.Input[str] passphrase: Passphrase on key.
        """
        KeyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            content=content,
            name=name,
            full_path=full_path,
            partition=partition,
            passphrase=passphrase,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             content: pulumi.Input[str],
             name: pulumi.Input[str],
             full_path: Optional[pulumi.Input[str]] = None,
             partition: Optional[pulumi.Input[str]] = None,
             passphrase: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("content", content)
        _setter("name", name)
        if full_path is not None:
            _setter("full_path", full_path)
        if partition is not None:
            _setter("partition", partition)
        if passphrase is not None:
            _setter("passphrase", passphrase)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        Content of SSL certificate key present on local Disk
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the SSL Certificate key to be Imported on to BIGIP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> Optional[pulumi.Input[str]]:
        """
        Full Path Name of ssl key
        """
        return pulumi.get(self, "full_path")

    @full_path.setter
    def full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_path", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Partition of ssl certificate key
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter
    def passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Passphrase on key.
        """
        return pulumi.get(self, "passphrase")

    @passphrase.setter
    def passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passphrase", value)


@pulumi.input_type
class _KeyState:
    def __init__(__self__, *,
                 content: Optional[pulumi.Input[str]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Key resources.
        :param pulumi.Input[str] content: Content of SSL certificate key present on local Disk
        :param pulumi.Input[str] full_path: Full Path Name of ssl key
        :param pulumi.Input[str] name: Name of the SSL Certificate key to be Imported on to BIGIP
        :param pulumi.Input[str] partition: Partition of ssl certificate key
        :param pulumi.Input[str] passphrase: Passphrase on key.
        """
        _KeyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            content=content,
            full_path=full_path,
            name=name,
            partition=partition,
            passphrase=passphrase,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             content: Optional[pulumi.Input[str]] = None,
             full_path: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             partition: Optional[pulumi.Input[str]] = None,
             passphrase: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if content is not None:
            _setter("content", content)
        if full_path is not None:
            _setter("full_path", full_path)
        if name is not None:
            _setter("name", name)
        if partition is not None:
            _setter("partition", partition)
        if passphrase is not None:
            _setter("passphrase", passphrase)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        Content of SSL certificate key present on local Disk
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> Optional[pulumi.Input[str]]:
        """
        Full Path Name of ssl key
        """
        return pulumi.get(self, "full_path")

    @full_path.setter
    def full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_path", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the SSL Certificate key to be Imported on to BIGIP
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Partition of ssl certificate key
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter
    def passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Passphrase on key.
        """
        return pulumi.get(self, "passphrase")

    @passphrase.setter
    def passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passphrase", value)


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `ssl.Key` This resource will import SSL certificate key on BIG-IP LTM.
        Certificate key can be imported from certificate key files on the local disk, in PEM format

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_key = f5bigip.ssl.Key("test-key",
            name="serverkey.key",
            content=(lambda path: open(path).read())("serverkey.key"),
            partition="Common")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: Content of SSL certificate key present on local Disk
        :param pulumi.Input[str] full_path: Full Path Name of ssl key
        :param pulumi.Input[str] name: Name of the SSL Certificate key to be Imported on to BIGIP
        :param pulumi.Input[str] partition: Partition of ssl certificate key
        :param pulumi.Input[str] passphrase: Passphrase on key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ssl.Key` This resource will import SSL certificate key on BIG-IP LTM.
        Certificate key can be imported from certificate key files on the local disk, in PEM format

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_key = f5bigip.ssl.Key("test-key",
            name="serverkey.key",
            content=(lambda path: open(path).read())("serverkey.key"),
            partition="Common")
        ```

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            KeyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyArgs.__new__(KeyArgs)

            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = None if content is None else pulumi.Output.secret(content)
            __props__.__dict__["full_path"] = full_path
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["partition"] = partition
            __props__.__dict__["passphrase"] = None if passphrase is None else pulumi.Output.secret(passphrase)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["content", "passphrase"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Key, __self__).__init__(
            'f5bigip:ssl/key:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            content: Optional[pulumi.Input[str]] = None,
            full_path: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            partition: Optional[pulumi.Input[str]] = None,
            passphrase: Optional[pulumi.Input[str]] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: Content of SSL certificate key present on local Disk
        :param pulumi.Input[str] full_path: Full Path Name of ssl key
        :param pulumi.Input[str] name: Name of the SSL Certificate key to be Imported on to BIGIP
        :param pulumi.Input[str] partition: Partition of ssl certificate key
        :param pulumi.Input[str] passphrase: Passphrase on key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeyState.__new__(_KeyState)

        __props__.__dict__["content"] = content
        __props__.__dict__["full_path"] = full_path
        __props__.__dict__["name"] = name
        __props__.__dict__["partition"] = partition
        __props__.__dict__["passphrase"] = passphrase
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        Content of SSL certificate key present on local Disk
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> pulumi.Output[str]:
        """
        Full Path Name of ssl key
        """
        return pulumi.get(self, "full_path")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the SSL Certificate key to be Imported on to BIGIP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> pulumi.Output[Optional[str]]:
        """
        Partition of ssl certificate key
        """
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter
    def passphrase(self) -> pulumi.Output[Optional[str]]:
        """
        Passphrase on key.
        """
        return pulumi.get(self, "passphrase")

