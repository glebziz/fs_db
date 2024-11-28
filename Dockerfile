FROM golang:1.23 AS build

WORKDIR /build
COPY go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/fs_db ./cmd/fs_db/.
RUN mkdir -p /var/lib/fs_db

FROM scratch

COPY config/docker_config.yaml /etc/fs_db/config.yaml
COPY --from=build /var/lib/fs_db /var/lib/fs_db
COPY --from=build /bin/fs_db /bin/fs_db

EXPOSE 8888

ENTRYPOINT ["/bin/fs_db", "--config", "/etc/fs_db/config.yaml"]