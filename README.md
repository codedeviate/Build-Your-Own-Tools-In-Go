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
- [x] [go-make](./go-make) :green_apple:
- [ ] [watcher](./watcher)

### Basic Network Tools
- [x] [Ping](./ping) :penguin: :green_apple:
- [x] [Whois](./whois) :penguin: :green_apple:
- [x] [Dig/nslookup](./dnslookup) :penguin: :green_apple:
- [ ] [Traceroute](./traceroute) :chicken: :apple:
- [x] [Netstat](./netstat) :penguin: :green_apple:
- [x] [Netscan](./netscan) :penguin: :green_apple:

### Advanced Network Tools
- [x] [Port Scanner](./portscanner) :penguin: :green_apple:
- [x] [IP Geolocation](./ipgeolocation) :penguin: :green_apple:
- [x] [Port Knocking](./portknocking) :penguin: :green_apple:
- [x] [Packet Sniffer](./packetsniffer) :penguin: :green_apple:
- [x] [Proxy Server](./proxyserver) :penguin: :green_apple:
- [x] [Port Forwarding](./portforwarding) :penguin: :green_apple:
- [ ] [Bandwidth Monitor](./bandwidthmonitor) :chicken: :apple:
- [x] [Web Crawler](./webcrawler) :penguin: :green_apple:
- [x] [Web Scraper](./webscraper) :penguin: :green_apple:

### Network Services
- [x] [HTTP Client](./httpclient) :penguin: :green_apple:
- [x] [HTTP Server](./httpserver) :penguin: :green_apple:
- [x] [HTTPS Client](./httpsclient) :penguin: :green_apple:
- [x] [HTTPS Server](./httpsserver) :penguin: :green_apple:
- [x] [FTP Client](./ftpclient) :green_apple:
- [x] [FTP Server](./ftpserver) :green_apple:
- [ ] [FTPS Client](./ftpsclient) :apple:
- [ ] [FTPS Server](./ftpsserver) :apple:
- [x] [SMTP Client](./smtpclient) :penguin: :green_apple:
- [x] [SMTP Server](./smtpserver) :penguin: :green_apple:
- [x] [Telnet Client](./telnetclient) :penguin: :green_apple:
- [x] [Telnet Server](./telnetserver) :penguin: :green_apple:
- [ ] [SSH Client](./sshclient) :apple:
- [ ] [SSH Server](./sshserver) :apple:
- [ ] [DNS Client](./dnsclient) :penguin: :apple:
- [ ] [DNS Server](./dnsserver) :apple:
- [ ] [DHCP Client](./dhcpclient)
- [ ] [DHCP Server](./dhcpserver)
- [ ] [VPN Client](./vpnclient)
- [ ] [VPN Server](./vpnserver)

### Security Tools
- [ ] [Firewall](./firewall)
- [ ] [Load Balancer](./loadbalancer)
- [ ] [Reverse Proxy](./reverseproxy)
