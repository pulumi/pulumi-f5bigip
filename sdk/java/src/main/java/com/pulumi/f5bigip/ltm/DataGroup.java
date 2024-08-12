// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.f5bigip.ltm;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.f5bigip.Utilities;
import com.pulumi.f5bigip.ltm.DataGroupArgs;
import com.pulumi.f5bigip.ltm.inputs.DataGroupState;
import com.pulumi.f5bigip.ltm.outputs.DataGroupRecord;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * `f5bigip.ltm.DataGroup` Manages internal (in-line) datagroup configuration
 * 
 * Resource should be named with their`full path`. The full path is the combination of the `partition + name` of the resource, for example `/Common/my-datagroup`.
 * 
 */
@ResourceType(type="f5bigip:ltm/dataGroup:DataGroup")
public class DataGroup extends com.pulumi.resources.CustomResource {
    /**
     * Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
     * 
     */
    @Export(name="internal", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> internal;

    /**
     * @return Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
     * 
     */
    public Output<Optional<Boolean>> internal() {
        return Codegen.optional(this.internal);
    }
    /**
     * Name of the datagroup
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the datagroup
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
     * 
     */
    @Export(name="records", refs={List.class,DataGroupRecord.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataGroupRecord>> records;

    /**
     * @return a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
     * 
     */
    public Output<Optional<List<DataGroupRecord>>> records() {
        return Codegen.optional(this.records);
    }
    /**
     * Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format &#34;key separator value&#34;. For example, `foo := bar`.
     * This should be used in conjunction with `internal` attribute set `false`
     * 
     */
    @Export(name="recordsSrc", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> recordsSrc;

    /**
     * @return Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format &#34;key separator value&#34;. For example, `foo := bar`.
     * This should be used in conjunction with `internal` attribute set `false`
     * 
     */
    public Output<Optional<String>> recordsSrc() {
        return Codegen.optional(this.recordsSrc);
    }
    /**
     * datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataGroup(java.lang.String name) {
        this(name, DataGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataGroup(java.lang.String name, DataGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataGroup(java.lang.String name, DataGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/dataGroup:DataGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DataGroup(java.lang.String name, Output<java.lang.String> id, @Nullable DataGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("f5bigip:ltm/dataGroup:DataGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static DataGroupArgs makeArgs(DataGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DataGroupArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DataGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable DataGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataGroup(name, id, state, options);
    }
}
