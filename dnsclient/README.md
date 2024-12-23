# DNS Client

This is a simple DNS Client tool that queries the DNS server for information about a domain.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-dns-client.html)

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
