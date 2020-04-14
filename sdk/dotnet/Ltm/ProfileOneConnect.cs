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
    /// `f5bigip.ltm.ProfileOneConnect` Configures a custom profile_oneconnect for use by health checks.
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
    /// </summary>
    public partial class ProfileOneConnect : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Output("defaultsFrom")]
        public Output<string?> DefaultsFrom { get; private set; } = null!;

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
        /// </summary>
        [Output("idleTimeoutOverride")]
        public Output<string?> IdleTimeoutOverride { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
        /// </summary>
        [Output("maxAge")]
        public Output<int?> MaxAge { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
        /// </summary>
        [Output("maxReuse")]
        public Output<int?> MaxReuse { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
        /// </summary>
        [Output("maxSize")]
        public Output<int?> MaxSize { get; private set; } = null!;

        /// <summary>
        /// Name of the profile_oneconnect
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Output("partition")]
        public Output<string?> Partition { get; private set; } = null!;

        /// <summary>
        /// Specify if you want to share the pool, default value is "disabled"
        /// </summary>
        [Output("sharePools")]
        public Output<string?> SharePools { get; private set; } = null!;

        /// <summary>
        /// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        /// </summary>
        [Output("sourceMask")]
        public Output<string?> SourceMask { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileOneConnect resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileOneConnect(string name, ProfileOneConnectArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileOneConnect:ProfileOneConnect", name, args ?? new ProfileOneConnectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileOneConnect(string name, Input<string> id, ProfileOneConnectState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:ltm/profileOneConnect:ProfileOneConnect", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProfileOneConnect resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileOneConnect Get(string name, Input<string> id, ProfileOneConnectState? state = null, CustomResourceOptions? options = null)
        {
            return new ProfileOneConnect(name, id, state, options);
        }
    }

    public sealed class ProfileOneConnectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
        /// </summary>
        [Input("idleTimeoutOverride")]
        public Input<string>? IdleTimeoutOverride { get; set; }

        /// <summary>
        /// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
        /// </summary>
        [Input("maxReuse")]
        public Input<int>? MaxReuse { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
        /// </summary>
        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        /// <summary>
        /// Name of the profile_oneconnect
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specify if you want to share the pool, default value is "disabled"
        /// </summary>
        [Input("sharePools")]
        public Input<string>? SharePools { get; set; }

        /// <summary>
        /// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        /// </summary>
        [Input("sourceMask")]
        public Input<string>? SourceMask { get; set; }

        public ProfileOneConnectArgs()
        {
        }
    }

    public sealed class ProfileOneConnectState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the profile that you want to use as the parent profile. Your new profile inherits all settings and values from the parent profile specified.
        /// </summary>
        [Input("defaultsFrom")]
        public Input<string>? DefaultsFrom { get; set; }

        /// <summary>
        /// Specifies the number of seconds that a connection is idle before the connection flow is eligible for deletion. Possible values are disabled, indefinite, or a numeric value that you specify. The default value is disabled.
        /// </summary>
        [Input("idleTimeoutOverride")]
        public Input<string>? IdleTimeoutOverride { get; set; }

        /// <summary>
        /// Specifies the maximum age in number of seconds allowed for a connection in the connection reuse pool. For any connection with an age higher than this value, the system removes that connection from the reuse pool. The default value is 86400.
        /// </summary>
        [Input("maxAge")]
        public Input<int>? MaxAge { get; set; }

        /// <summary>
        /// Specifies the maximum number of times that a server-side connection can be reused. The default value is 1000.
        /// </summary>
        [Input("maxReuse")]
        public Input<int>? MaxReuse { get; set; }

        /// <summary>
        /// Specifies the maximum number of connections that the system holds in the connection reuse pool. If the pool is already full, then the server-side connection closes after the response is completed. The default value is 10000.
        /// </summary>
        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        /// <summary>
        /// Name of the profile_oneconnect
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Displays the administrative partition within which this profile resides
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// Specify if you want to share the pool, default value is "disabled"
        /// </summary>
        [Input("sharePools")]
        public Input<string>? SharePools { get; set; }

        /// <summary>
        /// Specifies a source IP mask. The default value is 0.0.0.0. The system applies the value of this option to the source address to determine its eligibility for reuse. A mask of 0.0.0.0 causes the system to share reused connections across all clients. A host mask (all 1's in binary), causes the system to share only those reused connections originating from the same client IP address.
        /// </summary>
        [Input("sourceMask")]
        public Input<string>? SourceMask { get; set; }

        public ProfileOneConnectState()
        {
        }
    }
}
