// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `Do` provides details about bigip do resource
//
// This resource is helpful to configure do declarative JSON on BIG-IP.
type Do struct {
	pulumi.CustomResourceState

	// Name of the of the Declarative DO JSON file
	DoJson pulumi.StringOutput `pulumi:"doJson"`
	// unique identifier for DO resource
	//
	// Deprecated: this attribute is no longer in use
	TenantName pulumi.StringPtrOutput `pulumi:"tenantName"`
	// DO json
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewDo registers a new resource with the given unique name, arguments, and options.
func NewDo(ctx *pulumi.Context,
	name string, args *DoArgs, opts ...pulumi.ResourceOption) (*Do, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DoJson == nil {
		return nil, errors.New("invalid value for required argument 'DoJson'")
	}
	var resource Do
	err := ctx.RegisterResource("f5bigip:index/do:Do", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDo gets an existing Do resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DoState, opts ...pulumi.ResourceOption) (*Do, error) {
	var resource Do
	err := ctx.ReadResource("f5bigip:index/do:Do", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Do resources.
type doState struct {
	// Name of the of the Declarative DO JSON file
	DoJson *string `pulumi:"doJson"`
	// unique identifier for DO resource
	//
	// Deprecated: this attribute is no longer in use
	TenantName *string `pulumi:"tenantName"`
	// DO json
	Timeout *int `pulumi:"timeout"`
}

type DoState struct {
	// Name of the of the Declarative DO JSON file
	DoJson pulumi.StringPtrInput
	// unique identifier for DO resource
	//
	// Deprecated: this attribute is no longer in use
	TenantName pulumi.StringPtrInput
	// DO json
	Timeout pulumi.IntPtrInput
}

func (DoState) ElementType() reflect.Type {
	return reflect.TypeOf((*doState)(nil)).Elem()
}

type doArgs struct {
	// Name of the of the Declarative DO JSON file
	DoJson string `pulumi:"doJson"`
	// unique identifier for DO resource
	//
	// Deprecated: this attribute is no longer in use
	TenantName *string `pulumi:"tenantName"`
	// DO json
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a Do resource.
type DoArgs struct {
	// Name of the of the Declarative DO JSON file
	DoJson pulumi.StringInput
	// unique identifier for DO resource
	//
	// Deprecated: this attribute is no longer in use
	TenantName pulumi.StringPtrInput
	// DO json
	Timeout pulumi.IntPtrInput
}

func (DoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*doArgs)(nil)).Elem()
}

type DoInput interface {
	pulumi.Input

	ToDoOutput() DoOutput
	ToDoOutputWithContext(ctx context.Context) DoOutput
}

func (*Do) ElementType() reflect.Type {
	return reflect.TypeOf((*Do)(nil))
}

func (i *Do) ToDoOutput() DoOutput {
	return i.ToDoOutputWithContext(context.Background())
}

func (i *Do) ToDoOutputWithContext(ctx context.Context) DoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DoOutput)
}

type DoOutput struct {
	*pulumi.OutputState
}

func (DoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Do)(nil))
}

func (o DoOutput) ToDoOutput() DoOutput {
	return o
}

func (o DoOutput) ToDoOutputWithContext(ctx context.Context) DoOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DoOutput{})
}
