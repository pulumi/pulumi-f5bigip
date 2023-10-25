// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source (`fast.getGceServiceDiscovery`) to get the GCE Service discovery config to be used for `http`/`https` app deployment in FAST.
func GetGceServiceDiscovery(ctx *pulumi.Context, args *GetGceServiceDiscoveryArgs, opts ...pulumi.InvokeOption) (*GetGceServiceDiscoveryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGceServiceDiscoveryResult
	err := ctx.Invoke("f5bigip:fast/getGceServiceDiscovery:getGceServiceDiscovery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGceServiceDiscovery.
type GetGceServiceDiscoveryArgs struct {
	// Specifies whether to look for public or private IP addresses,default `private`.
	AddressRealm *string `pulumi:"addressRealm"`
	// Specifies whether you are updating your credentials,default `false`.
	CredentialUpdate *bool `pulumi:"credentialUpdate"`
	// Base 64 encoded service account credentials JSON.
	EncodedCredentials *string `pulumi:"encodedCredentials"`
	// Member is down when fewer than minimum monitors report it healthy.
	MinimumMonitors *string `pulumi:"minimumMonitors"`
	// Port to be used for AWS service discovery,default `80`.
	Port *int `pulumi:"port"`
	// For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
	ProjectId *string `pulumi:"projectId"`
	// GCE region in which ADC is running.
	Region string `pulumi:"region"`
	// The tag key associated with the node to add to this pool.
	TagKey string `pulumi:"tagKey"`
	// The tag value associated with the node to add to this pool.
	TagValue string  `pulumi:"tagValue"`
	Type     *string `pulumi:"type"`
	// Action to take when node cannot be detected,default `remove`.
	UndetectableAction *string `pulumi:"undetectableAction"`
	// Update interval for service discovery.
	UpdateInterval *string `pulumi:"updateInterval"`
}

// A collection of values returned by getGceServiceDiscovery.
type GetGceServiceDiscoveryResult struct {
	AddressRealm       *string `pulumi:"addressRealm"`
	CredentialUpdate   *bool   `pulumi:"credentialUpdate"`
	EncodedCredentials *string `pulumi:"encodedCredentials"`
	// The JSON for GCE service discovery block.
	GceSdJson string `pulumi:"gceSdJson"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	MinimumMonitors    *string `pulumi:"minimumMonitors"`
	Port               *int    `pulumi:"port"`
	ProjectId          *string `pulumi:"projectId"`
	Region             string  `pulumi:"region"`
	TagKey             string  `pulumi:"tagKey"`
	TagValue           string  `pulumi:"tagValue"`
	Type               *string `pulumi:"type"`
	UndetectableAction *string `pulumi:"undetectableAction"`
	UpdateInterval     *string `pulumi:"updateInterval"`
}

func GetGceServiceDiscoveryOutput(ctx *pulumi.Context, args GetGceServiceDiscoveryOutputArgs, opts ...pulumi.InvokeOption) GetGceServiceDiscoveryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGceServiceDiscoveryResult, error) {
			args := v.(GetGceServiceDiscoveryArgs)
			r, err := GetGceServiceDiscovery(ctx, &args, opts...)
			var s GetGceServiceDiscoveryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGceServiceDiscoveryResultOutput)
}

// A collection of arguments for invoking getGceServiceDiscovery.
type GetGceServiceDiscoveryOutputArgs struct {
	// Specifies whether to look for public or private IP addresses,default `private`.
	AddressRealm pulumi.StringPtrInput `pulumi:"addressRealm"`
	// Specifies whether you are updating your credentials,default `false`.
	CredentialUpdate pulumi.BoolPtrInput `pulumi:"credentialUpdate"`
	// Base 64 encoded service account credentials JSON.
	EncodedCredentials pulumi.StringPtrInput `pulumi:"encodedCredentials"`
	// Member is down when fewer than minimum monitors report it healthy.
	MinimumMonitors pulumi.StringPtrInput `pulumi:"minimumMonitors"`
	// Port to be used for AWS service discovery,default `80`.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// For Google Cloud Engine (GCE) only: The ID of the project in which the members are located.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// GCE region in which ADC is running.
	Region pulumi.StringInput `pulumi:"region"`
	// The tag key associated with the node to add to this pool.
	TagKey pulumi.StringInput `pulumi:"tagKey"`
	// The tag value associated with the node to add to this pool.
	TagValue pulumi.StringInput    `pulumi:"tagValue"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
	// Action to take when node cannot be detected,default `remove`.
	UndetectableAction pulumi.StringPtrInput `pulumi:"undetectableAction"`
	// Update interval for service discovery.
	UpdateInterval pulumi.StringPtrInput `pulumi:"updateInterval"`
}

func (GetGceServiceDiscoveryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGceServiceDiscoveryArgs)(nil)).Elem()
}

// A collection of values returned by getGceServiceDiscovery.
type GetGceServiceDiscoveryResultOutput struct{ *pulumi.OutputState }

func (GetGceServiceDiscoveryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGceServiceDiscoveryResult)(nil)).Elem()
}

func (o GetGceServiceDiscoveryResultOutput) ToGetGceServiceDiscoveryResultOutput() GetGceServiceDiscoveryResultOutput {
	return o
}

func (o GetGceServiceDiscoveryResultOutput) ToGetGceServiceDiscoveryResultOutputWithContext(ctx context.Context) GetGceServiceDiscoveryResultOutput {
	return o
}

func (o GetGceServiceDiscoveryResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetGceServiceDiscoveryResult] {
	return pulumix.Output[GetGceServiceDiscoveryResult]{
		OutputState: o.OutputState,
	}
}

func (o GetGceServiceDiscoveryResultOutput) AddressRealm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *string { return v.AddressRealm }).(pulumi.StringPtrOutput)
}

func (o GetGceServiceDiscoveryResultOutput) CredentialUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *bool { return v.CredentialUpdate }).(pulumi.BoolPtrOutput)
}

func (o GetGceServiceDiscoveryResultOutput) EncodedCredentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *string { return v.EncodedCredentials }).(pulumi.StringPtrOutput)
}

// The JSON for GCE service discovery block.
func (o GetGceServiceDiscoveryResultOutput) GceSdJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) string { return v.GceSdJson }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGceServiceDiscoveryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGceServiceDiscoveryResultOutput) MinimumMonitors() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *string { return v.MinimumMonitors }).(pulumi.StringPtrOutput)
}

func (o GetGceServiceDiscoveryResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o GetGceServiceDiscoveryResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o GetGceServiceDiscoveryResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetGceServiceDiscoveryResultOutput) TagKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) string { return v.TagKey }).(pulumi.StringOutput)
}

func (o GetGceServiceDiscoveryResultOutput) TagValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) string { return v.TagValue }).(pulumi.StringOutput)
}

func (o GetGceServiceDiscoveryResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetGceServiceDiscoveryResultOutput) UndetectableAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *string { return v.UndetectableAction }).(pulumi.StringPtrOutput)
}

func (o GetGceServiceDiscoveryResultOutput) UpdateInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGceServiceDiscoveryResult) *string { return v.UpdateInterval }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGceServiceDiscoveryResultOutput{})
}
