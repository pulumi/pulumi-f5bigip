// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.IpsecPolicy` Manage IPSec policies on a BIG-IP
 *
 * Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/test-policy)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test_policy = new f5bigip.IpsecPolicy("test-policy", {
 *     authAlgorithm: "sha1",
 *     description: "created by terraform provider",
 *     encryptAlgorithm: "3des",
 *     ipcomp: "deflate",
 *     lifetime: 3,
 *     mode: "tunnel",
 *     name: "/Common/test-policy",
 *     protocol: "esp",
 *     tunnelLocalAddress: "192.168.1.1",
 *     tunnelRemoteAddress: "10.10.1.1",
 * });
 * ```
 */
export class IpsecPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IpsecPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsecPolicyState, opts?: pulumi.CustomResourceOptions): IpsecPolicy {
        return new IpsecPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/ipsecPolicy:IpsecPolicy';

    /**
     * Returns true if the given object is an instance of IpsecPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpsecPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpsecPolicy.__pulumiType;
    }

    /**
     * Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
     * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
     */
    public readonly authAlgorithm!: pulumi.Output<string>;
    /**
     * Description of the IPSec policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
     * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
     */
    public readonly encryptAlgorithm!: pulumi.Output<string>;
    /**
     * Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
     */
    public readonly ipcomp!: pulumi.Output<string>;
    /**
     * Specifies the length of time before the IKE security association expires, in kilobytes.
     */
    public readonly kbLifetime!: pulumi.Output<number>;
    /**
     * Specifies the length of time before the IKE security association expires, in minutes.
     */
    public readonly lifetime!: pulumi.Output<number>;
    /**
     * Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
     * modp4096, modp6144, modp8192`
     */
    public readonly perfectForwardSecrecy!: pulumi.Output<string>;
    /**
     * Specifies the IPsec protocol. Valid choices are: `ah, esp`
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     */
    public readonly tunnelLocalAddress!: pulumi.Output<string>;
    /**
     * Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     */
    public readonly tunnelRemoteAddress!: pulumi.Output<string>;

    /**
     * Create a IpsecPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpsecPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpsecPolicyArgs | IpsecPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpsecPolicyState | undefined;
            inputs["authAlgorithm"] = state ? state.authAlgorithm : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encryptAlgorithm"] = state ? state.encryptAlgorithm : undefined;
            inputs["ipcomp"] = state ? state.ipcomp : undefined;
            inputs["kbLifetime"] = state ? state.kbLifetime : undefined;
            inputs["lifetime"] = state ? state.lifetime : undefined;
            inputs["mode"] = state ? state.mode : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["perfectForwardSecrecy"] = state ? state.perfectForwardSecrecy : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["tunnelLocalAddress"] = state ? state.tunnelLocalAddress : undefined;
            inputs["tunnelRemoteAddress"] = state ? state.tunnelRemoteAddress : undefined;
        } else {
            const args = argsOrState as IpsecPolicyArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["authAlgorithm"] = args ? args.authAlgorithm : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encryptAlgorithm"] = args ? args.encryptAlgorithm : undefined;
            inputs["ipcomp"] = args ? args.ipcomp : undefined;
            inputs["kbLifetime"] = args ? args.kbLifetime : undefined;
            inputs["lifetime"] = args ? args.lifetime : undefined;
            inputs["mode"] = args ? args.mode : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["perfectForwardSecrecy"] = args ? args.perfectForwardSecrecy : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["tunnelLocalAddress"] = args ? args.tunnelLocalAddress : undefined;
            inputs["tunnelRemoteAddress"] = args ? args.tunnelRemoteAddress : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IpsecPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsecPolicy resources.
 */
export interface IpsecPolicyState {
    /**
     * Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
     * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
     */
    readonly authAlgorithm?: pulumi.Input<string>;
    /**
     * Description of the IPSec policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
     * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
     */
    readonly encryptAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
     */
    readonly ipcomp?: pulumi.Input<string>;
    /**
     * Specifies the length of time before the IKE security association expires, in kilobytes.
     */
    readonly kbLifetime?: pulumi.Input<number>;
    /**
     * Specifies the length of time before the IKE security association expires, in minutes.
     */
    readonly lifetime?: pulumi.Input<number>;
    /**
     * Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
     * modp4096, modp6144, modp8192`
     */
    readonly perfectForwardSecrecy?: pulumi.Input<string>;
    /**
     * Specifies the IPsec protocol. Valid choices are: `ah, esp`
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     */
    readonly tunnelLocalAddress?: pulumi.Input<string>;
    /**
     * Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     */
    readonly tunnelRemoteAddress?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpsecPolicy resource.
 */
export interface IpsecPolicyArgs {
    /**
     * Specifies the algorithm to use for IKE authentication. Valid choices are: `sha1, sha256, sha384, sha512, aes-gcm128,
     * aes-gcm192, aes-gcm256, aes-gmac128, aes-gmac192, aes-gmac256`
     */
    readonly authAlgorithm?: pulumi.Input<string>;
    /**
     * Description of the IPSec policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the algorithm to use for IKE encryption. Valid choices are: `null, 3des, aes128, aes192, aes256, aes-gmac256,
     * aes-gmac192, aes-gmac128, aes-gcm256, aes-gcm192, aes-gcm256, aes-gcm128`
     */
    readonly encryptAlgorithm?: pulumi.Input<string>;
    /**
     * Specifies whether to use IPComp encapsulation. Valid choices are: `none", null", deflate`
     */
    readonly ipcomp?: pulumi.Input<string>;
    /**
     * Specifies the length of time before the IKE security association expires, in kilobytes.
     */
    readonly kbLifetime?: pulumi.Input<number>;
    /**
     * Specifies the length of time before the IKE security association expires, in minutes.
     */
    readonly lifetime?: pulumi.Input<number>;
    /**
     * Specifies the processing mode. Valid choices are: `transport, interface, isession, tunnel`
     */
    readonly mode?: pulumi.Input<string>;
    /**
     * Name of the IPSec policy,it should be "full path".The full path is the combination of the partition + name of the IPSec policy.(For example `/Common/test-policy`)
     */
    readonly name: pulumi.Input<string>;
    /**
     * Specifies the Diffie-Hellman group to use for IKE Phase 2 negotiation. Valid choices are: `none, modp768, modp1024, modp1536, modp2048, modp3072,
     * modp4096, modp6144, modp8192`
     */
    readonly perfectForwardSecrecy?: pulumi.Input<string>;
    /**
     * Specifies the IPsec protocol. Valid choices are: `ah, esp`
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * Specifies the local endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     */
    readonly tunnelLocalAddress?: pulumi.Input<string>;
    /**
     * Specifies the remote endpoint IP address of the IPsec tunnel. This parameter is only valid when mode is tunnel.
     */
    readonly tunnelRemoteAddress?: pulumi.Input<string>;
}
