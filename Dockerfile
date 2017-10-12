FROM golang:latest as builder
RUN mkdir -p /go/src/github.com/olsynt/kube-event-watcher
WORKDIR /go/src/github.com/olsynt/kube-event-watcher
COPY . .
RUN "./scripts/BUILD.sh"


FROM alpine:3.6

COPY --from=builder /go/src/github.com/olsynt/kube-event-watcher/dist/bin/kube-watch /usr/bin/kube-watch
ENV PATH $PATH:/usr/bin/kube-watch
ENTRYPOINT ["kube-watch"]

CMD ["--help"]