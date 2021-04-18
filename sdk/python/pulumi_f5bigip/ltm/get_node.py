# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetNodeResult',
    'AwaitableGetNodeResult',
    'get_node',
]

@pulumi.output_type
class GetNodeResult:
    """
    A collection of values returned by getNode.
    """
    def __init__(__self__, address=None, connection_limit=None, description=None, dynamic_ratio=None, fqdn=None, id=None, monitor=None, name=None, partition=None, rate_limit=None, ratio=None, state=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if connection_limit and not isinstance(connection_limit, int):
            raise TypeError("Expected argument 'connection_limit' to be a int")
        pulumi.set(__self__, "connection_limit", connection_limit)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dynamic_ratio and not isinstance(dynamic_ratio, int):
            raise TypeError("Expected argument 'dynamic_ratio' to be a int")
        pulumi.set(__self__, "dynamic_ratio", dynamic_ratio)
        if fqdn and not isinstance(fqdn, dict):
            raise TypeError("Expected argument 'fqdn' to be a dict")
        pulumi.set(__self__, "fqdn", fqdn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if monitor and not isinstance(monitor, str):
            raise TypeError("Expected argument 'monitor' to be a str")
        pulumi.set(__self__, "monitor", monitor)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        pulumi.set(__self__, "partition", partition)
        if rate_limit and not isinstance(rate_limit, str):
            raise TypeError("Expected argument 'rate_limit' to be a str")
        pulumi.set(__self__, "rate_limit", rate_limit)
        if ratio and not isinstance(ratio, int):
            raise TypeError("Expected argument 'ratio' to be a int")
        pulumi.set(__self__, "ratio", ratio)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> int:
        return pulumi.get(self, "connection_limit")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dynamicRatio")
    def dynamic_ratio(self) -> int:
        return pulumi.get(self, "dynamic_ratio")

    @property
    @pulumi.getter
    def fqdn(self) -> 'outputs.GetNodeFqdnResult':
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def monitor(self) -> str:
        return pulumi.get(self, "monitor")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> str:
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter(name="rateLimit")
    def rate_limit(self) -> str:
        return pulumi.get(self, "rate_limit")

    @property
    @pulumi.getter
    def ratio(self) -> int:
        return pulumi.get(self, "ratio")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")


class AwaitableGetNodeResult(GetNodeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodeResult(
            address=self.address,
            connection_limit=self.connection_limit,
            description=self.description,
            dynamic_ratio=self.dynamic_ratio,
            fqdn=self.fqdn,
            id=self.id,
            monitor=self.monitor,
            name=self.name,
            partition=self.partition,
            rate_limit=self.rate_limit,
            ratio=self.ratio,
            state=self.state)


def get_node(address: Optional[str] = None,
             description: Optional[str] = None,
             fqdn: Optional[pulumi.InputType['GetNodeFqdnArgs']] = None,
             name: Optional[str] = None,
             partition: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNodeResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['address'] = address
    __args__['description'] = description
    __args__['fqdn'] = fqdn
    __args__['name'] = name
    __args__['partition'] = partition
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('f5bigip:ltm/getNode:getNode', __args__, opts=opts, typ=GetNodeResult).value

    return AwaitableGetNodeResult(
        address=__ret__.address,
        connection_limit=__ret__.connection_limit,
        description=__ret__.description,
        dynamic_ratio=__ret__.dynamic_ratio,
        fqdn=__ret__.fqdn,
        id=__ret__.id,
        monitor=__ret__.monitor,
        name=__ret__.name,
        partition=__ret__.partition,
        rate_limit=__ret__.rate_limit,
        ratio=__ret__.ratio,
        state=__ret__.state)
