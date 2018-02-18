### docker build -t dialogflow-telegram:latest .
### docker run --rm -it dialogflow-telegram:latest --help

FROM golang:latest as builder

COPY . /go/src/dialogflow-telegram/
WORKDIR /go/src/dialogflow-telegram/
RUN go get
RUN CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s"

FROM alpine:3.5
COPY --from=builder /go/src/dialogflow-telegram/dialogflow-telegram /dialogflow-telegram
WORKDIR /
ENTRYPOINT ["/dialogflow-telegram"]
CMD ["--help"]