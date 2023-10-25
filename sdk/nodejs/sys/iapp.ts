// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.
 */
export class IApp extends pulumi.CustomResource {
    /**
     * Get an existing IApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IAppState, opts?: pulumi.CustomResourceOptions): IApp {
        return new IApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:sys/iApp:IApp';

    /**
     * Returns true if the given object is an instance of IApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IApp.__pulumiType;
    }

    /**
     * User defined description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly devicegroup!: pulumi.Output<string>;
    /**
     * Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
     */
    public readonly executeAction!: pulumi.Output<string>;
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
     */
    public readonly inheritedDevicegroup!: pulumi.Output<string | undefined>;
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
     */
    public readonly inheritedTrafficGroup!: pulumi.Output<string | undefined>;
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     */
    public readonly jsonfile!: pulumi.Output<string>;
    /**
     * string values
     */
    public readonly lists!: pulumi.Output<outputs.sys.IAppList[] | undefined>;
    /**
     * User defined generic data for the application service. It is a name and value pair.
     */
    public readonly metadatas!: pulumi.Output<outputs.sys.IAppMetadata[] | undefined>;
    /**
     * Name of the iApp.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Displays the administrative partition within which the application resides.
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
     */
    public readonly strictUpdates!: pulumi.Output<string | undefined>;
    public readonly tables!: pulumi.Output<outputs.sys.IAppTable[] | undefined>;
    /**
     * The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
     */
    public readonly template!: pulumi.Output<string | undefined>;
    /**
     * Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
     */
    public readonly templateModified!: pulumi.Output<string | undefined>;
    /**
     * Indicates any missing prerequisites associated with the template that defines this application.
     */
    public readonly templatePrerequisiteErrors!: pulumi.Output<string | undefined>;
    /**
     * The name of the traffic group that the application service is assigned to.
     */
    public readonly trafficGroup!: pulumi.Output<string | undefined>;
    public readonly variables!: pulumi.Output<outputs.sys.IAppVariable[] | undefined>;

    /**
     * Create a IApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAppArgs | IAppState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IAppState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["devicegroup"] = state ? state.devicegroup : undefined;
            resourceInputs["executeAction"] = state ? state.executeAction : undefined;
            resourceInputs["inheritedDevicegroup"] = state ? state.inheritedDevicegroup : undefined;
            resourceInputs["inheritedTrafficGroup"] = state ? state.inheritedTrafficGroup : undefined;
            resourceInputs["jsonfile"] = state ? state.jsonfile : undefined;
            resourceInputs["lists"] = state ? state.lists : undefined;
            resourceInputs["metadatas"] = state ? state.metadatas : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["strictUpdates"] = state ? state.strictUpdates : undefined;
            resourceInputs["tables"] = state ? state.tables : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
            resourceInputs["templateModified"] = state ? state.templateModified : undefined;
            resourceInputs["templatePrerequisiteErrors"] = state ? state.templatePrerequisiteErrors : undefined;
            resourceInputs["trafficGroup"] = state ? state.trafficGroup : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as IAppArgs | undefined;
            if ((!args || args.jsonfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jsonfile'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["devicegroup"] = args ? args.devicegroup : undefined;
            resourceInputs["executeAction"] = args ? args.executeAction : undefined;
            resourceInputs["inheritedDevicegroup"] = args ? args.inheritedDevicegroup : undefined;
            resourceInputs["inheritedTrafficGroup"] = args ? args.inheritedTrafficGroup : undefined;
            resourceInputs["jsonfile"] = args ? args.jsonfile : undefined;
            resourceInputs["lists"] = args ? args.lists : undefined;
            resourceInputs["metadatas"] = args ? args.metadatas : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["strictUpdates"] = args ? args.strictUpdates : undefined;
            resourceInputs["tables"] = args ? args.tables : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["templateModified"] = args ? args.templateModified : undefined;
            resourceInputs["templatePrerequisiteErrors"] = args ? args.templatePrerequisiteErrors : undefined;
            resourceInputs["trafficGroup"] = args ? args.trafficGroup : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IApp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IApp resources.
 */
export interface IAppState {
    /**
     * User defined description.
     */
    description?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    devicegroup?: pulumi.Input<string>;
    /**
     * Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
     */
    executeAction?: pulumi.Input<string>;
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
     */
    inheritedDevicegroup?: pulumi.Input<string>;
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
     */
    inheritedTrafficGroup?: pulumi.Input<string>;
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     */
    jsonfile?: pulumi.Input<string>;
    /**
     * string values
     */
    lists?: pulumi.Input<pulumi.Input<inputs.sys.IAppList>[]>;
    /**
     * User defined generic data for the application service. It is a name and value pair.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.sys.IAppMetadata>[]>;
    /**
     * Name of the iApp.
     */
    name?: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which the application resides.
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
     */
    strictUpdates?: pulumi.Input<string>;
    tables?: pulumi.Input<pulumi.Input<inputs.sys.IAppTable>[]>;
    /**
     * The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
     */
    template?: pulumi.Input<string>;
    /**
     * Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
     */
    templateModified?: pulumi.Input<string>;
    /**
     * Indicates any missing prerequisites associated with the template that defines this application.
     */
    templatePrerequisiteErrors?: pulumi.Input<string>;
    /**
     * The name of the traffic group that the application service is assigned to.
     */
    trafficGroup?: pulumi.Input<string>;
    variables?: pulumi.Input<pulumi.Input<inputs.sys.IAppVariable>[]>;
}

/**
 * The set of arguments for constructing a IApp resource.
 */
export interface IAppArgs {
    /**
     * User defined description.
     */
    description?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    devicegroup?: pulumi.Input<string>;
    /**
     * Run the specified template action associated with the application, this option can be specified in `json` with `executeAction`, value specified with `executeAction` attribute take precedence over `json` value
     */
    executeAction?: pulumi.Input<string>;
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
     */
    inheritedDevicegroup?: pulumi.Input<string>;
    /**
     * Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
     */
    inheritedTrafficGroup?: pulumi.Input<string>;
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     */
    jsonfile: pulumi.Input<string>;
    /**
     * string values
     */
    lists?: pulumi.Input<pulumi.Input<inputs.sys.IAppList>[]>;
    /**
     * User defined generic data for the application service. It is a name and value pair.
     */
    metadatas?: pulumi.Input<pulumi.Input<inputs.sys.IAppMetadata>[]>;
    /**
     * Name of the iApp.
     */
    name: pulumi.Input<string>;
    /**
     * Displays the administrative partition within which the application resides.
     */
    partition?: pulumi.Input<string>;
    /**
     * Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
     */
    strictUpdates?: pulumi.Input<string>;
    tables?: pulumi.Input<pulumi.Input<inputs.sys.IAppTable>[]>;
    /**
     * The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
     */
    template?: pulumi.Input<string>;
    /**
     * Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
     */
    templateModified?: pulumi.Input<string>;
    /**
     * Indicates any missing prerequisites associated with the template that defines this application.
     */
    templatePrerequisiteErrors?: pulumi.Input<string>;
    /**
     * The name of the traffic group that the application service is assigned to.
     */
    trafficGroup?: pulumi.Input<string>;
    variables?: pulumi.Input<pulumi.Input<inputs.sys.IAppVariable>[]>;
}
