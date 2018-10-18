FROM golang:1.11.1-alpine

COPY . /go/src/github.com/frouioui/orgpa-database-api

WORKDIR /go/src/github.com/frouioui/orgpa-database-api

EXPOSE 8000

CMD [ "go", "run", "main.go" ]
