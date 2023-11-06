# Skip Traversal of Files/Directories

## Skip Files
Vul traversals directories and looks for all lock files by default.
If your image contains lock files which are not maintained by you, you can skip the file.

```
$ vul image --skip-files "/Gemfile.lock" --skip-files "/var/lib/gems/2.5.0/gems/http_parser.rb-0.6.0/Gemfile.lock" quay.io/fluentd_elasticsearch/fluentd:v2.9.0
```

## Skip Directories
Vul traversals directories and look for all lock files by default.
If your image contains lock files which are not maintained by you, you can skip traversal in the specific directory.

```
$ vul image --skip-dirs /var/lib/gems/2.5.0/gems/fluent-plugin-detect-exceptions-0.0.13 --skip-dirs "/var/lib/gems/2.5.0/gems/http_parser.rb-0.6.0" quay.io/fluentd_elasticsearch/fluentd:v2.9.0
```
