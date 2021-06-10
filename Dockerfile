FROM golang:1.16 AS builder
ADD . /app
WORKDIR /app
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build
RUN chmod +x auth-faker

FROM scratch
COPY --from=builder /app/keys /keys
COPY --from=builder /app/auth-faker /bin/auth-faker
ENTRYPOINT [ "/bin/auth-faker" ]