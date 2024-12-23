---
title: "Practical AES Encryption, How to Decrypt It?"
author: "novran"
date: 2024-12-23 00:00:00 +0700
categories: [Cryptography]
tags: [Cryptography]
mermaid: true
image:
  path: /images/
  alt: Learn How to Decrypt AES
---

Motivasi saya membuat artikel ini adalah karena pada awalnya saya merasa bahwa belajar enkripsi memang sangat membosankan. Banyak artikel lebih menekankan konsep teoritis, tetapi kurang praktis dan teknis. Di sisi lain, saya menyadari bahwa enkripsi AES (Advanced Encryption Standard) tidaklah semembosankan itu. Sebagai seorang pentester, saya juga menemukan banyak aplikasi, khususnya aplikasi mobile, yang menggunakan mekanisme enkripsi AES pada data HTTP request sebelum dikirimkan ke API. Oleh karena itu, saya rasa dengan mempelajari AES ini dapat menjadi skillset tambahan yang berguna untuk melakukan pentest.

Agar tidak terlalu bertele-tele, berikut ini adalah poin-poin yang sudah saya rangkum agar lebih mudah dipahami.

- [X] [Symmetric Encryption](#symmetric-encryption): Encrypt dan Decrypt menggunakan kunci yang sama persis.
- [X] [AES Output Encoding](#aes-output-encoding): Base64 atau Hex?
- [X] [AES Key Length](#aes-key-length): Panjang kunci AES hanya 16, 24, atau 32.
- [X] [AES Operation Mode](#aes-operation-mode): ECB (tidak menggunakan IV) dan CBC (menggunakan IV).
- [X] [Initialization Vector (IV)](#initialization-vector-iv): Konsep "kunci kedua" yang digunakan oleh beberapa mode pada AES
- [X] [Cara mengenali AES atau bukan?](#cara-mengenali-aes-atau-bukan): Selalu berkelipatan 16 bytes (16, 32, 48, 64, dan seterusnya)
- [X] [Padding](#padding): Yang membuat AES selalu berkelipatan 16 bytes

## Symmetric Encryption

Symmetric Encryption adalah sebuah metode enkripsi di mana key yang digunakan untuk proses encrypt dan decrypt itu hanya menggunakan satu key (key yang sama). Jadi untuk mempelajari AES ini tidak terlalu sulit.

Berbeda dengan Asymmetric Encryption, yang menggunakan dua key yang berbeda, public key untuk encrypt dan private key decrypt (contohnya RSA).

## AES Output Encoding

Hasil enkripsi (output) atau (biasa disebut) ciphertext itu biasanya dibungkus dalam bentuk encoding, yang paling umum digunakan yaitu:
1. `Base64`
2. `Hex`

> Kenapa harus di-encode setelah selesai dienkripsi?

Karena kalau tidak di-encode maka bentuk aslinya adalah karakter-karakter yang tidak karuan bentuknya dan sangat riskan untuk pertukaran data karena beberapa protokol tidak akan mengenali karakter-karakter tersebut.

![AES without Encoding](/images/2024-12-23-practical-aes-encryption-and-how-to-decrypt-it-symmetric-encryption-random-characters-decoded.png)

## AES Key Length

Panjang kunci (Key Length) dalam AES hanya ada tiga varian:
- 16 byte (128 bit)
- 24 byte (192 bit)
- 32 byte (256 bit)

Panjang kunci menentukan tingkat keamanan AES (semakin panjang kuncinya, semakin sulit untuk dilakukan brute force).

Mungkin kalian pernah dengar tentang AES-128, AES-192, dan AES-256. Nah! Pada dasarnya itu hanya mengacu pada panjang kunci yang digunakan.

Contoh Key pada AES-128:
```python
key = b'KeepItSecret_16_'
```

Contoh Key pada AES-192:
```python
key = b'Key_Dengan_Panjang__24__'
```

Contoh Key pada AES-256:
```python
key = b'_Str0ng3st_AES_KeY_1s__32_bytes!'
```

## AES Operation Mode

Mode atau mode operasi pada AES ini yang nantinya akan berperan penting untuk menentukan bagaimana tiap blok (per-byte atau perhurufnya) akan diproses.

Pada dasarnya, terdapat banyak mode yang bisa digunakan pada AES. Seperti CTR, CFB, GCM dan lain-lain. Namun, di sini yang akan kita bahas itu hanya CBC dan ECB saja agar tidak terlalu overwhelming.

Perbedaan yang mencolok pada kedua mode tersebut adalah penggunaan IV (Initialization Vector). Yang di mana CBC itu memerlukan IV, sedangkan ECB tidak.

> IV ini bisa kita analogikan sebagai kunci kedua (second key) untuk tambahan keamanan

**ECB (Electronic Codebook)**

Setiap blok plaintext (per-byte atau perhuruf), akan langsung diproses begitu saja. Pola hasil enkripsi akan terlihat, karena blok dari plaintext (yang sama) akan menghasilkan blok ciphertext yang sama (juga).

**CBC (Cipher Block Chaining)**

Berbeda dengan ECB, pada CBC blok pertama akan di-encrypt menggunakan IV (Initialization Vector). Lalu, setiap blok plaintext-nya (yang belum dienkripsi) akan di-XOR dengan blok ciphertext sebelumnya, sehingga akan menghasilkan pola enkripsi yang tidak terlihat (karena acak).

Dengan penjelasan tersebut, maka kita dapat menyimpulkan bahwa mode ECB itu lebih lemah dari mode CBC dan mode-mode lainnya.

## Initialization Vector (IV)

Initialization Vector (IV) bisa kita analogikan sebagai kunci kedua (second key). Namun, lebih tepatnya lagi itu seperti `salt` pada konsep `hashing`. Yang di mana akan menghadirkan ketidakpastian karena akan menciptakan pola acak pada tiap-tiap blok yang terenkripsi.

Karakter pada IV memiliki panjang yang pasti, yaitu 16 byte dan tidak ada varian lain seperti key.

Sebagai contoh:
```python
init_vector = b'IV_would_be___16'
```

## Cara mengenali AES atau bukan?

Enkripsi AES akan selalu memproses data dalam kelipatan 16 byte. Jadi, meskipun data aslinya tidak berkelipatan 16 byte, hasil enkripsinya akan tetap menjadi kelipatan 16 byte. Selain itu, ukuran key-nya juga tidak memengaruhi ukuran data yang sudah diproses, di mana hasilnya akan tetap menjadi kelipatan 16 byte, apa pun bentuknya.

Sebagai contoh di sini, untuk menentukan AES atau bukan dapat menggunakan rumus `cipher_bytes % 16`, seperti script di bawah ini:
```python
import base64

def is_AES(ciphertext_bytes):
    if len(ciphertext_bytes) % 16 == 0:
    	return True
    else:
    	return False

base64_ciphertext = "Tra+6tlhtCU9PM7cKWM2QcK3ag22lxTivZmRdd9FZVCRgeM79lWVWbWAGevO3HoeTPV/rI5EabHM+TdOBvxkc/51qLQ0qUh4rlbELka4yKQ="
ciphertext_bytes = base64.b64decode(base64_ciphertext)

print('Byte Length {byte_length}, Is it AES? {is_AES}'.format(byte_length=len(ciphertext_bytes), is_AES=is_AES(ciphertext_bytes)))

# Byte Length 80, Is it AES? True
# 80 adalah kelipatan 16
```

## Padding

Jika ukuran data tidak dalam kelipatan 16 byte, maka tugas padding di sini akan menambahkan blok terakhir agar mencapai ukuran yang sesuai (kelipatan 16).

Misal, saya akan mencoba untuk mengenkripsi kata `ENCRYPT_ME` yang di mana kata tersebut hanya berukuran 10 bytes (atau 10 karakter).

| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 |
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
| E | N | C | R | Y | P | T | _ | M | E  |    |    |    |    |    |    |

Nah! Inilah fungsi padding, yang di mana akan mengisi sisa bloknya yang kosong (dari 11 sampai 16). Sehingga setelah proses enkripsinya selesai maka hasilnya akan tetap 16 bytes.

|  1   |  2   |  3   |  4   |  5   |  6   |  7   |  8   |  9   |  10  |  11  |  12  |  13  |  14  |  15  |  16  |
|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|
| 0x28 | 0x45 | 0x8b | 0xed | 0xbd | 0x21 | 0xb0 | 0x7d | 0x47 | 0x82 | 0xd0 | 0x56 | 0x36 | 0xcb | 0x5a | 0xff |
