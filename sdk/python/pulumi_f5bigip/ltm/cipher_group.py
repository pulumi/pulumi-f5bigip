# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['CipherGroupArgs', 'CipherGroup']

@pulumi.input_type
class CipherGroupArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[_builtins.str],
                 allows: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 ordering: Optional[pulumi.Input[_builtins.str]] = None,
                 requires: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        The set of arguments for constructing a CipherGroup resource.
        :param pulumi.Input[_builtins.str] name: Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] allows: Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        :param pulumi.Input[_builtins.str] description: Specifies descriptive text that identifies the cipher rule
        :param pulumi.Input[_builtins.str] ordering: Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] requires: Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        pulumi.set(__self__, "name", name)
        if allows is not None:
            pulumi.set(__self__, "allows", allows)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ordering is not None:
            pulumi.set(__self__, "ordering", ordering)
        if requires is not None:
            pulumi.set(__self__, "requires", requires)

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def allows(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        return pulumi.get(self, "allows")

    @allows.setter
    def allows(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "allows", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies descriptive text that identifies the cipher rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def ordering(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        """
        return pulumi.get(self, "ordering")

    @ordering.setter
    def ordering(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "ordering", value)

    @_builtins.property
    @pulumi.getter
    def requires(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        return pulumi.get(self, "requires")

    @requires.setter
    def requires(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "requires", value)


@pulumi.input_type
class _CipherGroupState:
    def __init__(__self__, *,
                 allows: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 ordering: Optional[pulumi.Input[_builtins.str]] = None,
                 requires: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering CipherGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] allows: Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        :param pulumi.Input[_builtins.str] description: Specifies descriptive text that identifies the cipher rule
        :param pulumi.Input[_builtins.str] name: Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        :param pulumi.Input[_builtins.str] ordering: Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] requires: Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        if allows is not None:
            pulumi.set(__self__, "allows", allows)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ordering is not None:
            pulumi.set(__self__, "ordering", ordering)
        if requires is not None:
            pulumi.set(__self__, "requires", requires)

    @_builtins.property
    @pulumi.getter
    def allows(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        return pulumi.get(self, "allows")

    @allows.setter
    def allows(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "allows", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies descriptive text that identifies the cipher rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def ordering(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        """
        return pulumi.get(self, "ordering")

    @ordering.setter
    def ordering(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "ordering", value)

    @_builtins.property
    @pulumi.getter
    def requires(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        return pulumi.get(self, "requires")

    @requires.setter
    def requires(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "requires", value)


@pulumi.type_token("f5bigip:ltm/cipherGroup:CipherGroup")
class CipherGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allows: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 ordering: Optional[pulumi.Input[_builtins.str]] = None,
                 requires: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        `ltm.CipherGroup` Manages F5 BIG-IP LTM cipher group using iControl REST.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_cipher_group = f5bigip.ltm.CipherGroup("test-cipher-group",
            name="/Common/test-cipher-group-01",
            allows=["/Common/f5-aes"],
            requires=["/Common/f5-quic"],
            ordering="speed")
        ```

        ## Importing

        An existing cipher group can be imported into this resource by supplying the cipher rule full path name ex : `/partition/name`
        An example is below:
        ```sh
        $ terraform import bigip_ltm_cipher_group.test_cipher_group /Common/test_cipher_group

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] allows: Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        :param pulumi.Input[_builtins.str] description: Specifies descriptive text that identifies the cipher rule
        :param pulumi.Input[_builtins.str] name: Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        :param pulumi.Input[_builtins.str] ordering: Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] requires: Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CipherGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.CipherGroup` Manages F5 BIG-IP LTM cipher group using iControl REST.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_cipher_group = f5bigip.ltm.CipherGroup("test-cipher-group",
            name="/Common/test-cipher-group-01",
            allows=["/Common/f5-aes"],
            requires=["/Common/f5-quic"],
            ordering="speed")
        ```

        ## Importing

        An existing cipher group can be imported into this resource by supplying the cipher rule full path name ex : `/partition/name`
        An example is below:
        ```sh
        $ terraform import bigip_ltm_cipher_group.test_cipher_group /Common/test_cipher_group

        ```

        :param str resource_name: The name of the resource.
        :param CipherGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CipherGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allows: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 ordering: Optional[pulumi.Input[_builtins.str]] = None,
                 requires: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CipherGroupArgs.__new__(CipherGroupArgs)

            __props__.__dict__["allows"] = allows
            __props__.__dict__["description"] = description
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["ordering"] = ordering
            __props__.__dict__["requires"] = requires
        super(CipherGroup, __self__).__init__(
            'f5bigip:ltm/cipherGroup:CipherGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allows: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            ordering: Optional[pulumi.Input[_builtins.str]] = None,
            requires: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None) -> 'CipherGroup':
        """
        Get an existing CipherGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] allows: Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        :param pulumi.Input[_builtins.str] description: Specifies descriptive text that identifies the cipher rule
        :param pulumi.Input[_builtins.str] name: Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        :param pulumi.Input[_builtins.str] ordering: Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] requires: Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CipherGroupState.__new__(_CipherGroupState)

        __props__.__dict__["allows"] = allows
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["ordering"] = ordering
        __props__.__dict__["requires"] = requires
        return CipherGroup(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def allows(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        return pulumi.get(self, "allows")

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies descriptive text that identifies the cipher rule
        """
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the Cipher group. Name should be in pattern `partition` + `cipher_group_name`
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def ordering(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
        """
        return pulumi.get(self, "ordering")

    @_builtins.property
    @pulumi.getter
    def requires(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
        """
        return pulumi.get(self, "requires")

