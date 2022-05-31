// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface EventServiceDiscoveryNode {
    id?: pulumi.Input<string>;
    ip?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
}
export namespace cm {
    export interface DeviceGroupDevice {
        /**
         * Is the name of the device Group
         */
        name?: pulumi.Input<string>;
        setSyncLeader?: pulumi.Input<boolean>;
    }
}

export namespace ltm {
    export interface DataGroupRecord {
        /**
         * , sets the value of the record's `data` attribute, specifying a value here will create a record in the form of `name := data`
         */
        data?: pulumi.Input<string>;
        /**
         * , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
         */
        name: pulumi.Input<string>;
    }

    export interface GetDataGroupRecordArgs {
        data?: pulumi.Input<string>;
        /**
         * Name of the datagroup
         */
        name: pulumi.Input<string>;
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
        autopopulate?: string;
        /**
         * The number of attempts to resolve a domain name.
         */
        downinterval?: number;
        /**
         * The amount of time before sending the next DNS query.
         */
        interval?: string;
        /**
         * Name of the node.
         */
        name?: string;
    }

    export interface GetNodeFqdnArgs {
        /**
         * The FQDN node's address family.
         */
        addressFamily?: pulumi.Input<string>;
        /**
         * Specifies if the node should scale to the IP address set returned by DNS.
         */
        autopopulate?: pulumi.Input<string>;
        /**
         * The number of attempts to resolve a domain name.
         */
        downinterval?: pulumi.Input<number>;
        /**
         * The amount of time before sending the next DNS query.
         */
        interval?: pulumi.Input<string>;
        /**
         * Name of the node.
         */
        name?: pulumi.Input<string>;
    }

    export interface GetPolicyRuleArgs {
        actions?: pulumi.Input<pulumi.Input<inputs.ltm.GetPolicyRuleActionArgs>[]>;
        conditions?: pulumi.Input<pulumi.Input<inputs.ltm.GetPolicyRuleConditionArgs>[]>;
        /**
         * Name of the policy which includes partion ( /partition/policy-name )
         */
        name: pulumi.Input<string>;
    }

    export interface GetPolicyRule {
        actions?: inputs.ltm.GetPolicyRuleAction[];
        conditions?: inputs.ltm.GetPolicyRuleCondition[];
        /**
         * Name of the policy which includes partion ( /partition/policy-name )
         */
        name: string;
    }

