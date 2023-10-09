## vul

Unified security scanner

### Synopsis

Scanner for vulnerabilities in container images, file systems, and Git repositories, as well as for configuration issues and hard-coded secrets

```
vul [global flags] command [flags] target
```

### Examples

```
  # Scan a container image
  $ vul image python:3.4-alpine

  # Scan a container image from a tar archive
  $ vul image --input ruby-3.1.tar

  # Scan local filesystem
  $ vul fs .

  # Run in server mode
  $ vul server
```

### Options

```
      --cache-dir string          cache directory (default "/path/to/cache")
  -c, --config string             config path (default "vul.yaml")
  -d, --debug                     debug mode
  -f, --format string             version format (json)
      --generate-default-config   write the default config to vul-default.yaml
  -h, --help                      help for vul
      --insecure                  allow insecure server connections
  -q, --quiet                     suppress progress bar and log output
      --timeout duration          timeout (default 5m0s)
  -v, --version                   show version
```

### SEE ALSO

* [vul aws](vul_aws.md)	 - [EXPERIMENTAL] Scan AWS account
* [vul config](vul_config.md)	 - Scan config files for misconfigurations
* [vul convert](vul_convert.md)	 - Convert Vul JSON report into a different format
* [vul filesystem](vul_filesystem.md)	 - Scan local filesystem
* [vul image](vul_image.md)	 - Scan a container image
* [vul kubernetes](vul_kubernetes.md)	 - [EXPERIMENTAL] Scan kubernetes cluster
* [vul module](vul_module.md)	 - Manage modules
* [vul plugin](vul_plugin.md)	 - Manage plugins
* [vul repository](vul_repository.md)	 - Scan a repository
* [vul rootfs](vul_rootfs.md)	 - Scan rootfs
* [vul sbom](vul_sbom.md)	 - Scan SBOM for vulnerabilities
* [vul server](vul_server.md)	 - Server mode
* [vul version](vul_version.md)	 - Print the version
* [vul vm](vul_vm.md)	 - [EXPERIMENTAL] Scan a virtual machine image

