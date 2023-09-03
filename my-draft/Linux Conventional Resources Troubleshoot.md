# Linux Conventional Resources Troubleshoot

**Check port on running services**

```
netstat -tlpn
```

**Check running services**

```
ps -ef | grep service-name
```

**Check Available Disk / Storage**

```
df -h
```

**Check running file based on active PID**

```
lsof -p <PID>
```

**Check memory used by the process**

```
ps -eo pid,cmd,%mem,%cpu --sort=-%mem | head -20
```

**Check CPU used by the process**

```
ps -eo pid,cmd,%mem,%cpu --sort=-%cpu | head -20
```
