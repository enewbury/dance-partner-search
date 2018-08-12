FROM golang:1.7.3
WORKDIR /go/src/github.com/enewbury/partner-search/
RUN go get
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/enewbury/partner-search/app .
ENTRYPOINT ["./app"]