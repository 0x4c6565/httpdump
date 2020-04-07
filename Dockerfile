FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/httpdump

FROM scratch
COPY --from=builder /build/httpdump /bin/httpdump
ENTRYPOINT ["/bin/httpdump"]