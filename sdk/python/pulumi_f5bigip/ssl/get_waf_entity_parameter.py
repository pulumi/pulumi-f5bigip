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
    'GetWafEntityParameterResult',
    'AwaitableGetWafEntityParameterResult',
    'get_waf_entity_parameter',
    'get_waf_entity_parameter_output',
]

@pulumi.output_type
class GetWafEntityParameterResult:
    """
    A collection of values returned by getWafEntityParameter.
    """
    def __init__(__self__, allow_empty_type=None, allow_repeated_parameter_name=None, attack_signatures_check=None, check_max_value_length=None, check_min_value_length=None, data_type=None, description=None, enable_regular_expression=None, id=None, is_base64=None, is_cookie=None, is_header=None, json=None, level=None, mandatory=None, metachars_on_parameter_value_check=None, name=None, parameter_location=None, perform_staging=None, sensitive_parameter=None, signature_overrides_disables=None, type=None, url=None, value_type=None):
        if allow_empty_type and not isinstance(allow_empty_type, bool):
            raise TypeError("Expected argument 'allow_empty_type' to be a bool")
        pulumi.set(__self__, "allow_empty_type", allow_empty_type)
        if allow_repeated_parameter_name and not isinstance(allow_repeated_parameter_name, bool):
            raise TypeError("Expected argument 'allow_repeated_parameter_name' to be a bool")
        pulumi.set(__self__, "allow_repeated_parameter_name", allow_repeated_parameter_name)
        if attack_signatures_check and not isinstance(attack_signatures_check, bool):
            raise TypeError("Expected argument 'attack_signatures_check' to be a bool")
        pulumi.set(__self__, "attack_signatures_check", attack_signatures_check)
        if check_max_value_length and not isinstance(check_max_value_length, bool):
            raise TypeError("Expected argument 'check_max_value_length' to be a bool")
        pulumi.set(__self__, "check_max_value_length", check_max_value_length)
        if check_min_value_length and not isinstance(check_min_value_length, bool):
            raise TypeError("Expected argument 'check_min_value_length' to be a bool")
        pulumi.set(__self__, "check_min_value_length", check_min_value_length)
        if data_type and not isinstance(data_type, str):
            raise TypeError("Expected argument 'data_type' to be a str")
        pulumi.set(__self__, "data_type", data_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_regular_expression and not isinstance(enable_regular_expression, bool):
            raise TypeError("Expected argument 'enable_regular_expression' to be a bool")
        pulumi.set(__self__, "enable_regular_expression", enable_regular_expression)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_base64 and not isinstance(is_base64, bool):
            raise TypeError("Expected argument 'is_base64' to be a bool")
        pulumi.set(__self__, "is_base64", is_base64)
        if is_cookie and not isinstance(is_cookie, bool):
            raise TypeError("Expected argument 'is_cookie' to be a bool")
        pulumi.set(__self__, "is_cookie", is_cookie)
        if is_header and not isinstance(is_header, bool):
            raise TypeError("Expected argument 'is_header' to be a bool")
        pulumi.set(__self__, "is_header", is_header)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if level and not isinstance(level, str):
            raise TypeError("Expected argument 'level' to be a str")
        pulumi.set(__self__, "level", level)
        if mandatory and not isinstance(mandatory, bool):
            raise TypeError("Expected argument 'mandatory' to be a bool")
        pulumi.set(__self__, "mandatory", mandatory)
        if metachars_on_parameter_value_check and not isinstance(metachars_on_parameter_value_check, bool):
            raise TypeError("Expected argument 'metachars_on_parameter_value_check' to be a bool")
        pulumi.set(__self__, "metachars_on_parameter_value_check", metachars_on_parameter_value_check)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parameter_location and not isinstance(parameter_location, str):
            raise TypeError("Expected argument 'parameter_location' to be a str")
        pulumi.set(__self__, "parameter_location", parameter_location)
        if perform_staging and not isinstance(perform_staging, bool):
            raise TypeError("Expected argument 'perform_staging' to be a bool")
        pulumi.set(__self__, "perform_staging", perform_staging)
        if sensitive_parameter and not isinstance(sensitive_parameter, bool):
            raise TypeError("Expected argument 'sensitive_parameter' to be a bool")
        pulumi.set(__self__, "sensitive_parameter", sensitive_parameter)
        if signature_overrides_disables and not isinstance(signature_overrides_disables, list):
            raise TypeError("Expected argument 'signature_overrides_disables' to be a list")
        pulumi.set(__self__, "signature_overrides_disables", signature_overrides_disables)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if url and not isinstance(url, dict):
            raise TypeError("Expected argument 'url' to be a dict")
        pulumi.set(__self__, "url", url)
        if value_type and not isinstance(value_type, str):
            raise TypeError("Expected argument 'value_type' to be a str")
        pulumi.set(__self__, "value_type", value_type)

    @property
    @pulumi.getter(name="allowEmptyType")
    def allow_empty_type(self) -> Optional[bool]:
        return pulumi.get(self, "allow_empty_type")

    @property
    @pulumi.getter(name="allowRepeatedParameterName")
    def allow_repeated_parameter_name(self) -> Optional[bool]:
        return pulumi.get(self, "allow_repeated_parameter_name")

    @property
    @pulumi.getter(name="attackSignaturesCheck")
    def attack_signatures_check(self) -> Optional[bool]:
        return pulumi.get(self, "attack_signatures_check")

    @property
    @pulumi.getter(name="checkMaxValueLength")
    def check_max_value_length(self) -> Optional[bool]:
        return pulumi.get(self, "check_max_value_length")

    @property
    @pulumi.getter(name="checkMinValueLength")
    def check_min_value_length(self) -> Optional[bool]:
        return pulumi.get(self, "check_min_value_length")

    @property
    @pulumi.getter(name="dataType")
    def data_type(self) -> Optional[str]:
        return pulumi.get(self, "data_type")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableRegularExpression")
    def enable_regular_expression(self) -> Optional[bool]:
        return pulumi.get(self, "enable_regular_expression")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isBase64")
    def is_base64(self) -> Optional[bool]:
        return pulumi.get(self, "is_base64")

    @property
    @pulumi.getter(name="isCookie")
    def is_cookie(self) -> Optional[bool]:
        return pulumi.get(self, "is_cookie")

    @property
    @pulumi.getter(name="isHeader")
    def is_header(self) -> Optional[bool]:
        return pulumi.get(self, "is_header")

    @property
    @pulumi.getter
    def json(self) -> str:
        return pulumi.get(self, "json")

    @property
    @pulumi.getter
    def level(self) -> Optional[str]:
        return pulumi.get(self, "level")

    @property
    @pulumi.getter
    def mandatory(self) -> Optional[bool]:
        return pulumi.get(self, "mandatory")

    @property
    @pulumi.getter(name="metacharsOnParameterValueCheck")
    def metachars_on_parameter_value_check(self) -> Optional[bool]:
        return pulumi.get(self, "metachars_on_parameter_value_check")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parameterLocation")
    def parameter_location(self) -> Optional[str]:
        return pulumi.get(self, "parameter_location")

    @property
    @pulumi.getter(name="performStaging")
    def perform_staging(self) -> Optional[bool]:
        return pulumi.get(self, "perform_staging")

    @property
    @pulumi.getter(name="sensitiveParameter")
    def sensitive_parameter(self) -> Optional[bool]:
        return pulumi.get(self, "sensitive_parameter")

    @property
    @pulumi.getter(name="signatureOverridesDisables")
    def signature_overrides_disables(self) -> Optional[Sequence[int]]:
        return pulumi.get(self, "signature_overrides_disables")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> Optional['outputs.GetWafEntityParameterUrlResult']:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> Optional[str]:
        return pulumi.get(self, "value_type")


class AwaitableGetWafEntityParameterResult(GetWafEntityParameterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWafEntityParameterResult(
            allow_empty_type=self.allow_empty_type,
            allow_repeated_parameter_name=self.allow_repeated_parameter_name,
            attack_signatures_check=self.attack_signatures_check,
            check_max_value_length=self.check_max_value_length,
            check_min_value_length=self.check_min_value_length,
            data_type=self.data_type,
            description=self.description,
            enable_regular_expression=self.enable_regular_expression,
            id=self.id,
            is_base64=self.is_base64,
            is_cookie=self.is_cookie,
            is_header=self.is_header,
            json=self.json,
            level=self.level,
            mandatory=self.mandatory,
            metachars_on_parameter_value_check=self.metachars_on_parameter_value_check,
            name=self.name,
            parameter_location=self.parameter_location,
            perform_staging=self.perform_staging,
            sensitive_parameter=self.sensitive_parameter,
            signature_overrides_disables=self.signature_overrides_disables,
            type=self.type,
            url=self.url,
            value_type=self.value_type)


def get_waf_entity_parameter(allow_empty_type: Optional[bool] = None,
                             allow_repeated_parameter_name: Optional[bool] = None,
                             attack_signatures_check: Optional[bool] = None,
                             check_max_value_length: Optional[bool] = None,
                             check_min_value_length: Optional[bool] = None,
                             data_type: Optional[str] = None,
                             description: Optional[str] = None,
                             enable_regular_expression: Optional[bool] = None,
                             is_base64: Optional[bool] = None,
                             is_cookie: Optional[bool] = None,
                             is_header: Optional[bool] = None,
                             json: Optional[str] = None,
                             level: Optional[str] = None,
                             mandatory: Optional[bool] = None,
                             metachars_on_parameter_value_check: Optional[bool] = None,
                             name: Optional[str] = None,
                             parameter_location: Optional[str] = None,
                             perform_staging: Optional[bool] = None,
                             sensitive_parameter: Optional[bool] = None,
                             signature_overrides_disables: Optional[Sequence[int]] = None,
                             type: Optional[str] = None,
                             url: Optional[pulumi.InputType['GetWafEntityParameterUrlArgs']] = None,
                             value_type: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWafEntityParameterResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['allowEmptyType'] = allow_empty_type
    __args__['allowRepeatedParameterName'] = allow_repeated_parameter_name
    __args__['attackSignaturesCheck'] = attack_signatures_check
    __args__['checkMaxValueLength'] = check_max_value_length
    __args__['checkMinValueLength'] = check_min_value_length
    __args__['dataType'] = data_type
    __args__['description'] = description
    __args__['enableRegularExpression'] = enable_regular_expression
    __args__['isBase64'] = is_base64
    __args__['isCookie'] = is_cookie
    __args__['isHeader'] = is_header
    __args__['json'] = json
    __args__['level'] = level
    __args__['mandatory'] = mandatory
    __args__['metacharsOnParameterValueCheck'] = metachars_on_parameter_value_check
    __args__['name'] = name
    __args__['parameterLocation'] = parameter_location
    __args__['performStaging'] = perform_staging
    __args__['sensitiveParameter'] = sensitive_parameter
    __args__['signatureOverridesDisables'] = signature_overrides_disables
    __args__['type'] = type
    __args__['url'] = url
    __args__['valueType'] = value_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:ssl/getWafEntityParameter:getWafEntityParameter', __args__, opts=opts, typ=GetWafEntityParameterResult).value

    return AwaitableGetWafEntityParameterResult(
        allow_empty_type=pulumi.get(__ret__, 'allow_empty_type'),
        allow_repeated_parameter_name=pulumi.get(__ret__, 'allow_repeated_parameter_name'),
        attack_signatures_check=pulumi.get(__ret__, 'attack_signatures_check'),
        check_max_value_length=pulumi.get(__ret__, 'check_max_value_length'),
        check_min_value_length=pulumi.get(__ret__, 'check_min_value_length'),
        data_type=pulumi.get(__ret__, 'data_type'),
        description=pulumi.get(__ret__, 'description'),
        enable_regular_expression=pulumi.get(__ret__, 'enable_regular_expression'),
        id=pulumi.get(__ret__, 'id'),
        is_base64=pulumi.get(__ret__, 'is_base64'),
        is_cookie=pulumi.get(__ret__, 'is_cookie'),
        is_header=pulumi.get(__ret__, 'is_header'),
        json=pulumi.get(__ret__, 'json'),
        level=pulumi.get(__ret__, 'level'),
        mandatory=pulumi.get(__ret__, 'mandatory'),
        metachars_on_parameter_value_check=pulumi.get(__ret__, 'metachars_on_parameter_value_check'),
        name=pulumi.get(__ret__, 'name'),
        parameter_location=pulumi.get(__ret__, 'parameter_location'),
        perform_staging=pulumi.get(__ret__, 'perform_staging'),
        sensitive_parameter=pulumi.get(__ret__, 'sensitive_parameter'),
        signature_overrides_disables=pulumi.get(__ret__, 'signature_overrides_disables'),
        type=pulumi.get(__ret__, 'type'),
        url=pulumi.get(__ret__, 'url'),
        value_type=pulumi.get(__ret__, 'value_type'))


@_utilities.lift_output_func(get_waf_entity_parameter)
def get_waf_entity_parameter_output(allow_empty_type: Optional[pulumi.Input[Optional[bool]]] = None,
                                    allow_repeated_parameter_name: Optional[pulumi.Input[Optional[bool]]] = None,
                                    attack_signatures_check: Optional[pulumi.Input[Optional[bool]]] = None,
                                    check_max_value_length: Optional[pulumi.Input[Optional[bool]]] = None,
                                    check_min_value_length: Optional[pulumi.Input[Optional[bool]]] = None,
                                    data_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    description: Optional[pulumi.Input[Optional[str]]] = None,
                                    enable_regular_expression: Optional[pulumi.Input[Optional[bool]]] = None,
                                    is_base64: Optional[pulumi.Input[Optional[bool]]] = None,
                                    is_cookie: Optional[pulumi.Input[Optional[bool]]] = None,
                                    is_header: Optional[pulumi.Input[Optional[bool]]] = None,
                                    json: Optional[pulumi.Input[Optional[str]]] = None,
                                    level: Optional[pulumi.Input[Optional[str]]] = None,
                                    mandatory: Optional[pulumi.Input[Optional[bool]]] = None,
                                    metachars_on_parameter_value_check: Optional[pulumi.Input[Optional[bool]]] = None,
                                    name: Optional[pulumi.Input[str]] = None,
                                    parameter_location: Optional[pulumi.Input[Optional[str]]] = None,
                                    perform_staging: Optional[pulumi.Input[Optional[bool]]] = None,
                                    sensitive_parameter: Optional[pulumi.Input[Optional[bool]]] = None,
                                    signature_overrides_disables: Optional[pulumi.Input[Optional[Sequence[int]]]] = None,
                                    type: Optional[pulumi.Input[Optional[str]]] = None,
                                    url: Optional[pulumi.Input[Optional[pulumi.InputType['GetWafEntityParameterUrlArgs']]]] = None,
                                    value_type: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWafEntityParameterResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
