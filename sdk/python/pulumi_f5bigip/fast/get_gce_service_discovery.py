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
    'GetGceServiceDiscoveryResult',
    'AwaitableGetGceServiceDiscoveryResult',
    'get_gce_service_discovery',
    'get_gce_service_discovery_output',
]

@pulumi.output_type
class GetGceServiceDiscoveryResult:
    """
    A collection of values returned by getGceServiceDiscovery.
    """
    def __init__(__self__, address_realm=None, credential_update=None, encoded_credentials=None, gce_sd_json=None, id=None, minimum_monitors=None, port=None, project_id=None, region=None, tag_key=None, tag_value=None, type=None, undetectable_action=None, update_interval=None):
        if address_realm and not isinstance(address_realm, str):
            raise TypeError("Expected argument 'address_realm' to be a str")
        pulumi.set(__self__, "address_realm", address_realm)
        if credential_update and not isinstance(credential_update, bool):
            raise TypeError("Expected argument 'credential_update' to be a bool")
        pulumi.set(__self__, "credential_update", credential_update)
        if encoded_credentials and not isinstance(encoded_credentials, str):
            raise TypeError("Expected argument 'encoded_credentials' to be a str")
        pulumi.set(__self__, "encoded_credentials", encoded_credentials)
        if gce_sd_json and not isinstance(gce_sd_json, str):
            raise TypeError("Expected argument 'gce_sd_json' to be a str")
        pulumi.set(__self__, "gce_sd_json", gce_sd_json)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if minimum_monitors and not isinstance(minimum_monitors, str):
            raise TypeError("Expected argument 'minimum_monitors' to be a str")
        pulumi.set(__self__, "minimum_monitors", minimum_monitors)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
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

    @_builtins.property
    @pulumi.getter(name="addressRealm")
    def address_realm(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "address_realm")

    @_builtins.property
    @pulumi.getter(name="credentialUpdate")
    def credential_update(self) -> Optional[_builtins.bool]:
        return pulumi.get(self, "credential_update")

    @_builtins.property
    @pulumi.getter(name="encodedCredentials")
    def encoded_credentials(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "encoded_credentials")

    @_builtins.property
    @pulumi.getter(name="gceSdJson")
    def gce_sd_json(self) -> _builtins.str:
        """
        The JSON for GCE service discovery block.
        """
        return pulumi.get(self, "gce_sd_json")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="minimumMonitors")
    def minimum_monitors(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "minimum_monitors")

    @_builtins.property
    @pulumi.getter
    def port(self) -> Optional[_builtins.int]:
        return pulumi.get(self, "port")

    @_builtins.property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "project_id")

    @_builtins.property
    @pulumi.getter
    def region(self) -> _builtins.str:
        return pulumi.get(self, "region")

    @_builtins.property
    @pulumi.getter(name="tagKey")
    def tag_key(self) -> _builtins.str:
        return pulumi.get(self, "tag_key")

    @_builtins.property
    @pulumi.getter(name="tagValue")
    def tag_value(self) -> _builtins.str:
        return pulumi.get(self, "tag_value")

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "type")

    @_builtins.property
    @pulumi.getter(name="undetectableAction")
    def undetectable_action(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "undetectable_action")

    @_builtins.property
    @pulumi.getter(name="updateInterval")
    def update_interval(self) -> Optional[_builtins.str]:
        return pulumi.get(self, "update_interval")


