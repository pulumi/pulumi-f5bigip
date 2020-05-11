# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Node(pulumi.CustomResource):
    address: pulumi.Output[str]
    """
    IP or hostname of the node
    """
    connection_limit: pulumi.Output[float]
    """
    Specifies the maximum number of connections allowed for the node or node address.
    """
    description: pulumi.Output[str]
    """
    User-defined description give ltm_node
    """
    dynamic_ratio: pulumi.Output[float]
    """
    Specifies the fixed ratio value used for a node during ratio load balancing.
    """
    fqdn: pulumi.Output[dict]
    monitor: pulumi.Output[str]
    """
    specifies the name of the monitor or monitor rule that you want to associate with the node.
    """
    name: pulumi.Output[str]
    """
    Name of the node
    """
    rate_limit: pulumi.Output[str]
    """
    Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
    """
    ratio: pulumi.Output[float]
    """
    Sets the ratio number for the node.
    """
    state: pulumi.Output[str]
    """
    Default is "user-up" you can set to "user-down" if you want to disable
    """
    def __init__(__self__, resource_name, opts=None, address=None, connection_limit=None, description=None, dynamic_ratio=None, fqdn=None, monitor=None, name=None, rate_limit=None, ratio=None, state=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Node resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IP or hostname of the node
        :param pulumi.Input[float] connection_limit: Specifies the maximum number of connections allowed for the node or node address.
        :param pulumi.Input[str] description: User-defined description give ltm_node
        :param pulumi.Input[float] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] monitor: specifies the name of the monitor or monitor rule that you want to associate with the node.
        :param pulumi.Input[str] name: Name of the node
        :param pulumi.Input[str] rate_limit: Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        :param pulumi.Input[float] ratio: Sets the ratio number for the node.
        :param pulumi.Input[str] state: Default is "user-up" you can set to "user-down" if you want to disable

        The **fqdn** object supports the following:

          * `addressFamily` (`pulumi.Input[str]`) - Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
          * `autopopulate` (`pulumi.Input[str]`)
          * `downinterval` (`pulumi.Input[float]`)
          * `interval` (`pulumi.Input[str]`) - Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
          * `name` (`pulumi.Input[str]`) - Name of the node
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if address is None:
                raise TypeError("Missing required property 'address'")
            __props__['address'] = address
            __props__['connection_limit'] = connection_limit
            __props__['description'] = description
            __props__['dynamic_ratio'] = dynamic_ratio
            __props__['fqdn'] = fqdn
            __props__['monitor'] = monitor
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['rate_limit'] = rate_limit
            __props__['ratio'] = ratio
            __props__['state'] = state
        super(Node, __self__).__init__(
            'f5bigip:ltm/node:Node',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, address=None, connection_limit=None, description=None, dynamic_ratio=None, fqdn=None, monitor=None, name=None, rate_limit=None, ratio=None, state=None):
        """
        Get an existing Node resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IP or hostname of the node
        :param pulumi.Input[float] connection_limit: Specifies the maximum number of connections allowed for the node or node address.
        :param pulumi.Input[str] description: User-defined description give ltm_node
        :param pulumi.Input[float] dynamic_ratio: Specifies the fixed ratio value used for a node during ratio load balancing.
        :param pulumi.Input[str] monitor: specifies the name of the monitor or monitor rule that you want to associate with the node.
        :param pulumi.Input[str] name: Name of the node
        :param pulumi.Input[str] rate_limit: Specifies the maximum number of connections per second allowed for a node or node address. The default value is 'disabled'.
        :param pulumi.Input[float] ratio: Sets the ratio number for the node.
        :param pulumi.Input[str] state: Default is "user-up" you can set to "user-down" if you want to disable

        The **fqdn** object supports the following:

          * `addressFamily` (`pulumi.Input[str]`) - Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
          * `autopopulate` (`pulumi.Input[str]`)
          * `downinterval` (`pulumi.Input[float]`)
          * `interval` (`pulumi.Input[str]`) - Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
          * `name` (`pulumi.Input[str]`) - Name of the node
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address"] = address
        __props__["connection_limit"] = connection_limit
        __props__["description"] = description
        __props__["dynamic_ratio"] = dynamic_ratio
        __props__["fqdn"] = fqdn
        __props__["monitor"] = monitor
        __props__["name"] = name
        __props__["rate_limit"] = rate_limit
        __props__["ratio"] = ratio
        __props__["state"] = state
        return Node(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

