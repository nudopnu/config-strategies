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
    -e APP_DATABASE_HOST="ENVIRONMENT" \
    -e APP_DATABASE_PORT="3333" \
    -e APP_DATABASE_DBNAME="ENVIRONMENT" \
    -e APP_RUNTIME_RUNTIME_SETUP="ENVIRONMENT" \
    golang_image
```