// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.ProfileFtp` Configures a custom profile_ftp.
 *
 * Resources should be named with their "full path". The full path is the combination of the partition + name (example: /Common/my-pool ) or  partition + directory + name of the resource  (example: /Common/test/my-pool )
 *
 * ## Example Usage
 *
 * ### For Bigip versions (14.x - 16.x)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const sanjose_ftp_profile = new f5bigip.ltm.ProfileFtp("sanjose-ftp-profile", {
 *     name: "/Common/sanjose-ftp-profile",
 *     defaultsFrom: "/Common/ftp",
 *     port: 2020,
 *     description: "test-tftp-profile",
 *     ftpsMode: "allow",
 *     enforceTlssessionReuse: "enabled",
 *     allowActiveMode: "enabled",
 * });
 * ```
 *
 * ### For Bigip versions (12.x - 13.x)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const sanjose_ftp_profile = new f5bigip.ltm.ProfileFtp("sanjose-ftp-profile", {
 *     name: "/Common/sanjose-ftp-profile",
 *     defaultsFrom: "/Common/ftp",
 *     port: 2020,
 *     description: "test-tftp-profile",
 *     allowFtps: "enabled",
 *     translateExtended: "enabled",
 * });
 * ```
 *
 * ## Common Arguments for all versions
 *
 * * `security` - (Optional)Specifies, when checked (enabled), that the system inspects FTP traffic for security vulnerabilities using an FTP security profile. This option is available only on systems licensed for BIG-IP ASM.
 *
 * * `port` - (Optional)Allows you to configure the FTP service to run on an alternate port. The default is 20.
 *
 * * `logProfile` - (Optional)Configures the ALG log profile that controls logging
 *
 * * `logPublisher` - (Optional)Configures the log publisher that handles events logging for this profile
 *
 * * `inheritParentProfile` - (Optional)Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses FastL4 only.
 *
 * * `description` - (Optional)User defined description for FTP profile
 */
export class ProfileFtp extends pulumi.CustomResource {
    /**
     * Get an existing ProfileFtp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfileFtpState, opts?: pulumi.CustomResourceOptions): ProfileFtp {
        return new ProfileFtp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/profileFtp:ProfileFtp';

    /**
     * Returns true if the given object is an instance of ProfileFtp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfileFtp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfileFtp.__pulumiType;
    }

    /**
     * Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled.
     */
    public readonly allowActiveMode!: pulumi.Output<string | undefined>;
    /**
     * Allows explicit FTPS negotiation
     */
    public readonly allowFtps!: pulumi.Output<string | undefined>;
    /**
     * The application service to which the object belongs.
     */
    public readonly appService!: pulumi.Output<string | undefined>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    public readonly defaultsFrom!: pulumi.Output<string>;
    /**
     * User defined description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default
     * value is unchecked (disabled).
     */
    public readonly enforceTlssessionReuse!: pulumi.Output<string | undefined>;
    /**
     * Allows explicit FTPS negotiation
     */
    public readonly ftpsMode!: pulumi.Output<string | undefined>;
    /**
     * Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
     * FastL4 only.
     */
    public readonly inheritParentProfile!: pulumi.Output<string | undefined>;
    /**
     * inherent vlan list
     */
    public readonly inheritVlanList!: pulumi.Output<string | undefined>;
    /**
     * Configures the ALG log profile that controls logging
     */
    public readonly logProfile!: pulumi.Output<string>;
    /**
     * Configures the log publisher that handles events logging for this profile
     */
    public readonly logPublisher!: pulumi.Output<string>;
    /**
     * Name of the profile_ftp
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    public readonly partition!: pulumi.Output<string>;
    /**
     * Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
     * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
     */
    public readonly security!: pulumi.Output<string>;
    /**
     * This setting is enabled by default, and thus, automatically translates RFC 2428 extended requests EPSV and EPRT to PASV
     * and PORT when communicating with IPv4 servers.
     */
    public readonly translateExtended!: pulumi.Output<string | undefined>;

