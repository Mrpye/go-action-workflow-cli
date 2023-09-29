package cmd

import (
	"reflect"
	"strings"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/Mrpye/hauler/pkg/client"
	"github.com/spf13/cobra"
)

func SaveTargetToConfig(w *client.Client, env string, target_name string, target interface{}) error {
	type_name := strings.ReplaceAll(strings.ToLower(reflect.TypeOf(target).String()), "*", "")

	//**********************************************
	//Use reflection to map the config to the target
	//**********************************************
	v := reflect.Indirect(reflect.ValueOf(target))
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		c := typeOfS.Field(i).Tag
		g := c.Get("yaml")
		if g == "" {
			continue
		}

		w.WriteConfig(type_name+"."+g, v.Field(i).Interface(), env)
	}
	return nil
}

// ************************************************
// BuildFlags - Build the flags based on the schema
// ************************************************
func BuildFlags(cmd *cobra.Command, schemas map[string]*workflow.Schema, prefix string) {
	for schema_key, config_schema := range schemas {
		//***************************************
		//Build the flag based on the schema type
		//***************************************
		switch config_schema.Type {
		case workflow.TypeString:
			var string_data string
			cmd.Flags().StringVarP(&string_data, prefix+schema_key, config_schema.Short, config_schema.Default.(string), config_schema.Description)
		case workflow.TypeBool:
			var bool_data bool
			cmd.Flags().BoolVarP(&bool_data, prefix+schema_key, config_schema.Short, config_schema.Default.(bool), config_schema.Description)
		case workflow.TypeInt:
			var int_data int
			cmd.Flags().IntVarP(&int_data, prefix+schema_key, config_schema.Short, config_schema.Default.(int), config_schema.Description)
		case workflow.TypeFloat:
			var float_data float64
			cmd.Flags().Float64VarP(&float_data, prefix+schema_key, config_schema.Short, config_schema.Default.(float64), config_schema.Description)
		case workflow.TypeList:
			var array_data []string
			cmd.Flags().StringArrayVarP(&array_data, prefix+schema_key, config_schema.Short, config_schema.Default.([]string), config_schema.Description)
		case workflow.TypeMap:
			var array_data []string
			cmd.Flags().StringArrayVarP(&array_data, prefix+schema_key, config_schema.Short, config_schema.Default.([]string), config_schema.Description)
		}

		if config_schema.Required {
			cmd.MarkFlagRequired(prefix + schema_key)

		}

	}
}
