# Packet sniffer

This is a simple packet sniffer that captures packets from the network and displays the information in a human-readable format.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-packet-sniffer.html)

## Installation

To install the required dependencies, run the following command:

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
make run INTERFACE=<interface>
```

Where `<interface>` is the network interface to capture packets from.

**Note:** You need to run the command as root/administrator to be able to capture packets from the network.

## Clean

To clean the project, run the following command:

```bash
make clean
```
