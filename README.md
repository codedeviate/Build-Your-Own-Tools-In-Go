# Build-Your-Own-Tools-In-Go

This is the code for the tools in the course [Build You Own Tools In Go](https://codedeviate.github.io/aicollection/go-build-your-own-tools.html)

Since this is a work in prograss, the tools are not complete yet. I will update the status of each tool as they are completed.

The current status of each tool will show at the end of the tool name in the list below. The status icons are as follows:

**Current status**

| Status icon     | Description                             |
| --------------- | --------------------------------------- |
| :green_apple:   | Working on macOS                        |
| :apple:         | Problems on macOS                       |
| :penguin:       | Working on Linux                        |
| :chicken:       | Problems on Linux                       |
| :window:        | Working on Windows                      |
| :headstone:     | Problems on Windows                     |
| :link:          | Code is in sync with the course         |
| :grey_question: | Code is not yet tested on this platform |
| :skull:         | Code is not working on this platform    |

The status field have a check mark.
- :white_check_mark: Code is working on all platforms
- :heavy_check_mark: Code is working on at least one platform
- :x: Code doesn't work on any platform
- :grey_question: The code isn't confirmed working on any platform
- :heavy_minus_sign: Code isn't ready yet or hasn't not been started
- :eight_pointed_black_star: One or more platforms lacks support

> **Clarification Regarding Status**
> If the status indicates an issue on a specific platform, it does not automatically mean the problem is severe. It could range from us being unable to verify that it works to confirming that it does not work.


## Tools

### Development Tools

4 tools (3 working on Linux, 4 working on macOS, 1 working on Windows)

| Status             | Module                                         | Linux           | Mac           | Windows         |
|--------------------|------------------------------------------------|-----------------|---------------|-----------------|
| :white_check_mark: | [go-make](./go-make)                           | :penguin:       | :green_apple: | :window:        |
| :heavy_check_mark: | [watcher](./watcher)                           | :penguin:       | :green_apple: | :grey_question: |
| :heavy_check_mark: | [textdiff](./textdiff)                         | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [executable-inspector](./executable-inspector) | :penguin:       | :green_apple: | :grey_question: |

### Basic Network Tools

9 tools (7 working on Linux, 9 working on macOS, 4 working on Windows)

| Status                     | Module                      | Linux           | Mac           | Windows         |
|----------------------------|-----------------------------|-----------------|---------------|-----------------|
| :white_check_mark:         | [Ping](./ping)              | :penguin:       | :green_apple: | :window:        |
| :white_check_mark:         | [Whois](./whois)            | :penguin:       | :green_apple: | :window:        |
| :white_check_mark:         | [Dig/nslookup](./dnslookup) | :penguin:       | :green_apple: | :window:        |
| :heavy_check_mark:         | [Traceroute](./traceroute)  | :chicken:       | :green_apple: | :grey_question: |
| :eight_pointed_black_star: | [Netstat](./netstat)        | :penguin:       | :green_apple: | :skull:         |
| :white_check_mark:         | [Netscan](./netscan)        | :penguin:       | :green_apple: | :window:        |
| :heavy_check_mark:         | [Subsystem](./subsystem)    | :penguin:       | :green_apple: | :grey_question: |
| :heavy_check_mark:         | [Ifinfo](./ifinfo)          | :penguin:       | :green_apple: | :grey_question: |
| :heavy_check_mark:         | [Certinfo](./certinfo)      | :grey_question: | :green_apple: | :grey_question: |

### Advanced Network Tools

9 tools (8 working on Linux, 9 working on macOS, 7 working on Windows)

| Status             | Module                                  | Linux     | Mac           | Windows     |
|--------------------|-----------------------------------------|-----------|---------------|-------------|
| :white_check_mark: | [Port Scanner](./portscanner)           | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [IP Geolocation](./ipgeolocation)       | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Port Knocking](./portknocking)         | :penguin: | :green_apple: | :window:    |
| :heavy_check_mark: | [Packet Sniffer](./packetsniffer)       | :penguin: | :green_apple: | :headstone: |
| :white_check_mark: | [Proxy Server](./proxyserver)           | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Port Forwarding](./portforwarding)     | :penguin: | :green_apple: | :window:    |
| :heavy_check_mark: | [Bandwidth Monitor](./bandwidthmonitor) | :chicken: | :green_apple: | :headstone: |
| :white_check_mark: | [Web Crawler](./webcrawler)             | :penguin: | :green_apple: | :window:    |
| :white_check_mark: | [Web Scraper](./webscraper)             | :penguin: | :green_apple: | :window:    |

