// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`fast.getConsulServiceDiscovery`) to get the Consul Service discovery config to be used for `http`/`https` app deployment in FAST.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/fast"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fast.GetConsulServiceDiscovery(ctx, &fast.GetConsulServiceDiscoveryArgs{
//				Port: 8080,
//				Uri:  "https://192.0.2.100:8500/v1/catalog/nodes",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetConsulServiceDiscovery(ctx *pulumi.Context, args *GetConsulServiceDiscoveryArgs, opts ...pulumi.InvokeOption) (*GetConsulServiceDiscoveryResult, error) {
	var rv GetConsulServiceDiscoveryResult
	err := ctx.Invoke("f5bigip:fast/getConsulServiceDiscovery:getConsulServiceDiscovery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConsulServiceDiscovery.
type GetConsulServiceDiscoveryArgs struct {
	// Specifies whether to look for public or private IP addresses,default `private`.
	AddressRealm *string `pulumi:"addressRealm"`
	// Specifies whether you are updating your credentials,default `false`.
	CredentialUpdate *bool `pulumi:"credentialUpdate"`
	// Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
	EncodedToken *string `pulumi:"encodedToken"`
	// Custom JMESPath Query.
	JmesPathQuery *string `pulumi:"jmesPathQuery"`
	// Member is down when fewer than minimum monitors report it healthy.
	MinimumMonitors *string `pulumi:"minimumMonitors"`
	// Port to be used for AWS service discovery,default `80`.
	Port int `pulumi:"port"`
	// If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
	RejectUnauthorized *bool `pulumi:"rejectUnauthorized"`
	// CA Bundle to validate server certificates.
	TrustCa *string `pulumi:"trustCa"`
	Type    *string `pulumi:"type"`
	// Action to take when node cannot be detected,default `remove`.
	UndetectableAction *string `pulumi:"undetectableAction"`
	// Update interval for service discovery.
	UpdateInterval *string `pulumi:"updateInterval"`
	// The location of the node data.
	Uri string `pulumi:"uri"`
}

// A collection of values returned by getConsulServiceDiscovery.
type GetConsulServiceDiscoveryResult struct {
	AddressRealm *string `pulumi:"addressRealm"`
	// The JSON for Hashicorp Consul service discovery block.
	ConsulSdJson     string  `pulumi:"consulSdJson"`
	CredentialUpdate *bool   `pulumi:"credentialUpdate"`
	EncodedToken     *string `pulumi:"encodedToken"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	JmesPathQuery      *string `pulumi:"jmesPathQuery"`
	MinimumMonitors    *string `pulumi:"minimumMonitors"`
	Port               int     `pulumi:"port"`
	RejectUnauthorized *bool   `pulumi:"rejectUnauthorized"`
	TrustCa            *string `pulumi:"trustCa"`
	Type               *string `pulumi:"type"`
	UndetectableAction *string `pulumi:"undetectableAction"`
	UpdateInterval     *string `pulumi:"updateInterval"`
	Uri                string  `pulumi:"uri"`
}

func GetConsulServiceDiscoveryOutput(ctx *pulumi.Context, args GetConsulServiceDiscoveryOutputArgs, opts ...pulumi.InvokeOption) GetConsulServiceDiscoveryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetConsulServiceDiscoveryResult, error) {
			args := v.(GetConsulServiceDiscoveryArgs)
			r, err := GetConsulServiceDiscovery(ctx, &args, opts...)
			var s GetConsulServiceDiscoveryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetConsulServiceDiscoveryResultOutput)
}

// A collection of arguments for invoking getConsulServiceDiscovery.
type GetConsulServiceDiscoveryOutputArgs struct {
	// Specifies whether to look for public or private IP addresses,default `private`.
	AddressRealm pulumi.StringPtrInput `pulumi:"addressRealm"`
	// Specifies whether you are updating your credentials,default `false`.
	CredentialUpdate pulumi.BoolPtrInput `pulumi:"credentialUpdate"`
	// Base 64 encoded bearer token to make requests to the Consul API. Will be stored in the declaration in an encrypted format.
	EncodedToken pulumi.StringPtrInput `pulumi:"encodedToken"`
	// Custom JMESPath Query.
	JmesPathQuery pulumi.StringPtrInput `pulumi:"jmesPathQuery"`
	// Member is down when fewer than minimum monitors report it healthy.
	MinimumMonitors pulumi.StringPtrInput `pulumi:"minimumMonitors"`
	// Port to be used for AWS service discovery,default `80`.
	Port pulumi.IntInput `pulumi:"port"`
	// If true, the server certificate is verified against the list of supplied/default CAs when making requests to the Consul API.
	RejectUnauthorized pulumi.BoolPtrInput `pulumi:"rejectUnauthorized"`
	// CA Bundle to validate server certificates.
	TrustCa pulumi.StringPtrInput `pulumi:"trustCa"`
	Type    pulumi.StringPtrInput `pulumi:"type"`
	// Action to take when node cannot be detected,default `remove`.
	UndetectableAction pulumi.StringPtrInput `pulumi:"undetectableAction"`
	// Update interval for service discovery.
	UpdateInterval pulumi.StringPtrInput `pulumi:"updateInterval"`
	// The location of the node data.
	Uri pulumi.StringInput `pulumi:"uri"`
}

func (GetConsulServiceDiscoveryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConsulServiceDiscoveryArgs)(nil)).Elem()
}

// A collection of values returned by getConsulServiceDiscovery.
type GetConsulServiceDiscoveryResultOutput struct{ *pulumi.OutputState }

func (GetConsulServiceDiscoveryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConsulServiceDiscoveryResult)(nil)).Elem()
}

func (o GetConsulServiceDiscoveryResultOutput) ToGetConsulServiceDiscoveryResultOutput() GetConsulServiceDiscoveryResultOutput {
	return o
}

func (o GetConsulServiceDiscoveryResultOutput) ToGetConsulServiceDiscoveryResultOutputWithContext(ctx context.Context) GetConsulServiceDiscoveryResultOutput {
	return o
}

func (o GetConsulServiceDiscoveryResultOutput) AddressRealm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.AddressRealm }).(pulumi.StringPtrOutput)
}

// The JSON for Hashicorp Consul service discovery block.
func (o GetConsulServiceDiscoveryResultOutput) ConsulSdJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) string { return v.ConsulSdJson }).(pulumi.StringOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) CredentialUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *bool { return v.CredentialUpdate }).(pulumi.BoolPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) EncodedToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.EncodedToken }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetConsulServiceDiscoveryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) JmesPathQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.JmesPathQuery }).(pulumi.StringPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) MinimumMonitors() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.MinimumMonitors }).(pulumi.StringPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) RejectUnauthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *bool { return v.RejectUnauthorized }).(pulumi.BoolPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) TrustCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.TrustCa }).(pulumi.StringPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) UndetectableAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.UndetectableAction }).(pulumi.StringPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) UpdateInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) *string { return v.UpdateInterval }).(pulumi.StringPtrOutput)
}

func (o GetConsulServiceDiscoveryResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GetConsulServiceDiscoveryResult) string { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetConsulServiceDiscoveryResultOutput{})
}
