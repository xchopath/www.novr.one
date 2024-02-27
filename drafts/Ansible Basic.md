# Ansible Basic

#### Ansible Installation

```
apt-get install ansible
```

#### Create a ansible config

File `~/.ansible.cfg`

```
[defaults]
inventory = ~/.ansible/hosts

[inventory]
enable_plugins = yaml, ini
```

#### Create an inventory file

File `~/.ansible/hosts`

```
[hostgroup]
foo.domain.com ansible_port=22 ansible_user=root ansible_ssh_private_key_file=~/.ssh/id_rsa
192.168.12.34 ansible_port=2222 ansible_user=root ansible_ssh_private_key_file=~/.ssh/another_id_rsa
```

#### Run a remote command

```
ansible foo.domain.com -a 'echo "hello world"'
```
