# NTLM Challenge/Response Is Not Secure

NTLM (New Technology LAN Manager) adalah protokol otentikasi yang dikembangkan oleh Microsoft untuk digunakan dalam lingkungan jaringan Windows. NTLM Challenge/Response adalah salah satu metode autentikasi yang digunakan dalam protokol NTLM.

Dalam NTLM Challenge/Response, proses autentikasi melibatkan tiga entitas:
- Client
- Server
- Domain Controller

Langkah-langkah komunikasi Challenge/Response.
1. Client mengirim permintaan otentikasi ke server yang meminta koneksi atau sumber daya yang terlindungi.
2. Server mengirim tantangan (challenge) ke Client. Tantangan ini berupa bilangan acak yang dihasilkan oleh server.
3. Client menghasilkan respons (response) terhadap tantangan server. Respons ini dihasilkan dengan menggunakan hash dari kata sandi pengguna yang dikombinasikan dengan nilai tantangan yang diterima.
4. Client mengirim respons tersebut ke server.
5. Server mengirim respons Client ke Domain Controller.
6. Domain Controller mengambil hash kata sandi yang sesuai untuk pengguna dari basis data keamanan dan membandingkannya dengan respons yang diterima dari Client. Jika respons sesuai, autentikasi dianggap berhasil.

Metode **NTLM Challenge/Response** ini memiliki beberapa kekurangan, seperti rentan terhadap serangan "Replay Attack" di mana _penyerang mencuri respons yang valid yang di mana berisi HASH dari Password client_. Oleh karena itu, metode autentikasi yang lebih aman seperti `Kerberos` telah diperkenalkan sebagai alternatif untuk `NTLM`.
