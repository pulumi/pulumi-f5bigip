import * as pulumi from "@pulumi/pulumi";
import * as f5bigip from "@pulumi/f5bigip";
import * as virtualappliance from "./virtualappliance";
import * as backendinstances from "./loadbalancedinstances";
import * as https from "https";
import * as fetch from "node-fetch";

function sleep(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

// Wait for the virtual appliance to be avaiabe to log in to.
const virtualapplicanceAddress = pulumi
    .all([virtualappliance.f5Address, virtualappliance.f5Password])
    .apply(async ([address, password]) => {
        let ready = false;
        const agent = new https.Agent({
            rejectUnauthorized: false,
        });
        const token = Buffer.from(`admin:${password}`).toString("base64");
        while (!ready) {
            try {
                const res = await fetch(`${address}/mgmt/tm/ltm`, {
                    agent,
                    headers: {
                        Authorization: `Basic ${token}`,
                    },
                });
                if (res.status == 200) {
                    console.log("bigip virtual appliance is ready");
                    ready = true;
                }
            } catch (err) {
                console.log(`caught exception: ${err}`);
            }
            // Keep trying
            console.log("waiting for bigip virtual appliance address to be ready");
            sleep(10000);
        }
        return address;
    });

const f5bigipProvider = new f5bigip.Provider("bigipProvider", {
    address: virtualapplicanceAddress,
    password: virtualappliance.f5Password,
    username: "admin",
});

const monitor = new f5bigip.ltm.Monitor(
    "backend",
    {
        name: "/Common/backend",
        parent: "/Common/http",
        send: "GET /\r\n",
        timeout: 5,
        interval: 10,
    },
    { provider: f5bigipProvider },
);

const pool = new f5bigip.ltm.Pool(
    "backend",
    {
        name: "/Common/backend",
        monitors: [monitor.name],
        allowNat: "yes",
        allowSnat: "yes",
    },
    { provider: f5bigipProvider },
);

backendinstances.instancePublicIps.map((backendAddress, i) => {
    const node = new f5bigip.ltm.Node(
        `backend-${i}`,
        {
            name: `/Common/node-${i}`,
            address: backendAddress.apply((x) => x.toString()),
        },
        { provider: f5bigipProvider },
    );

    new f5bigip.ltm.PoolAttachment(
        `backend-${i}`,
        {
            pool: pool.name,
            node: node.name.apply((x: string) => x + ":80"),
        },
        { provider: f5bigipProvider },
    );
});

new f5bigip.ltm.VirtualServer(
    "backend",
    {
        pool: pool.name,
        name: "/Common/backend",
        destination: virtualappliance.f5PrivateIp,
        port: 80,
        sourceAddressTranslation: "automap",
    },
    { provider: f5bigipProvider, dependsOn: [pool] },
);

export let instanceIp = virtualappliance.f5Address;
export let password = virtualappliance.f5Password;
