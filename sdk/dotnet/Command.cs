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
    /// `f5bigip.Command` Run TMSH commands on F5 devices
    /// 
    /// This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device
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
    ///         //create ltm node
    ///         var test_command = new F5BigIP.Command("test-command", new F5BigIP.CommandArgs
    ///         {
    ///             Commands = 
    ///             {
    ///                 "delete ltm node 10.10.10.70",
    ///             },
    ///             When = "destroy",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [F5BigIPResourceType("f5bigip:index/command:Command")]
    public partial class Command : Pulumi.CustomResource
    {
        /// <summary>
        /// The resulting output from the `commands` executed
        /// </summary>
        [Output("commandResults")]
        public Output<ImmutableArray<string>> CommandResults { get; private set; } = null!;

        /// <summary>
        /// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        /// </summary>
        [Output("commands")]
        public Output<ImmutableArray<string>> Commands { get; private set; } = null!;

        [Output("when")]
        public Output<string?> When { get; private set; } = null!;


        /// <summary>
        /// Create a Command resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Command(string name, CommandArgs args, CustomResourceOptions? options = null)
            : base("f5bigip:index/command:Command", name, args ?? new CommandArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Command(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
            : base("f5bigip:index/command:Command", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Command resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Command Get(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
        {
            return new Command(name, id, state, options);
        }
    }

    public sealed class CommandArgs : Pulumi.ResourceArgs
    {
        [Input("commandResults")]
        private InputList<string>? _commandResults;

        /// <summary>
        /// The resulting output from the `commands` executed
        /// </summary>
        public InputList<string> CommandResults
        {
            get => _commandResults ?? (_commandResults = new InputList<string>());
            set => _commandResults = value;
        }

        [Input("commands", required: true)]
        private InputList<string>? _commands;

        /// <summary>
        /// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("when")]
        public Input<string>? When { get; set; }

        public CommandArgs()
        {
        }
    }

    public sealed class CommandState : Pulumi.ResourceArgs
    {
        [Input("commandResults")]
        private InputList<string>? _commandResults;

        /// <summary>
        /// The resulting output from the `commands` executed
        /// </summary>
        public InputList<string> CommandResults
        {
            get => _commandResults ?? (_commandResults = new InputList<string>());
            set => _commandResults = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `command_result`
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("when")]
        public Input<string>? When { get; set; }

        public CommandState()
        {
        }
    }
}
