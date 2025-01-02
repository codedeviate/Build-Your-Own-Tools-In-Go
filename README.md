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

If the checkbox is ticked then the tools is working on atleast one platform.

If there is no status icon, then the tool has not been started yet.

> **Clarification Regarding Status**
> If the status indicates an issue on a specific platform, it does not automatically mean the problem is severe. It could range from us being unable to verify that it works to confirming that it does not work.


## Tools

### Development Tools
- [x] [go-make](./go-make) :penguin: :green_apple: :window:
- [x] [watcher](./watcher) :penguin: :green_apple:

### Basic Network Tools
- [x] [Ping](./ping) :penguin: :green_apple: :window:
- [x] [Whois](./whois) :penguin: :green_apple: :window:
- [x] [Dig/nslookup](./dnslookup) :penguin: :green_apple: :window:
- [ ] [Traceroute](./traceroute) :chicken: :apple:
- [x] [Netstat](./netstat) :penguin: :green_apple: :headstone:
- [x] [Netscan](./netscan) :penguin: :green_apple: :window:

### Advanced Network Tools
- [x] [Port Scanner](./portscanner) :penguin: :green_apple: :window:
- [x] [IP Geolocation](./ipgeolocation) :penguin: :green_apple: :window:
- [x] [Port Knocking](./portknocking) :penguin: :green_apple: :window:
- [x] [Packet Sniffer](./packetsniffer) :penguin: :green_apple: :headstone:
- [x] [Proxy Server](./proxyserver) :penguin: :green_apple: :window:
- [x] [Port Forwarding](./portforwarding) :penguin: :green_apple: :window:
- [ ] [Bandwidth Monitor](./bandwidthmonitor) :chicken: :apple: :headstone:
- [x] [Web Crawler](./webcrawler) :penguin: :green_apple: :window:
- [x] [Web Scraper](./webscraper) :penguin: :green_apple: :window:

### Network Services
- [x] [HTTP Client](./httpclient) :penguin: :green_apple: :window:
- [x] [HTTP Server](./httpserver) :penguin: :green_apple: :window:
- [x] [HTTPS Client](./httpsclient) :penguin: :green_apple: :window:
- [x] [HTTPS Server](./httpsserver) :penguin: :green_apple: :window:
- [x] [FTP Client](./ftpclient) :green_apple: :window:
- [x] [FTP Server](./ftpserver) :green_apple: :window:
- [ ] [FTPS Client](./ftpsclient) :apple: :headstone:
- [ ] [FTPS Server](./ftpsserver) :apple: :headstone:
- [x] [SMTP Client](./smtpclient) :penguin: :green_apple: :window:
- [x] [SMTP Server](./smtpserver) :penguin: :green_apple: :window:
- [x] [Telnet Client](./telnetclient) :penguin: :green_apple: :window:
- [x] [Telnet Server](./telnetserver) :penguin: :green_apple: :window:
- [ ] [SSH Client](./sshclient) :apple:
- [ ] [SSH Server](./sshserver) :apple:
- [ ] [DNS Client](./dnsclient) :penguin: :apple:+
- [ ] [DNS Server](./dnsserver) :apple: :headstone:
- [ ] [DHCP Client](./dhcpclient)
- [ ] [DHCP Server](./dhcpserver)
- [ ] [VPN Client](./vpnclient)
- [ ] [VPN Server](./vpnserver)

### Security Tools
- [ ] [Firewall](./firewall)
- [ ] [Load Balancer](./loadbalancer)
- [ ] [Reverse Proxy](./reverseproxy)

### Miscellaneous Tools
- [x] [jq](./jq) :green_apple: :window:
- [ ] [Log analyzer](./loganalyzer)
- [ ] [File Compressor](./filecompressor)
- [ ] [Task Scheduler](./taskscheduler)
- [ ] [Text Processor](./textprocessor)
- [x] [Encode/Decode](./encodedecode) :green_apple: :headstone:
- [x] [Hash](./hash) :green_apple: :window:
