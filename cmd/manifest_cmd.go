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

	go_file "github.com/Mrpye/golib/file"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/spf13/cobra"
)

func CMD_Manifest_INIT() *cobra.Command {
	meta := workflow.CreateMetaData()
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create a manifest file",
		Long:  "Create a manifest file",
		RunE: func(cmd *cobra.Command, args []string) error {

			//*****************
			//Check if file exists
			//*****************
			if go_file.FileExists(manifest_file) {
				return errors.New("manifest file already exists")
			}

			//**********************************
			//Create the manifest with meta data
			//**********************************
			wf_client.Workflow.ManifestCreateMeta(meta)

			//*****************
			//Save the manifest
			//*****************
			err := wf_client.Workflow.SaveManifest(manifest_file)
			if err != nil {
				return err
			}
			return nil
		},
	}

	//*****************
	//Parse the flags
	//*****************
	err := gpflag.ParseTo(meta, cmd.Flags())
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	return cmd

}

func CMD_Manifest_Delete() *cobra.Command {
	var confirm bool
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a manifest file",
		Long:  "Delete a manifest file",
		RunE: func(cmd *cobra.Command, args []string) error {
			//*****************
			//Check if file exists
			//*****************
			if !go_file.FileExists(manifest_file) {
				return errors.New("manifest file does not exist")
			}

			//***********************
			//Check if user confirmed
			//***********************
			if !confirm {
				return errors.New("must confirm deletion")
			}

			//*****************
			//Delete the file
			//*****************
			err := os.Remove(manifest_file)
			if err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().BoolVarP(&confirm, "confirm", "c", false, "must")
	return cmd

}

func CMD_Manifest_BuildAnswer() *cobra.Command {
	var answer_file string
	var cmd = &cobra.Command{
		Use:   "answer",
		Short: "create answer file",
		Long:  "create answer file",
		RunE: func(cmd *cobra.Command, args []string) error {

			if manifest_file == "" {
				return errors.New("manifest file not specified")
			}

			//*****************
			//Load the manifest
			//*****************
			err := wf_client.Workflow.LoadManifest(manifest_file)
			if err != nil {
				return err
			}

			//*********************
			//Build the answer file
			//*********************
			answer, err := wf_client.Workflow.BuildAnswerFile()
			if err != nil {
				return err
			}

			//********************
			//Save the answer file
			//********************
			err = answer.Save(answer_file)
			if err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&answer_file, "answer-file", "", "", "The filename for the answer file")
	cmd.MarkFlagRequired("answer-file")
	cmd.MarkFlagFilename("answer-file", "yml", "yaml")
	return cmd

}

func init() {

	manifest_Cmd.AddCommand(CMD_Manifest_Delete())
	manifest_Cmd.AddCommand(CMD_Manifest_INIT())
	manifest_Cmd.AddCommand(CMD_Manifest_BuildAnswer())

}
