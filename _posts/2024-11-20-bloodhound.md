---
title: Bloodhound
author: "novran"
date: 2024-11-20 00:00:00 +0700
categories: [Active Directory Pentest]
tags: [Active Directory Pentest, Red Team]
mermaid: true
image:
  path: /images/2024-11-20-bloodhound-logo.png
  alt: Bloodhound
---

BloodHound adalah alat (tool) yang dirancang untuk **memetakan jalur eksploitasi** di jaringan yang menggunakan Active Directory. Bloodhound akan _memetakan hubungan antara User, Group, dan perangkat_ di dalam jaringan yang tergabung di Active Directory. Dengan ini akan membantu untuk mengidentifikasi jalur eksploitasi yang dapat dimanfaatkan untuk meningkatkan hak akses atau aset penting yang ada di dalam jaringan tersebut.


Pada Bloodhound terdapat 2 Tools yang diimplementasi.

1. **Bloodhound** - Untuk memvisualisasikan grafis (jalur eksploitasi) di Active Directory.
2. **Bloodhound Ingestor** - Untuk mengumpulkan data yang ada (sedetail mungkin) di Active Directory.

> Untuk menjalankan Ingestor, diperlukan setidaknya **satu Domain User** (tidak masalah meskipun Low User), atau bisa juga menggunakan **satu mesin** yang telah tergabung dengan Active Directory (Join Domain).

## Bloodhound (Server)

Instalasi menggunakan Docker:

```bash
mkdir ~/bloodhound-docker
cd ~/bloodhound-docker
wget "https://raw.githubusercontent.com/SpecterOps/bloodhound/main/examples/docker-compose/docker-compose.yml"
sudo docker compose up -d --build
```

Untuk mengakses aplikasinya dapat melalui URL <http://localhost:8080> dan untuk memeriksa kredensial (saat awal instalasi) dapat menggunakan command di bawah ini.

```bash
sudo docker compose logs | grep "Initial Password Set"
```

Kemudian, Login menggunakan `admin` dan password yang diberikan.

## Bloodhound Ingestor

Terdapat beberapa Ingestor yang bisa kita gunakan, seperti `BloodHound.py` dan [SharpHound](https://github.com/BloodHoundAD/SharpHound/releases).

### BloodHound.py

Instalasi:

```bash
git clone --branch bloodhound-ce https://github.com/dirkjanm/BloodHound.py.git
```

```bash
cd BloodHound.py
python3 -m pip install .
```

Eksekusi:

```bash
python3 bloodhound.py -d <DOMAIN> -c All -u '<DOMAIN_USER>' -p '<PASSWORD>' -ns <DC_IP> --zip
```

> `<DC_IP>` - IP Domain Controller

Setelah berhasil menjalankan Ingestor-nya, hasil yang akan didapatkan yaitu berupa file `*.zip` yang nantinya akan kita Upload ke Bloodhound untuk dipetakan.

## Visualization

Pertama-tama pergi ke menu di pojok kanan atas atau `Gear Logo` > `Administration` > `File Ingest` > `Upload File(s)`.

![Upload](/images/2024-11-20-bloodhound-upload-file.png)

Kemudian upload file `*.zip` yang dihasilkan Ingestor dan tunggu hingga Bloodhound selesai memprosesnya.

**Play with Cypher Query**

Hal yang tidak kalah penting saat mengoperasikan Bloodhound yaitu bermain dengan Cypher Query, yang berfungsi untuk memudahkan kita untuk menemukan target mana saja yang perlu kita tuju (High-Value Target) dan menghindari target yang tidak terlalu penting.

![Play with Cypher Queries](/images/2024-11-20-bloodhound-play-with-cypher.png)
