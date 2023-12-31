FROM golang:latest AS build

WORKDIR /build
COPY . .
RUN go build -o /bin/fs_db -ldflags='-s -w -extldflags "-static"' ./cmd/fs_db/main.go
RUN mkdir -p /var/lib/fs_db

FROM scratch

COPY config/docker_config.yaml /etc/fs_db/config.yaml
COPY --from=build /var/lib/fs_db /var/lib/fs_db
COPY --from=build /bin/fs_db /bin/fs_db
ENTRYPOINT ["/bin/fs_db", "--config", "/etc/fs_db/config.yaml"]