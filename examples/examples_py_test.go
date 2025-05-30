// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build python || all
// +build python all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	return getBaseOptions(t).With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})
}
