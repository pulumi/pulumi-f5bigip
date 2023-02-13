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
    'GetWafEntityUrlResult',
    'AwaitableGetWafEntityUrlResult',
    'get_waf_entity_url',
    'get_waf_entity_url_output',
]

@pulumi.output_type
class GetWafEntityUrlResult:
    """
    A collection of values returned by getWafEntityUrl.
    """
    def __init__(__self__, description=None, id=None, json=None, method=None, method_overrides=None, name=None, perform_staging=None, protocol=None, signature_overrides_disables=None, type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if method and not isinstance(method, str):
            raise TypeError("Expected argument 'method' to be a str")
        pulumi.set(__self__, "method", method)
        if method_overrides and not isinstance(method_overrides, list):
            raise TypeError("Expected argument 'method_overrides' to be a list")
        pulumi.set(__self__, "method_overrides", method_overrides)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if perform_staging and not isinstance(perform_staging, bool):
            raise TypeError("Expected argument 'perform_staging' to be a bool")
        pulumi.set(__self__, "perform_staging", perform_staging)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if signature_overrides_disables and not isinstance(signature_overrides_disables, list):
            raise TypeError("Expected argument 'signature_overrides_disables' to be a list")
        pulumi.set(__self__, "signature_overrides_disables", signature_overrides_disables)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def json(self) -> str:
        """
        Json string representing created WAF entity URL declaration in JSON format
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def method(self) -> Optional[str]:
        return pulumi.get(self, "method")

    @property
    @pulumi.getter(name="methodOverrides")
    def method_overrides(self) -> Optional[Sequence['outputs.GetWafEntityUrlMethodOverrideResult']]:
        return pulumi.get(self, "method_overrides")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="performStaging")
    def perform_staging(self) -> Optional[bool]:
        return pulumi.get(self, "perform_staging")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="signatureOverridesDisables")
    def signature_overrides_disables(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "signature_overrides_disables")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")


class AwaitableGetWafEntityUrlResult(GetWafEntityUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWafEntityUrlResult(
            description=self.description,
            id=self.id,
            json=self.json,
            method=self.method,
            method_overrides=self.method_overrides,
            name=self.name,
            perform_staging=self.perform_staging,
            protocol=self.protocol,
            signature_overrides_disables=self.signature_overrides_disables,
            type=self.type)


def get_waf_entity_url(description: Optional[str] = None,
                       method: Optional[str] = None,
                       method_overrides: Optional[Sequence[pulumi.InputType['GetWafEntityUrlMethodOverrideArgs']]] = None,
                       name: Optional[str] = None,
                       perform_staging: Optional[bool] = None,
                       protocol: Optional[str] = None,
                       signature_overrides_disables: Optional[Sequence[int]] = None,
                       type: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWafEntityUrlResult:
    """
    Use this data source (_ssl_get_waf_pb_suggestions_) to create JSON for WAF URL to later use with an existing WAF policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    w_afurl1 = f5bigip.ssl.get_waf_entity_url(description="this is a test",
        method_overrides=[
            f5bigip.ssl.GetWafEntityUrlMethodOverrideArgs(
                allow=False,
                method="BCOPY",
            ),
            f5bigip.ssl.GetWafEntityUrlMethodOverrideArgs(
                allow=True,
                method="BDELETE",
            ),
        ],
        name="/foobar",
        perform_staging=True,
        protocol="HTTP",
        signature_overrides_disables=[
            12345678,
            87654321,
        ],
        type="explicit")
    ```


    :param str description: A description of the URL.
    :param str method: Specifies an HTTP method.
    :param Sequence[pulumi.InputType['GetWafEntityUrlMethodOverrideArgs']] method_overrides: A list of methods that are allowed or disallowed for a specific URL.
    :param str name: WAF entity URL name.
    :param bool perform_staging: If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
    :param str protocol: Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
    :param Sequence[int] signature_overrides_disables: List of Attack Signature Ids which are disabled for this particular URL.
    :param str type: Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['method'] = method
    __args__['methodOverrides'] = method_overrides
    __args__['name'] = name
    __args__['performStaging'] = perform_staging
    __args__['protocol'] = protocol
    __args__['signatureOverridesDisables'] = signature_overrides_disables
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:ssl/getWafEntityUrl:getWafEntityUrl', __args__, opts=opts, typ=GetWafEntityUrlResult).value

    return AwaitableGetWafEntityUrlResult(
        description=__ret__.description,
        id=__ret__.id,
        json=__ret__.json,
        method=__ret__.method,
        method_overrides=__ret__.method_overrides,
        name=__ret__.name,
        perform_staging=__ret__.perform_staging,
        protocol=__ret__.protocol,
        signature_overrides_disables=__ret__.signature_overrides_disables,
        type=__ret__.type)


@_utilities.lift_output_func(get_waf_entity_url)
def get_waf_entity_url_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                              method: Optional[pulumi.Input[Optional[str]]] = None,
                              method_overrides: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetWafEntityUrlMethodOverrideArgs']]]]] = None,
                              name: Optional[pulumi.Input[str]] = None,
                              perform_staging: Optional[pulumi.Input[Optional[bool]]] = None,
                              protocol: Optional[pulumi.Input[Optional[str]]] = None,
                              signature_overrides_disables: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                              type: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWafEntityUrlResult]:
    """
    Use this data source (_ssl_get_waf_pb_suggestions_) to create JSON for WAF URL to later use with an existing WAF policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    w_afurl1 = f5bigip.ssl.get_waf_entity_url(description="this is a test",
        method_overrides=[
            f5bigip.ssl.GetWafEntityUrlMethodOverrideArgs(
                allow=False,
                method="BCOPY",
            ),
            f5bigip.ssl.GetWafEntityUrlMethodOverrideArgs(
                allow=True,
                method="BDELETE",
            ),
        ],
        name="/foobar",
        perform_staging=True,
        protocol="HTTP",
        signature_overrides_disables=[
            12345678,
            87654321,
        ],
        type="explicit")
    ```


    :param str description: A description of the URL.
    :param str method: Specifies an HTTP method.
    :param Sequence[pulumi.InputType['GetWafEntityUrlMethodOverrideArgs']] method_overrides: A list of methods that are allowed or disallowed for a specific URL.
    :param str name: WAF entity URL name.
    :param bool perform_staging: If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
    :param str protocol: Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
    :param Sequence[int] signature_overrides_disables: List of Attack Signature Ids which are disabled for this particular URL.
    :param str type: Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
    """
    ...
