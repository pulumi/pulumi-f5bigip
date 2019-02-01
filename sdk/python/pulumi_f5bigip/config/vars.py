# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

__config__ = pulumi.Config('f5bigip')

address = __config__.require('address')
"""
Domain name/IP of the BigIP
"""

login_ref = __config__.get('loginRef')
"""
Login reference for token authentication (see BIG-IP REST docs for details)
"""

password = __config__.require('password')
"""
The user's password
"""

token_auth = __config__.get('tokenAuth')
"""
Enable to use an external authentication source (LDAP, TACACS, etc)
"""

username = __config__.require('username')
"""
Username with API access to the BigIP
"""

