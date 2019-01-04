FROM golang:1.11.4-alpine
EXPOSE 8080
RUN apk add --update git; \
    mkdir -p ${GOPATH}/rest-api; \
    go get -u github.com/gorilla/mux
WORKDIR ${GOPATH}/rest-api/
COPY rest-api.go ${GOPATH}/rest-api/
RUN go build -o rest-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/rest-api/rest-api .
CMD [ "/app/rest-api" ]