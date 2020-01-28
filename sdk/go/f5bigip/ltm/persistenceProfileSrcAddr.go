// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures a source address persistence profile
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
// `hashAlgorithm` (Optional) Specify the hash algorithm
// 
// `mask` (Optional) Identify a range of source IP addresses to manage together as a single source address affinity persistent connection when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
// 
// `mapProxies` (Optional) (enabled or disabled) Directs all to the same single pool member
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/ltm_persistence_profile_srcaddr.html.markdown.
type PersistenceProfileSrcAddr struct {
	pulumi.CustomResourceState

	AppService pulumi.StringPtrOutput `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringPtrOutput `pulumi:"hashAlgorithm"`
	// To enable _ disable directs all to the same single pool member
	MapProxies pulumi.StringPtrOutput `pulumi:"mapProxies"`
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask pulumi.StringPtrOutput `pulumi:"mask"`
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

// NewPersistenceProfileSrcAddr registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileSrcAddr(ctx *pulumi.Context,
	name string, args *PersistenceProfileSrcAddrArgs, opts ...pulumi.ResourceOption) (*PersistenceProfileSrcAddr, error) {
	if args == nil || args.DefaultsFrom == nil {
		return nil, errors.New("missing required argument 'DefaultsFrom'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &PersistenceProfileSrcAddrArgs{}
	}
	var resource PersistenceProfileSrcAddr
	err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistenceProfileSrcAddr gets an existing PersistenceProfileSrcAddr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileSrcAddr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistenceProfileSrcAddrState, opts ...pulumi.ResourceOption) (*PersistenceProfileSrcAddr, error) {
	var resource PersistenceProfileSrcAddr
	err := ctx.ReadResource("f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistenceProfileSrcAddr resources.
type persistenceProfileSrcAddrState struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm *string `pulumi:"hashAlgorithm"`
	// To enable _ disable directs all to the same single pool member
	MapProxies *string `pulumi:"mapProxies"`
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask *string `pulumi:"mask"`
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

type PersistenceProfileSrcAddrState struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringPtrInput
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringPtrInput
	// To enable _ disable directs all to the same single pool member
	MapProxies pulumi.StringPtrInput
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask pulumi.StringPtrInput
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

func (PersistenceProfileSrcAddrState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSrcAddrState)(nil)).Elem()
}

type persistenceProfileSrcAddrArgs struct {
	AppService *string `pulumi:"appService"`
	// Inherit defaults from parent profile
	DefaultsFrom string `pulumi:"defaultsFrom"`
	// Specify the hash algorithm
	HashAlgorithm *string `pulumi:"hashAlgorithm"`
	// To enable _ disable directs all to the same single pool member
	MapProxies *string `pulumi:"mapProxies"`
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask *string `pulumi:"mask"`
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

// The set of arguments for constructing a PersistenceProfileSrcAddr resource.
type PersistenceProfileSrcAddrArgs struct {
	AppService pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringInput
	// Specify the hash algorithm
	HashAlgorithm pulumi.StringPtrInput
	// To enable _ disable directs all to the same single pool member
	MapProxies pulumi.StringPtrInput
	// Identify a range of source IP addresses to manage together as a single source address affinity persistent connection
	// when connecting to the pool. Must be a valid IPv4 or IPv6 mask.
	Mask pulumi.StringPtrInput
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

func (PersistenceProfileSrcAddrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileSrcAddrArgs)(nil)).Elem()
}

