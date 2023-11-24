// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.CipherGroup` Manages F5 BIG-IP LTM cipher group using iControl REST.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const test_cipher_group = new f5bigip.ltm.CipherGroup("test-cipher-group", {
 *     allows: ["/Common/f5-aes"],
 *     name: "/Common/test-cipher-group-01",
 *     ordering: "speed",
 *     requires: ["/Common/f5-quic"],
 * });
 * ```
 */
export class CipherGroup extends pulumi.CustomResource {
    /**
     * Get an existing CipherGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CipherGroupState, opts?: pulumi.CustomResourceOptions): CipherGroup {
        return new CipherGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/cipherGroup:CipherGroup';

    /**
     * Returns true if the given object is an instance of CipherGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CipherGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CipherGroup.__pulumiType;
    }

    /**
     * Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     */
    public readonly allows!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies descriptive text that identifies the cipher rule
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the Cipher group. Name should be in pattern `partition` + `cipherGroupName`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
     */
    public readonly ordering!: pulumi.Output<string | undefined>;
    /**
     * Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     */
    public readonly requires!: pulumi.Output<string[] | undefined>;

    /**
     * Create a CipherGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CipherGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CipherGroupArgs | CipherGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CipherGroupState | undefined;
            resourceInputs["allows"] = state ? state.allows : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ordering"] = state ? state.ordering : undefined;
            resourceInputs["requires"] = state ? state.requires : undefined;
        } else {
            const args = argsOrState as CipherGroupArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["allows"] = args ? args.allows : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ordering"] = args ? args.ordering : undefined;
            resourceInputs["requires"] = args ? args.requires : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CipherGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CipherGroup resources.
 */
export interface CipherGroupState {
    /**
     * Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     */
    allows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies descriptive text that identifies the cipher rule
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Cipher group. Name should be in pattern `partition` + `cipherGroupName`
     */
    name?: pulumi.Input<string>;
    /**
     * Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
     */
    ordering?: pulumi.Input<string>;
    /**
     * Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     */
    requires?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a CipherGroup resource.
 */
export interface CipherGroupArgs {
    /**
     * Specifies the configuration of the allowed groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no allowed ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     */
    allows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies descriptive text that identifies the cipher rule
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Cipher group. Name should be in pattern `partition` + `cipherGroupName`
     */
    name: pulumi.Input<string>;
    /**
     * Controls the order of the Cipher String list in the Cipher Audit section. Options are Default, Speed, Strength, FIPS, and Hardware. The rules are processed in the order listed. The default is `default`.
     */
    ordering?: pulumi.Input<string>;
    /**
     * Specifies the configuration of the restrict groups of ciphers. You can select a cipher rule from the Available Cipher Rules list. To have no restricted ciphers, omit this attribute in the config or set it to an empty set like, `[]`.
     */
    requires?: pulumi.Input<pulumi.Input<string>[]>;
}
