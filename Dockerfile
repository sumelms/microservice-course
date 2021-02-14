# Build
FROM golang:1.14-alpine AS builder

RUN apk --no-cache add git curl openssh make

WORKDIR $GOPATH/src/github.com/sumelms/microservice-catalog
ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -ldflags '-extldflags "-static"' -o bin/sumelms-catalog cmd/server/main.go

# Main
FROM registry.access.redhat.com/ubi8/ubi-minimal

WORKDIR /root/
RUN mkdir -p ./cmd/sumelms

COPY --from=builder /go/src/github.com/sumelms/microservice-catalog/bin/sumelms-catalog .

EXPOSE 8080

CMD ["./sumelms-catalog"]