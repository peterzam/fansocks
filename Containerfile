FROM docker.io/library/golang:1.18.3-alpine as builder
RUN apk --no-cache add git
RUN git clone --branch main https://codeberg.org/peterzam/fansocks.git
WORKDIR /go/fansocks
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./fansocks

FROM scratch
COPY --from=builder /go/fansocks/fansocks /
ENTRYPOINT ["/fansocks"]
