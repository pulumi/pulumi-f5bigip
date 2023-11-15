// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures a cookie persistence profile
//
// ## Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/ltm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ltm.NewPersistenceProfileCookie(ctx, "testPpcookie", &ltm.PersistenceProfileCookieArgs{
//				Name:                       pulumi.String("/Common/terraform_cookie"),
//				DefaultsFrom:               pulumi.String("/Common/cookie"),
//				MatchAcrossPools:           pulumi.String("enabled"),
//				MatchAcrossServices:        pulumi.String("enabled"),
//				MatchAcrossVirtuals:        pulumi.String("enabled"),
//				Timeout:                    pulumi.Int(3600),
//				OverrideConnLimit:          pulumi.String("enabled"),
//				AlwaysSend:                 pulumi.String("enabled"),
//				CookieEncryption:           pulumi.String("required"),
//				CookieEncryptionPassphrase: pulumi.String("iam"),
//				CookieName:                 pulumi.String("ham"),
//				Expiration:                 pulumi.String("1:0:0"),
//				HashLength:                 pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
// `method` (Optional) Specifies the type of cookie processing that the system uses. The default value is insert
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
type PersistenceProfileCookie struct {
	pulumi.CustomResourceState

	// To enable _ disable always sending cookies
	AlwaysSend pulumi.StringOutput `pulumi:"alwaysSend"`
	AppService pulumi.StringOutput `pulumi:"appService"`
	// To required, preferred, or disabled policy for cookie encryption
	CookieEncryption pulumi.StringOutput `pulumi:"cookieEncryption"`
	// Passphrase for encrypted cookies
	CookieEncryptionPassphrase pulumi.StringOutput `pulumi:"cookieEncryptionPassphrase"`
	// Name of the cookie to track persistence
	CookieName pulumi.StringOutput `pulumi:"cookieName"`
	// Inherit defaults from parent profile
	DefaultsFrom pulumi.StringOutput `pulumi:"defaultsFrom"`
	// Expiration TTL for cookie specified in D:H:M:S or in seconds
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// Length of hash to apply to cookie
	HashLength pulumi.IntOutput `pulumi:"hashLength"`
	// Number of characters to skip in the cookie for the hash
	HashOffset pulumi.IntOutput `pulumi:"hashOffset"`
	// To enable _ disable sending only over http
	Httponly pulumi.StringOutput `pulumi:"httponly"`
	// To enable _ disable match across pools with given persistence record
	MatchAcrossPools pulumi.StringOutput `pulumi:"matchAcrossPools"`
	// To enable _ disable match across services with given persistence record
	MatchAcrossServices pulumi.StringOutput `pulumi:"matchAcrossServices"`
	// To enable _ disable match across virtual servers with given persistence record
	MatchAcrossVirtuals pulumi.StringOutput `pulumi:"matchAcrossVirtuals"`
	// Specifies the type of cookie processing that the system uses
	Method pulumi.StringOutput `pulumi:"method"`
	// To enable _ disable
	Mirror pulumi.StringOutput `pulumi:"mirror"`
	// Name of the persistence profile
	Name pulumi.StringOutput `pulumi:"name"`
	// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
	// limits remain hard limits and are not overridden.
	OverrideConnLimit pulumi.StringOutput `pulumi:"overrideConnLimit"`
	// Timeout for persistence of the session
	Timeout pulumi.IntOutput `pulumi:"timeout"`
}

// NewPersistenceProfileCookie registers a new resource with the given unique name, arguments, and options.
func NewPersistenceProfileCookie(ctx *pulumi.Context,
	name string, args *PersistenceProfileCookieArgs, opts ...pulumi.ResourceOption) (*PersistenceProfileCookie, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultsFrom == nil {
		return nil, errors.New("invalid value for required argument 'DefaultsFrom'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Specifies the type of cookie processing that the system uses
	Method *string `pulumi:"method"`
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
	// Specifies the type of cookie processing that the system uses
	Method pulumi.StringPtrInput
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
	// Specifies the type of cookie processing that the system uses
	Method *string `pulumi:"method"`
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
	// Specifies the type of cookie processing that the system uses
	Method pulumi.StringPtrInput
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

type PersistenceProfileCookieInput interface {
	pulumi.Input

	ToPersistenceProfileCookieOutput() PersistenceProfileCookieOutput
	ToPersistenceProfileCookieOutputWithContext(ctx context.Context) PersistenceProfileCookieOutput
}

func (*PersistenceProfileCookie) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileCookie)(nil)).Elem()
}

