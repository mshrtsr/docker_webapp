FROM golang:latest as builder

WORKDIR /go
ADD . /go

RUN mkdir  %USERPROFILE%\go\src\postgresqlgo
RUN cd %USERPROFILE%\go\src\postgresqlgo
RUN set GOPATH=%USERPROFILE%\go
RUN go get github.com/lib/pq

RUN ["make"]

FROM alpine:latest

WORKDIR /go
COPY --from=builder /go/webapp_linux_amd64 /go/webapp_linux_amd64
CMD ["./webapp_linux_amd64"]