FROM golang:alpine AS builder
RUN mkdir /work
WORKDIR /work
COPY . .
RUN go get -d -v
RUN go build -o /work/server

FROM alpine
COPY --from=builder /work/server/ /bin/server
EXPOSE 8000
ENTRYPOINT ["/bin/server"]
