// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ltm.DataGroup` Manages internal (in-line) datagroup configuration
//
// Resource should be named with their`full path`. The full path is the combination of the `partition + name` of the resource, for example `/Common/my-datagroup`.
type DataGroup struct {
	pulumi.CustomResourceState

	// Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
	Internal pulumi.BoolPtrOutput `pulumi:"internal"`
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name pulumi.StringOutput `pulumi:"name"`
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records DataGroupRecordArrayOutput `pulumi:"records"`
	// Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
	// This should be used in conjunction with `internal` attribute set `false`
	RecordsSrc pulumi.StringPtrOutput `pulumi:"recordsSrc"`
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataGroup registers a new resource with the given unique name, arguments, and options.
func NewDataGroup(ctx *pulumi.Context,
	name string, args *DataGroupArgs, opts ...pulumi.ResourceOption) (*DataGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource DataGroup
	err := ctx.RegisterResource("f5bigip:ltm/dataGroup:DataGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataGroup gets an existing DataGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataGroupState, opts ...pulumi.ResourceOption) (*DataGroup, error) {
	var resource DataGroup
	err := ctx.ReadResource("f5bigip:ltm/dataGroup:DataGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataGroup resources.
type dataGroupState struct {
	// Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
	Internal *bool `pulumi:"internal"`
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name *string `pulumi:"name"`
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records []DataGroupRecord `pulumi:"records"`
	// Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
	// This should be used in conjunction with `internal` attribute set `false`
	RecordsSrc *string `pulumi:"recordsSrc"`
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type *string `pulumi:"type"`
}

type DataGroupState struct {
	// Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
	Internal pulumi.BoolPtrInput
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name pulumi.StringPtrInput
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records DataGroupRecordArrayInput
	// Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
	// This should be used in conjunction with `internal` attribute set `false`
	RecordsSrc pulumi.StringPtrInput
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type pulumi.StringPtrInput
}

func (DataGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataGroupState)(nil)).Elem()
}

type dataGroupArgs struct {
	// Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
	Internal *bool `pulumi:"internal"`
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name string `pulumi:"name"`
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records []DataGroupRecord `pulumi:"records"`
	// Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
	// This should be used in conjunction with `internal` attribute set `false`
	RecordsSrc *string `pulumi:"recordsSrc"`
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DataGroup resource.
type DataGroupArgs struct {
	// Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
	Internal pulumi.BoolPtrInput
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name pulumi.StringInput
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records DataGroupRecordArrayInput
	// Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
	// This should be used in conjunction with `internal` attribute set `false`
	RecordsSrc pulumi.StringPtrInput
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type pulumi.StringInput
}

func (DataGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataGroupArgs)(nil)).Elem()
}

type DataGroupInput interface {
	pulumi.Input

	ToDataGroupOutput() DataGroupOutput
	ToDataGroupOutputWithContext(ctx context.Context) DataGroupOutput
}

func (*DataGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DataGroup)(nil)).Elem()
}

func (i *DataGroup) ToDataGroupOutput() DataGroupOutput {
	return i.ToDataGroupOutputWithContext(context.Background())
}

func (i *DataGroup) ToDataGroupOutputWithContext(ctx context.Context) DataGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataGroupOutput)
}

// DataGroupArrayInput is an input type that accepts DataGroupArray and DataGroupArrayOutput values.
// You can construct a concrete instance of `DataGroupArrayInput` via:
//
//          DataGroupArray{ DataGroupArgs{...} }
type DataGroupArrayInput interface {
	pulumi.Input

	ToDataGroupArrayOutput() DataGroupArrayOutput
	ToDataGroupArrayOutputWithContext(context.Context) DataGroupArrayOutput
}

type DataGroupArray []DataGroupInput

func (DataGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataGroup)(nil)).Elem()
}

func (i DataGroupArray) ToDataGroupArrayOutput() DataGroupArrayOutput {
	return i.ToDataGroupArrayOutputWithContext(context.Background())
}

