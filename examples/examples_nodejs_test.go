// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestAccVirtualAppliance(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "virtualappliance"),
			// TODO[pulumi/pulumi-f5bigip#21]: Can we get this to a state where the empty preview and update actually
			// have no changes?
			AllowEmptyPreviewChanges: true,
			AllowEmptyUpdateChanges:  true,
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	region := getRegion(t)
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"aws:region":       region,
			"f5bigip:address":  "",
			"f5bigip:password": "",
			"f5bigip:username": "",
		},
		Dependencies: []string{
			"@pulumi/f5bigip",
		},
	})

	return baseJS
}
