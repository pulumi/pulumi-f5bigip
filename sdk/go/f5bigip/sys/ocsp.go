// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `sys.Ocsp` Manages F5 BIG-IP OCSP responder using iControl REST.
type Ocsp struct {
	pulumi.CustomResourceState

	// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
	CacheErrorTimeout pulumi.IntPtrOutput `pulumi:"cacheErrorTimeout"`
	// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
	CacheTimeout pulumi.StringPtrOutput `pulumi:"cacheTimeout"`
	// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
	ClockSkew pulumi.IntPtrOutput `pulumi:"clockSkew"`
	// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
	ConcurrentConnectionsLimit pulumi.IntPtrOutput `pulumi:"concurrentConnectionsLimit"`
	// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
	ConnectionTimeout pulumi.IntPtrOutput `pulumi:"connectionTimeout"`
	// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
	DnsResolver pulumi.StringPtrOutput `pulumi:"dnsResolver"`
	// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a passphrase used to sign an OCSP request.
	Passphrase pulumi.StringPtrOutput `pulumi:"passphrase"`
	// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
	ProxyServerPool pulumi.StringPtrOutput `pulumi:"proxyServerPool"`
	// Specifies the URL of the OCSP responder.
	ResponderUrl pulumi.StringPtrOutput `pulumi:"responderUrl"`
	// Specifies the route domain for the OCSP responder.
	RouteDomain pulumi.StringPtrOutput `pulumi:"routeDomain"`
	// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
	SignHash pulumi.StringPtrOutput `pulumi:"signHash"`
	// Specifies the certificate used to sign the OCSP request.
	SignerCert pulumi.StringPtrOutput `pulumi:"signerCert"`
	// Specifies the key used to sign the OCSP request.
	SignerKey pulumi.StringPtrOutput `pulumi:"signerKey"`
	// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
	StatusAge pulumi.IntPtrOutput `pulumi:"statusAge"`
	// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
	StrictRespCertCheck pulumi.StringPtrOutput `pulumi:"strictRespCertCheck"`
	// Specifies the certificates used for validating the OCSP response.
	TrustedResponders pulumi.StringPtrOutput `pulumi:"trustedResponders"`
}

