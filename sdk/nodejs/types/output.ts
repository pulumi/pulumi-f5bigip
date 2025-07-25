// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface As3DeleteApps {
    /**
     * List of application names to delete from the specified tenant.
     *
     * > `deleteApps` cannot be used together with `as3Json`.
     */
    apps: string[];
    /**
     * Name of the tenant containing the apps to delete.
     */
    tenantName: string;
}

export interface EventServiceDiscoveryNode {
    /**
     * name of node
     */
    id?: string;
    /**
     * ip of nonde
     */
    ip?: string;
    /**
     * port
     */
    port?: number;
}

export interface FastHttpAppMonitor {
    /**
     * Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
     */
    interval?: number;
    /**
     * set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
     */
    monitorAuth?: boolean;
    /**
     * password for web access on FAST-Generated Pool Monitor.
     */
    password?: string;
    /**
     * The presence of this string anywhere in the HTTP response implies availability.
     */
    response?: string;
    /**
     * Specify data to be sent during each health check for FAST-Generated Pool Monitor.
     */
    sendString?: string;
    /**
     * username for web access on FAST-Generated Pool Monitor.
     */
    username?: string;
}

export interface FastHttpAppPoolMember {
    /**
     * List of server address to be used for FAST-Generated Pool.
     */
    addresses: string[];
    /**
     * connectionLimit value to be used for FAST-Generated Pool.
     */
    connectionLimit: number;
    /**
     * port number of serviceport to be used for FAST-Generated Pool.
     */
    port?: number;
    /**
     * priorityGroup value to be used for FAST-Generated Pool.
     */
    priorityGroup: number;
    /**
     * shareNodes value to be used for FAST-Generated Pool.
     */
    shareNodes: boolean;
}

export interface FastHttpAppVirtualServer {
    /**
     * IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     */
    ip: string;
    /**
     * Port number to used for accessing virtual server/application
     */
    port: number;
}

export interface FastHttpAppWafSecurityPolicy {
    /**
     * Setting `true` will enable FAST to create WAF Security Policy.
     */
    enable: boolean;
}

export interface FastHttpsAppMonitor {
    /**
     * Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
     */
    interval?: number;
    /**
     * set `true` if the servers require login credentials for web access on FAST-Generated Pool Monitor. default is `false`.
     */
    monitorAuth?: boolean;
    /**
     * password for web access on FAST-Generated Pool Monitor.
     */
    password?: string;
    /**
     * The presence of this string anywhere in the HTTP response implies availability.
     */
    response?: string;
    /**
     * Specify data to be sent during each health check for FAST-Generated Pool Monitor.
     */
    sendString?: string;
    /**
     * username for web access on FAST-Generated Pool Monitor.
     */
    username?: string;
}

export interface FastHttpsAppPoolMember {
    /**
     * List of server address to be used for FAST-Generated Pool.
     */
    addresses: string[];
    /**
     * connectionLimit value to be used for FAST-Generated Pool.
     */
    connectionLimit: number;
    /**
     * port number of serviceport to be used for FAST-Generated Pool.
     */
    port?: number;
    /**
     * priorityGroup value to be used for FAST-Generated Pool.
     */
    priorityGroup: number;
    /**
     * shareNodes value to be used for FAST-Generated Pool.
     */
    shareNodes: boolean;
}

export interface FastHttpsAppTlsClientProfile {
    /**
     * Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.
     */
    tlsCertName: string;
    /**
     * Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile.
     */
    tlsKeyName: string;
}

export interface FastHttpsAppTlsServerProfile {
    /**
     * Name of existing BIG-IP SSL certificate to be used for FAST-Generated TLS Server Profile.
     */
    tlsCertName: string;
    /**
     * Name of existing BIG-IP SSL Key to be used for FAST-Generated TLS Server Profile.
     */
    tlsKeyName: string;
}

export interface FastHttpsAppVirtualServer {
    /**
     * IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     */
    ip: string;
    /**
     * Port number to used for accessing virtual server/application
     */
    port: number;
}

export interface FastHttpsAppWafSecurityPolicy {
    /**
     * Setting `true` will enable FAST to create WAF Security Policy.
     */
    enable: boolean;
}

export interface FastTcpAppMonitor {
    /**
     * Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
     */
    interval?: number;
}

