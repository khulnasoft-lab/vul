# Configuration
Vul can be configured using the following ways. Each item takes precedence over the item below it:

- CLI flags
- Environment variables
- Configuration file

## CLI Flags
You can view the list of available flags using the `--help` option.
For more details, please refer to [the CLI reference](../references/configuration/cli/vul.md).

## Environment Variables
Vul can be customized by environment variables.
The environment variable key is the flag name converted by the following procedure.

- Add `VUL_` prefix
- Make it all uppercase
- Replace `-` with `_`

For example,

- `--debug` => `VUL_DEBUG`
- `--cache-dir` => `VUL_CACHE_DIR`

```
$ VUL_DEBUG=true VUL_SEVERITY=CRITICAL vul image alpine:3.15
```

## Configuration File
By default, Vul reads the `vul.yaml` file.
For more details, please refer to [the page](../references/configuration/config-file.md).
