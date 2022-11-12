FROM golang:1.19.3-alpine

WORKDIR /usr/src/app

RUN apk update && apk add git

COPY . .

# CMD ["go", "run", "main.go"]