export interface FastTcpAppPoolMember {
    /**
     * List of server address to be used for FAST-Generated Pool.
     */
    addresses: string[];
    /**
     * connectionLimit value to be used for FAST-Generated Pool.
     */
    connectionLimit?: number;
    /**
     * port number of serviceport to be used for FAST-Generated Pool.
     */
    port?: number;
    /**
     * priorityGroup value to be used for FAST-Generated Pool.
     */
    priorityGroup?: number;
    /**
     * shareNodes value to be used for FAST-Generated Pool.
     */
    shareNodes?: boolean;
}

export interface FastTcpAppVirtualServer {
    /**
     * IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     */
    ip: string;
    /**
     * Port number to used for accessing virtual server/application
     */
    port: number;
}

export interface FastUdpAppMonitor {
    /**
     * The presence of this optional string is required in the response, if specified it confirms availability.
     */
    expectedResponse?: string;
    /**
     * Set the time between health checks,in seconds for FAST-Generated Pool Monitor.
     */
    interval?: number;
    /**
     * Optional data to be sent during each health check.
     */
    sendString?: string;
}

export interface FastUdpAppPoolMember {
    /**
     * List of server address to be used for FAST-Generated Pool.
     */
    addresses: string[];
    /**
     * connectionLimit value to be used for FAST-Generated Pool.
     */
    connectionLimit?: number;
    /**
     * port number of serviceport to be used for FAST-Generated Pool.
     */
    port?: number;
    /**
     * priorityGroup value to be used for FAST-Generated Pool.
     */
    priorityGroup?: number;
    /**
     * shareNodes value to be used for FAST-Generated Pool.
     */
    shareNodes?: boolean;
}

export interface FastUdpAppVirtualServer {
    /**
     * IP4/IPv6 address to be used for virtual server ex: `10.1.1.1`
     */
    ip: string;
    /**
     * Port number to used for accessing virtual server/application
     */
    port: number;
}

export interface SaasBotDefenseProfileProtectedEndpoint {
    /**
     * Specifies the path to the web page to be protected by BD. For example, `/login`.
     */
    endpoint: string;
    /**
     * hostname or IP address of the web page to be protected by the Bot Defense
     */
    host: string;
    /**
     * Specifies whether the BIG-IP or F5 XC Bot Defense handles mitigation of malicious HTTP requests. This field is enabled only if the Service Level field is set to Advanced/Premium
     */
    mitigationAction: string;
    /**
     * Unique name for the protected endpoint
     */
    name: string;
    /**
     * POST field to protect the path when it has a POST method, `enabled` or `disabled`
     */
    post: string;
    /**
     * PUT field to protect the path when it has a PUT method,`enabled` or `disabled`
     */
    put: string;
}

export interface WafPolicyFileType {
    /**
     * Determines whether the file type is allowed or disallowed. In either of these cases the VIOL_FILETYPE violation is issued (if enabled) for an incoming request- 
     * * No allowed file type matched the file type of the request.
     * * The file type of the request matched a disallowed file type.
     */
    allowed?: boolean;
    /**
     * Specifies the file type name as appearing in the URL extension.
     */
    name?: string;
    /**
     * Determines the type of the name attribute. Only when setting the type to `wildcard` will the special wildcard characters in the name be interpreted as such
     */
    type?: string;
}

export interface WafPolicyGraphqlProfile {
    /**
     * Specifies when checked (enabled) that you want attack signatures and threat campaigns to be detected on this GraphQL profile and possibly override the security policy settings of an attack signature or threat campaign specifically for this GraphQL profile. After you enable this setting, the system displays a list of attack signatures and and threat campaigns. The default is enabled
     */
    attackSignaturesCheck?: boolean;
    /**
     * defense_attributes settings for policy
     */
    defenseAttributes?: outputs.WafPolicyGraphqlProfileDefenseAttribute[];
    /**
     * Specifies when checked (enabled) that the system enforces the security policy settings of a meta character for the GraphQL profile. After you enable this setting, the system displays a list of meta characters. The default is enabled
     */
    metacharElementcheck?: boolean;
    /**
     * The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     */
    name: string;
}

