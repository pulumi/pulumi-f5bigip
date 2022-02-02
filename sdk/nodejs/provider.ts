// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the bigip package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
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
     * Domain name/IP of the BigIP
     */
    public readonly address!: pulumi.Output<string | undefined>;
    /**
     * Login reference for token authentication (see BIG-IP REST docs for details)
     */
    public readonly loginRef!: pulumi.Output<string | undefined>;
    /**
     * The user's password. Leave empty if using token_value
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Management Port to connect to Bigip
     */
    public readonly port!: pulumi.Output<string | undefined>;
    /**
     * A token generated outside the provider, in place of password
     */
    public readonly tokenValue!: pulumi.Output<string | undefined>;
    /**
     * Username with API access to the BigIP
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["loginRef"] = args ? args.loginRef : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["teemDisable"] = pulumi.output(args ? args.teemDisable : undefined).apply(JSON.stringify);
            resourceInputs["tokenAuth"] = pulumi.output(args ? args.tokenAuth : undefined).apply(JSON.stringify);
            resourceInputs["tokenValue"] = args ? args.tokenValue : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Domain name/IP of the BigIP
     */
    address?: pulumi.Input<string>;
    /**
     * Login reference for token authentication (see BIG-IP REST docs for details)
     */
    loginRef?: pulumi.Input<string>;
    /**
     * The user's password. Leave empty if using token_value
     */
    password?: pulumi.Input<string>;
    /**
     * Management Port to connect to Bigip
     */
    port?: pulumi.Input<string>;
    /**
     * If this flag set to true,sending telemetry data to TEEM will be disabled
     */
    teemDisable?: pulumi.Input<boolean>;
    /**
     * Enable to use an external authentication source (LDAP, TACACS, etc)
     */
    tokenAuth?: pulumi.Input<boolean>;
    /**
     * A token generated outside the provider, in place of password
     */
    tokenValue?: pulumi.Input<string>;
    /**
     * Username with API access to the BigIP
     */
    username?: pulumi.Input<string>;
}
