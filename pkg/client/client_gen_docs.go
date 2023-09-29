package client

import (
	"bytes"
	"reflect"
	"strings"
	"text/template"

	"github.com/Mrpye/go-action-workflow/workflow"
	"github.com/Mrpye/golib/encrypt"
	"github.com/Mrpye/golib/file"
)

// const doc_gen_template = "IyBDdXN0b20gQWN0aW9ucwoKe3tyYW5nZSAkYWN0aW9uX2dyb3VwX2tleSwgJGFjdGlvbl9ncm91cCA6PSAuR2V0R3JvdXBlZEFjdGlvbnN9fQojIyBQbHVnaW4gTGlicmFyeTogIHt7JGFjdGlvbl9ncm91cF9rZXl9fQoKe3tyYW5nZSAkYWN0aW9uX2tleSwgJGFjdGlvbiA6PSAkYWN0aW9uX2dyb3VwfX0KIyMjIEFjdGlvbjoge3skYWN0aW9uX2tleX19Cnt7JGFjdGlvbi5Mb25nfX0KCjxkZXRhaWxzPgo8c3VtbWFyeT5JbmZvPC9zdW1tYXJ5PgoKfFByb3BlcnR5fERlc2NyaXB0aW9ufAp8LS0tLXwtLS0tfAp8QWN0aW9ufHt7JGFjdGlvbl9rZXl9fXwKfERlc2NyaXB0aW9ufHt7JGFjdGlvbi5TaG9ydH19fAp8VGFyZ2V0fFt7eyRhY3Rpb24uVGFyZ2V0fX1dKCN0YXJnZXQte3tyZXBsYWNlICRhY3Rpb24uVGFyZ2V0IGAuYCBgYH19KXwKfElubGluZVBhcmFtc3x7eyRhY3Rpb24uSW5saW5lUGFyYW1zfX0KfFByb2Nlc3NSZXN1bHRzfHt7JGFjdGlvbi5Qcm9jZXNzUmVzdWx0c319fAoKIyMjIENvbmZpZyBQYXJhbWV0ZXJzCgp8cGFyYW1ldGVyfFR5cGV8RGVzY3JpcHRpb258UmVxdWlyZWR8U2hvcnQgRmxhZ3xEZWZhdWx0fAp8LS0tLXwtLS0tfC0tLS18LS0tLXwtLS0tfC0tLS18Cnt7cmFuZ2UgJGksICRlIDo9ICRhY3Rpb24uQ29uZmlnU2NoZW1hfX18e3skaX19fHt7JGUuVHlwZS5TdHJpbmd9fXx7eyRlLkRlc2NyaXB0aW9ufX18e3skZS5SZXF1aXJlZH19fHt7JGUuU2hvcnR9fXx8e3skZS5EZWZhdWx0fX18Cnt7ZW5kfX0KCjwvZGV0YWlscz4KCi0tLQp7e2VuZH19Cnt7ZW5kfX0KCgoKCgojIFRhcmdldHMKe3tyYW5nZSAkdGFyZ2V0X2tleSwgJHRhcmdldCA6PSAuVGFyZ2V0U2NoZW1lfX0KIyMgVGFyZ2V0IHt7JHRhcmdldF9rZXl9fQoKPGRldGFpbHM+CjxzdW1tYXJ5PkluZm88L3N1bW1hcnk+Cgp8RmllbGR8anNvbnxjbGkgZmxhZ3xEZXNjfAp8LS0tLXwtLS0tfC0tLS18LS0tLXwKe3tyYW5nZSAkaSwgJGUgOj0gJHRhcmdldC5HZXRUYXJnZXRNYXB9fXx7eyRpfX18e3t0YWcgJHRhcmdldC5UYXJnZXQgJGkgYHlhbWxgfX18e3t0YWcgJHRhcmdldC5UYXJnZXQgJGkgYGZsYWdgfX18e3t0YWcgJHRhcmdldC5UYXJnZXQgJGkgYGRlc2NgfX18Cnt7ZW5kfX0KCjwvZGV0YWlscz4KCi0tLQp7e2VuZH19CgoK"
const doc_gen_template = "IyBQbHVnaW4gTGlicmFyeSBEb2N1bWVudGF0aW9uCkJlbG93IGFyZSB0aGUgaW5zdHJ1Y3Rpb25zIGZvciB1c2luZyB0aGUgcGx1Z2lucyBsaWJyYXJ5IGluIHlvdXIgd29ya2Zsb3cuIEJlbG93IGFyZSB0aGUgc2V0dGluZyBhbmQgcGFyYW1ldGVycyBmb3IgZWFjaCBvZiB0aGUgYWN0aW9ucywgdGFyZ2V0IGFuZCBmdW5jdGlvbiBpbmNsdWRlZCBpbiB0aGUgbGlicmFyeS4KCi0tLQoKe3tyYW5nZSAkYWN0aW9uX2dyb3VwX2tleSwgJGFjdGlvbl9ncm91cCA6PSAuR2V0R3JvdXBlZEFjdGlvbnN9fQojICoqUGx1Z2luIExpYnJhcnkqKjogIHt7JGFjdGlvbl9ncm91cF9rZXl9fQoKIyBBY3Rpb25zOgoKe3tyYW5nZSAkYWN0aW9uX2tleSwgJGFjdGlvbiA6PSAkYWN0aW9uX2dyb3VwfX0KIyMgKipBY3Rpb24qKjoge3skYWN0aW9uX2tleX19Cnt7JGFjdGlvbi5Mb25nfX0KCjxkZXRhaWxzPgo8c3VtbWFyeT5JbmZvPC9zdW1tYXJ5PgoKfFByb3BlcnR5fERlc2NyaXB0aW9ufAp8LS0tLXwtLS0tfAp8QWN0aW9ufHt7JGFjdGlvbl9rZXl9fXwKfERlc2NyaXB0aW9ufHt7JGFjdGlvbi5TaG9ydH19fAp8VGFyZ2V0fFt7eyRhY3Rpb24uVGFyZ2V0fX1dKCN0YXJnZXQte3tyZXBsYWNlICRhY3Rpb24uVGFyZ2V0IGAuYCBgYH19KXwKfElubGluZVBhcmFtc3x7eyRhY3Rpb24uSW5saW5lUGFyYW1zfX0KfFByb2Nlc3NSZXN1bHRzfHt7JGFjdGlvbi5Qcm9jZXNzUmVzdWx0c319fAoKCiMjIyBDb25maWcgUGFyYW1ldGVycwoKfHBhcmFtZXRlcnxUeXBlfERlc2NyaXB0aW9ufFJlcXVpcmVkfFNob3J0IEZsYWd8RGVmYXVsdHwKfC0tLS18LS0tLXwtLS0tfC0tLS18LS0tLXwtLS0tfAp7e3JhbmdlICRpLCAkZSA6PSAkYWN0aW9uLkNvbmZpZ1NjaGVtYX19fHt7JGl9fXx7eyRlLlR5cGUuU3RyaW5nfX18e3skZS5EZXNjcmlwdGlvbn19fHt7JGUuUmVxdWlyZWR9fXx7eyRlLlNob3J0fX18fHt7JGUuRGVmYXVsdH19fAp7e2VuZH19Cgo8L2RldGFpbHM+CgotLS0Ke3tlbmR9fQoKIyBGdW5jdGlvbnM6CgojIyB7eyRhY3Rpb25fZ3JvdXBfa2V5fX0gRnVuY3Rpb25zOgoKe3tyYW5nZSAkZnVuY3Rpb25fa2V5LCAkZnVuY3Rpb24gOj0gJC5HZXRGdW5jdGlvbnNCeUxpYnJhcnkgJGFjdGlvbl9ncm91cF9rZXl9fQojIyAqKkZ1bmN0aW9uKio6IHt7JGZ1bmN0aW9uX2tleX19Cnt7JGZ1bmN0aW9uLkRlc2NyaXB0aW9ufX0KCjxkZXRhaWxzPgo8c3VtbWFyeT5JbmZvPC9zdW1tYXJ5PgoKfFByb3BlcnR5fERlc2NyaXB0aW9ufAp8LS0tLXwtLS0tfAp8RnVuY3Rpb258e3skZnVuY3Rpb25fa2V5fX18CnxEZXNjcmlwdGlvbnx7eyRmdW5jdGlvbi5EZXNjcmlwdGlvbn19fAoKIyMjIFBhcmFtZXRlcnMKfHBhcmFtZXRlcnxUeXBlfERlc2NyaXB0aW9ufFJlcXVpcmVkfAp8LS0tLXwtLS0tfC0tLS18LS0tLXwKe3tyYW5nZSAkaSwgJGUgOj0gJGZ1bmN0aW9uLlBhcmFtZXRlclNjaGVtYX19fHt7JGl9fXx7eyRlLlR5cGUuU3RyaW5nfX18e3skZS5EZXNjcmlwdGlvbn19fHt7JGUuUmVxdWlyZWR9fXwKe3tlbmR9fQoKPC9kZXRhaWxzPgoKLS0tCnt7ZW5kfX0KLS0tCnt7ZW5kfX0KCiMgVGFyZ2V0czoKe3tyYW5nZSAkdGFyZ2V0X2tleSwgJHRhcmdldCA6PSAuVGFyZ2V0U2NoZW1lfX0KIyMgKipUYXJnZXQqKiB7eyR0YXJnZXRfa2V5fX0KCjxkZXRhaWxzPgo8c3VtbWFyeT5JbmZvPC9zdW1tYXJ5PgoKfEZpZWxkfGpzb258Y2xpIGZsYWd8RGVzY3wKfC0tLS18LS0tLXwtLS0tfC0tLS18Cnt7cmFuZ2UgJGksICRlIDo9ICR0YXJnZXQuR2V0VGFyZ2V0TWFwfX18e3skaX19fHt7dGFnICR0YXJnZXQuVGFyZ2V0ICRpIGB5YW1sYH19fHt7dGFnICR0YXJnZXQuVGFyZ2V0ICRpIGBmbGFnYH19fHt7dGFnICR0YXJnZXQuVGFyZ2V0ICRpIGBkZXNjYH19fAp7e2VuZH19Cgo8L2RldGFpbHM+CgotLS0Ke3tlbmR9fQoKCg=="

