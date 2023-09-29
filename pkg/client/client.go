package client

import (
	"github.com/Mrpye/go-action-workflow/workflow"
)

// Client is the main client object
type Client struct {
	ReadConfig     workflow.ReadConfigFunc             //workflow.ReadConfigFunc
	WriteConfig    workflow.WriteConfigFunc            //workflow.WriteConfigFunc
	DeleteConfig   workflow.DeleteConfigFunc           //workflow.DeleteConfigFunc
	FunctionScheme map[string]*workflow.FunctionSchema //map of action schemas
	TargetScheme   map[string]*workflow.TargetSchema   //map of target schemas
	ActionScheme   map[string]*workflow.ActionSchema   //map of action schemas
	Workflow       *workflow.Workflow                  //the workflow object
}

// AddAction adds an action and target schema to the client
// and adds the action to the workflow
func (dv *Client) AddActionSchema(sch workflow.SchemaEndpoint) {

	//******************************
	//Add the schema to the workflow
	//******************************
	dv.Workflow.AddActionSchema(sch)

	//***********************************
	//Add the target schema to the client
	//***********************************
	scheme := sch.GetTargetSchema()
	for key := range scheme {
		target_schema := scheme[key]
		dv.TargetScheme[key] = &target_schema
	}

	//********************************
	//Add the actions to the workflow
	//********************************
	actions := sch.GetActionSchema()
	for key := range actions {
		action_schema := actions[key]
		dv.ActionScheme[key] = &action_schema
	}

	//***********************************
	//Add the function map to the client
	//***********************************
	function_schemes := sch.GetFunctionMap()
	for key := range function_schemes {
		function_scheme := function_schemes[key]
		dv.FunctionScheme[key] = &function_scheme
	}
}

/*func UpdateConfig(m interface{}, config map[string]interface{}) {
	v := reflect.Indirect(reflect.ValueOf(m))
	//Test if the target_action is set
	if val, ok := config["target_action"]; ok {
		//do something here
		if val == "update" || val == "create" {
			for k, e := range config {
				if strings.Contains(k, "target_") {
					parts := strings.Split(k, "_")
					f := v.FieldByName(parts[1])
					if f.Kind() == reflect.Int {
						f.SetInt(e.(int64))
					} else if f.Kind() == reflect.String {
						f.SetString(e.(string))
					} else if f.Kind() == reflect.Bool {
						f.SetBool(e.(bool))
					}
				}
			}
		}
	}
}*/

func (dv *Client) AddTargetMapFunction(f workflow.TargetMapFunc) {
	dv.Workflow.TargetMapFunc = f
}

// AddReadConfig adds a read config function to the client
// and adds the function to the workflow
func (dv *Client) AddReadConfig(f workflow.ReadConfigFunc) {
	dv.ReadConfig = f
	dv.Workflow.ReadConfigFunc["viper"] = f
}

// AddWriteConfig adds a write config function to the client
func (dv *Client) AddWriteConfig(f workflow.WriteConfigFunc) {
	dv.WriteConfig = f
}

// AddDeleteConfig adds a delete config function to the client
func (dv *Client) AddDeleteConfig(f workflow.DeleteConfigFunc) {
	dv.DeleteConfig = f
}

// NewClient creates a new client object
func NewClient() *Client {
	//************************
	//Create the client object
	//************************
	client := &Client{
		TargetScheme:   make(map[string]*workflow.TargetSchema),
		ActionScheme:   make(map[string]*workflow.ActionSchema),
		FunctionScheme: make(map[string]*workflow.FunctionSchema),
		Workflow:       workflow.CreateWorkflow(),
	}

	//**********************************
	//Only show errors and print actions
	//**********************************
	client.Workflow.LogLevel = workflow.LOG_INFO

	return client
}
