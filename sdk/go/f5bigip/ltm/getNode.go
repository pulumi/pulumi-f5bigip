// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source (`ltm.Node`) to get the ltm node details available on BIG-IP
func LookupNode(ctx *pulumi.Context, args *LookupNodeArgs, opts ...pulumi.InvokeOption) (*LookupNodeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNodeResult
	err := ctx.Invoke("f5bigip:ltm/getNode:getNode", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNode.
type LookupNodeArgs struct {
	// The address of the node.
	Address *string `pulumi:"address"`
	// User defined description of the node.
	Description *string      `pulumi:"description"`
	Fqdn        *GetNodeFqdn `pulumi:"fqdn"`
	// Full path of the node (partition and name)
	FullPath *string `pulumi:"fullPath"`
	// Name of the node.
	Name string `pulumi:"name"`
	// partition of the node.
	Partition string `pulumi:"partition"`
}

// A collection of values returned by getNode.
type LookupNodeResult struct {
	// The address of the node.
	Address *string `pulumi:"address"`
	// Node connection limit.
	ConnectionLimit int `pulumi:"connectionLimit"`
	// User defined description of the node.
	Description *string `pulumi:"description"`
	// The dynamic ratio number for the node.
	DynamicRatio int         `pulumi:"dynamicRatio"`
	Fqdn         GetNodeFqdn `pulumi:"fqdn"`
	// Full path of the node (partition and name)
	FullPath *string `pulumi:"fullPath"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specifies the health monitors the system currently uses to monitor this node.
	Monitor   string `pulumi:"monitor"`
	Name      string `pulumi:"name"`
	Partition string `pulumi:"partition"`
	// Node rate limit.
	RateLimit string `pulumi:"rateLimit"`
	// Node ratio weight.
	Ratio   int    `pulumi:"ratio"`
	Session string `pulumi:"session"`
	// The current state of the node.
	State string `pulumi:"state"`
}

func LookupNodeOutput(ctx *pulumi.Context, args LookupNodeOutputArgs, opts ...pulumi.InvokeOption) LookupNodeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNodeResultOutput, error) {
			args := v.(LookupNodeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("f5bigip:ltm/getNode:getNode", args, LookupNodeResultOutput{}, options).(LookupNodeResultOutput), nil
		}).(LookupNodeResultOutput)
}

// A collection of arguments for invoking getNode.
type LookupNodeOutputArgs struct {
	// The address of the node.
	Address pulumi.StringPtrInput `pulumi:"address"`
	// User defined description of the node.
	Description pulumi.StringPtrInput `pulumi:"description"`
	Fqdn        GetNodeFqdnPtrInput   `pulumi:"fqdn"`
	// Full path of the node (partition and name)
	FullPath pulumi.StringPtrInput `pulumi:"fullPath"`
	// Name of the node.
	Name pulumi.StringInput `pulumi:"name"`
	// partition of the node.
	Partition pulumi.StringInput `pulumi:"partition"`
}

func (LookupNodeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeArgs)(nil)).Elem()
}

// A collection of values returned by getNode.
type LookupNodeResultOutput struct{ *pulumi.OutputState }

func (LookupNodeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeResult)(nil)).Elem()
}

func (o LookupNodeResultOutput) ToLookupNodeResultOutput() LookupNodeResultOutput {
	return o
}

func (o LookupNodeResultOutput) ToLookupNodeResultOutputWithContext(ctx context.Context) LookupNodeResultOutput {
	return o
}

// The address of the node.
func (o LookupNodeResultOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeResult) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// Node connection limit.
func (o LookupNodeResultOutput) ConnectionLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeResult) int { return v.ConnectionLimit }).(pulumi.IntOutput)
}

// User defined description of the node.
func (o LookupNodeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The dynamic ratio number for the node.
func (o LookupNodeResultOutput) DynamicRatio() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeResult) int { return v.DynamicRatio }).(pulumi.IntOutput)
}

func (o LookupNodeResultOutput) Fqdn() GetNodeFqdnOutput {
	return o.ApplyT(func(v LookupNodeResult) GetNodeFqdn { return v.Fqdn }).(GetNodeFqdnOutput)
}

// Full path of the node (partition and name)
func (o LookupNodeResultOutput) FullPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeResult) *string { return v.FullPath }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNodeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the health monitors the system currently uses to monitor this node.
func (o LookupNodeResultOutput) Monitor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeResult) string { return v.Monitor }).(pulumi.StringOutput)
}

func (o LookupNodeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNodeResultOutput) Partition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeResult) string { return v.Partition }).(pulumi.StringOutput)
}

// Node rate limit.
func (o LookupNodeResultOutput) RateLimit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeResult) string { return v.RateLimit }).(pulumi.StringOutput)
}

// Node ratio weight.
func (o LookupNodeResultOutput) Ratio() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeResult) int { return v.Ratio }).(pulumi.IntOutput)
}

func (o LookupNodeResultOutput) Session() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeResult) string { return v.Session }).(pulumi.StringOutput)
}

// The current state of the node.
func (o LookupNodeResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodeResultOutput{})
}
