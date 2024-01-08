---
layout: post
title: "DNS Attack: LLMNR / NBT-NS Poisoning (Network Sniffing)"
date: 2024-01-05 00:00:00 +0700
categories: "Active-Directory-Pentesting"
---

LLMNR (Link-Local Multicast Name Resolution) dan NBT-NS (NetBIOS Name Service) Poisoning merupakan teknik serangan dalam jaringan komputer yang dimanfaatkan untuk mencuri informasi dengan cara memalsukan hostname dan melakukan Man-in-The-Middle Attack.

LLMNR digunakan untuk menemukan alamat IP dari nama host dalam jaringan lokal tanpa bergantung pada server DNS, sementara NBT-NS berfungsi untuk mengonversi nama NetBIOS menjadi alamat IP.

Dalam serangan ini, Attacker akan memanipulasi respons LLMNR atau NBT-NS korban dengan berpura-pura menjadi host yang dituju, namun dengan catatan bahwa host yang dituju tidak tercatat pada DNS Server (dengan kata lain host yang tidak valid), sehingga memungkinkan mereka untuk merekayasa lalu lintas jaringan dan mencuri kredensial. 

# Proof of Concept

## 1. Setup Listener (Fake Services)

Tool ini diperlukan untuk menangkap permintaan yang masuk.

### Install Responder
```sh
git clone https://github.com/lgandx/Responder
cd Responder/
pip3 install -r requirements.txt
```

### Run Responder
```sh
sudo python3 Responder.py -I <network interface>
```

![Running Responder](https://github.com/xchopath/www.novr.one/assets/44427665/1726817f-832d-4b24-8f4c-87e172581a32)

<br/>

## 2. Attack Scenario

Attacker membiarkan User mengakses alamat (melalui hostname) yang tidak valid, maka DNS server tidak mengenali hostname yang diminta.

![LLMNR NBT-NS Poisoning (1)](https://github.com/xchopath/www.novr.one/assets/44427665/e7dc5884-501d-45dd-9471-b0113c174b82)

Setelah itu, User akan melakukan _Broadcast_ untuk mencari alamatnya (secara otomatis). Namun tanpa disadari, Attacker telah menyiapkan _Listener_ untuk _Capture_ permintaan pengguna dan berpura-pura menjadi alamat yang dituju.

![LLMNR NBT-NS Poisoning (2)](https://github.com/xchopath/www.novr.one/assets/44427665/0e685ca6-03a0-44cf-921f-465df145248a)

Kemudian, jaringan User dan Attacker akan berkomunikasi dan User akan memberikan _Credential_ kepada Attacker berupa akses _NTLM (Network Time Lockout and Login Mechanism)_.

![LLMNR NBT-NS Poisoning (3)](https://github.com/xchopath/www.novr.one/assets/44427665/f3f37d89-46d5-4647-9598-f3cf19aca284)

Berikut ini adalah contoh _Credential_ yang berhasil di-capture menggunakan _Responder_.

![image](https://github.com/xchopath/www.novr.one/assets/44427665/ed04a15a-1eae-45e2-886a-f9053368b473)
