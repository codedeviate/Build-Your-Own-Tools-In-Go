# Text Dff

This is a simple text diff tool that compares two text files and shows the differences between them.

[Lesson in course](https://codedeviate.github.io/aicollection/go-course-dev-tools-text-diff.html)

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
make run OLD_FILE=<file1> NEW_FILE=<file2>
```

Where `<file1>` and `<file2>` are the files to compare.

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
