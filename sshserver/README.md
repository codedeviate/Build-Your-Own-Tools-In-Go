# SSH Server

This is a simple SSH Server tool that listens for SSH requests on a port and displays the request information.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-ssh-server.html)

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
make run PORT=<port> KEY=<key>
```

Where `<port>` is the port to listen on and `<key>` is the path to the private key file.

## Clean

To clean the project, run the following command:

```bash
make clean
```

Run `make setup` again before building the project.
