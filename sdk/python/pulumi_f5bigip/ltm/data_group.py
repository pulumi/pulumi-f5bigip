# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class DataGroup(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
    """
    records: pulumi.Output[list]
    """
    a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
    
      * `data` (`str`) - , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
      * `name` (`str`) - , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
    """
    type: pulumi.Output[str]
    """
    datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
    """
    def __init__(__self__, resource_name, opts=None, name=None, records=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.DataGroup` Manages internal (in-line) datagroup configuration
        
        Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-datagroup.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[list] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        
        The **records** object supports the following:
        
          * `data` (`pulumi.Input[str]`) - , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
          * `name` (`pulumi.Input[str]`) - , sets the value of the record's `name` attribute, must be of type defined in `type` attribute

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_datagroup.html.markdown.
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

            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['records'] = records
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(DataGroup, __self__).__init__(
            'f5bigip:ltm/dataGroup:DataGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, records=None, type=None):
        """
        Get an existing DataGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
        :param pulumi.Input[list] records: a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
        :param pulumi.Input[str] type: datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
        
        The **records** object supports the following:
        
          * `data` (`pulumi.Input[str]`) - , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
          * `name` (`pulumi.Input[str]`) - , sets the value of the record's `name` attribute, must be of type defined in `type` attribute

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_datagroup.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["name"] = name
        __props__["records"] = records
        __props__["type"] = type
        return DataGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

