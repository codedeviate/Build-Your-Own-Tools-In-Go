# Text processor

This is a simple text processor tool that reads a file and converts it to uppercase.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-text-processor.html)

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
make run --file=<filename>
```

Where `<filename>` is the text file to process.

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
