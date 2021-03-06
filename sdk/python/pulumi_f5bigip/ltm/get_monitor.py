# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetMonitorResult',
    'AwaitableGetMonitorResult',
    'get_monitor',
]

@pulumi.output_type
class GetMonitorResult:
    """
    A collection of values returned by getMonitor.
    """
    def __init__(__self__, adaptive=None, adaptive_limit=None, database=None, defaults_from=None, destination=None, filename=None, id=None, interval=None, ip_dscp=None, manual_resume=None, mode=None, name=None, partition=None, receive_disable=None, reverse=None, time_until_up=None, timeout=None, transparent=None, username=None):
        if adaptive and not isinstance(adaptive, str):
            raise TypeError("Expected argument 'adaptive' to be a str")
        pulumi.set(__self__, "adaptive", adaptive)
        if adaptive_limit and not isinstance(adaptive_limit, int):
            raise TypeError("Expected argument 'adaptive_limit' to be a int")
        pulumi.set(__self__, "adaptive_limit", adaptive_limit)
        if database and not isinstance(database, str):
            raise TypeError("Expected argument 'database' to be a str")
        pulumi.set(__self__, "database", database)
        if defaults_from and not isinstance(defaults_from, str):
            raise TypeError("Expected argument 'defaults_from' to be a str")
        pulumi.set(__self__, "defaults_from", defaults_from)
        if destination and not isinstance(destination, str):
            raise TypeError("Expected argument 'destination' to be a str")
        pulumi.set(__self__, "destination", destination)
        if filename and not isinstance(filename, str):
            raise TypeError("Expected argument 'filename' to be a str")
        pulumi.set(__self__, "filename", filename)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interval and not isinstance(interval, int):
            raise TypeError("Expected argument 'interval' to be a int")
        pulumi.set(__self__, "interval", interval)
        if ip_dscp and not isinstance(ip_dscp, int):
            raise TypeError("Expected argument 'ip_dscp' to be a int")
        pulumi.set(__self__, "ip_dscp", ip_dscp)
        if manual_resume and not isinstance(manual_resume, str):
            raise TypeError("Expected argument 'manual_resume' to be a str")
        pulumi.set(__self__, "manual_resume", manual_resume)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        pulumi.set(__self__, "partition", partition)
        if receive_disable and not isinstance(receive_disable, str):
            raise TypeError("Expected argument 'receive_disable' to be a str")
        pulumi.set(__self__, "receive_disable", receive_disable)
        if reverse and not isinstance(reverse, str):
            raise TypeError("Expected argument 'reverse' to be a str")
        pulumi.set(__self__, "reverse", reverse)
        if time_until_up and not isinstance(time_until_up, int):
            raise TypeError("Expected argument 'time_until_up' to be a int")
        pulumi.set(__self__, "time_until_up", time_until_up)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
        if transparent and not isinstance(transparent, str):
            raise TypeError("Expected argument 'transparent' to be a str")
        pulumi.set(__self__, "transparent", transparent)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def adaptive(self) -> str:
        return pulumi.get(self, "adaptive")

    @property
    @pulumi.getter(name="adaptiveLimit")
    def adaptive_limit(self) -> int:
        return pulumi.get(self, "adaptive_limit")

    @property
    @pulumi.getter
    def database(self) -> str:
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> str:
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter
    def destination(self) -> str:
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def interval(self) -> int:
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="ipDscp")
    def ip_dscp(self) -> int:
        return pulumi.get(self, "ip_dscp")

    @property
    @pulumi.getter(name="manualResume")
    def manual_resume(self) -> str:
        return pulumi.get(self, "manual_resume")

    @property
    @pulumi.getter
    def mode(self) -> str:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> str:
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter(name="receiveDisable")
    def receive_disable(self) -> str:
        return pulumi.get(self, "receive_disable")

    @property
    @pulumi.getter
    def reverse(self) -> str:
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter(name="timeUntilUp")
    def time_until_up(self) -> int:
        return pulumi.get(self, "time_until_up")

    @property
    @pulumi.getter
    def timeout(self) -> int:
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter
    def transparent(self) -> str:
        return pulumi.get(self, "transparent")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


class AwaitableGetMonitorResult(GetMonitorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMonitorResult(
            adaptive=self.adaptive,
            adaptive_limit=self.adaptive_limit,
            database=self.database,
            defaults_from=self.defaults_from,
            destination=self.destination,
            filename=self.filename,
            id=self.id,
            interval=self.interval,
            ip_dscp=self.ip_dscp,
            manual_resume=self.manual_resume,
            mode=self.mode,
            name=self.name,
            partition=self.partition,
            receive_disable=self.receive_disable,
            reverse=self.reverse,
            time_until_up=self.time_until_up,
            timeout=self.timeout,
            transparent=self.transparent,
            username=self.username)


def get_monitor(name: Optional[str] = None,
                partition: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMonitorResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['partition'] = partition
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('f5bigip:ltm/getMonitor:getMonitor', __args__, opts=opts, typ=GetMonitorResult).value

    return AwaitableGetMonitorResult(
        adaptive=__ret__.adaptive,
        adaptive_limit=__ret__.adaptive_limit,
        database=__ret__.database,
        defaults_from=__ret__.defaults_from,
        destination=__ret__.destination,
        filename=__ret__.filename,
        id=__ret__.id,
        interval=__ret__.interval,
        ip_dscp=__ret__.ip_dscp,
        manual_resume=__ret__.manual_resume,
        mode=__ret__.mode,
        name=__ret__.name,
        partition=__ret__.partition,
        receive_disable=__ret__.receive_disable,
        reverse=__ret__.reverse,
        time_until_up=__ret__.time_until_up,
        timeout=__ret__.timeout,
        transparent=__ret__.transparent,
        username=__ret__.username)
