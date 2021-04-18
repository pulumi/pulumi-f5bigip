// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIrule(ctx *pulumi.Context, args *GetIruleArgs, opts ...pulumi.InvokeOption) (*GetIruleResult, error) {
	var rv GetIruleResult
	err := ctx.Invoke("f5bigip:ltm/getIrule:getIrule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIrule.
type GetIruleArgs struct {
	Name      string `pulumi:"name"`
	Partition string `pulumi:"partition"`
}

// A collection of values returned by getIrule.
type GetIruleResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	Irule     string `pulumi:"irule"`
	Name      string `pulumi:"name"`
	Partition string `pulumi:"partition"`
}
