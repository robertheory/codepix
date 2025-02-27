FROM golang:1.23.5

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN apt-get update && \
	apt-get install build-essential protobuf-compiler librdkafka-dev -y
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/spf13/cobra-cli@latest
RUN go install github.com/ktr0731/evans@latest

CMD ["tail", "-f", "/dev/null"]
