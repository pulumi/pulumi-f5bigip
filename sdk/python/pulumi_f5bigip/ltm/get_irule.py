# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetIruleResult',
    'AwaitableGetIruleResult',
    'get_irule',
    'get_irule_output',
]

@pulumi.output_type
class GetIruleResult:
    """
    A collection of values returned by getIrule.
    """
    def __init__(__self__, id=None, irule=None, name=None, partition=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if irule and not isinstance(irule, str):
            raise TypeError("Expected argument 'irule' to be a str")
        pulumi.set(__self__, "irule", irule)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        pulumi.set(__self__, "partition", partition)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def irule(self) -> Optional[str]:
        """
        Irule configured on bigip
        """
        return pulumi.get(self, "irule")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of irule configured on bigip with full path
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partition(self) -> str:
        """
        Bigip partition in which rule is configured
        """
        return pulumi.get(self, "partition")


class AwaitableGetIruleResult(GetIruleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIruleResult(
            id=self.id,
            irule=self.irule,
            name=self.name,
            partition=self.partition)


def get_irule(irule: Optional[str] = None,
              name: Optional[str] = None,
              partition: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIruleResult:
    """
    Use this data source (`ltm.IRule`) to get the ltm irule details available on BIG-IP

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    test = f5bigip.ltm.get_irule(name="terraform_irule",
        partition="Common")
    pulumi.export("bigipIrule", test.irule)
    ```
    <!--End PulumiCodeChooser -->


    :param str irule: Irule configured on bigip
    :param str name: Name of the irule
    :param str partition: partition of the ltm irule
    """
    __args__ = dict()
    __args__['irule'] = irule
    __args__['name'] = name
    __args__['partition'] = partition
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('f5bigip:ltm/getIrule:getIrule', __args__, opts=opts, typ=GetIruleResult).value

    return AwaitableGetIruleResult(
        id=pulumi.get(__ret__, 'id'),
        irule=pulumi.get(__ret__, 'irule'),
        name=pulumi.get(__ret__, 'name'),
        partition=pulumi.get(__ret__, 'partition'))


@_utilities.lift_output_func(get_irule)
def get_irule_output(irule: Optional[pulumi.Input[Optional[str]]] = None,
                     name: Optional[pulumi.Input[str]] = None,
                     partition: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIruleResult]:
    """
    Use this data source (`ltm.IRule`) to get the ltm irule details available on BIG-IP

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_f5bigip as f5bigip

    test = f5bigip.ltm.get_irule(name="terraform_irule",
        partition="Common")
    pulumi.export("bigipIrule", test.irule)
    ```
    <!--End PulumiCodeChooser -->


    :param str irule: Irule configured on bigip
    :param str name: Name of the irule
    :param str partition: partition of the ltm irule
    """
    ...
