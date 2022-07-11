###################
# Builder
###################
FROM golang AS builder

RUN apt update && apt install -y unzip

ENV GO111MODULE=on
ENV PROTOC_VERSION=3.15.8

# Install dependencies
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0 \
    && curl -L https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -o /tmp/protoc.zip \
    && unzip -o /tmp/protoc.zip -d /usr/local bin/protoc \
    && unzip -o /tmp/protoc.zip -d /usr/local 'include/*' \
    && rm /tmp/protoc.zip \
    && go get google.golang.org/protobuf/cmd/protoc-gen-go \
    && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

WORKDIR $GOPATH/src/github.com/sumelms/sumelms-course
ADD . .

# Build
RUN make build

###################
# Microservice
###################
FROM registry.access.redhat.com/ubi8/ubi-minimal

WORKDIR /root/
RUN mkdir -p ./cmd/sumelms

COPY --from=builder /go/src/github.com/sumelms/sumelms-course/bin/sumelms-course .

EXPOSE 8080

CMD ["./sumelms-course"]