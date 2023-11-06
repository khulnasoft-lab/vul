# FAQ

## Scan
### analyze error: timeout: context deadline exceeded
Your scan may time out. Java takes a particularly long time to scan. Try increasing the value of the ---timeout option such as `--timeout 15m`.

## Homebrew
### Error: Your macOS keychain GitHub credentials do not have sufficient scope!

```
$ brew tap khulnasoft-lab/vul
Error: Your macOS keychain GitHub credentials do not have sufficient scope!
Scopes they need: none
Scopes they have:
Create a personal access token:
https://github.com/settings/tokens/new?scopes=gist,public_repo&description=Homebrew
echo 'export HOMEBREW_GITHUB_API_TOKEN=your_token_here' >> ~/.zshrc
```

Try:

```
$ printf "protocol=https\nhost=github.com\n" | git credential-osxkeychain erase
```

### Error: khulnasoft-lab/vul/vul 64 already installed

```
$ brew upgrade
...
Error: khulnasoft-lab/vul/vul 64 already installed
```

Try:

```
$ brew unlink vul && brew uninstall vul
($ rm -rf /usr/local/Cellar/vul/64)
$ brew install khulnasoft-lab/vul/vul
```

### Error: x509: certificate signed by unknown authority
`VUL_INSECURE` can be used to allow insecure connections to a container registry when using SSL.

```
$ VUL_INSECURE=true vul image [YOUR_IMAGE]
```

### Running in parallel takes same time as series run
When running vul on multiple images simultaneously, it will take same time as running vul in series.  
This is because of a limitation of boltdb.
> Bolt obtains a file lock on the data file so multiple processes cannot open the same database at the same time. Opening an already open Bolt database will cause it to hang until the other process closes it.
Reference : [boltdb: Opening a database][boltdb].

[boltdb]: https://github.com/boltdb/bolt#opening-a-database

## Others
### GitHub Rate limiting

Specify GITHUB_TOKEN for authentication
https://developer.github.com/v3/#rate-limiting

```
$ GITHUB_TOKEN=XXXXXXXXXX vul alpine:3.10
```

### Unknown error

Try again with `--reset` option:

```
$ vul image --reset
```
