// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package main

import (
	"os"

	"github.com/jackfrancis/kubectl-plugin-sandbox/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
