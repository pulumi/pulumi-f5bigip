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
    'GetConsulServiceDiscoveryResult',
    'AwaitableGetConsulServiceDiscoveryResult',
    'get_consul_service_discovery',
    'get_consul_service_discovery_output',
]

@pulumi.output_type
class GetConsulServiceDiscoveryResult:
    """
    A collection of values returned by getConsulServiceDiscovery.
    """
    def __init__(__self__, address_realm=None, consul_sd_json=None, credential_update=None, encoded_token=None, id=None, jmes_path_query=None, minimum_monitors=None, port=None, reject_unauthorized=None, trust_ca=None, type=None, undetectable_action=None, update_interval=None, uri=None):
        if address_realm and not isinstance(address_realm, str):
            raise TypeError("Expected argument 'address_realm' to be a str")
        pulumi.set(__self__, "address_realm", address_realm)
        if consul_sd_json and not isinstance(consul_sd_json, str):
            raise TypeError("Expected argument 'consul_sd_json' to be a str")
        pulumi.set(__self__, "consul_sd_json", consul_sd_json)
        if credential_update and not isinstance(credential_update, bool):
            raise TypeError("Expected argument 'credential_update' to be a bool")
        pulumi.set(__self__, "credential_update", credential_update)
        if encoded_token and not isinstance(encoded_token, str):
            raise TypeError("Expected argument 'encoded_token' to be a str")
        pulumi.set(__self__, "encoded_token", encoded_token)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if jmes_path_query and not isinstance(jmes_path_query, str):
            raise TypeError("Expected argument 'jmes_path_query' to be a str")
        pulumi.set(__self__, "jmes_path_query", jmes_path_query)
        if minimum_monitors and not isinstance(minimum_monitors, str):
            raise TypeError("Expected argument 'minimum_monitors' to be a str")
        pulumi.set(__self__, "minimum_monitors", minimum_monitors)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if reject_unauthorized and not isinstance(reject_unauthorized, bool):
            raise TypeError("Expected argument 'reject_unauthorized' to be a bool")
        pulumi.set(__self__, "reject_unauthorized", reject_unauthorized)
        if trust_ca and not isinstance(trust_ca, str):
            raise TypeError("Expected argument 'trust_ca' to be a str")
        pulumi.set(__self__, "trust_ca", trust_ca)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if undetectable_action and not isinstance(undetectable_action, str):
            raise TypeError("Expected argument 'undetectable_action' to be a str")
        pulumi.set(__self__, "undetectable_action", undetectable_action)
        if update_interval and not isinstance(update_interval, str):
            raise TypeError("Expected argument 'update_interval' to be a str")
        pulumi.set(__self__, "update_interval", update_interval)
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="addressRealm")
    def address_realm(self) -> Optional[str]:
        return pulumi.get(self, "address_realm")

    @property
    @pulumi.getter(name="consulSdJson")
    def consul_sd_json(self) -> str:
        """
        The JSON for Hashicorp Consul service discovery block.
        """
        return pulumi.get(self, "consul_sd_json")

    @property
    @pulumi.getter(name="credentialUpdate")
    def credential_update(self) -> Optional[bool]:
        return pulumi.get(self, "credential_update")

    @property
    @pulumi.getter(name="encodedToken")
    def encoded_token(self) -> Optional[str]:
        return pulumi.get(self, "encoded_token")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="jmesPathQuery")
    def jmes_path_query(self) -> Optional[str]:
        return pulumi.get(self, "jmes_path_query")

    @property
    @pulumi.getter(name="minimumMonitors")
    def minimum_monitors(self) -> Optional[str]:
        return pulumi.get(self, "minimum_monitors")

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="rejectUnauthorized")
    def reject_unauthorized(self) -> Optional[bool]:
        return pulumi.get(self, "reject_unauthorized")

    @property
    @pulumi.getter(name="trustCa")
    def trust_ca(self) -> Optional[str]:
        return pulumi.get(self, "trust_ca")

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

    @property
    @pulumi.getter
    def uri(self) -> str:
        return pulumi.get(self, "uri")


