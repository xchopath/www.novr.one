# Linux Samba shared to Windows Virtual Machine

### Install Samba
```
sudo apt-get install samba -y
```

### Add User for Samba
```
sudo adduser smbshare
sudo mkdir /usr/share/Documents
chown -R smbshare. /usr/share/Documents
```

### Configuration for Samba

Edit `sudo nano /etc/samba/smb.conf`

```
[Documents]
  path = /usr/share/Documents
  browsable = yes
  writable = yes
  read only = no
  guest ok = yes
  force create mode = 0666
  force directory mode = 0777
  hosts allow = 127.0.0.1
  hosts deny = ALL
```

### Restart Samba

```
sudo systemctl restart smbd
```

### Add User to Samba

```
sudo smbpasswd -a smbshare
sudo smbpasswd -e smbshare
```

### CMD Notes:

Equivalent `rm -rf` in Windows CMD:
```
rd /s /q "C:\Users\vboxuser\Documents"
```

SymLink In Windows:
```
mklink /D C:\Users\vboxuser\Documents \\10.0.2.2\Documents\
```
