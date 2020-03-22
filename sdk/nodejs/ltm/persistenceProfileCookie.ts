// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Configures a cookie persistence profile
 * 
 * ## Reference
 * 
 * `name` - (Required) Name of the virtual address
 * 
 * `defaultsFrom` - (Required) Parent cookie persistence profile
 * 
 * `matchAcrossPools` (Optional) (enabled or disabled) match across pools with given persistence record
 * 
 * `matchAcrossServices` (Optional) (enabled or disabled) match across services with given persistence record
 * 
 * `matchAcrossVirtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
 * 
 * `mirror` (Optional) (enabled or disabled) mirror persistence record
 * 
 * `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
 * 
 * `overrideConnLimit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
 * 
 * `alwaysSend` (Optional) (enabled or disabled) always send cookies
 * 
 * `cookieEncryption` (Optional) (required, preferred, or disabled) To required, preferred, or disabled policy for cookie encryption
 * 
 * `cookieEncryptionPassphrase` (Optional) (required, preferred, or disabled) Passphrase for encrypted cookies. The field is encrypted on the server and will always return differently then set.
 * If this is configured specify `ignoreChanges` under the `lifecycle` block to ignore returned encrypted value.
 * 
 * `cookieName` (Optional) Name of the cookie to track persistence
 * 
 * `expiration` (Optional) Expiration TTL for cookie specified in DAY:HOUR:MIN:SECONDS (Examples: 1:0:0:0 one day, 1:0:0 one hour, 30:0 thirty minutes)
 * 
 * `hashLength` (Optional) (Integer) Length of hash to apply to cookie
 * 
 * `hashOffset` (Optional) (Integer) Number of characters to skip in the cookie for the hash
 * 
 * `httponly` (Optional) (enabled or disabled) Sending only over http
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_persistence_profile_cookie.html.markdown.
 */
export class PersistenceProfileCookie extends pulumi.CustomResource {
    /**
     * Get an existing PersistenceProfileCookie resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PersistenceProfileCookieState, opts?: pulumi.CustomResourceOptions): PersistenceProfileCookie {
        return new PersistenceProfileCookie(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie';

    /**
     * Returns true if the given object is an instance of PersistenceProfileCookie.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PersistenceProfileCookie {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PersistenceProfileCookie.__pulumiType;
    }

    /**
     * To enable _ disable always sending cookies
     */
    public readonly alwaysSend!: pulumi.Output<string | undefined>;
    public readonly appService!: pulumi.Output<string | undefined>;
    /**
     * To required, preferred, or disabled policy for cookie encryption
     */
    public readonly cookieEncryption!: pulumi.Output<string | undefined>;
    /**
     * Passphrase for encrypted cookies
     */
    public readonly cookieEncryptionPassphrase!: pulumi.Output<string | undefined>;
    /**
     * Name of the cookie to track persistence
     */
    public readonly cookieName!: pulumi.Output<string | undefined>;
    /**
     * Inherit defaults from parent profile
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * Expiration TTL for cookie specified in D:H:M:S or in seconds
     */
    public readonly expiration!: pulumi.Output<string | undefined>;
    /**
     * Length of hash to apply to cookie
     */
    public readonly hashLength!: pulumi.Output<number | undefined>;
    /**
     * Number of characters to skip in the cookie for the hash
     */
    public readonly hashOffset!: pulumi.Output<number | undefined>;
    /**
     * To enable _ disable sending only over http
     */
    public readonly httponly!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    public readonly matchAcrossPools!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    public readonly matchAcrossServices!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across virtual servers with given persistence record
     */
    public readonly matchAcrossVirtuals!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable
     */
    public readonly mirror!: pulumi.Output<string | undefined>;
    /**
     * Name of the persistence profile
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     */
    public readonly overrideConnLimit!: pulumi.Output<string | undefined>;
    /**
     * Timeout for persistence of the session
     */
    public readonly timeout!: pulumi.Output<number | undefined>;

    /**
     * Create a PersistenceProfileCookie resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PersistenceProfileCookieArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PersistenceProfileCookieArgs | PersistenceProfileCookieState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PersistenceProfileCookieState | undefined;
            inputs["alwaysSend"] = state ? state.alwaysSend : undefined;
            inputs["appService"] = state ? state.appService : undefined;
            inputs["cookieEncryption"] = state ? state.cookieEncryption : undefined;
            inputs["cookieEncryptionPassphrase"] = state ? state.cookieEncryptionPassphrase : undefined;
            inputs["cookieName"] = state ? state.cookieName : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["expiration"] = state ? state.expiration : undefined;
            inputs["hashLength"] = state ? state.hashLength : undefined;
            inputs["hashOffset"] = state ? state.hashOffset : undefined;
            inputs["httponly"] = state ? state.httponly : undefined;
            inputs["matchAcrossPools"] = state ? state.matchAcrossPools : undefined;
            inputs["matchAcrossServices"] = state ? state.matchAcrossServices : undefined;
            inputs["matchAcrossVirtuals"] = state ? state.matchAcrossVirtuals : undefined;
            inputs["mirror"] = state ? state.mirror : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["overrideConnLimit"] = state ? state.overrideConnLimit : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as PersistenceProfileCookieArgs | undefined;
            if (!args || args.defaultsFrom === undefined) {
                throw new Error("Missing required property 'defaultsFrom'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["alwaysSend"] = args ? args.alwaysSend : undefined;
            inputs["appService"] = args ? args.appService : undefined;
            inputs["cookieEncryption"] = args ? args.cookieEncryption : undefined;
            inputs["cookieEncryptionPassphrase"] = args ? args.cookieEncryptionPassphrase : undefined;
            inputs["cookieName"] = args ? args.cookieName : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["expiration"] = args ? args.expiration : undefined;
            inputs["hashLength"] = args ? args.hashLength : undefined;
            inputs["hashOffset"] = args ? args.hashOffset : undefined;
            inputs["httponly"] = args ? args.httponly : undefined;
            inputs["matchAcrossPools"] = args ? args.matchAcrossPools : undefined;
            inputs["matchAcrossServices"] = args ? args.matchAcrossServices : undefined;
            inputs["matchAcrossVirtuals"] = args ? args.matchAcrossVirtuals : undefined;
            inputs["mirror"] = args ? args.mirror : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["overrideConnLimit"] = args ? args.overrideConnLimit : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PersistenceProfileCookie.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PersistenceProfileCookie resources.
 */
export interface PersistenceProfileCookieState {
    /**
     * To enable _ disable always sending cookies
     */
    readonly alwaysSend?: pulumi.Input<string>;
    readonly appService?: pulumi.Input<string>;
    /**
     * To required, preferred, or disabled policy for cookie encryption
     */
    readonly cookieEncryption?: pulumi.Input<string>;
    /**
     * Passphrase for encrypted cookies
     */
    readonly cookieEncryptionPassphrase?: pulumi.Input<string>;
    /**
     * Name of the cookie to track persistence
     */
    readonly cookieName?: pulumi.Input<string>;
    /**
     * Inherit defaults from parent profile
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * Expiration TTL for cookie specified in D:H:M:S or in seconds
     */
    readonly expiration?: pulumi.Input<string>;
    /**
     * Length of hash to apply to cookie
     */
    readonly hashLength?: pulumi.Input<number>;
    /**
     * Number of characters to skip in the cookie for the hash
     */
    readonly hashOffset?: pulumi.Input<number>;
    /**
     * To enable _ disable sending only over http
     */
    readonly httponly?: pulumi.Input<string>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    readonly matchAcrossPools?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossServices?: pulumi.Input<string>;
    /**
     * To enable _ disable match across virtual servers with given persistence record
     */
    readonly matchAcrossVirtuals?: pulumi.Input<string>;
    /**
     * To enable _ disable
     */
    readonly mirror?: pulumi.Input<string>;
    /**
     * Name of the persistence profile
     */
    readonly name?: pulumi.Input<string>;
    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     */
    readonly overrideConnLimit?: pulumi.Input<string>;
    /**
     * Timeout for persistence of the session
     */
    readonly timeout?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PersistenceProfileCookie resource.
 */
export interface PersistenceProfileCookieArgs {
    /**
     * To enable _ disable always sending cookies
     */
    readonly alwaysSend?: pulumi.Input<string>;
    readonly appService?: pulumi.Input<string>;
    /**
     * To required, preferred, or disabled policy for cookie encryption
     */
    readonly cookieEncryption?: pulumi.Input<string>;
    /**
     * Passphrase for encrypted cookies
     */
    readonly cookieEncryptionPassphrase?: pulumi.Input<string>;
    /**
     * Name of the cookie to track persistence
     */
    readonly cookieName?: pulumi.Input<string>;
    /**
     * Inherit defaults from parent profile
     */
    readonly defaultsFrom: pulumi.Input<string>;
    /**
     * Expiration TTL for cookie specified in D:H:M:S or in seconds
     */
    readonly expiration?: pulumi.Input<string>;
    /**
     * Length of hash to apply to cookie
     */
    readonly hashLength?: pulumi.Input<number>;
    /**
     * Number of characters to skip in the cookie for the hash
     */
    readonly hashOffset?: pulumi.Input<number>;
    /**
     * To enable _ disable sending only over http
     */
    readonly httponly?: pulumi.Input<string>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    readonly matchAcrossPools?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossServices?: pulumi.Input<string>;
    /**
     * To enable _ disable match across virtual servers with given persistence record
     */
    readonly matchAcrossVirtuals?: pulumi.Input<string>;
    /**
     * To enable _ disable
     */
    readonly mirror?: pulumi.Input<string>;
    /**
     * Name of the persistence profile
     */
    readonly name: pulumi.Input<string>;
    /**
     * To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
     * limits remain hard limits and are not overridden.
     */
    readonly overrideConnLimit?: pulumi.Input<string>;
    /**
     * Timeout for persistence of the session
     */
    readonly timeout?: pulumi.Input<number>;
}
