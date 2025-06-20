// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Ltm
{
    /// <summary>
    /// `f5bigip.ltm.ProfileHttpCompress`  Virtual server HTTP compression profile configuration
    /// 
    /// Resources should be named with their `full path`.The full path is the combination of the `partition + name` (example: `/Common/my-httpcompresprofile` ) or  `partition + directory + name` of the resource  (example: `/Common/test/my-httpcompresprofile`)
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
    ///     var sjhttpcompression = new F5BigIP.Ltm.ProfileHttpCompress("sjhttpcompression", new()
    ///     {
    ///         Name = "/Common/sjhttpcompression2",
    ///         DefaultsFrom = "/Common/httpcompression",
    ///         UriExcludes = new[]
    ///         {
    ///             "www.abc.f5.com",
    ///             "www.abc2.f5.com",
    ///         },
    ///         UriIncludes = new[]
    ///         {
    ///             "www.xyzbc.cisco.com",
    ///         },
    ///         ContentTypeIncludes = new[]
    ///         {
    ///             "nicecontent.com",
    ///         },
    ///         ContentTypeExcludes = new[]
    ///         {
    ///             "nicecontentexclude.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BIG-IP LTM HTTP Compress profiles can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import f5bigip:ltm/profileHttpCompress:ProfileHttpCompress test-httpcomprs_import /Common/test-httpcomprs
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:ltm/profileHttpCompress:ProfileHttpCompress")]
    public partial class ProfileHttpCompress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the maximum number of compressed bytes that the system buffers before inserting a Content-Length header (which specifies the compressed size) into the response. The default is `4096` bytes.
        /// </summary>
        [Output("compressionBuffersize")]
        public Output<int> CompressionBuffersize { get; private set; } = null!;

        /// <summary>
        /// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        /// </summary>
        [Output("contentTypeExcludes")]
        public Output<ImmutableArray<string>> ContentTypeExcludes { get; private set; } = null!;

        /// <summary>
        /// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        /// </summary>
        [Output("contentTypeIncludes")]
        public Output<ImmutableArray<string>> ContentTypeIncludes { get; private set; } = null!;

        /// <summary>
        /// Specifies, when checked (enabled), that the system monitors the percent CPU usage and adjusts compression rates automatically when the CPU usage reaches either the CPU Saver High Threshold or the CPU Saver Low Threshold. The default is `enabled`.
        /// </summary>
        [Output("cpuSaver")]
        public Output<string> CpuSaver { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specifies the degree to which the system compresses the content. Higher compression levels cause the compression process to be slower. The default is 1 - Least Compression (Fastest)
        /// </summary>
        [Output("gzipCompressionLevel")]
        public Output<int> GzipCompressionLevel { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of bytes of memory that the system uses for internal compression buffers when compressing a server response. The default is `8 kilobytes/8192 bytes`.
        /// </summary>
        [Output("gzipMemoryLevel")]
        public Output<int> GzipMemoryLevel { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The default is `16` kilobytes
        /// </summary>
        [Output("gzipWindowSize")]
        public Output<int> GzipWindowSize { get; private set; } = null!;

        /// <summary>
        /// Specifies, when checked (enabled), that the system does not remove the Accept-Encoding: header from an HTTP request. The default is `disabled`.
        /// </summary>
        [Output("keepAcceptEncoding")]
        public Output<string> KeepAcceptEncoding { get; private set; } = null!;

        /// <summary>
        /// Name of the LTM http compress profile,named with their `full path`.The full path is the combination of the `partition + name` (example: `/Common/my-httpcompresprofile` ) or  `partition + directory + name` of the resource  (example: `my-httpcompresprofile`)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
        /// </summary>
        [Output("uriExcludes")]
        public Output<ImmutableArray<string>> UriExcludes { get; private set; } = null!;

        /// <summary>
        /// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
        /// </summary>
        [Output("uriIncludes")]
        public Output<ImmutableArray<string>> UriIncludes { get; private set; } = null!;

        /// <summary>
        /// Specifies, when checked (enabled), that the system inserts a Vary header into cacheable server responses. The default is `enabled`.
        /// </summary>
        [Output("varyHeader")]
        public Output<string> VaryHeader { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileHttpCompress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileHttpCompress(string name, ProfileHttpCompressArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileHttpCompress:ProfileHttpCompress", name, args ?? new ProfileHttpCompressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileHttpCompress(string name, Input<string> id, ProfileHttpCompressState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileHttpCompress:ProfileHttpCompress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileHttpCompress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileHttpCompress Get(string name, Input<string> id, ProfileHttpCompressState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileHttpCompress(name, id, state, options);
        }
    }

    public sealed class ProfileHttpCompressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the maximum number of compressed bytes that the system buffers before inserting a Content-Length header (which specifies the compressed size) into the response. The default is `4096` bytes.
        /// </summary>
        [Input("compressionBuffersize")]
        public Input<int>? CompressionBuffersize { get; set; }

        [Input("contentTypeExcludes")]
        private InputList<string>? _contentTypeExcludes;

        /// <summary>
        /// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        /// </summary>
        public InputList<string> ContentTypeExcludes
        {
            get => _contentTypeExcludes ?? (_contentTypeExcludes = new InputList<string>());
            set => _contentTypeExcludes = value;
        }

        [Input("contentTypeIncludes")]
        private InputList<string>? _contentTypeIncludes;

        /// <summary>
        /// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        /// </summary>
        public InputList<string> ContentTypeIncludes
        {
            get => _contentTypeIncludes ?? (_contentTypeIncludes = new InputList<string>());
            set => _contentTypeIncludes = value;
        }

        /// <summary>
        /// Specifies, when checked (enabled), that the system monitors the percent CPU usage and adjusts compression rates automatically when the CPU usage reaches either the CPU Saver High Threshold or the CPU Saver Low Threshold. The default is `enabled`.
        /// </summary>
        [Input("cpuSaver")]
        public Input<string>? CpuSaver { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies the degree to which the system compresses the content. Higher compression levels cause the compression process to be slower. The default is 1 - Least Compression (Fastest)
        /// </summary>
        [Input("gzipCompressionLevel")]
        public Input<int>? GzipCompressionLevel { get; set; }

        /// <summary>
        /// Specifies the number of bytes of memory that the system uses for internal compression buffers when compressing a server response. The default is `8 kilobytes/8192 bytes`.
        /// </summary>
        [Input("gzipMemoryLevel")]
        public Input<int>? GzipMemoryLevel { get; set; }

        /// <summary>
        /// Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The default is `16` kilobytes
        /// </summary>
        [Input("gzipWindowSize")]
        public Input<int>? GzipWindowSize { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system does not remove the Accept-Encoding: header from an HTTP request. The default is `disabled`.
        /// </summary>
        [Input("keepAcceptEncoding")]
        public Input<string>? KeepAcceptEncoding { get; set; }

        /// <summary>
        /// Name of the LTM http compress profile,named with their `full path`.The full path is the combination of the `partition + name` (example: `/Common/my-httpcompresprofile` ) or  `partition + directory + name` of the resource  (example: `my-httpcompresprofile`)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("uriExcludes")]
        private InputList<string>? _uriExcludes;

        /// <summary>
        /// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
        /// </summary>
        public InputList<string> UriExcludes
        {
            get => _uriExcludes ?? (_uriExcludes = new InputList<string>());
            set => _uriExcludes = value;
        }

        [Input("uriIncludes")]
        private InputList<string>? _uriIncludes;

        /// <summary>
        /// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
        /// </summary>
        public InputList<string> UriIncludes
        {
            get => _uriIncludes ?? (_uriIncludes = new InputList<string>());
            set => _uriIncludes = value;
        }

        /// <summary>
        /// Specifies, when checked (enabled), that the system inserts a Vary header into cacheable server responses. The default is `enabled`.
        /// </summary>
        [Input("varyHeader")]
        public Input<string>? VaryHeader { get; set; }

        public ProfileHttpCompressArgs()
        {
        }
        public static new ProfileHttpCompressArgs Empty => new ProfileHttpCompressArgs();
    }

    public sealed class ProfileHttpCompressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the maximum number of compressed bytes that the system buffers before inserting a Content-Length header (which specifies the compressed size) into the response. The default is `4096` bytes.
        /// </summary>
        [Input("compressionBuffersize")]
        public Input<int>? CompressionBuffersize { get; set; }

        [Input("contentTypeExcludes")]
        private InputList<string>? _contentTypeExcludes;

        /// <summary>
        /// Excludes a specified list of content types from compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        /// </summary>
        public InputList<string> ContentTypeExcludes
        {
            get => _contentTypeExcludes ?? (_contentTypeExcludes = new InputList<string>());
            set => _contentTypeExcludes = value;
        }

        [Input("contentTypeIncludes")]
        private InputList<string>? _contentTypeIncludes;

        /// <summary>
        /// Specifies a list of content types for compression of HTTP Content-Type responses. Use a string list to specify a list of content types you want to compress.
        /// </summary>
        public InputList<string> ContentTypeIncludes
        {
            get => _contentTypeIncludes ?? (_contentTypeIncludes = new InputList<string>());
            set => _contentTypeIncludes = value;
        }

        /// <summary>
        /// Specifies, when checked (enabled), that the system monitors the percent CPU usage and adjusts compression rates automatically when the CPU usage reaches either the CPU Saver High Threshold or the CPU Saver Low Threshold. The default is `enabled`.
        /// </summary>
        [Input("cpuSaver")]
        public Input<string>? CpuSaver { get; set; }

        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies the degree to which the system compresses the content. Higher compression levels cause the compression process to be slower. The default is 1 - Least Compression (Fastest)
        /// </summary>
        [Input("gzipCompressionLevel")]
        public Input<int>? GzipCompressionLevel { get; set; }

        /// <summary>
        /// Specifies the number of bytes of memory that the system uses for internal compression buffers when compressing a server response. The default is `8 kilobytes/8192 bytes`.
        /// </summary>
        [Input("gzipMemoryLevel")]
        public Input<int>? GzipMemoryLevel { get; set; }

        /// <summary>
        /// Specifies the number of kilobytes in the window size that the system uses when compressing a server response. The default is `16` kilobytes
        /// </summary>
        [Input("gzipWindowSize")]
        public Input<int>? GzipWindowSize { get; set; }

        /// <summary>
        /// Specifies, when checked (enabled), that the system does not remove the Accept-Encoding: header from an HTTP request. The default is `disabled`.
        /// </summary>
        [Input("keepAcceptEncoding")]
        public Input<string>? KeepAcceptEncoding { get; set; }

        /// <summary>
        /// Name of the LTM http compress profile,named with their `full path`.The full path is the combination of the `partition + name` (example: `/Common/my-httpcompresprofile` ) or  `partition + directory + name` of the resource  (example: `my-httpcompresprofile`)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("uriExcludes")]
        private InputList<string>? _uriExcludes;

        /// <summary>
        /// Disables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you do not want to compress.
        /// </summary>
        public InputList<string> UriExcludes
        {
            get => _uriExcludes ?? (_uriExcludes = new InputList<string>());
            set => _uriExcludes = value;
        }

        [Input("uriIncludes")]
        private InputList<string>? _uriIncludes;

        /// <summary>
        /// Enables compression on a specified list of HTTP Request-URI responses. Use a regular expression to specify a list of URIs you want to compress.
        /// </summary>
        public InputList<string> UriIncludes
        {
            get => _uriIncludes ?? (_uriIncludes = new InputList<string>());
            set => _uriIncludes = value;
        }

        /// <summary>
        /// Specifies, when checked (enabled), that the system inserts a Vary header into cacheable server responses. The default is `enabled`.
        /// </summary>
        [Input("varyHeader")]
        public Input<string>? VaryHeader { get; set; }

        public ProfileHttpCompressState()
        {
        }
        public static new ProfileHttpCompressState Empty => new ProfileHttpCompressState();
    }
}
