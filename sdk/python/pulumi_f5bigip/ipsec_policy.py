# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IpsecPolicyArgs', 'IpsecPolicy']

@pulumi.input_type
class IpsecPolicyArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_algorithm: Optional[pulumi.Input[str]] = None,
                 ipcomp: Optional[pulumi.Input[str]] = None,
                 kb_lifetime: Optional[pulumi.Input[int]] = None,
                 lifetime: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 perfect_forward_secrecy: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 tunnel_local_address: Optional[pulumi.Input[str]] = None,
                 tunnel_remote_address: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpsecPolicy resource.
        :param pulumi.Input[str] name: Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        :param pulumi.Input[str] auth_algorithm: Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
               aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        :param pulumi.Input[str] description: Description of the IPSec policy.
        :param pulumi.Input[str] encrypt_algorithm: Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
               aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        :param pulumi.Input[str] ipcomp: Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        :param pulumi.Input[int] kb_lifetime: Specifies the length of time before the IKE security association expires, in kilobytes.
        :param pulumi.Input[int] lifetime: Specifies the length of time before the IKE security association expires, in minutes.
        :param pulumi.Input[str] mode: Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        :param pulumi.Input[str] perfect_forward_secrecy: Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
               modp4096, modp6144, modp8192`
        :param pulumi.Input[str] protocol: Specifies the IPsec protocol. Valid choices are: `ah, esp`
        :param pulumi.Input[str] tunnel_local_address: Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        :param pulumi.Input[str] tunnel_remote_address: Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        IpsecPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            auth_algorithm=auth_algorithm,
            description=description,
            encrypt_algorithm=encrypt_algorithm,
            ipcomp=ipcomp,
            kb_lifetime=kb_lifetime,
            lifetime=lifetime,
            mode=mode,
            perfect_forward_secrecy=perfect_forward_secrecy,
            protocol=protocol,
            tunnel_local_address=tunnel_local_address,
            tunnel_remote_address=tunnel_remote_address,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: pulumi.Input[str],
             auth_algorithm: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             encrypt_algorithm: Optional[pulumi.Input[str]] = None,
             ipcomp: Optional[pulumi.Input[str]] = None,
             kb_lifetime: Optional[pulumi.Input[int]] = None,
             lifetime: Optional[pulumi.Input[int]] = None,
             mode: Optional[pulumi.Input[str]] = None,
             perfect_forward_secrecy: Optional[pulumi.Input[str]] = None,
             protocol: Optional[pulumi.Input[str]] = None,
             tunnel_local_address: Optional[pulumi.Input[str]] = None,
             tunnel_remote_address: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("name", name)
        if auth_algorithm is not None:
            _setter("auth_algorithm", auth_algorithm)
        if description is not None:
            _setter("description", description)
        if encrypt_algorithm is not None:
            _setter("encrypt_algorithm", encrypt_algorithm)
        if ipcomp is not None:
            _setter("ipcomp", ipcomp)
        if kb_lifetime is not None:
            _setter("kb_lifetime", kb_lifetime)
        if lifetime is not None:
            _setter("lifetime", lifetime)
        if mode is not None:
            _setter("mode", mode)
        if perfect_forward_secrecy is not None:
            _setter("perfect_forward_secrecy", perfect_forward_secrecy)
        if protocol is not None:
            _setter("protocol", protocol)
        if tunnel_local_address is not None:
            _setter("tunnel_local_address", tunnel_local_address)
        if tunnel_remote_address is not None:
            _setter("tunnel_remote_address", tunnel_remote_address)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="authAlgorithm")
    def auth_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
        aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        """
        return pulumi.get(self, "auth_algorithm")

    @auth_algorithm.setter
    def auth_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_algorithm", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the IPSec policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptAlgorithm")
    def encrypt_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
        aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        """
        return pulumi.get(self, "encrypt_algorithm")

    @encrypt_algorithm.setter
    def encrypt_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encrypt_algorithm", value)

    @property
    @pulumi.getter
    def ipcomp(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        """
        return pulumi.get(self, "ipcomp")

    @ipcomp.setter
    def ipcomp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipcomp", value)

    @property
    @pulumi.getter(name="kbLifetime")
    def kb_lifetime(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the length of time before the IKE security association expires, in kilobytes.
        """
        return pulumi.get(self, "kb_lifetime")

    @kb_lifetime.setter
    def kb_lifetime(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kb_lifetime", value)

    @property
    @pulumi.getter
    def lifetime(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the length of time before the IKE security association expires, in minutes.
        """
        return pulumi.get(self, "lifetime")

    @lifetime.setter
    def lifetime(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lifetime", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="perfectForwardSecrecy")
    def perfect_forward_secrecy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
        modp4096, modp6144, modp8192`
        """
        return pulumi.get(self, "perfect_forward_secrecy")

    @perfect_forward_secrecy.setter
    def perfect_forward_secrecy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "perfect_forward_secrecy", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IPsec protocol. Valid choices are: `ah, esp`
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="tunnelLocalAddress")
    def tunnel_local_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        return pulumi.get(self, "tunnel_local_address")

    @tunnel_local_address.setter
    def tunnel_local_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tunnel_local_address", value)

    @property
    @pulumi.getter(name="tunnelRemoteAddress")
    def tunnel_remote_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        return pulumi.get(self, "tunnel_remote_address")

    @tunnel_remote_address.setter
    def tunnel_remote_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tunnel_remote_address", value)


@pulumi.input_type
class _IpsecPolicyState:
    def __init__(__self__, *,
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_algorithm: Optional[pulumi.Input[str]] = None,
                 ipcomp: Optional[pulumi.Input[str]] = None,
                 kb_lifetime: Optional[pulumi.Input[int]] = None,
                 lifetime: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 perfect_forward_secrecy: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 tunnel_local_address: Optional[pulumi.Input[str]] = None,
                 tunnel_remote_address: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpsecPolicy resources.
        :param pulumi.Input[str] auth_algorithm: Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
               aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        :param pulumi.Input[str] description: Description of the IPSec policy.
        :param pulumi.Input[str] encrypt_algorithm: Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
               aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        :param pulumi.Input[str] ipcomp: Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        :param pulumi.Input[int] kb_lifetime: Specifies the length of time before the IKE security association expires, in kilobytes.
        :param pulumi.Input[int] lifetime: Specifies the length of time before the IKE security association expires, in minutes.
        :param pulumi.Input[str] mode: Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        :param pulumi.Input[str] name: Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        :param pulumi.Input[str] perfect_forward_secrecy: Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
               modp4096, modp6144, modp8192`
        :param pulumi.Input[str] protocol: Specifies the IPsec protocol. Valid choices are: `ah, esp`
        :param pulumi.Input[str] tunnel_local_address: Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        :param pulumi.Input[str] tunnel_remote_address: Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        _IpsecPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            auth_algorithm=auth_algorithm,
            description=description,
            encrypt_algorithm=encrypt_algorithm,
            ipcomp=ipcomp,
            kb_lifetime=kb_lifetime,
            lifetime=lifetime,
            mode=mode,
            name=name,
            perfect_forward_secrecy=perfect_forward_secrecy,
            protocol=protocol,
            tunnel_local_address=tunnel_local_address,
            tunnel_remote_address=tunnel_remote_address,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             auth_algorithm: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             encrypt_algorithm: Optional[pulumi.Input[str]] = None,
             ipcomp: Optional[pulumi.Input[str]] = None,
             kb_lifetime: Optional[pulumi.Input[int]] = None,
             lifetime: Optional[pulumi.Input[int]] = None,
             mode: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             perfect_forward_secrecy: Optional[pulumi.Input[str]] = None,
             protocol: Optional[pulumi.Input[str]] = None,
             tunnel_local_address: Optional[pulumi.Input[str]] = None,
             tunnel_remote_address: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if auth_algorithm is not None:
            _setter("auth_algorithm", auth_algorithm)
        if description is not None:
            _setter("description", description)
        if encrypt_algorithm is not None:
            _setter("encrypt_algorithm", encrypt_algorithm)
        if ipcomp is not None:
            _setter("ipcomp", ipcomp)
        if kb_lifetime is not None:
            _setter("kb_lifetime", kb_lifetime)
        if lifetime is not None:
            _setter("lifetime", lifetime)
        if mode is not None:
            _setter("mode", mode)
        if name is not None:
            _setter("name", name)
        if perfect_forward_secrecy is not None:
            _setter("perfect_forward_secrecy", perfect_forward_secrecy)
        if protocol is not None:
            _setter("protocol", protocol)
        if tunnel_local_address is not None:
            _setter("tunnel_local_address", tunnel_local_address)
        if tunnel_remote_address is not None:
            _setter("tunnel_remote_address", tunnel_remote_address)

    @property
    @pulumi.getter(name="authAlgorithm")
    def auth_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
        aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        """
        return pulumi.get(self, "auth_algorithm")

    @auth_algorithm.setter
    def auth_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_algorithm", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the IPSec policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptAlgorithm")
    def encrypt_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
        aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        """
        return pulumi.get(self, "encrypt_algorithm")

    @encrypt_algorithm.setter
    def encrypt_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encrypt_algorithm", value)

    @property
    @pulumi.getter
    def ipcomp(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        """
        return pulumi.get(self, "ipcomp")

    @ipcomp.setter
    def ipcomp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipcomp", value)

    @property
    @pulumi.getter(name="kbLifetime")
    def kb_lifetime(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the length of time before the IKE security association expires, in kilobytes.
        """
        return pulumi.get(self, "kb_lifetime")

    @kb_lifetime.setter
    def kb_lifetime(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "kb_lifetime", value)

    @property
    @pulumi.getter
    def lifetime(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the length of time before the IKE security association expires, in minutes.
        """
        return pulumi.get(self, "lifetime")

    @lifetime.setter
    def lifetime(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lifetime", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="perfectForwardSecrecy")
    def perfect_forward_secrecy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
        modp4096, modp6144, modp8192`
        """
        return pulumi.get(self, "perfect_forward_secrecy")

    @perfect_forward_secrecy.setter
    def perfect_forward_secrecy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "perfect_forward_secrecy", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the IPsec protocol. Valid choices are: `ah, esp`
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="tunnelLocalAddress")
    def tunnel_local_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        return pulumi.get(self, "tunnel_local_address")

    @tunnel_local_address.setter
    def tunnel_local_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tunnel_local_address", value)

    @property
    @pulumi.getter(name="tunnelRemoteAddress")
    def tunnel_remote_address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        return pulumi.get(self, "tunnel_remote_address")

    @tunnel_remote_address.setter
    def tunnel_remote_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tunnel_remote_address", value)


class IpsecPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_algorithm: Optional[pulumi.Input[str]] = None,
                 ipcomp: Optional[pulumi.Input[str]] = None,
                 kb_lifetime: Optional[pulumi.Input[int]] = None,
                 lifetime: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 perfect_forward_secrecy: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 tunnel_local_address: Optional[pulumi.Input[str]] = None,
                 tunnel_remote_address: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `IpsecPolicy` Manage IPSec policies on a BIG-IP

        Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-policy)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_policy = f5bigip.IpsecPolicy("test-policy",
            auth_algorithm="sha1",
            description="created by terraform provider",
            encrypt_algorithm="3des",
            ipcomp="deflate",
            lifetime=3,
            mode="tunnel",
            name="/Common/test-policy",
            protocol="esp",
            tunnel_local_address="192.168.1.1",
            tunnel_remote_address="10.10.1.1")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_algorithm: Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
               aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        :param pulumi.Input[str] description: Description of the IPSec policy.
        :param pulumi.Input[str] encrypt_algorithm: Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
               aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        :param pulumi.Input[str] ipcomp: Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        :param pulumi.Input[int] kb_lifetime: Specifies the length of time before the IKE security association expires, in kilobytes.
        :param pulumi.Input[int] lifetime: Specifies the length of time before the IKE security association expires, in minutes.
        :param pulumi.Input[str] mode: Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        :param pulumi.Input[str] name: Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        :param pulumi.Input[str] perfect_forward_secrecy: Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
               modp4096, modp6144, modp8192`
        :param pulumi.Input[str] protocol: Specifies the IPsec protocol. Valid choices are: `ah, esp`
        :param pulumi.Input[str] tunnel_local_address: Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        :param pulumi.Input[str] tunnel_remote_address: Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpsecPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `IpsecPolicy` Manage IPSec policies on a BIG-IP

        Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-policy)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip

        test_policy = f5bigip.IpsecPolicy("test-policy",
            auth_algorithm="sha1",
            description="created by terraform provider",
            encrypt_algorithm="3des",
            ipcomp="deflate",
            lifetime=3,
            mode="tunnel",
            name="/Common/test-policy",
            protocol="esp",
            tunnel_local_address="192.168.1.1",
            tunnel_remote_address="10.10.1.1")
        ```

        :param str resource_name: The name of the resource.
        :param IpsecPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpsecPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IpsecPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_algorithm: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encrypt_algorithm: Optional[pulumi.Input[str]] = None,
                 ipcomp: Optional[pulumi.Input[str]] = None,
                 kb_lifetime: Optional[pulumi.Input[int]] = None,
                 lifetime: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 perfect_forward_secrecy: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 tunnel_local_address: Optional[pulumi.Input[str]] = None,
                 tunnel_remote_address: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpsecPolicyArgs.__new__(IpsecPolicyArgs)

            __props__.__dict__["auth_algorithm"] = auth_algorithm
            __props__.__dict__["description"] = description
            __props__.__dict__["encrypt_algorithm"] = encrypt_algorithm
            __props__.__dict__["ipcomp"] = ipcomp
            __props__.__dict__["kb_lifetime"] = kb_lifetime
            __props__.__dict__["lifetime"] = lifetime
            __props__.__dict__["mode"] = mode
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["perfect_forward_secrecy"] = perfect_forward_secrecy
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["tunnel_local_address"] = tunnel_local_address
            __props__.__dict__["tunnel_remote_address"] = tunnel_remote_address
        super(IpsecPolicy, __self__).__init__(
            'f5bigip:index/ipsecPolicy:IpsecPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_algorithm: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encrypt_algorithm: Optional[pulumi.Input[str]] = None,
            ipcomp: Optional[pulumi.Input[str]] = None,
            kb_lifetime: Optional[pulumi.Input[int]] = None,
            lifetime: Optional[pulumi.Input[int]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            perfect_forward_secrecy: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            tunnel_local_address: Optional[pulumi.Input[str]] = None,
            tunnel_remote_address: Optional[pulumi.Input[str]] = None) -> 'IpsecPolicy':
        """
        Get an existing IpsecPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_algorithm: Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
               aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        :param pulumi.Input[str] description: Description of the IPSec policy.
        :param pulumi.Input[str] encrypt_algorithm: Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
               aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        :param pulumi.Input[str] ipcomp: Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        :param pulumi.Input[int] kb_lifetime: Specifies the length of time before the IKE security association expires, in kilobytes.
        :param pulumi.Input[int] lifetime: Specifies the length of time before the IKE security association expires, in minutes.
        :param pulumi.Input[str] mode: Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        :param pulumi.Input[str] name: Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        :param pulumi.Input[str] perfect_forward_secrecy: Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
               modp4096, modp6144, modp8192`
        :param pulumi.Input[str] protocol: Specifies the IPsec protocol. Valid choices are: `ah, esp`
        :param pulumi.Input[str] tunnel_local_address: Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        :param pulumi.Input[str] tunnel_remote_address: Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpsecPolicyState.__new__(_IpsecPolicyState)

        __props__.__dict__["auth_algorithm"] = auth_algorithm
        __props__.__dict__["description"] = description
        __props__.__dict__["encrypt_algorithm"] = encrypt_algorithm
        __props__.__dict__["ipcomp"] = ipcomp
        __props__.__dict__["kb_lifetime"] = kb_lifetime
        __props__.__dict__["lifetime"] = lifetime
        __props__.__dict__["mode"] = mode
        __props__.__dict__["name"] = name
        __props__.__dict__["perfect_forward_secrecy"] = perfect_forward_secrecy
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["tunnel_local_address"] = tunnel_local_address
        __props__.__dict__["tunnel_remote_address"] = tunnel_remote_address
        return IpsecPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authAlgorithm")
    def auth_algorithm(self) -> pulumi.Output[str]:
        """
        Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
        aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
        """
        return pulumi.get(self, "auth_algorithm")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the IPSec policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptAlgorithm")
    def encrypt_algorithm(self) -> pulumi.Output[str]:
        """
        Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
        aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
        """
        return pulumi.get(self, "encrypt_algorithm")

    @property
    @pulumi.getter
    def ipcomp(self) -> pulumi.Output[str]:
        """
        Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
        """
        return pulumi.get(self, "ipcomp")

    @property
    @pulumi.getter(name="kbLifetime")
    def kb_lifetime(self) -> pulumi.Output[int]:
        """
        Specifies the length of time before the IKE security association expires, in kilobytes.
        """
        return pulumi.get(self, "kb_lifetime")

    @property
    @pulumi.getter
    def lifetime(self) -> pulumi.Output[int]:
        """
        Specifies the length of time before the IKE security association expires, in minutes.
        """
        return pulumi.get(self, "lifetime")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[str]:
        """
        Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="perfectForwardSecrecy")
    def perfect_forward_secrecy(self) -> pulumi.Output[str]:
        """
        Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
        modp4096, modp6144, modp8192`
        """
        return pulumi.get(self, "perfect_forward_secrecy")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Specifies the IPsec protocol. Valid choices are: `ah, esp`
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="tunnelLocalAddress")
    def tunnel_local_address(self) -> pulumi.Output[str]:
        """
        Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        return pulumi.get(self, "tunnel_local_address")

    @property
    @pulumi.getter(name="tunnelRemoteAddress")
    def tunnel_remote_address(self) -> pulumi.Output[str]:
        """
        Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
        """
        return pulumi.get(self, "tunnel_remote_address")

