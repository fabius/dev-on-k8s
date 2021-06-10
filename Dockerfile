FROM golang:1.16 as build
COPY main.go .
RUN go build -o myservice main.go
CMD ./myservice
