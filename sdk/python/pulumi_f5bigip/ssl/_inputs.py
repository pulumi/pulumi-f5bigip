# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetWafEntityUrlMethodOverrideArgs',
]

@pulumi.input_type
class GetWafEntityUrlMethodOverrideArgs:
    def __init__(__self__, *,
                 allow: bool,
                 method: str):
        """
        :param bool allow: Specifies that the system allows or disallows a method for this URL
        :param str method: Specifies an HTTP method.
        """
        pulumi.set(__self__, "allow", allow)
        pulumi.set(__self__, "method", method)

    @property
    @pulumi.getter
    def allow(self) -> bool:
        """
        Specifies that the system allows or disallows a method for this URL
        """
        return pulumi.get(self, "allow")

    @allow.setter
    def allow(self, value: bool):
        pulumi.set(self, "allow", value)

    @property
    @pulumi.getter
    def method(self) -> str:
        """
        Specifies an HTTP method.
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: str):
        pulumi.set(self, "method", value)


