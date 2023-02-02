FROM alpine
RUN apk --no-cache add ca-certificates git bzr mercurial make musl-dev go nodejs nodejs-npm python3

ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN go get -u github.com/Masterminds/glide/...

ENV D=/go/src/github.com/sportorg/online
RUN go get -u github.com/golang/dep/...
ADD ./ $D/

WORKDIR $D/

ENV GO111MODULE on
RUN go get

RUN mkdir /root/.npm-global && npm config set prefix '/root/.npm-global'
ENV PATH="/root/.npm-global/bin:${PATH}"
ENV NPM_CONFIG_LOGLEVEL warn
ENV NPM_CONFIG_PREFIX=/root/.npm-global
RUN npm install -g npm@latest
RUN cd ./web && npm install

EXPOSE 8080
CMD ["python3", "./tools/server.py"]