FROM golang:1.11.1-alpine

ENV GOPATH=/go:/Orgpa

COPY . /Orgpa/src/orgpa-database-api

WORKDIR /Orgpa/src/orgpa-database-api

EXPOSE 8000

CMD [ "go", "run", "main.go" ]
