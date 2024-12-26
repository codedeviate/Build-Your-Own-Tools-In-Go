# FTPS Server

This is a simple FTPS Server tool that listens for FTPS requests on a port and displays the request information.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-ftps-server.html)

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
make run PORT=<port>
```

Where `<port>` is the port to listen on.

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
