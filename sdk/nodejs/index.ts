// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { As3Args, As3State } from "./as3";
export type As3 = import("./as3").As3;
export const As3: typeof import("./as3").As3 = null as any;
utilities.lazyLoad(exports, ["As3"], () => require("./as3"));

export { BigIqAs3Args, BigIqAs3State } from "./bigIqAs3";
export type BigIqAs3 = import("./bigIqAs3").BigIqAs3;
export const BigIqAs3: typeof import("./bigIqAs3").BigIqAs3 = null as any;
utilities.lazyLoad(exports, ["BigIqAs3"], () => require("./bigIqAs3"));

export { CommandArgs, CommandState } from "./command";
export type Command = import("./command").Command;
export const Command: typeof import("./command").Command = null as any;
utilities.lazyLoad(exports, ["Command"], () => require("./command"));

export { CommonLicenseManageBigIqArgs, CommonLicenseManageBigIqState } from "./commonLicenseManageBigIq";
export type CommonLicenseManageBigIq = import("./commonLicenseManageBigIq").CommonLicenseManageBigIq;
export const CommonLicenseManageBigIq: typeof import("./commonLicenseManageBigIq").CommonLicenseManageBigIq = null as any;
utilities.lazyLoad(exports, ["CommonLicenseManageBigIq"], () => require("./commonLicenseManageBigIq"));

export { DoArgs, DoState } from "./do";
export type Do = import("./do").Do;
export const Do: typeof import("./do").Do = null as any;
utilities.lazyLoad(exports, ["Do"], () => require("./do"));

export { EventServiceDiscoveryArgs, EventServiceDiscoveryState } from "./eventServiceDiscovery";
export type EventServiceDiscovery = import("./eventServiceDiscovery").EventServiceDiscovery;
export const EventServiceDiscovery: typeof import("./eventServiceDiscovery").EventServiceDiscovery = null as any;
utilities.lazyLoad(exports, ["EventServiceDiscovery"], () => require("./eventServiceDiscovery"));

export { FastApplicationArgs, FastApplicationState } from "./fastApplication";
export type FastApplication = import("./fastApplication").FastApplication;
export const FastApplication: typeof import("./fastApplication").FastApplication = null as any;
utilities.lazyLoad(exports, ["FastApplication"], () => require("./fastApplication"));

export { FastHttpAppArgs, FastHttpAppState } from "./fastHttpApp";
export type FastHttpApp = import("./fastHttpApp").FastHttpApp;
export const FastHttpApp: typeof import("./fastHttpApp").FastHttpApp = null as any;
utilities.lazyLoad(exports, ["FastHttpApp"], () => require("./fastHttpApp"));

export { FastHttpsAppArgs, FastHttpsAppState } from "./fastHttpsApp";
export type FastHttpsApp = import("./fastHttpsApp").FastHttpsApp;
export const FastHttpsApp: typeof import("./fastHttpsApp").FastHttpsApp = null as any;
utilities.lazyLoad(exports, ["FastHttpsApp"], () => require("./fastHttpsApp"));

export { FastTcpAppArgs, FastTcpAppState } from "./fastTcpApp";
export type FastTcpApp = import("./fastTcpApp").FastTcpApp;
export const FastTcpApp: typeof import("./fastTcpApp").FastTcpApp = null as any;
utilities.lazyLoad(exports, ["FastTcpApp"], () => require("./fastTcpApp"));

export { FastTemplateArgs, FastTemplateState } from "./fastTemplate";
export type FastTemplate = import("./fastTemplate").FastTemplate;
export const FastTemplate: typeof import("./fastTemplate").FastTemplate = null as any;
utilities.lazyLoad(exports, ["FastTemplate"], () => require("./fastTemplate"));

export { FastUdpAppArgs, FastUdpAppState } from "./fastUdpApp";
export type FastUdpApp = import("./fastUdpApp").FastUdpApp;
export const FastUdpApp: typeof import("./fastUdpApp").FastUdpApp = null as any;
utilities.lazyLoad(exports, ["FastUdpApp"], () => require("./fastUdpApp"));

export { IpsecPolicyArgs, IpsecPolicyState } from "./ipsecPolicy";
export type IpsecPolicy = import("./ipsecPolicy").IpsecPolicy;
export const IpsecPolicy: typeof import("./ipsecPolicy").IpsecPolicy = null as any;
utilities.lazyLoad(exports, ["IpsecPolicy"], () => require("./ipsecPolicy"));

export { IpsecProfileArgs, IpsecProfileState } from "./ipsecProfile";
export type IpsecProfile = import("./ipsecProfile").IpsecProfile;
export const IpsecProfile: typeof import("./ipsecProfile").IpsecProfile = null as any;
utilities.lazyLoad(exports, ["IpsecProfile"], () => require("./ipsecProfile"));

