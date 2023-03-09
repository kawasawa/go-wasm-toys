FROM golang:1.20 AS remote-container

# golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN golangci-lint --version

# go-tools
RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/shurcooL/goexec@latest

WORKDIR /workspace
