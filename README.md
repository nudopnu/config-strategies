# Configuration loading strategies
Here are some opinionated reference implementations for configuration management in containers. The idea is that an app running in a container should read its configuration from three-fold:

1. A default conifg file
2. An overriding config file
3. Overriding environment variables

The implementation details are a matter of taste, here are the chosen opinions:

- using `.toml` files for configuration is a good balance between simplicity and expressiveness
- using an override config file in a seperate folder allows for easier mounts
- environment variables should take precedence over all config files
- environment variables should be prefixed and in SCREAMING_SNAKE_CASE

## Golang
The Golang version uses (Viper)[https://github.com/spf13/viper] to load the config files.

```bash
# Default config
docker build -t golang_image golang && docker run --rm golang_image

# Mounted override file
docker build -t golang_image golang && docker run --rm golang_image

# Overriding environment variables

```