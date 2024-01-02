# APK File Recognition

Pada awalnya saya bingung saat melakukan Test pada aplikasi Mobile Android, dikarenakan terdapat banyak sekali variasi tutorial yang dipecah berdasarkan Framework yang digunakan, seperti *How to SSL Pinning in Flutter* dan lainnya. Dari situlah muncul satu pertanyaan di otak saya.

- Bagaimana cara kita mengetahui Framework apa yang digunakan pada APK tersebut?

Setelah saya berdiskusi dengan beberapa teman saya, akhirnya mendapat pencerahan, kita dapat mengetahuinya dengan cara membaca **Common Files** pada file APK-nya setelah diekstraksi.

### APK Extraction

Sebelum mencocokkan Common Files yang di dalamnya, berikut ini adalah Command untuk ekstraksi APK menggunakan apktool.

```
apktool d Sample.apk -o ./OutputDir
```

## Software & Framework Classification

### 1. Flutter
```
lib/armeabi-v7a/libflutter.so
lib/arm64-v8a/libflutter.so
lib/x86/libflutter.so
lib/x86_64/libflutter.so
```

Bahasa pemrograman yang digunakan `DART`.

### 2. React Native
```
lib/armeabi-v7a/libreactnativejni.so
lib/arm64-v8a/libreactnativejni.so
lib/x86/libreactnativejni.so
lib/x86_64/libreactnativejni.so
assets/index.android.bundle
```

Bahasa pemrograman yang digunakan `React Native` (mirip seperti JavaScript).

### 3. Kotlin
```
kotlin/*
smali/kotlin/*
original/META-INF/kotlin-stdlib*
```

Bahasa pemrograman yang digunakan `Java`.

----------

Sebagai catatan, jika Common Files di atas tidak berhasil ditemukan, maka periksa kembali File-File yang terdapat di dalam APK-nya. Bilamana terdapat File yang memiliki ekstensi `.java` kemungkinan APK tersebut dibuat menggunakan `Java (Native)`, yang di mana saat pengetesan nanti, kurang lebih akan mirip seperti Kotlin nantinya.
