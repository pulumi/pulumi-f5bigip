// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`fast.getAwsServiceDiscovery`) to get the AWS Service discovery config to be used for `http`/`https` app deployment in FAST.
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
//			_, err := fast.GetAwsServiceDiscovery(ctx, &fast.GetAwsServiceDiscoveryArgs{
//				TagKey:   "testawstagkey",
//				TagValue: "testawstagvalue",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAwsServiceDiscovery(ctx *pulumi.Context, args *GetAwsServiceDiscoveryArgs, opts ...pulumi.InvokeOption) (*GetAwsServiceDiscoveryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAwsServiceDiscoveryResult
	err := ctx.Invoke("f5bigip:fast/getAwsServiceDiscovery:getAwsServiceDiscovery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAwsServiceDiscovery.
type GetAwsServiceDiscoveryArgs struct {
	// Specifies whether to look for public or private IP addresses,default `private`.
	AddressRealm *string `pulumi:"addressRealm"`
	// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
	AwsAccessKey *string `pulumi:"awsAccessKey"`
	// AWS region in which ADC is running,default Empty string.
	AwsRegion *string `pulumi:"awsRegion"`
	// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
	AwsSecretAccessKey *string `pulumi:"awsSecretAccessKey"`
	// Specifies whether you are updating your credentials,default `false`.
	CredentialUpdate *bool `pulumi:"credentialUpdate"`
	// AWS externalID field.
	ExternalId *string `pulumi:"externalId"`
	// Member is down when fewer than minimum monitors report it healthy.
	MinimumMonitors *string `pulumi:"minimumMonitors"`
	// Port to be used for AWS service discovery,default `80`.
	Port *int `pulumi:"port"`
	// Assume a role (also requires the `externalId` field)
	RoleArn *string `pulumi:"roleArn"`
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

// A collection of values returned by getAwsServiceDiscovery.
type GetAwsServiceDiscoveryResult struct {
	AddressRealm *string `pulumi:"addressRealm"`
	AwsAccessKey *string `pulumi:"awsAccessKey"`
	AwsRegion    string  `pulumi:"awsRegion"`
	// The JSON for AWS service discovery block.
	AwsSdJson          string  `pulumi:"awsSdJson"`
	AwsSecretAccessKey *string `pulumi:"awsSecretAccessKey"`
	CredentialUpdate   *bool   `pulumi:"credentialUpdate"`
	ExternalId         string  `pulumi:"externalId"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	MinimumMonitors    *string `pulumi:"minimumMonitors"`
	Port               *int    `pulumi:"port"`
	RoleArn            string  `pulumi:"roleArn"`
	TagKey             string  `pulumi:"tagKey"`
	TagValue           string  `pulumi:"tagValue"`
	Type               *string `pulumi:"type"`
	UndetectableAction *string `pulumi:"undetectableAction"`
	UpdateInterval     *string `pulumi:"updateInterval"`
}

func GetAwsServiceDiscoveryOutput(ctx *pulumi.Context, args GetAwsServiceDiscoveryOutputArgs, opts ...pulumi.InvokeOption) GetAwsServiceDiscoveryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAwsServiceDiscoveryResult, error) {
			args := v.(GetAwsServiceDiscoveryArgs)
			r, err := GetAwsServiceDiscovery(ctx, &args, opts...)
			var s GetAwsServiceDiscoveryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAwsServiceDiscoveryResultOutput)
}

// A collection of arguments for invoking getAwsServiceDiscovery.
type GetAwsServiceDiscoveryOutputArgs struct {
	// Specifies whether to look for public or private IP addresses,default `private`.
	AddressRealm pulumi.StringPtrInput `pulumi:"addressRealm"`
	// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
	AwsAccessKey pulumi.StringPtrInput `pulumi:"awsAccessKey"`
	// AWS region in which ADC is running,default Empty string.
	AwsRegion pulumi.StringPtrInput `pulumi:"awsRegion"`
	// Information for discovering AWS nodes that are not in the same region as your BIG-IP (also requires the `awsSecretAccessKey` field)
	AwsSecretAccessKey pulumi.StringPtrInput `pulumi:"awsSecretAccessKey"`
	// Specifies whether you are updating your credentials,default `false`.
	CredentialUpdate pulumi.BoolPtrInput `pulumi:"credentialUpdate"`
	// AWS externalID field.
	ExternalId pulumi.StringPtrInput `pulumi:"externalId"`
	// Member is down when fewer than minimum monitors report it healthy.
	MinimumMonitors pulumi.StringPtrInput `pulumi:"minimumMonitors"`
	// Port to be used for AWS service discovery,default `80`.
	Port pulumi.IntPtrInput `pulumi:"port"`
	// Assume a role (also requires the `externalId` field)
	RoleArn pulumi.StringPtrInput `pulumi:"roleArn"`
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

func (GetAwsServiceDiscoveryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAwsServiceDiscoveryArgs)(nil)).Elem()
}

// A collection of values returned by getAwsServiceDiscovery.
type GetAwsServiceDiscoveryResultOutput struct{ *pulumi.OutputState }

func (GetAwsServiceDiscoveryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAwsServiceDiscoveryResult)(nil)).Elem()
}

func (o GetAwsServiceDiscoveryResultOutput) ToGetAwsServiceDiscoveryResultOutput() GetAwsServiceDiscoveryResultOutput {
	return o
}

func (o GetAwsServiceDiscoveryResultOutput) ToGetAwsServiceDiscoveryResultOutputWithContext(ctx context.Context) GetAwsServiceDiscoveryResultOutput {
	return o
}

func (o GetAwsServiceDiscoveryResultOutput) AddressRealm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *string { return v.AddressRealm }).(pulumi.StringPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) AwsAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *string { return v.AwsAccessKey }).(pulumi.StringPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) AwsRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) string { return v.AwsRegion }).(pulumi.StringOutput)
}

// The JSON for AWS service discovery block.
func (o GetAwsServiceDiscoveryResultOutput) AwsSdJson() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) string { return v.AwsSdJson }).(pulumi.StringOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) AwsSecretAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *string { return v.AwsSecretAccessKey }).(pulumi.StringPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) CredentialUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *bool { return v.CredentialUpdate }).(pulumi.BoolPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) string { return v.ExternalId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAwsServiceDiscoveryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) MinimumMonitors() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *string { return v.MinimumMonitors }).(pulumi.StringPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) TagKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) string { return v.TagKey }).(pulumi.StringOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) TagValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) string { return v.TagValue }).(pulumi.StringOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) UndetectableAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *string { return v.UndetectableAction }).(pulumi.StringPtrOutput)
}

func (o GetAwsServiceDiscoveryResultOutput) UpdateInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAwsServiceDiscoveryResult) *string { return v.UpdateInterval }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAwsServiceDiscoveryResultOutput{})
}
