// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the bigip package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/index.html.markdown.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'f5bigip';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            if (!args || args.address === undefined) {
                throw new Error("Missing required property 'address'");
            }
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            if (!args || args.username === undefined) {
                throw new Error("Missing required property 'username'");
            }
            inputs["address"] = args ? args.address : undefined;
            inputs["loginRef"] = args ? args.loginRef : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["tokenAuth"] = pulumi.output(args ? args.tokenAuth : undefined).apply(JSON.stringify);
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Domain name/IP of the BigIP
     */
    readonly address: pulumi.Input<string>;
    /**
     * Login reference for token authentication (see BIG-IP REST docs for details)
     */
    readonly loginRef?: pulumi.Input<string>;
    /**
     * The user's password
     */
    readonly password: pulumi.Input<string>;
    /**
     * Enable to use an external authentication source (LDAP, TACACS, etc)
     */
    readonly tokenAuth?: pulumi.Input<boolean>;
    /**
     * Username with API access to the BigIP
     */
    readonly username: pulumi.Input<string>;
}
