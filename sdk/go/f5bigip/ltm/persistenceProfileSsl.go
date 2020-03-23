// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures an SSL persistence profile
//
// ## Reference
//
// `name` - (Required) Name of the virtual address
//
// `defaultsFrom` - (Required) Parent cookie persistence profile
//
// `matchAcrossPools` (Optional) (enabled or disabled) match across pools with given persistence record
//
// `matchAcrossServices` (Optional) (enabled or disabled) match across services with given persistence record
//
// `matchAcrossVirtuals` (Optional) (enabled or disabled) match across virtual servers with given persistence record
//
// `mirror` (Optional) (enabled or disabled) mirror persistence record
//
// `timeout` (Optional) (enabled or disabled) Timeout for persistence of the session in seconds
//
// `overrideConnLimit` (Optional) (enabled or disabled) Enable or dissable pool member connection limits are overridden for persisted clients. Per-virtual connection limits remain hard limits and are not overridden.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_persistence_profile_ssl.html.markdown.
type PersistenceProfileSsl struct {
	pulumi.CustomResourceState

	AppService pulumi.StringPtrOutput `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrOutput `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrOutput `pulumi:"matchAcrossServices"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals pulumi.StringPtrOutput `pulumi:"matchAcrossVirtuals"`
	// To enable _ disable
	Mirror pulumi.StringPtrOutput `pulumi:"mirror"`
	// Name of the persistence profile
	Name pulumi.StringOutput `pulumi:"name"`
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit pulumi.StringPtrOutput `pulumi:"overrideConnLimit"`
	// Timeout for persistence of the session
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewPersistenceProfileSsl registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileSsl(ctx *pulumi.Context,
	name string, args *PersistenceProfileSslArgs, opts ...pulumi.ResourceOption) (*PersistenceProfileSsl, error) {
	if args == nil || args.DefaultsFrom == nil {
		return nil, errors.New("missing required argument 'DefaultsFrom'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &PersistenceProfileSslArgs{}
	}
	var resource PersistenceProfileSsl
	err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistenceProfileSsl gets an existing PersistenceProfileSsl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileSsl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistenceProfileSslState, opts ...pulumi.ResourceOption) (*PersistenceProfileSsl, error) {
	var resource PersistenceProfileSsl
	err := ctx.ReadResource("f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistenceProfileSsl resources.
type persistenceProfileSslState struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools *string `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices *string `pulumi:"matchAcrossServices"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals *string `pulumi:"matchAcrossVirtuals"`
	// To enable _ disable
	Mirror *string `pulumi:"mirror"`
	// Name of the persistence profile
	Name *string `pulumi:"name"`
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit *string `pulumi:"overrideConnLimit"`
	// Timeout for persistence of the session
	Timeout *int `pulumi:"timeout"`
}

type PersistenceProfileSslState struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringPtrInput
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals pulumi.StringPtrInput
	// To enable _ disable
	Mirror pulumi.StringPtrInput
	// Name of the persistence profile
	Name pulumi.StringPtrInput
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit pulumi.StringPtrInput
	// Timeout for persistence of the session
	Timeout pulumi.IntPtrInput
}

func (PersistenceProfileSslState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSslState)(nil)).Elem()
}

type persistenceProfileSslArgs struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom string `pulumi:"defaultsFrom"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools *string `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices *string `pulumi:"matchAcrossServices"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals *string `pulumi:"matchAcrossVirtuals"`
	// To enable _ disable
	Mirror *string `pulumi:"mirror"`
	// Name of the persistence profile
	Name string `pulumi:"name"`
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit *string `pulumi:"overrideConnLimit"`
	// Timeout for persistence of the session
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a PersistenceProfileSsl resource.
type PersistenceProfileSslArgs struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringInput
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossVirtuals pulumi.StringPtrInput
	// To enable _ disable
	Mirror pulumi.StringPtrInput
	// Name of the persistence profile
	Name pulumi.StringInput
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit pulumi.StringPtrInput
	// Timeout for persistence of the session
	Timeout pulumi.IntPtrInput
}

func (PersistenceProfileSslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSslArgs)(nil)).Elem()
}

