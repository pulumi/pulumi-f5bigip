# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'VlanInterface',
]

@pulumi.output_type
class VlanInterface(dict):
    def __init__(__self__, *,
                 tagged: Optional[bool] = None,
                 vlanport: Optional[str] = None):
        """
        :param bool tagged: Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        :param str vlanport: Physical or virtual port used for traffic
        """
        if tagged is not None:
            pulumi.set(__self__, "tagged", tagged)
        if vlanport is not None:
            pulumi.set(__self__, "vlanport", vlanport)

    @property
    @pulumi.getter
    def tagged(self) -> Optional[bool]:
        """
        Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
        """
        return pulumi.get(self, "tagged")

    @property
    @pulumi.getter
    def vlanport(self) -> Optional[str]:
        """
        Physical or virtual port used for traffic
        """
        return pulumi.get(self, "vlanport")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