export interface WafPolicyGraphqlProfileDefenseAttribute {
    /**
     * Introspection queries can also be enforced to prevent attackers from using them to
     * understand the API structure and potentially breach an application.
     */
    allowIntrospectionQueries?: boolean;
    /**
     * Specifies the highest number of batched queries allowed by the security policy.
     */
    maximumBatchedQueries?: string;
    /**
     * Specifies the greatest nesting depth found in the GraphQL structure allowed by the security policy.
     */
    maximumStructureDepth?: string;
    /**
     * Specifies the longest length, in bytes, allowed by the security policy of the request payload, or parameter value, where the GraphQL data was found.
     */
    maximumTotalLength?: string;
    /**
     * Specifies the longest length (in bytes) of the longest GraphQL element value in the document allowed by the security policy.
     */
    maximumValueLength?: string;
    /**
     * Specifies, when checked (enabled), that the system does not report when the security enforcer encounters warnings while parsing GraphQL content. Specifies when cleared (disabled), that the security policy reports when the security enforcer encounters warnings while parsing GraphQL content. The default setting is disabled.
     */
    tolerateParsingWarnings?: boolean;
}

export interface WafPolicyHostName {
    /**
     * The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     */
    name?: string;
}

export interface WafPolicyIpException {
    /**
     * Specifies how the system responds to blocking requests sent from this IP address. Possible options [`always`, `never`, `policy-default`].
     */
    blockRequests?: string;
    /**
     * Specifies the description of the policy.
     */
    description?: string;
    /**
     * Specifies when enabled that the system considers this IP address legitimate and does not take it into account when performing brute force prevention.
     */
    ignoreAnomalies?: boolean;
    /**
     * Specifies when enabled that the system considers this IP address legitimate even if it is found in the IP Intelligence database (a database of questionable IP addresses).
     */
    ignoreIpreputation?: boolean;
    /**
     * Specifies the IP address that you want the system to trust.
     */
    ipAddress: string;
    /**
     * Specifies the netmask of the exceptional IP address. This is an optional field.
     */
    ipMask: string;
    /**
     * Specifies when enabled the Policy Builder considers traffic from this IP address as being safe.
     */
    trustedbyPolicybuilder?: boolean;
}

export interface WafPolicyPolicyBuilder {
    /**
     * learning mode setting for policy-builder, possible options: [`automatic`,`disabled`, `manual`]
     */
    learningMode?: string;
}

export interface WafPolicySignaturesSetting {
    placesignaturesInStaging?: boolean;
    /**
     * setting true will enforce all signature from staging
     */
    signatureStaging?: boolean;
}

export namespace cm {
    export interface DeviceGroupDevice {
        /**
         * Is the name of the device Group
         */
        name?: string;
        /**
         * Name of origin
         */
        setSyncLeader?: boolean;
    }

}

export namespace ltm {
    export interface DataGroupRecord {
        /**
         * , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
         */
        data?: string;
        /**
         * , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
         */
        name: string;
    }

    export interface GetDataGroupRecord {
        data?: string;
        /**
         * Name of the datagroup
         */
        name: string;
    }

    export interface GetNodeFqdn {
        /**
         * The FQDN node's address family.
         */
        addressFamily?: string;
        /**
         * Specifies if the node should scale to the IP address set returned by DNS.
         */
        autopopulate: string;
        /**
         * The number of attempts to resolve a domain name.
         */
        downinterval: number;
        /**
         * The amount of time before sending the next DNS query.
         */
        interval: string;
        /**
         * Name of the node.
         */
        name?: string;
    }

    export interface GetPolicyRule {
        actions?: outputs.ltm.GetPolicyRuleAction[];
        conditions?: outputs.ltm.GetPolicyRuleCondition[];
        /**
         * Name of the policy which includes partion ( /partition/policy-name )
         */
        name: string;
    }