// NewOcsp registers a new resource with the given unique name, arguments, and options.
func NewOcsp(ctx *pulumi.Context,
	name string, args *OcspArgs, opts ...pulumi.ResourceOption) (*Ocsp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Passphrase != nil {
		args.Passphrase = pulumi.ToSecret(args.Passphrase).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"passphrase",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Ocsp
	err := ctx.RegisterResource("f5bigip:sys/ocsp:Ocsp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOcsp gets an existing Ocsp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOcsp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OcspState, opts ...pulumi.ResourceOption) (*Ocsp, error) {
	var resource Ocsp
	err := ctx.ReadResource("f5bigip:sys/ocsp:Ocsp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ocsp resources.
type ocspState struct {
	// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
	CacheErrorTimeout *int `pulumi:"cacheErrorTimeout"`
	// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
	CacheTimeout *string `pulumi:"cacheTimeout"`
	// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
	ClockSkew *int `pulumi:"clockSkew"`
	// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
	ConcurrentConnectionsLimit *int `pulumi:"concurrentConnectionsLimit"`
	// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
	DnsResolver *string `pulumi:"dnsResolver"`
	// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
	Name *string `pulumi:"name"`
	// Specifies a passphrase used to sign an OCSP request.
	Passphrase *string `pulumi:"passphrase"`
	// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
	ProxyServerPool *string `pulumi:"proxyServerPool"`
	// Specifies the URL of the OCSP responder.
	ResponderUrl *string `pulumi:"responderUrl"`
	// Specifies the route domain for the OCSP responder.
	RouteDomain *string `pulumi:"routeDomain"`
	// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
	SignHash *string `pulumi:"signHash"`
	// Specifies the certificate used to sign the OCSP request.
	SignerCert *string `pulumi:"signerCert"`
	// Specifies the key used to sign the OCSP request.
	SignerKey *string `pulumi:"signerKey"`
	// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
	StatusAge *int `pulumi:"statusAge"`
	// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
	StrictRespCertCheck *string `pulumi:"strictRespCertCheck"`
	// Specifies the certificates used for validating the OCSP response.
	TrustedResponders *string `pulumi:"trustedResponders"`
}

type OcspState struct {
	// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
	CacheErrorTimeout pulumi.IntPtrInput
	// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
	CacheTimeout pulumi.StringPtrInput
	// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
	ClockSkew pulumi.IntPtrInput
	// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
	ConcurrentConnectionsLimit pulumi.IntPtrInput
	// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
	ConnectionTimeout pulumi.IntPtrInput
	// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
	DnsResolver pulumi.StringPtrInput
	// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
	Name pulumi.StringPtrInput
	// Specifies a passphrase used to sign an OCSP request.
	Passphrase pulumi.StringPtrInput
	// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
	ProxyServerPool pulumi.StringPtrInput
	// Specifies the URL of the OCSP responder.
	ResponderUrl pulumi.StringPtrInput
	// Specifies the route domain for the OCSP responder.
	RouteDomain pulumi.StringPtrInput
	// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
	SignHash pulumi.StringPtrInput
	// Specifies the certificate used to sign the OCSP request.
	SignerCert pulumi.StringPtrInput
	// Specifies the key used to sign the OCSP request.
	SignerKey pulumi.StringPtrInput
	// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
	StatusAge pulumi.IntPtrInput
	// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
	StrictRespCertCheck pulumi.StringPtrInput
	// Specifies the certificates used for validating the OCSP response.
	TrustedResponders pulumi.StringPtrInput
}

func (OcspState) ElementType() reflect.Type {
	return reflect.TypeOf((*ocspState)(nil)).Elem()
}

type ocspArgs struct {
	// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
	CacheErrorTimeout *int `pulumi:"cacheErrorTimeout"`
	// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
	CacheTimeout *string `pulumi:"cacheTimeout"`
	// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
	ClockSkew *int `pulumi:"clockSkew"`
	// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
	ConcurrentConnectionsLimit *int `pulumi:"concurrentConnectionsLimit"`
	// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
	ConnectionTimeout *int `pulumi:"connectionTimeout"`
	// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
	DnsResolver *string `pulumi:"dnsResolver"`
	// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
	Name string `pulumi:"name"`
	// Specifies a passphrase used to sign an OCSP request.
	Passphrase *string `pulumi:"passphrase"`
	// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
	ProxyServerPool *string `pulumi:"proxyServerPool"`
	// Specifies the URL of the OCSP responder.
	ResponderUrl *string `pulumi:"responderUrl"`
	// Specifies the route domain for the OCSP responder.
	RouteDomain *string `pulumi:"routeDomain"`
	// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
	SignHash *string `pulumi:"signHash"`
	// Specifies the certificate used to sign the OCSP request.
	SignerCert *string `pulumi:"signerCert"`
	// Specifies the key used to sign the OCSP request.
	SignerKey *string `pulumi:"signerKey"`
	// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
	StatusAge *int `pulumi:"statusAge"`
	// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
	StrictRespCertCheck *string `pulumi:"strictRespCertCheck"`
	// Specifies the certificates used for validating the OCSP response.
	TrustedResponders *string `pulumi:"trustedResponders"`
}

// The set of arguments for constructing a Ocsp resource.
type OcspArgs struct {
	// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
	CacheErrorTimeout pulumi.IntPtrInput
	// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
	CacheTimeout pulumi.StringPtrInput
	// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
	ClockSkew pulumi.IntPtrInput
	// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
	ConcurrentConnectionsLimit pulumi.IntPtrInput
	// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
	ConnectionTimeout pulumi.IntPtrInput
	// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
	DnsResolver pulumi.StringPtrInput
	// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
	Name pulumi.StringInput
	// Specifies a passphrase used to sign an OCSP request.
	Passphrase pulumi.StringPtrInput
	// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
	ProxyServerPool pulumi.StringPtrInput
	// Specifies the URL of the OCSP responder.
	ResponderUrl pulumi.StringPtrInput
	// Specifies the route domain for the OCSP responder.
	RouteDomain pulumi.StringPtrInput
	// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
	SignHash pulumi.StringPtrInput
	// Specifies the certificate used to sign the OCSP request.
	SignerCert pulumi.StringPtrInput
	// Specifies the key used to sign the OCSP request.
	SignerKey pulumi.StringPtrInput
	// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
	StatusAge pulumi.IntPtrInput
	// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
	StrictRespCertCheck pulumi.StringPtrInput
	// Specifies the certificates used for validating the OCSP response.
	TrustedResponders pulumi.StringPtrInput
}

func (OcspArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ocspArgs)(nil)).Elem()
}

type OcspInput interface {
	pulumi.Input

	ToOcspOutput() OcspOutput
	ToOcspOutputWithContext(ctx context.Context) OcspOutput
}

func (*Ocsp) ElementType() reflect.Type {
	return reflect.TypeOf((**Ocsp)(nil)).Elem()
}

func (i *Ocsp) ToOcspOutput() OcspOutput {
	return i.ToOcspOutputWithContext(context.Background())
}

func (i *Ocsp) ToOcspOutputWithContext(ctx context.Context) OcspOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OcspOutput)
}

