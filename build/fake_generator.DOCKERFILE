FROM golang:latest as builder
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY ./ .
RUN ls
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fake_generator  github.com/alichz2001/fidibo/cmd/fake_data_generator

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/fake_generator /bin/fake_generator
RUN chmod +x /bin/fake_generator
