# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetWafEntityParameterUrlResult',
    'GetWafEntityUrlMethodOverrideResult',
]

@pulumi.output_type
class GetWafEntityParameterUrlResult(dict):
    def __init__(__self__, *,
                 method: str,
                 name: str,
                 protocol: str,
                 type: str):
        GetWafEntityParameterUrlResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            method=method,
            name=name,
            protocol=protocol,
            type=type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             method: Optional[str] = None,
             name: Optional[str] = None,
             protocol: Optional[str] = None,
             type: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if method is None:
            raise TypeError("Missing 'method' argument")
        if name is None:
            raise TypeError("Missing 'name' argument")
        if protocol is None:
            raise TypeError("Missing 'protocol' argument")
        if type is None:
            raise TypeError("Missing 'type' argument")

        _setter("method", method)
        _setter("name", name)
        _setter("protocol", protocol)
        _setter("type", type)

    @property
    @pulumi.getter
    def method(self) -> str:
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class GetWafEntityUrlMethodOverrideResult(dict):
    def __init__(__self__, *,
                 allow: bool,
                 method: str):
        """
        :param bool allow: Specifies that the system allows or disallows a method for this URL
        :param str method: Specifies an HTTP method.
        """
        GetWafEntityUrlMethodOverrideResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            allow=allow,
            method=method,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             allow: Optional[bool] = None,
             method: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if allow is None:
            raise TypeError("Missing 'allow' argument")
        if method is None:
            raise TypeError("Missing 'method' argument")

        _setter("allow", allow)
        _setter("method", method)

    @property
    @pulumi.getter
    def allow(self) -> bool:
        """
        Specifies that the system allows or disallows a method for this URL
        """
        return pulumi.get(self, "allow")

    @property
    @pulumi.getter
    def method(self) -> str:
        """
        Specifies an HTTP method.
        """
        return pulumi.get(self, "method")


