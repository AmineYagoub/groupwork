FROM golang:buster as builder

WORKDIR /app
ADD . .
RUN go build -o app
EXPOSE 8020
CMD ["./app"]