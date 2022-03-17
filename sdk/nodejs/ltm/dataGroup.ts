// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * `f5bigip.ltm.DataGroup` Manages internal (in-line) datagroup configuration
 *
 * Resource should be named with their`full path`. The full path is the combination of the `partition + name` of the resource, for example `/Common/my-datagroup`.
 */
export class DataGroup extends pulumi.CustomResource {
    /**
     * Get an existing DataGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataGroupState, opts?: pulumi.CustomResourceOptions): DataGroup {
        return new DataGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'f5bigip:ltm/dataGroup:DataGroup';

    /**
     * Returns true if the given object is an instance of DataGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataGroup.__pulumiType;
    }

    /**
     * Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
     */
    public readonly internal!: pulumi.Output<boolean | undefined>;
    /**
     * , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
     */
    public readonly records!: pulumi.Output<outputs.ltm.DataGroupRecord[] | undefined>;
    /**
     * Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
     * This should be used in conjunction with `internal` attribute set `false`
     */
    public readonly recordsSrc!: pulumi.Output<string | undefined>;
    /**
     * datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a DataGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataGroupArgs | DataGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataGroupState | undefined;
            resourceInputs["internal"] = state ? state.internal : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["records"] = state ? state.records : undefined;
            resourceInputs["recordsSrc"] = state ? state.recordsSrc : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as DataGroupArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["internal"] = args ? args.internal : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["records"] = args ? args.records : undefined;
            resourceInputs["recordsSrc"] = args ? args.recordsSrc : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataGroup resources.
 */
export interface DataGroupState {
    /**
     * Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
     */
    internal?: pulumi.Input<boolean>;
    /**
     * , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
     */
    name?: pulumi.Input<string>;
    /**
     * a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
     */
    records?: pulumi.Input<pulumi.Input<inputs.ltm.DataGroupRecord>[]>;
    /**
     * Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
     * This should be used in conjunction with `internal` attribute set `false`
     */
    recordsSrc?: pulumi.Input<string>;
    /**
     * datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataGroup resource.
 */
export interface DataGroupArgs {
    /**
     * Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
     */
    internal?: pulumi.Input<boolean>;
    /**
     * , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
     */
    name: pulumi.Input<string>;
    /**
     * a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
     */
    records?: pulumi.Input<pulumi.Input<inputs.ltm.DataGroupRecord>[]>;
    /**
     * Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
     * This should be used in conjunction with `internal` attribute set `false`
     */
    recordsSrc?: pulumi.Input<string>;
    /**
     * datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
     */
    type: pulumi.Input<string>;
}
