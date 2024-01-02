# Decode APK File

Untuk melakukan Decode File APK ke bentuk Source Code yang hampir mirip dengan aslinya menggunakan `apktool`.

```
apktool d InsecureAndroidApp.apk
```

# The Findings

## Android allowBackup is Enabled
AndroidManifest.xml
```
 <application
    android:allowBackup="true"
```
Bilamana ditemukan Flag `` pada file `AndroidManifest.xml` dapat dipastikan bahwa APK tersebut memiliki mekanisme Backup. Hal ini termasuk dalam kategori **OWASP - M2: Insecure Data Storage** yang dapat dimasukan ke dalam laporan Penetration Test.

## Android Debuggable is Enabled
AndroidManifest.xml
```
 <application 
    android:debuggable="true"
```
Pada Tag tersebut menentukan apakah aplikasi dapat di-debug atau tidak. Jika aplikasi dapat di-debug maka dapat memberikan banyak informasi kepada Attacker.

------------------

## Cleartext Traffic is Permitted
```
grep -Rnai 'cleartextTrafficPermitted'
```
Bilamana menemukan Flag `cleartextTrafficPermitted="true"` artinya aplikasi mengizinkan Client untuk berkomunikasi menggunakan jalur Non-HTTPS (Cleartext).
