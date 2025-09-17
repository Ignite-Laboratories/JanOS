ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app

COPY {{.Depth}} .
RUN go mod download && go mod verify
RUN go build -v -o /run-app ./{{.Target}}


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
