FROM golang:alpine as builder

LABEL maintainer "Yannick Foeillet <bzhtux@gmail.com>"

RUN mkdir -p $GOPATH/src/github.com/bzhtux/go-genpassword \
    && apk add git

COPY ./main.go $GOPATH/src/github.com/bzhtux/go-genpassword

WORKDIR $GOPATH/src/github.com/bzhtux/go-genpassword

RUN go get ./... \
    && go build -o /tmp/gen_pass . \
    && ls -ltr /tmp/

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /tmp/gen_pass /usr/bin/
# WORKDIR /app
ENTRYPOINT ["/usr/bin/gen_pass"]