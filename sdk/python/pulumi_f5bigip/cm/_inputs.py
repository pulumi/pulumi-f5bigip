# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'DeviceGroupDeviceArgs',
    'DeviceGroupDeviceArgsDict',
]

MYPY = False

if not MYPY:
    class DeviceGroupDeviceArgsDict(TypedDict):
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        Is the name of the device Group
        """
        set_sync_leader: NotRequired[pulumi.Input[builtins.bool]]
        """
        Name of origin
        """
elif False:
    DeviceGroupDeviceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DeviceGroupDeviceArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 set_sync_leader: Optional[pulumi.Input[builtins.bool]] = None):
        """
        :param pulumi.Input[builtins.str] name: Is the name of the device Group
        :param pulumi.Input[builtins.bool] set_sync_leader: Name of origin
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if set_sync_leader is not None:
            pulumi.set(__self__, "set_sync_leader", set_sync_leader)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Is the name of the device Group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="setSyncLeader")
    def set_sync_leader(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Name of origin
        """
        return pulumi.get(self, "set_sync_leader")

    @set_sync_leader.setter
    def set_sync_leader(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "set_sync_leader", value)


