# Bandwidth monitor

This is a simple Bandwidth monitor tool that monitors the bandwidth usage on a host.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-bandwidth-monitor.html)

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
make run INTERFACE=<interface>
```

Where `<interface>` is the network interface to monitor the bandwidth usage.

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
