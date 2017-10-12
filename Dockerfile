FROM golang:latest as builder
RUN mkdir -p /go/src/github.com/codefresh-io/kube-watch
WORKDIR /go/src/github.com/codefresh-io/kube-watch
COPY . .
RUN "./scripts/BUILD.sh"


FROM alpine:3.6

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/codefresh-io/kube-watch/dist/bin/kube-watch /usr/bin/kube-watch
ENV PATH $PATH:/usr/bin/kube-watch
ENTRYPOINT ["kube-watch"]

CMD ["--help"]