func (i *Ocsp) ToOutput(ctx context.Context) pulumix.Output[*Ocsp] {
	return pulumix.Output[*Ocsp]{
		OutputState: i.ToOcspOutputWithContext(ctx).OutputState,
	}
}

// OcspArrayInput is an input type that accepts OcspArray and OcspArrayOutput values.
// You can construct a concrete instance of `OcspArrayInput` via:
//
//	OcspArray{ OcspArgs{...} }
type OcspArrayInput interface {
	pulumi.Input

	ToOcspArrayOutput() OcspArrayOutput
	ToOcspArrayOutputWithContext(context.Context) OcspArrayOutput
}

type OcspArray []OcspInput

func (OcspArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ocsp)(nil)).Elem()
}

func (i OcspArray) ToOcspArrayOutput() OcspArrayOutput {
	return i.ToOcspArrayOutputWithContext(context.Background())
}

func (i OcspArray) ToOcspArrayOutputWithContext(ctx context.Context) OcspArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OcspArrayOutput)
}

func (i OcspArray) ToOutput(ctx context.Context) pulumix.Output[[]*Ocsp] {
	return pulumix.Output[[]*Ocsp]{
		OutputState: i.ToOcspArrayOutputWithContext(ctx).OutputState,
	}
}

// OcspMapInput is an input type that accepts OcspMap and OcspMapOutput values.
// You can construct a concrete instance of `OcspMapInput` via:
//
//	OcspMap{ "key": OcspArgs{...} }
type OcspMapInput interface {
	pulumi.Input

	ToOcspMapOutput() OcspMapOutput
	ToOcspMapOutputWithContext(context.Context) OcspMapOutput
}

type OcspMap map[string]OcspInput

func (OcspMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ocsp)(nil)).Elem()
}

func (i OcspMap) ToOcspMapOutput() OcspMapOutput {
	return i.ToOcspMapOutputWithContext(context.Background())
}

func (i OcspMap) ToOcspMapOutputWithContext(ctx context.Context) OcspMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OcspMapOutput)
}

func (i OcspMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ocsp] {
	return pulumix.Output[map[string]*Ocsp]{
		OutputState: i.ToOcspMapOutputWithContext(ctx).OutputState,
	}
}

type OcspOutput struct{ *pulumi.OutputState }

func (OcspOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ocsp)(nil)).Elem()
}

func (o OcspOutput) ToOcspOutput() OcspOutput {
	return o
}

func (o OcspOutput) ToOcspOutputWithContext(ctx context.Context) OcspOutput {
	return o
}

func (o OcspOutput) ToOutput(ctx context.Context) pulumix.Output[*Ocsp] {
	return pulumix.Output[*Ocsp]{
		OutputState: o.OutputState,
	}
}

// Specifies the lifetime of an error response in the cache, in seconds. This value must be greater than connection_timeout. The default value is `3600`.
func (o OcspOutput) CacheErrorTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.IntPtrOutput { return v.CacheErrorTimeout }).(pulumi.IntPtrOutput)
}

// Specifies the lifetime of the OCSP response in the cache, in seconds. The default value is `indefinite`.
func (o OcspOutput) CacheTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.CacheTimeout }).(pulumi.StringPtrOutput)
}

// Specifies the time interval that the BIG-IP system allows for clock skew, in seconds. The default value is `300`.
func (o OcspOutput) ClockSkew() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.IntPtrOutput { return v.ClockSkew }).(pulumi.IntPtrOutput)
}

// Specifies the maximum number of connections per second allowed for the OCSP certificate validator. The default value is `50`.
func (o OcspOutput) ConcurrentConnectionsLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.IntPtrOutput { return v.ConcurrentConnectionsLimit }).(pulumi.IntPtrOutput)
}

