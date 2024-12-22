# Ping

This is a simple ping tool that sends ICMP echo requests to a host and displays the response time.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-ping.html)

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
make run TARGET=<target host>
```

Where `<target host>` is the hostname to ping.

## Clean

To clean the project, run the following command:

```bash
make clean
```

Run `make setup` again before building the project.
