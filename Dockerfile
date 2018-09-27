FROM golang:1.11.0-alpine3.8 as builder
WORKDIR /go-service-boilerplate
COPY . /go-service-boilerplate
RUN CGO_ENABLED=0 go build --mod=vendor -ldflags "-s -w" -o go_service

# final stage
FROM alpine:3.8
WORKDIR /
COPY --from=builder /go-service-boilerplate/go_service /go_service
ENV SERVICE_ADDR=":9000"
EXPOSE 9000
CMD ["/go_service"]
