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

// `ltm.CipherRule` Manages F5 BIG-IP LTM cipher rule using iControl REST.
//
// ## Example Usage
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
//			_, err := ltm.NewCipherRule(ctx, "test_cipher_rule", &ltm.CipherRuleArgs{
//				Name:                pulumi.String("/Common/test_cipher_rule"),
//				Cipher:              pulumi.String("TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384"),
//				DhGroups:            pulumi.String("P256:P384:FFDHE2048:FFDHE3072:FFDHE4096"),
//				SignatureAlgorithms: pulumi.String("DEFAULT"),
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
// ## Importing
//
// An existing cipher rule can be imported into this resource by supplying the cipher rule full path name  ex : `/partition/name`
// An example is below:
// ```sh
// $ terraform import bigip_ltm_cipher_rule.test_cipher_rule /Common/test_cipher_rule
// ```
type CipherRule struct {
	pulumi.CustomResourceState

	// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
	Cipher pulumi.StringOutput `pulumi:"cipher"`
	// The Partition in which the Cipher Rule will be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the DH Groups algorithms, separated by colons (:).
	DhGroups pulumi.StringOutput `pulumi:"dhGroups"`
	// Name of the Cipher Rule. Name should be in pattern `partition` + `cipherRuleName`
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Signature Algorithms, separated by colons (:).
	SignatureAlgorithms pulumi.StringOutput `pulumi:"signatureAlgorithms"`
}

// NewCipherRule registers a new resource with the given unique name, arguments, and options.
func NewCipherRule(ctx *pulumi.Context,
	name string, args *CipherRuleArgs, opts ...pulumi.ResourceOption) (*CipherRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cipher == nil {
		return nil, errors.New("invalid value for required argument 'Cipher'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CipherRule
	err := ctx.RegisterResource("f5bigip:ltm/cipherRule:CipherRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCipherRule gets an existing CipherRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCipherRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CipherRuleState, opts ...pulumi.ResourceOption) (*CipherRule, error) {
	var resource CipherRule
	err := ctx.ReadResource("f5bigip:ltm/cipherRule:CipherRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CipherRule resources.
type cipherRuleState struct {
	// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
	Cipher *string `pulumi:"cipher"`
	// The Partition in which the Cipher Rule will be created.
	Description *string `pulumi:"description"`
	// Specifies the DH Groups algorithms, separated by colons (:).
	DhGroups *string `pulumi:"dhGroups"`
	// Name of the Cipher Rule. Name should be in pattern `partition` + `cipherRuleName`
	Name *string `pulumi:"name"`
	// Specifies the Signature Algorithms, separated by colons (:).
	SignatureAlgorithms *string `pulumi:"signatureAlgorithms"`
}

type CipherRuleState struct {
	// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
	Cipher pulumi.StringPtrInput
	// The Partition in which the Cipher Rule will be created.
	Description pulumi.StringPtrInput
	// Specifies the DH Groups algorithms, separated by colons (:).
	DhGroups pulumi.StringPtrInput
	// Name of the Cipher Rule. Name should be in pattern `partition` + `cipherRuleName`
	Name pulumi.StringPtrInput
	// Specifies the Signature Algorithms, separated by colons (:).
	SignatureAlgorithms pulumi.StringPtrInput
}

func (CipherRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*cipherRuleState)(nil)).Elem()
}

type cipherRuleArgs struct {
	// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
	Cipher string `pulumi:"cipher"`
	// The Partition in which the Cipher Rule will be created.
	Description *string `pulumi:"description"`
	// Specifies the DH Groups algorithms, separated by colons (:).
	DhGroups *string `pulumi:"dhGroups"`
	// Name of the Cipher Rule. Name should be in pattern `partition` + `cipherRuleName`
	Name string `pulumi:"name"`
	// Specifies the Signature Algorithms, separated by colons (:).
	SignatureAlgorithms *string `pulumi:"signatureAlgorithms"`
}

// The set of arguments for constructing a CipherRule resource.
type CipherRuleArgs struct {
	// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
	Cipher pulumi.StringInput
	// The Partition in which the Cipher Rule will be created.
	Description pulumi.StringPtrInput
	// Specifies the DH Groups algorithms, separated by colons (:).
	DhGroups pulumi.StringPtrInput
	// Name of the Cipher Rule. Name should be in pattern `partition` + `cipherRuleName`
	Name pulumi.StringInput
	// Specifies the Signature Algorithms, separated by colons (:).
	SignatureAlgorithms pulumi.StringPtrInput
}

func (CipherRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cipherRuleArgs)(nil)).Elem()
}

type CipherRuleInput interface {
	pulumi.Input

	ToCipherRuleOutput() CipherRuleOutput
	ToCipherRuleOutputWithContext(ctx context.Context) CipherRuleOutput
}

func (*CipherRule) ElementType() reflect.Type {
	return reflect.TypeOf((**CipherRule)(nil)).Elem()
}

func (i *CipherRule) ToCipherRuleOutput() CipherRuleOutput {
	return i.ToCipherRuleOutputWithContext(context.Background())
}

func (i *CipherRule) ToCipherRuleOutputWithContext(ctx context.Context) CipherRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CipherRuleOutput)
}

