# Encode/decode

This is a simple encode/decode tool that encodes and decodes strings using base64 and url encoding.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-encode-decode.html)

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
make run INPUT=<data to encode>
```

Where `<data to encode>` is the data to encode/decode.

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