// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_persistence_profile_ssl.html.markdown.
 */
export class PersistenceProfileSsl extends pulumi.CustomResource {
    /**
     * Get an existing PersistenceProfileSsl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PersistenceProfileSslState, opts?: pulumi.CustomResourceOptions): PersistenceProfileSsl {
        return new PersistenceProfileSsl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl';

    /**
     * Returns true if the given object is an instance of PersistenceProfileSsl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PersistenceProfileSsl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PersistenceProfileSsl.__pulumiType;
    }

    public readonly appService!: pulumi.Output<string | undefined>;
    /**
     * Inherit defaults from parent profile
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    public readonly matchAcrossPools!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    public readonly matchAcrossServices!: pulumi.Output<string | undefined>;
    /**
     * To enable _ disable match across services with given persistence record
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
     * Create a PersistenceProfileSsl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PersistenceProfileSslArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PersistenceProfileSslArgs | PersistenceProfileSslState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PersistenceProfileSslState | undefined;
            inputs["appService"] = state ? state.appService : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["matchAcrossPools"] = state ? state.matchAcrossPools : undefined;
            inputs["matchAcrossServices"] = state ? state.matchAcrossServices : undefined;
            inputs["matchAcrossVirtuals"] = state ? state.matchAcrossVirtuals : undefined;
            inputs["mirror"] = state ? state.mirror : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["overrideConnLimit"] = state ? state.overrideConnLimit : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
        } else {
            const args = argsOrState as PersistenceProfileSslArgs | undefined;
            if (!args || args.defaultsFrom === undefined) {
                throw new Error("Missing required property 'defaultsFrom'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["appService"] = args ? args.appService : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
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
        super(PersistenceProfileSsl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PersistenceProfileSsl resources.
 */
export interface PersistenceProfileSslState {
    readonly appService?: pulumi.Input<string>;
    /**
     * Inherit defaults from parent profile
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    readonly matchAcrossPools?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossServices?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
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
 * The set of arguments for constructing a PersistenceProfileSsl resource.
 */
export interface PersistenceProfileSslArgs {
    readonly appService?: pulumi.Input<string>;
    /**
     * Inherit defaults from parent profile
     */
    readonly defaultsFrom: pulumi.Input<string>;
    /**
     * To enable _ disable match across pools with given persistence record
     */
    readonly matchAcrossPools?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
     */
    readonly matchAcrossServices?: pulumi.Input<string>;
    /**
     * To enable _ disable match across services with given persistence record
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
