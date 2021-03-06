// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.ProfileHttp` Configures a custom profile_http for use by health checks.
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using F5BigIP = Pulumi.F5BigIP;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var sanjose_http = new F5BigIP.Ltm.ProfileHttp("sanjose-http", new F5BigIP.Ltm.ProfileHttpArgs
    ///         {
    ///             DefaultsFrom = "/Common/http",
    ///             Description = "some http",
    ///             FallbackHost = "titanic",
    ///             FallbackStatusCodes = 
    ///             {
    ///                 "400",
    ///                 "500",
    ///                 "300",
    ///             },
    ///             Name = "/Common/sanjose-http",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// BIG-IP LTM http profiles can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import f5bigip:ltm/profileHttp:ProfileHttp test-http /Common/test-http
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/profileHttp:ProfileHttp")]
    public partial class ProfileHttp : Pulumi.CustomResource
    {
        /// <summary>
        /// Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
        /// XFF (X-forwarded-for) headers, if they exist.
        /// </summary>
        [Output("acceptXff")]
        public Output<string> AcceptXff { get; private set; } = null!;

        /// <summary>
        /// The application service to which the object belongs.
        /// </summary>
        [Output("appService")]
        public Output<string?> AppService { get; private set; } = null!;

        /// <summary>
        /// Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        /// </summary>
        [Output("basicAuthRealm")]
        public Output<string> BasicAuthRealm { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// User defined description
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies a passphrase for the cookie encryption
        /// </summary>
        [Output("encryptCookieSecret")]
        public Output<string?> EncryptCookieSecret { get; private set; } = null!;

        /// <summary>
        /// Encrypts specified cookies that the BIG-IP system sends to a client system
        /// </summary>
        [Output("encryptCookies")]
        public Output<ImmutableArray<string>> EncryptCookies { get; private set; } = null!;

        /// <summary>
        /// Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        /// </summary>
        [Output("fallbackHost")]
        public Output<string> FallbackHost { get; private set; } = null!;

        /// <summary>
        /// Specifies one or more three-digit status codes that can be returned by an HTTP server.
        /// </summary>
        [Output("fallbackStatusCodes")]
        public Output<ImmutableArray<string>> FallbackStatusCodes { get; private set; } = null!;

        /// <summary>
        /// Specifies the header string that you want to erase from an HTTP request. You can also specify none
        /// </summary>
        [Output("headErase")]
        public Output<string> HeadErase { get; private set; } = null!;

        /// <summary>
        /// Specifies a quoted header string that you want to insert into an HTTP request
        /// </summary>
        [Output("headInsert")]
        public Output<string> HeadInsert { get; private set; } = null!;

        /// <summary>
        /// When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        /// </summary>
        [Output("insertXforwardedFor")]
        public Output<string> InsertXforwardedFor { get; private set; } = null!;

        /// <summary>
        /// Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
        /// </summary>
        [Output("lwsSeparator")]
        public Output<string> LwsSeparator { get; private set; } = null!;

        /// <summary>
        /// Name of the profile_http
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
        /// </summary>
        [Output("oneconnectTransformations")]
        public Output<string> OneconnectTransformations { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of HTTP proxy.
        /// </summary>
        [Output("proxyType")]
        public Output<string> ProxyType { get; private set; } = null!;

        /// <summary>
        /// Specifies which of the application HTTP redirects the system rewrites to HTTPS.
        /// </summary>
        [Output("redirectRewrite")]
        public Output<string> RedirectRewrite { get; private set; } = null!;

        /// <summary>
        /// Specifies how to handle chunked and unchunked requests.
        /// </summary>
        [Output("requestChunking")]
        public Output<string> RequestChunking { get; private set; } = null!;

        /// <summary>
        /// Specifies how to handle chunked and unchunked responses.
        /// </summary>
        [Output("responseChunking")]
        public Output<string> ResponseChunking { get; private set; } = null!;

        /// <summary>
        /// Specifies headers that the BIG-IP system allows in an HTTP response.
        /// </summary>
        [Output("responseHeadersPermitteds")]
        public Output<ImmutableArray<string>> ResponseHeadersPermitteds { get; private set; } = null!;

        /// <summary>
        /// Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
        /// string is specified, then no Server header will be added to such responses
        /// </summary>
        [Output("serverAgentName")]
        public Output<string> ServerAgentName { get; private set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides.
        /// </summary>
        [Output("tmPartition")]
        public Output<string?> TmPartition { get; private set; } = null!;

        /// <summary>
        /// Specifies the hostname to include into Via header
        /// </summary>
        [Output("viaHostName")]
        public Output<string> ViaHostName { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to append, remove, or preserve a Via header in an HTTP request
        /// </summary>
        [Output("viaRequest")]
        public Output<string> ViaRequest { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to append, remove, or preserve a Via header in an HTTP request
        /// </summary>
        [Output("viaResponse")]
        public Output<string> ViaResponse { get; private set; } = null!;

        /// <summary>
        /// Specifies alternative XFF headers instead of the default X-forwarded-for header
        /// </summary>
        [Output("xffAlternativeNames")]
        public Output<ImmutableArray<string>> XffAlternativeNames { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileHttp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileHttp(string name, ProfileHttpArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileHttp:ProfileHttp", name, args ?? new ProfileHttpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileHttp(string name, Input<string> id, ProfileHttpState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileHttp:ProfileHttp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileHttp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileHttp Get(string name, Input<string> id, ProfileHttpState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileHttp(name, id, state, options);
        }
    }

    public sealed class ProfileHttpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
        /// XFF (X-forwarded-for) headers, if they exist.
        /// </summary>
        [Input("acceptXff")]
        public Input<string>? AcceptXff { get; set; }

        /// <summary>
        /// The application service to which the object belongs.
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        /// </summary>
        [Input("basicAuthRealm")]
        public Input<string>? BasicAuthRealm { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies a passphrase for the cookie encryption
        /// </summary>
        [Input("encryptCookieSecret")]
        public Input<string>? EncryptCookieSecret { get; set; }

        [Input("encryptCookies")]
        private InputList<string>? _encryptCookies;

        /// <summary>
        /// Encrypts specified cookies that the BIG-IP system sends to a client system
        /// </summary>
        public InputList<string> EncryptCookies
        {
            get => _encryptCookies ?? (_encryptCookies = new InputList<string>());
            set => _encryptCookies = value;
        }

        /// <summary>
        /// Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        /// </summary>
        [Input("fallbackHost")]
        public Input<string>? FallbackHost { get; set; }

        [Input("fallbackStatusCodes")]
        private InputList<string>? _fallbackStatusCodes;

        /// <summary>
        /// Specifies one or more three-digit status codes that can be returned by an HTTP server.
        /// </summary>
        public InputList<string> FallbackStatusCodes
        {
            get => _fallbackStatusCodes ?? (_fallbackStatusCodes = new InputList<string>());
            set => _fallbackStatusCodes = value;
        }

        /// <summary>
        /// Specifies the header string that you want to erase from an HTTP request. You can also specify none
        /// </summary>
        [Input("headErase")]
        public Input<string>? HeadErase { get; set; }

        /// <summary>
        /// Specifies a quoted header string that you want to insert into an HTTP request
        /// </summary>
        [Input("headInsert")]
        public Input<string>? HeadInsert { get; set; }

        /// <summary>
        /// When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        /// </summary>
        [Input("insertXforwardedFor")]
        public Input<string>? InsertXforwardedFor { get; set; }

        /// <summary>
        /// Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
        /// </summary>
        [Input("lwsSeparator")]
        public Input<string>? LwsSeparator { get; set; }

        /// <summary>
        /// Name of the profile_http
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
        /// </summary>
        [Input("oneconnectTransformations")]
        public Input<string>? OneconnectTransformations { get; set; }

        /// <summary>
        /// Specifies the type of HTTP proxy.
        /// </summary>
        [Input("proxyType")]
        public Input<string>? ProxyType { get; set; }

        /// <summary>
        /// Specifies which of the application HTTP redirects the system rewrites to HTTPS.
        /// </summary>
        [Input("redirectRewrite")]
        public Input<string>? RedirectRewrite { get; set; }

        /// <summary>
        /// Specifies how to handle chunked and unchunked requests.
        /// </summary>
        [Input("requestChunking")]
        public Input<string>? RequestChunking { get; set; }

        /// <summary>
        /// Specifies how to handle chunked and unchunked responses.
        /// </summary>
        [Input("responseChunking")]
        public Input<string>? ResponseChunking { get; set; }

        [Input("responseHeadersPermitteds")]
        private InputList<string>? _responseHeadersPermitteds;

        /// <summary>
        /// Specifies headers that the BIG-IP system allows in an HTTP response.
        /// </summary>
        public InputList<string> ResponseHeadersPermitteds
        {
            get => _responseHeadersPermitteds ?? (_responseHeadersPermitteds = new InputList<string>());
            set => _responseHeadersPermitteds = value;
        }

        /// <summary>
        /// Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
        /// string is specified, then no Server header will be added to such responses
        /// </summary>
        [Input("serverAgentName")]
        public Input<string>? ServerAgentName { get; set; }

        /// <summary>
        /// Displays the administrative partition within which this profile resides.
        /// </summary>
        [Input("tmPartition")]
        public Input<string>? TmPartition { get; set; }

        /// <summary>
        /// Specifies the hostname to include into Via header
        /// </summary>
        [Input("viaHostName")]
        public Input<string>? ViaHostName { get; set; }

        /// <summary>
        /// Specifies whether to append, remove, or preserve a Via header in an HTTP request
        /// </summary>
        [Input("viaRequest")]
        public Input<string>? ViaRequest { get; set; }

        /// <summary>
        /// Specifies whether to append, remove, or preserve a Via header in an HTTP request
        /// </summary>
        [Input("viaResponse")]
        public Input<string>? ViaResponse { get; set; }

        [Input("xffAlternativeNames")]
        private InputList<string>? _xffAlternativeNames;

        /// <summary>
        /// Specifies alternative XFF headers instead of the default X-forwarded-for header
        /// </summary>
        public InputList<string> XffAlternativeNames
        {
            get => _xffAlternativeNames ?? (_xffAlternativeNames = new InputList<string>());
            set => _xffAlternativeNames = value;
        }

        public ProfileHttpArgs()
        {
        }
    }

    public sealed class ProfileHttpState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables or disables trusting the client IP address, and statistics from the client IP address, based on the request's
        /// XFF (X-forwarded-for) headers, if they exist.
        /// </summary>
        [Input("acceptXff")]
        public Input<string>? AcceptXff { get; set; }

        /// <summary>
        /// The application service to which the object belongs.
        /// </summary>
        [Input("appService")]
        public Input<string>? AppService { get; set; }

        /// <summary>
        /// Specifies a quoted string for the basic authentication realm. The system sends this string to a client whenever authorization fails. The default value is none
        /// </summary>
        [Input("basicAuthRealm")]
        public Input<string>? BasicAuthRealm { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// User defined description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies a passphrase for the cookie encryption
        /// </summary>
        [Input("encryptCookieSecret")]
        public Input<string>? EncryptCookieSecret { get; set; }

        [Input("encryptCookies")]
        private InputList<string>? _encryptCookies;

        /// <summary>
        /// Encrypts specified cookies that the BIG-IP system sends to a client system
        /// </summary>
        public InputList<string> EncryptCookies
        {
            get => _encryptCookies ?? (_encryptCookies = new InputList<string>());
            set => _encryptCookies = value;
        }

        /// <summary>
        /// Specifies an HTTP fallback host. HTTP redirection allows you to redirect HTTP traffic to another protocol identifier, host name, port number
        /// </summary>
        [Input("fallbackHost")]
        public Input<string>? FallbackHost { get; set; }

        [Input("fallbackStatusCodes")]
        private InputList<string>? _fallbackStatusCodes;

        /// <summary>
        /// Specifies one or more three-digit status codes that can be returned by an HTTP server.
        /// </summary>
        public InputList<string> FallbackStatusCodes
        {
            get => _fallbackStatusCodes ?? (_fallbackStatusCodes = new InputList<string>());
            set => _fallbackStatusCodes = value;
        }

        /// <summary>
        /// Specifies the header string that you want to erase from an HTTP request. You can also specify none
        /// </summary>
        [Input("headErase")]
        public Input<string>? HeadErase { get; set; }

        /// <summary>
        /// Specifies a quoted header string that you want to insert into an HTTP request
        /// </summary>
        [Input("headInsert")]
        public Input<string>? HeadInsert { get; set; }

        /// <summary>
        /// When using connection pooling, which allows clients to make use of other client requests' server-side connections, you can insert the X-Forwarded-For header and specify a client IP address
        /// </summary>
        [Input("insertXforwardedFor")]
        public Input<string>? InsertXforwardedFor { get; set; }

        /// <summary>
        /// Specifies a quoted header string that you want to insert into an HTTP request. You can also specify none.
        /// </summary>
        [Input("lwsSeparator")]
        public Input<string>? LwsSeparator { get; set; }

        /// <summary>
        /// Name of the profile_http
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Enables the system to perform HTTP header transformations for the purpose of  keeping server-side connections open. This feature requires configuration of a OneConnect profile
        /// </summary>
        [Input("oneconnectTransformations")]
        public Input<string>? OneconnectTransformations { get; set; }

        /// <summary>
        /// Specifies the type of HTTP proxy.
        /// </summary>
        [Input("proxyType")]
        public Input<string>? ProxyType { get; set; }

        /// <summary>
        /// Specifies which of the application HTTP redirects the system rewrites to HTTPS.
        /// </summary>
        [Input("redirectRewrite")]
        public Input<string>? RedirectRewrite { get; set; }

        /// <summary>
        /// Specifies how to handle chunked and unchunked requests.
        /// </summary>
        [Input("requestChunking")]
        public Input<string>? RequestChunking { get; set; }

        /// <summary>
        /// Specifies how to handle chunked and unchunked responses.
        /// </summary>
        [Input("responseChunking")]
        public Input<string>? ResponseChunking { get; set; }

        [Input("responseHeadersPermitteds")]
        private InputList<string>? _responseHeadersPermitteds;

        /// <summary>
        /// Specifies headers that the BIG-IP system allows in an HTTP response.
        /// </summary>
        public InputList<string> ResponseHeadersPermitteds
        {
            get => _responseHeadersPermitteds ?? (_responseHeadersPermitteds = new InputList<string>());
            set => _responseHeadersPermitteds = value;
        }

        /// <summary>
        /// Specifies the value of the Server header in responses that the BIG-IP itself generates. The default is BigIP. If no
        /// string is specified, then no Server header will be added to such responses
        /// </summary>
        [Input("serverAgentName")]
        public Input<string>? ServerAgentName { get; set; }

        /// <summary>
        /// Displays the administrative partition within which this profile resides.
        /// </summary>
        [Input("tmPartition")]
        public Input<string>? TmPartition { get; set; }

        /// <summary>
        /// Specifies the hostname to include into Via header
        /// </summary>
        [Input("viaHostName")]
        public Input<string>? ViaHostName { get; set; }

        /// <summary>
        /// Specifies whether to append, remove, or preserve a Via header in an HTTP request
        /// </summary>
        [Input("viaRequest")]
        public Input<string>? ViaRequest { get; set; }

        /// <summary>
        /// Specifies whether to append, remove, or preserve a Via header in an HTTP request
        /// </summary>
        [Input("viaResponse")]
        public Input<string>? ViaResponse { get; set; }

        [Input("xffAlternativeNames")]
        private InputList<string>? _xffAlternativeNames;

        /// <summary>
        /// Specifies alternative XFF headers instead of the default X-forwarded-for header
        /// </summary>
        public InputList<string> XffAlternativeNames
        {
            get => _xffAlternativeNames ?? (_xffAlternativeNames = new InputList<string>());
            set => _xffAlternativeNames = value;
        }

        public ProfileHttpState()
        {
        }
    }
}
