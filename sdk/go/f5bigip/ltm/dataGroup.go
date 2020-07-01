// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
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
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &DataGroupArgs{}
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
