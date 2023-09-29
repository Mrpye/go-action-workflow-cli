# Plugin Library Documentation
Below are the instructions for using the plugins library in your workflow. Below are the setting and parameters for each of the actions, target and function included in the library.

---


# **Plugin Library**:  Inbuilt

# Actions:


## **Action**: action
Runs a Global Action

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|action|
|Description|Runs a Global Action|
|Target|[](#target-)|
|InlineParams|true
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|action_key|String|Message to print|true|a|||


</details>

---

## **Action**: condition
Run an action based on a condition

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|condition|
|Description|Run an action based on a condition|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|condition|String|The condition to evaluate must evaluate to a boolean|true|c|||
|fail|String|Action or Action Key to run if the condition failed|true|f|||
|pass|String|Action or Action Key to run if the condition passes|true|p|||


</details>

---

## **Action**: end
End the workflow

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|end|
|Description|End the workflow|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|


</details>

---

## **Action**: fail
Fail the workflow

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|fail|
|Description|Fail the workflow|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|


</details>

---

## **Action**: file_append
Append a file

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|file_append|
|Description|Append a file|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|content|String|Content to append to the file|true|c|||
|source_file|String|Source file to append to|true|s|||


</details>

---

## **Action**: file_copy
Copy a file

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|file_copy|
|Description|Copy a file|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|dest_file|String|Destination file to copy to|true|d|||
|source_file|String|Source file to copy|true|s|||


</details>

---

## **Action**: file_create
Create a file

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|file_create|
|Description|Create a file|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|content|String|Content to add to the file|true|c|||
|source_file|String|Source file to create|true|s|||


</details>

---

## **Action**: file_delete
Delete a file

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|file_delete|
|Description|Delete a file|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|source_file|String|Source file to rename|true|s|||


</details>

---

## **Action**: file_rename
Rename a file

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|file_rename|
|Description|Rename a file|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|dest_file|String|Destination file to rename to|true|d|||
|source_file|String|Source file to rename|true|s|||


</details>

---

## **Action**: file_template
Create a file using a template

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|file_template|
|Description|Create a file using a template|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|data|Map|Map Data to use in the template|true|d||[]|
|file|String|file to create|true|f|||
|template|String|Template file|true|t|||


</details>

---

## **Action**: for
for loop

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|for|
|Description|for loop|
|Target|[](#target-)|
|InlineParams|true
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|from|String|From Value|true|f|||
|to|String|To Value|true|t|||
|variable|String|Variable for the loop|true|v|||


</details>

---

## **Action**: goto
goto

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|goto|
|Description|goto|
|Target|[](#target-)|
|InlineParams|true
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|key|String|Action to goto|true|i|||


</details>

---

## **Action**: next
next loop

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|next|
|Description|next loop|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|


</details>

---

## **Action**: parallel
Store a value in the data bucket

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|parallel|
|Description|Store a value in the data bucket|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|actions|List|A list of actions to run in parallel|true|a||[]|


</details>

---

## **Action**: print
Print a value

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|print|
|Description|Print a value|
|Target|[](#target-)|
|InlineParams|true
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|message|String|Message to print|true|m|||


</details>

---

## **Action**: store
Store a value in the data bucket

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|store|
|Description|Store a value in the data bucket|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|bucket|String|The bucket to store the value in|true|b|||
|key|String|The key to store the value with|true|i|||
|value|String|The value to store|true|v|||


</details>

---

## **Action**: sub_workflow
Set Environment Variables

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|sub_workflow|
|Description|Set Environment Variables|
|Target|[](#target-)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|env_vars|Map|Map of environment variables to set|true|e||[]|


</details>

---


# Functions:

## Inbuilt Functions:


---

# **Plugin Library**:  action_docker

# Actions:


## **Action**: docker_reg_download
Download a docker image from a registry

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|docker_reg_download|
|Description|Download a docker image from a registry|
|Target|[action_docker.registry](#target-action_dockerregistry)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|folder|String|Folder to save the images to|true|f|||
|images|List|List of images to download|true|i||[]|
|target_name|String|The target to use|false|t|||


</details>

---

## **Action**: docker_reg_upload
Upload a docker image to a registry

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|docker_reg_upload|
|Description|Upload a docker image to a registry|
|Target|[action_docker.registry](#target-action_dockerregistry)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|folder|String|Folder to load the images from|true|f|||
|images|List|List of images to upload|true|i||[]|
|import_all|Bool|Import all images in the folder|false|a||false|
|target_name|String|The target to use|false|t|||


</details>

---


# Functions:

## action_docker Functions:


## **Function**: image_account
gets the account name from the image path

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|image_account|
|Description|gets the account name from the image path|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|image|String|docker.io/[circleci]/slim-base:latest|true|


</details>

---

## **Function**: image_name
gets the name of the image from the image path

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|image_name|
|Description|gets the name of the image from the image path|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|image|String|docker.io/circleci/[slim-base]:latest|true|


</details>

---

## **Function**: image_name_tag
gets the name and tag from the image path

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|image_name_tag|
|Description|gets the name and tag from the image path|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|image|String|docker.io/circleci/[slim-base:latest]|true|


</details>

---

## **Function**: image_registry
gets the registry name from the image path

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|image_registry|
|Description|gets the registry name from the image path|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|image|String|docker.io/circleci/slim-base]:latest|true|


</details>

---

## **Function**: image_shortname
gets the account name from the image path

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|image_shortname|
|Description|gets the account name from the image path|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|image|String|docker.io/[circleci/slim-base]:latest|true|


</details>

---

## **Function**: image_tag
gets the tag  from the image path

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|image_tag|
|Description|gets the tag  from the image path|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|image|String|docker.io/circleci/slim-base:[latest]|true|


</details>

---

## **Function**: remap_image
remaps the docker image

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|remap_image|
|Description|remaps the docker image|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|*workflow.Workflow|Workflow|work flow engine use (get_wf) to get the workflow engine|true|
|image|String|name of the image|true|
|no_tag|Bool|true not include the tag in the returned value|true|
|original|String|the original image path to use if use_original=true else build path based on target|false|
|target_name|String|name of the target or just use ``|false|
|use_original|Bool|use the original image path|true|


</details>

---

## **Function**: remap_image2
remaps the docker image no option for using default image path

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Function|remap_image2|
|Description|remaps the docker image no option for using default image path|

### Parameters
|parameter|Type|Description|Required|
|----|----|----|----|
|*workflow.Workflow|Workflow|work flow engine use (get_wf) to get the workflow engine|true|
|image|String|name of the image|true|
|no_tag|Bool|true not include the tag in the returned value|true|
|target_name|String|name of the target or just use ``|false|


</details>

---

---

# **Plugin Library**:  action_git

# Actions:


## **Action**: git_download
Apply or delete a yaml file to a k8 cluster

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|git_download|
|Description|Apply or delete a yaml file to a k8 cluster|
|Target|[action_git.git](#target-action_gitgit)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|files|List|List of files to download|false|f||[]|
|target_name|String|The target to use|false|t|||


</details>

---


# Functions:

## action_git Functions:


---

# **Plugin Library**:  action_govc

# Actions:


## **Action**: govc
Run GOVC commands

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|govc|
|Description|Run GOVC commands|
|Target|[action_govc.vcenter](#target-action_govcvcenter)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|command|String|Command to run|true|c|||
|target_name|String|The target to use|false|t|||


</details>

---


# Functions:

## action_govc Functions:


---

# **Plugin Library**:  action_k8

# Actions:


## **Action**: k8_copy
Copy a file from a pod to the local machine or vice versa

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_copy|
|Description|Copy a file from a pod to the local machine or vice versa|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|container_name|String|container name|true|c|||
|dest|String|destination file path|true|d|||
|namespace|String|Namespace to use|false|s||default|
|src|String|source file path|true|f|||
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_create_ns
Create a namespace

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_create_ns|
|Description|Create a namespace|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_demon_set
Delete a demon set

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_demon_set|
|Description|Delete a demon set|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the demon set|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_deployment
Delete Deployment

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_deployment|
|Description|Delete Deployment|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the Deployment|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_ns
Delete a namespace

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_ns|
|Description|Delete a namespace|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_pod
Delete Pod

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_pod|
|Description|Delete Pod|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the Pod|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_pv
Delete a PV

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_pv|
|Description|Delete a PV|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the PV|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_pvc
Delete a PVC

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_pvc|
|Description|Delete a PVC|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the PVC|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_secret
Delete a secret

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_secret|
|Description|Delete a secret|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the secret|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_service
Delete Service

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_service|
|Description|Delete Service|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the service|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_delete_stateful_set
Delete a stateful set

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_delete_stateful_set|
|Description|Delete a stateful set|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|Name of the stateful set|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_get_service_ip
Get the IP of a service

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_get_service_ip|
|Description|Get the IP of a service|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|name of service can use regex|true|n|||
|namespace|String|Namespace to use|false|s||default|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_get_ws_items
Get items in a workspace

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_get_ws_items|
|Description|Get items in a workspace|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|namespace|String|Namespace to use|false|S||default|
|target_name|String|The target name to use if not default target type|false|t|||
|workspace|String|workspace to use|true|w|||


</details>

---

## **Action**: k8_helm_add_repo
Add a helm repo

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_helm_add_repo|
|Description|Add a helm repo|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|name|String|repo name|false|a|||
|namespace|String|Namespace to use|false|s||default|
|password|String|Password|false|p|||
|target_name|String|The target name to use if not default target type|false|t|||
|url|String|Url of repo|false|||h|
|use_config|Bool|Use the target  config|false|c||false|
|username|String|Username|false|u|||


</details>

---

## **Action**: k8_helm_delete
Delete a helm chart

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_helm_delete|
|Description|Delete a helm chart|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|namespace|String|Namespace to use|false|s||default|
|release_name|String|release name to use|true|n|||
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_helm_deploy_upgrade
Deploy or upgrade a helm chart

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_helm_deploy_upgrade|
|Description|Deploy or upgrade a helm chart|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|chart_Path|String|chart path|false|c|||
|namespace|String|Namespace to use|false|s||default|
|release_name|String|release name|false|n|||
|target_name|String|The target name to use if not default target type|false|t|||
|upgrade|Bool|chart path|false|u||false|


</details>

---

## **Action**: k8_pod_exec
Execute a command in a pod

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_pod_exec|
|Description|Execute a command in a pod|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|command|String|command to execute|true|c|||
|namespace|String|Namespace to use|false|s||default|
|pod_name|String|pod name|true|n|||
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_wait
Wait for a k8 resource to be in a complete state

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_wait|
|Description|Wait for a k8 resource to be in a complete state|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|checks|List|Checks to run|false|c||[replica:nginx2(.*) stateful:nginx3(.*) demon:nginx4(.*) service:nginx(.*)]|
|namespace|String|Namespace to use|false|s||default|
|not_running|Bool|All checks not running|false|x||false|
|retry|Int|retry count|false|r||10|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---

## **Action**: k8_yaml
Apply or delete a yaml file to a k8 cluster

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|k8_yaml|
|Description|Apply or delete a yaml file to a k8 cluster|
|Target|[action_k8.k8](#target-action_k8k8)|
|InlineParams|false
|ProcessResults|false|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|delete|Bool|If true, delete the deployment|false|d||false|
|manifest|String|k8 manifest to apply or delete as a file path or object|true|m|||
|namespace|String|Namespace to use|false|s||default|
|process_tokens|Bool|Should tokens be processed in the k8 manifest|false|p||true|
|target_name|String|The target name to use if not default target type|false|t|||


</details>

---


# Functions:

## action_k8 Functions:


---

# **Plugin Library**:  action_scp

# Actions:


## **Action**: scp_download
Download file to scp server

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|scp_download|
|Description|Download file to scp server|
|Target|[action_scp.scp](#target-action_scpscp)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|dest|String|Destination file to download to|true|d|||
|file|String|Source file to download from the server|true|f|||
|target_name|String|The target to use|false|t|||


</details>

---

## **Action**: scp_upload
Upload file to scp server

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|scp_upload|
|Description|Upload file to scp server|
|Target|[action_scp.scp](#target-action_scpscp)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|dest|String|Destination file to upload to on the server|true|d|||
|file|String|Source file to upload|true|f|||
|target_name|String|The target to use|false|t|||


</details>

---


# Functions:

## action_scp Functions:


---

# **Plugin Library**:  action_ssh

# Actions:


## **Action**: ssh_download
Download a file to a SSH server

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|ssh_download|
|Description|Download a file to a SSH server|
|Target|[action_ssh.ssh](#target-action_sshssh)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|dest|String|Destination file to download to|true|d|||
|source|String|Source file to download|true|s|||
|target_name|String|The target to use|false|t|||


</details>

---

## **Action**: ssh_run_cmd
Run a command on a SSH server

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|ssh_run_cmd|
|Description|Run a command on a SSH server|
|Target|[action_ssh.ssh](#target-action_sshssh)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|command|String|command to run on the server|true|c|||
|target_name|String|The target to use|false|t|||


</details>

---

## **Action**: ssh_run_script
Run a script on a SSH server

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|ssh_run_script|
|Description|Run a script on a SSH server|
|Target|[action_ssh.ssh](#target-action_sshssh)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|script|String|Script to run on the server|true|s|||
|target_name|String|The target to use|false|t|||


</details>

---

## **Action**: ssh_run_script_file
Run a script file on a SSH server

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|ssh_run_script_file|
|Description|Run a script file on a SSH server|
|Target|[action_ssh.ssh](#target-action_sshssh)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|file|String|Script file to run on the server|true|f|||
|target_name|String|The target to use|false|t|||


</details>

---

## **Action**: ssh_upload
Upload a file to a SSH server

<details>
<summary>Info</summary>

|Property|Description|
|----|----|
|Action|ssh_upload|
|Description|Upload a file to a SSH server|
|Target|[action_ssh.ssh](#target-action_sshssh)|
|InlineParams|false
|ProcessResults|true|


### Config Parameters

|parameter|Type|Description|Required|Short Flag|Default|
|----|----|----|----|----|----|
|dest|String|Destination file to download to|true|d|||
|source|String|Source file to upload|true|s|||
|target_name|String|The target to use|false|t|||


</details>

---


# Functions:

## action_ssh Functions:


---


# Targets:

## **Target** action_docker.registry

<details>
<summary>Info</summary>

|Field|json|cli flag|Desc|
|----|----|----|----|
|Host|host|host h|the host url|
|IgnoreSSL|ignore_ssl|ignore_ssl i|Ignore SSL|
|Library|library|library l|Library to use|
|Password|password|password p|Password for the registry|
|UserName|user|user u|Username for the registry|


</details>

---

## **Target** action_git.git

<details>
<summary>Info</summary>

|Field|json|cli flag|Desc|
|----|----|----|----|
|Authorization|authorization|auth a|The auth token|
|Email|email|email e|The Email address|
|Host|host|host h|the host url|
|Ssh|ssh|ssh s|The SSH key|
|User|user|user u|The user name|


</details>

---

## **Target** action_govc.vcenter

<details>
<summary>Info</summary>

|Field|json|cli flag|Desc|
|----|----|----|----|
|DataCenter|data_center|data_center c|Data Center|
|DataStore|data_store|data_store s|Data Store|
|Host|host|host h|Url to vcenter|
|IgnoreSSL|ignore_ssl|ignore_ssl i|Ignore SSL|
|Network|network|network n|Network|
|Password|password|password p|The password for the vcenter|
|ResourcePool|resource_pool|resource_pool r|Resource Pool|
|User|user|user u|The user name for the vcenter|


</details>

---

## **Target** action_k8.k8

<details>
<summary>Info</summary>

|Field|json|cli flag|Desc|
|----|----|----|----|
|Authorization|authorization|auth a|The authorization token|
|ConfigPath|config_path|config_path p|The path to the kube config file|
|DefaultContext|default_context|context c|The default context to use|
|Host|host|host h|The host to connect to|
|Ignore_ssl|ignore_ssl|ignore_ssl i|If true, ignore the ssl connection|
|UseTokenConnection|use_token_connection|conn-type u|Connection type if true, use the token connection, otherwise use the kube config file|


</details>

---

## **Target** action_scp.scp

<details>
<summary>Info</summary>

|Field|json|cli flag|Desc|
|----|----|----|----|
|Host|host|host h|SCP host|
|Password|password|password p|Password for SCP|
|User|user|user u|User for SCP|


</details>

---

## **Target** action_ssh.ssh

<details>
<summary>Info</summary>

|Field|json|cli flag|Desc|
|----|----|----|----|
|Host|host|host h|SSH host|
|Password|password|password p|Password for SSH|
|PrivateKeyFile|private_key_file|private_key_file s|private key file for SSH|
|User|user|user u|User for SSH|


</details>

---



