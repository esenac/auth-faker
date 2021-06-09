FROM golang:1.16 AS builder
ADD . /app
WORKDIR /app
RUN go build

FROM scratch
COPY --from=builder /app/keys ./keys
COPY --from=builder /app/auth-faker .
ENTRYPOINT [ "auth-faker" ]