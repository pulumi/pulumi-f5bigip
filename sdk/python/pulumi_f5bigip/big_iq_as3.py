# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BigIqAs3Args', 'BigIqAs3']

@pulumi.input_type
class BigIqAs3Args:
    def __init__(__self__, *,
                 as3_json: pulumi.Input[str],
                 bigiq_address: pulumi.Input[str],
                 bigiq_password: pulumi.Input[str],
                 bigiq_user: pulumi.Input[str],
                 bigiq_login_ref: Optional[pulumi.Input[str]] = None,
                 bigiq_port: Optional[pulumi.Input[str]] = None,
                 bigiq_token_auth: Optional[pulumi.Input[bool]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BigIqAs3 resource.
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[str] bigiq_address: Address of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_password: Password of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_user: User name  of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_login_ref: BIGIQ Login reference for token authentication
        :param pulumi.Input[str] bigiq_port: type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        :param pulumi.Input[bool] bigiq_token_auth: type `bool`, if set to `true` enables Token based Authentication,default is `false`
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to `true`
               
               * `bigiq_example.json` - Example  AS3 Declarative JSON file
               
               ```json
               {
               "class": "AS3",
               "action": "deploy",
               "persist": true,
               "declaration": {
               "class": "ADC",
               "schemaVersion": "3.7.0",
               "id": "example-declaration-01",
               "label": "Task1",
               "remark": "Task 1 - HTTP Application Service",
               "target": {
               "address": "xx.xxx.xx.xxx"
               },
               "Task1": {
               "class": "Tenant",
               "MyWebApp1http": {
               "class": "Application",
               "template": "http",
               
               
               "serviceMain": {
               "class": "Service_HTTP",
               "virtualAddresses": [
               "10.1.2.10"
               ],
               "pool": "web_pool"
               },
               "web_pool": {
               "class": "Pool",
               "monitors": [
               "http"
               ],
               "members": [
               {
               "servicePort": 80,
               "serverAddresses": [
               "192.0.2.33",
               "192.0.2.13"
               ],
               "shareNodes": true
               }
               ]
               }
               }
               }
               }
               }
               ```
               
               * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
               
               >  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        :param pulumi.Input[str] tenant_list: Name of Tenant
        """
        pulumi.set(__self__, "as3_json", as3_json)
        pulumi.set(__self__, "bigiq_address", bigiq_address)
        pulumi.set(__self__, "bigiq_password", bigiq_password)
        pulumi.set(__self__, "bigiq_user", bigiq_user)
        if bigiq_login_ref is not None:
            pulumi.set(__self__, "bigiq_login_ref", bigiq_login_ref)
        if bigiq_port is not None:
            pulumi.set(__self__, "bigiq_port", bigiq_port)
        if bigiq_token_auth is not None:
            pulumi.set(__self__, "bigiq_token_auth", bigiq_token_auth)
        if ignore_metadata is not None:
            pulumi.set(__self__, "ignore_metadata", ignore_metadata)
        if tenant_list is not None:
            pulumi.set(__self__, "tenant_list", tenant_list)

    @property
    @pulumi.getter(name="as3Json")
    def as3_json(self) -> pulumi.Input[str]:
        """
        Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        """
        return pulumi.get(self, "as3_json")

    @as3_json.setter
    def as3_json(self, value: pulumi.Input[str]):
        pulumi.set(self, "as3_json", value)

    @property
    @pulumi.getter(name="bigiqAddress")
    def bigiq_address(self) -> pulumi.Input[str]:
        """
        Address of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_address")

    @bigiq_address.setter
    def bigiq_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "bigiq_address", value)

    @property
    @pulumi.getter(name="bigiqPassword")
    def bigiq_password(self) -> pulumi.Input[str]:
        """
        Password of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_password")

    @bigiq_password.setter
    def bigiq_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "bigiq_password", value)

    @property
    @pulumi.getter(name="bigiqUser")
    def bigiq_user(self) -> pulumi.Input[str]:
        """
        User name  of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_user")

    @bigiq_user.setter
    def bigiq_user(self, value: pulumi.Input[str]):
        pulumi.set(self, "bigiq_user", value)

    @property
    @pulumi.getter(name="bigiqLoginRef")
    def bigiq_login_ref(self) -> Optional[pulumi.Input[str]]:
        """
        BIGIQ Login reference for token authentication
        """
        return pulumi.get(self, "bigiq_login_ref")

    @bigiq_login_ref.setter
    def bigiq_login_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bigiq_login_ref", value)

    @property
    @pulumi.getter(name="bigiqPort")
    def bigiq_port(self) -> Optional[pulumi.Input[str]]:
        """
        type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        """
        return pulumi.get(self, "bigiq_port")

    @bigiq_port.setter
    def bigiq_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bigiq_port", value)

    @property
    @pulumi.getter(name="bigiqTokenAuth")
    def bigiq_token_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        type `bool`, if set to `true` enables Token based Authentication,default is `false`
        """
        return pulumi.get(self, "bigiq_token_auth")

    @bigiq_token_auth.setter
    def bigiq_token_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bigiq_token_auth", value)

    @property
    @pulumi.getter(name="ignoreMetadata")
    def ignore_metadata(self) -> Optional[pulumi.Input[bool]]:
        """
        Set True if you want to ignore metadata changes during update. By default it is set to `true`

        * `bigiq_example.json` - Example  AS3 Declarative JSON file

        ```json
        {
        "class": "AS3",
        "action": "deploy",
        "persist": true,
        "declaration": {
        "class": "ADC",
        "schemaVersion": "3.7.0",
        "id": "example-declaration-01",
        "label": "Task1",
        "remark": "Task 1 - HTTP Application Service",
        "target": {
        "address": "xx.xxx.xx.xxx"
        },
        "Task1": {
        "class": "Tenant",
        "MyWebApp1http": {
        "class": "Application",
        "template": "http",


        "serviceMain": {
        "class": "Service_HTTP",
        "virtualAddresses": [
        "10.1.2.10"
        ],
        "pool": "web_pool"
        },
        "web_pool": {
        "class": "Pool",
        "monitors": [
        "http"
        ],
        "members": [
        {
        "servicePort": 80,
        "serverAddresses": [
        "192.0.2.33",
        "192.0.2.13"
        ],
        "shareNodes": true
        }
        ]
        }
        }
        }
        }
        }
        ```

        * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html

        >  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        """
        return pulumi.get(self, "ignore_metadata")

    @ignore_metadata.setter
    def ignore_metadata(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_metadata", value)

    @property
    @pulumi.getter(name="tenantList")
    def tenant_list(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_list")

    @tenant_list.setter
    def tenant_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_list", value)


@pulumi.input_type
class _BigIqAs3State:
    def __init__(__self__, *,
                 as3_json: Optional[pulumi.Input[str]] = None,
                 bigiq_address: Optional[pulumi.Input[str]] = None,
                 bigiq_login_ref: Optional[pulumi.Input[str]] = None,
                 bigiq_password: Optional[pulumi.Input[str]] = None,
                 bigiq_port: Optional[pulumi.Input[str]] = None,
                 bigiq_token_auth: Optional[pulumi.Input[bool]] = None,
                 bigiq_user: Optional[pulumi.Input[str]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BigIqAs3 resources.
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[str] bigiq_address: Address of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_login_ref: BIGIQ Login reference for token authentication
        :param pulumi.Input[str] bigiq_password: Password of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_port: type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        :param pulumi.Input[bool] bigiq_token_auth: type `bool`, if set to `true` enables Token based Authentication,default is `false`
        :param pulumi.Input[str] bigiq_user: User name  of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to `true`
               
               * `bigiq_example.json` - Example  AS3 Declarative JSON file
               
               ```json
               {
               "class": "AS3",
               "action": "deploy",
               "persist": true,
               "declaration": {
               "class": "ADC",
               "schemaVersion": "3.7.0",
               "id": "example-declaration-01",
               "label": "Task1",
               "remark": "Task 1 - HTTP Application Service",
               "target": {
               "address": "xx.xxx.xx.xxx"
               },
               "Task1": {
               "class": "Tenant",
               "MyWebApp1http": {
               "class": "Application",
               "template": "http",
               
               
               "serviceMain": {
               "class": "Service_HTTP",
               "virtualAddresses": [
               "10.1.2.10"
               ],
               "pool": "web_pool"
               },
               "web_pool": {
               "class": "Pool",
               "monitors": [
               "http"
               ],
               "members": [
               {
               "servicePort": 80,
               "serverAddresses": [
               "192.0.2.33",
               "192.0.2.13"
               ],
               "shareNodes": true
               }
               ]
               }
               }
               }
               }
               }
               ```
               
               * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
               
               >  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        :param pulumi.Input[str] tenant_list: Name of Tenant
        """
        if as3_json is not None:
            pulumi.set(__self__, "as3_json", as3_json)
        if bigiq_address is not None:
            pulumi.set(__self__, "bigiq_address", bigiq_address)
        if bigiq_login_ref is not None:
            pulumi.set(__self__, "bigiq_login_ref", bigiq_login_ref)
        if bigiq_password is not None:
            pulumi.set(__self__, "bigiq_password", bigiq_password)
        if bigiq_port is not None:
            pulumi.set(__self__, "bigiq_port", bigiq_port)
        if bigiq_token_auth is not None:
            pulumi.set(__self__, "bigiq_token_auth", bigiq_token_auth)
        if bigiq_user is not None:
            pulumi.set(__self__, "bigiq_user", bigiq_user)
        if ignore_metadata is not None:
            pulumi.set(__self__, "ignore_metadata", ignore_metadata)
        if tenant_list is not None:
            pulumi.set(__self__, "tenant_list", tenant_list)

    @property
    @pulumi.getter(name="as3Json")
    def as3_json(self) -> Optional[pulumi.Input[str]]:
        """
        Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        """
        return pulumi.get(self, "as3_json")

    @as3_json.setter
    def as3_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "as3_json", value)

    @property
    @pulumi.getter(name="bigiqAddress")
    def bigiq_address(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_address")

    @bigiq_address.setter
    def bigiq_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bigiq_address", value)

    @property
    @pulumi.getter(name="bigiqLoginRef")
    def bigiq_login_ref(self) -> Optional[pulumi.Input[str]]:
        """
        BIGIQ Login reference for token authentication
        """
        return pulumi.get(self, "bigiq_login_ref")

    @bigiq_login_ref.setter
    def bigiq_login_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bigiq_login_ref", value)

    @property
    @pulumi.getter(name="bigiqPassword")
    def bigiq_password(self) -> Optional[pulumi.Input[str]]:
        """
        Password of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_password")

    @bigiq_password.setter
    def bigiq_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bigiq_password", value)

    @property
    @pulumi.getter(name="bigiqPort")
    def bigiq_port(self) -> Optional[pulumi.Input[str]]:
        """
        type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        """
        return pulumi.get(self, "bigiq_port")

    @bigiq_port.setter
    def bigiq_port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bigiq_port", value)

    @property
    @pulumi.getter(name="bigiqTokenAuth")
    def bigiq_token_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        type `bool`, if set to `true` enables Token based Authentication,default is `false`
        """
        return pulumi.get(self, "bigiq_token_auth")

    @bigiq_token_auth.setter
    def bigiq_token_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bigiq_token_auth", value)

    @property
    @pulumi.getter(name="bigiqUser")
    def bigiq_user(self) -> Optional[pulumi.Input[str]]:
        """
        User name  of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_user")

    @bigiq_user.setter
    def bigiq_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bigiq_user", value)

    @property
    @pulumi.getter(name="ignoreMetadata")
    def ignore_metadata(self) -> Optional[pulumi.Input[bool]]:
        """
        Set True if you want to ignore metadata changes during update. By default it is set to `true`

        * `bigiq_example.json` - Example  AS3 Declarative JSON file

        ```json
        {
        "class": "AS3",
        "action": "deploy",
        "persist": true,
        "declaration": {
        "class": "ADC",
        "schemaVersion": "3.7.0",
        "id": "example-declaration-01",
        "label": "Task1",
        "remark": "Task 1 - HTTP Application Service",
        "target": {
        "address": "xx.xxx.xx.xxx"
        },
        "Task1": {
        "class": "Tenant",
        "MyWebApp1http": {
        "class": "Application",
        "template": "http",


        "serviceMain": {
        "class": "Service_HTTP",
        "virtualAddresses": [
        "10.1.2.10"
        ],
        "pool": "web_pool"
        },
        "web_pool": {
        "class": "Pool",
        "monitors": [
        "http"
        ],
        "members": [
        {
        "servicePort": 80,
        "serverAddresses": [
        "192.0.2.33",
        "192.0.2.13"
        ],
        "shareNodes": true
        }
        ]
        }
        }
        }
        }
        }
        ```

        * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html

        >  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        """
        return pulumi.get(self, "ignore_metadata")

    @ignore_metadata.setter
    def ignore_metadata(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_metadata", value)

    @property
    @pulumi.getter(name="tenantList")
    def tenant_list(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_list")

    @tenant_list.setter
    def tenant_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_list", value)


class BigIqAs3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 as3_json: Optional[pulumi.Input[str]] = None,
                 bigiq_address: Optional[pulumi.Input[str]] = None,
                 bigiq_login_ref: Optional[pulumi.Input[str]] = None,
                 bigiq_password: Optional[pulumi.Input[str]] = None,
                 bigiq_port: Optional[pulumi.Input[str]] = None,
                 bigiq_token_auth: Optional[pulumi.Input[bool]] = None,
                 bigiq_user: Optional[pulumi.Input[str]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        `BigIqAs3` provides details about bigiq as3 resource

        This resource is helpful to configure as3 declarative JSON on BIG-IP through BIG-IQ.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip
        import pulumi_std as std

        # Example Usage for json file
        exampletask = f5bigip.BigIqAs3("exampletask",
            bigiq_address="xx.xx.xxx.xx",
            bigiq_user="xxxxx",
            bigiq_password="xxxxxxxxx",
            as3_json=std.file(input="bigiq_example.json").result)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[str] bigiq_address: Address of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_login_ref: BIGIQ Login reference for token authentication
        :param pulumi.Input[str] bigiq_password: Password of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_port: type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        :param pulumi.Input[bool] bigiq_token_auth: type `bool`, if set to `true` enables Token based Authentication,default is `false`
        :param pulumi.Input[str] bigiq_user: User name  of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to `true`
               
               * `bigiq_example.json` - Example  AS3 Declarative JSON file
               
               ```json
               {
               "class": "AS3",
               "action": "deploy",
               "persist": true,
               "declaration": {
               "class": "ADC",
               "schemaVersion": "3.7.0",
               "id": "example-declaration-01",
               "label": "Task1",
               "remark": "Task 1 - HTTP Application Service",
               "target": {
               "address": "xx.xxx.xx.xxx"
               },
               "Task1": {
               "class": "Tenant",
               "MyWebApp1http": {
               "class": "Application",
               "template": "http",
               
               
               "serviceMain": {
               "class": "Service_HTTP",
               "virtualAddresses": [
               "10.1.2.10"
               ],
               "pool": "web_pool"
               },
               "web_pool": {
               "class": "Pool",
               "monitors": [
               "http"
               ],
               "members": [
               {
               "servicePort": 80,
               "serverAddresses": [
               "192.0.2.33",
               "192.0.2.13"
               ],
               "shareNodes": true
               }
               ]
               }
               }
               }
               }
               }
               ```
               
               * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
               
               >  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        :param pulumi.Input[str] tenant_list: Name of Tenant
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BigIqAs3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `BigIqAs3` provides details about bigiq as3 resource

        This resource is helpful to configure as3 declarative JSON on BIG-IP through BIG-IQ.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_f5bigip as f5bigip
        import pulumi_std as std

        # Example Usage for json file
        exampletask = f5bigip.BigIqAs3("exampletask",
            bigiq_address="xx.xx.xxx.xx",
            bigiq_user="xxxxx",
            bigiq_password="xxxxxxxxx",
            as3_json=std.file(input="bigiq_example.json").result)
        ```

        :param str resource_name: The name of the resource.
        :param BigIqAs3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BigIqAs3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 as3_json: Optional[pulumi.Input[str]] = None,
                 bigiq_address: Optional[pulumi.Input[str]] = None,
                 bigiq_login_ref: Optional[pulumi.Input[str]] = None,
                 bigiq_password: Optional[pulumi.Input[str]] = None,
                 bigiq_port: Optional[pulumi.Input[str]] = None,
                 bigiq_token_auth: Optional[pulumi.Input[bool]] = None,
                 bigiq_user: Optional[pulumi.Input[str]] = None,
                 ignore_metadata: Optional[pulumi.Input[bool]] = None,
                 tenant_list: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BigIqAs3Args.__new__(BigIqAs3Args)

            if as3_json is None and not opts.urn:
                raise TypeError("Missing required property 'as3_json'")
            __props__.__dict__["as3_json"] = as3_json
            if bigiq_address is None and not opts.urn:
                raise TypeError("Missing required property 'bigiq_address'")
            __props__.__dict__["bigiq_address"] = bigiq_address
            __props__.__dict__["bigiq_login_ref"] = None if bigiq_login_ref is None else pulumi.Output.secret(bigiq_login_ref)
            if bigiq_password is None and not opts.urn:
                raise TypeError("Missing required property 'bigiq_password'")
            __props__.__dict__["bigiq_password"] = None if bigiq_password is None else pulumi.Output.secret(bigiq_password)
            __props__.__dict__["bigiq_port"] = None if bigiq_port is None else pulumi.Output.secret(bigiq_port)
            __props__.__dict__["bigiq_token_auth"] = None if bigiq_token_auth is None else pulumi.Output.secret(bigiq_token_auth)
            if bigiq_user is None and not opts.urn:
                raise TypeError("Missing required property 'bigiq_user'")
            __props__.__dict__["bigiq_user"] = None if bigiq_user is None else pulumi.Output.secret(bigiq_user)
            __props__.__dict__["ignore_metadata"] = ignore_metadata
            __props__.__dict__["tenant_list"] = tenant_list
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["bigiqLoginRef", "bigiqPassword", "bigiqPort", "bigiqTokenAuth", "bigiqUser"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(BigIqAs3, __self__).__init__(
            'f5bigip:index/bigIqAs3:BigIqAs3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            as3_json: Optional[pulumi.Input[str]] = None,
            bigiq_address: Optional[pulumi.Input[str]] = None,
            bigiq_login_ref: Optional[pulumi.Input[str]] = None,
            bigiq_password: Optional[pulumi.Input[str]] = None,
            bigiq_port: Optional[pulumi.Input[str]] = None,
            bigiq_token_auth: Optional[pulumi.Input[bool]] = None,
            bigiq_user: Optional[pulumi.Input[str]] = None,
            ignore_metadata: Optional[pulumi.Input[bool]] = None,
            tenant_list: Optional[pulumi.Input[str]] = None) -> 'BigIqAs3':
        """
        Get an existing BigIqAs3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] as3_json: Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        :param pulumi.Input[str] bigiq_address: Address of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_login_ref: BIGIQ Login reference for token authentication
        :param pulumi.Input[str] bigiq_password: Password of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[str] bigiq_port: type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        :param pulumi.Input[bool] bigiq_token_auth: type `bool`, if set to `true` enables Token based Authentication,default is `false`
        :param pulumi.Input[str] bigiq_user: User name  of the BIG-IQ to which your targer BIG-IP is attached
        :param pulumi.Input[bool] ignore_metadata: Set True if you want to ignore metadata changes during update. By default it is set to `true`
               
               * `bigiq_example.json` - Example  AS3 Declarative JSON file
               
               ```json
               {
               "class": "AS3",
               "action": "deploy",
               "persist": true,
               "declaration": {
               "class": "ADC",
               "schemaVersion": "3.7.0",
               "id": "example-declaration-01",
               "label": "Task1",
               "remark": "Task 1 - HTTP Application Service",
               "target": {
               "address": "xx.xxx.xx.xxx"
               },
               "Task1": {
               "class": "Tenant",
               "MyWebApp1http": {
               "class": "Application",
               "template": "http",
               
               
               "serviceMain": {
               "class": "Service_HTTP",
               "virtualAddresses": [
               "10.1.2.10"
               ],
               "pool": "web_pool"
               },
               "web_pool": {
               "class": "Pool",
               "monitors": [
               "http"
               ],
               "members": [
               {
               "servicePort": 80,
               "serverAddresses": [
               "192.0.2.33",
               "192.0.2.13"
               ],
               "shareNodes": true
               }
               ]
               }
               }
               }
               }
               }
               ```
               
               * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html
               
               >  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        :param pulumi.Input[str] tenant_list: Name of Tenant
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BigIqAs3State.__new__(_BigIqAs3State)

        __props__.__dict__["as3_json"] = as3_json
        __props__.__dict__["bigiq_address"] = bigiq_address
        __props__.__dict__["bigiq_login_ref"] = bigiq_login_ref
        __props__.__dict__["bigiq_password"] = bigiq_password
        __props__.__dict__["bigiq_port"] = bigiq_port
        __props__.__dict__["bigiq_token_auth"] = bigiq_token_auth
        __props__.__dict__["bigiq_user"] = bigiq_user
        __props__.__dict__["ignore_metadata"] = ignore_metadata
        __props__.__dict__["tenant_list"] = tenant_list
        return BigIqAs3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="as3Json")
    def as3_json(self) -> pulumi.Output[str]:
        """
        Path/Filename of Declarative AS3 JSON which is a json file used with builtin ```file``` function
        """
        return pulumi.get(self, "as3_json")

    @property
    @pulumi.getter(name="bigiqAddress")
    def bigiq_address(self) -> pulumi.Output[str]:
        """
        Address of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_address")

    @property
    @pulumi.getter(name="bigiqLoginRef")
    def bigiq_login_ref(self) -> pulumi.Output[Optional[str]]:
        """
        BIGIQ Login reference for token authentication
        """
        return pulumi.get(self, "bigiq_login_ref")

    @property
    @pulumi.getter(name="bigiqPassword")
    def bigiq_password(self) -> pulumi.Output[str]:
        """
        Password of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_password")

    @property
    @pulumi.getter(name="bigiqPort")
    def bigiq_port(self) -> pulumi.Output[Optional[str]]:
        """
        type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
        """
        return pulumi.get(self, "bigiq_port")

    @property
    @pulumi.getter(name="bigiqTokenAuth")
    def bigiq_token_auth(self) -> pulumi.Output[Optional[bool]]:
        """
        type `bool`, if set to `true` enables Token based Authentication,default is `false`
        """
        return pulumi.get(self, "bigiq_token_auth")

    @property
    @pulumi.getter(name="bigiqUser")
    def bigiq_user(self) -> pulumi.Output[str]:
        """
        User name  of the BIG-IQ to which your targer BIG-IP is attached
        """
        return pulumi.get(self, "bigiq_user")

    @property
    @pulumi.getter(name="ignoreMetadata")
    def ignore_metadata(self) -> pulumi.Output[Optional[bool]]:
        """
        Set True if you want to ignore metadata changes during update. By default it is set to `true`

        * `bigiq_example.json` - Example  AS3 Declarative JSON file

        ```json
        {
        "class": "AS3",
        "action": "deploy",
        "persist": true,
        "declaration": {
        "class": "ADC",
        "schemaVersion": "3.7.0",
        "id": "example-declaration-01",
        "label": "Task1",
        "remark": "Task 1 - HTTP Application Service",
        "target": {
        "address": "xx.xxx.xx.xxx"
        },
        "Task1": {
        "class": "Tenant",
        "MyWebApp1http": {
        "class": "Application",
        "template": "http",


        "serviceMain": {
        "class": "Service_HTTP",
        "virtualAddresses": [
        "10.1.2.10"
        ],
        "pool": "web_pool"
        },
        "web_pool": {
        "class": "Pool",
        "monitors": [
        "http"
        ],
        "members": [
        {
        "servicePort": 80,
        "serverAddresses": [
        "192.0.2.33",
        "192.0.2.13"
        ],
        "shareNodes": true
        }
        ]
        }
        }
        }
        }
        }
        ```

        * `AS3 documentation` - https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/big-iq.html

        >  **Note:** This resource does not support `teanat_filter` parameter as BIG-IP As3 resource
        """
        return pulumi.get(self, "ignore_metadata")

    @property
    @pulumi.getter(name="tenantList")
    def tenant_list(self) -> pulumi.Output[str]:
        """
        Name of Tenant
        """
        return pulumi.get(self, "tenant_list")

