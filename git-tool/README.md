# Git tool

This is a simple git tool that can perform basic git commands.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-git-tool.html)

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
make run --command=<command> [--arguments=<arguments>]
```

Where `<command>` is the git command to run and `<arguments>` are the arguments to pass to the command. Depending on command, the arguments may be optional.
Supported commands are:
- `status`
- `commit`
- `push`
- `pull`
- `clone`

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
