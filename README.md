<p align="center">
  <img src=".github/sumelms.svg" />
</p>

<p align="center">
  <a href="https://travis-ci.com/sumelms/microservice-catalog">
    <img alt="Travis" src="https://travis-ci.com/sumelms/microservice-catalog.svg?branch=main">
  </a>  
  <a href="https://codecov.io/gh/sumelms/microservice-catalog">
    <img src="https://codecov.io/gh/sumelms/backend/microservice-catalog/main/graph/badge.svg?token=8E92BS3SR9" />
  </a>
  <img alt="GitHub" src="https://img.shields.io/github/license/sumelms/microservice-catalog">
  <a href="https://discord.gg/Yh9q9cd">
    <img alt="Discord" src="https://img.shields.io/discord/726500188021063682">
  </a>
</p>

## About Sumé LMS

> Note: This repository contains the **catalog microservice** of the Sumé LMS. If you are looking for more information
> about the application, we strongly recommend you to [check the documentation](https://www.sumelms.com/docs).

Sumé LMS is a modern and open-source learning management system that uses modern technologies to deliver performance
and scalability to your learning environment.

- Compatible with SCORM and xAPI (TinCan)
- Flexible and modular
- Open-source and Free
- Fast and modern
- Easy to install and run
- Designed for microservices
- REST API based application
- and more.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Prepare](#prepare)
- [Building](#building)
- [Running](#running)
- [Configuring](#configuring)
- [Testing](#testing)
- [Contributing](#contributing)
- [Code of Conduct](#code-of-conduct)
- [Team](#team)
- [Support](#support)
- [License](#license)

## Prerequisites

- Go >= 1.14.6
- PostgreSQL >= 9.5

## Prepare

Clone the repository

```bash
$ git clone [git@github.com](mailto:git@github.com):sumelms/microservice-catalog.git
```

Access the project folder, and download the Go dependencies

```bash
$ go get ./...
```

It may take a while to download all the dependencies, then you are [ready to build](#building).

## Building

There are two ways that you can use to build this microservice. The first one will build it using your own machine,
while the second one will build it using a docker instance. Also, you can build the docker image to use it with
[Docker](https://www.docker.com/) and [Kubernetes](https://kubernetes.io/), but it is up to you.

Here are the following instructions for each available option:

### Local build

It should be pretty simple, once all the dependencies are download just run the following command:

```bash
$ make build
```

It will generate an executable file at the `/bin` directory inside the project folder, and If everything works, you can
now [run the microservice](#local-run).

### Docker build

At this point, I'll assume that you have installed and configure the Docker in your system, but if it is not the case,
visit the [https://docs.docker.com/get-started/](https://docs.docker.com/get-started/).

```bash
$ make docker-build
```

If everything works, you can now [run the microservice using the docker image](#docker-run).

## Running

OK! Now you build it you need to run the microservice. That should also be pretty easy.

### Local run

If you want to run the microservice locally, you may need to have a **Postgres** instance running and accessible
from your machine, and you may need to first [configure it](#configuring). Then you can run it, you just need to
execute the following command:

```bash
$ make run
```

Once it is running you can test it: http://localhost:8080/health

### Docker run

If you want to run the microservice using Docker, the easiest way to do it is using docker swarm.

First you need to initialize the docker swarm

```bash
$ docker swarm init
```

Keep in mind that it will load the `config/config.yml` file from the project. If you want to change some
configurations you can set the environment variables in your `docker-compose.yml` file, or edit the configuration file.

Once initialized you need to deploy your containers:

```bash
$ docker stack deploy -c docker-compose.yml sumelms
```

That is it, if everything works it should be now running. You can check it using the following command:

```bash
$ docker service ls
```

If the services are correctly working you should see two containers running with 1 replica each. Now, you need to get
the IP address to access the microservice. In order to do it, you can use the following command:

```bash
$ docker system info | grep "Node Address"
```

Once you have the IP address you can now access the endpoint: http://<docker-ip>:8080/health

> NOTE: You can remove/shutdown the deployment with: `$ docker stack rm sumelms`

## Configuring

You can easily configure the application editing the `config/config.yml` file or using environment variables. We do
strongly recommend that you use the configuration file instead of the environment variables. Again, it is up to you
to decide. If you want to use the variables, be sure to prefix it all with `SUMELMS_`.

The list of the environment variables and it's default values:

```bash
SUMELMS_SERVER_HTTP_PORT = 8080
SUMELMS_DATABASE_DRIVER = "postgres"
SUMELMS_DATABASE_HOST = "localhost"
SUMELMS_DATABASE_PORT = 5432
SUMELMS_DATABASE_USER = nil
SUMELMS_DATABASE_PASSWORD = nil
SUMELMS_DATABASE_DATABASE = "sumelms_catalog"
```

> We are using [configuro](https://github.com/sherifabdlnaby/configuro) to manage the configuration, so the precedence
> order to configuration is: _Environment variables > .env > Config File > Value set in Struct before loading._

## Testing

You can run all the tests with one single command:

```bash
$ make test
```

## Contributing

Thank you for considering contributing to the project. In order to ensure that the Sumé LMS community is welcome to
all make sure to read our [Contributor Guideline](https://www.sumelms.com/docs/development/contribute).

## Code of Conduct

Would you like to contribute and participate in our communities? Please read our [Code of Conduct](https://www.sumelms.com/docs/conduct).

## Team

### Core

- Ricardo Lüders (@rluders)
- Ariane Rocha (@arianerocha)

### Contributors

...

## Support

### Discussion

You can reach us or get community support in our [Discord server](https://discord.gg/Yh9q9cd). This is the best way to
find help and get in touch with the community.

### Bugs or feature requests

If you found a bug or have a feature request, the best way to do
it is [opening an issue](https://github.com/sumelms/microservice-catalog/issues).


## License

This project licensed by the Apache License 2.0. For more information check the LICENSE file.
