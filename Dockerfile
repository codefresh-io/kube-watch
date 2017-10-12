FROM golang:latest
RUN mkdir -p /go/src/github.com/olsynt/kube-event-watcher
WORKDIR /go/src/github.com/olsynt/kube-event-watcher
COPY . .
RUN "./scripts/BUILD.sh"
ENTRYPOINT ["./kube-watch"]
CMD ["--help"]