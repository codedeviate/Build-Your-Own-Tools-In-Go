# Executable Inspector

This is a simple executable inspector tool that displays information about an executable file.

[Lesson in course](https://codedeviate.github.io/aicollection/go-course-dev-tools-detect-executable-format.html)

## Installation

To set up the project and install the required dependencies, run the following command:

```bash
make setup
```

## Build

To build the project, run the following command:

```bash
make build
```

## Build All

To build all of the project executables, run the following command:

```bash
make build-all
```

This will build executables for the following platforms, each with a unique name:
- Windows on AMD64 (64-bit)
- Windows on 386 (32-bit)
- Windows on ARM (32-bit)
- Windows on ARM64 (64-bit)
- Linux on AMD64 (64-bit)
- Linux on 386 (32-bit)
- Linux on ARM (32-bit)
- Linux on ARM64 (64-bit)
- Darwin on AMD64 (64-bit)
- Darwin on ARM64 (64-bit)
- FreeBSD on AMD64 (64-bit)
- FreeBSD on 386 (32-bit)
- FreeBSD on ARM (32-bit)
- FreeBSD on ARM64 (64-bit)

This is mainly for testing purposes.

Please note that some platforms may not be supported by your Go compiler.

If the `make` command fails, verify that the platform is supported by your Go compiler. You can do this by running the following command:

```bash
go tool dist list
```

This command will list all supported platforms.

## Run

To run the project, use the following command:

```bash
make run
```

This will run the executable on the current platform, inspecting all generated executables.

## Clean

To clean the project, run the following command:

```bash
make clean
```

You will need to run `make setup` again before building the project.

## Complete

For testing purposes, you can run the following command:

```bash
make complete
```

This command will clean the project, set it up, build it, and then run it with some basic test parameters.
