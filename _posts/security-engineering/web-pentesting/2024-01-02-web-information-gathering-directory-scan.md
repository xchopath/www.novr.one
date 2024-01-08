---
layout: post
title: "Web Directory Scanning"
date: 2024-01-02 08:00:00 +0700
categories: "Web-Pentesting"
---

### Wordlist Setup

```bash
git clone https://github.com/danielmiessler/SecLists
sudo mv SecLists /usr/share/
```

<br/>

## Directory Enumeration (Fuzzing)

### Gobuster

Install
```bash
go install github.com/OJ/gobuster/v3@latest
sudo cp ~/go/bin/gobuster /usr/local/bin/
```

Run
```bash
gobuster dir -u http://target.com -w /usr/share/SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt
```

### FFuF

Install
```bash
go install github.com/ffuf/ffuf/v2@latest
sudo cp ~/go/bin/ffuf /usr/local/bin/
```

Run
```bash
ffuf -recursion-depth 3 -t 100 -w /usr/share/SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt -u http://target.com/FUZZ -r
```

-----

<br/>

## Web Crawling

### Katana

Install
```bash
go install github.com/projectdiscovery/katana/cmd/katana@latest
sudo cp ~/go/bin/katana /usr/local/bin/
```

Run
```bash
katana -u target.com
```