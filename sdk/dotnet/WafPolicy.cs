// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP
{
    /// <summary>
    /// `f5bigip.WafPolicy` Manages a WAF Policy resource with its adjustments and modifications on a BIG-IP.
    /// It outputs an up-to-date WAF Policy in a JSON format
    /// 
    /// * [Declarative WAF documentation](https://clouddocs.f5.com/products/waf-declarative-policy/declarative_policy_v16_1.html)
    /// 
    /// &gt; **NOTE** This Resource Requires F5 BIG-IP v16.x above version, and ASM need to be provisioned.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var param1 = F5BigIP.Ssl.GetWafEntityParameter.Invoke(new()
    ///     {
    ///         Name = "Param1",
    ///         Type = "explicit",
    ///         DataType = "alpha-numeric",
    ///         PerformStaging = true,
    ///     });
    /// 
    ///     var param2 = F5BigIP.Ssl.GetWafEntityParameter.Invoke(new()
    ///     {
    ///         Name = "Param2",
    ///         Type = "explicit",
    ///         DataType = "alpha-numeric",
    ///         PerformStaging = true,
    ///     });
    /// 
    ///     var uRL = F5BigIP.Ssl.GetWafEntityUrl.Invoke(new()
    ///     {
    ///         Name = "URL1",
    ///         Protocol = "http",
    ///     });
    /// 
    ///     var uRL2 = F5BigIP.Ssl.GetWafEntityUrl.Invoke(new()
    ///     {
    ///         Name = "URL2",
    ///     });
    /// 
    ///     var test_awaf = new F5BigIP.WafPolicy("test-awaf", new()
    ///     {
    ///         Name = "testpolicyravi",
    ///         Partition = "Common",
    ///         TemplateName = "POLICY_TEMPLATE_RAPID_DEPLOYMENT",
    ///         ApplicationLanguage = "utf-8",
    ///         EnforcementMode = "blocking",
    ///         ServerTechnologies = new[]
    ///         {
    ///             "MySQL",
    ///             "Unix/Linux",
    ///             "MongoDB",
    ///         },
    ///         Parameters = new[]
    ///         {
    ///             param1.Apply(getWafEntityParameterResult =&gt; getWafEntityParameterResult.Json),
    ///             param2.Apply(getWafEntityParameterResult =&gt; getWafEntityParameterResult.Json),
    ///         },
    ///         Urls = new[]
    ///         {
    ///             uRL.Apply(getWafEntityUrlResult =&gt; getWafEntityUrlResult.Json),
    ///             uRL2.Apply(getWafEntityUrlResult =&gt; getWafEntityUrlResult.Json),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// An existing WAF Policy or if the WAF Policy has been manually created or modified on the BIG-IP WebUI, it can be imported using its `id`. e.g
    /// 
    /// ```sh
    ///  $ pulumi import f5bigip:index/wafPolicy:WafPolicy example &lt;id&gt;
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/wafPolicy:WafPolicy")]
    public partial class WafPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
        /// </summary>
        [Output("applicationLanguage")]
        public Output<string?> ApplicationLanguage { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
        /// </summary>
        [Output("caseInsensitive")]
        public Output<bool?> CaseInsensitive { get; private set; } = null!;

        /// <summary>
        /// Specifies the description of the policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
        /// </summary>
        [Output("enablePassivemode")]
        public Output<bool?> EnablePassivemode { get; private set; } = null!;

        /// <summary>
        /// How the system processes a request that triggers a security policy violation
        /// </summary>
        [Output("enforcementMode")]
        public Output<string?> EnforcementMode { get; private set; } = null!;

        /// <summary>
        /// `file_types` takes list of file-types options to be used for policy builder.
        /// See file types below for more details.
        /// </summary>
        [Output("fileTypes")]
        public Output<ImmutableArray<Outputs.WafPolicyFileType>> FileTypes { get; private set; } = null!;

        /// <summary>
        /// `graphql_profiles` takes list of graphql profile options to be used for policy builder.
        /// See graphql profiles below for more details.
        /// </summary>
        [Output("graphqlProfiles")]
        public Output<ImmutableArray<Outputs.WafPolicyGraphqlProfile>> GraphqlProfiles { get; private set; } = null!;

        /// <summary>
        /// specify the list of host name that is used to access the application
        /// </summary>
        [Output("hostNames")]
        public Output<ImmutableArray<Outputs.WafPolicyHostName>> HostNames { get; private set; } = null!;

        /// <summary>
        /// `ip_exceptions` takes list of IP address exception,An IP address exception is an IP address that you want the system to treat in a specific way for a security policy.For example, you can specify IP addresses from which the system should always trust traffic.
        /// See IP Exceptions below for more details.
        /// </summary>
        [Output("ipExceptions")]
        public Output<ImmutableArray<Outputs.WafPolicyIpException>> IpExceptions { get; private set; } = null!;

        /// <summary>
        /// the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
        /// section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
        /// provided by the BIG-IP.
        /// </summary>
        [Output("modifications")]
        public Output<ImmutableArray<string>> Modifications { get; private set; } = null!;

        /// <summary>
        /// The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This section defines the Link for open api files on the policy.
        /// </summary>
        [Output("openApiFiles")]
        public Output<ImmutableArray<string>> OpenApiFiles { get; private set; } = null!;

        /// <summary>
        /// This section defines parameters that the security policy permits in requests.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<string>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Specifies the partition of the policy. Default is `Common`
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;

        /// <summary>
        /// `policy_builder` block will provide `learning_mode` options to be used for policy builder.
        /// See policy builder below for more details.
        /// </summary>
        [Output("policyBuilders")]
        public Output<ImmutableArray<Outputs.WafPolicyPolicyBuilder>> PolicyBuilders { get; private set; } = null!;

        /// <summary>
        /// Exported WAF policy deployed on BIGIP.
        /// </summary>
        [Output("policyExportJson")]
        public Output<string> PolicyExportJson { get; private set; } = null!;

        /// <summary>
        /// The id of the A.WAF Policy as it would be calculated on the BIG-IP.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
        /// </summary>
        [Output("policyImportJson")]
        public Output<string?> PolicyImportJson { get; private set; } = null!;

        /// <summary>
        /// When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
        /// </summary>
        [Output("protocolIndependent")]
        public Output<bool?> ProtocolIndependent { get; private set; } = null!;

        /// <summary>
        /// The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
        /// </summary>
        [Output("serverTechnologies")]
        public Output<ImmutableArray<string>> ServerTechnologies { get; private set; } = null!;

        /// <summary>
        /// Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
        /// </summary>
        [Output("signatureSets")]
        public Output<ImmutableArray<string>> SignatureSets { get; private set; } = null!;

        /// <summary>
        /// This section defines the properties of a signature on the policy.
        /// </summary>
        [Output("signatures")]
        public Output<ImmutableArray<string>> Signatures { get; private set; } = null!;

        /// <summary>
        /// bulk signature setting
        /// </summary>
        [Output("signaturesSettings")]
        public Output<ImmutableArray<Outputs.WafPolicySignaturesSetting>> SignaturesSettings { get; private set; } = null!;

        /// <summary>
        /// Specifies the Link of the template used for the policy creation.
        /// </summary>
        [Output("templateLink")]
        public Output<string?> TemplateLink { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the template used for the policy creation.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// The type of policy you want to create. The default policy type is `security`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
        /// </summary>
        [Output("urls")]
        public Output<ImmutableArray<string>> Urls { get; private set; } = null!;


        /// <summary>
        /// Create a WafPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WafPolicy(string name, WafPolicyArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/wafPolicy:WafPolicy", name, args ?? new WafPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WafPolicy(string name, Input<string> id, WafPolicyState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/wafPolicy:WafPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WafPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WafPolicy Get(string name, Input<string> id, WafPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new WafPolicy(name, id, state, options);
        }
    }

    public sealed class WafPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
        /// </summary>
        [Input("applicationLanguage")]
        public Input<string>? ApplicationLanguage { get; set; }

        /// <summary>
        /// Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
        /// </summary>
        [Input("caseInsensitive")]
        public Input<bool>? CaseInsensitive { get; set; }

        /// <summary>
        /// Specifies the description of the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
        /// </summary>
        [Input("enablePassivemode")]
        public Input<bool>? EnablePassivemode { get; set; }

        /// <summary>
        /// How the system processes a request that triggers a security policy violation
        /// </summary>
        [Input("enforcementMode")]
        public Input<string>? EnforcementMode { get; set; }

        [Input("fileTypes")]
        private InputList<Inputs.WafPolicyFileTypeArgs>? _fileTypes;

        /// <summary>
        /// `file_types` takes list of file-types options to be used for policy builder.
        /// See file types below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyFileTypeArgs> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<Inputs.WafPolicyFileTypeArgs>());
            set => _fileTypes = value;
        }

        [Input("graphqlProfiles")]
        private InputList<Inputs.WafPolicyGraphqlProfileArgs>? _graphqlProfiles;

        /// <summary>
        /// `graphql_profiles` takes list of graphql profile options to be used for policy builder.
        /// See graphql profiles below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyGraphqlProfileArgs> GraphqlProfiles
        {
            get => _graphqlProfiles ?? (_graphqlProfiles = new InputList<Inputs.WafPolicyGraphqlProfileArgs>());
            set => _graphqlProfiles = value;
        }

        [Input("hostNames")]
        private InputList<Inputs.WafPolicyHostNameArgs>? _hostNames;

        /// <summary>
        /// specify the list of host name that is used to access the application
        /// </summary>
        public InputList<Inputs.WafPolicyHostNameArgs> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<Inputs.WafPolicyHostNameArgs>());
            set => _hostNames = value;
        }

        [Input("ipExceptions")]
        private InputList<Inputs.WafPolicyIpExceptionArgs>? _ipExceptions;

        /// <summary>
        /// `ip_exceptions` takes list of IP address exception,An IP address exception is an IP address that you want the system to treat in a specific way for a security policy.For example, you can specify IP addresses from which the system should always trust traffic.
        /// See IP Exceptions below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyIpExceptionArgs> IpExceptions
        {
            get => _ipExceptions ?? (_ipExceptions = new InputList<Inputs.WafPolicyIpExceptionArgs>());
            set => _ipExceptions = value;
        }

        [Input("modifications")]
        private InputList<string>? _modifications;

        /// <summary>
        /// the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
        /// section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
        /// provided by the BIG-IP.
        /// </summary>
        public InputList<string> Modifications
        {
            get => _modifications ?? (_modifications = new InputList<string>());
            set => _modifications = value;
        }

        /// <summary>
        /// The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("openApiFiles")]
        private InputList<string>? _openApiFiles;

        /// <summary>
        /// This section defines the Link for open api files on the policy.
        /// </summary>
        public InputList<string> OpenApiFiles
        {
            get => _openApiFiles ?? (_openApiFiles = new InputList<string>());
            set => _openApiFiles = value;
        }

        [Input("parameters")]
        private InputList<string>? _parameters;

        /// <summary>
        /// This section defines parameters that the security policy permits in requests.
        /// </summary>
        public InputList<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Specifies the partition of the policy. Default is `Common`
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        [Input("policyBuilders")]
        private InputList<Inputs.WafPolicyPolicyBuilderArgs>? _policyBuilders;

        /// <summary>
        /// `policy_builder` block will provide `learning_mode` options to be used for policy builder.
        /// See policy builder below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyPolicyBuilderArgs> PolicyBuilders
        {
            get => _policyBuilders ?? (_policyBuilders = new InputList<Inputs.WafPolicyPolicyBuilderArgs>());
            set => _policyBuilders = value;
        }

        /// <summary>
        /// The id of the A.WAF Policy as it would be calculated on the BIG-IP.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
        /// </summary>
        [Input("policyImportJson")]
        public Input<string>? PolicyImportJson { get; set; }

        /// <summary>
        /// When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
        /// </summary>
        [Input("protocolIndependent")]
        public Input<bool>? ProtocolIndependent { get; set; }

        [Input("serverTechnologies")]
        private InputList<string>? _serverTechnologies;

        /// <summary>
        /// The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
        /// </summary>
        public InputList<string> ServerTechnologies
        {
            get => _serverTechnologies ?? (_serverTechnologies = new InputList<string>());
            set => _serverTechnologies = value;
        }

        [Input("signatureSets")]
        private InputList<string>? _signatureSets;

        /// <summary>
        /// Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
        /// </summary>
        public InputList<string> SignatureSets
        {
            get => _signatureSets ?? (_signatureSets = new InputList<string>());
            set => _signatureSets = value;
        }

        [Input("signatures")]
        private InputList<string>? _signatures;

        /// <summary>
        /// This section defines the properties of a signature on the policy.
        /// </summary>
        public InputList<string> Signatures
        {
            get => _signatures ?? (_signatures = new InputList<string>());
            set => _signatures = value;
        }

        [Input("signaturesSettings")]
        private InputList<Inputs.WafPolicySignaturesSettingArgs>? _signaturesSettings;

        /// <summary>
        /// bulk signature setting
        /// </summary>
        public InputList<Inputs.WafPolicySignaturesSettingArgs> SignaturesSettings
        {
            get => _signaturesSettings ?? (_signaturesSettings = new InputList<Inputs.WafPolicySignaturesSettingArgs>());
            set => _signaturesSettings = value;
        }

        /// <summary>
        /// Specifies the Link of the template used for the policy creation.
        /// </summary>
        [Input("templateLink")]
        public Input<string>? TemplateLink { get; set; }

        /// <summary>
        /// Specifies the name of the template used for the policy creation.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// The type of policy you want to create. The default policy type is `security`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("urls")]
        private InputList<string>? _urls;

        /// <summary>
        /// In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        public WafPolicyArgs()
        {
        }
        public static new WafPolicyArgs Empty => new WafPolicyArgs();
    }

    public sealed class WafPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The character encoding for the web application. The character encoding determines how the policy processes the character sets. The default is `utf-8`
        /// </summary>
        [Input("applicationLanguage")]
        public Input<string>? ApplicationLanguage { get; set; }

        /// <summary>
        /// Specifies whether the security policy treats microservice URLs, file types, URLs, and parameters as case sensitive or not. When this setting is enabled, the system stores these security policy elements in lowercase in the security policy configuration
        /// </summary>
        [Input("caseInsensitive")]
        public Input<bool>? CaseInsensitive { get; set; }

        /// <summary>
        /// Specifies the description of the policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Passive Mode allows the policy to be associated with a Performance L4 Virtual Server (using a FastL4 profile). With FastL4, traffic is analyzed but is not modified in any way.
        /// </summary>
        [Input("enablePassivemode")]
        public Input<bool>? EnablePassivemode { get; set; }

        /// <summary>
        /// How the system processes a request that triggers a security policy violation
        /// </summary>
        [Input("enforcementMode")]
        public Input<string>? EnforcementMode { get; set; }

        [Input("fileTypes")]
        private InputList<Inputs.WafPolicyFileTypeGetArgs>? _fileTypes;

        /// <summary>
        /// `file_types` takes list of file-types options to be used for policy builder.
        /// See file types below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyFileTypeGetArgs> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<Inputs.WafPolicyFileTypeGetArgs>());
            set => _fileTypes = value;
        }

        [Input("graphqlProfiles")]
        private InputList<Inputs.WafPolicyGraphqlProfileGetArgs>? _graphqlProfiles;

        /// <summary>
        /// `graphql_profiles` takes list of graphql profile options to be used for policy builder.
        /// See graphql profiles below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyGraphqlProfileGetArgs> GraphqlProfiles
        {
            get => _graphqlProfiles ?? (_graphqlProfiles = new InputList<Inputs.WafPolicyGraphqlProfileGetArgs>());
            set => _graphqlProfiles = value;
        }

        [Input("hostNames")]
        private InputList<Inputs.WafPolicyHostNameGetArgs>? _hostNames;

        /// <summary>
        /// specify the list of host name that is used to access the application
        /// </summary>
        public InputList<Inputs.WafPolicyHostNameGetArgs> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<Inputs.WafPolicyHostNameGetArgs>());
            set => _hostNames = value;
        }

        [Input("ipExceptions")]
        private InputList<Inputs.WafPolicyIpExceptionGetArgs>? _ipExceptions;

        /// <summary>
        /// `ip_exceptions` takes list of IP address exception,An IP address exception is an IP address that you want the system to treat in a specific way for a security policy.For example, you can specify IP addresses from which the system should always trust traffic.
        /// See IP Exceptions below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyIpExceptionGetArgs> IpExceptions
        {
            get => _ipExceptions ?? (_ipExceptions = new InputList<Inputs.WafPolicyIpExceptionGetArgs>());
            set => _ipExceptions = value;
        }

        [Input("modifications")]
        private InputList<string>? _modifications;

        /// <summary>
        /// the modifications section includes actions that modify the declarative policy as it is defined in the adjustments
        /// section. The modifications section is updated manually, with the changes generally driven by the learning suggestions
        /// provided by the BIG-IP.
        /// </summary>
        public InputList<string> Modifications
        {
            get => _modifications ?? (_modifications = new InputList<string>());
            set => _modifications = value;
        }

        /// <summary>
        /// The unique user-given name of the policy. Policy names cannot contain spaces or special characters. Allowed characters are a-z, A-Z, 0-9, dot, dash (-), colon (:) and underscore (_).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("openApiFiles")]
        private InputList<string>? _openApiFiles;

        /// <summary>
        /// This section defines the Link for open api files on the policy.
        /// </summary>
        public InputList<string> OpenApiFiles
        {
            get => _openApiFiles ?? (_openApiFiles = new InputList<string>());
            set => _openApiFiles = value;
        }

        [Input("parameters")]
        private InputList<string>? _parameters;

        /// <summary>
        /// This section defines parameters that the security policy permits in requests.
        /// </summary>
        public InputList<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Specifies the partition of the policy. Default is `Common`
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        [Input("policyBuilders")]
        private InputList<Inputs.WafPolicyPolicyBuilderGetArgs>? _policyBuilders;

        /// <summary>
        /// `policy_builder` block will provide `learning_mode` options to be used for policy builder.
        /// See policy builder below for more details.
        /// </summary>
        public InputList<Inputs.WafPolicyPolicyBuilderGetArgs> PolicyBuilders
        {
            get => _policyBuilders ?? (_policyBuilders = new InputList<Inputs.WafPolicyPolicyBuilderGetArgs>());
            set => _policyBuilders = value;
        }

        /// <summary>
        /// Exported WAF policy deployed on BIGIP.
        /// </summary>
        [Input("policyExportJson")]
        public Input<string>? PolicyExportJson { get; set; }

        /// <summary>
        /// The id of the A.WAF Policy as it would be calculated on the BIG-IP.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The payload of the WAF Policy to be used for IMPORT on to BIG-IP.
        /// </summary>
        [Input("policyImportJson")]
        public Input<string>? PolicyImportJson { get; set; }

        /// <summary>
        /// When creating a security policy, you can determine whether a security policy differentiates between HTTP and HTTPS URLs. If enabled, the security policy differentiates between HTTP and HTTPS URLs. If disabled, the security policy configures URLs without specifying a specific protocol. This is useful for applications that behave the same for HTTP and HTTPS, and it keeps the security policy from including the same URL twice.
        /// </summary>
        [Input("protocolIndependent")]
        public Input<bool>? ProtocolIndependent { get; set; }

        [Input("serverTechnologies")]
        private InputList<string>? _serverTechnologies;

        /// <summary>
        /// The server technology is a server-side application, framework, web server or operating system type that is configured in the policy in order to adapt the policy to the checks needed for the respective technology.
        /// </summary>
        public InputList<string> ServerTechnologies
        {
            get => _serverTechnologies ?? (_serverTechnologies = new InputList<string>());
            set => _serverTechnologies = value;
        }

        [Input("signatureSets")]
        private InputList<string>? _signatureSets;

        /// <summary>
        /// Defines behavior when signatures found within a signature-set are detected in a request. Settings are culmulative, so if a signature is found in any set with block enabled, that signature will have block enabled.
        /// </summary>
        public InputList<string> SignatureSets
        {
            get => _signatureSets ?? (_signatureSets = new InputList<string>());
            set => _signatureSets = value;
        }

        [Input("signatures")]
        private InputList<string>? _signatures;

        /// <summary>
        /// This section defines the properties of a signature on the policy.
        /// </summary>
        public InputList<string> Signatures
        {
            get => _signatures ?? (_signatures = new InputList<string>());
            set => _signatures = value;
        }

        [Input("signaturesSettings")]
        private InputList<Inputs.WafPolicySignaturesSettingGetArgs>? _signaturesSettings;

        /// <summary>
        /// bulk signature setting
        /// </summary>
        public InputList<Inputs.WafPolicySignaturesSettingGetArgs> SignaturesSettings
        {
            get => _signaturesSettings ?? (_signaturesSettings = new InputList<Inputs.WafPolicySignaturesSettingGetArgs>());
            set => _signaturesSettings = value;
        }

        /// <summary>
        /// Specifies the Link of the template used for the policy creation.
        /// </summary>
        [Input("templateLink")]
        public Input<string>? TemplateLink { get; set; }

        /// <summary>
        /// Specifies the name of the template used for the policy creation.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        /// <summary>
        /// The type of policy you want to create. The default policy type is `security`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("urls")]
        private InputList<string>? _urls;

        /// <summary>
        /// In a security policy, you can manually specify the HTTP URLs that are allowed (or disallowed) in traffic to the web application being protected. If you are using automatic policy building (and the policy includes learning URLs), the system can determine which URLs to add, based on legitimate traffic.
        /// </summary>
        public InputList<string> Urls
        {
            get => _urls ?? (_urls = new InputList<string>());
            set => _urls = value;
        }

        public WafPolicyState()
        {
        }
        public static new WafPolicyState Empty => new WafPolicyState();
    }
}
