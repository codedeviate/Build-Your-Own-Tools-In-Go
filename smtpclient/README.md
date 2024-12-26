# SMTP CLient

This is a simple SMTP Client tool that sends an email to a recipient.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-smtp-client.html)

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
make run SERVER=<server> PORT=<port> USER=<user> PASSWORD=<password> FROM=<from> TO=<to> SUBJECT=<subject> BODY=<body>
```

Where `<server>` is the server to listen on, `<port>` is the port to listen on, `<user>` is the user to authenticate with, `<password>` is the password to authenticate with, `<from>` is the email address to send from, `<to>` is the email address to send to, `<subject>` is the email subject, and `<body>` is the email body.

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
