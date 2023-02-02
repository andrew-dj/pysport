FROM golang:1-alpine3.11 AS build-go
RUN apk --no-cache add git bzr mercurial
ENV D=/go/src/github.com/sportorg/online
RUN go get -u github.com/golang/dep/...
ADD ./ $D/
ENV GO111MODULE on
RUN cd $D \
    && go get -u \
    && go build ./cmd/sportorg


FROM node:alpine AS build-node
# This is required due to this issue: https://github.com/nodejs/node-gyp/issues/1236#issuecomment-309401410
RUN mkdir /root/.npm-global && npm config set prefix '/root/.npm-global'
ENV PATH="/root/.npm-global/bin:${PATH}"
ENV NPM_CONFIG_LOGLEVEL warn
ENV NPM_CONFIG_PREFIX=/root/.npm-global
RUN npm install -g npm@latest
RUN mkdir -p /web
ADD ./web/ /web/
RUN cd /web && npm install
RUN cd /web && npm run build --prod --progress


FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /usr/share/sportorg/
COPY --from=build-go /go/src/github.com/sportorg/online/sportorg ./sportorg
COPY --from=build-go /go/src/github.com/sportorg/online/data ./data
COPY --from=build-node /web/dist ./web/dist
EXPOSE 8080
CMD ["./sportorg"]