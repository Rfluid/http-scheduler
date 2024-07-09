## Introduction

<!-- Write introduction -->

This project is a simple HTTP server that schedules tasks using a Redis database and the [schedule](https://github.com/Rfluid/scheduler) package.

## Setup

To run the default setup just hit

```bash
sudo make setup
```

The setup command may not work if some of your installations are not in the default path, in this case you can run the commands manually.
Please ensure that you have `make` installed on your system to use these commands.

### Running the Redis

To run the `http-scheduler-redis` container with a Redis database, use the following Docker command:

```bash
docker run -d --name http-scheduler-redis -p 6379:6379 -e REDIS_PASSWORD=redis redis
```

This command will start a Redis container with the default port `6379` and the password set to `redis`. This is the default at [`env.local`](./env.local) file.

## Initializing the Setup

To initialize the setup, run the following command:

```bash
sudo make init-setup
```

Or start the database container and other resources manually.

## Running the Project

This project uses a Makefile for managing build tasks.

### Building the Project

To build the project, use the following command:

```bash
make build
```

This command will compile the project and create an executable file. In production, you must use the executable file to run the project.

### Running the Project

After building the project, you can run it using the following command:

```bash
make run
```

This command will start the project.