    export interface GetPolicyRuleActionArgs {
        appService?: pulumi.Input<string>;
        application?: pulumi.Input<string>;
        asm?: pulumi.Input<boolean>;
        avr?: pulumi.Input<boolean>;
        cache?: pulumi.Input<boolean>;
        carp?: pulumi.Input<boolean>;
        category?: pulumi.Input<string>;
        classify?: pulumi.Input<boolean>;
        clonePool?: pulumi.Input<string>;
        code?: pulumi.Input<number>;
        compress?: pulumi.Input<boolean>;
        connection?: pulumi.Input<boolean>;
        content?: pulumi.Input<string>;
        cookieHash?: pulumi.Input<boolean>;
        cookieInsert?: pulumi.Input<boolean>;
        cookiePassive?: pulumi.Input<boolean>;
        cookieRewrite?: pulumi.Input<boolean>;
        decompress?: pulumi.Input<boolean>;
        defer?: pulumi.Input<boolean>;
        destinationAddress?: pulumi.Input<boolean>;
        disable?: pulumi.Input<boolean>;
        domain?: pulumi.Input<string>;
        enable?: pulumi.Input<boolean>;
        expiry?: pulumi.Input<string>;
        expirySecs?: pulumi.Input<number>;
        expression?: pulumi.Input<string>;
        extension?: pulumi.Input<string>;
        facility?: pulumi.Input<string>;
        forward?: pulumi.Input<boolean>;
        fromProfile?: pulumi.Input<string>;
        hash?: pulumi.Input<boolean>;
        host?: pulumi.Input<string>;
        http?: pulumi.Input<boolean>;
        httpBasicAuth?: pulumi.Input<boolean>;
        httpCookie?: pulumi.Input<boolean>;
        httpHeader?: pulumi.Input<boolean>;
        httpHost?: pulumi.Input<boolean>;
        httpReferer?: pulumi.Input<boolean>;
        httpReply?: pulumi.Input<boolean>;
        httpSetCookie?: pulumi.Input<boolean>;
        httpUri?: pulumi.Input<boolean>;
        ifile?: pulumi.Input<string>;
        insert?: pulumi.Input<boolean>;
        internalVirtual?: pulumi.Input<string>;
        ipAddress?: pulumi.Input<string>;
        key?: pulumi.Input<string>;
        l7dos?: pulumi.Input<boolean>;
        length?: pulumi.Input<number>;
        location?: pulumi.Input<string>;
        log?: pulumi.Input<boolean>;
        ltmPolicy?: pulumi.Input<boolean>;
        member?: pulumi.Input<string>;
        message?: pulumi.Input<string>;
        netmask?: pulumi.Input<string>;
        nexthop?: pulumi.Input<string>;
        node?: pulumi.Input<string>;
        offset?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
        pem?: pulumi.Input<boolean>;
        persist?: pulumi.Input<boolean>;
        pin?: pulumi.Input<boolean>;
        policy?: pulumi.Input<string>;
        pool?: pulumi.Input<string>;
        port?: pulumi.Input<number>;
        priority?: pulumi.Input<string>;
        profile?: pulumi.Input<string>;
        protocol?: pulumi.Input<string>;
        queryString?: pulumi.Input<string>;
        rateclass?: pulumi.Input<string>;
        redirect?: pulumi.Input<boolean>;
        remove?: pulumi.Input<boolean>;
        replace?: pulumi.Input<boolean>;
        request?: pulumi.Input<boolean>;
        requestAdapt?: pulumi.Input<boolean>;
        reset?: pulumi.Input<boolean>;
        response?: pulumi.Input<boolean>;
        responseAdapt?: pulumi.Input<boolean>;
        scheme?: pulumi.Input<string>;
        script?: pulumi.Input<string>;
        select?: pulumi.Input<boolean>;
        serverSsl?: pulumi.Input<boolean>;
        setVariable?: pulumi.Input<boolean>;
        shutdown?: pulumi.Input<boolean>;
        snat?: pulumi.Input<string>;
        snatpool?: pulumi.Input<string>;
        sourceAddress?: pulumi.Input<boolean>;
        sslClientHello?: pulumi.Input<boolean>;
        sslServerHandshake?: pulumi.Input<boolean>;
        sslServerHello?: pulumi.Input<boolean>;
        sslSessionId?: pulumi.Input<boolean>;
        status?: pulumi.Input<number>;
        tcl?: pulumi.Input<boolean>;
        tcpNagle?: pulumi.Input<boolean>;
        text?: pulumi.Input<string>;
        timeout?: pulumi.Input<number>;
        tmName?: pulumi.Input<string>;
        uie?: pulumi.Input<boolean>;
        universal?: pulumi.Input<boolean>;
        value?: pulumi.Input<string>;
        virtual?: pulumi.Input<string>;
        vlan?: pulumi.Input<string>;
        vlanId?: pulumi.Input<number>;
        wam?: pulumi.Input<boolean>;
        write?: pulumi.Input<boolean>;
    }

    export interface GetPolicyRuleAction {
        appService?: string;
        application?: string;
        asm?: boolean;
        avr?: boolean;
        cache?: boolean;
        carp?: boolean;
        category?: string;
        classify?: boolean;
        clonePool?: string;
        code?: number;
        compress?: boolean;
        connection?: boolean;
        content?: string;
        cookieHash?: boolean;
        cookieInsert?: boolean;
        cookiePassive?: boolean;
        cookieRewrite?: boolean;
        decompress?: boolean;
        defer?: boolean;
        destinationAddress?: boolean;
        disable?: boolean;
        domain?: string;
        enable?: boolean;
        expiry?: string;
        expirySecs?: number;
        expression?: string;
        extension?: string;
        facility?: string;
        forward?: boolean;
        fromProfile?: string;
        hash?: boolean;
        host?: string;
        http?: boolean;
        httpBasicAuth?: boolean;
        httpCookie?: boolean;
        httpHeader?: boolean;
        httpHost?: boolean;
        httpReferer?: boolean;
        httpReply?: boolean;
        httpSetCookie?: boolean;
        httpUri?: boolean;
        ifile?: string;
        insert?: boolean;
        internalVirtual?: string;
        ipAddress?: string;
        key?: string;
        l7dos?: boolean;
        length?: number;
        location?: string;
        log?: boolean;
        ltmPolicy?: boolean;
        member?: string;
        message?: string;
        netmask?: string;
        nexthop?: string;
        node?: string;
        offset?: number;
        path?: string;
        pem?: boolean;
        persist?: boolean;
        pin?: boolean;
        policy?: string;
        pool?: string;
        port?: number;
        priority?: string;
        profile?: string;
        protocol?: string;
        queryString?: string;
        rateclass?: string;
        redirect?: boolean;
        remove?: boolean;
        replace?: boolean;
        request?: boolean;
        requestAdapt?: boolean;
        reset?: boolean;
        response?: boolean;
        responseAdapt?: boolean;
        scheme?: string;
        script?: string;
        select?: boolean;
        serverSsl?: boolean;
        setVariable?: boolean;
        shutdown?: boolean;
        snat?: string;
        snatpool?: string;
        sourceAddress?: boolean;
        sslClientHello?: boolean;
        sslServerHandshake?: boolean;
        sslServerHello?: boolean;
        sslSessionId?: boolean;
        status?: number;
        tcl?: boolean;
        tcpNagle?: boolean;
        text?: string;
        timeout?: number;
        tmName?: string;
        uie?: boolean;
        universal?: boolean;
        value?: string;
        virtual?: string;
        vlan?: string;
        vlanId?: number;
        wam?: boolean;
        write?: boolean;
    }

