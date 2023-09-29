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
	"strings"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/Mrpye/golib/convert"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func Manifest_Job_Create() *cobra.Command {
	job := workflow.CreateJob()
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a job",
		Long:  "Create a job",
		RunE: func(cmd *cobra.Command, args []string) error {
			//*****************
			//Check for job key
			//*****************
			if job.Key == "" {
				return errors.New("must specify a job key -k --job-key")
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
			err = wf_client.Workflow.ManifestAddJob(job)
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
	err := gpflag.ParseTo(job, cmd.Flags())
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return cmd

}

func Manifest_Job_Delete() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete [Key]",
		Short: "Create a job",
		Long:  "Create a job",
		RunE: func(cmd *cobra.Command, args []string) error {
			//*****************
			//Check for job key
			//*****************
			if len(args) != 1 {
				return errors.New("must specify a job key")
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
			wf_client.Workflow.ManifestDeleteJob(args[0])
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

func Manifest_Job_List() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "This command will list the jobs in the manifest",
		Long:  "This command will list the jobs in the manifest",
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
			table.SetHeader([]string{"Key ", "Title", "Description", "Sub Workflow"})

			for _, v := range wf_client.Workflow.Manifest.Jobs {
				row := []string{v.Key, v.Title, v.Description, convert.ToString(v.IsSubWorkflow)}
				table.Append(row)
			}
			table.Render() // Send output

			return nil
		},
	}

	return cmd
}

func Manifest_Job_Run() *cobra.Command {
	var params []string
	var log int64
	var answer_file string
	var cmd = &cobra.Command{
		Use:   "run [Key] [env]",
		Short: "Create a job",
		Long:  "Create a job",
		RunE: func(cmd *cobra.Command, args []string) error {

			//*****************
			//Check for job key
			//*****************
			if len(args) != 2 {
				return errors.New("must specify a job key and a environment")
			}

			//*****************
			//Set the log level
			//*****************
			wf_client.Workflow.LogLevel = log

			//*****************
			//Load the manifest
			//*****************
			err := wf_client.Workflow.LoadManifest(manifest_file)
			if err != nil {
				return err
			}

			//*****************************
			//See if we have an answer file
			//*****************************
			if answer_file != "" {
				err = wf_client.Workflow.LoadAnswerFile(answer_file)
				if err != nil {
					return err
				}
			}

			//*****************
			//Parse the params
			//*****************
			params_map := make(map[string]interface{})
			for _, v := range params {
				parts := strings.Split(v, "=")
				if len(parts) == 2 {
					params_map[parts[0]] = parts[1]
				}
			}

			//***************************
			//Add the job to the manifest
			//***************************
			err = wf_client.Workflow.RunJob(args[0], map[string]interface{}{"env": args[1]}, params_map)
			if err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&answer_file, "answer-file", "", "", "The filename for the answer file")
	cmd.Flags().StringSliceVarP(&params, "set", "s", []string{}, "")
	cmd.Flags().Int64VarP(&log, "log-level", "l", 0, "Log level 0 = quite, 1 = info, 2 = verbose")
	return cmd

}

func init() {
	job_Cmd.AddCommand(Manifest_Job_Run())
	job_Cmd.AddCommand(Manifest_Job_Create())
	job_Cmd.AddCommand(Manifest_Job_Delete())
	job_Cmd.AddCommand(Manifest_Job_List())
}
