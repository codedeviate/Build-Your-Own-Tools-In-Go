# Port knocking

This is a simple Port knocking tool that sends a sequence of packets to a host to open a port.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-port-knocking.html)

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
make run HOSTNAME=<hostname> DELAY_MS=<delay> PORT_1=<port 1> PORT_2=<port 2>
```

Where `<hostname>` is the hostname to send the packets to, `<delay>` is the delay between sending packets, `<port 1>`, and `<port 2>` are the ports to open.

## Clean

To clean the project, run the following command:

```bash
make clean
```

Run `make setup` again before building the project.
