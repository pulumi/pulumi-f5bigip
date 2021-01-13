// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.CommonLicenseManageBigIq` This Resource is used for BIGIP/Provider License Management from BIGIQ
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * // MANAGED Regkey Pool
 * const testExampleCommonLicenseManageBigIq = new f5bigip.CommonLicenseManageBigIq("testExampleCommonLicenseManageBigIq", {
 *     bigiqAddress: _var.bigiq,
 *     bigiqUser: _var.bigiq_un,
 *     bigiqPassword: _var.bigiq_pw,
 *     licensePoolname: "regkeypool_name",
 *     assignmentType: "MANAGED",
 * });
 * // UNMANAGED Regkey Pool
 * const testExampleIndex_commonLicenseManageBigIqCommonLicenseManageBigIq = new f5bigip.CommonLicenseManageBigIq("testExampleIndex/commonLicenseManageBigIqCommonLicenseManageBigIq", {
 *     bigiqAddress: _var.bigiq,
 *     bigiqUser: _var.bigiq_un,
 *     bigiqPassword: _var.bigiq_pw,
 *     licensePoolname: "regkeypool_name",
 *     assignmentType: "UNMANAGED",
 * });
 * // UNMANAGED Utility Pool
 * const testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq = new f5bigip.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq", {
 *     bigiqAddress: _var.bigiq,
 *     bigiqUser: _var.bigiq_un,
 *     bigiqPassword: _var.bigiq_pw,
 *     licensePoolname: "utilitypool_name",
 *     assignmentType: "UNMANAGED",
 *     unitOfMeasure: "yearly",
 *     skukeyword1: "BTHSM200M",
 * });
 * // UNREACHABLE Regkey Pool
 * const testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq1 = new f5bigip.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq1", {
 *     bigiqAddress: "xxx.xxx.xxx.xxx",
 *     bigiqUser: "xxxx",
 *     bigiqPassword: "xxxxx",
 *     licensePoolname: "regkey_pool_name",
 *     assignmentType: "UNREACHABLE",
 *     macAddress: "FA:16:3E:1B:6D:32",
 *     hypervisor: "azure",
 * });
 * // MANAGED Purchased Pool
 * const testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq2 = new f5bigip.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq2", {
 *     bigiqAddress: _var.bigiq,
 *     bigiqUser: _var.bigiq_un,
 *     bigiqPassword: _var.bigiq_pw,
 *     licensePoolname: "purchased_pool_name",
 *     assignmentType: "MANAGED",
 * });
 * // UNMANAGED Purchased Pool
 * const testExampleF5bigipIndex_commonLicenseManageBigIqCommonLicenseManageBigIq3 = new f5bigip.CommonLicenseManageBigIq("testExampleF5bigipIndex/commonLicenseManageBigIqCommonLicenseManageBigIq3", {
 *     bigiqAddress: _var.bigiq,
 *     bigiqUser: _var.bigiq_un,
 *     bigiqPassword: _var.bigiq_pw,
 *     licensePoolname: "purchased_pool_name",
 *     assignmentType: "UNMANAGED",
 * });
 * ```
 */
export class CommonLicenseManageBigIq extends pulumi.CustomResource {
    /**
     * Get an existing CommonLicenseManageBigIq resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CommonLicenseManageBigIqState, opts?: pulumi.CustomResourceOptions): CommonLicenseManageBigIq {
        return new CommonLicenseManageBigIq(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/commonLicenseManageBigIq:CommonLicenseManageBigIq';

    /**
     * Returns true if the given object is an instance of CommonLicenseManageBigIq.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CommonLicenseManageBigIq {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CommonLicenseManageBigIq.__pulumiType;
    }

    /**
     * The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
     */
    public readonly assignmentType!: pulumi.Output<string>;
    /**
     * BIGIQ License Manager IP Address, variable type `string`
     */
    public readonly bigiqAddress!: pulumi.Output<string>;
    /**
     * BIGIQ Login reference for token authentication
     */
    public readonly bigiqLoginRef!: pulumi.Output<string | undefined>;
    /**
     * BIGIQ License Manager password.  variable type `string`
     */
    public readonly bigiqPassword!: pulumi.Output<string>;
    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     */
    public readonly bigiqPort!: pulumi.Output<string | undefined>;
    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     */
    public readonly bigiqTokenAuth!: pulumi.Output<boolean | undefined>;
    /**
     * BIGIQ License Manager username, variable type `string`
     */
    public readonly bigiqUser!: pulumi.Output<string>;
    /**
     * Status of Licence Assignment
     */
    public readonly deviceLicenseStatus!: pulumi.Output<string>;
    /**
     * Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
     */
    public readonly hypervisor!: pulumi.Output<string | undefined>;
    /**
     * License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * A name given to the license pool. type `string`
     */
    public readonly licensePoolname!: pulumi.Output<string>;
    /**
     * MAC address of the BIG-IP. type `string`
     */
    public readonly macAddress!: pulumi.Output<string | undefined>;
    /**
     * An optional offering name. type `string`
     */
    public readonly skukeyword1!: pulumi.Output<string | undefined>;
    /**
     * An optional offering name. type `string`
     */
    public readonly skukeyword2!: pulumi.Output<string | undefined>;
    /**
     * For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
     */
    public readonly tenant!: pulumi.Output<string | undefined>;
    /**
     * The units used to measure billing. For example, “hourly” or “daily”. Type `string`
     */
    public readonly unitOfMeasure!: pulumi.Output<string | undefined>;

    /**
     * Create a CommonLicenseManageBigIq resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CommonLicenseManageBigIqArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CommonLicenseManageBigIqArgs | CommonLicenseManageBigIqState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CommonLicenseManageBigIqState | undefined;
            inputs["assignmentType"] = state ? state.assignmentType : undefined;
            inputs["bigiqAddress"] = state ? state.bigiqAddress : undefined;
            inputs["bigiqLoginRef"] = state ? state.bigiqLoginRef : undefined;
            inputs["bigiqPassword"] = state ? state.bigiqPassword : undefined;
            inputs["bigiqPort"] = state ? state.bigiqPort : undefined;
            inputs["bigiqTokenAuth"] = state ? state.bigiqTokenAuth : undefined;
            inputs["bigiqUser"] = state ? state.bigiqUser : undefined;
            inputs["deviceLicenseStatus"] = state ? state.deviceLicenseStatus : undefined;
            inputs["hypervisor"] = state ? state.hypervisor : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["licensePoolname"] = state ? state.licensePoolname : undefined;
            inputs["macAddress"] = state ? state.macAddress : undefined;
            inputs["skukeyword1"] = state ? state.skukeyword1 : undefined;
            inputs["skukeyword2"] = state ? state.skukeyword2 : undefined;
            inputs["tenant"] = state ? state.tenant : undefined;
            inputs["unitOfMeasure"] = state ? state.unitOfMeasure : undefined;
        } else {
            const args = argsOrState as CommonLicenseManageBigIqArgs | undefined;
            if ((!args || args.assignmentType === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'assignmentType'");
            }
            if ((!args || args.bigiqAddress === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'bigiqAddress'");
            }
            if ((!args || args.bigiqPassword === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'bigiqPassword'");
            }
            if ((!args || args.bigiqUser === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'bigiqUser'");
            }
            if ((!args || args.licensePoolname === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'licensePoolname'");
            }
            inputs["assignmentType"] = args ? args.assignmentType : undefined;
            inputs["bigiqAddress"] = args ? args.bigiqAddress : undefined;
            inputs["bigiqLoginRef"] = args ? args.bigiqLoginRef : undefined;
            inputs["bigiqPassword"] = args ? args.bigiqPassword : undefined;
            inputs["bigiqPort"] = args ? args.bigiqPort : undefined;
            inputs["bigiqTokenAuth"] = args ? args.bigiqTokenAuth : undefined;
            inputs["bigiqUser"] = args ? args.bigiqUser : undefined;
            inputs["deviceLicenseStatus"] = args ? args.deviceLicenseStatus : undefined;
            inputs["hypervisor"] = args ? args.hypervisor : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["licensePoolname"] = args ? args.licensePoolname : undefined;
            inputs["macAddress"] = args ? args.macAddress : undefined;
            inputs["skukeyword1"] = args ? args.skukeyword1 : undefined;
            inputs["skukeyword2"] = args ? args.skukeyword2 : undefined;
            inputs["tenant"] = args ? args.tenant : undefined;
            inputs["unitOfMeasure"] = args ? args.unitOfMeasure : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CommonLicenseManageBigIq.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CommonLicenseManageBigIq resources.
 */
export interface CommonLicenseManageBigIqState {
    /**
     * The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
     */
    readonly assignmentType?: pulumi.Input<string>;
    /**
     * BIGIQ License Manager IP Address, variable type `string`
     */
    readonly bigiqAddress?: pulumi.Input<string>;
    /**
     * BIGIQ Login reference for token authentication
     */
    readonly bigiqLoginRef?: pulumi.Input<string>;
    /**
     * BIGIQ License Manager password.  variable type `string`
     */
    readonly bigiqPassword?: pulumi.Input<string>;
    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     */
    readonly bigiqPort?: pulumi.Input<string>;
    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     */
    readonly bigiqTokenAuth?: pulumi.Input<boolean>;
    /**
     * BIGIQ License Manager username, variable type `string`
     */
    readonly bigiqUser?: pulumi.Input<string>;
    /**
     * Status of Licence Assignment
     */
    readonly deviceLicenseStatus?: pulumi.Input<string>;
    /**
     * Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
     */
    readonly hypervisor?: pulumi.Input<string>;
    /**
     * License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
     */
    readonly key?: pulumi.Input<string>;
    /**
     * A name given to the license pool. type `string`
     */
    readonly licensePoolname?: pulumi.Input<string>;
    /**
     * MAC address of the BIG-IP. type `string`
     */
    readonly macAddress?: pulumi.Input<string>;
    /**
     * An optional offering name. type `string`
     */
    readonly skukeyword1?: pulumi.Input<string>;
    /**
     * An optional offering name. type `string`
     */
    readonly skukeyword2?: pulumi.Input<string>;
    /**
     * For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
     */
    readonly tenant?: pulumi.Input<string>;
    /**
     * The units used to measure billing. For example, “hourly” or “daily”. Type `string`
     */
    readonly unitOfMeasure?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CommonLicenseManageBigIq resource.
 */
export interface CommonLicenseManageBigIqArgs {
    /**
     * The type of assignment, which is determined by whether the BIG-IP is unreachable, unmanaged, or managed by BIG-IQ. Possible values: “UNREACHABLE”, “UNMANAGED”, or “MANAGED”.
     */
    readonly assignmentType: pulumi.Input<string>;
    /**
     * BIGIQ License Manager IP Address, variable type `string`
     */
    readonly bigiqAddress: pulumi.Input<string>;
    /**
     * BIGIQ Login reference for token authentication
     */
    readonly bigiqLoginRef?: pulumi.Input<string>;
    /**
     * BIGIQ License Manager password.  variable type `string`
     */
    readonly bigiqPassword: pulumi.Input<string>;
    /**
     * type `int`, BIGIQ License Manager Port number, specify if port is other than `443`
     */
    readonly bigiqPort?: pulumi.Input<string>;
    /**
     * type `bool`, if set to `true` enables Token based Authentication,default is `false`
     */
    readonly bigiqTokenAuth?: pulumi.Input<boolean>;
    /**
     * BIGIQ License Manager username, variable type `string`
     */
    readonly bigiqUser: pulumi.Input<string>;
    /**
     * Status of Licence Assignment
     */
    readonly deviceLicenseStatus?: pulumi.Input<string>;
    /**
     * Identifies the platform running the BIG-IP VE. Possible values: “aws”, “azure”, “gce”, “vmware”, “hyperv”, “kvm”, or “xen”. type `string`
     */
    readonly hypervisor?: pulumi.Input<string>;
    /**
     * License Assignment is done with specified `key`, supported only with RegKeypool type License assignement. type `string`
     */
    readonly key?: pulumi.Input<string>;
    /**
     * A name given to the license pool. type `string`
     */
    readonly licensePoolname: pulumi.Input<string>;
    /**
     * MAC address of the BIG-IP. type `string`
     */
    readonly macAddress?: pulumi.Input<string>;
    /**
     * An optional offering name. type `string`
     */
    readonly skukeyword1?: pulumi.Input<string>;
    /**
     * An optional offering name. type `string`
     */
    readonly skukeyword2?: pulumi.Input<string>;
    /**
     * For an unreachable BIG-IP, you can provide an optional description for the assignment in this field.
     */
    readonly tenant?: pulumi.Input<string>;
    /**
     * The units used to measure billing. For example, “hourly” or “daily”. Type `string`
     */
    readonly unitOfMeasure?: pulumi.Input<string>;
}
