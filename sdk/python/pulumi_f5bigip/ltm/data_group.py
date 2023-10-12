# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DataGroupArgs', 'DataGroup']

@pulumi.input_type
class DataGroupArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 internal: Optional[pulumi.Input[bool]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]] = None,
                 records_src: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataGroup resource.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        :param pulumi.Input[bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        """
        DataGroupArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            type=type,
            internal=internal,
            records=records,
            records_src=records_src,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: pulumi.Input[str],
             type: pulumi.Input[str],
             internal: Optional[pulumi.Input[bool]] = None,
             records: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]] = None,
             records_src: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("name", name)
        _setter("type", type)
        if internal is not None:
            _setter("internal", internal)
        if records is not None:
            _setter("records", records)
        if records_src is not None:
            _setter("records_src", records_src)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def internal(self) -> Optional[pulumi.Input[bool]]:
        """
        Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        """
        return pulumi.get(self, "internal")

    @internal.setter
    def internal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal", value)

    @property
    @pulumi.getter
    def records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]:
        """
        a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]):
        pulumi.set(self, "records", value)

    @property
    @pulumi.getter(name="recordsSrc")
    def records_src(self) -> Optional[pulumi.Input[str]]:
        """
        Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
        This should be used in conjunction with `internal` attribute set `false`
        """
        return pulumi.get(self, "records_src")

    @records_src.setter
    def records_src(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "records_src", value)


@pulumi.input_type
class _DataGroupState:
    def __init__(__self__, *,
                 internal: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]] = None,
                 records_src: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataGroup resources.
        :param pulumi.Input[bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        _DataGroupState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            internal=internal,
            name=name,
            records=records,
            records_src=records_src,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             internal: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             records: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]] = None,
             records_src: Optional[pulumi.Input[str]] = None,
             type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if internal is not None:
            _setter("internal", internal)
        if name is not None:
            _setter("name", name)
        if records is not None:
            _setter("records", records)
        if records_src is not None:
            _setter("records_src", records_src)
        if type is not None:
            _setter("type", type)

    @property
    @pulumi.getter
    def internal(self) -> Optional[pulumi.Input[bool]]:
        """
        Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        """
        return pulumi.get(self, "internal")

    @internal.setter
    def internal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]:
        """
        a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]):
        pulumi.set(self, "records", value)

    @property
    @pulumi.getter(name="recordsSrc")
    def records_src(self) -> Optional[pulumi.Input[str]]:
        """
        Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
        This should be used in conjunction with `internal` attribute set `false`
        """
        return pulumi.get(self, "records_src")

    @records_src.setter
    def records_src(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "records_src", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class DataGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internal: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]]] = None,
                 records_src: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `ltm.DataGroup` Manages internal (in-line) datagroup configuration

        Resource should be named with their`full path`. The full path is the combination of the `partition + name` of the resource, for example `/Common/my-datagroup`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DataGroupArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 internal: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]]] = None,
                 records_src: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            internal: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]]] = None,
            records_src: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'DataGroup':
        """
        Get an existing DataGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] internal: Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[str] records_src: Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
               This should be used in conjunction with `internal` attribute set `false`
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataGroupState.__new__(_DataGroupState)

        __props__.__dict__["internal"] = internal
        __props__.__dict__["name"] = name
        __props__.__dict__["records"] = records
        __props__.__dict__["records_src"] = records_src
        __props__.__dict__["type"] = type
        return DataGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def internal(self) -> pulumi.Output[Optional[bool]]:
        """
        Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
        """
        return pulumi.get(self, "internal")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def records(self) -> pulumi.Output[Optional[Sequence['outputs.DataGroupRecord']]]:
        """
        a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        return pulumi.get(self, "records")

    @property
    @pulumi.getter(name="recordsSrc")
    def records_src(self) -> pulumi.Output[Optional[str]]:
        """
        Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
        This should be used in conjunction with `internal` attribute set `false`
        """
        return pulumi.get(self, "records_src")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        return pulumi.get(self, "type")