    /**
     * Create a ProfileFtp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfileFtpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfileFtpArgs | ProfileFtpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfileFtpState | undefined;
            resourceInputs["allowActiveMode"] = state ? state.allowActiveMode : undefined;
            resourceInputs["allowFtps"] = state ? state.allowFtps : undefined;
            resourceInputs["appService"] = state ? state.appService : undefined;
            resourceInputs["defaultsFrom"] = state ? state.defaultsFrom : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enforceTlssessionReuse"] = state ? state.enforceTlssessionReuse : undefined;
            resourceInputs["ftpsMode"] = state ? state.ftpsMode : undefined;
            resourceInputs["inheritParentProfile"] = state ? state.inheritParentProfile : undefined;
            resourceInputs["inheritVlanList"] = state ? state.inheritVlanList : undefined;
            resourceInputs["logProfile"] = state ? state.logProfile : undefined;
            resourceInputs["logPublisher"] = state ? state.logPublisher : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["security"] = state ? state.security : undefined;
            resourceInputs["translateExtended"] = state ? state.translateExtended : undefined;
        } else {
            const args = argsOrState as ProfileFtpArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["allowActiveMode"] = args ? args.allowActiveMode : undefined;
            resourceInputs["allowFtps"] = args ? args.allowFtps : undefined;
            resourceInputs["appService"] = args ? args.appService : undefined;
            resourceInputs["defaultsFrom"] = args ? args.defaultsFrom : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enforceTlssessionReuse"] = args ? args.enforceTlssessionReuse : undefined;
            resourceInputs["ftpsMode"] = args ? args.ftpsMode : undefined;
            resourceInputs["inheritParentProfile"] = args ? args.inheritParentProfile : undefined;
            resourceInputs["inheritVlanList"] = args ? args.inheritVlanList : undefined;
            resourceInputs["logProfile"] = args ? args.logProfile : undefined;
            resourceInputs["logPublisher"] = args ? args.logPublisher : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["security"] = args ? args.security : undefined;
            resourceInputs["translateExtended"] = args ? args.translateExtended : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProfileFtp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfileFtp resources.
 */
export interface ProfileFtpState {
    /**
     * Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled.
     */
    allowActiveMode?: pulumi.Input<string>;
    /**
     * Allows explicit FTPS negotiation
     */
    allowFtps?: pulumi.Input<string>;
    /**
     * The application service to which the object belongs.
     */
    appService?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * User defined description
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default
     * value is unchecked (disabled).
     */
    enforceTlssessionReuse?: pulumi.Input<string>;
    /**
     * Allows explicit FTPS negotiation
     */
    ftpsMode?: pulumi.Input<string>;
    /**
     * Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
     * FastL4 only.
     */
    inheritParentProfile?: pulumi.Input<string>;
    /**
     * inherent vlan list
     */
    inheritVlanList?: pulumi.Input<string>;
    /**
     * Configures the ALG log profile that controls logging
     */
    logProfile?: pulumi.Input<string>;
    /**
     * Configures the log publisher that handles events logging for this profile
     */
    logPublisher?: pulumi.Input<string>;
    /**
     * Name of the profile_ftp
     */
    name?: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
     */
    port?: pulumi.Input<number>;
    /**
     * Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
     * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
     */
    security?: pulumi.Input<string>;
    /**
     * This setting is enabled by default, and thus, automatically translates RFC 2428 extended requests EPSV and EPRT to PASV
     * and PORT when communicating with IPv4 servers.
     */
    translateExtended?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProfileFtp resource.
 */
export interface ProfileFtpArgs {
    /**
     * Specifies, when selected (enabled), that the system allows FTP Active Transfer mode. The default value is enabled.
     */
    allowActiveMode?: pulumi.Input<string>;
    /**
     * Allows explicit FTPS negotiation
     */
    allowFtps?: pulumi.Input<string>;
    /**
     * The application service to which the object belongs.
     */
    appService?: pulumi.Input<string>;
    /**
     * Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
     */
    defaultsFrom?: pulumi.Input<string>;
    /**
     * User defined description
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies, when selected (enabled), that the system enforces the data connection to reuse a TLS session. The default
     * value is unchecked (disabled).
     */
    enforceTlssessionReuse?: pulumi.Input<string>;
    /**
     * Allows explicit FTPS negotiation
     */
    ftpsMode?: pulumi.Input<string>;
    /**
     * Enables the FTP data channel to inherit the TCP profile used by the control channel.If disabled,the data channel uses
     * FastL4 only.
     */
    inheritParentProfile?: pulumi.Input<string>;
    /**
     * inherent vlan list
     */
    inheritVlanList?: pulumi.Input<string>;
    /**
     * Configures the ALG log profile that controls logging
     */
    logProfile?: pulumi.Input<string>;
    /**
     * Configures the log publisher that handles events logging for this profile
     */
    logPublisher?: pulumi.Input<string>;
    /**
     * Name of the profile_ftp
     */
    name: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which this profile resides
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies a service for the data channel port used for this FTP profile. The default port is ftp-data.
     */
    port?: pulumi.Input<number>;
    /**
     * Enables secure FTP traffic for the BIG-IP Application Security Manager. You can set the security option only if the
     * system is licensed for the BIG-IP Application Security Manager. The default value is disabled.
     */
    security?: pulumi.Input<string>;
    /**
     * This setting is enabled by default, and thus, automatically translates RFC 2428 extended requests EPSV and EPRT to PASV
     * and PORT when communicating with IPv4 servers.
     */
    translateExtended?: pulumi.Input<string>;
}
