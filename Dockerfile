FROM golang:latest
RUN mkdir /kube-watch
WORKDIR /kube-watch
COPY . .
RUN curl https://glide.sh/get | sh
RUN glide install
ENV GOPATH=/go:/kube-watch/vendor
RUN cd vendor && bash -c "shopt -s extglob dotglob" && mkdir src && mv !(src) src && cd ..
RUN "./scripts/BUILD.sh"
CMD ["./kube-watch"]