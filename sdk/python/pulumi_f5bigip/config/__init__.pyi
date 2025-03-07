# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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

address: Optional[str]
"""
Domain name/IP of the BigIP
"""

apiRetries: Optional[int]
"""
Amount of times to retry AS3 API requests. Default: 10.
"""

apiTimeout: Optional[int]
"""
A timeout for AS3 requests, represented as a number of seconds. Default: 60
"""

loginRef: Optional[str]
"""
Login reference for token authentication (see BIG-IP REST docs for details)
"""

password: Optional[str]
"""
The user's password. Leave empty if using token_value
"""

port: Optional[str]
"""
Management Port to connect to Bigip
"""

teemDisable: Optional[bool]
"""
If this flag set to true,sending telemetry data to TEEM will be disabled
"""

tokenAuth: Optional[bool]
"""
Enable to use token authentication. Can be set via the BIGIP_TOKEN_AUTH environment variable
"""

tokenTimeout: Optional[int]
"""
A lifespan to request for the AS3 auth token, represented as a number of seconds. Default: 1200
"""

tokenValue: Optional[str]
"""
A token generated outside the provider, in place of password
"""

trustedCertPath: Optional[str]
"""
Valid Trusted Certificate path
"""

username: Optional[str]
"""
Username with API access to the BigIP
"""

validateCertsDisable: Optional[bool]
"""
If set to true, Disables TLS certificate check on BIG-IP. Default : True
"""

