# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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

__all__ = [
    'GetMonitorResult',
    'AwaitableGetMonitorResult',
    'get_monitor',
    'get_monitor_output',
]

@pulumi.output_type
class GetMonitorResult:
    """
    A collection of values returned by getMonitor.
    """
    def __init__(__self__, adaptive=None, adaptive_limit=None, base=None, chase_referrals=None, database=None, defaults_from=None, destination=None, filename=None, filter=None, id=None, interval=None, ip_dscp=None, mandatory_attributes=None, manual_resume=None, mode=None, name=None, partition=None, receive_disable=None, reverse=None, security=None, time_until_up=None, timeout=None, transparent=None, username=None):
        if adaptive and not isinstance(adaptive, str):
            raise TypeError("Expected argument 'adaptive' to be a str")
        pulumi.set(__self__, "adaptive", adaptive)
        if adaptive_limit and not isinstance(adaptive_limit, int):
            raise TypeError("Expected argument 'adaptive_limit' to be a int")
        pulumi.set(__self__, "adaptive_limit", adaptive_limit)
        if base and not isinstance(base, str):
            raise TypeError("Expected argument 'base' to be a str")
        pulumi.set(__self__, "base", base)
        if chase_referrals and not isinstance(chase_referrals, str):
            raise TypeError("Expected argument 'chase_referrals' to be a str")
        pulumi.set(__self__, "chase_referrals", chase_referrals)
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
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interval and not isinstance(interval, int):
            raise TypeError("Expected argument 'interval' to be a int")
        pulumi.set(__self__, "interval", interval)
        if ip_dscp and not isinstance(ip_dscp, int):
            raise TypeError("Expected argument 'ip_dscp' to be a int")
        pulumi.set(__self__, "ip_dscp", ip_dscp)
        if mandatory_attributes and not isinstance(mandatory_attributes, str):
            raise TypeError("Expected argument 'mandatory_attributes' to be a str")
        pulumi.set(__self__, "mandatory_attributes", mandatory_attributes)
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
        if security and not isinstance(security, str):
            raise TypeError("Expected argument 'security' to be a str")
        pulumi.set(__self__, "security", security)
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
        """
        Displays whether adaptive response time monitoring is enabled for this monitor.
        """
        return pulumi.get(self, "adaptive")

    @property
    @pulumi.getter(name="adaptiveLimit")
    def adaptive_limit(self) -> int:
        """
        Displays whether adaptive response time monitoring is enabled for this monitor.
        """
        return pulumi.get(self, "adaptive_limit")

    @property
    @pulumi.getter
    def base(self) -> str:
        return pulumi.get(self, "base")

    @property
    @pulumi.getter(name="chaseReferrals")
    def chase_referrals(self) -> str:
        return pulumi.get(self, "chase_referrals")

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
        """
        id will be full path name of ltm monitor.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> str:
        return pulumi.get(self, "filter")

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
        """
        Specifies, in seconds, the frequency at which the system issues the monitor check when either the resource is down or the status of the resource is unknown.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter(name="ipDscp")
    def ip_dscp(self) -> int:
        """
        Displays the differentiated services code point (DSCP). DSCP is a 6-bit value in the Differentiated Services (DS) field of the IP header.
        """
        return pulumi.get(self, "ip_dscp")

    @property
    @pulumi.getter(name="mandatoryAttributes")
    def mandatory_attributes(self) -> str:
        return pulumi.get(self, "mandatory_attributes")

    @property
    @pulumi.getter(name="manualResume")
    def manual_resume(self) -> str:
        """
        Displays whether the system automatically changes the status of a resource to Enabled at the next successful monitor check.
        """
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
        """
        Instructs the system to mark the target resource down when the test is successful.
        """
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def security(self) -> str:
        return pulumi.get(self, "security")

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
        """
        Displays whether the monitor operates in transparent mode.
        """
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
            base=self.base,
            chase_referrals=self.chase_referrals,
            database=self.database,
            defaults_from=self.defaults_from,
            destination=self.destination,
            filename=self.filename,
            filter=self.filter,
            id=self.id,
            interval=self.interval,
            ip_dscp=self.ip_dscp,
            mandatory_attributes=self.mandatory_attributes,
            manual_resume=self.manual_resume,
            mode=self.mode,
            name=self.name,
            partition=self.partition,
            receive_disable=self.receive_disable,
            reverse=self.reverse,
            security=self.security,
            time_until_up=self.time_until_up,
            timeout=self.timeout,
            transparent=self.transparent,
            username=self.username)


