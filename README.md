Loggy
=====
Log files, the reader of the files with the same name spread through the hosts in the cloud.

## Usage

```bash
loggy [scope] [remote_file] [local_file]

```

### Config

The pool of SSH nodes has to be stored in the `~/.config/loggy/pool.yml` file 
in the structured way to get accessed by the scope argument.

```yml
foo:
  staging:
    - user@node1.example.com
    - user@node2.example.com
    - user@node3.example.com
  production:
    - user1@node.example.com:2222
    - user2@node.example.com:2223
```

### Commands Example

```bash
loggy foo.staging var/log/system.log system.log
```

```bash
loggy foo.staging --tail -n 1000 var/log/system.log system.log
```

```bash
loggy foo.staging --head -n 1000 /var/www/html/var/log/system.log ~/Downloads/system.log
```

```bash
loggy foo.staging -i ~/.ssh/custom_rsa var/log/system.log system.log
```

```bash 
loggy foo.staging -p ~/pool.yml var/log/system.log system.log
```

## Test

```bash
go test ./...
```