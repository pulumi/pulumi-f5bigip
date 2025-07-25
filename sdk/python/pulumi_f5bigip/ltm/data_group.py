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
from . import outputs
from ._inputs import *

__all__ = ['DataGroupArgs', 'DataGroup']

@pulumi.input_type
class DataGroupArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[_builtins.str],
                 type: pulumi.Input[_builtins.str],
                 internal: Optional[pulumi.Input[_builtins.bool]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]] = None,
                 records_src: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a DataGroup resource.
        :param pulumi.Input[_builtins.str] name: Name of the datagroup
        :param pulumi.Input[_builtins.str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        :param pulumi.Input[_builtins.bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[_builtins.str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if internal is not None:
            pulumi.set(__self__, "internal", internal)
        if records is not None:
            pulumi.set(__self__, "records", records)
        if records_src is not None:
            pulumi.set(__self__, "records_src", records_src)

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the datagroup
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Input[_builtins.str]:
        """
        datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter
    def internal(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        """
        return pulumi.get(self, "internal")

    @internal.setter
    def internal(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "internal", value)

    @_builtins.property
    @pulumi.getter
    def records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]:
        """
        a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]):
        pulumi.set(self, "records", value)

    @_builtins.property
    @pulumi.getter(name="recordsSrc")
    def records_src(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
        This should be used in conjunction with `internal` attribute set `false`
        """
        return pulumi.get(self, "records_src")

    @records_src.setter
    def records_src(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "records_src", value)


@pulumi.input_type
class _DataGroupState:
    def __init__(__self__, *,
                 internal: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]] = None,
                 records_src: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering DataGroup resources.
        :param pulumi.Input[_builtins.bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[_builtins.str] name: Name of the datagroup
        :param pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[_builtins.str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        :param pulumi.Input[_builtins.str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        if internal is not None:
            pulumi.set(__self__, "internal", internal)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if records is not None:
            pulumi.set(__self__, "records", records)
        if records_src is not None:
            pulumi.set(__self__, "records_src", records_src)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @_builtins.property
    @pulumi.getter
    def internal(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        """
        return pulumi.get(self, "internal")

    @internal.setter
    def internal(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "internal", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the datagroup
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]:
        """
        a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]):
        pulumi.set(self, "records", value)

    @_builtins.property
    @pulumi.getter(name="recordsSrc")
    def records_src(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
        This should be used in conjunction with `internal` attribute set `false`
        """
        return pulumi.get(self, "records_src")

    @records_src.setter
    def records_src(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "records_src", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("f5bigip:ltm/dataGroup:DataGroup")
class DataGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internal: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DataGroupRecordArgs', 'DataGroupRecordArgsDict']]]]] = None,
                 records_src: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        `ltm.DataGroup` Manages internal (in-line) datagroup configuration

        Resource should be named with their`full path`. The full path is the combination of the `partition + name` of the resource, for example `/Common/my-datagroup`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[_builtins.str] name: Name of the datagroup
        :param pulumi.Input[Sequence[pulumi.Input[Union['DataGroupRecordArgs', 'DataGroupRecordArgsDict']]]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[_builtins.str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        :param pulumi.Input[_builtins.str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.DataGroup` Manages internal (in-line) datagroup configuration

        Resource should be named with their`full path`. The full path is the combination of the `partition + name` of the resource, for example `/Common/my-datagroup`.

        :param str resource_name: The name of the resource.
        :param DataGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internal: Optional[pulumi.Input[_builtins.bool]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DataGroupRecordArgs', 'DataGroupRecordArgsDict']]]]] = None,
                 records_src: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataGroupArgs.__new__(DataGroupArgs)

            __props__.__dict__["internal"] = internal
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["records"] = records
            __props__.__dict__["records_src"] = records_src
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(DataGroup, __self__).__init__(
            'f5bigip:ltm/dataGroup:DataGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            internal: Optional[pulumi.Input[_builtins.bool]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DataGroupRecordArgs', 'DataGroupRecordArgsDict']]]]] = None,
            records_src: Optional[pulumi.Input[_builtins.str]] = None,
            type: Optional[pulumi.Input[_builtins.str]] = None) -> 'DataGroup':
        """
        Get an existing DataGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[_builtins.str] name: Name of the datagroup
        :param pulumi.Input[Sequence[pulumi.Input[Union['DataGroupRecordArgs', 'DataGroupRecordArgsDict']]]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[_builtins.str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        :param pulumi.Input[_builtins.str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataGroupState.__new__(_DataGroupState)

        __props__.__dict__["internal"] = internal
        __props__.__dict__["name"] = name
        __props__.__dict__["records"] = records
        __props__.__dict__["records_src"] = records_src
        __props__.__dict__["type"] = type
        return DataGroup(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def internal(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        """
        return pulumi.get(self, "internal")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the datagroup
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def records(self) -> pulumi.Output[Optional[Sequence['outputs.DataGroupRecord']]]:
        """
        a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        return pulumi.get(self, "records")

    @_builtins.property
    @pulumi.getter(name="recordsSrc")
    def records_src(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
        This should be used in conjunction with `internal` attribute set `false`
        """
        return pulumi.get(self, "records_src")

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Output[_builtins.str]:
        """
        datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        return pulumi.get(self, "type")

