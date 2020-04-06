// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ltm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures a cookie persistence profile
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
// `alwaysSend` (Optional) (enabled or disabled) always send cookies
//
// `cookieEncryption` (Optional) (required, preferred, or disabled) To required, preferred, or disabled policy for cookie encryption
//
// `cookieEncryptionPassphrase` (Optional) (required, preferred, or disabled) Passphrase for encrypted cookies. The field is encrypted on the server and will always return differently then set.
// If this is configured specify `ignoreChanges` under the `lifecycle` block to ignore returned encrypted value.
//
// `cookieName` (Optional) Name of the cookie to track persistence
//
// `expiration` (Optional) Expiration TTL for cookie specified in DAY:HOUR:MIN:SECONDS (Examples: 1:0:0:0 one day, 1:0:0 one hour, 30:0 thirty minutes)
//
// `hashLength` (Optional) (Integer) Length of hash to apply to cookie
//
// `hashOffset` (Optional) (Integer) Number of characters to skip in the cookie for the hash
//
// `httponly` (Optional) (enabled or disabled) Sending only over http
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/r/bigip_ltm_persistence_profile_cookie.html.markdown.
type PersistenceProfileCookie struct {
	pulumi.CustomResourceState

	// To enable _ disable always sending cookies
	AlwaysSend pulumi.StringPtrOutput `pulumi:"alwaysSend"`
	AppService pulumi.StringPtrOutput `pulumi:"appService"`
	// To required, preferred, or disabled policy for cookie encryption
	CookieEncryption pulumi.StringPtrOutput `pulumi:"cookieEncryption"`
	// Passphrase for encrypted cookies
	CookieEncryptionPassphrase pulumi.StringPtrOutput `pulumi:"cookieEncryptionPassphrase"`
	// Name of the cookie to track persistence
	CookieName pulumi.StringPtrOutput `pulumi:"cookieName"`
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Expiration TTL for cookie specified in D:H:M:S or in seconds
	Expiration pulumi.StringPtrOutput `pulumi:"expiration"`
	// Length of hash to apply to cookie
	HashLength pulumi.IntPtrOutput `pulumi:"hashLength"`
	// Number of characters to skip in the cookie for the hash
	HashOffset pulumi.IntPtrOutput `pulumi:"hashOffset"`
	// To enable _ disable sending only over http
	Httponly pulumi.StringPtrOutput `pulumi:"httponly"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrOutput `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrOutput `pulumi:"matchAcrossServices"`
	// To enable _ disable match across virtual servers with given persistence record
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

// NewPersistenceProfileCookie registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileCookie(ctx *pulumi.Context,
	name string, args *PersistenceProfileCookieArgs, opts ...pulumi.ResourceOption) (*PersistenceProfileCookie, error) {
	if args == nil || args.DefaultsFrom == nil {
		return nil, errors.New("missing required argument 'DefaultsFrom'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &PersistenceProfileCookieArgs{}
	}
	var resource PersistenceProfileCookie
	err := ctx.RegisterResource("f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPersistenceProfileCookie gets an existing PersistenceProfileCookie resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPersistenceProfileCookie(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PersistenceProfileCookieState, opts ...pulumi.ResourceOption) (*PersistenceProfileCookie, error) {
	var resource PersistenceProfileCookie
	err := ctx.ReadResource("f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PersistenceProfileCookie resources.
type persistenceProfileCookieState struct {
	// To enable _ disable always sending cookies
	AlwaysSend *string `pulumi:"alwaysSend"`
	AppService *string `pulumi:"appService"`
	// To required, preferred, or disabled policy for cookie encryption
	CookieEncryption *string `pulumi:"cookieEncryption"`
	// Passphrase for encrypted cookies
	CookieEncryptionPassphrase *string `pulumi:"cookieEncryptionPassphrase"`
	// Name of the cookie to track persistence
	CookieName *string `pulumi:"cookieName"`
	// Inherit defaults from parent profile
	DefaultsFrom *string `pulumi:"defaultsFrom"`
	// Expiration TTL for cookie specified in D:H:M:S or in seconds
	Expiration *string `pulumi:"expiration"`
	// Length of hash to apply to cookie
	HashLength *int `pulumi:"hashLength"`
	// Number of characters to skip in the cookie for the hash
	HashOffset *int `pulumi:"hashOffset"`
	// To enable _ disable sending only over http
	Httponly *string `pulumi:"httponly"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools *string `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices *string `pulumi:"matchAcrossServices"`
	// To enable _ disable match across virtual servers with given persistence record
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

type PersistenceProfileCookieState struct {
	// To enable _ disable always sending cookies
	AlwaysSend pulumi.StringPtrInput
	AppService pulumi.StringPtrInput
	// To required, preferred, or disabled policy for cookie encryption
	CookieEncryption pulumi.StringPtrInput
	// Passphrase for encrypted cookies
	CookieEncryptionPassphrase pulumi.StringPtrInput
	// Name of the cookie to track persistence
	CookieName pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringPtrInput
	// Expiration TTL for cookie specified in D:H:M:S or in seconds
	Expiration pulumi.StringPtrInput
	// Length of hash to apply to cookie
	HashLength pulumi.IntPtrInput
	// Number of characters to skip in the cookie for the hash
	HashOffset pulumi.IntPtrInput
	// To enable _ disable sending only over http
	Httponly pulumi.StringPtrInput
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrInput
	// To enable _ disable match across virtual servers with given persistence record
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

func (PersistenceProfileCookieState) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileCookieState)(nil)).Elem()
}

type persistenceProfileCookieArgs struct {
	// To enable _ disable always sending cookies
	AlwaysSend *string `pulumi:"alwaysSend"`
	AppService *string `pulumi:"appService"`
	// To required, preferred, or disabled policy for cookie encryption
	CookieEncryption *string `pulumi:"cookieEncryption"`
	// Passphrase for encrypted cookies
	CookieEncryptionPassphrase *string `pulumi:"cookieEncryptionPassphrase"`
	// Name of the cookie to track persistence
	CookieName *string `pulumi:"cookieName"`
	// Inherit defaults from parent profile
	DefaultsFrom string `pulumi:"defaultsFrom"`
	// Expiration TTL for cookie specified in D:H:M:S or in seconds
	Expiration *string `pulumi:"expiration"`
	// Length of hash to apply to cookie
	HashLength *int `pulumi:"hashLength"`
	// Number of characters to skip in the cookie for the hash
	HashOffset *int `pulumi:"hashOffset"`
	// To enable _ disable sending only over http
	Httponly *string `pulumi:"httponly"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools *string `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices *string `pulumi:"matchAcrossServices"`
	// To enable _ disable match across virtual servers with given persistence record
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

// The set of arguments for constructing a PersistenceProfileCookie resource.
type PersistenceProfileCookieArgs struct {
	// To enable _ disable always sending cookies
	AlwaysSend pulumi.StringPtrInput
	AppService pulumi.StringPtrInput
	// To required, preferred, or disabled policy for cookie encryption
	CookieEncryption pulumi.StringPtrInput
	// Passphrase for encrypted cookies
	CookieEncryptionPassphrase pulumi.StringPtrInput
	// Name of the cookie to track persistence
	CookieName pulumi.StringPtrInput
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringInput
	// Expiration TTL for cookie specified in D:H:M:S or in seconds
	Expiration pulumi.StringPtrInput
	// Length of hash to apply to cookie
	HashLength pulumi.IntPtrInput
	// Number of characters to skip in the cookie for the hash
	HashOffset pulumi.IntPtrInput
	// To enable _ disable sending only over http
	Httponly pulumi.StringPtrInput
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringPtrInput
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringPtrInput
	// To enable _ disable match across virtual servers with given persistence record
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

func (PersistenceProfileCookieArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*persistenceProfileCookieArgs)(nil)).Elem()
}
