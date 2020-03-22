// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package f5bigip

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `.As3` provides details about bigip as3 resource
//
// This resource is helpful to configure as3 declarative JSON on BIG-IP.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_as3.html.markdown.
type As3 struct {
	pulumi.CustomResourceState

	// Path/Filename of Declarative AS3 JSON which can be a template/json file used with builtin ```templatefile``` function (or) ```file``` function
	As3Json pulumi.StringOutput `pulumi:"as3Json"`
	// Name of Tenant
	TenantName pulumi.StringOutput `pulumi:"tenantName"`
}

// NewAs3 registers a new resource with the given unique name, arguments, and options.
func NewAs3(ctx *pulumi.Context,
	name string, args *As3Args, opts ...pulumi.ResourceOption) (*As3, error) {
	if args == nil || args.As3Json == nil {
		return nil, errors.New("missing required argument 'As3Json'")
	}
	if args == nil || args.TenantName == nil {
		return nil, errors.New("missing required argument 'TenantName'")
	}
	if args == nil {
		args = &As3Args{}
	}
	var resource As3
	err := ctx.RegisterResource("f5bigip:index/as3:As3", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAs3 gets an existing As3 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAs3(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *As3State, opts ...pulumi.ResourceOption) (*As3, error) {
	var resource As3
	err := ctx.ReadResource("f5bigip:index/as3:As3", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering As3 resources.
type as3State struct {
	// Path/Filename of Declarative AS3 JSON which can be a template/json file used with builtin ```templatefile``` function (or) ```file``` function
	As3Json *string `pulumi:"as3Json"`
	// Name of Tenant
	TenantName *string `pulumi:"tenantName"`
}

type As3State struct {
	// Path/Filename of Declarative AS3 JSON which can be a template/json file used with builtin ```templatefile``` function (or) ```file``` function
	As3Json pulumi.StringPtrInput
	// Name of Tenant
	TenantName pulumi.StringPtrInput
}

func (As3State) ElementType() reflect.Type {
	return reflect.TypeOf((*as3State)(nil)).Elem()
}

type as3Args struct {
	// Path/Filename of Declarative AS3 JSON which can be a template/json file used with builtin ```templatefile``` function (or) ```file``` function
	As3Json string `pulumi:"as3Json"`
	// Name of Tenant
	TenantName string `pulumi:"tenantName"`
}

// The set of arguments for constructing a As3 resource.
type As3Args struct {
	// Path/Filename of Declarative AS3 JSON which can be a template/json file used with builtin ```templatefile``` function (or) ```file``` function
	As3Json pulumi.StringInput
	// Name of Tenant
	TenantName pulumi.StringInput
}

func (As3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*as3Args)(nil)).Elem()
}

