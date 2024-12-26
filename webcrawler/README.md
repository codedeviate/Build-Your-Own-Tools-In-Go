# Web crawler

This is a simple Web crawler tool that crawls a website and extracts links from the website.

[Lesson in course](https://codedeviate.github.io/aicollection/go-tools-web-crawler.html)

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
make run URL=<url>
```

Where `<url>` is the URL of the website to crawl.

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
