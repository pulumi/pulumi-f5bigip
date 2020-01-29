// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package f5bigip

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The provider type for the bigip package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-bigip/blob/master/website/docs/index.html.markdown.
type Provider struct {
	pulumi.ProviderResourceState

}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil || args.Address == nil {
		return nil, errors.New("missing required argument 'Address'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &ProviderArgs{}
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:f5bigip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Domain name/IP of the BigIP
	Address string `pulumi:"address"`
	// Login reference for token authentication (see BIG-IP REST docs for details)
	LoginRef *string `pulumi:"loginRef"`
	// The user's password
	Password string `pulumi:"password"`
	// Management Port to connect to Bigip
	Port *string `pulumi:"port"`
	// Enable to use an external authentication source (LDAP, TACACS, etc)
	TokenAuth *bool `pulumi:"tokenAuth"`
	// Username with API access to the BigIP
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Domain name/IP of the BigIP
	Address pulumi.StringInput
	// Login reference for token authentication (see BIG-IP REST docs for details)
	LoginRef pulumi.StringPtrInput
	// The user's password
	Password pulumi.StringInput
	// Management Port to connect to Bigip
	Port pulumi.StringPtrInput
	// Enable to use an external authentication source (LDAP, TACACS, etc)
	TokenAuth pulumi.BoolPtrInput
	// Username with API access to the BigIP
	Username pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

