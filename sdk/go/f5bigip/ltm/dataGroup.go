// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ltm.DataGroup` Manages internal (in-line) datagroup configuration
//
// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-datagroup.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-f5bigip/sdk/v2/go/f5bigip/ltm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ltm.NewDataGroup(ctx, "datagroup", &ltm.DataGroupArgs{
// 			Name: pulumi.String("/Common/dgx2"),
// 			Records: ltm.DataGroupRecordArray{
// 				&ltm.DataGroupRecordArgs{
// 					Data: pulumi.String("pool1"),
// 					Name: pulumi.String("abc.com"),
// 				},
// 				&ltm.DataGroupRecordArgs{
// 					Data: pulumi.String("123"),
// 					Name: pulumi.String("test"),
// 				},
// 			},
// 			Type: pulumi.String("string"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DataGroup struct {
	pulumi.CustomResourceState

	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name pulumi.StringOutput `pulumi:"name"`
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records DataGroupRecordArrayOutput `pulumi:"records"`
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
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name *string `pulumi:"name"`
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records []DataGroupRecord `pulumi:"records"`
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type *string `pulumi:"type"`
}

type DataGroupState struct {
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name pulumi.StringPtrInput
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records DataGroupRecordArrayInput
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type pulumi.StringPtrInput
}

func (DataGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataGroupState)(nil)).Elem()
}

type dataGroupArgs struct {
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name string `pulumi:"name"`
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records []DataGroupRecord `pulumi:"records"`
	// datagroup type (applies to the `name` field of the record), supports: `string`, `ip` or `integer`
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DataGroup resource.
type DataGroupArgs struct {
	// , sets the value of the record's `name` attribute, must be of type defined in `type` attribute
	Name pulumi.StringInput
	// a set of `name` and `data` attributes, name must be of type specified by the `type` attributed (`string`, `ip` and `integer`), data is optional and can take any value, multiple `record` sets can be specified as needed.
	Records DataGroupRecordArrayInput
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
	return reflect.TypeOf((*DataGroup)(nil))
}

func (i *DataGroup) ToDataGroupOutput() DataGroupOutput {
	return i.ToDataGroupOutputWithContext(context.Background())
}

func (i *DataGroup) ToDataGroupOutputWithContext(ctx context.Context) DataGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataGroupOutput)
}

func (i *DataGroup) ToDataGroupPtrOutput() DataGroupPtrOutput {
	return i.ToDataGroupPtrOutputWithContext(context.Background())
}

func (i *DataGroup) ToDataGroupPtrOutputWithContext(ctx context.Context) DataGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataGroupPtrOutput)
}

type DataGroupPtrInput interface {
	pulumi.Input

	ToDataGroupPtrOutput() DataGroupPtrOutput
	ToDataGroupPtrOutputWithContext(ctx context.Context) DataGroupPtrOutput
}

type dataGroupPtrType DataGroupArgs

func (*dataGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataGroup)(nil))
}

func (i *dataGroupPtrType) ToDataGroupPtrOutput() DataGroupPtrOutput {
	return i.ToDataGroupPtrOutputWithContext(context.Background())
}

func (i *dataGroupPtrType) ToDataGroupPtrOutputWithContext(ctx context.Context) DataGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataGroupPtrOutput)
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
	return reflect.TypeOf(([]*DataGroup)(nil))
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
	return reflect.TypeOf((map[string]*DataGroup)(nil))
}

func (i DataGroupMap) ToDataGroupMapOutput() DataGroupMapOutput {
	return i.ToDataGroupMapOutputWithContext(context.Background())
}

func (i DataGroupMap) ToDataGroupMapOutputWithContext(ctx context.Context) DataGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataGroupMapOutput)
}

type DataGroupOutput struct {
	*pulumi.OutputState
}

func (DataGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataGroup)(nil))
}

func (o DataGroupOutput) ToDataGroupOutput() DataGroupOutput {
	return o
}

func (o DataGroupOutput) ToDataGroupOutputWithContext(ctx context.Context) DataGroupOutput {
	return o
}

func (o DataGroupOutput) ToDataGroupPtrOutput() DataGroupPtrOutput {
	return o.ToDataGroupPtrOutputWithContext(context.Background())
}

func (o DataGroupOutput) ToDataGroupPtrOutputWithContext(ctx context.Context) DataGroupPtrOutput {
	return o.ApplyT(func(v DataGroup) *DataGroup {
		return &v
	}).(DataGroupPtrOutput)
}

type DataGroupPtrOutput struct {
	*pulumi.OutputState
}

func (DataGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataGroup)(nil))
}

func (o DataGroupPtrOutput) ToDataGroupPtrOutput() DataGroupPtrOutput {
	return o
}

func (o DataGroupPtrOutput) ToDataGroupPtrOutputWithContext(ctx context.Context) DataGroupPtrOutput {
	return o
}

type DataGroupArrayOutput struct{ *pulumi.OutputState }

func (DataGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataGroup)(nil))
}

func (o DataGroupArrayOutput) ToDataGroupArrayOutput() DataGroupArrayOutput {
	return o
}

func (o DataGroupArrayOutput) ToDataGroupArrayOutputWithContext(ctx context.Context) DataGroupArrayOutput {
	return o
}

func (o DataGroupArrayOutput) Index(i pulumi.IntInput) DataGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataGroup {
		return vs[0].([]DataGroup)[vs[1].(int)]
	}).(DataGroupOutput)
}

type DataGroupMapOutput struct{ *pulumi.OutputState }

func (DataGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DataGroup)(nil))
}

func (o DataGroupMapOutput) ToDataGroupMapOutput() DataGroupMapOutput {
	return o
}

func (o DataGroupMapOutput) ToDataGroupMapOutputWithContext(ctx context.Context) DataGroupMapOutput {
	return o
}

func (o DataGroupMapOutput) MapIndex(k pulumi.StringInput) DataGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DataGroup {
		return vs[0].(map[string]DataGroup)[vs[1].(string)]
	}).(DataGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(DataGroupOutput{})
	pulumi.RegisterOutputType(DataGroupPtrOutput{})
	pulumi.RegisterOutputType(DataGroupArrayOutput{})
	pulumi.RegisterOutputType(DataGroupMapOutput{})
}
