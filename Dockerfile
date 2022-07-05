FROM golang:1.18-alpine3.16 as build-stage

COPY ./ /go/src/as-authentication
WORKDIR /go/src/as-authentication

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o as-authentication

FROM alpine:latest as production-stage

RUN apk --no-cache add ca-certificates

COPY --from=build-stage /go/src/as-authentication /as-authentication
WORKDIR /as-authentication

CMD ["./as-authentication"]