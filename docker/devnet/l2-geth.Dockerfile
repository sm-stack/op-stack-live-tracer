FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Install git to clone the repository
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

# Clone the op-geth repository and checkout to Holocene release
RUN git clone https://github.com/ethereum-optimism/op-geth.git && \
    cd op-geth && \
    git checkout v1.101411.4

# Copy the tracer file to a desired location inside the container
COPY tracers/simple-tracer/simple.go /app/op-geth/eth/tracers/live/simple.go

# Copy the entrypoint script to the container
COPY scripts/devnet/l2-op-geth-entrypoint.sh /entrypoint.sh

# Set the default command to keep the container running
VOLUME ["/db"]

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]