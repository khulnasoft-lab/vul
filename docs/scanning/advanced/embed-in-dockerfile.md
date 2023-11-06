# Embed in Dockerfile

Scan your image as part of the build process by embedding Vul in the
Dockerfile. This approach can be used to update Dockerfiles currently using
KhulnaSoft’s [Microscanner][microscanner].

```bash
$ cat Dockerfile
FROM alpine:3.7

RUN apk add curl \
    && curl -sfL https://raw.githubusercontent.com/khulnasoft-lab/vul/main/contrib/install.sh | sh -s -- -b /usr/local/bin \
    && vul filesystem --exit-code 1 --no-progress /

$ docker build -t vulnerable-image .
```
Alternatively you can use Vul in a multistage build. Thus avoiding the
insecure `curl | sh`. Also the image is not changed.
```bash
[...]
# Run vulnerability scan on build image
FROM build AS vulnscan
COPY --from=aquasec/vul:latest /usr/local/bin/vul /usr/local/bin/vul
RUN vul filesystem --exit-code 1 --no-progress /
[...]
```

[microscanner]: https://github.com/aquasecurity/microscanner
