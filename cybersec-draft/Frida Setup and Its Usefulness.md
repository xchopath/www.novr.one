# Frida Setup and Its Usefulness

![Frida Tool and Frida Server](https://github.com/xchopath/www.novr.one/assets/44427665/10f77004-8a16-4bb6-a416-62f947f0f77f)

- Apa fungsi Frida?

Banyak pertanyaan yang muncul saat saya belum mengenal Frida.

Untuk saya pribadi, biasanya lebih sering menggunakan Frida untuk:
- Bypass Root Detection
- Bypass Emulator Detection
- Bypass SSL Pinning

Dengan Frida, hal tersebut dapat dilakukan dengan cara memodifikasi aplikasi Android yang sedang berjalan dan memanipulasi aktivitasnya. Hal ini dinamakan Hooking (dalam istilah Reverse Engineering). Hooking adalah teknik yang digunakan untuk memantau dan mengintersep pemanggilan Function atau Method di dalam sebuah aplikasi. Dengan Frida Tool ini kita dapat menginjeksi sebuah Script (JavaScript) ke dalam aplikasi Android yang sedang berjalan, tujuan skrip tersebut biasanya digunakan untuk memantau dan memodifikasi aktivitas di dalam aplikasi secara Real-Time.


# Frida Setup

### Install Frida Tool in Computer Device:

```
virtualenv frida
source frida/bin/activate
pip3 install frida-tools frida==16.1.4
```

### Download Frida Server
- <https://github.com/frida/frida/releases>

Sesuaikan versinya dengan Frida Tool yang digunakan (misal: 16.1.4).

### Install to Android Device via ADB

```
adb push frida-server-* /sdcard/
echo -n "chmod +x /sdcard/frida-server-*" | adb shell
```

Jalankan Frida Server:
```
adb shell
/sdcard/frida-server-*
```

# Basic Command Cheatsheet

Melihat device yang terhubung dengan Frida Client.
```
frida-ls-devices
```

Menampilkan daftar lengkap proses yang sedang berjalan di perangkat Android yang sudah terhubung dengan Frida.
```
frida-ps -Uai
```

Menginjeksi Script ke dalam APK yang berjalan.
```
frida -l <script> -U -f <identifier>
```
