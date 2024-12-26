# Dig/nslookup

This is a simple dig/nslookup tool that queries the DNS server for information about a domain.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-dig-nslookup.html)

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
make run HOST=<target host>
```

Where `<target host>` is the hostname to query the DNS server for information.

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
