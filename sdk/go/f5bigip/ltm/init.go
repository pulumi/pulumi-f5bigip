// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ltm

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-f5bigip/sdk/v3/go/f5bigip/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "f5bigip:ltm/cipherGroup:CipherGroup":
		r = &CipherGroup{}
	case "f5bigip:ltm/cipherRule:CipherRule":
		r = &CipherRule{}
	case "f5bigip:ltm/dataGroup:DataGroup":
		r = &DataGroup{}
	case "f5bigip:ltm/iRule:IRule":
		r = &IRule{}
	case "f5bigip:ltm/monitor:Monitor":
		r = &Monitor{}
	case "f5bigip:ltm/node:Node":
		r = &Node{}
	case "f5bigip:ltm/persistenceProfileCookie:PersistenceProfileCookie":
		r = &PersistenceProfileCookie{}
	case "f5bigip:ltm/persistenceProfileDstAddr:PersistenceProfileDstAddr":
		r = &PersistenceProfileDstAddr{}
	case "f5bigip:ltm/persistenceProfileSrcAddr:PersistenceProfileSrcAddr":
		r = &PersistenceProfileSrcAddr{}
	case "f5bigip:ltm/persistenceProfileSsl:PersistenceProfileSsl":
		r = &PersistenceProfileSsl{}
	case "f5bigip:ltm/policy:Policy":
		r = &Policy{}
	case "f5bigip:ltm/pool:Pool":
		r = &Pool{}
	case "f5bigip:ltm/poolAttachment:PoolAttachment":
		r = &PoolAttachment{}
	case "f5bigip:ltm/profileClientSsl:ProfileClientSsl":
		r = &ProfileClientSsl{}
	case "f5bigip:ltm/profileFastHttp:ProfileFastHttp":
		r = &ProfileFastHttp{}
	case "f5bigip:ltm/profileFastL4:ProfileFastL4":
		r = &ProfileFastL4{}
	case "f5bigip:ltm/profileFtp:ProfileFtp":
		r = &ProfileFtp{}
	case "f5bigip:ltm/profileHttp2:ProfileHttp2":
		r = &ProfileHttp2{}
	case "f5bigip:ltm/profileHttp:ProfileHttp":
		r = &ProfileHttp{}
	case "f5bigip:ltm/profileHttpCompress:ProfileHttpCompress":
		r = &ProfileHttpCompress{}
	case "f5bigip:ltm/profileOneConnect:ProfileOneConnect":
		r = &ProfileOneConnect{}
	case "f5bigip:ltm/profileServerSsl:ProfileServerSsl":
		r = &ProfileServerSsl{}
	case "f5bigip:ltm/profileTcp:ProfileTcp":
		r = &ProfileTcp{}
	case "f5bigip:ltm/profileWebAcceleration:ProfileWebAcceleration":
		r = &ProfileWebAcceleration{}
	case "f5bigip:ltm/snat:Snat":
		r = &Snat{}
	case "f5bigip:ltm/snatPool:SnatPool":
		r = &SnatPool{}
	case "f5bigip:ltm/virtualAddress:VirtualAddress":
		r = &VirtualAddress{}
	case "f5bigip:ltm/virtualServer:VirtualServer":
		r = &VirtualServer{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/cipherGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/cipherRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/dataGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/iRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/monitor",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/node",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/persistenceProfileCookie",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/persistenceProfileDstAddr",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/persistenceProfileSrcAddr",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/persistenceProfileSsl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/policy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/pool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/poolAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileClientSsl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileFastHttp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileFastL4",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileFtp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileHttp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileHttp2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileHttpCompress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileOneConnect",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileServerSsl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileTcp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/profileWebAcceleration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/snat",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/snatPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/virtualAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"f5bigip",
		"ltm/virtualServer",
		&module{version},
	)
}
