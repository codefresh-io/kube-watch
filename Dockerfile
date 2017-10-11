FROM golang:latest
RUN mkdir /kube-watch
WORKDIR /kube-watch
COPY . .
RUN curl https://glide.sh/get | sh
RUN glide install
ENV GOPATH=/go:/kube-watch/vendor
RUN cd vendor && find . -maxdepth 1 | grep -v ./src | sed -n '1!p' | xargs -i mv {} ./src && cd ..
# RUN "./scripts/BUILD.sh"
# CMD ["./kube-watch"]