FROM golang:alpine3.13 as builder

RUN apk update && apk add git
COPY . $GOPATH/src/github.com/lacazethomas/wakapi-stats/
WORKDIR $GOPATH/src/github.com/lacazethomas/wakapi-stats/
RUN go build -o /go/bin/wakapi-stats

ENV ENVIRONMENT prod

FROM alpine:3.13.1
EXPOSE 8080
COPY --from=builder /go/bin/wakapi-stats /bin/wakapi-stats
COPY colors.json colors.json
COPY favicon.ico favicon.ico

ENTRYPOINT ["/bin/wakapi-stats"]