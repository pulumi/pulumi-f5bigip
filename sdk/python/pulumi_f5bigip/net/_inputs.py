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

__all__ = [
    'VlanInterfaceArgs',
    'VlanInterfaceArgsDict',
]

MYPY = False

if not MYPY:
    class VlanInterfaceArgsDict(TypedDict):
        tagged: NotRequired[pulumi.Input[_builtins.bool]]
        """
        Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        """
        vlanport: NotRequired[pulumi.Input[_builtins.str]]
        """
        Physical or virtual port used for traffic
        """
elif False:
    VlanInterfaceArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class VlanInterfaceArgs:
    def __init__(__self__, *,
                 tagged: Optional[pulumi.Input[_builtins.bool]] = None,
                 vlanport: Optional[pulumi.Input[_builtins.str]] = None):
        """
        :param pulumi.Input[_builtins.bool] tagged: Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        :param pulumi.Input[_builtins.str] vlanport: Physical or virtual port used for traffic
        """
        if tagged is not None:
            pulumi.set(__self__, "tagged", tagged)
        if vlanport is not None:
            pulumi.set(__self__, "vlanport", vlanport)

    @_builtins.property
    @pulumi.getter
    def tagged(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        """
        return pulumi.get(self, "tagged")

    @tagged.setter
    def tagged(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "tagged", value)

    @_builtins.property
    @pulumi.getter
    def vlanport(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Physical or virtual port used for traffic
        """
        return pulumi.get(self, "vlanport")

    @vlanport.setter
    def vlanport(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "vlanport", value)


