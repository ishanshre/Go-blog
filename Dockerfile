FROM golang:1.20.2-bullseye

LABEL maintainer="Ishan Shrestha"

#RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . /app/

RUN go mod download
RUN go mod vendor
RUN go mod verify

RUN go build -o main /app/cmd/Go-blog/main.go

EXPOSE 8000

CMD [ "/app/main" ]
