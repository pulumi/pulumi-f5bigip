# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class PoolAttachment(pulumi.CustomResource):
    node: pulumi.Output[str]
    """
    Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
    """
    pool: pulumi.Output[str]
    """
    Name of the pool in /Partition/Name format
    """
    def __init__(__self__, resource_name, opts=None, node=None, pool=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a PoolAttachment resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] node: Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
        :param pulumi.Input[str] pool: Name of the pool in /Partition/Name format
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

            if node is None:
                raise TypeError("Missing required property 'node'")
            __props__['node'] = node
            if pool is None:
                raise TypeError("Missing required property 'pool'")
            __props__['pool'] = pool
        super(PoolAttachment, __self__).__init__(
            'f5bigip:ltm/poolAttachment:PoolAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, node=None, pool=None):
        """
        Get an existing PoolAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] node: Node to add to the pool in /Partition/NodeName:Port format (e.g. /Common/Node01:80)
        :param pulumi.Input[str] pool: Name of the pool in /Partition/Name format
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["node"] = node
        __props__["pool"] = pool
        return PoolAttachment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

