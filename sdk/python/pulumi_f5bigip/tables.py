# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

_SNAKE_TO_CAMEL_CASE_TABLE = {
    "accept_xff": "acceptXff",
    "activation_modes": "activationModes",
    "adaptive_limit": "adaptiveLimit",
    "advertize_route": "advertizeRoute",
    "alert_timeout": "alertTimeout",
    "allow_nat": "allowNat",
    "allow_non_ssl": "allowNonSsl",
    "allow_snat": "allowSnat",
    "always_send": "alwaysSend",
    "app_service": "appService",
    "as3_json": "as3Json",
    "auth_passwordencrypted": "authPasswordencrypted",
    "auth_protocol": "authProtocol",
    "authenticate_depth": "authenticateDepth",
    "auto_delete": "autoDelete",
    "auto_sync": "autoSync",
    "basic_auth_realm": "basicAuthRealm",
    "ca_file": "caFile",
    "cache_size": "cacheSize",
    "cache_timeout": "cacheTimeout",
    "cert_extension_includes": "certExtensionIncludes",
    "cert_key_chains": "certKeyChains",
    "cert_life_span": "certLifeSpan",
    "cert_lookup_by_ipaddr_port": "certLookupByIpaddrPort",
    "client_cert_ca": "clientCertCa",
    "client_profiles": "clientProfiles",
    "client_timeout": "clientTimeout",
    "close_wait_timeout": "closeWaitTimeout",
    "concurrent_streams_per_connection": "concurrentStreamsPerConnection",
    "configsync_ip": "configsyncIp",
    "conn_limit": "connLimit",
    "connection_idle_timeout": "connectionIdleTimeout",
    "connection_limit": "connectionLimit",
    "connpool_maxreuse": "connpoolMaxreuse",
    "connpool_maxsize": "connpoolMaxsize",
    "connpool_minsize": "connpoolMinsize",
    "connpool_replenish": "connpoolReplenish",
    "connpool_step": "connpoolStep",
    "connpoolidle_timeoutoverride": "connpoolidleTimeoutoverride",
    "content_type_excludes": "contentTypeExcludes",
    "content_type_includes": "contentTypeIncludes",
    "cookie_encryption": "cookieEncryption",
    "cookie_encryption_passphrase": "cookieEncryptionPassphrase",
    "cookie_name": "cookieName",
    "cpu_ratio": "cpuRatio",
    "crl_file": "crlFile",
    "default_persistence_profile": "defaultPersistenceProfile",
    "defaults_from": "defaultsFrom",
    "deferred_accept": "deferredAccept",
    "disk_ratio": "diskRatio",
    "do_json": "doJson",
    "dynamic_ratio": "dynamicRatio",
    "encrypt_cookie_secret": "encryptCookieSecret",
    "encrypt_cookies": "encryptCookies",
    "engine_id": "engineId",
    "execute_action": "executeAction",
    "expire_cert_response_control": "expireCertResponseControl",
    "explicitflow_migration": "explicitflowMigration",
    "fallback_host": "fallbackHost",
    "fallback_persistence_profile": "fallbackPersistenceProfile",
    "fallback_status_codes": "fallbackStatusCodes",
    "fast_open": "fastOpen",
    "finwait_timeout": "finwaitTimeout",
    "forward_proxy_bypass_default_action": "forwardProxyBypassDefaultAction",
    "full_load_on_sync": "fullLoadOnSync",
    "full_path": "fullPath",
    "generic_alert": "genericAlert",
    "handshake_timeout": "handshakeTimeout",
    "hardware_syncookie": "hardwareSyncookie",
    "hash_algorithm": "hashAlgorithm",
    "hash_length": "hashLength",
    "hash_offset": "hashOffset",
    "head_erase": "headErase",
    "head_insert": "headInsert",
    "header_table_size": "headerTableSize",
    "icmp_echo": "icmpEcho",
    "idle_timeout": "idleTimeout",
    "idle_timeout_override": "idleTimeoutOverride",
    "incremental_config": "incrementalConfig",
    "inherit_cert_keychain": "inheritCertKeychain",
    "inherited_devicegroup": "inheritedDevicegroup",
    "inherited_traffic_group": "inheritedTrafficGroup",
    "insert_xforwarded_for": "insertXforwardedFor",
    "ip_dscp": "ipDscp",
    "ip_protocol": "ipProtocol",
    "iptos_toclient": "iptosToclient",
    "iptos_toserver": "iptosToserver",
    "keepalive_interval": "keepaliveInterval",
    "load_balancing_mode": "loadBalancingMode",
    "lws_separator": "lwsSeparator",
    "manual_resume": "manualResume",
    "map_proxies": "mapProxies",
    "match_across_pools": "matchAcrossPools",
    "match_across_services": "matchAcrossServices",
    "match_across_virtuals": "matchAcrossVirtuals",
    "max_age": "maxAge",
    "max_reuse": "maxReuse",
    "max_size": "maxSize",
    "maxheader_size": "maxheaderSize",
    "memory_ratio": "memoryRatio",
    "mirror_ip": "mirrorIp",
    "mirror_secondary_ip": "mirrorSecondaryIp",
    "mod_ssl_methods": "modSslMethods",
    "name_servers": "nameServers",
    "network_failover": "networkFailover",
    "number_of_dots": "numberOfDots",
    "oneconnect_transformations": "oneconnectTransformations",
    "override_conn_limit": "overrideConnLimit",
    "peer_cert_mode": "peerCertMode",
    "persistence_profiles": "persistenceProfiles",
    "privacy_password": "privacyPassword",
    "privacy_password_encrypted": "privacyPasswordEncrypted",
    "privacy_protocol": "privacyProtocol",
    "proxy_ca_cert": "proxyCaCert",
    "proxy_ca_key": "proxyCaKey",
    "proxy_ca_passphrase": "proxyCaPassphrase",
    "proxy_ssl": "proxySsl",
    "proxy_ssl_passthrough": "proxySslPassthrough",
    "proxy_type": "proxyType",
    "published_copy": "publishedCopy",
    "rate_limit": "rateLimit",
    "receive_disable": "receiveDisable",
    "redirect_rewrite": "redirectRewrite",
    "registration_key": "registrationKey",
    "renegotiate_period": "renegotiatePeriod",
    "renegotiate_size": "renegotiateSize",
    "request_chunking": "requestChunking",
    "reselect_tries": "reselectTries",
    "response_chunking": "responseChunking",
    "response_headers_permitteds": "responseHeadersPermitteds",
    "retain_certificate": "retainCertificate",
    "save_on_auto_sync": "saveOnAutoSync",
    "secure_renegotiation": "secureRenegotiation",
    "security_level": "securityLevel",
    "security_name": "securityName",
    "server_agent_name": "serverAgentName",
    "server_name": "serverName",
    "server_profiles": "serverProfiles",
    "service_down_action": "serviceDownAction",
    "session_mirroring": "sessionMirroring",
    "session_ticket": "sessionTicket",
    "share_pools": "sharePools",
    "slow_ramp_time": "slowRampTime",
    "sni_default": "sniDefault",
    "sni_require": "sniRequire",
    "source_address_translation": "sourceAddressTranslation",
    "source_mask": "sourceMask",
    "ssl_forward_proxy": "sslForwardProxy",
    "ssl_forward_proxy_bypass": "sslForwardProxyBypass",
    "ssl_sign_hash": "sslSignHash",
    "strict_resume": "strictResume",
    "strict_updates": "strictUpdates",
    "sys_contact": "sysContact",
    "sys_location": "sysLocation",
    "template_modified": "templateModified",
    "template_prerequisite_errors": "templatePrerequisiteErrors",
    "tenant_filter": "tenantFilter",
    "tenant_list": "tenantList",
    "tenant_name": "tenantName",
    "time_until_up": "timeUntilUp",
    "tm_options": "tmOptions",
    "tm_partition": "tmPartition",
    "traffic_group": "trafficGroup",
    "translate_address": "translateAddress",
    "translate_port": "translatePort",
    "unclean_shutdown": "uncleanShutdown",
    "untrusted_cert_response_control": "untrustedCertResponseControl",
    "uri_excludes": "uriExcludes",
    "uri_includes": "uriIncludes",
    "via_host_name": "viaHostName",
    "via_request": "viaRequest",
    "via_response": "viaResponse",
    "vlans_enabled": "vlansEnabled",
    "xff_alternative_names": "xffAlternativeNames",
}

