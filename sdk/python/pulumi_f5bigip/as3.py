# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class As3(pulumi.CustomResource):
    as3_json: pulumi.Output[str]
    """
    Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
    """
    tenant_filter: pulumi.Output[str]
    """
    If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified 
    """
    tenant_list: pulumi.Output[str]
    """
    Name of Tenant
    """
    tenant_name: pulumi.Output[str]
    """
    Name of Tenant
    """
    def __init__(__self__, resource_name, opts=None, as3_json=None, tenant_filter=None, tenant_list=None, tenant_name=None, __props__=None, __name__=None, __opts__=None):
        """
        `.As3` provides details about bigip as3 resource

        This resource is helpful to configure as3 declarative JSON on BIG-IP.

        ## Example Usage 

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        # Example Usage for json file with tenant filter
        as3_example1 = f5bigip.As3("as3-example1",
            as3_json=(lambda path: open(path).read())("example2.json"),
            tenant_filter="Sample_03")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[str] tenant_filter: If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified 
        :param pulumi.Input[str] tenant_list: Name of Tenant
        :param pulumi.Input[str] tenant_name: Name of Tenant
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

            if as3_json is None:
                raise TypeError("Missing required property 'as3_json'")
            __props__['as3_json'] = as3_json
            __props__['tenant_filter'] = tenant_filter
            __props__['tenant_list'] = tenant_list
            __props__['tenant_name'] = tenant_name
        super(As3, __self__).__init__(
            'f5bigip:index/as3:As3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, as3_json=None, tenant_filter=None, tenant_list=None, tenant_name=None):
        """
        Get an existing As3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[str] tenant_filter: If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified 
        :param pulumi.Input[str] tenant_list: Name of Tenant
        :param pulumi.Input[str] tenant_name: Name of Tenant
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["as3_json"] = as3_json
        __props__["tenant_filter"] = tenant_filter
        __props__["tenant_list"] = tenant_list
        __props__["tenant_name"] = tenant_name
        return As3(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