    export interface GetPolicyRuleCondition {
        address?: boolean;
        all?: boolean;
        appService?: string;
        browserType?: boolean;
        browserVersion?: boolean;
        caseInsensitive?: boolean;
        caseSensitive?: boolean;
        cipher?: boolean;
        cipherBits?: boolean;
        clientSsl?: boolean;
        code?: boolean;
        commonName?: boolean;
        contains?: boolean;
        continent?: boolean;
        countryCode?: boolean;
        countryName?: boolean;
        cpuUsage?: boolean;
        deviceMake?: boolean;
        deviceModel?: boolean;
        domain?: boolean;
        endsWith?: boolean;
        equals?: boolean;
        expiry?: boolean;
        extension?: boolean;
        external?: boolean;
        geoip?: boolean;
        greater?: boolean;
        greaterOrEqual?: boolean;
        host?: boolean;
        httpBasicAuth?: boolean;
        httpCookie?: boolean;
        httpHeader?: boolean;
        httpHost?: boolean;
        httpMethod?: boolean;
        httpReferer?: boolean;
        httpSetCookie?: boolean;
        httpStatus?: boolean;
        httpUri?: boolean;
        httpUserAgent?: boolean;
        httpVersion?: boolean;
        index?: number;
        internal?: boolean;
        isp?: boolean;
        last15secs?: boolean;
        last1min?: boolean;
        last5mins?: boolean;
        less?: boolean;
        lessOrEqual?: boolean;
        local?: boolean;
        major?: boolean;
        matches?: boolean;
        minor?: boolean;
        missing?: boolean;
        mss?: boolean;
        not?: boolean;
        org?: boolean;
        password?: boolean;
        path?: boolean;
        pathSegment?: boolean;
        port?: boolean;
        present?: boolean;
        protocol?: boolean;
        queryParameter?: boolean;
        queryString?: boolean;
        regionCode?: boolean;
        regionName?: boolean;
        remote?: boolean;
        request?: boolean;
        response?: boolean;
        routeDomain?: boolean;
        rtt?: boolean;
        scheme?: boolean;
        serverName?: boolean;
        sslCert?: boolean;
        sslClientHello?: boolean;
        sslExtension?: boolean;
        sslServerHandshake?: boolean;
        sslServerHello?: boolean;
        startsWith?: boolean;
        tcp?: boolean;
        text?: boolean;
        tmName?: string;
        unnamedQueryParameter?: boolean;
        userAgentToken?: boolean;
        username?: boolean;
        value?: boolean;
        values?: string[];
        version?: boolean;
        vlan?: boolean;
        vlanId?: boolean;
    }

