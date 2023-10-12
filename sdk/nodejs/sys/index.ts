// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BigIpLicenseArgs, BigIpLicenseState } from "./bigIpLicense";
export type BigIpLicense = import("./bigIpLicense").BigIpLicense;
export const BigIpLicense: typeof import("./bigIpLicense").BigIpLicense = null as any;
utilities.lazyLoad(exports, ["BigIpLicense"], () => require("./bigIpLicense"));

export { DnsArgs, DnsState } from "./dns";
export type Dns = import("./dns").Dns;
export const Dns: typeof import("./dns").Dns = null as any;
utilities.lazyLoad(exports, ["Dns"], () => require("./dns"));

export { IAppArgs, IAppState } from "./iapp";
export type IApp = import("./iapp").IApp;
export const IApp: typeof import("./iapp").IApp = null as any;
utilities.lazyLoad(exports, ["IApp"], () => require("./iapp"));

export { NtpArgs, NtpState } from "./ntp";
export type Ntp = import("./ntp").Ntp;
export const Ntp: typeof import("./ntp").Ntp = null as any;
utilities.lazyLoad(exports, ["Ntp"], () => require("./ntp"));

export { OcspArgs, OcspState } from "./ocsp";
export type Ocsp = import("./ocsp").Ocsp;
export const Ocsp: typeof import("./ocsp").Ocsp = null as any;
utilities.lazyLoad(exports, ["Ocsp"], () => require("./ocsp"));

export { ProvisionArgs, ProvisionState } from "./provision";
export type Provision = import("./provision").Provision;
export const Provision: typeof import("./provision").Provision = null as any;
utilities.lazyLoad(exports, ["Provision"], () => require("./provision"));

export { SnmpArgs, SnmpState } from "./snmp";
export type Snmp = import("./snmp").Snmp;
export const Snmp: typeof import("./snmp").Snmp = null as any;
utilities.lazyLoad(exports, ["Snmp"], () => require("./snmp"));

export { SnmpTrapsArgs, SnmpTrapsState } from "./snmpTraps";
export type SnmpTraps = import("./snmpTraps").SnmpTraps;
export const SnmpTraps: typeof import("./snmpTraps").SnmpTraps = null as any;
utilities.lazyLoad(exports, ["SnmpTraps"], () => require("./snmpTraps"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "f5bigip:sys/bigIpLicense:BigIpLicense":
                return new BigIpLicense(name, <any>undefined, { urn })
            case "f5bigip:sys/dns:Dns":
                return new Dns(name, <any>undefined, { urn })
            case "f5bigip:sys/iApp:IApp":
                return new IApp(name, <any>undefined, { urn })
            case "f5bigip:sys/ntp:Ntp":
                return new Ntp(name, <any>undefined, { urn })
            case "f5bigip:sys/ocsp:Ocsp":
                return new Ocsp(name, <any>undefined, { urn })
            case "f5bigip:sys/provision:Provision":
                return new Provision(name, <any>undefined, { urn })
            case "f5bigip:sys/snmp:Snmp":
                return new Snmp(name, <any>undefined, { urn })
            case "f5bigip:sys/snmpTraps:SnmpTraps":
                return new SnmpTraps(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("f5bigip", "sys/bigIpLicense", _module)
pulumi.runtime.registerResourceModule("f5bigip", "sys/dns", _module)
pulumi.runtime.registerResourceModule("f5bigip", "sys/iApp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "sys/ntp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "sys/ocsp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "sys/provision", _module)
pulumi.runtime.registerResourceModule("f5bigip", "sys/snmp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "sys/snmpTraps", _module)