class AwaitableGetConsulServiceDiscoveryResult(GetConsulServiceDiscoveryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConsulServiceDiscoveryResult(
            address_realm=self.address_realm,
            consul_sd_json=self.consul_sd_json,
            credential_update=self.credential_update,
            encoded_token=self.encoded_token,
            id=self.id,
            jmes_path_query=self.jmes_path_query,
            minimum_monitors=self.minimum_monitors,
            port=self.port,
            reject_unauthorized=self.reject_unauthorized,
            trust_ca=self.trust_ca,
            type=self.type,
            undetectable_action=self.undetectable_action,
            update_interval=self.update_interval,
            uri=self.uri)


def get_consul_service_discovery(address_realm: Optional[str] = None,
                                 credential_update: Optional[bool] = None,
                                 encoded_token: Optional[str] = None,
                                 jmes_path_query: Optional[str] = None,
                                 minimum_monitors: Optional[str] = None,
                                 port: Optional[int] = None,
                                 reject_unauthorized: Optional[bool] = None,
                                 trust_ca: Optional[str] = None,
                                 type: Optional[str] = None,
                                 undetectable_action: Optional[str] = None,
                                 update_interval: Optional[str] = None,
                                 uri: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConsulServiceDiscoveryResult:
    """
    Use this data source (`fast_get_consul_service_discovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.


    :param str address_realm: Specifies whether to look for public or private IP addresses,default `private`.
    :param bool credential_update: Specifies whether you are updating your credentials,default `false`.
    :param str encoded_token: Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
    :param str jmes_path_query: Custom JMESPath Query.
    :param str minimum_monitors: Member is down when fewer than minimum monitors report it healthy.
    :param int port: Port to be used for AWS service discovery,default `80`.
    :param bool reject_unauthorized: If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
    :param str trust_ca: CA Bundle to validate server certificates.
    :param str undetectable_action: Action to take when node cannot be detected,default `remove`.
    :param str update_interval: Update interval for service discovery.
    :param str uri: The location of the node data.
    """
    __args__ = dict()
    __args__['addressRealm'] = address_realm
    __args__['credentialUpdate'] = credential_update
    __args__['encodedToken'] = encoded_token
    __args__['jmesPathQuery'] = jmes_path_query
    __args__['minimumMonitors'] = minimum_monitors
    __args__['port'] = port
    __args__['rejectUnauthorized'] = reject_unauthorized
    __args__['trustCa'] = trust_ca
    __args__['type'] = type
    __args__['undetectableAction'] = undetectable_action
    __args__['updateInterval'] = update_interval
    __args__['uri'] = uri
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery', __args__, opts=opts, typ=GetConsulServiceDiscoveryResult).value

    return AwaitableGetConsulServiceDiscoveryResult(
        address_realm=pulumi.get(__ret__, 'address_realm'),
        consul_sd_json=pulumi.get(__ret__, 'consul_sd_json'),
        credential_update=pulumi.get(__ret__, 'credential_update'),
        encoded_token=pulumi.get(__ret__, 'encoded_token'),
        id=pulumi.get(__ret__, 'id'),
        jmes_path_query=pulumi.get(__ret__, 'jmes_path_query'),
        minimum_monitors=pulumi.get(__ret__, 'minimum_monitors'),
        port=pulumi.get(__ret__, 'port'),
        reject_unauthorized=pulumi.get(__ret__, 'reject_unauthorized'),
        trust_ca=pulumi.get(__ret__, 'trust_ca'),
        type=pulumi.get(__ret__, 'type'),
        undetectable_action=pulumi.get(__ret__, 'undetectable_action'),
        update_interval=pulumi.get(__ret__, 'update_interval'),
        uri=pulumi.get(__ret__, 'uri'))


@_utilities.lift_output_func(get_consul_service_discovery)
def get_consul_service_discovery_output(address_realm: Optional[pulumi.Input[Optional[str]]] = None,
                                        credential_update: Optional[pulumi.Input[Optional[bool]]] = None,
                                        encoded_token: Optional[pulumi.Input[Optional[str]]] = None,
                                        jmes_path_query: Optional[pulumi.Input[Optional[str]]] = None,
                                        minimum_monitors: Optional[pulumi.Input[Optional[str]]] = None,
                                        port: Optional[pulumi.Input[int]] = None,
                                        reject_unauthorized: Optional[pulumi.Input[Optional[bool]]] = None,
                                        trust_ca: Optional[pulumi.Input[Optional[str]]] = None,
                                        type: Optional[pulumi.Input[Optional[str]]] = None,
                                        undetectable_action: Optional[pulumi.Input[Optional[str]]] = None,
                                        update_interval: Optional[pulumi.Input[Optional[str]]] = None,
                                        uri: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConsulServiceDiscoveryResult]:
    """
    Use this data source (`fast_get_consul_service_discovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.


    :param str address_realm: Specifies whether to look for public or private IP addresses,default `private`.
    :param bool credential_update: Specifies whether you are updating your credentials,default `false`.
    :param str encoded_token: Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
    :param str jmes_path_query: Custom JMESPath Query.
    :param str minimum_monitors: Member is down when fewer than minimum monitors report it healthy.
    :param int port: Port to be used for AWS service discovery,default `80`.
    :param bool reject_unauthorized: If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
    :param str trust_ca: CA Bundle to validate server certificates.
    :param str undetectable_action: Action to take when node cannot be detected,default `remove`.
    :param str update_interval: Update interval for service discovery.
    :param str uri: The location of the node data.
    """
    ...
