---
layout: post
title: "Service and Protocol Login Commands (Cheatsheet)"
date: 2024-01-26 00:00:00 +0700
categories: "Infrastructure-Pentest"
---

## MySQL Login (Port 3306)

```
mysql -u <username> -h <host> -P <port> <database> --password='<password>'
```

## PostgreSQL Login (Port 5432)

```
PGPASSWORD='<password>' psql -h <host> -p <port> -U <username> <database>
```

## Redis Login (Port 6379)

```
redis-cli -h <host> -p <port> -a '<password>'
```