    export interface GetPolicyRuleConditionArgs {
        address?: pulumi.Input<boolean>;
        all?: pulumi.Input<boolean>;
        appService?: pulumi.Input<string>;
        browserType?: pulumi.Input<boolean>;
        browserVersion?: pulumi.Input<boolean>;
        caseInsensitive?: pulumi.Input<boolean>;
        caseSensitive?: pulumi.Input<boolean>;
        cipher?: pulumi.Input<boolean>;
        cipherBits?: pulumi.Input<boolean>;
        clientSsl?: pulumi.Input<boolean>;
        code?: pulumi.Input<boolean>;
        commonName?: pulumi.Input<boolean>;
        contains?: pulumi.Input<boolean>;
        continent?: pulumi.Input<boolean>;
        countryCode?: pulumi.Input<boolean>;
        countryName?: pulumi.Input<boolean>;
        cpuUsage?: pulumi.Input<boolean>;
        deviceMake?: pulumi.Input<boolean>;
        deviceModel?: pulumi.Input<boolean>;
        domain?: pulumi.Input<boolean>;
        endsWith?: pulumi.Input<boolean>;
        equals?: pulumi.Input<boolean>;
        expiry?: pulumi.Input<boolean>;
        extension?: pulumi.Input<boolean>;
        external?: pulumi.Input<boolean>;
        geoip?: pulumi.Input<boolean>;
        greater?: pulumi.Input<boolean>;
        greaterOrEqual?: pulumi.Input<boolean>;
        host?: pulumi.Input<boolean>;
        httpBasicAuth?: pulumi.Input<boolean>;
        httpCookie?: pulumi.Input<boolean>;
        httpHeader?: pulumi.Input<boolean>;
        httpHost?: pulumi.Input<boolean>;
        httpMethod?: pulumi.Input<boolean>;
        httpReferer?: pulumi.Input<boolean>;
        httpSetCookie?: pulumi.Input<boolean>;
        httpStatus?: pulumi.Input<boolean>;
        httpUri?: pulumi.Input<boolean>;
        httpUserAgent?: pulumi.Input<boolean>;
        httpVersion?: pulumi.Input<boolean>;
        index?: pulumi.Input<number>;
        internal?: pulumi.Input<boolean>;
        isp?: pulumi.Input<boolean>;
        last15secs?: pulumi.Input<boolean>;
        last1min?: pulumi.Input<boolean>;
        last5mins?: pulumi.Input<boolean>;
        less?: pulumi.Input<boolean>;
        lessOrEqual?: pulumi.Input<boolean>;
        local?: pulumi.Input<boolean>;
        major?: pulumi.Input<boolean>;
        matches?: pulumi.Input<boolean>;
        minor?: pulumi.Input<boolean>;
        missing?: pulumi.Input<boolean>;
        mss?: pulumi.Input<boolean>;
        not?: pulumi.Input<boolean>;
        org?: pulumi.Input<boolean>;
        password?: pulumi.Input<boolean>;
        path?: pulumi.Input<boolean>;
        pathSegment?: pulumi.Input<boolean>;
        port?: pulumi.Input<boolean>;
        present?: pulumi.Input<boolean>;
        protocol?: pulumi.Input<boolean>;
        queryParameter?: pulumi.Input<boolean>;
        queryString?: pulumi.Input<boolean>;
        regionCode?: pulumi.Input<boolean>;
        regionName?: pulumi.Input<boolean>;
        remote?: pulumi.Input<boolean>;
        request?: pulumi.Input<boolean>;
        response?: pulumi.Input<boolean>;
        routeDomain?: pulumi.Input<boolean>;
        rtt?: pulumi.Input<boolean>;
        scheme?: pulumi.Input<boolean>;
        serverName?: pulumi.Input<boolean>;
        sslCert?: pulumi.Input<boolean>;
        sslClientHello?: pulumi.Input<boolean>;
        sslExtension?: pulumi.Input<boolean>;
        sslServerHandshake?: pulumi.Input<boolean>;
        sslServerHello?: pulumi.Input<boolean>;
        startsWith?: pulumi.Input<boolean>;
        tcp?: pulumi.Input<boolean>;
        text?: pulumi.Input<boolean>;
        tmName?: pulumi.Input<string>;
        unnamedQueryParameter?: pulumi.Input<boolean>;
        userAgentToken?: pulumi.Input<boolean>;
        username?: pulumi.Input<boolean>;
        value?: pulumi.Input<boolean>;
        values?: pulumi.Input<pulumi.Input<string>[]>;
        version?: pulumi.Input<boolean>;
        vlan?: pulumi.Input<boolean>;
        vlanId?: pulumi.Input<boolean>;
    }

