FROM golang:1.13-buster as build

WORKDIR /go/src/X-X-X
ADD . .

RUN go mod download
RUN go build -o /go/main

FROM gcr.io/distroless/base-debian10
WORKDIR /go/
COPY --from=build /go/main .
COPY *.env ./

CMD ["./main"]