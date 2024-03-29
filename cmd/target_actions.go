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

var target_Cmd = &cobra.Command{
	Use:   "target",
	Short: "Manage targets",
	Long:  ``,
}

// target_createCmd represents the create command
var target_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a target",
	Long:  ``,
}

// target_updateCmd represents the edit command
var target_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a target",
	Long:  ``,
}

var target_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a target",
	Long:  ``,
}

func init() {

	//target_Cmd.AddCommand(Target_List())
	target_Cmd.AddCommand(target_createCmd)
	target_Cmd.AddCommand(target_updateCmd)
	target_Cmd.AddCommand(target_deleteCmd)
	rootCmd.AddCommand(target_Cmd)
}
