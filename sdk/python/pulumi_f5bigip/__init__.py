# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
from . import _utilities
import typing
# Export this package's modules as members:
from .as3 import *
from .big_iq_as3 import *
from .command import *
from .common_license_manage_big_iq import *
from .do import *
from .event_service_discovery import *
from .fast_application import *
from .fast_http_app import *
from .fast_https_app import *
from .fast_tcp_app import *
from .fast_template import *
from .fast_udp_app import *
from .ipsec_policy import *
from .ipsec_profile import *
from .net_ike_peer import *
from .net_tunnel import *
from .partition import *
from .provider import *
from .saas_bot_defense_profile import *
from .ssl_key_cert import *
from .traffic_selector import *
from .waf_policy import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_f5bigip.cm as __cm
    cm = __cm
    import pulumi_f5bigip.config as __config
    config = __config
    import pulumi_f5bigip.fast as __fast
    fast = __fast
    import pulumi_f5bigip.ltm as __ltm
    ltm = __ltm
    import pulumi_f5bigip.net as __net
    net = __net
    import pulumi_f5bigip.ssl as __ssl
    ssl = __ssl
    import pulumi_f5bigip.sys as __sys
    sys = __sys
    import pulumi_f5bigip.vcmp as __vcmp
    vcmp = __vcmp
