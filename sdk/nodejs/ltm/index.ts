// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./dataGroup";
export * from "./getDataGroup";
export * from "./getIrule";
export * from "./getMonitor";
export * from "./getNode";
export * from "./getPool";
export * from "./irule";
export * from "./monitor";
export * from "./node";
export * from "./persistenceProfileCookie";
export * from "./persistenceProfileDstAddr";
export * from "./persistenceProfileSrcAddr";
export * from "./persistenceProfileSsl";
export * from "./policy";
export * from "./pool";
export * from "./poolAttachment";
export * from "./profileClientSsl";
export * from "./profileFastHttp";
export * from "./profileFastL4";
export * from "./profileHttp";
export * from "./profileHttp2";
export * from "./profileHttpCompress";
export * from "./profileOneConnect";
export * from "./profileServerSsl";
export * from "./profileTcp";
export * from "./snat";
export * from "./snatPool";
export * from "./virtualAddress";
export * from "./virtualServer";

// Import resources to register:
import { DataGroup } from "./dataGroup";
import { IRule } from "./irule";
import { Monitor } from "./monitor";
import { Node } from "./node";
import { PersistenceProfileCookie } from "./persistenceProfileCookie";
import { PersistenceProfileDstAddr } from "./persistenceProfileDstAddr";
import { PersistenceProfileSrcAddr } from "./persistenceProfileSrcAddr";
import { PersistenceProfileSsl } from "./persistenceProfileSsl";
import { Policy } from "./policy";
import { Pool } from "./pool";
import { PoolAttachment } from "./poolAttachment";
import { ProfileClientSsl } from "./profileClientSsl";
import { ProfileFastHttp } from "./profileFastHttp";
import { ProfileFastL4 } from "./profileFastL4";
import { ProfileHttp2 } from "./profileHttp2";
import { ProfileHttp } from "./profileHttp";
import { ProfileHttpCompress } from "./profileHttpCompress";
import { ProfileOneConnect } from "./profileOneConnect";
import { ProfileServerSsl } from "./profileServerSsl";
import { ProfileTcp } from "./profileTcp";
import { Snat } from "./snat";
import { SnatPool } from "./snatPool";
import { VirtualAddress } from "./virtualAddress";
import { VirtualServer } from "./virtualServer";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "f5bigip:ltm/dataGroup:DataGroup":
                return new DataGroup(name, <any>undefined, { urn })
            case "f5bigip:ltm/iRule:IRule":
                return new IRule(name, <any>undefined, { urn })
            case "f5bigip:ltm/monitor:Monitor":
                return new Monitor(name, <any>undefined, { urn })
            case "f5bigip:ltm/node:Node":
                return new Node(name, <any>undefined, { urn })
            case "f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie":
                return new PersistenceProfileCookie(name, <any>undefined, { urn })
            case "f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr":
                return new PersistenceProfileDstAddr(name, <any>undefined, { urn })
            case "f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr":
                return new PersistenceProfileSrcAddr(name, <any>undefined, { urn })
            case "f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl":
                return new PersistenceProfileSsl(name, <any>undefined, { urn })
            case "f5bigip:ltm/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "f5bigip:ltm/pool:Pool":
                return new Pool(name, <any>undefined, { urn })
            case "f5bigip:ltm/poolAttachment:PoolAttachment":
                return new PoolAttachment(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileClientSsl:ProfileClientSsl":
                return new ProfileClientSsl(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileFastHttp:ProfileFastHttp":
                return new ProfileFastHttp(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileFastL4:ProfileFastL4":
                return new ProfileFastL4(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileHttp2:ProfileHttp2":
                return new ProfileHttp2(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileHttp:ProfileHttp":
                return new ProfileHttp(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileHttpCompress:ProfileHttpCompress":
                return new ProfileHttpCompress(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileOneConnect:ProfileOneConnect":
                return new ProfileOneConnect(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileServerSsl:ProfileServerSsl":
                return new ProfileServerSsl(name, <any>undefined, { urn })
            case "f5bigip:ltm/profileTcp:ProfileTcp":
                return new ProfileTcp(name, <any>undefined, { urn })
            case "f5bigip:ltm/snat:Snat":
                return new Snat(name, <any>undefined, { urn })
            case "f5bigip:ltm/snatPool:SnatPool":
                return new SnatPool(name, <any>undefined, { urn })
            case "f5bigip:ltm/virtualAddress:VirtualAddress":
                return new VirtualAddress(name, <any>undefined, { urn })
            case "f5bigip:ltm/virtualServer:VirtualServer":
                return new VirtualServer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("f5bigip", "ltm/dataGroup", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/iRule", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/monitor", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/node", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/persistenceProfileCookie", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/persistenceProfileDstAddr", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/persistenceProfileSrcAddr", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/persistenceProfileSsl", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/policy", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/pool", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/poolAttachment", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileClientSsl", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileFastHttp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileFastL4", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileHttp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileHttp2", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileHttpCompress", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileOneConnect", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileServerSsl", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/profileTcp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/snat", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/snatPool", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/virtualAddress", _module)
pulumi.runtime.registerResourceModule("f5bigip", "ltm/virtualServer", _module)
