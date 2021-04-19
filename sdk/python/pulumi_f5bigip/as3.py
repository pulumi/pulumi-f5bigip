# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['As3Args', 'As3']

@pulumi.input_type
class As3Args:
    def __init__(__self__, *,
                 as3_json: pulumi.Input[str],
                 application_list: Optional[pulumi.Input[str]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_filter: Optional[pulumi.Input[str]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None,
                 tenant_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a As3 resource.
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[str] application_list: Name of Application
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to false
        :param pulumi.Input[str] tenant_filter: If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
        :param pulumi.Input[str] tenant_list: Name of Tenant
        :param pulumi.Input[str] tenant_name: Name of Tenant
        """
        pulumi.set(__self__, "as3_json", as3_json)
        if application_list is not None:
            pulumi.set(__self__, "application_list", application_list)
        if ignore_metadata is not None:
            pulumi.set(__self__, "ignore_metadata", ignore_metadata)
        if tenant_filter is not None:
            pulumi.set(__self__, "tenant_filter", tenant_filter)
        if tenant_list is not None:
            pulumi.set(__self__, "tenant_list", tenant_list)
        if tenant_name is not None:
            warnings.warn("""this attribute is no longer in use""", DeprecationWarning)
            pulumi.log.warn("""tenant_name is deprecated: this attribute is no longer in use""")
        if tenant_name is not None:
            pulumi.set(__self__, "tenant_name", tenant_name)

    @property
    @pulumi.getter(name="as3Json")
    def as3_json(self) -> pulumi.Input[str]:
        """
        Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        """
        return pulumi.get(self, "as3_json")

    @as3_json.setter
    def as3_json(self, value: pulumi.Input[str]):
        pulumi.set(self, "as3_json", value)

    @property
    @pulumi.getter(name="applicationList")
    def application_list(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Application
        """
        return pulumi.get(self, "application_list")

    @application_list.setter
    def application_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_list", value)

    @property
    @pulumi.getter(name="ignoreMetadata")
    def ignore_metadata(self) -> Optional[pulumi.Input[bool]]:
        """
        Set True if you want to ignore metadata changes during update. By default it is set to false
        """
        return pulumi.get(self, "ignore_metadata")

    @ignore_metadata.setter
    def ignore_metadata(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_metadata", value)

    @property
    @pulumi.getter(name="tenantFilter")
    def tenant_filter(self) -> Optional[pulumi.Input[str]]:
        """
        If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
        """
        return pulumi.get(self, "tenant_filter")

    @tenant_filter.setter
    def tenant_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_filter", value)

    @property
    @pulumi.getter(name="tenantList")
    def tenant_list(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_list")

    @tenant_list.setter
    def tenant_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_list", value)

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_name")

    @tenant_name.setter
    def tenant_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_name", value)


@pulumi.input_type
class _As3State:
    def __init__(__self__, *,
                 application_list: Optional[pulumi.Input[str]] = None,
                 as3_json: Optional[pulumi.Input[str]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_filter: Optional[pulumi.Input[str]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None,
                 tenant_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering As3 resources.
        :param pulumi.Input[str] application_list: Name of Application
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to false
        :param pulumi.Input[str] tenant_filter: If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
        :param pulumi.Input[str] tenant_list: Name of Tenant
        :param pulumi.Input[str] tenant_name: Name of Tenant
        """
        if application_list is not None:
            pulumi.set(__self__, "application_list", application_list)
        if as3_json is not None:
            pulumi.set(__self__, "as3_json", as3_json)
        if ignore_metadata is not None:
            pulumi.set(__self__, "ignore_metadata", ignore_metadata)
        if tenant_filter is not None:
            pulumi.set(__self__, "tenant_filter", tenant_filter)
        if tenant_list is not None:
            pulumi.set(__self__, "tenant_list", tenant_list)
        if tenant_name is not None:
            warnings.warn("""this attribute is no longer in use""", DeprecationWarning)
            pulumi.log.warn("""tenant_name is deprecated: this attribute is no longer in use""")
        if tenant_name is not None:
            pulumi.set(__self__, "tenant_name", tenant_name)

    @property
    @pulumi.getter(name="applicationList")
    def application_list(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Application
        """
        return pulumi.get(self, "application_list")

    @application_list.setter
    def application_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_list", value)

    @property
    @pulumi.getter(name="as3Json")
    def as3_json(self) -> Optional[pulumi.Input[str]]:
        """
        Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        """
        return pulumi.get(self, "as3_json")

    @as3_json.setter
    def as3_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "as3_json", value)

    @property
    @pulumi.getter(name="ignoreMetadata")
    def ignore_metadata(self) -> Optional[pulumi.Input[bool]]:
        """
        Set True if you want to ignore metadata changes during update. By default it is set to false
        """
        return pulumi.get(self, "ignore_metadata")

    @ignore_metadata.setter
    def ignore_metadata(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_metadata", value)

    @property
    @pulumi.getter(name="tenantFilter")
    def tenant_filter(self) -> Optional[pulumi.Input[str]]:
        """
        If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
        """
        return pulumi.get(self, "tenant_filter")

    @tenant_filter.setter
    def tenant_filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_filter", value)

    @property
    @pulumi.getter(name="tenantList")
    def tenant_list(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_list")

    @tenant_list.setter
    def tenant_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_list", value)

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_name")

    @tenant_name.setter
    def tenant_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_name", value)


class As3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_list: Optional[pulumi.Input[str]] = None,
                 as3_json: Optional[pulumi.Input[str]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_filter: Optional[pulumi.Input[str]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None,
                 tenant_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `As3` provides details about bigip as3 resource

        This resource is helpful to configure as3 declarative JSON on BIG-IP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        # Example Usage for json file
        as3_example1_as3 = f5bigip.As3("as3-example1As3", as3_json=(lambda path: open(path).read())("example1.json"))
        # Example Usage for json file with tenant filter
        as3_example1_index_as3_as3 = f5bigip.As3("as3-example1Index/as3As3",
            as3_json=(lambda path: open(path).read())("example2.json"),
            tenant_filter="Sample_03")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_list: Name of Application
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to false
        :param pulumi.Input[str] tenant_filter: If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
        :param pulumi.Input[str] tenant_list: Name of Tenant
        :param pulumi.Input[str] tenant_name: Name of Tenant
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: As3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `As3` provides details about bigip as3 resource

        This resource is helpful to configure as3 declarative JSON on BIG-IP.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        # Example Usage for json file
        as3_example1_as3 = f5bigip.As3("as3-example1As3", as3_json=(lambda path: open(path).read())("example1.json"))
        # Example Usage for json file with tenant filter
        as3_example1_index_as3_as3 = f5bigip.As3("as3-example1Index/as3As3",
            as3_json=(lambda path: open(path).read())("example2.json"),
            tenant_filter="Sample_03")
        ```

        :param str resource_name: The name of the resource.
        :param As3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(As3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_list: Optional[pulumi.Input[str]] = None,
                 as3_json: Optional[pulumi.Input[str]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_filter: Optional[pulumi.Input[str]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None,
                 tenant_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = As3Args.__new__(As3Args)

            __props__.__dict__["application_list"] = application_list
            if as3_json is None and not opts.urn:
                raise TypeError("Missing required property 'as3_json'")
            __props__.__dict__["as3_json"] = as3_json
            __props__.__dict__["ignore_metadata"] = ignore_metadata
            __props__.__dict__["tenant_filter"] = tenant_filter
            __props__.__dict__["tenant_list"] = tenant_list
            if tenant_name is not None and not opts.urn:
                warnings.warn("""this attribute is no longer in use""", DeprecationWarning)
                pulumi.log.warn("""tenant_name is deprecated: this attribute is no longer in use""")
            __props__.__dict__["tenant_name"] = tenant_name
        super(As3, __self__).__init__(
            'f5bigip:index/as3:As3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_list: Optional[pulumi.Input[str]] = None,
            as3_json: Optional[pulumi.Input[str]] = None,
            ignore_metadata: Optional[pulumi.Input[bool]] = None,
            tenant_filter: Optional[pulumi.Input[str]] = None,
            tenant_list: Optional[pulumi.Input[str]] = None,
            tenant_name: Optional[pulumi.Input[str]] = None) -> 'As3':
        """
        Get an existing As3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_list: Name of Application
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to false
        :param pulumi.Input[str] tenant_filter: If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
        :param pulumi.Input[str] tenant_list: Name of Tenant
        :param pulumi.Input[str] tenant_name: Name of Tenant
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _As3State.__new__(_As3State)

        __props__.__dict__["application_list"] = application_list
        __props__.__dict__["as3_json"] = as3_json
        __props__.__dict__["ignore_metadata"] = ignore_metadata
        __props__.__dict__["tenant_filter"] = tenant_filter
        __props__.__dict__["tenant_list"] = tenant_list
        __props__.__dict__["tenant_name"] = tenant_name
        return As3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationList")
    def application_list(self) -> pulumi.Output[str]:
        """
        Name of Application
        """
        return pulumi.get(self, "application_list")

    @property
    @pulumi.getter(name="as3Json")
    def as3_json(self) -> pulumi.Output[str]:
        """
        Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        """
        return pulumi.get(self, "as3_json")

    @property
    @pulumi.getter(name="ignoreMetadata")
    def ignore_metadata(self) -> pulumi.Output[Optional[bool]]:
        """
        Set True if you want to ignore metadata changes during update. By default it is set to false
        """
        return pulumi.get(self, "ignore_metadata")

    @property
    @pulumi.getter(name="tenantFilter")
    def tenant_filter(self) -> pulumi.Output[Optional[str]]:
        """
        If there are muntiple tenants in a json this attribute helps the user to set a particular tenant to which he want to reflect the changes. Other tenants will neither be created nor be modified
        """
        return pulumi.get(self, "tenant_filter")

    @property
    @pulumi.getter(name="tenantList")
    def tenant_list(self) -> pulumi.Output[str]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_list")

    @property
    @pulumi.getter(name="tenantName")
    def tenant_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_name")

