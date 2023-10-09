FROM alpine:3.18.4
RUN apk --no-cache add ca-certificates git
COPY vul /usr/local/bin/vul
COPY contrib/*.tpl contrib/
ENTRYPOINT ["vul"]
