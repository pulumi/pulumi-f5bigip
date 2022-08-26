# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'EventServiceDiscoveryNodeArgs',
    'FastHttpAppFastCreateMonitorArgs',
    'FastHttpAppFastCreatePoolMemberArgs',
    'FastHttpAppVirtualServerArgs',
    'FastHttpsAppCreateTlsServerProfileArgs',
    'FastHttpsAppFastCreateMonitorArgs',
    'FastHttpsAppFastCreatePoolMemberArgs',
    'FastHttpsAppVirtualServerArgs',
    'FastTcpAppFastCreateMonitorArgs',
    'FastTcpAppFastCreatePoolMemberArgs',
    'FastTcpAppVirtualServerArgs',
    'WafPolicyFileTypeArgs',
    'WafPolicyGraphqlProfileArgs',
    'WafPolicyPolicyBuilderArgs',
    'WafPolicySignaturesSettingArgs',
]

@pulumi.input_type
class EventServiceDiscoveryNodeArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        if id is not None:
            pulumi.set(__self__, "id", id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class FastHttpAppFastCreateMonitorArgs:
    def __init__(__self__, *,
                 interval: Optional[pulumi.Input[int]] = None,
                 monitor_auth: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 response: Optional[pulumi.Input[str]] = None,
                 send_string: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] interval: Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        :param pulumi.Input[bool] monitor_auth: set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
        :param pulumi.Input[str] password: password for web access on FAST-Generated Pool Monitor.
        :param pulumi.Input[str] response: The presence of this string anywhere in the HTTP response implies availability.
        :param pulumi.Input[str] send_string: Specify data to be sent during each health check for FAST-Generated Pool Monitor.
        :param pulumi.Input[str] username: username for web access on FAST-Generated Pool Monitor.
        """
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if monitor_auth is not None:
            pulumi.set(__self__, "monitor_auth", monitor_auth)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if response is not None:
            pulumi.set(__self__, "response", response)
        if send_string is not None:
            pulumi.set(__self__, "send_string", send_string)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter(name="monitorAuth")
    def monitor_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
        """
        return pulumi.get(self, "monitor_auth")

    @monitor_auth.setter
    def monitor_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "monitor_auth", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        password for web access on FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def response(self) -> Optional[pulumi.Input[str]]:
        """
        The presence of this string anywhere in the HTTP response implies availability.
        """
        return pulumi.get(self, "response")

    @response.setter
    def response(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response", value)

    @property
    @pulumi.getter(name="sendString")
    def send_string(self) -> Optional[pulumi.Input[str]]:
        """
        Specify data to be sent during each health check for FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "send_string")

    @send_string.setter
    def send_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "send_string", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        username for web access on FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class FastHttpAppFastCreatePoolMemberArgs:
    def __init__(__self__, *,
                 addresses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 share_nodes: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] addresses: List of server address to be used for FAST-Generated Pool.
        :param pulumi.Input[int] connection_limit: connectionLimit value to be used for FAST-Generated Pool.
        :param pulumi.Input[int] port: port number of serviceport to be used for FAST-Generated Pool.
        :param pulumi.Input[int] priority_group: priorityGroup value to be used for FAST-Generated Pool.
        :param pulumi.Input[bool] share_nodes: shareNodes value to be used for FAST-Generated Pool.
        """
        pulumi.set(__self__, "addresses", addresses)
        if connection_limit is not None:
            pulumi.set(__self__, "connection_limit", connection_limit)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority_group is not None:
            pulumi.set(__self__, "priority_group", priority_group)
        if share_nodes is not None:
            pulumi.set(__self__, "share_nodes", share_nodes)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of server address to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[int]]:
        """
        connectionLimit value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        port number of serviceport to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="priorityGroup")
    def priority_group(self) -> Optional[pulumi.Input[int]]:
        """
        priorityGroup value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "priority_group")

    @priority_group.setter
    def priority_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority_group", value)

    @property
    @pulumi.getter(name="shareNodes")
    def share_nodes(self) -> Optional[pulumi.Input[bool]]:
        """
        shareNodes value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "share_nodes")

    @share_nodes.setter
    def share_nodes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "share_nodes", value)


@pulumi.input_type
class FastHttpAppVirtualServerArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 port: pulumi.Input[int]):
        """
        :param pulumi.Input[str] ip: IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
        :param pulumi.Input[int] port: -(Optional , `int`) Port number to used for accessing virtual server/application
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        -(Optional , `int`) Port number to used for accessing virtual server/application
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class FastHttpsAppCreateTlsServerProfileArgs:
    def __init__(__self__, *,
                 tls_cert_name: pulumi.Input[str],
                 tls_key_name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] tls_cert_name: Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.
        :param pulumi.Input[str] tls_key_name: Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile.
        """
        pulumi.set(__self__, "tls_cert_name", tls_cert_name)
        pulumi.set(__self__, "tls_key_name", tls_key_name)

    @property
    @pulumi.getter(name="tlsCertName")
    def tls_cert_name(self) -> pulumi.Input[str]:
        """
        Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.
        """
        return pulumi.get(self, "tls_cert_name")

    @tls_cert_name.setter
    def tls_cert_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "tls_cert_name", value)

    @property
    @pulumi.getter(name="tlsKeyName")
    def tls_key_name(self) -> pulumi.Input[str]:
        """
        Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile.
        """
        return pulumi.get(self, "tls_key_name")

    @tls_key_name.setter
    def tls_key_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "tls_key_name", value)


@pulumi.input_type
class FastHttpsAppFastCreateMonitorArgs:
    def __init__(__self__, *,
                 interval: Optional[pulumi.Input[int]] = None,
                 monitor_auth: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 response: Optional[pulumi.Input[str]] = None,
                 send_string: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] interval: Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        :param pulumi.Input[bool] monitor_auth: set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
        :param pulumi.Input[str] password: password for web access on FAST-Generated Pool Monitor.
        :param pulumi.Input[str] response: The presence of this string anywhere in the HTTP response implies availability.
        :param pulumi.Input[str] send_string: Specify data to be sent during each health check for FAST-Generated Pool Monitor.
        :param pulumi.Input[str] username: username for web access on FAST-Generated Pool Monitor.
        """
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if monitor_auth is not None:
            pulumi.set(__self__, "monitor_auth", monitor_auth)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if response is not None:
            pulumi.set(__self__, "response", response)
        if send_string is not None:
            pulumi.set(__self__, "send_string", send_string)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter(name="monitorAuth")
    def monitor_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
        """
        return pulumi.get(self, "monitor_auth")

    @monitor_auth.setter
    def monitor_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "monitor_auth", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        password for web access on FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def response(self) -> Optional[pulumi.Input[str]]:
        """
        The presence of this string anywhere in the HTTP response implies availability.
        """
        return pulumi.get(self, "response")

    @response.setter
    def response(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response", value)

    @property
    @pulumi.getter(name="sendString")
    def send_string(self) -> Optional[pulumi.Input[str]]:
        """
        Specify data to be sent during each health check for FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "send_string")

    @send_string.setter
    def send_string(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "send_string", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        username for web access on FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class FastHttpsAppFastCreatePoolMemberArgs:
    def __init__(__self__, *,
                 addresses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 share_nodes: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] addresses: List of server address to be used for FAST-Generated Pool.
        :param pulumi.Input[int] connection_limit: connectionLimit value to be used for FAST-Generated Pool.
        :param pulumi.Input[int] port: port number of serviceport to be used for FAST-Generated Pool.
        :param pulumi.Input[int] priority_group: priorityGroup value to be used for FAST-Generated Pool.
        :param pulumi.Input[bool] share_nodes: shareNodes value to be used for FAST-Generated Pool.
        """
        pulumi.set(__self__, "addresses", addresses)
        if connection_limit is not None:
            pulumi.set(__self__, "connection_limit", connection_limit)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority_group is not None:
            pulumi.set(__self__, "priority_group", priority_group)
        if share_nodes is not None:
            pulumi.set(__self__, "share_nodes", share_nodes)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of server address to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[int]]:
        """
        connectionLimit value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        port number of serviceport to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="priorityGroup")
    def priority_group(self) -> Optional[pulumi.Input[int]]:
        """
        priorityGroup value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "priority_group")

    @priority_group.setter
    def priority_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority_group", value)

    @property
    @pulumi.getter(name="shareNodes")
    def share_nodes(self) -> Optional[pulumi.Input[bool]]:
        """
        shareNodes value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "share_nodes")

    @share_nodes.setter
    def share_nodes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "share_nodes", value)


