# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Policy']


class Policy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 controls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 published_copy: Optional[pulumi.Input[str]] = None,
                 requires: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyRuleArgs']]]]] = None,
                 strategy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        `ltm.Policy` Configures Virtual Server

        For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_policy = f5bigip.ltm.Policy("test-policy",
            name="my_policy",
            strategy="first-match",
            requires=["http"],
            published_copy="Drafts/my_policy",
            controls=["forwarding"],
            rules=[f5bigip.ltm.PolicyRuleArgs(
                name="rule6",
                actions=[f5bigip.ltm.PolicyRuleActionArgs(
                    tm_name="20",
                    forward=True,
                    pool="/Common/mypool",
                )],
            )],
            opts=ResourceOptions(depends_on=[bigip_ltm_pool["mypool"]]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] controls: Specifies the controls
        :param pulumi.Input[str] name: Name of the Policy
        :param pulumi.Input[str] published_copy: If you want to publish the policy else it will be deployed in Drafts mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requires: Specifies the protocol
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyRuleArgs']]]] rules: Rules can be applied using the policy
        :param pulumi.Input[str] strategy: Specifies the match strategy
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

            __props__['controls'] = controls
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['published_copy'] = published_copy
            __props__['requires'] = requires
            __props__['rules'] = rules
            __props__['strategy'] = strategy
        super(Policy, __self__).__init__(
            'f5bigip:ltm/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            controls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            published_copy: Optional[pulumi.Input[str]] = None,
            requires: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyRuleArgs']]]]] = None,
            strategy: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] controls: Specifies the controls
        :param pulumi.Input[str] name: Name of the Policy
        :param pulumi.Input[str] published_copy: If you want to publish the policy else it will be deployed in Drafts mode.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requires: Specifies the protocol
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PolicyRuleArgs']]]] rules: Rules can be applied using the policy
        :param pulumi.Input[str] strategy: Specifies the match strategy
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["controls"] = controls
        __props__["name"] = name
        __props__["published_copy"] = published_copy
        __props__["requires"] = requires
        __props__["rules"] = rules
        __props__["strategy"] = strategy
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def controls(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies the controls
        """
        return pulumi.get(self, "controls")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the Policy
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publishedCopy")
    def published_copy(self) -> pulumi.Output[Optional[str]]:
        """
        If you want to publish the policy else it will be deployed in Drafts mode.
        """
        return pulumi.get(self, "published_copy")

    @property
    @pulumi.getter
    def requires(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies the protocol
        """
        return pulumi.get(self, "requires")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.PolicyRule']]]:
        """
        Rules can be applied using the policy
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def strategy(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the match strategy
        """
        return pulumi.get(self, "strategy")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

