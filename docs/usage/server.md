# Server

```bash
NAME:
   vul server - server mode

USAGE:
   vul server [command options] [arguments...]

OPTIONS:
   --skip-update          skip db update (default: false) [$VUL_SKIP_UPDATE]
   --download-db-only     download/update vulnerability database but don't run a scan (default: false) [$VUL_DOWNLOAD_DB_ONLY]
   --reset                remove all caches and database (default: false) [$VUL_RESET]
   --cache-backend value  cache backend (e.g. redis://localhost:6379) (default: "fs") [$VUL_CACHE_BACKEND]
   --token value          for authentication [$VUL_TOKEN]
   --token-header value   specify a header name for token (default: "Vul-Token") [$VUL_TOKEN_HEADER]
   --listen value         listen address (default: "localhost:4954") [$VUL_LISTEN]
   --help, -h             show help (default: false)
```
