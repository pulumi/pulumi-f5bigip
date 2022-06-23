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
    'GetWafSignaturesResult',
    'AwaitableGetWafSignaturesResult',
    'get_waf_signatures',
    'get_waf_signatures_output',
]

@pulumi.output_type
class GetWafSignaturesResult:
    """
    A collection of values returned by getWafSignatures.
    """
    def __init__(__self__, accuracy=None, description=None, id=None, name=None, risk=None, signature_id=None, system_signature_id=None, type=None):
        if accuracy and not isinstance(accuracy, str):
            raise TypeError("Expected argument 'accuracy' to be a str")
        pulumi.set(__self__, "accuracy", accuracy)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if risk and not isinstance(risk, str):
            raise TypeError("Expected argument 'risk' to be a str")
        pulumi.set(__self__, "risk", risk)
        if signature_id and not isinstance(signature_id, int):
            raise TypeError("Expected argument 'signature_id' to be a int")
        pulumi.set(__self__, "signature_id", signature_id)
        if system_signature_id and not isinstance(system_signature_id, str):
            raise TypeError("Expected argument 'system_signature_id' to be a str")
        pulumi.set(__self__, "system_signature_id", system_signature_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def accuracy(self) -> str:
        """
        The relative detection accuracy of the signature.
        """
        return pulumi.get(self, "accuracy")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the signature.
        """
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
    def name(self) -> str:
        """
        Name of the signature as configured on the system.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def risk(self) -> str:
        """
        The relative risk level of the attack that matches this signature.
        """
        return pulumi.get(self, "risk")

    @property
    @pulumi.getter(name="signatureId")
    def signature_id(self) -> int:
        """
        ID of the signature in the database.
        """
        return pulumi.get(self, "signature_id")

    @property
    @pulumi.getter(name="systemSignatureId")
    def system_signature_id(self) -> str:
        """
        System generated ID of the signature.
        """
        return pulumi.get(self, "system_signature_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of the signature.
        """
        return pulumi.get(self, "type")


class AwaitableGetWafSignaturesResult(GetWafSignaturesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWafSignaturesResult(
            accuracy=self.accuracy,
            description=self.description,
            id=self.id,
            name=self.name,
            risk=self.risk,
            signature_id=self.signature_id,
            system_signature_id=self.system_signature_id,
            type=self.type)


def get_waf_signatures(accuracy: Optional[str] = None,
                       description: Optional[str] = None,
                       name: Optional[str] = None,
                       risk: Optional[str] = None,
                       signature_id: Optional[int] = None,
                       system_signature_id: Optional[str] = None,
                       type: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWafSignaturesResult:
    """
    Use this data source (`ssl.get_waf_signatures`) to get the details of attack signatures available on BIG-IP WAF

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    w_afsig1 = f5bigip.ssl.get_waf_signatures(signature_id=200104004)
    ```


    :param str accuracy: The relative detection accuracy of the signature.
    :param str description: Description of the signature.
    :param str name: Name of the signature as configured on the system.
    :param str risk: The relative risk level of the attack that matches this signature.
    :param int signature_id: ID of the signature in the BIG-IP WAF database.
    :param str system_signature_id: System generated ID of the signature.
    :param str type: Type of the signature.
    """
    __args__ = dict()
    __args__['accuracy'] = accuracy
    __args__['description'] = description
    __args__['name'] = name
    __args__['risk'] = risk
    __args__['signatureId'] = signature_id
    __args__['systemSignatureId'] = system_signature_id
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:ssl/getWafSignatures:getWafSignatures', __args__, opts=opts, typ=GetWafSignaturesResult).value

    return AwaitableGetWafSignaturesResult(
        accuracy=__ret__.accuracy,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        risk=__ret__.risk,
        signature_id=__ret__.signature_id,
        system_signature_id=__ret__.system_signature_id,
        type=__ret__.type)


@_utilities.lift_output_func(get_waf_signatures)
def get_waf_signatures_output(accuracy: Optional[pulumi.Input[Optional[str]]] = None,
                              description: Optional[pulumi.Input[Optional[str]]] = None,
                              name: Optional[pulumi.Input[Optional[str]]] = None,
                              risk: Optional[pulumi.Input[Optional[str]]] = None,
                              signature_id: Optional[pulumi.Input[int]] = None,
                              system_signature_id: Optional[pulumi.Input[Optional[str]]] = None,
                              type: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWafSignaturesResult]:
    """
    Use this data source (`ssl.get_waf_signatures`) to get the details of attack signatures available on BIG-IP WAF

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    w_afsig1 = f5bigip.ssl.get_waf_signatures(signature_id=200104004)
    ```


    :param str accuracy: The relative detection accuracy of the signature.
    :param str description: Description of the signature.
    :param str name: Name of the signature as configured on the system.
    :param str risk: The relative risk level of the attack that matches this signature.
    :param int signature_id: ID of the signature in the BIG-IP WAF database.
    :param str system_signature_id: System generated ID of the signature.
    :param str type: Type of the signature.
    """
    ...
