# Firewalld Commands (firewall-cmd)

### List all firewalld rules

```
firewall-cmd --list-all
```

### Open port to public zone

```
firewall-cmd --zone=public --add-port=80/tcp --permanent
firewall-cmd --reload
```

### Whitelist port to specific IP using rich rule

```
firewall-cmd --permanent --zone=public --add-rich-rule='rule family="ipv4" source address="192.168.100.50" port port="5432" protocol="tcp" accept'
firewall-cmd --reload
```

### Remove port from public zone

```
firewall-cmd --zone=public --remove-port=10050/tcp
firewall-cmd --runtime-to-permanent 
firewall-cmd --reload 
```
