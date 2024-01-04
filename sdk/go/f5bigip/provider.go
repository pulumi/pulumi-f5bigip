// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package f5bigip

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the bigip package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// Domain name/IP of the BigIP
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// Login reference for token authentication (see BIG-IP REST docs for details)
	LoginRef pulumi.StringPtrOutput `pulumi:"loginRef"`
	// The user's password. Leave empty if using token_value
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Management Port to connect to Bigip
	Port pulumi.StringPtrOutput `pulumi:"port"`
	// A token generated outside the provider, in place of password
	TokenValue pulumi.StringPtrOutput `pulumi:"tokenValue"`
	// Valid Trusted Certificate path
	TrustedCertPath pulumi.StringPtrOutput `pulumi:"trustedCertPath"`
	// Username with API access to the BigIP
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:f5bigip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Domain name/IP of the BigIP
	Address *string `pulumi:"address"`
	// Amount of times to retry AS3 API requests. Default: 10.
	ApiRetries *int `pulumi:"apiRetries"`
	// A timeout for AS3 requests, represented as a number of seconds. Default: 60
	ApiTimeout *int `pulumi:"apiTimeout"`
	// Login reference for token authentication (see BIG-IP REST docs for details)
	LoginRef *string `pulumi:"loginRef"`
	// The user's password. Leave empty if using token_value
	Password *string `pulumi:"password"`
	// Management Port to connect to Bigip
	Port *string `pulumi:"port"`
	// If this flag set to true,sending telemetry data to TEEM will be disabled
	TeemDisable *bool `pulumi:"teemDisable"`
	// Enable to use an external authentication source (LDAP, TACACS, etc)
	TokenAuth *bool `pulumi:"tokenAuth"`
	// A lifespan to request for the AS3 auth token, represented as a number of seconds. Default: 1200
	TokenTimeout *int `pulumi:"tokenTimeout"`
	// A token generated outside the provider, in place of password
	TokenValue *string `pulumi:"tokenValue"`
	// Valid Trusted Certificate path
	TrustedCertPath *string `pulumi:"trustedCertPath"`
	// Username with API access to the BigIP
	Username *string `pulumi:"username"`
	// If set to true, Disables TLS certificate check on BIG-IP. Default : True
	ValidateCertsDisable *bool `pulumi:"validateCertsDisable"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Domain name/IP of the BigIP
	Address pulumi.StringPtrInput
	// Amount of times to retry AS3 API requests. Default: 10.
	ApiRetries pulumi.IntPtrInput
	// A timeout for AS3 requests, represented as a number of seconds. Default: 60
	ApiTimeout pulumi.IntPtrInput
	// Login reference for token authentication (see BIG-IP REST docs for details)
	LoginRef pulumi.StringPtrInput
	// The user's password. Leave empty if using token_value
	Password pulumi.StringPtrInput
	// Management Port to connect to Bigip
	Port pulumi.StringPtrInput
	// If this flag set to true,sending telemetry data to TEEM will be disabled
	TeemDisable pulumi.BoolPtrInput
	// Enable to use an external authentication source (LDAP, TACACS, etc)
	TokenAuth pulumi.BoolPtrInput
	// A lifespan to request for the AS3 auth token, represented as a number of seconds. Default: 1200
	TokenTimeout pulumi.IntPtrInput
	// A token generated outside the provider, in place of password
	TokenValue pulumi.StringPtrInput
	// Valid Trusted Certificate path
	TrustedCertPath pulumi.StringPtrInput
	// Username with API access to the BigIP
	Username pulumi.StringPtrInput
	// If set to true, Disables TLS certificate check on BIG-IP. Default : True
	ValidateCertsDisable pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// Domain name/IP of the BigIP
func (o ProviderOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

// Login reference for token authentication (see BIG-IP REST docs for details)
func (o ProviderOutput) LoginRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.LoginRef }).(pulumi.StringPtrOutput)
}

// The user's password. Leave empty if using token_value
func (o ProviderOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Management Port to connect to Bigip
func (o ProviderOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Port }).(pulumi.StringPtrOutput)
}

// A token generated outside the provider, in place of password
func (o ProviderOutput) TokenValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TokenValue }).(pulumi.StringPtrOutput)
}

// Valid Trusted Certificate path
func (o ProviderOutput) TrustedCertPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.TrustedCertPath }).(pulumi.StringPtrOutput)
}

// Username with API access to the BigIP
func (o ProviderOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
