# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DataGroupArgs', 'DataGroup']

@pulumi.input_type
class DataGroupArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 records: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]] = None):
        """
        The set of arguments for constructing a DataGroup resource.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        :param pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if records is not None:
            pulumi.set(__self__, "records", records)

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
    def records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]:
        """
        a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataGroupRecordArgs']]]]):
        pulumi.set(self, "records", value)


class DataGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.DataGroup` Manages internal (in-line) datagroup configuration

        Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-datagroup.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        datagroup = f5bigip.ltm.DataGroup("datagroup",
            name="/Common/dgx2",
            records=[
                f5bigip.ltm.DataGroupRecordArgs(
                    data="pool1",
                    name="abc.com",
                ),
                f5bigip.ltm.DataGroupRecordArgs(
                    data="123",
                    name="test",
                ),
            ],
            type="string")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
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

        Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-datagroup.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        datagroup = f5bigip.ltm.DataGroup("datagroup",
            name="/Common/dgx2",
            records=[
                f5bigip.ltm.DataGroupRecordArgs(
                    data="pool1",
                    name="abc.com",
                ),
                f5bigip.ltm.DataGroupRecordArgs(
                    data="123",
                    name="test",
                ),
            ],
            type="string")
        ```

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
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['records'] = records
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(DataGroup, __self__).__init__(
            'f5bigip:ltm/dataGroup:DataGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            records: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'DataGroup':
        """
        Get an existing DataGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataGroupRecordArgs']]]] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["records"] = records
        __props__["type"] = type
        return DataGroup(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

