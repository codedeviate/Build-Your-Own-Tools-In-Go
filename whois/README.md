# Whois

This is a simple whois tool that queries the whois database for information about a domain.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-whois.html)

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
make run TARGET=<target domain>
```

Where `<target domain>` is the domain to query the whois database for.

## Clean

To clean the project, run the following command:

```bash
make clean
```

Run `make setup` again before building the project.
