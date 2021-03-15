# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Do']


class Do(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 do_json: Optional[pulumi.Input[str]] = None,
                 tenant_name: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `Do` provides details about bigip do resource

        This resource is helpful to configure do declarative JSON on BIG-IP.
        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        do_example = f5bigip.Do("do-example",
            do_json=(lambda path: open(path).read())("example.json"),
            timeout=15)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] do_json: Name of the of the Declarative DO JSON file
        :param pulumi.Input[str] tenant_name: unique identifier for DO resource
        :param pulumi.Input[int] timeout: DO json
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

            if do_json is None and not opts.urn:
                raise TypeError("Missing required property 'do_json'")
            __props__['do_json'] = do_json
            if tenant_name is not None and not opts.urn:
                warnings.warn("""this attribute is no longer in use""", DeprecationWarning)
                pulumi.log.warn("""tenant_name is deprecated: this attribute is no longer in use""")
            __props__['tenant_name'] = tenant_name
            __props__['timeout'] = timeout
        super(Do, __self__).__init__(
            'f5bigip:index/do:Do',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            do_json: Optional[pulumi.Input[str]] = None,
            tenant_name: Optional[pulumi.Input[str]] = None,
            timeout: Optional[pulumi.Input[int]] = None) -> 'Do':
        """
        Get an existing Do resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] do_json: Name of the of the Declarative DO JSON file
        :param pulumi.Input[str] tenant_name: unique identifier for DO resource
        :param pulumi.Input[int] timeout: DO json
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["do_json"] = do_json
        __props__["tenant_name"] = tenant_name
        __props__["timeout"] = timeout
        return Do(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="doJson")
    def do_json(self) -> pulumi.Output[str]:
        """
        Name of the of the Declarative DO JSON file
        """
        return pulumi.get(self, "do_json")

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> pulumi.Output[Optional[str]]:
        """
        unique identifier for DO resource
        """
        return pulumi.get(self, "tenant_name")

    @property
    @pulumi.getter
    def timeout(self) -> pulumi.Output[Optional[int]]:
        """
        DO json
        """
        return pulumi.get(self, "timeout")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

