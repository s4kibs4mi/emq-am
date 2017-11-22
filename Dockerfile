FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /app
COPY etc /etc/config/
COPY bin/emq-am /app
WORKDIR /app
EXPOSE 8090
ENTRYPOINT ["./emq-am", "serve"]
