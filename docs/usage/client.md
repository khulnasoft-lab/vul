# Client

```bash
NAME:
   vul client - client mode

USAGE:
   vul client [command options] image_name

OPTIONS:
   --template value, -t value  output template [$VUL_TEMPLATE]
   --format value, -f value    format (table, json, template) (default: "table") [$VUL_FORMAT]
   --input value, -i value     input file path instead of image name [$VUL_INPUT]
   --severity value, -s value  severities of vulnerabilities to be displayed (comma separated) (default: "UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL") [$VUL_SEVERITY]
   --output value, -o value    output file name [$VUL_OUTPUT]
   --exit-code value           Exit code when vulnerabilities were found (default: 0) [$VUL_EXIT_CODE]
   --clear-cache, -c           clear image caches without scanning (default: false) [$VUL_CLEAR_CACHE]
   --ignore-unfixed            display only fixed vulnerabilities (default: false) [$VUL_IGNORE_UNFIXED]
   --removed-pkgs              detect vulnerabilities of removed packages (only for Alpine) (default: false) [$VUL_REMOVED_PKGS]
   --vuln-type value           comma-separated list of vulnerability types (os,library) (default: "os,library") [$VUL_VULN_TYPE]
   --ignorefile value          specify .vulignore file (default: ".vulignore") [$VUL_IGNOREFILE]
   --timeout value             timeout (default: 5m0s) [$VUL_TIMEOUT]
   --ignore-policy value       specify the Rego file to evaluate each vulnerability [$VUL_IGNORE_POLICY]
   --token value               for authentication [$VUL_TOKEN]
   --token-header value        specify a header name for token (default: "Vul-Token") [$VUL_TOKEN_HEADER]
   --remote value              server address (default: "http://localhost:4954") [$VUL_REMOTE]
   --custom-headers value      custom headers [$VUL_CUSTOM_HEADERS]
   --help, -h                  show help (default: false)
```
