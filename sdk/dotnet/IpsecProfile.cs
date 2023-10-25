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
    /// `f5bigip.IpsecProfile` Manage IPSec Profiles on a BIG-IP
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
    ///     var azurevWANProfile = new F5BigIP.IpsecProfile("azurevWANProfile", new()
    ///     {
    ///         Description = "mytestipsecprofile",
    ///         Name = "/Common/Mytestipsecprofile",
    ///         TrafficSelector = "test-trafficselector",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/ipsecProfile:IpsecProfile")]
    public partial class IpsecProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies descriptive text that identifies the IPsec interface tunnel profile.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        /// </summary>
        [Output("parentProfile")]
        public Output<string?> ParentProfile { get; private set; } = null!;

        /// <summary>
        /// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        /// </summary>
        [Output("trafficSelector")]
        public Output<string> TrafficSelector { get; private set; } = null!;


        /// <summary>
        /// Create a IpsecProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpsecProfile(string name, IpsecProfileArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/ipsecProfile:IpsecProfile", name, args ?? new IpsecProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpsecProfile(string name, Input<string> id, IpsecProfileState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/ipsecProfile:IpsecProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpsecProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpsecProfile Get(string name, Input<string> id, IpsecProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new IpsecProfile(name, id, state, options);
        }
    }

    public sealed class IpsecProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies descriptive text that identifies the IPsec interface tunnel profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        /// </summary>
        [Input("parentProfile")]
        public Input<string>? ParentProfile { get; set; }

        /// <summary>
        /// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        /// </summary>
        [Input("trafficSelector")]
        public Input<string>? TrafficSelector { get; set; }

        public IpsecProfileArgs()
        {
        }
        public static new IpsecProfileArgs Empty => new IpsecProfileArgs();
    }

    public sealed class IpsecProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies descriptive text that identifies the IPsec interface tunnel profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Displays the name of the IPsec interface tunnel profile,it should be "full path".The full path is the combination of the partition + name of the IPSec profile.(For example `/Common/test-profile`)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the profile from which this profile inherits settings. The default is the system-supplied `/Common/ipsec` profile
        /// </summary>
        [Input("parentProfile")]
        public Input<string>? ParentProfile { get; set; }

        /// <summary>
        /// Specifies the traffic selector for the IPsec interface tunnel to which the profile is applied
        /// </summary>
        [Input("trafficSelector")]
        public Input<string>? TrafficSelector { get; set; }

        public IpsecProfileState()
        {
        }
        public static new IpsecProfileState Empty => new IpsecProfileState();
    }
}
