# Compress

This is a simple decompression tool that decompresses a file using the rar, 7z or zip algorithm.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-decompression.html)

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
make run FILE=<file> DEST=<dest>
```

Where `<file>` is the file to decompress and `<dest>` is the destination directory.

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
