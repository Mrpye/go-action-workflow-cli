package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/Mrpye/golib/convert"
	"github.com/octago/sflags/gen/gpflag"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// createCmd represents the create command
func Manifest_Action_List() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list [Job Key]",
		Short: "This command will list actions in a job",
		Long:  "This command will list actions in a job",
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
			job := wf_client.Workflow.ManifestGetJob(args[0])
			if job == nil {
				return errors.New("job not found")
			}
			//*****************
			//Print the results
			//*****************
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Key ", "Action", "Description", "Fail Action", "ContinueOnError", "Disabled", "Config"})

			for _, v := range job.Actions {
				conf, _ := json.Marshal(v.Config)
				row := []string{v.Key, v.Action, v.Description, v.Fail, convert.ToString(v.ContinueOnError), convert.ToString(v.Disabled), string(conf)}
				table.Append(row)
			}
			table.Render() // Send output

			return nil
		},
	}

	return cmd
}

// createCmd represents the create command for a action
func Build_CreateActionCommand(actions map[string]*workflow.ActionSchema) {

	//******************************************************
	//Loop through the actions and create a command for each
	//******************************************************
	for key, action := range actions {

		//************************
		//Create the action object
		//and target_name variable
		//************************
		action_obj := workflow.CreateAction()
		//var target_name string //store the action name

		var cmd = &cobra.Command{
			Use:   key + " [[job_key add to job][] global action]",
			Short: action.Short,
			Long:  action.Long,
			PreRun: func(cmd *cobra.Command, args []string) {
				if disable_required {
					cmd.Flags().VisitAll(func(f *pflag.Flag) {
						if f.Name != "help" && f.Name != "file" {
							cmd.Flags().SetAnnotation(f.Name, cobra.BashCompOneRequiredFlag, []string{"false"})
						}
					})
				}

			},
			RunE: func(cmd *cobra.Command, args []string) error {
				//************************
				//Check for an environment
				//************************

				if len(args) == 0 {
					value, err := cmd.Flags().GetString("key")
					if err != nil {
						return err
					}
					if value == "" {
						return errors.New("this will be added as a global action but you must specify an action key using -k --key")
					}
				}
				/*if manifest_file == "" {
					return errors.New("must specify a manifest file using -f --file")
				}*/

				//**********************
				//Create the action name
				//**********************
				parts := strings.Split(cmd.Use, " ")
				action_n := parts[0]
				//*********************
				//Get the action object
				//*********************
				action = wf_client.ActionScheme[action_n]

				//************************
				//Create the config object
				//************************
				tmp_config := make(map[string]interface{})
				if action.ProcessResults {
					tmp_config["result_format"] = "none"
					tmp_config["result_action"] = "js"
					tmp_config["result_js"] = "function ActionResults(store, result) {\nreturn true\n}"
				}

				//****************
				//Write the config
				//****************
				var err error
				cmd.Flags().VisitAll(func(f *pflag.Flag) {
					if f.Name != "help" && f.Name != "file" && strings.HasPrefix(f.Name, "cfg-") {
						cnf_name := strings.TrimPrefix(f.Name, "cfg-")
						if f.Value.Type() == "stringArray" {
							new_val := strings.ReplaceAll(f.Value.String(), "[", "")
							new_val = strings.ReplaceAll(new_val, "]", "")
							if strings.Contains(new_val, "=") {
								tmpmap := make(map[string]string)
								for _, v := range strings.Split(new_val, ",") {
									parts := strings.Split(v, "=")
									tmpmap[parts[0]] = parts[1]
								}
								tmp_config[cnf_name] = tmpmap
							} else {
								tmp_config[cnf_name] = strings.Split(new_val, ",")
							}
						} else {
							tmp_config[cnf_name] = f.Value
						}
					}
				})

				//*******************
				//Set the action name
				//and config object
				//*******************
				if action.InlineParams {
					action_obj.Action = action.InlineFormat(tmp_config)
				} else {
					action_obj.Action = action_n
					action_obj.Config = tmp_config
				}
				//*****************
				//load the manifest
				//*****************
				err = wf_client.Workflow.LoadManifest(manifest_file)
				if err != nil {
					return err
				}
				//*************************
				//Add the action to the job
				//*************************
				if len(args) == 0 {
					err = wf_client.Workflow.ManifestAddGlobalAction(action_obj)
					if err != nil {
						return err
					}
				} else {
					err = wf_client.Workflow.ManifestAddActionToJob(args[0], action_obj)
					if err != nil {
						return err
					}
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

		//********************************************************
		//Loop through the target schema and create flags for each
		//********************************************************
		BuildFlags(cmd, action.ConfigSchema, "cfg-")

		gpflag.ParseTo(action_obj, cmd.Flags())

		//***************************************
		//Add the command to the action_createCmd
		//***************************************
		action_createCmd.AddCommand(cmd)

	}
}

// createCmd represents the create command
func Manifest_Action_Delete() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete [Job Key] [action key]",
		Short: "This command will delete actions in a job",
		Long:  "This command will delete actions in a job",
		RunE: func(cmd *cobra.Command, args []string) error {

			//*****************
			//Check for job key
			//*****************
			if len(args) != 2 {
				return errors.New("must specify a job key and action key")
			}

			//*****************
			//Load the manifest
			//*****************
			err := wf_client.Workflow.LoadManifest(manifest_file)
			if err != nil {
				return err
			}

			//*****************
			//Delete the action
			//*****************
			err = wf_client.Workflow.ManifestDeleteActionFromJob(args[0], args[1])
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

func init() {
	action_Cmd.AddCommand(Manifest_Action_List())
	action_Cmd.AddCommand(Manifest_Action_Delete())

}
