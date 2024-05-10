// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.WafPolicyArgs;
import com.pulumi.f5bigip.inputs.WafPolicyState;
import com.pulumi.f5bigip.outputs.WafPolicyFileType;
import com.pulumi.f5bigip.outputs.WafPolicyGraphqlProfile;
import com.pulumi.f5bigip.outputs.WafPolicyHostName;
import com.pulumi.f5bigip.outputs.WafPolicyIpException;
import com.pulumi.f5bigip.outputs.WafPolicyPolicyBuilder;
import com.pulumi.f5bigip.outputs.WafPolicySignaturesSetting;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.WafPolicy` Manages a WAF Policy resource with its adjustments and modifications on a BIG-IP.
 * It outputs an up-to-date WAF Policy in a JSON format
 * 
 * * [Declarative WAF documentation](https://clouddocs.f5.com/products/waf-declarative-policy/declarative_policy_v16_1.html)
 * 
 * &gt; **NOTE** This Resource Requires F5 BIG-IP v16.x above version, and ASM need to be provisioned.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.f5bigip.ssl.SslFunctions;
 * import com.pulumi.f5bigip.ssl.inputs.GetWafEntityParameterArgs;
 * import com.pulumi.f5bigip.ssl.inputs.GetWafEntityUrlArgs;
 * import com.pulumi.f5bigip.WafPolicy;
 * import com.pulumi.f5bigip.WafPolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var param1 = SslFunctions.getWafEntityParameter(GetWafEntityParameterArgs.builder()
 *             .name("Param1")
 *             .type("explicit")
 *             .dataType("alpha-numeric")
 *             .performStaging(true)
 *             .build());
 * 
 *         final var param2 = SslFunctions.getWafEntityParameter(GetWafEntityParameterArgs.builder()
 *             .name("Param2")
 *             .type("explicit")
 *             .dataType("alpha-numeric")
 *             .performStaging(true)
 *             .build());
 * 
 *         final var URL = SslFunctions.getWafEntityUrl(GetWafEntityUrlArgs.builder()
 *             .name("URL1")
 *             .protocol("http")
 *             .build());
 * 
 *         final var URL2 = SslFunctions.getWafEntityUrl(GetWafEntityUrlArgs.builder()
 *             .name("URL2")
 *             .build());
 * 
 *         var test_awaf = new WafPolicy("test-awaf", WafPolicyArgs.builder()        
 *             .name("testpolicyravi")
 *             .partition("Common")
 *             .templateName("POLICY_TEMPLATE_RAPID_DEPLOYMENT")
 *             .applicationLanguage("utf-8")
 *             .enforcementMode("blocking")
 *             .serverTechnologies(            
 *                 "MySQL",
 *                 "Unix/Linux",
 *                 "MongoDB")
 *             .parameters(            
 *                 param1.applyValue(getWafEntityParameterResult -> getWafEntityParameterResult.json()),
 *                 param2.applyValue(getWafEntityParameterResult -> getWafEntityParameterResult.json()))
 *             .urls(            
 *                 URL.applyValue(getWafEntityUrlResult -> getWafEntityUrlResult.json()),
 *                 URL2.applyValue(getWafEntityUrlResult -> getWafEntityUrlResult.json()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * An existing WAF Policy or if the WAF Policy has been manually created or modified on the BIG-IP WebUI, it can be imported using its `id`.
 * 
 * e.g:
 * 
 * ```sh
 * $ pulumi import f5bigip:index/wafPolicy:WafPolicy example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="f5bigip:index/wafPolicy:WafPolicy")
