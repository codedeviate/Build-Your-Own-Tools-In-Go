# Port scanner

This is a simple port scanner tool that scans for open ports on a host.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-port-scanner.html)

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
make run PROTOCOL=<protocol> HOST=<target host>
```

Where `<protocol>` is the protocol to use for scanning and `<target host>` is the hostname to scan for open ports.

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