export { NetIkePeerArgs, NetIkePeerState } from "./netIkePeer";
export type NetIkePeer = import("./netIkePeer").NetIkePeer;
export const NetIkePeer: typeof import("./netIkePeer").NetIkePeer = null as any;
utilities.lazyLoad(exports, ["NetIkePeer"], () => require("./netIkePeer"));

export { NetTunnelArgs, NetTunnelState } from "./netTunnel";
export type NetTunnel = import("./netTunnel").NetTunnel;
export const NetTunnel: typeof import("./netTunnel").NetTunnel = null as any;
utilities.lazyLoad(exports, ["NetTunnel"], () => require("./netTunnel"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { TrafficSelectorArgs, TrafficSelectorState } from "./trafficSelector";
export type TrafficSelector = import("./trafficSelector").TrafficSelector;
export const TrafficSelector: typeof import("./trafficSelector").TrafficSelector = null as any;
utilities.lazyLoad(exports, ["TrafficSelector"], () => require("./trafficSelector"));

export { WafPolicyArgs, WafPolicyState } from "./wafPolicy";
export type WafPolicy = import("./wafPolicy").WafPolicy;
export const WafPolicy: typeof import("./wafPolicy").WafPolicy = null as any;
utilities.lazyLoad(exports, ["WafPolicy"], () => require("./wafPolicy"));


// Export sub-modules:
import * as cm from "./cm";
import * as config from "./config";
import * as fast from "./fast";
import * as ltm from "./ltm";
import * as net from "./net";
import * as ssl from "./ssl";
import * as sys from "./sys";
import * as types from "./types";
import * as vcmp from "./vcmp";

export {
    cm,
    config,
    fast,
    ltm,
    net,
    ssl,
    sys,
    types,
    vcmp,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "f5bigip:index/as3:As3":
                return new As3(name, <any>undefined, { urn })
            case "f5bigip:index/bigIqAs3:BigIqAs3":
                return new BigIqAs3(name, <any>undefined, { urn })
            case "f5bigip:index/command:Command":
                return new Command(name, <any>undefined, { urn })
            case "f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq":
                return new CommonLicenseManageBigIq(name, <any>undefined, { urn })
            case "f5bigip:index/do:Do":
                return new Do(name, <any>undefined, { urn })
            case "f5bigip:index/eventServiceDiscovery:EventServiceDiscovery":
                return new EventServiceDiscovery(name, <any>undefined, { urn })
            case "f5bigip:index/fastApplication:FastApplication":
                return new FastApplication(name, <any>undefined, { urn })
            case "f5bigip:index/fastHttpApp:FastHttpApp":
                return new FastHttpApp(name, <any>undefined, { urn })
            case "f5bigip:index/fastHttpsApp:FastHttpsApp":
                return new FastHttpsApp(name, <any>undefined, { urn })
            case "f5bigip:index/fastTcpApp:FastTcpApp":
                return new FastTcpApp(name, <any>undefined, { urn })
            case "f5bigip:index/fastTemplate:FastTemplate":
                return new FastTemplate(name, <any>undefined, { urn })
            case "f5bigip:index/fastUdpApp:FastUdpApp":
                return new FastUdpApp(name, <any>undefined, { urn })
            case "f5bigip:index/ipsecPolicy:IpsecPolicy":
                return new IpsecPolicy(name, <any>undefined, { urn })
            case "f5bigip:index/ipsecProfile:IpsecProfile":
                return new IpsecProfile(name, <any>undefined, { urn })
            case "f5bigip:index/netIkePeer:NetIkePeer":
                return new NetIkePeer(name, <any>undefined, { urn })
            case "f5bigip:index/netTunnel:NetTunnel":
                return new NetTunnel(name, <any>undefined, { urn })
            case "f5bigip:index/trafficSelector:TrafficSelector":
                return new TrafficSelector(name, <any>undefined, { urn })
            case "f5bigip:index/wafPolicy:WafPolicy":
                return new WafPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("f5bigip", "index/as3", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/bigIqAs3", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/command", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/commonLicenseManageBigIq", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/do", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/eventServiceDiscovery", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/fastApplication", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/fastHttpApp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/fastHttpsApp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/fastTcpApp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/fastTemplate", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/fastUdpApp", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/ipsecPolicy", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/ipsecProfile", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/netIkePeer", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/netTunnel", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/trafficSelector", _module)
pulumi.runtime.registerResourceModule("f5bigip", "index/wafPolicy", _module)
pulumi.runtime.registerResourcePackage("f5bigip", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:f5bigip") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
