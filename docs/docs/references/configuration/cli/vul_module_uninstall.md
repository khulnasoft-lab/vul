## vul module uninstall

Uninstall a module

```
vul module uninstall [flags] REPOSITORY
```

### Options

```
  -h, --help   help for uninstall
```

### Options inherited from parent commands

```
      --cache-dir string          cache directory (default "/path/to/cache")
  -c, --config string             config path (default "vul.yaml")
  -d, --debug                     debug mode
      --enable-modules strings    [EXPERIMENTAL] module names to enable
      --generate-default-config   write the default config to vul-default.yaml
      --insecure                  allow insecure server connections
      --module-dir string         specify directory to the wasm modules that will be loaded (default "$HOME/.vul/modules")
  -q, --quiet                     suppress progress bar and log output
      --timeout duration          timeout (default 5m0s)
  -v, --version                   show version
```

### SEE ALSO

* [vul module](vul_module.md)	 - Manage modules

