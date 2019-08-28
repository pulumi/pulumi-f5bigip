// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";

export namespace cm {
    export interface DeviceGroupDevice {
        /**
         * Is the name of the device Group
         */
        name?: string;
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

    export interface NodeFqdn {
        /**
         * Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
         */
        addressFamily?: string;
        autopopulate: string;
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
        actions?: outputs.ltm.PolicyRuleAction[];
        conditions?: outputs.ltm.PolicyRuleCondition[];
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
        forward: boolean;
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
        /**
         * If Rule is used then you need to provide the tmName it can be any value
         */
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
        clientSsl: boolean;
        code: boolean;
        commonName: boolean;
        contains: boolean;
        continent: boolean;
        countryCode: boolean;
        countryName: boolean;
        cpuUsage: boolean;
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
        /**
         * If Rule is used then you need to provide the tmName it can be any value
         */
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

    export interface SnatOrigin {
        appService?: string;
        /**
         * Name of the snat
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

export namespace sys {
    export interface IAppList {
        encrypted?: string;
        value?: string;
    }

    export interface IAppMetadata {
        persists?: string;
        value?: string;
    }

    export interface IAppTable {
        columnNames?: string[];
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
        encrypted?: string;
        /**
         * Name of the iApp.
         */
        name?: string;
        value?: string;
    }
}
