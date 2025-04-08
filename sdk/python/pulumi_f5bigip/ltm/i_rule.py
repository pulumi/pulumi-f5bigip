# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['IRuleArgs', 'IRule']

@pulumi.input_type
class IRuleArgs:
    def __init__(__self__, *,
                 irule: pulumi.Input[builtins.str],
                 name: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a IRule resource.
        :param pulumi.Input[builtins.str] irule: Body of the iRule
        :param pulumi.Input[builtins.str] name: Name of the iRule
        """
        pulumi.set(__self__, "irule", irule)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def irule(self) -> pulumi.Input[builtins.str]:
        """
        Body of the iRule
        """
        return pulumi.get(self, "irule")

    @irule.setter
    def irule(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "irule", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the iRule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IRuleState:
    def __init__(__self__, *,
                 irule: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering IRule resources.
        :param pulumi.Input[builtins.str] irule: Body of the iRule
        :param pulumi.Input[builtins.str] name: Name of the iRule
        """
        if irule is not None:
            pulumi.set(__self__, "irule", irule)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def irule(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Body of the iRule
        """
        return pulumi.get(self, "irule")

    @irule.setter
    def irule(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "irule", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the iRule
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


class IRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 irule: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        `ltm.IRule` Creates iRule on BIG-IP F5 device

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip
        import pulumi_std as std

        # Loading from a file is the preferred method
        rule = f5bigip.ltm.IRule("rule",
            name="/Common/terraform_irule",
            irule=std.file(input="myirule.tcl").result)
        rule2 = f5bigip.ltm.IRule("rule2",
            name="/Common/terraform_irule2",
            irule=\"\"\"when CLIENT_ACCEPTED {
             log local0. "test"
           }
        \"\"\")
        ```

        ##myirule.tcl

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] irule: Body of the iRule
        :param pulumi.Input[builtins.str] name: Name of the iRule
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.IRule` Creates iRule on BIG-IP F5 device

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip
        import pulumi_std as std

        # Loading from a file is the preferred method
        rule = f5bigip.ltm.IRule("rule",
            name="/Common/terraform_irule",
            irule=std.file(input="myirule.tcl").result)
        rule2 = f5bigip.ltm.IRule("rule2",
            name="/Common/terraform_irule2",
            irule=\"\"\"when CLIENT_ACCEPTED {
             log local0. "test"
           }
        \"\"\")
        ```

        ##myirule.tcl

        :param str resource_name: The name of the resource.
        :param IRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 irule: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IRuleArgs.__new__(IRuleArgs)

            if irule is None and not opts.urn:
                raise TypeError("Missing required property 'irule'")
            __props__.__dict__["irule"] = irule
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
        super(IRule, __self__).__init__(
            'f5bigip:ltm/iRule:IRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            irule: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None) -> 'IRule':
        """
        Get an existing IRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] irule: Body of the iRule
        :param pulumi.Input[builtins.str] name: Name of the iRule
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IRuleState.__new__(_IRuleState)

        __props__.__dict__["irule"] = irule
        __props__.__dict__["name"] = name
        return IRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def irule(self) -> pulumi.Output[builtins.str]:
        """
        Body of the iRule
        """
        return pulumi.get(self, "irule")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the iRule
        """
        return pulumi.get(self, "name")