public class WafPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
     * 
     */
    @Export(name="applicationLanguage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> applicationLanguage;

    /**
     * @return The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
     * 
     */
    public Output<Optional<String>> applicationLanguage() {
        return Codegen.optional(this.applicationLanguage);
    }
    /**
     * Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
     * 
     */
    @Export(name="caseInsensitive", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> caseInsensitive;

    /**
     * @return Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
     * 
     */
    public Output<Optional<Boolean>> caseInsensitive() {
        return Codegen.optional(this.caseInsensitive);
    }
    /**
     * Specifies the description of the policy.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return Specifies the description of the policy.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
     * 
     */
    @Export(name="enablePassivemode", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enablePassivemode;

    /**
     * @return Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
     * 
     */
    public Output<Optional<Boolean>> enablePassivemode() {
        return Codegen.optional(this.enablePassivemode);
    }
    /**
     * How the system processes a request that triggers a security policy violation
     * 
     */
    @Export(name="enforcementMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enforcementMode;

    /**
     * @return How the system processes a request that triggers a security policy violation
     * 
     */
    public Output<Optional<String>> enforcementMode() {
        return Codegen.optional(this.enforcementMode);
    }
    /**
     * `file_types` takes list of file-types options to be used for policy builder.
     * See file types below for more details.
     * 
     */
    @Export(name="fileTypes", refs={List.class,WafPolicyFileType.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WafPolicyFileType>> fileTypes;

    /**
     * @return `file_types` takes list of file-types options to be used for policy builder.
     * See file types below for more details.
     * 
     */
    public Output<Optional<List<WafPolicyFileType>>> fileTypes() {
        return Codegen.optional(this.fileTypes);
    }
    /**
     * `graphql_profiles` takes list of graphql profile options to be used for policy builder.
     * See graphql profiles below for more details.
     * 
     */
    @Export(name="graphqlProfiles", refs={List.class,WafPolicyGraphqlProfile.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WafPolicyGraphqlProfile>> graphqlProfiles;

    /**
     * @return `graphql_profiles` takes list of graphql profile options to be used for policy builder.
     * See graphql profiles below for more details.
     * 
     */
    public Output<Optional<List<WafPolicyGraphqlProfile>>> graphqlProfiles() {
        return Codegen.optional(this.graphqlProfiles);
    }
    /**
     * specify the list of host name that is used to access the application
     * 
     */
    @Export(name="hostNames", refs={List.class,WafPolicyHostName.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WafPolicyHostName>> hostNames;

    /**
     * @return specify the list of host name that is used to access the application
     * 
     */
    public Output<Optional<List<WafPolicyHostName>>> hostNames() {
        return Codegen.optional(this.hostNames);
    }
    /**
     * `ip_exceptions` takes list of IP address exception,An IP address exception is an IP address that you want the system to treat in a specific way for a security policy.For example, you can specify IP addresses from which the system should always trust traffic.
     * See IP Exceptions below for more details.
     * 
     */
    @Export(name="ipExceptions", refs={List.class,WafPolicyIpException.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WafPolicyIpException>> ipExceptions;

    /**
     * @return `ip_exceptions` takes list of IP address exception,An IP address exception is an IP address that you want the system to treat in a specific way for a security policy.For example, you can specify IP addresses from which the system should always trust traffic.
     * See IP Exceptions below for more details.
     * 
     */
    public Output<Optional<List<WafPolicyIpException>>> ipExceptions() {
        return Codegen.optional(this.ipExceptions);
    }
    /**
     * the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
     * section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
     * provided by the BIG-IP.
     * 
     */
    @Export(name="modifications", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> modifications;

    /**
     * @return the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
     * section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
     * provided by the BIG-IP.
     * 
     */
    public Output<Optional<List<String>>> modifications() {
        return Codegen.optional(this.modifications);
    }
    /**
     * The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * This section defines the Link for open api files on the policy.
     * 
     */
    @Export(name="openApiFiles", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> openApiFiles;

    /**
     * @return This section defines the Link for open api files on the policy.
     * 
     */
    public Output<Optional<List<String>>> openApiFiles() {
        return Codegen.optional(this.openApiFiles);
    }
    /**
     * This section defines parameters that the security policy permits in requests.
     * 
     */
    @Export(name="parameters", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> parameters;

    /**
     * @return This section defines parameters that the security policy permits in requests.
     * 
     */
    public Output<Optional<List<String>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * Specifies the partition of the policy. Default is `Common`
     * 
     */
    @Export(name="partition", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partition;

    /**
     * @return Specifies the partition of the policy. Default is `Common`
     * 
     */
    public Output<Optional<String>> partition() {
        return Codegen.optional(this.partition);
    }
    /**
     * `policy_builder` block will provide `learning_mode` options to be used for policy builder.
     * See policy builder below for more details.
     * 
     */
    @Export(name="policyBuilders", refs={List.class,WafPolicyPolicyBuilder.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WafPolicyPolicyBuilder>> policyBuilders;

    /**
     * @return `policy_builder` block will provide `learning_mode` options to be used for policy builder.
     * See policy builder below for more details.
     * 
     */
    public Output<Optional<List<WafPolicyPolicyBuilder>>> policyBuilders() {
        return Codegen.optional(this.policyBuilders);
    }
    /**
     * Exported WAF policy deployed on BIGIP.
     * 
     */
    @Export(name="policyExportJson", refs={String.class}, tree="[0]")
    private Output<String> policyExportJson;

    /**
     * @return Exported WAF policy deployed on BIGIP.
     * 
     */
    public Output<String> policyExportJson() {
        return this.policyExportJson;
    }
    /**
     * The id of the A.WAF Policy as it would be calculated on the BIG-IP.
     * 
     */
    @Export(name="policyId", refs={String.class}, tree="[0]")
    private Output<String> policyId;

    /**
     * @return The id of the A.WAF Policy as it would be calculated on the BIG-IP.
     * 
     */
    public Output<String> policyId() {
        return this.policyId;
    }
    /**
     * The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
     * 
     */
    @Export(name="policyImportJson", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> policyImportJson;

    /**
     * @return The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
     * 
     */
    public Output<Optional<String>> policyImportJson() {
        return Codegen.optional(this.policyImportJson);
    }
    /**
     * When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
     * 
     */
    @Export(name="protocolIndependent", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> protocolIndependent;

    /**
     * @return When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
     * 
     */
    public Output<Optional<Boolean>> protocolIndependent() {
        return Codegen.optional(this.protocolIndependent);
    }
    /**
     * The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
     * 
     */
    @Export(name="serverTechnologies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> serverTechnologies;

    /**
     * @return The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
     * 
     */
    public Output<Optional<List<String>>> serverTechnologies() {
        return Codegen.optional(this.serverTechnologies);
    }
    /**
     * Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
     * 
     */
    @Export(name="signatureSets", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> signatureSets;

    /**
     * @return Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
     * 
     */
    public Output<Optional<List<String>>> signatureSets() {
        return Codegen.optional(this.signatureSets);
    }
    /**
     * This section defines the properties of a signature on the policy.
     * 
     */
    @Export(name="signatures", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> signatures;

    /**
     * @return This section defines the properties of a signature on the policy.
     * 
     */
    public Output<Optional<List<String>>> signatures() {
        return Codegen.optional(this.signatures);
    }
    /**
     * bulk signature setting
     * 
     */
    @Export(name="signaturesSettings", refs={List.class,WafPolicySignaturesSetting.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WafPolicySignaturesSetting>> signaturesSettings;

    /**
     * @return bulk signature setting
     * 
     */
    public Output<Optional<List<WafPolicySignaturesSetting>>> signaturesSettings() {
        return Codegen.optional(this.signaturesSettings);
    }
    /**
     * Specifies the Link of the template used for the policy creation.
     * 
     */
    @Export(name="templateLink", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateLink;

    /**
     * @return Specifies the Link of the template used for the policy creation.
     * 
     */
    public Output<Optional<String>> templateLink() {
        return Codegen.optional(this.templateLink);
    }
    /**
     * Specifies the name of the template used for the policy creation.
     * 
     */
    @Export(name="templateName", refs={String.class}, tree="[0]")
    private Output<String> templateName;

    /**
     * @return Specifies the name of the template used for the policy creation.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }
    /**
     * The type of policy you want to create. The default policy type is `security`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of policy you want to create. The default policy type is `security`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
     * 
     */
    @Export(name="urls", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> urls;

    /**
     * @return In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
     * 
     */
    public Output<Optional<List<String>>> urls() {
        return Codegen.optional(this.urls);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WafPolicy(String name) {
        this(name, WafPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WafPolicy(String name, WafPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WafPolicy(String name, WafPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/wafPolicy:WafPolicy", name, args == null ? WafPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private WafPolicy(String name, Output<String> id, @Nullable WafPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:index/wafPolicy:WafPolicy", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static WafPolicy get(String name, Output<String> id, @Nullable WafPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WafPolicy(name, id, state, options);
    }
}
