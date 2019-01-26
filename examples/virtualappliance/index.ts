import * as pulumi from "@pulumi/pulumi";
import * as f5bigip from "@pulumi/f5bigip";
import * as virutalapplicance from "./virtualappliance";
import * as backendinstances from "./loadbalancedinstances";
import fetch from "node-fetch";
const https = require("https");
const agent = new https.Agent({
    rejectUnauthorized: false
})

const config = new pulumi.Config();

const baseTags = {
    project: `${pulumi.getProject()}-${pulumi.getStack()}`,
};

const virtualapplicanceAddress = virutalapplicance.f5Address.apply(async (address) => {
    let ready = false;
    while (!ready) {
        try {
            const res = await fetch(address, { agent });
            if (res.status == 200) {
                ready = true;
            }
            else {
                console.log("waiting for bigip virtual appliance address to be healthy");
            }
        } catch {
            // Keep trying on failure
            console.log("bigip virtual appliance not yet ready")

            continue;
        }
    }
    return address;
});

const f5bigipProvider = new f5bigip.Provider("bigipProvider", {
    address: virtualapplicanceAddress,
    password: virutalapplicance.f5Password,
    username: "admin",
});

const monitor = new f5bigip.ltm.Monitor("backend", {
    name: "/Common/backend",
    parent: "/Common/http",
    send: "GET /\r\n",
    timeout: 5,
    interval: 10,
}, { provider: f5bigipProvider });

const pool = new f5bigip.ltm.Pool("backend", {
    name: "/Common/backend",
    monitors: [monitor.name],
    allowNat: "yes",
    allowSnat: "yes",
}, { provider: f5bigipProvider });

const poolAttachments = backendinstances.instancePublicIps.map((backendAddress, i) => {
    const applicationPoolAttachment = new f5bigip.ltm.PoolAttachment(`backend-${i}`, {
        pool: pool.name,
        node: pulumi.interpolate`/Common/${backendAddress}:80`,
    }, { provider: f5bigipProvider });
    return applicationPoolAttachment;
});

const virtualServer = new f5bigip.ltm.VirtualServer("backend", {
    pool: pool.name,
    name: "/Common/backend",
    destination: virutalapplicance.f5PrivateIp,
    port: 80,
    sourceAddressTranslation: "automap",
}, { provider: f5bigipProvider, dependsOn: [pool] });

export let password = virutalapplicance.f5Password;
