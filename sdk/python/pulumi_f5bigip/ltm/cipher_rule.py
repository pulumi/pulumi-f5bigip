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

__all__ = ['CipherRuleArgs', 'CipherRule']

@pulumi.input_type
class CipherRuleArgs:
    def __init__(__self__, *,
                 cipher: pulumi.Input[builtins.str],
                 name: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dh_groups: Optional[pulumi.Input[builtins.str]] = None,
                 signature_algorithms: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a CipherRule resource.
        :param pulumi.Input[builtins.str] cipher: Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        :param pulumi.Input[builtins.str] name: Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        :param pulumi.Input[builtins.str] description: The Partition in which the Cipher Rule will be created.
        :param pulumi.Input[builtins.str] dh_groups: Specifies the DH Groups algorithms, separated by colons (:).
        :param pulumi.Input[builtins.str] signature_algorithms: Specifies the Signature Algorithms, separated by colons (:).
        """
        pulumi.set(__self__, "cipher", cipher)
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dh_groups is not None:
            pulumi.set(__self__, "dh_groups", dh_groups)
        if signature_algorithms is not None:
            pulumi.set(__self__, "signature_algorithms", signature_algorithms)

    @property
    @pulumi.getter
    def cipher(self) -> pulumi.Input[builtins.str]:
        """
        Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        """
        return pulumi.get(self, "cipher")

    @cipher.setter
    def cipher(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cipher", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Partition in which the Cipher Rule will be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dhGroups")
    def dh_groups(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the DH Groups algorithms, separated by colons (:).
        """
        return pulumi.get(self, "dh_groups")

    @dh_groups.setter
    def dh_groups(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dh_groups", value)

    @property
    @pulumi.getter(name="signatureAlgorithms")
    def signature_algorithms(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the Signature Algorithms, separated by colons (:).
        """
        return pulumi.get(self, "signature_algorithms")

    @signature_algorithms.setter
    def signature_algorithms(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "signature_algorithms", value)


@pulumi.input_type
class _CipherRuleState:
    def __init__(__self__, *,
                 cipher: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dh_groups: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 signature_algorithms: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering CipherRule resources.
        :param pulumi.Input[builtins.str] cipher: Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        :param pulumi.Input[builtins.str] description: The Partition in which the Cipher Rule will be created.
        :param pulumi.Input[builtins.str] dh_groups: Specifies the DH Groups algorithms, separated by colons (:).
        :param pulumi.Input[builtins.str] name: Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        :param pulumi.Input[builtins.str] signature_algorithms: Specifies the Signature Algorithms, separated by colons (:).
        """
        if cipher is not None:
            pulumi.set(__self__, "cipher", cipher)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dh_groups is not None:
            pulumi.set(__self__, "dh_groups", dh_groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if signature_algorithms is not None:
            pulumi.set(__self__, "signature_algorithms", signature_algorithms)

    @property
    @pulumi.getter
    def cipher(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        """
        return pulumi.get(self, "cipher")

    @cipher.setter
    def cipher(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cipher", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Partition in which the Cipher Rule will be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dhGroups")
    def dh_groups(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the DH Groups algorithms, separated by colons (:).
        """
        return pulumi.get(self, "dh_groups")

    @dh_groups.setter
    def dh_groups(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "dh_groups", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="signatureAlgorithms")
    def signature_algorithms(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the Signature Algorithms, separated by colons (:).
        """
        return pulumi.get(self, "signature_algorithms")

    @signature_algorithms.setter
    def signature_algorithms(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "signature_algorithms", value)


class CipherRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cipher: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dh_groups: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 signature_algorithms: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        `ltm.CipherRule` Manages F5 BIG-IP LTM cipher rule using iControl REST.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_cipher_rule = f5bigip.ltm.CipherRule("test_cipher_rule",
            name="/Common/test_cipher_rule",
            cipher="TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384",
            dh_groups="P256:P384:FFDHE2048:FFDHE3072:FFDHE4096",
            signature_algorithms="DEFAULT")
        ```

        ## Importing

        An existing cipher rule can be imported into this resource by supplying the cipher rule full path name  ex : `/partition/name`
        An example is below:
        ```sh
        $ terraform import bigip_ltm_cipher_rule.test_cipher_rule /Common/test_cipher_rule
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cipher: Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        :param pulumi.Input[builtins.str] description: The Partition in which the Cipher Rule will be created.
        :param pulumi.Input[builtins.str] dh_groups: Specifies the DH Groups algorithms, separated by colons (:).
        :param pulumi.Input[builtins.str] name: Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        :param pulumi.Input[builtins.str] signature_algorithms: Specifies the Signature Algorithms, separated by colons (:).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CipherRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.CipherRule` Manages F5 BIG-IP LTM cipher rule using iControl REST.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_cipher_rule = f5bigip.ltm.CipherRule("test_cipher_rule",
            name="/Common/test_cipher_rule",
            cipher="TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384",
            dh_groups="P256:P384:FFDHE2048:FFDHE3072:FFDHE4096",
            signature_algorithms="DEFAULT")
        ```

        ## Importing

        An existing cipher rule can be imported into this resource by supplying the cipher rule full path name  ex : `/partition/name`
        An example is below:
        ```sh
        $ terraform import bigip_ltm_cipher_rule.test_cipher_rule /Common/test_cipher_rule
        ```

        :param str resource_name: The name of the resource.
        :param CipherRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CipherRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cipher: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 dh_groups: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 signature_algorithms: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CipherRuleArgs.__new__(CipherRuleArgs)

            if cipher is None and not opts.urn:
                raise TypeError("Missing required property 'cipher'")
            __props__.__dict__["cipher"] = cipher
            __props__.__dict__["description"] = description
            __props__.__dict__["dh_groups"] = dh_groups
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["signature_algorithms"] = signature_algorithms
        super(CipherRule, __self__).__init__(
            'f5bigip:ltm/cipherRule:CipherRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cipher: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            dh_groups: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            signature_algorithms: Optional[pulumi.Input[builtins.str]] = None) -> 'CipherRule':
        """
        Get an existing CipherRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cipher: Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        :param pulumi.Input[builtins.str] description: The Partition in which the Cipher Rule will be created.
        :param pulumi.Input[builtins.str] dh_groups: Specifies the DH Groups algorithms, separated by colons (:).
        :param pulumi.Input[builtins.str] name: Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        :param pulumi.Input[builtins.str] signature_algorithms: Specifies the Signature Algorithms, separated by colons (:).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CipherRuleState.__new__(_CipherRuleState)

        __props__.__dict__["cipher"] = cipher
        __props__.__dict__["description"] = description
        __props__.__dict__["dh_groups"] = dh_groups
        __props__.__dict__["name"] = name
        __props__.__dict__["signature_algorithms"] = signature_algorithms
        return CipherRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cipher(self) -> pulumi.Output[builtins.str]:
        """
        Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
        """
        return pulumi.get(self, "cipher")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The Partition in which the Cipher Rule will be created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dhGroups")
    def dh_groups(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the DH Groups algorithms, separated by colons (:).
        """
        return pulumi.get(self, "dh_groups")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the Cipher Rule. Name should be in pattern `partition` + `cipher_rule_name`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="signatureAlgorithms")
    def signature_algorithms(self) -> pulumi.Output[builtins.str]:
        """
        Specifies the Signature Algorithms, separated by colons (:).
        """
        return pulumi.get(self, "signature_algorithms")

