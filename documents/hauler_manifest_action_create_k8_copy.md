## hauler manifest action create k8_copy

Copy a file from a pod to the local machine or vice versa

### Synopsis

Copy a file from a pod to the local machine or vice versa

```
hauler manifest action create k8_copy [[job_key add to job][] global action] [flags]
```

### Options

```
  -c, --cfg-container_name string   container name
  -d, --cfg-dest string             destination file path
  -s, --cfg-namespace string        Namespace to use (default "default")
  -f, --cfg-src string              source file path
  -t, --cfg-target_name string      The target name to use if not default target type
      --desc string                 description for the job
      --fail string                 The action to run if this action fails
  -k, --key string                  key for the job or global action
```

### Options inherited from parent commands

```
      --disable_required   disable the required flag for actions
      --file string        manifest file to use or document folder to use
      --help               help for hauler
```

### SEE ALSO

* [hauler manifest action create](hauler_manifest_action_create.md)	 - Create a action

###### Auto generated by spf13/cobra on 22-Apr-2023