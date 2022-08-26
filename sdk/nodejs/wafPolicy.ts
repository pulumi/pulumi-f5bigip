// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `f5bigip.WafPolicy` Manages a WAF Policy resource with its adjustments and modifications on a BIG-IP.
 * It outputs an up-to-date WAF Policy in a JSON format
 *
 * * [Declarative WAF documentation](https://clouddocs.f5.com/products/waf-declarative-policy/declarative_policy_v16_1.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as f5bigip from "@pulumi/f5bigip";
 *
 * const param1 = f5bigip.ssl.getWafEntityParameter({
 *     name: "Param1",
 *     type: "explicit",
 *     dataType: "alpha-numeric",
 *     performStaging: true,
 * });
 * const param2 = f5bigip.ssl.getWafEntityParameter({
 *     name: "Param2",
 *     type: "explicit",
 *     dataType: "alpha-numeric",
 *     performStaging: true,
 * });
 * const uRL = f5bigip.ssl.getWafEntityUrl({
 *     name: "URL1",
 *     protocol: "http",
 * });
 * const uRL2 = f5bigip.ssl.getWafEntityUrl({
 *     name: "URL2",
 * });
 * const test_awaf = new f5bigip.WafPolicy("test-awaf", {
 *     name: "testpolicyravi",
 *     partition: "Common",
 *     templateName: "POLICY_TEMPLATE_RAPID_DEPLOYMENT",
 *     applicationLanguage: "utf-8",
 *     enforcementMode: "blocking",
 *     serverTechnologies: [
 *         "MySQL",
 *         "Unix/Linux",
 *         "MongoDB",
 *     ],
 *     parameters: [
 *         param1.then(param1 => param1.json),
 *         param2.then(param2 => param2.json),
 *     ],
 *     urls: [
 *         uRL.then(uRL => uRL.json),
 *         uRL2.then(uRL2 => uRL2.json),
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * An existing WAF Policy or if the WAF Policy has been manually created or modified on the BIG-IP WebUI, it can be imported using its `id`. e.g
 *
 * ```sh
 *  $ pulumi import f5bigip:index/wafPolicy:WafPolicy example <id>
 * ```
 */
export class WafPolicy extends pulumi.CustomResource {
    /**
     * Get an existing WafPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WafPolicyState, opts?: pulumi.CustomResourceOptions): WafPolicy {
        return new WafPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/wafPolicy:WafPolicy';

    /**
     * Returns true if the given object is an instance of WafPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WafPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WafPolicy.__pulumiType;
    }

    /**
     * The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
     */
    public readonly applicationLanguage!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
     */
    public readonly caseInsensitive!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the description of the policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
     */
    public readonly enablePassivemode!: pulumi.Output<boolean | undefined>;
    /**
     * How the system processes a request that triggers a security policy violation
     */
    public readonly enforcementMode!: pulumi.Output<string | undefined>;
    /**
     * `fileTypes` takes list of file-types options to be used for policy builder.
     * See file types below for more details.
     */
    public readonly fileTypes!: pulumi.Output<outputs.WafPolicyFileType[] | undefined>;
    /**
     * `graphqlProfiles` takes list of graphql profile options to be used for policy builder.
     * See graphql profiles below for more details.
     */
    public readonly graphqlProfiles!: pulumi.Output<outputs.WafPolicyGraphqlProfile[] | undefined>;
    /**
     * the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
     * section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
     * provided by the BIG-IP.
     */
    public readonly modifications!: pulumi.Output<string[] | undefined>;
    /**
     * The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * This section defines the Link for open api files on the policy.
     */
    public readonly openApiFiles!: pulumi.Output<string[] | undefined>;
    /**
     * This section defines parameters that the security policy permits in requests.
     */
    public readonly parameters!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the partition of the policy. Default is `Common`
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * `policyBuilder` block will provide `learningMode` options to be used for policy builder.
     * See policy builder below for more details.
     */
    public readonly policyBuilders!: pulumi.Output<outputs.WafPolicyPolicyBuilder[] | undefined>;
    /**
     * Exported WAF policy deployed on BIGIP.
     */
    public readonly policyExportJson!: pulumi.Output<string>;
    /**
     * The id of the A.WAF Policy as it would be calculated on the BIG-IP.
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
     */
    public readonly policyImportJson!: pulumi.Output<string | undefined>;
    /**
     * When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
     */
    public readonly protocolIndependent!: pulumi.Output<boolean | undefined>;
    /**
     * The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
     */
    public readonly serverTechnologies!: pulumi.Output<string[] | undefined>;
    /**
     * Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
     */
    public readonly signatureSets!: pulumi.Output<string[] | undefined>;
    /**
     * This section defines the properties of a signature on the policy.
     */
    public readonly signatures!: pulumi.Output<string[] | undefined>;
    /**
     * bulk signature setting
     */
    public readonly signaturesSettings!: pulumi.Output<outputs.WafPolicySignaturesSetting[] | undefined>;
    /**
     * Specifies the name of the template used for the policy creation.
     */
    public readonly templateName!: pulumi.Output<string>;
    /**
     * The type of policy you want to create. The default policy type is `security`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
     */
    public readonly urls!: pulumi.Output<string[] | undefined>;

    /**
     * Create a WafPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WafPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WafPolicyArgs | WafPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WafPolicyState | undefined;
            resourceInputs["applicationLanguage"] = state ? state.applicationLanguage : undefined;
            resourceInputs["caseInsensitive"] = state ? state.caseInsensitive : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enablePassivemode"] = state ? state.enablePassivemode : undefined;
            resourceInputs["enforcementMode"] = state ? state.enforcementMode : undefined;
            resourceInputs["fileTypes"] = state ? state.fileTypes : undefined;
            resourceInputs["graphqlProfiles"] = state ? state.graphqlProfiles : undefined;
            resourceInputs["modifications"] = state ? state.modifications : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["openApiFiles"] = state ? state.openApiFiles : undefined;
            resourceInputs["parameters"] = state ? state.parameters : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["policyBuilders"] = state ? state.policyBuilders : undefined;
            resourceInputs["policyExportJson"] = state ? state.policyExportJson : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["policyImportJson"] = state ? state.policyImportJson : undefined;
            resourceInputs["protocolIndependent"] = state ? state.protocolIndependent : undefined;
            resourceInputs["serverTechnologies"] = state ? state.serverTechnologies : undefined;
            resourceInputs["signatureSets"] = state ? state.signatureSets : undefined;
            resourceInputs["signatures"] = state ? state.signatures : undefined;
            resourceInputs["signaturesSettings"] = state ? state.signaturesSettings : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["urls"] = state ? state.urls : undefined;
        } else {
            const args = argsOrState as WafPolicyArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["applicationLanguage"] = args ? args.applicationLanguage : undefined;
            resourceInputs["caseInsensitive"] = args ? args.caseInsensitive : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enablePassivemode"] = args ? args.enablePassivemode : undefined;
            resourceInputs["enforcementMode"] = args ? args.enforcementMode : undefined;
            resourceInputs["fileTypes"] = args ? args.fileTypes : undefined;
            resourceInputs["graphqlProfiles"] = args ? args.graphqlProfiles : undefined;
            resourceInputs["modifications"] = args ? args.modifications : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["openApiFiles"] = args ? args.openApiFiles : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["policyBuilders"] = args ? args.policyBuilders : undefined;
            resourceInputs["policyExportJson"] = args ? args.policyExportJson : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["policyImportJson"] = args ? args.policyImportJson : undefined;
            resourceInputs["protocolIndependent"] = args ? args.protocolIndependent : undefined;
            resourceInputs["serverTechnologies"] = args ? args.serverTechnologies : undefined;
            resourceInputs["signatureSets"] = args ? args.signatureSets : undefined;
            resourceInputs["signatures"] = args ? args.signatures : undefined;
            resourceInputs["signaturesSettings"] = args ? args.signaturesSettings : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["urls"] = args ? args.urls : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WafPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WafPolicy resources.
 */
export interface WafPolicyState {
    /**
     * The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
     */
    applicationLanguage?: pulumi.Input<string>;
    /**
     * Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
     */
    caseInsensitive?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
     */
    enablePassivemode?: pulumi.Input<boolean>;
    /**
     * How the system processes a request that triggers a security policy violation
     */
    enforcementMode?: pulumi.Input<string>;
    /**
     * `fileTypes` takes list of file-types options to be used for policy builder.
     * See file types below for more details.
     */
    fileTypes?: pulumi.Input<pulumi.Input<inputs.WafPolicyFileType>[]>;
    /**
     * `graphqlProfiles` takes list of graphql profile options to be used for policy builder.
     * See graphql profiles below for more details.
     */
    graphqlProfiles?: pulumi.Input<pulumi.Input<inputs.WafPolicyGraphqlProfile>[]>;
    /**
     * the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
     * section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
     * provided by the BIG-IP.
     */
    modifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     */
    name?: pulumi.Input<string>;
    /**
     * This section defines the Link for open api files on the policy.
     */
    openApiFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This section defines parameters that the security policy permits in requests.
     */
    parameters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the partition of the policy. Default is `Common`
     */
    partition?: pulumi.Input<string>;
    /**
     * `policyBuilder` block will provide `learningMode` options to be used for policy builder.
     * See policy builder below for more details.
     */
    policyBuilders?: pulumi.Input<pulumi.Input<inputs.WafPolicyPolicyBuilder>[]>;
    /**
     * Exported WAF policy deployed on BIGIP.
     */
    policyExportJson?: pulumi.Input<string>;
    /**
     * The id of the A.WAF Policy as it would be calculated on the BIG-IP.
     */
    policyId?: pulumi.Input<string>;
    /**
     * The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
     */
    policyImportJson?: pulumi.Input<string>;
    /**
     * When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
     */
    protocolIndependent?: pulumi.Input<boolean>;
    /**
     * The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
     */
    serverTechnologies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
     */
    signatureSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This section defines the properties of a signature on the policy.
     */
    signatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * bulk signature setting
     */
    signaturesSettings?: pulumi.Input<pulumi.Input<inputs.WafPolicySignaturesSetting>[]>;
    /**
     * Specifies the name of the template used for the policy creation.
     */
    templateName?: pulumi.Input<string>;
    /**
     * The type of policy you want to create. The default policy type is `security`.
     */
    type?: pulumi.Input<string>;
    /**
     * In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
     */
    urls?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a WafPolicy resource.
 */
export interface WafPolicyArgs {
    /**
     * The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
     */
    applicationLanguage?: pulumi.Input<string>;
    /**
     * Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
     */
    caseInsensitive?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
     */
    enablePassivemode?: pulumi.Input<boolean>;
    /**
     * How the system processes a request that triggers a security policy violation
     */
    enforcementMode?: pulumi.Input<string>;
    /**
     * `fileTypes` takes list of file-types options to be used for policy builder.
     * See file types below for more details.
     */
    fileTypes?: pulumi.Input<pulumi.Input<inputs.WafPolicyFileType>[]>;
    /**
     * `graphqlProfiles` takes list of graphql profile options to be used for policy builder.
     * See graphql profiles below for more details.
     */
    graphqlProfiles?: pulumi.Input<pulumi.Input<inputs.WafPolicyGraphqlProfile>[]>;
    /**
     * the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
     * section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
     * provided by the BIG-IP.
     */
    modifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     */
    name: pulumi.Input<string>;
    /**
     * This section defines the Link for open api files on the policy.
     */
    openApiFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This section defines parameters that the security policy permits in requests.
     */
    parameters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the partition of the policy. Default is `Common`
     */
    partition?: pulumi.Input<string>;
    /**
     * `policyBuilder` block will provide `learningMode` options to be used for policy builder.
     * See policy builder below for more details.
     */
    policyBuilders?: pulumi.Input<pulumi.Input<inputs.WafPolicyPolicyBuilder>[]>;
    /**
     * Exported WAF policy deployed on BIGIP.
     */
    policyExportJson?: pulumi.Input<string>;
    /**
     * The id of the A.WAF Policy as it would be calculated on the BIG-IP.
     */
    policyId?: pulumi.Input<string>;
    /**
     * The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
     */
    policyImportJson?: pulumi.Input<string>;
    /**
     * When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
     */
    protocolIndependent?: pulumi.Input<boolean>;
    /**
     * The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
     */
    serverTechnologies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
     */
    signatureSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * This section defines the properties of a signature on the policy.
     */
    signatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * bulk signature setting
     */
    signaturesSettings?: pulumi.Input<pulumi.Input<inputs.WafPolicySignaturesSetting>[]>;
    /**
     * Specifies the name of the template used for the policy creation.
     */
    templateName: pulumi.Input<string>;
    /**
     * The type of policy you want to create. The default policy type is `security`.
     */
    type?: pulumi.Input<string>;
    /**
     * In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
     */
    urls?: pulumi.Input<pulumi.Input<string>[]>;
}
