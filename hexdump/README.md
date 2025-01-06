# Hexdump

This is a simple hexdump tool that displays the hexadecimal representation of a file.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-hexdump.html)

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
make run FILE=<filename>
```

Where `<filename>` is the file to display the hexadecimal representation of.

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
