---
layout: post
title: "Subdomain Enumeration (DNS Exfiltration)"
date: 2024-01-02 10:00:00 +0700
categories: "Web-Pentesting"
---

<br/>

#### Wordlist Setup

```bash
git clone https://github.com/danielmiessler/SecLists
sudo mv SecLists /usr/share/
```

<br/>

## Gobuster

Install
```bash
go install github.com/OJ/gobuster/v3@latest
sudo cp ~/go/bin/gobuster /usr/local/bin/
```

Run
```sh
gobuster dns -d target.com -w /usr/share/SecLists/Discovery/DNS/subdomains-top1million-110000.txt -t 100 --wildcard
```
