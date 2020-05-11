# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Vlan(pulumi.CustomResource):
    interfaces: pulumi.Output[list]
    """
    Specifies which interfaces you want this VLAN to use for traffic management.

      * `tagged` (`bool`) - Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
      * `vlanport` (`str`) - Physical or virtual port used for traffic
    """
    name: pulumi.Output[str]
    """
    Name of the vlan
    """
    tag: pulumi.Output[float]
    """
    Specifies a number that the system adds into the header of any frame passing through the VLAN.
    """
    def __init__(__self__, resource_name, opts=None, interfaces=None, name=None, tag=None, __props__=None, __name__=None, __opts__=None):
        """
        `net.Vlan` Manages a vlan configuration

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.


        ## Example Usage



        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        vlan1 = f5bigip.net.Vlan("vlan1",
            interfaces=[{
                "tagged": False,
                "vlanport": 1.2,
            }],
            name="/Common/Internal",
            tag=101)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] interfaces: Specifies which interfaces you want this VLAN to use for traffic management.
        :param pulumi.Input[str] name: Name of the vlan
        :param pulumi.Input[float] tag: Specifies a number that the system adds into the header of any frame passing through the VLAN.

        The **interfaces** object supports the following:

          * `tagged` (`pulumi.Input[bool]`) - Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
          * `vlanport` (`pulumi.Input[str]`) - Physical or virtual port used for traffic
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

            __props__['interfaces'] = interfaces
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['tag'] = tag
        super(Vlan, __self__).__init__(
            'f5bigip:net/vlan:Vlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, interfaces=None, name=None, tag=None):
        """
        Get an existing Vlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] interfaces: Specifies which interfaces you want this VLAN to use for traffic management.
        :param pulumi.Input[str] name: Name of the vlan
        :param pulumi.Input[float] tag: Specifies a number that the system adds into the header of any frame passing through the VLAN.

        The **interfaces** object supports the following:

          * `tagged` (`pulumi.Input[bool]`) - Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
          * `vlanport` (`pulumi.Input[str]`) - Physical or virtual port used for traffic
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["interfaces"] = interfaces
        __props__["name"] = name
        __props__["tag"] = tag
        return Vlan(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

