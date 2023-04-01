# MySQL Multi-Master Load Balancing with HAProxy

### Preparation

### Servers

- 10.10.10.1 (MySQL Server #1)
- 10.10.10.2 (MySQL Server #2)
- 10.10.10.3 (HAProxy Server)

### Preparation
- Install MySQL Server version 5.7 on `10.10.10.1` (MySQL Server #1) and `10.10.10.2` (MySQL Server #2).
- Install HAProxy on `10.10.10.3` (HAProxy Server).

# Setup Multi-Master Replication on MySQL Servers

### Configuration

- 10.10.10.1 (MySQL Server #1): `my.cnf` configuration:

```
[mysqld]

bind-address             = 0.0.0.0
server_id                = 1
log_bin                  = /var/lib/mysql/mysql-bin.log
log_bin_index            = /var/lib/mysql/mysql-bin.log.index
relay_log                = /var/lib/mysql/mysql-relay-bin
relay_log_index          = /var/lib/mysql/mysql-relay-bin.index
expire_logs_days         = 10
max_binlog_size          = 100M
log_slave_updates        = 1
auto-increment-increment = 2
auto-increment-offset    = 1
```

- 10.10.10.2 (MySQL Server #2): `my.cnf` configuration:

```
[mysqld]

bind-address             = 0.0.0.0
server_id                = 2
log_bin                  = /var/lib/mysql/mysql-bin.log
log_bin_index            = /var/lib/mysql/mysql-bin.log.index
relay_log                = /var/lib/mysql/mysql-relay-bin
relay_log_index          = /var/lib/mysql/mysql-relay-bin.index
expire_logs_days         = 10
max_binlog_size          = 100M
log_slave_updates        = 1
auto-increment-increment = 2
auto-increment-offset    = 2
```

Restart MySQL service on both servers.

```
mysql-server-1:~$ sudo systemctl restart mysqld
```

```
mysql-server-2:~$ sudo systemctl restart mysqld
```

### Grant replica user to both servers

Add MySQL Server 2 address on MySQL Server 1 & MySQL Server 1 address on MySQL Server 2. 

- 10.10.10.1 (MySQL Server #1)

```
MySQL [(none)]> GRANT REPLICATION SLAVE ON *.* TO 'replica'@'10.10.10.2' IDENTIFIED BY 'strongpassword';
```

- 10.10.10.2 (MySQL Server #2)

```
MySQL [(none)]> GRANT REPLICATION SLAVE ON *.* TO 'replica'@'10.10.10.1' IDENTIFIED BY 'strongpassword';
```

### Get binlog file name and position id

- 10.10.10.1 (MySQL Server #1)

```
MySQL [(none)]> SHOW MASTER STATUS;
+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000001 |      446 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+
```

- 10.10.10.2 (MySQL Server #2)

```
MySQL [(none)]> SHOW MASTER STATUS;
+------------------+----------+--------------+------------------+-------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+------------------+----------+--------------+------------------+-------------------+
| mysql-bin.000003 |      524 |              |                  |                   |
+------------------+----------+--------------+------------------+-------------------+
```

### Reconfigure Slave

Do not forget to adjust `master_log_file` and `master_log_pos` on both servers.

- 10.10.10.1 (MySQL Server #1)

```
MySQL [(none)]> STOP SLAVE;

MySQL [(none)]> CHANGE MASTER TO master_host='10.10.10.2', master_port=3306, master_user='replica', master_password='strongpassword', master_log_file='mysql-bin.000003', master_log_pos=524;

MySQL [(none)]> START SLAVE;
```

- 10.10.10.2 (MySQL Server #2)

```
MySQL [(none)]> STOP SLAVE;

MySQL [(none)]> CHANGE MASTER TO master_host='10.10.10.1', master_port=3306, master_user='replica', master_password='strongpassword', master_log_file='mysql-bin.000001', master_log_pos=446;

MySQL [(none)]> START SLAVE;
```

### Testing Replication

- 10.10.10.1 (MySQL Server #1)

```
mysql-server-1:~$ mysql -u root -p

MySQL [(none)]> CREATE DATABASE test_mmr;
```

- 10.10.10.2 (MySQL Server #2)

```
mysql-server-2:~$ mysql -u root -p

MySQL [(none)]> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| test_mmr           |
+--------------------+
```


# HAProxy Settings

- 10.10.10.3 (HAProxy Server)

Edit `/etc/haproxy/haproxy.cfg`:

```
listen lbmysqlserver
    bind *:3306
    balance leastconn
    mode tcp
    option tcpka
    server mysql-master-1 10.10.10.1:3306 check weight 1
    server mysql-master-2 10.10.10.2:3306 check weight 1
```

Restart HAProxy service:

```
haproxy:~$ sudo systemctl restart haproxy
```

### Testing HAProxy Load Balance

- First

```
$ mysql -u root -h 10.10.10.3 -p

MySQL [(none)]> SHOW VARIABLES LIKE 'server_id';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| server_id     | 1     |
+---------------+-------+
```

- Second

```
$ mysql -u root -h 10.10.10.3 -p

MySQL [(none)]> SHOW VARIABLES LIKE 'server_id';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| server_id     | 2     |
+---------------+-------+
```
