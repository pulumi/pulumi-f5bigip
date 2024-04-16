# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SslKeyCertArgs', 'SslKeyCert']

@pulumi.input_type
class SslKeyCertArgs:
    def __init__(__self__, *,
                 cert_content: pulumi.Input[str],
                 cert_name: pulumi.Input[str],
                 key_content: pulumi.Input[str],
                 key_name: pulumi.Input[str],
                 cert_full_path: Optional[pulumi.Input[str]] = None,
                 cert_monitoring_type: Optional[pulumi.Input[str]] = None,
                 cert_ocsp: Optional[pulumi.Input[str]] = None,
                 issuer_cert: Optional[pulumi.Input[str]] = None,
                 key_full_path: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SslKeyCert resource.
        :param pulumi.Input[str] cert_content: The content of the cert.
        :param pulumi.Input[str] cert_name: Name of the SSL certificate to be Imported on to BIGIP.
        :param pulumi.Input[str] key_content: The content of the key.
        :param pulumi.Input[str] key_name: Name of the SSL key to be Imported on to BIGIP.
        :param pulumi.Input[str] cert_full_path: full path of the SSL certificate on the BIGIP.
        :param pulumi.Input[str] cert_monitoring_type: Specifies the type of monitoring used.
        :param pulumi.Input[str] cert_ocsp: Specifies the OCSP responder.
        :param pulumi.Input[str] issuer_cert: Specifies the issuer certificate.
        :param pulumi.Input[str] key_full_path: full path of the SSL key on the BIGIP.
        :param pulumi.Input[str] partition: Partition on to SSL certificate and key to be imported.
        :param pulumi.Input[str] passphrase: Passphrase on the SSL key.
        """
        pulumi.set(__self__, "cert_content", cert_content)
        pulumi.set(__self__, "cert_name", cert_name)
        pulumi.set(__self__, "key_content", key_content)
        pulumi.set(__self__, "key_name", key_name)
        if cert_full_path is not None:
            pulumi.set(__self__, "cert_full_path", cert_full_path)
        if cert_monitoring_type is not None:
            pulumi.set(__self__, "cert_monitoring_type", cert_monitoring_type)
        if cert_ocsp is not None:
            pulumi.set(__self__, "cert_ocsp", cert_ocsp)
        if issuer_cert is not None:
            pulumi.set(__self__, "issuer_cert", issuer_cert)
        if key_full_path is not None:
            pulumi.set(__self__, "key_full_path", key_full_path)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if passphrase is not None:
            pulumi.set(__self__, "passphrase", passphrase)

    @property
    @pulumi.getter(name="certContent")
    def cert_content(self) -> pulumi.Input[str]:
        """
        The content of the cert.
        """
        return pulumi.get(self, "cert_content")

    @cert_content.setter
    def cert_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert_content", value)

    @property
    @pulumi.getter(name="certName")
    def cert_name(self) -> pulumi.Input[str]:
        """
        Name of the SSL certificate to be Imported on to BIGIP.
        """
        return pulumi.get(self, "cert_name")

    @cert_name.setter
    def cert_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert_name", value)

    @property
    @pulumi.getter(name="keyContent")
    def key_content(self) -> pulumi.Input[str]:
        """
        The content of the key.
        """
        return pulumi.get(self, "key_content")

    @key_content.setter
    def key_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_content", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Input[str]:
        """
        Name of the SSL key to be Imported on to BIGIP.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="certFullPath")
    def cert_full_path(self) -> Optional[pulumi.Input[str]]:
        """
        full path of the SSL certificate on the BIGIP.
        """
        return pulumi.get(self, "cert_full_path")

    @cert_full_path.setter
    def cert_full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_full_path", value)

    @property
    @pulumi.getter(name="certMonitoringType")
    def cert_monitoring_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of monitoring used.
        """
        return pulumi.get(self, "cert_monitoring_type")

    @cert_monitoring_type.setter
    def cert_monitoring_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_monitoring_type", value)

    @property
    @pulumi.getter(name="certOcsp")
    def cert_ocsp(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the OCSP responder.
        """
        return pulumi.get(self, "cert_ocsp")

    @cert_ocsp.setter
    def cert_ocsp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_ocsp", value)

    @property
    @pulumi.getter(name="issuerCert")
    def issuer_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the issuer certificate.
        """
        return pulumi.get(self, "issuer_cert")

    @issuer_cert.setter
    def issuer_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer_cert", value)

    @property
    @pulumi.getter(name="keyFullPath")
    def key_full_path(self) -> Optional[pulumi.Input[str]]:
        """
        full path of the SSL key on the BIGIP.
        """
        return pulumi.get(self, "key_full_path")

    @key_full_path.setter
    def key_full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_full_path", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Partition on to SSL certificate and key to be imported.
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter
    def passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Passphrase on the SSL key.
        """
        return pulumi.get(self, "passphrase")

    @passphrase.setter
    def passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passphrase", value)


