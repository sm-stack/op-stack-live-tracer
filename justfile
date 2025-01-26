# Define variables for convenience
set dotenv-load
COMPOSE_FILE := "docker-compose.dev.yml"
DEV_COMPOSE_FILE := "docker-compose.dev.yml"

# Task to bring up the docker-compose stack
default:
    @just up

#################### Devnet tasks ####################

# Bring up the devnet
devnet-up:
    docker compose -f {{DEV_COMPOSE_FILE}} up -d

# Shut down all services
devnet-down:
    docker compose -f {{DEV_COMPOSE_FILE}} stop

# Rebuild the services if needed
devnet-build:
    docker compose -f {{DEV_COMPOSE_FILE}} build

# Remove volumes (for cleanup purposes)
devnet-clean:
    docker compose -f {{DEV_COMPOSE_FILE}} down -v
