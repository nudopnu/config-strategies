# Configuration loading strategies

Here are some opinionated reference implementations for configuration management in containers. The idea is that an app running in a container should read its configuration from three-fold:

1. A default conifg file
2. An overriding config file
3. Overriding environment variables

The implementation details are a matter of taste, here are the chosen options:

- using `.toml` files for configuration as balance between simplicity and expressiveness
- using an override config file with an `-override` suffix
- precedence for environment variables over all config files
- prefixed environment variables in SCREAMING_SNAKE_CASE
- allowing sensitive configs (secrets) only via environment variables or as separate files that have to be specified
- when providing a secret via a file, the corresponding environment variable is suffixed by `_FILE` (e.g. `APP_DATABASE_PASSWORD_FILE` containing the database password)

## Golang

The Golang version uses [Viper](https://github.com/spf13/viper) to load the config files.

```bash
# Build the golang image
docker build -t golang_image golang

# Run with default config
docker run --rm golang_image

# Run with a mounted override file
MSYS_NO_PATHCONV=1 docker run --rm \
    -v "$(pwd)/mounted-config.toml:/app/config-override/config.toml" \
    golang_image

# Run with overriding environment variables
docker run --rm \
    -e APP_SERVER_HOST="ENVIRONMENT" \
    -e APP_SERVER_PORT="3333" \
    -e APP_SERVER_JWT_SECRET="ENVIRONMENT" \
    -e APP_DATABASE_HOST="ENVIRONMENT" \
    -e APP_DATABASE_PORT="3333" \
    -e APP_DATABASE_DBNAME="ENVIRONMENT" \
    -e APP_DATABASE_PASSWORD="ENVIRONMENT" \
    -e APP_RUNTIME_RUNTIME_SETUP="ENVIRONMENT" \
    golang_image
```

## Docker Compose

The docker compose setup requires sensitive configuration data to be set via environment variables. You can test it by generating a `.env` file:

```bash
echo "JWT_SECRET=\"$(openssl rand 64 | openssl enc -A -base64)\"" > .env
echo "DATABASE_PASSWORD=\"$(tr -dc A-Za-z0-9 </dev/urandom | head -c 13)\"" >> .env
```

And then run

```bash
docker compose up
```

## Docker Stack

The docker stack setup requires sensitive configuration data to be set via docker secrets. 

```bash
# if not already a swarm
docker swarm init

# Build the golang image
docker build -t golang_image golang

# generate jwt secret
openssl rand 64 | openssl enc -A -base64 | docker secret create jwt_secret -

# generate database password secret
tr -dc A-Za-z0-9 </dev/urandom | head -c 13 | docker secret create database_password -

# deploy stack
docker stack deploy -c docker-stack.yml stack

# inspect the logs
docker service logs stack_golang --follow

# exit swarm mode
docker swarm leave --force
```