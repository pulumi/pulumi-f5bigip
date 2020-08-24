# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Route']


class Route(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gw: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `net.Route` Manages a route configuration

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        route2 = f5bigip.net.Route("route2",
            gw="1.1.1.2",
            name="external-route",
            network="10.10.10.0/24")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gw: Gateway address
        :param pulumi.Input[str] name: Name of the route
        :param pulumi.Input[str] network: Specifies a gateway address for the route.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['gw'] = gw
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if network is None:
                raise TypeError("Missing required property 'network'")
            __props__['network'] = network
        super(Route, __self__).__init__(
            'f5bigip:net/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            gw: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[str]] = None) -> 'Route':
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gw: Gateway address
        :param pulumi.Input[str] name: Name of the route
        :param pulumi.Input[str] network: Specifies a gateway address for the route.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["gw"] = gw
        __props__["name"] = name
        __props__["network"] = network
        return Route(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def gw(self) -> Optional[str]:
        """
        Gateway address
        """
        return pulumi.get(self, "gw")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the route
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        Specifies a gateway address for the route.
        """
        return pulumi.get(self, "network")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

