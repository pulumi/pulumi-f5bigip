// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileHttp` Configures a custom profileHttp for use by health checks.
 *
 * For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const sanjose_http = new f5bigip.ltm.ProfileHttp("sanjose-http", {
 *     defaultsFrom: "/Common/http",
 *     description: "some http",
 *     fallbackHost: "titanic",
 *     fallbackStatusCodes: [
 *         "400",
 *         "500",
 *         "300",
 *     ],
 *     name: "/Common/sanjose-http",
 * });
 * ```
 *
 * ## Import
 *
 * BIG-IP LTM http profiles can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import f5bigip:ltm/profileHttp:ProfileHttp test-http /Common/test-http
 * ```
 */
export class ProfileHttp extends pulumi.CustomResource {
    /**
     * Get an existing ProfileHttp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileHttpState, opts?: pulumi.CustomResourceOptions): ProfileHttp {
        return new ProfileHttp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileHttp:ProfileHttp';

    /**
     * Returns true if the given object is an instance of ProfileHttp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileHttp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileHttp.__pulumiType;
    }

    /**
     * Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
     * XFF (X-forwarded-for) headers, if they exist.
     */
    public readonly acceptXff!: pulumi.Output<string>;
    /**
     * The application service to which the object belongs.
     */
    public readonly appService!: pulumi.Output<string | undefined>;
    /**
     * Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
     */
    public readonly basicAuthRealm!: pulumi.Output<string>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * User defined description
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Specifies a passphrase for the cookie encryption
     */
    public readonly encryptCookieSecret!: pulumi.Output<string | undefined>;
    /**
     * Encrypts specified cookies that the BIG-IP system sends to a client system
     */
    public readonly encryptCookies!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
     */
    public readonly fallbackHost!: pulumi.Output<string>;
    /**
     * Specifies one or more three-digit status codes that can be returned by an HTTP server.
     */
    public readonly fallbackStatusCodes!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the header string that you want to erase from an HTTP request. You can also specify none
     */
    public readonly headErase!: pulumi.Output<string>;
    /**
     * Specifies a quoted header string that you want to insert into an HTTP request
     */
    public readonly headInsert!: pulumi.Output<string>;
    /**
     * When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
     */
    public readonly insertXforwardedFor!: pulumi.Output<string>;
    /**
     * Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
     */
    public readonly lwsSeparator!: pulumi.Output<string>;
    /**
     * Name of the profile_http
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
     */
    public readonly oneconnectTransformations!: pulumi.Output<string>;
    /**
     * Specifies the type of HTTP proxy.
     */
    public readonly proxyType!: pulumi.Output<string>;
    /**
     * Specifies which of the application HTTP redirects the system rewrites to HTTPS.
     */
    public readonly redirectRewrite!: pulumi.Output<string>;
    /**
     * Specifies how to handle chunked and unchunked requests.
     */
    public readonly requestChunking!: pulumi.Output<string>;
    /**
     * Specifies how to handle chunked and unchunked responses.
     */
    public readonly responseChunking!: pulumi.Output<string>;
    /**
     * Specifies headers that the BIG-IP system allows in an HTTP response.
     */
    public readonly responseHeadersPermitteds!: pulumi.Output<string[]>;
    /**
     * Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
     * string is specified, then no Server header will be added to such responses
     */
    public readonly serverAgentName!: pulumi.Output<string>;
    /**
     * Displays the administrative partition within which this profile resides.
     */
    public readonly tmPartition!: pulumi.Output<string | undefined>;
    /**
     * Specifies the hostname to include into Via header
     */
    public readonly viaHostName!: pulumi.Output<string>;
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     */
    public readonly viaRequest!: pulumi.Output<string>;
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     */
    public readonly viaResponse!: pulumi.Output<string>;
    /**
     * Specifies alternative XFF headers instead of the default X-forwarded-for header
     */
    public readonly xffAlternativeNames!: pulumi.Output<string[]>;

