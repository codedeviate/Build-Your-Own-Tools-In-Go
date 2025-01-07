# File compressor

This is a simple file compressor tool that compresses or decompresses a file using the gzip algorithm.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-file-compressor.html)

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
make run --<compress|decompress> source=<source file> --target=<target file>
```

Where `<compress|decompress>` is the action which can be either `compress` or `decompress`, `<source file>` is the file to compress or decompress and `<target file>` is the destination file.

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
