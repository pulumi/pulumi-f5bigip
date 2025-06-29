# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
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
    def __init__(__self__, cross_origin_requests_enforcements=None, description=None, id=None, json=None, method=None, method_overrides=None, name=None, perform_staging=None, protocol=None, signature_overrides_disables=None, type=None):
        if cross_origin_requests_enforcements and not isinstance(cross_origin_requests_enforcements, list):
            raise TypeError("Expected argument 'cross_origin_requests_enforcements' to be a list")
        pulumi.set(__self__, "cross_origin_requests_enforcements", cross_origin_requests_enforcements)
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
    @pulumi.getter(name="crossOriginRequestsEnforcements")
    def cross_origin_requests_enforcements(self) -> Optional[Sequence['outputs.GetWafEntityUrlCrossOriginRequestsEnforcementResult']]:
        return pulumi.get(self, "cross_origin_requests_enforcements")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def json(self) -> builtins.str:
        """
        Json string representing created WAF entity URL declaration in JSON format
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def method(self) -> Optional[builtins.str]:
        return pulumi.get(self, "method")

    @property
    @pulumi.getter(name="methodOverrides")
    def method_overrides(self) -> Optional[Sequence['outputs.GetWafEntityUrlMethodOverrideResult']]:
        return pulumi.get(self, "method_overrides")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="performStaging")
    def perform_staging(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "perform_staging")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[builtins.str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="signatureOverridesDisables")
    def signature_overrides_disables(self) -> Optional[Sequence[builtins.int]]:
        return pulumi.get(self, "signature_overrides_disables")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")


class AwaitableGetWafEntityUrlResult(GetWafEntityUrlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWafEntityUrlResult(
            cross_origin_requests_enforcements=self.cross_origin_requests_enforcements,
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


def get_waf_entity_url(cross_origin_requests_enforcements: Optional[Sequence[Union['GetWafEntityUrlCrossOriginRequestsEnforcementArgs', 'GetWafEntityUrlCrossOriginRequestsEnforcementArgsDict']]] = None,
                       description: Optional[builtins.str] = None,
                       method: Optional[builtins.str] = None,
                       method_overrides: Optional[Sequence[Union['GetWafEntityUrlMethodOverrideArgs', 'GetWafEntityUrlMethodOverrideArgsDict']]] = None,
                       name: Optional[builtins.str] = None,
                       perform_staging: Optional[builtins.bool] = None,
                       protocol: Optional[builtins.str] = None,
                       signature_overrides_disables: Optional[Sequence[builtins.int]] = None,
                       type: Optional[builtins.str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWafEntityUrlResult:
    """
    Use this data source (`ssl_get_waf_pb_suggestions`) to create JSON for WAF URL to later use with an existing WAF policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    wafurl1 = f5bigip.ssl.get_waf_entity_url(name="/foobar",
        description="this is a test",
        type="explicit",
        protocol="HTTP",
        perform_staging=True,
        signature_overrides_disables=[
            12345678,
            87654321,
        ],
        method_overrides=[
            {
                "allow": False,
                "method": "BCOPY",
            },
            {
                "allow": True,
                "method": "BDELETE",
            },
        ],
        cross_origin_requests_enforcements=[
            {
                "include_subdomains": True,
                "origin_name": "app1.com",
                "origin_port": "80",
                "origin_protocol": "http",
            },
            {
                "include_subdomains": True,
                "origin_name": "app2.com",
                "origin_port": "443",
                "origin_protocol": "http",
            },
        ])
    ```


    :param Sequence[Union['GetWafEntityUrlCrossOriginRequestsEnforcementArgs', 'GetWafEntityUrlCrossOriginRequestsEnforcementArgsDict']] cross_origin_requests_enforcements: A list of options that enables your web-application to share data with a website hosted on a
           different domain.
    :param builtins.str description: A description of the URL.
    :param builtins.str method: Select a Method for the URL to create an API endpoint. Default is : *.
    :param Sequence[Union['GetWafEntityUrlMethodOverrideArgs', 'GetWafEntityUrlMethodOverrideArgsDict']] method_overrides: A list of methods that are allowed or disallowed for a specific URL.
    :param builtins.str name: WAF entity URL name.
    :param builtins.bool perform_staging: If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
    :param builtins.str protocol: Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
    :param Sequence[builtins.int] signature_overrides_disables: List of Attack Signature Ids which are disabled for this particular URL.
    :param builtins.str type: Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
    """
    __args__ = dict()
    __args__['crossOriginRequestsEnforcements'] = cross_origin_requests_enforcements
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
        cross_origin_requests_enforcements=pulumi.get(__ret__, 'cross_origin_requests_enforcements'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        json=pulumi.get(__ret__, 'json'),
        method=pulumi.get(__ret__, 'method'),
        method_overrides=pulumi.get(__ret__, 'method_overrides'),
        name=pulumi.get(__ret__, 'name'),
        perform_staging=pulumi.get(__ret__, 'perform_staging'),
        protocol=pulumi.get(__ret__, 'protocol'),
        signature_overrides_disables=pulumi.get(__ret__, 'signature_overrides_disables'),
        type=pulumi.get(__ret__, 'type'))
def get_waf_entity_url_output(cross_origin_requests_enforcements: Optional[pulumi.Input[Optional[Sequence[Union['GetWafEntityUrlCrossOriginRequestsEnforcementArgs', 'GetWafEntityUrlCrossOriginRequestsEnforcementArgsDict']]]]] = None,
                              description: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              method: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              method_overrides: Optional[pulumi.Input[Optional[Sequence[Union['GetWafEntityUrlMethodOverrideArgs', 'GetWafEntityUrlMethodOverrideArgsDict']]]]] = None,
                              name: Optional[pulumi.Input[builtins.str]] = None,
                              perform_staging: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                              protocol: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              signature_overrides_disables: Optional[pulumi.Input[Optional[Sequence[builtins.int]]]] = None,
                              type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetWafEntityUrlResult]:
    """
    Use this data source (`ssl_get_waf_pb_suggestions`) to create JSON for WAF URL to later use with an existing WAF policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    wafurl1 = f5bigip.ssl.get_waf_entity_url(name="/foobar",
        description="this is a test",
        type="explicit",
        protocol="HTTP",
        perform_staging=True,
        signature_overrides_disables=[
            12345678,
            87654321,
        ],
        method_overrides=[
            {
                "allow": False,
                "method": "BCOPY",
            },
            {
                "allow": True,
                "method": "BDELETE",
            },
        ],
        cross_origin_requests_enforcements=[
            {
                "include_subdomains": True,
                "origin_name": "app1.com",
                "origin_port": "80",
                "origin_protocol": "http",
            },
            {
                "include_subdomains": True,
                "origin_name": "app2.com",
                "origin_port": "443",
                "origin_protocol": "http",
            },
        ])
    ```


    :param Sequence[Union['GetWafEntityUrlCrossOriginRequestsEnforcementArgs', 'GetWafEntityUrlCrossOriginRequestsEnforcementArgsDict']] cross_origin_requests_enforcements: A list of options that enables your web-application to share data with a website hosted on a
           different domain.
    :param builtins.str description: A description of the URL.
    :param builtins.str method: Select a Method for the URL to create an API endpoint. Default is : *.
    :param Sequence[Union['GetWafEntityUrlMethodOverrideArgs', 'GetWafEntityUrlMethodOverrideArgsDict']] method_overrides: A list of methods that are allowed or disallowed for a specific URL.
    :param builtins.str name: WAF entity URL name.
    :param builtins.bool perform_staging: If true then any violation associated to the respective URL will not be enforced, and the request will not be considered illegal.
    :param builtins.str protocol: Specifies whether the protocol for the URL is 'http' or 'https'. Default is: http.
    :param Sequence[builtins.int] signature_overrides_disables: List of Attack Signature Ids which are disabled for this particular URL.
    :param builtins.str type: Specifies whether the parameter is an 'explicit' or a 'wildcard' attribute. Default is: wildcard.
    """
    __args__ = dict()
    __args__['crossOriginRequestsEnforcements'] = cross_origin_requests_enforcements
    __args__['description'] = description
    __args__['method'] = method
    __args__['methodOverrides'] = method_overrides
    __args__['name'] = name
    __args__['performStaging'] = perform_staging
    __args__['protocol'] = protocol
    __args__['signatureOverridesDisables'] = signature_overrides_disables
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('f5bigip:ssl/getWafEntityUrl:getWafEntityUrl', __args__, opts=opts, typ=GetWafEntityUrlResult)
    return __ret__.apply(lambda __response__: GetWafEntityUrlResult(
        cross_origin_requests_enforcements=pulumi.get(__response__, 'cross_origin_requests_enforcements'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        json=pulumi.get(__response__, 'json'),
        method=pulumi.get(__response__, 'method'),
        method_overrides=pulumi.get(__response__, 'method_overrides'),
        name=pulumi.get(__response__, 'name'),
        perform_staging=pulumi.get(__response__, 'perform_staging'),
        protocol=pulumi.get(__response__, 'protocol'),
        signature_overrides_disables=pulumi.get(__response__, 'signature_overrides_disables'),
        type=pulumi.get(__response__, 'type')))