@pulumi.input_type
class _SslKeyCertState:
    def __init__(__self__, *,
                 cert_content: Optional[pulumi.Input[str]] = None,
                 cert_full_path: Optional[pulumi.Input[str]] = None,
                 cert_monitoring_type: Optional[pulumi.Input[str]] = None,
                 cert_name: Optional[pulumi.Input[str]] = None,
                 cert_ocsp: Optional[pulumi.Input[str]] = None,
                 issuer_cert: Optional[pulumi.Input[str]] = None,
                 key_content: Optional[pulumi.Input[str]] = None,
                 key_full_path: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SslKeyCert resources.
        :param pulumi.Input[str] cert_content: The content of the cert.
        :param pulumi.Input[str] cert_full_path: full path of the SSL certificate on the BIGIP.
        :param pulumi.Input[str] cert_monitoring_type: Specifies the type of monitoring used.
        :param pulumi.Input[str] cert_name: Name of the SSL certificate to be Imported on to BIGIP.
        :param pulumi.Input[str] cert_ocsp: Specifies the OCSP responder.
        :param pulumi.Input[str] issuer_cert: Specifies the issuer certificate.
        :param pulumi.Input[str] key_content: The content of the key.
        :param pulumi.Input[str] key_full_path: full path of the SSL key on the BIGIP.
        :param pulumi.Input[str] key_name: Name of the SSL key to be Imported on to BIGIP.
        :param pulumi.Input[str] partition: Partition on to SSL certificate and key to be imported.
        :param pulumi.Input[str] passphrase: Passphrase on the SSL key.
        """
        if cert_content is not None:
            pulumi.set(__self__, "cert_content", cert_content)
        if cert_full_path is not None:
            pulumi.set(__self__, "cert_full_path", cert_full_path)
        if cert_monitoring_type is not None:
            pulumi.set(__self__, "cert_monitoring_type", cert_monitoring_type)
        if cert_name is not None:
            pulumi.set(__self__, "cert_name", cert_name)
        if cert_ocsp is not None:
            pulumi.set(__self__, "cert_ocsp", cert_ocsp)
        if issuer_cert is not None:
            pulumi.set(__self__, "issuer_cert", issuer_cert)
        if key_content is not None:
            pulumi.set(__self__, "key_content", key_content)
        if key_full_path is not None:
            pulumi.set(__self__, "key_full_path", key_full_path)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if partition is not None:
            pulumi.set(__self__, "partition", partition)
        if passphrase is not None:
            pulumi.set(__self__, "passphrase", passphrase)

    @property
    @pulumi.getter(name="certContent")
    def cert_content(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the cert.
        """
        return pulumi.get(self, "cert_content")

    @cert_content.setter
    def cert_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_content", value)

    @property
    @pulumi.getter(name="certFullPath")
    def cert_full_path(self) -> Optional[pulumi.Input[str]]:
        """
        full path of the SSL certificate on the BIGIP.
        """
        return pulumi.get(self, "cert_full_path")

    @cert_full_path.setter
    def cert_full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_full_path", value)

    @property
    @pulumi.getter(name="certMonitoringType")
    def cert_monitoring_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of monitoring used.
        """
        return pulumi.get(self, "cert_monitoring_type")

    @cert_monitoring_type.setter
    def cert_monitoring_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_monitoring_type", value)

    @property
    @pulumi.getter(name="certName")
    def cert_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the SSL certificate to be Imported on to BIGIP.
        """
        return pulumi.get(self, "cert_name")

    @cert_name.setter
    def cert_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_name", value)

    @property
    @pulumi.getter(name="certOcsp")
    def cert_ocsp(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the OCSP responder.
        """
        return pulumi.get(self, "cert_ocsp")

    @cert_ocsp.setter
    def cert_ocsp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_ocsp", value)

    @property
    @pulumi.getter(name="issuerCert")
    def issuer_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the issuer certificate.
        """
        return pulumi.get(self, "issuer_cert")

    @issuer_cert.setter
    def issuer_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer_cert", value)

    @property
    @pulumi.getter(name="keyContent")
    def key_content(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the key.
        """
        return pulumi.get(self, "key_content")

    @key_content.setter
    def key_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_content", value)

    @property
    @pulumi.getter(name="keyFullPath")
    def key_full_path(self) -> Optional[pulumi.Input[str]]:
        """
        full path of the SSL key on the BIGIP.
        """
        return pulumi.get(self, "key_full_path")

    @key_full_path.setter
    def key_full_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_full_path", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the SSL key to be Imported on to BIGIP.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter
    def partition(self) -> Optional[pulumi.Input[str]]:
        """
        Partition on to SSL certificate and key to be imported.
        """
        return pulumi.get(self, "partition")

    @partition.setter
    def partition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "partition", value)

    @property
    @pulumi.getter
    def passphrase(self) -> Optional[pulumi.Input[str]]:
        """
        Passphrase on the SSL key.
        """
        return pulumi.get(self, "passphrase")

    @passphrase.setter
    def passphrase(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "passphrase", value)


class SslKeyCert(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert_content: Optional[pulumi.Input[str]] = None,
                 cert_full_path: Optional[pulumi.Input[str]] = None,
                 cert_monitoring_type: Optional[pulumi.Input[str]] = None,
                 cert_name: Optional[pulumi.Input[str]] = None,
                 cert_ocsp: Optional[pulumi.Input[str]] = None,
                 issuer_cert: Optional[pulumi.Input[str]] = None,
                 key_content: Optional[pulumi.Input[str]] = None,
                 key_full_path: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `SslKeyCert` This resource will import SSL certificate and key on BIG-IP LTM.
        The certificate and the key can be imported from files on the local disk, in PEM format

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_content: The content of the cert.
        :param pulumi.Input[str] cert_full_path: full path of the SSL certificate on the BIGIP.
        :param pulumi.Input[str] cert_monitoring_type: Specifies the type of monitoring used.
        :param pulumi.Input[str] cert_name: Name of the SSL certificate to be Imported on to BIGIP.
        :param pulumi.Input[str] cert_ocsp: Specifies the OCSP responder.
        :param pulumi.Input[str] issuer_cert: Specifies the issuer certificate.
        :param pulumi.Input[str] key_content: The content of the key.
        :param pulumi.Input[str] key_full_path: full path of the SSL key on the BIGIP.
        :param pulumi.Input[str] key_name: Name of the SSL key to be Imported on to BIGIP.
        :param pulumi.Input[str] partition: Partition on to SSL certificate and key to be imported.
        :param pulumi.Input[str] passphrase: Passphrase on the SSL key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SslKeyCertArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `SslKeyCert` This resource will import SSL certificate and key on BIG-IP LTM.
        The certificate and the key can be imported from files on the local disk, in PEM format

        :param str resource_name: The name of the resource.
        :param SslKeyCertArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SslKeyCertArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cert_content: Optional[pulumi.Input[str]] = None,
                 cert_full_path: Optional[pulumi.Input[str]] = None,
                 cert_monitoring_type: Optional[pulumi.Input[str]] = None,
                 cert_name: Optional[pulumi.Input[str]] = None,
                 cert_ocsp: Optional[pulumi.Input[str]] = None,
                 issuer_cert: Optional[pulumi.Input[str]] = None,
                 key_content: Optional[pulumi.Input[str]] = None,
                 key_full_path: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 partition: Optional[pulumi.Input[str]] = None,
                 passphrase: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SslKeyCertArgs.__new__(SslKeyCertArgs)

            if cert_content is None and not opts.urn:
                raise TypeError("Missing required property 'cert_content'")
            __props__.__dict__["cert_content"] = None if cert_content is None else pulumi.Output.secret(cert_content)
            __props__.__dict__["cert_full_path"] = cert_full_path
            __props__.__dict__["cert_monitoring_type"] = cert_monitoring_type
            if cert_name is None and not opts.urn:
                raise TypeError("Missing required property 'cert_name'")
            __props__.__dict__["cert_name"] = cert_name
            __props__.__dict__["cert_ocsp"] = cert_ocsp
            __props__.__dict__["issuer_cert"] = issuer_cert
            if key_content is None and not opts.urn:
                raise TypeError("Missing required property 'key_content'")
            __props__.__dict__["key_content"] = None if key_content is None else pulumi.Output.secret(key_content)
            __props__.__dict__["key_full_path"] = key_full_path
            if key_name is None and not opts.urn:
                raise TypeError("Missing required property 'key_name'")
            __props__.__dict__["key_name"] = key_name
            __props__.__dict__["partition"] = partition
            __props__.__dict__["passphrase"] = None if passphrase is None else pulumi.Output.secret(passphrase)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["certContent", "keyContent", "passphrase"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SslKeyCert, __self__).__init__(
            'f5bigip:index/sslKeyCert:SslKeyCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cert_content: Optional[pulumi.Input[str]] = None,
            cert_full_path: Optional[pulumi.Input[str]] = None,
            cert_monitoring_type: Optional[pulumi.Input[str]] = None,
            cert_name: Optional[pulumi.Input[str]] = None,
            cert_ocsp: Optional[pulumi.Input[str]] = None,
            issuer_cert: Optional[pulumi.Input[str]] = None,
            key_content: Optional[pulumi.Input[str]] = None,
            key_full_path: Optional[pulumi.Input[str]] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            partition: Optional[pulumi.Input[str]] = None,
            passphrase: Optional[pulumi.Input[str]] = None) -> 'SslKeyCert':
        """
        Get an existing SslKeyCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert_content: The content of the cert.
        :param pulumi.Input[str] cert_full_path: full path of the SSL certificate on the BIGIP.
        :param pulumi.Input[str] cert_monitoring_type: Specifies the type of monitoring used.
        :param pulumi.Input[str] cert_name: Name of the SSL certificate to be Imported on to BIGIP.
        :param pulumi.Input[str] cert_ocsp: Specifies the OCSP responder.
        :param pulumi.Input[str] issuer_cert: Specifies the issuer certificate.
        :param pulumi.Input[str] key_content: The content of the key.
        :param pulumi.Input[str] key_full_path: full path of the SSL key on the BIGIP.
        :param pulumi.Input[str] key_name: Name of the SSL key to be Imported on to BIGIP.
        :param pulumi.Input[str] partition: Partition on to SSL certificate and key to be imported.
        :param pulumi.Input[str] passphrase: Passphrase on the SSL key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SslKeyCertState.__new__(_SslKeyCertState)

        __props__.__dict__["cert_content"] = cert_content
        __props__.__dict__["cert_full_path"] = cert_full_path
        __props__.__dict__["cert_monitoring_type"] = cert_monitoring_type
        __props__.__dict__["cert_name"] = cert_name
        __props__.__dict__["cert_ocsp"] = cert_ocsp
        __props__.__dict__["issuer_cert"] = issuer_cert
        __props__.__dict__["key_content"] = key_content
        __props__.__dict__["key_full_path"] = key_full_path
        __props__.__dict__["key_name"] = key_name
        __props__.__dict__["partition"] = partition
        __props__.__dict__["passphrase"] = passphrase
        return SslKeyCert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certContent")
    def cert_content(self) -> pulumi.Output[str]:
        """
        The content of the cert.
        """
        return pulumi.get(self, "cert_content")

    @property
    @pulumi.getter(name="certFullPath")
    def cert_full_path(self) -> pulumi.Output[str]:
        """
        full path of the SSL certificate on the BIGIP.
        """
        return pulumi.get(self, "cert_full_path")

    @property
    @pulumi.getter(name="certMonitoringType")
    def cert_monitoring_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the type of monitoring used.
        """
        return pulumi.get(self, "cert_monitoring_type")

    @property
    @pulumi.getter(name="certName")
    def cert_name(self) -> pulumi.Output[str]:
        """
        Name of the SSL certificate to be Imported on to BIGIP.
        """
        return pulumi.get(self, "cert_name")

    @property
    @pulumi.getter(name="certOcsp")
    def cert_ocsp(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the OCSP responder.
        """
        return pulumi.get(self, "cert_ocsp")

    @property
    @pulumi.getter(name="issuerCert")
    def issuer_cert(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the issuer certificate.
        """
        return pulumi.get(self, "issuer_cert")

    @property
    @pulumi.getter(name="keyContent")
    def key_content(self) -> pulumi.Output[str]:
        """
        The content of the key.
        """
        return pulumi.get(self, "key_content")

    @property
    @pulumi.getter(name="keyFullPath")
    def key_full_path(self) -> pulumi.Output[str]:
        """
        full path of the SSL key on the BIGIP.
        """
        return pulumi.get(self, "key_full_path")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[str]:
        """
        Name of the SSL key to be Imported on to BIGIP.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter
    def partition(self) -> pulumi.Output[Optional[str]]:
        """
        Partition on to SSL certificate and key to be imported.
        """
        return pulumi.get(self, "partition")

    @property
    @pulumi.getter
    def passphrase(self) -> pulumi.Output[Optional[str]]:
        """
        Passphrase on the SSL key.
        """
        return pulumi.get(self, "passphrase")

