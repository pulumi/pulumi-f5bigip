# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['EventServiceDiscoveryArgs', 'EventServiceDiscovery']

@pulumi.input_type
class EventServiceDiscoveryArgs:
    def __init__(__self__, *,
                 taskid: pulumi.Input[builtins.str],
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]]] = None):
        """
        The set of arguments for constructing a EventServiceDiscovery resource.
        :param pulumi.Input[builtins.str] taskid: servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        :param pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]] nodes: Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
               
               For more information, please refer below document
               https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
               
               Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
               
               With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
               
               When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
               
               For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
               
               Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
               First we show the initial declaration to POST to the BIG-IP system.
               
               {
               "class": "ADC",
               "schemaVersion": "3.9.0",
               "id": "Pool",
               "Sample_event_sd": {
               "class": "Tenant",
               "My_app": {
               "class": "Application",
               "My_pool": {
               "class": "Pool",
               "members": [
               {
               "servicePort": 8080,
               "addressDiscovery": "static",
               "serverAddresses": [
               "192.0.2.2"
               ]
               },
               {
               "servicePort": 8080,
               "addressDiscovery": "event"
               }
               ]
               }
               }
               }
               }
               
               
               Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        """
        pulumi.set(__self__, "taskid", taskid)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)

    @property
    @pulumi.getter
    def taskid(self) -> pulumi.Input[builtins.str]:
        """
        servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        """
        return pulumi.get(self, "taskid")

    @taskid.setter
    def taskid(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "taskid", value)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]]]:
        """
        Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)

        For more information, please refer below document
        https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery

        Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.

        With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.

        When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.

        For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes

        Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
        First we show the initial declaration to POST to the BIG-IP system.

        {
        "class": "ADC",
        "schemaVersion": "3.9.0",
        "id": "Pool",
        "Sample_event_sd": {
        "class": "Tenant",
        "My_app": {
        "class": "Application",
        "My_pool": {
        "class": "Pool",
        "members": [
        {
        "servicePort": 8080,
        "addressDiscovery": "static",
        "serverAddresses": [
        "192.0.2.2"
        ]
        },
        {
        "servicePort": 8080,
        "addressDiscovery": "event"
        }
        ]
        }
        }
        }
        }


        Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]]]):
        pulumi.set(self, "nodes", value)


@pulumi.input_type
class _EventServiceDiscoveryState:
    def __init__(__self__, *,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]]] = None,
                 taskid: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering EventServiceDiscovery resources.
        :param pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]] nodes: Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
               
               For more information, please refer below document
               https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
               
               Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
               
               With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
               
               When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
               
               For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
               
               Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
               First we show the initial declaration to POST to the BIG-IP system.
               
               {
               "class": "ADC",
               "schemaVersion": "3.9.0",
               "id": "Pool",
               "Sample_event_sd": {
               "class": "Tenant",
               "My_app": {
               "class": "Application",
               "My_pool": {
               "class": "Pool",
               "members": [
               {
               "servicePort": 8080,
               "addressDiscovery": "static",
               "serverAddresses": [
               "192.0.2.2"
               ]
               },
               {
               "servicePort": 8080,
               "addressDiscovery": "event"
               }
               ]
               }
               }
               }
               }
               
               
               Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        :param pulumi.Input[builtins.str] taskid: servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        """
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if taskid is not None:
            pulumi.set(__self__, "taskid", taskid)

    @property
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]]]:
        """
        Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)

        For more information, please refer below document
        https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery

        Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.

        With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.

        When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.

        For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes

        Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
        First we show the initial declaration to POST to the BIG-IP system.

        {
        "class": "ADC",
        "schemaVersion": "3.9.0",
        "id": "Pool",
        "Sample_event_sd": {
        "class": "Tenant",
        "My_app": {
        "class": "Application",
        "My_pool": {
        "class": "Pool",
        "members": [
        {
        "servicePort": 8080,
        "addressDiscovery": "static",
        "serverAddresses": [
        "192.0.2.2"
        ]
        },
        {
        "servicePort": 8080,
        "addressDiscovery": "event"
        }
        ]
        }
        }
        }
        }


        Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        """
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventServiceDiscoveryNodeArgs']]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def taskid(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        """
        return pulumi.get(self, "taskid")

    @taskid.setter
    def taskid(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "taskid", value)


