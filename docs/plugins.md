# Plugins
Vul provides a plugin feature to allow others to extend the Vul CLI without the need to change the Vulcode base.
This plugin system was inspired by the plugin system used in [kubectl][kubectl], [Helm][helm], and [Conftest][conftest].

## Overview
Vul plugins are add-on tools that integrate seamlessly with Vul.
They provide a way to extend the core feature set of Vul, but without requiring every new feature to be written in Go and added to the core tool.

- They can be added and removed from a Vul installation without impacting the core Vul tool.
- They can be written in any programming language.
- They integrate with Vul, and will show up in Vul help and subcommands.

!!! warning
    Vul plugins available in public are not audited for security.
    You should install and run third-party plugins at your own risk, since they are arbitrary programs running on your machine.


## Installing a Plugin
A plugin can be installed using the `vul plugin install` command.
This command takes a url and will download the plugin and install it in the plugin cache.

Vul adheres to the XDG specification, so the location depends on whether XDG_DATA_HOME is set.
Vul will now search XDG_DATA_HOME for the location of the Vul plugins cache.
The preference order is as follows:

- XDG_DATA_HOME if set and .vul/plugins exists within the XDG_DATA_HOME dir
- ~/.vul/plugins

Under the hood Vul leverages [go-getter][go-getter] to download plugins.
This means the following protocols are supported for downloading plugins:

- OCI Registries
- Local Files
- Git
- HTTP/HTTPS
- Mercurial
- Amazon S3
- Google Cloud Storage

For example, to download the Kubernetes Vul plugin you can execute the following command:

```bash
$ vul plugin install github.com/khulnasoft-lab/vul-plugin-kubectl
```
## Using Plugins
Once the plugin is installed, Vul will load all available plugins in the cache on the start of the next Vul execution.
A plugin will be made in the Vul CLI based on the plugin name.
To display all plugins, you can list them by `vul --help`

```bash
$ vul --help
NAME:
   vul - A simple and comprehensive vulnerability scanner for containers

USAGE:
   vul [global options] command [command options] target

VERSION:
   dev

COMMANDS:
   image, i          scan an image
   filesystem, fs    scan local filesystem
   repository, repo  scan remote repository
   client, c         client mode
   server, s         server mode
   plugin, p         manage plugins
   kubectl           scan kubectl resources
   help, h           Shows a list of commands or help for one command
```

As shown above, `kubectl` subcommand exists in the `COMMANDS` section.
To call the kubectl plugin and scan existing Kubernetes deployments, you can execute the following command:

```
$ vul kubectl deployment <deployment-id> -- --ignore-unfixed --severity CRITICAL
```

Internally the kubectl plugin calls the kubectl binary to fetch information about that deployment and passes the using images to Vul.
You can see the detail [here][vul-plugin-kubectl].

If you want to omit even the subcommand, you can use `VUL_RUN_AS_PLUGIN` environment variable.

```bash
$ VUL_RUN_AS_PLUGIN=kubectl vul job your-job -- --format json
```

## Installing and Running Plugins on the fly
`vul plugin run` installs a plugin and runs it on the fly.
If the plugin is already present in the cache, the installation is skipped.

```bash
vul plugin run github.com/khulnasoft-lab/vul-plugin-kubectl pod your-pod -- --exit-code 1
```

## Uninstalling Plugins
Specify a plugin name with `vul plugin uninstall` command.

```bash
$ vul plugin uninstall kubectl
```

## Building Plugins
Each plugin has a top-level directory, and then a plugin.yaml file.

```bash
your-plugin/
  |
  |- plugin.yaml
  |- your-plugin.sh
```

In the example above, the plugin is contained inside of a directory named `your-plugin`.
It has two files: plugin.yaml (required) and an executable script, your-plugin.sh (optional).

The core of a plugin is a simple YAML file named plugin.yaml.
Here is an example YAML of vul-plugin-kubectl plugin that adds support for Kubernetes scanning.

```yaml
name: "kubectl"
repository: github.com/khulnasoft-lab/vul-plugin-kubectl
version: "0.1.0"
usage: scan kubectl resources
description: |-
  A Vul plugin that scans the images of a kubernetes resource.
  Usage: vul kubectl TYPE[.VERSION][.GROUP] NAME
platforms:
  - selector: # optional
      os: darwin
      arch: amd64
    uri: ./vul-kubectl # where the execution file is (local file, http, git, etc.)
    bin: ./vul-kubectl # path to the execution file
  - selector: # optional
      os: linux
      arch: amd64
    uri: https://github.com/khulnasoft-lab/vul-plugin-kubectl/releases/download/v0.1.0/vul-kubectl.tar.gz
    bin: ./vul-kubectl
```

The `plugin.yaml` field should contain the following information:

- name: The name of the plugin. This also determines how the plugin will be made available in the Vul CLI. For example, if the plugin is named kubectl, you can call the plugin with `vul kubectl`. (required)
- version: The version of the plugin. (required)
- usage: A short usage description. (required)
- description: A long description of the plugin. This is where you could provide a helpful documentation of your plugin. (required)
- platforms: (required)
  - selector: The OS/Architecture specific variations of a execution file. (optional)
    - os: OS information based on GOOS (linux, darwin, etc.) (optional)
    - arch: The architecture information based on GOARCH (amd64, arm64, etc.) (optional)
  - uri: Where the executable file is. Relative path from the root directory of the plugin or remote URL such as HTTP and S3. (required)
  - bin: Which file to call when the plugin is executed. Relative path from the root directory of the plugin. (required)

The following rules will apply in deciding which platform to select:

- If both `os` and `arch` under `selector` match the current platform, search will stop and the platform will be used.
- If `selector` is not present, the platform will be used.
- If `os` matches and there is no more specific `arch` match, the platform will be used.
- If no `platform` match is found, Vul will exit with an error.

After determining platform, Vul will download the execution file from `uri` and store it in the plugin cache.
When the plugin is called via Vul CLI, `bin` command will be executed.

The plugin is responsible for handling flags and arguments. Any arguments are passed to the plugin from the `vul` command.

## Example
https://github.com/khulnasoft-lab/vul-plugin-kubectl

[kubectl]: https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/
[helm]: https://helm.sh/docs/topics/plugins/
[conftest]: https://www.conftest.dev/plugins/
[go-getter]: https://github.com/hashicorp/go-getter
[vul-plugin-kubectl]: https://github.com/khulnasoft-lab/vul-plugin-kubectl

