# vi:syntax=dockerfile

FROM golang:1.19.1-alpine as go

ENV ROOT=/go/src/blog
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY . ${ROOT}
RUN ["go", "mod", "download"]
EXPOSE 8080

ENTRYPOINT ["go","run","./src/main.go"]
