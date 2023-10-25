// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.IpsecProfile` Manage IPSec Profiles on a BIG-IP
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const azurevWANProfile = new f5bigip.IpsecProfile("azurevWANProfile", {
 *     description: "mytestipsecprofile",
 *     name: "/Common/Mytestipsecprofile",
 *     trafficSelector: "test-trafficselector",
 * });
 * ```
 */
export class IpsecProfile extends pulumi.CustomResource {
    /**
     * Get an existing IpsecProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpsecProfileState, opts?: pulumi.CustomResourceOptions): IpsecProfile {
        return new IpsecProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/ipsecProfile:IpsecProfile';

    /**
     * Returns true if the given object is an instance of IpsecProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpsecProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpsecProfile.__pulumiType;
    }

    /**
     * Specifies descriptive text that identifies the IPsec interface tunnel profile.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
     */
    public readonly parentProfile!: pulumi.Output<string | undefined>;
    /**
     * Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
     */
    public readonly trafficSelector!: pulumi.Output<string>;

    /**
     * Create a IpsecProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpsecProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpsecProfileArgs | IpsecProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpsecProfileState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentProfile"] = state ? state.parentProfile : undefined;
            resourceInputs["trafficSelector"] = state ? state.trafficSelector : undefined;
        } else {
            const args = argsOrState as IpsecProfileArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentProfile"] = args ? args.parentProfile : undefined;
            resourceInputs["trafficSelector"] = args ? args.trafficSelector : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpsecProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpsecProfile resources.
 */
export interface IpsecProfileState {
    /**
     * Specifies descriptive text that identifies the IPsec interface tunnel profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
     */
    parentProfile?: pulumi.Input<string>;
    /**
     * Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
     */
    trafficSelector?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpsecProfile resource.
 */
export interface IpsecProfileArgs {
    /**
     * Specifies descriptive text that identifies the IPsec interface tunnel profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
     */
    name: pulumi.Input<string>;
    /**
     * Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
     */
    parentProfile?: pulumi.Input<string>;
    /**
     * Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
     */
    trafficSelector?: pulumi.Input<string>;
}
