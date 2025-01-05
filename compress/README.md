# Compress

This is a simple compress tool that compresses or decompress a file using the gzip, deflate, lzw, zlib or snappy algorithm.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-compress.html)

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
make run ACTION=<action> TYPE=<type> FILE=<file>
```

Where `<action>` is either `compress` or `decompress`, `<type>` is the compression type and `<file>` is the file to compress or decompress.

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
