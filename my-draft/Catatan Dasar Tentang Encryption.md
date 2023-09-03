# Catatan Dasar Tentang Encryption

## Encryption

Dalam enkripsi, terdapat dua jenis algoritma, yaitu Symmetric Encryption dan Asymmetric Encryption. Symmetric Encryption menggunakan Key yang sama untuk mengenkripsi dan membongkar, sedangkan Asymmetric Encryption menggunakan dua Key yang berbeda, yaitu Public Key dan Private Key. Public Key digunakan untuk mengenkripsi dan Private Key digunakan untuk membongkar.

![Symmetric-vs-Asymmetric-Encryption](https://github.com/xchopath/www.novr.one/assets/44427665/21e2c677-3fac-47c6-8a75-ae45ec99e603)

#### Contoh Algoritma Yang Sering Digunakan

Symmetric Encryption:
- AES (Advanced Encryption Standard)
- DES (Data Encryption Standard)
- Blowfish
- Twofish
- Camellia

Asymmetric Encryption:
- RSA (Rivest-Shamir-Adleman)
- Diffie-Hellman
- ElGamal
- ECDSA (Elliptic Curve Digital Signature Algorithm)
- DSA (Digital Signature Algorithm)

## Encryption Modes

Setiap mode pada enkripsi memiliki cara khusus dan memiliki aturannya masing-masing pada saat proses enkripsi. Misalnya, ada mode yang memecah data menjadi block-block kecil, ada yang mencampurkan data dengan block sebelumnya, dan ada juga yang mengubah data menjadi aliran (Streaming) itu tergantung dari mode yang dipilih.

Encryption Modes ini biasa digunakan dalam Symmetric Encryption (bukan dalam Asymmetric Encryption), karena Asymmetric Encryption menggunakan metode yang berbeda, yaitu dengan konsep Public Key dan Private Key yang sebelumnya dijelaskan.

Berikut ini adalah mode-mode yang dapat digunakan pada Symmetric Encryption.

#### Modes

1. **CBC (Cipher Block Chaining)**:
    - Menggunakan Initialization Vector (IV).
    - Blok pesan dienkripsi dengan XOR-ing dengan blok sebelumnya.
    - Membuat ketergantungan antarblok untuk keamanan tambahan.
    - Digunakan dalam enkripsi blok simetris.

2. **ECB (Electronic Codebook Mode)**:
    - Setiap blok pesan dienkripsi secara independen.
    - Pola dalam pesan asli akan terlihat dalam cipher teks.
    - Kurang aman dibandingkan dengan mode-mode lainnya.

3. **CFB (Cipher Feedback Mode)**:
    - Mengubah algoritma enkripsi menjadi aliran (stream).
    - Digunakan untuk enkripsi aliran data bit demi bit.
    - Ketergantungan antarblok tergantung pada parameter konfigurasi.

4. **OFB (Output Feedback Mode)**:
    - Seperti CFB, mengubah algoritma menjadi aliran.
    - Tidak ada ketergantungan antarblok.
    - Cocok untuk enkripsi aliran data.

5. **CTR (Counter Mode)**:
    - Menggunakan counter sebagai input ke enkripsi.
    - Menghasilkan aliran cipher yang digabungkan dengan pesan menggunakan XOR.
    - Efisien dan digunakan dalam enkripsi berkecepatan tinggi.
  
#### Initialization Vector (IV)

Initialization Vector (IV) adalah nilai acak yang diperlukan untuk memulai proses enkripsi dan digunakan dalam mode-mode enkripsi yang mengharuskan vektor inisialisasi, terutama dalam mode-mode enkripsi blok seperti CBC (Cipher Block Chaining), CFB (Cipher Feedback), dan OFB (Output Feedback).

Namun, tidak semua mode enkripsi memerlukan IV. Sebagai contoh, dalam mode ECB (Electronic Codebook), IV tidak digunakan karena setiap blok pesan dienkripsi secara independen, sehingga tidak ada ketergantungan antarblok yang memerlukan IV. Jadi, IV adalah elemen yang khusus untuk mode-mode enkripsi tertentu, terutama dalam konteks Symmetric Encryption.