func (i *PersistenceProfileCookie) ToPersistenceProfileCookieOutput() PersistenceProfileCookieOutput {
	return i.ToPersistenceProfileCookieOutputWithContext(context.Background())
}

func (i *PersistenceProfileCookie) ToPersistenceProfileCookieOutputWithContext(ctx context.Context) PersistenceProfileCookieOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileCookieOutput)
}

// PersistenceProfileCookieArrayInput is an input type that accepts PersistenceProfileCookieArray and PersistenceProfileCookieArrayOutput values.
// You can construct a concrete instance of `PersistenceProfileCookieArrayInput` via:
//
//	PersistenceProfileCookieArray{ PersistenceProfileCookieArgs{...} }
type PersistenceProfileCookieArrayInput interface {
	pulumi.Input

	ToPersistenceProfileCookieArrayOutput() PersistenceProfileCookieArrayOutput
	ToPersistenceProfileCookieArrayOutputWithContext(context.Context) PersistenceProfileCookieArrayOutput
}

type PersistenceProfileCookieArray []PersistenceProfileCookieInput

func (PersistenceProfileCookieArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileCookie)(nil)).Elem()
}

func (i PersistenceProfileCookieArray) ToPersistenceProfileCookieArrayOutput() PersistenceProfileCookieArrayOutput {
	return i.ToPersistenceProfileCookieArrayOutputWithContext(context.Background())
}

func (i PersistenceProfileCookieArray) ToPersistenceProfileCookieArrayOutputWithContext(ctx context.Context) PersistenceProfileCookieArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileCookieArrayOutput)
}

// PersistenceProfileCookieMapInput is an input type that accepts PersistenceProfileCookieMap and PersistenceProfileCookieMapOutput values.
// You can construct a concrete instance of `PersistenceProfileCookieMapInput` via:
//
//	PersistenceProfileCookieMap{ "key": PersistenceProfileCookieArgs{...} }
type PersistenceProfileCookieMapInput interface {
	pulumi.Input

	ToPersistenceProfileCookieMapOutput() PersistenceProfileCookieMapOutput
	ToPersistenceProfileCookieMapOutputWithContext(context.Context) PersistenceProfileCookieMapOutput
}

type PersistenceProfileCookieMap map[string]PersistenceProfileCookieInput

func (PersistenceProfileCookieMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileCookie)(nil)).Elem()
}

func (i PersistenceProfileCookieMap) ToPersistenceProfileCookieMapOutput() PersistenceProfileCookieMapOutput {
	return i.ToPersistenceProfileCookieMapOutputWithContext(context.Background())
}

func (i PersistenceProfileCookieMap) ToPersistenceProfileCookieMapOutputWithContext(ctx context.Context) PersistenceProfileCookieMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PersistenceProfileCookieMapOutput)
}

type PersistenceProfileCookieOutput struct{ *pulumi.OutputState }

func (PersistenceProfileCookieOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistenceProfileCookie)(nil)).Elem()
}

func (o PersistenceProfileCookieOutput) ToPersistenceProfileCookieOutput() PersistenceProfileCookieOutput {
	return o
}

func (o PersistenceProfileCookieOutput) ToPersistenceProfileCookieOutputWithContext(ctx context.Context) PersistenceProfileCookieOutput {
	return o
}

// To enable _ disable always sending cookies
func (o PersistenceProfileCookieOutput) AlwaysSend() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.AlwaysSend }).(pulumi.StringOutput)
}

func (o PersistenceProfileCookieOutput) AppService() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.AppService }).(pulumi.StringOutput)
}

// To required, preferred, or disabled policy for cookie encryption
func (o PersistenceProfileCookieOutput) CookieEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.CookieEncryption }).(pulumi.StringOutput)
}

// Passphrase for encrypted cookies
func (o PersistenceProfileCookieOutput) CookieEncryptionPassphrase() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.CookieEncryptionPassphrase }).(pulumi.StringOutput)
}

// Name of the cookie to track persistence
func (o PersistenceProfileCookieOutput) CookieName() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.CookieName }).(pulumi.StringOutput)
}