@pulumi.input_type
class FastHttpsAppVirtualServerArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 port: pulumi.Input[int]):
        """
        :param pulumi.Input[str] ip: IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
        :param pulumi.Input[int] port: -(Optional , `int`) Port number to used for accessing virtual server/application
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        -(Optional , `int`) Port number to used for accessing virtual server/application
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class FastTcpAppFastCreateMonitorArgs:
    def __init__(__self__, *,
                 interval: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[int] interval: Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        """
        if interval is not None:
            pulumi.set(__self__, "interval", interval)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)


@pulumi.input_type
class FastTcpAppFastCreatePoolMemberArgs:
    def __init__(__self__, *,
                 addresses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 connection_limit: Optional[pulumi.Input[int]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 priority_group: Optional[pulumi.Input[int]] = None,
                 share_nodes: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] addresses: List of server address to be used for FAST-Generated Pool.
        :param pulumi.Input[int] connection_limit: connectionLimit value to be used for FAST-Generated Pool.
        :param pulumi.Input[int] port: port number of serviceport to be used for FAST-Generated Pool.
        :param pulumi.Input[int] priority_group: priorityGroup value to be used for FAST-Generated Pool.
        :param pulumi.Input[bool] share_nodes: shareNodes value to be used for FAST-Generated Pool.
        """
        pulumi.set(__self__, "addresses", addresses)
        if connection_limit is not None:
            pulumi.set(__self__, "connection_limit", connection_limit)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if priority_group is not None:
            pulumi.set(__self__, "priority_group", priority_group)
        if share_nodes is not None:
            pulumi.set(__self__, "share_nodes", share_nodes)

    @property
    @pulumi.getter
    def addresses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of server address to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> Optional[pulumi.Input[int]]:
        """
        connectionLimit value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "connection_limit")

    @connection_limit.setter
    def connection_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "connection_limit", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        port number of serviceport to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="priorityGroup")
    def priority_group(self) -> Optional[pulumi.Input[int]]:
        """
        priorityGroup value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "priority_group")

    @priority_group.setter
    def priority_group(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority_group", value)

    @property
    @pulumi.getter(name="shareNodes")
    def share_nodes(self) -> Optional[pulumi.Input[bool]]:
        """
        shareNodes value to be used for FAST-Generated Pool.
        """
        return pulumi.get(self, "share_nodes")

    @share_nodes.setter
    def share_nodes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "share_nodes", value)


@pulumi.input_type
class FastTcpAppVirtualServerArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 port: pulumi.Input[int]):
        """
        :param pulumi.Input[str] ip: IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
        :param pulumi.Input[int] port: -(Optional , `int`) Port number to used for accessing virtual server/application
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        -(Optional , `int`) Port number to used for accessing virtual server/application
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class WafPolicyFileTypeArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Specifies the file type name as appearing in the URL extension.
        :param pulumi.Input[str] type: Determines the type of the name attribute. Only when setting the type to `wildcard` will the special wildcard characters in the name be interpreted as such
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the file type name as appearing in the URL extension.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Determines the type of the name attribute. Only when setting the type to `wildcard` will the special wildcard characters in the name be interpreted as such
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class WafPolicyGraphqlProfileArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class WafPolicyPolicyBuilderArgs:
    def __init__(__self__, *,
                 learning_mode: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] learning_mode: learning mode setting for policy-builder, possible options: [`automatic`,`disabled`, `manual`]
        """
        if learning_mode is not None:
            pulumi.set(__self__, "learning_mode", learning_mode)

    @property
    @pulumi.getter(name="learningMode")
    def learning_mode(self) -> Optional[pulumi.Input[str]]:
        """
        learning mode setting for policy-builder, possible options: [`automatic`,`disabled`, `manual`]
        """
        return pulumi.get(self, "learning_mode")

    @learning_mode.setter
    def learning_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "learning_mode", value)


@pulumi.input_type
class WafPolicySignaturesSettingArgs:
    def __init__(__self__, *,
                 placesignatures_in_staging: Optional[pulumi.Input[bool]] = None,
                 signature_staging: Optional[pulumi.Input[bool]] = None):
        if placesignatures_in_staging is not None:
            pulumi.set(__self__, "placesignatures_in_staging", placesignatures_in_staging)
        if signature_staging is not None:
            pulumi.set(__self__, "signature_staging", signature_staging)

    @property
    @pulumi.getter(name="placesignaturesInStaging")
    def placesignatures_in_staging(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "placesignatures_in_staging")

    @placesignatures_in_staging.setter
    def placesignatures_in_staging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "placesignatures_in_staging", value)

    @property
    @pulumi.getter(name="signatureStaging")
    def signature_staging(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "signature_staging")

    @signature_staging.setter
    def signature_staging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "signature_staging", value)


