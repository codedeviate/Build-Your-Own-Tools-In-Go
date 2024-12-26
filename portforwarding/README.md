# Port forwarding

This is a simple Fort forwarding tool that forwards packets to another target.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-port-forwarding.html)

## Installation

Setup the project and install the required dependencies, run the following command:

```bash
make setup
```

## Build

To build the project, run the following command:

```bash
make build
```

## Run

To run the project, run the following command:

```bash
make run LISTEN_PORT=<port> TARGET=<target>
```

Where `<port>` is the port to listen on and `<target>` is the target to forward packets to.

## Clean

To clean the project, run the following command:

```bash
make clean
```

Run `make setup` again before building the project.

## Complete

For testing purposes, you can run the following command:

```bash
make complete
```

It will run clean, setup, build and finally run the project with some basic test parameters.
