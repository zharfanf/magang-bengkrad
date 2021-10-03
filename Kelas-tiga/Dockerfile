# +build linux,amd64,go1.16.5,!cgo
FROM golang:alpine as builder

#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \

WORKDIR /build

COPY . .

#ENV db.url = ""

RUN go mod download

#RUN go install .

RUN go build -o build .

FROM scratch

COPY --from=builder /build/main /

ENTRYPOINT ["/main"]