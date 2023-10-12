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
    'GetWafPbSuggestionsResult',
    'AwaitableGetWafPbSuggestionsResult',
    'get_waf_pb_suggestions',
    'get_waf_pb_suggestions_output',
]

@pulumi.output_type
class GetWafPbSuggestionsResult:
    """
    A collection of values returned by getWafPbSuggestions.
    """
    def __init__(__self__, id=None, json=None, minimum_learning_score=None, partition=None, policy_id=None, policy_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if json and not isinstance(json, str):
            raise TypeError("Expected argument 'json' to be a str")
        pulumi.set(__self__, "json", json)
        if minimum_learning_score and not isinstance(minimum_learning_score, int):
            raise TypeError("Expected argument 'minimum_learning_score' to be a int")
        pulumi.set(__self__, "minimum_learning_score", minimum_learning_score)
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        pulumi.set(__self__, "partition", partition)
        if policy_id and not isinstance(policy_id, str):
            raise TypeError("Expected argument 'policy_id' to be a str")
        pulumi.set(__self__, "policy_id", policy_id)
        if policy_name and not isinstance(policy_name, str):
            raise TypeError("Expected argument 'policy_name' to be a str")
        pulumi.set(__self__, "policy_name", policy_name)

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
        Json string representing exported PB suggestions ready to be used in WAF policy declaration
        """
        return pulumi.get(self, "json")

    @property
    @pulumi.getter(name="minimumLearningScore")
    def minimum_learning_score(self) -> int:
        return pulumi.get(self, "minimum_learning_score")

    @property
    @pulumi.getter
    def partition(self) -> str:
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> str:
        """
        System generated id of the WAF policy
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> str:
        return pulumi.get(self, "policy_name")


class AwaitableGetWafPbSuggestionsResult(GetWafPbSuggestionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWafPbSuggestionsResult(
            id=self.id,
            json=self.json,
            minimum_learning_score=self.minimum_learning_score,
            partition=self.partition,
            policy_id=self.policy_id,
            policy_name=self.policy_name)


def get_waf_pb_suggestions(minimum_learning_score: Optional[int] = None,
                           partition: Optional[str] = None,
                           policy_id: Optional[str] = None,
                           policy_name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWafPbSuggestionsResult:
    """
    Use this data source (`ssl_get_waf_pb_suggestions`) to export PB suggestions from an existing WAF policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    p_bwaf1 = f5bigip.ssl.get_waf_pb_suggestions(minimum_learning_score=20,
        partition="Common",
        policy_name="protect_me_policy")
    ```


    :param int minimum_learning_score: The minimum learning score for suggestions.
    :param str partition: Partition on which WAF policy is located.
    :param str policy_id: System generated id of the WAF policy
    :param str policy_name: WAF policy name from which PB suggestions should be exported.
    """
    __args__ = dict()
    __args__['minimumLearningScore'] = minimum_learning_score
    __args__['partition'] = partition
    __args__['policyId'] = policy_id
    __args__['policyName'] = policy_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:ssl/getWafPbSuggestions:getWafPbSuggestions', __args__, opts=opts, typ=GetWafPbSuggestionsResult).value

    return AwaitableGetWafPbSuggestionsResult(
        id=pulumi.get(__ret__, 'id'),
        json=pulumi.get(__ret__, 'json'),
        minimum_learning_score=pulumi.get(__ret__, 'minimum_learning_score'),
        partition=pulumi.get(__ret__, 'partition'),
        policy_id=pulumi.get(__ret__, 'policy_id'),
        policy_name=pulumi.get(__ret__, 'policy_name'))


@_utilities.lift_output_func(get_waf_pb_suggestions)
def get_waf_pb_suggestions_output(minimum_learning_score: Optional[pulumi.Input[int]] = None,
                                  partition: Optional[pulumi.Input[str]] = None,
                                  policy_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  policy_name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWafPbSuggestionsResult]:
    """
    Use this data source (`ssl_get_waf_pb_suggestions`) to export PB suggestions from an existing WAF policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    p_bwaf1 = f5bigip.ssl.get_waf_pb_suggestions(minimum_learning_score=20,
        partition="Common",
        policy_name="protect_me_policy")
    ```


    :param int minimum_learning_score: The minimum learning score for suggestions.
    :param str partition: Partition on which WAF policy is located.
    :param str policy_id: System generated id of the WAF policy
    :param str policy_name: WAF policy name from which PB suggestions should be exported.
    """
    ...
