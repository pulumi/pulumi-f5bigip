// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `f5bigip.sys.IApp` resource helps you to deploy Application Services template that can be used to automate and orchestrate Layer 4-7 applications service deployments using F5 Network.  
 * 
 * ## Example Usage of Json file
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 * 
 *  * `description` - User defined description.
 *  * `deviceGroup` - The name of the device group that the application service is assigned to.
 *  * `executeAction` - Run the specified template action associated with the application.
 *  * `inheritedDevicegroup`- Read-only. Shows whether the application folder will automatically remain with the same device-group as its parent folder. Use 'device-group default' or 'device-group non-default' to set this.
 *  * `inheritedTrafficGroup` - Read-only. Shows whether the application folder will automatically remain with the same traffic-group as its parent folder. Use 'traffic-group default' or 'traffic-group non-default' to set this.
 *  * `partition` - Displays the administrative partition within which the application resides.
 *  * `strictUpdates` - Specifies whether configuration objects contained in the application may be directly modified, outside the context of the system's application management interfaces.
 *  * `template` - The template defines the configuration for the application. This may be changed after the application has been created to move the application to a new template.
 *  * `templateModified` - Indicates that the application template used to deploy the application has been modified. The application should be updated to make use of the latest changes.
 *  * `templatePrerequisiteErrors` - Indicates any missing prerequisites associated with the template that defines this application.
 *  * `trafficGroup` - The name of the traffic group that the application service is assigned to.
 *  * `lists` - string values
 *  * `metadata` - User defined generic data for the application service. It is a name and value pair.
 *  * `tables` - Values provided like pool name, nodes etc.
 *  * `variables` - Name, values, encrypted or not
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_sys_iapp.html.markdown.
 */
export class IApp extends pulumi.CustomResource {
    /**
     * Get an existing IApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * Address of the Iapp which needs to be Iappensed
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly devicegroup!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly executeAction!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly inheritedDevicegroup!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly inheritedTrafficGroup!: pulumi.Output<string | undefined>;
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     */
    public readonly jsonfile!: pulumi.Output<string | undefined>;
    public readonly lists!: pulumi.Output<outputs.sys.IAppList[] | undefined>;
    public readonly metadatas!: pulumi.Output<outputs.sys.IAppMetadata[] | undefined>;
    /**
     * Name of the iApp.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Address of the Iapp which needs to be Iappensed
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly strictUpdates!: pulumi.Output<string | undefined>;
    public readonly tables!: pulumi.Output<outputs.sys.IAppTable[] | undefined>;
    /**
     * BIG-IP password
     */
    public readonly template!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly templateModified!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
     */
    public readonly templatePrerequisiteErrors!: pulumi.Output<string | undefined>;
    /**
     * BIG-IP password
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
    constructor(name: string, args?: IAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IAppArgs | IAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IAppState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["devicegroup"] = state ? state.devicegroup : undefined;
            inputs["executeAction"] = state ? state.executeAction : undefined;
            inputs["inheritedDevicegroup"] = state ? state.inheritedDevicegroup : undefined;
            inputs["inheritedTrafficGroup"] = state ? state.inheritedTrafficGroup : undefined;
            inputs["jsonfile"] = state ? state.jsonfile : undefined;
            inputs["lists"] = state ? state.lists : undefined;
            inputs["metadatas"] = state ? state.metadatas : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["partition"] = state ? state.partition : undefined;
            inputs["strictUpdates"] = state ? state.strictUpdates : undefined;
            inputs["tables"] = state ? state.tables : undefined;
            inputs["template"] = state ? state.template : undefined;
            inputs["templateModified"] = state ? state.templateModified : undefined;
            inputs["templatePrerequisiteErrors"] = state ? state.templatePrerequisiteErrors : undefined;
            inputs["trafficGroup"] = state ? state.trafficGroup : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as IAppArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["devicegroup"] = args ? args.devicegroup : undefined;
            inputs["executeAction"] = args ? args.executeAction : undefined;
            inputs["inheritedDevicegroup"] = args ? args.inheritedDevicegroup : undefined;
            inputs["inheritedTrafficGroup"] = args ? args.inheritedTrafficGroup : undefined;
            inputs["jsonfile"] = args ? args.jsonfile : undefined;
            inputs["lists"] = args ? args.lists : undefined;
            inputs["metadatas"] = args ? args.metadatas : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["partition"] = args ? args.partition : undefined;
            inputs["strictUpdates"] = args ? args.strictUpdates : undefined;
            inputs["tables"] = args ? args.tables : undefined;
            inputs["template"] = args ? args.template : undefined;
            inputs["templateModified"] = args ? args.templateModified : undefined;
            inputs["templatePrerequisiteErrors"] = args ? args.templatePrerequisiteErrors : undefined;
            inputs["trafficGroup"] = args ? args.trafficGroup : undefined;
            inputs["variables"] = args ? args.variables : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IApp resources.
 */
export interface IAppState {
    /**
     * Address of the Iapp which needs to be Iappensed
     */
    readonly description?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly devicegroup?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly executeAction?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly inheritedDevicegroup?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly inheritedTrafficGroup?: pulumi.Input<string>;
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     */
    readonly jsonfile?: pulumi.Input<string>;
    readonly lists?: pulumi.Input<pulumi.Input<inputs.sys.IAppList>[]>;
    readonly metadatas?: pulumi.Input<pulumi.Input<inputs.sys.IAppMetadata>[]>;
    /**
     * Name of the iApp.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Address of the Iapp which needs to be Iappensed
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly strictUpdates?: pulumi.Input<string>;
    readonly tables?: pulumi.Input<pulumi.Input<inputs.sys.IAppTable>[]>;
    /**
     * BIG-IP password
     */
    readonly template?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly templateModified?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly templatePrerequisiteErrors?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly trafficGroup?: pulumi.Input<string>;
    readonly variables?: pulumi.Input<pulumi.Input<inputs.sys.IAppVariable>[]>;
}

/**
 * The set of arguments for constructing a IApp resource.
 */
export interface IAppArgs {
    /**
     * Address of the Iapp which needs to be Iappensed
     */
    readonly description?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly devicegroup?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly executeAction?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly inheritedDevicegroup?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly inheritedTrafficGroup?: pulumi.Input<string>;
    /**
     * Refer to the Json file which will be deployed on F5 BIG-IP.
     */
    readonly jsonfile?: pulumi.Input<string>;
    readonly lists?: pulumi.Input<pulumi.Input<inputs.sys.IAppList>[]>;
    readonly metadatas?: pulumi.Input<pulumi.Input<inputs.sys.IAppMetadata>[]>;
    /**
     * Name of the iApp.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Address of the Iapp which needs to be Iappensed
     */
    readonly partition?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly strictUpdates?: pulumi.Input<string>;
    readonly tables?: pulumi.Input<pulumi.Input<inputs.sys.IAppTable>[]>;
    /**
     * BIG-IP password
     */
    readonly template?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly templateModified?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly templatePrerequisiteErrors?: pulumi.Input<string>;
    /**
     * BIG-IP password
     */
    readonly trafficGroup?: pulumi.Input<string>;
    readonly variables?: pulumi.Input<pulumi.Input<inputs.sys.IAppVariable>[]>;
}
