FROM inf-repo-registry.cn-wulanchabu.cr.aliyuncs.com/infly-dev/golang:1.23.0 AS builder

WORKDIR /k8s-project

ENV GOPROXY="https://goproxy.cn,direct"
ENV GOPRIVATE=git.infly.tech/*
ENV GONOSUMDB=git.infly.tech

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -v -o ./bin/inspire-app main.go
RUN chmod +x bin/app

FROM inf-repo-registry.cn-wulanchabu.cr.aliyuncs.com/infly-dev/inspire-inspection-base:0.0.3

COPY --from=builder /k8s-project/bin/inspire-app /bin/inspire-app

ENTRYPOINT ["/bin/inspire-app"]