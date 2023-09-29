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
	"errors"
	"log"
	"os"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/Mrpye/golib/convert"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func Manifest_Create_Parameter() *cobra.Command {
	var value string
	param := workflow.CreateParameter()
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a parameter",
		Long:  "Create a parameter",
		RunE: func(cmd *cobra.Command, args []string) error {

			//********************
			//Check for key is set
			//********************
			if param.Key == "" {
				return errors.New("must specify a parameter key -k --key")
			}

			//*****************
			//Load the manifest
			//*****************
			err := wf_client.Workflow.LoadManifest(manifest_file)
			if err != nil {
				return err
			}

			//*************************************
			//Convert the value to the correct type
			//*************************************
			switch param.InputType {
			case "string":
				param.Value = convert.ToString(value)
			case "int":
				param.Value = convert.ToInt(value)
			case "float":
				param.Value = convert.ToFloat64(value)
			case "bool":
				param.Value = convert.ToBool(value)
			default:
				param.Value = convert.ToString(value)
			}
			//***************************
			//Add the job to the manifest
			//***************************
			err = wf_client.Workflow.ManifestCreateParameter(param)
			if err != nil {
				return err
			}

			//*****************
			//Save the manifest
			//*****************
			err = wf_client.Workflow.SaveManifest(manifest_file)
			if err != nil {
				return err
			}

			return nil
		},
	}
	//***************
	//Parse the flags
	//***************
	cmd.Flags().StringVarP(&value, "value", "v", "", "The value for the parameter")
	err := gpflag.ParseTo(param, cmd.Flags())
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return cmd

}

func Manifest_Delete_Parameter() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete [Key]",
		Short: "Delete a Parameter",
		Long:  "Delete a Parameter",
		RunE: func(cmd *cobra.Command, args []string) error {
			//*****************
			//Check for job key
			//*****************
			if len(args) != 1 {
				return errors.New("must specify a parameter key")
			}

			//*****************
			//Load the manifest
			//*****************
			err := wf_client.Workflow.LoadManifest(manifest_file)
			if err != nil {
				return err
			}
			//***************************
			//Add the job to the manifest
			//***************************
			err = wf_client.Workflow.ManifestDeleteParameter(args[0])
			if err != nil {
				return err
			}

			//*****************
			//Save the manifest
			//*****************
			err = wf_client.Workflow.SaveManifest(manifest_file)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd

}

func Manifest_Parameter_List() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "This command will list the parameters in the manifest",
		Long:  "This command will list the parameters in the manifest",
		RunE: func(cmd *cobra.Command, args []string) error {

			//*****************
			//Load the manifest
			//*****************
			err := wf_client.Workflow.LoadManifest(manifest_file)
			if err != nil {
				return err
			}

			//*****************
			//Print the results
			//*****************
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Key ", "Title", "Description", "Value"})

			for _, v := range wf_client.Workflow.Manifest.Parameters {
				row := []string{v.Key, v.Title, v.Description, convert.ToString(v.Value)}
				table.Append(row)
			}
			table.Render() // Send output

			return nil
		},
	}

	return cmd
}

func init() {
	parameter_Cmd.AddCommand(Manifest_Create_Parameter())
	parameter_Cmd.AddCommand(Manifest_Delete_Parameter())
	parameter_Cmd.AddCommand(Manifest_Parameter_List())
}
