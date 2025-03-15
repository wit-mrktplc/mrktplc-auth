# The auth service

This is a simple auth service that authenticates our users based on enabled domains.

The domains are stored in the environment variable `AUTHORIZED_DOMAINS` as a comma-separated list.

## Steps to run

There is a makefile so it is extremely simple:
1. `make build` to build the docker image
2. `make deploy` to forward the ports and start the service
