# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ProfileHttp']


class ProfileHttp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_xff: Optional[pulumi.Input[str]] = None,
                 app_service: Optional[pulumi.Input[str]] = None,
                 basic_auth_realm: Optional[pulumi.Input[str]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_cookie_secret: Optional[pulumi.Input[str]] = None,
                 encrypt_cookies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 fallback_host: Optional[pulumi.Input[str]] = None,
                 fallback_status_codes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 head_erase: Optional[pulumi.Input[str]] = None,
                 head_insert: Optional[pulumi.Input[str]] = None,
                 insert_xforwarded_for: Optional[pulumi.Input[str]] = None,
                 lws_separator: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 oneconnect_transformations: Optional[pulumi.Input[str]] = None,
                 proxy_type: Optional[pulumi.Input[str]] = None,
                 redirect_rewrite: Optional[pulumi.Input[str]] = None,
                 request_chunking: Optional[pulumi.Input[str]] = None,
                 response_chunking: Optional[pulumi.Input[str]] = None,
                 response_headers_permitteds: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 server_agent_name: Optional[pulumi.Input[str]] = None,
                 tm_partition: Optional[pulumi.Input[str]] = None,
                 via_host_name: Optional[pulumi.Input[str]] = None,
                 via_request: Optional[pulumi.Input[str]] = None,
                 via_response: Optional[pulumi.Input[str]] = None,
                 xff_alternative_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.ProfileHttp` Configures a custom profile_http for use by health checks.

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        sanjose_http = f5bigip.ltm.ProfileHttp("sanjose-http",
            defaults_from="/Common/http",
            description="some http",
            fallback_host="titanic",
            fallback_status_codes=[
                "400",
                "500",
                "300",
            ],
            name="/Common/sanjose-http")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_xff: Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
               XFF (X-forwarded-for) headers, if they exist.
        :param pulumi.Input[str] app_service: The application service to which the object belongs.
        :param pulumi.Input[str] basic_auth_realm: Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] description: User defibned description
        :param pulumi.Input[str] encrypt_cookie_secret: Specifies a passphrase for the cookie encryption
        :param pulumi.Input[List[pulumi.Input[str]]] encrypt_cookies: Encrypts specified cookies that the BIG-IP system sends to a client system
        :param pulumi.Input[str] fallback_host: Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        :param pulumi.Input[List[pulumi.Input[str]]] fallback_status_codes: Specifies one or more three-digit status codes that can be returned by an HTTP server.
        :param pulumi.Input[str] head_erase: Specifies the header string that you want to erase from an HTTP request. You can also specify none
        :param pulumi.Input[str] head_insert: Specifies a quoted header string that you want to insert into an HTTP request
        :param pulumi.Input[str] insert_xforwarded_for: When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        :param pulumi.Input[str] lws_separator: Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
        :param pulumi.Input[str] name: Name of the profile_http
        :param pulumi.Input[str] oneconnect_transformations: Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
        :param pulumi.Input[str] proxy_type: Specifies the type of HTTP proxy.
        :param pulumi.Input[str] redirect_rewrite: Specifies which of the application HTTP redirects the system rewrites to HTTPS.
        :param pulumi.Input[str] request_chunking: Specifies how to handle chunked and unchunked requests.
        :param pulumi.Input[str] response_chunking: Specifies how to handle chunked and unchunked responses.
        :param pulumi.Input[List[pulumi.Input[str]]] response_headers_permitteds: Specifies headers that the BIG-IP system allows in an HTTP response.
        :param pulumi.Input[str] server_agent_name: Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
               string is specified, then no Server header will be added to such responses
        :param pulumi.Input[str] tm_partition: Displays the administrative partition within which this profile resides.
        :param pulumi.Input[str] via_host_name: Specifies the hostname to include into Via header
        :param pulumi.Input[str] via_request: Specifies whether to append, remove, or preserve a Via header in an HTTP request
        :param pulumi.Input[str] via_response: Specifies whether to append, remove, or preserve a Via header in an HTTP request
        :param pulumi.Input[List[pulumi.Input[str]]] xff_alternative_names: Specifies alternative XFF headers instead of the default X-forwarded-for header
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accept_xff: Optional[pulumi.Input[str]] = None,
            app_service: Optional[pulumi.Input[str]] = None,
            basic_auth_realm: Optional[pulumi.Input[str]] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encrypt_cookie_secret: Optional[pulumi.Input[str]] = None,
            encrypt_cookies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            fallback_host: Optional[pulumi.Input[str]] = None,
            fallback_status_codes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            head_erase: Optional[pulumi.Input[str]] = None,
            head_insert: Optional[pulumi.Input[str]] = None,
            insert_xforwarded_for: Optional[pulumi.Input[str]] = None,
            lws_separator: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            oneconnect_transformations: Optional[pulumi.Input[str]] = None,
            proxy_type: Optional[pulumi.Input[str]] = None,
            redirect_rewrite: Optional[pulumi.Input[str]] = None,
            request_chunking: Optional[pulumi.Input[str]] = None,
            response_chunking: Optional[pulumi.Input[str]] = None,
            response_headers_permitteds: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            server_agent_name: Optional[pulumi.Input[str]] = None,
            tm_partition: Optional[pulumi.Input[str]] = None,
            via_host_name: Optional[pulumi.Input[str]] = None,
            via_request: Optional[pulumi.Input[str]] = None,
            via_response: Optional[pulumi.Input[str]] = None,
            xff_alternative_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'ProfileHttp':
        """
        Get an existing ProfileHttp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accept_xff: Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
               XFF (X-forwarded-for) headers, if they exist.
        :param pulumi.Input[str] app_service: The application service to which the object belongs.
        :param pulumi.Input[str] basic_auth_realm: Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] description: User defibned description
        :param pulumi.Input[str] encrypt_cookie_secret: Specifies a passphrase for the cookie encryption
        :param pulumi.Input[List[pulumi.Input[str]]] encrypt_cookies: Encrypts specified cookies that the BIG-IP system sends to a client system
        :param pulumi.Input[str] fallback_host: Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        :param pulumi.Input[List[pulumi.Input[str]]] fallback_status_codes: Specifies one or more three-digit status codes that can be returned by an HTTP server.
        :param pulumi.Input[str] head_erase: Specifies the header string that you want to erase from an HTTP request. You can also specify none
        :param pulumi.Input[str] head_insert: Specifies a quoted header string that you want to insert into an HTTP request
        :param pulumi.Input[str] insert_xforwarded_for: When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        :param pulumi.Input[str] lws_separator: Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
        :param pulumi.Input[str] name: Name of the profile_http
        :param pulumi.Input[str] oneconnect_transformations: Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
        :param pulumi.Input[str] proxy_type: Specifies the type of HTTP proxy.
        :param pulumi.Input[str] redirect_rewrite: Specifies which of the application HTTP redirects the system rewrites to HTTPS.
        :param pulumi.Input[str] request_chunking: Specifies how to handle chunked and unchunked requests.
        :param pulumi.Input[str] response_chunking: Specifies how to handle chunked and unchunked responses.
        :param pulumi.Input[List[pulumi.Input[str]]] response_headers_permitteds: Specifies headers that the BIG-IP system allows in an HTTP response.
        :param pulumi.Input[str] server_agent_name: Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
               string is specified, then no Server header will be added to such responses
        :param pulumi.Input[str] tm_partition: Displays the administrative partition within which this profile resides.
        :param pulumi.Input[str] via_host_name: Specifies the hostname to include into Via header
        :param pulumi.Input[str] via_request: Specifies whether to append, remove, or preserve a Via header in an HTTP request
        :param pulumi.Input[str] via_response: Specifies whether to append, remove, or preserve a Via header in an HTTP request
        :param pulumi.Input[List[pulumi.Input[str]]] xff_alternative_names: Specifies alternative XFF headers instead of the default X-forwarded-for header
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

    @property
    @pulumi.getter(name="acceptXff")
    def accept_xff(self) -> Optional[str]:
        """
        Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
        XFF (X-forwarded-for) headers, if they exist.
        """
        return pulumi.get(self, "accept_xff")

    @property
    @pulumi.getter(name="appService")
    def app_service(self) -> Optional[str]:
        """
        The application service to which the object belongs.
        """
        return pulumi.get(self, "app_service")

    @property
    @pulumi.getter(name="basicAuthRealm")
    def basic_auth_realm(self) -> Optional[str]:
        """
        Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        """
        return pulumi.get(self, "basic_auth_realm")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> str:
        """
        Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        User defibned description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptCookieSecret")
    def encrypt_cookie_secret(self) -> Optional[str]:
        """
        Specifies a passphrase for the cookie encryption
        """
        return pulumi.get(self, "encrypt_cookie_secret")

    @property
    @pulumi.getter(name="encryptCookies")
    def encrypt_cookies(self) -> Optional[List[str]]:
        """
        Encrypts specified cookies that the BIG-IP system sends to a client system
        """
        return pulumi.get(self, "encrypt_cookies")

    @property
    @pulumi.getter(name="fallbackHost")
    def fallback_host(self) -> Optional[str]:
        """
        Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        """
        return pulumi.get(self, "fallback_host")

    @property
    @pulumi.getter(name="fallbackStatusCodes")
    def fallback_status_codes(self) -> Optional[List[str]]:
        """
        Specifies one or more three-digit status codes that can be returned by an HTTP server.
        """
        return pulumi.get(self, "fallback_status_codes")

    @property
    @pulumi.getter(name="headErase")
    def head_erase(self) -> Optional[str]:
        """
        Specifies the header string that you want to erase from an HTTP request. You can also specify none
        """
        return pulumi.get(self, "head_erase")

    @property
    @pulumi.getter(name="headInsert")
    def head_insert(self) -> Optional[str]:
        """
        Specifies a quoted header string that you want to insert into an HTTP request
        """
        return pulumi.get(self, "head_insert")

    @property
    @pulumi.getter(name="insertXforwardedFor")
    def insert_xforwarded_for(self) -> Optional[str]:
        """
        When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        """
        return pulumi.get(self, "insert_xforwarded_for")

    @property
    @pulumi.getter(name="lwsSeparator")
    def lws_separator(self) -> Optional[str]:
        """
        Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
        """
        return pulumi.get(self, "lws_separator")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the profile_http
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="oneconnectTransformations")
    def oneconnect_transformations(self) -> Optional[str]:
        """
        Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
        """
        return pulumi.get(self, "oneconnect_transformations")

    @property
    @pulumi.getter(name="proxyType")
    def proxy_type(self) -> Optional[str]:
        """
        Specifies the type of HTTP proxy.
        """
        return pulumi.get(self, "proxy_type")

    @property
    @pulumi.getter(name="redirectRewrite")
    def redirect_rewrite(self) -> Optional[str]:
        """
        Specifies which of the application HTTP redirects the system rewrites to HTTPS.
        """
        return pulumi.get(self, "redirect_rewrite")

    @property
    @pulumi.getter(name="requestChunking")
    def request_chunking(self) -> Optional[str]:
        """
        Specifies how to handle chunked and unchunked requests.
        """
        return pulumi.get(self, "request_chunking")

    @property
    @pulumi.getter(name="responseChunking")
    def response_chunking(self) -> Optional[str]:
        """
        Specifies how to handle chunked and unchunked responses.
        """
        return pulumi.get(self, "response_chunking")

    @property
    @pulumi.getter(name="responseHeadersPermitteds")
    def response_headers_permitteds(self) -> Optional[List[str]]:
        """
        Specifies headers that the BIG-IP system allows in an HTTP response.
        """
        return pulumi.get(self, "response_headers_permitteds")

    @property
    @pulumi.getter(name="serverAgentName")
    def server_agent_name(self) -> Optional[str]:
        """
        Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
        string is specified, then no Server header will be added to such responses
        """
        return pulumi.get(self, "server_agent_name")

    @property
    @pulumi.getter(name="tmPartition")
    def tm_partition(self) -> Optional[str]:
        """
        Displays the administrative partition within which this profile resides.
        """
        return pulumi.get(self, "tm_partition")

    @property
    @pulumi.getter(name="viaHostName")
    def via_host_name(self) -> Optional[str]:
        """
        Specifies the hostname to include into Via header
        """
        return pulumi.get(self, "via_host_name")

    @property
    @pulumi.getter(name="viaRequest")
    def via_request(self) -> Optional[str]:
        """
        Specifies whether to append, remove, or preserve a Via header in an HTTP request
        """
        return pulumi.get(self, "via_request")

    @property
    @pulumi.getter(name="viaResponse")
    def via_response(self) -> Optional[str]:
        """
        Specifies whether to append, remove, or preserve a Via header in an HTTP request
        """
        return pulumi.get(self, "via_response")

    @property
    @pulumi.getter(name="xffAlternativeNames")
    def xff_alternative_names(self) -> Optional[List[str]]:
        """
        Specifies alternative XFF headers instead of the default X-forwarded-for header
        """
        return pulumi.get(self, "xff_alternative_names")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

