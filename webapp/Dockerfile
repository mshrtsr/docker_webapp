FROM golang:latest

WORKDIR /go
ADD . /go

RUN mkdir  %USERPROFILE%\go\src\postgresqlgo
RUN cd %USERPROFILE%\go\src\postgresqlgo
RUN set GOPATH=%USERPROFILE%\go
RUN go get github.com/lib/pq

CMD ["go", "run", "wait_for_postgres.go"]
CMD ["go", "run", "webapp.go"]
