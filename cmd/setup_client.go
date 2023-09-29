package cmd

import (
	"github.com/Mrpye/go-action-workflow/actions/action_condition"
	"github.com/Mrpye/go-action-workflow/actions/action_default"
	"github.com/Mrpye/go-action-workflow/actions/action_docker"
	"github.com/Mrpye/go-action-workflow/actions/action_file"
	"github.com/Mrpye/go-action-workflow/actions/action_git"
	"github.com/Mrpye/go-action-workflow/actions/action_govc"
	"github.com/Mrpye/go-action-workflow/actions/action_js"
	"github.com/Mrpye/go-action-workflow/actions/action_k8"
	"github.com/Mrpye/go-action-workflow/actions/action_parallel_workflow"
	"github.com/Mrpye/go-action-workflow/actions/action_scp"
	"github.com/Mrpye/go-action-workflow/actions/action_ssh"
	"github.com/Mrpye/go-action-workflow/actions/action_store"
	"github.com/Mrpye/go-action-workflow/actions/action_sub_workflow"
	"github.com/Mrpye/go-action-workflow/actions/action_system"
	"github.com/Mrpye/hauler/config"
	"github.com/Mrpye/hauler/pkg/client"
)

// *********************************
// Create the client global variable
// *********************************
var wf_client *client.Client

func init() {
	//*****************
	//Create the client
	//*****************
	wf_client = client.NewClient()

	//********************************
	//Add the config read/write funcs
	//********************************
	wf_client.AddReadConfig(config.ReadViperConfig)
	wf_client.AddWriteConfig(config.WriteViperConfig)
	wf_client.AddDeleteConfig(config.DeleteViperConfig)

	//************************************************
	//Add the function to map the config to the target
	//************************************************
	wf_client.AddTargetMapFunction(config.MapConfigToTarget)

	//*****************************************************
	//Add the function to map for the template function map
	//*****************************************************
	wf_client.Workflow.SetTemplateFuncMap(wf_client.Workflow.GetInbuiltTemplateFuncMap())

	//***********************************************
	//Add the target schemes and action to the client
	//***********************************************
	wf_client.AddActionSchema(action_default.GetSchema())
	wf_client.AddActionSchema(action_k8.GetSchema())
	wf_client.AddActionSchema(action_git.GetSchema())
	wf_client.AddActionSchema(action_file.GetSchema())
	wf_client.AddActionSchema(action_condition.GetSchema())
	wf_client.AddActionSchema(action_docker.GetSchema())
	wf_client.AddActionSchema(action_js.GetSchema())
	wf_client.AddActionSchema(action_parallel_workflow.GetSchema())
	wf_client.AddActionSchema(action_store.GetSchema())
	wf_client.AddActionSchema(action_sub_workflow.GetSchema())
	wf_client.AddActionSchema(action_store.GetSchema())
	wf_client.AddActionSchema(action_govc.GetSchema())
	wf_client.AddActionSchema(action_scp.GetSchema())
	wf_client.AddActionSchema(action_ssh.GetSchema())
	wf_client.AddActionSchema(action_system.GetSchema())

	//*************************
	//Build the target commands
	//*************************
	Build_CreateTargetCommand(wf_client.TargetScheme)
	Build_UpdateTargetCommand(wf_client.TargetScheme)
	Build_DeleteTargetCommand(wf_client.TargetScheme)

	//*************************
	//Build the action commands
	//*************************
	Build_CreateActionCommand(wf_client.ActionScheme)

}
