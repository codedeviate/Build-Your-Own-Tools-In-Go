# Executable inspector

This is a simple executable inspector tool that displays information about an executable file.

[Lesson in course](https://codedeviate.github.io/aicollection/go-course-dev-tools-detect-executable-format.html)

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

## Build all

To build all of the projects executables, run the following command:

```bash
make build-all
```

This will build the executables for the following platforms:
- Windows on Amd64
- Windows on 386
- Windoww on Arm
- Windows on Arm64
- Linux on Amd64
- Linux on 386
- Linux on Arm
- Linux on Arm64
- Darwin on Amd64
- Darwin on Arm64
- FreeBSD on Amd64
- FreeBSD on 386
- FreeBSD on Arm
- FreeBSD on Arm64

## Run

To run the project, run the following command:

```bash
make run
```

This will run the executable on the current platform inspecting all generated executables.

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
