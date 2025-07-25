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
from . import outputs
from ._inputs import *

__all__ = ['SnatArgs', 'Snat']

@pulumi.input_type
class SnatArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[_builtins.str],
                 origins: pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]],
                 autolasthop: Optional[pulumi.Input[_builtins.str]] = None,
                 full_path: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 partition: Optional[pulumi.Input[_builtins.str]] = None,
                 snatpool: Optional[pulumi.Input[_builtins.str]] = None,
                 sourceport: Optional[pulumi.Input[_builtins.str]] = None,
                 translation: Optional[pulumi.Input[_builtins.str]] = None,
                 vlans: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 vlansdisabled: Optional[pulumi.Input[_builtins.bool]] = None):
        """
        The set of arguments for constructing a Snat resource.
        :param pulumi.Input[_builtins.str] name: Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        :param pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]] origins: Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
        :param pulumi.Input[_builtins.str] autolasthop: Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
        :param pulumi.Input[_builtins.str] full_path: Fullpath
        :param pulumi.Input[_builtins.str] mirror: Enables or disables mirroring of SNAT connections.
        :param pulumi.Input[_builtins.str] partition: Partition or path to which the SNAT belongs
        :param pulumi.Input[_builtins.str] snatpool: Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
        :param pulumi.Input[_builtins.str] sourceport: Specifies how the SNAT object handles the client's source port. The default is `preserve`.
        :param pulumi.Input[_builtins.str] translation: Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] vlans: Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
        :param pulumi.Input[_builtins.bool] vlansdisabled: Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "origins", origins)
        if autolasthop is not None:
            pulumi.set(__self__, "autolasthop", autolasthop)
        if full_path is not None:
            pulumi.set(__self__, "full_path", full_path)
        if mirror is not None:
            pulumi.set(__self__, "mirror", mirror)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if snatpool is not None:
            pulumi.set(__self__, "snatpool", snatpool)
        if sourceport is not None:
            pulumi.set(__self__, "sourceport", sourceport)
        if translation is not None:
            pulumi.set(__self__, "translation", translation)
        if vlans is not None:
            pulumi.set(__self__, "vlans", vlans)
        if vlansdisabled is not None:
            pulumi.set(__self__, "vlansdisabled", vlansdisabled)

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def origins(self) -> pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]]:
        """
        Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
        """
        return pulumi.get(self, "origins")

    @origins.setter
    def origins(self, value: pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]]):
        pulumi.set(self, "origins", value)

    @_builtins.property
    @pulumi.getter
    def autolasthop(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
        """
        return pulumi.get(self, "autolasthop")

    @autolasthop.setter
    def autolasthop(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "autolasthop", value)

    @_builtins.property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Fullpath
        """
        return pulumi.get(self, "full_path")

    @full_path.setter
    def full_path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "full_path", value)

    @_builtins.property
    @pulumi.getter
    def mirror(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Enables or disables mirroring of SNAT connections.
        """
        return pulumi.get(self, "mirror")

    @mirror.setter
    def mirror(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mirror", value)

    @_builtins.property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Partition or path to which the SNAT belongs
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "partition", value)

    @_builtins.property
    @pulumi.getter
    def snatpool(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
        """
        return pulumi.get(self, "snatpool")

    @snatpool.setter
    def snatpool(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "snatpool", value)

    @_builtins.property
    @pulumi.getter
    def sourceport(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies how the SNAT object handles the client's source port. The default is `preserve`.
        """
        return pulumi.get(self, "sourceport")

    @sourceport.setter
    def sourceport(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "sourceport", value)

    @_builtins.property
    @pulumi.getter
    def translation(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
        """
        return pulumi.get(self, "translation")

    @translation.setter
    def translation(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "translation", value)

    @_builtins.property
    @pulumi.getter
    def vlans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
        """
        return pulumi.get(self, "vlans")

    @vlans.setter
    def vlans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "vlans", value)

    @_builtins.property
    @pulumi.getter
    def vlansdisabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
        """
        return pulumi.get(self, "vlansdisabled")

    @vlansdisabled.setter
    def vlansdisabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "vlansdisabled", value)


@pulumi.input_type
class _SnatState:
    def __init__(__self__, *,
                 autolasthop: Optional[pulumi.Input[_builtins.str]] = None,
                 full_path: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 origins: Optional[pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]]] = None,
                 partition: Optional[pulumi.Input[_builtins.str]] = None,
                 snatpool: Optional[pulumi.Input[_builtins.str]] = None,
                 sourceport: Optional[pulumi.Input[_builtins.str]] = None,
                 translation: Optional[pulumi.Input[_builtins.str]] = None,
                 vlans: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 vlansdisabled: Optional[pulumi.Input[_builtins.bool]] = None):
        """
        Input properties used for looking up and filtering Snat resources.
        :param pulumi.Input[_builtins.str] autolasthop: Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
        :param pulumi.Input[_builtins.str] full_path: Fullpath
        :param pulumi.Input[_builtins.str] mirror: Enables or disables mirroring of SNAT connections.
        :param pulumi.Input[_builtins.str] name: Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        :param pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]] origins: Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
        :param pulumi.Input[_builtins.str] partition: Partition or path to which the SNAT belongs
        :param pulumi.Input[_builtins.str] snatpool: Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
        :param pulumi.Input[_builtins.str] sourceport: Specifies how the SNAT object handles the client's source port. The default is `preserve`.
        :param pulumi.Input[_builtins.str] translation: Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] vlans: Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
        :param pulumi.Input[_builtins.bool] vlansdisabled: Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
        """
        if autolasthop is not None:
            pulumi.set(__self__, "autolasthop", autolasthop)
        if full_path is not None:
            pulumi.set(__self__, "full_path", full_path)
        if mirror is not None:
            pulumi.set(__self__, "mirror", mirror)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if origins is not None:
            pulumi.set(__self__, "origins", origins)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if snatpool is not None:
            pulumi.set(__self__, "snatpool", snatpool)
        if sourceport is not None:
            pulumi.set(__self__, "sourceport", sourceport)
        if translation is not None:
            pulumi.set(__self__, "translation", translation)
        if vlans is not None:
            pulumi.set(__self__, "vlans", vlans)
        if vlansdisabled is not None:
            pulumi.set(__self__, "vlansdisabled", vlansdisabled)

    @_builtins.property
    @pulumi.getter
    def autolasthop(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
        """
        return pulumi.get(self, "autolasthop")

    @autolasthop.setter
    def autolasthop(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "autolasthop", value)

    @_builtins.property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Fullpath
        """
        return pulumi.get(self, "full_path")

    @full_path.setter
    def full_path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "full_path", value)

    @_builtins.property
    @pulumi.getter
    def mirror(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Enables or disables mirroring of SNAT connections.
        """
        return pulumi.get(self, "mirror")

    @mirror.setter
    def mirror(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mirror", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def origins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]]]:
        """
        Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
        """
        return pulumi.get(self, "origins")

    @origins.setter
    def origins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnatOriginArgs']]]]):
        pulumi.set(self, "origins", value)

    @_builtins.property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Partition or path to which the SNAT belongs
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "partition", value)

    @_builtins.property
    @pulumi.getter
    def snatpool(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
        """
        return pulumi.get(self, "snatpool")

    @snatpool.setter
    def snatpool(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "snatpool", value)

    @_builtins.property
    @pulumi.getter
    def sourceport(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies how the SNAT object handles the client's source port. The default is `preserve`.
        """
        return pulumi.get(self, "sourceport")

    @sourceport.setter
    def sourceport(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "sourceport", value)

    @_builtins.property
    @pulumi.getter
    def translation(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
        """
        return pulumi.get(self, "translation")

    @translation.setter
    def translation(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "translation", value)

    @_builtins.property
    @pulumi.getter
    def vlans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
        """
        return pulumi.get(self, "vlans")

    @vlans.setter
    def vlans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "vlans", value)

    @_builtins.property
    @pulumi.getter
    def vlansdisabled(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
        """
        return pulumi.get(self, "vlansdisabled")

    @vlansdisabled.setter
    def vlansdisabled(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "vlansdisabled", value)


@pulumi.type_token("f5bigip:ltm/snat:Snat")
class Snat(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autolasthop: Optional[pulumi.Input[_builtins.str]] = None,
                 full_path: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 origins: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SnatOriginArgs', 'SnatOriginArgsDict']]]]] = None,
                 partition: Optional[pulumi.Input[_builtins.str]] = None,
                 snatpool: Optional[pulumi.Input[_builtins.str]] = None,
                 sourceport: Optional[pulumi.Input[_builtins.str]] = None,
                 translation: Optional[pulumi.Input[_builtins.str]] = None,
                 vlans: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 vlansdisabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 __props__=None):
        """
        `ltm.Snat` Manages a SNAT configuration

        For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource.For example `/Common/test-snat`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_snat = f5bigip.ltm.Snat("test-snat",
            name="/Common/test-snat",
            translation="/Common/136.1.1.2",
            sourceport="preserve",
            origins=[{
                "name": "0.0.0.0/0",
            }],
            vlans=["/Common/internal"],
            vlansdisabled=False)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] autolasthop: Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
        :param pulumi.Input[_builtins.str] full_path: Fullpath
        :param pulumi.Input[_builtins.str] mirror: Enables or disables mirroring of SNAT connections.
        :param pulumi.Input[_builtins.str] name: Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['SnatOriginArgs', 'SnatOriginArgsDict']]]] origins: Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
        :param pulumi.Input[_builtins.str] partition: Partition or path to which the SNAT belongs
        :param pulumi.Input[_builtins.str] snatpool: Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
        :param pulumi.Input[_builtins.str] sourceport: Specifies how the SNAT object handles the client's source port. The default is `preserve`.
        :param pulumi.Input[_builtins.str] translation: Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] vlans: Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
        :param pulumi.Input[_builtins.bool] vlansdisabled: Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnatArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `ltm.Snat` Manages a SNAT configuration

        For resources should be named with their `full path`. The full path is the combination of the `partition + name` of the resource.For example `/Common/test-snat`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_snat = f5bigip.ltm.Snat("test-snat",
            name="/Common/test-snat",
            translation="/Common/136.1.1.2",
            sourceport="preserve",
            origins=[{
                "name": "0.0.0.0/0",
            }],
            vlans=["/Common/internal"],
            vlansdisabled=False)
        ```

        :param str resource_name: The name of the resource.
        :param SnatArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnatArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autolasthop: Optional[pulumi.Input[_builtins.str]] = None,
                 full_path: Optional[pulumi.Input[_builtins.str]] = None,
                 mirror: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 origins: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SnatOriginArgs', 'SnatOriginArgsDict']]]]] = None,
                 partition: Optional[pulumi.Input[_builtins.str]] = None,
                 snatpool: Optional[pulumi.Input[_builtins.str]] = None,
                 sourceport: Optional[pulumi.Input[_builtins.str]] = None,
                 translation: Optional[pulumi.Input[_builtins.str]] = None,
                 vlans: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 vlansdisabled: Optional[pulumi.Input[_builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnatArgs.__new__(SnatArgs)

            __props__.__dict__["autolasthop"] = autolasthop
            __props__.__dict__["full_path"] = full_path
            __props__.__dict__["mirror"] = mirror
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if origins is None and not opts.urn:
                raise TypeError("Missing required property 'origins'")
            __props__.__dict__["origins"] = origins
            __props__.__dict__["partition"] = partition
            __props__.__dict__["snatpool"] = snatpool
            __props__.__dict__["sourceport"] = sourceport
            __props__.__dict__["translation"] = translation
            __props__.__dict__["vlans"] = vlans
            __props__.__dict__["vlansdisabled"] = vlansdisabled
        super(Snat, __self__).__init__(
            'f5bigip:ltm/snat:Snat',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autolasthop: Optional[pulumi.Input[_builtins.str]] = None,
            full_path: Optional[pulumi.Input[_builtins.str]] = None,
            mirror: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            origins: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SnatOriginArgs', 'SnatOriginArgsDict']]]]] = None,
            partition: Optional[pulumi.Input[_builtins.str]] = None,
            snatpool: Optional[pulumi.Input[_builtins.str]] = None,
            sourceport: Optional[pulumi.Input[_builtins.str]] = None,
            translation: Optional[pulumi.Input[_builtins.str]] = None,
            vlans: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            vlansdisabled: Optional[pulumi.Input[_builtins.bool]] = None) -> 'Snat':
        """
        Get an existing Snat resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] autolasthop: Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
        :param pulumi.Input[_builtins.str] full_path: Fullpath
        :param pulumi.Input[_builtins.str] mirror: Enables or disables mirroring of SNAT connections.
        :param pulumi.Input[_builtins.str] name: Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['SnatOriginArgs', 'SnatOriginArgsDict']]]] origins: Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
        :param pulumi.Input[_builtins.str] partition: Partition or path to which the SNAT belongs
        :param pulumi.Input[_builtins.str] snatpool: Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
        :param pulumi.Input[_builtins.str] sourceport: Specifies how the SNAT object handles the client's source port. The default is `preserve`.
        :param pulumi.Input[_builtins.str] translation: Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] vlans: Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
        :param pulumi.Input[_builtins.bool] vlansdisabled: Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnatState.__new__(_SnatState)

        __props__.__dict__["autolasthop"] = autolasthop
        __props__.__dict__["full_path"] = full_path
        __props__.__dict__["mirror"] = mirror
        __props__.__dict__["name"] = name
        __props__.__dict__["origins"] = origins
        __props__.__dict__["partition"] = partition
        __props__.__dict__["snatpool"] = snatpool
        __props__.__dict__["sourceport"] = sourceport
        __props__.__dict__["translation"] = translation
        __props__.__dict__["vlans"] = vlans
        __props__.__dict__["vlansdisabled"] = vlansdisabled
        return Snat(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def autolasthop(self) -> pulumi.Output[_builtins.str]:
        """
        Specifies whether to automatically map last hop for pools or not. The default is to use next level's default.
        """
        return pulumi.get(self, "autolasthop")

    @_builtins.property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Fullpath
        """
        return pulumi.get(self, "full_path")

    @_builtins.property
    @pulumi.getter
    def mirror(self) -> pulumi.Output[_builtins.str]:
        """
        Enables or disables mirroring of SNAT connections.
        """
        return pulumi.get(self, "mirror")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def origins(self) -> pulumi.Output[Sequence['outputs.SnatOrigin']]:
        """
        Specifies, for each SNAT that you create, the origin addresses that are to be members of that SNAT. Specify origin addresses by their IP addresses and service ports
        """
        return pulumi.get(self, "origins")

    @_builtins.property
    @pulumi.getter
    def partition(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Partition or path to which the SNAT belongs
        """
        return pulumi.get(self, "partition")

    @_builtins.property
    @pulumi.getter
    def snatpool(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the name of a SNAT pool. You can only use this option when `automap` and `translation` are not used.
        """
        return pulumi.get(self, "snatpool")

    @_builtins.property
    @pulumi.getter
    def sourceport(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies how the SNAT object handles the client's source port. The default is `preserve`.
        """
        return pulumi.get(self, "sourceport")

    @_builtins.property
    @pulumi.getter
    def translation(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the IP address configured for translation. Note that translated addresses are outside the traffic management system. You can only use this option when `automap` and `snatpool` are not used.
        """
        return pulumi.get(self, "translation")

    @_builtins.property
    @pulumi.getter
    def vlans(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Specifies the available VLANs or tunnels and those for which the SNAT is enabled or disabled.
        """
        return pulumi.get(self, "vlans")

    @_builtins.property
    @pulumi.getter
    def vlansdisabled(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Specifies the VLANs or tunnels for which the SNAT is enabled or disabled. The default is `true`, vlandisabled on VLANS specified by `vlans`,if set to `false` vlanEnabled set on VLANS specified by `vlans` .
        """
        return pulumi.get(self, "vlansdisabled")

