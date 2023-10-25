// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `f5bigip.Command` Run TMSH commands on F5 devices
 *
 * This resource is helpful to send TMSH command to an BIG-IP node and returns the results read from the device
 */
export class Command extends pulumi.CustomResource {
    /**
     * Get an existing Command resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CommandState, opts?: pulumi.CustomResourceOptions): Command {
        return new Command(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:index/command:Command';

    /**
     * Returns true if the given object is an instance of Command.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Command {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Command.__pulumiType;
    }

    /**
     * The resulting output from the `commands` executed
     */
    public readonly commandResults!: pulumi.Output<string[]>;
    /**
     * The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
     */
    public readonly commands!: pulumi.Output<string[]>;
    public readonly when!: pulumi.Output<string | undefined>;

    /**
     * Create a Command resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CommandArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CommandArgs | CommandState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CommandState | undefined;
            resourceInputs["commandResults"] = state ? state.commandResults : undefined;
            resourceInputs["commands"] = state ? state.commands : undefined;
            resourceInputs["when"] = state ? state.when : undefined;
        } else {
            const args = argsOrState as CommandArgs | undefined;
            if ((!args || args.commands === undefined) && !opts.urn) {
                throw new Error("Missing required property 'commands'");
            }
            resourceInputs["commandResults"] = args ? args.commandResults : undefined;
            resourceInputs["commands"] = args ? args.commands : undefined;
            resourceInputs["when"] = args ? args.when : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Command.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Command resources.
 */
export interface CommandState {
    /**
     * The resulting output from the `commands` executed
     */
    commandResults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
     */
    commands?: pulumi.Input<pulumi.Input<string>[]>;
    when?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Command resource.
 */
export interface CommandArgs {
    /**
     * The resulting output from the `commands` executed
     */
    commandResults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The commands to send to the remote BIG-IP device over the configured provider. The resulting output from the command is returned and added to `commandResult`
     */
    commands: pulumi.Input<pulumi.Input<string>[]>;
    when?: pulumi.Input<string>;
}
