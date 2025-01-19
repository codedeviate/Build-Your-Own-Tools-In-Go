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

This will build the executables for the following platforms, each with a unique name:
- Windows on Amd64 (64-bits)
- Windows on 386 (32-bits)
- Windoww on Arm (32-bits)
- Windows on Arm64 (64-bits)
- Linux on Amd64 (64-bits)
- Linux on 386 (32-bits)
- Linux on Arm (32-bits)
- Linux on Arm64 (64-bits)
- Darwin on Amd64 (64-bits)
- Darwin on Arm64 (64-bits)
- FreeBSD on Amd64 (64-bits)
- FreeBSD on 386 (32-bits)
- FreeBSD on Arm (32-bits)
- FreeBSD on Arm64 (64-bits)

This is mainly for testing purposes.

Please note that some platforms may not be supported by your Go compiler.

If the make command fails, you might need to verify that the platform is supported by your Go compiler. This can be done by running the following command:

```bash
go tool dist list
```

that will list all supported platforms.

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
