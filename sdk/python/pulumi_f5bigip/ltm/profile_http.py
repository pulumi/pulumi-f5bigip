# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ProfileHttp(pulumi.CustomResource):
    accept_xff: pulumi.Output[str]
    app_service: pulumi.Output[str]
    basic_auth_realm: pulumi.Output[str]
    """
    Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
    """
    defaults_from: pulumi.Output[str]
    """
    Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
    """
    description: pulumi.Output[str]
    encrypt_cookie_secret: pulumi.Output[str]
    encrypt_cookies: pulumi.Output[list]
    fallback_host: pulumi.Output[str]
    """
    Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
    """
    fallback_status_codes: pulumi.Output[list]
    """
    Specifies one or more three-digit status codes that can be returned by an HTTP server.
    """
    head_erase: pulumi.Output[str]
    head_insert: pulumi.Output[str]
    """
    Specifies a quoted header string that you want to insert into an HTTP request
    """
    insert_xforwarded_for: pulumi.Output[str]
    """
    When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
    """
    lws_separator: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    Name of the profile_http
    """
    oneconnect_transformations: pulumi.Output[str]
    """
    Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
    """
    proxy_type: pulumi.Output[str]
    redirect_rewrite: pulumi.Output[str]
    request_chunking: pulumi.Output[str]
    response_chunking: pulumi.Output[str]
    response_headers_permitteds: pulumi.Output[list]
    server_agent_name: pulumi.Output[str]
    tm_partition: pulumi.Output[str]
    via_host_name: pulumi.Output[str]
    via_request: pulumi.Output[str]
    via_response: pulumi.Output[str]
    xff_alternative_names: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, accept_xff=None, app_service=None, basic_auth_realm=None, defaults_from=None, description=None, encrypt_cookie_secret=None, encrypt_cookies=None, fallback_host=None, fallback_status_codes=None, head_erase=None, head_insert=None, insert_xforwarded_for=None, lws_separator=None, name=None, oneconnect_transformations=None, proxy_type=None, redirect_rewrite=None, request_chunking=None, response_chunking=None, response_headers_permitteds=None, server_agent_name=None, tm_partition=None, via_host_name=None, via_request=None, via_response=None, xff_alternative_names=None, __props__=None, __name__=None, __opts__=None):
        """
        `ltm.ProfileHttp` Configures a custom profile_http for use by health checks.
        
        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] basic_auth_realm: Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] fallback_host: Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        :param pulumi.Input[list] fallback_status_codes: Specifies one or more three-digit status codes that can be returned by an HTTP server.
        :param pulumi.Input[str] head_insert: Specifies a quoted header string that you want to insert into an HTTP request
        :param pulumi.Input[str] insert_xforwarded_for: When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        :param pulumi.Input[str] name: Name of the profile_http
        :param pulumi.Input[str] oneconnect_transformations: Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_http.html.markdown.
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

            __props__['accept_xff'] = accept_xff
            __props__['app_service'] = app_service
            __props__['basic_auth_realm'] = basic_auth_realm
            if defaults_from is None:
                raise TypeError("Missing required property 'defaults_from'")
            __props__['defaults_from'] = defaults_from
            __props__['description'] = description
            __props__['encrypt_cookie_secret'] = encrypt_cookie_secret
            __props__['encrypt_cookies'] = encrypt_cookies
            __props__['fallback_host'] = fallback_host
            __props__['fallback_status_codes'] = fallback_status_codes
            __props__['head_erase'] = head_erase
            __props__['head_insert'] = head_insert
            __props__['insert_xforwarded_for'] = insert_xforwarded_for
            __props__['lws_separator'] = lws_separator
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['oneconnect_transformations'] = oneconnect_transformations
            __props__['proxy_type'] = proxy_type
            __props__['redirect_rewrite'] = redirect_rewrite
            __props__['request_chunking'] = request_chunking
            __props__['response_chunking'] = response_chunking
            __props__['response_headers_permitteds'] = response_headers_permitteds
            __props__['server_agent_name'] = server_agent_name
            __props__['tm_partition'] = tm_partition
            __props__['via_host_name'] = via_host_name
            __props__['via_request'] = via_request
            __props__['via_response'] = via_response
            __props__['xff_alternative_names'] = xff_alternative_names
        super(ProfileHttp, __self__).__init__(
            'f5bigip:ltm/profileHttp:ProfileHttp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accept_xff=None, app_service=None, basic_auth_realm=None, defaults_from=None, description=None, encrypt_cookie_secret=None, encrypt_cookies=None, fallback_host=None, fallback_status_codes=None, head_erase=None, head_insert=None, insert_xforwarded_for=None, lws_separator=None, name=None, oneconnect_transformations=None, proxy_type=None, redirect_rewrite=None, request_chunking=None, response_chunking=None, response_headers_permitteds=None, server_agent_name=None, tm_partition=None, via_host_name=None, via_request=None, via_response=None, xff_alternative_names=None):
        """
        Get an existing ProfileHttp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] basic_auth_realm: Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] fallback_host: Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        :param pulumi.Input[list] fallback_status_codes: Specifies one or more three-digit status codes that can be returned by an HTTP server.
        :param pulumi.Input[str] head_insert: Specifies a quoted header string that you want to insert into an HTTP request
        :param pulumi.Input[str] insert_xforwarded_for: When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        :param pulumi.Input[str] name: Name of the profile_http
        :param pulumi.Input[str] oneconnect_transformations: Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile

        > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_profile_http.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["accept_xff"] = accept_xff
        __props__["app_service"] = app_service
        __props__["basic_auth_realm"] = basic_auth_realm
        __props__["defaults_from"] = defaults_from
        __props__["description"] = description
        __props__["encrypt_cookie_secret"] = encrypt_cookie_secret
        __props__["encrypt_cookies"] = encrypt_cookies
        __props__["fallback_host"] = fallback_host
        __props__["fallback_status_codes"] = fallback_status_codes
        __props__["head_erase"] = head_erase
        __props__["head_insert"] = head_insert
        __props__["insert_xforwarded_for"] = insert_xforwarded_for
        __props__["lws_separator"] = lws_separator
        __props__["name"] = name
        __props__["oneconnect_transformations"] = oneconnect_transformations
        __props__["proxy_type"] = proxy_type
        __props__["redirect_rewrite"] = redirect_rewrite
        __props__["request_chunking"] = request_chunking
        __props__["response_chunking"] = response_chunking
        __props__["response_headers_permitteds"] = response_headers_permitteds
        __props__["server_agent_name"] = server_agent_name
        __props__["tm_partition"] = tm_partition
        __props__["via_host_name"] = via_host_name
        __props__["via_request"] = via_request
        __props__["via_response"] = via_response
        __props__["xff_alternative_names"] = xff_alternative_names
        return ProfileHttp(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

