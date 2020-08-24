# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'address',
    'login_ref',
    'password',
    'port',
    'teem_disable',
    'token_auth',
    'username',
]

__config__ = pulumi.Config('f5bigip')

address = __config__.get('address')
"""
Domain name/IP of the BigIP
"""

login_ref = __config__.get('loginRef')
"""
Login reference for token authentication (see BIG-IP REST docs for details)
"""

password = __config__.get('password')
"""
The user's password
"""

port = __config__.get('port')
"""
Management Port to connect to Bigip
"""

teem_disable = __config__.get('teemDisable')
"""
If this flag set to true,sending telemetry data to TEEM will be disabled
"""

token_auth = __config__.get('tokenAuth')
"""
Enable to use an external authentication source (LDAP, TACACS, etc)
"""

username = __config__.get('username')
"""
Username with API access to the BigIP
"""

