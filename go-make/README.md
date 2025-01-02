# Go-make

This is a simple make tool that reads a makefile and executes the tasks defined in it.

[Lesson in course](https://codedeviate.github.io/aicollection/go-course-dev-tools-make.html)

## Structure

```plaintext
go-make/
├── main.go
├── cmd/
│   └── root.go
├── vars/
│   └── vars.go
├── makefile/
│   ├── parser.go
│   └── task.go
├── executor/
│   └── runner.go
└── utils/
    ├── logger.go
    └── stack.go
```

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
make run
```

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
