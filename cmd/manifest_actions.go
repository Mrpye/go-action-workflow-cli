/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var manifest_Cmd = &cobra.Command{
	Use:   "manifest",
	Short: "Manage the manifest",
	Long:  ``,
}

/*
// manifest_createCmd represents the create command

	var manifest_createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a manifest",
		Long:  ``,
	}

// manifest_updateCmd represents the edit command

	var manifest_updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a manifest",
		Long:  ``,
	}

	var manifest_deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a manifest",
		Long:  ``,
	}
*/
func init() {
	//manifest_Cmd.AddCommand(manifest_createCmd)
	//manifest_Cmd.AddCommand(manifest_updateCmd)
	//manifest_Cmd.AddCommand(manifest_deleteCmd)
	rootCmd.AddCommand(manifest_Cmd)

	//******************************************
	// Create a PersistentFlags for the manifest
	// And make it required
	//******************************************
	manifest_Cmd.PersistentFlags().StringVarP(&manifest_file, "file", "", "", "manifest file to use or document folder to use")
	manifest_Cmd.MarkPersistentFlagRequired("file")
	manifest_Cmd.MarkPersistentFlagFilename("file", "yml", "yaml")
	//********************************************
	// This flag will disable the required flags
	//So if you want to just get template and edit
	// directly in file, you can
	//********************************************
	manifest_Cmd.PersistentFlags().BoolVarP(&disable_required, "disable_required", "", false, "disable the required flag for actions")
}
