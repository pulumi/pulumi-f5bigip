// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class SaasBotDefenseProfile extends pulumi.CustomResource {
    /**
     * Get an existing SaasBotDefenseProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SaasBotDefenseProfileState, opts?: pulumi.CustomResourceOptions): SaasBotDefenseProfile {
        return new SaasBotDefenseProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/saasBotDefenseProfile:SaasBotDefenseProfile';

    /**
     * Returns true if the given object is an instance of SaasBotDefenseProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SaasBotDefenseProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SaasBotDefenseProfile.__pulumiType;
    }

    /**
     * Specifies the API key, enter the value provided by F5 Support.
     */
    public readonly apiKey!: pulumi.Output<string>;
    /**
     * Specifies the Bot Defense API application ID, enter the value provided by F5 Support
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * Distributed Cloud Services Bot Defense parent profile from which this profile will inherit settings.
     */
    public readonly defaultsFrom!: pulumi.Output<string | undefined>;
    /**
     * Specifies descriptive text that identifies the BD profile.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Unique name for the Distributed Cloud Services Bot Defense profile
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Use these settings to configure which pages on the website will be protected by BD
     */
    public readonly protectedEndpoints!: pulumi.Output<outputs.SaasBotDefenseProfileProtectedEndpoint[]>;
    /**
     * Specifies the web hostname to which API requests are made
     */
    public readonly shapeProtectionPool!: pulumi.Output<string>;
    /**
     * Specifies a server-side SSL profile that is different from what the application pool uses
     */
    public readonly sslProfile!: pulumi.Output<string>;
    /**
     * Specifies the tenant ID, enter the value provided by F5 Support
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a SaasBotDefenseProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SaasBotDefenseProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SaasBotDefenseProfileArgs | SaasBotDefenseProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SaasBotDefenseProfileState | undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
            resourceInputs["applicationId"] = state ? state.applicationId : undefined;
            resourceInputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protectedEndpoints"] = state ? state.protectedEndpoints : undefined;
            resourceInputs["shapeProtectionPool"] = state ? state.shapeProtectionPool : undefined;
            resourceInputs["sslProfile"] = state ? state.sslProfile : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as SaasBotDefenseProfileArgs | undefined;
            if ((!args || args.apiKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiKey'");
            }
            if ((!args || args.applicationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationId'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.protectedEndpoints === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protectedEndpoints'");
            }
            if ((!args || args.shapeProtectionPool === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shapeProtectionPool'");
            }
            if ((!args || args.sslProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sslProfile'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["apiKey"] = args?.apiKey ? pulumi.secret(args.apiKey) : undefined;
            resourceInputs["applicationId"] = args?.applicationId ? pulumi.secret(args.applicationId) : undefined;
            resourceInputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protectedEndpoints"] = args ? args.protectedEndpoints : undefined;
            resourceInputs["shapeProtectionPool"] = args ? args.shapeProtectionPool : undefined;
            resourceInputs["sslProfile"] = args ? args.sslProfile : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiKey", "applicationId"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SaasBotDefenseProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SaasBotDefenseProfile resources.
 */
export interface SaasBotDefenseProfileState {
    /**
     * Specifies the API key, enter the value provided by F5 Support.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * Specifies the Bot Defense API application ID, enter the value provided by F5 Support
     */
    applicationId?: pulumi.Input<string>;
    /**
     * Distributed Cloud Services Bot Defense parent profile from which this profile will inherit settings.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies descriptive text that identifies the BD profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique name for the Distributed Cloud Services Bot Defense profile
     */
    name?: pulumi.Input<string>;
    /**
     * Use these settings to configure which pages on the website will be protected by BD
     */
    protectedEndpoints?: pulumi.Input<pulumi.Input<inputs.SaasBotDefenseProfileProtectedEndpoint>[]>;
    /**
     * Specifies the web hostname to which API requests are made
     */
    shapeProtectionPool?: pulumi.Input<string>;
    /**
     * Specifies a server-side SSL profile that is different from what the application pool uses
     */
    sslProfile?: pulumi.Input<string>;
    /**
     * Specifies the tenant ID, enter the value provided by F5 Support
     */
    tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SaasBotDefenseProfile resource.
 */
export interface SaasBotDefenseProfileArgs {
    /**
     * Specifies the API key, enter the value provided by F5 Support.
     */
    apiKey: pulumi.Input<string>;
    /**
     * Specifies the Bot Defense API application ID, enter the value provided by F5 Support
     */
    applicationId: pulumi.Input<string>;
    /**
     * Distributed Cloud Services Bot Defense parent profile from which this profile will inherit settings.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * Specifies descriptive text that identifies the BD profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique name for the Distributed Cloud Services Bot Defense profile
     */
    name: pulumi.Input<string>;
    /**
     * Use these settings to configure which pages on the website will be protected by BD
     */
    protectedEndpoints: pulumi.Input<pulumi.Input<inputs.SaasBotDefenseProfileProtectedEndpoint>[]>;
    /**
     * Specifies the web hostname to which API requests are made
     */
    shapeProtectionPool: pulumi.Input<string>;
    /**
     * Specifies a server-side SSL profile that is different from what the application pool uses
     */
    sslProfile: pulumi.Input<string>;
    /**
     * Specifies the tenant ID, enter the value provided by F5 Support
     */
    tenantId: pulumi.Input<string>;
}
