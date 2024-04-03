FROM golang:alpine as build

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s"

FROM scratch

COPY --from=build build/warp-plus warp-plus

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./warp-plus"]