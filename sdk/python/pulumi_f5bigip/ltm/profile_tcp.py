# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProfileTcp(pulumi.CustomResource):
    close_wait_timeout: pulumi.Output[float]
    """
    Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
    """
    defaults_from: pulumi.Output[str]
    """
    Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
    """
    deferred_accept: pulumi.Output[str]
    """
    Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
    """
    fast_open: pulumi.Output[str]
    """
    When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
    """
    finwait2timeout: pulumi.Output[float]
    """
    Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
    """
    finwait_timeout: pulumi.Output[float]
    """
    Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
    """
    idle_timeout: pulumi.Output[float]
    """
    Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
    """
    keepalive_interval: pulumi.Output[float]
    """
    Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
    """
    name: pulumi.Output[str]
    """
    Name of the profile_tcp
    """
    partition: pulumi.Output[str]
    """
    Displays the administrative partition within which this profile resides
    """
    def __init__(__self__, resource_name, opts=None, close_wait_timeout=None, defaults_from=None, deferred_accept=None, fast_open=None, finwait2timeout=None, finwait_timeout=None, idle_timeout=None, keepalive_interval=None, name=None, partition=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.ProfileTcp` Configures a custom profile_tcp for use by health checks.
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] close_wait_timeout: Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] deferred_accept: Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
        :param pulumi.Input[str] fast_open: When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
        :param pulumi.Input[float] finwait2timeout: Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
        :param pulumi.Input[float] finwait_timeout: Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
        :param pulumi.Input[float] idle_timeout: Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
        :param pulumi.Input[float] keepalive_interval: Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        :param pulumi.Input[str] name: Name of the profile_tcp
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_tcp.html.markdown.
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

            __props__['close_wait_timeout'] = close_wait_timeout
            __props__['defaults_from'] = defaults_from
            __props__['deferred_accept'] = deferred_accept
            __props__['fast_open'] = fast_open
            __props__['finwait2timeout'] = finwait2timeout
            __props__['finwait_timeout'] = finwait_timeout
            __props__['idle_timeout'] = idle_timeout
            __props__['keepalive_interval'] = keepalive_interval
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['partition'] = partition
        super(ProfileTcp, __self__).__init__(
            'f5bigip:ltm/profileTcp:ProfileTcp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, close_wait_timeout=None, defaults_from=None, deferred_accept=None, fast_open=None, finwait2timeout=None, finwait_timeout=None, idle_timeout=None, keepalive_interval=None, name=None, partition=None):
        """
        Get an existing ProfileTcp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] close_wait_timeout: Specifies the number of seconds that a connection remains in a LAST-ACK state before quitting. A value of 0 represents a term of forever (or until the maxrtx of the FIN state). The default value is 5 seconds.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] deferred_accept: Specifies, when enabled, that the system defers allocation of the connection chain context until the client response is received. This option is useful for dealing with 3-way handshake DOS attacks. The default value is disabled.
        :param pulumi.Input[str] fast_open: When enabled, permits TCP Fast Open, allowing properly equipped TCP clients to send data with the SYN packet.
        :param pulumi.Input[float] finwait2timeout: Specifies the number of seconds that a connection is in the FIN-WAIT-2 state before quitting. The default value is 300 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state).
        :param pulumi.Input[float] finwait_timeout: Specifies the number of seconds that a connection is in the FIN-WAIT-1 or closing state before quitting. The default value is 5 seconds. A value of 0 (zero) represents a term of forever (or until the maxrtx of the FIN state). You can also specify immediate or indefinite.
        :param pulumi.Input[float] idle_timeout: Specifies the number of seconds that a connection is idle before the connection is eligible for deletion. The default value is 300 seconds.
        :param pulumi.Input[float] keepalive_interval: Specifies the keep alive probe interval, in seconds. The default value is 1800 seconds.
        :param pulumi.Input[str] name: Name of the profile_tcp
        :param pulumi.Input[str] partition: Displays the administrative partition within which this profile resides

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_tcp.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["close_wait_timeout"] = close_wait_timeout
        __props__["defaults_from"] = defaults_from
        __props__["deferred_accept"] = deferred_accept
        __props__["fast_open"] = fast_open
        __props__["finwait2timeout"] = finwait2timeout
        __props__["finwait_timeout"] = finwait_timeout
        __props__["idle_timeout"] = idle_timeout
        __props__["keepalive_interval"] = keepalive_interval
        __props__["name"] = name
        __props__["partition"] = partition
        return ProfileTcp(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