func (m *Client) GetFunctionsByLibrary(lib string) map[string]*workflow.FunctionSchema {
	grouped := make(map[string]*workflow.FunctionSchema)
	for key, f := range m.FunctionScheme {
		parts := strings.Split(f.Target, ".")
		group := parts[0]
		if group == "" {
			group = "General"
		}
		if strings.EqualFold(group, lib) {
			grouped[key] = f
		}
	}
	return grouped
}

func (m *Client) GetGroupedActions() map[string]map[string]*workflow.ActionSchema {
	grouped := make(map[string]map[string]*workflow.ActionSchema)
	for key, action := range m.ActionScheme {
		parts := strings.Split(action.Target, ".")
		group := parts[0]
		if group == "" {
			group = "Inbuilt"
		}
		if _, ok := grouped[group]; !ok {
			grouped[group] = make(map[string]*workflow.ActionSchema)
		}
		grouped[group][key] = action
	}
	return grouped
}

func GetLib(TargetName string) string {
	parts := strings.Split(TargetName, ".")
	return parts[0]
}

func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}

func GetTag(t interface{}, field_name string, tag_name string) string {
	field, ok := reflect.TypeOf(t).Elem().FieldByName(field_name)
	if !ok {
		panic("Field not found")
	}
	return getStructTag(field, tag_name)
}

func GetObjectName(t interface{}) string {
	type_name := strings.ReplaceAll(strings.ToLower(reflect.TypeOf(t).String()), "*", "")
	return type_name
}

func (m *Client) doc_gen_funcMap() template.FuncMap {
	return template.FuncMap{
		"lc":       strings.ToLower, //Lowercase a string
		"uc":       strings.ToUpper, //Uppercase a string
		"tag":      GetTag,
		"lib":      GetLib,
		"replace":  strings.ReplaceAll,
		"obj_name": GetObjectName,
	}
}
func (m *Client) BuildActionDoc(output string) error {

	template_doc, err := encrypt.Base64DecString(doc_gen_template)
	if err != nil {
		return err
	}

	//********************************
	//Create a new template and parse
	//********************************
	tmpl, err := template.New("CodeRun").Funcs(m.doc_gen_funcMap()).Parse(template_doc)
	if err != nil {
		return err
	}

	//**************************************
	//Run the template to verify the output.
	//**************************************
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, m)
	if err != nil {
		return err
	}

	//*************
	//Save the file
	//*************
	err = file.SaveStringToFile(output, tpl.String())
	if err != nil {
		return err
	}

	//******************
	//Return the result
	//******************
	return nil
}
