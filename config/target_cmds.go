package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/Mrpye/golib/log"
	"github.com/spf13/viper"
)

// DeleteViperConfig deletes a key from the config file
// - key: the key to delete
// - custom: the custom data to use
func DeleteViperConfig(key string, custom ...string) error {
	if key == "" {
		return errors.New("key is empty")
	}
	base_path := fmt.Sprintf("targets.%s", custom[0])

	map_data := viper.Get(base_path)
	if map_data != nil {
		delete(map_data.(map[string]interface{}), key)
		viper.WriteConfig()
	}
	return nil
}

// ReadViperConfig reads a key from the config file
// - key: the key to read
// - data_type: the data type to return
// - custom: the custom data to use
func ReadViperConfig(key string, data_type string, custom ...string) (interface{}, error) {
	if key == "" {
		return nil, errors.New("key is empty")
	}

	base_path := fmt.Sprintf("%s.%s.%s", "targets", custom[0], key)

	switch data_type {
	case "string":
		return viper.GetString(base_path), nil
	case "int":
		return viper.GetInt(base_path), nil
	case "bool":
		return viper.GetBool(base_path), nil
	case "float64":
		return viper.GetFloat64(base_path), nil
	default:
		return viper.Get(base_path), nil
	}
}

// WriteViperConfig writes a key to the config file
// - key: the key to write
// - value: the value to write
// - custom: the custom data to use
func WriteViperConfig(key string, value interface{}, custom ...string) error {
	if key == "" {
		return errors.New("key is empty")
	}

	base_path := fmt.Sprintf("%s.%s.%s", "targets", custom[0], strings.ToLower(key))
	viper.Set(base_path, value)
	viper.WriteConfig()
	return nil
}

// MapConfigToTarget maps the config to the target
// - w: the workflow
// - m: the template data
// - target: the target to map to
// - returns: the target and error
func MapConfigToTarget(w *workflow.Workflow, m interface{}, target interface{}) (interface{}, error) {
	//*****************************
	//Get the env from runtime_vars
	//*****************************
	env, err := w.GetRuntimeVar("env")
	if err != nil {
		return nil, err
	}

	env = env.(string)
	if w.LogLevel == workflow.LOG_VERBOSE {
		log.LogVerbose(fmt.Sprintf("env: %s\n", env))
	}

	//*******************
	//Get the target_name
	//*******************
	target_name := ""
	var model *workflow.TemplateData
	switch v := m.(type) {
	case *workflow.TemplateData:
		model = v
		target_name, err = w.GetConfigTokenString("target_name", v, false)
		if err != nil {
			return nil, err
		}
	case string:
		target_name = v
	}

	type_name := strings.ReplaceAll(strings.ToLower(reflect.TypeOf(target).String()), "*", "")
	if target_name == "" {
		target_name = type_name
	} else {
		target_name = type_name + "_" + target_name
	}
	if w.LogLevel == workflow.LOG_VERBOSE {
		log.LogVerbose(fmt.Sprintf("Target name: %s\n", target_name))
	}

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
		//******************************************
		//Remap the config value to the target filed
		//******************************************
		if v.Field(i).Kind() == reflect.Int {
			temp_config_val, _ := w.GetConfigValue("", target_name+"."+g, "int", env.(string))
			if w.LogLevel == workflow.LOG_VERBOSE {
				log.LogVerbose(fmt.Sprintf("%s: %s\n", g, temp_config_val))
			}

			result, err := w.GetConfigTokenInt("target_"+g, model, true)
			if err != nil {
				v.Field(i).SetInt(temp_config_val.(int64))
			} else {
				v.Field(i).SetInt(int64(result))
			}

		} else if v.Field(i).Kind() == reflect.String {
			temp_config_val, _ := w.GetConfigValue("", target_name+"."+g, "string", env.(string))
			if w.LogLevel == workflow.LOG_VERBOSE {
				log.LogVerbose(fmt.Sprintf("%s: %v\n", g, temp_config_val))
			}
			result, err := w.GetConfigTokenString("target_"+g, model, true)
			if err != nil {
				v.Field(i).SetString(temp_config_val.(string))
			} else {
				v.Field(i).SetString(result)
			}

		} else if v.Field(i).Kind() == reflect.Bool {
			temp_config_val, _ := w.GetConfigValue("", target_name+"."+g, "bool", env.(string))
			if w.LogLevel == workflow.LOG_VERBOSE {
				log.LogVerbose(fmt.Sprintf("%s: %v\n", g, temp_config_val))
			}
			result, err := w.GetConfigTokenBool("target_"+g, model, true)
			if err != nil {
				v.Field(i).SetBool(temp_config_val.(bool))
			} else {
				v.Field(i).SetBool(result)
			}

		} else if v.Field(i).Kind() == reflect.Float64 {
			temp_config_val, _ := w.GetConfigValue("", target_name+"."+g, "float", env.(string))
			if w.LogLevel == workflow.LOG_VERBOSE {
				log.LogVerbose(fmt.Sprintf("%s: %v\n", g, temp_config_val))
			}
			result, err := w.GetConfigTokenFloat("target_"+g, model, true)
			if err != nil {
				v.Field(i).SetFloat(temp_config_val.(float64))
			} else {
				v.Field(i).SetFloat(result)
			}

		}
	}

	//See if to override the target values
	switch mv := m.(type) {
	case *workflow.TemplateData:
		//********************************
		//See if there is a current action
		//********************************
		if mv.CurrentAction == nil {
			return target, nil
		}
		//********************************
		//See if there target_* overrides
		//********************************
		v := reflect.Indirect(reflect.ValueOf(target))
		for k, e := range mv.CurrentAction.Config {
			if strings.Contains(k, "target_") {
				parts := strings.Split(k, "_")

				f := v.FieldByName(parts[1])
				field, ok := reflect.TypeOf(target).Elem().FieldByName(parts[1])
				if ok {
					g := field.Tag.Get("yaml")
					if g == "" {
						continue
					}
				}

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

	return target, nil
}

func GetTarget() {

}

func init() {
	//***********
	//Setup viper
	//***********
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()

}
