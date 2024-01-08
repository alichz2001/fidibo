FROM golang:latest as builder
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY ./ .
RUN ls
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server  github.com/alichz2001/fidibo/cmd/server

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/server /bin/server
RUN chmod +x /bin/server
ENTRYPOINT /bin/server