# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IpsecProfileArgs', 'IpsecProfile']

@pulumi.input_type
class IpsecProfileArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 parent_profile: Optional[pulumi.Input[str]] = None,
                 traffic_selector: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpsecProfile resource.
        :param pulumi.Input[str] name: Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        :param pulumi.Input[str] description: Specifies descriptive text that identifies the IPsec interface tunnel profile.
        :param pulumi.Input[str] parent_profile: Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        :param pulumi.Input[str] traffic_selector: Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        """
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if parent_profile is not None:
            pulumi.set(__self__, "parent_profile", parent_profile)
        if traffic_selector is not None:
            pulumi.set(__self__, "traffic_selector", traffic_selector)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies descriptive text that identifies the IPsec interface tunnel profile.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="parentProfile")
    def parent_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        """
        return pulumi.get(self, "parent_profile")

    @parent_profile.setter
    def parent_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_profile", value)

    @property
    @pulumi.getter(name="trafficSelector")
    def traffic_selector(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        """
        return pulumi.get(self, "traffic_selector")

    @traffic_selector.setter
    def traffic_selector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_selector", value)


@pulumi.input_type
class _IpsecProfileState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_profile: Optional[pulumi.Input[str]] = None,
                 traffic_selector: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpsecProfile resources.
        :param pulumi.Input[str] description: Specifies descriptive text that identifies the IPsec interface tunnel profile.
        :param pulumi.Input[str] name: Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        :param pulumi.Input[str] parent_profile: Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        :param pulumi.Input[str] traffic_selector: Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_profile is not None:
            pulumi.set(__self__, "parent_profile", parent_profile)
        if traffic_selector is not None:
            pulumi.set(__self__, "traffic_selector", traffic_selector)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies descriptive text that identifies the IPsec interface tunnel profile.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentProfile")
    def parent_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        """
        return pulumi.get(self, "parent_profile")

    @parent_profile.setter
    def parent_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_profile", value)

    @property
    @pulumi.getter(name="trafficSelector")
    def traffic_selector(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        """
        return pulumi.get(self, "traffic_selector")

    @traffic_selector.setter
    def traffic_selector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_selector", value)


class IpsecProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_profile: Optional[pulumi.Input[str]] = None,
                 traffic_selector: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `IpsecProfile` Manage IPSec Profiles on a BIG-IP

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        azurev_wan_profile = f5bigip.IpsecProfile("azurevWANProfile",
            description="mytestipsecprofile",
            name="/Common/Mytestipsecprofile",
            traffic_selector="test-trafficselector")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies descriptive text that identifies the IPsec interface tunnel profile.
        :param pulumi.Input[str] name: Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        :param pulumi.Input[str] parent_profile: Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        :param pulumi.Input[str] traffic_selector: Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpsecProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `IpsecProfile` Manage IPSec Profiles on a BIG-IP

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        azurev_wan_profile = f5bigip.IpsecProfile("azurevWANProfile",
            description="mytestipsecprofile",
            name="/Common/Mytestipsecprofile",
            traffic_selector="test-trafficselector")
        ```

        :param str resource_name: The name of the resource.
        :param IpsecProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpsecProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_profile: Optional[pulumi.Input[str]] = None,
                 traffic_selector: Optional[pulumi.Input[str]] = None,
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
            __props__ = IpsecProfileArgs.__new__(IpsecProfileArgs)

            __props__.__dict__["description"] = description
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_profile"] = parent_profile
            __props__.__dict__["traffic_selector"] = traffic_selector
        super(IpsecProfile, __self__).__init__(
            'f5bigip:index/ipsecProfile:IpsecProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_profile: Optional[pulumi.Input[str]] = None,
            traffic_selector: Optional[pulumi.Input[str]] = None) -> 'IpsecProfile':
        """
        Get an existing IpsecProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies descriptive text that identifies the IPsec interface tunnel profile.
        :param pulumi.Input[str] name: Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        :param pulumi.Input[str] parent_profile: Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        :param pulumi.Input[str] traffic_selector: Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpsecProfileState.__new__(_IpsecProfileState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["parent_profile"] = parent_profile
        __props__.__dict__["traffic_selector"] = traffic_selector
        return IpsecProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Specifies descriptive text that identifies the IPsec interface tunnel profile.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentProfile")
    def parent_profile(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        """
        return pulumi.get(self, "parent_profile")

    @property
    @pulumi.getter(name="trafficSelector")
    def traffic_selector(self) -> pulumi.Output[str]:
        """
        Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        """
        return pulumi.get(self, "traffic_selector")

