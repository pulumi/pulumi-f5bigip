# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ProvisionArgs', 'Provision']

@pulumi.input_type
class ProvisionArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 cpu_ratio: Optional[pulumi.Input[int]] = None,
                 disk_ratio: Optional[pulumi.Input[int]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 memory_ratio: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Provision resource.
        :param pulumi.Input[str] name: Name of module to provision in BIG-IP. 
               possible options:
               * afm
               * am
               * apm
               * cgnat
               * asm
               * avr
               * dos
               * fps
               * gtm
               * ilx
               * lc
               * ltm
               * pem
               * sslo
               * swg
               * urldb
        :param pulumi.Input[int] cpu_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[int] disk_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[str] level: Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
               default is `nominal`
               possible options:
               * nominal
               * minimum
               * none
               * dedicated
        :param pulumi.Input[int] memory_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        ProvisionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            cpu_ratio=cpu_ratio,
            disk_ratio=disk_ratio,
            full_path=full_path,
            level=level,
            memory_ratio=memory_ratio,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: pulumi.Input[str],
             cpu_ratio: Optional[pulumi.Input[int]] = None,
             disk_ratio: Optional[pulumi.Input[int]] = None,
             full_path: Optional[pulumi.Input[str]] = None,
             level: Optional[pulumi.Input[str]] = None,
             memory_ratio: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("name", name)
        if cpu_ratio is not None:
            _setter("cpu_ratio", cpu_ratio)
        if disk_ratio is not None:
            _setter("disk_ratio", disk_ratio)
        if full_path is not None:
            _setter("full_path", full_path)
        if level is not None:
            _setter("level", level)
        if memory_ratio is not None:
            _setter("memory_ratio", memory_ratio)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of module to provision in BIG-IP. 
        possible options:
        * afm
        * am
        * apm
        * cgnat
        * asm
        * avr
        * dos
        * fps
        * gtm
        * ilx
        * lc
        * ltm
        * pem
        * sslo
        * swg
        * urldb
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="cpuRatio")
    def cpu_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "cpu_ratio")

    @cpu_ratio.setter
    def cpu_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu_ratio", value)

    @property
    @pulumi.getter(name="diskRatio")
    def disk_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "disk_ratio")

    @disk_ratio.setter
    def disk_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_ratio", value)

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "full_path")

    @full_path.setter
    def full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_path", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
        default is `nominal`
        possible options:
        * nominal
        * minimum
        * none
        * dedicated
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter(name="memoryRatio")
    def memory_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "memory_ratio")

    @memory_ratio.setter
    def memory_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_ratio", value)


