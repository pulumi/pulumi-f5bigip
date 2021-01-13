# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ProfileHttpCompress']


class ProfileHttpCompress(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type_excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 content_type_includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 defaults_from: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 uri_excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 uri_includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.ProfileHttpCompress`  Virtual server HTTP compression profile configuration

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        sjhttpcompression = f5bigip.ltm.ProfileHttpCompress("sjhttpcompression",
            content_type_excludes=["nicecontentexclude.com"],
            content_type_includes=["nicecontent.com"],
            defaults_from="/Common/httpcompression",
            name="/Common/sjhttpcompression2",
            uri_excludes=[
                "www.abc.f5.com",
                "www.abc2.f5.com",
            ],
            uri_includes=["www.xyzbc.cisco.com"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] content_type_excludes: Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] content_type_includes: Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] name: Name of the profile_httpcompress
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_excludes: Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_includes: Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
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

            __props__['content_type_excludes'] = content_type_excludes
            __props__['content_type_includes'] = content_type_includes
            __props__['defaults_from'] = defaults_from
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['uri_excludes'] = uri_excludes
            __props__['uri_includes'] = uri_includes
        super(ProfileHttpCompress, __self__).__init__(
            'f5bigip:ltm/profileHttpCompress:ProfileHttpCompress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            content_type_excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            content_type_includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            defaults_from: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            uri_excludes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            uri_includes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ProfileHttpCompress':
        """
        Get an existing ProfileHttpCompress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] content_type_excludes: Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] content_type_includes: Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        :param pulumi.Input[str] defaults_from: Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        :param pulumi.Input[str] name: Name of the profile_httpcompress
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_excludes: Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_includes: Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["content_type_excludes"] = content_type_excludes
        __props__["content_type_includes"] = content_type_includes
        __props__["defaults_from"] = defaults_from
        __props__["name"] = name
        __props__["uri_excludes"] = uri_excludes
        __props__["uri_includes"] = uri_includes
        return ProfileHttpCompress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentTypeExcludes")
    def content_type_excludes(self) -> pulumi.Output[Sequence[str]]:
        """
        Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        """
        return pulumi.get(self, "content_type_excludes")

    @property
    @pulumi.getter(name="contentTypeIncludes")
    def content_type_includes(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        """
        return pulumi.get(self, "content_type_includes")

    @property
    @pulumi.getter(name="defaultsFrom")
    def defaults_from(self) -> pulumi.Output[str]:
        """
        Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        """
        return pulumi.get(self, "defaults_from")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the profile_httpcompress
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="uriExcludes")
    def uri_excludes(self) -> pulumi.Output[Sequence[str]]:
        """
        Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
        """
        return pulumi.get(self, "uri_excludes")

    @property
    @pulumi.getter(name="uriIncludes")
    def uri_includes(self) -> pulumi.Output[Sequence[str]]:
        """
        Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
        """
        return pulumi.get(self, "uri_includes")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

