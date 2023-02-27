FROM golang:1.20.1-alpine3.17 as build
WORKDIR /flink-deployer
COPY . .
RUN go mod download
RUN go build ./cmd/cli

FROM alpine:3.17
WORKDIR /flink-deployer
COPY --from=build /flink-deployer/cli .
VOLUME [ "/data/flink" ]
ENTRYPOINT [ "/flink-deployer/cli" ]
CMD [ "help" ]
