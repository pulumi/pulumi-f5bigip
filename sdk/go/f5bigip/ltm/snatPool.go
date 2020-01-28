// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `ltm.SnatPool` Collections of SNAT translation addresses
// 
// Resource should be named with their "full path". The full path is the combination of the partition + name of the resource, for example /Common/my-snatpool. 
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_snatpool.html.markdown.
type SnatPool struct {
	pulumi.CustomResourceState

	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Name of the snatpool
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSnatPool registers a new resource with the given unique name, arguments, and options.
func NewSnatPool(ctx *pulumi.Context,
	name string, args *SnatPoolArgs, opts ...pulumi.ResourceOption) (*SnatPool, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &SnatPoolArgs{}
	}
	var resource SnatPool
	err := ctx.RegisterResource("f5bigip:ltm/snatPool:SnatPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnatPool gets an existing SnatPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnatPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnatPoolState, opts ...pulumi.ResourceOption) (*SnatPool, error) {
	var resource SnatPool
	err := ctx.ReadResource("f5bigip:ltm/snatPool:SnatPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnatPool resources.
type snatPoolState struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members []string `pulumi:"members"`
	// Name of the snatpool
	Name *string `pulumi:"name"`
}

type SnatPoolState struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members pulumi.StringArrayInput
	// Name of the snatpool
	Name pulumi.StringPtrInput
}

func (SnatPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*snatPoolState)(nil)).Elem()
}

type snatPoolArgs struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members []string `pulumi:"members"`
	// Name of the snatpool
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a SnatPool resource.
type SnatPoolArgs struct {
	// Specifies a translation address to add to or delete from a SNAT pool (at least one address is required)
	Members pulumi.StringArrayInput
	// Name of the snatpool
	Name pulumi.StringInput
}

func (SnatPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snatPoolArgs)(nil)).Elem()
}

