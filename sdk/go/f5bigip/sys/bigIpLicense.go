// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sys

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type BigIpLicense struct {
	pulumi.CustomResourceState

	// Tmsh command to execute tmsh commands like install
	Command pulumi.StringOutput `pulumi:"command"`
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey pulumi.StringOutput `pulumi:"registrationKey"`
}

// NewBigIpLicense registers a new resource with the given unique name, arguments, and options.
func NewBigIpLicense(ctx *pulumi.Context,
	name string, args *BigIpLicenseArgs, opts ...pulumi.ResourceOption) (*BigIpLicense, error) {
	if args == nil || args.Command == nil {
		return nil, errors.New("missing required argument 'Command'")
	}
	if args == nil || args.RegistrationKey == nil {
		return nil, errors.New("missing required argument 'RegistrationKey'")
	}
	if args == nil {
		args = &BigIpLicenseArgs{}
	}
	var resource BigIpLicense
	err := ctx.RegisterResource("f5bigip:sys/bigIpLicense:BigIpLicense", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigIpLicense gets an existing BigIpLicense resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigIpLicense(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigIpLicenseState, opts ...pulumi.ResourceOption) (*BigIpLicense, error) {
	var resource BigIpLicense
	err := ctx.ReadResource("f5bigip:sys/bigIpLicense:BigIpLicense", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigIpLicense resources.
type bigIpLicenseState struct {
	// Tmsh command to execute tmsh commands like install
	Command *string `pulumi:"command"`
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey *string `pulumi:"registrationKey"`
}

type BigIpLicenseState struct {
	// Tmsh command to execute tmsh commands like install
	Command pulumi.StringPtrInput
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey pulumi.StringPtrInput
}

func (BigIpLicenseState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigIpLicenseState)(nil)).Elem()
}

type bigIpLicenseArgs struct {
	// Tmsh command to execute tmsh commands like install
	Command string `pulumi:"command"`
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey string `pulumi:"registrationKey"`
}

// The set of arguments for constructing a BigIpLicense resource.
type BigIpLicenseArgs struct {
	// Tmsh command to execute tmsh commands like install
	Command pulumi.StringInput
	// A unique Key F5 provides for Licensing BIG-IP
	RegistrationKey pulumi.StringInput
}

func (BigIpLicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigIpLicenseArgs)(nil)).Elem()
}
