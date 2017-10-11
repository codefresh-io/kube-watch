#
# ----- Go Builder Image ------
#
FROM golang:1.8-alpine3.6 
#AS builder

# install required packages: curl and bash
RUN apk add --no-cache bash curl git || apk update && apk upgrade

# github-release - Github Release and upload artifacts
ARG GITHUB_RELEASE=v0.7.2
RUN curl -s -LS "https://github.com/aktau/github-release/releases/download/${GITHUB_RELEASE}/linux-amd64-github-release.tar.bz2" | tar -xj -C /tmp && \
    mv /tmp/bin/linux/amd64/github-release /usr/local/bin/

# install gosu
# gosu version and sha256
ARG GOSU_VERSION=1.10
ARG GOSU_SHA_256=5b3b03713a888cee84ecbf4582b21ac9fd46c3d935ff2d7ea25dd5055d302d3c
RUN curl -s -o /tmp/gosu-amd64 -LS "https://github.com/tianon/gosu/releases/download/${GOSU_VERSION}/gosu-amd64" && \
    echo "${GOSU_SHA_256}  gosu-amd64" > /tmp/gosu-amd64.sha256 && \
    cd /tmp && sha256sum -c gosu-amd64.sha256 && \
    mv /tmp/gosu-amd64 /usr/local/bin/gosu && \
    chmod +x /usr/local/bin/gosu

# set working directory
RUN mkdir -p /go/src/github.com/olsynt/kube-event-watcher
WORKDIR /go/src/github.com/olsynt/kube-event-watcher

# copy sources
COPY . .

# set entrypoint to bash
ENTRYPOINT ["/bin/bash"]
RUN curl https://glide.sh/get | sh
ENV GOPATH "/go:/go/src/github.com/olsynt/kube-event-watcher/vendor"
RUN glide install
# RUN scripts/BUILD.sh

#
# ------ Kube-Watch runtime image ------
#
# FROM alpine:3.6

# add root certificates
# RUN apk add --no-cache ca-certificates

# copy gosu
# COPY --from=builder /usr/local/bin/gosu /usr/local/bin/gosu
# add user:group, gosu nobody
# RUN addgroup kube-watch && adduser -s /bin/bash -D -G kube-watch kube-watch && gosu nobody true

# COPY --from=builder /go/src/github.com/olsynt/kube-watch/dist/bin/kube-watch /usr/kube-event-watcher

# CMD ["kube-watch", "--help"]

