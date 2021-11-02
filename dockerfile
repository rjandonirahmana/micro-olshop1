FROM golang:1.17-alpine

RUN apk update

WORKDIR /app

COPY . .

RUN go mod tidy
EXPOSE 6060

RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]
