// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `f5bigip.sys.Snmp` provides details bout how to enable "ilx", "asm" "apm" resource on BIG-IP
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const snmp = new f5bigip.sys.Snmp("snmp", {
 *     allowedaddresses: ["202.10.10.2"],
 *     sysContact: " NetOPsAdmin s.shitole@f5.com",
 *     sysLocation: "SeattleHQ",
 * });
 * ```
 */
export class Snmp extends pulumi.CustomResource {
    /**
     * Get an existing Snmp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnmpState, opts?: pulumi.CustomResourceOptions): Snmp {
        return new Snmp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:sys/snmp:Snmp';

    /**
     * Returns true if the given object is an instance of Snmp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snmp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snmp.__pulumiType;
    }

    /**
     * Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
     */
    public readonly allowedaddresses!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the contact information for the system administrator.
     */
    public readonly sysContact!: pulumi.Output<string | undefined>;
    /**
     * Describes the system's physical location.
     */
    public readonly sysLocation!: pulumi.Output<string | undefined>;

    /**
     * Create a Snmp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnmpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnmpArgs | SnmpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnmpState | undefined;
            resourceInputs["allowedaddresses"] = state ? state.allowedaddresses : undefined;
            resourceInputs["sysContact"] = state ? state.sysContact : undefined;
            resourceInputs["sysLocation"] = state ? state.sysLocation : undefined;
        } else {
            const args = argsOrState as SnmpArgs | undefined;
            resourceInputs["allowedaddresses"] = args ? args.allowedaddresses : undefined;
            resourceInputs["sysContact"] = args ? args.sysContact : undefined;
            resourceInputs["sysLocation"] = args ? args.sysLocation : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Snmp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snmp resources.
 */
export interface SnmpState {
    /**
     * Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
     */
    allowedaddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the contact information for the system administrator.
     */
    sysContact?: pulumi.Input<string>;
    /**
     * Describes the system's physical location.
     */
    sysLocation?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snmp resource.
 */
export interface SnmpArgs {
    /**
     * Configures hosts or networks from which snmpd can accept traffic. Entries go directly into hosts.allow.
     */
    allowedaddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the contact information for the system administrator.
     */
    sysContact?: pulumi.Input<string>;
    /**
     * Describes the system's physical location.
     */
    sysLocation?: pulumi.Input<string>;
}