@pulumi.input_type
class _ProvisionState:
    def __init__(__self__, *,
                 cpu_ratio: Optional[pulumi.Input[int]] = None,
                 disk_ratio: Optional[pulumi.Input[int]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 memory_ratio: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Provision resources.
        :param pulumi.Input[int] cpu_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[int] disk_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[str] level: Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
               default is `nominal`
               possible options:
               * nominal
               * minimum
               * none
               * dedicated
        :param pulumi.Input[int] memory_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[str] name: Name of module to provision in BIG-IP. 
               possible options:
               * afm
               * am
               * apm
               * cgnat
               * asm
               * avr
               * dos
               * fps
               * gtm
               * ilx
               * lc
               * ltm
               * pem
               * sslo
               * swg
               * urldb
        """
        _ProvisionState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            cpu_ratio=cpu_ratio,
            disk_ratio=disk_ratio,
            full_path=full_path,
            level=level,
            memory_ratio=memory_ratio,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             cpu_ratio: Optional[pulumi.Input[int]] = None,
             disk_ratio: Optional[pulumi.Input[int]] = None,
             full_path: Optional[pulumi.Input[str]] = None,
             level: Optional[pulumi.Input[str]] = None,
             memory_ratio: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if cpu_ratio is not None:
            _setter("cpu_ratio", cpu_ratio)
        if disk_ratio is not None:
            _setter("disk_ratio", disk_ratio)
        if full_path is not None:
            _setter("full_path", full_path)
        if level is not None:
            _setter("level", level)
        if memory_ratio is not None:
            _setter("memory_ratio", memory_ratio)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="cpuRatio")
    def cpu_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "cpu_ratio")

    @cpu_ratio.setter
    def cpu_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cpu_ratio", value)

    @property
    @pulumi.getter(name="diskRatio")
    def disk_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "disk_ratio")

    @disk_ratio.setter
    def disk_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disk_ratio", value)

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "full_path")

    @full_path.setter
    def full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_path", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        """
        Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
        default is `nominal`
        possible options:
        * nominal
        * minimum
        * none
        * dedicated
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter(name="memoryRatio")
    def memory_ratio(self) -> Optional[pulumi.Input[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "memory_ratio")

    @memory_ratio.setter
    def memory_ratio(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_ratio", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of module to provision in BIG-IP. 
        possible options:
        * afm
        * am
        * apm
        * cgnat
        * asm
        * avr
        * dos
        * fps
        * gtm
        * ilx
        * lc
        * ltm
        * pem
        * sslo
        * swg
        * urldb
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Provision(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu_ratio: Optional[pulumi.Input[int]] = None,
                 disk_ratio: Optional[pulumi.Input[int]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 memory_ratio: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `sys.Provision` Manage BIG-IP module provisioning. This resource will only provision at the standard levels of Dedicated, Nominal, and Minimum.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        gtm = f5bigip.sys.Provision("gtm",
            cpu_ratio=0,
            disk_ratio=0,
            level="nominal",
            memory_ratio=0,
            name="gtm")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cpu_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[int] disk_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[str] level: Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
               default is `nominal`
               possible options:
               * nominal
               * minimum
               * none
               * dedicated
        :param pulumi.Input[int] memory_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[str] name: Name of module to provision in BIG-IP. 
               possible options:
               * afm
               * am
               * apm
               * cgnat
               * asm
               * avr
               * dos
               * fps
               * gtm
               * ilx
               * lc
               * ltm
               * pem
               * sslo
               * swg
               * urldb
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProvisionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `sys.Provision` Manage BIG-IP module provisioning. This resource will only provision at the standard levels of Dedicated, Nominal, and Minimum.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        gtm = f5bigip.sys.Provision("gtm",
            cpu_ratio=0,
            disk_ratio=0,
            level="nominal",
            memory_ratio=0,
            name="gtm")
        ```

        :param str resource_name: The name of the resource.
        :param ProvisionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProvisionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProvisionArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu_ratio: Optional[pulumi.Input[int]] = None,
                 disk_ratio: Optional[pulumi.Input[int]] = None,
                 full_path: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 memory_ratio: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProvisionArgs.__new__(ProvisionArgs)

            __props__.__dict__["cpu_ratio"] = cpu_ratio
            __props__.__dict__["disk_ratio"] = disk_ratio
            __props__.__dict__["full_path"] = full_path
            __props__.__dict__["level"] = level
            __props__.__dict__["memory_ratio"] = memory_ratio
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
        super(Provision, __self__).__init__(
            'f5bigip:sys/provision:Provision',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cpu_ratio: Optional[pulumi.Input[int]] = None,
            disk_ratio: Optional[pulumi.Input[int]] = None,
            full_path: Optional[pulumi.Input[str]] = None,
            level: Optional[pulumi.Input[str]] = None,
            memory_ratio: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Provision':
        """
        Get an existing Provision resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cpu_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[int] disk_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[str] level: Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
               default is `nominal`
               possible options:
               * nominal
               * minimum
               * none
               * dedicated
        :param pulumi.Input[int] memory_ratio: Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        :param pulumi.Input[str] name: Name of module to provision in BIG-IP. 
               possible options:
               * afm
               * am
               * apm
               * cgnat
               * asm
               * avr
               * dos
               * fps
               * gtm
               * ilx
               * lc
               * ltm
               * pem
               * sslo
               * swg
               * urldb
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProvisionState.__new__(_ProvisionState)

        __props__.__dict__["cpu_ratio"] = cpu_ratio
        __props__.__dict__["disk_ratio"] = disk_ratio
        __props__.__dict__["full_path"] = full_path
        __props__.__dict__["level"] = level
        __props__.__dict__["memory_ratio"] = memory_ratio
        __props__.__dict__["name"] = name
        return Provision(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cpuRatio")
    def cpu_ratio(self) -> pulumi.Output[Optional[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "cpu_ratio")

    @property
    @pulumi.getter(name="diskRatio")
    def disk_ratio(self) -> pulumi.Output[Optional[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "disk_ratio")

    @property
    @pulumi.getter(name="fullPath")
    def full_path(self) -> pulumi.Output[str]:
        return pulumi.get(self, "full_path")

    @property
    @pulumi.getter
    def level(self) -> pulumi.Output[Optional[str]]:
        """
        Sets the provisioning level for the requested modules. Changing the level for one module may require modifying the level of another module. For example, changing one module to `dedicated` requires setting all others to `none`. Setting the level of a module to `none` means the module is not activated.
        default is `nominal`
        possible options:
        * nominal
        * minimum
        * none
        * dedicated
        """
        return pulumi.get(self, "level")

    @property
    @pulumi.getter(name="memoryRatio")
    def memory_ratio(self) -> pulumi.Output[Optional[int]]:
        """
        Use this option only when the level option is set to custom.F5 Networks recommends that you do not modify this option. The default value is none
        """
        return pulumi.get(self, "memory_ratio")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of module to provision in BIG-IP. 
        possible options:
        * afm
        * am
        * apm
        * cgnat
        * asm
        * avr
        * dos
        * fps
        * gtm
        * ilx
        * lc
        * ltm
        * pem
        * sslo
        * swg
        * urldb
        """
        return pulumi.get(self, "name")

