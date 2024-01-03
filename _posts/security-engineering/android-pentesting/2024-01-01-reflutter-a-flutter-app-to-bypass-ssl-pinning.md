---
layout: post
title: "ReFlutter a Flutter App to Bypass SSL Pinning"
date: 2024-01-01 08:00:00 +0700
categories: "Android-Pentesting"
---

<br/>

## ReFlutter

Install
```
pip3 install reflutter==0.7.8
```

Run
```
reflutter myapp.apk
```

Kemudian, masukkan IP Proxy BurpSuite yang nantinya akan digunakan untuk Intercept.

![ReFlutter](https://github.com/xchopath/www.novr.one/assets/44427665/7c6cee3d-ec94-48f6-aeb2-1e667cd9a757)

Ingat bahwa konfigurasi network BurpSuite harus disesuaikan dengan instruksi ReFLutter `<ipburp>:8083`.

<br/>

## Setup Uber APK Signer

Repository
- <https://github.com/patrickfav/uber-apk-signer/releases/latest>

Download
```
wget https://github.com/patrickfav/uber-apk-signer/releases/download/v1.3.0/uber-apk-signer-1.3.0.jar
sudo mv uber-apk-signer-1.3.0.jar /usr/share/
```

<br/>

## Inject Uber Signer to APK File

```
java -jar /usr/share/uber-apk-signer-1.3.0.jar -a release.RE.apk
```

![Inject Uber Signer to APK File](https://github.com/xchopath/www.novr.one/assets/44427665/ae8c9556-6533-48e8-b818-c22e1dd9a253)

<br/>

File APK terbaru yang akan digunakan (ReFluttered & Signed).
```
release.RE-aligned-debugSigned.apk
```
