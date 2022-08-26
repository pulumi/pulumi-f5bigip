# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetDataGroupResult',
    'AwaitableGetDataGroupResult',
    'get_data_group',
    'get_data_group_output',
]

@pulumi.output_type
class GetDataGroupResult:
    """
    A collection of values returned by getDataGroup.
    """
    def __init__(__self__, id=None, name=None, partition=None, records=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        pulumi.set(__self__, "partition", partition)
        if records and not isinstance(records, list):
            raise TypeError("Expected argument 'records' to be a list")
        pulumi.set(__self__, "records", records)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> str:
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter
    def records(self) -> Sequence['outputs.GetDataGroupRecordResult']:
        """
        Specifies record of type (string/ip/integer)
        """
        return pulumi.get(self, "records")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Data Group type (string, ip, integer)"
        """
        return pulumi.get(self, "type")


class AwaitableGetDataGroupResult(GetDataGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataGroupResult(
            id=self.id,
            name=self.name,
            partition=self.partition,
            records=self.records,
            type=self.type)


def get_data_group(name: Optional[str] = None,
                   partition: Optional[str] = None,
                   records: Optional[Sequence[pulumi.InputType['GetDataGroupRecordArgs']]] = None,
                   type: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataGroupResult:
    """
    Use this data source (`ltm.DataGroup`) to get the data group details available on BIG-IP

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    d_g__tc3 = f5bigip.ltm.get_data_group(name="test-dg",
        partition="Common")
    ```


    :param str name: Name of the datagroup
    :param str partition: partition of the datagroup
    :param Sequence[pulumi.InputType['GetDataGroupRecordArgs']] records: Specifies record of type (string/ip/integer)
    :param str type: The Data Group type (string, ip, integer)"
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['partition'] = partition
    __args__['records'] = records
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:ltm/getDataGroup:getDataGroup', __args__, opts=opts, typ=GetDataGroupResult).value

    return AwaitableGetDataGroupResult(
        id=__ret__.id,
        name=__ret__.name,
        partition=__ret__.partition,
        records=__ret__.records,
        type=__ret__.type)


@_utilities.lift_output_func(get_data_group)
def get_data_group_output(name: Optional[pulumi.Input[str]] = None,
                          partition: Optional[pulumi.Input[str]] = None,
                          records: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetDataGroupRecordArgs']]]]] = None,
                          type: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataGroupResult]:
    """
    Use this data source (`ltm.DataGroup`) to get the data group details available on BIG-IP

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    d_g__tc3 = f5bigip.ltm.get_data_group(name="test-dg",
        partition="Common")
    ```


    :param str name: Name of the datagroup
    :param str partition: partition of the datagroup
    :param Sequence[pulumi.InputType['GetDataGroupRecordArgs']] records: Specifies record of type (string/ip/integer)
    :param str type: The Data Group type (string, ip, integer)"
    """
    ...
