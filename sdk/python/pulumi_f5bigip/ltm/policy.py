# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Policy(pulumi.CustomResource):
    controls: pulumi.Output[list]
    """
    Specifies the controls
    """
    name: pulumi.Output[str]
    """
    Name of the Policy
    """
    published_copy: pulumi.Output[str]
    """
    If you want to publish the policy else it will be deployed in Drafts mode.
    """
    requires: pulumi.Output[list]
    """
    Specifies the protocol
    """
    rules: pulumi.Output[list]
    """
    Rules can be applied using the policy

      * `actions` (`list`)
        * `app_service` (`str`)
        * `application` (`str`)
        * `asm` (`bool`)
        * `avr` (`bool`)
        * `cache` (`bool`)
        * `carp` (`bool`)
        * `category` (`str`)
        * `classify` (`bool`)
        * `clonePool` (`str`)
        * `code` (`float`)
        * `compress` (`bool`)
        * `content` (`str`)
        * `cookieHash` (`bool`)
        * `cookieInsert` (`bool`)
        * `cookiePassive` (`bool`)
        * `cookieRewrite` (`bool`)
        * `decompress` (`bool`)
        * `defer` (`bool`)
        * `destinationAddress` (`bool`)
        * `disable` (`bool`)
        * `domain` (`str`)
        * `enable` (`bool`)
        * `expiry` (`str`)
        * `expirySecs` (`float`)
        * `expression` (`str`)
        * `extension` (`str`)
        * `facility` (`str`)
        * `forward` (`bool`) - This action will affect forwarding.
        * `fromProfile` (`str`)
        * `hash` (`bool`)
        * `host` (`str`)
        * `http` (`bool`)
        * `httpBasicAuth` (`bool`)
        * `httpCookie` (`bool`)
        * `httpHeader` (`bool`)
        * `httpHost` (`bool`)
        * `httpReferer` (`bool`)
        * `httpReply` (`bool`)
        * `httpSetCookie` (`bool`)
        * `httpUri` (`bool`)
        * `ifile` (`str`)
        * `insert` (`bool`)
        * `internalVirtual` (`str`)
        * `ipAddress` (`str`)
        * `key` (`str`)
        * `l7dos` (`bool`)
        * `length` (`float`)
        * `location` (`str`)
        * `log` (`bool`)
        * `ltmPolicy` (`bool`)
        * `member` (`str`)
        * `message` (`str`)
        * `netmask` (`str`)
        * `nexthop` (`str`)
        * `node` (`str`)
        * `offset` (`float`)
        * `path` (`str`)
        * `pem` (`bool`)
        * `persist` (`bool`)
        * `pin` (`bool`)
        * `policy` (`str`)
        * `pool` (`str`) - This action will direct the stream to this pool.
        * `port` (`float`)
        * `priority` (`str`)
        * `profile` (`str`)
        * `protocol` (`str`)
        * `queryString` (`str`)
        * `rateclass` (`str`)
        * `redirect` (`bool`)
        * `remove` (`bool`)
        * `replace` (`bool`)
        * `request` (`bool`)
        * `requestAdapt` (`bool`)
        * `reset` (`bool`)
        * `response` (`bool`)
        * `responseAdapt` (`bool`)
        * `scheme` (`str`)
        * `script` (`str`)
        * `select` (`bool`)
        * `serverSsl` (`bool`)
        * `setVariable` (`bool`)
        * `snat` (`str`)
        * `snatpool` (`str`)
        * `sourceAddress` (`bool`)
        * `sslClientHello` (`bool`)
        * `sslServerHandshake` (`bool`)
        * `sslServerHello` (`bool`)
        * `sslSessionId` (`bool`)
        * `status` (`float`)
        * `tcl` (`bool`)
        * `tcpNagle` (`bool`)
        * `text` (`str`)
        * `timeout` (`float`)
        * `tmName` (`str`) - If Rule is used then you need to provide the tm_name it can be any value
        * `uie` (`bool`)
        * `universal` (`bool`)
        * `value` (`str`)
        * `virtual` (`str`)
        * `vlan` (`str`)
        * `vlanId` (`float`)
        * `wam` (`bool`)
        * `write` (`bool`)

      * `conditions` (`list`)
        * `address` (`bool`)
        * `all` (`bool`)
        * `app_service` (`str`)
        * `browserType` (`bool`)
        * `browserVersion` (`bool`)
        * `caseInsensitive` (`bool`)
        * `caseSensitive` (`bool`)
        * `cipher` (`bool`)
        * `cipherBits` (`bool`)
        * `clientSsl` (`bool`)
        * `code` (`bool`)
        * `commonName` (`bool`)
        * `contains` (`bool`)
        * `continent` (`bool`)
        * `countryCode` (`bool`)
        * `countryName` (`bool`)
        * `cpuUsage` (`bool`)
        * `deviceMake` (`bool`)
        * `deviceModel` (`bool`)
        * `domain` (`bool`)
        * `endsWith` (`bool`)
        * `equals` (`bool`)
        * `expiry` (`bool`)
        * `extension` (`bool`)
        * `external` (`bool`)
        * `geoip` (`bool`)
        * `greater` (`bool`)
        * `greaterOrEqual` (`bool`)
        * `host` (`bool`)
        * `httpBasicAuth` (`bool`)
        * `httpCookie` (`bool`)
        * `httpHeader` (`bool`)
        * `httpHost` (`bool`)
        * `httpMethod` (`bool`)
        * `httpReferer` (`bool`)
        * `httpSetCookie` (`bool`)
        * `httpStatus` (`bool`)
        * `httpUri` (`bool`)
        * `httpUserAgent` (`bool`)
        * `httpVersion` (`bool`)
        * `index` (`float`)
        * `internal` (`bool`)
        * `isp` (`bool`)
        * `last15secs` (`bool`)
        * `last1min` (`bool`)
        * `last5mins` (`bool`)
        * `less` (`bool`)
        * `lessOrEqual` (`bool`)
        * `local` (`bool`)
        * `major` (`bool`)
        * `matches` (`bool`)
        * `minor` (`bool`)
        * `missing` (`bool`)
        * `mss` (`bool`)
        * `not` (`bool`)
        * `org` (`bool`)
        * `password` (`bool`)
        * `path` (`bool`)
        * `pathSegment` (`bool`)
        * `port` (`bool`)
        * `present` (`bool`)
        * `protocol` (`bool`)
        * `queryParameter` (`bool`)
        * `queryString` (`bool`)
        * `regionCode` (`bool`)
        * `regionName` (`bool`)
        * `remote` (`bool`)
        * `request` (`bool`)
        * `response` (`bool`)
        * `routeDomain` (`bool`)
        * `rtt` (`bool`)
        * `scheme` (`bool`)
        * `server_name` (`bool`)
        * `sslCert` (`bool`)
        * `sslClientHello` (`bool`)
        * `sslExtension` (`bool`)
        * `sslServerHandshake` (`bool`)
        * `sslServerHello` (`bool`)
        * `startsWith` (`bool`)
        * `tcp` (`bool`)
        * `text` (`bool`)
        * `tmName` (`str`) - If Rule is used then you need to provide the tm_name it can be any value
        * `unnamedQueryParameter` (`bool`)
        * `userAgentToken` (`bool`)
        * `username` (`bool`)
        * `value` (`bool`)
        * `values` (`list`)
        * `version` (`bool`)
        * `vlan` (`bool`)
        * `vlanId` (`bool`)

      * `name` (`str`) - Name of the Policy
    """
    strategy: pulumi.Output[str]
    """
    Specifies the match strategy
    """
    def __init__(__self__, resource_name, opts=None, controls=None, name=None, published_copy=None, requires=None, rules=None, strategy=None, __props__=None, __name__=None, __opts__=None):
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
            rules=[{
                "name": "rule6",
                "actions": [{
                    "tmName": "20",
                    "forward": True,
                    "pool": "/Common/mypool",
                }],
            }],
            opts=ResourceOptions(depends_on=[bigip_ltm_pool["mypool"]]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] controls: Specifies the controls
        :param pulumi.Input[str] name: Name of the Policy
        :param pulumi.Input[str] published_copy: If you want to publish the policy else it will be deployed in Drafts mode.
        :param pulumi.Input[list] requires: Specifies the protocol
        :param pulumi.Input[list] rules: Rules can be applied using the policy
        :param pulumi.Input[str] strategy: Specifies the match strategy

        The **rules** object supports the following:

          * `actions` (`pulumi.Input[list]`)
            * `app_service` (`pulumi.Input[str]`)
            * `application` (`pulumi.Input[str]`)
            * `asm` (`pulumi.Input[bool]`)
            * `avr` (`pulumi.Input[bool]`)
            * `cache` (`pulumi.Input[bool]`)
            * `carp` (`pulumi.Input[bool]`)
            * `category` (`pulumi.Input[str]`)
            * `classify` (`pulumi.Input[bool]`)
            * `clonePool` (`pulumi.Input[str]`)
            * `code` (`pulumi.Input[float]`)
            * `compress` (`pulumi.Input[bool]`)
            * `content` (`pulumi.Input[str]`)
            * `cookieHash` (`pulumi.Input[bool]`)
            * `cookieInsert` (`pulumi.Input[bool]`)
            * `cookiePassive` (`pulumi.Input[bool]`)
            * `cookieRewrite` (`pulumi.Input[bool]`)
            * `decompress` (`pulumi.Input[bool]`)
            * `defer` (`pulumi.Input[bool]`)
            * `destinationAddress` (`pulumi.Input[bool]`)
            * `disable` (`pulumi.Input[bool]`)
            * `domain` (`pulumi.Input[str]`)
            * `enable` (`pulumi.Input[bool]`)
            * `expiry` (`pulumi.Input[str]`)
            * `expirySecs` (`pulumi.Input[float]`)
            * `expression` (`pulumi.Input[str]`)
            * `extension` (`pulumi.Input[str]`)
            * `facility` (`pulumi.Input[str]`)
            * `forward` (`pulumi.Input[bool]`) - This action will affect forwarding.
            * `fromProfile` (`pulumi.Input[str]`)
            * `hash` (`pulumi.Input[bool]`)
            * `host` (`pulumi.Input[str]`)
            * `http` (`pulumi.Input[bool]`)
            * `httpBasicAuth` (`pulumi.Input[bool]`)
            * `httpCookie` (`pulumi.Input[bool]`)
            * `httpHeader` (`pulumi.Input[bool]`)
            * `httpHost` (`pulumi.Input[bool]`)
            * `httpReferer` (`pulumi.Input[bool]`)
            * `httpReply` (`pulumi.Input[bool]`)
            * `httpSetCookie` (`pulumi.Input[bool]`)
            * `httpUri` (`pulumi.Input[bool]`)
            * `ifile` (`pulumi.Input[str]`)
            * `insert` (`pulumi.Input[bool]`)
            * `internalVirtual` (`pulumi.Input[str]`)
            * `ipAddress` (`pulumi.Input[str]`)
            * `key` (`pulumi.Input[str]`)
            * `l7dos` (`pulumi.Input[bool]`)
            * `length` (`pulumi.Input[float]`)
            * `location` (`pulumi.Input[str]`)
            * `log` (`pulumi.Input[bool]`)
            * `ltmPolicy` (`pulumi.Input[bool]`)
            * `member` (`pulumi.Input[str]`)
            * `message` (`pulumi.Input[str]`)
            * `netmask` (`pulumi.Input[str]`)
            * `nexthop` (`pulumi.Input[str]`)
            * `node` (`pulumi.Input[str]`)
            * `offset` (`pulumi.Input[float]`)
            * `path` (`pulumi.Input[str]`)
            * `pem` (`pulumi.Input[bool]`)
            * `persist` (`pulumi.Input[bool]`)
            * `pin` (`pulumi.Input[bool]`)
            * `policy` (`pulumi.Input[str]`)
            * `pool` (`pulumi.Input[str]`) - This action will direct the stream to this pool.
            * `port` (`pulumi.Input[float]`)
            * `priority` (`pulumi.Input[str]`)
            * `profile` (`pulumi.Input[str]`)
            * `protocol` (`pulumi.Input[str]`)
            * `queryString` (`pulumi.Input[str]`)
            * `rateclass` (`pulumi.Input[str]`)
            * `redirect` (`pulumi.Input[bool]`)
            * `remove` (`pulumi.Input[bool]`)
            * `replace` (`pulumi.Input[bool]`)
            * `request` (`pulumi.Input[bool]`)
            * `requestAdapt` (`pulumi.Input[bool]`)
            * `reset` (`pulumi.Input[bool]`)
            * `response` (`pulumi.Input[bool]`)
            * `responseAdapt` (`pulumi.Input[bool]`)
            * `scheme` (`pulumi.Input[str]`)
            * `script` (`pulumi.Input[str]`)
            * `select` (`pulumi.Input[bool]`)
            * `serverSsl` (`pulumi.Input[bool]`)
            * `setVariable` (`pulumi.Input[bool]`)
            * `snat` (`pulumi.Input[str]`)
            * `snatpool` (`pulumi.Input[str]`)
            * `sourceAddress` (`pulumi.Input[bool]`)
            * `sslClientHello` (`pulumi.Input[bool]`)
            * `sslServerHandshake` (`pulumi.Input[bool]`)
            * `sslServerHello` (`pulumi.Input[bool]`)
            * `sslSessionId` (`pulumi.Input[bool]`)
            * `status` (`pulumi.Input[float]`)
            * `tcl` (`pulumi.Input[bool]`)
            * `tcpNagle` (`pulumi.Input[bool]`)
            * `text` (`pulumi.Input[str]`)
            * `timeout` (`pulumi.Input[float]`)
            * `tmName` (`pulumi.Input[str]`) - If Rule is used then you need to provide the tm_name it can be any value
            * `uie` (`pulumi.Input[bool]`)
            * `universal` (`pulumi.Input[bool]`)
            * `value` (`pulumi.Input[str]`)
            * `virtual` (`pulumi.Input[str]`)
            * `vlan` (`pulumi.Input[str]`)
            * `vlanId` (`pulumi.Input[float]`)
            * `wam` (`pulumi.Input[bool]`)
            * `write` (`pulumi.Input[bool]`)

          * `conditions` (`pulumi.Input[list]`)
            * `address` (`pulumi.Input[bool]`)
            * `all` (`pulumi.Input[bool]`)
            * `app_service` (`pulumi.Input[str]`)
            * `browserType` (`pulumi.Input[bool]`)
            * `browserVersion` (`pulumi.Input[bool]`)
            * `caseInsensitive` (`pulumi.Input[bool]`)
            * `caseSensitive` (`pulumi.Input[bool]`)
            * `cipher` (`pulumi.Input[bool]`)
            * `cipherBits` (`pulumi.Input[bool]`)
            * `clientSsl` (`pulumi.Input[bool]`)
            * `code` (`pulumi.Input[bool]`)
            * `commonName` (`pulumi.Input[bool]`)
            * `contains` (`pulumi.Input[bool]`)
            * `continent` (`pulumi.Input[bool]`)
            * `countryCode` (`pulumi.Input[bool]`)
            * `countryName` (`pulumi.Input[bool]`)
            * `cpuUsage` (`pulumi.Input[bool]`)
            * `deviceMake` (`pulumi.Input[bool]`)
            * `deviceModel` (`pulumi.Input[bool]`)
            * `domain` (`pulumi.Input[bool]`)
            * `endsWith` (`pulumi.Input[bool]`)
            * `equals` (`pulumi.Input[bool]`)
            * `expiry` (`pulumi.Input[bool]`)
            * `extension` (`pulumi.Input[bool]`)
            * `external` (`pulumi.Input[bool]`)
            * `geoip` (`pulumi.Input[bool]`)
            * `greater` (`pulumi.Input[bool]`)
            * `greaterOrEqual` (`pulumi.Input[bool]`)
            * `host` (`pulumi.Input[bool]`)
            * `httpBasicAuth` (`pulumi.Input[bool]`)
            * `httpCookie` (`pulumi.Input[bool]`)
            * `httpHeader` (`pulumi.Input[bool]`)
            * `httpHost` (`pulumi.Input[bool]`)
            * `httpMethod` (`pulumi.Input[bool]`)
            * `httpReferer` (`pulumi.Input[bool]`)
            * `httpSetCookie` (`pulumi.Input[bool]`)
            * `httpStatus` (`pulumi.Input[bool]`)
            * `httpUri` (`pulumi.Input[bool]`)
            * `httpUserAgent` (`pulumi.Input[bool]`)
            * `httpVersion` (`pulumi.Input[bool]`)
            * `index` (`pulumi.Input[float]`)
            * `internal` (`pulumi.Input[bool]`)
            * `isp` (`pulumi.Input[bool]`)
            * `last15secs` (`pulumi.Input[bool]`)
            * `last1min` (`pulumi.Input[bool]`)
            * `last5mins` (`pulumi.Input[bool]`)
            * `less` (`pulumi.Input[bool]`)
            * `lessOrEqual` (`pulumi.Input[bool]`)
            * `local` (`pulumi.Input[bool]`)
            * `major` (`pulumi.Input[bool]`)
            * `matches` (`pulumi.Input[bool]`)
            * `minor` (`pulumi.Input[bool]`)
            * `missing` (`pulumi.Input[bool]`)
            * `mss` (`pulumi.Input[bool]`)
            * `not` (`pulumi.Input[bool]`)
            * `org` (`pulumi.Input[bool]`)
            * `password` (`pulumi.Input[bool]`)
            * `path` (`pulumi.Input[bool]`)
            * `pathSegment` (`pulumi.Input[bool]`)
            * `port` (`pulumi.Input[bool]`)
            * `present` (`pulumi.Input[bool]`)
            * `protocol` (`pulumi.Input[bool]`)
            * `queryParameter` (`pulumi.Input[bool]`)
            * `queryString` (`pulumi.Input[bool]`)
            * `regionCode` (`pulumi.Input[bool]`)
            * `regionName` (`pulumi.Input[bool]`)
            * `remote` (`pulumi.Input[bool]`)
            * `request` (`pulumi.Input[bool]`)
            * `response` (`pulumi.Input[bool]`)
            * `routeDomain` (`pulumi.Input[bool]`)
            * `rtt` (`pulumi.Input[bool]`)
            * `scheme` (`pulumi.Input[bool]`)
            * `server_name` (`pulumi.Input[bool]`)
            * `sslCert` (`pulumi.Input[bool]`)
            * `sslClientHello` (`pulumi.Input[bool]`)
            * `sslExtension` (`pulumi.Input[bool]`)
            * `sslServerHandshake` (`pulumi.Input[bool]`)
            * `sslServerHello` (`pulumi.Input[bool]`)
            * `startsWith` (`pulumi.Input[bool]`)
            * `tcp` (`pulumi.Input[bool]`)
            * `text` (`pulumi.Input[bool]`)
            * `tmName` (`pulumi.Input[str]`) - If Rule is used then you need to provide the tm_name it can be any value
            * `unnamedQueryParameter` (`pulumi.Input[bool]`)
            * `userAgentToken` (`pulumi.Input[bool]`)
            * `username` (`pulumi.Input[bool]`)
            * `value` (`pulumi.Input[bool]`)
            * `values` (`pulumi.Input[list]`)
            * `version` (`pulumi.Input[bool]`)
            * `vlan` (`pulumi.Input[bool]`)
            * `vlanId` (`pulumi.Input[bool]`)

          * `name` (`pulumi.Input[str]`) - Name of the Policy
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, controls=None, name=None, published_copy=None, requires=None, rules=None, strategy=None):
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] controls: Specifies the controls
        :param pulumi.Input[str] name: Name of the Policy
        :param pulumi.Input[str] published_copy: If you want to publish the policy else it will be deployed in Drafts mode.
        :param pulumi.Input[list] requires: Specifies the protocol
        :param pulumi.Input[list] rules: Rules can be applied using the policy
        :param pulumi.Input[str] strategy: Specifies the match strategy

        The **rules** object supports the following:

          * `actions` (`pulumi.Input[list]`)
            * `app_service` (`pulumi.Input[str]`)
            * `application` (`pulumi.Input[str]`)
            * `asm` (`pulumi.Input[bool]`)
            * `avr` (`pulumi.Input[bool]`)
            * `cache` (`pulumi.Input[bool]`)
            * `carp` (`pulumi.Input[bool]`)
            * `category` (`pulumi.Input[str]`)
            * `classify` (`pulumi.Input[bool]`)
            * `clonePool` (`pulumi.Input[str]`)
            * `code` (`pulumi.Input[float]`)
            * `compress` (`pulumi.Input[bool]`)
            * `content` (`pulumi.Input[str]`)
            * `cookieHash` (`pulumi.Input[bool]`)
            * `cookieInsert` (`pulumi.Input[bool]`)
            * `cookiePassive` (`pulumi.Input[bool]`)
            * `cookieRewrite` (`pulumi.Input[bool]`)
            * `decompress` (`pulumi.Input[bool]`)
            * `defer` (`pulumi.Input[bool]`)
            * `destinationAddress` (`pulumi.Input[bool]`)
            * `disable` (`pulumi.Input[bool]`)
            * `domain` (`pulumi.Input[str]`)
            * `enable` (`pulumi.Input[bool]`)
            * `expiry` (`pulumi.Input[str]`)
            * `expirySecs` (`pulumi.Input[float]`)
            * `expression` (`pulumi.Input[str]`)
            * `extension` (`pulumi.Input[str]`)
            * `facility` (`pulumi.Input[str]`)
            * `forward` (`pulumi.Input[bool]`) - This action will affect forwarding.
            * `fromProfile` (`pulumi.Input[str]`)
            * `hash` (`pulumi.Input[bool]`)
            * `host` (`pulumi.Input[str]`)
            * `http` (`pulumi.Input[bool]`)
            * `httpBasicAuth` (`pulumi.Input[bool]`)
            * `httpCookie` (`pulumi.Input[bool]`)
            * `httpHeader` (`pulumi.Input[bool]`)
            * `httpHost` (`pulumi.Input[bool]`)
            * `httpReferer` (`pulumi.Input[bool]`)
            * `httpReply` (`pulumi.Input[bool]`)
            * `httpSetCookie` (`pulumi.Input[bool]`)
            * `httpUri` (`pulumi.Input[bool]`)
            * `ifile` (`pulumi.Input[str]`)
            * `insert` (`pulumi.Input[bool]`)
            * `internalVirtual` (`pulumi.Input[str]`)
            * `ipAddress` (`pulumi.Input[str]`)
            * `key` (`pulumi.Input[str]`)
            * `l7dos` (`pulumi.Input[bool]`)
            * `length` (`pulumi.Input[float]`)
            * `location` (`pulumi.Input[str]`)
            * `log` (`pulumi.Input[bool]`)
            * `ltmPolicy` (`pulumi.Input[bool]`)
            * `member` (`pulumi.Input[str]`)
            * `message` (`pulumi.Input[str]`)
            * `netmask` (`pulumi.Input[str]`)
            * `nexthop` (`pulumi.Input[str]`)
            * `node` (`pulumi.Input[str]`)
            * `offset` (`pulumi.Input[float]`)
            * `path` (`pulumi.Input[str]`)
            * `pem` (`pulumi.Input[bool]`)
            * `persist` (`pulumi.Input[bool]`)
            * `pin` (`pulumi.Input[bool]`)
            * `policy` (`pulumi.Input[str]`)
            * `pool` (`pulumi.Input[str]`) - This action will direct the stream to this pool.
            * `port` (`pulumi.Input[float]`)
            * `priority` (`pulumi.Input[str]`)
            * `profile` (`pulumi.Input[str]`)
            * `protocol` (`pulumi.Input[str]`)
            * `queryString` (`pulumi.Input[str]`)
            * `rateclass` (`pulumi.Input[str]`)
            * `redirect` (`pulumi.Input[bool]`)
            * `remove` (`pulumi.Input[bool]`)
            * `replace` (`pulumi.Input[bool]`)
            * `request` (`pulumi.Input[bool]`)
            * `requestAdapt` (`pulumi.Input[bool]`)
            * `reset` (`pulumi.Input[bool]`)
            * `response` (`pulumi.Input[bool]`)
            * `responseAdapt` (`pulumi.Input[bool]`)
            * `scheme` (`pulumi.Input[str]`)
            * `script` (`pulumi.Input[str]`)
            * `select` (`pulumi.Input[bool]`)
            * `serverSsl` (`pulumi.Input[bool]`)
            * `setVariable` (`pulumi.Input[bool]`)
            * `snat` (`pulumi.Input[str]`)
            * `snatpool` (`pulumi.Input[str]`)
            * `sourceAddress` (`pulumi.Input[bool]`)
            * `sslClientHello` (`pulumi.Input[bool]`)
            * `sslServerHandshake` (`pulumi.Input[bool]`)
            * `sslServerHello` (`pulumi.Input[bool]`)
            * `sslSessionId` (`pulumi.Input[bool]`)
            * `status` (`pulumi.Input[float]`)
            * `tcl` (`pulumi.Input[bool]`)
            * `tcpNagle` (`pulumi.Input[bool]`)
            * `text` (`pulumi.Input[str]`)
            * `timeout` (`pulumi.Input[float]`)
            * `tmName` (`pulumi.Input[str]`) - If Rule is used then you need to provide the tm_name it can be any value
            * `uie` (`pulumi.Input[bool]`)
            * `universal` (`pulumi.Input[bool]`)
            * `value` (`pulumi.Input[str]`)
            * `virtual` (`pulumi.Input[str]`)
            * `vlan` (`pulumi.Input[str]`)
            * `vlanId` (`pulumi.Input[float]`)
            * `wam` (`pulumi.Input[bool]`)
            * `write` (`pulumi.Input[bool]`)

          * `conditions` (`pulumi.Input[list]`)
            * `address` (`pulumi.Input[bool]`)
            * `all` (`pulumi.Input[bool]`)
            * `app_service` (`pulumi.Input[str]`)
            * `browserType` (`pulumi.Input[bool]`)
            * `browserVersion` (`pulumi.Input[bool]`)
            * `caseInsensitive` (`pulumi.Input[bool]`)
            * `caseSensitive` (`pulumi.Input[bool]`)
            * `cipher` (`pulumi.Input[bool]`)
            * `cipherBits` (`pulumi.Input[bool]`)
            * `clientSsl` (`pulumi.Input[bool]`)
            * `code` (`pulumi.Input[bool]`)
            * `commonName` (`pulumi.Input[bool]`)
            * `contains` (`pulumi.Input[bool]`)
            * `continent` (`pulumi.Input[bool]`)
            * `countryCode` (`pulumi.Input[bool]`)
            * `countryName` (`pulumi.Input[bool]`)
            * `cpuUsage` (`pulumi.Input[bool]`)
            * `deviceMake` (`pulumi.Input[bool]`)
            * `deviceModel` (`pulumi.Input[bool]`)
            * `domain` (`pulumi.Input[bool]`)
            * `endsWith` (`pulumi.Input[bool]`)
            * `equals` (`pulumi.Input[bool]`)
            * `expiry` (`pulumi.Input[bool]`)
            * `extension` (`pulumi.Input[bool]`)
            * `external` (`pulumi.Input[bool]`)
            * `geoip` (`pulumi.Input[bool]`)
            * `greater` (`pulumi.Input[bool]`)
            * `greaterOrEqual` (`pulumi.Input[bool]`)
            * `host` (`pulumi.Input[bool]`)
            * `httpBasicAuth` (`pulumi.Input[bool]`)
            * `httpCookie` (`pulumi.Input[bool]`)
            * `httpHeader` (`pulumi.Input[bool]`)
            * `httpHost` (`pulumi.Input[bool]`)
            * `httpMethod` (`pulumi.Input[bool]`)
            * `httpReferer` (`pulumi.Input[bool]`)
            * `httpSetCookie` (`pulumi.Input[bool]`)
            * `httpStatus` (`pulumi.Input[bool]`)
            * `httpUri` (`pulumi.Input[bool]`)
            * `httpUserAgent` (`pulumi.Input[bool]`)
            * `httpVersion` (`pulumi.Input[bool]`)
            * `index` (`pulumi.Input[float]`)
            * `internal` (`pulumi.Input[bool]`)
            * `isp` (`pulumi.Input[bool]`)
            * `last15secs` (`pulumi.Input[bool]`)
            * `last1min` (`pulumi.Input[bool]`)
            * `last5mins` (`pulumi.Input[bool]`)
            * `less` (`pulumi.Input[bool]`)
            * `lessOrEqual` (`pulumi.Input[bool]`)
            * `local` (`pulumi.Input[bool]`)
            * `major` (`pulumi.Input[bool]`)
            * `matches` (`pulumi.Input[bool]`)
            * `minor` (`pulumi.Input[bool]`)
            * `missing` (`pulumi.Input[bool]`)
            * `mss` (`pulumi.Input[bool]`)
            * `not` (`pulumi.Input[bool]`)
            * `org` (`pulumi.Input[bool]`)
            * `password` (`pulumi.Input[bool]`)
            * `path` (`pulumi.Input[bool]`)
            * `pathSegment` (`pulumi.Input[bool]`)
            * `port` (`pulumi.Input[bool]`)
            * `present` (`pulumi.Input[bool]`)
            * `protocol` (`pulumi.Input[bool]`)
            * `queryParameter` (`pulumi.Input[bool]`)
            * `queryString` (`pulumi.Input[bool]`)
            * `regionCode` (`pulumi.Input[bool]`)
            * `regionName` (`pulumi.Input[bool]`)
            * `remote` (`pulumi.Input[bool]`)
            * `request` (`pulumi.Input[bool]`)
            * `response` (`pulumi.Input[bool]`)
            * `routeDomain` (`pulumi.Input[bool]`)
            * `rtt` (`pulumi.Input[bool]`)
            * `scheme` (`pulumi.Input[bool]`)
            * `server_name` (`pulumi.Input[bool]`)
            * `sslCert` (`pulumi.Input[bool]`)
            * `sslClientHello` (`pulumi.Input[bool]`)
            * `sslExtension` (`pulumi.Input[bool]`)
            * `sslServerHandshake` (`pulumi.Input[bool]`)
            * `sslServerHello` (`pulumi.Input[bool]`)
            * `startsWith` (`pulumi.Input[bool]`)
            * `tcp` (`pulumi.Input[bool]`)
            * `text` (`pulumi.Input[bool]`)
            * `tmName` (`pulumi.Input[str]`) - If Rule is used then you need to provide the tm_name it can be any value
            * `unnamedQueryParameter` (`pulumi.Input[bool]`)
            * `userAgentToken` (`pulumi.Input[bool]`)
            * `username` (`pulumi.Input[bool]`)
            * `value` (`pulumi.Input[bool]`)
            * `values` (`pulumi.Input[list]`)
            * `version` (`pulumi.Input[bool]`)
            * `vlan` (`pulumi.Input[bool]`)
            * `vlanId` (`pulumi.Input[bool]`)

          * `name` (`pulumi.Input[str]`) - Name of the Policy
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