else:
    cm = _utilities.lazy_import('pulumi_f5bigip.cm')
    config = _utilities.lazy_import('pulumi_f5bigip.config')
    fast = _utilities.lazy_import('pulumi_f5bigip.fast')
    ltm = _utilities.lazy_import('pulumi_f5bigip.ltm')
    net = _utilities.lazy_import('pulumi_f5bigip.net')
    ssl = _utilities.lazy_import('pulumi_f5bigip.ssl')
    sys = _utilities.lazy_import('pulumi_f5bigip.sys')
    vcmp = _utilities.lazy_import('pulumi_f5bigip.vcmp')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "f5bigip",
  "mod": "cm/device",
  "fqn": "pulumi_f5bigip.cm",
  "classes": {
   "f5bigip:cm/device:Device": "Device"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "cm/deviceGroup",
  "fqn": "pulumi_f5bigip.cm",
  "classes": {
   "f5bigip:cm/deviceGroup:DeviceGroup": "DeviceGroup"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/as3",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/as3:As3": "As3"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/bigIqAs3",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/bigIqAs3:BigIqAs3": "BigIqAs3"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/command",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/command:Command": "Command"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/commonLicenseManageBigIq",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq": "CommonLicenseManageBigIq"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/do",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/do:Do": "Do"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/eventServiceDiscovery",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/eventServiceDiscovery:EventServiceDiscovery": "EventServiceDiscovery"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/fastApplication",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/fastApplication:FastApplication": "FastApplication"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/fastHttpApp",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/fastHttpApp:FastHttpApp": "FastHttpApp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/fastHttpsApp",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/fastHttpsApp:FastHttpsApp": "FastHttpsApp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/fastTcpApp",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/fastTcpApp:FastTcpApp": "FastTcpApp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/fastTemplate",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/fastTemplate:FastTemplate": "FastTemplate"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/fastUdpApp",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/fastUdpApp:FastUdpApp": "FastUdpApp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/ipsecPolicy",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/ipsecPolicy:IpsecPolicy": "IpsecPolicy"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/ipsecProfile",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/ipsecProfile:IpsecProfile": "IpsecProfile"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/netIkePeer",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/netIkePeer:NetIkePeer": "NetIkePeer"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/netTunnel",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/netTunnel:NetTunnel": "NetTunnel"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/partition",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/partition:Partition": "Partition"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/saasBotDefenseProfile",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/saasBotDefenseProfile:SaasBotDefenseProfile": "SaasBotDefenseProfile"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/sslKeyCert",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/sslKeyCert:SslKeyCert": "SslKeyCert"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/trafficSelector",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/trafficSelector:TrafficSelector": "TrafficSelector"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "index/wafPolicy",
  "fqn": "pulumi_f5bigip",
  "classes": {
   "f5bigip:index/wafPolicy:WafPolicy": "WafPolicy"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/cipherGroup",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/cipherGroup:CipherGroup": "CipherGroup"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/cipherRule",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/cipherRule:CipherRule": "CipherRule"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/dataGroup",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/dataGroup:DataGroup": "DataGroup"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/iRule",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/iRule:IRule": "IRule"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/monitor",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/monitor:Monitor": "Monitor"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/node",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/node:Node": "Node"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/persistenceProfileCookie",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie": "PersistenceProfileCookie"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/persistenceProfileDstAddr",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr": "PersistenceProfileDstAddr"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/persistenceProfileSrcAddr",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr": "PersistenceProfileSrcAddr"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/persistenceProfileSsl",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl": "PersistenceProfileSsl"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/policy",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/policy:Policy": "Policy"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/pool",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/pool:Pool": "Pool"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/poolAttachment",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/poolAttachment:PoolAttachment": "PoolAttachment"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileBotDefense",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileBotDefense:ProfileBotDefense": "ProfileBotDefense"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileClientSsl",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileClientSsl:ProfileClientSsl": "ProfileClientSsl"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileFastHttp",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileFastHttp:ProfileFastHttp": "ProfileFastHttp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileFastL4",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileFastL4:ProfileFastL4": "ProfileFastL4"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileFtp",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileFtp:ProfileFtp": "ProfileFtp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileHttp",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileHttp:ProfileHttp": "ProfileHttp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileHttp2",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileHttp2:ProfileHttp2": "ProfileHttp2"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileHttpCompress",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileHttpCompress:ProfileHttpCompress": "ProfileHttpCompress"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileOneConnect",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileOneConnect:ProfileOneConnect": "ProfileOneConnect"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileRewrite",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileRewrite:ProfileRewrite": "ProfileRewrite"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileRewriteUriRules",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileRewriteUriRules:ProfileRewriteUriRules": "ProfileRewriteUriRules"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileServerSsl",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileServerSsl:ProfileServerSsl": "ProfileServerSsl"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileTcp",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileTcp:ProfileTcp": "ProfileTcp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/profileWebAcceleration",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/profileWebAcceleration:ProfileWebAcceleration": "ProfileWebAcceleration"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/requestLogProfile",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/requestLogProfile:RequestLogProfile": "RequestLogProfile"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/snat",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/snat:Snat": "Snat"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/snatPool",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/snatPool:SnatPool": "SnatPool"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/virtualAddress",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/virtualAddress:VirtualAddress": "VirtualAddress"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ltm/virtualServer",
  "fqn": "pulumi_f5bigip.ltm",
  "classes": {
   "f5bigip:ltm/virtualServer:VirtualServer": "VirtualServer"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "net/route",
  "fqn": "pulumi_f5bigip.net",
  "classes": {
   "f5bigip:net/route:Route": "Route"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "net/selfIp",
  "fqn": "pulumi_f5bigip.net",
  "classes": {
   "f5bigip:net/selfIp:SelfIp": "SelfIp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "net/vlan",
  "fqn": "pulumi_f5bigip.net",
  "classes": {
   "f5bigip:net/vlan:Vlan": "Vlan"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ssl/certificate",
  "fqn": "pulumi_f5bigip.ssl",
  "classes": {
   "f5bigip:ssl/certificate:Certificate": "Certificate"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "ssl/key",
  "fqn": "pulumi_f5bigip.ssl",
  "classes": {
   "f5bigip:ssl/key:Key": "Key"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/bigIpLicense",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/bigIpLicense:BigIpLicense": "BigIpLicense"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/dns",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/dns:Dns": "Dns"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/iApp",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/iApp:IApp": "IApp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/ntp",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/ntp:Ntp": "Ntp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/ocsp",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/ocsp:Ocsp": "Ocsp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/provision",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/provision:Provision": "Provision"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/snmp",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/snmp:Snmp": "Snmp"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "sys/snmpTraps",
  "fqn": "pulumi_f5bigip.sys",
  "classes": {
   "f5bigip:sys/snmpTraps:SnmpTraps": "SnmpTraps"
  }
 },
 {
  "pkg": "f5bigip",
  "mod": "vcmp/guest",
  "fqn": "pulumi_f5bigip.vcmp",
  "classes": {
   "f5bigip:vcmp/guest:Guest": "Guest"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "f5bigip",
  "token": "pulumi:providers:f5bigip",
  "fqn": "pulumi_f5bigip",
  "class": "Provider"
 }
]
"""
)