    /**
     * Create a ProfileHttp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileHttpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileHttpArgs | ProfileHttpState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProfileHttpState | undefined;
            inputs["acceptXff"] = state ? state.acceptXff : undefined;
            inputs["appService"] = state ? state.appService : undefined;
            inputs["basicAuthRealm"] = state ? state.basicAuthRealm : undefined;
            inputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["encryptCookieSecret"] = state ? state.encryptCookieSecret : undefined;
            inputs["encryptCookies"] = state ? state.encryptCookies : undefined;
            inputs["fallbackHost"] = state ? state.fallbackHost : undefined;
            inputs["fallbackStatusCodes"] = state ? state.fallbackStatusCodes : undefined;
            inputs["headErase"] = state ? state.headErase : undefined;
            inputs["headInsert"] = state ? state.headInsert : undefined;
            inputs["insertXforwardedFor"] = state ? state.insertXforwardedFor : undefined;
            inputs["lwsSeparator"] = state ? state.lwsSeparator : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["oneconnectTransformations"] = state ? state.oneconnectTransformations : undefined;
            inputs["proxyType"] = state ? state.proxyType : undefined;
            inputs["redirectRewrite"] = state ? state.redirectRewrite : undefined;
            inputs["requestChunking"] = state ? state.requestChunking : undefined;
            inputs["responseChunking"] = state ? state.responseChunking : undefined;
            inputs["responseHeadersPermitteds"] = state ? state.responseHeadersPermitteds : undefined;
            inputs["serverAgentName"] = state ? state.serverAgentName : undefined;
            inputs["tmPartition"] = state ? state.tmPartition : undefined;
            inputs["viaHostName"] = state ? state.viaHostName : undefined;
            inputs["viaRequest"] = state ? state.viaRequest : undefined;
            inputs["viaResponse"] = state ? state.viaResponse : undefined;
            inputs["xffAlternativeNames"] = state ? state.xffAlternativeNames : undefined;
        } else {
            const args = argsOrState as ProfileHttpArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["acceptXff"] = args ? args.acceptXff : undefined;
            inputs["appService"] = args ? args.appService : undefined;
            inputs["basicAuthRealm"] = args ? args.basicAuthRealm : undefined;
            inputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["encryptCookieSecret"] = args ? args.encryptCookieSecret : undefined;
            inputs["encryptCookies"] = args ? args.encryptCookies : undefined;
            inputs["fallbackHost"] = args ? args.fallbackHost : undefined;
            inputs["fallbackStatusCodes"] = args ? args.fallbackStatusCodes : undefined;
            inputs["headErase"] = args ? args.headErase : undefined;
            inputs["headInsert"] = args ? args.headInsert : undefined;
            inputs["insertXforwardedFor"] = args ? args.insertXforwardedFor : undefined;
            inputs["lwsSeparator"] = args ? args.lwsSeparator : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["oneconnectTransformations"] = args ? args.oneconnectTransformations : undefined;
            inputs["proxyType"] = args ? args.proxyType : undefined;
            inputs["redirectRewrite"] = args ? args.redirectRewrite : undefined;
            inputs["requestChunking"] = args ? args.requestChunking : undefined;
            inputs["responseChunking"] = args ? args.responseChunking : undefined;
            inputs["responseHeadersPermitteds"] = args ? args.responseHeadersPermitteds : undefined;
            inputs["serverAgentName"] = args ? args.serverAgentName : undefined;
            inputs["tmPartition"] = args ? args.tmPartition : undefined;
            inputs["viaHostName"] = args ? args.viaHostName : undefined;
            inputs["viaRequest"] = args ? args.viaRequest : undefined;
            inputs["viaResponse"] = args ? args.viaResponse : undefined;
            inputs["xffAlternativeNames"] = args ? args.xffAlternativeNames : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProfileHttp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileHttp resources.
 */
export interface ProfileHttpState {
    /**
     * Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
     * XFF (X-forwarded-for) headers, if they exist.
     */
    readonly acceptXff?: pulumi.Input<string>;
    /**
     * The application service to which the object belongs.
     */
    readonly appService?: pulumi.Input<string>;
    /**
     * Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
     */
    readonly basicAuthRealm?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * User defined description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies a passphrase for the cookie encryption
     */
    readonly encryptCookieSecret?: pulumi.Input<string>;
    /**
     * Encrypts specified cookies that the BIG-IP system sends to a client system
     */
    readonly encryptCookies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
     */
    readonly fallbackHost?: pulumi.Input<string>;
    /**
     * Specifies one or more three-digit status codes that can be returned by an HTTP server.
     */
    readonly fallbackStatusCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the header string that you want to erase from an HTTP request. You can also specify none
     */
    readonly headErase?: pulumi.Input<string>;
    /**
     * Specifies a quoted header string that you want to insert into an HTTP request
     */
    readonly headInsert?: pulumi.Input<string>;
    /**
     * When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
     */
    readonly insertXforwardedFor?: pulumi.Input<string>;
    /**
     * Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
     */
    readonly lwsSeparator?: pulumi.Input<string>;
    /**
     * Name of the profile_http
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
     */
    readonly oneconnectTransformations?: pulumi.Input<string>;
    /**
     * Specifies the type of HTTP proxy.
     */
    readonly proxyType?: pulumi.Input<string>;
    /**
     * Specifies which of the application HTTP redirects the system rewrites to HTTPS.
     */
    readonly redirectRewrite?: pulumi.Input<string>;
    /**
     * Specifies how to handle chunked and unchunked requests.
     */
    readonly requestChunking?: pulumi.Input<string>;
    /**
     * Specifies how to handle chunked and unchunked responses.
     */
    readonly responseChunking?: pulumi.Input<string>;
    /**
     * Specifies headers that the BIG-IP system allows in an HTTP response.
     */
    readonly responseHeadersPermitteds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
     * string is specified, then no Server header will be added to such responses
     */
    readonly serverAgentName?: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides.
     */
    readonly tmPartition?: pulumi.Input<string>;
    /**
     * Specifies the hostname to include into Via header
     */
    readonly viaHostName?: pulumi.Input<string>;
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     */
    readonly viaRequest?: pulumi.Input<string>;
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     */
    readonly viaResponse?: pulumi.Input<string>;
    /**
     * Specifies alternative XFF headers instead of the default X-forwarded-for header
     */
    readonly xffAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ProfileHttp resource.
 */
export interface ProfileHttpArgs {
    /**
     * Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
     * XFF (X-forwarded-for) headers, if they exist.
     */
    readonly acceptXff?: pulumi.Input<string>;
    /**
     * The application service to which the object belongs.
     */
    readonly appService?: pulumi.Input<string>;
    /**
     * Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
     */
    readonly basicAuthRealm?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    readonly defaultsFrom?: pulumi.Input<string>;
    /**
     * User defined description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies a passphrase for the cookie encryption
     */
    readonly encryptCookieSecret?: pulumi.Input<string>;
    /**
     * Encrypts specified cookies that the BIG-IP system sends to a client system
     */
    readonly encryptCookies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
     */
    readonly fallbackHost?: pulumi.Input<string>;
    /**
     * Specifies one or more three-digit status codes that can be returned by an HTTP server.
     */
    readonly fallbackStatusCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the header string that you want to erase from an HTTP request. You can also specify none
     */
    readonly headErase?: pulumi.Input<string>;
    /**
     * Specifies a quoted header string that you want to insert into an HTTP request
     */
    readonly headInsert?: pulumi.Input<string>;
    /**
     * When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
     */
    readonly insertXforwardedFor?: pulumi.Input<string>;
    /**
     * Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
     */
    readonly lwsSeparator?: pulumi.Input<string>;
    /**
     * Name of the profile_http
     */
    readonly name: pulumi.Input<string>;
    /**
     * Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
     */
    readonly oneconnectTransformations?: pulumi.Input<string>;
    /**
     * Specifies the type of HTTP proxy.
     */
    readonly proxyType?: pulumi.Input<string>;
    /**
     * Specifies which of the application HTTP redirects the system rewrites to HTTPS.
     */
    readonly redirectRewrite?: pulumi.Input<string>;
    /**
     * Specifies how to handle chunked and unchunked requests.
     */
    readonly requestChunking?: pulumi.Input<string>;
    /**
     * Specifies how to handle chunked and unchunked responses.
     */
    readonly responseChunking?: pulumi.Input<string>;
    /**
     * Specifies headers that the BIG-IP system allows in an HTTP response.
     */
    readonly responseHeadersPermitteds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
     * string is specified, then no Server header will be added to such responses
     */
    readonly serverAgentName?: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides.
     */
    readonly tmPartition?: pulumi.Input<string>;
    /**
     * Specifies the hostname to include into Via header
     */
    readonly viaHostName?: pulumi.Input<string>;
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     */
    readonly viaRequest?: pulumi.Input<string>;
    /**
     * Specifies whether to append, remove, or preserve a Via header in an HTTP request
     */
    readonly viaResponse?: pulumi.Input<string>;
    /**
     * Specifies alternative XFF headers instead of the default X-forwarded-for header
     */
    readonly xffAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
}
