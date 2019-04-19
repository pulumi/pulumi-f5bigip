# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProfileFastL4(pulumi.CustomResource):
    client_timeout: pulumi.Output[float]
    """
    Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
    """
    defaults_from: pulumi.Output[str]
    """
    Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
    """
    explicitflow_migration: pulumi.Output[str]
    """
    Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
    """
    hardware_syncookie: pulumi.Output[str]
    """
    Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
    """
    idle_timeout: pulumi.Output[str]
    """
    Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
    """
    iptos_toclient: pulumi.Output[str]
    """
    Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
    """
    iptos_toserver: pulumi.Output[str]
    """
    Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
    """
    keepalive_interval: pulumi.Output[str]
    """
    Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
    """
    name: pulumi.Output[str]
    """
    Name of the profile_fastl4
    """
    partition: pulumi.Output[str]
    """
    Displays the administrative partition within which this profile resides
    """
    def __init__(__self__, resource_name, opts=None, client_timeout=None, defaults_from=None, explicitflow_migration=None, hardware_syncookie=None, idle_timeout=None, iptos_toclient=None, iptos_toserver=None, keepalive_interval=None, name=None, partition=None, __name__=None, __opts__=None):
        """
        `bigip_ltm_profile_fastl4` Configures a custom profile_fastl4 for use by health checks.
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] client_timeout: Specifies late binding client timeout in seconds. This setting specifies the number of seconds allowed for a client to transmit enough data to select a server when late binding is enabled. If it expires timeout-recovery mode will dictate what action to take.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] explicitflow_migration: Enables or disables late binding explicit flow migration that allows iRules to control when flows move from software to hardware. Explicit flow migration is disabled by default hence BIG-IP automatically migrates flows from software to hardware.
        :param pulumi.Input[str] hardware_syncookie: Enables or disables hardware SYN cookie support when PVA10 is present on the system. Note that when you set the hardware syncookie option to enabled, you may also want to set the following bigdb database variables using the "/sys modify db" command, based on your requirements: pva.SynCookies.Full.ConnectionThreshold (default: 500000), pva.SynCookies.Assist.ConnectionThreshold (default: 500000) pva.SynCookies.ClientWindow (default: 0). The default value is disabled.
        :param pulumi.Input[str] idle_timeout: Specifies an idle timeout in seconds. This setting specifies the number of seconds that a connection is idle before the connection is eligible for deletion.When you specify an idle timeout for the Fast L4 profile, the value must be greater than the bigdb database variable Pva.Scrub time in msec for it to work properly.The default value is 300 seconds.
        :param pulumi.Input[str] iptos_toclient: Specifies an IP ToS number for the client side. This option specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to clients. The default value is 65535 (pass-through), which indicates, do not modify.
        :param pulumi.Input[str] iptos_toserver: Specifies an IP ToS number for the server side. This setting specifies the Type of Service level that the traffic management system assigns to IP packets when sending them to servers. The default value is 65535 (pass-through), which indicates, do not modify.
        :param pulumi.Input[str] keepalive_interval: Specifies the keep alive probe interval, in seconds. The default value is disabled (0 seconds).
        :param pulumi.Input[str] name: Name of the profile_fastl4
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides
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

        __props__['client_timeout'] = client_timeout

        __props__['defaults_from'] = defaults_from

        __props__['explicitflow_migration'] = explicitflow_migration

        __props__['hardware_syncookie'] = hardware_syncookie

        __props__['idle_timeout'] = idle_timeout

        __props__['iptos_toclient'] = iptos_toclient

        __props__['iptos_toserver'] = iptos_toserver

        __props__['keepalive_interval'] = keepalive_interval

        if name is None:
            raise TypeError("Missing required property 'name'")
        __props__['name'] = name

        __props__['partition'] = partition

        super(ProfileFastL4, __self__).__init__(
            'f5bigip:ltm/profileFastL4:ProfileFastL4',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

