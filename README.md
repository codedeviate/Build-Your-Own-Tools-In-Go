# Build-Your-Own-Tools-In-Go

This is the code for the tools in the course [Build You Own Tools In Go](https://codedeviate.github.io/aicollection/go-build-your-own-tools.html)

Since this is a work in prograss, the tools are not complete yet. I will update the status of each tool as they are completed.

The current status of each tool will show at the end of the tool name in the list below. The status icons are as follows:

**Current status**

| Status icon   | Description                     |
| ------------- | ------------------------------- |
| :green_apple: | Working on macOS                |
| :apple:       | Problems on macOS               |
| :penguin:     | Working on Linux                |
| :chicken:     | Problems on Linux               |
| :window:      | Working on Windows              |
| :headstone:   | Problems on Windows             |
| :link:        | Code is in sync with the course |

The status field have a check mark.
- :white_check_mark: Code is working on all platforms
- :grey_question: Code is working on at least one platform
- :x: Code doesn't work on any platform
- :heavy_minus_sign: Code isn't ready yet or hasn't not been started

> **Clarification Regarding Status**
> If the status indicates an issue on a specific platform, it does not automatically mean the problem is severe. It could range from us being unable to verify that it works to confirming that it does not work.


## Tools

### Development Tools

| Status             | Module               | Linux     | Mac           | Windows  |
|--------------------|----------------------|-----------|---------------|----------|
| :white_check_mark: | [go-make](./go-make) | :penguin: | :green_apple: | :window: |
| :grey_question:    | [watcher](./watcher) | :penguin: | :green_apple: |          |

### Basic Network Tools

| Status             | Module                      | Linux     | Mac           | Windows     |
|--------------------|-----------------------------|-----------|---------------|-------------|
| :white_check_mark: | [Ping](./ping)              | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Whois](./whois)            | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Dig/nslookup](./dnslookup) | :penguin: | :green_apple: | :window:    |
| :grey_question:    | [Traceroute](./traceroute)  | :chicken: | :apple:       |             |
| :grey_question:    | [Netstat](./netstat)        | :penguin: | :green_apple: | :headstone: |
| :white_check_mark: | [Netscan](./netscan)        | :penguin: | :green_apple: | :window:    |

### Advanced Network Tools

| Status             | Module                                  | Linux     | Mac           | Windows     |
|--------------------|-----------------------------------------|-----------|---------------|-------------|
| :white_check_mark: | [Port Scanner](./portscanner)           | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [IP Geolocation](./ipgeolocation)       | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Port Knocking](./portknocking)         | :penguin: | :green_apple: | :window:    |
| :grey_question:    | [Packet Sniffer](./packetsniffer)       | :penguin: | :green_apple: | :headstone: |
| :white_check_mark: | [Proxy Server](./proxyserver)           | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Port Forwarding](./portforwarding)     | :penguin: | :green_apple: | :window:    |
| :heavy_minus_sign: | [Bandwidth Monitor](./bandwidthmonitor) | :chicken: | :apple:       | :headstone: |
| :white_check_mark: | [Web Crawler](./webcrawler)             | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Web Scraper](./webscraper)             | :penguin: | :green_apple: | :window:    |

### Network Services

| Status             | Module                          | Linux     | Mac           | Windows     |
|--------------------|---------------------------------|-----------|---------------|-------------|
| :white_check_mark: | [HTTP Client](./httpclient)     | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [HTTP Server](./httpserver)     | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [HTTPS Client](./httpsclient)   | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [HTTPS Server](./httpsserver)   | :penguin: | :green_apple: | :window:    |
| :grey_question:    | [FTP Server](./ftpserver)       |           | :green_apple: | :window:    |
| :grey_question:    | [FTPS Client](./ftpsclient)     |           | :apple:       | :headstone: |
| :grey_question:    | [FTPS Server](./ftpsserver)     |           | :apple:       | :headstone: |
| :white_check_mark: | [SMTP Client](./smtpclient)     | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [SMTP Server](./smtpserver)     | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Telnet Client](./telnetclient) | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Telnet Server](./telnetserver) | :penguin: | :green_apple: | :window:    |
| :heavy_minus_sign: | [SSH Client](./sshclient)       |           | :apple:       |             |
| :heavy_minus_sign: | [SSH Server](./sshserver)       |           | :apple:       |             |
| :white_check_mark: | [DNS Client](./dnsclient)       | :penguin: | :apple:       |             |
| :heavy_minus_sign: | [DNS Server](./dnsserver)       |           | :apple:       | :headstone: |
| :heavy_minus_sign: | [DHCP Client](./dhcpclient)     |           |               |             |
| :heavy_minus_sign: | [DHCP Server](./dhcpserver)     |           |               |             |
| :heavy_minus_sign: | [VPN Client](./vpnclient)       |           |               |             |
| :heavy_minus_sign: | [VPN Server](./vpnserver)       |           |               |             |

### Security Tools

| Status             | Module                          | Linux  | Mac  | Windows  |
|--------------------|---------------------------------|--------|------|----------|
| :heavy_minus_sign: | [Firewall](./firewall)          |        |      |          |
| :heavy_minus_sign: | [Load Balancer](./loadbalancer) |        |      |          |
| :heavy_minus_sign: | [Reverse Proxy](./reverseproxy) |        |      |          |

### Miscellaneous Tools

| Status             | Module                              | Linux  | Mac           | Windows     |
|--------------------|-------------------------------------|--------|---------------|-------------|
| :grey_question:    | [jq](./jq)                          |        | :green_apple: | :window:    |
| :heavy_minus_sign: | [Log analyzer](./loganalyzer)       |        |               |             |
| :heavy_minus_sign: | [File Compressor](./filecompressor) |        |               |             |
| :heavy_minus_sign: | [Task Scheduler](./taskscheduler)   |        |               |             |
| :heavy_minus_sign: | [Text Processor](./textprocessor)   |        |               |             |
| :grey_question:    | [Encode/Decode](./encodedecode)     |        | :green_apple: | :headstone: |
| :grey_question:    | [Hash](./hash)                      |        | :green_apple: | :window:    |
