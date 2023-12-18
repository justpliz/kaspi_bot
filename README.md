# kaspi_bot
## Getting started

1. Change the config according to the local environment. The config should work with `docker-compose.yaml.example` out of the box. In order to use it:
    1. `cp docker-compose.yaml{.example,}`
    1. Create a folder for docker volumes: `mkdir ~/.docker-conf`
    1. Run `docker-compose up` to run containerized project dependencies.