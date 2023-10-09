# Custom Data

Custom policies may require additional data in order to determine an answer.

For example, an allowed list of resources that can be created. 
Instead of hardcoding this information inside your policy, Vul allows passing paths to data files with the `--data` flag.

Given the following yaml file:

```bash
$ cd examples/misconf/custom-data
$ cat data/ports.yaml                                                                                                                                                                      [~/src/github.com/khulnasoft-lab/vul/examples/misconf/custom-data]
services:
  ports:
    - "20"
    - "20/tcp"
    - "20/udp"
    - "23"
    - "23/tcp"
```

This can be imported into your policy:

```rego
import data.services

ports := services.ports
```

Then, you need to pass data paths through `--data` option.
Vul recursively searches the specified paths for JSON (`*.json`) and YAML (`*.yaml`) files.

```bash
$ vul conf --policy ./policy --data data --namespaces user ./configs
```