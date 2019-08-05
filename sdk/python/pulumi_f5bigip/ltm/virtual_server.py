# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VirtualServer(pulumi.CustomResource):
    client_profiles: pulumi.Output[list]
    """
    List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
    """
    destination: pulumi.Output[str]
    """
    Destination IP
    """
    fallback_persistence_profile: pulumi.Output[str]
    """
    Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
    """
    ip_protocol: pulumi.Output[str]
    irules: pulumi.Output[list]
    mask: pulumi.Output[str]
    """
    Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
    """
    name: pulumi.Output[str]
    persistence_profiles: pulumi.Output[list]
    """
    List of persistence profiles associated with the Virtual Server.
    """
    policies: pulumi.Output[list]
    pool: pulumi.Output[str]
    """
    Default pool name
    """
    port: pulumi.Output[float]
    """
    Listen port for the virtual server
    """
    profiles: pulumi.Output[list]
    """
    List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
    """
    server_profiles: pulumi.Output[list]
    """
    List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
    """
    snatpool: pulumi.Output[str]
    """
    Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
    """
    source: pulumi.Output[str]
    """
    Specifies an IP address or network from which the virtual server will accept traffic.
    """
    source_address_translation: pulumi.Output[str]
    """
    Can be either omitted for none or the values automap or snat
    """
    translate_address: pulumi.Output[str]
    """
    Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
    """
    translate_port: pulumi.Output[str]
    """
    Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
    """
    vlans: pulumi.Output[list]
    """
    The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
    """
    vlans_enabled: pulumi.Output[bool]
    """
    Enables the virtual server on the VLANs specified by the VLANs option.
    """
    def __init__(__self__, resource_name, opts=None, client_profiles=None, destination=None, fallback_persistence_profile=None, ip_protocol=None, irules=None, mask=None, name=None, persistence_profiles=None, policies=None, pool=None, port=None, profiles=None, server_profiles=None, snatpool=None, source=None, source_address_translation=None, translate_address=None, translate_port=None, vlans=None, vlans_enabled=None, __name__=None, __opts__=None):
        """
        `bigip_ltm_virtual_server` Configures Virtual Server
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] client_profiles: List of client context profiles associated on the virtual server. Not mutually exclusive with profiles and server_profiles
        :param pulumi.Input[str] destination: Destination IP
        :param pulumi.Input[str] fallback_persistence_profile: Specifies a fallback persistence profile for the Virtual Server to use when the default persistence profile is not available.
        :param pulumi.Input[str] mask: Mask can either be in CIDR notation or decimal, i.e.: 24 or 255.255.255.0. A CIDR mask of 0 is the same as 0.0.0.0
        :param pulumi.Input[list] persistence_profiles: List of persistence profiles associated with the Virtual Server.
        :param pulumi.Input[str] pool: Default pool name
        :param pulumi.Input[float] port: Listen port for the virtual server
        :param pulumi.Input[list] profiles: List of profiles associated both client and server contexts on the virtual server. This includes protocol, ssl, http, etc.
        :param pulumi.Input[list] server_profiles: List of server context profiles associated on the virtual server. Not mutually exclusive with profiles and client_profiles
        :param pulumi.Input[str] snatpool: Specifies the name of an existing SNAT pool that you want the virtual server to use to implement selective and intelligent SNATs. DEPRECATED - see Virtual Server Property Groups source-address-translation
        :param pulumi.Input[str] source: Specifies an IP address or network from which the virtual server will accept traffic.
        :param pulumi.Input[str] source_address_translation: Can be either omitted for none or the values automap or snat
        :param pulumi.Input[str] translate_address: Enables or disables address translation for the virtual server. Turn address translation off for a virtual server if you want to use the virtual server to load balance connections to any address. This option is useful when the system is load balancing devices that have the same IP address.
        :param pulumi.Input[str] translate_port: Enables or disables port translation. Turn port translation off for a virtual server if you want to use the virtual server to load balance connections to any service
        :param pulumi.Input[list] vlans: The virtual server is enabled/disabled on this set of VLANs. See vlans-disabled and vlans-enabled.
        :param pulumi.Input[bool] vlans_enabled: Enables the virtual server on the VLANs specified by the VLANs option.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_virtual_server.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['client_profiles'] = client_profiles

        if destination is None:
            raise TypeError("Missing required property 'destination'")
        __props__['destination'] = destination

        __props__['fallback_persistence_profile'] = fallback_persistence_profile

        __props__['ip_protocol'] = ip_protocol

        __props__['irules'] = irules

        __props__['mask'] = mask

        if name is None:
            raise TypeError("Missing required property 'name'")
        __props__['name'] = name

        __props__['persistence_profiles'] = persistence_profiles

        __props__['policies'] = policies

        __props__['pool'] = pool

        if port is None:
            raise TypeError("Missing required property 'port'")
        __props__['port'] = port

        __props__['profiles'] = profiles

        __props__['server_profiles'] = server_profiles

        __props__['snatpool'] = snatpool

        __props__['source'] = source

        __props__['source_address_translation'] = source_address_translation

        __props__['translate_address'] = translate_address

        __props__['translate_port'] = translate_port

        __props__['vlans'] = vlans

        __props__['vlans_enabled'] = vlans_enabled

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(VirtualServer, __self__).__init__(
            'f5bigip:ltm/virtualServer:VirtualServer',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

