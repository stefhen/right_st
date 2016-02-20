# right_st

right_st is a tool for managing RightScale ServerTemplate and RightScripts. The tool is able to upload and download ServerTemplate and RightScripts using RightScale's API. This tool can easily be hooked into Travis CI or other build systems to manage these design objects if stored in Github. See below for usage examples

## Managing RightScripts

RightScripts consist of a script body, attachments, and metadata. Metadata is the list of inputs, description, and name of the RightScript. Metadata is expected to be embedded as a YAML formatted comment at the top of the RightScript. Metadata format is as follows:

| Field | Format | Description |
| ----- | ------ | ----------- |
| RightScript Name | String | Name of RightScript. Name must be unique for your account. |
| Description | String | Description field for the RightScript |
| Inputs | Hash of String -> Input | The hash key is the input name. The hash value is an Input definition (defined below) |
| Attachments | Array of Strings | Each string is a filename of an attachment file. Attachments must be placed in an "attachments/" subdirectory |

Input definition format is as follows:

| Field | Format | Description |
| ----- | ------ | ----------- |
| Input Type | String | "single" or "array" |
| Category | String | Category to group Input under |
| Description | String | Description field for the Input |
| Default | String | Default value. Value must be in [Inputs 2.0 format](http://reference.rightscale.com/api1.5/resources/ResourceInputs.html) from RightScale API which consists of a type followed by a colon then the value. I.e. "text:Foo", "ignore", "env:ENV_VAR" |
| Required | Boolean | "true" or "false". Whether or not the Input is required |
| Advanced | Boolean | "true" or "false". Whether or not the Input is advanced (hidden by default) |


Example RightScript is as follows. This RightScript has one attachment, which must be located at "attachments/foo" relative
to the script.
```bash
#! /bin/bash -e

# ---
# RightScript Name: Run Foo Tool
# Description: Runs attached foo executable with input
# Inputs:
#   FOO_PARAM:
#     Input Type: single
#     Category: RightScale
#     Description: A parameter to the foo tool
#     Default: "text:bar"
#     Required: false
#     Advanced: true
# Attachments:
# - foo
# ...
#

cp -f $RS_ATTACH_DIR/foo /usr/local/bin/foo
chmod a+x /usr/local/bin/foo
foo $FOO_PARAM
```
 
### Usage
RightScript related commands are as follows:

~~~bash
right_st rightscript show <name|href|id>
  Show a single RightScript and its attachments

right_st rightscript upload [<flags>] <path>...
  Upload a RightScript

right_st rightscript download <name|href|id> [<path>]
  Download a RightScript to a file or files

right_st rightscript scaffold [<flags>] <path>...
  Add RightScript YAML metadata comments to a file or files

right_st rightscript validate <path>...
  Validate RightScript YAML metadata comments in a file or files
~~~


## Managing ServerTemplates

ServerTemplates are defined by a YAML format representing the ServerTemplate. The following keys are supported:

| Field | Format | Description |
| ----- | ------ | ----------- |
| Name | String | Name of the ServerTemplate. Name must be unique for your account. |
| Description | String | Description field for the ServerTemplate. |
| RightScripts | Hash | The hash key is the sequence type, one of "Boot", "Operational", or "Decommission". The hash value is a array of strings, where each string is a relative pathname to a RightScript on disk. |
| Inputs | Hash of String -> String | The hash key is the input name. The hash value is the default value. Note this inputs array is much simpler than the Input definition in RightScripts - only default values can be overriden in a ServerTemplate. |
| MultiCloudImages | Array of MultiCloudImages | An array of MultiCloudImage definitions. A MultiCloudImage definition is a hash specifying a MCI. Currently the only MultiCloudImage definition format is by Href. See example below |

Here is an example ServerTemplate yaml file:
```yaml
Name: My ServerTemplate
Description: This is an example Description
Inputs:
  FIRST_INPUT: "text:overriding value"
  SECOND_INPUT: "env:RS_UUID"
RightScripts:
  Boot:
  - path/to/script1.sh
  - path/to/script2.sh
  Operational:
  - path/to/script1.sh
  Decommission:
  - path/to/decom_script1.sh
MultiCloudImages:
- Href: /api/multi_cloud_images/403042003
```

### Usage

TBD

## Contributors

This tool is maintained by [Douglas Thrift (douglaswth)](https://github.com/douglaswth),
[Peter Schroeter (psschroeter)](https://github.com/psschroeter), and [Lopaka Delp (lopaka)](https://github.com/lopaka).

## License

The `right_st` source code is subject to the MIT license, see the
[LICENSE](https://github.com/douglaswth/right_st/blob/master/LICENSE) file.
