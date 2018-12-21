// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    public readonly controls: pulumi.Output<string[]>;
    /**
     * Name of the Policy
     */
    public readonly name: pulumi.Output<string>;
    /**
     * Publish the Policy
     */
    public readonly publishedCopy: pulumi.Output<string | undefined>;
    public readonly requires: pulumi.Output<string[]>;
    public readonly rules: pulumi.Output<{ actions?: { appService: string, application: string, asm: boolean, avr: boolean, cache: boolean, carp: boolean, category: string, classify: boolean, clonePool: string, code: number, compress: boolean, content: string, cookieHash: boolean, cookieInsert: boolean, cookiePassive: boolean, cookieRewrite: boolean, decompress: boolean, defer: boolean, destinationAddress: boolean, disable: boolean, domain: string, enable: boolean, expiry: string, expirySecs: number, expression: string, extension: string, facility: string, forward: boolean, fromProfile: string, hash: boolean, host: string, http: boolean, httpBasicAuth: boolean, httpCookie: boolean, httpHeader: boolean, httpHost?: boolean, httpReferer: boolean, httpReply: boolean, httpSetCookie: boolean, httpUri: boolean, ifile: string, insert: boolean, internalVirtual: string, ipAddress: string, key: string, l7dos: boolean, length: number, location: string, log: boolean, ltmPolicy: boolean, member: string, message: string, netmask: string, nexthop: string, node: string, offset: number, path: string, pem: boolean, persist: boolean, pin: boolean, policy: string, pool: string, port: number, priority: string, profile: string, protocol: string, queryString: string, rateclass: string, redirect: boolean, remove: boolean, replace: boolean, request: boolean, requestAdapt: boolean, reset: boolean, response: boolean, responseAdapt: boolean, scheme: string, script: string, select: boolean, serverSsl: boolean, setVariable: boolean, snat: string, snatpool: string, sourceAddress: boolean, sslClientHello: boolean, sslServerHandshake: boolean, sslServerHello: boolean, sslSessionId: boolean, status: number, tcl: boolean, tcpNagle: boolean, text: string, timeout: number, tmName: string, uie: boolean, universal: boolean, value: string, virtual: string, vlan: string, vlanId: number, wam: boolean, write: boolean }[], conditions?: { address: boolean, all: boolean, appService: string, browserType: boolean, browserVersion: boolean, caseInsensitive: boolean, caseSensitive: boolean, cipher: boolean, cipherBits: boolean, clientSsl: boolean, code: boolean, commonName: boolean, contains: boolean, continent: boolean, countryCode: boolean, countryName: boolean, cpuUsage: boolean, deviceMake: boolean, deviceModel: boolean, domain: boolean, endsWith: boolean, equals: boolean, expiry: boolean, extension: boolean, external: boolean, geoip: boolean, greater: boolean, greaterOrEqual: boolean, host: boolean, httpBasicAuth: boolean, httpCookie: boolean, httpHeader: boolean, httpHost: boolean, httpMethod: boolean, httpReferer: boolean, httpSetCookie: boolean, httpStatus: boolean, httpUri: boolean, httpUserAgent: boolean, httpVersion: boolean, index: number, internal: boolean, isp: boolean, last15secs: boolean, last1min: boolean, last5mins: boolean, less: boolean, lessOrEqual: boolean, local: boolean, major: boolean, matches: boolean, minor: boolean, missing: boolean, mss: boolean, not: boolean, org: boolean, password: boolean, path: boolean, pathSegment: boolean, port: boolean, present: boolean, protocol: boolean, queryParameter: boolean, queryString: boolean, regionCode: boolean, regionName: boolean, remote: boolean, request: boolean, response: boolean, routeDomain: boolean, rtt: boolean, scheme: boolean, serverName: boolean, sslCert: boolean, sslClientHello: boolean, sslExtension: boolean, sslServerHandshake: boolean, sslServerHello: boolean, startsWith: boolean, tcp: boolean, text: boolean, tmName: string, unnamedQueryParameter: boolean, userAgentToken: boolean, username: boolean, value: boolean, values: string[], version: boolean, vlan: boolean, vlanId: boolean }[], name: string }[]>;
    /**
     * Policy Strategy (i.e. /Common/first-match)
     */
    public readonly strategy: pulumi.Output<string | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: PolicyState = argsOrState as PolicyState | undefined;
            inputs["controls"] = state ? state.controls : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publishedCopy"] = state ? state.publishedCopy : undefined;
            inputs["requires"] = state ? state.requires : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["strategy"] = state ? state.strategy : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.controls === undefined) {
                throw new Error("Missing required property 'controls'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.requires === undefined) {
                throw new Error("Missing required property 'requires'");
            }
            if (!args || args.rules === undefined) {
                throw new Error("Missing required property 'rules'");
            }
            inputs["controls"] = args ? args.controls : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publishedCopy"] = args ? args.publishedCopy : undefined;
            inputs["requires"] = args ? args.requires : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["strategy"] = args ? args.strategy : undefined;
        }
        super("f5bigip:ltm/policy:Policy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    readonly controls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Policy
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Publish the Policy
     */
    readonly publishedCopy?: pulumi.Input<string>;
    readonly requires?: pulumi.Input<pulumi.Input<string>[]>;
    readonly rules?: pulumi.Input<pulumi.Input<{ actions?: pulumi.Input<pulumi.Input<{ appService?: pulumi.Input<string>, application?: pulumi.Input<string>, asm?: pulumi.Input<boolean>, avr?: pulumi.Input<boolean>, cache?: pulumi.Input<boolean>, carp?: pulumi.Input<boolean>, category?: pulumi.Input<string>, classify?: pulumi.Input<boolean>, clonePool?: pulumi.Input<string>, code?: pulumi.Input<number>, compress?: pulumi.Input<boolean>, content?: pulumi.Input<string>, cookieHash?: pulumi.Input<boolean>, cookieInsert?: pulumi.Input<boolean>, cookiePassive?: pulumi.Input<boolean>, cookieRewrite?: pulumi.Input<boolean>, decompress?: pulumi.Input<boolean>, defer?: pulumi.Input<boolean>, destinationAddress?: pulumi.Input<boolean>, disable?: pulumi.Input<boolean>, domain?: pulumi.Input<string>, enable?: pulumi.Input<boolean>, expiry?: pulumi.Input<string>, expirySecs?: pulumi.Input<number>, expression?: pulumi.Input<string>, extension?: pulumi.Input<string>, facility?: pulumi.Input<string>, forward?: pulumi.Input<boolean>, fromProfile?: pulumi.Input<string>, hash?: pulumi.Input<boolean>, host?: pulumi.Input<string>, http?: pulumi.Input<boolean>, httpBasicAuth?: pulumi.Input<boolean>, httpCookie?: pulumi.Input<boolean>, httpHeader?: pulumi.Input<boolean>, httpHost?: pulumi.Input<boolean>, httpReferer?: pulumi.Input<boolean>, httpReply?: pulumi.Input<boolean>, httpSetCookie?: pulumi.Input<boolean>, httpUri?: pulumi.Input<boolean>, ifile?: pulumi.Input<string>, insert?: pulumi.Input<boolean>, internalVirtual?: pulumi.Input<string>, ipAddress?: pulumi.Input<string>, key?: pulumi.Input<string>, l7dos?: pulumi.Input<boolean>, length?: pulumi.Input<number>, location?: pulumi.Input<string>, log?: pulumi.Input<boolean>, ltmPolicy?: pulumi.Input<boolean>, member?: pulumi.Input<string>, message?: pulumi.Input<string>, netmask?: pulumi.Input<string>, nexthop?: pulumi.Input<string>, node?: pulumi.Input<string>, offset?: pulumi.Input<number>, path?: pulumi.Input<string>, pem?: pulumi.Input<boolean>, persist?: pulumi.Input<boolean>, pin?: pulumi.Input<boolean>, policy?: pulumi.Input<string>, pool?: pulumi.Input<string>, port?: pulumi.Input<number>, priority?: pulumi.Input<string>, profile?: pulumi.Input<string>, protocol?: pulumi.Input<string>, queryString?: pulumi.Input<string>, rateclass?: pulumi.Input<string>, redirect?: pulumi.Input<boolean>, remove?: pulumi.Input<boolean>, replace?: pulumi.Input<boolean>, request?: pulumi.Input<boolean>, requestAdapt?: pulumi.Input<boolean>, reset?: pulumi.Input<boolean>, response?: pulumi.Input<boolean>, responseAdapt?: pulumi.Input<boolean>, scheme?: pulumi.Input<string>, script?: pulumi.Input<string>, select?: pulumi.Input<boolean>, serverSsl?: pulumi.Input<boolean>, setVariable?: pulumi.Input<boolean>, snat?: pulumi.Input<string>, snatpool?: pulumi.Input<string>, sourceAddress?: pulumi.Input<boolean>, sslClientHello?: pulumi.Input<boolean>, sslServerHandshake?: pulumi.Input<boolean>, sslServerHello?: pulumi.Input<boolean>, sslSessionId?: pulumi.Input<boolean>, status?: pulumi.Input<number>, tcl?: pulumi.Input<boolean>, tcpNagle?: pulumi.Input<boolean>, text?: pulumi.Input<string>, timeout?: pulumi.Input<number>, tmName?: pulumi.Input<string>, uie?: pulumi.Input<boolean>, universal?: pulumi.Input<boolean>, value?: pulumi.Input<string>, virtual?: pulumi.Input<string>, vlan?: pulumi.Input<string>, vlanId?: pulumi.Input<number>, wam?: pulumi.Input<boolean>, write?: pulumi.Input<boolean> }>[]>, conditions?: pulumi.Input<pulumi.Input<{ address?: pulumi.Input<boolean>, all?: pulumi.Input<boolean>, appService?: pulumi.Input<string>, browserType?: pulumi.Input<boolean>, browserVersion?: pulumi.Input<boolean>, caseInsensitive?: pulumi.Input<boolean>, caseSensitive?: pulumi.Input<boolean>, cipher?: pulumi.Input<boolean>, cipherBits?: pulumi.Input<boolean>, clientSsl?: pulumi.Input<boolean>, code?: pulumi.Input<boolean>, commonName?: pulumi.Input<boolean>, contains?: pulumi.Input<boolean>, continent?: pulumi.Input<boolean>, countryCode?: pulumi.Input<boolean>, countryName?: pulumi.Input<boolean>, cpuUsage?: pulumi.Input<boolean>, deviceMake?: pulumi.Input<boolean>, deviceModel?: pulumi.Input<boolean>, domain?: pulumi.Input<boolean>, endsWith?: pulumi.Input<boolean>, equals?: pulumi.Input<boolean>, expiry?: pulumi.Input<boolean>, extension?: pulumi.Input<boolean>, external?: pulumi.Input<boolean>, geoip?: pulumi.Input<boolean>, greater?: pulumi.Input<boolean>, greaterOrEqual?: pulumi.Input<boolean>, host?: pulumi.Input<boolean>, httpBasicAuth?: pulumi.Input<boolean>, httpCookie?: pulumi.Input<boolean>, httpHeader?: pulumi.Input<boolean>, httpHost?: pulumi.Input<boolean>, httpMethod?: pulumi.Input<boolean>, httpReferer?: pulumi.Input<boolean>, httpSetCookie?: pulumi.Input<boolean>, httpStatus?: pulumi.Input<boolean>, httpUri?: pulumi.Input<boolean>, httpUserAgent?: pulumi.Input<boolean>, httpVersion?: pulumi.Input<boolean>, index?: pulumi.Input<number>, internal?: pulumi.Input<boolean>, isp?: pulumi.Input<boolean>, last15secs?: pulumi.Input<boolean>, last1min?: pulumi.Input<boolean>, last5mins?: pulumi.Input<boolean>, less?: pulumi.Input<boolean>, lessOrEqual?: pulumi.Input<boolean>, local?: pulumi.Input<boolean>, major?: pulumi.Input<boolean>, matches?: pulumi.Input<boolean>, minor?: pulumi.Input<boolean>, missing?: pulumi.Input<boolean>, mss?: pulumi.Input<boolean>, not?: pulumi.Input<boolean>, org?: pulumi.Input<boolean>, password?: pulumi.Input<boolean>, path?: pulumi.Input<boolean>, pathSegment?: pulumi.Input<boolean>, port?: pulumi.Input<boolean>, present?: pulumi.Input<boolean>, protocol?: pulumi.Input<boolean>, queryParameter?: pulumi.Input<boolean>, queryString?: pulumi.Input<boolean>, regionCode?: pulumi.Input<boolean>, regionName?: pulumi.Input<boolean>, remote?: pulumi.Input<boolean>, request?: pulumi.Input<boolean>, response?: pulumi.Input<boolean>, routeDomain?: pulumi.Input<boolean>, rtt?: pulumi.Input<boolean>, scheme?: pulumi.Input<boolean>, serverName?: pulumi.Input<boolean>, sslCert?: pulumi.Input<boolean>, sslClientHello?: pulumi.Input<boolean>, sslExtension?: pulumi.Input<boolean>, sslServerHandshake?: pulumi.Input<boolean>, sslServerHello?: pulumi.Input<boolean>, startsWith?: pulumi.Input<boolean>, tcp?: pulumi.Input<boolean>, text?: pulumi.Input<boolean>, tmName?: pulumi.Input<string>, unnamedQueryParameter?: pulumi.Input<boolean>, userAgentToken?: pulumi.Input<boolean>, username?: pulumi.Input<boolean>, value?: pulumi.Input<boolean>, values?: pulumi.Input<pulumi.Input<string>[]>, version?: pulumi.Input<boolean>, vlan?: pulumi.Input<boolean>, vlanId?: pulumi.Input<boolean> }>[]>, name: pulumi.Input<string> }>[]>;
    /**
     * Policy Strategy (i.e. /Common/first-match)
     */
    readonly strategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    readonly controls: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the Policy
     */
    readonly name: pulumi.Input<string>;
    /**
     * Publish the Policy
     */
    readonly publishedCopy?: pulumi.Input<string>;
    readonly requires: pulumi.Input<pulumi.Input<string>[]>;
    readonly rules: pulumi.Input<pulumi.Input<{ actions?: pulumi.Input<pulumi.Input<{ appService?: pulumi.Input<string>, application?: pulumi.Input<string>, asm?: pulumi.Input<boolean>, avr?: pulumi.Input<boolean>, cache?: pulumi.Input<boolean>, carp?: pulumi.Input<boolean>, category?: pulumi.Input<string>, classify?: pulumi.Input<boolean>, clonePool?: pulumi.Input<string>, code?: pulumi.Input<number>, compress?: pulumi.Input<boolean>, content?: pulumi.Input<string>, cookieHash?: pulumi.Input<boolean>, cookieInsert?: pulumi.Input<boolean>, cookiePassive?: pulumi.Input<boolean>, cookieRewrite?: pulumi.Input<boolean>, decompress?: pulumi.Input<boolean>, defer?: pulumi.Input<boolean>, destinationAddress?: pulumi.Input<boolean>, disable?: pulumi.Input<boolean>, domain?: pulumi.Input<string>, enable?: pulumi.Input<boolean>, expiry?: pulumi.Input<string>, expirySecs?: pulumi.Input<number>, expression?: pulumi.Input<string>, extension?: pulumi.Input<string>, facility?: pulumi.Input<string>, forward?: pulumi.Input<boolean>, fromProfile?: pulumi.Input<string>, hash?: pulumi.Input<boolean>, host?: pulumi.Input<string>, http?: pulumi.Input<boolean>, httpBasicAuth?: pulumi.Input<boolean>, httpCookie?: pulumi.Input<boolean>, httpHeader?: pulumi.Input<boolean>, httpHost?: pulumi.Input<boolean>, httpReferer?: pulumi.Input<boolean>, httpReply?: pulumi.Input<boolean>, httpSetCookie?: pulumi.Input<boolean>, httpUri?: pulumi.Input<boolean>, ifile?: pulumi.Input<string>, insert?: pulumi.Input<boolean>, internalVirtual?: pulumi.Input<string>, ipAddress?: pulumi.Input<string>, key?: pulumi.Input<string>, l7dos?: pulumi.Input<boolean>, length?: pulumi.Input<number>, location?: pulumi.Input<string>, log?: pulumi.Input<boolean>, ltmPolicy?: pulumi.Input<boolean>, member?: pulumi.Input<string>, message?: pulumi.Input<string>, netmask?: pulumi.Input<string>, nexthop?: pulumi.Input<string>, node?: pulumi.Input<string>, offset?: pulumi.Input<number>, path?: pulumi.Input<string>, pem?: pulumi.Input<boolean>, persist?: pulumi.Input<boolean>, pin?: pulumi.Input<boolean>, policy?: pulumi.Input<string>, pool?: pulumi.Input<string>, port?: pulumi.Input<number>, priority?: pulumi.Input<string>, profile?: pulumi.Input<string>, protocol?: pulumi.Input<string>, queryString?: pulumi.Input<string>, rateclass?: pulumi.Input<string>, redirect?: pulumi.Input<boolean>, remove?: pulumi.Input<boolean>, replace?: pulumi.Input<boolean>, request?: pulumi.Input<boolean>, requestAdapt?: pulumi.Input<boolean>, reset?: pulumi.Input<boolean>, response?: pulumi.Input<boolean>, responseAdapt?: pulumi.Input<boolean>, scheme?: pulumi.Input<string>, script?: pulumi.Input<string>, select?: pulumi.Input<boolean>, serverSsl?: pulumi.Input<boolean>, setVariable?: pulumi.Input<boolean>, snat?: pulumi.Input<string>, snatpool?: pulumi.Input<string>, sourceAddress?: pulumi.Input<boolean>, sslClientHello?: pulumi.Input<boolean>, sslServerHandshake?: pulumi.Input<boolean>, sslServerHello?: pulumi.Input<boolean>, sslSessionId?: pulumi.Input<boolean>, status?: pulumi.Input<number>, tcl?: pulumi.Input<boolean>, tcpNagle?: pulumi.Input<boolean>, text?: pulumi.Input<string>, timeout?: pulumi.Input<number>, tmName?: pulumi.Input<string>, uie?: pulumi.Input<boolean>, universal?: pulumi.Input<boolean>, value?: pulumi.Input<string>, virtual?: pulumi.Input<string>, vlan?: pulumi.Input<string>, vlanId?: pulumi.Input<number>, wam?: pulumi.Input<boolean>, write?: pulumi.Input<boolean> }>[]>, conditions?: pulumi.Input<pulumi.Input<{ address?: pulumi.Input<boolean>, all?: pulumi.Input<boolean>, appService?: pulumi.Input<string>, browserType?: pulumi.Input<boolean>, browserVersion?: pulumi.Input<boolean>, caseInsensitive?: pulumi.Input<boolean>, caseSensitive?: pulumi.Input<boolean>, cipher?: pulumi.Input<boolean>, cipherBits?: pulumi.Input<boolean>, clientSsl?: pulumi.Input<boolean>, code?: pulumi.Input<boolean>, commonName?: pulumi.Input<boolean>, contains?: pulumi.Input<boolean>, continent?: pulumi.Input<boolean>, countryCode?: pulumi.Input<boolean>, countryName?: pulumi.Input<boolean>, cpuUsage?: pulumi.Input<boolean>, deviceMake?: pulumi.Input<boolean>, deviceModel?: pulumi.Input<boolean>, domain?: pulumi.Input<boolean>, endsWith?: pulumi.Input<boolean>, equals?: pulumi.Input<boolean>, expiry?: pulumi.Input<boolean>, extension?: pulumi.Input<boolean>, external?: pulumi.Input<boolean>, geoip?: pulumi.Input<boolean>, greater?: pulumi.Input<boolean>, greaterOrEqual?: pulumi.Input<boolean>, host?: pulumi.Input<boolean>, httpBasicAuth?: pulumi.Input<boolean>, httpCookie?: pulumi.Input<boolean>, httpHeader?: pulumi.Input<boolean>, httpHost?: pulumi.Input<boolean>, httpMethod?: pulumi.Input<boolean>, httpReferer?: pulumi.Input<boolean>, httpSetCookie?: pulumi.Input<boolean>, httpStatus?: pulumi.Input<boolean>, httpUri?: pulumi.Input<boolean>, httpUserAgent?: pulumi.Input<boolean>, httpVersion?: pulumi.Input<boolean>, index?: pulumi.Input<number>, internal?: pulumi.Input<boolean>, isp?: pulumi.Input<boolean>, last15secs?: pulumi.Input<boolean>, last1min?: pulumi.Input<boolean>, last5mins?: pulumi.Input<boolean>, less?: pulumi.Input<boolean>, lessOrEqual?: pulumi.Input<boolean>, local?: pulumi.Input<boolean>, major?: pulumi.Input<boolean>, matches?: pulumi.Input<boolean>, minor?: pulumi.Input<boolean>, missing?: pulumi.Input<boolean>, mss?: pulumi.Input<boolean>, not?: pulumi.Input<boolean>, org?: pulumi.Input<boolean>, password?: pulumi.Input<boolean>, path?: pulumi.Input<boolean>, pathSegment?: pulumi.Input<boolean>, port?: pulumi.Input<boolean>, present?: pulumi.Input<boolean>, protocol?: pulumi.Input<boolean>, queryParameter?: pulumi.Input<boolean>, queryString?: pulumi.Input<boolean>, regionCode?: pulumi.Input<boolean>, regionName?: pulumi.Input<boolean>, remote?: pulumi.Input<boolean>, request?: pulumi.Input<boolean>, response?: pulumi.Input<boolean>, routeDomain?: pulumi.Input<boolean>, rtt?: pulumi.Input<boolean>, scheme?: pulumi.Input<boolean>, serverName?: pulumi.Input<boolean>, sslCert?: pulumi.Input<boolean>, sslClientHello?: pulumi.Input<boolean>, sslExtension?: pulumi.Input<boolean>, sslServerHandshake?: pulumi.Input<boolean>, sslServerHello?: pulumi.Input<boolean>, startsWith?: pulumi.Input<boolean>, tcp?: pulumi.Input<boolean>, text?: pulumi.Input<boolean>, tmName?: pulumi.Input<string>, unnamedQueryParameter?: pulumi.Input<boolean>, userAgentToken?: pulumi.Input<boolean>, username?: pulumi.Input<boolean>, value?: pulumi.Input<boolean>, values?: pulumi.Input<pulumi.Input<string>[]>, version?: pulumi.Input<boolean>, vlan?: pulumi.Input<boolean>, vlanId?: pulumi.Input<boolean> }>[]>, name: pulumi.Input<string> }>[]>;
    /**
     * Policy Strategy (i.e. /Common/first-match)
     */
    readonly strategy?: pulumi.Input<string>;
}
