FROM golang:alpine3.13 as builder

RUN apk update && apk add git
COPY . $GOPATH/src/github.com/lacazethomas/wakapi-stats/
WORKDIR $GOPATH/src/github.com/lacazethomas/wakapi-stats/
RUN go build -o /go/bin/wakapi-stats

ENV ENVIRONMENT prod
ENV API_URL ''
ENV API_KEY ''
ENV AccessToken ""
ENV Branch "main"
ENV Message "Update images from wakapi-stats"
ENV CommitName ""
ENV CommitEmail ""
ENV UserName ""
ENV RepoName ""

FROM alpine:3.13.1
COPY --from=builder /go/bin/wakapi-stats /bin/wakapi-stats

ENTRYPOINT ["/bin/wakapi-stats"]