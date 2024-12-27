# Wathcer

This is a simple wathcer tool that monitors the file changes in a directory.

It uses a json config file to specify glob patterns to watch, event types to react to and the command to run when a event is fired.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-dev-tools-watcher.html)

## Config file

```json
{
  "*.txt": {
    "WRITE": [
      "echo Text file changed: {file}"
    ]
  },
  "logs/*.log": {
    "CHMOD": [
      "echo Log file modified: {file}"
    ]
  }
}
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