func (i DataGroupArray) ToDataGroupArrayOutputWithContext(ctx context.Context) DataGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataGroupArrayOutput)
}

// DataGroupMapInput is an input type that accepts DataGroupMap and DataGroupMapOutput values.
// You can construct a concrete instance of `DataGroupMapInput` via:
//
//          DataGroupMap{ "key": DataGroupArgs{...} }
type DataGroupMapInput interface {
	pulumi.Input

	ToDataGroupMapOutput() DataGroupMapOutput
	ToDataGroupMapOutputWithContext(context.Context) DataGroupMapOutput
}

type DataGroupMap map[string]DataGroupInput

func (DataGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataGroup)(nil)).Elem()
}

func (i DataGroupMap) ToDataGroupMapOutput() DataGroupMapOutput {
	return i.ToDataGroupMapOutputWithContext(context.Background())
}

func (i DataGroupMap) ToDataGroupMapOutputWithContext(ctx context.Context) DataGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataGroupMapOutput)
}

type DataGroupOutput struct{ *pulumi.OutputState }

func (DataGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataGroup)(nil)).Elem()
}

func (o DataGroupOutput) ToDataGroupOutput() DataGroupOutput {
	return o
}

func (o DataGroupOutput) ToDataGroupOutputWithContext(ctx context.Context) DataGroupOutput {
	return o
}

// Set `false` if you want to Create External Datagroups. default is `true`,means creates internal datagroup.
func (o DataGroupOutput) Internal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataGroup) pulumi.BoolPtrOutput { return v.Internal }).(pulumi.BoolPtrOutput)
}

// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
func (o DataGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
func (o DataGroupOutput) Records() DataGroupRecordArrayOutput {
	return o.ApplyT(func(v *DataGroup) DataGroupRecordArrayOutput { return v.Records }).(DataGroupRecordArrayOutput)
}

// Path to a file with records in it,The file should be well-formed,it includes records, one per line,that resemble the following format "key separator value". For example, `foo := bar`.
// This should be used in conjunction with `internal` attribute set `false`
func (o DataGroupOutput) RecordsSrc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataGroup) pulumi.StringPtrOutput { return v.RecordsSrc }).(pulumi.StringPtrOutput)
}

// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
func (o DataGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DataGroupArrayOutput struct{ *pulumi.OutputState }

func (DataGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataGroup)(nil)).Elem()
}

func (o DataGroupArrayOutput) ToDataGroupArrayOutput() DataGroupArrayOutput {
	return o
}

func (o DataGroupArrayOutput) ToDataGroupArrayOutputWithContext(ctx context.Context) DataGroupArrayOutput {
	return o
}

func (o DataGroupArrayOutput) Index(i pulumi.IntInput) DataGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataGroup {
		return vs[0].([]*DataGroup)[vs[1].(int)]
	}).(DataGroupOutput)
}

type DataGroupMapOutput struct{ *pulumi.OutputState }

func (DataGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataGroup)(nil)).Elem()
}

func (o DataGroupMapOutput) ToDataGroupMapOutput() DataGroupMapOutput {
	return o
}

func (o DataGroupMapOutput) ToDataGroupMapOutputWithContext(ctx context.Context) DataGroupMapOutput {
	return o
}

func (o DataGroupMapOutput) MapIndex(k pulumi.StringInput) DataGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataGroup {
		return vs[0].(map[string]*DataGroup)[vs[1].(string)]
	}).(DataGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataGroupInput)(nil)).Elem(), &DataGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataGroupArrayInput)(nil)).Elem(), DataGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataGroupMapInput)(nil)).Elem(), DataGroupMap{})
	pulumi.RegisterOutputType(DataGroupOutput{})
	pulumi.RegisterOutputType(DataGroupArrayOutput{})
	pulumi.RegisterOutputType(DataGroupMapOutput{})
}
