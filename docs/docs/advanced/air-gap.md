# Air-Gapped Environment

Vul can be used in air-gapped environments. Note that an allowlist is [here][allowlist].

## Air-Gapped Environment for vulnerabilities

### Download the vulnerability database
At first, you need to download the vulnerability database for use in air-gapped environments.

=== "Vul"

    ```
    VUL_TEMP_DIR=$(mktemp -d)
    vul --cache-dir $VUL_TEMP_DIR image --download-db-only
    tar -cf ./db.tar.gz -C $VUL_TEMP_DIR/db metadata.json vul.db
    rm -rf $VUL_TEMP_DIR
    ```

=== "oras >= v0.13.0"
    Please follow [oras installation instruction][oras].

    Download `db.tar.gz`:

    ```
    $ oras pull ghcr.io/khulnasoft-lab/vul-db:2
    ```

=== "oras < v0.13.0"
    Please follow [oras installation instruction][oras].

    Download `db.tar.gz`:

    ```
    $ oras pull -a ghcr.io/khulnasoft-lab/vul-db:2
    ```

### Download the Java index database[^1]
Java users also need to download the Java index database for use in air-gapped environments.

!!! note
    You container image may contain JAR files even though you don't use Java directly.
    In that case, you also need to download the Java index database.

=== "Vul"

    ```
    VUL_TEMP_DIR=$(mktemp -d)
    vul --cache-dir $VUL_TEMP_DIR image --download-java-db-only
    tar -cf ./javadb.tar.gz -C $VUL_TEMP_DIR/java-db metadata.json vul-java.db
    rm -rf $VUL_TEMP_DIR
    ```
=== "oras >= v0.13.0"
    Please follow [oras installation instruction][oras].

    Download `javadb.tar.gz`:

    ```
    $ oras pull ghcr.io/khulnasoft-lab/vul-java-db:1
    ```

=== "oras < v0.13.0"
    Please follow [oras installation instruction][oras].

    Download `javadb.tar.gz`:

    ```
    $ oras pull -a ghcr.io/khulnasoft-lab/vul-java-db:1
    ```


### Transfer the DB files into the air-gapped environment
The way of transfer depends on the environment.

=== "Vulnerability db"
    ```
    $ rsync -av -e ssh /path/to/db.tar.gz [user]@[host]:dst
    ```

=== "Java index db[^1]"
    ```
    $ rsync -av -e ssh /path/to/javadb.tar.gz [user]@[host]:dst
    ```

### Put the DB files in Vul's cache directory
You have to know where to put the DB files. The following command shows the default cache directory.

```
$ ssh user@host
$ vul -h | grep cache
   --cache-dir value  cache directory (default: "/home/myuser/.cache/vul") [$VUL_CACHE_DIR]
```
=== "Vulnerability db"
    Put the DB file in the cache directory + `/db`.
    
    ```
    $ mkdir -p /home/myuser/.cache/vul/db
    $ cd /home/myuser/.cache/vul/db
    $ tar xvf /path/to/db.tar.gz -C /home/myuser/.cache/vul/db
    x vul.db
    x metadata.json
    $ rm /path/to/db.tar.gz
    ```

=== "Java index db[^1]"
    Put the DB file in the cache directory + `/java-db`.

    ```
    $ mkdir -p /home/myuser/.cache/vul/java-db
    $ cd /home/myuser/.cache/vul/java-db
    $ tar xvf /path/to/javadb.tar.gz -C /home/myuser/.cache/vul/java-db
    x vul-java.db
    x metadata.json
    $ rm /path/to/javadb.tar.gz
    ```



In an air-gapped environment it is your responsibility to update the Vul databases on a regular basis, so that the scanner can detect recently-identified vulnerabilities. 

### Run Vul with the specific flags.
In an air-gapped environment, you have to specify `--skip-db-update` and `--skip-java-db-update`[^1] so that Vul doesn't attempt to download the latest database files.
In addition, if you want to scan `pom.xml` dependencies, you need to specify `--offline-scan` since Vul tries to issue API requests for scanning Java applications by default.

```
$ vul image --skip-db-update --skip-java-db-update --offline-scan alpine:3.12
```

## Air-Gapped Environment for misconfigurations

No special measures are required to detect misconfigurations in an air-gapped environment.

### Run Vul with `--skip-policy-update` option
In an air-gapped environment, specify `--skip-policy-update` so that Vul doesn't attempt to download the latest misconfiguration policies.

```
$ vul conf --skip-policy-update /path/to/conf
```

[allowlist]: ../references/troubleshooting.md
[oras]: https://oras.land/cli/

[^1]: This is only required to scan `jar` files. More information about `Java index db` [here](../coverage/language/java.md)