    export interface GetPolicyRuleAction {
        appService: string;
        application: string;
        asm: boolean;
        avr: boolean;
        cache: boolean;
        carp: boolean;
        category: string;
        classify: boolean;
        clonePool: string;
        code: number;
        compress: boolean;
        connection: boolean;
        content: string;
        cookieHash: boolean;
        cookieInsert: boolean;
        cookiePassive: boolean;
        cookieRewrite: boolean;
        decompress: boolean;
        defer: boolean;
        destinationAddress: boolean;
        disable: boolean;
        domain: string;
        enable: boolean;
        expiry: string;
        expirySecs: number;
        expression: string;
        extension: string;
        facility: string;
        forward?: boolean;
        fromProfile: string;
        hash: boolean;
        host: string;
        http: boolean;
        httpBasicAuth: boolean;
        httpCookie: boolean;
        httpHeader: boolean;
        httpHost?: boolean;
        httpReferer: boolean;
        httpReply: boolean;
        httpSetCookie: boolean;
        httpUri: boolean;
        ifile: string;
        insert: boolean;
        internalVirtual: string;
        ipAddress: string;
        key: string;
        l7dos: boolean;
        length: number;
        location: string;
        log: boolean;
        ltmPolicy: boolean;
        member: string;
        message: string;
        netmask: string;
        nexthop: string;
        node: string;
        offset: number;
        path: string;
        pem: boolean;
        persist: boolean;
        pin: boolean;
        policy: string;
        pool: string;
        port: number;
        priority: string;
        profile: string;
        protocol: string;
        queryString: string;
        rateclass: string;
        redirect: boolean;
        remove: boolean;
        replace: boolean;
        request: boolean;
        requestAdapt: boolean;
        reset: boolean;
        response: boolean;
        responseAdapt: boolean;
        scheme: string;
        script: string;
        select: boolean;
        serverSsl: boolean;
        setVariable: boolean;
        shutdown: boolean;
        snat: string;
        snatpool: string;
        sourceAddress: boolean;
        sslClientHello: boolean;
        sslServerHandshake: boolean;
        sslServerHello: boolean;
        sslSessionId: boolean;
        status: number;
        tcl: boolean;
        tcpNagle: boolean;
        text: string;
        timeout: number;
        tmName: string;
        uie: boolean;
        universal: boolean;
        value: string;
        virtual: string;
        vlan: string;
        vlanId: number;
        wam: boolean;
        write: boolean;
    }

    export interface GetPolicyRuleCondition {
        address: boolean;
        all: boolean;
        appService: string;
        browserType: boolean;
        browserVersion: boolean;
        caseInsensitive: boolean;
        caseSensitive: boolean;
        cipher: boolean;
        cipherBits: boolean;
        clientSsl: boolean;
        code: boolean;
        commonName: boolean;
        contains: boolean;
        continent: boolean;
        countryCode: boolean;
        countryName: boolean;
        cpuUsage: boolean;
        datagroup: string;
        deviceMake: boolean;
        deviceModel: boolean;
        domain: boolean;
        endsWith: boolean;
        equals: boolean;
        expiry: boolean;
        extension: boolean;
        external: boolean;
        geoip: boolean;
        greater: boolean;
        greaterOrEqual: boolean;
        host: boolean;
        httpBasicAuth: boolean;
        httpCookie: boolean;
        httpHeader: boolean;
        httpHost: boolean;
        httpMethod: boolean;
        httpReferer: boolean;
        httpSetCookie: boolean;
        httpStatus: boolean;
        httpUri: boolean;
        httpUserAgent: boolean;
        httpVersion: boolean;
        index: number;
        internal: boolean;
        isp: boolean;
        last15secs: boolean;
        last1min: boolean;
        last5mins: boolean;
        less: boolean;
        lessOrEqual: boolean;
        local: boolean;
        major: boolean;
        matches: boolean;
        minor: boolean;
        missing: boolean;
        mss: boolean;
        not: boolean;
        org: boolean;
        password: boolean;
        path: boolean;
        pathSegment: boolean;
        port: boolean;
        present: boolean;
        protocol: boolean;
        queryParameter: boolean;
        queryString: boolean;
        regionCode: boolean;
        regionName: boolean;
        remote: boolean;
        request: boolean;
        response: boolean;
        routeDomain: boolean;
        rtt: boolean;
        scheme: boolean;
        serverName: boolean;
        sslCert: boolean;
        sslClientHello: boolean;
        sslExtension: boolean;
        sslServerHandshake: boolean;
        sslServerHello: boolean;
        startsWith: boolean;
        tcp: boolean;
        text: boolean;
        tmName: string;
        unnamedQueryParameter: boolean;
        userAgentToken: boolean;
        username: boolean;
        value: boolean;
        values: string[];
        version: boolean;
        vlan: boolean;
        vlanId: boolean;
    }

    export interface NodeFqdn {
        /**
         * Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
         */
        addressFamily: string;
        /**
         * Specifies whether the node should scale to the IP address set returned by DNS.
         */
        autopopulate: string;
        /**
         * Specifies the number of attempts to resolve a domain name. The default is 5.
         */
        downinterval: number;
        /**
         * Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
         */
        interval: string;
        /**
         * Name of the node
         */
        name?: string;
    }

