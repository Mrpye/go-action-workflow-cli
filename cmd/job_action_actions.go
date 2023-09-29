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

var action_Cmd = &cobra.Command{
	Use:   "action",
	Short: "Manage Actions",
	Long:  ``,
}

// action_createCmd represents the create command
var action_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a action",
	Long:  ``,
}

/*
// action_updateCmd represents the edit command

	var action_updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a action",
		Long:  ``,
	}
*/

/*var action_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a action",
	Long:  ``,
}*/

func init() {
	action_Cmd.AddCommand(action_createCmd)
	//action_Cmd.AddCommand(action_updateCmd)
	//action_Cmd.AddCommand(action_deleteCmd)
	manifest_Cmd.AddCommand(action_Cmd)
}