_CAMEL_TO_SNAKE_CASE_TABLE = {
    "acceptXff": "accept_xff",
    "activationModes": "activation_modes",
    "adaptiveLimit": "adaptive_limit",
    "advertizeRoute": "advertize_route",
    "alertTimeout": "alert_timeout",
    "allowNat": "allow_nat",
    "allowNonSsl": "allow_non_ssl",
    "allowSnat": "allow_snat",
    "alwaysSend": "always_send",
    "appService": "app_service",
    "as3Json": "as3_json",
    "authPasswordencrypted": "auth_passwordencrypted",
    "authProtocol": "auth_protocol",
    "authenticateDepth": "authenticate_depth",
    "autoDelete": "auto_delete",
    "autoSync": "auto_sync",
    "basicAuthRealm": "basic_auth_realm",
    "caFile": "ca_file",
    "cacheSize": "cache_size",
    "cacheTimeout": "cache_timeout",
    "certExtensionIncludes": "cert_extension_includes",
    "certKeyChains": "cert_key_chains",
    "certLifeSpan": "cert_life_span",
    "certLookupByIpaddrPort": "cert_lookup_by_ipaddr_port",
    "clientCertCa": "client_cert_ca",
    "clientProfiles": "client_profiles",
    "clientTimeout": "client_timeout",
    "closeWaitTimeout": "close_wait_timeout",
    "concurrentStreamsPerConnection": "concurrent_streams_per_connection",
    "configsyncIp": "configsync_ip",
    "connLimit": "conn_limit",
    "connectionIdleTimeout": "connection_idle_timeout",
    "connectionLimit": "connection_limit",
    "connpoolMaxreuse": "connpool_maxreuse",
    "connpoolMaxsize": "connpool_maxsize",
    "connpoolMinsize": "connpool_minsize",
    "connpoolReplenish": "connpool_replenish",
    "connpoolStep": "connpool_step",
    "connpoolidleTimeoutoverride": "connpoolidle_timeoutoverride",
    "contentTypeExcludes": "content_type_excludes",
    "contentTypeIncludes": "content_type_includes",
    "cookieEncryption": "cookie_encryption",
    "cookieEncryptionPassphrase": "cookie_encryption_passphrase",
    "cookieName": "cookie_name",
    "cpuRatio": "cpu_ratio",
    "crlFile": "crl_file",
    "defaultPersistenceProfile": "default_persistence_profile",
    "defaultsFrom": "defaults_from",
    "deferredAccept": "deferred_accept",
    "diskRatio": "disk_ratio",
    "doJson": "do_json",
    "dynamicRatio": "dynamic_ratio",
    "encryptCookieSecret": "encrypt_cookie_secret",
    "encryptCookies": "encrypt_cookies",
    "engineId": "engine_id",
    "executeAction": "execute_action",
    "expireCertResponseControl": "expire_cert_response_control",
    "explicitflowMigration": "explicitflow_migration",
    "fallbackHost": "fallback_host",
    "fallbackPersistenceProfile": "fallback_persistence_profile",
    "fallbackStatusCodes": "fallback_status_codes",
    "fastOpen": "fast_open",
    "finwaitTimeout": "finwait_timeout",
    "forwardProxyBypassDefaultAction": "forward_proxy_bypass_default_action",
    "fullLoadOnSync": "full_load_on_sync",
    "fullPath": "full_path",
    "genericAlert": "generic_alert",
    "handshakeTimeout": "handshake_timeout",
    "hardwareSyncookie": "hardware_syncookie",
    "hashAlgorithm": "hash_algorithm",
    "hashLength": "hash_length",
    "hashOffset": "hash_offset",
    "headErase": "head_erase",
    "headInsert": "head_insert",
    "headerTableSize": "header_table_size",
    "icmpEcho": "icmp_echo",
    "idleTimeout": "idle_timeout",
    "idleTimeoutOverride": "idle_timeout_override",
    "incrementalConfig": "incremental_config",
    "inheritCertKeychain": "inherit_cert_keychain",
    "inheritedDevicegroup": "inherited_devicegroup",
    "inheritedTrafficGroup": "inherited_traffic_group",
    "insertXforwardedFor": "insert_xforwarded_for",
    "ipDscp": "ip_dscp",
    "ipProtocol": "ip_protocol",
    "iptosToclient": "iptos_toclient",
    "iptosToserver": "iptos_toserver",
    "keepaliveInterval": "keepalive_interval",
    "loadBalancingMode": "load_balancing_mode",
    "lwsSeparator": "lws_separator",
    "manualResume": "manual_resume",
    "mapProxies": "map_proxies",
    "matchAcrossPools": "match_across_pools",
    "matchAcrossServices": "match_across_services",
    "matchAcrossVirtuals": "match_across_virtuals",
    "maxAge": "max_age",
    "maxReuse": "max_reuse",
    "maxSize": "max_size",
    "maxheaderSize": "maxheader_size",
    "memoryRatio": "memory_ratio",
    "mirrorIp": "mirror_ip",
    "mirrorSecondaryIp": "mirror_secondary_ip",
    "modSslMethods": "mod_ssl_methods",
    "nameServers": "name_servers",
    "networkFailover": "network_failover",
    "numberOfDots": "number_of_dots",
    "oneconnectTransformations": "oneconnect_transformations",
    "overrideConnLimit": "override_conn_limit",
    "peerCertMode": "peer_cert_mode",
    "persistenceProfiles": "persistence_profiles",
    "privacyPassword": "privacy_password",
    "privacyPasswordEncrypted": "privacy_password_encrypted",
    "privacyProtocol": "privacy_protocol",
    "proxyCaCert": "proxy_ca_cert",
    "proxyCaKey": "proxy_ca_key",
    "proxyCaPassphrase": "proxy_ca_passphrase",
    "proxySsl": "proxy_ssl",
    "proxySslPassthrough": "proxy_ssl_passthrough",
    "proxyType": "proxy_type",
    "publishedCopy": "published_copy",
    "rateLimit": "rate_limit",
    "receiveDisable": "receive_disable",
    "redirectRewrite": "redirect_rewrite",
    "registrationKey": "registration_key",
    "renegotiatePeriod": "renegotiate_period",
    "renegotiateSize": "renegotiate_size",
    "requestChunking": "request_chunking",
    "reselectTries": "reselect_tries",
    "responseChunking": "response_chunking",
    "responseHeadersPermitteds": "response_headers_permitteds",
    "retainCertificate": "retain_certificate",
    "saveOnAutoSync": "save_on_auto_sync",
    "secureRenegotiation": "secure_renegotiation",
    "securityLevel": "security_level",
    "securityName": "security_name",
    "serverAgentName": "server_agent_name",
    "serverName": "server_name",
    "serverProfiles": "server_profiles",
    "serviceDownAction": "service_down_action",
    "sessionMirroring": "session_mirroring",
    "sessionTicket": "session_ticket",
    "sharePools": "share_pools",
    "slowRampTime": "slow_ramp_time",
    "sniDefault": "sni_default",
    "sniRequire": "sni_require",
    "sourceAddressTranslation": "source_address_translation",
    "sourceMask": "source_mask",
    "sslForwardProxy": "ssl_forward_proxy",
    "sslForwardProxyBypass": "ssl_forward_proxy_bypass",
    "sslSignHash": "ssl_sign_hash",
    "strictResume": "strict_resume",
    "strictUpdates": "strict_updates",
    "sysContact": "sys_contact",
    "sysLocation": "sys_location",
    "templateModified": "template_modified",
    "templatePrerequisiteErrors": "template_prerequisite_errors",
    "tenantFilter": "tenant_filter",
    "tenantList": "tenant_list",
    "tenantName": "tenant_name",
    "timeUntilUp": "time_until_up",
    "tmOptions": "tm_options",
    "tmPartition": "tm_partition",
    "trafficGroup": "traffic_group",
    "translateAddress": "translate_address",
    "translatePort": "translate_port",
    "uncleanShutdown": "unclean_shutdown",
    "untrustedCertResponseControl": "untrusted_cert_response_control",
    "uriExcludes": "uri_excludes",
    "uriIncludes": "uri_includes",
    "viaHostName": "via_host_name",
    "viaRequest": "via_request",
    "viaResponse": "via_response",
    "vlansEnabled": "vlans_enabled",
    "xffAlternativeNames": "xff_alternative_names",
}