    export interface NodeFqdn {
        /**
         * Specifies the node's address family. The default is 'unspecified', or IP-agnostic. This needs to be specified inside the fqdn (fully qualified domain name).
         */
        addressFamily?: pulumi.Input<string>;
        autopopulate?: pulumi.Input<string>;
        downinterval?: pulumi.Input<number>;
        /**
         * Specifies the amount of time before sending the next DNS query. Default is 3600. This needs to be specified inside the fqdn (fully qualified domain name).
         */
        interval?: pulumi.Input<string>;
        /**
         * Name of the node
         */
        name?: pulumi.Input<string>;
    }

    export interface PolicyRule {
        actions?: pulumi.Input<pulumi.Input<inputs.ltm.PolicyRuleAction>[]>;
        conditions?: pulumi.Input<pulumi.Input<inputs.ltm.PolicyRuleCondition>[]>;
        /**
         * Name of the Policy ( policy name should be in full path which is combination of partition and policy name )
         */
        name: pulumi.Input<string>;
    }

    export interface PolicyRuleAction {
        appService?: pulumi.Input<string>;
        application?: pulumi.Input<string>;
        asm?: pulumi.Input<boolean>;
        avr?: pulumi.Input<boolean>;
        cache?: pulumi.Input<boolean>;
        carp?: pulumi.Input<boolean>;
        category?: pulumi.Input<string>;
        classify?: pulumi.Input<boolean>;
        clonePool?: pulumi.Input<string>;
        code?: pulumi.Input<number>;
        compress?: pulumi.Input<boolean>;
        connection?: pulumi.Input<boolean>;
        content?: pulumi.Input<string>;
        cookieHash?: pulumi.Input<boolean>;
        cookieInsert?: pulumi.Input<boolean>;
        cookiePassive?: pulumi.Input<boolean>;
        cookieRewrite?: pulumi.Input<boolean>;
        decompress?: pulumi.Input<boolean>;
        defer?: pulumi.Input<boolean>;
        destinationAddress?: pulumi.Input<boolean>;
        disable?: pulumi.Input<boolean>;
        domain?: pulumi.Input<string>;
        enable?: pulumi.Input<boolean>;
        expiry?: pulumi.Input<string>;
        expirySecs?: pulumi.Input<number>;
        expression?: pulumi.Input<string>;
        extension?: pulumi.Input<string>;
        facility?: pulumi.Input<string>;
        /**
         * This action will affect forwarding.
         */
        forward?: pulumi.Input<boolean>;
        fromProfile?: pulumi.Input<string>;
        hash?: pulumi.Input<boolean>;
        host?: pulumi.Input<string>;
        http?: pulumi.Input<boolean>;
        httpBasicAuth?: pulumi.Input<boolean>;
        httpCookie?: pulumi.Input<boolean>;
        httpHeader?: pulumi.Input<boolean>;
        httpHost?: pulumi.Input<boolean>;
        httpReferer?: pulumi.Input<boolean>;
        httpReply?: pulumi.Input<boolean>;
        httpSetCookie?: pulumi.Input<boolean>;
        httpUri?: pulumi.Input<boolean>;
        ifile?: pulumi.Input<string>;
        insert?: pulumi.Input<boolean>;
        internalVirtual?: pulumi.Input<string>;
        ipAddress?: pulumi.Input<string>;
        key?: pulumi.Input<string>;
        l7dos?: pulumi.Input<boolean>;
        length?: pulumi.Input<number>;
        location?: pulumi.Input<string>;
        log?: pulumi.Input<boolean>;
        ltmPolicy?: pulumi.Input<boolean>;
        member?: pulumi.Input<string>;
        message?: pulumi.Input<string>;
        netmask?: pulumi.Input<string>;
        nexthop?: pulumi.Input<string>;
        node?: pulumi.Input<string>;
        offset?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
        pem?: pulumi.Input<boolean>;
        persist?: pulumi.Input<boolean>;
        pin?: pulumi.Input<boolean>;
        policy?: pulumi.Input<string>;
        /**
         * This action will direct the stream to this pool.
         */
        pool?: pulumi.Input<string>;
        port?: pulumi.Input<number>;
        priority?: pulumi.Input<string>;
        profile?: pulumi.Input<string>;
        protocol?: pulumi.Input<string>;
        queryString?: pulumi.Input<string>;
        rateclass?: pulumi.Input<string>;
        redirect?: pulumi.Input<boolean>;
        remove?: pulumi.Input<boolean>;
        replace?: pulumi.Input<boolean>;
        request?: pulumi.Input<boolean>;
        requestAdapt?: pulumi.Input<boolean>;
        reset?: pulumi.Input<boolean>;
        response?: pulumi.Input<boolean>;
        responseAdapt?: pulumi.Input<boolean>;
        scheme?: pulumi.Input<string>;
        script?: pulumi.Input<string>;
        select?: pulumi.Input<boolean>;
        serverSsl?: pulumi.Input<boolean>;
        setVariable?: pulumi.Input<boolean>;
        shutdown?: pulumi.Input<boolean>;
        snat?: pulumi.Input<string>;
        snatpool?: pulumi.Input<string>;
        sourceAddress?: pulumi.Input<boolean>;
        sslClientHello?: pulumi.Input<boolean>;
        sslServerHandshake?: pulumi.Input<boolean>;
        sslServerHello?: pulumi.Input<boolean>;
        sslSessionId?: pulumi.Input<boolean>;
        status?: pulumi.Input<number>;
        tcl?: pulumi.Input<boolean>;
        tcpNagle?: pulumi.Input<boolean>;
        text?: pulumi.Input<string>;
        timeout?: pulumi.Input<number>;
        tmName?: pulumi.Input<string>;
        uie?: pulumi.Input<boolean>;
        universal?: pulumi.Input<boolean>;
        value?: pulumi.Input<string>;
        virtual?: pulumi.Input<string>;
        vlan?: pulumi.Input<string>;
        vlanId?: pulumi.Input<number>;
        wam?: pulumi.Input<boolean>;
        write?: pulumi.Input<boolean>;
    }

