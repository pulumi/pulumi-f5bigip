# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Key(pulumi.CustomResource):
    content: pulumi.Output[str]
    """
    Content of SSL certificate key present on local Disk
    """
    name: pulumi.Output[str]
    """
    Name of SSL Certificate key with .key extension
    """
    partition: pulumi.Output[str]
    """
    Partition on to SSL Certificate key to be imported
    """
    def __init__(__self__, resource_name, opts=None, content=None, name=None, partition=None, __props__=None, __name__=None, __opts__=None):
        """
        `ssl.Key` This resource will import SSL certificate key on BIG-IP LTM. 
        Certificate key can be imported from certificate key files on the local disk, in PEM format




        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ssl_key.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: Content of SSL certificate key present on local Disk
        :param pulumi.Input[str] name: Name of SSL Certificate key with .key extension
        :param pulumi.Input[str] partition: Partition on to SSL Certificate key to be imported
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

            if content is None:
                raise TypeError("Missing required property 'content'")
            __props__['content'] = content
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['partition'] = partition
        super(Key, __self__).__init__(
            'f5bigip:ssl/key:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, content=None, name=None, partition=None):
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: Content of SSL certificate key present on local Disk
        :param pulumi.Input[str] name: Name of SSL Certificate key with .key extension
        :param pulumi.Input[str] partition: Partition on to SSL Certificate key to be imported
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content"] = content
        __props__["name"] = name
        __props__["partition"] = partition
        return Key(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
