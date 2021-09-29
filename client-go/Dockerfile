
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...
RUN ls -al
RUN ls -al ..
RUN ls -al /
RUN ls -al /go
RUN ls -al ../app
RUN ls -al /go/bin
RUN ls -al /go/bin/app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=client-go Version=1.0.0
