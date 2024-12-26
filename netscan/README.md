# Netscan

This is a simple netscan tool that displays the network connections on a host.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-netscan.html)

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
make run HOST=<host> START_PORT=<start port> END_PORT=<end port>
```

Where `<host>` is the hostname to scan, `<start_port>` is the starting port number, and `<end_port>` is the ending port number.

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
