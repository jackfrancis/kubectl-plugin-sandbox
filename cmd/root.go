// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/cluster-api/api/v1alpha3"
)

// NewRootCmd returns the root command for AKS Engine.
func NewRootCmd() *cobra.Command {
	var version string
	rootCmd := &cobra.Command{
		Use:   "kubectl-plugin-sandbox",
		Short: "kubectl plugin sandbox",
		Long:  "kubectl plugin sandbox",
		RunE: func(cmd *cobra.Command, args []string) error {
			var groupVersion schema.GroupVersion
			switch version {
			case "v1alpha2":
				groupVersion = v1alpha2.GroupVersion
			case "v1alpha3":
				groupVersion = v1alpha3.GroupVersion
			default:
				groupVersion = schema.GroupVersion{
					Version: "unknown",
				}
			}
			fmt.Println(groupVersion.Version)
			return nil
		},
	}
	rootCmd.Flags().StringVarP(&version, "version", "v", "", "which version to print")
	return rootCmd
}