    export interface PolicyRuleCondition {
        address?: pulumi.Input<boolean>;
        all?: pulumi.Input<boolean>;
        appService?: pulumi.Input<string>;
        browserType?: pulumi.Input<boolean>;
        browserVersion?: pulumi.Input<boolean>;
        caseInsensitive?: pulumi.Input<boolean>;
        caseSensitive?: pulumi.Input<boolean>;
        cipher?: pulumi.Input<boolean>;
        cipherBits?: pulumi.Input<boolean>;
        clientSsl?: pulumi.Input<boolean>;
        code?: pulumi.Input<boolean>;
        commonName?: pulumi.Input<boolean>;
        contains?: pulumi.Input<boolean>;
        continent?: pulumi.Input<boolean>;
        countryCode?: pulumi.Input<boolean>;
        countryName?: pulumi.Input<boolean>;
        cpuUsage?: pulumi.Input<boolean>;
        deviceMake?: pulumi.Input<boolean>;
        deviceModel?: pulumi.Input<boolean>;
        domain?: pulumi.Input<boolean>;
        endsWith?: pulumi.Input<boolean>;
        equals?: pulumi.Input<boolean>;
        expiry?: pulumi.Input<boolean>;
        extension?: pulumi.Input<boolean>;
        external?: pulumi.Input<boolean>;
        geoip?: pulumi.Input<boolean>;
        greater?: pulumi.Input<boolean>;
        greaterOrEqual?: pulumi.Input<boolean>;
        host?: pulumi.Input<boolean>;
        httpBasicAuth?: pulumi.Input<boolean>;
        httpCookie?: pulumi.Input<boolean>;
        httpHeader?: pulumi.Input<boolean>;
        httpHost?: pulumi.Input<boolean>;
        httpMethod?: pulumi.Input<boolean>;
        httpReferer?: pulumi.Input<boolean>;
        httpSetCookie?: pulumi.Input<boolean>;
        httpStatus?: pulumi.Input<boolean>;
        httpUri?: pulumi.Input<boolean>;
        httpUserAgent?: pulumi.Input<boolean>;
        httpVersion?: pulumi.Input<boolean>;
        index?: pulumi.Input<number>;
        internal?: pulumi.Input<boolean>;
        isp?: pulumi.Input<boolean>;
        last15secs?: pulumi.Input<boolean>;
        last1min?: pulumi.Input<boolean>;
        last5mins?: pulumi.Input<boolean>;
        less?: pulumi.Input<boolean>;
        lessOrEqual?: pulumi.Input<boolean>;
        local?: pulumi.Input<boolean>;
        major?: pulumi.Input<boolean>;
        matches?: pulumi.Input<boolean>;
        minor?: pulumi.Input<boolean>;
        missing?: pulumi.Input<boolean>;
        mss?: pulumi.Input<boolean>;
        not?: pulumi.Input<boolean>;
        org?: pulumi.Input<boolean>;
        password?: pulumi.Input<boolean>;
        path?: pulumi.Input<boolean>;
        pathSegment?: pulumi.Input<boolean>;
        port?: pulumi.Input<boolean>;
        present?: pulumi.Input<boolean>;
        protocol?: pulumi.Input<boolean>;
        queryParameter?: pulumi.Input<boolean>;
        queryString?: pulumi.Input<boolean>;
        regionCode?: pulumi.Input<boolean>;
        regionName?: pulumi.Input<boolean>;
        remote?: pulumi.Input<boolean>;
        request?: pulumi.Input<boolean>;
        response?: pulumi.Input<boolean>;
        routeDomain?: pulumi.Input<boolean>;
        rtt?: pulumi.Input<boolean>;
        scheme?: pulumi.Input<boolean>;
        serverName?: pulumi.Input<boolean>;
        sslCert?: pulumi.Input<boolean>;
        sslClientHello?: pulumi.Input<boolean>;
        sslExtension?: pulumi.Input<boolean>;
        sslServerHandshake?: pulumi.Input<boolean>;
        sslServerHello?: pulumi.Input<boolean>;
        startsWith?: pulumi.Input<boolean>;
        tcp?: pulumi.Input<boolean>;
        text?: pulumi.Input<boolean>;
        tmName?: pulumi.Input<string>;
        unnamedQueryParameter?: pulumi.Input<boolean>;
        userAgentToken?: pulumi.Input<boolean>;
        username?: pulumi.Input<boolean>;
        value?: pulumi.Input<boolean>;
        values?: pulumi.Input<pulumi.Input<string>[]>;
        version?: pulumi.Input<boolean>;
        vlan?: pulumi.Input<boolean>;
        vlanId?: pulumi.Input<boolean>;
    }

