// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.F5BigIP.Net
{
    /// <summary>
    /// `f5bigip.net.Route` Manages a route configuration
    /// 
    /// For resources should be named with their "full path". The full path is the combination of the partition + name of the resource. For example /Common/my-pool.
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
    ///     var route2 = new F5BigIP.Net.Route("route2", new()
    ///     {
    ///         Name = "/Common/external-route",
    ///         Network = "10.10.10.0/24",
    ///         Gw = "1.1.1.2",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:net/route:Route")]
    public partial class Route : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies a gateway address for the route.
        /// </summary>
        [Output("gw")]
        public Output<string?> Gw { get; private set; } = null!;

        /// <summary>
        /// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The destination subnet and netmask for the route.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// reject route
        /// </summary>
        [Output("reject")]
        public Output<bool?> Reject { get; private set; } = null!;

        /// <summary>
        /// tunnel_ref to route traffic
        /// </summary>
        [Output("tunnelRef")]
        public Output<string?> TunnelRef { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:net/route:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:net/route:Route", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, RouteState? state = null, CustomResourceOptions? options = null)
        {
            return new Route(name, id, state, options);
        }
    }

    public sealed class RouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a gateway address for the route.
        /// </summary>
        [Input("gw")]
        public Input<string>? Gw { get; set; }

        /// <summary>
        /// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The destination subnet and netmask for the route.
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// reject route
        /// </summary>
        [Input("reject")]
        public Input<bool>? Reject { get; set; }

        /// <summary>
        /// tunnel_ref to route traffic
        /// </summary>
        [Input("tunnelRef")]
        public Input<string>? TunnelRef { get; set; }

        public RouteArgs()
        {
        }
        public static new RouteArgs Empty => new RouteArgs();
    }

    public sealed class RouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies a gateway address for the route.
        /// </summary>
        [Input("gw")]
        public Input<string>? Gw { get; set; }

        /// <summary>
        /// Name of the route.Name of Route should be full path,full path is the combination of the `partition + route name`,For ex: `/Common/test-net-route`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The destination subnet and netmask for the route.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// reject route
        /// </summary>
        [Input("reject")]
        public Input<bool>? Reject { get; set; }

        /// <summary>
        /// tunnel_ref to route traffic
        /// </summary>
        [Input("tunnelRef")]
        public Input<string>? TunnelRef { get; set; }

        public RouteState()
        {
        }
        public static new RouteState Empty => new RouteState();
    }
}
