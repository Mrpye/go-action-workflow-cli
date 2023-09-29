package cmd

import (
	"errors"
	"reflect"
	"strings"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
func Target_List() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list [env]",
		Short: "list targets in an environment",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

// createCmd represents the create command for a target
func Build_CreateTargetCommand(targets map[string]*workflow.TargetSchema) {

	var target_name string //store the target name

	//******************************************************
	//Loop through the targets and create a command for each
	//******************************************************
	for key := range targets {

		target_obj := targets[key]

		var cmd = &cobra.Command{
			Use:   key + " [environment]",
			Short: target_obj.Short,
			Long:  target_obj.Long,
			RunE: func(cmd *cobra.Command, args []string) error {
				//************************
				//Check for an environment
				//************************
				if len(args) != 1 {
					return errors.New("must provide an environment")
				}

				//**********************
				//Create the action name
				//**********************
				parts := strings.Split(cmd.Use, " ")
				target_obj = wf_client.TargetScheme[parts[0]]

				//**********************
				//Create the target name
				//**********************
				target_n := strings.ReplaceAll(strings.ToLower(reflect.TypeOf(target_obj.Target).String()), "*", "")
				if target_name == "" {
					target_name = target_n
				} else {
					target_name = target_n + "_" + target_name
				}

				//***********************************
				//Write target_type and target_action
				//***********************************
				wf_client.WriteConfig(target_n+".target_type", target_n, args[0])
				//wf_client.WriteConfig(target_n+".target_action", target_obj.Action, args[0])
				SaveTargetToConfig(wf_client, args[0], target_n, target_obj.Target)

				return nil
			},
		}
		//********************************************************
		//Loop through the target schema and create flags for each
		//********************************************************
		//BuildFlags(cmd, target_obj.Schema, "")
		gpflag.ParseTo(target_obj.Target, cmd.Flags())

		//***************************************
		//Add the target_name flag to the command
		//***************************************
		cmd.Flags().StringVarP(&target_name, "target_name", "", "", "Set the target name")

		//***************************************
		//Add the command to the target_createCmd
		//***************************************
		target_createCmd.AddCommand(cmd)

	}
}

// updateCmd represents the update command for targets
func Build_UpdateTargetCommand(targets map[string]*workflow.TargetSchema) {

	var target_name string //store the target name

	//******************************************************
	//Loop through the targets and create a command for each
	//******************************************************
	for key := range targets {
		target_obj := targets[key]
		var cmd = &cobra.Command{
			Use:   target_obj.Action + " [environment]",
			Short: target_obj.Short,
			Long:  target_obj.Long,
			RunE: func(cmd *cobra.Command, args []string) error {
				//************************
				//Check for an environment
				//************************
				if len(args) != 1 {
					return errors.New("must provide an environment")
				}

				//**********************
				//Create the action name
				//**********************
				parts := strings.Split(cmd.Use, " ")
				target_obj = wf_client.TargetScheme[parts[0]]

				//**********************
				//Create the target name
				//**********************
				target_n := strings.ReplaceAll(strings.ToLower(reflect.TypeOf(target_obj.Target).String()), "*", "")
				if target_name == "" {
					target_name = target_n
				} else {
					target_name = target_n + "_" + target_name
				}
				//************************
				//See if the target exists
				//************************
				value, err := wf_client.ReadConfig(target_n+".target_type", "", args[0])
				if err != nil {
					return err
				}
				if value == nil {
					return errors.New("target does not exist")
				}

				//***********************************
				//Write target_type and target_action
				//***********************************
				wf_client.WriteConfig(target_n+".target_type", target_n, args[0])
				//wf_client.WriteConfig(target_n+".target_action", target_obj.Action, args[0])
				SaveTargetToConfig(wf_client, args[0], target_n, target_obj.Target)

				return err
			},
		}

		//********************************************************
		//Loop through the target schema and create flags for each
		//********************************************************
		gpflag.ParseTo(target_obj.Target, cmd.Flags())

		//***************************************
		//Add the target_name flag to the command
		//***************************************
		cmd.Flags().StringVarP(&target_name, "target_name", "", "", "Set the target name")

		//***************************************
		//Add the command to the target_createCmd
		//***************************************
		target_updateCmd.AddCommand(cmd)
	}
}

// deleteCmd represents the delete command for targets
func Build_DeleteTargetCommand(targets map[string]*workflow.TargetSchema) {

	//******************************************************
	//Loop through the targets and create a command for each
	//******************************************************
	for _, target := range targets {

		var cmd = &cobra.Command{
			Use:   target.Action + " [environment]",
			Short: target.Short,
			Long:  target.Long,
			RunE: func(cmd *cobra.Command, args []string) error {
				//************************
				//Check for an environment
				//************************
				if len(args) != 1 {
					return errors.New("must provide an environment")
				}
				//*****************************************
				///Delete the target type to the config file
				//*****************************************
				return wf_client.DeleteConfig(target.Action, args[0])
			},
		}
		//***************************************
		//Add the command to the target_createCmd
		//***************************************
		target_deleteCmd.AddCommand(cmd)
	}
}