def get_monitor(name: Optional[str] = None,
                partition: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMonitorResult:
    """
    Use this data source (`ltm.Monitor`) to get the ltm monitor details available on BIG-IP

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    monitor__tc1 = f5bigip.ltm.get_monitor(name="test-monitor",
        partition="Common")
    ```


    :param str name: Name of the ltm monitor
    :param str partition: partition of the ltm monitor
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['partition'] = partition
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:ltm/getMonitor:getMonitor', __args__, opts=opts, typ=GetMonitorResult).value

    return AwaitableGetMonitorResult(
        adaptive=pulumi.get(__ret__, 'adaptive'),
        adaptive_limit=pulumi.get(__ret__, 'adaptive_limit'),
        base=pulumi.get(__ret__, 'base'),
        chase_referrals=pulumi.get(__ret__, 'chase_referrals'),
        database=pulumi.get(__ret__, 'database'),
        defaults_from=pulumi.get(__ret__, 'defaults_from'),
        destination=pulumi.get(__ret__, 'destination'),
        filename=pulumi.get(__ret__, 'filename'),
        filter=pulumi.get(__ret__, 'filter'),
        id=pulumi.get(__ret__, 'id'),
        interval=pulumi.get(__ret__, 'interval'),
        ip_dscp=pulumi.get(__ret__, 'ip_dscp'),
        mandatory_attributes=pulumi.get(__ret__, 'mandatory_attributes'),
        manual_resume=pulumi.get(__ret__, 'manual_resume'),
        mode=pulumi.get(__ret__, 'mode'),
        name=pulumi.get(__ret__, 'name'),
        partition=pulumi.get(__ret__, 'partition'),
        receive_disable=pulumi.get(__ret__, 'receive_disable'),
        reverse=pulumi.get(__ret__, 'reverse'),
        security=pulumi.get(__ret__, 'security'),
        time_until_up=pulumi.get(__ret__, 'time_until_up'),
        timeout=pulumi.get(__ret__, 'timeout'),
        transparent=pulumi.get(__ret__, 'transparent'),
        username=pulumi.get(__ret__, 'username'))
def get_monitor_output(name: Optional[pulumi.Input[str]] = None,
                       partition: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMonitorResult]:
    """
    Use this data source (`ltm.Monitor`) to get the ltm monitor details available on BIG-IP

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    monitor__tc1 = f5bigip.ltm.get_monitor(name="test-monitor",
        partition="Common")
    ```


    :param str name: Name of the ltm monitor
    :param str partition: partition of the ltm monitor
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['partition'] = partition
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('f5bigip:ltm/getMonitor:getMonitor', __args__, opts=opts, typ=GetMonitorResult)
    return __ret__.apply(lambda __response__: GetMonitorResult(
        adaptive=pulumi.get(__response__, 'adaptive'),
        adaptive_limit=pulumi.get(__response__, 'adaptive_limit'),
        base=pulumi.get(__response__, 'base'),
        chase_referrals=pulumi.get(__response__, 'chase_referrals'),
        database=pulumi.get(__response__, 'database'),
        defaults_from=pulumi.get(__response__, 'defaults_from'),
        destination=pulumi.get(__response__, 'destination'),
        filename=pulumi.get(__response__, 'filename'),
        filter=pulumi.get(__response__, 'filter'),
        id=pulumi.get(__response__, 'id'),
        interval=pulumi.get(__response__, 'interval'),
        ip_dscp=pulumi.get(__response__, 'ip_dscp'),
        mandatory_attributes=pulumi.get(__response__, 'mandatory_attributes'),
        manual_resume=pulumi.get(__response__, 'manual_resume'),
        mode=pulumi.get(__response__, 'mode'),
        name=pulumi.get(__response__, 'name'),
        partition=pulumi.get(__response__, 'partition'),
        receive_disable=pulumi.get(__response__, 'receive_disable'),
        reverse=pulumi.get(__response__, 'reverse'),
        security=pulumi.get(__response__, 'security'),
        time_until_up=pulumi.get(__response__, 'time_until_up'),
        timeout=pulumi.get(__response__, 'timeout'),
        transparent=pulumi.get(__response__, 'transparent'),
        username=pulumi.get(__response__, 'username')))