class EventServiceDiscovery(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventServiceDiscoveryNodeArgs', 'EventServiceDiscoveryNodeArgsDict']]]]] = None,
                 taskid: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test = f5bigip.EventServiceDiscovery("test",
            taskid="~Sample_event_sd~My_app~My_pool",
            nodes=[
                {
                    "id": "newNode1",
                    "ip": "192.168.2.3",
                    "port": 8080,
                },
                {
                    "id": "newNode2",
                    "ip": "192.168.2.4",
                    "port": 8080,
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventServiceDiscoveryNodeArgs', 'EventServiceDiscoveryNodeArgsDict']]]] nodes: Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
               
               For more information, please refer below document
               https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
               
               Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
               
               With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
               
               When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
               
               For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
               
               Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
               First we show the initial declaration to POST to the BIG-IP system.
               
               {
               "class": "ADC",
               "schemaVersion": "3.9.0",
               "id": "Pool",
               "Sample_event_sd": {
               "class": "Tenant",
               "My_app": {
               "class": "Application",
               "My_pool": {
               "class": "Pool",
               "members": [
               {
               "servicePort": 8080,
               "addressDiscovery": "static",
               "serverAddresses": [
               "192.0.2.2"
               ]
               },
               {
               "servicePort": 8080,
               "addressDiscovery": "event"
               }
               ]
               }
               }
               }
               }
               
               
               Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        :param pulumi.Input[builtins.str] taskid: servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventServiceDiscoveryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test = f5bigip.EventServiceDiscovery("test",
            taskid="~Sample_event_sd~My_app~My_pool",
            nodes=[
                {
                    "id": "newNode1",
                    "ip": "192.168.2.3",
                    "port": 8080,
                },
                {
                    "id": "newNode2",
                    "ip": "192.168.2.4",
                    "port": 8080,
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param EventServiceDiscoveryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventServiceDiscoveryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventServiceDiscoveryNodeArgs', 'EventServiceDiscoveryNodeArgsDict']]]]] = None,
                 taskid: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventServiceDiscoveryArgs.__new__(EventServiceDiscoveryArgs)

            __props__.__dict__["nodes"] = nodes
            if taskid is None and not opts.urn:
                raise TypeError("Missing required property 'taskid'")
            __props__.__dict__["taskid"] = taskid
        super(EventServiceDiscovery, __self__).__init__(
            'f5bigip:index/eventServiceDiscovery:EventServiceDiscovery',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            nodes: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EventServiceDiscoveryNodeArgs', 'EventServiceDiscoveryNodeArgsDict']]]]] = None,
            taskid: Optional[pulumi.Input[builtins.str]] = None) -> 'EventServiceDiscovery':
        """
        Get an existing EventServiceDiscovery resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EventServiceDiscoveryNodeArgs', 'EventServiceDiscoveryNodeArgsDict']]]] nodes: Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)
               
               For more information, please refer below document
               https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery
               
               Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.
               
               With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.
               
               When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.
               
               For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes
               
               Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
               First we show the initial declaration to POST to the BIG-IP system.
               
               {
               "class": "ADC",
               "schemaVersion": "3.9.0",
               "id": "Pool",
               "Sample_event_sd": {
               "class": "Tenant",
               "My_app": {
               "class": "Application",
               "My_pool": {
               "class": "Pool",
               "members": [
               {
               "servicePort": 8080,
               "addressDiscovery": "static",
               "serverAddresses": [
               "192.0.2.2"
               ]
               },
               {
               "servicePort": 8080,
               "addressDiscovery": "event"
               }
               ]
               }
               }
               }
               }
               
               
               Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        :param pulumi.Input[builtins.str] taskid: servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventServiceDiscoveryState.__new__(_EventServiceDiscoveryState)

        __props__.__dict__["nodes"] = nodes
        __props__.__dict__["taskid"] = taskid
        return EventServiceDiscovery(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Optional[Sequence['outputs.EventServiceDiscoveryNode']]]:
        """
        Map of node which will be added to pool which will be having node name(id),node address(ip) and node port(port)

        For more information, please refer below document
        https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/declarations/discovery.html?highlight=service%20discovery#event-driven-service-discovery

        Below example shows how to use event-driven service discovery, introduced in AS3 3.9.0.

        With event-driven service discovery, you POST a declaration with the addressDiscovery property set to event. This creates a new endpoint which you can use to add nodes that does not require an AS3 declaration, so it can be more efficient than using PATCH or POST to add nodes.

        When you use the event value for addressDiscovery, the system creates the new endpoint with the following syntax: https://<host>/mgmt/shared/service-discovery/task/~<tenant name>~<application name>~<pool name>/nodes.

        For example, in the following declaration, assuming 192.0.2.14 is our BIG-IP, the endpoint that is created is: https://192.0.2.14/mgmt/shared/service-discovery/task/~Sample_event_sd~My_app~My_pool/nodes

        Once the endpoint is created( taskid ), you can use it to add nodes to the BIG-IP pool
        First we show the initial declaration to POST to the BIG-IP system.

        {
        "class": "ADC",
        "schemaVersion": "3.9.0",
        "id": "Pool",
        "Sample_event_sd": {
        "class": "Tenant",
        "My_app": {
        "class": "Application",
        "My_pool": {
        "class": "Pool",
        "members": [
        {
        "servicePort": 8080,
        "addressDiscovery": "static",
        "serverAddresses": [
        "192.0.2.2"
        ]
        },
        {
        "servicePort": 8080,
        "addressDiscovery": "event"
        }
        ]
        }
        }
        }
        }


        Once the declaration has been sent to the BIG-IP, we can use taskid/id ( ~Sample_event_sd~My_app~My_pool" ) and node list for the resource to dynamically update the node list.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter
    def taskid(self) -> pulumi.Output[builtins.str]:
        """
        servicediscovery endpoint ( Below example shows how to create endpoing using AS3 )
        """
        return pulumi.get(self, "taskid")