// Inherit defaults from parent profile
func (o PersistenceProfileCookieOutput) DefaultsFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.DefaultsFrom }).(pulumi.StringOutput)
}

// Expiration TTL for cookie specified in D:H:M:S or in seconds
func (o PersistenceProfileCookieOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

// Length of hash to apply to cookie
func (o PersistenceProfileCookieOutput) HashLength() pulumi.IntOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.IntOutput { return v.HashLength }).(pulumi.IntOutput)
}

// Number of characters to skip in the cookie for the hash
func (o PersistenceProfileCookieOutput) HashOffset() pulumi.IntOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.IntOutput { return v.HashOffset }).(pulumi.IntOutput)
}

// To enable _ disable sending only over http
func (o PersistenceProfileCookieOutput) Httponly() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.Httponly }).(pulumi.StringOutput)
}

// To enable _ disable match across pools with given persistence record
func (o PersistenceProfileCookieOutput) MatchAcrossPools() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.MatchAcrossPools }).(pulumi.StringOutput)
}

// To enable _ disable match across services with given persistence record
func (o PersistenceProfileCookieOutput) MatchAcrossServices() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.MatchAcrossServices }).(pulumi.StringOutput)
}

// To enable _ disable match across virtual servers with given persistence record
func (o PersistenceProfileCookieOutput) MatchAcrossVirtuals() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.MatchAcrossVirtuals }).(pulumi.StringOutput)
}

// Specifies the type of cookie processing that the system uses
func (o PersistenceProfileCookieOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.Method }).(pulumi.StringOutput)
}

// To enable _ disable
func (o PersistenceProfileCookieOutput) Mirror() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.Mirror }).(pulumi.StringOutput)
}

// Name of the persistence profile
func (o PersistenceProfileCookieOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// To enable _ disable that pool member connection limits are overridden for persisted clients. Per-virtual connection
// limits remain hard limits and are not overridden.
func (o PersistenceProfileCookieOutput) OverrideConnLimit() pulumi.StringOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.StringOutput { return v.OverrideConnLimit }).(pulumi.StringOutput)
}

// Timeout for persistence of the session
func (o PersistenceProfileCookieOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *PersistenceProfileCookie) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

type PersistenceProfileCookieArrayOutput struct{ *pulumi.OutputState }

func (PersistenceProfileCookieArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PersistenceProfileCookie)(nil)).Elem()
}

func (o PersistenceProfileCookieArrayOutput) ToPersistenceProfileCookieArrayOutput() PersistenceProfileCookieArrayOutput {
	return o
}

func (o PersistenceProfileCookieArrayOutput) ToPersistenceProfileCookieArrayOutputWithContext(ctx context.Context) PersistenceProfileCookieArrayOutput {
	return o
}

func (o PersistenceProfileCookieArrayOutput) Index(i pulumi.IntInput) PersistenceProfileCookieOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PersistenceProfileCookie {
		return vs[0].([]*PersistenceProfileCookie)[vs[1].(int)]
	}).(PersistenceProfileCookieOutput)
}

type PersistenceProfileCookieMapOutput struct{ *pulumi.OutputState }

func (PersistenceProfileCookieMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PersistenceProfileCookie)(nil)).Elem()
}

func (o PersistenceProfileCookieMapOutput) ToPersistenceProfileCookieMapOutput() PersistenceProfileCookieMapOutput {
	return o
}

func (o PersistenceProfileCookieMapOutput) ToPersistenceProfileCookieMapOutputWithContext(ctx context.Context) PersistenceProfileCookieMapOutput {
	return o
}

func (o PersistenceProfileCookieMapOutput) MapIndex(k pulumi.StringInput) PersistenceProfileCookieOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PersistenceProfileCookie {
		return vs[0].(map[string]*PersistenceProfileCookie)[vs[1].(string)]
	}).(PersistenceProfileCookieOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileCookieInput)(nil)).Elem(), &PersistenceProfileCookie{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileCookieArrayInput)(nil)).Elem(), PersistenceProfileCookieArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PersistenceProfileCookieMapInput)(nil)).Elem(), PersistenceProfileCookieMap{})
	pulumi.RegisterOutputType(PersistenceProfileCookieOutput{})
	pulumi.RegisterOutputType(PersistenceProfileCookieArrayOutput{})
	pulumi.RegisterOutputType(PersistenceProfileCookieMapOutput{})
}
