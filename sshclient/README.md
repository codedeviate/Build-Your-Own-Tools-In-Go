# SSH Client

This is a simple SSH client tool that connects to an SSH server and executes a command.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-ssh-client.html)

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
make run USER=<user> PASSWORD=<password> HOST=<host> PORT=<port> COMMAND=<command>
```

Where `<user>` is the username, `<password>` is the password, `<host>` is the hostname or IP address of the server, `<port>` is the port to connect to, and `<command>` is the command to execute.

## Clean

To clean the project, run the following command:

```bash
make clean
```

Run `make setup` again before building the project.
