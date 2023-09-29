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
	"log"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/spf13/cobra"
)

func Build_Create_Meta() *cobra.Command {
	meta := workflow.CreateMetaData()
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a meta data",
		Long:  "Create a meta data",
		RunE: func(cmd *cobra.Command, args []string) error {

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
			wf_client.Workflow.ManifestCreateMeta(meta)
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
	err := gpflag.ParseTo(meta, cmd.Flags())
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	return cmd

}

func Build_Update_Meta() *cobra.Command {
	meta := workflow.CreateMetaData()
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "Update a meta data",
		Long:  "Update a meta data",
		RunE: func(cmd *cobra.Command, args []string) error {

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
			wf_client.Workflow.ManifestUpdateMeta(meta)
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
	err := gpflag.ParseTo(meta, cmd.Flags())
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	return cmd

}

func init() {
	meta_Cmd.AddCommand(Build_Create_Meta())
	meta_Cmd.AddCommand(Build_Update_Meta())

}