    export interface PolicyRule {
        /**
         * Block type. See action block for more details.
         */
        actions?: outputs.ltm.PolicyRuleAction[];
        /**
         * Block type. See condition block for more details.
         */
        conditions?: outputs.ltm.PolicyRuleCondition[];
        /**
         * Specifies descriptive text that identifies the irule attached to policy.
         */
        description?: string;
        /**
         * Name of Rule to be applied in policy.
         */
        name: string;
    }

    export interface PolicyRuleAction {
        appService: string;
        application: string;
        asm: boolean;
        avr: boolean;
        cache: boolean;
        carp: boolean;
        category: string;
        classify: boolean;
        clonePool: string;
        code: number;
        compress: boolean;
        /**
         * This action is set to `true` by default, it needs to be explicitly set to `false` for actions it conflicts with.
         */
        connection?: boolean;
        content: string;
        cookieHash: boolean;
        cookieInsert: boolean;
        cookiePassive: boolean;
        cookieRewrite: boolean;
        decompress: boolean;
        defer: boolean;
        destinationAddress: boolean;
        disable: boolean;
        domain: string;
        enable: boolean;
        expiry: string;
        expirySecs: number;
        expression: string;
        extension: string;
        facility: string;
        /**
         * This action will affect forwarding.
         */
        forward?: boolean;
        fromProfile: string;
        hash: boolean;
        host: string;
        http: boolean;
        httpBasicAuth: boolean;
        httpCookie: boolean;
        httpHeader: boolean;
        httpHost?: boolean;
        httpReferer: boolean;
        httpReply: boolean;
        httpSetCookie: boolean;
        httpUri: boolean;
        ifile: string;
        insert: boolean;
        internalVirtual: string;
        ipAddress: string;
        key: string;
        l7dos: boolean;
        length: number;
        location: string;
        log: boolean;
        ltmPolicy: boolean;
        member: string;
        message: string;
        netmask: string;
        nexthop: string;
        node: string;
        offset: number;
        path: string;
        pem: boolean;
        persist: boolean;
        pin: boolean;
        policy: string;
        /**
         * This action will direct the stream to this pool.
         */
        pool: string;
        port: number;
        priority: string;
        profile: string;
        protocol: string;
        queryString: string;
        rateclass: string;
        redirect: boolean;
        remove: boolean;
        replace: boolean;
        request: boolean;
        requestAdapt: boolean;
        reset: boolean;
        response: boolean;
        responseAdapt: boolean;
        scheme: string;
        script: string;
        select: boolean;
        serverSsl: boolean;
        setVariable: boolean;
        shutdown?: boolean;
        snat: string;
        snatpool: string;
        sourceAddress: boolean;
        sslClientHello: boolean;
        sslServerHandshake: boolean;
        sslServerHello: boolean;
        sslSessionId: boolean;
        status: number;
        tcl: boolean;
        tcpNagle: boolean;
        text: string;
        timeout: number;
        tmName: string;
        uie: boolean;
        universal: boolean;
        value: string;
        virtual: string;
        vlan: string;
        vlanId: number;
        wam: boolean;
        write: boolean;
    }

    export interface PolicyRuleCondition {
        address: boolean;
        all: boolean;
        appService: string;
        browserType: boolean;
        browserVersion: boolean;
        caseInsensitive: boolean;
        caseSensitive: boolean;
        cipher: boolean;
        cipherBits: boolean;
        clientAccepted: boolean;
        clientSsl: boolean;
        code: boolean;
        commonName: boolean;
        contains: boolean;
        continent: boolean;
        countryCode: boolean;
        countryName: boolean;
        cpuUsage: boolean;
        datagroup?: string;
        deviceMake: boolean;
        deviceModel: boolean;
        domain: boolean;
        endsWith: boolean;
        equals: boolean;
        exists: boolean;
        expiry: boolean;
        extension: boolean;
        external: boolean;
        geoip: boolean;
        greater: boolean;
        greaterOrEqual: boolean;
        host: boolean;
        httpBasicAuth: boolean;
        httpCookie: boolean;
        httpHeader: boolean;
        httpHost: boolean;
        httpMethod: boolean;
        httpReferer: boolean;
        httpSetCookie: boolean;
        httpStatus: boolean;
        httpUri: boolean;
        httpUserAgent: boolean;
        httpVersion: boolean;
        index: number;
        internal: boolean;
        isp: boolean;
        last15secs: boolean;
        last1min: boolean;
        last5mins: boolean;
        less: boolean;
        lessOrEqual: boolean;
        local: boolean;
        major: boolean;
        matches: boolean;
        minor: boolean;
        missing: boolean;
        mss: boolean;
        not: boolean;
        org: boolean;
        password: boolean;
        path: boolean;
        pathSegment: boolean;
        port: boolean;
        present: boolean;
        protocol: boolean;
        queryParameter: boolean;
        queryString: boolean;
        regionCode: boolean;
        regionName: boolean;
        remote: boolean;
        request: boolean;
        response: boolean;
        routeDomain: boolean;
        rtt: boolean;
        scheme: boolean;
        serverName: boolean;
        sslCert: boolean;
        sslClientHello: boolean;
        sslExtension: boolean;
        sslServerHandshake: boolean;
        sslServerHello: boolean;
        startsWith: boolean;
        tcp: boolean;
        text: boolean;
        tmName: string;
        unnamedQueryParameter: boolean;
        userAgentToken: boolean;
        username: boolean;
        value: boolean;
        values: string[];
        version: boolean;
        vlan: boolean;
        vlanId: boolean;
    }

