# jq

This is a simple jq tool that reads a JSON or XML file and presents the content in a structured form.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-jg.html)

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
make run --file <filename> --query <query> --format <format>
```

Where `<filename>` is the JSON, CSV or XML file to process, `<query>` is the query to filter the content and `<format>` is the format of the input.

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
