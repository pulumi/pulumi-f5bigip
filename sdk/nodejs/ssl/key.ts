// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ssl.Key` This resource will import SSL certificate key on BIG-IP LTM. 
 * Certificate key can be imported from certificate key files on the local disk, in PEM format
 * 
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ssl_key.html.markdown.
 */
export class Key extends pulumi.CustomResource {
    /**
     * Get an existing Key resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyState, opts?: pulumi.CustomResourceOptions): Key {
        return new Key(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ssl/key:Key';

    /**
     * Returns true if the given object is an instance of Key.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Key {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Key.__pulumiType;
    }

    /**
     * Content of SSL certificate key present on local Disk
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Name of SSL Certificate key with .key extension
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Partition on to SSL Certificate key to be imported
     */
    public readonly partition!: pulumi.Output<string | undefined>;

    /**
     * Create a Key resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyArgs | KeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as KeyState | undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partition"] = state ? state.partition : undefined;
        } else {
            const args = argsOrState as KeyArgs | undefined;
            if (!args || args.content === undefined) {
                throw new Error("Missing required property 'content'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["content"] = args ? args.content : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partition"] = args ? args.partition : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Key.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Key resources.
 */
export interface KeyState {
    /**
     * Content of SSL certificate key present on local Disk
     */
    readonly content?: pulumi.Input<string>;
    /**
     * Name of SSL Certificate key with .key extension
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Partition on to SSL Certificate key to be imported
     */
    readonly partition?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Key resource.
 */
export interface KeyArgs {
    /**
     * Content of SSL certificate key present on local Disk
     */
    readonly content: pulumi.Input<string>;
    /**
     * Name of SSL Certificate key with .key extension
     */
    readonly name: pulumi.Input<string>;
    /**
     * Partition on to SSL Certificate key to be imported
     */
    readonly partition?: pulumi.Input<string>;
}
