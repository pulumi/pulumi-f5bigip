# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Monitor(pulumi.CustomResource):
    adaptive: pulumi.Output[str]
    """
    ftp adaptive
    """
    adaptive_limit: pulumi.Output[float]
    """
    Integer value
    """
    compatibility: pulumi.Output[str]
    """
    Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
    """
    database: pulumi.Output[str]
    """
    Specifies the database in which the user is created
    """
    defaults_from: pulumi.Output[str]
    """
    Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
    """
    destination: pulumi.Output[str]
    """
    Specify an alias address for monitoring
    """
    filename: pulumi.Output[str]
    """
    Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
    """
    interval: pulumi.Output[float]
    """
    Check interval in seconds
    """
    ip_dscp: pulumi.Output[float]
    manual_resume: pulumi.Output[str]
    mode: pulumi.Output[str]
    """
    Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
    """
    name: pulumi.Output[str]
    """
    Name of the monitor
    """
    parent: pulumi.Output[str]
    """
    Existing LTM monitor to inherit from
    """
    password: pulumi.Output[str]
    """
    Specifies the password if the monitored target requires authentication 
    """
    receive: pulumi.Output[str]
    """
    Expected response string
    """
    receive_disable: pulumi.Output[str]
    """
    Expected response string.
    """
    reverse: pulumi.Output[str]
    send: pulumi.Output[str]
    """
    Request string to send
    """
    time_until_up: pulumi.Output[float]
    """
    Time in seconds
    """
    timeout: pulumi.Output[float]
    """
    Timeout in seconds
    """
    transparent: pulumi.Output[str]
    username: pulumi.Output[str]
    """
    Specifies the user name if the monitored target requires authentication
    """
    def __init__(__self__, resource_name, opts=None, adaptive=None, adaptive_limit=None, compatibility=None, database=None, defaults_from=None, destination=None, filename=None, interval=None, ip_dscp=None, manual_resume=None, mode=None, name=None, parent=None, password=None, receive=None, receive_disable=None, reverse=None, send=None, time_until_up=None, timeout=None, transparent=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.Monitor` Configures a custom monitor for use by health checks.

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_monitor.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adaptive: ftp adaptive
        :param pulumi.Input[float] adaptive_limit: Integer value
        :param pulumi.Input[str] compatibility: Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        :param pulumi.Input[str] database: Specifies the database in which the user is created
        :param pulumi.Input[str] defaults_from: Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
        :param pulumi.Input[str] destination: Specify an alias address for monitoring
        :param pulumi.Input[str] filename: Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
        :param pulumi.Input[float] interval: Check interval in seconds
        :param pulumi.Input[str] mode: Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        :param pulumi.Input[str] name: Name of the monitor
        :param pulumi.Input[str] parent: Existing LTM monitor to inherit from
        :param pulumi.Input[str] password: Specifies the password if the monitored target requires authentication 
        :param pulumi.Input[str] receive: Expected response string
        :param pulumi.Input[str] receive_disable: Expected response string.
        :param pulumi.Input[str] send: Request string to send
        :param pulumi.Input[float] time_until_up: Time in seconds
        :param pulumi.Input[float] timeout: Timeout in seconds
        :param pulumi.Input[str] username: Specifies the user name if the monitored target requires authentication
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

            __props__['adaptive'] = adaptive
            __props__['adaptive_limit'] = adaptive_limit
            __props__['compatibility'] = compatibility
            __props__['database'] = database
            __props__['defaults_from'] = defaults_from
            __props__['destination'] = destination
            __props__['filename'] = filename
            __props__['interval'] = interval
            __props__['ip_dscp'] = ip_dscp
            __props__['manual_resume'] = manual_resume
            __props__['mode'] = mode
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if parent is None:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['password'] = password
            __props__['receive'] = receive
            __props__['receive_disable'] = receive_disable
            __props__['reverse'] = reverse
            __props__['send'] = send
            __props__['time_until_up'] = time_until_up
            __props__['timeout'] = timeout
            __props__['transparent'] = transparent
            __props__['username'] = username
        super(Monitor, __self__).__init__(
            'f5bigip:ltm/monitor:Monitor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, adaptive=None, adaptive_limit=None, compatibility=None, database=None, defaults_from=None, destination=None, filename=None, interval=None, ip_dscp=None, manual_resume=None, mode=None, name=None, parent=None, password=None, receive=None, receive_disable=None, reverse=None, send=None, time_until_up=None, timeout=None, transparent=None, username=None):
        """
        Get an existing Monitor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adaptive: ftp adaptive
        :param pulumi.Input[float] adaptive_limit: Integer value
        :param pulumi.Input[str] compatibility: Specifies, when enabled, that the SSL options setting (in OpenSSL) is set to ALL. Accepts 'enabled' or 'disabled' values, the default value is 'enabled'.
        :param pulumi.Input[str] database: Specifies the database in which the user is created
        :param pulumi.Input[str] defaults_from: Existing monitor to inherit from. Must be one of /Common/http, /Common/https, /Common/icmp or /Common/gateway-icmp.
        :param pulumi.Input[str] destination: Specify an alias address for monitoring
        :param pulumi.Input[str] filename: Specifies the full path and file name of the file that the system attempts to download. The health check is successful if the system can download the file.
        :param pulumi.Input[float] interval: Check interval in seconds
        :param pulumi.Input[str] mode: Specifies the data transfer process (DTP) mode. The default value is passive. The options are passive (Specifies that the monitor sends a data transfer request to the FTP server. When the FTP server receives the request, the FTP server then initiates and establishes the data connection.) and active (Specifies that the monitor initiates and establishes the data connection with the FTP server.).
        :param pulumi.Input[str] name: Name of the monitor
        :param pulumi.Input[str] parent: Existing LTM monitor to inherit from
        :param pulumi.Input[str] password: Specifies the password if the monitored target requires authentication 
        :param pulumi.Input[str] receive: Expected response string
        :param pulumi.Input[str] receive_disable: Expected response string.
        :param pulumi.Input[str] send: Request string to send
        :param pulumi.Input[float] time_until_up: Time in seconds
        :param pulumi.Input[float] timeout: Timeout in seconds
        :param pulumi.Input[str] username: Specifies the user name if the monitored target requires authentication
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["adaptive"] = adaptive
        __props__["adaptive_limit"] = adaptive_limit
        __props__["compatibility"] = compatibility
        __props__["database"] = database
        __props__["defaults_from"] = defaults_from
        __props__["destination"] = destination
        __props__["filename"] = filename
        __props__["interval"] = interval
        __props__["ip_dscp"] = ip_dscp
        __props__["manual_resume"] = manual_resume
        __props__["mode"] = mode
        __props__["name"] = name
        __props__["parent"] = parent
        __props__["password"] = password
        __props__["receive"] = receive
        __props__["receive_disable"] = receive_disable
        __props__["reverse"] = reverse
        __props__["send"] = send
        __props__["time_until_up"] = time_until_up
        __props__["timeout"] = timeout
        __props__["transparent"] = transparent
        __props__["username"] = username
        return Monitor(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

