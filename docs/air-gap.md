# Air-Gapped Environment

Vul can be used in air-gapped environments.

## Download the vulnerability database
At first, you need to download the vulnerability database for use in air-gapped environments.
Go to [vul-db][vul-db] and download `vul-offline.db.tgz` in the latest release.
If you download `vul-light-offline.db.tgz`, you have to run Vul with `--light` option.

```
$ wget https://github.com/khulnasoft-lab/vul-db/releases/latest/download/vul-offline.db.tgz
```

## Transfer the DB file into the air-gapped environment
The way of transfer depends on the environment.

```
$ rsync -av -e ssh /path/to/vul-offline.db.tgz [user]@[host]:dst
```

## Put the DB file in Vul's cache directory
You have to know where to put the DB file. The following command shows the default cache directory.

```
$ ssh user@host
$ vul -h | grep cache
   --cache-dir value  cache directory (default: "/home/myuser/.cache/vul") [$VUL_CACHE_DIR]
```

Put the DB file in the cache directory + `/db`.

```
$ mkdir -p /home/myuser/.cache/vul/db
$ cd /home/myuser/.cache/vul/db
$ mv /path/to/vul-offline.db.tgz .
```

Then, decompress it.
`vul-offline.db.tgz` file includes two files, `vul.db` and `metadata.json`.

```
$ tar xvf vul-offline.db.tgz
x vul.db
x metadata.json
$ rm vul-offline.db.tgz
```

In an air-gapped environment it is your responsibility to update the Vul database on a regular basis, so that the scanner can detect recently-identified vulnerabilities. 

## Run Vul with --skip-update option
In an air-gapped environment, specify `--skip-update` so that Vul doesn't attempt to download the latest database file.

```
$ vul image --skip-update alpine:3.12
```

[vul-db]: https://github.com/khulnasoft-lab/vul-db/releases
