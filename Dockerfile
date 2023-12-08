FROM alpine:3.19
RUN apk --no-cache add ca-certificates git
COPY vul /usr/local/bin/vul
COPY contrib/*.tpl contrib/
ENTRYPOINT ["vul"]
