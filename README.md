# tmhc-norvegicus
Super bad-ass TMHC project.

This project can be found on https://github.com/hook-s3c/tmhc-norvegicus 

## Overview

Pentest tool and framework for rapid, accessible development. A project for our community.

## Quickstart

### Standardized Docker environment; 

For dependencies you will need to install Docker and Docker-Compose.
Instructions can be found on the wiki.

```bash
user@host$ docker-compose up                         # start the container
user@host$ docker exec -t -i $(docker ps -lq) bash -l  # attach to the container to get a bash prompt
root@contaainer$ cd /root/tmhc-norvegicus && make build
```

You can build, test and edit the project within the container.

### Baremetal;

Optionally you should be able to install all the components directly on your machine.

Dependencies;

- Go (`apt-get install golang` - Ubuntu)
- Make (`apt-get install make` - Ubuntu)


## Disclaimer

Don't use for evil, always seek permission of the owner of the platform you wish to test.
Don't do crime, the government hates competition.