    export interface ProfileClientSslCertKeyChain {
        /**
         * Specifies the name of the certificate that the system uses for client-side SSL processing. The default is `default`
         */
        cert?: string;
        /**
         * Specifies a certificate chain file that a server can use for authentication. The default is `None`.
         */
        chain?: string;
        /**
         * Specifies the file name of the SSL key. The default is `default`
         */
        key?: string;
        /**
         * Name of Cert-key-chain
         */
        name?: string;
        /**
         * Type the name of the pass phrase used to encrypt the key.
         */
        passphrase: string;
    }

    export interface ProfileHttpEnforcement {
        /**
         * Specifies which HTTP methods count as being known. Removing RFC-defined methods from this list will cause the HTTP filter to not recognize them. Default value is [CONNECT DELETE GET HEAD LOCK OPTIONS POST PROPFIND PUT TRACE UNLOCK].If no value is specified while creating, then default value will be assigned by BigIP. In order to remove it, [""] list is to be passed. If knownMethods is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value , we need to pass [CONNECT DELETE GET HEAD LOCK OPTIONS POST PROPFIND PUT TRACE UNLOCK] explicitly.
         */
        knownMethods: string[];
        /**
         * Specifies the maximum number of headers allowed in HTTP request/response. The default is 64 headers.If no value is specified while creating, then default value will be assigned by BigIP. If maxHeaderCount is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "64" explicitly.
         */
        maxHeaderCount: number;
        /**
         * Specifies the maximum header size. The default value is 32768. If no string is specified while creating, then default value will be assigned by BigIP. If maxHeaderSize is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "32768" explicitly.
         */
        maxHeaderSize: number;
        /**
         * Specifies whether to allow, reject or switch to pass-through mode when an unknown HTTP method is parsed. Default value is "allow". If no string is specified while creating, then default value will be assigned by BigIP. If unknownMethod is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "allow" explicitly.
         */
        unknownMethod: string;
    }

    export interface ProfileHttpHttpStrictTransportSecurity {
        /**
         * The Include Subdomains setting applies the HSTS policy to the HSTS host and its subdomains. The default is "enabled". If no string is specified during Create, then default value will be assigned by BigIp. If includeSubdomains is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "enabled" explicitly.
         */
        includeSubdomains: string;
        /**
         * The Maximum Age value specifies the length of time, in seconds, that HSTS functionality requests that clients only use HTTPS to connect to the current host and any subdomains of the current host's domain name.  The default is 16070400 seconds. If no value is specified during Create, then default value will be assigned by BigIp. If maximumAge is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value , we need to pass 16070400 explicitly.
         */
        maximumAge: number;
        /**
         * The Mode setting enables and disables HSTS functionality within the HTTP profile. The default is "disabled". If no string is specified during Create, then default value will be assigned by BigIp. If mode is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "disabled" explicitly.
         */
        mode: string;
        /**
         * An HSTS preload list is a list of domains built into a web browser. When you enable the Preload setting, the domain for the web site that this HTTP profile is associated with is submitted for inclusion in the browser's preload list. The default is "disabled". If no string is specified during Create, then default value will be assigned by BigIp. If preload is commented (or not passed) during the update call, then no changes would be applied and previous value will persist. In order to put default value, we need to pass "disabled" explicitly.
         */
        preload: string;
    }