// Specifies the time interval that the BIG-IP system waits for before ending the connection to the OCSP responder, in seconds. The default value is `8`.
func (o OcspOutput) ConnectionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.IntPtrOutput { return v.ConnectionTimeout }).(pulumi.IntPtrOutput)
}

// Specifies the internal DNS resolver the BIG-IP system uses to fetch the OCSP response.
func (o OcspOutput) DnsResolver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.DnsResolver }).(pulumi.StringPtrOutput)
}

// Name of the OCSP Responder. Name should be in pattern `/partition/ocsp_name`.
func (o OcspOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies a passphrase used to sign an OCSP request.
func (o OcspOutput) Passphrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.Passphrase }).(pulumi.StringPtrOutput)
}

// Specifies the proxy server pool the BIG-IP system uses to fetch the OCSP response.
func (o OcspOutput) ProxyServerPool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.ProxyServerPool }).(pulumi.StringPtrOutput)
}

// Specifies the URL of the OCSP responder.
func (o OcspOutput) ResponderUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.ResponderUrl }).(pulumi.StringPtrOutput)
}

// Specifies the route domain for the OCSP responder.
func (o OcspOutput) RouteDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.RouteDomain }).(pulumi.StringPtrOutput)
}

// Specifies the hash algorithm used to sign the OCSP request. The default value is `sha256`.
func (o OcspOutput) SignHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.SignHash }).(pulumi.StringPtrOutput)
}

// Specifies the certificate used to sign the OCSP request.
func (o OcspOutput) SignerCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.SignerCert }).(pulumi.StringPtrOutput)
}

// Specifies the key used to sign the OCSP request.
func (o OcspOutput) SignerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.SignerKey }).(pulumi.StringPtrOutput)
}

// Specifies the maximum allowed lag time that the BIG-IP system accepts for the 'thisUpdate' time in the OCSP response, in seconds. The default value is `0`.
func (o OcspOutput) StatusAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.IntPtrOutput { return v.StatusAge }).(pulumi.IntPtrOutput)
}

// Specifies whether the responder's certificate is checked for an OCSP signing extension. The default value is `enabled`.
func (o OcspOutput) StrictRespCertCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.StrictRespCertCheck }).(pulumi.StringPtrOutput)
}

// Specifies the certificates used for validating the OCSP response.
func (o OcspOutput) TrustedResponders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ocsp) pulumi.StringPtrOutput { return v.TrustedResponders }).(pulumi.StringPtrOutput)
}

type OcspArrayOutput struct{ *pulumi.OutputState }

func (OcspArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ocsp)(nil)).Elem()
}

func (o OcspArrayOutput) ToOcspArrayOutput() OcspArrayOutput {
	return o
}

func (o OcspArrayOutput) ToOcspArrayOutputWithContext(ctx context.Context) OcspArrayOutput {
	return o
}

func (o OcspArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Ocsp] {
	return pulumix.Output[[]*Ocsp]{
		OutputState: o.OutputState,
	}
}

func (o OcspArrayOutput) Index(i pulumi.IntInput) OcspOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ocsp {
		return vs[0].([]*Ocsp)[vs[1].(int)]
	}).(OcspOutput)
}

type OcspMapOutput struct{ *pulumi.OutputState }

func (OcspMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ocsp)(nil)).Elem()
}

func (o OcspMapOutput) ToOcspMapOutput() OcspMapOutput {
	return o
}

func (o OcspMapOutput) ToOcspMapOutputWithContext(ctx context.Context) OcspMapOutput {
	return o
}

func (o OcspMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Ocsp] {
	return pulumix.Output[map[string]*Ocsp]{
		OutputState: o.OutputState,
	}
}

func (o OcspMapOutput) MapIndex(k pulumi.StringInput) OcspOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ocsp {
		return vs[0].(map[string]*Ocsp)[vs[1].(string)]
	}).(OcspOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OcspInput)(nil)).Elem(), &Ocsp{})
	pulumi.RegisterInputType(reflect.TypeOf((*OcspArrayInput)(nil)).Elem(), OcspArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OcspMapInput)(nil)).Elem(), OcspMap{})
	pulumi.RegisterOutputType(OcspOutput{})
	pulumi.RegisterOutputType(OcspArrayOutput{})
	pulumi.RegisterOutputType(OcspMapOutput{})
}
