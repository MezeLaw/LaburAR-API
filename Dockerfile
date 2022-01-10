FROM golang:1.16.7-alpine3.13

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o LaburAR cmd/api/api.go
EXPOSE 8080
CMD ["./LaburAR"]