    export interface ProfileClientSslCertKeyChain {
        /**
         * Specifies a cert name for use.
         */
        cert?: pulumi.Input<string>;
        /**
         * Contains a certificate chain that is relevant to the certificate and key mentioned earlier.This key is optional
         */
        chain?: pulumi.Input<string>;
        /**
         * Contains a key name
         */
        key?: pulumi.Input<string>;
        /**
         * Specifies the name of the profile.Name of Profile should be full path.The full path is the combination of the `partition + profile name`,For example `/Common/test-clientssl-profile`.
         */
        name?: pulumi.Input<string>;
        passphrase?: pulumi.Input<string>;
    }

    export interface SnatOrigin {
        appService?: pulumi.Input<string>;
        /**
         * Name of the SNAT, name of SNAT should be full path. Full path is the combination of the `partition + SNAT name`,For example `/Common/test-snat`.
         */
        name?: pulumi.Input<string>;
    }

}

export namespace net {
    export interface VlanInterface {
        /**
         * Specifies a list of tagged interfaces or trunks associated with this VLAN. Note that you can associate tagged interfaces or trunks with any number of VLANs.
         */
        tagged?: pulumi.Input<boolean>;
        /**
         * Physical or virtual port used for traffic
         */
        vlanport?: pulumi.Input<string>;
    }
}

export namespace ssl {
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

    export interface GetWafEntityUrlMethodOverrideArgs {
        /**
         * Specifies that the system allows or disallows a method for this URL
         */
        allow: pulumi.Input<boolean>;
        /**
         * Specifies an HTTP method.
         */
        method: pulumi.Input<string>;
    }
}

export namespace sys {
    export interface IAppList {
        encrypted?: pulumi.Input<string>;
        value?: pulumi.Input<string>;
    }

    export interface IAppMetadata {
        persists?: pulumi.Input<string>;
        value?: pulumi.Input<string>;
    }

    export interface IAppTable {
        columnNames?: pulumi.Input<pulumi.Input<string>[]>;
        encryptedColumns?: pulumi.Input<string>;
        /**
         * Name of the iApp.
         */
        name?: pulumi.Input<string>;
        rows?: pulumi.Input<pulumi.Input<inputs.sys.IAppTableRow>[]>;
    }

    export interface IAppTableRow {
        rows?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface IAppVariable {
        encrypted?: pulumi.Input<string>;
        /**
         * Name of the iApp.
         */
        name?: pulumi.Input<string>;
        value?: pulumi.Input<string>;
    }
}
