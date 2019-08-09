# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Provision(pulumi.CustomResource):
    cpu_ratio: pulumi.Output[float]
    disk_ratio: pulumi.Output[float]
    full_path: pulumi.Output[str]
    level: pulumi.Output[str]
    memory_ratio: pulumi.Output[float]
    name: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, cpu_ratio=None, disk_ratio=None, full_path=None, level=None, memory_ratio=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        `sys.Provision` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/sys_provision.html.markdown.
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

            __props__['cpu_ratio'] = cpu_ratio
            __props__['disk_ratio'] = disk_ratio
            __props__['full_path'] = full_path
            __props__['level'] = level
            __props__['memory_ratio'] = memory_ratio
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
        super(Provision, __self__).__init__(
            'f5bigip:sys/provision:Provision',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cpu_ratio=None, disk_ratio=None, full_path=None, level=None, memory_ratio=None, name=None):
        """
        Get an existing Provision resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/sys_provision.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["cpu_ratio"] = cpu_ratio
        __props__["disk_ratio"] = disk_ratio
        __props__["full_path"] = full_path
        __props__["level"] = level
        __props__["memory_ratio"] = memory_ratio
        __props__["name"] = name
        return Provision(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

