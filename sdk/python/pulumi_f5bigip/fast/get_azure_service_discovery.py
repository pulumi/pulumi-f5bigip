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
    'GetAzureServiceDiscoveryResult',
    'AwaitableGetAzureServiceDiscoveryResult',
    'get_azure_service_discovery',
    'get_azure_service_discovery_output',
]

@pulumi.output_type
class GetAzureServiceDiscoveryResult:
    """
    A collection of values returned by getAzureServiceDiscovery.
    """
    def __init__(__self__, address_realm=None, azure_sd_json=None, credential_update=None, id=None, minimum_monitors=None, port=None, resource_group=None, subscription_id=None, tag_key=None, tag_value=None, type=None, undetectable_action=None, update_interval=None):
        if address_realm and not isinstance(address_realm, str):
            raise TypeError("Expected argument 'address_realm' to be a str")
        pulumi.set(__self__, "address_realm", address_realm)
        if azure_sd_json and not isinstance(azure_sd_json, str):
            raise TypeError("Expected argument 'azure_sd_json' to be a str")
        pulumi.set(__self__, "azure_sd_json", azure_sd_json)
        if credential_update and not isinstance(credential_update, bool):
            raise TypeError("Expected argument 'credential_update' to be a bool")
        pulumi.set(__self__, "credential_update", credential_update)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if minimum_monitors and not isinstance(minimum_monitors, str):
            raise TypeError("Expected argument 'minimum_monitors' to be a str")
        pulumi.set(__self__, "minimum_monitors", minimum_monitors)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if resource_group and not isinstance(resource_group, str):
            raise TypeError("Expected argument 'resource_group' to be a str")
        pulumi.set(__self__, "resource_group", resource_group)
        if subscription_id and not isinstance(subscription_id, str):
            raise TypeError("Expected argument 'subscription_id' to be a str")
        pulumi.set(__self__, "subscription_id", subscription_id)
        if tag_key and not isinstance(tag_key, str):
            raise TypeError("Expected argument 'tag_key' to be a str")
        pulumi.set(__self__, "tag_key", tag_key)
        if tag_value and not isinstance(tag_value, str):
            raise TypeError("Expected argument 'tag_value' to be a str")
        pulumi.set(__self__, "tag_value", tag_value)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if undetectable_action and not isinstance(undetectable_action, str):
            raise TypeError("Expected argument 'undetectable_action' to be a str")
        pulumi.set(__self__, "undetectable_action", undetectable_action)
        if update_interval and not isinstance(update_interval, str):
            raise TypeError("Expected argument 'update_interval' to be a str")
        pulumi.set(__self__, "update_interval", update_interval)

    @property
    @pulumi.getter(name="addressRealm")
    def address_realm(self) -> Optional[str]:
        return pulumi.get(self, "address_realm")

    @property
    @pulumi.getter(name="azureSdJson")
    def azure_sd_json(self) -> str:
        """
        The JSON for Azure service discovery block.
        """
        return pulumi.get(self, "azure_sd_json")

    @property
    @pulumi.getter(name="credentialUpdate")
    def credential_update(self) -> Optional[bool]:
        return pulumi.get(self, "credential_update")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="minimumMonitors")
    def minimum_monitors(self) -> Optional[str]:
        return pulumi.get(self, "minimum_monitors")

    @property
    @pulumi.getter
    def port(self) -> Optional[int]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> str:
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="tagKey")
    def tag_key(self) -> Optional[str]:
        return pulumi.get(self, "tag_key")

    @property
    @pulumi.getter(name="tagValue")
    def tag_value(self) -> Optional[str]:
        return pulumi.get(self, "tag_value")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="undetectableAction")
    def undetectable_action(self) -> Optional[str]:
        return pulumi.get(self, "undetectable_action")

    @property
    @pulumi.getter(name="updateInterval")
    def update_interval(self) -> Optional[str]:
        return pulumi.get(self, "update_interval")


