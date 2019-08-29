# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Snat(pulumi.CustomResource):
    autolasthop: pulumi.Output[str]
    full_path: pulumi.Output[str]
    mirror: pulumi.Output[str]
    """
    Enables or disables mirroring of SNAT connections.
    """
    name: pulumi.Output[str]
    """
    Name of the snat
    """
    origins: pulumi.Output[list]
    """
    IP or hostname of the snat
    
      * `app_service` (`str`)
      * `name` (`str`) - Name of the snat
    """
    partition: pulumi.Output[str]
    """
    Displays the administrative partition within which this profile resides
    """
    snatpool: pulumi.Output[str]
    """
    Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
    """
    sourceport: pulumi.Output[str]
    """
    Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
    """
    translation: pulumi.Output[str]
    """
    Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
    """
    vlans: pulumi.Output[list]
    """
    Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
    """
    vlansdisabled: pulumi.Output[bool]
    """
    Disables the SNAT on all VLANs.
    """
    def __init__(__self__, resource_name, opts=None, autolasthop=None, full_path=None, mirror=None, name=None, origins=None, partition=None, snatpool=None, sourceport=None, translation=None, vlans=None, vlansdisabled=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.Snat` Manages a snat configuration
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mirror: Enables or disables mirroring of SNAT connections.
        :param pulumi.Input[str] name: Name of the snat
        :param pulumi.Input[list] origins: IP or hostname of the snat
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] snatpool: Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
        :param pulumi.Input[str] sourceport: Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
        :param pulumi.Input[str] translation: Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
        :param pulumi.Input[list] vlans: Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
        :param pulumi.Input[bool] vlansdisabled: Disables the SNAT on all VLANs.
        
        The **origins** object supports the following:
        
          * `app_service` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - Name of the snat

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_snat.html.markdown.
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

            __props__['autolasthop'] = autolasthop
            __props__['full_path'] = full_path
            __props__['mirror'] = mirror
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if origins is None:
                raise TypeError("Missing required property 'origins'")
            __props__['origins'] = origins
            __props__['partition'] = partition
            __props__['snatpool'] = snatpool
            __props__['sourceport'] = sourceport
            __props__['translation'] = translation
            __props__['vlans'] = vlans
            __props__['vlansdisabled'] = vlansdisabled
        super(Snat, __self__).__init__(
            'f5bigip:ltm/snat:Snat',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, autolasthop=None, full_path=None, mirror=None, name=None, origins=None, partition=None, snatpool=None, sourceport=None, translation=None, vlans=None, vlansdisabled=None):
        """
        Get an existing Snat resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mirror: Enables or disables mirroring of SNAT connections.
        :param pulumi.Input[str] name: Name of the snat
        :param pulumi.Input[list] origins: IP or hostname of the snat
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
        :param pulumi.Input[str] snatpool: Specifies the name of a SNAT pool. You can only use this option when automap and translation are not used.
        :param pulumi.Input[str] sourceport: Specifies whether the system preserves the source port of the connection. The default is preserve. Use of the preserve-strict setting should be restricted to UDP only under very special circumstances such as nPath or transparent (that is, no translation of any other L3/L4 field), where there is a 1:1 relationship between virtual IP addresses and node addresses, or when clustered multi-processing (CMP) is disabled. The change setting is useful for obfuscating internal network addresses.
        :param pulumi.Input[str] translation: Specifies the name of a translated IP address. Note that translated addresses are outside the traffic management system. You can only use this option when automap and snatpool are not used.
        :param pulumi.Input[list] vlans: Specifies the name of the VLAN to which you want to assign the SNAT. The default is vlans-enabled.
        :param pulumi.Input[bool] vlansdisabled: Disables the SNAT on all VLANs.
        
        The **origins** object supports the following:
        
          * `app_service` (`pulumi.Input[str]`)
          * `name` (`pulumi.Input[str]`) - Name of the snat

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_snat.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["autolasthop"] = autolasthop
        __props__["full_path"] = full_path
        __props__["mirror"] = mirror
        __props__["name"] = name
        __props__["origins"] = origins
        __props__["partition"] = partition
        __props__["snatpool"] = snatpool
        __props__["sourceport"] = sourceport
        __props__["translation"] = translation
        __props__["vlans"] = vlans
        __props__["vlansdisabled"] = vlansdisabled
        return Snat(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