class AwaitableGetGceServiceDiscoveryResult(GetGceServiceDiscoveryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGceServiceDiscoveryResult(
            address_realm=self.address_realm,
            credential_update=self.credential_update,
            encoded_credentials=self.encoded_credentials,
            gce_sd_json=self.gce_sd_json,
            id=self.id,
            minimum_monitors=self.minimum_monitors,
            port=self.port,
            project_id=self.project_id,
            region=self.region,
            tag_key=self.tag_key,
            tag_value=self.tag_value,
            type=self.type,
            undetectable_action=self.undetectable_action,
            update_interval=self.update_interval)


def get_gce_service_discovery(address_realm: Optional[_builtins.str] = None,
                              credential_update: Optional[_builtins.bool] = None,
                              encoded_credentials: Optional[_builtins.str] = None,
                              minimum_monitors: Optional[_builtins.str] = None,
                              port: Optional[_builtins.int] = None,
                              project_id: Optional[_builtins.str] = None,
                              region: Optional[_builtins.str] = None,
                              tag_key: Optional[_builtins.str] = None,
                              tag_value: Optional[_builtins.str] = None,
                              type: Optional[_builtins.str] = None,
                              undetectable_action: Optional[_builtins.str] = None,
                              update_interval: Optional[_builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGceServiceDiscoveryResult:
    """
    Use this data source (`fast_get_gce_service_discovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    tc3 = f5bigip.fast.get_gce_service_discovery(tag_key="testgcetag",
        tag_value="testgcevalue",
        region="testgceregion")
    ```


    :param _builtins.str address_realm: Specifies whether to look for public or private IP addresses,default `private`.
    :param _builtins.bool credential_update: Specifies whether you are updating your credentials,default `false`.
    :param _builtins.str encoded_credentials: Base 64 encoded service account credentials JSON.
    :param _builtins.str minimum_monitors: Member is down when fewer than minimum monitors report it healthy.
    :param _builtins.int port: Port to be used for AWS service discovery,default `80`.
    :param _builtins.str project_id: For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
    :param _builtins.str region: GCE region in which ADC is running.
    :param _builtins.str tag_key: The tag key associated with the node to add to this pool.
    :param _builtins.str tag_value: The tag value associated with the node to add to this pool.
    :param _builtins.str undetectable_action: Action to take when node cannot be detected,default `remove`.
    :param _builtins.str update_interval: Update interval for service discovery.
    """
    __args__ = dict()
    __args__['addressRealm'] = address_realm
    __args__['credentialUpdate'] = credential_update
    __args__['encodedCredentials'] = encoded_credentials
    __args__['minimumMonitors'] = minimum_monitors
    __args__['port'] = port
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['tagKey'] = tag_key
    __args__['tagValue'] = tag_value
    __args__['type'] = type
    __args__['undetectableAction'] = undetectable_action
    __args__['updateInterval'] = update_interval
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery', __args__, opts=opts, typ=GetGceServiceDiscoveryResult).value

    return AwaitableGetGceServiceDiscoveryResult(
        address_realm=pulumi.get(__ret__, 'address_realm'),
        credential_update=pulumi.get(__ret__, 'credential_update'),
        encoded_credentials=pulumi.get(__ret__, 'encoded_credentials'),
        gce_sd_json=pulumi.get(__ret__, 'gce_sd_json'),
        id=pulumi.get(__ret__, 'id'),
        minimum_monitors=pulumi.get(__ret__, 'minimum_monitors'),
        port=pulumi.get(__ret__, 'port'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        tag_key=pulumi.get(__ret__, 'tag_key'),
        tag_value=pulumi.get(__ret__, 'tag_value'),
        type=pulumi.get(__ret__, 'type'),
        undetectable_action=pulumi.get(__ret__, 'undetectable_action'),
        update_interval=pulumi.get(__ret__, 'update_interval'))
def get_gce_service_discovery_output(address_realm: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                     credential_update: Optional[pulumi.Input[Optional[_builtins.bool]]] = None,
                                     encoded_credentials: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                     minimum_monitors: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                     port: Optional[pulumi.Input[Optional[_builtins.int]]] = None,
                                     project_id: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                     region: Optional[pulumi.Input[_builtins.str]] = None,
                                     tag_key: Optional[pulumi.Input[_builtins.str]] = None,
                                     tag_value: Optional[pulumi.Input[_builtins.str]] = None,
                                     type: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                     undetectable_action: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                     update_interval: Optional[pulumi.Input[Optional[_builtins.str]]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGceServiceDiscoveryResult]:
    """
    Use this data source (`fast_get_gce_service_discovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    tc3 = f5bigip.fast.get_gce_service_discovery(tag_key="testgcetag",
        tag_value="testgcevalue",
        region="testgceregion")
    ```


    :param _builtins.str address_realm: Specifies whether to look for public or private IP addresses,default `private`.
    :param _builtins.bool credential_update: Specifies whether you are updating your credentials,default `false`.
    :param _builtins.str encoded_credentials: Base 64 encoded service account credentials JSON.
    :param _builtins.str minimum_monitors: Member is down when fewer than minimum monitors report it healthy.
    :param _builtins.int port: Port to be used for AWS service discovery,default `80`.
    :param _builtins.str project_id: For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
    :param _builtins.str region: GCE region in which ADC is running.
    :param _builtins.str tag_key: The tag key associated with the node to add to this pool.
    :param _builtins.str tag_value: The tag value associated with the node to add to this pool.
    :param _builtins.str undetectable_action: Action to take when node cannot be detected,default `remove`.
    :param _builtins.str update_interval: Update interval for service discovery.
    """
    __args__ = dict()
    __args__['addressRealm'] = address_realm
    __args__['credentialUpdate'] = credential_update
    __args__['encodedCredentials'] = encoded_credentials
    __args__['minimumMonitors'] = minimum_monitors
    __args__['port'] = port
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['tagKey'] = tag_key
    __args__['tagValue'] = tag_value
    __args__['type'] = type
    __args__['undetectableAction'] = undetectable_action
    __args__['updateInterval'] = update_interval
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery', __args__, opts=opts, typ=GetGceServiceDiscoveryResult)
    return __ret__.apply(lambda __response__: GetGceServiceDiscoveryResult(
        address_realm=pulumi.get(__response__, 'address_realm'),
        credential_update=pulumi.get(__response__, 'credential_update'),
        encoded_credentials=pulumi.get(__response__, 'encoded_credentials'),
        gce_sd_json=pulumi.get(__response__, 'gce_sd_json'),
        id=pulumi.get(__response__, 'id'),
        minimum_monitors=pulumi.get(__response__, 'minimum_monitors'),
        port=pulumi.get(__response__, 'port'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        tag_key=pulumi.get(__response__, 'tag_key'),
        tag_value=pulumi.get(__response__, 'tag_value'),
        type=pulumi.get(__response__, 'type'),
        undetectable_action=pulumi.get(__response__, 'undetectable_action'),
        update_interval=pulumi.get(__response__, 'update_interval')))
