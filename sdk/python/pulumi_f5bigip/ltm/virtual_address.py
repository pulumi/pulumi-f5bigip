# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VirtualAddress(pulumi.CustomResource):
    advertize_route: pulumi.Output[bool]
    """
    Enabled dynamic routing of the address
    """
    arp: pulumi.Output[bool]
    """
    Enable or disable ARP for the virtual address
    """
    auto_delete: pulumi.Output[bool]
    """
    Automatically delete the virtual address with the virtual server
    """
    conn_limit: pulumi.Output[float]
    """
    Max number of connections for virtual address
    """
    enabled: pulumi.Output[bool]
    """
    Enable or disable the virtual address
    """
    icmp_echo: pulumi.Output[bool]
    """
    Enable/Disable ICMP response to the virtual address
    """
    name: pulumi.Output[str]
    """
    Name of the virtual address
    """
    traffic_group: pulumi.Output[str]
    """
    Specify the partition and traffic group
    """
    def __init__(__self__, resource_name, opts=None, advertize_route=None, arp=None, auto_delete=None, conn_limit=None, enabled=None, icmp_echo=None, name=None, traffic_group=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.VirtualAddress` Configures Virtual Server

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        vs_va = f5bigip.ltm.VirtualAddress("vsVa",
            advertize_route=True,
            name="/Common/vs_va")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] advertize_route: Enabled dynamic routing of the address
        :param pulumi.Input[bool] arp: Enable or disable ARP for the virtual address
        :param pulumi.Input[bool] auto_delete: Automatically delete the virtual address with the virtual server
        :param pulumi.Input[float] conn_limit: Max number of connections for virtual address
        :param pulumi.Input[bool] enabled: Enable or disable the virtual address
        :param pulumi.Input[bool] icmp_echo: Enable/Disable ICMP response to the virtual address
        :param pulumi.Input[str] name: Name of the virtual address
        :param pulumi.Input[str] traffic_group: Specify the partition and traffic group
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

            __props__['advertize_route'] = advertize_route
            __props__['arp'] = arp
            __props__['auto_delete'] = auto_delete
            __props__['conn_limit'] = conn_limit
            __props__['enabled'] = enabled
            __props__['icmp_echo'] = icmp_echo
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['traffic_group'] = traffic_group
        super(VirtualAddress, __self__).__init__(
            'f5bigip:ltm/virtualAddress:VirtualAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, advertize_route=None, arp=None, auto_delete=None, conn_limit=None, enabled=None, icmp_echo=None, name=None, traffic_group=None):
        """
        Get an existing VirtualAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] advertize_route: Enabled dynamic routing of the address
        :param pulumi.Input[bool] arp: Enable or disable ARP for the virtual address
        :param pulumi.Input[bool] auto_delete: Automatically delete the virtual address with the virtual server
        :param pulumi.Input[float] conn_limit: Max number of connections for virtual address
        :param pulumi.Input[bool] enabled: Enable or disable the virtual address
        :param pulumi.Input[bool] icmp_echo: Enable/Disable ICMP response to the virtual address
        :param pulumi.Input[str] name: Name of the virtual address
        :param pulumi.Input[str] traffic_group: Specify the partition and traffic group
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["advertize_route"] = advertize_route
        __props__["arp"] = arp
        __props__["auto_delete"] = auto_delete
        __props__["conn_limit"] = conn_limit
        __props__["enabled"] = enabled
        __props__["icmp_echo"] = icmp_echo
        __props__["name"] = name
        __props__["traffic_group"] = traffic_group
        return VirtualAddress(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
