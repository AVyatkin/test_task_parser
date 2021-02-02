FROM golang:1.10

WORKDIR /go/cmd/postcode_parser
COPY ./cmd/postcode_parser .
CMD ["/go/cmd/postcode_parser/run.sh"]