### Network Services

19 tools (9 working on Linux, 11 working on macOS, 10 working on Windows)

| Status             | Module                          | Linux           | Mac             | Windows         |
|--------------------|---------------------------------|-----------------|-----------------|-----------------|
| :white_check_mark: | [HTTP Client](./httpclient)     | :penguin:       | :green_apple:   | :window:        |
| :white_check_mark: | [HTTP Server](./httpserver)     | :penguin:       | :green_apple:   | :window:        |
| :white_check_mark: | [HTTPS Client](./httpsclient)   | :penguin:       | :green_apple:   | :window:        |
| :white_check_mark: | [HTTPS Server](./httpsserver)   | :penguin:       | :green_apple:   | :window:        |
| :heavy_check_mark: | [FTP Client](./ftpclient)       | :grey_question: | :green_apple:   | :window:        |
| :heavy_check_mark: | [FTP Server](./ftpserver)       | :grey_question: | :green_apple:   | :window:        |
| :x:                | [FTPS Client](./ftpsclient)     | :grey_question: | :apple:         | :headstone:     |
| :x:                | [FTPS Server](./ftpsserver)     | :grey_question: | :apple:         | :headstone:     |
| :white_check_mark: | [SMTP Client](./smtpclient)     | :penguin:       | :green_apple:   | :window:        |
| :white_check_mark: | [SMTP Server](./smtpserver)     | :penguin:       | :green_apple:   | :window:        |
| :white_check_mark: | [Telnet Client](./telnetclient) | :penguin:       | :green_apple:   | :window:        |
| :white_check_mark: | [Telnet Server](./telnetserver) | :penguin:       | :green_apple:   | :window:        |
| :heavy_minus_sign: | [SSH Client](./sshclient)       | :grey_question: | :apple:         | :grey_question: |
| :heavy_minus_sign: | [SSH Server](./sshserver)       | :grey_question: | :apple:         | :grey_question: |
| :white_check_mark: | [DNS Client](./dnsclient)       | :penguin:       | :green_apple:   | :grey_question: |
| :heavy_minus_sign: | [DNS Server](./dnsserver)       | :grey_question: | :apple:         | :headstone:     |
| :heavy_minus_sign: | [DHCP Client](./dhcpclient)     | :grey_question: | :apple:         | :headstone:     |
| :heavy_minus_sign: | [DHCP Server](./dhcpserver)     | :grey_question: | :apple:         | :headstone:     |
| :grey_question:    | [VPN Client](./vpnclient)       | :grey_question: | :grey_question: | :grey_question: |
| :grey_question:    | [VPN Server](./vpnserver)       | :grey_question: | :grey_question: | :grey_question: |

### Security Tools

3 tools (0 working on Linux, 2 working on macOS, 0 working on Windows)

| Status             | Module                          | Linux           | Mac             | Windows         |
|--------------------|---------------------------------|-----------------|-----------------|-----------------|
| :grey_question:    | [Firewall](./firewall)          | :grey_question: | :grey_question: | :grey_question: |
| :heavy_check_mark: | [Load Balancer](./loadbalancer) | :grey_question: | :green_apple:   | :grey_question: |
| :heavy_check_mark: | [Reverse Proxy](./reverseproxy) | :grey_question: | :green_apple:   | :grey_question: |

### Miscellaneous Tools

12 tools (0 working on Linux, 12 working on macOS, 2 working on Windows)

| Status             | Module                              | Linux           | Mac           | Windows         |
|--------------------|-------------------------------------|-----------------|---------------|-----------------|
| :heavy_check_mark: | [jq](./jq)                          | :grey_question: | :green_apple: | :window:        |
| :heavy_check_mark: | [Log analyzer](./loganalyzer)       | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [File Compressor](./filecompressor) | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [Task Scheduler](./taskscheduler)   | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [Text Processor](./textprocessor)   | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [Encode/Decode](./encodedecode)     | :grey_question: | :green_apple: | :headstone:     |
| :heavy_check_mark: | [Hash](./hash)                      | :grey_question: | :green_apple: | :window:        |
| :heavy_check_mark: | [Strings](./strings)                | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [Decompression](./decompression)    | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [Compression](./compression)        | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [Hexdump](./hexdump)                | :grey_question: | :green_apple: | :grey_question: |
| :heavy_check_mark: | [git-tool](./git-tool)              | :grey_question: | :green_apple: | :grey_question: |