// CipherRuleArrayInput is an input type that accepts CipherRuleArray and CipherRuleArrayOutput values.
// You can construct a concrete instance of `CipherRuleArrayInput` via:
//
//	CipherRuleArray{ CipherRuleArgs{...} }
type CipherRuleArrayInput interface {
	pulumi.Input

	ToCipherRuleArrayOutput() CipherRuleArrayOutput
	ToCipherRuleArrayOutputWithContext(context.Context) CipherRuleArrayOutput
}

type CipherRuleArray []CipherRuleInput

func (CipherRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CipherRule)(nil)).Elem()
}

func (i CipherRuleArray) ToCipherRuleArrayOutput() CipherRuleArrayOutput {
	return i.ToCipherRuleArrayOutputWithContext(context.Background())
}

func (i CipherRuleArray) ToCipherRuleArrayOutputWithContext(ctx context.Context) CipherRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CipherRuleArrayOutput)
}

// CipherRuleMapInput is an input type that accepts CipherRuleMap and CipherRuleMapOutput values.
// You can construct a concrete instance of `CipherRuleMapInput` via:
//
//	CipherRuleMap{ "key": CipherRuleArgs{...} }
type CipherRuleMapInput interface {
	pulumi.Input

	ToCipherRuleMapOutput() CipherRuleMapOutput
	ToCipherRuleMapOutputWithContext(context.Context) CipherRuleMapOutput
}

type CipherRuleMap map[string]CipherRuleInput

func (CipherRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CipherRule)(nil)).Elem()
}

func (i CipherRuleMap) ToCipherRuleMapOutput() CipherRuleMapOutput {
	return i.ToCipherRuleMapOutputWithContext(context.Background())
}

func (i CipherRuleMap) ToCipherRuleMapOutputWithContext(ctx context.Context) CipherRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CipherRuleMapOutput)
}

type CipherRuleOutput struct{ *pulumi.OutputState }

func (CipherRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CipherRule)(nil)).Elem()
}

func (o CipherRuleOutput) ToCipherRuleOutput() CipherRuleOutput {
	return o
}

func (o CipherRuleOutput) ToCipherRuleOutputWithContext(ctx context.Context) CipherRuleOutput {
	return o
}

// Specifies one or more Cipher Suites used,this is a colon (:) separated string of cipher suites. example, `TLS13-AES128-GCM-SHA256:TLS13-AES256-GCM-SHA384`.
func (o CipherRuleOutput) Cipher() pulumi.StringOutput {
	return o.ApplyT(func(v *CipherRule) pulumi.StringOutput { return v.Cipher }).(pulumi.StringOutput)
}

// The Partition in which the Cipher Rule will be created.
func (o CipherRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CipherRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the DH Groups algorithms, separated by colons (:).
func (o CipherRuleOutput) DhGroups() pulumi.StringOutput {
	return o.ApplyT(func(v *CipherRule) pulumi.StringOutput { return v.DhGroups }).(pulumi.StringOutput)
}

// Name of the Cipher Rule. Name should be in pattern `partition` + `cipherRuleName`
func (o CipherRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CipherRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the Signature Algorithms, separated by colons (:).
func (o CipherRuleOutput) SignatureAlgorithms() pulumi.StringOutput {
	return o.ApplyT(func(v *CipherRule) pulumi.StringOutput { return v.SignatureAlgorithms }).(pulumi.StringOutput)
}

type CipherRuleArrayOutput struct{ *pulumi.OutputState }

func (CipherRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CipherRule)(nil)).Elem()
}

func (o CipherRuleArrayOutput) ToCipherRuleArrayOutput() CipherRuleArrayOutput {
	return o
}

func (o CipherRuleArrayOutput) ToCipherRuleArrayOutputWithContext(ctx context.Context) CipherRuleArrayOutput {
	return o
}

func (o CipherRuleArrayOutput) Index(i pulumi.IntInput) CipherRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CipherRule {
		return vs[0].([]*CipherRule)[vs[1].(int)]
	}).(CipherRuleOutput)
}

type CipherRuleMapOutput struct{ *pulumi.OutputState }

func (CipherRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CipherRule)(nil)).Elem()
}

func (o CipherRuleMapOutput) ToCipherRuleMapOutput() CipherRuleMapOutput {
	return o
}

func (o CipherRuleMapOutput) ToCipherRuleMapOutputWithContext(ctx context.Context) CipherRuleMapOutput {
	return o
}

func (o CipherRuleMapOutput) MapIndex(k pulumi.StringInput) CipherRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CipherRule {
		return vs[0].(map[string]*CipherRule)[vs[1].(string)]
	}).(CipherRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CipherRuleInput)(nil)).Elem(), &CipherRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*CipherRuleArrayInput)(nil)).Elem(), CipherRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CipherRuleMapInput)(nil)).Elem(), CipherRuleMap{})
	pulumi.RegisterOutputType(CipherRuleOutput{})
	pulumi.RegisterOutputType(CipherRuleArrayOutput{})
	pulumi.RegisterOutputType(CipherRuleMapOutput{})
}