class AwaitableGetAzureServiceDiscoveryResult(GetAzureServiceDiscoveryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAzureServiceDiscoveryResult(
            address_realm=self.address_realm,
            azure_sd_json=self.azure_sd_json,
            credential_update=self.credential_update,
            id=self.id,
            minimum_monitors=self.minimum_monitors,
            port=self.port,
            resource_group=self.resource_group,
            subscription_id=self.subscription_id,
            tag_key=self.tag_key,
            tag_value=self.tag_value,
            type=self.type,
            undetectable_action=self.undetectable_action,
            update_interval=self.update_interval)


def get_azure_service_discovery(address_realm: Optional[str] = None,
                                credential_update: Optional[bool] = None,
                                minimum_monitors: Optional[str] = None,
                                port: Optional[int] = None,
                                resource_group: Optional[str] = None,
                                subscription_id: Optional[str] = None,
                                tag_key: Optional[str] = None,
                                tag_value: Optional[str] = None,
                                type: Optional[str] = None,
                                undetectable_action: Optional[str] = None,
                                update_interval: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAzureServiceDiscoveryResult:
    """
    Use this data source (`fast_get_azure_service_discovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    t_c3 = f5bigip.fast.get_azure_service_discovery(resource_group="testazurerg",
        subscription_id="testazuresid",
        tag_key="testazuretag",
        tag_value="testazurevalue")
    ```


    :param str address_realm: Specifies whether to look for public or private IP addresses,default `private`.
    :param bool credential_update: Specifies whether you are updating your credentials,default `false`.
    :param str minimum_monitors: Member is down when fewer than minimum monitors report it healthy.
    :param int port: Port to be used for Azure service discovery,default `80`.
    :param str resource_group: Azure Resource Group name.
    :param str subscription_id: Azure subscription ID.
    :param str tag_key: The tag key associated with the node to add to this pool.
    :param str tag_value: The tag value associated with the node to add to this pool.
    :param str undetectable_action: Action to take when node cannot be detected,default `remove`.
    :param str update_interval: Update interval for service discovery.
    """
    __args__ = dict()
    __args__['addressRealm'] = address_realm
    __args__['credentialUpdate'] = credential_update
    __args__['minimumMonitors'] = minimum_monitors
    __args__['port'] = port
    __args__['resourceGroup'] = resource_group
    __args__['subscriptionId'] = subscription_id
    __args__['tagKey'] = tag_key
    __args__['tagValue'] = tag_value
    __args__['type'] = type
    __args__['undetectableAction'] = undetectable_action
    __args__['updateInterval'] = update_interval
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:fast/getAzureServiceDiscovery:getAzureServiceDiscovery', __args__, opts=opts, typ=GetAzureServiceDiscoveryResult).value

    return AwaitableGetAzureServiceDiscoveryResult(
        address_realm=__ret__.address_realm,
        azure_sd_json=__ret__.azure_sd_json,
        credential_update=__ret__.credential_update,
        id=__ret__.id,
        minimum_monitors=__ret__.minimum_monitors,
        port=__ret__.port,
        resource_group=__ret__.resource_group,
        subscription_id=__ret__.subscription_id,
        tag_key=__ret__.tag_key,
        tag_value=__ret__.tag_value,
        type=__ret__.type,
        undetectable_action=__ret__.undetectable_action,
        update_interval=__ret__.update_interval)


@_utilities.lift_output_func(get_azure_service_discovery)
def get_azure_service_discovery_output(address_realm: Optional[pulumi.Input[Optional[str]]] = None,
                                       credential_update: Optional[pulumi.Input[Optional[bool]]] = None,
                                       minimum_monitors: Optional[pulumi.Input[Optional[str]]] = None,
                                       port: Optional[pulumi.Input[Optional[int]]] = None,
                                       resource_group: Optional[pulumi.Input[str]] = None,
                                       subscription_id: Optional[pulumi.Input[str]] = None,
                                       tag_key: Optional[pulumi.Input[Optional[str]]] = None,
                                       tag_value: Optional[pulumi.Input[Optional[str]]] = None,
                                       type: Optional[pulumi.Input[Optional[str]]] = None,
                                       undetectable_action: Optional[pulumi.Input[Optional[str]]] = None,
                                       update_interval: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAzureServiceDiscoveryResult]:
    """
    Use this data source (`fast_get_azure_service_discovery`) to get the Azure Service discovery config to be used for `http`/`https` app deployment in FAST.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    t_c3 = f5bigip.fast.get_azure_service_discovery(resource_group="testazurerg",
        subscription_id="testazuresid",
        tag_key="testazuretag",
        tag_value="testazurevalue")
    ```


    :param str address_realm: Specifies whether to look for public or private IP addresses,default `private`.
    :param bool credential_update: Specifies whether you are updating your credentials,default `false`.
    :param str minimum_monitors: Member is down when fewer than minimum monitors report it healthy.
    :param int port: Port to be used for Azure service discovery,default `80`.
    :param str resource_group: Azure Resource Group name.
    :param str subscription_id: Azure subscription ID.
    :param str tag_key: The tag key associated with the node to add to this pool.
    :param str tag_value: The tag value associated with the node to add to this pool.
    :param str undetectable_action: Action to take when node cannot be detected,default `remove`.
    :param str update_interval: Update interval for service discovery.
    """
    ...
