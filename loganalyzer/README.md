# Log analyzer

This is a simple log analyzer tool that reads a log file and displays the number of occurrences of each unique log message.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-log-analyzer.html)

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
make run --file=<filename> --type=<type>
```

Where `<filename>` is the log file to analyze and `<type>` is the type of log file (e.g., apache_access, nginx_access, etc.).

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