    export interface ProfileRewriteCookieRule {
        clientDomain: string;
        clientPath: string;
        /**
         * Name of the cookie rewrite rule.
         */
        ruleName: string;
        serverDomain: string;
        serverPath: string;
    }

    export interface ProfileRewriteRequest {
        /**
         * Enable to add the X-Forwarded For (XFF) header, to specify the originating IP address of the client. Valid choices are: `enabled, disabled`
         */
        insertXfwdFor?: string;
        /**
         * Enable to add the X-Forwarded Host header, to specify the originating host of the client. Valid choices are: `enabled, disabled`
         */
        insertXfwdHost?: string;
        /**
         * Enable to add the X-Forwarded Proto header, to specify the originating protocol of the client. Valid choices are: `enabled, disabled`
         */
        insertXfwdProtocol?: string;
        /**
         * Enable to rewrite headers in Request settings. Valid choices are: `enabled, disabled`
         */
        rewriteHeaders?: string;
    }

    export interface ProfileRewriteResponse {
        /**
         * Enable to rewrite links in content in the response. Valid choices are: `enabled, disabled`
         */
        rewriteContent?: string;
        /**
         * Enable to rewrite headers in the response. Valid choices are: `enabled, disabled`
         */
        rewriteHeaders?: string;
    }

    export interface ProfileRewriteUriRulesClient {
        /**
         * Host part of the uri, e.g. `www.foo.com`.
         */
        host: string;
        /**
         * Path part of the uri, must always end with `/`. Default value is: `/`
         */
        path?: string;
        /**
         * Port part of the uri. Default value is: `none`
         */
        port?: string;
        /**
         * Scheme part of the uri, e.g. `https`, `ftp`.
         */
        scheme: string;
    }

    export interface ProfileRewriteUriRulesServer {
        /**
         * Host part of the uri, e.g. `www.foo.com`.
         */
        host: string;
        /**
         * Path part of the uri, must always end with `/`. Default value is: `/`
         */
        path?: string;
        /**
         * Port part of the uri. Default value is: `none`
         */
        port?: string;
        /**
         * Scheme part of the uri, e.g. `https`, `ftp`.
         */
        scheme: string;
    }

    export interface SnatOrigin {
        /**
         * app service
         */
        appService?: string;
        /**
         * Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
         */
        name?: string;
    }

}

export namespace net {
    export interface VlanInterface {
        /**
         * Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
         */
        tagged?: boolean;
        /**
         * Physical or virtual port used for traffic
         */
        vlanport?: string;
    }

}

export namespace ssl {
    export interface GetWafEntityParameterUrl {
        method: string;
        name: string;
        protocol: string;
        type: string;
    }

    export interface GetWafEntityUrlCrossOriginRequestsEnforcement {
        /**
         * Determines whether the subdomains are allowed to receive data from the web application.
         */
        includeSubdomains?: boolean;
        /**
         * Specifies the name of the origin with which you want to share your data.
         */
        originName: string;
        /**
         * Specifies the port that other web applications are allowed to use to request data from your web application.
         */
        originPort: string;
        /**
         * Specifies the protocol that other web applications are allowed to use to request data from your web application.
         */
        originProtocol: string;
    }

    export interface GetWafEntityUrlMethodOverride {
        /**
         * Specifies that the system allows or disallows a method for this URL
         */
        allow: boolean;
        /**
         * Specifies an HTTP method.
         */
        method: string;
    }

}

export namespace sys {
    export interface IAppList {
        /**
         * Name of origin
         */
        encrypted?: string;
        /**
         * Name of origin
         */
        value?: string;
    }

    export interface IAppMetadata {
        /**
         * Name of origin
         */
        persists?: string;
        /**
         * Name of origin
         */
        value?: string;
    }

    export interface IAppTable {
        columnNames?: string[];
        /**
         * Name of origin
         */
        encryptedColumns?: string;
        /**
         * Name of the iApp.
         */
        name?: string;
        rows?: outputs.sys.IAppTableRow[];
    }

    export interface IAppTableRow {
        rows?: string[];
    }

    export interface IAppVariable {
        /**
         * Name of origin
         */
        encrypted?: string;
        /**
         * Name of the iApp.
         */
        name?: string;
        /**
         * Name of origin
         */
        value?: string;
    }

}
