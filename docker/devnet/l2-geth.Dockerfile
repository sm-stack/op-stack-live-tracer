FROM golang:1.23-alpine AS builder

RUN apk add --no-cache gcc musl-dev linux-headers git

WORKDIR /app

RUN git clone https://github.com/ethereum-optimism/op-geth.git && \
    cd op-geth && \
    git checkout v1.101411.4

COPY tracers/simple-tracer/simple.go /app/op-geth/eth/tracers/live/simple.go

# Build op-geth
RUN cd op-geth && go run build/ci.go install -static ./cmd/geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/op-geth/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp

ENTRYPOINT ["geth"]